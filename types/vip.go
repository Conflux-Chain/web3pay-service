package types

import (
	"github.com/Conflux-Chain/web3pay-service/contract"
	"github.com/ethereum/go-ethereum/common"
)

// VipInfo VIP subscription information
type VipInfo struct {
	contract.ICardTrackerVipInfo // VIP card info

	Account common.Address // VIP account address
}
