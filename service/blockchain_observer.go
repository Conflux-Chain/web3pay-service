package service

import (
	"math/big"
	"time"

	"github.com/Conflux-Chain/web3pay-service/contract"
	"github.com/Conflux-Chain/web3pay-service/util"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go/types"
	"github.com/sirupsen/logrus"
)

func (bs *BlockchainService) addStatusConfirmTask(app, addr common.Address) bool {
	logrus.WithFields(logrus.Fields{
		"app": app, "address": addr,
	}).Debug("Blockchain service adding APP account status confirmation task")

	if len(bs.appAccountStatusConfirmQueue) >= statusConfirmQueueSize {
		logrus.Warn("APP account status confirmation queue full")
		return false
	}

	bs.appAccountStatusConfirmQueue <- [2]common.Address{app, addr}
	return true
}

// implements `ContractEventObserver` interface

func (bs *BlockchainService) StatusConfirmQueue() <-chan [2]common.Address {
	return bs.appAccountStatusConfirmQueue
}

func (bs *BlockchainService) OnConfirmStatus(
	app, addr common.Address, balance *big.Int, frozen, block int64,
) error {
	account, err := bs.UpdateAccountStatus(app, addr, balance, &frozen, &block)

	logrus.WithFields(logrus.Fields{
		"app": app, "address": addr,
		"frozen": frozen, "block": block,
		"newBalance":    balance.Int64(),
		"updateAccount": account,
	}).WithError(err).Debug("Blockchain service `OnConfirmStatus` event handled")

	return err
}

func (bs *BlockchainService) OnAppCreated(event *contract.AppRegistryCreated, rawlog *types.Log) error {
	bs.workerPool.Submit(func() {
		for {
			blockNum := int64(rawlog.BlockNumber)
			baseCallOpt := &bind.CallOpts{
				BlockNumber: big.NewInt(blockNum),
			}

			logger := logrus.WithFields(logrus.Fields{
				"app":            event.App,
				"owner":          event.Owner,
				"operator":       event.Operator,
				"apiWeightToken": event.ApiWeightToken,
			})

			pendingSec, err := bs.provider.GetApiWeightTokenPendingSeconds(baseCallOpt, event.ApiWeightToken)
			if err != nil {
				logger.Error("Blockchain service failed to get pending seconds")
				time.Sleep(time.Second)
				continue
			}

			resources, err := bs.provider.GetApiWeightTokenResources(baseCallOpt, event.ApiWeightToken)
			if err != nil {
				logger.WithError(err).Error("Blockchain service failed to get new created APP resources")
				time.Sleep(time.Second)
				continue
			}

			bs.appMutex.Lock()
			defer bs.appMutex.Unlock()

			bs.appBaseMap[event.App] = AppBase{
				UpdateBlock:    blockNum,
				Addr:           event.App,
				ApiWeightToken: event.ApiWeightToken,
				PendingSeconds: pendingSec,
				Resources:      resources,
			}

			logrus.WithField("appBase", bs.appBaseMap[event.App]).
				Debug("Blockchain service `OnAppCreated` event handled")
			return
		}
	})

	return nil
}

func (bs *BlockchainService) OnDeposit(event *contract.AppDeposit, rawlog *types.Log) error {
	if event.TokenId.Cmp(big.NewInt(contract.TokenIdCoin)) != 0 && // vip coin
		event.TokenId.Cmp(big.NewInt(contract.TokenIdAirdrop)) != 0 { // airdrop
		return nil
	}

	return bs.deposit(&DepositRequest{
		App:         rawlog.Address,
		Address:     event.Receiver,
		Amount:      event.Amount,
		TxHash:      rawlog.TxHash,
		BlockHash:   rawlog.BlockHash,
		BlockNumber: int64(rawlog.BlockNumber),
		SubmitAt:    time.Now(),
	})
}

