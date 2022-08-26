package worker

import (
	"fmt"
	"math/big"
	"time"

	"github.com/Conflux-Chain/go-conflux-util/viper"
	"github.com/Conflux-Chain/web3pay-service/blockchain"
	"github.com/Conflux-Chain/web3pay-service/contract"
	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/service"
	"github.com/Conflux-Chain/web3pay-service/store/sqlite"
	"github.com/Conflux-Chain/web3pay-service/util"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	// channel size for billing settlement
	billingQueueSize = 5000
)

type settlementConfig struct {
	PollingInterval         time.Duration `default:"1m"`
	PollingBatchSize        int           `default:"20"`
	MaxSettlementRetries    uint          `default:"5"`
	MinConfirmedBlocks      uint64        `default:"30"`
	MaxPendingAwaitDuration time.Duration `default:"30s"`
}

type BillTask struct {
	*model.Bill
	Statements  map[uint32]int64
	TryTimes    uint
	CreatedAt   time.Time
	SubmittedAt time.Time
}

func NewBillTask(bill *model.Bill, statements map[uint32]int64) *BillTask {
	return &BillTask{
		Bill:       bill,
		Statements: statements,
		CreatedAt:  time.Now(),
	}
}

func (t *BillTask) isExpired(lifetime time.Duration) bool {
	return time.Now().After(t.SubmittedAt.Add(lifetime))
}

type BlockchainWorker struct {
	settlementConfig
	provider            *blockchain.Provider
	sqliteStore         *sqlite.SqliteStore
	billSvc             *service.BillingService
	chainSvc            *service.BlockchainService
	billSettlementQueue chan []*BillTask
	billConfirmQueue    chan []*BillTask
}

func MustNewBlockchainWorkerFromViper(
	provider *blockchain.Provider, sqliteStore *sqlite.SqliteStore,
	billSvc *service.BillingService, chainSvc *service.BlockchainService,
) *BlockchainWorker {
	var conf settlementConfig
	viper.MustUnmarshalKey("worker", &conf)

	// TODO: check bill submitting/submitted tasks status for action before run?
	return &BlockchainWorker{
		settlementConfig: conf,
		sqliteStore:      sqliteStore, provider: provider,
		billSvc: billSvc, chainSvc: chainSvc,
		billSettlementQueue: make(chan []*BillTask, billingQueueSize),
		billConfirmQueue:    make(chan []*BillTask, billingQueueSize),
	}
}

func (worker *BlockchainWorker) Run() {
	go worker.poll()
	go worker.confirm()

	worker.settle()
}

func (worker *BlockchainWorker) confirm() {
	// TODO: task confirmations are order independent, multiple consumers can be utilized
	// to increase workload capacity.
	for billTasks := range worker.billConfirmQueue {
		worker.confirmBillTasks(billTasks)
	}
}

func (worker *BlockchainWorker) confirmBillTasks(billTasks []*BillTask) {
	var successTasks, reconfirmTasks, retryTasks []*BillTask

	for i := range billTasks {
		logger := logrus.WithFields(logrus.Fields{
			"taskBillID": billTasks[i].ID, "txnHash": billTasks[i].TxnHash,
		})

		txnHash := common.HexToHash(billTasks[i].TxnHash)
		txn, err := worker.provider.TransactionByHash(txnHash)
		if err != nil {
			logger.WithError(err).
				Info("Blockchain worker failed to get txn by hash for task confirmation")
			reconfirmTasks = append(reconfirmTasks, billTasks[i])

			continue
		}

		// transaction dropped?
		if txn == nil {
			logger.Info("Blockchain worker got txn dropped for task confirmation")
			billTasks[i].Memo = "transaction dropped"
			retryTasks = append(retryTasks, billTasks[i])

			continue
		}

		// transaction pending (not mined or executed)?
		if txn.BlockHash == nil {
			if billTasks[i].isExpired(worker.MaxPendingAwaitDuration) { // pending too long?
				logger.WithField("submittedAt", billTasks[i].SubmittedAt).
					Info("Blockchain worker got txn pending too long for task confirmation")
				billTasks[i].Memo = "transaction pending too long"
				retryTasks = append(retryTasks, billTasks[i])
			} else {
				logger.Debug("Blockchain worker got txn pending for task confirmation")
				reconfirmTasks = append(reconfirmTasks, billTasks[i])
			}

			continue
		}

		// transaction failed?
		if txn.Status == nil || *txn.Status != ethtypes.ReceiptStatusSuccessful {
			logger.WithField("txnStatus", txn.Status).
				Info("Blockchain worker got non-successful txn for task confirmation")
			billTasks[i].Memo = fmt.Sprintf("transaction execution failed with status %v", txn.Status)
			retryTasks = append(retryTasks, billTasks[i])

			continue
		}

		latestBlock, err := worker.provider.BlockNumber()
		if err != nil {
			logger.WithError(err).
				Info("Blockchain worker failed to get latest block number for task confirmation")
			reconfirmTasks = append(reconfirmTasks, billTasks[i])

			continue
		}

		// enough blocks confirmed?
		targetBlockNumber := (txn.BlockNumber.Uint64() + worker.MinConfirmedBlocks)
		if targetBlockNumber > latestBlock.Uint64() {
			leftConfirmBlocks := targetBlockNumber - latestBlock.Uint64()

			logrus.WithFields(logrus.Fields{
				"txnBlockNumber":    txn.BlockNumber.Uint64(),
				"latestBlockNumber": latestBlock.Uint64(),
				"taskBillId":        billTasks[i].ID,
				"leftConfirmBlocks": leftConfirmBlocks,
			}).Debug("Blockchain worker got not enough blocks for task confirmation")
			reconfirmTasks = append(reconfirmTasks, billTasks[i])
			continue
		}

		logger.Debug("Blockchain worker finished task confirmation")
		successTasks = append(successTasks, billTasks[i])
	}

	if len(successTasks) > 0 {
		worker.finishTasks(successTasks)
	}

	if len(retryTasks) > 0 {
		worker.billSettlementQueue <- retryTasks
	}

	if len(reconfirmTasks) > 0 {
		// wait for a while to reconfirm tasks
		time.AfterFunc(2*time.Second, func() {
			worker.billConfirmQueue <- reconfirmTasks
		})
	}
}

