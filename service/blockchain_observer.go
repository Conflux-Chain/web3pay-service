package service

import (
	"math/big"
	"time"

	"github.com/Conflux-Chain/web3pay-service/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go/types"
	"github.com/sirupsen/logrus"
)

func (bs *BlockchainService) addStatusConfirmTask(coin, addr common.Address) {
	logrus.WithFields(logrus.Fields{
		"appCoin": coin, "address": addr,
	}).Debug("Blockchain service adding APP coin status confirmation task")

	bs.appCoinStatusConfirmQueue <- [2]common.Address{coin, addr}
}

// implements `ContractEventObserver` interface

func (bs *BlockchainService) StatusConfirmQueue() <-chan [2]common.Address {
	return bs.appCoinStatusConfirmQueue
}

func (bs *BlockchainService) OnConfirmStatus(
	coin, addr common.Address, balance *big.Int, frozen, block int64,
) error {
	account, err := bs.UpdateAccountStatus(coin, addr, balance, &frozen, &block)

	logrus.WithFields(logrus.Fields{
		"appCoin": coin, "address": addr,
		"frozen": frozen, "block": block,
		"newBalance":    balance.Int64(),
		"updateAccount": account,
	}).WithError(err).Debug("Blockchain service `OnConfirmStatus` event handled")

	return err
}

func (bs *BlockchainService) OnAppCreated(event *contract.ControllerAPPCREATED, rawlog *types.Log) error {
	bs.workerPool.Submit(func() {
		for {
			coin := event.Addr
			baseCallOpt := &bind.CallOpts{
				BlockNumber: big.NewInt(int64(rawlog.BlockNumber)),
			}

			resources, err := bs.provider.GetAppCoinResources(baseCallOpt, coin)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"AppCoin": event.Addr, "AppCoinOwner": event.AppOwner,
				}).WithError(err).Error("Blockchain service failed to get new created APP coin resources")
				time.Sleep(time.Second)
				continue
			}

			bs.appCoinMutex.Lock()
			defer bs.appCoinMutex.Unlock()

			bs.appCoinBaseMap[coin] = AppCoinBase{
				Addr: coin, Owner: event.AppOwner, Resources: resources,
			}

			logrus.WithField("appCoinBase", bs.appCoinBaseMap[coin]).
				Debug("Blockchain service `OnAppCreated` event handled")
			return
		}
	})

	return nil
}

func (bs *BlockchainService) OnDrop(event *contract.AirdropDrop, rawlog *types.Log) error {
	if event.Amount.Cmp(big.NewInt(0)) == 0 {
		return nil
	}

	depositReq := &DepositRequest{
		Coin:        rawlog.Address,
		Address:     event.To,
		Amount:      event.Amount,
		TxHash:      rawlog.TxHash,
		BlockHash:   rawlog.BlockHash,
		BlockNumber: int64(rawlog.BlockNumber),
		SubmitAt:    time.Now(),
	}
	err := bs.DepositPending(depositReq)

	logrus.WithFields(logrus.Fields{
		"depositRequest": depositReq,
		"amount":         depositReq.Amount.Int64(),
	}).WithError(err).Debug("Blockchain service `OnDrop` event handled")

	return err
}

func (bs *BlockchainService) OnTransfer(event *contract.APPCoinTransferSingle, rawlog *types.Log) error {
	if event.Value.Cmp(big.NewInt(0)) == 0 {
		return nil
	}

	depositReq := &DepositRequest{
		Coin:        rawlog.Address,
		Address:     event.To,
		Amount:      event.Value,
		TxHash:      rawlog.TxHash,
		BlockHash:   rawlog.BlockHash,
		BlockNumber: int64(rawlog.BlockNumber),
		SubmitAt:    time.Now(),
	}
	err := bs.DepositPending(depositReq)

	logrus.WithFields(logrus.Fields{
		"depositRequest": depositReq,
		"amount":         depositReq.Amount.Int64(),
	}).WithError(err).Debug("Blockchain service `OnTransfer` event handled")

	return err
}

func (bs *BlockchainService) OnFrozen(event *contract.APPCoinFrozen, rawlog *types.Log) error {
	coin := rawlog.Address
	addr := event.Addr
	// TODO: extracted from event? but not provided yet
	newStatus := int64(1)
	account, err := bs.UpdateAccountStatus(coin, addr, nil, &newStatus, nil)

	logrus.WithFields(logrus.Fields{
		"appCoin":       coin,
		"address":       addr,
		"frozen":        newStatus,
		"updateAccount": account,
	}).WithError(err).Debug("Blockchain service `OnFrozen` event handled")

	return err
}

func (bs *BlockchainService) OnWithdraw(event *contract.APPCoinWithdraw, rawlog *types.Log) error {
	coin := rawlog.Address
	addr := event.Account
	_, err := bs.DeleteAccountStatus(rawlog.Address, event.Account)

	logrus.WithFields(logrus.Fields{
		"appCoin": coin,
		"address": addr,
		"amount":  event.Amount.Int64(),
	}).WithError(err).Debug("Blockchain service `OnWithdraw` event handled")

	return err
}

func (bs *BlockchainService) OnResourceChanged(event *contract.APPCoinResourceChanged, rawlog *types.Log) error {
	bs.workerPool.Submit(func() {
		for {
			coin := rawlog.Address
			baseCallOpt := &bind.CallOpts{
				BlockNumber: big.NewInt(int64(rawlog.BlockNumber)),
			}

			resources, err := bs.provider.GetAppCoinResources(baseCallOpt, coin)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"resourceId":     event.Id,
					"resourceWeight": event.Weight.Int64(),
					"Op":             event.Op,
				}).WithError(err).Error("Blockchain service failed to get changed APP coin resources")
				time.Sleep(time.Second)
				continue
			}

			bs.appCoinMutex.Lock()
			defer bs.appCoinMutex.Unlock()

			if apiCoinBase, ok := bs.appCoinBaseMap[coin]; ok {
				apiCoinBase.Resources = resources
			}

			logrus.WithField("appCoinBase", bs.appCoinBaseMap[coin]).
				Debug("Blockchain service `OnResourceChanged` event handled")
			return
		}
	})

	return nil
}

func (bs *BlockchainService) OnReorgRevert(revertToBlock int64) error {
	for { // endless retry until success
		err := bs.memStore.DeleteAccountsAfterBlock(revertToBlock)
		if err == nil {
			break
		}

		logrus.WithField("revertToBlock", revertToBlock).
			WithError(err).Error("Blockchain service failed to reorg revert APP coin accounts")
		time.Sleep(time.Second)
	}

	logrus.WithField("revertToBlock", revertToBlock).
		Debug("Blockchain service `OnReorgRevert` event handled")

	return nil
}
