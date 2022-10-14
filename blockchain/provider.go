package blockchain

import (
	"math/big"
	"sync"

	"github.com/Conflux-Chain/web3pay-service/contract"
	"github.com/Conflux-Chain/web3pay-service/util"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/client"
	"github.com/openweb3/web3go/signers"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	// default list page size
	defaultListAppPageSize      = 50
	defaultListResourcePageSize = 50

	// skip blocks ahead of latest block number to reduce chain reorg
	// while sync or state call.
	skipBlocksAheadOfLeatestBlock = 45
)

type contractBindCallContext struct {
	signerAddress  common.Address
	contractClient *web3go.ClientForContract
	signer         bind.SignerFn
}

// Provider provides blockchain operations.
type Provider struct {
	*client.RpcEthClient
	Config

	operatorAddr         common.Address
	bindCallContext      *contractBindCallContext
	appRegistry          *contract.AppRegistry
	mutex                sync.Mutex
	apps                 sync.Map // common.Address => *contract.APP
	apiWeightTokens      sync.Map // common.Address => *contract.ApiWeightToken
	referenceBlockNumber int64    // reference block number for ops (eg., sync)
}

func MustNewProvider(config *Config) *Provider {
	// parse operator address from private key
	operatorAddr, err := util.AddressFromEcdsaPrivateKeyString(config.OperatorPrivateKey)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to parse operator address")
	}

	// sign manager
	signerMgr := signers.MustNewSignerManagerByPrivateKeyStrings([]string{config.OperatorPrivateKey})
	signerAddr := signerMgr.List()[0].Address()

	// eth client
	w3c := util.MustNewEthClientFromViper(func(opt *web3go.ClientOption) {
		opt.SignerManager = signerMgr
	})

	latestBlockNumber, err := w3c.Eth.BlockNumber()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to get latest block number")
	}

	clientForContract, singerFn := w3c.ToClientForContract()

	// init AppRegistry contract stub
	appRegistryCaller, err := contract.NewAppRegistry(config.AppRegistryContractAddr, clientForContract)
	if err != nil {
		logrus.WithField("appRegistryAddr", config.AppRegistryContractAddr).
			WithError(err).
			Fatal("Failed to initialize `AppRegistry` contract")
	}

	refBlockNum := latestBlockNumber.Int64() - 2*skipBlocksAheadOfLeatestBlock

	return &Provider{
		RpcEthClient: w3c.Eth,
		Config:       *config,
		operatorAddr: operatorAddr,
		bindCallContext: &contractBindCallContext{
			signerAddress:  signerAddr,
			contractClient: clientForContract,
			signer:         singerFn,
		},
		appRegistry:          appRegistryCaller,
		referenceBlockNumber: refBlockNum,
	}
}

// OperatorAddress returns the operator address.
func (p *Provider) OperatorAddress() common.Address {
	return p.operatorAddr
}

// ReferenceBlockNumber returns reference block number.
func (p *Provider) ReferenceBlockNumber() int64 {
	return p.referenceBlockNumber
}

// GetAppContract gets `APP` contract stub.
func (p *Provider) GetAppContract(appAddr common.Address) (*contract.App, error) {
	if v, ok := p.apps.Load(appAddr); ok {
		return v.(*contract.App), nil
	}

	p.mutex.Lock()
	defer p.mutex.Unlock()

	if v, ok := p.apps.Load(appAddr); ok { // double check
		return v.(*contract.App), nil
	}

	appCaller, err := contract.NewApp(appAddr, p.bindCallContext.contractClient)
	if err != nil {
		logrus.WithField("appAddr", appAddr).WithError(err).Info("Failed to create APP contract stub")
		return nil, err
	}

	p.apps.Store(appAddr, appCaller)
	return appCaller, nil
}

