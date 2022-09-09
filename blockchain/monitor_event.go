package blockchain

import (
	"github.com/Conflux-Chain/web3pay-service/contract"
	"github.com/Conflux-Chain/web3pay-service/util"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (m *Monitor) handleAirdropEvent(log *types.Log) (bool, error) {
	airdropAbi, err := contract.AirdropMetaData.GetAbi()
	if err != nil {
		return false, errors.WithMessage(err, "failed to get airdrop contract ABI")
	}

	// `Drop` event concerned only
	airdropDropEventId := airdropAbi.Events[contract.EventAirdropDrop].ID
	if log.Topics[0] != airdropDropEventId {
		return false, nil
	}

	eventAirdropDrop, err := contract.UnpackAirdropDrop(airdropAbi, log)
	if err != nil {
		return false, err
	}

	logger := logrus.WithFields(logrus.Fields{
		"airdropTo": eventAirdropDrop.To,
		"Amount":    eventAirdropDrop.Amount.Int64(),
		"Reason":    eventAirdropDrop.Reason,
	})

	if err := m.contractEventObserver.OnDrop(eventAirdropDrop, log); err != nil {
		logger.WithError(err).Info("Monitor failed to handle airdrop event")
		return false, err
	}

	logger.Debug("Monitor handled airdrop event")
	return true, nil
}

func (m *Monitor) handleAppCoinEvent(log *types.Log) (bool, error) {
	appCoinAbi, err := contract.APPCoinMetaData.GetAbi()
	if err != nil {
		return false, errors.WithMessage(err, "failed to get APP coin contract ABI")
	}

	switch log.Topics[0] {
	case appCoinAbi.Events[contract.EventAppCoinTransferSingle].ID:
		// transfer
		err = m.handleAppCoinTransferSingle(appCoinAbi, log)
	case appCoinAbi.Events[contract.EventAppCoinFrozen].ID:
		// frozen
		err = m.handleAppCoinFrozen(appCoinAbi, log)
	case appCoinAbi.Events[contract.EventAppCointWithdraw].ID:
		// withdraw
		err = m.handleAppCoinWithdraw(appCoinAbi, log)
	case appCoinAbi.Events[contract.EventAppCoinResourceChanged].ID:
		// resource changed
		err = m.handleAppCoinResourceChanged(appCoinAbi, log)
	case appCoinAbi.Events[contract.EventAppOwnerChanged].ID:
		// owner changed
		err = m.handleAppOwnerChanged(appCoinAbi, log)
	default: // maybe airdrop event?
		return m.handleAirdropEvent(log)
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (m *Monitor) handleAppCoinResourceChanged(appCoinAbi *abi.ABI, log *types.Log) error {
	eventAppCoinResrcChanged, err := contract.UnpackAppCoinResourceChanged(appCoinAbi, log)
	if err != nil {
		return err
	}

	logger := logrus.WithFields(logrus.Fields{
		"resourceId":     eventAppCoinResrcChanged.Id,
		"resourceWeight": eventAppCoinResrcChanged.Weight.Int64(),
		"Op":             eventAppCoinResrcChanged.Op,
	})

	if err := m.contractEventObserver.OnResourceChanged(eventAppCoinResrcChanged, log); err != nil {
		logger.WithError(err).Info("Monitor failed to handle APP coin resource changed event")
		return err
	}

	logger.Debug("Monitor handled APP coin resource changed event")
	return nil
}

func (m *Monitor) handleAppOwnerChanged(appCoinAbi *abi.ABI, log *types.Log) error {
	eventAppOwnerChanged, err := contract.UnpackAPPCoinAppOwnerChanged(appCoinAbi, log)
	if err != nil {
		return err
	}

	logger := logrus.WithFields(logrus.Fields{
		"appCoinContract": log.Address,
		"newOwner":        eventAppOwnerChanged.To,
	})

	if err := m.contractEventObserver.OnAppOwnerChanged(eventAppOwnerChanged, log); err != nil {
		logger.WithError(err).Info("Monitor failed to handle APP coin owner changed event")
		return err
	}

	logger.Debug("Monitor handled APP coin owner changed event")
	return nil
}

func (m *Monitor) handleAppCoinWithdraw(appCoinAbi *abi.ABI, log *types.Log) error {
	eventAppCoinWithdraw, err := contract.UnpackAppCoinWithdraw(appCoinAbi, log)
	if err != nil {
		return err
	}

	logger := logrus.WithFields(logrus.Fields{
		"withdrawAccount": eventAppCoinWithdraw.Account,
		"withdrawAmount":  eventAppCoinWithdraw.Amount.Int64(),
	})

	if err := m.contractEventObserver.OnWithdraw(eventAppCoinWithdraw, log); err != nil {
		logger.WithError(err).Info("Monitor failed to handle APP coin withdraw event")
		return err
	}

	logger.Debug("Monitor handled APP coin withdraw event")
	return nil
}

func (m *Monitor) handleAppCoinFrozen(appCoinAbi *abi.ABI, log *types.Log) error {
	eventAppCoinFrozen, err := contract.UnpackAppCoinFrozen(appCoinAbi, log)
	if err != nil {
		return err
	}

	logger := logrus.WithField("frozenAddress", eventAppCoinFrozen.Addr)

	if err := m.contractEventObserver.OnFrozen(eventAppCoinFrozen, log); err != nil {
		logger.WithError(err).Info("Monitor failed to handle APPCoin frozen event")
		return err
	}

	logger.Debug("Monitor handled APPCoin frozen event")
	return nil
}

func (m *Monitor) handleAppCoinTransferSingle(appCoinAbi *abi.ABI, log *types.Log) error {
	eventAppCoinTransfer, err := contract.UnpackAPPCoinTransferSingle(appCoinAbi, log)
	if err != nil {
		return err
	}

	if !util.IsZeroAddress(eventAppCoinTransfer.From) &&
		!util.IsZeroAddress(eventAppCoinTransfer.To) { // not a minted event or burnt event?
		return nil
	}

	logger := logrus.WithFields(logrus.Fields{
		"transferOperator": eventAppCoinTransfer.Operator,
		"transferFrom":     eventAppCoinTransfer.From,
		"transferTo":       eventAppCoinTransfer.To,
		"transferId":       eventAppCoinTransfer.Id,
		"transferValue":    eventAppCoinTransfer.Value.Int64(),
	})

	if err := m.contractEventObserver.OnTransfer(eventAppCoinTransfer, log); err != nil {
		logger.WithError(err).Info("Monitor failed to handle APP coin transfer event")
		return err
	}

	logger.Debug("Monitor handled APP coin transfer event")
	return nil
}

func (m *Monitor) handleControllerEvent(log *types.Log) (bool, error) {
	controllerAbi, err := contract.ControllerMetaData.GetAbi()
	if err != nil {
		return false, errors.WithMessage(err, "failed to get controller contract ABI")
	}

	// `APP_CREATED` event concerned only
	appCreatedEventId := controllerAbi.Events[contract.EventControllerAppCreated].ID
	if log.Topics[0] != appCreatedEventId {
		return false, nil
	}

	eventAppCreated, err := contract.UnpackControllerAPPCREATED(controllerAbi, log)
	if err != nil {
		return false, err
	}

	logger := logrus.WithFields(logrus.Fields{
		"AppCoin": eventAppCreated.Addr, "AppCoinOwner": eventAppCreated.AppOwner,
	})

	if m.FilterCreatorAddress != nil && *m.FilterCreatorAddress != eventAppCreated.AppOwner {
		// not an APP coin by concerned creator
		logger.Debug("Monitor skipped APPCreated event due to not a concerned APP coin creator")
		return false, nil
	}

	// add observing for new created APP coin
	for i := range m.AppCoinAddresses {
		if m.AppCoinAddresses[i] == eventAppCreated.Addr { // already exists?
			logger.Debug("Monitor skipped APPCreated event due to already existed")
			return false, nil
		}
	}

	m.AppCoinAddresses = append(m.AppCoinAddresses, eventAppCreated.Addr)

	if err := m.contractEventObserver.OnAppCreated(eventAppCreated, log); err != nil {
		logger.WithError(err).Info("Monitor failed to handle APPCreated event")
		return false, err
	}

	logger.Debug("Monitor handled APPCreated event")
	return true, nil
}

func (m *Monitor) reorgRevert(revertToBlock int64) error {
	logrus.WithField("revertToBlock", revertToBlock).Info("Monitor reorg reverted")

	// remove block hash of reverted block from cache window
	m.blockWindow.popn(revertToBlock)

	// reset syncer start block
	m.SyncFromBlockNumber = revertToBlock

	// notify observer for reorg revert
	return m.contractEventObserver.OnReorgRevert(revertToBlock)
}
