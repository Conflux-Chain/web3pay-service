package blockchain

import (
	"math/big"
	"sync"

	"github.com/Conflux-Chain/go-conflux-util/viper"
	"github.com/Conflux-Chain/web3pay-service/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/client"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	// default list page size
	defaultListAppPageSize      = 50
	defaultListResourcePageSize = 50
)

type contractBindCallContext struct {
	contractClient *web3go.ClientForContract
	signer         bind.SignerFn
}

// Provider provides blockchain operations.
type Provider struct {
	*client.RpcEthClient

	config               *Config
	bindCallContext      *contractBindCallContext
	controller           *contract.Controller
	mutex                sync.Mutex
	appCoins             sync.Map // common.Address => *contract.APPCoin
	referenceBlockNumber int64    // reference block number for ops (eg., sync)
}

func MustNewProviderFromViper(w3c *web3go.Client) *Provider {
	var conf Config
	viper.MustUnmarshalKey("blockchain", &conf)

	latestBlockNumber, err := w3c.Eth.BlockNumber()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to get latest block number")
	}

	clientForContract, singerFn := w3c.ToClientForContract()
	ctrlAddr := common.HexToAddress(conf.ControllerContractAddr)

	// init controller contract stub
	ctrlCaller, err := contract.NewController(ctrlAddr, clientForContract)
	if err != nil {
		logrus.WithField("config", conf).
			WithError(err).
			Fatal("Failed to initialize controller contract")
	}

	refBlockNum := latestBlockNumber.Int64() - 2*skipBlocksNearHeadOfLatest

	return &Provider{
		RpcEthClient: w3c.Eth,
		config:       &conf,
		bindCallContext: &contractBindCallContext{
			contractClient: clientForContract,
			signer:         singerFn,
		},
		controller:           ctrlCaller,
		referenceBlockNumber: refBlockNum,
	}
}

func (p *Provider) ControllerAddress() common.Address {
	return common.HexToAddress(p.config.ControllerContractAddr)
}

func (p *Provider) ControllerContract() *contract.Controller {
	return p.controller
}

func (p *Provider) ReferenceBlockNumber() int64 {
	return p.referenceBlockNumber
}

func (p *Provider) GetAppCoinContract(appCoinAddr common.Address) (*contract.APPCoin, error) {
	if v, ok := p.appCoins.Load(appCoinAddr); ok {
		return v.(*contract.APPCoin), nil
	}

	p.mutex.Lock()
	defer p.mutex.Unlock()

	if v, ok := p.appCoins.Load(appCoinAddr); ok { // double check
		return v.(*contract.APPCoin), nil
	}

	appCoinCaller, err := contract.NewAPPCoin(appCoinAddr, p.bindCallContext.contractClient)
	if err != nil {
		logrus.WithField("appCoinAddr", appCoinAddr).
			WithError(err).
			Info("Failed to create APP coin contract")

		return nil, err
	}

	p.appCoins.Store(appCoinAddr, appCoinCaller)
	return appCoinCaller, nil
}

func (p *Provider) GetAppCoinFrozenStatus(
	callOpts *bind.CallOpts, coin, address common.Address) (uint64, error) {
	appCoinContract, err := p.GetAppCoinContract(coin)
	if err != nil {
		return 0, errors.WithMessage(err, "failed to get APP coin contract")
	}

	fronzen, err := appCoinContract.FrozenMap(callOpts, address)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"coin": coin, "address": address,
		}).WithError(err).Info("Failed to get APP coin fronze status")

		return 0, errors.WithMessage(err, "failed to get APP coin fronze status")
	}

	return fronzen.Uint64(), nil
}

func (p *Provider) GetAppCoinBalance(
	callOpts *bind.CallOpts, coin, address common.Address) (uint64, error) {
	appCoinContract, err := p.GetAppCoinContract(coin)
	if err != nil {
		return 0, errors.WithMessage(err, "failed to get APP coin contract")
	}

	balance, err := appCoinContract.BalanceOf(callOpts, address)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"coin": coin, "address": address,
		}).WithError(err).Info("Failed to get APP coin balance")

		return 0, errors.WithMessage(err, "failed to get APP coin balance")
	}

	return balance.Uint64(), nil
}

