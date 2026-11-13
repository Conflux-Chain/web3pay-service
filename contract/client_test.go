package contract

import (
	"bytes"
	"context"
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/holiman/uint256"
	"github.com/openweb3/web3go"
	"github.com/stretchr/testify/require"
)

func TestContractTransactionByHash(t *testing.T) {
	to := common.HexToAddress("0x1111111111111111111111111111111111111111")
	legacy := types.NewTransaction(1, to, big.NewInt(10), 21000, big.NewInt(2), nil)
	type testCase struct {
		name        string
		tx          *types.Transaction
		blockNumber interface{}
		missing     bool
		rpcError    bool
		malformed   bool
	}
	tests := []testCase{
		{name: "pending"},
		{name: "mined", blockNumber: "0x64"},
		{name: "not found", missing: true},
		{name: "RPC error", rpcError: true},
		{name: "malformed transaction", malformed: true},
	}
	accessList := types.AccessList{{Address: to, StorageKeys: []common.Hash{common.HexToHash("0x42")}}}
	for _, fixture := range []struct {
		name string
		tx   *types.Transaction
	}{
		{"access list", types.NewTx(&types.AccessListTx{
			ChainID: big.NewInt(71), Nonce: 2, GasPrice: big.NewInt(3), Gas: 50000,
			To: &to, Value: big.NewInt(11), Data: []byte{1, 2}, AccessList: accessList,
		})},
		{"dynamic fee", types.NewTx(&types.DynamicFeeTx{
			ChainID: big.NewInt(71), Nonce: 3, GasTipCap: big.NewInt(2), GasFeeCap: big.NewInt(5), Gas: 50000,
			To: &to, Value: big.NewInt(12), Data: []byte{3, 4}, AccessList: accessList,
		})},
		{"blob", types.NewTx(&types.BlobTx{
			ChainID: uint256.NewInt(71), Nonce: 4, GasTipCap: uint256.NewInt(2), GasFeeCap: uint256.NewInt(5), Gas: 50000,
			To: to, Value: uint256.NewInt(13), Data: []byte{5, 6}, AccessList: accessList,
			BlobFeeCap: uint256.NewInt(7), BlobHashes: []common.Hash{common.HexToHash("0x0100000000000000000000000000000000000000000000000000000000000042")},
		})},
		{"EIP-7702", types.NewTx(&types.SetCodeTx{
			ChainID: uint256.NewInt(71), Nonce: 5, GasTipCap: uint256.NewInt(2), GasFeeCap: uint256.NewInt(5), Gas: 100000,
			To: to, Value: uint256.NewInt(14), Data: []byte{7, 8}, AccessList: accessList,
			AuthList: []types.SetCodeAuthorization{{ChainID: *uint256.NewInt(71), Address: to, Nonce: 6}},
		})},
	} {
		tests = append(tests,
			testCase{name: fixture.name + " pending", tx: fixture.tx},
			testCase{name: fixture.name + " mined", tx: fixture.tx, blockNumber: "0x64"},
		)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := tt.tx
			if tx == nil {
				tx = legacy
			}
			encoded, err := tx.MarshalJSON()
			require.NoError(t, err)
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
				require.Equal(t, tx.Type(), found.Type())
				require.Equal(t, tx.Hash(), found.Hash())
				require.Equal(t, tx.ChainId(), found.ChainId())
				require.Equal(t, tx.Nonce(), found.Nonce())
				require.Equal(t, tx.To(), found.To())
				require.Equal(t, tx.Value(), found.Value())
				require.Equal(t, tx.Gas(), found.Gas())
				require.Equal(t, tx.GasFeeCap(), found.GasFeeCap())
				require.Equal(t, tx.GasTipCap(), found.GasTipCap())
				require.True(t, bytes.Equal(tx.Data(), found.Data()))
				require.Equal(t, tx.AccessList(), found.AccessList())
				if tx.Type() == types.BlobTxType {
					require.Equal(t, tx.BlobGasFeeCap(), found.BlobGasFeeCap())
					require.Equal(t, tx.BlobHashes(), found.BlobHashes())
				}
				if tx.Type() == types.SetCodeTxType {
					require.Equal(t, tx.SetCodeAuthorizations(), found.SetCodeAuthorizations())
				}
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
