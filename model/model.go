package model

import (
	"math"
	"math/big"
	"time"

	"github.com/shopspring/decimal"
)

const (
	BillStatusCreated = iota
	BillStatusSubmitting
	BillStatusSubmitted
	BillStatusFailed
)

// Bill bills to settle on blockchain
type Bill struct {
	ID uint64
	// APP coin contract address
	Coin string `gorm:"size:64;not null;index:idx_coin_addr,priority:1"`
	// account address
	Address string `gorm:"size:64;not null;index:idx_coin_addr,priority:2"`
	// total deduction fee
	Fee decimal.Decimal `gorm:"size:128;type:string"`
	// 0 - created, 1 - submitting, 2 - submitted
	// record shall be deleted if settlement confirmed
	Status uint8 `gorm:"default:0"`
	// transaction hash if submitted
	TxnHash string `gorm:"size:64;type:string"`
	// memo
	Memo string `gorm:"size:256;type:string"`
	// create date
	CreatedAt time.Time
	// update date
	UpdatedAt time.Time
}

var All = []interface{}{
	&Bill{},
}

type AppCoinAccount struct {
	// APP coin contract address
	Coin string
	// account address
	Address string
	// frozen status, 0 means not frozen
	Frozen int64
	// deduction fee
	Fee decimal.Decimal
	// pending balance
	Balance decimal.Decimal
	// the confirmed block number, math.MaxInt64 means not confirmed
	ConfirmedBlock int64
}

func NewAppCoinAccount(coin, address string, frozen int64, balance *big.Int) *AppCoinAccount {
	account := AppCoinAccount{
		Coin:           coin,
		Address:        address,
		Frozen:         frozen,
		Balance:        decimal.NewFromBigInt(balance, 0),
		Fee:            decimal.NewFromInt(0),
		ConfirmedBlock: math.MaxInt64, // unconfirmed status
	}

	return &account
}

func (account *AppCoinAccount) TotalBalance() *big.Int {
	res := account.Balance.Sub(account.Fee)
	return res.BigInt()
}

func (account *AppCoinAccount) IsFrozen() bool {
	return account.Frozen > 0
}

func (account *AppCoinAccount) IsConfirmed() bool {
	return account.ConfirmedBlock != math.MaxInt64
}

func (account *AppCoinAccount) IncreaseFee(delta *big.Int) {
	account.Fee = account.Fee.Add(decimal.NewFromBigInt(delta, 0))
}

func (account *AppCoinAccount) DecreaseFee(delta *big.Int) bool {
	deltaD := decimal.NewFromBigInt(delta, 0)
	newFee, overflow := decreasePositiveDecimal(account.Fee, deltaD)
	account.Fee = newFee

	return overflow
}

func (account *AppCoinAccount) IncreaseBalance(delta *big.Int) {
	account.Balance = account.Balance.Add(decimal.NewFromBigInt(delta, 0))
}

func (account *AppCoinAccount) DecreaseBalance(delta *big.Int) bool {
	deltaD := decimal.NewFromBigInt(delta, 0)
	newBalance, overflow := decreasePositiveDecimal(account.Balance, deltaD)
	account.Balance = newBalance

	return overflow
}

func decreasePositiveDecimal(value, delta decimal.Decimal) (decimal.Decimal, bool) {
	if value.Cmp(delta) <= 0 {
		return decimal.Zero.Copy(), true
	}

	return value.Sub(delta), false
}
