package service

import (
	"sync"
	"time"

	"github.com/Conflux-Chain/web3pay-service/blockchain"
	"github.com/Conflux-Chain/web3pay-service/store/memdb"
	"github.com/Conflux-Chain/web3pay-service/store/sqlite"
	myqueue "github.com/MoeYang/go-queue"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gammazero/workerpool"
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
)

const (
	statusConfirmQueueSize  = 5000
	workerPoolSize          = 2
	delayQueueSize          = 5000
	depositTxnHashCacheSize = 1_500
)

type BlockchainService struct {
	sqliteStore                  *sqlite.SqliteStore
	memStore                     *memdb.MemStore
	provider                     *blockchain.Provider
	appBaseMap                   map[common.Address]AppBase
	appMutex                     sync.Mutex
	appAccountStatusConfirmQueue chan [2]common.Address
	workerPool                   *workerpool.WorkerPool
	delayQueue                   *myqueue.DelayQueue
	depositTxnHashCache          *cache.Cache
}

func NewBlockchainService(
	sqliteStore *sqlite.SqliteStore, memStore *memdb.MemStore, provider *blockchain.Provider,
) (*BlockchainService, error) {
	bs := &BlockchainService{
		provider:                     provider,
		sqliteStore:                  sqliteStore,
		memStore:                     memStore,
		appBaseMap:                   make(map[common.Address]AppBase),
		appAccountStatusConfirmQueue: make(chan [2]common.Address, statusConfirmQueueSize),
		workerPool:                   workerpool.New(workerPoolSize),
		delayQueue:                   myqueue.NewDelayQueue(delayQueueSize),

		// Create a deposit transaction hash cache with a default expiration time of 5 minutes,
		// and purges expired items every 10 minutes
		depositTxnHashCache: cache.New(5*time.Minute, 10*time.Minute),
	}

	if err := bs.initApps(); err != nil {
		return nil, errors.WithMessage(err, "failed to initialize APPs")
	}

	go bs.delayQueue.Poll()
	go bs.Deposit()
	go bs.delayExecResourceConfig()
	go bs.checkOperatorBalance()

	return bs, nil
}
