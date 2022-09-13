package service

import (
	"sync"
	"time"

	"github.com/Conflux-Chain/web3pay-service/blockchain"
	"github.com/Conflux-Chain/web3pay-service/store/memdb"
	"github.com/Conflux-Chain/web3pay-service/store/sqlite"
	"github.com/Conflux-Chain/web3pay-service/util"
	myqueue "github.com/MoeYang/go-queue"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gammazero/workerpool"
	lru "github.com/hashicorp/golang-lru"
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	signatureAddressCacheSize = 50_000
	statusConfirmQueueSize    = 5000
	workerPoolSize            = 2
	delayQueueSize            = 5000
	depositTxnHashCacheSize   = 1_500
)

type BlockchainService struct {
	sigAddrCache              *lru.Cache // sha3(sig) => addr
	sqliteStore               *sqlite.SqliteStore
	memStore                  *memdb.MemStore
	provider                  *blockchain.Provider
	appCoinBaseMap            map[common.Address]AppCoinBase
	appCoinMutex              sync.Mutex
	appCoinStatusConfirmQueue chan [2]common.Address
	workerPool                *workerpool.WorkerPool
	delayQueue                *myqueue.DelayQueue
	depositTxnHashCache       *cache.Cache
}

func NewBlockchainService(
	sqliteStore *sqlite.SqliteStore, memStore *memdb.MemStore, provider *blockchain.Provider,
) (*BlockchainService, error) {
	bs := &BlockchainService{
		provider:                  provider,
		sqliteStore:               sqliteStore,
		memStore:                  memStore,
		appCoinBaseMap:            make(map[common.Address]AppCoinBase),
		appCoinStatusConfirmQueue: make(chan [2]common.Address, statusConfirmQueueSize),
		workerPool:                workerpool.New(workerPoolSize),
		delayQueue:                myqueue.NewDelayQueue(delayQueueSize),

		// Create a deposit transaction hash cache with a default expiration time of 5 minutes,
		// and which purges expired items every 10 minutes
		depositTxnHashCache: cache.New(5*time.Minute, 10*time.Minute),
	}

	lruCache, err := lru.New(signatureAddressCacheSize)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to init sig addr cache")
	}
	bs.sigAddrCache = lruCache

	if err := bs.initAppCoins(); err != nil {
		return nil, errors.WithMessage(err, "failed to initialize APP coins")
	}

	go bs.delayQueue.Poll()
	go bs.Deposit()
	go bs.delayExecResourceConfig()
	go bs.checkOperatorBalance()

	return bs, nil
}

// RecoverAddressBySignature recovers signer address from message and signature.
// Also cache the recovered address for later use to improve performance.
func (svc *BlockchainService) RecoverAddressBySignature(msg, sig string) (string, error) {
	logger := logrus.WithFields(logrus.Fields{
		"msg": msg, "sig": sig,
	})

	cacheKey := crypto.Keccak256Hash([]byte(sig))

	val, ok := svc.sigAddrCache.Get(cacheKey)
	if ok { // hit in cache
		addr := val.(string)

		logger.WithFields(logrus.Fields{
			"addr": addr, "cacheKey": cacheKey,
		}).Debug("Get address by signagure from the cache")

		return addr, nil
	}

	lockKey := util.MutexKey(cacheKey.String())
	util.KLock(lockKey)
	defer util.KUnlock(lockKey)

	if val, ok := svc.sigAddrCache.Get(cacheKey); ok { // double checking
		return val.(string), nil
	}

	addr, err := util.RecoverAddress(msg, sig)
	if err != nil {
		logger.WithError(err).Debug("Failed to recover address by signature")

		return "", err
	}

	logger.WithFields(logrus.Fields{
		"addr": addr, "cacheKey": cacheKey,
	}).Debug("Address recovered from signature")

	svc.sigAddrCache.Add(cacheKey, addr)

	return addr, nil
}
