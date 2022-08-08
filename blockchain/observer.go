package blockchain

import (
	"math/big"

	"github.com/Conflux-Chain/web3pay-service/contract"
	"github.com/ethereum/go-ethereum/common"
)

type ContractEventObserver interface {
	// controller
	OnAppCreated(event *contract.ControllerAPPCREATED) error
	// app coin
	OnTransfer(event *contract.APPCoinTransferSingle) error
	OnFrozen(event *contract.APPCoinFrozen) error
	OnWithdraw(event *contract.APPCoinWithdraw) error
	OnResourceChanged(event *contract.APPCoinResourceChanged) error
	StatusConfirmQueue() <-chan [2]common.Address // [coin, addr]
	OnConfirmStatus(coin, addr common.Address, balance *big.Int, frozen, block int64) error
	// airdrop
	OnDrop(event *contract.AirdropDrop) error
	// reorg revert
	OnReorgRevert(revertToBlock int64) error
}
