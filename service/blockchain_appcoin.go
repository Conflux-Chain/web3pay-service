package service

import (
	"math/big"
	"reflect"

	"github.com/Conflux-Chain/web3pay-service/contract"
	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
)

const (
	defaultResourceId = "default"
)

type resourceConfig = contract.AppConfigConfigEntry
type AppCoinBase struct {
	Addr      common.Address            // contract address
	Owner     common.Address            // owner address
	Resources map[string]resourceConfig // resource config
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
		resourceId = defaultResourceId
	}

	svc.appCoinMutex.Lock()
	defer svc.appCoinMutex.Unlock()

	appCoin, ok := svc.appCoinBaseMap[coin]
	if !ok {
		return nil, model.ErrAppCoinNotFound
	}

	resrc, ok := appCoin.Resources[resourceId]
	if !ok {
		if resourceId != defaultResourceId { // not existed?
			// use default resource
			return svc.GetAppCoinResourceWithId(coin, defaultResourceId)
		}

		return nil, model.ErrAppCoinResourceNotFound
	}

	return &resrc, nil
}

func (svc *BlockchainService) GetAppCoinOwner(coin common.Address) (*common.Address, error) {
	svc.appCoinMutex.Lock()
	defer svc.appCoinMutex.Unlock()

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
