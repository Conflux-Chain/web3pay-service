package client

import (
	"context"

	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/service"
	"github.com/openweb3/go-rpc-provider"
	"github.com/pkg/errors"
)

type CtxKey string

const (
	CtxKeyBillingStatus = CtxKey("Web3Pay-Billing-Status")
)

var (
	errCustomerKeyNotProvided = errors.New("customer key not provided")
)

// BillingStatus billing result
type BillingStatus struct {
	Data  interface{}
	Error error
}

func NewBillingStatusWithError(err error) *BillingStatus {
	return &BillingStatus{Error: err}
}

func NewBillingStatusWithReceipt(data interface{}) *BillingStatus {
	return &BillingStatus{Data: data}
}

// Success checks if billing succeeded
func (bs *BillingStatus) Success() bool {
	return bs.Error == nil
}

// BusinessError returns business error as it is otherwise nil
func (bs *BillingStatus) BusinessError() (*model.BusinessError, bool) {
	var bzerr model.BusinessError
	if errors.As(bs.Error, &bzerr) {
		return &bzerr, true
	}

	return nil, false
}

// Receipt returns billing receipt
func (bs *BillingStatus) Receipt() (*service.BillingReceipt, bool) {
	receipt, ok := bs.Data.(*service.BillingReceipt)
	return receipt, ok
}

// BillingStatusFromContext returns billing status from context
func BillingStatusFromContext(ctx context.Context) (*BillingStatus, bool) {
	bs, ok := ctx.Value(CtxKeyBillingStatus).(*BillingStatus)
	return bs, ok
}

// BillingMiddleware provides billing RPC middleware for openweb3
func BillingMiddleware(client *Client, customerKeyProvider func(ctx context.Context) (string, bool)) rpc.HandleCallMsgMiddleware {
	return func(next rpc.HandleCallMsgFunc) rpc.HandleCallMsgFunc {
		wrapup := func(ctx context.Context, msg *rpc.JsonRpcMessage, bs *BillingStatus) *rpc.JsonRpcMessage {
			ctx = context.WithValue(ctx, CtxKeyBillingStatus, bs)
			return next(ctx, msg)
		}

		return func(ctx context.Context, msg *rpc.JsonRpcMessage) *rpc.JsonRpcMessage {
			customerKey, ok := customerKeyProvider(ctx)
			if !ok { // customer key provided?
				return wrapup(ctx, msg, NewBillingStatusWithError(errCustomerKeyNotProvided))
			}

			receipt, err := client.Bill(ctx, msg.Method, false, customerKey)
			if err != nil { // billing failed?
				err = errors.WithMessage(err, "web3pay billing failed")
				return wrapup(ctx, msg, NewBillingStatusWithError(err))
			}

			return wrapup(ctx, msg, NewBillingStatusWithReceipt(receipt))
		}
	}
}
