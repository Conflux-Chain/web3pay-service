package blockchain

import (
	"container/list"
	"math/big"
	"time"

	"github.com/Conflux-Chain/web3pay-service/contract"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	blockWinCapacity = 100

	syncIntervalNormal  = time.Second * 1
	syncIntervalCatchUp = time.Millisecond * 100
)

type MonitorConfig struct {
	ControllerAddress   common.Address   // controller contract
	AppCoinAddresses    []common.Address // APP coin contract list
	SyncFromBlockNumber int64            // the block number to start sync from
	SyncIntervalNormal  time.Duration    // interval to sync data in normal status
	SyncIntervalCatchUp time.Duration    // interval to sync data in catching up mode
}

func NewMonitorConfig(
	syncStartBlock int64, controllerAddr common.Address, appCoinAddrs []common.Address) *MonitorConfig {
	return &MonitorConfig{
		ControllerAddress:   controllerAddr,
		AppCoinAddresses:    appCoinAddrs,
		SyncFromBlockNumber: syncStartBlock,
		SyncIntervalNormal:  syncIntervalNormal,
		SyncIntervalCatchUp: syncIntervalCatchUp,
	}
}

// Monitor sync blockchain event logs to monitor contract events.
type Monitor struct {
	*MonitorConfig                              // monitor configurations
	provider              *Provider             // blockchain ops provider
	blockWindow           *blockHashWindow      // window to cache sequent block hashes
	contractEventObserver ContractEventObserver // contract event observer
}

func MustNewMonitor(provider *Provider, eventObserver ContractEventObserver) *Monitor {
	refBlockNumber := provider.ReferenceBlockNumber()
	baseCallOpt := &bind.CallOpts{
		BlockNumber: big.NewInt(refBlockNumber),
	}

	appCoinAddrs, err := provider.ListTrackedAppCoins(baseCallOpt)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to get APP coin list to init monitor")
	}

	config := NewMonitorConfig(refBlockNumber+1, stdConf.controllerContractAddr, appCoinAddrs)
	logrus.WithField("monitorConfig", config).Debug("Monitor config loaded")

	return &Monitor{
		MonitorConfig:         config,
		provider:              provider,
		blockWindow:           newBlockHashWindow(blockWinCapacity),
		contractEventObserver: eventObserver,
	}
}

func (m *Monitor) Sync() {
	logrus.WithField("syncFromBlock", m.SyncFromBlockNumber).
		Debug("Monitor starting to sync blockchain data")

	ticker := time.NewTicker(m.SyncIntervalCatchUp)
	defer ticker.Stop()

	confirmQueue := m.contractEventObserver.StatusConfirmQueue()
	confirmTasks := list.New()

	for {
		select {
		case <-ticker.C:
			complete, err := m.syncOnce(confirmTasks)
			if err != nil || complete {
				ticker.Reset(m.SyncIntervalNormal)
			} else {
				ticker.Reset(m.SyncIntervalCatchUp)
			}

			if err != nil {
				logrus.WithField("syncBlock", m.SyncFromBlockNumber).
					WithError(err).
					Error("Monitor failed to sync blockchain data")
			}
		case task := <-confirmQueue:
			confirmTasks.PushBack(task)
		}
	}
}

func (m *Monitor) syncOnce(confirmTasks *list.List) (bool, error) {
	latestBlockBigInt, err := m.provider.BlockNumber()
	if err != nil {
		return false, errors.WithMessage(err, "failed to query latest block")
	}

	syncBlockNum := (types.BlockNumber)(m.SyncFromBlockNumber)
	ceilBlockNumber := latestBlockBigInt.Int64() - skipBlocksAheadOfLeatestBlock

	logger := logrus.WithFields(logrus.Fields{
		"latestBlockNumber": latestBlockBigInt.Int64(),
		"ceilBlockNumber":   ceilBlockNumber,
		"syncBlockNo":       syncBlockNum,
	})

	if m.SyncFromBlockNumber > ceilBlockNumber { // already catched up to ceil
		logger.Debug("Monitor skipped sync due to already catched up")
		return true, nil
	}

	block, err := m.provider.BlockByNumber(syncBlockNum, false)
	if err != nil {
		logger.WithError(err).Info("Monitor failed to get block by number")
		return false, errors.WithMessage(err, "failed to get block by number")
	}

	prevBlockNum := int64(syncBlockNum) - 1
	prevBlockHash, exist := m.blockWindow.getBlockHash(prevBlockNum)

	if exist && block.ParentHash != prevBlockHash { // parent hash not matched
		err := m.reorgRevert(prevBlockNum)

		logger.WithFields(logrus.Fields{
			"prevBlockHash":   prevBlockHash,
			"parentBlockHash": block.ParentHash,
		}).WithError(err).Info("Monitor reorg revert due to parent hash mismatch")

		return false, errors.WithMessage(err, "failed to reorg revert")
	}

	// build log filters
	filterAddrs := []common.Address{m.ControllerAddress}
	filterAddrs = append(filterAddrs, m.AppCoinAddresses...)

	// TODO: filter topics too?
	logFilter := types.FilterQuery{
		FromBlock: &syncBlockNum, ToBlock: &syncBlockNum, Addresses: filterAddrs,
	}

	logs, err := m.provider.Logs(logFilter)
	if err != nil {
		logger.WithField("logFilter", logFilter).
			WithError(err).
			Info("Monitor failed to get event logs")
		return false, errors.WithMessage(err, "failed to get event logs")
	}

	logger.WithField("numLogs", len(logs)).Debug("Monitor fetched block event logs")

	for i := range logs {
		if logs[i].BlockHash != block.Hash || logs[i].BlockNumber != block.Number.Uint64() {
			// block number or hash not matched, chain reorg during fetch? have a retry
			return false, nil
		}

		// handle controller or APP coin contract events
		eventCategory, consumed, err := "", false, (error)(nil)

		if logs[i].Address == m.ControllerAddress {
			eventCategory = "controller"
			consumed, err = m.handleControllerEvent(&logs[i])
		} else {
			eventCategory = "appCoin"
			consumed, err = m.handleAppCoinEvent(&logs[i])
		}

		if err != nil {
			logger.WithFields(logrus.Fields{
				"log":           logs[i],
				"eventCategory": eventCategory,
			}).WithError(err).Info("Monitor failed to handle event log")

			return false, errors.WithMessage(err, "failed to handle event log")
		}

		logger.WithFields(logrus.Fields{
			"logIndex":      i,
			"eventCategory": eventCategory,
			"consumed":      consumed,
		}).Debug("Monitor handled event log")
	}

	if err := m.blockWindow.push(m.SyncFromBlockNumber, block.Hash, block.ParentHash); err != nil {
		logger.WithError(err).Info("Monitor failed to push block into cache window")
		m.blockWindow.reset()
	}

	// Confirm subscribed APP coin account status
	baseCallOpt := &bind.CallOpts{BlockNumber: big.NewInt(m.SyncFromBlockNumber)}

	// TODO: use batch call?
	for v := confirmTasks.Front(); v != nil; {
		task := v.Value.([2]common.Address)
		coin, addr := task[0], task[1]

		balance, frozen, err := m.provider.GetAppCoinBalanceAndFrozenStatus(baseCallOpt, coin, addr)
		if err != nil {
			logger.WithField("confirmTask", task).
				WithError(err).
				Info("Monitor failed to fetch APP coin account status for confirming task")
			v = v.Next()
			continue
		}

		err = m.contractEventObserver.OnConfirmStatus(coin, addr, balance, frozen, m.SyncFromBlockNumber)
		if err != nil {
			logger.WithField("confirmTask", task).
				WithError(err).
				Info("Monitor failed to confirm APP coin account status")
			v = v.Next()
			continue
		}

		nv := v.Next()
		confirmTasks.Remove(v)
		v = nv
	}

	m.SyncFromBlockNumber++
	return int64(syncBlockNum) == ceilBlockNumber, nil
}