// GetApiWeightTokenContract gets `ApiWeightToken` contract stub.
func (p *Provider) GetApiWeightTokenContract(awtAddr common.Address) (*contract.ApiWeightToken, error) {
	if v, ok := p.apiWeightTokens.Load(awtAddr); ok {
		return v.(*contract.ApiWeightToken), nil
	}

	p.mutex.Lock()
	defer p.mutex.Unlock()

	if v, ok := p.apiWeightTokens.Load(awtAddr); ok { // double check
		return v.(*contract.ApiWeightToken), nil
	}

	awtCaller, err := contract.NewApiWeightToken(awtAddr, p.bindCallContext.contractClient)
	if err != nil {
		logrus.WithField("apiWeightToken", awtAddr).WithError(err).Info("Failed to create `ApiWeightToken` contract stub")
		return nil, err
	}

	p.apiWeightTokens.Store(awtAddr, awtCaller)
	return awtCaller, nil
}

// GetCardShopContract gets `CardShop` contract stub.
func (p *Provider) GetCardShopContract(csAddr common.Address) (*contract.CardShop, error) {
	csCaller, err := contract.NewCardShop(csAddr, p.bindCallContext.contractClient)
	if err != nil {
		logrus.WithField("cardShop", csAddr).WithError(err).Info("Failed to create `CardShop` contract stub")
		return nil, err
	}

	return csCaller, nil
}

// BatchChargeAppBills batch charges APP bills.
func (p *Provider) BatchChargeAppBills(
	opts *bind.TransactOpts, app common.Address, requests []contract.IAppConfigChargeRequest) (*types.Transaction, error) {
	appContract, err := p.GetAppContract(app)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get APP contract")
	}

	// bind call context
	if opts == nil {
		opts = &bind.TransactOpts{}
	}

	opts.From = p.bindCallContext.signerAddress
	opts.Signer = p.bindCallContext.signer

	// estimate gas && gas price
	gasPrice, err := p.GasPrice()
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get gas price")
	}

	appAbi, err := contract.AppMetaData.GetAbi()
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get APP contract ABI")
	}

	data, err := appAbi.Pack(contract.MethodAppChargeBatch, requests)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to pack ABI data")
	}

	latestBlockNumberOrHash := types.BlockNumberOrHashWithNumber(types.LatestBlockNumber)
	gas, err := p.EstimateGas(types.CallRequest{
		From:     &opts.From,
		To:       &app,
		GasPrice: gasPrice,
		Data:     data,
	}, &latestBlockNumberOrHash)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to estimate gas")
	}

	opts.GasPrice = gasPrice
	opts.GasLimit = gas.Uint64()

	return appContract.ChargeBatch(opts, requests)
}

// FlushApiWeightTokenPendingConfig flushes `ApiWeightToken` pending configurations.
func (p *Provider) FlushApiWeightTokenPendingConfig(opts *bind.TransactOpts, awtAddr common.Address) (*types.Transaction, error) {
	awtContract, err := p.GetApiWeightTokenContract(awtAddr)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get `ApiWeightToken` contract")
	}

	// bind call context
	if opts == nil {
		opts = &bind.TransactOpts{}
	}

	opts.From = p.bindCallContext.signerAddress
	opts.Signer = p.bindCallContext.signer

	// estimate gas && gas price
	gasPrice, err := p.GasPrice()
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get gas price")
	}

	awtAbi, err := contract.ApiWeightTokenMetaData.GetAbi()
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get `ApiWeightToken` contract ABI")
	}

	data, err := awtAbi.Pack(contract.MethodApiWeightTokenFlushPendingConfig)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to pack ABI data")
	}

	latestBlockNumberOrHash := types.BlockNumberOrHashWithNumber(types.LatestBlockNumber)
	gas, err := p.EstimateGas(types.CallRequest{
		From:     &opts.From,
		To:       &awtAddr,
		GasPrice: gasPrice,
		Data:     data,
	}, &latestBlockNumberOrHash)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to estimate gas")
	}

	opts.GasPrice = gasPrice
	opts.GasLimit = gas.Uint64()

	return awtContract.FlushPendingConfig(opts)
}

