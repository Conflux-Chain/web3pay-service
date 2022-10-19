package client

import (
	"context"
	"time"

	"github.com/Conflux-Chain/web3pay-service/client/jsonrpc"
	"github.com/Conflux-Chain/web3pay-service/contract"
	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/service"
	"github.com/Conflux-Chain/web3pay-service/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mcuadros/go-defaults"
	"github.com/openweb3/web3go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	jrpcMethodBilling      = "web3pay.Bill"
	jrpcMethodBillingBatch = "web3pay.BillBatch"
)

type VipInfo = contract.ICardTrackerVipInfo

type BillingClientConfig struct {
	Gateway    string        // API gateway endpoint
	BillingKey string        // billing auth key
	PingTest   bool          // test ping gateway?
	Timeout    time.Duration `default:"200ms"` // request timeout, default 200ms
}

// BillingClient client for billing
type BillingClient struct {
	*BillingClientConfig
	jrpcClient jsonrpc.RPCClient // JSON-RPC request client
}

func NewBillingClient(conf BillingClientConfig) (*BillingClient, error) {
	defaults.SetDefaults(&conf)

	rpcClient := jsonrpc.NewClientWithOpts(conf.Gateway,
		&jsonrpc.RPCClientOpts{
			Timeout: conf.Timeout,
			CustomHeaders: map[string]string{
				model.AuthHeaderBillingKey: conf.BillingKey,
			},
		},
	)

	if conf.PingTest {
		// test ping
		_, err := rpcClient.Call(context.Background(), jrpcMethodBilling)
		if err != nil {
			return nil, errors.WithMessage(err, "failed to dial payment gateway")
		}
	}

	client := &BillingClient{
		BillingClientConfig: &conf,
		jrpcClient:          rpcClient,
	}
	return client, nil
}

func (c *BillingClient) Bill(
	ctx context.Context, resourceId string, dryRun bool, apikey string) (*service.BillingReceipt, error) {

	ctx = jsonrpc.NewContextWithCustomHeaders(ctx, map[string]string{
		model.AuthHeaderApiKey: apikey,
	})

	args := &service.BillingRequest{
		DryRun:     dryRun,
		ResourceId: resourceId,
	}

	var receipt service.BillingReceipt
	if err := c.doCall(ctx, &receipt, jrpcMethodBilling, args); err != nil {
		return nil, err
	}

	return &receipt, nil
}

func (c *BillingClient) BillBatch(
	ctx context.Context, resourceUses map[string]int64, dryRun bool, apikey string) (*service.BillingBatchReceipt, error) {

	ctx = jsonrpc.NewContextWithCustomHeaders(ctx, map[string]string{
		model.AuthHeaderApiKey: apikey,
	})

	args := &service.BillingBatchRequest{
		DryRun:       dryRun,
		ResourceUses: resourceUses,
	}

	var receipt service.BillingBatchReceipt
	if err := c.doCall(ctx, &receipt, jrpcMethodBillingBatch, args); err != nil {
		return nil, err
	}

	return &receipt, nil
}

func (c *BillingClient) doCall(ctx context.Context, out interface{}, method string, args interface{}) error {
	var reply *model.BusinessError

	// call payment gateway
	if err := c.jrpcClient.CallFor(ctx, &reply, method, []interface{}{args}); err != nil {
		logrus.WithField("args", args).
			WithError(err).
			Debug("Web3Pay client failed to request payment gateway")
		return errors.WithMessage(err, "failed to request payment gateway")
	}

	// handle business error
	if reply.Code != model.ErrNil.Code {
		logrus.WithFields(logrus.Fields{
			"args":       args,
			"errCode":    reply.Code,
			"errMessage": reply.Message,
			"errData":    reply.Data,
		}).Debug("Web3Pay client encountered internal business error")
		return errors.WithMessage(reply, "internal business error")
	}

	return reply.GetObject(out)
}

type VipSubscriptionClientConfig struct {
	*web3go.ClientOption

	ChainRpcUrl          string        // blockchain network RPC endpoint
	AppContract          string        // App contract address
	VipInfoCacheSize     int           `default:"5000"` // VIP info cache size
	VipInfoExpirationTTL time.Duration `default:"15m"`  // VIP info cache expiration time
}

// VipSubscriptionClient client to get VIP subscription info
type VipSubscriptionClient struct {
	*web3go.Client
	*VipSubscriptionClientConfig

	// card tracker contract stub
	cardTracker *contract.CardTracker
	// VIP info cache, api key => VIP info
	vipInfoCache *util.ExpirableLruCache
}

func NewVipSubscriptionClient(conf VipSubscriptionClientConfig) (*VipSubscriptionClient, error) {
	defaults.SetDefaults(&conf)

	option := web3go.ClientOption{}
	if conf.ClientOption != nil {
		option = *conf.ClientOption
	}

	w3c, err := web3go.NewClientWithOption(conf.ChainRpcUrl, option)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to new `web3go` client")
	}

	clientForContract, _ := w3c.ToClientForContract()

	// App contract stub
	app := common.HexToAddress(conf.AppContract)
	appContract, err := contract.NewApp(app, clientForContract)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to init `App` contract")
	}

	// CardShop contract stub
	cardShop, err := appContract.CardShop(nil)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get `App` card shop")
	}

	cardShopContract, err := contract.NewCardShop(cardShop, clientForContract)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to init `App` card shop contract")
	}

	// CardTracker contract stub
	cardTracker, err := cardShopContract.Tracker(nil)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get `App` card tracker")
	}

	cardTrackerContract, err := contract.NewCardTracker(cardTracker, clientForContract)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to init `App` card tracker contract")
	}

	lruCache, _ := util.NewExpirableLruCache(conf.VipInfoCacheSize, conf.VipInfoExpirationTTL)
	return &VipSubscriptionClient{
		Client:                      w3c,
		VipSubscriptionClientConfig: &conf,
		cardTracker:                 cardTrackerContract,
		vipInfoCache:                lruCache,
	}, nil
}

func (c *VipSubscriptionClient) GetVipSubscriptionInfo(apiKey string) (*VipInfo, error) {
	account, err := util.GetAddrByApiKey(c.AppContract, apiKey)
	if err != nil {
		return nil, model.ErrAuth.WithData(err.Error())
	}

	if v, ok := c.vipInfoCache.Get(apiKey); ok { // hit in cache
		return v.(*VipInfo), nil
	}

	lockKey := util.MutexKey("vipsub/" + apiKey)
	util.KLock(lockKey)
	defer util.KUnlock(lockKey)

	if v, ok := c.vipInfoCache.Get(apiKey); ok { // double check
		return v.(*VipInfo), nil
	}

	vi, err := c.cardTracker.GetVipInfo(nil, account)
	if err != nil {
		return nil, err
	}

	c.vipInfoCache.Add(apiKey, &vi)
	return &vi, nil
}
