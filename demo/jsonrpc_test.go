package demo

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Conflux-Chain/web3pay-service/client/middleware"
	"github.com/ethereum/go-ethereum/node"
	"github.com/openweb3/go-rpc-provider"
	"github.com/stretchr/testify/require"
)

func TestDemoRPCWithGzipAndBillingContext(t *testing.T) {
	rpcServer := rpc.NewServer()
	defer rpcServer.Stop()
	require.NoError(t, rpcServer.RegisterName("demo", &jrDemoApi{}))
	handler := node.NewHTTPHandlerStack(rpcServer, []string{"*"}, []string{"*"}, nil, false)
	handler = middleware.HttpInjectContextMiddleware(middleware.ApiKeyContextInjector(GetApiKey))(handler)
	for _, encoding := range []string{"", "gzip"} {
		t.Run("encoding="+encoding, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/test-key", strings.NewReader(`{"jsonrpc":"2.0","id":1,"method":"demo_testBilling","params":[]}`))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Accept-Encoding", encoding)
			ctx := context.WithValue(req.Context(), middleware.CtxKeyBillingStatus,
				middleware.NewBillingStatusWithReceipt("test-key", nil))
			recorder := httptest.NewRecorder()
			handler.ServeHTTP(recorder, req.WithContext(ctx))
			require.Equal(t, http.StatusOK, recorder.Code)
			var body io.Reader = recorder.Body
			if encoding == "gzip" {
				require.Equal(t, "gzip", recorder.Header().Get("Content-Encoding"))
				reader, err := gzip.NewReader(body)
				require.NoError(t, err)
				defer reader.Close()
				body = reader
			}
			var response struct {
				ID     int             `json:"id"`
				Result string          `json:"result"`
				Error  json.RawMessage `json:"error"`
			}
			require.NoError(t, json.NewDecoder(body).Decode(&response))
			require.Equal(t, 1, response.ID)
			require.Equal(t, "billing succeeded", response.Result)
			require.Empty(t, response.Error)
		})
	}
}
