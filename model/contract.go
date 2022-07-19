package model

type Contract struct {
	ID uint64
	// contract address for the APP token
	Address string `gorm:"size:66;not null;index"`
}
