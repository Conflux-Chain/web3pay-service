package service

import (
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type DepositRequest struct {
	Coin        common.Address
	Address     common.Address
	Amount      *big.Int
	TxHash      common.Hash
	BlockHash   common.Hash
	BlockNumber uint64
	SubmitAt    time.Time
}

func (bs *BlockchainService) Deposit(request *DepositRequest) error {
	// TODO: add deposit logic here
	return errors.New("not supported yet")
}
