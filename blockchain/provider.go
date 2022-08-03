package blockchain

import (
	"math/big"
	"sync"

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

	// skip blocks ahead of latest block number to reduce chain reorg
	// while sync or state call.
	skipBlocksAheadOfLeatestBlock = 35
)

type contractBindCallContext struct {
	contractClient *web3go.ClientForContract
	signer         bind.SignerFn
}

// Provider provides blockchain operations.
type Provider struct {
	*client.RpcEthClient
	bindCallContext      *contractBindCallContext
	controller           *contract.Controller
	mutex                sync.Mutex
	appCoins             sync.Map // common.Address => *contract.APPCoin
	referenceBlockNumber int64    // reference block number for ops (eg., sync)
}

func MustNewProvider(w3c *web3go.Client) *Provider {
	latestBlockNumber, err := w3c.Eth.BlockNumber()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to get latest block number")
	}

	clientForContract, singerFn := w3c.ToClientForContract()

	// init controller contract stub
	ctrlCaller, err := contract.NewController(stdConf.controllerContractAddr, clientForContract)
	if err != nil {
		logrus.WithField("ctrlAddr", stdConf.controllerContractAddr).
			WithError(err).
			Fatal("Failed to initialize controller contract")
	}

	refBlockNum := latestBlockNumber.Int64() - 2*skipBlocksAheadOfLeatestBlock

	return &Provider{
		RpcEthClient: w3c.Eth,
		bindCallContext: &contractBindCallContext{
			contractClient: clientForContract,
			signer:         singerFn,
		},
		controller:           ctrlCaller,
		referenceBlockNumber: refBlockNum,
	}
}

// ReferenceBlockNumber returns reference block number.
func (p *Provider) ReferenceBlockNumber() int64 {
	return p.referenceBlockNumber
}

// GetAppCoinContract gets APP coin contract caller.
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

// GetAppCoinFrozenStatus gets APP coin frozen status for specific address.
func (p *Provider) GetAppCoinFrozenStatus(
	callOpts *bind.CallOpts, coin, address common.Address) (int64, error) {
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

	return fronzen.Int64(), nil
}

func (p *Provider) GetAppCoinBalanceAndFrozenStatus(
	callOpts *bind.CallOpts, coin, address common.Address) (int64, int64, error) {
	balance, err := p.GetAppCoinBalance(callOpts, coin, address)
	if err != nil {
		return 0, 0, errors.WithMessage(err, "failed to get balance")
	}

	frozen, err := p.GetAppCoinFrozenStatus(callOpts, coin, address)
	if err != nil {
		return 0, 0, errors.WithMessage(err, "failed to get frozen status")
	}

	return balance, frozen, err
}

// GetAppCoinBalance gets APP coin balance for specific address.
func (p *Provider) GetAppCoinBalance(
	callOpts *bind.CallOpts, coin, address common.Address) (int64, error) {
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

	return balance.Int64(), nil
}

// GetAppCoinContractOwner gets APP coin contract owner.
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

// ListTrackedAppCoins lists all tracked APP coin contracts.
func (p *Provider) ListTrackedAppCoins(callOpts *bind.CallOpts) ([]common.Address, error) {
	var appCoinAddrs []common.Address

	err := p.IterateTrackedAppCoins(callOpts, func(coin common.Address) error {
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

// IterateTrackedAppCoins iterates all tracked APP coin contracts.
func (p *Provider) IterateTrackedAppCoins(
	callOpts *bind.CallOpts, iterator func(common.Address) error) error {
	if stdConf.creatorAddr != nil {
		return p.iterateControllerAppCoins(callOpts, iterator, *stdConf.creatorAddr)
	}

	return p.iterateControllerAppCoins(callOpts, iterator)
}

// iterateControllerAppCoins iterates all APP coin contracts created by controller contracts filtered by creator.
func (p *Provider) iterateControllerAppCoins(
	callOpts *bind.CallOpts, iterator func(common.Address) error, creators ...common.Address) error {
	for offset := int64(0); ; {
		var appContractAddrs []common.Address
		var total *big.Int
		var listErr error

		if len(creators) > 0 {
			appInfos, err := p.controller.ListAppByCreator(
				callOpts, creators[0], uint32(offset), big.NewInt(int64(defaultListAppPageSize)),
			)

			if listErr = err; listErr == nil {
				total = appInfos.Total
				for i := range appInfos.Apps {
					appContractAddrs = append(appContractAddrs, appInfos.Apps[i].Addr)
				}
			}
		} else {
			appContractAddrs, total, listErr = p.controller.ListApp(
				callOpts, big.NewInt(offset), big.NewInt(int64(defaultListAppPageSize)),
			)
		}

		if listErr != nil {
			logrus.WithField("filterCreator", creators).
				WithError(listErr).
				Info("Failed to list APP contracts by controller")
			return errors.WithMessage(listErr, "failed to list APP contracts")
		}

		for i := range appContractAddrs {
			if err := iterator(appContractAddrs[i]); err != nil {
				logrus.WithFields(logrus.Fields{
					"filterCreator": creators, "appContractAddr": appContractAddrs[i],
				}).WithError(listErr).Info("Failed to iterate APP contracts by controller")
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
