package blockchain

import (
	"math/big"

	"github.com/Conflux-Chain/web3pay-service/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go/types"
)

type ContractEventObserver interface {
	// controller
	OnAppCreated(event *contract.ControllerAPPCREATED, rawlog *types.Log) error
	// app coin
	OnTransfer(event *contract.APPCoinTransferSingle, rawlog *types.Log) error
	OnFrozen(event *contract.APPCoinFrozen, rawlog *types.Log) error
	OnWithdraw(event *contract.APPCoinWithdraw, rawlog *types.Log) error
	OnResourceChanged(event *contract.APPCoinResourceChanged, rawlog *types.Log) error
	StatusConfirmQueue() <-chan [2]common.Address // [coin, addr]
	OnConfirmStatus(coin, addr common.Address, balance *big.Int, frozen, block int64) error
	// airdrop
	OnDrop(event *contract.AirdropDrop, rawlog *types.Log) error
	// reorg revert
	OnReorgRevert(revertToBlock int64) error
}
