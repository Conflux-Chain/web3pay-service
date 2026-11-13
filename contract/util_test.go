package contract

import (
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/types"
	"github.com/stretchr/testify/require"
)

func TestAppCallThroughWeb3go(t *testing.T) {
	appAddress := common.HexToAddress("0x1111111111111111111111111111111111111111")
	shopAddress := common.HexToAddress("0x2222222222222222222222222222222222222222")
	type rpcRequest struct {
		ID     json.RawMessage   `json:"id"`
		Method string            `json:"method"`
		Params []json.RawMessage `json:"params"`
	}
	requests := make(chan rpcRequest, 1)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req rpcRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		requests <- req
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]interface{}{
			"jsonrpc": "2.0", "id": req.ID,
			"result": hexutil.Encode(common.LeftPadBytes(shopAddress.Bytes(), 32)),
		}); err != nil {
			t.Error(err)
		}
	}))
	defer srv.Close()

	client, err := web3go.NewClient(srv.URL)
	require.NoError(t, err)
	backend, _ := NewClientForContract(client)
	app, err := NewApp(appAddress, backend)
	require.NoError(t, err)
	shop, err := app.CardShop(&bind.CallOpts{BlockNumber: big.NewInt(100)})
	require.NoError(t, err)
	require.Equal(t, shopAddress, shop)

	req := <-requests
	require.Equal(t, "eth_call", req.Method)
	require.GreaterOrEqual(t, len(req.Params), 2)
	var call map[string]interface{}
	require.NoError(t, json.Unmarshal(req.Params[0], &call))
	require.Equal(t, appAddress.Hex(), common.HexToAddress(call["to"].(string)).Hex())
	require.Equal(t, hexutil.Encode(crypto.Keccak256([]byte("cardShop()"))[:4]), call["data"])
	require.JSONEq(t, `"0x64"`, string(req.Params[1]))
}

func TestUnpackAppDepositCompatibility(t *testing.T) {
	appABI, err := AppMetaData.GetAbi()
	require.NoError(t, err)
	operator := common.HexToAddress("0x1111111111111111111111111111111111111111")
	receiver := common.HexToAddress("0x2222222222222222222222222222222222222222")
	log := &types.Log{
		Topics: []common.Hash{
			common.HexToHash("0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7"),
			common.BytesToHash(operator.Bytes()), common.BytesToHash(receiver.Bytes()), common.BigToHash(big.NewInt(1)),
		},
		Data: common.LeftPadBytes(big.NewInt(12345).Bytes(), 32),
	}
	event, err := UnpackAppDeposit(appABI, log)
	require.NoError(t, err)
	require.Equal(t, operator, event.Operator)
	require.Equal(t, receiver, event.Receiver)
	require.Equal(t, int64(1), event.TokenId.Int64())
	require.Equal(t, int64(12345), event.Amount.Int64())

	log.Topics[0] = common.Hash{}
	_, err = UnpackAppDeposit(appABI, log)
	require.ErrorIs(t, err, errEventSigMismatch)
}
