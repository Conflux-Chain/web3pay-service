package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/Conflux-Chain/web3pay-service/contract"
	"github.com/Conflux-Chain/web3pay-service/util"
	"github.com/ethereum/go-ethereum/common"
	lru "github.com/hashicorp/golang-lru"
	"github.com/openweb3/go-rpc-provider"
	"github.com/openweb3/web3go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	// VIP subscription status context key
	CtxKeySubscriptionStatus = CtxKey("Web3Pay-Subscription-Status")

	// VIP subscription info cache expiration TTL
	vipInfoExpirationTTL = 15 * time.Minute
)

var (
	chainRpcError = errors.New("blockchain RPC error")
)

// VipInfo VIP subscription information
type VipInfo = contract.ICardTrackerVipInfo

// VipSubscriptionClient client to get VIP subscription info
type VipSubscriptionClient struct {
	*web3go.Client

	// App contract address
	app common.Address
	// card tracker contract stub
	cardTracker *contract.CardTracker
	// VIP info cache
	vipInfoCache *util.ExpirableLruCache
}

func NewVipSubscriptionClient(w3c *web3go.Client, app common.Address) (*VipSubscriptionClient, error) {
	clientForContract, _ := w3c.ToClientForContract()

	// App contract stub
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

	lruCache, _ := util.NewExpirableLruCache(apiKeyCacheSize, vipInfoExpirationTTL)
	return &VipSubscriptionClient{
		Client: w3c, app: app,
		cardTracker:  cardTrackerContract,
		vipInfoCache: lruCache,
	}, nil
}

func (c *VipSubscriptionClient) GetVipSubscriptionInfo(apiKey string) (*VipInfo, error) {
	account, err := util.GetAddrByApiKey(c.app.String(), apiKey)
	if err != nil {
		return nil, errors.WithMessage(errInvalidApiKey, err.Error())
	}

	if v, ok := c.vipInfoCache.Get(apiKey); ok { // hit in cache
		return v.(*VipInfo), nil
	}

	lockKey := util.MutexKey(apiKey)
	util.KLock(lockKey)
	defer util.KUnlock(lockKey)

	if v, ok := c.vipInfoCache.Get(apiKey); ok { // double check
		return v.(*VipInfo), nil
	}

	vi, err := c.cardTracker.GetVipInfo(nil, account)
	if err != nil {
		return nil, errors.WithMessage(chainRpcError, err.Error())
	}

	c.vipInfoCache.Add(apiKey, &vi)
	return &vi, err
}

// VipSubscriptionStatus VIP subscription status
type VipSubscriptionStatus struct {
	VipInfo   *VipInfo // VIP info
	Error     error    // subscription error
	skipError bool     // skip any error pretending to be successful
	apiKey    string   // subscription API key
}

// create subscription status from error
func NewVipSubscriptionStatusWithError(apiKey string, err error) *VipSubscriptionStatus {
	return &VipSubscriptionStatus{apiKey: apiKey, Error: err}
}

// create subscription status from subscription VIP info
func NewVipSubscriptionStatusWithInfo(apiKey string, vi *VipInfo) *VipSubscriptionStatus {
	return &VipSubscriptionStatus{apiKey: apiKey, VipInfo: vi}
}

func (ss *VipSubscriptionStatus) GetVipInfo() (*VipInfo, error) {
	if ss.Error != nil && !ss.skipError {
		return nil, ss.Error
	}

	return ss.VipInfo, nil
}

type VipSubscriptionMiddlewareOption struct {
	// client to request VIP subscription info from the blockchain network (mandatory)
	Client *VipSubscriptionClient
	// provider to get API key from context, use `GetApiKeyFromContext` if not provided
	ApiKeyProvider authKeyProvider
	// whether to propagate RPC errors of the requested blockchain network endpoint
	// to the already subscribed users, default as false which will pretend nothing wrong
	// happened except some error logs will be printed. This is usually used to mitigate
	// side effects such as blocking paid user from access due to internal server errors.
	PropagateChainRpcError bool
}

func NewVipSubscriptionMiddlewareOptionWithClient(client *VipSubscriptionClient) *VipSubscriptionMiddlewareOption {
	return (&VipSubscriptionMiddlewareOption{Client: client}).InitDefault()
}

func (option *VipSubscriptionMiddlewareOption) InitDefault() *VipSubscriptionMiddlewareOption {
	if option.ApiKeyProvider == nil {
		option.ApiKeyProvider = GetApiKeyFromContext
	}
	return option
}

