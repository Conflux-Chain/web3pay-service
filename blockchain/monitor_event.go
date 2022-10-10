package blockchain

import (
	"github.com/Conflux-Chain/web3pay-service/contract"
	"github.com/Conflux-Chain/web3pay-service/util"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (m *Monitor) handleVipCoinEvent(log *types.Log) (bool, error) {
	vcAbi, err := contract.VipCoinMetaData.GetAbi()
	if err != nil {
		return false, errors.WithMessage(err, "failed to get `VipCoin` contract ABI")
	}

	// `TransferSingle` event concerned only
	tsEventId := vcAbi.Events[contract.EventVipCoinTransferSingle].ID
	if log.Topics[0] != tsEventId {
		return false, nil
	}

	if err := m.handleVipCoinTransferSingle(vcAbi, log); err != nil {
		return false, err
	}

	return true, nil
}

func (m *Monitor) handleVipCoinTransferSingle(vcAbi *abi.ABI, log *types.Log) error {
	eventTransfer, err := contract.UnpackVipCoinTransferSingle(vcAbi, log)
	if err != nil {
		return err
	}

	logger := logrus.WithFields(logrus.Fields{
		"from":  eventTransfer.From,
		"to":    eventTransfer.To,
		"id":    eventTransfer.Id.Int64(),
		"value": eventTransfer.Value.Int64(),
	})

	if err := m.contractEventObserver.OnTransfer(eventTransfer, log); err != nil {
		logger.WithError(err).Info("Monitor failed to handle VipCoin transfer event")
		return err
	}

	logger.Debug("Monitor handled VipCoin transfer event")
	return nil
}

func (m *Monitor) handleApiWeightTokenEvent(log *types.Log) (bool, error) {
	awtAbi, err := contract.ApiWeightTokenMetaData.GetAbi()
	if err != nil {
		return false, errors.WithMessage(err, "failed to get `ApiWeightToken` contract ABI")
	}

	// `ResourceChanged` event concerned only
	resrcChangedEventId := awtAbi.Events[contract.EventApiTokenWeightResourceChanged].ID
	if log.Topics[0] != resrcChangedEventId {
		return false, nil
	}

	if err := m.handleApiTokenWeightResourceChanged(awtAbi, log); err != nil {
		return false, err
	}

	return true, nil
}

func (m *Monitor) handleApiTokenWeightResourceChanged(awtAbi *abi.ABI, log *types.Log) error {
	eventResrcChanged, err := contract.UnpackApiWeightTokenResourceChanged(awtAbi, log)
	if err != nil {
		return err
	}

	logger := logrus.WithFields(logrus.Fields{
		"resourceId":     eventResrcChanged.Id,
		"resourceWeight": eventResrcChanged.Weight.Int64(),
		"Op":             eventResrcChanged.Op,
	})

	if err := m.contractEventObserver.OnResourceChanged(eventResrcChanged, log); err != nil {
		logger.WithError(err).Info("Monitor failed to handle APP config resource changed event")
		return err
	}

	logger.Debug("Monitor handled APP config resource changed event")
	return nil
}

func (m *Monitor) handleAppEvent(log *types.Log) (bool, error) {
	appAbi, err := contract.AppMetaData.GetAbi()
	if err != nil {
		return false, errors.WithMessage(err, "failed to get App contract ABI")
	}

	switch log.Topics[0] {
	case appAbi.Events[contract.EventAppDeposit].ID:
		// deposit
		err = m.handleAppDeposit(appAbi, log)
	case appAbi.Events[contract.EventAppWithdraw].ID:
		// withdraw
		err = m.handleAppWithdraw(appAbi, log)
	case appAbi.Events[contract.EventAppFrozen].ID:
		// frozen
		err = m.handleAppFrozen(appAbi, log)
	default:
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (m *Monitor) handleAppFrozen(appAbi *abi.ABI, log *types.Log) error {
	eventAppFrozen, err := contract.UnpackAppFrozen(appAbi, log)
	if err != nil {
		return err
	}

	logger := logrus.WithField("frozenAccount", eventAppFrozen.Account)

	if err := m.contractEventObserver.OnFrozen(eventAppFrozen, log); err != nil {
		logger.WithError(err).Info("Monitor failed to handle APP frozen event")
		return err
	}

	logger.Debug("Monitor handled APP frozen event")
	return nil
}

func (m *Monitor) handleAppWithdraw(appAbi *abi.ABI, log *types.Log) error {
	eventAppWithdraw, err := contract.UnpackAppWithdraw(appAbi, log)
	if err != nil {
		return err
	}

	logger := logrus.WithFields(logrus.Fields{
		"withdrawOperator": eventAppWithdraw.Operator,
		"withdrawReceiver": eventAppWithdraw.Receiver,
		"withdrawAccount":  eventAppWithdraw.Account,
		"withdrawAmount":   eventAppWithdraw.Amount.Int64(),
	})

	if err := m.contractEventObserver.OnWithdraw(eventAppWithdraw, log); err != nil {
		logger.WithError(err).Info("Monitor failed to handle APP withdraw event")
		return err
	}

	logger.Debug("Monitor handled APP withdraw event")
	return nil
}

func (m *Monitor) handleAppDeposit(appAbi *abi.ABI, log *types.Log) error {
	eventAppDeposit, err := contract.UnpackAppDeposit(appAbi, log)
	if err != nil {
		return err
	}

	if !util.IsZeroAddress(eventAppDeposit.Receiver) { // receiver address must not be zero
		return nil
	}

	logger := logrus.WithFields(logrus.Fields{
		"depositOperator": eventAppDeposit.Operator,
		"depositReceiver": eventAppDeposit.Receiver,
		"depositAmount":   eventAppDeposit.Amount.Int64(),
		"tokenID":         eventAppDeposit.TokenId.Int64(),
	})

	if err := m.contractEventObserver.OnDeposit(eventAppDeposit, log); err != nil {
		logger.WithError(err).Info("Monitor failed to handle APP deposit event")
		return err
	}

	logger.Debug("Monitor handled APP deposit event")
	return nil
}

func (m *Monitor) handleAppRegistryEvent(log *types.Log) (bool, error) {
	appRegistryAbi, err := contract.AppRegistryMetaData.GetAbi()
	if err != nil {
		return false, errors.WithMessage(err, "failed to get APP registry contract ABI")
	}

	// `Created` event concerned only
	createdEventId := appRegistryAbi.Events[contract.EventAppRegistryCreated].ID
	if log.Topics[0] != createdEventId {
		return false, nil
	}

	return m.handleAppRegistryCreated(appRegistryAbi, log)
}

func (m *Monitor) handleAppRegistryCreated(appRegistryAbi *abi.ABI, log *types.Log) (bool, error) {
	eventCreated, err := contract.UnpackAppRegistryCreated(appRegistryAbi, log)
	if err != nil {
		return false, err
	}

	logger := logrus.WithFields(logrus.Fields{
		"App":            eventCreated.App,
		"Operator":       eventCreated.Operator,
		"Owner":          eventCreated.Owner,
		"ApiWeightToken": eventCreated.ApiWeightToken,
	})

	if m.FilterOwnerAddress != nil && *m.FilterOwnerAddress != eventCreated.Owner {
		// not an APP by concerned owner
		logger.Debug("Monitor skipped `AppRegistryCreated` event due to not of concerned owner")
		return false, nil
	}

	// add observing for new created APP
	for i := range m.AppAddresses {
		if m.AppAddresses[i] == eventCreated.App { // already exists?
			logger.Debug("Monitor skipped `AppRegistryCreated` event due to APP already existed")
			return false, nil
		}
	}

	m.AppAddresses = append(m.AppAddresses, eventCreated.App)
	m.ApiWeightTokens = append(m.ApiWeightTokens, eventCreated.ApiWeightToken)
	// TODO: extract VipCoin from event log
	// m.VipCoins = append(m.VipCoins, eventCreated.VipCoin)

	if err := m.contractEventObserver.OnAppCreated(eventCreated, log); err != nil {
		logger.WithError(err).Info("Monitor failed to handle `AppRegistryCreated` event")
		return false, err
	}

	logger.Debug("Monitor handled `AppRegistryCreated` event")
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
