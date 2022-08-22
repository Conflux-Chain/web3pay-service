package client

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/Conflux-Chain/web3pay-service/client/jsonrpc"
	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/service"
	"github.com/Conflux-Chain/web3pay-service/util"
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
	Timeout    time.Duration // request timeout
}

type Client struct {
	*ClientConfig
	jrpcClient jsonrpc.RPCClient // JSON-RPC request client
}

func NewClient(conf *ClientConfig) (*Client, error) {
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
		ClientConfig: conf,
		jrpcClient:   rpcClient,
	}
	return client, nil
}

func (c *Client) Bill(
	ctx context.Context, resourceId string, dryRun bool, customerKey string) (*service.BillingReceipt, error) {

	ctx = jsonrpc.NewContextWithCustomHeaders(ctx, map[string]string{
		model.AuthHeaderCustomerKey: customerKey,
	})

	args := &service.BillingRequest{
		DryRun:     dryRun,
		ResourceId: resourceId,
	}

	reply, err := c.doCall(ctx, jrpcMethodBilling, args)
	if err != nil {
		return nil, err
	}

	var receipt service.BillingReceipt
	if err := reply.GetObject(&receipt); err != nil {
		return nil, errors.WithMessage(err, "failed to convert bill receipt")
	}

	return &receipt, nil
}

func (c *Client) BillBatch(
	ctx context.Context, resourceUses map[string]int64, dryRun bool, customerKey string) (*service.BillingBatchReceipt, error) {

	ctx = jsonrpc.NewContextWithCustomHeaders(ctx, map[string]string{
		model.AuthHeaderCustomerKey: customerKey,
	})

	args := &service.BillingBatchRequest{
		DryRun:       dryRun,
		ResourceUses: resourceUses,
	}

	reply, err := c.doCall(ctx, jrpcMethodBillingBatch, args)
	if err != nil {
		return nil, err
	}

	var receipt service.BillingBatchReceipt
	if err := reply.GetObject(&receipt); err != nil {
		return nil, errors.WithMessage(err, "failed to convert bill batch receipt")
	}

	return &receipt, nil
}

func (c *Client) doCall(ctx context.Context, method string, args interface{}) (*model.BusinessError, error) {
	var reply *model.BusinessError

	// call payment gateway
	if err := c.jrpcClient.CallFor(ctx, &reply, method, []interface{}{args}); err != nil {
		logrus.WithField("args", args).
			WithError(err).
			Debug("Web3Pay client failed to request payment gateway")
		return nil, errors.WithMessage(err, "failed to request payment gateway")
	}

	// handle business error
	if reply.Code != model.ErrNil.Code {
		logrus.WithFields(logrus.Fields{
			"args":       args,
			"errCode":    reply.Code,
			"errMessage": reply.Message,
			"errData":    reply.Data,
		}).Debug("Web3Pay client encountered internal business error")
		return nil, errors.WithMessage(reply, "internal business error")
	}

	return reply, nil
}

// BuildBillingKey utility function to help build billing key with specified
// APP coin contract address and its owner private key text.
func BuildBillingKey(appCoinContract string, ownerPrivateKeyText string) (string, error) {
	return BuildAuthKey(appCoinContract, ownerPrivateKeyText)
}

// BuildAuthKey utility function to help build auth key with specific message
// signatured with some private key text.
func BuildAuthKey(message string, privateKeyText string) (string, error) {
	// load private key
	privateKey, err := util.EcdsaPrivateKeyFromString(privateKeyText)
	if err != nil {
		return "", err
	}

	// create signature
	sig, _, err := util.PersonalSign(message, privateKey)
	if err != nil {
		return "", errors.WithMessage(err, "failed to create signature")
	}

	// json marshal auth key
	authKeyObj, err := json.Marshal(model.AuthKeyObject{
		Msg: message, Sig: sig,
	})
	if err != nil {
		return "", errors.WithMessage(err, "failed to json marshal auth key object")
	}

	// base64 encoding auth key json
	billKey := base64.StdEncoding.EncodeToString(authKeyObj)
	return billKey, nil
}