// GetAppCoinContractOwner gets concerned APP coin contract owner.
func (p *Provider) GetAppCoinContractOwner(
	callOpts *bind.CallOpts, coin common.Address) (*common.Address, error) {
	appCoinContract, err := p.GetAppCoinContract(coin)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get APP coin contract")
	}

	// fetch APP coin contract owner
	appCoinOwner, err := appCoinContract.AppOwner(callOpts)
	if err != nil {
		logrus.WithField("appCoinAddr", coin).
			WithError(err).
			Info("Failed to get APP coin contract owner")

		return nil, errors.WithMessage(err, "failed to get APP coin contract owner")
	}

	return &appCoinOwner, nil
}

// ListControllerAppCoins lists addresses of all APP coins created by the controller contract.
// TODO: support config to filter by specific contract creator.
func (p *Provider) ListControllerAppCoins(
	callOpts *bind.CallOpts, creator ...common.Address) ([]common.Address, error) {
	var appCoinAddrs []common.Address

	err := p.IterateControllerApps(callOpts, func(coin common.Address) error {
		appCoinAddrs = append(appCoinAddrs, coin)
		return nil
	})

	return appCoinAddrs, err
}

// GetAppCoinResources gets concerned APP coin resources with specified id.
func (p *Provider) GetAppCoinResources(
	callOpts *bind.CallOpts, coin common.Address) (map[string]contract.AppConfigConfigEntry, error) {
	appResources := make(map[string]contract.AppConfigConfigEntry)

	// iterate all resources under APP coin
	err := p.IterateAppCoinResources(callOpts, coin, func(confEntry contract.AppConfigConfigEntry) error {
		appResources[confEntry.ResourceId] = confEntry
		return nil
	})

	if err != nil {
		logrus.WithField("appCoin", coin).
			WithError(err).
			Info("Failed to get APP coin resources")

		return nil, errors.WithMessage(err, "failed to get APP coin resources")
	}

	return appResources, nil
}

// IterateControllerApps iterates all APP contracts created by controller contract.
// TODO: support config to filter by specific contract creator.
func (p *Provider) IterateControllerApps(
	callOpts *bind.CallOpts, iterator func(common.Address) error, creator ...common.Address) error {
	for offset := int64(0); ; {
		appContractAddrs, total, err := p.controller.ListApp(
			callOpts, big.NewInt(offset), big.NewInt(int64(defaultListAppPageSize)),
		)

		if err != nil {
			logrus.WithError(err).Info("Failed to list APP contracts by controller")
			return errors.WithMessage(err, "failed to list APP contracts")
		}

		for i := range appContractAddrs {
			if err := iterator(appContractAddrs[i]); err != nil {
				logrus.WithError(err).Info("Failed to iterate APP contracts by controller")
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
	callOpts *bind.CallOpts, coin common.Address, iterator func(confEntry contract.AppConfigConfigEntry) error,
) error {
	appCoinContract, err := p.GetAppCoinContract(coin)
	if err != nil {
		return errors.WithMessage(err, "failed to get APP coin contract")
	}

	for offset := int64(0); ; {
		configEntries, total, err := appCoinContract.ListResources(
			callOpts, big.NewInt(offset), big.NewInt(int64(defaultListResourcePageSize)),
		)

		if err != nil {
			logrus.WithField("coin", coin).
				WithError(err).
				Info("Failed to list resources for APP coin")
			return errors.WithMessage(err, "failed to list resources")
		}

		for i := range configEntries {
			if err := iterator(configEntries[i]); err != nil {
				logrus.WithField("coin", coin).
					WithError(err).
					Info("Failed to iterate APP coin resource")
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
