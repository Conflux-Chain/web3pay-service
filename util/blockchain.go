package util

import (
	"crypto/ecdsa"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

// IsZeroAddress validate if it's a 0 address
func IsZeroAddress(iaddress interface{}) bool {
	var address common.Address
	switch v := iaddress.(type) {
	case string:
		address = common.HexToAddress(v)
	case common.Address:
		address = v
	default:
		return false
	}

	zeroAddressBytes := common.FromHex("0x0000000000000000000000000000000000000000")
	addressBytes := address.Bytes()
	return reflect.DeepEqual(addressBytes, zeroAddressBytes)
}

func EcdsaPrivateKeyFromString(privateKeyStr string) (*ecdsa.PrivateKey, error) {
	if len(privateKeyStr) >= 2 && privateKeyStr[0:2] == "0x" {
		privateKeyStr = privateKeyStr[2:]
	}

	opPrivateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, errors.WithMessage(err, "invalid HEX format of private key")
	}

	return opPrivateKey, nil
}

func AddressFromEcdsaPrivateKey(privateKey *ecdsa.PrivateKey) (common.Address, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, errors.New("not a valid ECDSA private key")
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA), nil
}

func AddressFromEcdsaPrivateKeyString(privateKeyStr string) (common.Address, error) {
	privKey, err := EcdsaPrivateKeyFromString(privateKeyStr)
	if err != nil {
		return common.Address{}, errors.WithMessage(err, "failed to parse private key")
	}

	return AddressFromEcdsaPrivateKey(privKey)
}
