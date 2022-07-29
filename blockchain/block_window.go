package blockchain

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

// blockHashWindow caches sequent block hashes with limited capacity.
type blockHashWindow struct {
	// hashmap to cache hash of block (block number => block hash)
	blockToHash map[int64]common.Hash
	// maximum number of blocks to hold
	capacity uint32
	// cached block range
	blockFrom, blockTo int64
}

func newBlockHashWindow(capacity uint32) *blockHashWindow {
	win := &blockHashWindow{capacity: capacity}
	win.reset()

	return win
}

func (win *blockHashWindow) getBlockHash(block int64) (common.Hash, bool) {
	blockHash, ok := win.blockToHash[block]
	return blockHash, ok
}

func (win *blockHashWindow) reset() {
	win.blockFrom = -1
	win.blockTo = -1

	win.blockToHash = make(map[int64]common.Hash)
}

func (win *blockHashWindow) push(blockNum int64, blockHash, parentHash common.Hash) error {
	if win.size() > 0 { // validate incoming block
		if (win.blockTo + 1) != blockNum {
			return errors.Errorf(
				"incontinuous epoch pushed, expect %v got %v", win.blockTo+1, blockNum,
			)
		}

		latestHash, ok := win.blockToHash[win.blockTo]
		if !ok || parentHash != latestHash {
			return errors.Errorf(
				"mismatched parent hash, expect %v got %v", latestHash, parentHash,
			)
		}
	}

	// reclaim in case of memory blast
	for win.size() != 0 && win.size() >= win.capacity {
		delete(win.blockToHash, win.blockFrom)
		win.blockFrom++
	}

	// cache store block hash
	win.blockToHash[blockNum] = blockHash
	win.expandTo(blockNum)

	return nil
}

func (win *blockHashWindow) expandTo(newBlock int64) {
	if !win.isSet() {
		win.blockFrom, win.blockTo = newBlock, newBlock
	} else if win.blockTo < newBlock {
		win.blockTo = newBlock
	}
}

func (win *blockHashWindow) popn(blockUntil int64) {
	if win.size() == 0 || win.blockTo < blockUntil {
		return
	}

	for win.blockTo >= blockUntil {
		delete(win.blockToHash, win.blockTo)
		win.blockTo--

		if win.size() == 0 {
			win.reset()
			return
		}
	}
}

func (win *blockHashWindow) isSet() bool {
	return win.blockFrom != -1 && win.blockTo != -1
}

func (win *blockHashWindow) size() uint32 {
	if !win.isSet() || win.blockFrom > win.blockTo {
		return 0
	}

	return uint32(win.blockTo - win.blockFrom + 1)
}
