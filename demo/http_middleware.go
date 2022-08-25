package demo

import (
	"context"
	"net/http"
	"strings"
)

type CtxKey string

const (
	CtxCustomerKey = CtxKey("Web3Pay-Customer-Key")
)

// Inject values into context for static RPC call middlewares, e.g. billing
func httpMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, CtxCustomerKey, GetCustomerKey(r))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetCustomerKey(r *http.Request) string {
	if r == nil || r.URL == nil {
		return ""
	}

	// customer key path pattern:
	// http://example.com/${customerKey}...
	key := strings.TrimLeft(r.URL.Path, "/")
	if idx := strings.Index(key, "/"); idx > 0 {
		key = key[:idx]
	}

	return key
}

func GetCustomerKeyFromContext(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(CtxCustomerKey).(string)
	return val, ok
}
