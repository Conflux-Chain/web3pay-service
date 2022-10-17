package cmd

import (
	"strings"

	"github.com/Conflux-Chain/web3pay-service/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type keygenConfig struct {
	KeyType     string // "billing" or "api" auth key
	AppContract string // APP contract address
	PrivateKey  string // private key
}

var (
	kgconfig  keygenConfig
	genKeyCmd = &cobra.Command{
		Use:   "genkey",
		Short: "Generate billing or API auth key",
		Run:   genAuthKey,
	}
)

func init() {
	rootCmd.AddCommand(genKeyCmd)

	// auth key type
	genKeyCmd.Flags().StringVarP(
		&kgconfig.KeyType, "type", "t", "billing", "(billing or api) auth key",
	)

	// APP contract
	genKeyCmd.Flags().StringVarP(
		&kgconfig.AppContract, "app", "a", "", "APP contract address",
	)
	genKeyCmd.MarkFlagRequired("app")

	// Private keyphrase
	genKeyCmd.Flags().StringVarP(
		&kgconfig.PrivateKey, "privk", "k", "", "private key",
	)
	genKeyCmd.MarkFlagRequired("privk")
}

func genAuthKey(cmd *cobra.Command, args []string) {
	keyBuilder := util.BuildBillingKey
	if strings.EqualFold(kgconfig.KeyType, "api") { // API key
		keyBuilder = util.BuildApiKey
	} else {
		kgconfig.KeyType = "billing"
	}

	authKey, err := keyBuilder(kgconfig.AppContract, kgconfig.PrivateKey)
	logrus.WithFields(logrus.Fields{
		"APPContract": kgconfig.AppContract,
		"AuthKeyType": kgconfig.KeyType,
		"AuthKey":     authKey,
	}).WithError(err).Info("Auth key generated")
}
