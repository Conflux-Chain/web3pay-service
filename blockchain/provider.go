package blockchain

import (
	"math/big"

	"github.com/Conflux-Chain/go-conflux-util/viper"
	"github.com/Conflux-Chain/web3pay-service/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	listAppPageSize      = 50
	listResourcePageSize = 50
)

// Provider provides blockchain data.
type Provider struct {
	w3client *web3go.Client
	conf     *Config

	bindCallContext *contractBindCallContext
	controller      *controllerContractObj
	appCoins        map[common.Address]*appCoinContractObj
}

func MustNewProviderFromViper(w3c *web3go.Client) *Provider {
	var conf Config
	viper.MustUnmarshalKey("blockchain", &conf)

	clientForContract, singerFn := w3c.ToClientForContract()
	ctrlAddr := common.HexToAddress(conf.ControllerContractAddr)

	// init controller contract stub
	ctrlCaller, err := contract.NewController(ctrlAddr, clientForContract)
	if err != nil {
		logrus.WithField("config", conf).
			WithError(err).
			Fatal("Failed to initialize controller contract")
	}

	p := &Provider{
		w3client:        w3c,
		conf:            &conf,
		bindCallContext: &contractBindCallContext{clientForContract, singerFn},
		appCoins:        make(map[common.Address]*appCoinContractObj),
		controller:      newControllerContractObj(&ctrlAddr, nil, ctrlCaller),
	}

	// iterate all controller APPs to init APP contract instances
	if err := p.IterateControllerApps(func(addr common.Address) error {
		appCoinCaller, err := contract.NewAPPCoin(addr, clientForContract)
		if err != nil {
			logrus.WithField("addr", addr.String()).
				WithError(err).
				Debug("Failed to initialize APP coin contract")

			return errors.WithMessage(err, "failed to initialize APP coin contract")
		}

		// fetch APP coin contract owner
		appCoinOwner, err := appCoinCaller.AppOwner(nil)
		if err != nil {
			logrus.WithField("addr", addr.String()).
				WithError(err).
				Debug("Failed to get APP coin contract owner")

			return errors.WithMessage(err, "failed to get APP coin contract owner")
		}

		coinContractObj := newAppCoinContractObj(&addr, &appCoinOwner, appCoinCaller)
		p.appCoins[addr] = coinContractObj

		// iterate all resources under APP coin
		if err := p.IterateAppCoinResources(addr, func(confEntry contract.AppConfigConfigEntry) error {
			coinContractObj.resources[confEntry.ResourceId] = confEntry
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
		for coin, contractObj := range p.appCoins {
			logrus.WithFields(logrus.Fields{
				"appCoin": coin, "resources": contractObj.resources,
			}).Debug("APP coin resources loaded")
		}

		logrus.WithFields(logrus.Fields{
			"config":             conf,
			"controllerAddr":     ctrlAddr,
			"controllerAppCoins": p.appCoins,
		}).Debug("Blockchain data provider initialized")
	}

	return p
}

// GetAppCoinContractOwner gets concerned APP coin contract owner.
func (p *Provider) GetAppCoinContractOwner(contractAddr common.Address) *common.Address {
	if contractObj, ok := p.appCoins[contractAddr]; ok {
		return contractObj.owner
	}
	return nil
}

// IterateControllerApps iterates all APP contracts created by controller contract.
// TODO: support config to filter by specific contract creator
func (p *Provider) IterateControllerApps(iterator func(common.Address) error, creator ...common.Address) error {
	for offset := int64(0); ; {
		appContractAddrs, total, err := p.controller.stub.ListApp(
			nil, big.NewInt(offset), big.NewInt(int64(listAppPageSize)),
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
	coin common.Address, iterator func(confEntry contract.AppConfigConfigEntry) error,
) error {
	contractObj, ok := p.appCoins[coin]
	if !ok {
		return errors.New("APP coin contract not found")
	}

	for offset := int64(0); ; {
		configEntries, total, err := contractObj.stub.ListResources(
			nil, big.NewInt(offset), big.NewInt(int64(listResourcePageSize)),
		)

		if err != nil {
			logrus.WithField("coin", coin).WithError(err).Debug("Failed to list resources for APP coin")
			return errors.WithMessage(err, "failed to list resources")
		}

		for i := range configEntries {
			if err := iterator(configEntries[i]); err != nil {
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
