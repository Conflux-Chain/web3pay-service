package api

import (
	"github.com/gorilla/mux"
)

type controllerFactory struct {
	billing *BillingController
}

func newRouter() *mux.Router {
	r := mux.NewRouter()

	// middlewares
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(LogTracingMiddleware)

	factory := controllerFactory{
		billing: NewBillingController(),
	}

	r.HandleFunc("/billing", Wrap(factory.billing.Charge, "web3pay/api/billing")).Methods("POST")

	return r
}
