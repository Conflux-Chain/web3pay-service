package blockchain

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/Conflux-Chain/go-conflux-util/viper"
	"github.com/Conflux-Chain/web3pay-service/contract"
	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	// default list page size
	defaultListAppPageSize      = 50
	defaultListResourcePageSize = 50

	// default contract call timeout
	defaultContractCallTimeout = 3 * time.Second

	// number of diff blocks before latest as the base block
	baseDiffBlockAheadOfLatest = 40
)

// Provider provides blockchain data.
type Provider struct {
	w3client *web3go.Client
	conf     *Config

	bindCallContext *contractBindCallContext
	controller      *controllerContractObj
	appCoins        sync.Map // common.Address => *appCoinContractObj
}

func MustNewProviderFromViper(w3c *web3go.Client) *Provider {
	var conf Config
	viper.MustUnmarshalKey("blockchain", &conf)

	clientForContract, singerFn := w3c.ToClientForContract()
	ctrlAddr := common.HexToAddress(conf.ControllerContractAddr)

	latestBlockNo, err := w3c.Eth.BlockNumber()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to get latest block number")
	}

	// init controller contract stub
	ctrlCaller, err := contract.NewController(ctrlAddr, clientForContract)
	if err != nil {
		logrus.WithField("config", conf).
			WithError(err).
			Fatal("Failed to initialize controller contract")
	}

	p := &Provider{
		w3client: w3c,
		conf:     &conf,
		bindCallContext: &contractBindCallContext{
			baseBlockNumber: latestBlockNo.Int64() - baseDiffBlockAheadOfLatest,
			contractClient:  clientForContract,
			signer:          singerFn,
		},
		controller: newControllerContractObj(&ctrlAddr, nil, ctrlCaller),
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultContractCallTimeout)
	defer cancel()

	baseCallOpt := &bind.CallOpts{
		BlockNumber: big.NewInt(p.bindCallContext.baseBlockNumber),
		Context:     ctx,
	}

	// iterate all controller APPs to init APP contract instances
	if err := p.IterateControllerApps(baseCallOpt, func(addr common.Address) error {
		appCoinCaller, err := contract.NewAPPCoin(addr, clientForContract)
		if err != nil {
			logrus.WithField("addr", addr.String()).
				WithError(err).
				Debug("Failed to initialize APP coin contract")

			return errors.WithMessage(err, "failed to initialize APP coin contract")
		}

		// fetch APP coin contract owner
		appCoinOwner, err := appCoinCaller.AppOwner(baseCallOpt)
		if err != nil {
			logrus.WithField("addr", addr.String()).
				WithError(err).
				Debug("Failed to get APP coin contract owner")

			return errors.WithMessage(err, "failed to get APP coin contract owner")
		}

		coinContractObj := newAppCoinContractObj(&addr, &appCoinOwner, appCoinCaller)
		p.appCoins.Store(addr, coinContractObj)

		// iterate all resources under APP coin
		if err := p.IterateAppCoinResources(baseCallOpt, addr, func(confEntry *contract.AppConfigConfigEntry) error {
			coinContractObj.resources.Store(confEntry.ResourceId, confEntry)
			return nil
		}); err != nil {
			logrus.WithField("appCoin", addr.String()).
				WithError(err).
				Debug("Failed to get APP coin resources")

			return errors.WithMessage(err, "failed to get APP coin resources")
		}

		return nil

	}); err != nil {
		logrus.WithError(err).Fatal("Failed to initialize APP coin contract list")
	}

	if logrus.IsLevelEnabled(logrus.DebugLevel) {
		p.appCoins.Range(func(coin, contractV any) bool {
			contractObj := contractV.(*appCoinContractObj)
			resourceMaps := make(map[string]contract.AppConfigConfigEntry)

			contractObj.resources.Range(func(key, value any) bool {
				conf := value.(*contract.AppConfigConfigEntry)
				resourceMaps[key.(string)] = *conf
				return true
			})

			logrus.WithFields(logrus.Fields{
				"appCoin":   coin,
				"resources": resourceMaps,
			}).Debug("APP coin resources loaded")

			return true
		})

		logrus.WithFields(logrus.Fields{
			"controllerAddr":  ctrlAddr,
			"baseBlockNumber": p.bindCallContext.baseBlockNumber,
		}).Debug("Blockchain data provider initialized")
	}

	return p
}

func (p *Provider) GetBaseBlockNumber() int64 {
	return p.bindCallContext.baseBlockNumber
}