// Openweb3VipSubscriptionMiddleware provides VIP subscription RPC middleware for openweb3.
func Openweb3VipSubscriptionMiddleware(option *VipSubscriptionMiddlewareOption) Ow3Middleware {
	// cache subscription VIP API keys
	vipApiKeyCache, _ := lru.New(apiKeyCacheSize)

	return func(next rpc.HandleCallMsgFunc) rpc.HandleCallMsgFunc {
		wrapup := func(ctx context.Context, msg *rpc.JsonRpcMessage, ss *VipSubscriptionStatus) *rpc.JsonRpcMessage {
			// inject subscription status to context
			ctx = context.WithValue(ctx, CtxKeySubscriptionStatus, ss)

			if ss.Error == nil {
				logrus.WithFields(logrus.Fields{
					"apiKey":  ss.apiKey,
					"vipInfo": ss.VipInfo,
				}).Debug("VIP subscription middleware called successfully")
				vipApiKeyCache.Add(ss.apiKey, ss.VipInfo)
				return next(ctx, msg)
			}

			// handle blockchain RPC error
			if err := ss.Error; errors.Is(err, chainRpcError) {
				if !option.PropagateChainRpcError {
					if v, ok := vipApiKeyCache.Get(ss.apiKey); ok {
						ss.VipInfo = v.(*VipInfo)
						ss.skipError = true
					}
				}

				logrus.WithFields(logrus.Fields{
					"msg":                    msg,
					"skipError":              ss.skipError,
					"PropagateChainRpcError": option.PropagateChainRpcError,
				}).WithError(err).Error("VIP subscription middleware chain RPC error")
			}

			return next(ctx, msg)
		}

		return func(ctx context.Context, msg *rpc.JsonRpcMessage) *rpc.JsonRpcMessage {
			apiKey, ok := option.ApiKeyProvider(ctx)
			if !ok || len(apiKey) == 0 { // api key provided?
				ss := NewVipSubscriptionStatusWithError(apiKey, errInvalidApiKey)
				return wrapup(ctx, msg, ss)
			}

			vi, err := option.Client.GetVipSubscriptionInfo(apiKey)
			if err != nil {
				ss := NewVipSubscriptionStatusWithError(apiKey, err)
				return wrapup(ctx, msg, ss)
			}

			return wrapup(ctx, msg, NewVipSubscriptionStatusWithInfo(apiKey, vi))
		}
	}
}

// HttpVipSubscriptionMiddleware provides VIP subscription RPC middleware for net/http.
func HttpVipSubscriptionMiddleware(option *VipSubscriptionMiddlewareOption) HttpMiddleware {
	// cache subscription VIP API keys
	vipApiKeyCache, _ := lru.New(apiKeyCacheSize)

	return func(next http.Handler) http.Handler {
		wrapup := func(w http.ResponseWriter, r *http.Request, ss *VipSubscriptionStatus) {
			// inject subscription status to context
			ctx := context.WithValue(r.Context(), CtxKeySubscriptionStatus, ss)
			r = r.WithContext(ctx)

			if ss.Error == nil {
				logrus.WithFields(logrus.Fields{
					"apiKey":  ss.apiKey,
					"vipInfo": ss.VipInfo,
				}).Debug("VIP subscription middleware called successfully")
				vipApiKeyCache.Add(ss.apiKey, ss.VipInfo)
				next.ServeHTTP(w, r)
				return
			}

			// handle blockchain RPC error
			if err := ss.Error; errors.Is(err, chainRpcError) {
				if !option.PropagateChainRpcError {
					if v, ok := vipApiKeyCache.Get(ss.apiKey); ok {
						ss.VipInfo = v.(*VipInfo)
						ss.skipError = true
					}
				}

				logrus.WithFields(logrus.Fields{
					"request":                r,
					"skipError":              ss.skipError,
					"PropagateChainRpcError": option.PropagateChainRpcError,
				}).WithError(err).Error("VIP subscription middleware chain RPC error")
			}

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			apiKey, ok := option.ApiKeyProvider(ctx)
			if !ok || len(apiKey) == 0 { // api key provided?
				ss := NewVipSubscriptionStatusWithError(apiKey, errInvalidApiKey)
				wrapup(w, r, ss)
				return
			}

			vi, err := option.Client.GetVipSubscriptionInfo(apiKey)
			if err != nil {
				ss := NewVipSubscriptionStatusWithError(apiKey, err)
				wrapup(w, r, ss)
				return
			}

			wrapup(w, r, NewVipSubscriptionStatusWithInfo(apiKey, vi))
		})
	}
}
