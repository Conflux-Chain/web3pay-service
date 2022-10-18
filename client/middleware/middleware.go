package middleware

import (
	"context"
	"net/http"

	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/openweb3/go-rpc-provider"
)

const (
	// API key
	CtxApiKey = CtxKey("Web3Pay-Api-Key")

	// default API key LRU cache size
	apiKeyCacheSize = 5000
)

var (
	// errors
	errInvalidApiKey = model.ErrAuth.WithData("Invalid API key")
)

type (
	// CtxKey context key
	CtxKey string

	// authKeyProvider retrieves auth key from context
	authKeyProvider func(ctx context.Context) (string, bool)
	// httpContextInjector injects data into context to return new context
	httpContextInjector = func(ctx context.Context, r *http.Request) context.Context

	// openweb3 middleware
	Ow3Middleware = rpc.HandleCallMsgMiddleware
	// net/http middleware
	HttpMiddleware = func(next http.Handler) http.Handler
)

// HttpInjectContextMiddleware injects contextual data
func HttpInjectContextMiddleware(injectors ...httpContextInjector) HttpMiddleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			for i := range injectors {
				ctx = injectors[i](ctx, r)
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// handy utility functions

// ApiKeyContextInjector returns context injector by extracting API key from request
func ApiKeyContextInjector(kextractor func(r *http.Request) string) httpContextInjector {
	return func(ctx context.Context, r *http.Request) context.Context {
		return context.WithValue(ctx, CtxApiKey, kextractor(r))
	}
}

// GetApiKeyFromContext gets API key from context
func GetApiKeyFromContext(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(CtxApiKey).(string)
	return val, ok
}
