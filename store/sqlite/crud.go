package sqlite

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Crud struct {
	store                  *SqliteStore
	errEntityNotFound      error
	errEntityAlreadyExists error
}

func NewCrud(store *SqliteStore, errEntityNotFound, errEntityAlreadyExists error) *Crud {
	return &Crud{store, errEntityNotFound, errEntityAlreadyExists}
}

func (crud *Crud) Exists(modelPtr interface{}, whereQuery string, args ...interface{}) (bool, error) {
	err := crud.store.Where(whereQuery, args...).First(modelPtr).Error
	if err == nil {
		return true, nil
	}

	if crud.store.IsRecordNotFound(err) {
		return false, nil
	}

	return false, err
}

func (crud *Crud) Get(modelPtr interface{}, whereQuery string, args ...interface{}) error {
	exists, err := crud.Exists(modelPtr, whereQuery, args...)
	if err != nil {
		return err
	}

	if !exists {
		return crud.errEntityNotFound
	}

	return nil
}

func (crud *Crud) GetById(modelPtr interface{}, id uint64) error {
	return crud.Get(modelPtr, "id = ?", id)
}

func (crud *Crud) RequireAbsent(modelPtr interface{}, whereQuery string, args ...interface{}) error {
	exists, err := crud.Exists(modelPtr, whereQuery, args...)
	if err != nil {
		return err
	}

	if exists {
		return crud.errEntityAlreadyExists
	}

	return nil
}

func (crud *Crud) List(db *gorm.DB, idDESC bool, offset, limit int, destSlice interface{}) (total int64, err error) {
	var orderBy string

	if idDESC {
		orderBy = "id DESC"
	} else {
		orderBy = "id ASC"
	}

	return crud.ListByOrder(db, orderBy, offset, limit, destSlice)
}

func (*Crud) ListByOrder(db *gorm.DB, orderBy string, offset, limit int, destSlice interface{}) (total int64, err error) {
	if err = db.Count(&total).Error; err != nil {
		return 0, err
	}

	if !logrus.IsLevelEnabled(logrus.TraceLevel) && (total == 0 || total <= int64(offset)) {
		return total, nil
	}

	if err = db.Order(orderBy).Offset(offset).Limit(limit).Find(destSlice).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func List(
	db *gorm.DB, idDESC bool, offset, limit int, destSlice interface{},
) (total int64, err error) {
	var orderBy string

	if idDESC {
		orderBy = "id DESC"
	} else {
		orderBy = "id ASC"
	}

	return ListByOrder(db, orderBy, offset, limit, destSlice)
}

func ListByOrder(
	db *gorm.DB, orderBy string, offset, limit int, destSlice interface{},
) (total int64, err error) {
	if err = db.Count(&total).Error; err != nil {
		return 0, err
	}

	if err = db.Order(orderBy).Offset(offset).Limit(limit).Find(destSlice).Error; err != nil {
		return 0, err
	}

	return total, nil
}
