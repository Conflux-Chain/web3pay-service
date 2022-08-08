package worker

import (
	"time"

	"github.com/Conflux-Chain/web3pay-service/blockchain"
	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/service"
	"github.com/Conflux-Chain/web3pay-service/store/sqlite"
	"github.com/Conflux-Chain/web3pay-service/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	pollingInterval         = time.Second * 90
	pollingBatchSize        = 20
	billSettlementQueueSize = 5000
)

type BlockchainWorker struct {
	provider            *blockchain.Provider
	sqliteStore         *sqlite.SqliteStore
	billSvc             *service.BillingService
	chainSvc            *service.BlockchainService
	billSettlementQueue chan []*model.Bill
}

func NewBlockchainWorker(
	provider *blockchain.Provider, sqliteStore *sqlite.SqliteStore,
	billSvc *service.BillingService, chainSvc *service.BlockchainService,
) *BlockchainWorker {
	// TODO: also check submitting/submitted tasks status?
	return &BlockchainWorker{
		sqliteStore: sqliteStore, provider: provider,
		billSvc: billSvc, chainSvc: chainSvc,
		billSettlementQueue: make(chan []*model.Bill, billSettlementQueueSize),
	}
}

func (worker *BlockchainWorker) Run() {
	go worker.poll()

	for bills := range worker.billSettlementQueue {
		// TODO:
		// 1. split bills group by APP coin address;
		// 2. call batch charge contract method for bill settlement;
		// 3. update bill submitted status once txn submitted ok;
		// 4. txn finality check to update bill status and account balance;
		_ = bills
	}
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
	if len(w.billSettlementQueue) >= billSettlementQueueSize { // channel full?
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

	w.billSettlementQueue <- bills

	return nil
}
