package service

import (
	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/store/sqlite"
	"github.com/Conflux-Chain/web3pay-service/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

const (
	KeyBillingChargeLocker = util.MutexKey("BillingChargeLocker")
)

type BillingChargeRequest struct {
	ResourceId string         `json:"resourceId"`
	DryRun     bool           `json:"dryRun"`
	AppCoin    common.Address `json:"-"`
	Customer   common.Address `json:"-"`
}

type BillingChargeReceipt struct {
	ResourceId string `json:"resourceId"`
	Fee        string `json:"fee"`
}

type BillingChargeBatchRequest struct {
	ResourceUses map[string]int64 `json:"resourceUses"`
	DryRun       bool             `json:"dryRun"`
	AppCoin      common.Address   `json:"-"`
	Customer     common.Address   `json:"-"`
}

type BillingChargeBatchReceipt struct {
	ResourceFees map[string]decimal.Decimal `json:"resourceFees"`
}

type BillingService struct {
	chainSvc    *BlockchainService
	sqliteStore *sqlite.SqliteStore
	// map{ bill ID => map{ resource index = > use count } }
	billingStatements map[uint64]map[uint32]int64
}

func NewBillingService(sqliteStore *sqlite.SqliteStore, chainSvc *BlockchainService) *BillingService {
	return &BillingService{
		sqliteStore:       sqliteStore,
		chainSvc:          chainSvc,
		billingStatements: make(map[uint64]map[uint32]int64),
	}
}

func (bs *BillingService) ChargeBatch(ctx context.Context, req *BillingChargeBatchRequest) (*BillingChargeBatchReceipt, error) {
	var totalCost decimal.Decimal
	resourceCosts := make(map[string]decimal.Decimal)
	statements := make(map[uint32]int64)

	for resourceId, cnt := range req.ResourceUses {
		resource, err := bs.chainSvc.GetAppCoinResourceWithId(req.AppCoin, resourceId)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"resourceId": resourceId,
				"appCoin":    req.AppCoin,
			}).WithError(err).Debug("Billing charge failed to retrieve APP coin resource")
			return nil, err
		}

		delta := decimal.NewFromBigInt(resource.Weight, 0).Mul(decimal.NewFromInt(cnt))
		cost := resourceCosts[resource.ResourceId].Add(delta)
		resourceCosts[resource.ResourceId] = cost

		statements[resource.Index] += cnt
		totalCost = totalCost.Add(delta)
	}

	logger := logrus.WithFields(logrus.Fields{
		"request":       req,
		"totalCost":     totalCost,
		"resourceCosts": resourceCosts,
	})

	if totalCost.IsZero() { // no fee charged
		logger.Debug("Billing charge skipped with no fee to be charged")
		return &BillingChargeBatchReceipt{ResourceFees: resourceCosts}, nil
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
	if appCoinAccount.TotalBalance().Cmp(totalCost.BigInt()) < 0 {
		logger.Debug("Billing charge skipped due to insufficient balance")
		return nil, model.ErrInsufficentBalance
	}

	if req.DryRun { // for simulation only?
		logger.Debug("Billing charge skipped for dry run")
		return &BillingChargeBatchReceipt{ResourceFees: resourceCosts}, nil
	}

	util.KLock(KeyBillingChargeLocker)
	defer util.KUnlock(KeyBillingChargeLocker)

	coin, addr := appCoinAccount.Coin, appCoinAccount.Address
	if err = bs.sqliteStore.Transaction(func(tx *gorm.DB) error {
		fee := totalCost.BigInt()

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
			bs.collectBillStatements(bill.ID, statements)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &BillingChargeBatchReceipt{ResourceFees: resourceCosts}, nil
}

func (bs *BillingService) Charge(ctx context.Context, req *BillingChargeRequest) (*BillingChargeReceipt, error) {
	batchReceipt, err := bs.ChargeBatch(ctx, &BillingChargeBatchRequest{
		ResourceUses: map[string]int64{req.ResourceId: 1},
		DryRun:       req.DryRun,
		AppCoin:      req.AppCoin,
		Customer:     req.Customer,
	})
	if err != nil {
		return nil, err
	}

	for resourceId, fee := range batchReceipt.ResourceFees {
		return &BillingChargeReceipt{
			ResourceId: resourceId,
			Fee:        fee.String(),
		}, nil
	}

	return nil, nil
}

func (bs *BillingService) GetAndDelStatements(billID uint64) map[uint32]int64 {
	statements := bs.billingStatements[billID]
	delete(bs.billingStatements, billID)

	return statements
}

func (bs *BillingService) collectBillStatements(billID uint64, resourceStatements map[uint32]int64) {
	if _, ok := bs.billingStatements[billID]; !ok {
		bs.billingStatements[billID] = resourceStatements
		return
	}

	for resourceIndex, useCnt := range resourceStatements {
		bs.billingStatements[billID][resourceIndex] += useCnt
	}
}
