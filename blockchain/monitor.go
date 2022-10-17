package blockchain

import (
	"container/list"
	"context"
	"math/big"
	"time"

	"github.com/Conflux-Chain/web3pay-service/metrics"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	blockWinCapacity = 200

	syncIntervalNormal  = time.Second * 1
	syncIntervalCatchUp = time.Millisecond * 300
)

var (
	errChainReorged = errors.New("chain re-orged")
	errGetLogs      = errors.New("getLogs error")
)

type monitorConfig struct {
	AppRegistryAddress  common.Address   // APP registry contract address
	FilterOwnerAddress  *common.Address  // filtering APP owner
	AppAddresses        []common.Address //  APP contract list
	ApiWeightTokens     []common.Address // ApiWeightToken list
	VipCoins            []common.Address // VipCoin list
	SyncFromBlockNumber int64            // the block number to start sync from
	SyncIntervalNormal  time.Duration    // interval to sync data in normal status
	SyncIntervalCatchUp time.Duration    // interval to sync data in catching up mode
	UseFastCatchUp      bool             // whether to use fast catchup
}

// Monitor sync blockchain event logs to monitor contract events.
type Monitor struct {
	*monitorConfig                              // monitor configurations
	provider              *Provider             // blockchain ops provider
	blockWindow           *blockHashWindow      // window to cache sequent block hashes
	contractEventObserver ContractEventObserver // contract event observer
}

func MustNewMonitor(config *Config, provider *Provider, eventObserver ContractEventObserver) *Monitor {
	refBlockNumber := provider.ReferenceBlockNumber()
	baseCallOpt := &bind.CallOpts{
		BlockNumber: big.NewInt(refBlockNumber),
	}

	appAddrs, awtAddrs, vcAddrs, err := provider.ListTrackedApps(baseCallOpt)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to get APP list to init monitor")
	}

	monConfig := &monitorConfig{
		AppRegistryAddress:  config.AppRegistryContractAddr,
		FilterOwnerAddress:  config.OwnerAddr,
		AppAddresses:        appAddrs,
		ApiWeightTokens:     awtAddrs,
		VipCoins:            vcAddrs,
		SyncFromBlockNumber: refBlockNumber + 1,
		SyncIntervalNormal:  syncIntervalNormal,
		SyncIntervalCatchUp: syncIntervalCatchUp,
	}
	logrus.WithField("monitorConfig", config).Debug("Monitor config loaded")

	return &Monitor{
		monitorConfig:         monConfig,
		provider:              provider,
		blockWindow:           newBlockHashWindow(blockWinCapacity),
		contractEventObserver: eventObserver,
	}
}

