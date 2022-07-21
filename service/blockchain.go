package service

import (
	"reflect"

	"github.com/Conflux-Chain/web3pay-service/blockchain"
	"github.com/Conflux-Chain/web3pay-service/util"
	"github.com/ethereum/go-ethereum/common"
	lru "github.com/hashicorp/golang-lru"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	signatureAddressCacheSize = 1_000_000
)

type BlockchainService struct {
	cache    *lru.Cache
	provider *blockchain.Provider
}

func NewBlockchainService(provider *blockchain.Provider) (*BlockchainService, error) {
	lruCache, err := lru.New(signatureAddressCacheSize)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to init LRU cache")
	}

	return &BlockchainService{
		cache:    lruCache,
		provider: provider,
	}, nil
}

// ValidateAppCoinContractOwner validates if the specific blockchain address is the owner for
// the contract of the specified address.
func (svc *BlockchainService) ValidateAppCoinContractOwner(contractAddr, owner common.Address) error {
	contractOwner := svc.provider.GetAppCoinContractOwner(contractAddr)

	logrus.WithFields(logrus.Fields{
		"contractAddr":          contractAddr,
		"expectedContractOwner": owner,
		"actualContractOwner":   contractOwner,
	}).Debug("Validating APP coin contract owner...")

	if contractOwner == nil {
		return errors.New("contract not found")
	}

	if !reflect.DeepEqual(contractOwner, &owner) {
		return errors.New("invalid contract owner")
	}

	return nil
}

// RecoverAddressBySignature recovers signer address from message and signature.
// Also cache the recovered address for later use to improve performance.
func (svc *BlockchainService) RecoverAddressBySignature(msg, sig string) (string, error) {
	logger := logrus.WithFields(logrus.Fields{
		"msg": msg, "sig": sig,
	})

	val, ok := svc.cache.Get(sig)
	if ok { // hit in cache
		addr := val.(string)
		logger.WithField("addr", addr).Debug("Get address by signagure from the cache")

		return addr, nil
	}

	addr, err := util.RecoverAddress(msg, sig)
	if err != nil {
		logger.WithError(err).Debug("Failed to recover address by signature")

		return "", err
	}

	logger.WithField("addr", addr).Debug("Address recovered from signature")
	svc.cache.Add(sig, addr)

	return addr, nil
}