func (worker *BlockchainWorker) finishTasks(tasks []*BillTask) {
	delBillIds := []uint64{}
	for i := range tasks {
		delBillIds = append(delBillIds, tasks[i].ID)

		exempted, err := worker.chainSvc.WriteOffAccountFee(
			common.HexToAddress(tasks[i].Coin),
			common.HexToAddress(tasks[i].Address),
			tasks[i].Fee.BigInt(),
		)

		if err != nil {
			logrus.WithField("taskBillId", tasks[i].ID).
				WithError(err).
				Error("Blockchain worker failed to write off bill task fee")

			continue
		}

		logrus.WithFields(logrus.Fields{
			"taskBillId": tasks[i].ID,
			"writtenOff": exempted,
		}).Debug("Blockchain worker written off bill fee")
	}

	util.KLock(service.KeyBillingLocker)
	defer util.KUnlock(service.KeyBillingLocker)

	if err := worker.sqliteStore.Delete(&model.Bill{}, delBillIds).Error; err != nil {
		logrus.WithField("delBillIds", delBillIds).
			WithError(err).
			Error("Blockchain worker failed to delete task bills")
	}
}

func (worker *BlockchainWorker) settle() {
	for billTasks := range worker.billSettlementQueue {
		var todoTasks, deadTasks []*BillTask

		// filter dead tasks and todo tasks
		for i := range billTasks {
			if billTasks[i].TryTimes > worker.MaxSettlementRetries {
				logrus.WithFields(logrus.Fields{
					"bill":       *billTasks[i].Bill,
					"statements": billTasks[i].Statements,
					"tryTimes":   billTasks[i].TryTimes,
				}).Warn("Blockchain worker dumped (dead) bill tasks after too many retries")
				deadTasks = append(deadTasks, billTasks[i])
				continue
			}

			todoTasks = append(todoTasks, billTasks[i])
		}

		if len(deadTasks) > 0 {
			// update bill failed status
			worker.updateBillTaskStatus(deadTasks, model.BillStatusFailed)
		}

		if len(todoTasks) == 0 {
			return
		}

		successTasks, failureTasks := worker.settleBillTasks(todoTasks)

		if len(successTasks) > 0 {
			// update bill submitted status
			worker.updateBillTaskStatus(successTasks, model.BillStatusSubmitted)
			// put into confirm queue
			worker.billConfirmQueue <- successTasks
		}

		if len(failureTasks) > 0 {
			// put into retry queue
			worker.billSettlementQueue <- failureTasks
		}
	}
}

