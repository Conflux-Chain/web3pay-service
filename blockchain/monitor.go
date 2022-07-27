package blockchain

import (
	"time"

	"github.com/Conflux-Chain/web3pay-service/contract"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	blockWinCapacity = 100
)

type EventSourcer interface {
	GetApiCoinControllerContract() controllerContractObj
	GetAppCoinContractList() []*appCoinContractObj
}

// Monitor sync blockchain event logs to monitor contract events.
type Monitor struct {
	sourcer  EventSourcer   // event sourcer
	w3client *web3go.Client // eth client

	blockWindow         *blockHashWindow // window to cache sequent block hashes
	syncFromBlockNumber int64            // the block number to sync from
	syncIntervalNormal  time.Duration    // interval to sync data in normal status
	syncIntervalCatchUp time.Duration    // interval to sync data in catching up mode
}

func NewMonitor(sourcer EventSourcer, w3c *web3go.Client) *Monitor {
	return &Monitor{
		sourcer:             sourcer,
		w3client:            w3c,
		blockWindow:         newBlockHashWindow(blockWinCapacity),
		syncIntervalNormal:  time.Second,
		syncIntervalCatchUp: time.Millisecond * 50,
	}
}

func (m *Monitor) Sync(fromBlock int64) {
	logrus.WithField("fromBlock", fromBlock).
		Debug("Monitor starting to sync blockchain data")

	m.syncFromBlockNumber = fromBlock

	ticker := time.NewTicker(m.syncIntervalCatchUp)
	defer ticker.Stop()

	for range ticker.C {
		complete, err := m.syncOnce()
		if err != nil || complete {
			ticker.Reset(m.syncIntervalNormal)
		} else {
			ticker.Reset(m.syncIntervalCatchUp)
		}

		if err != nil {
			logrus.WithField("syncBlock", m.syncFromBlockNumber).
				WithError(err).
				Error("Monitor failed to sync blockchain data")
		}
	}
}

func (m *Monitor) syncOnce() (bool, error) {
	latestBlockNumber, err := m.w3client.Eth.BlockNumber()
	if err != nil {
		return false, errors.WithMessage(err, "failed to query latest block")
	}

	syncBlockNum := (types.BlockNumber)(m.syncFromBlockNumber)
	latestBlockNum := latestBlockNumber.Int64()

	logger := logrus.WithFields(logrus.Fields{
		"latestBlockNumber": latestBlockNum,
		"syncBlockNo":       syncBlockNum,
	})

	if m.syncFromBlockNumber > latestBlockNum { // already catched up
		logger.Debug("Monitor skipped sync due to already catched up")
		return true, nil
	}

	block, err := m.w3client.Eth.BlockByNumber(syncBlockNum, false)
	if err != nil {
		logger.WithError(err).Info("Monitor failed to get block by number")
		return false, errors.WithMessage(err, "failed to get block by number")
	}

	prevBlockNum := int64(syncBlockNum) - 1
	prevBlockHash, exist := m.blockWindow.getBlockHash(prevBlockNum)

	if exist && block.ParentHash != prevBlockHash { // parent hash not matched
		if err := m.reorgRevert(prevBlockNum); err != nil {
			logger.WithFields(logrus.Fields{
				"prevBlockHash":   prevBlockHash,
				"parentBlockHash": block.ParentHash,
			}).WithError(err).Info("Monitor failed to reorg revert due to parent hash mismatch")

			return false, errors.WithMessage(err, "failed to reorg revert")
		}

		logger.WithFields(logrus.Fields{
			"prevBlockHash": prevBlockHash,
			"prevBlockNum":  prevBlockNum,
		}).Info("Monitor reorg reverted prev block due to parent hash mismatch")
		return false, nil
	}

	// TODO: maybe we can use a sync filter for this?
	ctrlContract := m.sourcer.GetApiCoinControllerContract()
	coinContractList := m.sourcer.GetAppCoinContractList()

	filterAddrs := []common.Address{*ctrlContract.addr}
	filterAddrsMap := make(map[common.Address]bool)
	for i := range coinContractList {
		filterAddrs = append(filterAddrs, *coinContractList[i].addr)
		filterAddrsMap[*coinContractList[i].addr] = true
	}

	// TODO: filter topics too?
	logFilter := types.FilterQuery{
		FromBlock: &syncBlockNum,
		ToBlock:   &syncBlockNum,
		Addresses: filterAddrs,
	}

	logs, err := m.w3client.Eth.Logs(logFilter)
	if err != nil {
		logger.WithField("logFilter", logFilter).
			WithError(err).
			Info("Monitor failed to get event logs")
		return false, errors.WithMessage(err, "failed to get event logs")
	}

	for i := range logs {
		if logs[i].BlockHash != block.Hash {
			// block hash not matched, chain reorg during fetch? have a retry
			return false, nil
		}

		eventCategory, consumed, err := "", false, (error)(nil)
		if logs[i].Address == *ctrlContract.addr {
			eventCategory = "controller"
			consumed, err = m.handleControllerEvent(&logs[i])
		} else if filterAddrsMap[logs[i].Address] {
			eventCategory = "appCoin"
			consumed, err = m.handleAppCoinEvent(&logs[i])
		} else { // not concerned
			continue
		}

		if err != nil {
			logger.WithFields(logrus.Fields{
				"log":           logs[i],
				"eventCategory": eventCategory,
			}).WithError(err).Info("Monitor failed to handle log event")

			return false, errors.WithMessage(err, "failed to handle log event")
		}

		logger.WithFields(logrus.Fields{
			"log":           logs[i],
			"eventCategory": eventCategory,
			"consumed":      consumed,
		}).Debug("Monitor handled controller event")
	}

	if err := m.blockWindow.push(m.syncFromBlockNumber, block.Hash, block.ParentHash); err != nil {
		logger.WithField("blockNumber", syncBlockNum).
			WithError(err).
			Info("Monitor failed to push block into cache window")
		m.blockWindow.reset()
	}

	m.syncFromBlockNumber++
	return false, nil
}

