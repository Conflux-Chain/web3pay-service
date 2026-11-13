package contract

import (
	"context"
	"encoding/json"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/openweb3/web3go"
)

// ClientForContract adds geth's transaction lookup to the web3go contract backend.
type ClientForContract struct {
	*web3go.ClientForContract
	client *web3go.Client
}

var _ bind.ContractBackend = (*ClientForContract)(nil)

func NewClientForContract(client *web3go.Client) (*ClientForContract, bind.SignerFn) {
	backend, signer := client.ToClientForContract()
	return &ClientForContract{ClientForContract: backend, client: client}, signer
}

func (c *ClientForContract) TransactionByHash(ctx context.Context, hash common.Hash) (*types.Transaction, bool, error) {
	if err := ctx.Err(); err != nil {
		return nil, false, err
	}
	var result json.RawMessage
	if err := c.client.CallContext(ctx, &result, "eth_getTransactionByHash", hash); err != nil {
		return nil, false, err
	}
	if len(result) == 0 {
		return nil, false, ethereum.NotFound
	}
	var tx *types.Transaction
	if err := json.Unmarshal(result, &tx); err != nil {
		return nil, false, err
	}
	if tx == nil {
		return nil, false, ethereum.NotFound
	}
	var metadata struct {
		BlockNumber *hexutil.Big `json:"blockNumber"`
	}
	if err := json.Unmarshal(result, &metadata); err != nil {
		return nil, false, err
	}
	return tx, metadata.BlockNumber == nil, nil
}
