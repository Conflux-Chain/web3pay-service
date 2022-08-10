package worker

import (
	"math/big"
	"time"

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
	pollingInterval  = time.Second * 90
	pollingBatchSize = 20
	billQueueSize    = 5000

	maxSettlementRetries    = 5
	minConfirmBlocks        = 30
	maxPendingAwaitDuration = time.Second * 30
)

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

type BlockchainWorker struct {
	provider            *blockchain.Provider
	sqliteStore         *sqlite.SqliteStore
	billSvc             *service.BillingService
	chainSvc            *service.BlockchainService
	billSettlementQueue chan []*BillTask
	billConfirmQueue    chan []*BillTask
	billRetryQueue      chan []*BillTask
}

func NewBlockchainWorker(
	provider *blockchain.Provider, sqliteStore *sqlite.SqliteStore,
	billSvc *service.BillingService, chainSvc *service.BlockchainService,
) *BlockchainWorker {
	// TODO: also check submitting/submitted tasks status?
	return &BlockchainWorker{
		sqliteStore: sqliteStore, provider: provider,
		billSvc: billSvc, chainSvc: chainSvc,
		billSettlementQueue: make(chan []*BillTask, billQueueSize),
		billConfirmQueue:    make(chan []*BillTask, billQueueSize),
		billRetryQueue:      make(chan []*BillTask, billQueueSize),
	}
}

func (worker *BlockchainWorker) Run() {
	go worker.poll()
	go worker.confirm()
	go worker.retry()

	worker.settle()
}

func (worker *BlockchainWorker) confirm() {
	for billTasks := range worker.billConfirmQueue {
		worker.confirmBillTasks(billTasks)
	}
}

