package service

import (
	"time"

	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/store/sqlite"
	"github.com/Conflux-Chain/web3pay-service/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	errInsufficentBalance = errors.New("insufficient balance")
	errAccountAddrFrozen  = errors.New("account address fronzen")
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
	chainSvc *BlockchainService
	store    *sqlite.SqliteStore
	kmutex   *util.KMutex
}

func NewBillingService(store *sqlite.SqliteStore, chainSvc *BlockchainService) *BillingService {
	return &BillingService{
		chainSvc: chainSvc,
		store:    store,
		kmutex:   util.NewKMutex(),
	}
}

func (bs *BillingService) Charge(ctx context.Context, req *ChargeRequest) (*ChargeReceipt, error) {
	if len(req.ResourceId) == 0 { // if resourceId is empty, use default resource
		req.ResourceId = "default"
	}

	resource, err := bs.chainSvc.GetAppCoinResourceWithId(req.ContractAddr, req.ResourceId)
	if err != nil {
		return nil, err
	}

	appCoinStatus, err := bs.chainSvc.GetAppCoinStatusOfAddr(req.ContractAddr, req.CustomerAddr)
	if err != nil {
		return nil, err
	}

	if appCoinStatus.Frozen > 0 {
		return nil, errAccountAddrFrozen
	}

	if resource.Weight > uint32(appCoinStatus.Balance) {
		return nil, errInsufficentBalance
	}

	if req.DryRun { // for simulation only?
		return &ChargeReceipt{
			ResourceId: req.ResourceId,
			Fee:        uint64(resource.Weight),
			Balance:    appCoinStatus.Balance - uint64(resource.Weight),
		}, nil
	}

	var newBalance uint64
	if err := bs.store.Transaction(func(tx *gorm.DB) error {
		statement := model.BillingStatement{
			Contract:  req.ContractAddr.String(),
			Address:   req.CustomerAddr.String(),
			Fee:       uint64(resource.Weight),
			Calls:     1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "contract"}, {Name: "address"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"fee":        gorm.Expr("fee + ?", resource.Weight),
				"calls":      gorm.Expr("calls + 1"),
				"updated_at": time.Now(),
			}),
		}).Create(&statement).Error; err != nil {
			return errors.WithMessage(err, "failed to upsert billing statement")
		}

		newBalance, err = bs.chainSvc.DeductAppCoinBalance(req.ContractAddr, req.CustomerAddr, uint64(resource.Weight))
		return err
	}); err != nil {
		return nil, err
	}

	return &ChargeReceipt{
		ResourceId: req.ResourceId,
		Fee:        uint64(resource.Weight),
		Balance:    newBalance,
	}, nil
}