func (m *Monitor) handleAppCoinEvent(log *types.Log) (handled bool, err error) {
	appCoinAbi, err := contract.APPCoinMetaData.GetAbi()
	if err != nil {
		return false, errors.WithMessage(err, "failed to get APP coin contract ABI")
	}

	if len(log.Topics) == 0 {
		return false, nil
	}

	var event string

	switch log.Topics[0] {
	case appCoinAbi.Events[contract.EventAppCoinMinted].ID: // minted
		event = contract.EventAppCoinMinted
		err = m.handleAppCoinMinted(appCoinAbi, log)
	case appCoinAbi.Events[contract.EventAppCoinFrozen].ID: // frozen
		event = contract.EventAppCoinFrozen
		err = m.handleAppCoinFrozen(appCoinAbi, log)
	case appCoinAbi.Events[contract.EventAppCointWithdraw].ID: // withdraw
		event = contract.EventAppCointWithdraw
		err = m.handleAppCoinWithdraw(appCoinAbi, log)
	case appCoinAbi.Events[contract.EventAppCoinResourceChanged].ID: // resource changed
		event = contract.EventAppCoinResourceChanged
		err = m.handleAppCoinResourceChanged(appCoinAbi, log)
	default:
		return false, nil
	}

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"log":   log,
			"event": event,
		}).WithError(err).Info("Monitor failed to handle APP coin event")
		return false, err
	}

	logrus.WithFields(logrus.Fields{
		"log":   log,
		"event": event,
	}).Debug("Monitor handled APP coin event")
	return true, nil
}

func (m *Monitor) handleAppCoinResourceChanged(appCoinAbi *abi.ABI, log *types.Log) error {
	eventAppCoinWithdraw, err := contract.UnpackAppCoinWithdraw(appCoinAbi, log)
	if err != nil {
		return err
	}

	// TODO: notify observers for update
	_ = eventAppCoinWithdraw
	return nil
}

func (m *Monitor) handleAppCoinWithdraw(appCoinAbi *abi.ABI, log *types.Log) error {
	eventAppCoinWithdraw, err := contract.UnpackAppCoinWithdraw(appCoinAbi, log)
	if err != nil {
		return err
	}

	// TODO: notify observers for update
	_ = eventAppCoinWithdraw
	return nil
}

func (m *Monitor) handleAppCoinFrozen(appCoinAbi *abi.ABI, log *types.Log) error {
	eventAppCoinFrozen, err := contract.UnpackAppCoinFrozen(appCoinAbi, log)
	if err != nil {
		return err
	}

	// TODO: notify observers for update
	_ = eventAppCoinFrozen
	return nil
}

func (m *Monitor) handleAppCoinMinted(appCoinAbi *abi.ABI, log *types.Log) error {
	eventAppCoinMinted, err := contract.UnpackAppCoinMinted(appCoinAbi, log)
	if err != nil {
		return err
	}

	// TODO: notify observers for update
	_ = eventAppCoinMinted
	return nil
}

func (m *Monitor) handleControllerEvent(log *types.Log) (bool, error) {
	controllerAbi, err := contract.ControllerMetaData.GetAbi()
	if err != nil {
		return false, errors.WithMessage(err, "failed to get controller contract ABI")
	}

	// `APP_CREATED` event concerned only
	topic0 := controllerAbi.Events[contract.EventControllerAppCreated].ID
	if len(log.Topics) == 0 || log.Topics[0] != topic0 {
		return false, nil
	}

	eventAppCreated, err := contract.UnpackControllerAPPCREATED(controllerAbi, log)
	if err != nil {
		return false, err
	}

	// TODO: notify observers for update
	_ = eventAppCreated
	return true, nil
}

func (m *Monitor) reorgRevert(revertToBlock int64) error {
	// remove block hash of reverted block from cache window
	m.blockWindow.popn(revertToBlock)

	// update syncer start block
	m.syncFromBlockNumber = revertToBlock

	// TODO: notify observer for reorg revert
	logrus.WithFields(logrus.Fields{
		"revertToBlock":       revertToBlock,
		"syncFromBlockNumber": m.syncFromBlockNumber,
	}).Info("Monitor reorg reverted")

	return nil
}
