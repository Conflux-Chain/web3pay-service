package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Conflux-Chain/web3pay-service/client/jsonrpc"
	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/service"
	"github.com/gorilla/rpc/v2"
	gjson "github.com/gorilla/rpc/v2/json"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type JrBillingApi struct {
	billingSvc *service.BillingService
}

// JSON-RPC Billing API can be requested like:
// {"jsonrpc":"2.0","method":"billing.Charge","params":[{ "dryRun": true, "resourceId": "default"}],"id":1}
func (api *JrBillingApi) Charge(r *http.Request, args *service.BillingChargeRequest, reply **model.BusinessError) error {
	ctx := r.Context()
	args.AppCoin = contractAddrFromContext(ctx)
	args.Customer = customerAddrFromContext(ctx)

	logger := logrus.WithFields(logrus.Fields{
		"args":      args,
		"requestId": requestIdFromContext(ctx),
	})

	receipt, err := api.billingSvc.Charge(ctx, args)
	if err != nil {
		logger.WithError(err).Debug("Billing charge failed")

		if bizerr, ok := err.(*model.BusinessError); ok {
			*reply = bizerr
			return nil
		}

		*reply = model.ErrInternalServer.WithData(err.Error())
		return nil
	}

	logger.WithField("receipt", receipt).Debug("Billing charge done")
	*reply = model.ErrNil.WithData(receipt)

	return nil
}

func newJsonRpcServer(svcFactory *service.Factory) *rpc.Server {
	// Create JSON-RPC server
	srv := rpc.NewServer()

	// Register the type of data requested as JSON
	srv.RegisterCodec(gjson.NewCodec(), "application/json")
	srv.RegisterService(&JrBillingApi{billingSvc: svcFactory.Billing}, "billing")

	return srv
}

func respJsonRpcError(ctx context.Context, rw http.ResponseWriter, err error) {
	resp := jsonrpc.RPCResponse{JSONRPC: "2.0"}
	msg := jsonRpcMessageFromContext(ctx)
	if msg != nil {
		resp.ID = msg.ID
	}

	if bizerr, ok := err.(*model.BusinessError); ok { // business error?
		resp.Result = bizerr
	} else { // non-buisiness error?
		rpcerr, ok := err.(*jsonrpc.RPCError)
		if !ok {
			rpcerr = &jsonrpc.RPCError{
				Code:    -1,
				Message: "Bad request",
				Data:    err.Error(),
			}
		}

		resp.Error = rpcerr
		rw.WriteHeader(400)
	}

	if encodingErr := json.NewEncoder(rw).Encode(&resp); encodingErr != nil {
		panic(errors.WithMessage(encodingErr, "json encoding error"))
	}
}
