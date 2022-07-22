package model

import (
	"time"
)

type AppCoinAddrStatus struct {
	Fronzen uint64
	Balance uint64
}

// BillingStatement gross billing statements
type BillingStatement struct {
	ID       uint64
	Contract string `gorm:"size:64;not null;index:idx_contract_addr,priority:1"` // APP coin contract
	Address  string `gorm:"size:64;not null;index:idx_contract_addr,priority:2"` // customer address
	Fee      uint64 `gorm:"default:0"`                                           // total deduction fee
	Calls    uint64 `gorm:"default:0"`                                           // API call times
	// 0 - created, 1 - chain submmitting, 2 - chain submitted
	Status    uint8 `gorm:"default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

var All = []interface{}{
	&BillingStatement{},
}
