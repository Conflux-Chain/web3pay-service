package rpc

import (
	"context"
	"fmt"

	"github.com/Conflux-Chain/web3pay-service/api"
	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/openweb3/go-rpc-provider"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// ResourceIdMapper maps RPC method to resource ID
type ResourceIdMapper func(method string) string

// PaymentMwProvider payment middleware provider
type PaymentMwProvider struct {
	gateWayURL       string
	client           *rpc.Client
	fallback         rpc.HandleCallMsgMiddleware
	resourceIdMapper ResourceIdMapper
}

// NewPaymentMwProvider new payment middleware provider
func NewPaymentMwProvider(gateWayUrl string) (*PaymentMwProvider, error) {
	client, err := rpc.DialHTTP(gateWayUrl)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to dial payment gateway")
	}

	return &PaymentMwProvider{
		gateWayURL: gateWayUrl,
		client:     client,
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
		var args api.JrChargeArgs
		var resp struct {
			Result api.JrChargeRely `json:"result"`
			Error  interface{}      `json:"error"`
			// ignore `id`
		}

		// map method to resource ID
		args.ResourceId = pmp.getResourceId(msg.Method)

		// call payment gateway for billing charge
		if err := pmp.client.Call(&resp, "billing.Charge", args); err != nil {
			logrus.WithField("args", args).
				WithError(err).
				Error("Billing charge middleware failed to request payment gateway")

			if pmp.fallback != nil {
				return pmp.fallback(next)(ctx, msg)
			}

			return msg.ErrorResponse(err)
		}

		// handle internal error for payment gateway
		if resp.Error != nil {
			logrus.WithFields(logrus.Fields{
				"args":  args,
				"error": resp.Error,
			}).Warn("Billing charge middleware encountered internal payment gateway error")

			if pmp.fallback != nil {
				return pmp.fallback(next)(ctx, msg)
			}

			err := fmt.Errorf("bad payment gateway: %v", resp.Error)
			return msg.ErrorResponse(err)
		}

		// handle business error for payment gateway
		if resp.Result.Code != model.ErrNil.Code {
			logrus.WithFields(logrus.Fields{
				"args":  args,
				"reply": resp.Result,
			}).Debug("Billing charge middleware failed to billing charge from payment gateway")

			if pmp.fallback != nil {
				return pmp.fallback(next)(ctx, msg)
			}

			return msg.ErrorResponse(resp.Result)
		}

		return next(ctx, msg)
	}
}
