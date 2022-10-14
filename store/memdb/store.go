package memdb

import (
	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/store"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hashicorp/go-memdb"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.uber.org/multierr"
)

const (
	AppAccountDBName = "account"
)

var (
	dbSchema *memdb.DBSchema
)

func init() {
	tblSchemas := map[string]*memdb.TableSchema{
		AppAccountDBName: {
			Name: AppAccountDBName,
			Indexes: map[string]*memdb.IndexSchema{
				"id": {
					Name:   "id",
					Unique: true,
					Indexer: &memdb.CompoundIndex{
						Indexes: []memdb.Indexer{
							&memdb.StringFieldIndex{Field: "App"},
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

func (ms *MemStore) SaveAccountWithTxn(txn *memdb.Txn, account *model.AppAccount) error {
	return txn.Insert("account", account)
}

func (ms *MemStore) SaveAccount(account *model.AppAccount) error {
	txn := ms.Txn(true)

	if err := txn.Insert("account", account); err != nil {
		txn.Abort()
		return err
	}

	txn.Commit()
	return nil
}

func (ms *MemStore) DeleteAccount(app, address common.Address) (*model.AppAccount, bool, error) {
	account, ok, err := ms.GetAccount(app, address)
	if err != nil {
		return nil, false, errors.WithMessage(err, "failed to get APP account")
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

func (ms *MemStore) GetAccount(app, address common.Address) (*model.AppAccount, bool, error) {
	account, err := ms.Txn(false).First("account", "id", app.String(), address.String())
	if err != nil || account == nil {
		return nil, false, err
	}

	return account.(*model.AppAccount), true, nil
}

func (ms *MemStore) DeleteAccountsAfterBlock(blockNumber int64) error {
	txn := ms.Txn(false)

	// range scan over account status with `confirmedBlock` behind block
	it, err := txn.LowerBound(AppAccountDBName, "block", blockNumber)
	if err != nil {
		logrus.WithField("blockNumber", blockNumber).
			WithError(err).
			Error("MemDB failed to get APP accounts with lower bound of block number")
		return err
	}

	toDeleteAccounts := []*model.AppAccount{}
	for obj := it.Next(); obj != nil; obj = it.Next() {
		account := obj.(*model.AppAccount)
		toDeleteAccounts = append(toDeleteAccounts, account)
	}

	if len(toDeleteAccounts) == 0 {
		return nil
	}

	txn = ms.Txn(true)
	defer txn.Commit()

	var finalErr error
	for _, account := range toDeleteAccounts {
		err := txn.Delete(AppAccountDBName, account)
		if err != nil {
			logrus.WithField("account", *account).
				WithError(err).
				Error("MemDB failed to delete APP account")
			finalErr = multierr.Combine(finalErr, err)

			continue
		}

		logrus.WithField("account", *account).Debug("MemDB deleted APP account")
	}

	return finalErr
}