func (bs *BlockchainService) OnTransfer(event *contract.VipCoinTransferSingle, rawlog *types.Log) error {
	logger := logrus.WithFields(logrus.Fields{
		"vipCoin":        rawlog.Address,
		"accountAddress": event.From,
		"amount":         event.Value.Int64(),
		"blockNumber":    rawlog.BlockNumber,
		"operator":       event.Operator,
	})

	if event.Id.Cmp(big.NewInt(contract.TokenIdCoin)) != 0 && // vip coin
		event.Id.Cmp(big.NewInt(contract.TokenIdAirdrop)) != 0 { // airdrop
		logger.Debug("Blockchain service skipped `OnTransfer` event due to not a VIP or Airdrop coin")
		return nil
	}

	if !util.IsZeroAddress(event.To) { // only burnt event are concerned
		logger.Debug("Blockchain service skipped `OnTransfer` event due to no burnt event")
		return nil
	}

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

func (bs *BlockchainService) deposit(depositReq *DepositRequest) error {
	err := bs.DepositPending(depositReq)

	logrus.WithFields(logrus.Fields{
		"depositRequest": depositReq,
		"amount":         depositReq.Amount.Int64(),
	}).WithError(err).Debug("Blockchain service pending deposited balance for `OnDeposit` event")
	return err
}

func (bs *BlockchainService) OnFrozen(event *contract.AppFrozen, rawlog *types.Log) error {
	app := rawlog.Address
	addr := event.Account
	// TODO: extracted from event? but not provided by event body yet
	newStatus := int64(1)
	account, err := bs.UpdateAccountStatus(app, addr, nil, &newStatus, nil)

	logrus.WithFields(logrus.Fields{
		"app":           app,
		"address":       addr,
		"frozen":        newStatus,
		"updateAccount": account,
	}).WithError(err).Debug("Blockchain service `OnFrozen` event handled")

	return err
}

func (bs *BlockchainService) OnWithdraw(event *contract.AppWithdraw, rawlog *types.Log) error {
	app := rawlog.Address
	addr := event.Account
	_, err := bs.DeleteAccountStatus(rawlog.Address, event.Account)

	logrus.WithFields(logrus.Fields{
		"app":     app,
		"address": addr,
		"amount":  event.Amount.Int64(),
	}).WithError(err).Debug("Blockchain service `OnWithdraw` event handled")

	return err
}

func (bs *BlockchainService) OnResourceChanged(event *contract.ApiWeightTokenResourceChanged, rawlog *types.Log) error {
	bs.workerPool.Submit(func() {
		for {
			apiWeightToken := rawlog.Address
			blockNum := int64(rawlog.BlockNumber)
			baseCallOpt := &bind.CallOpts{BlockNumber: big.NewInt(blockNum)}
			logger := logrus.WithFields(logrus.Fields{
				"apiWeightToken": apiWeightToken,
				"resourceId":     event.Id,
				"resourceWeight": event.Weight.Int64(),
				"Operation":      event.Op,
			})

			resources, err := bs.provider.GetApiWeightTokenResources(baseCallOpt, apiWeightToken)
			if err != nil {
				logger.WithError(err).Error("Blockchain service failed to get config resources")
				time.Sleep(time.Second)
				continue
			}

			bs.appMutex.Lock()
			defer bs.appMutex.Unlock()

			var app common.Address
			for _app, ab := range bs.appBaseMap {
				if ab.ApiWeightToken == apiWeightToken {
					app = _app
					break
				}
			}

			appBase, ok := bs.appBaseMap[app]
			if !ok { // // APP not found?
				logger.Error("Blockchain service failed to update config resources for non-existed app")
				return
			}

			// in case of stale block update
			if appBase.UpdateBlock >= blockNum {
				return
			}

			appBase.Resources = resources
			appBase.UpdateBlock = blockNum
			bs.appBaseMap[app] = appBase

			logrus.WithField("appBase", appBase).
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
			WithError(err).Error("Blockchain service failed to reorg revert APP accounts")
		time.Sleep(time.Second)
	}

	logrus.WithField("revertToBlock", revertToBlock).
		Info("Blockchain service `OnReorgRevert` event handled")

	return nil
}
