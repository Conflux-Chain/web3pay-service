package model

import (
	"time"

	"github.com/ethereum/go-ethereum/common/math"
)

// Bill bills to settle on blockchain
type Bill struct {
	ID      uint64
	Coin    string `gorm:"size:64;not null;uniqueIndex:idx_coin_addr,priority:1"` // APP coin contract address
	Address string `gorm:"size:64;not null;uniqueIndex:idx_coin_addr,priority:2"` // account address
	Fee     int64  `gorm:"default:0"`                                             // total deduction fee

	Status uint8 `gorm:"default:0"` // 0 - created, 1 - submmitting, 2 - submitted

	CreatedAt time.Time // create date
	UpdatedAt time.Time // update date
}

var All = []interface{}{
	&Bill{},
}

type AppCoinAccount struct {
	Coin           string // APP coin contract address
	Address        string // account address
	Frozen         int64  // frozen status
	Fee            int64  // deduction fee
	Balance        int64  // pending balance
	ConfirmedBlock int64  // the confirmed block number, math.MaxInt64 means not confirmed
}

func NewAppCoinAccount(coin, address string, frozen, balance int64) *AppCoinAccount {
	account := AppCoinAccount{
		Coin:           coin,
		Address:        address,
		Frozen:         frozen,
		Balance:        balance,
		ConfirmedBlock: math.MaxInt64, // unconfirmed status
	}

	return &account
}

func (account *AppCoinAccount) TotalBalance() int64 {
	return account.Balance - account.Fee
}

func (account *AppCoinAccount) IsFrozen() bool {
	return account.Frozen > 0
}

func (account *AppCoinAccount) IsConfirmed() bool {
	return account.ConfirmedBlock != math.MaxInt64
}

func (account *AppCoinAccount) IncreaseFee(addFee int64) int64 {
	account.Fee += addFee
	return account.Fee
}
