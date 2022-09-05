package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/Conflux-Chain/web3pay-service/client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type keygenConfig struct {
	KeyType         string // "billing" or "customer" auth key
	AppCoinContract string // APP coin contract address
	PrivateKey      string // private key
}

var (
	kgconfig  keygenConfig
	genKeyCmd = &cobra.Command{
		Use:   "genkey",
		Short: "Generate billing or customer auth key",
		Run:   genAuthKey,
	}
)

func init() {
	rootCmd.AddCommand(genKeyCmd)

	// auth key type
	genKeyCmd.Flags().StringVarP(
		&kgconfig.KeyType, "type", "t", "billing", "auth key type",
	)

	// APP coin contract
	genKeyCmd.Flags().StringVarP(
		&kgconfig.AppCoinContract, "app", "a", "", "APP coin contract address",
	)
	genKeyCmd.MarkFlagRequired("app")

	// Private keyphrase
	genKeyCmd.Flags().StringVarP(
		&kgconfig.PrivateKey, "privk", "k", "", "private key",
	)
	genKeyCmd.MarkFlagRequired("privk")
}

func genAuthKey(cmd *cobra.Command, args []string) {
	msg := kgconfig.AppCoinContract
	if !strings.EqualFold(kgconfig.KeyType, "billing") {
		msg = fmt.Sprintf("%s_%v", msg, time.Now().Unix())
		kgconfig.KeyType = "customer"
	}

	key, err := client.BuildAuthKey(msg, kgconfig.PrivateKey)
	logrus.WithFields(logrus.Fields{
		"APPCoinContract": kgconfig.AppCoinContract,
		"AuthKeyType":     kgconfig.KeyType,
		"AuthKey":         key,
		"Message":         msg,
	}).WithError(err).Info("Auth key generated")
}
