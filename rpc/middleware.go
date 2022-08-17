package rpc

import (
	"context"
	"time"

	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/service"
	"github.com/openweb3/go-rpc-provider"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/ybbus/jsonrpc/v3"
)

const (
	jsonRpcMethodBillingCharge = "billing.Charge"
)

// ResourceIdMapper maps RPC method to resource ID
type ResourceIdMapper func(method string) string

// PaymentMwProvider payment middleware provider
type PaymentMwProvider struct {
	gateWayURL       string
	client           jsonrpc.RPCClient
	fallback         rpc.HandleCallMsgMiddleware
	resourceIdMapper ResourceIdMapper
}

// NewPaymentMwProvider new payment middleware provider
func NewPaymentMwProvider(gateWayUrl string) (*PaymentMwProvider, error) {
	rpcClient := jsonrpc.NewClient(gateWayUrl)

	// try to dial
	ctx, cancer := context.WithTimeout(context.Background(), time.Second)
	defer cancer()

	if _, err := rpcClient.Call(ctx, jsonRpcMethodBillingCharge); err != nil {
		return nil, errors.WithMessage(err, "failed to dial payment gateway")
	}

	return &PaymentMwProvider{
		gateWayURL: gateWayUrl,
		client:     rpcClient,
	}, nil
}

// SetFallback sets fallback handling middleware if billing charge failed.
func (pmp *PaymentMwProvider) SetFallback(fallback rpc.HandleCallMsgMiddleware) {
	pmp.fallback = fallback
}

// SetResourceIdMapper sets mapper function to get resource ID by method.
func (pmp *PaymentMwProvider) SetResourceIdMapper(mapper ResourceIdMapper) {
	pmp.resourceIdMapper = mapper
}

func (pmp *PaymentMwProvider) getResourceId(method string) string {
	if pmp.resourceIdMapper != nil {
		return pmp.resourceIdMapper(method)
	}

	return method
}

// BillingCharge returns billing charge middleware.
func (pmp *PaymentMwProvider) BillingCharge(next rpc.HandleCallMsgFunc) rpc.HandleCallMsgFunc {
	return func(ctx context.Context, msg *rpc.JsonRpcMessage) *rpc.JsonRpcMessage {
		var args service.BillingChargeRequest
		var reply *model.BusinessError

		// map method to resource ID
		args.ResourceId = pmp.getResourceId(msg.Method)

		// TODO: add auth headers

		// call payment gateway for billing charge
		if err := pmp.client.CallFor(ctx, &reply, jsonRpcMethodBillingCharge, args); err != nil {
			logrus.WithField("args", args).
				WithError(err).
				Error("Billing charge middleware failed to request payment gateway")

			if pmp.fallback != nil {
				return pmp.fallback(next)(ctx, msg)
			}

			return msg.ErrorResponse(err)
		}

		// handle business error for payment gateway
		if reply.Code != model.ErrNil.Code {
			logrus.WithFields(logrus.Fields{
				"args":       args,
				"errCode":    reply.Code,
				"errMessage": reply.Message,
				"errData":    reply.Data,
			}).Debug("Billing charge middleware failed to billing charge from payment gateway")

			if pmp.fallback != nil {
				return pmp.fallback(next)(ctx, msg)
			}

			return msg.ErrorResponse(errors.WithMessage(reply, "billing charge middleware error"))
		}

		logrus.WithFields(logrus.Fields{
			"args":    args,
			"receipt": reply.Data,
		}).Debug("Billing charge middleware charged successfully")

		return next(ctx, msg)
	}
}
