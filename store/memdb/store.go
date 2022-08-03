package memdb

import (
	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/store"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hashicorp/go-memdb"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	AppCoinAccountDBName = "account"
)

var (
	dbSchema *memdb.DBSchema
)

func init() {
	tblSchemas := map[string]*memdb.TableSchema{
		AppCoinAccountDBName: {
			Name: AppCoinAccountDBName,
			Indexes: map[string]*memdb.IndexSchema{
				"id": {
					Name:   "id",
					Unique: true,
					Indexer: &memdb.CompoundIndex{
						Indexes: []memdb.Indexer{
							&memdb.StringFieldIndex{Field: "Coin"},
							&memdb.StringFieldIndex{Field: "Address"},
						},
					},
				},
				"block": {
					Name:    "block",
					Unique:  false,
					Indexer: &memdb.IntFieldIndex{Field: "ConfirmedBlock"},
				},
			},
		},
	}

	dbSchema = &memdb.DBSchema{Tables: tblSchemas}
}

type MemStore struct {
	*memdb.MemDB
}

func MustNewStoreFromViper() *MemStore {
	db, err := memdb.NewMemDB(dbSchema)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to new memory store")
	}

	return &MemStore{MemDB: db}
}

func (ms *MemStore) IsRecordNotFound(err error) bool {
	return errors.Is(err, memdb.ErrNotFound) || err == store.ErrNotFound
}

func (ms *MemStore) Close() error {
	// nothing to do
	return nil
}

func (ms *MemStore) SaveAccount(account *model.AppCoinAccount) error {
	txn := ms.Txn(true)

	if err := txn.Insert("account", account); err != nil {
		txn.Abort()
		return err
	}

	txn.Commit()
	return nil
}

func (ms *MemStore) DeleteAccount(appCoin, address common.Address) (*model.AppCoinAccount, bool, error) {
	account, ok, err := ms.GetAccount(appCoin, address)
	if err != nil {
		return nil, false, errors.WithMessage(err, "failed to get APP coin account")
	}

	if !ok {
		return nil, false, nil
	}

	tx := ms.Txn(true)
	if err := tx.Delete("account", account); err != nil {
		tx.Abort()
		return nil, false, err
	}

	tx.Commit()
	return account, true, nil
}

func (ms *MemStore) GetAccount(appCoin, address common.Address) (*model.AppCoinAccount, bool, error) {
	account, err := ms.Txn(false).First("account", "id", appCoin.String(), address.String())
	if ms.IsRecordNotFound(err) {
		return nil, false, nil
	}

	if err == nil {
		return account.(*model.AppCoinAccount), true, nil
	}

	return nil, false, err
}

func (ms *MemStore) DeleteAccountsAfterBlock(blockNumber int64) error {
	txn := ms.Txn(false)

	// range scan over account status with `confirmedBlock` behind block
	it, err := txn.LowerBound(AppCoinAccountDBName, "block", blockNumber)
	if err != nil {
		logrus.WithField("blockNumber", blockNumber).
			WithError(err).Error("Failed to get APP coin accounts with lower bound")
		return err
	}

	toDeleteAccounts := []*model.AppCoinAccount{}
	for obj := it.Next(); obj != nil; obj = it.Next() {
		account := obj.(*model.AppCoinAccount)
		toDeleteAccounts = append(toDeleteAccounts, account)
	}

	if len(toDeleteAccounts) == 0 {
		return nil
	}

	txn = ms.Txn(true)
	defer txn.Commit()

	for i := range toDeleteAccounts {
		txn.Delete(AppCoinAccountDBName, toDeleteAccounts[i])
	}

	logrus.WithFields(logrus.Fields{
		"blockNumber":      blockNumber,
		"toDeleteAccounts": toDeleteAccounts,
	}).Debug("MemDB deleted accounts with confirmed block behind")

	return nil
}
