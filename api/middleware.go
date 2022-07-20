package api

import (
	"context"
	"math"
	"net/http"

	mathutil "github.com/Conflux-Chain/go-conflux-util/math"
	"github.com/Conflux-Chain/web3pay-service/util"
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

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// extract signature from header
		appContract, ownerSignature := r.Header.Get("App-Contract"), r.Header.Get("Owner-Signature")
		if len(appContract) == 0 || len(ownerSignature) == 0 {
			respJsonError(w, errAuth.withData("incomplete contract signature information"))
			return
		}

		appNonce, customerSignature := r.Header.Get("App-Nonce"), r.Header.Get("Customer-Signature")
		if len(appNonce) == 0 || len(customerSignature) == 0 {
			respJsonError(w, errAuth.withData("incomplete customer signature information"))
			return
		}

		// authenticate the APP contract owner
		contractOwnerAddr, err := util.RecoverAddress(appContract, ownerSignature)
		if err != nil {
			respJsonError(w, errAuth.withData(err.Error()))
			return
		}

		// TODO:
		// 1. validate contract owner
		// 2. cache contract owner signature for performance

		logrus.WithFields(logrus.Fields{
			"contractOwnerAddr": contractOwnerAddr,
			"appContract":       appContract,
			"ownerSignature":    ownerSignature,
		}).Debug("Extract contract owner address from auth signature")

		// authenticate the customer
		customerAddr, err := util.RecoverAddress(appNonce, customerSignature)
		if err != nil {
			respJsonError(w, errAuth.withData(err.Error()))
			return
		}

		// TODO: cache customer signature for performance

		logrus.WithFields(logrus.Fields{
			"customerAddr":      customerAddr,
			"appNonce":          appNonce,
			"customerSignature": customerSignature,
		}).Debug("Extract customer address from auth signature")

		next.ServeHTTP(w, r)
	})
}
