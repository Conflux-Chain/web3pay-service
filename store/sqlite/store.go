package sqlite

import (
	"github.com/Conflux-Chain/web3pay-service/store"
	"github.com/pkg/errors"
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
