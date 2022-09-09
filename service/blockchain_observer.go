package service

import (
	"math/big"
	"time"

	"github.com/Conflux-Chain/web3pay-service/contract"
	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/util"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go/types"
	"github.com/sirupsen/logrus"
)

func (bs *BlockchainService) addStatusConfirmTask(coin, addr common.Address) bool {
	logrus.WithFields(logrus.Fields{
		"appCoin": coin, "address": addr,
	}).Debug("Blockchain service adding APP coin status confirmation task")

	if len(bs.appCoinStatusConfirmQueue) >= statusConfirmQueueSize {
		logrus.Warn("APP coin status confirmation queue full")
		return false
	}

	bs.appCoinStatusConfirmQueue <- [2]common.Address{coin, addr}
	return true
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
			blockNum := int64(rawlog.BlockNumber)
			baseCallOpt := &bind.CallOpts{
				BlockNumber: big.NewInt(blockNum),
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
				UpdateBlock: blockNum, Addr: coin,
				Owner: event.AppOwner, Resources: resources,
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
	// only minted and burnt event for `FT_ID` token concerned only
	if event.Value.Cmp(big.NewInt(0)) == 0 || event.Id.Cmp(big.NewInt(contract.FT_ID)) != 0 {
		return nil
	}

	// for minted event, depositing balance
	if util.IsZeroAddress(event.From) {
		return bs.deposit(&DepositRequest{
			Coin:        rawlog.Address,
			Address:     event.To,
			Amount:      event.Value,
			TxHash:      rawlog.TxHash,
			BlockHash:   rawlog.BlockHash,
			BlockNumber: int64(rawlog.BlockNumber),
			SubmitAt:    time.Now(),
		})
	}

	// for burnt event, deducting account balance
	if util.IsZeroAddress(event.To) {
		logger := logrus.WithFields(logrus.Fields{
			"appCoinAddress": rawlog.Address,
			"accountAddress": event.From,
			"amount":         event.Value.Int64(),
			"blockNumber":    rawlog.BlockNumber,
			"operator":       event.Operator,
		})

		// skip transfer burnt event from transactions initiated by our operator
		if event.Operator == bs.provider.OperatorAddress() {
			logger.Debug("Blockchain service skipped transfer burnt event due to inner operator")
			return nil
		}

		decreased, err := bs.DecreaseAccountBalance(
			rawlog.Address, event.From, event.Value, int64(rawlog.BlockNumber),
		)

		logger.WithField("decreased", decreased).
			WithError(err).
			Debug("Blockchain service `OnTransfer` burnt event handled")
		return err
	}

	return nil
}

func (bs *BlockchainService) deposit(depositReq *DepositRequest) error {
	err := bs.DepositPending(depositReq)

	logrus.WithFields(logrus.Fields{
		"depositRequest": depositReq,
		"amount":         depositReq.Amount.Int64(),
	}).WithError(err).Debug("Blockchain service pending deposited balance for `OnTransfer` event")
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

func (bs *BlockchainService) OnAppOwnerChanged(event *contract.APPCoinAppOwnerChanged, rawlog *types.Log) error {
	bs.appCoinMutex.Lock()
	defer bs.appCoinMutex.Unlock()

	appCoin, ok := bs.appCoinBaseMap[rawlog.Address]
	if !ok {
		return model.ErrAppCoinNotFound
	}

	appCoin.Owner = event.To
	return nil
}

func (bs *BlockchainService) OnResourceChanged(event *contract.APPCoinResourceChanged, rawlog *types.Log) error {
	bs.workerPool.Submit(func() {
		for {
			coin := rawlog.Address
			blockNum := int64(rawlog.BlockNumber)
			baseCallOpt := &bind.CallOpts{BlockNumber: big.NewInt(blockNum)}

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

			// not our concerned APP coin
			appCoinBase, ok := bs.appCoinBaseMap[coin]
			if !ok {
				return
			}

			// in case of stale block update
			if appCoinBase.UpdateBlock >= blockNum {
				return
			}

			appCoinBase.Resources = resources
			appCoinBase.UpdateBlock = blockNum
			bs.appCoinBaseMap[coin] = appCoinBase

			logrus.WithField("appCoinBase", appCoinBase).
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
		Info("Blockchain service `OnReorgRevert` event handled")

	return nil
}
