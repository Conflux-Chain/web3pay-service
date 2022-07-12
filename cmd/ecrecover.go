package cmd

import (
	"crypto/ecdsa"
	"github.com/Conflux-Chain/web3pay-service/util"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sirupsen/logrus"
	"log"

	"github.com/spf13/cobra"
)

// ecrecoverCmd represents the ecrecover command
var ecrecoverCmd = &cobra.Command{
	Use:        "recover <signature> <message>",
	Short:      "Recover address",
	Long:       `Recover address from message and its signature`,
	Args:       cobra.ExactArgs(2),
	ArgAliases: []string{"signature", "message"},
	Run: func(cmd *cobra.Command, args []string) {
		logrus.WithFields(logrus.Fields{"args": args}).Debug("recover called")
		signature, message := args[0], args[1]
		address, err := util.RecoverAddress(message, signature)
		if err != nil {
			logrus.WithFields(logrus.Fields{"error": err}).Fatal("recover fail", err)
		}
		logrus.WithFields(logrus.Fields{"address": address}).Info("recovered")
	},
}

var signCmd = &cobra.Command{
	Use:        "sign <privateKey> <message>",
	Short:      "Sign message",
	Long:       `Sign a message with private key in metamask style`,
	Args:       cobra.ExactArgs(2),
	ArgAliases: []string{"privateKey", "message"},
	Run: func(cmd *cobra.Command, args []string) {
		logrus.WithFields(logrus.Fields{"args": args}).Debug("sign called")
		pkString, message := args[0], args[1]
		privateKey, err := crypto.HexToECDSA(pkString)
		if err != nil {
			log.Fatal(err)
		}
		publicKey := privateKey.Public()
		publicKeyECDSA := publicKey.(*ecdsa.PublicKey)
		addressSig := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
		logrus.WithFields(logrus.Fields{"address": addressSig}).Info("sign with ")

		input := message
		signature, _, err := util.PersonalSign(input, privateKey)
		if err != nil {
			log.Fatal(err)
		}

		logrus.WithFields(logrus.Fields{"signature ": signature, "length ": len(signature)}).Info("result")
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
