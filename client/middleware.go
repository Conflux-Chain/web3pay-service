package client

import (
	"context"
	"time"

	"github.com/openweb3/go-rpc-provider"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	errCustomerKeyNotProvided = errors.New("web3pay customer key not provided")
)

// authKeyProvider provides auth key from context value
type authKeyProvider func(ctx context.Context) (string, bool)

type RpcMiddlewareProviderOption struct {
	// request timeout to payment gateway
	Timeout time.Duration

	// web3pay billing key
	BillingKey string
	// web3pay customer key getter
	CustomerKey authKeyProvider

	// fallback for billing middleware on failure
	BillingFallbackMw      rpc.HandleCallMsgMiddleware
	BillingBatchFallbackMw rpc.HandleBatchMiddleware
}

// RpcMiddlewareProvider provides RPC middleware for openweb3.
type RpcMiddlewareProvider struct {
	RpcMiddlewareProviderOption
	client *Client
}

func NewRpcMiddlewareProvider(
	paymentGateway string, option RpcMiddlewareProviderOption) (*RpcMiddlewareProvider, error) {

	// new web3pay client
	client, err := NewClientWithOption(paymentGateway, ClientOption{
		BillingKey: option.BillingKey,
		Timeout:    option.Timeout,
	})

	if err != nil {
		return nil, errors.WithMessage(err, "failed to create web3pay client")
	}

	provider := &RpcMiddlewareProvider{
		RpcMiddlewareProviderOption: option,
		client:                      client,
	}
	return provider, nil
}

// BillingMiddleware provides fee billing middleware.
func (p *RpcMiddlewareProvider) BillingMiddleware(next rpc.HandleCallMsgFunc) rpc.HandleCallMsgFunc {
	// fallback handler if billing failed
	fbhandler := func(ctx context.Context, msg *rpc.JsonRpcMessage, err error) *rpc.JsonRpcMessage {
		logger := logrus.WithField("method", msg.Method).WithError(err)

		if p.BillingFallbackMw != nil {
			logger.Debug("Billing middleware switching over to fallback middleware on failure")
			return p.BillingFallbackMw(next)(ctx, msg)
		}

		logger.Debug("Billing middleware billing failed")
		return msg.ErrorResponse(err)
	}

	return func(ctx context.Context, msg *rpc.JsonRpcMessage) *rpc.JsonRpcMessage {
		customerKey, ok := p.CustomerKey(ctx)
		if !ok { // customer key provided?
			return fbhandler(ctx, msg, errCustomerKeyNotProvided)
		}

		receipt, err := p.client.Bill(ctx, msg.Method, false, customerKey)
		if err != nil { // billing failed?
			return fbhandler(ctx, msg, errors.WithMessage(err, "web3pay billing failure"))
		}

		logrus.WithFields(logrus.Fields{
			"receipt": receipt,
			"method":  msg.Method,
		}).Debug("Billing middleware billing ok")

		return next(ctx, msg)
	}
}

// BillingMiddleware provides batch fee billing middleware.
func (p *RpcMiddlewareProvider) BillingBatchMiddleware(next rpc.HandleBatchFunc) rpc.HandleBatchFunc {
	// fallback handler if billing failed
	fbhandler := func(ctx context.Context, msgs []*rpc.JsonRpcMessage, err error) []*rpc.JsonRpcMessage {
		logger := logrus.WithField("batch", len(msgs)).WithError(err)

		if p.BillingBatchFallbackMw != nil {
			logger.Debug("Billing middleware switching over to fallback middleware on failure")
			return p.BillingBatchFallbackMw(next)(ctx, msgs)
		}

		var responses []*rpc.JsonRpcMessage
		for _, v := range msgs {
			responses = append(responses, v.ErrorResponse(err))
		}

		logger.Debug("Billing middleware billing failed")
		return responses
	}

	return func(ctx context.Context, msgs []*rpc.JsonRpcMessage) []*rpc.JsonRpcMessage {
		customerKey, ok := p.CustomerKey(ctx)
		if !ok { // customer key provided?
			return fbhandler(ctx, msgs, errCustomerKeyNotProvided)
		}

		// collect RPC method uses
		msgCalls := make(map[string]int64)
		for i := range msgs {
			msgCalls[msgs[i].Method]++
		}

		receipt, err := p.client.BillBatch(ctx, msgCalls, true, customerKey)
		if err != nil { // batch billing failed?
			return fbhandler(ctx, msgs, errors.WithMessage(err, "web3pay billing failure"))
		}

		logrus.WithFields(logrus.Fields{
			"receipt": receipt,
			"batch":   len(msgs),
		}).Debug("Billing middleware billing ok")

		return next(ctx, msgs)
	}
}
