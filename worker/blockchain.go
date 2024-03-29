package worker

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"sync"
	"time"

	"github.com/Conflux-Chain/go-conflux-util/viper"
	"github.com/Conflux-Chain/web3pay-service/blockchain"
	"github.com/Conflux-Chain/web3pay-service/contract"
	"github.com/Conflux-Chain/web3pay-service/metrics"
	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/service"
	"github.com/Conflux-Chain/web3pay-service/store/sqlite"
	"github.com/Conflux-Chain/web3pay-service/util"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/schollz/jsonstore"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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
	PersistOnShutdown       bool          `default:"false"`
	PersistencePath         string        `default:"./wdata"`
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

func (worker *BlockchainWorker) Run(ctx context.Context) {
	var wg sync.WaitGroup

	// start polling
	wg.Add(1)
	go func() {
		defer wg.Done()
		worker.poll(ctx)
	}()

	// start confirming
	wg.Add(1)
	go func() {
		defer wg.Done()
		worker.confirm(ctx)
	}()

	// start settling
	wg.Add(1)
	go func() {
		defer wg.Done()
		worker.settle(ctx)
	}()

	wg.Wait()

	// persist billing status after shut down
	if worker.PersistOnShutdown {
		worker.persist()
	}
}

func (worker *BlockchainWorker) confirm(ctx context.Context) {
	// TODO: task confirmations are order independent, multiple consumers can be utilized
	// to increase workload capacity.
	for {
		select {
		case billTasks := <-worker.billConfirmQueue:
			worker.confirmBillTasks(billTasks)
		case <-ctx.Done():
			logrus.Info("Blockchain worker confirming completed")
			return
		}
	}
}

func (worker *BlockchainWorker) confirmBillTasks(billTasks []*BillTask) {
	var successTasks, reconfirmTasks, retryTasks []*BillTask

	start := time.Now()
	defer metrics.Worker.ConfirmOnceQps().UpdateSince(start)

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

	metrics.Worker.UpdateConfirmOnceSize(
		len(billTasks), len(successTasks), len(retryTasks), len(reconfirmTasks),
	)

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
			common.HexToAddress(tasks[i].App),
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

	for _, billId := range delBillIds {
		worker.billSvc.DelStatements(billId)
	}

	if err := worker.sqliteStore.Delete(&model.Bill{}, delBillIds).Error; err != nil {
		logrus.WithField("delBillIds", delBillIds).
			WithError(err).
			Error("Blockchain worker failed to delete task bills")
	}
}

func (worker *BlockchainWorker) settle(ctx context.Context) {
	for {
		select {
		case billTasks := <-worker.billSettlementQueue:
			worker.settleBillTasks(billTasks)
		case <-ctx.Done():
			logrus.Info("Blockchain worker settling completed")
			return
		}
	}
}

func (worker *BlockchainWorker) settleBillTasks(billTasks []*BillTask) {
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

	successTasks, failureTasks := worker.doSettleBillTasks(todoTasks)

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

func (worker *BlockchainWorker) doSettleBillTasks(billTasks []*BillTask) (successTasks, failureTasks []*BillTask) {
	start := time.Now()
	defer metrics.Worker.SettleOnceQps().UpdateSince(start)

	// split tasks && requests group by APP for batch operation
	appTasks := make(map[string][]*BillTask)
	appRequests := make(map[string][]contract.IAppConfigChargeRequest)

	for _, task := range billTasks {
		var details []contract.IAppConfigResourceUseDetail
		for resourceIndex, callTimes := range task.Statements {
			details = append(details, contract.IAppConfigResourceUseDetail{
				Id:    resourceIndex,
				Times: big.NewInt(callTimes),
			})
		}

		appRequests[task.App] = append(appRequests[task.App], contract.IAppConfigChargeRequest{
			Account:   common.HexToAddress(task.Address),
			Amount:    task.Fee.BigInt(),
			UseDetail: details,
		})

		appTasks[task.App] = append(appTasks[task.App], task)
	}

	// call batch charge contract method for settlement
	for app, req := range appRequests {
		tasks := appTasks[app]
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
			"APP":                  app,
			"batchChargeBillTasks": taskInfos,
		})

		// prepare the transaction to be sent
		sendTxn, err := worker.provider.BatchChargeAppBills(
			&bind.TransactOpts{NoSend: true}, common.HexToAddress(app), req,
		)

		var rawTxnData []byte
		var txnHash common.Hash
		var isRawTxnSendErr bool

		if err != nil {
			err = errors.WithMessage(err, "failed to prepare the batch bill charging transaction")
			goto failure
		}

		// RLP encoding the to send transaction
		rawTxnData, err = sendTxn.MarshalBinary()
		if err != nil {
			err = errors.WithMessage(err, "failed to RLP encoding the batch bill charging transaction")
			goto failure
		}

		// submmit to the transaction pool
		txnHash, err = worker.provider.SendRawTransaction(rawTxnData)
		if err == nil && txnHash != sendTxn.Hash() {
			err = errors.New("unexpected transaction hash after raw txn sent")
			goto failure
		}

		if err != nil {
			err = errors.New("failed to send the batch bill charging raw transaction")
			isRawTxnSendErr, txnHash = true, sendTxn.Hash()
			goto failure
		}

	success:
		logger.WithField("txnHash", txnHash).Debug("Blockchain worker batch bill charged requests called")

		for i := range tasks { // update task info
			tasks[i].TxnHash = txnHash.String()
			tasks[i].SubmittedAt = time.Now()
		}

		successTasks = append(successTasks, tasks...)
		continue

	failure:
		if isRawTxnSendErr {
			// distinguish between different error types
			// https://developer.confluxnetwork.org/sending-tx/en/transaction_send_common_error
			switch {
			case isTxnAlreadyExistError(err):
				// txn already exists
				logger.WithError(err).Info("Blockchain worker found duplicate txn for bill charging request")
				goto success
			case isTxnPollFullError(err):
				// txn pool full
			case isTxnNonceAlreadyExistsError(err):
				// txn nonce already exists
			case isTxnNonceTooStaleError(err):
				// txn nonce too stale
			case isTxnNonceTooFutureError(err):
				// txn nonce too future
			case isTxnGasTooSmallError(err):
				// txn gas too small
			case isTxnGasTooLargeError(err):
				// txn gas too large
			case isTxnGasPriceIsZeroError(err):
				// txn gas price is zero
			default:
				// In case of the same billing task submitted multiple times, we'd better to ensure the
				// transaction not put into the transaction pool yet due to some network failure.
				for {
					txn, err := worker.provider.TransactionByHash(txnHash)
					if err == nil {
						if txn != nil { // txn already exists
							goto success
						}
						break
					}

					logger.WithError(err).Info("Blockchain worker failed to check batch bill charging transaction")
					time.Sleep(time.Second)
				}
			}
		}

		logger.WithError(err).Info("Blockchain worker failed to request batch billing charge")

		failureTasks = append(failureTasks, tasks...)
		for i := range tasks { // update task memo
			tasks[i].Memo = err.Error()
		}
	}

	metrics.Worker.UpdateSettleOnceSize(len(billTasks), len(successTasks), len(failureTasks))
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

