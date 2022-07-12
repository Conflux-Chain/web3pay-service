/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"

	"github.com/spf13/cobra"
)

// Returns a signature string
func PersonalSign(message string, privateKey *ecdsa.PrivateKey) (string, common.Hash, error) {
	fullMessage := BuildPersonalSignMessage(message)
	hash := crypto.Keccak256Hash([]byte(fullMessage))
	signatureBytes, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return "", hash, err
	}
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
	decode[64] -= 27
	hash := crypto.Keccak256Hash([]byte(BuildPersonalSignMessage(message)))
	sigPublicKey, err := crypto.SigToPub(hash.Bytes(), decode)
	if err != nil {
		return "", err
	}
	address := crypto.PubkeyToAddress(*sigPublicKey).Hex()
	return address, nil
}

// ecrecoverCmd represents the ecrecover command
var ecrecoverCmd = &cobra.Command{
	Use:        "recover signature message",
	Short:      "Recover address",
	Long:       `Recover address from message and its signature`,
	Args:       cobra.ExactArgs(2),
	ArgAliases: []string{"signature", "message"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ecrecover called", args)
		signature, message := args[0], args[1]
		address, err := RecoverAddress(message, signature)
		if err != nil {
			log.Fatal("recover fail", err)
		}
		fmt.Printf("recovered address: %s\n", address)
	},
}

var signCmd = &cobra.Command{
	Use:        "sign privateKey message",
	Short:      "Sign message",
	Long:       `Sign a message with private key in metamask style`,
	Args:       cobra.ExactArgs(2),
	ArgAliases: []string{"privateKey", "message"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sign called", args)
		pkString, message := args[0], args[1]
		privateKey, err := crypto.HexToECDSA(pkString)
		if err != nil {
			log.Fatal(err)
		}
		publicKey := privateKey.Public()
		publicKeyECDSA := publicKey.(*ecdsa.PublicKey)
		addressSig := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
		fmt.Printf("sign with address %s\n", addressSig)

		input := message
		signature, _, err := PersonalSign(input, privateKey)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("signature", signature, "length", len(signature))
	},
}

func init() {
	rootCmd.AddCommand(ecrecoverCmd)
	rootCmd.AddCommand(signCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ecrecoverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ecrecoverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
