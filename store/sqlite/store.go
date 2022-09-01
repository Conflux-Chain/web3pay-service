package sqlite

import (
	"math/big"
	"time"

	"github.com/Conflux-Chain/web3pay-service/metrics"
	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/store"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type SqliteStore struct {
	*gorm.DB
}

func MustNewStoreFromViper(tables ...interface{}) *SqliteStore {
	config := MustNewConfigFromViper()
	return config.MustOpenOrCreate(tables...)
}

func NewSqliteStore(db *gorm.DB) *SqliteStore {
	return &SqliteStore{DB: db}
}

func (ms *SqliteStore) IsRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound) || err == store.ErrNotFound
}

func (ms *SqliteStore) Close() error {
	if db, err := ms.DB.DB(); err != nil {
		return err
	} else {
		return db.Close()
	}
}

func (ms *SqliteStore) GetBillFee(coin, addr string) (*big.Int, error) {
	var bills []*model.Bill

	if err := ms.Find(&bills, "coin = ? AND address = ?", coin, addr).Error; err != nil {
		return nil, err
	}

	var totalFees decimal.Decimal
	for i := range bills {
		totalFees = totalFees.Add(bills[i].Fee)
	}

	return totalFees.BigInt(), nil
}

func (ms *SqliteStore) GetBill(coin, addr string, status int) (*model.Bill, bool, error) {
	bill := &model.Bill{}

	err := ms.First(bill, "coin = ? AND address = ? AND status = ?", coin, addr, status).Error
	if ms.IsRecordNotFound(err) {
		return nil, false, nil
	}

	if err != nil {
		return nil, false, err
	}

	return bill, true, nil
}

func (ms *SqliteStore) UpsertBill(tx *gorm.DB, coin, addr string, fee *big.Int) (*model.Bill, error) {
	start := time.Now()

	bill, existed, err := ms.GetBill(coin, addr, model.BillStatusCreated)
	if err != nil {
		metrics.Store.UpsertBillQps(err).UpdateSince(start)
		return nil, errors.WithMessage(err, "failed to load bill")
	}

	if !existed { // insert
		bill = &model.Bill{
			Coin:      coin,
			Address:   addr,
			Fee:       decimal.NewFromBigInt(fee, 0),
			Status:    model.BillStatusCreated,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := tx.Create(bill).Error; err != nil {
			metrics.Store.UpsertBillQps(err).UpdateSince(start)
			return nil, errors.WithMessage(err, "failed to create bill")
		}

		return bill, nil
	}

	// update
	bill.Fee = bill.Fee.Add(decimal.NewFromBigInt(fee, 0))
	if err := tx.Model(&model.Bill{}).Where("id = ?", bill.ID).Update("fee", bill.Fee).Error; err != nil {
		metrics.Store.UpsertBillQps(err).UpdateSince(start)
		return nil, errors.WithMessage(err, "failed to update bill")
	}

	metrics.Store.UpsertBillQps(nil).UpdateSince(start)
	return bill, nil
}
