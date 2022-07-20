package api

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"math"
	"net/http"

	mathutil "github.com/Conflux-Chain/go-conflux-util/math"
	"github.com/Conflux-Chain/web3pay-service/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type reqIdCtxKey string

func LogTracingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqId := mathutil.RandUint64(uint64(math.MaxUint32))
		newCtx := context.WithValue(r.Context(), reqIdCtxKey("reqId"), reqId)

		newReq := r.WithContext(newCtx)
		next.ServeHTTP(w, newReq)
	})
}

func requestIdFromContext(ctx context.Context) uint64 {
	if reqId, ok := ctx.Value(reqIdCtxKey("reqId")).(uint64); ok {
		return reqId
	}

	return 0
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.WithFields(logrus.Fields{
			"reqID":  requestIdFromContext(r.Context()),
			"method": r.Method,
		}).Debug(r.RequestURI)

		next.ServeHTTP(w, r)
	})
}

type authKey struct {
	Msg string // signed message
	Sig string // signature
}

func parseAuthKey(r *http.Request, headerKey string) (*authKey, error) {
	keyJson, err := base64.StdEncoding.DecodeString(r.Header.Get(headerKey))
	if err != nil {
		return nil, errors.WithMessage(err, "base64 decode error")
	}

	var key authKey
	if err := json.Unmarshal(keyJson, &key); err != nil {
		return nil, errors.WithMessage(err, "json decode error")
	}

	if len(key.Msg) == 0 {
		return nil, errors.New("msg not provided")
	}

	if len(key.Sig) == 0 {
		return nil, errors.New("sig not provided")
	}

	return &key, err
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		billingKey, err := parseAuthKey(r, "Billing-Key")
		if err != nil {
			err = errors.WithMessage(err, "billing key parse error")
			respJsonError(w, errAuth.withData(err.Error()))
			return
		}

		customerKey, err := parseAuthKey(r, "Customer-Key")
		if err != nil {
			err = errors.WithMessage(err, "customer key parse error")
			respJsonError(w, errAuth.withData(err.Error()))
			return
		}

		// authenticate contract owner
		contractOwnerAddr, err := util.RecoverAddress(billingKey.Msg, billingKey.Sig)
		if err != nil {
			respJsonError(w, errAuth.withData(err.Error()))
			return
		}

		// TODO:
		// 1. validate contract owner
		// 2. cache contract owner signature for performance

		logrus.WithFields(logrus.Fields{
			"contractOwnerAddr": contractOwnerAddr,
			"msg":               billingKey.Msg,
			"sig":               billingKey.Sig,
		}).Debug("Extract contract owner address from auth key")

		// authenticate customer
		customerAddr, err := util.RecoverAddress(customerKey.Msg, customerKey.Sig)
		if err != nil {
			respJsonError(w, errAuth.withData(err.Error()))
			return
		}

		// TODO: cache customer signature for performance

		logrus.WithFields(logrus.Fields{
			"customerAddr": customerAddr,
			"msg":          customerKey.Msg,
			"sig":          customerKey.Sig,
		}).Debug("Extract customer address from auth key")

		next.ServeHTTP(w, r)
	})
}
