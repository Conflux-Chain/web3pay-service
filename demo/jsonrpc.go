package demo

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	web3pay "github.com/Conflux-Chain/web3pay-service/client"
	"github.com/Conflux-Chain/web3pay-service/client/jsonrpc"
	"github.com/ethereum/go-ethereum/node"
	"github.com/openweb3/go-rpc-provider"
	"github.com/pkg/errors"
)

// jrDemoApi provides a demo JSON-RPC billed service.
type jrDemoApi struct{}

func (api *jrDemoApi) TestBilling(ctx context.Context) (string, error) {
	bs, ok := web3pay.BillingStatusFromContext(ctx)
	if !ok {
		return "", errors.New("billing middleware not enabled")
	}

	if bs.Success() {
		return "billing succeeded", nil
	}

	return "", errors.WithMessage(bs.Error, "billing failed")
}

// RunJsonRpcServiceProvider runs a demo JSON-RPC server to provide billed service.
func RunJsonRpcServiceProvider(config web3pay.ClientConfig, port int) error {
	// create web3pay client
	client, err := web3pay.NewClient(config)
	if err != nil {
		return errors.WithMessage(err, "failed to new web3pay client")
	}

	// hook web3pay billing middleware for go-rpc-provider
	mwoption := web3pay.DefaultOw3BillingMiddlewareOption(client)
	rpc.HookHandleCallMsg(web3pay.Openweb3BillingMiddleware(mwoption))

	// create JSON-RPC server
	handler := rpc.NewServer()
	if err := handler.RegisterName("demo", &jrDemoApi{}); err != nil {
		return errors.WithMessage(err, "failed to register demo API service")
	}

	ckContextInjector := web3pay.CustomerKeyContextInjector(GetCustomerKey)
	ctxInjectMw := web3pay.HttpInjectContextMiddleware(ckContextInjector)
	server := http.Server{
		Handler: ctxInjectMw(node.NewHTTPHandlerStack(handler, []string{"*"}, []string{"*"})),
	}

	// listen endpoint
	endpoint := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", endpoint)
	if err != nil {
		return errors.WithMessage(err, "failed to listen to endpoint")
	}

	// serve JSON-RPC service
	if err := server.Serve(listener); err != http.ErrServerClosed {
		return err
	}

	return nil
}

// RunJsonRpcServiceConsumer runs a JSON-RPC consumer once to test the demo billed service provider.
func RunJsonRpcServiceConsumer(customerKey string, srvPort int) (interface{}, error) {
	rpcSrvUrl := fmt.Sprintf("http://127.0.0.1:%d/%s", srvPort, url.QueryEscape(customerKey))
	rpcClient := jsonrpc.NewClientWithOpts(
		rpcSrvUrl, &jsonrpc.RPCClientOpts{Timeout: time.Second},
	)

	// call billed service provider
	var reply string
	err := rpcClient.CallFor(context.Background(), &reply, "demo_testBilling", []interface{}{})
	if err != nil {
		return nil, err
	}

	return reply, nil
}