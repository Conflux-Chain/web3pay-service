package client

import (
	"context"
	"net/http"

	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/service"
	lru "github.com/hashicorp/golang-lru"
	"github.com/openweb3/go-rpc-provider"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type CtxKey string

const (
	// context key
	CtxKeyBillingStatus = CtxKey("Web3Pay-Billing-Status")
	CtxCustomerKey      = CtxKey("Web3Pay-Customer-Key")

	// default customer key LRU cache size
	customerKeyCacheSize = 2000
)

var (
	// errors
	errCustomerKeyNotProvided = model.ErrAuth.WithData("customer key not provided")

	// cache VIP customer keys
	vipCustomerKeyCache, _ = lru.New(customerKeyCacheSize)
)

// BillingStatus billing result
type BillingStatus struct {
	Receipt     *service.BillingReceipt // billing receipt
	Error       error                   // billing error
	skipError   bool                    // skip any error pretending to be successful
	customerKey string                  // billed customer key
}

// create billing status from error
func NewBillingStatusWithError(customerKey string, err error) *BillingStatus {
	return &BillingStatus{customerKey: customerKey, Error: err}
}

// create billing status from receipt
func NewBillingStatusWithReceipt(customerKey string, receipt *service.BillingReceipt) *BillingStatus {
	return &BillingStatus{customerKey: customerKey, Receipt: receipt}
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
	if !ok || bzerr.Code == model.ErrInternalServer.Code {
		return bs.Error, true
	}

	return nil, false
}

// BusinessError returns business error as it is otherwise nil
func (bs *BillingStatus) BusinessError() (*model.BusinessError, bool) {
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

// authKeyProvider retrieves auth key from context
type authKeyProvider func(ctx context.Context) (string, bool)

type BillingMiddlewareOption struct {
	// we3pay client to request billing gateway (mandatory)
	Client *Client
	// provider to get customer key from context, use `GetCustomerKeyFromContext` if not provided
	CustomerKeyProvider authKeyProvider
	// whether to propagate internal server errors of the requested billing gateway
	// to the already billed users, default as false which will pretend nothing wrong
	// happened except some error logs will be printed. This is usually used to mitigate
	// side effects such as blocking paid user from access due to internal server errors.
	PropagateInternalServerError bool
}

func NewBillingMiddlewareOptionWithClient(client *Client) *BillingMiddlewareOption {
	return (&BillingMiddlewareOption{Client: client}).InitDefault()
}

func (option *BillingMiddlewareOption) InitDefault() *BillingMiddlewareOption {
	if option.CustomerKeyProvider == nil {
		option.CustomerKeyProvider = GetCustomerKeyFromContext
	}
	return option
}

// providing openweb3 middleware

type Ow3Middleware = rpc.HandleCallMsgMiddleware
type Ow3ResourceIdMapper func(msg *rpc.JsonRpcMessage) string

type Ow3BillingMiddlewareOption struct {
	*BillingMiddlewareOption
	// gets resource ID from json rpc message
	ResourceIdMapper Ow3ResourceIdMapper
}

func NewOw3BillingMiddlewareOptionWithClient(client *Client) *Ow3BillingMiddlewareOption {
	return &Ow3BillingMiddlewareOption{
		BillingMiddlewareOption: NewBillingMiddlewareOptionWithClient(client),
	}
}

// Openweb3BillingMiddleware provides billing RPC middleware for openweb3.
func Openweb3BillingMiddleware(option *Ow3BillingMiddlewareOption) Ow3Middleware {
	return func(next rpc.HandleCallMsgFunc) rpc.HandleCallMsgFunc {
		wrapup := func(ctx context.Context, msg *rpc.JsonRpcMessage, bs *BillingStatus) *rpc.JsonRpcMessage {
			// inject billing status to context
			ctx = context.WithValue(ctx, CtxKeyBillingStatus, bs)

			if bs.Error == nil { // billing successfully?
				logrus.WithField("receipt", bs.Receipt).Debug("Billing middleware billed successfully")
				vipCustomerKeyCache.Add(bs.customerKey, struct{}{})
				return next(ctx, msg)
			}

			// handle gateway internal server error
			if err, ok := bs.InternalServerError(); ok {
				if !option.PropagateInternalServerError {
					_, bs.skipError = vipCustomerKeyCache.Get(bs.customerKey)
				}

				logrus.WithFields(logrus.Fields{
					"msg": msg, "skipError": bs.skipError,
					"propagateInternalServerError": option.PropagateInternalServerError,
				}).WithError(err).Error("Billing middleware internal server error")
			}

			return next(ctx, msg)
		}

		return func(ctx context.Context, msg *rpc.JsonRpcMessage) *rpc.JsonRpcMessage {
			customerKey, ok := option.CustomerKeyProvider(ctx)
			if !ok || len(customerKey) == 0 { // customer key provided?
				bs := NewBillingStatusWithError(customerKey, errCustomerKeyNotProvided)
				return wrapup(ctx, msg, bs)
			}

			// mapping resource ID
			resourceId := msg.Method
			if option.ResourceIdMapper != nil {
				resourceId = option.ResourceIdMapper(msg)
			}

			receipt, err := option.Client.Bill(ctx, resourceId, false, customerKey)
			if err != nil { // billing failed?
				err = errors.WithMessage(err, "web3pay billing failed")
				return wrapup(ctx, msg, NewBillingStatusWithError(customerKey, err))
			}

			return wrapup(ctx, msg, NewBillingStatusWithReceipt(customerKey, receipt))
		}
	}
}

// providing net/http middleware

type HttpMiddleware = func(next http.Handler) http.Handler
type HttpResourceIdMapper func(req *http.Request) string

type HttpBillingMiddlewareOption struct {
	*BillingMiddlewareOption
	// gets resource ID from http request
	ResourceIdMapper HttpResourceIdMapper
}

func NewHttpBillingMiddlewareOptionWithClient(client *Client) *HttpBillingMiddlewareOption {
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
	return func(next http.Handler) http.Handler {
		wrapup := func(w http.ResponseWriter, r *http.Request, bs *BillingStatus) {
			// inject billing status to context
			ctx := context.WithValue(r.Context(), CtxKeyBillingStatus, bs)
			r = r.WithContext(ctx)

			if bs.Error == nil { // billing successfull
				logrus.WithField("receipt", bs.Receipt).Debug("Billing middleware billed successfully")
				vipCustomerKeyCache.Add(bs.customerKey, struct{}{})
				next.ServeHTTP(w, r)
				return
			}

			// handle gateway internal server error
			if err, ok := bs.InternalServerError(); ok {
				if !option.PropagateInternalServerError {
					_, bs.skipError = vipCustomerKeyCache.Get(bs.customerKey)
				}

				logrus.WithFields(logrus.Fields{
					"request": r, "skipError": bs.skipError,
					"propagateInternalServerError": option.PropagateInternalServerError,
				}).WithError(err).Error("Billing middleware internal server error")
			}

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			customerKey, ok := option.CustomerKeyProvider(ctx)
			if !ok || len(customerKey) == 0 { // customer key provided?
				bs := NewBillingStatusWithError(customerKey, errCustomerKeyNotProvided)
				wrapup(w, r, bs)
				return
			}

			// mapping resource ID
			var resourceId string
			if option.ResourceIdMapper != nil {
				resourceId = option.ResourceIdMapper(r)
			}

			receipt, err := option.Client.Bill(ctx, resourceId, false, customerKey)
			if err != nil { // billing failed?
				err = errors.WithMessage(err, "web3pay billing failed")
				wrapup(w, r, NewBillingStatusWithError(customerKey, err))
				return
			}

			wrapup(w, r, NewBillingStatusWithReceipt(customerKey, receipt))
		})
	}
}

// httpContextInjector injects data into context to return new context
type httpContextInjector = func(ctx context.Context, r *http.Request) context.Context

// HttpInjectContextMiddleware injects contextual data
func HttpInjectContextMiddleware(injectors ...httpContextInjector) HttpMiddleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			for i := range injectors {
				ctx = injectors[i](ctx, r)
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// handy utility functions

// CustomerKeyContextInjector returns context injector by extracting customer key from request
func CustomerKeyContextInjector(ckExtractor func(r *http.Request) string) httpContextInjector {
	return func(ctx context.Context, r *http.Request) context.Context {
		return context.WithValue(ctx, CtxCustomerKey, ckExtractor(r))
	}
}

// GetCustomerKeyFromContext gets customer key from context
func GetCustomerKeyFromContext(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(CtxCustomerKey).(string)
	return val, ok
}
