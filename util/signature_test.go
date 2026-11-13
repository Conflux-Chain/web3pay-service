package util

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/openweb3/web3go/signers"
	"github.com/stretchr/testify/require"
)

const testSigningKey = "0000000000000000000000000000000000000000000000000000000000000001"

// These vectors were recorded with geth v1.15.11 and web3go v0.3.3.
func TestPersonalSignCompatibility(t *testing.T) {
	key, err := EcdsaPrivateKeyFromString(testSigningKey)
	require.NoError(t, err)
	sig, hash, err := PersonalSign("web3pay dependency upgrade", key)
	require.NoError(t, err)
	require.Equal(t, "0xee595b2ca3681811965898f716e876b608767267579aee1fce67676d627a8c01", hash.Hex())
	require.Equal(t, "0x558503f5c88d771c3062bfd1045a9490f992c59dfcba64fac4b0f687ea2c46fa74d68d5e41ba1ed0ad1c3bb914d6db5b6ecb32be2a726982b5a20145d70e9ce21b", sig)
	address, err := RecoverAddress("web3pay dependency upgrade", sig)
	require.NoError(t, err)
	require.Equal(t, "0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf", address)
}

func TestTransactionSigningCompatibility(t *testing.T) {
	key, err := EcdsaPrivateKeyFromString(testSigningKey)
	require.NoError(t, err)
	signer := signers.NewPrivateKeySigner(key)
	to := common.HexToAddress("0x3535353535353535353535353535353535353535")
	tests := []struct {
		name    string
		chainID int64
		tx      *types.Transaction
		raw     string
	}{
		{
			name: "legacy", chainID: 1,
			tx:  types.NewTransaction(9, to, big.NewInt(1000000000000000000), 21000, big.NewInt(20000000000), nil),
			raw: "0xf86c098504a817c800825208943535353535353535353535353535353535353535880de0b6b3a76400008025a0c44b8c8b4f0b9ac6a006efb6ef7b4726114e1502cef33614b9186709364d1757a0393a09d1679f54c8ad8775a112b524d67cf1760513f99eb0eaee991c9cfe21b9",
		},
		{
			name: "dynamic fee", chainID: 71,
			tx: types.NewTx(&types.DynamicFeeTx{
				ChainID: big.NewInt(71), Nonce: 7, To: &to, Value: big.NewInt(1234),
				Gas: 25000, GasTipCap: big.NewInt(1000000000), GasFeeCap: big.NewInt(30000000000),
			}),
			raw: "0x02f86d4707843b9aca008506fc23ac008261a89435353535353535353535353535353535353535358204d280c001a011ebd1ab22569690136ba2dd3aefb8c8cf637a59d039ec9ffb5565a98c3a0321a07ebface7d54d4e6b70cf792e36cea39a533ea70197ad33b5b13c8fac04f1f66f",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chainID := big.NewInt(tt.chainID)
			signed, err := signer.SignTransaction(tt.tx, chainID)
			require.NoError(t, err)
			raw, err := signed.MarshalBinary()
			require.NoError(t, err)
			require.Equal(t, tt.raw, hexutil.Encode(raw))
			sender, err := types.Sender(types.LatestSignerForChainID(chainID), signed)
			require.NoError(t, err)
			require.Equal(t, signer.Address(), sender)
		})
	}
}