func (m *Monitor) Sync(ctx context.Context) {
	logrus.WithField("syncFromBlock", m.SyncFromBlockNumber).
		Debug("Monitor starting to sync blockchain data")

	if m.UseFastCatchUp {
		m.fastCatchup()
	}

	ticker := time.NewTicker(m.SyncIntervalNormal)
	defer ticker.Stop()

	confirmQueue := m.contractEventObserver.StatusConfirmQueue()
	confirmTasks := list.New()

	for {
		select {
		case <-ctx.Done():
			logrus.Info("Monitor sync completed")
			return
		case <-ticker.C:
			start := time.Now()
			complete, err := m.syncOnce(confirmTasks)
			metrics.Monitor.SyncOnceQps(err).UpdateSince(start)

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

// fastCatchup fast catches up to the latest block number.
func (m *Monitor) fastCatchup() {
	for {
		latestBlockBigInt, err := m.provider.BlockNumber()
		if err != nil {
			logrus.WithError(err).Error("Monitor failed to query latest block during fast catchup")
			continue
		}

		goAfterBlockNumber := latestBlockBigInt.Int64() - skipBlocksAheadOfLeatestBlock
		if m.SyncFromBlockNumber > goAfterBlockNumber { // already catched up?
			logrus.WithField("latestBlockNumber", goAfterBlockNumber).Debug("Monitor fast catched up latest block")
			return
		}

		// sync to the target block number with binary probing in case of error resulted in by
		// loose `getLogs` search filter.
		for {
			err := m.catchupOnce(goAfterBlockNumber)
			if err == nil || errors.Is(err, errChainReorged) { // success or reorg?
				break
			}

			if errors.Is(err, errGetLogs) { // getLogs query error
				// narrow down `getLogs` search filter in case of error caused by huge query/result set
				goAfterBlockNumber = (m.SyncFromBlockNumber + goAfterBlockNumber) / 2

				logrus.WithField("newTarget", goAfterBlockNumber).
					Info("Monitor adjusted target block number due to getLogs error during fast catchup")
				continue
			}

			logrus.WithFields(logrus.Fields{
				"start": m.SyncFromBlockNumber,
				"end":   goAfterBlockNumber,
			}).WithError(err).Error("Monitor failed to sync once during fast catchup")

			time.Sleep(time.Second)
		}
	}
}

func (m *Monitor) catchupOnce(target int64) error {
	start := (types.BlockNumber)(m.SyncFromBlockNumber)
	end := (types.BlockNumber)(target)

	startBlock, err := m.provider.BlockByNumber(start, false)
	if err != nil {
		return errors.WithMessage(err, "failed to get block by number")
	}

	endBlock := startBlock
	if start != end {
		endBlock, err = m.provider.BlockByNumber(end, false)
		if err != nil {
			return errors.WithMessage(err, "failed to get block by number")
		}
	}

	logger := logrus.WithFields(logrus.Fields{
		"start": start, "end": end,
	})

	// parent hash checking
	prevBlockNum := int64(start) - 1
	prevBlockHash, exist := m.blockWindow.getBlockHash(prevBlockNum)

	if exist && startBlock.ParentHash != prevBlockHash { // parent hash not matched
		logger.WithFields(logrus.Fields{
			"prevBlockHash":   prevBlockHash,
			"parentBlockHash": startBlock.ParentHash,
		}).Info("Monitor parent hash mismatch detected")

		if err := m.reorgRevert(prevBlockNum); err != nil {
			logger.WithError(err).
				Error("Monitor failed to reorg revert due to parent hash mismatch")
		}

		return errChainReorged
	}

	// get event logs

	// build log filters
	filterAddrs := []common.Address{m.AppRegistryAddress}
	filterAddrs = append(filterAddrs, m.AppAddresses...)
	filterAddrs = append(filterAddrs, m.ApiWeightTokens...)
	filterAddrs = append(filterAddrs, m.VipCoins...)

	logFilter := types.FilterQuery{
		FromBlock: &start, ToBlock: &end, Addresses: filterAddrs,
	}

	logs, err := m.provider.Logs(logFilter)
	if err != nil {
		logrus.WithField("logFilter", logFilter).WithError(err).Info("Monitor failed to get event logs")
		return errGetLogs
	}

	logger.WithField("numLogs", len(logs)).Debug("Monitor fetched block event logs during catchup")

	validate := func(log *types.Log, block *types.Block) bool {
		return log.BlockNumber != block.Number.Uint64() || log.BlockHash == block.Hash
	}

	for i := range logs {
		if !validate(&logs[i], startBlock) ||
			(startBlock != endBlock && !validate(&logs[i], endBlock)) {
			logger.WithFields(logrus.Fields{
				"log":        logs[i],
				"startBlock": startBlock,
				"endBlock":   endBlock,
			}).Info("Monitor detected mismatched block hash for event log")

			// block hash not matched, chain reorg during fetch?
			return errChainReorged
		}

		// handle contract events
		var err error
		switch {
		case m.isAppRegistryEvent(&logs[i]): // `AppRegistry` related event
			_, err = m.handleAppRegistryEvent(&logs[i])
		case m.isAppEvent(&logs[i]): // `App` related event
			_, err = m.handleAppEvent(&logs[i])
		case m.isApiWeightTokenEvent(&logs[i]): // `ApiWeightToken` related event
			_, err = m.handleApiWeightTokenEvent(&logs[i])
		case m.isVipCoinEvent((&logs[i])): // `VipCoin` related event
			_, err = m.handleVipCoinEvent(&logs[i])
		default:
			logrus.WithField("log", logs[i]).Warn("Monitor detected unconcerned event log")
		}

		if err != nil {
			logger.WithField("log", logs[i]).WithError(err).Info("Monitor failed to handle event log")
			return errors.WithMessage(err, "failed to handle event log")
		}
	}

	blockHash, parentHash := startBlock.Hash, startBlock.ParentHash
	for i := start; i <= end; i++ {
		err := m.blockWindow.push(i.Int64(), blockHash, parentHash)
		if err != nil {
			logger.WithField("blockNum", i).
				WithError(err).
				Error("Monitor failed to push block into cache window")
			m.blockWindow.reset()
		}

		parentHash = blockHash
		blockHash = common.Hash{}

		if i == end-1 {
			blockHash = endBlock.Hash
		}
	}

	m.SyncFromBlockNumber = end.Int64() + 1
	return nil
}

func (m *Monitor) syncOnce(confirmTasks *list.List) (bool, error) {
	latestBlockBigInt, err := m.provider.BlockNumber()
	if err != nil {
		return false, errors.WithMessage(err, "failed to query latest block")
	}

	syncBlockNum := (types.BlockNumber)(m.SyncFromBlockNumber)
	goAfterBlockNumber := latestBlockBigInt.Int64() - skipBlocksAheadOfLeatestBlock

	logger := logrus.WithFields(logrus.Fields{
		"latestBlockNumber":  latestBlockBigInt.Int64(),
		"goAfterBlockNumber": goAfterBlockNumber,
		"syncBlockNo":        syncBlockNum,
	})

	if m.SyncFromBlockNumber > goAfterBlockNumber { // already catched up to ceil
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
		logger.WithFields(logrus.Fields{
			"prevBlockHash":   prevBlockHash,
			"parentBlockHash": block.ParentHash,
		}).Info("Monitor parent hash mismatch detected")

		if err := m.reorgRevert(prevBlockNum); err != nil {
			logger.WithError(err).Info("Monitor failed to reorg revert due to parent hash mismatch")
			return false, errors.WithMessage(err, "failed to reorg revert")
		}

		return false, nil
	}

	// build log filters
	filterAddrs := []common.Address{m.AppRegistryAddress}
	filterAddrs = append(filterAddrs, m.AppAddresses...)
	filterAddrs = append(filterAddrs, m.ApiWeightTokens...)
	filterAddrs = append(filterAddrs, m.VipCoins...)

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

	if len(logs) > 0 {
		logger.WithField("numLogs", len(logs)).Debug("Monitor fetched block event logs")
	}

	for i := range logs {
		if logs[i].BlockHash != block.Hash || logs[i].BlockNumber != block.Number.Uint64() {
			// block number or hash not matched, chain reorg during fetch? have a retry
			return false, errors.New("mismatched block hash or number of event log")
		}

		// handle contract events
		var err error
		switch {
		case m.isAppRegistryEvent(&logs[i]): // `AppRegistry` related event
			_, err = m.handleAppRegistryEvent(&logs[i])
		case m.isAppEvent(&logs[i]): // `App` related event
			_, err = m.handleAppEvent(&logs[i])
		case m.isApiWeightTokenEvent(&logs[i]): // `ApiWeightToken` related event
			_, err = m.handleApiWeightTokenEvent(&logs[i])
		case m.isVipCoinEvent((&logs[i])): // `VipCoin` related event
			_, err = m.handleVipCoinEvent(&logs[i])
		default:
			logrus.WithField("log", logs[i]).Warn("Monitor detected unconcerned event log")
		}

		if err != nil {
			logger.WithField("log", logs[i]).WithError(err).Info("Monitor failed to handle event log")
			return false, errors.WithMessage(err, "failed to handle event log")
		}
	}

	if err := m.blockWindow.push(m.SyncFromBlockNumber, block.Hash, block.ParentHash); err != nil {
		logger.WithError(err).Info("Monitor failed to push block into cache window")
		m.blockWindow.reset()
	}

	m.confirmSubscribedAccountStatus(confirmTasks)

	m.SyncFromBlockNumber++
	return int64(syncBlockNum) == goAfterBlockNumber, nil
}

func (m *Monitor) confirmSubscribedAccountStatus(confirmTasks *list.List) {
	// Confirm subscribed APP account status
	baseCallOpt := &bind.CallOpts{BlockNumber: big.NewInt(m.SyncFromBlockNumber)}

	// TODO: use batch call or multiple call?
	for v := confirmTasks.Front(); v != nil; {
		start := time.Now()
		task := v.Value.([2]common.Address)
		app, addr := task[0], task[1]

		balance, frozen, err := m.provider.GetAppAccountBalanceAndFrozenStatus(baseCallOpt, app, addr)
		if err != nil {
			logrus.WithField("confirmTask", task).
				WithError(err).
				Info("Monitor failed to fetch APP account status for confirming task")
			metrics.Monitor.ConfirmQps(err).UpdateSince(start)
			v = v.Next()
			continue
		}

		err = m.contractEventObserver.OnConfirmStatus(app, addr, balance, frozen, m.SyncFromBlockNumber)
		if err != nil {
			logrus.WithField("confirmTask", task).
				WithError(err).
				Info("Monitor failed to confirm APP account status")
			metrics.Monitor.ConfirmQps(err).UpdateSince(start)
			v = v.Next()
			continue
		}

		metrics.Monitor.ConfirmQps(err).UpdateSince(start)
		nv := v.Next()
		confirmTasks.Remove(v)
		v = nv
	}
}

func (m *Monitor) isAppRegistryEvent(log *types.Log) bool {
	return log.Address == m.AppRegistryAddress
}

func (m *Monitor) isAppEvent(log *types.Log) bool {
	return inAddressSlice(m.AppAddresses, log.Address)
}

func (m *Monitor) isApiWeightTokenEvent(log *types.Log) bool {
	return inAddressSlice(m.ApiWeightTokens, log.Address)
}

func (m *Monitor) isVipCoinEvent(log *types.Log) bool {
	return inAddressSlice(m.VipCoins, log.Address)
}

func inAddressSlice(addrs []common.Address, target common.Address) bool {
	for _, addr := range addrs {
		if addr == target {
			return true
		}
	}

	return false
}
