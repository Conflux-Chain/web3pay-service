package middleware

import (
	"context"
	"net/http"

	"github.com/Conflux-Chain/web3pay-service/client"
	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/service"
	lru "github.com/hashicorp/golang-lru"
	"github.com/openweb3/go-rpc-provider"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	// billing status context key
	CtxKeyBillingStatus = CtxKey("Web3Pay-Billing-Status")
)

// BillingStatus billing result
type BillingStatus struct {
	Receipt   *service.BillingReceipt // billing receipt
	Error     error                   // billing error
	skipError bool                    // skip any error pretending to be successful
	apiKey    string                  // billed API key
}

// create billing status from error
func NewBillingStatusWithError(apiKey string, err error) *BillingStatus {
	return &BillingStatus{apiKey: apiKey, Error: err}
}

// create billing status from receipt
func NewBillingStatusWithReceipt(apiKey string, receipt *service.BillingReceipt) *BillingStatus {
	return &BillingStatus{apiKey: apiKey, Receipt: receipt}
}

// Success checks if billing successful
func (bs *BillingStatus) Success() bool {
	return bs.Error == nil || bs.skipError
}

// InternalServerError return internal server error as it is otherwise nil
func (bs *BillingStatus) InternalServerError() (error, bool) {
	if bs.Error == nil {
		return nil, false
	}

	bzerr, ok := bs.BusinessError()
	if !ok { // regarded as internal server error if not business error
		return bs.Error, true
	}

	if err := bzerr.(*model.BusinessError); err.Code == model.ErrInternalServer.Code {
		return bs.Error, true
	}

	return nil, false
}

// BusinessError returns business error as it is otherwise nil
func (bs *BillingStatus) BusinessError() (error, bool) {
	var bzerr *model.BusinessError
	if errors.As(bs.Error, &bzerr) {
		return bzerr, true
	}

	return nil, false
}

// BillingStatusFromContext returns billing status from context
func BillingStatusFromContext(ctx context.Context) (*BillingStatus, bool) {
	bs, ok := ctx.Value(CtxKeyBillingStatus).(*BillingStatus)
	return bs, ok
}

type BillingMiddlewareOption struct {
	// we3pay client to request billing gateway (mandatory)
	Client *client.BillingClient
	// provider to get API key from context, use `GetApiKeyFromContext` if not provided
	ApiKeyProvider authKeyProvider
	// whether to propagate internal server errors of the requested billing gateway
	// to the already billed users, default as false which will pretend nothing wrong
	// happened except some error logs will be printed. This is usually used to mitigate
	// side effects such as blocking paid user from access due to internal server errors.
	PropagateInternalServerError bool
}

func NewBillingMiddlewareOptionWithClient(client *client.BillingClient) *BillingMiddlewareOption {
	return (&BillingMiddlewareOption{Client: client}).InitDefault()
}

func (option *BillingMiddlewareOption) InitDefault() *BillingMiddlewareOption {
	if option.ApiKeyProvider == nil {
		option.ApiKeyProvider = GetApiKeyFromContext
	}
	return option
}

// providing openweb3 middleware

type Ow3ResourceIdMapper func(msg *rpc.JsonRpcMessage) string

type Ow3BillingMiddlewareOption struct {
	*BillingMiddlewareOption
	// gets resource ID from json rpc message
	ResourceIdMapper Ow3ResourceIdMapper
}

func NewOw3BillingMiddlewareOptionWithClient(client *client.BillingClient) *Ow3BillingMiddlewareOption {
	return &Ow3BillingMiddlewareOption{
		BillingMiddlewareOption: NewBillingMiddlewareOptionWithClient(client),
	}
}