func (worker *BlockchainWorker) poll(ctx context.Context) {
	ticker := time.NewTicker(worker.PollingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := worker.pollOnce(); err != nil {
				logrus.WithError(err).
					Error("Blockchain worker failed to poll bill tasks for settlement")
			}
		case <-ctx.Done():
			logrus.Info("Blockchain worker polling completed")
			return
		}
	}
}

// poll polls for bills to settle on the blockchain.
func (w *BlockchainWorker) pollOnce() error {
	if overloaded, reason := w.isOverloaded(); overloaded { // overload protection
		logrus.WithField("reason", reason).Warn("Blockchain worker skipped polling due to overload")
		return nil
	}

	start := time.Now()

	util.KLock(service.KeyBillingLocker)
	defer util.KUnlock(service.KeyBillingLocker)

	var bills []*model.Bill
	db := w.sqliteStore.Order("id ASC").Limit(w.PollingBatchSize)

	res := db.Find(&bills, "status = ?", model.BillStatusCreated)
	if err := res.Error; err != nil {
		metrics.Worker.PollOnceQps(err).UpdateSince(start)
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
		metrics.Worker.PollOnceQps(err).UpdateSince(start)
		return errors.WithMessage(err, "failed to update bill status")
	}

	billTasks := make([]*BillTask, 0, len(bills))
	for _, bill := range bills {
		statements := w.billSvc.GetStatements(bill.ID)
		billTasks = append(billTasks, NewBillTask(bill, statements))
	}

	logrus.WithField("taskBillIds", updateBillIds).
		Debug("Blockchain worker polled task bills for settlement")
	w.billSettlementQueue <- billTasks

	metrics.Worker.PollOnceQps(nil).UpdateSince(start)
	metrics.Worker.PollOnceSize().Update(int64(len(billTasks)))

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

func (w *BlockchainWorker) persist() {
	var bills []*model.Bill
	ks := new(jsonstore.JSONStore)

	// iterate all existed bills
	res := w.sqliteStore.FindInBatches(&bills, w.PollingBatchSize, func(tx *gorm.DB, batch int) error {
		for _, bill := range bills {
			task := &BillTask{
				Bill:       bill,
				Statements: w.billSvc.GetStatements(bill.ID),
			}

			if err := ks.Set(fmt.Sprintf("bill:%v", bill.ID), task); err != nil {
				logrus.WithField("task", task).Error("Failed to json store set billing task")
			}
		}

		return nil
	})

	if res.Error != nil {
		logrus.WithError(res.Error).Error("Failed to list all bills for persistence")
	}

	if len(ks.Data) == 0 { // no bills?
		return
	}

	// prepare persistence folder
	err := os.MkdirAll(w.PersistencePath, os.ModePerm)
	if err != nil {
		logrus.WithField("path", w.PersistencePath).Error("Failed to prepare path folder")
		return
	}

	// save file
	fn := fmt.Sprintf("%v/ps.%v.json.gz", w.PersistencePath, time.Now().Unix())
	if err := jsonstore.Save(ks, fn); err != nil {
		logrus.WithField("fileName", fn).WithError(err).Error("Failed to save json store file")
	}
}
