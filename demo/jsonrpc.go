package demo

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	web3pay "github.com/Conflux-Chain/web3pay-service/client"
	"github.com/Conflux-Chain/web3pay-service/client/jsonrpc"
	"github.com/Conflux-Chain/web3pay-service/client/middleware"
	"github.com/ethereum/go-ethereum/node"
	"github.com/openweb3/go-rpc-provider"
	"github.com/pkg/errors"
)

// jrDemoApi provides a demo JSON-RPC billed service.
type jrDemoApi struct{}

func (api *jrDemoApi) TestBilling(ctx context.Context) (string, error) {
	bs, ok := middleware.BillingStatusFromContext(ctx)
	if !ok {
		return "", errors.New("billing middleware not enabled")
	}

	if bs.Success() {
		return "billing succeeded", nil
	}

	return "", errors.WithMessage(bs.Error, "billing failed")
}

func (api *jrDemoApi) TestSubscription(ctx context.Context) (string, error) {
	ss, ok := middleware.VipSubscriptionStatusFromContext(ctx)
	if !ok {
		return "", errors.New("VIP subscription middleware  not enabled")
	}

	vi, err := ss.GetVipInfo()
	if err != nil {
		return "", errors.WithMessage(ss.Error, "VIP subscription middleware failed")
	}

	vidata, err := json.Marshal(vi)
	return string(vidata), err
}

// RunBillingJsonRpcServiceProvider runs a demo JSON-RPC server to provide billing service.
func RunBillingJsonRpcServiceProvider(config web3pay.BillingClientConfig, port int) error {
	// create web3pay client
	client, err := web3pay.NewBillingClient(config)
	if err != nil {
		return errors.WithMessage(err, "failed to new web3pay client")
	}

	// hook web3pay billing middleware for go-rpc-provider
	mwoption := middleware.NewOw3BillingMiddlewareOptionWithClient(client)
	rpc.HookHandleCallMsg(middleware.Openweb3BillingMiddleware(mwoption))

	return runJsonRpcDemoApiServer(port)
}

// RunSubscriptionJsonRpcServiceProvider runs a demo JSON-RPC server to provide VIP subscription service.
func RunSubscriptionJsonRpcServiceProvider(config web3pay.VipSubscriptionClientConfig, port int) error {
	// create web3pay VIP subscription client
	client, err := web3pay.NewVipSubscriptionClient(config)
	if err != nil {
		return errors.WithMessage(err, "failed to new web3pay VIP subscription client")
	}

	// hook web3pay billing middleware for go-rpc-provider
	mwoption := middleware.NewVipSubscriptionMiddlewareOptionWithClient(client)
	rpc.HookHandleCallMsg(middleware.Openweb3VipSubscriptionMiddleware(mwoption))

	return runJsonRpcDemoApiServer(port)
}

func runJsonRpcDemoApiServer(port int) error {
	// create JSON-RPC server
	handler := rpc.NewServer()
	if err := handler.RegisterName("demo", &jrDemoApi{}); err != nil {
		return errors.WithMessage(err, "failed to register demo API service")
	}

	kContextInjector := middleware.ApiKeyContextInjector(GetApiKey)
	ctxInjectMw := middleware.HttpInjectContextMiddleware(kContextInjector)
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

// RunBillingJsonRpcServiceConsumer runs a JSON-RPC consumer once to test the demo billing service provider.
func RunBillingJsonRpcServiceConsumer(apiKey string, srvPort int) (interface{}, error) {
	return callJsonRpcServiceProvider(apiKey, srvPort, "demo_testBilling")
}

// RunSubscriptionJsonRpcServiceConsumer runs a JSON-RPC consumer once to test the demo VIP subscription service provider.
func RunSubscriptionJsonRpcServiceConsumer(apiKey string, srvPort int) (interface{}, error) {
	return callJsonRpcServiceProvider(apiKey, srvPort, "demo_testSubscription")
}

func callJsonRpcServiceProvider(apiKey string, srvPort int, method string) (interface{}, error) {
	rpcSrvUrl := fmt.Sprintf("http://127.0.0.1:%d/%s", srvPort, url.QueryEscape(apiKey))
	rpcClient := jsonrpc.NewClientWithOpts(
		rpcSrvUrl, &jsonrpc.RPCClientOpts{Timeout: time.Second},
	)

	// call service provider
	var reply string
	err := rpcClient.CallFor(context.Background(), &reply, method, []interface{}{})
	if err != nil {
		return nil, err
	}

	return reply, nil
}
