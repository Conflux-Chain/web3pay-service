package service

import (
	"fmt"
	"math/big"

	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/store/sqlite"
	"github.com/Conflux-Chain/web3pay-service/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

const (
	KeyBillingChargeLocker = util.MutexKey("BillingChargeLocker")
)

type ChargeRequest struct {
	ResourceId string
	DryRun     bool
	AppCoin    common.Address
	Customer   common.Address
}

type ChargeReceipt struct {
	ResourceId string
	Fee        string
}

type BillingService struct {
	chainSvc    *BlockchainService
	sqliteStore *sqlite.SqliteStore
	// map{ bill ID => map{ resource index = > API call times } }
	billingStatements map[uint64]map[uint32]int64
}

func NewBillingService(sqliteStore *sqlite.SqliteStore, chainSvc *BlockchainService) *BillingService {
	return &BillingService{
		sqliteStore:       sqliteStore,
		chainSvc:          chainSvc,
		billingStatements: make(map[uint64]map[uint32]int64),
	}
}

func (bs *BillingService) Charge(ctx context.Context, req *ChargeRequest) (*ChargeReceipt, error) {
	logger := logrus.WithField("request", req)

	resource, err := bs.chainSvc.GetAppCoinResourceWithId(req.AppCoin, req.ResourceId)
	if err != nil {
		logger.WithError(err).Debug("Billing charge failed to retrieve APP coin resource")
		return nil, err
	}

	fee := resource.Weight
	logger = logger.WithField("fee", fee.String())

	if fee.Cmp(big.NewInt(0)) <= 0 { // no fee charged
		logger.WithError(err).Debug("Billing charge skipped with no fee to be charged")

		return &ChargeReceipt{
			ResourceId: resource.ResourceId,
			Fee:        fmt.Sprintf("%v", fee.String()),
		}, nil
	}

	// get account status
	appCoinAccount, err := bs.chainSvc.GetOrFetchAccountStatus(req.AppCoin, req.Customer)
	if err != nil {
		logger.WithError(err).Info("Billing charge failed to get account status")
		return nil, err
	}

	logger = logger.WithField("account", appCoinAccount)

	// account frozen?
	if appCoinAccount.IsFrozen() {
		logger.Debug("Billing charge skipped due to customer account frozen")
		return nil, model.ErrAccountAddrFrozen
	}

	// insufficient balance?
	if appCoinAccount.TotalBalance().Cmp(fee) < 0 {
		logger.Debug("Billing charge skipped due to insufficient balance")
		return nil, model.ErrInsufficentBalance
	}

	if req.DryRun { // for simulation only?
		logger.Debug("Billing charge skipped for dry run")
		return &ChargeReceipt{
			ResourceId: resource.ResourceId,
			Fee:        fmt.Sprintf("%v", fee.String()),
		}, nil
	}

	util.KLock(KeyBillingChargeLocker)
	defer util.KUnlock(KeyBillingChargeLocker)

	coin, addr := appCoinAccount.Coin, appCoinAccount.Address
	if err = bs.sqliteStore.Transaction(func(tx *gorm.DB) error {
		bill, err := bs.sqliteStore.UpsertBill(tx, coin, addr, fee)
		if err != nil {
			logger.WithError(err).Error("Billing charge failed to upsert bill")
			return err
		}

		deducted, err := bs.chainSvc.WithholdAccountFee(req.AppCoin, req.Customer, fee)
		if err != nil {
			logger.WithError(err).Error("Billing charge failed to withhold account fee")
			return err
		}

		if deducted {
			bs.recordApiCallOnce(bill.ID, resource.Index)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &ChargeReceipt{
		ResourceId: resource.ResourceId,
		Fee:        fee.String(),
	}, nil
}

func (bs *BillingService) GetAndDelStatements(billID uint64) map[uint32]int64 {
	statements := bs.billingStatements[billID]
	delete(bs.billingStatements, billID)

	return statements
}

func (bs *BillingService) recordApiCallOnce(billID uint64, resourceIndex uint32) {
	if _, ok := bs.billingStatements[billID]; !ok {
		bs.billingStatements[billID] = make(map[uint32]int64)
	}

	bs.billingStatements[billID][resourceIndex]++
}
