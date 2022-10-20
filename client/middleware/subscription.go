package middleware

import (
	"context"
	"net/http"

	"github.com/Conflux-Chain/web3pay-service/client"
	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/types"
	lru "github.com/hashicorp/golang-lru"
	"github.com/openweb3/go-rpc-provider"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	// VIP subscription status context key
	CtxKeyVipSubscriptionStatus = CtxKey("Web3Pay-Vip-Subscription-Status")
)

// VipSubscriptionStatus VIP subscription status
type VipSubscriptionStatus struct {
	VipInfo   *types.VipInfo // VIP info
	Error     error          // subscription error
	skipError bool           // skip any error pretending to be successful
	apiKey    string         // subscription API key
}

// create subscription status from error
func NewVipSubscriptionStatusWithError(apiKey string, err error) *VipSubscriptionStatus {
	return &VipSubscriptionStatus{apiKey: apiKey, Error: err}
}

// create subscription status from subscription VIP info
func NewVipSubscriptionStatusWithInfo(apiKey string, vi *types.VipInfo) *VipSubscriptionStatus {
	return &VipSubscriptionStatus{apiKey: apiKey, VipInfo: vi}
}

func (ss *VipSubscriptionStatus) GetVipInfo() (*types.VipInfo, error) {
	if ss.Error != nil && !ss.skipError {
		return nil, ss.Error
	}

	return ss.VipInfo, nil
}

// BusinessError returns business error as it is otherwise nil
func (bs *VipSubscriptionStatus) BusinessError() (*model.BusinessError, bool) {
	var bzerr *model.BusinessError
	if errors.As(bs.Error, &bzerr) {
		return bzerr, true
	}

	return nil, false
}

// VipSubscriptionStatusFromContext returns VIP subscription status from context
func VipSubscriptionStatusFromContext(ctx context.Context) (*VipSubscriptionStatus, bool) {
	ss, ok := ctx.Value(CtxKeyVipSubscriptionStatus).(*VipSubscriptionStatus)
	return ss, ok
}

type VipSubscriptionMiddlewareOption struct {
	// client to request VIP subscription info from the blockchain network (mandatory)
	Client *client.VipSubscriptionClient
	// provider to get API key from context, use `GetApiKeyFromContext` if not provided
	ApiKeyProvider authKeyProvider
	// whether to propagate non-business error to the already subscribed users, default as
	// false which will pretend nothing wrong happened except some error logs will be printed.
	// This is usually used to mitigate side effects such as blocking paid user from access
	// due to some server errors (such as timeout etc.,)
	PropagateNonBusinessError bool
}

func NewVipSubscriptionMiddlewareOptionWithClient(client *client.VipSubscriptionClient) *VipSubscriptionMiddlewareOption {
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
			ctx = context.WithValue(ctx, CtxKeyVipSubscriptionStatus, ss)

			if ss.Error == nil {
				logrus.WithFields(logrus.Fields{
					"apiKey":  ss.apiKey,
					"vipInfo": ss.VipInfo,
				}).Debug("VIP subscription middleware called successfully")
				vipApiKeyCache.Add(ss.apiKey, ss.VipInfo)
				return next(ctx, msg)
			}

			// handle non business error
			if err, ok := ss.BusinessError(); !ok {
				if !option.PropagateNonBusinessError {
					if v, ok := vipApiKeyCache.Get(ss.apiKey); ok {
						ss.VipInfo = v.(*types.VipInfo)
						ss.skipError = true
					}
				}

				logrus.WithFields(logrus.Fields{
					"msg":                       msg,
					"skipError":                 ss.skipError,
					"propagateNonBusinessError": option.PropagateNonBusinessError,
				}).WithError(err).Error("VIP subscription middleware non-business error")
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
			ctx := context.WithValue(r.Context(), CtxKeyVipSubscriptionStatus, ss)
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

			// handle non business error
			if err, ok := ss.BusinessError(); !ok {
				if !option.PropagateNonBusinessError {
					if v, ok := vipApiKeyCache.Get(ss.apiKey); ok {
						ss.VipInfo = v.(*types.VipInfo)
						ss.skipError = true
					}
				}

				logrus.WithFields(logrus.Fields{
					"request":                   r,
					"skipError":                 ss.skipError,
					"propagateNonBusinessError": option.PropagateNonBusinessError,
				}).WithError(err).Error("VIP subscription middleware non-business error")
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
