package service

import (
	"math/big"

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
	KeyBillingLocker = util.MutexKey("BillingLocker")
)

type BillingRequest struct {
	ResourceId string         `json:"resourceId"`
	DryRun     bool           `json:"dryRun"`
	App        common.Address `json:"-"`
	Customer   common.Address `json:"-"`
}

type BillingReceipt struct {
	App          common.Address             `json:"app,omitempty"`
	Customer     common.Address             `json:"customer,omitempty"`
	Balance      string                     `json:"balance,omitempty"`
	ResourceFees map[string]decimal.Decimal `json:"resourceFees,omitempty"` // resource ID => fee
}

type BillingBatchRequest struct {
	ResourceUses map[string]int64 `json:"resourceUses"`
	DryRun       bool             `json:"dryRun"`
	App          common.Address   `json:"-"`
	Customer     common.Address   `json:"-"`
}

type BillingBatchReceipt = BillingReceipt

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

func (bs *BillingService) BillBatch(ctx context.Context, req *BillingBatchRequest) (*BillingReceipt, error) {
	if err := bs.chainSvc.ValidateBillingApp(req.App); err != nil {
		return nil, err
	}

	var totalCost decimal.Decimal
	resourceCosts := make(map[string]decimal.Decimal)
	statements := make(map[uint32]int64)

	for resourceId, cnt := range req.ResourceUses {
		resource, err := bs.chainSvc.GetAppConfigResourceWithId(req.App, resourceId)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"resourceId": resourceId, "app": req.App,
			}).WithError(err).Debug("Failed to retrieve APP config resource for billing")
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

	// get account status
	appAccount, err := bs.chainSvc.GetOrFetchAppAccountStatus(req.App, req.Customer)
	if err != nil {
		logger.WithError(err).Info("Billing failed to get account status")
		return nil, err
	}

	logger = logger.WithField("account", appAccount)
	receipt := &BillingReceipt{
		App: req.App, Customer: req.Customer, ResourceFees: resourceCosts,
	}

	// account frozen?
	if appAccount.IsFrozen() {
		logger.Debug("Billing skipped due to customer account frozen")
		return nil, model.ErrAccountFrozen.WithData(receipt)
	}

	receipt.Balance = appAccount.TotalBalance().String()

	// insufficient balance?
	if appAccount.TotalBalance().Cmp(totalCost.BigInt()) < 0 {
		logger.Debug("Billing skipped due to insufficient balance")
		return nil, model.ErrInsufficentBalance.WithData(receipt)
	}

	if req.DryRun { // for simulation only?
		logger.Debug("Billing skipped for dry run")
		receipt.Balance = big.NewInt(0).
			Sub(appAccount.TotalBalance(), totalCost.BigInt()).String()
		return receipt, nil
	}

	util.KLock(KeyBillingLocker)
	defer util.KUnlock(KeyBillingLocker)

	app, addr := appAccount.App, appAccount.Address
	if err = bs.sqliteStore.Transaction(func(tx *gorm.DB) error {
		fee := totalCost.BigInt()

		bill, err := bs.sqliteStore.UpsertBill(tx, app, addr, fee)
		if err != nil {
			logger.WithError(err).Error("Billing failed to upsert bill")
			return err
		}

		deducted, err := bs.chainSvc.WithholdAccountFee(req.App, req.Customer, fee)
		if err != nil {
			logger.WithError(err).Error("Billing failed to withhold account fee")
			return err
		}

		if deducted {
			bs.collectBillStatements(bill.ID, statements)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	receipt.Balance = appAccount.TotalBalance().String()
	return receipt, nil
}

func (bs *BillingService) Bill(ctx context.Context, req *BillingRequest) (*BillingReceipt, error) {
	return bs.BillBatch(ctx, &BillingBatchRequest{
		ResourceUses: map[string]int64{req.ResourceId: 1},
		DryRun:       req.DryRun,
		App:          req.App,
		Customer:     req.Customer,
	})
}

func (bs *BillingService) GetAndDelStatements(billID uint64) map[uint32]int64 {
	statements := bs.billingStatements[billID]
	delete(bs.billingStatements, billID)

	return statements
}

func (bs *BillingService) GetStatements(billID uint64) map[uint32]int64 {
	return bs.billingStatements[billID]
}

func (bs *BillingService) DelStatements(billID uint64) {
	delete(bs.billingStatements, billID)
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
