package contract

import (
	"context"
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/openweb3/web3go"
	"github.com/stretchr/testify/require"
)

func TestContractTransactionByHash(t *testing.T) {
	to := common.HexToAddress("0x1111111111111111111111111111111111111111")
	tx := types.NewTransaction(1, to, big.NewInt(10), 21000, big.NewInt(2), nil)
	encoded, err := tx.MarshalJSON()
	require.NoError(t, err)
	for _, tt := range []struct {
		name        string
		blockNumber interface{}
		missing     bool
		rpcError    bool
		malformed   bool
	}{
		{name: "pending"},
		{name: "mined", blockNumber: "0x64"},
		{name: "not found", missing: true},
		{name: "RPC error", rpcError: true},
		{name: "malformed transaction", malformed: true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				var req struct {
					ID     json.RawMessage `json:"id"`
					Method string          `json:"method"`
					Params []common.Hash   `json:"params"`
				}
				if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
					t.Error(err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				if req.Method != "eth_getTransactionByHash" || len(req.Params) != 1 || req.Params[0] != tx.Hash() {
					t.Error("unexpected transaction lookup", req)
				}
				var result map[string]interface{}
				if !tt.missing {
					if err := json.Unmarshal(encoded, &result); err != nil {
						t.Error(err)
					}
					result["blockNumber"] = tt.blockNumber
					if tt.malformed {
						delete(result, "gas")
					}
				}
				response := map[string]interface{}{"jsonrpc": "2.0", "id": req.ID, "result": result}
				if tt.rpcError {
					delete(response, "result")
					response["error"] = map[string]interface{}{"code": -32000, "message": "lookup unavailable"}
				}
				w.Header().Set("Content-Type", "application/json")
				if err := json.NewEncoder(w).Encode(response); err != nil {
					t.Error(err)
				}
			}))
			defer srv.Close()
			client, err := web3go.NewClient(srv.URL)
			require.NoError(t, err)
			backend, _ := NewClientForContract(client)
			found, pending, err := backend.TransactionByHash(context.Background(), tx.Hash())
			switch {
			case tt.missing:
				require.ErrorIs(t, err, ethereum.NotFound)
			case tt.rpcError:
				require.ErrorContains(t, err, "lookup unavailable")
			case tt.malformed:
				require.Error(t, err)
			default:
				require.NoError(t, err)
				require.Equal(t, tx.Hash(), found.Hash())
				require.Equal(t, tt.blockNumber == nil, pending)
			}
			if err != nil {
				require.Nil(t, found)
				require.False(t, pending)
			}
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			_, _, err = backend.TransactionByHash(ctx, tx.Hash())
			require.ErrorIs(t, err, context.Canceled)
		})
	}
}
