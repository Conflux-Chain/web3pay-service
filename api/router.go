package api

import (
	"net/http"

	"github.com/Conflux-Chain/web3pay-service/service"
	"github.com/gorilla/mux"
)

type chainedRouter struct {
	routers []*mux.Router
}

func (cr *chainedRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, r := range cr.routers {
		var match mux.RouteMatch
		if r.Match(req, &match) { // matched
			r.ServeHTTP(w, req)
			return
		}
	}

	// fallback
	if len(cr.routers) > 0 {
		cr.routers[len(cr.routers)-1].ServeHTTP(w, req)
	}
}

func newChainedRouter(svcFactory *service.Factory) *chainedRouter {
	return &chainedRouter{routers: []*mux.Router{
		newJsonRpcRouter(svcFactory),
		newRestfulRouter(svcFactory),
	}}
}

func newRestfulRouter(svcFactory *service.Factory) *mux.Router {
	r := mux.NewRouter()

	// TODO:
	// 1. add metrics middleware
	// 2. add CORS middleware?
	r.Use(LogTracingMiddleware("REST"))
	r.Use(LoggingMiddleware)
	r.Use(AuthMiddleware(r, svcFactory.Blockchain, respJsonError))

	billingCtr := NewBillingController(svcFactory.Billing)
	r.HandleFunc("/billing", Wrap(billingCtr.Bill, "web3pay/api/billing")).
		Methods("POST").
		Headers("Content-Type", "application/json")

	return r
}

func newJsonRpcRouter(svcFactory *service.Factory) *mux.Router {
	r := mux.NewRouter()

	// TODO:
	// 1. add metrics middleware
	// 2. add CORS middleware?
	r.Use(JsonRpcValidationMiddleware)
	r.Use(LogTracingMiddleware("JRPC"))
	r.Use(LoggingMiddleware)
	r.Use(AuthMiddleware(r, svcFactory.Blockchain, respJsonRpcError))

	// json rpc
	r.Handle("/", newJsonRpcServer(svcFactory))

	return r
}
