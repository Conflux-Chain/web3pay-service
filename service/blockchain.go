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
)

type BlockchainService struct {
	sigAddrCache *lru.Cache // sha3(sig) => addr
	provider     *blockchain.Provider
	kmutex       *util.KMutex
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

	return bs, nil
}

func (svc *BlockchainService) GetAppCoinResourceWithId(
	appCoin common.Address, resourceId string) (*contract.AppConfigConfigEntry, error) {
	if len(resourceId) == 0 { // if resourceId is empty, use default resource
		resourceId = "default"
	}

	return svc.provider.GetAppCoinResource(appCoin, resourceId)
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