// GetAppApiWeightToken gets `ApiWeightToken` for specified app.
func (p *Provider) GetAppApiWeightToken(
	callOpts *bind.CallOpts, app common.Address) (common.Address, error) {
	appContract, err := p.GetAppContract(app)
	if err != nil {
		return common.Address{}, errors.WithMessage(err, "failed to get APP contract")
	}

	awtAddr, err := appContract.GetApiWeightToken(callOpts)
	if err != nil {
		logrus.WithField("app", app).WithError(err).Info("Failed to get `ApiWeightToken` for APP")
		return common.Address{}, errors.WithMessage(err, "Failed to get `ApiWeightToken`")
	}

	return awtAddr, nil
}

// GetAppVipCoin gets `VipCoin` for specified app.
func (p *Provider) GetAppVipCoin(
	callOpts *bind.CallOpts, app common.Address) (common.Address, error) {
	appContract, err := p.GetAppContract(app)
	if err != nil {
		return common.Address{}, errors.WithMessage(err, "failed to get APP contract")
	}

	vcAddr, err := appContract.GetVipCoin(callOpts)
	if err != nil {
		logrus.WithField("app", app).WithError(err).Info("Failed to get `VipCoin` for APP")
		return common.Address{}, errors.WithMessage(err, "Failed to get `VipCoin`")
	}

	return vcAddr, nil
}

// GetAppAccountBalanceAndFrozenStatus combo-gets APP account balance and frozen status.
func (p *Provider) GetAppAccountBalanceAndFrozenStatus(
	callOpts *bind.CallOpts, app, address common.Address) (*big.Int, int64, error) {
	balance, err := p.GetAppAccountBalance(callOpts, app, address)
	if err != nil {
		return nil, 0, errors.WithMessage(err, "failed to get APP account balance")
	}

	frozen, err := p.GetAppAccountFrozenStatus(callOpts, app, address)
	if err != nil {
		return nil, 0, errors.WithMessage(err, "failed to get APP account frozen status")
	}

	return balance, frozen, err
}

// GetAppAccountFrozenStatus gets APP account frozen status for specific address.
func (p *Provider) GetAppAccountFrozenStatus(
	callOpts *bind.CallOpts, app, address common.Address) (int64, error) {
	appContract, err := p.GetAppContract(app)
	if err != nil {
		return 0, errors.WithMessage(err, "failed to get APP contract")
	}

	fronzen, err := appContract.WithdrawSchedules(callOpts, address)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"app": app, "address": address,
		}).WithError(err).Info("Failed to get APP account frozen status")

		return 0, errors.WithMessage(err, "failed to get APP account frozen status")
	}

	return fronzen.Int64(), nil
}

// GetAppBalance gets APP account balance for specific address.
func (p *Provider) GetAppAccountBalance(
	callOpts *bind.CallOpts, app, address common.Address) (*big.Int, error) {
	appContract, err := p.GetAppContract(app)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get APP contract")
	}

	depositAmount, airdropAmount, err := appContract.BalanceOf(callOpts, address)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"app":     app,
			"address": address,
		}).WithError(err).Info("Failed to get APP balance")

		return nil, errors.WithMessage(err, "failed to get APP balance")
	}

	return big.NewInt(0).Add(depositAmount, airdropAmount), nil
}

// GetApiWeightTokenPendingSeconds gets pending delay seconds for `ApiWeightToken` resources.
func (p *Provider) GetApiWeightTokenPendingSeconds(callOpts *bind.CallOpts, awtAddr common.Address) (*big.Int, error) {
	awtContract, err := p.GetApiWeightTokenContract(awtAddr)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get `ApiWeightToken` contract")
	}

	return awtContract.PendingSeconds(callOpts)
}

// ListTrackedApps lists all tracked APP contracts.
func (p *Provider) ListTrackedApps(callOpts *bind.CallOpts) (apps, apiWeightTokens, vipCoins []common.Address, err error) {
	err = p.IterateTrackedApps(callOpts, func(app common.Address) error {
		awt, err := p.GetAppApiWeightToken(callOpts, app)
		if err != nil {
			return errors.WithMessage(err, "failed to get App `ApiWeightToken`")
		}

		vc, err := p.GetAppVipCoin(callOpts, app)
		if err != nil {
			return errors.WithMessage(err, "failed to get App `VipCoin`")
		}

		apps = append(apps, app)
		apiWeightTokens = append(apiWeightTokens, awt)
		vipCoins = append(vipCoins, vc)

		return nil
	})

	return
}

