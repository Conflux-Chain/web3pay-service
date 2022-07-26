package service

import (
	"time"

	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/store/sqlite"
	"github.com/Conflux-Chain/web3pay-service/util"
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
	ContractAddr common.Address
	CustomerAddr common.Address
}

type ChargeReceipt struct {
	ResourceId string
	Fee        uint64
	Balance    uint64
}

type BillingService struct {
	accountSvc *AccountService
	chainSvc   *BlockchainService
	store      *sqlite.SqliteStore
	kmutex     *util.KMutex
}

func NewBillingService(store *sqlite.SqliteStore, accountSvc *AccountService, chainSvc *BlockchainService) *BillingService {
	return &BillingService{
		accountSvc: accountSvc,
		chainSvc:   chainSvc,
		store:      store,
		kmutex:     util.NewKMutex(),
	}
}

func (bs *BillingService) Charge(ctx context.Context, req *ChargeRequest) (*ChargeReceipt, error) {
	resource, err := bs.chainSvc.GetAppCoinResourceWithId(req.ContractAddr, req.ResourceId)
	if err != nil {
		return nil, err
	}

	appCoinStatus, err := bs.accountSvc.GetStatus(req.ContractAddr, req.CustomerAddr)
	if err != nil {
		return nil, err
	}

	// account frozen?
	if appCoinStatus.Frozen > 0 {
		return nil, model.ErrAccountAddrFrozen
	}

	// insufficient balance?
	if resource.Weight > uint32(appCoinStatus.Balance) {
		return nil, model.ErrInsufficentBalance
	}

	if req.DryRun { // for simulation only?
		return &ChargeReceipt{
			ResourceId: resource.ResourceId,
			Fee:        uint64(resource.Weight),
			Balance:    appCoinStatus.Balance - uint64(resource.Weight),
		}, nil
	}

	var newBalance uint64
	statement := model.BillingStatement{
		Contract:  req.ContractAddr.String(),
		Address:   req.CustomerAddr.String(),
		Fee:       uint64(resource.Weight),
		Calls:     1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	logger := logrus.WithFields(logrus.Fields{
		"statement": statement, "resource": resource, "status": appCoinStatus,
	})

	if err := bs.store.Transaction(func(tx *gorm.DB) error {
		if err := tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "contract"}, {Name: "address"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"fee":        gorm.Expr("fee + ?", resource.Weight),
				"calls":      gorm.Expr("calls + 1"),
				"updated_at": time.Now(),
			}),
		}).Create(&statement).Error; err != nil {
			logger.WithError(err).Info("Failed to upsert billing statement")
			return errors.WithMessage(err, "failed to upsert billing statement")
		}

		newBalance, err := bs.accountSvc.DeductBalance(req.ContractAddr, req.CustomerAddr, uint64(resource.Weight))
		if err != nil {
			logger.WithField("newBalance", newBalance).Info("Failed to deduct account balance")
		}
		return err
	}); err != nil {
		return nil, err
	}

	return &ChargeReceipt{
		ResourceId: resource.ResourceId,
		Fee:        uint64(resource.Weight),
		Balance:    newBalance,
	}, nil
}
