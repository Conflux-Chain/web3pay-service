package model

import (
	"time"
)

// BillingStatement gross billing statements
type BillingStatement struct {
	ID       uint64
	Address  string `gorm:"size:64;not null;uniqueIndex:idx_contract_addr,priority:1"` // customer address
	Contract string `gorm:"size:64;not null;uniqueIndex:idx_contract_addr,priority:2"` // APP coin contract
	Fee      uint64 `gorm:"default:0"`                                                 // total deduction fee
	Calls    uint64 `gorm:"default:0"`                                                 // API call times
	// 0 - created, 1 - chain submmitting, 2 - chain submitted
	Status    uint8 `gorm:"default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

var All = []interface{}{
	&BillingStatement{},
}
