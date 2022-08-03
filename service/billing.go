package service

import (
	"time"

	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/store/sqlite"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ChargeRequest struct {
	ResourceId   string
	DryRun       bool
	AppCoinAddr  common.Address
	CustomerAddr common.Address
}

type ChargeReceipt struct {
	ResourceId string
	Fee        uint64
}

type BillingService struct {
	chainSvc *BlockchainService
	store    *sqlite.SqliteStore
}

func NewBillingService(store *sqlite.SqliteStore, chainSvc *BlockchainService) *BillingService {
	return &BillingService{
		store: store, chainSvc: chainSvc,
	}
}

func (bs *BillingService) Charge(ctx context.Context, req *ChargeRequest) (*ChargeReceipt, error) {
	resource, err := bs.chainSvc.GetAppCoinResourceWithId(req.AppCoinAddr, req.ResourceId)
	if err != nil {
		return nil, err
	}

	appCoinAccount, err := bs.chainSvc.GetAccountStatus(req.AppCoinAddr, req.CustomerAddr)
	if err != nil {
		return nil, err
	}

	// account frozen?
	if appCoinAccount.IsFrozen() {
		return nil, model.ErrAccountAddrFrozen
	}

	// insufficient balance?
	if int64(resource.Weight) > appCoinAccount.TotalBalance() {
		return nil, model.ErrInsufficentBalance
	}

	if req.DryRun { // for simulation only?
		return &ChargeReceipt{
			ResourceId: resource.ResourceId,
			Fee:        uint64(resource.Weight),
		}, nil
	}

	bill := model.Bill{
		Coin:      appCoinAccount.Coin,
		Address:   appCoinAccount.Address,
		Fee:       int64(resource.Weight),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	logger := logrus.WithFields(logrus.Fields{
		"statement": bill,
		"resource":  resource,
		"account":   appCoinAccount,
	})

	// TODO: need some benchmarking here in case of poor performance.
	if err := bs.store.Transaction(func(tx *gorm.DB) error {
		res := tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "coin"}, {Name: "address"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"fee":        gorm.Expr("fee + ?", resource.Weight),
				"updated_at": time.Now(),
			}),
		}).Create(&bill)

		if err := res.Error; err != nil {
			logger.WithError(err).Error("Failed to upsert bill")
			return errors.WithMessage(err, "failed to upsert bill")
		}

		// TODO add billing statements

		return bs.chainSvc.DeductAccountBalance(req.AppCoinAddr, req.CustomerAddr, int64(resource.Weight))
	}); err != nil {
		return nil, err
	}

	return &ChargeReceipt{
		ResourceId: resource.ResourceId,
		Fee:        uint64(resource.Weight),
	}, nil
}
