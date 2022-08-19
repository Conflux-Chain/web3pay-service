package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/service"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type BillingController struct {
	billingSvc *service.BillingService
}

func NewBillingController(billingSvc *service.BillingService) *BillingController {
	return &BillingController{billingSvc: billingSvc}
}

func (bc *BillingController) Bill(hc *handlerContext) (interface{}, error) {
	var cr service.BillingRequest
	if _, err := jsonUnmarshalRequestBody(hc.r.Body, &cr); err != nil {
		return nil, model.ErrValidation.WithData(err.Error())
	}

	ctx := hc.r.Context()
	reqId := requestIdFromContext(ctx)
	cr.AppCoin = contractAddrFromContext(ctx)
	cr.Customer = customerAddrFromContext(ctx)

	logger := logrus.WithFields(logrus.Fields{
		"billRequest": cr,
		"requestId":   reqId,
	})

	receipt, err := bc.billingSvc.Bill(ctx, &cr)
	if err != nil {
		logger.WithError(err).Debug("Billing failed")
		return nil, err
	}

	logger.WithField("receipt", receipt).Debug("Billing done")
	return receipt, nil
}

type handlerContext struct {
	w http.ResponseWriter
	r *http.Request
}

type wrapper struct {
	inner func(hc *handlerContext) (interface{}, error)
	perf  metrics.Timer
}

func Wrap(
	controllerFunc func(hc *handlerContext) (interface{}, error),
	metricName string,
) func(w http.ResponseWriter, r *http.Request) {
	w := wrapper{
		controllerFunc,
		metrics.GetOrRegisterTimer(metricName, nil),
	}

	return w.wrap
}

func (w *wrapper) wrap(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	start := time.Now()
	rw.Header().Set("Content-Type", "application/json")

	result, err := w.inner(&handlerContext{w: rw, r: r})
	if err != nil {
		respJsonError(ctx, rw, err)
	} else {
		respJsonOK(ctx, rw, result)
		w.perf.UpdateSince(start)
	}
}

func respJsonOK(ctx context.Context, rw http.ResponseWriter, payload interface{}) {
	if err := json.NewEncoder(rw).Encode(model.ErrNil.WithData(payload)); err != nil {
		panic("json encoding error")
	}
}

func respJsonError(ctx context.Context, rw http.ResponseWriter, err error) {
	var encodingErr error

	switch e := err.(type) {
	case *model.BusinessError: // business error
		rw.WriteHeader(http.StatusOK)
		encodingErr = json.NewEncoder(rw).Encode(e)
	default: // internal server error
		rw.WriteHeader(http.StatusInternalServerError)
		encodingErr = json.NewEncoder(rw).Encode(model.ErrInternalServer.WithData(err.Error()))
	}

	if encodingErr != nil {
		panic(errors.WithMessage(encodingErr, "json encoding error"))
	}
}
