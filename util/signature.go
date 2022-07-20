package util

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

// Returns a signature string
func PersonalSign(message string, privateKey *ecdsa.PrivateKey) (string, common.Hash, error) {
	fullMessage := BuildPersonalSignMessage(message)
	hash := crypto.Keccak256Hash([]byte(fullMessage))

	signatureBytes, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return "", hash, err
	}

	// metamask/ethers uses V at the end. EIP-191.
	// refer: https://stackoverflow.com/questions/61682191/go-ethereum-sign-provides-different-signature-than-nodejs-ethers
	signatureBytes[64] += 27
	return hexutil.Encode(signatureBytes), hash, nil
}

func BuildPersonalSignMessage(message string) string {
	//It's metamask personal sign format.
	fullMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	return fullMessage
}

func RecoverAddress(message string, signature string) (string, error) {
	decode, err := hexutil.Decode(signature)
	if err != nil {
		return "", errors.WithMessage(err, "failed to decode signaute")
	}

	if len(decode) < 64 {
		return "", errors.New("no enough signature bytes")
	}

	decode[64] -= 27
	hash := crypto.Keccak256Hash([]byte(BuildPersonalSignMessage(message)))
	sigPublicKey, err := crypto.SigToPub(hash.Bytes(), decode)
	if err != nil {
		return "", errors.WithMessage(err, "failed to recover pubkey from signature")
	}

	address := crypto.PubkeyToAddress(*sigPublicKey).Hex()
	return address, nil
}
