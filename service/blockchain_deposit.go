package service

import (
	"math"
	"math/big"
	"sort"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/patrickmn/go-cache"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

var (
	// progressive deposit amount arrival latency settings
	progressiveAmountDecimals []decimal.Decimal
	progressiveArrivalLatency []time.Duration
)

func init() {
	// These configurations are based on Conflux eSpace network with `CFX` coin unit.
	// TODO: support to load from configuration file.
	depositProgressiveArrivalLatency := map[float32]time.Duration{
		10:              time.Second * 0,      // for deposit amount (0, 10], instantly
		100:             time.Second * 3,      // for deposit amount (10, 100], 3s
		1000:            time.Second * 10,     // for deposit amount (100, 1000], 10s
		10000:           time.Second * 30,     // for deposit amount (1000, 10000], 30s
		100000:          time.Second * 60 * 1, // for deposit amount (10000, 100000], 1min
		1000000:         time.Second * 60 * 3, // for deposit amount (100000, 1000000], 3min
		math.MaxFloat32: time.Second * 60 * 5, // for deposit amount > 1000000, 5 min
	}

	progressiveAmounts := make([]float64, 0, len(depositProgressiveArrivalLatency))
	for k := range depositProgressiveArrivalLatency {
		progressiveAmounts = append(progressiveAmounts, float64(k))
	}

	sort.Float64s(progressiveAmounts)
	for _, dpAmount := range progressiveAmounts {
		progressiveAmountDecimals = append(
			progressiveAmountDecimals,
			decimal.NewFromFloatWithExponent(dpAmount, 18),
		)

		progressiveArrivalLatency = append(
			progressiveArrivalLatency,
			depositProgressiveArrivalLatency[float32(dpAmount)],
		)
	}
}

type DepositRequest struct {
	Coin        common.Address
	Address     common.Address
	Amount      *big.Int
	TxHash      common.Hash
	BlockHash   common.Hash
	BlockNumber int64
	SubmitAt    time.Time
}

func (bs *BlockchainService) DepositPending(request *DepositRequest) error {
	var arrivalLatency time.Duration

	for i, dpAmountDecimal := range progressiveAmountDecimals {
		// deposit amount <= progressive amount
		if request.Amount.Cmp(dpAmountDecimal.BigInt()) <= 0 {
			arrivalLatency = progressiveArrivalLatency[i]
			break
		}
	}

	// add to delay queue for furthur handling
	bs.delayQueue.Offer(request, time.Now().Add(arrivalLatency))

	return nil
}

func (bs *BlockchainService) Deposit() {
	for v := range bs.delayQueue.C {
		depositReq := v.(*DepositRequest)

		logrus.WithFields(logrus.Fields{
			"depositRequest": depositReq,
			"amount":         depositReq.Amount.String(),
		}).Debug("Handling pending deposit request")

		// skip duplicate deposit transaction hash
		if _, ok := bs.depositTxnHashCache.Get(depositReq.TxHash.String()); ok {
			logrus.WithField("txHash", depositReq.TxHash).
				Debug("Deposit transaction hash already existed")
			continue
		}

		// validate block hash
		block, err := bs.provider.BlockByNumber(rpc.BlockNumber(depositReq.BlockNumber), false)
		if err != nil {
			logrus.WithField("blockNumber", depositReq.BlockNumber).
				Info("Failed to get block by number to validate deposit txn")

			// retry it later
			bs.delayQueue.Offer(depositReq, time.Now())
			continue
		}

		if block.Hash != depositReq.BlockHash {
			logrus.WithFields(logrus.Fields{
				"blockHash":           block.Hash,
				"depositReqBlockHash": depositReq.BlockHash,
			}).Info("Block hash doesn't match to validate deposit txn")
			continue
		}

		// validate transaction
		txn, err := bs.provider.TransactionByHash(depositReq.TxHash)
		if err != nil {
			logrus.WithField("txnHash", depositReq.TxHash).
				Info("Failed to get transaction by hash to validate deposit txn")
			// retry it later
			bs.delayQueue.Offer(depositReq, time.Now())
			continue
		}

		if txn == nil { // transaction missing?
			logrus.WithField("txnHash", depositReq.TxHash).
				Info("Transaction missing to valdiate deposit txn")
			continue
		}

		if txn.BlockNumber == nil || txn.BlockHash == nil { // not mined
			logrus.WithField("txnHash", depositReq.TxHash).
				Info("Transaction not mined (block number or hash is nil) to valdiate deposit txn")
			continue
		}

		if *txn.BlockHash != block.Hash {
			logrus.WithFields(logrus.Fields{
				"txnHash":      depositReq.TxHash,
				"txnBlockHash": *txn.BlockHash,
				"blockHash":    block.Hash,
			}).Info("Transaction hash not matched to valdiate deposit txn")
			continue
		}

		if txn.BlockNumber.Cmp(block.Number) != 0 {
			logrus.WithFields(logrus.Fields{
				"txnHash":        depositReq.TxHash,
				"txnBlockNumber": txn.BlockNumber.String(),
				"blockNumber":    block.Number.String(),
			}).Info("Transaction block number not matched to valdiate deposit txn")
			continue
		}

		if txn.Status != nil && *txn.Status != types.ReceiptStatusSuccessful {
			logrus.WithFields(logrus.Fields{
				"txnHash":   depositReq.TxHash,
				"txnStatus": *txn.Status,
			}).Info("Transaction not succeeded to valdiate deposit txn")
			continue
		}

		deposited, err := bs.IncreaseAccountBalance(
			depositReq.Coin, depositReq.Address, depositReq.Amount, depositReq.BlockNumber,
		)
		if err != nil {
			logrus.WithError(err).Info("Failed to increase account balance for deposit")
			// retry it later
			bs.delayQueue.Offer(depositReq, time.Now())
			continue
		}

		bs.depositTxnHashCache.Set(depositReq.TxHash.String(), struct{}{}, cache.DefaultExpiration)
		logrus.WithField("deposited", deposited).Debug("Deposit request handled")
	}
}
