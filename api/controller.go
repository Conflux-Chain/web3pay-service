package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/metrics"
	"github.com/sirupsen/logrus"
)

type BillingController struct{}

func NewBillingController() *BillingController {
	return &BillingController{}
}

type chargeRequest struct {
	ResourceId string `json:"resourceId"`
	DryRun     bool   `json:"dryRun"`
}

func (bc *BillingController) Charge(hc *handlerContext) (interface{}, error) {
	var cr chargeRequest
	if err := unmarshalRequestBody(hc.r, &cr); err != nil {
		return nil, errValidation.withData(err.Error())
	}

	// TODO: also validate `resourceId`
	if len(cr.ResourceId) == 0 {
		return nil, errValidation.withData("resource ID invalid")
	}

	logrus.WithFields(logrus.Fields{
		"reqID":         requestIdFromContext(hc.r.Context()),
		"chargeRequest": cr,
	}).Debug("New billing charge request received.")

	// TODO: billing charge logic here

	return nil, nil
}

func unmarshalRequestBody(r *http.Request, ptr interface{}) error {
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
