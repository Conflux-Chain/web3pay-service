package service

import (
	"math/big"
	"reflect"
	"sync"

	"github.com/Conflux-Chain/web3pay-service/blockchain"
	"github.com/Conflux-Chain/web3pay-service/contract"
	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/util"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	lru "github.com/hashicorp/golang-lru"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	signatureAddressCacheSize = 1_00_000

	// skip blocks ahead of latest block number to reduce chain reorg
	// while sync or state call.
	skipBlocksAheadOfLeatestBlock = 50
)

type resourceConfig = contract.AppConfigConfigEntry
type AppCoinBase struct {
	Addr      common.Address            // contract address
	Owner     common.Address            // owner address
	Resources map[string]resourceConfig // resource config
}

type BlockchainService struct {
	sigAddrCache   *lru.Cache // sha3(sig) => addr
	provider       *blockchain.Provider
	kmutex         *util.KMutex
	appCoinBaseMap map[common.Address]AppCoinBase
	mutex          sync.Mutex
}

func NewBlockchainService(provider *blockchain.Provider) (*BlockchainService, error) {
	bs := &BlockchainService{
		provider:       provider,
		kmutex:         util.NewKMutex(),
		appCoinBaseMap: make(map[common.Address]AppCoinBase),
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

func (bs *BlockchainService) initAppCoins() error {
	refBlockNumber := bs.provider.ReferenceBlockNumber()
	baseCallOpt := &bind.CallOpts{
		BlockNumber: big.NewInt(refBlockNumber),
	}

	// iterate all controller APPs to init APP contract instances
	err := bs.provider.IterateTrackedAppCoins(baseCallOpt, func(coin common.Address) error {
		owner, err := bs.provider.GetAppCoinContractOwner(baseCallOpt, coin)
		if err != nil {
			return err
		}

		resources, err := bs.provider.GetAppCoinResources(baseCallOpt, coin)
		if err != nil {
			return err
		}

		bs.appCoinBaseMap[coin] = AppCoinBase{
			Addr: coin, Owner: *owner, Resources: resources,
		}

		return nil
	})

	logrus.WithFields(logrus.Fields{
		"appCoinBases":         bs.appCoinBaseMap,
		"referenceBlockNumber": refBlockNumber,
	}).Debug("Blockchain service APP coin bases initialized")

	return err
}

func (svc *BlockchainService) GetAppCoinResourceWithId(
	coin common.Address, resourceId string) (*contract.AppConfigConfigEntry, error) {
	if len(resourceId) == 0 { // if resourceId is empty, use default resource
		resourceId = "default"
	}

	svc.mutex.Lock()
	defer svc.mutex.Unlock()

	appCoin, ok := svc.appCoinBaseMap[coin]
	if !ok {
		return nil, model.ErrAppCoinNotFound
	}

	resrc, ok := appCoin.Resources[resourceId]
	if !ok {
		return nil, model.ErrAppCoinResourceNotFound
	}

	return &resrc, nil
}

func (svc *BlockchainService) GetAppCoinOwner(coin common.Address) (*common.Address, error) {
	svc.mutex.Lock()
	defer svc.mutex.Unlock()

	appCoin, ok := svc.appCoinBaseMap[coin]
	if !ok {
		return nil, model.ErrAppCoinNotFound
	}

	return &appCoin.Owner, nil
}

// ValidateAppCoinOwner validates if the specific blockchain address is the owner for
// the APP coin contract of the specified address.
func (svc *BlockchainService) ValidateAppCoinOwner(contractAddr, owner common.Address) error {
	contractOwner, err := svc.GetAppCoinOwner(contractAddr)
	if err != nil {
		return err
	}

	logrus.WithFields(logrus.Fields{
		"contractAddr":          contractAddr,
		"expectedContractOwner": owner,
		"actualContractOwner":   contractOwner,
	}).Debug("Blockchain service validated APP coin owner")

	if !reflect.DeepEqual(contractOwner, &owner) {
		return model.ErrNotAnValidAppCoinOwner
	}

	return nil
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
