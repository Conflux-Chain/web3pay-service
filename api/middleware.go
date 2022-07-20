package api

import (
	"context"
	"math"
	"net/http"

	mathutil "github.com/Conflux-Chain/go-conflux-util/math"
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
