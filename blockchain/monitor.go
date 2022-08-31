package blockchain

import (
	"container/list"
	"math/big"
	"time"

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
	ControllerAddress    common.Address   // controller contract
	FilterCreatorAddress *common.Address  // filtering APP coin creator
	AppCoinAddresses     []common.Address // APP coin contract list
	SyncFromBlockNumber  int64            // the block number to start sync from
	SyncIntervalNormal   time.Duration    // interval to sync data in normal status
	SyncIntervalCatchUp  time.Duration    // interval to sync data in catching up mode
	UseFastCatchUp       bool             // whether to use fast catchup
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

	appCoinAddrs, err := provider.ListTrackedAppCoins(baseCallOpt)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to get APP coin list to init monitor")
	}

	monConfig := &monitorConfig{
		ControllerAddress:    config.ControllerContractAddr,
		FilterCreatorAddress: config.CreatorAddr,
		AppCoinAddresses:     appCoinAddrs,
		SyncFromBlockNumber:  refBlockNumber + 1,
		SyncIntervalNormal:   syncIntervalNormal,
		SyncIntervalCatchUp:  syncIntervalCatchUp,
	}
	logrus.WithField("monitorConfig", config).Debug("Monitor config loaded")

	return &Monitor{
		monitorConfig:         monConfig,
		provider:              provider,
		blockWindow:           newBlockHashWindow(blockWinCapacity),
		contractEventObserver: eventObserver,
	}
}

func (m *Monitor) Sync() {
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
	filterAddrs := []common.Address{m.ControllerAddress}
	filterAddrs = append(filterAddrs, m.AppCoinAddresses...)
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
		case logs[i].Address == m.ControllerAddress: // controller event
			_, err = m.handleControllerEvent(&logs[i])
		default: // APP coin or airdrop event
			_, err = m.handleAppCoinEvent(&logs[i])
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
			return false, errors.New("mismatched block hash or number of event log")
		}

		// handle contract events
		var err error
		switch {
		case logs[i].Address == m.ControllerAddress: // controller event
			_, err = m.handleControllerEvent(&logs[i])
		default: // APP coin or airdrop event
			_, err = m.handleAppCoinEvent(&logs[i])
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
	return int64(syncBlockNum) == goAfterBlockNumber, nil
}
