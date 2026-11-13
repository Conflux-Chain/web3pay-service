package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Conflux-Chain/web3pay-service/client"
	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/service"
	"github.com/Conflux-Chain/web3pay-service/types"
	"github.com/openweb3/go-rpc-provider"
	"github.com/stretchr/testify/require"
)

func TestOpenweb3BillingMiddleware(t *testing.T) {
	type billingRequest struct {
		ID     json.RawMessage          `json:"id"`
		Method string                   `json:"method"`
		Params []service.BillingRequest `json:"params"`
		apiKey string
	}
	requests := make(chan billingRequest, 3)
	replies := make(chan *model.BusinessError, 3)
	replies <- model.ErrNil.WithData(&service.BillingReceipt{Balance: "100"})
	replies <- model.ErrInsufficentBalance
	replies <- model.ErrInternalServer.WithData("gateway unavailable")
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req billingRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		req.apiKey = r.Header.Get(types.AuthHeaderApiKey)
		requests <- req
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]interface{}{
			"jsonrpc": "2.0", "id": req.ID, "result": <-replies,
		}); err != nil {
			t.Error(err)
		}
	}))
	defer srv.Close()
	billingClient, err := client.NewBillingClient(client.BillingClientConfig{Gateway: srv.URL})
	require.NoError(t, err)
	option := NewOw3BillingMiddlewareOptionWithClient(billingClient)
	var status *BillingStatus
	next := Openweb3BillingMiddleware(option)(func(ctx context.Context, msg *rpc.JsonRpcMessage) *rpc.JsonRpcMessage {
		var ok bool
		status, ok = BillingStatusFromContext(ctx)
		require.True(t, ok)
		return msg
	})
	msg := &rpc.JsonRpcMessage{Method: "eth_getBalance"}
	ctx := context.WithValue(context.Background(), CtxApiKey, "test-key")
	for i := 0; i < 3; i++ {
		require.Same(t, msg, next(ctx, msg))
		req := <-requests
		require.Equal(t, "test-key", req.apiKey)
		require.Equal(t, "web3pay.Bill", req.Method)
		require.Len(t, req.Params, 1)
		require.Equal(t, "eth_getBalance", req.Params[0].ResourceId)
		require.False(t, req.Params[0].DryRun)
		switch i {
		case 0:
			require.True(t, status.Success())
			require.Equal(t, "100", status.Receipt.Balance)
		case 1:
			require.False(t, status.Success())
			var businessError *model.BusinessError
			require.ErrorAs(t, status.Error, &businessError)
			require.Equal(t, model.ErrInsufficentBalance.Code, businessError.Code)
		case 2:
			require.True(t, status.Success(), "a previously billed key retains access during a gateway failure")
			require.Error(t, status.Error)
		}
	}
	require.Same(t, msg, next(context.Background(), msg))
	require.False(t, status.Success())
	require.ErrorIs(t, status.Error, errInvalidApiKey)
}