// IterateTrackedApps iterates all tracked APP contracts.
func (p *Provider) IterateTrackedApps(
	callOpts *bind.CallOpts, iterator func(common.Address) error) error {
	if p.OwnerAddr != nil {
		return p.iterateAppRegistryApps(callOpts, iterator, *p.OwnerAddr)
	}

	return p.iterateAppRegistryApps(callOpts, iterator)
}

// iterateAppRegistryApps iterates all APP contracts created within registry filtered by owner.
func (p *Provider) iterateAppRegistryApps(
	callOpts *bind.CallOpts, iterator func(common.Address) error, owners ...common.Address) error {
	for offset := int64(0); ; {
		var appInfos []contract.AppRegistryAppInfo
		var total *big.Int
		var listErr error

		if len(owners) > 0 {
			total, appInfos, listErr = p.appRegistry.ListByOwner(
				callOpts, owners[0], big.NewInt(offset), big.NewInt(defaultListAppPageSize),
			)
		} else {
			total, appInfos, listErr = p.appRegistry.List(
				callOpts, big.NewInt(offset), big.NewInt(int64(defaultListAppPageSize)),
			)
		}

		if listErr != nil {
			logrus.WithField("filterOwner", owners).WithError(listErr).Info("Failed to list APPs within registry")
			return errors.WithMessage(listErr, "failed to list APP contracts")
		}

		for i := range appInfos {
			if err := iterator(appInfos[i].Addr); err != nil {
				logrus.WithFields(logrus.Fields{
					"filterOwners":    owners,
					"appContractAddr": appInfos[i],
				}).WithError(listErr).Info("Failed to iterate APP contracts within registry")
				return errors.WithMessage(err, "failed to iterate APP contract")
			}
		}

		offset += int64(len(appInfos))
		if offset >= total.Int64() { // all done
			break
		}
	}

	return nil
}

// GetApiWeightTokenResources gets `ApiWeightToken` config resources.
func (p *Provider) GetApiWeightTokenResources(
	callOpts *bind.CallOpts, awtAddr common.Address) (map[string]contract.IAppConfigConfigEntry, error) {
	appResources := make(map[string]contract.IAppConfigConfigEntry)

	// iterate all resources under specified `ApiWeightToken` contract
	err := p.IterateApiWeightTokenResources(callOpts, awtAddr, func(confEntry contract.IAppConfigConfigEntry) (bool, error) {
		appResources[confEntry.ResourceId] = confEntry
		return false, nil
	})

	if err != nil {
		logrus.WithField("apiWeightToken", awtAddr).WithError(err).Info("Failed to get config resources")
		return nil, errors.WithMessage(err, "failed to get config resources")
	}

	return appResources, nil
}

// IterateApiWeightTokenResources iterates all resources under specified `ApiWeightToken`
func (p *Provider) IterateApiWeightTokenResources(
	callOpts *bind.CallOpts, awtAddr common.Address, iterator func(confEntry contract.IAppConfigConfigEntry) (bool, error),
) error {
	awtContract, err := p.GetApiWeightTokenContract(awtAddr)
	if err != nil {
		return errors.WithMessage(err, "failed to get `ApiWeightToken` contract")
	}

	for offset := int64(0); ; {
		configEntries, total, err := awtContract.ListResources(
			callOpts, big.NewInt(offset), big.NewInt(int64(defaultListResourcePageSize)),
		)

		if err != nil {
			logrus.WithField("apiWeightToken", awtAddr).WithError(err).Info("Failed to list APP config resources")
			return errors.WithMessage(err, "failed to list config resources")
		}

		for i := range configEntries {
			interrupt, err := iterator(configEntries[i])
			if err != nil {
				logrus.WithField("apiWeightToken", awtAddr).WithError(err).Info("Failed to iterate APP config resource")
				return errors.WithMessage(err, "failed to iterate config resource")
			}

			if interrupt {
				return nil
			}
		}

		offset += int64(len(configEntries))
		if offset >= total.Int64() { // all done
			break
		}
	}

	return nil
}
