package blockchain

import (
	"math/big"

	"github.com/Conflux-Chain/web3pay-service/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go/types"
)

type ContractEventObserver interface {
	// APP registry
	OnAppCreated(event *contract.AppRegistryCreated, rawlog *types.Log) error
	// App
	OnDeposit(event *contract.AppDeposit, rawlog *types.Log) error
	OnWithdraw(event *contract.AppWithdraw, rawlog *types.Log) error
	OnFrozen(event *contract.AppFrozen, rawlog *types.Log) error
	// ApiWeightToken
	OnResourceChanged(event *contract.ApiWeightTokenResourceChanged, rawlog *types.Log) error
	// VipCoin
	OnTransfer(event *contract.VipCoinTransferSingle, rawlog *types.Log) error
	// confirmation
	StatusConfirmQueue() <-chan [2]common.Address // [app, addr]
	OnConfirmStatus(app, addr common.Address, balance *big.Int, frozen, block int64) error
	// reorg revert
	OnReorgRevert(revertToBlock int64) error
}
