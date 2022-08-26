package service

import (
	"math/big"
	"reflect"
	"time"

	"github.com/Conflux-Chain/web3pay-service/contract"
	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	defaultResourceId = "default"

	// APP resource config delay exec interval
	delayConfigExecIntervalNormal = time.Minute * 5
	delayConfigExecIntervalPush   = time.Second * 15
)

type resourceConfig = contract.AppConfigConfigEntry
type AppCoinBase struct {
	Addr        common.Address            // contract address
	Owner       common.Address            // owner address
	Resources   map[string]resourceConfig // resource config
	UpdateBlock int64                     // the last block of which to update resources
}

func (bs *BlockchainService) initAppCoins() error {
	refBlockNumber := bs.provider.ReferenceBlockNumber()
	baseCallOpt := &bind.CallOpts{
		BlockNumber: big.NewInt(int64(refBlockNumber)),
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
			UpdateBlock: refBlockNumber, Addr: coin,
			Owner: *owner, Resources: resources,
		}

		return nil
	})
	if err != nil {
		return errors.WithMessage(err, "failed to init APP coins")
	}

	logrus.WithFields(logrus.Fields{
		"appCoinBases":         bs.appCoinBaseMap,
		"referenceBlockNumber": refBlockNumber,
	}).Debug("Blockchain service APP coin bases initialized")

	return nil
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
	if ok {
		return &resrc, nil
	}

	if resourceId != defaultResourceId { // retry default resource
		resrc, ok = appCoin.Resources[defaultResourceId]
		if ok {
			return &resrc, nil
		}
	}

	return nil, model.ErrAppCoinResourceNotFound
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

// delayExecResourceConfigOnce periodically executes configuring APP coin resources.
func (bs *BlockchainService) delayExecResourceConfig() {
	bs.execResourceConfigOnce()

	ticker := time.NewTicker(delayConfigExecIntervalNormal)
	defer ticker.Stop()

	for range ticker.C {
		err := bs.execResourceConfigOnce()
		if err != nil {
			logrus.WithError(err).Error("Failed to run delay exec resource config once")
			ticker.Reset(delayConfigExecIntervalPush)
		} else {
			ticker.Reset(delayConfigExecIntervalNormal)
		}
	}
}

func (bs *BlockchainService) execResourceConfigOnce() error {
	return bs.provider.IterateTrackedAppCoins(nil, func(coin common.Address) error {
		// get pending seconds config
		pendingSeconds, err := bs.provider.GetAppConfigPendingSeconds(nil, coin)
		if err != nil {
			logrus.WithField("appCoin", coin).
				WithError(err).
				Info("Failed to executing APP coin resource config")
			return errors.WithMessage(err, "failed to get pending seconds config")
		}

		// iterate all tracking APP coins
		return bs.provider.IterateAppCoinResources(nil, coin, func(confEntry contract.AppConfigConfigEntry) error {
			// checking pending operation code
			if confEntry.PendingOP == contract.OpCodeResourceConfigNoPending {
				logrus.WithFields(logrus.Fields{
					"appCoin":     coin,
					"configEntry": confEntry,
				}).Debug("Skipping execute APP coin resource config due to no pending operation")
				return nil
			}

			expireTimestamp := big.NewInt(0).Add(confEntry.SubmitSeconds, pendingSeconds)
			nowTimeStamp := time.Now().Unix()

			// check time up or not
			if expireTimestamp.Int64() > nowTimeStamp {
				logrus.WithFields(logrus.Fields{
					"submitTimeStamp": confEntry.SubmitSeconds.Int64(),
					"pendingSeconds":  pendingSeconds.Int64(),
					"nowTimeStamp":    nowTimeStamp,
				}).Debug("Skipping execute APP coin resource config due to time not up yet")
				return nil
			}

			// contract call `flushPendingResourceConfig`
			txn, err := bs.provider.FlushPendingResourceConfig(nil, coin)
			if err != nil {
				logrus.WithField("appCoin", coin).
					WithError(err).
					Info("Failed to flush APP coin pending resource configurations")
				return errors.WithMessage(err, "failed to flush pending resource config")
			}

			logrus.WithFields(logrus.Fields{
				"appCoin":        coin,
				"configEntry":    confEntry,
				"txnHash":        txn.Hash,
				"pendingSeconds": pendingSeconds.Int64(),
			}).Debug("APP coin resource config executed")
			return nil
		})
	})
}
