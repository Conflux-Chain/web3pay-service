package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"time"

	mathutil "github.com/Conflux-Chain/go-conflux-util/math"
	"github.com/Conflux-Chain/web3pay-service/client/jsonrpc"
	"github.com/Conflux-Chain/web3pay-service/metrics"
	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/service"
	"github.com/Conflux-Chain/web3pay-service/types.go"
	"github.com/Conflux-Chain/web3pay-service/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

type reqCtxKey string

var (
	ctxKeyRequestId       reqCtxKey = "reqId"
	ctxKeyMetricCollector reqCtxKey = "metricCollector"
	ctxKeyContractAddr    reqCtxKey = "contractAddr"
	ctxKeyCustomerAddr    reqCtxKey = "customerAddr"
	ctxKeyJsonRpcMsg      reqCtxKey = "jsonRpcMsg"
)

func metricCollectorFromContext(ctx context.Context) metrics.Collector {
	if c, ok := ctx.Value(ctxKeyMetricCollector).(metrics.Collector); ok {
		return c
	}

	return nil
}

func requestIdFromContext(ctx context.Context) string {
	return ctx.Value(ctxKeyRequestId).(string)
}

func contractAddrFromContext(ctx context.Context) common.Address {
	return ctx.Value(ctxKeyContractAddr).(common.Address)
}

func customerAddrFromContext(ctx context.Context) common.Address {
	return ctx.Value(ctxKeyCustomerAddr).(common.Address)
}

func jsonRpcMessageFromContext(ctx context.Context) *jsonrpc.RPCRequest {
	if msg, ok := ctx.Value(ctxKeyJsonRpcMsg).(*jsonrpc.RPCRequest); ok {
		return msg
	}

	return nil
}

func LogTracingMiddleware(prefix string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqId := mathutil.RandUint64(uint64(math.MaxUint32))
			newCtx := context.WithValue(r.Context(), ctxKeyRequestId, fmt.Sprintf("%v#%d", prefix, reqId))
			newReq := r.WithContext(newCtx)
			next.ServeHTTP(w, newReq)
		})
	}
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := logrus.WithFields(logrus.Fields{
			"reqID":      requestIdFromContext(r.Context()),
			"method":     r.Method,
			"requestUri": r.RequestURI,
		})

		logger.Debug("RPC enter")
		start := time.Now()
		next.ServeHTTP(w, r)
		logger.WithField("elapsed", time.Since(start)).Debug("RPC leave")
	})
}

func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// inject metrics collector
		collector := metrics.NewRpcCollector()
		ctx := context.WithValue(r.Context(), ctxKeyMetricCollector, collector)

		nw := negroni.NewResponseWriter(w)
		next.ServeHTTP(nw, r.WithContext(ctx))

		collector.Collect(metrics.CollectKeyStatusCode, nw.Status())
		metrics.RPC.UpdateWithCollector(collector)
	})
}

func parseBillingKey(r *http.Request) (*types.BillingAuthKey, error) {
	billingKeyStr := r.Header.Get(types.AuthHeaderBillingKey)
	if len(billingKeyStr) < 260 {
		return nil, errors.New("key bytes too short")
	}

	return util.ParseBillingKey(billingKeyStr)
}

func parseApiKey(r *http.Request) (*types.ApiAuthKey, error) {
	apiKeyStr := r.Header.Get(types.AuthHeaderApiKey)
	if len(apiKeyStr) < 89 {
		return nil, errors.New("key bytes too short")
	}

	return util.ParseApiKey(apiKeyStr)
}

type authErrorHandler func(ctx context.Context, w http.ResponseWriter, err error)

func AuthMiddleware(r *mux.Router, chainSvc *service.BlockchainService, handler authErrorHandler) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			billingKey, err := parseBillingKey(r)
			if err != nil {
				err = errors.WithMessage(err, "billing key parsed error")
				handler(ctx, w, model.ErrAuth.WithData(err.Error()))
				return
			}

			// authenticate contract operator
			op, err := util.GetAddrByBillingAuthKey(billingKey)
			if err != nil {
				handler(ctx, w, model.ErrAuth.WithData(err.Error()))
				return
			}

			// `msg` is the contract address
			contract := common.HexToAddress(billingKey.Msg)
			if err := chainSvc.ValidateAppOperator(contract, op); err != nil {
				handler(ctx, w, model.ErrAuth.WithData(err.Error()))
				return
			}

			apiKey, err := parseApiKey(r)
			if err != nil {
				err = errors.WithMessage(err, "API key parse error")
				handler(ctx, w, model.ErrAuth.WithData(err.Error()))
				return
			}

			// authenticate customer
			customer, err := util.GetAddrByApiAuthKey(contract.String(), apiKey)
			if err != nil {
				handler(ctx, w, model.ErrAuth.WithData(err.Error()))
				return
			}

			// fill info to new context
			ctx = context.WithValue(ctx, ctxKeyContractAddr, contract)
			ctx = context.WithValue(ctx, ctxKeyCustomerAddr, customer)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func JsonRpcValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req struct {
			jsonrpc.RPCRequest
			ID *int `json:"id,omitempty"`
		}

		bodyData, err := jsonUnmarshalRequestBody(r.Body, &req)
		if err != nil {
			respJsonRpcError(ctx, w, errors.WithMessage(err, "failed to parse request body"))
			return
		}

		if req.ID == nil {
			err := errors.WithMessage(errors.New("id not provided"), "invalid JSON-RPC request")
			respJsonRpcError(ctx, w, err)
			return
		}

		if len(req.Method) == 0 {
			err := errors.WithMessage(errors.New("method not provided"), "invalid JSON-RPC request")
			respJsonRpcError(ctx, w, err)
			return
		}

		metricCollectRpcModule(ctx, req.Method)

		// Set a new body with the same data we read before. This is crucial since we need to read it again later.
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyData))

		req.RPCRequest.ID = *req.ID
		ctx = context.WithValue(ctx, ctxKeyJsonRpcMsg, &req.RPCRequest)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func jsonUnmarshalRequestBody(reqBody io.ReadCloser, ptr interface{}) ([]byte, error) {
	bytes, err := ioutil.ReadAll(reqBody)
	defer reqBody.Close()

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bytes), ptr); err != nil {
		return nil, err
	}

	return bytes, nil
}

func metricCollectRpcError(ctx context.Context, err error) {
	if collector := metricCollectorFromContext(ctx); collector != nil {
		collector.Collect(metrics.CollectKeyRPCError, err)
	}
}

func metricCollectRpcModule(ctx context.Context, module string) {
	if collector := metricCollectorFromContext(ctx); collector != nil {
		collector.Collect(metrics.CollectKeyRPCModule, module)
	}
}