// Openweb3BillingMiddleware provides billing RPC middleware for openweb3.
func Openweb3BillingMiddleware(option *Ow3BillingMiddlewareOption) Ow3Middleware {
	// cache billing API keys
	billingApiKeyCache, _ := lru.New(apiKeyCacheSize)

	return func(next rpc.HandleCallMsgFunc) rpc.HandleCallMsgFunc {
		wrapup := func(ctx context.Context, msg *rpc.JsonRpcMessage, bs *BillingStatus) *rpc.JsonRpcMessage {
			// inject billing status to context
			ctx = context.WithValue(ctx, CtxKeyBillingStatus, bs)

			if bs.Success() { // billing successfully?
				logrus.WithField("receipt", bs.Receipt).Debug("Billing middleware billed successfully")
				billingApiKeyCache.Add(bs.apiKey, struct{}{})
				return next(ctx, msg)
			}

			// handle gateway internal server error
			if err, ok := bs.InternalServerError(); ok {
				if !option.PropagateInternalServerError {
					_, bs.skipError = billingApiKeyCache.Get(bs.apiKey)
				}

				logrus.WithFields(logrus.Fields{
					"msg":                          msg,
					"skipError":                    bs.skipError,
					"propagateInternalServerError": option.PropagateInternalServerError,
				}).WithError(err).Info("Billing middleware internal server error")
			}

			return next(ctx, msg)
		}

		return func(ctx context.Context, msg *rpc.JsonRpcMessage) *rpc.JsonRpcMessage {
			apiKey, ok := option.ApiKeyProvider(ctx)
			if !ok || len(apiKey) == 0 { // api key provided?
				bs := NewBillingStatusWithError(apiKey, errInvalidApiKey)
				return wrapup(ctx, msg, bs)
			}

			// mapping resource ID
			resourceId := msg.Method
			if option.ResourceIdMapper != nil {
				resourceId = option.ResourceIdMapper(msg)
			}

			receipt, err := option.Client.Bill(context.Background(), resourceId, false, apiKey)
			if err != nil { // billing failed?
				err = errors.WithMessage(err, "web3pay billing failed")
				return wrapup(ctx, msg, NewBillingStatusWithError(apiKey, err))
			}

			return wrapup(ctx, msg, NewBillingStatusWithReceipt(apiKey, receipt))
		}
	}
}

// providing net/http middleware

type HttpResourceIdMapper func(req *http.Request) string

type HttpBillingMiddlewareOption struct {
	*BillingMiddlewareOption
	// gets resource ID from http request
	ResourceIdMapper HttpResourceIdMapper
}

func NewHttpBillingMiddlewareOptionWithClient(client *client.BillingClient) *HttpBillingMiddlewareOption {
	option := &HttpBillingMiddlewareOption{
		BillingMiddlewareOption: NewBillingMiddlewareOptionWithClient(client),
	}

	return option.InitDefault()
}

func (option *HttpBillingMiddlewareOption) InitDefault() *HttpBillingMiddlewareOption {
	option.BillingMiddlewareOption.InitDefault()

	if option.ResourceIdMapper == nil {
		option.ResourceIdMapper = func(r *http.Request) string {
			return r.URL.Query().Get("rid")
		}
	}

	return option
}

// HttpBillingMiddleware provides billing RPC middleware for net/http.
func HttpBillingMiddleware(option *HttpBillingMiddlewareOption) HttpMiddleware {
	// cache billing API keys
	billingApiKeyCache, _ := lru.New(apiKeyCacheSize)

	return func(next http.Handler) http.Handler {
		wrapup := func(w http.ResponseWriter, r *http.Request, bs *BillingStatus) {
			// inject billing status to context
			ctx := context.WithValue(r.Context(), CtxKeyBillingStatus, bs)
			r = r.WithContext(ctx)

			if bs.Error == nil { // billing successfull
				logrus.WithField("receipt", bs.Receipt).Debug("Billing middleware billed successfully")
				billingApiKeyCache.Add(bs.apiKey, struct{}{})
				next.ServeHTTP(w, r)
				return
			}

			// handle gateway internal server error
			if err, ok := bs.InternalServerError(); ok {
				if !option.PropagateInternalServerError {
					_, bs.skipError = billingApiKeyCache.Get(bs.apiKey)
				}

				logrus.WithFields(logrus.Fields{
					"request":                      r,
					"skipError":                    bs.skipError,
					"propagateInternalServerError": option.PropagateInternalServerError,
				}).WithError(err).Info("Billing middleware internal server error")
			}

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			apiKey, ok := option.ApiKeyProvider(ctx)
			if !ok || len(apiKey) == 0 { // API key provided?
				bs := NewBillingStatusWithError(apiKey, errInvalidApiKey)
				wrapup(w, r, bs)
				return
			}

			// mapping resource ID
			var resourceId string
			if option.ResourceIdMapper != nil {
				resourceId = option.ResourceIdMapper(r)
			}

			receipt, err := option.Client.Bill(context.Background(), resourceId, false, apiKey)
			if err != nil { // billing failed?
				err = errors.WithMessage(err, "web3pay billing failed")
				wrapup(w, r, NewBillingStatusWithError(apiKey, err))
				return
			}

			wrapup(w, r, NewBillingStatusWithReceipt(apiKey, receipt))
		})
	}
}
