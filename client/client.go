package client

import (
	"context"
	"time"

	"github.com/Conflux-Chain/web3pay-service/client/jsonrpc"
	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/service"
	"github.com/mcuadros/go-defaults"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	jrpcMethodBilling      = "web3pay.Bill"
	jrpcMethodBillingBatch = "web3pay.BillBatch"
)

type ClientConfig struct {
	Gateway    string        // API gateway endpoint
	BillingKey string        // billing auth key
	PingTest   bool          // test ping gateway?
	Timeout    time.Duration `default:"200ms"` // request timeout, default 200ms
}

type Client struct {
	*ClientConfig
	jrpcClient jsonrpc.RPCClient // JSON-RPC request client
}

func NewClient(conf ClientConfig) (*Client, error) {
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

	client := &Client{
		ClientConfig: &conf,
		jrpcClient:   rpcClient,
	}
	return client, nil
}

func (c *Client) Bill(
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

func (c *Client) BillBatch(
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

func (c *Client) doCall(ctx context.Context, out interface{}, method string, args interface{}) error {
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
