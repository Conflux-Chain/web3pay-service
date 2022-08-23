package client

import (
	"context"

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

type authKeyProvider func(ctx context.Context) (string, bool)

// BillingMiddleware provides billing RPC middleware for openweb3.
//
// Parameter `ignoreInternalGatewayError` toggles whether to ignore internal gateway error for already billed users, default as true.
// This is usually used to mitigate side effects such as blocking paid user from access due to internal server errors.
func BillingMiddleware(
	client *Client, customerKeyProvider authKeyProvider, ignoreInternalGatewayError ...bool) rpc.HandleCallMsgMiddleware {

	ignoreInternalServerError := true
	if len(ignoreInternalGatewayError) > 0 {
		ignoreInternalServerError = ignoreInternalGatewayError[0]
	}

	return func(next rpc.HandleCallMsgFunc) rpc.HandleCallMsgFunc {
		wrapup := func(ctx context.Context, msg *rpc.JsonRpcMessage, bs *BillingStatus) *rpc.JsonRpcMessage {
			ctx = context.WithValue(ctx, CtxKeyBillingStatus, bs)

			if bs.Error == nil { // billing successfully?
				vipCustomerKeyCache.Add(bs.customerKey, struct{}{})
				return next(ctx, msg)
			}

			// handle gateway internal server error
			if err, ok := bs.InternalServerError(); ok {
				logrus.WithField("msg", msg).WithError(err).Error("Billing middleware internal server error")

				if ignoreInternalServerError {
					_, bs.skipError = vipCustomerKeyCache.Get(bs.customerKey)
				}
			}

			return next(ctx, msg)
		}

		return func(ctx context.Context, msg *rpc.JsonRpcMessage) *rpc.JsonRpcMessage {
			customerKey, ok := customerKeyProvider(ctx)
			if !ok { // customer key provided?
				bs := NewBillingStatusWithError(customerKey, errCustomerKeyNotProvided)
				return wrapup(ctx, msg, bs)
			}

			receipt, err := client.Bill(ctx, msg.Method, false, customerKey)
			if err != nil { // billing failed?
				err = errors.WithMessage(err, "web3pay billing failed")
				return wrapup(ctx, msg, NewBillingStatusWithError(customerKey, err))
			}

			return wrapup(ctx, msg, NewBillingStatusWithReceipt(customerKey, receipt))
		}
	}
}
