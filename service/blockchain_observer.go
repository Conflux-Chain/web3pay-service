package service

import (
	"math/big"
	"time"

	"github.com/Conflux-Chain/web3pay-service/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
)

func (bs *BlockchainService) addStatusConfirmTask(coin, addr common.Address) {
	bs.appCoinStatusConfirmQueue <- [2]common.Address{coin, addr}
}

// implements `ContractEventObserver` interface

func (bs *BlockchainService) StatusConfirmQueue() <-chan [2]common.Address {
	return bs.appCoinStatusConfirmQueue
}

func (bs *BlockchainService) OnConfirmStatus(coin, addr common.Address, balance *big.Int, frozen, block int64) error {
	_, err := bs.UpdateAccountStatus(coin, addr, balance, &frozen, &block)
	return err
}

func (bs *BlockchainService) OnAppCreated(event *contract.ControllerAPPCREATED) error {
	bs.workerPool.Submit(func() {
		for {
			coin := event.Addr
			baseCallOpt := &bind.CallOpts{
				BlockNumber: big.NewInt(int64(event.Raw.BlockNumber)),
			}

			resources, err := bs.provider.GetAppCoinResources(baseCallOpt, coin)
			if err != nil {
				logrus.WithField("event", event).
					WithError(err).
					Error("Failed to get APP coin resources on created")
				time.Sleep(time.Second)
				continue
			}

			bs.appCoinMutex.Lock()
			defer bs.appCoinMutex.Unlock()

			bs.appCoinBaseMap[coin] = AppCoinBase{
				Addr: coin, Owner: event.AppOwner, Resources: resources,
			}

			return
		}
	})

	return nil
}

func (bs *BlockchainService) OnMinted(event *contract.APPCoinMinted) error {
	if event.Amount.Cmp(big.NewInt(0)) == 0 {
		return nil
	}

	depositReq := &DepositRequest{
		Coin:        event.Raw.Address,
		Address:     event.To,
		Amount:      event.Amount,
		TxHash:      event.Raw.TxHash,
		BlockHash:   event.Raw.BlockHash,
		BlockNumber: event.Raw.BlockNumber,
		SubmitAt:    time.Now(),
	}

	return bs.Deposit(depositReq)
}

func (bs *BlockchainService) OnFrozen(event *contract.APPCoinFrozen) error {
	newStatus := int64(1) // TODO: extract from event? but not provided yet
	_, err := bs.UpdateAccountStatus(event.Raw.Address, event.Addr, nil, &newStatus, nil)
	return err
}

func (bs *BlockchainService) OnWithdraw(event *contract.APPCoinWithdraw) error {
	_, err := bs.DeleteAccountStatus(event.Raw.Address, event.Account)
	return err
}

func (bs *BlockchainService) OnResourceChanged(event *contract.APPCoinResourceChanged) error {
	bs.workerPool.Submit(func() {
		for {
			coin := event.Raw.Address
			baseCallOpt := &bind.CallOpts{
				BlockNumber: big.NewInt(int64(event.Raw.BlockNumber)),
			}

			resources, err := bs.provider.GetAppCoinResources(baseCallOpt, coin)
			if err != nil {
				logrus.WithField("event", event).
					WithError(err).
					Error("Failed to get APP coin resources on changed")
				time.Sleep(time.Second)
				continue
			}

			bs.appCoinMutex.Lock()
			defer bs.appCoinMutex.Unlock()

			if apiCoinBase, ok := bs.appCoinBaseMap[coin]; ok {
				apiCoinBase.Resources = resources
			}

			return
		}
	})

	return nil
}

func (bs *BlockchainService) OnReorgRevert(revertToBlock int64) error {
	for { // endless retry unless success
		err := bs.memStore.DeleteAccountsAfterBlock(revertToBlock)
		if err == nil {
			break
		}

		logrus.WithField("revertToBlock", revertToBlock).
			WithError(err).Error("Failed to reorg revert APP coin accounts")
		time.Sleep(time.Second)
	}

	return nil
}
