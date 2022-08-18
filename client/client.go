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
	jsonRpcMethodBillingCharge = "billing.Charge"
)

type ClientOption struct {
	BillingKey string        // billing auth key
	Timeout    time.Duration // request gateway timeout
}

type Client struct {
	jsonrpc.RPCClient // JSON-RPC request client
}

func NewClientWithOption(gateWayUrl string, option ClientOption) (*Client, error) {
	rpcClient := jsonrpc.NewClientWithOpts(gateWayUrl, &jsonrpc.RPCClientOpts{
		Timeout: option.Timeout,
		CustomHeaders: map[string]string{
			model.AuthHeaderBillingKey: option.BillingKey,
		},
	})

	// try to dial
	if _, err := rpcClient.Call(context.Background(), jsonRpcMethodBillingCharge); err != nil {
		return nil, errors.WithMessage(err, "failed to dial payment gateway")
	}

	return &Client{RPCClient: rpcClient}, nil
}

func (c *Client) ChargeBilling(
	ctx context.Context, resourceId string, dryRun bool, customerKey string) (*service.BillingChargeReceipt, error) {
	args := service.BillingChargeRequest{ResourceId: resourceId}
	ctx = jsonrpc.NewContextWithCustomHeaders(ctx, map[string]string{
		model.AuthHeaderCustomerKey: customerKey,
	})

	// call payment gateway for billing charge
	var reply *model.BusinessError
	if err := c.CallFor(ctx, &reply, jsonRpcMethodBillingCharge, []interface{}{args}); err != nil {
		logrus.WithField("args", args).
			WithError(err).
			Error("Web3pay client failed to request payment gateway")
		return nil, errors.WithMessage(err, "failed to request payment gateway")
	}

	// handle business error for payment gateway
	if reply.Code != model.ErrNil.Code {
		logrus.WithFields(logrus.Fields{
			"args":       args,
			"errCode":    reply.Code,
			"errMessage": reply.Message,
			"errData":    reply.Data,
		}).Debug("Web3pay client failed to billing charge from payment gateway")

		return nil, errors.WithMessage(reply, "failed to billing charge from payment gateway")
	}

	return reply.Data.(*service.BillingChargeReceipt), nil
}

// BuildBillingKey utility function to help build billing key with specified APP coin
// contract address and its owner private key text.
func BuildBillingKey(appCoin string, ownerPrivateKeyText string) (string, error) {
	// load APP coin contract owner private key
	privateKey, err := util.EcdsaPrivateKeyFromString(ownerPrivateKeyText)
	if err != nil {
		return "", err
	}

	// create signature
	sig, _, err := util.PersonalSign(appCoin, privateKey)
	if err != nil {
		return "", errors.WithMessage(err, "failed to create signature")
	}

	// json marshal auth key
	authKeyObj, err := json.Marshal(model.AuthKeyObject{
		Msg: appCoin, Sig: sig,
	})
	if err != nil {
		return "", errors.WithMessage(err, "failed to json marshal auth key object")
	}

	// base64 encoding auth key json
	billKey := base64.StdEncoding.EncodeToString(authKeyObj)
	return billKey, nil
}
