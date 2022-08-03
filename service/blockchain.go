package service

import (
	"sync"

	"github.com/Conflux-Chain/web3pay-service/blockchain"
	"github.com/Conflux-Chain/web3pay-service/store/memdb"
	"github.com/Conflux-Chain/web3pay-service/store/sqlite"
	"github.com/Conflux-Chain/web3pay-service/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gammazero/workerpool"
	lru "github.com/hashicorp/golang-lru"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	signatureAddressCacheSize = 1_00_000
)

type BlockchainService struct {
	sigAddrCache              *lru.Cache // sha3(sig) => addr
	kmutex                    *util.KMutex
	sqliteStore               *sqlite.SqliteStore
	memStore                  *memdb.MemStore
	provider                  *blockchain.Provider
	appCoinBaseMap            map[common.Address]AppCoinBase
	appCoinMutex              sync.Mutex
	appCoinStatusConfirmQueue chan [2]common.Address
	workerPool                *workerpool.WorkerPool
}

func NewBlockchainService(
	sqliteStore *sqlite.SqliteStore, memStore *memdb.MemStore, provider *blockchain.Provider,
) (*BlockchainService, error) {
	bs := &BlockchainService{
		provider:                  provider,
		kmutex:                    util.NewKMutex(),
		sqliteStore:               sqliteStore,
		memStore:                  memStore,
		appCoinBaseMap:            make(map[common.Address]AppCoinBase),
		appCoinStatusConfirmQueue: make(chan [2]common.Address, 10),
		workerPool:                workerpool.New(2),
	}

	lruCache, err := lru.New(signatureAddressCacheSize)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to init sig addr cache")
	}
	bs.sigAddrCache = lruCache

	if err := bs.initAppCoins(); err != nil {
		return nil, errors.WithMessage(err, "failed to initialize APP coins")
	}

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
	svc.kmutex.Lock(lockKey)
	defer svc.kmutex.Unlock(lockKey)

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
