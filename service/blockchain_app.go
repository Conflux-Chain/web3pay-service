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
	"github.com/spf13/viper"
)

const (
	defaultResourceId = "default"

	// APP resource config delay exec interval
	delayConfigExecIntervalNormal = time.Minute * 5
)

type resourceConfig = contract.IAppConfigConfigEntry
type AppBase struct {
	Addr           common.Address            // APP contract address
	ApiWeightToken common.Address            // `ApiWeightToken` contract address
	VipCoin        common.Address            // `VipCoin` contract address
	PendingSeconds *big.Int                  // pending seconds for delaying config
	Resources      map[string]resourceConfig // resource config
	UpdateBlock    int64                     // the last block of which to update resources
}

func (bs *BlockchainService) initApps() error {
	refBlockNumber := bs.provider.ReferenceBlockNumber()
	baseCallOpt := &bind.CallOpts{
		BlockNumber: big.NewInt(int64(refBlockNumber)),
	}

	// iterate all APPs to init APP contract instances
	err := bs.provider.IterateTrackedApps(baseCallOpt, func(app common.Address) error {
		appContract, err := bs.provider.GetAppContract(app)
		if err != nil {
			return errors.WithMessage(err, "failed to get `app` contract")
		}

		// check if the operator has charging role?
		chargeRole, err := appContract.CHARGEROLE(baseCallOpt)
		if err != nil {
			return errors.WithMessage(err, "failed to get APP `chargeRole`")
		}

		granted, err := appContract.HasRole(baseCallOpt, chargeRole, bs.provider.OperatorAddress())
		if err != nil {
			return errors.WithMessage(err, "failed to check APP `chargeRole`")
		}

		if !granted {
			return errors.New("`chargeRole` not granted for the operator")
		}

		vipCoinAddr, err := bs.provider.GetAppVipCoin(baseCallOpt, app)
		if err != nil {
			return err
		}

		awtAddr, err := bs.provider.GetAppApiWeightToken(baseCallOpt, app)
		if err != nil {
			return err
		}

		pendingSec, err := bs.provider.GetApiWeightTokenPendingSeconds(baseCallOpt, awtAddr)
		if err != nil {
			return errors.WithMessage(err, "failed to get pending seconds")
		}

		resources, err := bs.provider.GetApiWeightTokenResources(baseCallOpt, awtAddr)
		if err != nil {
			return err
		}

		bs.appBaseMap[app] = AppBase{
			Addr:           app,
			ApiWeightToken: awtAddr,
			VipCoin:        vipCoinAddr,
			PendingSeconds: pendingSec,
			Resources:      resources,
			UpdateBlock:    refBlockNumber,
		}

		return nil
	})
	if err != nil {
		return errors.WithMessage(err, "failed to init APPs")
	}

	logrus.WithFields(logrus.Fields{
		"appBases":             bs.appBaseMap,
		"referenceBlockNumber": refBlockNumber,
	}).Debug("Blockchain service APP bases initialized")

	return nil
}

func (svc *BlockchainService) GetAppConfigResourceWithId(
	app common.Address, resourceId string) (*contract.IAppConfigConfigEntry, error) {
	if len(resourceId) == 0 { // if resourceId is empty, use default resource
		resourceId = defaultResourceId
	}

	svc.appMutex.Lock()
	defer svc.appMutex.Unlock()

	ab, ok := svc.appBaseMap[app]
	if !ok {
		return nil, model.ErrAppNotFound
	}

	resrc, ok := ab.Resources[resourceId]
	if ok && resrc.PendingOP != contract.OpCodeResourceConfigAdd {
		return &resrc, nil
	}

	if resourceId != defaultResourceId { // retry default resource
		resrc, ok = ab.Resources[defaultResourceId]
		if ok {
			return &resrc, nil
		}
	}

	return nil, model.ErrResourceNotFound
}

// ValidateAppOperator validates if the account address is the valid operator for the APP contract.
func (svc *BlockchainService) ValidateAppOperator(appAddr, operator common.Address) error {
	if _, ok := svc.getAppBase(appAddr); !ok {
		return model.ErrAppNotFound
	}

	expecOperator := svc.provider.OperatorAddress()

	logrus.WithFields(logrus.Fields{
		"appContractAddr":          appAddr,
		"expectedContractOperator": expecOperator,
		"actualOperator":           operator,
	}).Debug("Blockchain service validated APP contract operator")

	if !reflect.DeepEqual(expecOperator, operator) {
		return model.ErrInvalidAppOperator
	}

	return nil
}

func (svc *BlockchainService) getAppBase(app common.Address) (*AppBase, bool) {
	svc.appMutex.Lock()
	defer svc.appMutex.Unlock()

	ab, ok := svc.appBaseMap[app]
	return &ab, ok
}

// delayExecResourceConfigOnce periodically executes APP config resources.
func (bs *BlockchainService) delayExecResourceConfig() {
	bs.execResourceConfigOnce()

	viper.SetDefault("blockchain.delayConfigExecInterval", delayConfigExecIntervalNormal)
	delayConfigExecIntervalNormal := viper.GetDuration("blockchain.delayConfigExecInterval")

	ticker := time.NewTicker(delayConfigExecIntervalNormal)
	defer ticker.Stop()

	for range ticker.C {
		err := bs.execResourceConfigOnce()
		if err != nil {
			logrus.WithError(err).Error("Failed to run delay exec resource config once")
			ticker.Reset(delayConfigExecIntervalNormal / 2)
		} else {
			ticker.Reset(delayConfigExecIntervalNormal)
		}
	}
}

func (bs *BlockchainService) execResourceConfigOnce() error {
	return bs.provider.IterateTrackedApps(nil, func(app common.Address) error {
		ab, ok := bs.getAppBase(app)
		if !ok { // APP not found yet
			return nil
		}

		// iterate all tracking APP
		return bs.provider.IterateApiWeightTokenResources(nil, ab.ApiWeightToken, func(confEntry contract.IAppConfigConfigEntry) (bool, error) {
			// checking pending operation code
			if confEntry.PendingOP == contract.OpCodeResourceConfigNoPending {
				logrus.WithFields(logrus.Fields{
					"app": app, "configEntry": confEntry,
				}).Debug("Skipping execute APP config resource config due to no pending operation")

				return false, nil
			}

			expireTimestamp := big.NewInt(0).Add(confEntry.SubmitSeconds, ab.PendingSeconds)
			nowTimeStamp := time.Now().Unix()

			// check time is up or not
			if expireTimestamp.Int64() > nowTimeStamp {
				logrus.WithFields(logrus.Fields{
					"submitTimeStamp": confEntry.SubmitSeconds.Int64(),
					"pendingSeconds":  ab.PendingSeconds.Int64(),
					"nowTimeStamp":    nowTimeStamp,
				}).Debug("Skipping execute APP resource config due to time is not up yet")

				return false, nil
			}

			// contract call `flushPendingConfig`
			txn, err := bs.provider.FlushApiWeightTokenPendingConfig(nil, ab.ApiWeightToken)
			if err != nil {
				logrus.WithField("app", app).
					WithError(err).
					Info("Failed to flush APP pending resource configurations")

				return false, errors.WithMessage(err, "failed to flush pending resource config")
			}

			logrus.WithFields(logrus.Fields{
				"app":            app,
				"configEntry":    confEntry,
				"txnHash":        txn.Hash,
				"pendingSeconds": ab.PendingSeconds.Int64(),
			}).Debug("APP pending resource config flushed")

			return true, nil
		})
	})
}
