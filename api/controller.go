package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Conflux-Chain/web3pay-service/service"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/sirupsen/logrus"
)

type BillingController struct {
	billingSvc *service.BillingService
}

func NewBillingController(billingSvc *service.BillingService) *BillingController {
	return &BillingController{billingSvc: billingSvc}
}

type chargeRequest struct {
	ResourceId string `json:"resourceId"`
	DryRun     bool   `json:"dryRun"`
}

func (bc *BillingController) Charge(hc *handlerContext) (interface{}, error) {
	var cr chargeRequest
	if err := jsonUnmarshalRequestBody(hc.r, &cr); err != nil {
		return nil, errValidation.withData(err.Error())
	}

	ctx := hc.r.Context()
	reqId := requestIdFromContext(ctx)
	contractAddr := contractAddrFromContext(ctx)
	customerAddr := customerAddrFromContext(ctx)

	chargeReq := &service.ChargeRequest{
		ResourceId:   cr.ResourceId,
		DryRun:       cr.DryRun,
		ContractAddr: contractAddr,
		CustomerAddr: customerAddr,
	}

	logger := logrus.WithFields(logrus.Fields{
		"chargeRequest": chargeReq,
		"requestId":     reqId,
	})

	receipt, err := bc.billingSvc.Charge(ctx, chargeReq)
	if err != nil {
		logger.WithError(err).Debug("Billing charge failed")
		return nil, err
	}

	logger.WithField("receipt", receipt).Debug("Billing charge done")
	return receipt, nil
}

func jsonUnmarshalRequestBody(r *http.Request, ptr interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(body), ptr); err != nil {
		return err
	}

	return nil
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
	start := time.Now()
	rw.Header().Set("Content-Type", "application/json")

	result, err := w.inner(&handlerContext{w: rw, r: r})
	if err != nil {
		respJsonError(rw, err)
	} else {
		respJsonOK(rw, result)
		w.perf.UpdateSince(start)
	}
}

func respJsonOK(rw http.ResponseWriter, payload interface{}) {
	if err := json.NewEncoder(rw).Encode(errNil.withData(payload)); err != nil {
		panic("json encoding error")
	}
}

func respJsonError(rw http.ResponseWriter, err error) {
	var encodingErr error

	switch e := err.(type) {
	case *businessError: // business error
		rw.WriteHeader(http.StatusOK)
		encodingErr = json.NewEncoder(rw).Encode(e)
	default: // internal server error
		rw.WriteHeader(http.StatusInternalServerError)
		encodingErr = json.NewEncoder(rw).Encode(errInternalServer.withData(err.Error()))
	}

	if encodingErr != nil {
		panic("json encoding error")
	}
}