func (worker *BlockchainWorker) settleBillTasks(billTasks []*BillTask) (successTasks, failureTasks []*BillTask) {
	// split tasks && requests group by APP coin for batch operation
	coinTasks := make(map[string][]*BillTask)
	coinRequests := make(map[string][]contract.APPCoinChargeRequest)

	for _, task := range billTasks {
		var details []contract.APPCoinResourceUseDetail
		for resourceIndex, callTimes := range task.Statements {
			details = append(details, contract.APPCoinResourceUseDetail{
				Id:    resourceIndex,
				Times: big.NewInt(callTimes),
			})
		}

		coinRequests[task.Coin] = append(coinRequests[task.Coin], contract.APPCoinChargeRequest{
			Account:   common.HexToAddress(task.Address),
			Amount:    task.Fee.BigInt(),
			UseDetail: details,
		})

		coinTasks[task.Coin] = append(coinTasks[task.Coin], task)
	}

	// call batch charge contract method for settlement
	for coin, req := range coinRequests {
		tasks := coinTasks[coin]
		taskInfos := make([]interface{}, 0, len(tasks))

		for i := range tasks {
			tasks[i].TryTimes++ // update try times

			taskInfos = append(taskInfos, struct {
				TaskBillId uint64
				Statements map[uint32]int64
				TryTimes   uint
			}{tasks[i].ID, tasks[i].Statements, tasks[i].TryTimes})
		}

		logger := logrus.WithFields(logrus.Fields{
			"APPCoin":              coin,
			"batchChargeBillTasks": taskInfos,
		})

		txn, err := worker.provider.BatchChargeAppCoinBills(
			&bind.TransactOpts{}, common.HexToAddress(coin), req,
		)

		// TODO: distinguish between different error types
		// https://developer.confluxnetwork.org/sending-tx/en/transaction_send_common_error
		// eg.,
		// 1. transaction pool full error
		// 2. duplicated txn error
		// 3. network error
		if err != nil {
			logger.WithError(err).Info("Blockchain worker failed to call batch bill charge request")
			failureTasks = append(failureTasks, tasks...)
			for i := range tasks { // update task memo
				tasks[i].Memo = err.Error()
			}

			continue
		}

		txnHash := txn.Hash().String()
		logger.WithField("txnHash", txnHash).Debug("Blockchain worker batch bill charged requests called")

		for i := range tasks { // update task info
			tasks[i].TxnHash = txnHash
			tasks[i].SubmittedAt = time.Now()
		}

		successTasks = append(successTasks, tasks...)
	}

	return
}

func (worker *BlockchainWorker) updateBillTaskStatus(tasks []*BillTask, status uint8) {
	util.KLock(service.KeyBillingLocker)
	defer util.KUnlock(service.KeyBillingLocker)

	for i := range tasks {
		tasks[i].Status = status

		updates := map[string]interface{}{
			"status":   status,
			"txn_hash": tasks[i].TxnHash,
			"memo":     tasks[i].Memo,
		}

		res := worker.sqliteStore.Model(&model.Bill{}).Where("id = ?", tasks[i].ID).Updates(updates)
		if err := res.Error; err != nil {
			logrus.WithField("task", tasks[i]).
				WithError(err).
				Error("Blockchain worker failed to update bill task status")
		}
	}
}

var pollSignal = make(chan time.Time, 1)

func TriggerPoll() {
	pollSignal <- time.Now()
}
func (worker *BlockchainWorker) poll() {
	ticker := time.NewTicker(worker.PollingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
		case <-pollSignal:
			fmt.Printf("hit ------ \n")
			if err := worker.pollOnce(); err != nil {
				logrus.WithError(err).
					Error("Blockchain worker failed to poll bill tasks for settlement")
			}
		}
	}
}

// poll polls for bills to settle on the blockchain.
func (w *BlockchainWorker) pollOnce() error {
	if overloaded, reason := w.isOverloaded(); overloaded { // overload protection
		logrus.WithField("reason", reason).Debug("Blockchain worker skipped polling due to overload")
		return nil
	}

	util.KLock(service.KeyBillingLocker)
	defer util.KUnlock(service.KeyBillingLocker)

	var bills []*model.Bill
	db := w.sqliteStore.Order("id ASC").Limit(w.PollingBatchSize)

	res := db.Find(&bills, "status = ?", model.BillStatusCreated)
	if err := res.Error; err != nil {
		return errors.WithMessage(err, "failed to load bills")
	}

	if len(bills) == 0 {
		logrus.Debug("Blockchain worker polled without any charge bills")
		return nil
	}

	updateBillIds := []uint64{}
	for i := range bills {
		updateBillIds = append(updateBillIds, bills[i].ID)
	}

	db = w.sqliteStore.Model(&model.Bill{}).Where("id IN ?", updateBillIds)
	res = db.Update("status", model.BillStatusSubmitting)
	if err := res.Error; err != nil {
		return errors.WithMessage(err, "failed to update bill status")
	}

	billTasks := make([]*BillTask, 0, len(bills))
	for _, bill := range bills {
		statements := w.billSvc.GetAndDelStatements(bill.ID)
		billTasks = append(billTasks, NewBillTask(bill, statements))
	}

	logrus.WithField("taskBillIds", updateBillIds).
		Debug("Blockchain worker polled task bills for settlement")
	w.billSettlementQueue <- billTasks

	return nil
}

// isOverloaded check if the worker is overloaded
func (w *BlockchainWorker) isOverloaded() (bool, string) {
	// cumulative ratio should not be more than 75%, otherwise regarded as overloaded
	maxCumulativeSize := billingQueueSize * 3 / 4

	if len(w.billSettlementQueue) > maxCumulativeSize {
		return true, "cumulative ratio for settlement queue too high"
	}

	if len(w.billConfirmQueue) > maxCumulativeSize {
		return true, "cumulative ratio for txn confirmation queue too high"
	}

	return false, ""
}
