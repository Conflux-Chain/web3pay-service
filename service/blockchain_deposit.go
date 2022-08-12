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
		100:             time.Second * 0,      // for deposit amount (0, 100], instantly
		1000:            time.Second * 3,      // for deposit amount (100, 1000], 3s
		10000:           time.Second * 5,      // for deposit amount (1000, 10000], 5s
		100000:          time.Second * 10,     // for deposit amount (10000, 100000], 10s
		1000000:         time.Second * 30,     // for deposit amount (100000, 1000000], 30s
		10000000:        time.Second * 60,     // for deposit amount (1000000, 10000000], 1min
		math.MaxFloat32: time.Second * 60 * 2, // for deposit amount > 10000000, 2 min
	}

	progressiveAmounts := make([]float64, 0, len(depositProgressiveArrivalLatency))
	for k := range depositProgressiveArrivalLatency {
		progressiveAmounts = append(progressiveAmounts, float64(k))
	}

	sort.Float64s(progressiveAmounts)
	for i := range progressiveAmounts {
		dpAmount := progressiveAmounts[i]
		progressiveAmountDecimals = append(
			progressiveAmountDecimals,
			decimal.NewFromFloat(dpAmount).Mul(decimal.New(1, 18)),
		)

		progressiveArrivalLatency = append(
			progressiveArrivalLatency,
			depositProgressiveArrivalLatency[float32(dpAmount)],
		)
	}

	logrus.WithFields(logrus.Fields{
		"progressiveArrivalLatency": progressiveArrivalLatency,
		"progressiveAmounts":        progressiveAmounts,
		"progressiveAmountDecimals": progressiveAmountDecimals,
	}).Debug("Loaded progressive deposit amount to arrival latency hierarchy settings")
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

	logrus.WithFields(logrus.Fields{
		"depositRequest":    request,
		"depositAmount":     request.Amount.String(),
		"arrivalLatency(s)": arrivalLatency.Seconds(),
	}).Debug("Blockchain service pending deposit request...")

	// add to delay queue for furthur handling
	bs.delayQueue.Offer(request, time.Now().Add(arrivalLatency))
	return nil
}

func (bs *BlockchainService) Deposit() {
	for v := range bs.delayQueue.C {
		depositReq := v.(*DepositRequest)

		logger := logrus.WithFields(logrus.Fields{
			"depositRequest": depositReq,
			"amount":         depositReq.Amount.String(),
		})

		// skip duplicate deposit transaction hash
		if _, ok := bs.depositTxnHashCache.Get(depositReq.TxHash.String()); ok {
			logger.Info("Deposit transaction hash already existed")
			continue
		}

		// validate block hash
		block, err := bs.provider.BlockByNumber(rpc.BlockNumber(depositReq.BlockNumber), false)
		if err != nil {
			logger.Info("Failed to get block by number to validate deposit txn")

			// retry it later
			bs.delayQueue.Offer(depositReq, time.Now())
			continue
		}

		if block.Hash != depositReq.BlockHash {
			logger.WithField("opponentBlockHash", block.Hash).
				Info("Block hash doesn't match to validate deposit txn")
			continue
		}

		// validate transaction
		txn, err := bs.provider.TransactionByHash(depositReq.TxHash)
		if err != nil {
			logger.Info("Failed to get transaction by hash to validate deposit txn")
			// retry it later
			bs.delayQueue.Offer(depositReq, time.Now())
			continue
		}

		if txn == nil { // transaction missing?
			logger.Info("Transaction missing to valdiate deposit txn")
			continue
		}

		if txn.BlockNumber == nil || txn.BlockHash == nil { // not mined?
			logger.Info("Transaction not mined (block number or hash is nil) to valdiate deposit txn")
			continue
		}

		if *txn.BlockHash != block.Hash {
			logger.WithField("txnBlockHash", *txn.BlockHash).
				Info("Transaction block hash not matched to valdiate deposit txn")
			continue
		}

		if txn.BlockNumber.Cmp(block.Number) != 0 {
			logger.WithField("txnBlockNumber", txn.BlockNumber.String()).
				Info("Transaction block hash not matched to valdiate deposit txn")
			continue
		}

		if txn.Status == nil || *txn.Status != types.ReceiptStatusSuccessful {
			logger.WithField("txnStatus", txn.Status).
				Info("Transaction status not successful to valdiate deposit txn")
			continue
		}

		deposited, err := bs.IncreaseAccountBalance(
			depositReq.Coin, depositReq.Address,
			depositReq.Amount, depositReq.BlockNumber,
		)
		if err != nil {
			logger.WithError(err).
				Error("Blockchain service failed to increase APP coin account balance")
			continue
		}

		logger.WithField("deposited", deposited).
			Debug("Blockchain service handled deposit request")

		// cache txn hash for dedupe
		bs.depositTxnHashCache.Set(depositReq.TxHash.String(), struct{}{}, cache.DefaultExpiration)
	}
}
