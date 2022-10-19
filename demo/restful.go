package demo

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	web3pay "github.com/Conflux-Chain/web3pay-service/client"
	"github.com/Conflux-Chain/web3pay-service/client/middleware"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

// demo billed RESTful service handler
func testBilling(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	ctx := req.Context()
	bs, ok := middleware.BillingStatusFromContext(ctx)
	if !ok {
		rw.Write([]byte("billing middleware not enabled"))
		return
	}

	if bs.Success() {
		rw.Write([]byte("billing succeeded"))
		return
	}

	rw.Write([]byte(errors.WithMessage(bs.Error, "billing failed").Error()))
}

// RunRestfulServiceProvider runs a demo RESTful server to provide billed service.
func RunRestfulServiceProvider(config web3pay.BillingClientConfig, port int) error {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(testBilling))

	// create web3pay client
	client, err := web3pay.NewBillingClient(config)
	if err != nil {
		return errors.WithMessage(err, "failed to new web3pay client")
	}

	// hook http server middleware handler
	mwOption := middleware.NewHttpBillingMiddlewareOptionWithClient(client)
	kContextInjector := middleware.ApiKeyContextInjector(GetApiKey)
	ctxInjectMw := middleware.HttpInjectContextMiddleware(kContextInjector)
	handler := ctxInjectMw(middleware.HttpBillingMiddleware(mwOption)(mux))

	// serve RESTful RPC service
	endpoint := fmt.Sprintf(":%d", port)
	if err := http.ListenAndServe(endpoint, handler); err != http.ErrServerClosed {
		return errors.WithMessage(err, "failed to listen and server endpoint")
	}

	return nil
}

// RunRestfulServiceConsumer runs a RESTful consumer once to test the demo billed service provider.
func RunRestfulServiceConsumer(apiKey string, srvPort int) (interface{}, error) {
	httpSrvUrl := fmt.Sprintf("http://127.0.0.1:%d/%s", srvPort, url.QueryEscape(apiKey))

	// call billed service provider
	resp, err := resty.New().SetTimeout(200 * time.Millisecond).R().Post(httpSrvUrl)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func GetApiKey(r *http.Request) string {
	if r == nil || r.URL == nil {
		return ""
	}

	// API key path pattern:
	// http://example.com/${apiKey}...
	key := strings.TrimLeft(r.URL.EscapedPath(), "/")
	if idx := strings.Index(key, "/"); idx > 0 {
		key = key[:idx]
	}

	if key, err := url.PathUnescape(key); err == nil {
		return key
	}

	return ""
}