func (p *Provider) GetAppCoinFronzenStatusOfAddr(
	callOpts *bind.CallOpts, coin, address common.Address) (uint64, error) {
	atv, ok := p.appCoins.Load(coin)
	if !ok {
		return 0, model.ErrAppCoinNotFound
	}

	contractObj := atv.(*appCoinContractObj)

	fronzen, err := contractObj.stub.FrozenMap(callOpts, address)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"coin": coin, "address": address,
		}).WithError(err).Info("Failed to get APP coin fronze status")

		return 0, errors.WithMessage(err, "failed to get APP coin fronze status")
	}

	return fronzen.Uint64(), nil
}

func (p *Provider) GetAppCoinBalanceOfAddr(
	callOpts *bind.CallOpts, coin, address common.Address) (uint64, error) {
	atv, ok := p.appCoins.Load(coin)
	if !ok {
		return 0, model.ErrAppCoinNotFound
	}

	contractObj := atv.(*appCoinContractObj)

	balance, err := contractObj.stub.BalanceOf(callOpts, address)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"coin": coin, "address": address,
		}).WithError(err).Info("Failed to get APP coin balance")

		return 0, errors.WithMessage(err, "failed to get APP coin balance")
	}

	return balance.Uint64(), nil
}

// GetAppCoinContractOwner gets concerned APP coin contract owner.
func (p *Provider) GetAppCoinContractOwner(coin common.Address) (*common.Address, error) {
	atv, ok := p.appCoins.Load(coin)
	if !ok {
		return nil, model.ErrAppCoinNotFound
	}

	contractObj := atv.(*appCoinContractObj)
	return contractObj.owner, nil
}

// GetAppCoinResources gets concerned APP coin resource with specified id.
func (p *Provider) GetAppCoinResource(coin common.Address, resourceId string) (*contract.AppConfigConfigEntry, error) {
	atv, ok := p.appCoins.Load(coin)
	if !ok {
		return nil, model.ErrAppCoinNotFound
	}

	contractObj := atv.(*appCoinContractObj)

	if v, ok := contractObj.resources.Load(resourceId); ok {
		return v.(*contract.AppConfigConfigEntry), nil
	}

	return nil, model.ErrAppCoinResourceNotFound
}

// IterateControllerApps iterates all APP contracts created by controller contract.
// TODO: support config to filter by specific contract creator
func (p *Provider) IterateControllerApps(
	callOpts *bind.CallOpts, iterator func(common.Address) error, creator ...common.Address) error {
	for offset := int64(0); ; {
		appContractAddrs, total, err := p.controller.stub.ListApp(
			callOpts, big.NewInt(offset), big.NewInt(int64(defaultListAppPageSize)),
		)

		if err != nil {
			logrus.WithError(err).Debug("Failed to list APP contracts by controller")
			return errors.WithMessage(err, "failed to list APP contracts")
		}

		for i := range appContractAddrs {
			if err := iterator(appContractAddrs[i]); err != nil {
				logrus.WithError(err).Debug("Failed to iterate APP contracts by controller")
				return errors.WithMessage(err, "failed to iterate APP contract")
			}
		}

		offset += int64(len(appContractAddrs))
		if offset >= total.Int64() { // all done
			break
		}
	}

	return nil
}

// IterateAppCoinResources iterates all resource under specified APP coin.
func (p *Provider) IterateAppCoinResources(
	callOpts *bind.CallOpts, coin common.Address, iterator func(confEntry *contract.AppConfigConfigEntry) error,
) error {
	atv, ok := p.appCoins.Load(coin)
	if !ok {
		return model.ErrAppCoinNotFound
	}

	contractObj := atv.(*appCoinContractObj)

	for offset := int64(0); ; {
		configEntries, total, err := contractObj.stub.ListResources(
			callOpts, big.NewInt(offset), big.NewInt(int64(defaultListResourcePageSize)),
		)

		if err != nil {
			logrus.WithField("coin", coin).WithError(err).Debug("Failed to list resources for APP coin")
			return errors.WithMessage(err, "failed to list resources")
		}

		for i := range configEntries {
			if err := iterator(&configEntries[i]); err != nil {
				logrus.WithField("coin", coin).WithError(err).Debug("Failed to iterate APP coin resource")
				return errors.WithMessage(err, "failed to iterate resource")
			}
		}

		offset += int64(len(configEntries))
		if offset >= total.Int64() { // all done
			break
		}
	}

	return nil
}
