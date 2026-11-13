package jsonrpc_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/openweb3/go-rpc-provider"
	"github.com/stretchr/testify/require"
)

func TestProviderBatchResponseMatching(t *testing.T) {
	for _, missing := range []bool{false, true} {
		name := "reversed responses"
		if missing {
			name = "missing response"
		}
		t.Run(name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				var requests []struct {
					ID json.RawMessage `json:"id"`
				}
				if err := json.NewDecoder(r.Body).Decode(&requests); err != nil || len(requests) != 2 {
					t.Error("expected a two-item JSON-RPC batch", err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				replies := []map[string]interface{}{
					{"jsonrpc": "2.0", "id": requests[1].ID, "result": "second"},
				}
				if !missing {
					replies = append(replies, map[string]interface{}{"jsonrpc": "2.0", "id": requests[0].ID, "result": "first"})
				}
				w.Header().Set("Content-Type", "application/json")
				if err := json.NewEncoder(w).Encode(replies); err != nil {
					t.Error(err)
				}
			}))
			defer srv.Close()
			client, err := rpc.DialHTTP(srv.URL)
			require.NoError(t, err)
			defer client.Close()
			var first, second string
			batch := []rpc.BatchElem{{Method: "first", Result: &first}, {Method: "second", Result: &second}}
			require.NoError(t, client.BatchCallContext(context.Background(), batch))
			require.NoError(t, batch[1].Error)
			require.Equal(t, "second", second)
			if missing {
				require.Error(t, batch[0].Error)
				require.Empty(t, first)
			} else {
				require.NoError(t, batch[0].Error)
				require.Equal(t, "first", first)
			}
		})
	}
}
