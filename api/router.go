package api

import (
	"github.com/Conflux-Chain/web3pay-service/service"
	"github.com/gorilla/mux"
)

func newRouter(svcFactory *service.Factory) *mux.Router {
	r := mux.NewRouter()

	// TODO:
	// 1. add metrics middleware
	// 2. add CORS middleware?
	r.Use(AuthMiddleware(r, svcFactory.Blockchain))
	r.Use(LogTracingMiddleware)
	r.Use(LoggingMiddleware)

	billingCtr := NewBillingController()
	r.HandleFunc("/billing", Wrap(billingCtr.Charge, "web3pay/api/billing")).Methods("POST")

	return r
}
