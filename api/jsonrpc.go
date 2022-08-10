package api

import (
	"net/http"

	"github.com/Conflux-Chain/web3pay-service/service"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/sirupsen/logrus"
)

type JrChargeArgs struct {
	chargeRequest
}

type JrChargeRely struct {
	service.ChargeReceipt
}

type JrBillingApi struct {
	billingSvc *service.BillingService
}

// JSON-RPC Billing API can be requested like:
// {"jsonrpc":"2.0","method":"billing.Charge","params":[{ "dryRun": true, "resourceId": "default"}],"id":1}
func (api *JrBillingApi) Charge(r *http.Request, args *JrChargeArgs, reply *JrChargeRely) error {
	ctx := r.Context()
	reqId := requestIdFromContext(ctx)
	contract := contractAddrFromContext(ctx)
	customer := customerAddrFromContext(ctx)

	chargeReq := &service.ChargeRequest{
		ResourceId: args.ResourceId,
		DryRun:     args.DryRun,
		AppCoin:    contract,
		Customer:   customer,
	}

	logger := logrus.WithFields(logrus.Fields{
		"chargeRequest": chargeReq,
		"requestId":     reqId,
		"isJsonRPC":     true,
	})

	receipt, err := api.billingSvc.Charge(ctx, chargeReq)
	if err != nil {
		logger.WithError(err).Debug("Billing charge failed")
		return err
	}

	reply.ChargeReceipt = *receipt
	logger.WithField("receipt", receipt).Debug("Billing charge done")

	return nil
}

func newJsonRpcServer(svcFactory *service.Factory) *rpc.Server {
	// Create JSON-RPC server
	srv := rpc.NewServer()

	// Register the type of data requested as JSON
	srv.RegisterCodec(json.NewCodec(), "application/json")
	srv.RegisterService(&JrBillingApi{billingSvc: svcFactory.Billing}, "billing")

	return srv
}