func (m *Monitor) handleAppCoinEvent(log *types.Log) (bool, error) {
	appCoinAbi, err := contract.APPCoinMetaData.GetAbi()
	if err != nil {
		return false, errors.WithMessage(err, "failed to get APP coin contract ABI")
	}

	var event string

	switch log.Topics[0] {
	case appCoinAbi.Events[contract.EventAppCoinMinted].ID:
		// minted
		event = contract.EventAppCoinMinted
		err = m.handleAppCoinMinted(appCoinAbi, log)
	case appCoinAbi.Events[contract.EventAppCoinFrozen].ID:
		// frozen
		event = contract.EventAppCoinFrozen
		err = m.handleAppCoinFrozen(appCoinAbi, log)
	case appCoinAbi.Events[contract.EventAppCointWithdraw].ID:
		// withdraw
		event = contract.EventAppCointWithdraw
		err = m.handleAppCoinWithdraw(appCoinAbi, log)
	case appCoinAbi.Events[contract.EventAppCoinResourceChanged].ID:
		// resource changed
		event = contract.EventAppCoinResourceChanged
		err = m.handleAppCoinResourceChanged(appCoinAbi, log)
	default: // not concerned
		return false, nil
	}

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"log": log, "event": event,
		}).WithError(err).Info("Monitor failed to handle APP coin event")

		return false, err
	}

	logrus.WithField("event", event).Debug("Monitor handled APP coin event")
	return true, nil
}

func (m *Monitor) handleAppCoinResourceChanged(appCoinAbi *abi.ABI, log *types.Log) error {
	eventAppCoinResrcChanged, err := contract.UnpackAppCoinResourceChanged(appCoinAbi, log)
	if err != nil {
		return err
	}

	return m.contractEventObserver.OnResourceChanged(eventAppCoinResrcChanged)
}

func (m *Monitor) handleAppCoinWithdraw(appCoinAbi *abi.ABI, log *types.Log) error {
	eventAppCoinWithdraw, err := contract.UnpackAppCoinWithdraw(appCoinAbi, log)
	if err != nil {
		return err
	}

	return m.contractEventObserver.OnWithdraw(eventAppCoinWithdraw)
}

func (m *Monitor) handleAppCoinFrozen(appCoinAbi *abi.ABI, log *types.Log) error {
	eventAppCoinFrozen, err := contract.UnpackAppCoinFrozen(appCoinAbi, log)
	if err != nil {
		return err
	}

	return m.contractEventObserver.OnFrozen(eventAppCoinFrozen)
}

func (m *Monitor) handleAppCoinMinted(appCoinAbi *abi.ABI, log *types.Log) error {
	eventAppCoinMinted, err := contract.UnpackAppCoinMinted(appCoinAbi, log)
	if err != nil {
		return err
	}

	return m.contractEventObserver.OnMinted(eventAppCoinMinted)
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

	if stdConf.creatorAddr != nil && *stdConf.creatorAddr != eventAppCreated.AppOwner {
		// not an APP coin by concerned creator
		return false, nil
	}

	// add observing for new created APP coin
	m.AppCoinAddresses = append(m.AppCoinAddresses, eventAppCreated.Addr)
	return true, m.contractEventObserver.OnAppCreated(eventAppCreated)
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
