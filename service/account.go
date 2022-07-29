package service

import (
	"github.com/Conflux-Chain/web3pay-service/blockchain"
	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	lru "github.com/hashicorp/golang-lru"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	accountStatusCacheSize = 1_00_000
)

// AccountStatus APP coin account status
type AccountStatus struct {
	Frozen  uint64
	Balance uint64
}

// AccountService APP coin account service
type AccountService struct {
	provider           *blockchain.Provider
	accountStatusCache *lru.Cache // sha3(coin, addr) => *AccountStatus
	kmutex             *util.KMutex
}

func NewAccountService(provider *blockchain.Provider) (*AccountService, error) {
	bs := &AccountService{
		provider: provider,
		kmutex:   util.NewKMutex(),
	}

	lruCache, err := lru.New(accountStatusCacheSize)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to init account status cache")
	}

	bs.accountStatusCache = lruCache

	return bs, nil
}

// DeductBalance deducts APP coin balance of specific address.
func (svc *AccountService) DeductBalance(appCoin, address common.Address, amount uint64) (uint64, error) {
	cacheKey := crypto.Keccak256Hash(appCoin[:], address[:])

	lockKey := util.MutexKey(cacheKey.String())
	svc.kmutex.Lock(lockKey)
	defer svc.kmutex.Unlock(lockKey)

	if val, ok := svc.accountStatusCache.Get(cacheKey); ok {
		addrStatus := val.(*AccountStatus)

		if addrStatus.Balance < amount {
			return 0, model.ErrInsufficentBalance
		}

		addrStatus.Balance -= amount
		return addrStatus.Balance, nil
	}

	return 0, model.ErrAppCoinAddrBalanceNotFound
}

// GetStatus gets APP coin status (balance, frozen status etc.,) of specific address.
func (svc *AccountService) GetStatus(appCoin, address common.Address) (*AccountStatus, error) {
	cacheKey := crypto.Keccak256Hash(appCoin[:], address[:])

	logger := logrus.WithFields(logrus.Fields{
		"appCoin": appCoin, "address": address, "cacheKey": cacheKey,
	})

	if val, ok := svc.accountStatusCache.Get(cacheKey); ok {
		status := val.(*AccountStatus)
		logger.WithField("status", status).Debug("App coin status for address hit in cache")
		return status, nil
	}

	lockKey := util.MutexKey(cacheKey.String())
	svc.kmutex.Lock(lockKey)
	defer svc.kmutex.Unlock(lockKey)

	balance, err := svc.provider.GetAppCoinBalance(nil, appCoin, address)
	if err != nil {
		logger.WithError(err).Info("Failed to get APP coin balance")
		return nil, errors.WithMessage(err, "failed to get APP coin balance")
	}

	frozen, err := svc.provider.GetAppCoinFrozenStatus(nil, appCoin, address)
	if err != nil {
		logger.WithError(err).Info("Failed to get APP coin frozen status")
		return nil, errors.WithMessage(err, "failed to get APP coin frozen status")
	}

	if val, ok := svc.accountStatusCache.Get(cacheKey); ok { // double checking
		return val.(*AccountStatus), nil
	}

	coinStatus := &AccountStatus{
		Balance: balance, Frozen: frozen,
	}

	logger.WithField("appCoinStatus", coinStatus).Debug("Fetched APP coin status for address")

	svc.accountStatusCache.Add(cacheKey, coinStatus)
	return coinStatus, err
}