func (worker *BlockchainWorker) confirmBillTasks(billTasks []*BillTask) {
	var successTasks, reconfirmTasks, retryTasks []*BillTask

	for i := range billTasks {
		logrus.WithField("task", billTasks[i]).
			Debug("Blockchain worker confirming bill task")

		txnHash := common.HexToHash(billTasks[i].TxnHash)
		txn, err := worker.provider.TransactionByHash(txnHash)
		if err != nil {
			logrus.WithField("task", billTasks[i]).WithError(err).
				Info("Blockchain worker failed to get txn by hash during confirm")

			reconfirmTasks = append(reconfirmTasks, billTasks[i])
			continue
		}

		if txn == nil { // transaction dropped?
			logrus.WithField("task", billTasks[i]).
				Debug("Blockchain worker got nil txn for confirming task")

			retryTasks = append(retryTasks, billTasks[i])
			continue
		}

		if txn.BlockHash == nil { // transaction pending (not mined or executed)?
			if time.Now().After(billTasks[i].SubmittedAt.Add(maxPendingAwaitDuration)) { // pending too long?
				logrus.WithField("task", billTasks[i]).
					Debug("Blockchain worker got timeout pending txn for confirming task")

				retryTasks = append(retryTasks, billTasks[i])
			} else {
				logrus.WithField("task", billTasks[i]).
					Debug("Blockchain worker got pending txn for confirming task")

				reconfirmTasks = append(reconfirmTasks, billTasks[i])
			}

			continue
		}

		if txn.Status == nil || *txn.Status != ethtypes.ReceiptStatusSuccessful { // transaction failed?
			logrus.WithFields(logrus.Fields{
				"task":      billTasks[i],
				"txnStatus": txn.Status,
			}).Debug("Blockchain worker got non-successful txn for confirming task")

			retryTasks = append(retryTasks, billTasks[i])
			continue
		}

		latestBlock, err := worker.provider.BlockNumber()
		if err != nil {
			logrus.WithField("task", billTasks[i]).WithError(err).
				Info("Blockchain worker failed to get latest block number during confirm")

			reconfirmTasks = append(reconfirmTasks, billTasks[i])
			continue
		}

		if (txn.BlockNumber.Uint64() + minConfirmBlocks) > latestBlock.Uint64() {
			logrus.WithField("task", billTasks[i]).
				Debug("Blockchain worker got not enough blocks for confirming task")

			// no enough blocks confirmed?
			reconfirmTasks = append(reconfirmTasks, billTasks[i])
			continue
		}

		logrus.WithField("task", billTasks[i]).Debug("Blockchain worker confirmed txn for task")
		successTasks = append(successTasks, billTasks[i])
	}

	if len(successTasks) > 0 {
		worker.finishTasks(successTasks)
	}

	if len(retryTasks) > 0 {
		worker.billRetryQueue <- retryTasks
	}

	if len(reconfirmTasks) > 0 {
		// wait for a while for re-confirm tasks
		time.AfterFunc(time.Second, func() {
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
			logrus.WithFields(logrus.Fields{
				"task":     tasks[i],
				"exempted": exempted,
			}).WithError(err).Info("Blockchain worker failed to exempt account fee")
		}
	}

	util.KLock(service.KeyBillingChargeLocker)
	defer util.KUnlock(service.KeyBillingChargeLocker)

	if err := worker.sqliteStore.Delete(&model.Bill{}, delBillIds).Error; err != nil {
		logrus.WithField("delBillIds", delBillIds).
			WithError(err).
			Info("Blockchain worker failed to delete bills")
	}
}

func (worker *BlockchainWorker) retry() {
	for billTasks := range worker.billRetryQueue {
		worker.retryBillTasks(billTasks)
	}
}

func (worker *BlockchainWorker) retryBillTasks(billTasks []*BillTask) {
	var doTasks, deadTasks, retryTasks []*BillTask

	// filter dead tasks
	for i := range billTasks {
		if billTasks[i].TryTimes > maxSettlementRetries {
			deadTasks = append(deadTasks, billTasks[i])
			continue
		}

		doTasks = append(doTasks, billTasks[i])
	}

	successTasks, failureTasks := worker.settleBillTasks(doTasks)

	// update bill submitted status
	if len(successTasks) > 0 {
		err := worker.updateBillTaskStatus(successTasks, model.BillStatusSubmitted)
		if err != nil {
			logrus.WithField("updateBillTasks", successTasks).
				WithError(err).
				Info("Blockchain worker failed to update bill submitted status during retry")
		}

		// put into confirm queue
		worker.billConfirmQueue <- successTasks
	}

	// filter retry tasks && dead tasks
	for i := range failureTasks {
		if failureTasks[i].TryTimes <= maxSettlementRetries {
			retryTasks = append(retryTasks, failureTasks[i])
			continue
		}

		logrus.WithField("task", failureTasks[i]).
			Info("Blockchain worker failed to handle (dead) bill task after too many retries")
		deadTasks = append(deadTasks, failureTasks[i])
	}

	// update bill failed status
	if len(deadTasks) > 0 {
		err := worker.updateBillTaskStatus(successTasks, model.BillStatusFailed)
		if err != nil {
			logrus.WithField("updateBillTasks", successTasks).
				WithError(err).
				Info("Blockchain worker failed to update bill failed status during retry")
		}
	}

	if len(retryTasks) > 0 {
		// put into retry queue
		worker.billRetryQueue <- retryTasks
	}
}

func (worker *BlockchainWorker) settle() {
	for billTasks := range worker.billSettlementQueue {
		successTasks, failureTasks := worker.settleBillTasks(billTasks)

		// update bill submitted status
		if len(successTasks) > 0 {
			err := worker.updateBillTaskStatus(successTasks, model.BillStatusSubmitted)
			if err != nil {
				logrus.WithField("updateBillTasks", successTasks).
					WithError(err).
					Info("Blockchain worker failed to update bill submitted status")
			}

			// put into confirm queue
			worker.billConfirmQueue <- successTasks
		}

		if len(failureTasks) > 0 {
			// put into retry queue
			worker.billRetryQueue <- failureTasks
		}
	}
}

func (worker *BlockchainWorker) settleBillTasks(billTasks []*BillTask) (successTasks, failureTasks []*BillTask) {
	// split tasks && requests group by APP coin for batch operation
	coinTasks := make(map[string][]*BillTask)
	coinRequests := make(map[string][]contract.APPCoinChargeRequest)

	for _, task := range billTasks {
		coinTasks[task.Coin] = append(coinTasks[task.Coin], task)

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

		task.TryTimes++
	}

	// call batch charge contract method for settlement
	for coin, req := range coinRequests {
		txn, err := worker.provider.BatchChargeAppCoinBills(
			&bind.TransactOpts{}, common.HexToAddress(coin), req,
		)

		tasks := coinTasks[coin]

		// TODO: distinguish between different error types
		// https://developer.confluxnetwork.org/sending-tx/en/transaction_send_common_error
		// eg.,
		// 1. transaction pool full error
		// 2. duplicated txn error
		// 3. network error
		if err != nil {
			logrus.WithField("tasks", tasks).
				WithError(err).
				Debug("Blockchain worker failed to submit batch charge APP coin request")
			failureTasks = append(failureTasks, tasks...)
			continue
		}

		txnHash := txn.Hash().String()
		for i := range tasks {
			tasks[i].TxnHash = txnHash
			tasks[i].SubmittedAt = time.Now()
		}

		successTasks = append(successTasks, tasks...)

		logrus.WithField("task", tasks).
			Debug("Blockchain worker succeeded to submit batch charge APP coin request")
	}

	return
}

func (worker *BlockchainWorker) updateBillTaskStatus(tasks []*BillTask, status uint8) error {
	updateBillIds := []uint64{}
	for i := range tasks {
		if tasks[i].Status == status {
			continue
		}

		updateBillIds = append(updateBillIds, tasks[i].ID)
		tasks[i].Status = status
	}

	if len(updateBillIds) == 0 {
		return nil
	}

	util.KLock(service.KeyBillingChargeLocker)
	defer util.KUnlock(service.KeyBillingChargeLocker)

	return worker.sqliteStore.Model(&model.Bill{}).
		Where("id IN ?", updateBillIds).
		Update("status", status).Error
}

func (worker *BlockchainWorker) poll() {
	ticker := time.NewTicker(pollingInterval)
	defer ticker.Stop()

	for range ticker.C {
		if err := worker.pollOnce(); err != nil {
			logrus.WithError(err).
				Error("Blockchain worker failed to poll bills for settlement")
		}
	}
}

// poll polls for bills to settle on the blockchain.
func (w *BlockchainWorker) pollOnce() error {
	if len(w.billSettlementQueue) >= billQueueSize { // channel full?
		logrus.Debug("Blockchain worker skipped to poll bill tasks due to settlement queue full")
		return nil
	}

	util.KLock(service.KeyBillingChargeLocker)
	defer util.KUnlock(service.KeyBillingChargeLocker)

	var bills []*model.Bill
	db := w.sqliteStore.Order("id ASC").Limit(pollingBatchSize)

	res := db.Find(&bills, "status = ?", model.BillStatusCreated)
	if err := res.Error; err != nil {
		return errors.WithMessage(err, "failed to load bills")
	}

	if len(bills) == 0 {
		logrus.Debug("Blockchain worker skipped without any concerned bills")
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

	logrus.WithFields(logrus.Fields{
		"batchBills":   bills,
		"updatedBills": res.RowsAffected,
	}).Debug("Blockchain worker fetched bills for settlement")

	billTasks := make([]*BillTask, 0, len(bills))
	for _, bill := range bills {
		statements := w.billSvc.GetAndDelStatements(bill.ID)
		billTasks = append(billTasks, NewBillTask(bill, statements))
	}

	w.billSettlementQueue <- billTasks

	return nil
}
