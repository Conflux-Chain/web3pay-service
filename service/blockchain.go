package service

import (
	"reflect"

	"github.com/Conflux-Chain/web3pay-service/blockchain"
	"github.com/Conflux-Chain/web3pay-service/contract"
	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	lru "github.com/hashicorp/golang-lru"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	signatureAddressCacheSize = 1_00_000
	coinAddrStatusCacheSize   = 1_00_000
)

var (
	errNotAnValidAppCoinOwner     = errors.New("not a valid APP coin contract owner")
	errAppCoinAddrBalanceNotFound = errors.New("address balance of APP coin not found")
)

type BlockchainService struct {
	sigAddrCache        *lru.Cache // sha3(sig) => addr
	coinAddrStatusCache *lru.Cache // sha3(coin, addr) => *model.AppCoinAddrStatus
	provider            *blockchain.Provider
	kmutex              *util.KMutex
}

func NewBlockchainService(provider *blockchain.Provider) (*BlockchainService, error) {
	bs := &BlockchainService{
		provider: provider,
		kmutex:   util.NewKMutex(),
	}

	lruCache, err := lru.New(signatureAddressCacheSize)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to init sig addr cache")
	}
	bs.sigAddrCache = lruCache

	lruCache, err = lru.New(coinAddrStatusCacheSize)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to init coin addr status cache")
	}
	bs.coinAddrStatusCache = lruCache

	return bs, nil
}

// DeductAppCoinBalance deducts APP coin balance of specific address.
func (svc *BlockchainService) DeductAppCoinBalance(appCoin, address common.Address, amount uint64) (uint64, error) {
	cacheKey := crypto.Keccak256Hash(appCoin[:], address[:])

	lockKey := util.MutexKey(cacheKey.String())
	svc.kmutex.Lock(lockKey)
	defer svc.kmutex.Unlock(lockKey)

	if val, ok := svc.coinAddrStatusCache.Get(cacheKey); ok {
		addrStatus := val.(*model.AppCoinAddrStatus)

		if addrStatus.Balance < amount {
			return 0, errInsufficentBalance
		}

		addrStatus.Balance -= amount
	}

	return 0, errAppCoinAddrBalanceNotFound
}

func (svc *BlockchainService) GetAppCoinResourceWithId(
	appCoin common.Address, resourceId string) (*contract.AppConfigConfigEntry, error) {
	return svc.provider.GetAppCoinResource(appCoin, resourceId)
}

// GetAppCoinBalanceOfAddress gets APP coin status (balance, frozen status etc.,) of specific address.
func (svc *BlockchainService) GetAppCoinStatusOfAddr(appCoin, address common.Address) (*model.AppCoinAddrStatus, error) {
	cacheKey := crypto.Keccak256Hash(appCoin[:], address[:])

	logger := logrus.WithFields(logrus.Fields{
		"appCoin": appCoin, "address": address, "cacheKey": cacheKey,
	})

	if val, ok := svc.coinAddrStatusCache.Get(cacheKey); ok {
		status := val.(*model.AppCoinAddrStatus)
		logger.WithField("status", status).Debug("App coin status for address hit in cache")
		return status, nil
	}

	lockKey := util.MutexKey(cacheKey.String())
	svc.kmutex.Lock(lockKey)
	defer svc.kmutex.Unlock(lockKey)

	balance, err := svc.provider.GetAppCoinBalanceOfAddr(appCoin, address)
	if err != nil {
		logger.WithError(err).Info("Failed to get APP coin balance")
		return nil, errors.WithMessage(err, "failed to get APP coin balance")
	}

	frozen, err := svc.provider.GetAppCoinFronzenStatusOfAddr(appCoin, address)
	if err != nil {
		logger.WithError(err).Info("Failed to get APP coin frozen status")
		return nil, errors.WithMessage(err, "failed to get APP coin frozen status")
	}

	if val, ok := svc.coinAddrStatusCache.Get(cacheKey); ok { // double checking
		return val.(*model.AppCoinAddrStatus), nil
	}

	coinStatus := &model.AppCoinAddrStatus{
		Balance: balance, Frozen: frozen,
	}

	logger.WithField("appCoinStatus", coinStatus).Debug("Fetched APP coin status for address")

	svc.coinAddrStatusCache.Add(cacheKey, coinStatus)
	return coinStatus, err
}

// ValidateAppCoinContractOwner validates if the specific blockchain address is the owner for
// the contract of the specified address.
func (svc *BlockchainService) ValidateAppCoinContractOwner(contractAddr, owner common.Address) error {
	contractOwner, err := svc.provider.GetAppCoinContractOwner(contractAddr)

	logrus.WithFields(logrus.Fields{
		"contractAddr":          contractAddr,
		"expectedContractOwner": owner,
		"actualContractOwner":   contractOwner,
	}).WithError(err).Debug("Get APP coin contract owner")

	if err != nil {
		return err
	}

	if !reflect.DeepEqual(contractOwner, &owner) {
		return errNotAnValidAppCoinOwner
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
