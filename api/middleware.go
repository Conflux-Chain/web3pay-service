package api

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"math"
	"net/http"

	mathutil "github.com/Conflux-Chain/go-conflux-util/math"
	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/service"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type reqCtxKey string

var (
	ctxKeyRequestId    reqCtxKey = "reqId"
	ctxKeyContractAddr reqCtxKey = "contractAddr"
	ctxKeyCustomerAddr reqCtxKey = "customerAddr"
)

func requestIdFromContext(ctx context.Context) uint64 {
	return ctx.Value(ctxKeyRequestId).(uint64)
}

func contractAddrFromContext(ctx context.Context) common.Address {
	return ctx.Value(ctxKeyContractAddr).(common.Address)
}

func customerAddrFromContext(ctx context.Context) common.Address {
	return ctx.Value(ctxKeyCustomerAddr).(common.Address)
}

func LogTracingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqId := mathutil.RandUint64(uint64(math.MaxUint32))
		newCtx := context.WithValue(r.Context(), ctxKeyRequestId, reqId)

		newReq := r.WithContext(newCtx)
		next.ServeHTTP(w, newReq)
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.WithFields(logrus.Fields{
			"reqID":        requestIdFromContext(r.Context()),
			"contractAddr": contractAddrFromContext(r.Context()),
			"customerAddr": customerAddrFromContext(r.Context()),
			"method":       r.Method,
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

func AuthMiddleware(r *mux.Router, chainSvc *service.BlockchainService) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			billingKey, err := parseAuthKey(r, "Billing-Key")
			if err != nil {
				err = errors.WithMessage(err, "billing key parsed error")
				respJsonError(w, model.ErrAuth.WithData(err.Error()))
				return
			}

			customerKey, err := parseAuthKey(r, "Customer-Key")
			if err != nil {
				err = errors.WithMessage(err, "customer key parse error")
				respJsonError(w, model.ErrAuth.WithData(err.Error()))
				return
			}

			// authenticate contract owner
			if !common.IsHexAddress(billingKey.Msg) { // `msg` part must be a valid hex address
				respJsonError(w, model.ErrAuth.WithData("invalid contract address"))
				return
			}

			ownerAddr, err := chainSvc.RecoverAddressBySignature(billingKey.Msg, billingKey.Sig)
			if err != nil {
				respJsonError(w, model.ErrAuth.WithData(err.Error()))
				return
			}

			contract, owner := common.HexToAddress(billingKey.Msg), common.HexToAddress(ownerAddr)
			if err := chainSvc.ValidateAppCoinContractOwner(contract, owner); err != nil {
				respJsonError(w, model.ErrAuth.WithData(err.Error()))
				return
			}

			// authenticate customer
			customerAddr, err := chainSvc.RecoverAddressBySignature(customerKey.Msg, customerKey.Sig)
			if err != nil {
				respJsonError(w, model.ErrAuth.WithData(err.Error()))
				return
			}

			customer := common.HexToAddress(customerAddr)

			// fill info to new context
			ctx := context.WithValue(r.Context(), ctxKeyContractAddr, contract)
			ctx = context.WithValue(ctx, ctxKeyCustomerAddr, customer)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
