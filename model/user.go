package model

type User struct {
	ID uint64
	// contract ID
	Cid uint64 `gorm:"not null;index"`
	// EIP55-compliant hex string representation of the address
	Address string `gorm:"size:66;not null;index"`
	// whether the user is frozen
	Fronze bool
	// balance of APP token
	Balance uint64
}
