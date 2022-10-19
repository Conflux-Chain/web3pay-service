package cmd

import (
	"strings"

	viperutil "github.com/Conflux-Chain/go-conflux-util/viper"
	"github.com/Conflux-Chain/web3pay-service/client"
	"github.com/Conflux-Chain/web3pay-service/demo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type demoConfig struct {
	client.BillingClientConfig         `mapstructure:",squash"`
	client.VipSubscriptionClientConfig `mapstructure:",squash"`
	ApiKey                             string
	ServerPort                         int
	RpcStyle                           string
	PayMode                            string
}

var (
	demoConf = demoConfig{}

	demoCmd = &cobra.Command{
		Use:   "demo",
		Short: "Run Web3Pay service provider and consumer demo",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	demoProviderCmd = &cobra.Command{
		Use:   "provider",
		Short: "Run Web3Pay demo service provider",
		Run:   startProvider,
	}

	demoConsumerCmd = &cobra.Command{
		Use:   "consumer",
		Short: "Run Web3Pay demo service consumer",
		Run:   startConsumer,
	}
)

func init() {
	demoCmd.AddCommand(demoProviderCmd)
	demoCmd.AddCommand(demoConsumerCmd)
	rootCmd.AddCommand(demoCmd)

	flags := demoCmd.PersistentFlags()

	// billing gateway URL
	flags.StringVarP(
		&demoConf.Gateway, "billing-gateway", "g", "", "billing gateway URL",
	)
	viper.BindPFlag("demo.gateway", flags.Lookup("billing-gateway"))

	// billing key
	flags.StringVarP(
		&demoConf.BillingKey, "billing-key", "b", "",
		"billing key (use 'ecrecover' command to generate one)",
	)
	viper.BindPFlag("demo.billingKey", flags.Lookup("billing-key"))

	// blockchain RPC endpoint for VIP subscription
	flags.StringVarP(
		&demoConf.ChainRpcUrl, "sub-chain-rpc-url", "r", "",
		"blockchain RPC endpoint for VIP subscription",
	)
	viper.BindPFlag("demo.chainRpcUrl", flags.Lookup("sub-chain-rpc-url"))

	// App contract address for VIP subscription
	flags.StringVarP(
		&demoConf.AppContract, "sub-app-contract", "c", "",
		"App contract for VIP subscription",
	)
	viper.BindPFlag("demo.appContract", flags.Lookup("sub-app-contract"))

	// API key
	flags.StringVarP(
		&demoConf.ApiKey, "api-key", "a", "",
		"API key (use 'genkey' command to generate one)",
	)
	viper.BindPFlag("demo.apiKey", flags.Lookup("api-key"))

	// demo provider server port
	flags.IntVarP(
		&demoConf.ServerPort, "port", "p", 38258,
		"listening port for the demo billed service provider",
	)
	viper.BindPFlag("demo.serverPort", flags.Lookup("port"))

	// RPC style
	flags.StringVarP(
		&demoConf.RpcStyle, "style", "s", "jsonrpc",
		"RPC style for the billed service demo, available options are 'jsonrpc' and 'restful'",
	)
	viper.BindPFlag("demo.rpcStyle", flags.Lookup("style"))

	// Mode
	flags.StringVarP(
		&demoConf.PayMode, "mode", "m", "billing",
		"Payment mode, available options are 'subscription' and 'billing'",
	)
	viper.BindPFlag("demo.payMode", flags.Lookup("mode"))
}

func startProvider(cmd *cobra.Command, args []string) {
	// load config from env vars or flags
	viperutil.MustUnmarshalKey("demo", &demoConf)

	if strings.EqualFold(demoConf.PayMode, "subscription") {
		if len(demoConf.ChainRpcUrl) == 0 {
			logrus.Info("Blockchain RPC endpoint must be provided")
			return
		}

		if len(demoConf.AppContract) == 0 {
			logrus.Info("Subscription App contract must be provided")
			return
		}

		if strings.EqualFold(demoConf.RpcStyle, "restful") {
			startSubscriptionRestfulProvider()
		} else {
			startSubscriptionJsonRpcProvider()
		}

		return
	}

	{ // billing
		if len(demoConf.Gateway) == 0 {
			logrus.Info("Billing gateway must be provided")
			return
		}

		if len(demoConf.BillingKey) == 0 {
			logrus.Info("Billing key must be provided")
			return
		}

		if strings.EqualFold(demoConf.RpcStyle, "restful") {
			startBillingRestfulProvider()
		} else {
			startBillingJsonRpcProvider()
		}
	}
}

func startBillingRestfulProvider() {
	logrus.WithField("listenPort", demoConf.ServerPort).
		Info("Starting demo billing RESTful service provider...")

	err := demo.RunBillingRestfulServiceProvider(demoConf.BillingClientConfig, demoConf.ServerPort)
	if err != nil {
		logrus.WithField("demoConfig", demoConf).
			Info("Failed to serve demo billing RESTful service provider")
	}
}

func startSubscriptionRestfulProvider() {
	logrus.WithField("listenPort", demoConf.ServerPort).
		Info("Starting demo VIP subscription RESTful service provider...")

	err := demo.RunSubscriptionRestfulServiceProvider(demoConf.VipSubscriptionClientConfig, demoConf.ServerPort)
	if err != nil {
		logrus.WithField("demoConfig", demoConf).
			Info("Failed to serve demo VIP subscription RESTful service provider")
	}
}

func startBillingJsonRpcProvider() {
	logrus.WithField("listenPort", demoConf.ServerPort).
		Info("Starting demo billing JSON-RPC service provider...")

	err := demo.RunBillingJsonRpcServiceProvider(demoConf.BillingClientConfig, demoConf.ServerPort)
	if err != nil {
		logrus.WithField("demoConfig", demoConf).
			Info("Failed to serve demo billing JSON-RPC service provider")
	}
}

func startSubscriptionJsonRpcProvider() {
	logrus.WithField("listenPort", demoConf.ServerPort).
		Info("Starting demo VIP subscription JSON-RPC service provider...")

	err := demo.RunSubscriptionJsonRpcServiceProvider(demoConf.VipSubscriptionClientConfig, demoConf.ServerPort)
	if err != nil {
		logrus.WithField("demoConfig", demoConf).
			Info("Failed to serve demo VIP subscription JSON-RPC service provider")
	}
}

func startConsumer(cmd *cobra.Command, args []string) {
	// load config from env vars or flags
	viperutil.MustUnmarshalKey("demo", &demoConf)

	if len(demoConf.ApiKey) == 0 {
		logrus.Info("Api key must be provided")
		return
	}

	if strings.EqualFold(demoConf.PayMode, "subscription") {
		if strings.EqualFold(demoConf.RpcStyle, "restful") {
			startSubscriptionRestfulConsumer()
		} else {
			startSubscriptionJsonRpcConsuemr()
		}

		return
	}

	{ // billing
		if strings.EqualFold(demoConf.RpcStyle, "restful") {
			startBillingRestfulConsumer()
		} else {
			startBillingJsonRpcConsumer()
		}
	}
}

func startBillingRestfulConsumer() {
	resp, err := demo.RunRestfulServiceConsumer(demoConf.ApiKey, demoConf.ServerPort)
	if err != nil {
		logrus.WithField("demoConfig", demoConf).
			WithError(err).
			Info("Failed to run demo billing RESTful service consumer")
		return
	}

	logrus.WithField("response", resp).Info("Run demo billing RESTful service consumer")
}

func startSubscriptionRestfulConsumer() {
	resp, err := demo.RunRestfulServiceConsumer(demoConf.ApiKey, demoConf.ServerPort)
	if err != nil {
		logrus.WithField("demoConfig", demoConf).
			WithError(err).
			Info("Failed to run demo VIP subscription RESTful service consumer")
		return
	}

	logrus.WithField("response", resp).Info("Run demo VIP subscription RESTful service consumer")
}

func startBillingJsonRpcConsumer() {
	resp, err := demo.RunBillingJsonRpcServiceConsumer(demoConf.ApiKey, demoConf.ServerPort)
	if err != nil {
		logrus.WithField("demoConfig", demoConf).
			WithError(err).
			Info("Failed to run demo billing JSON-RPC service consumer")
		return
	}

	logrus.WithField("response", resp).Info("Run demo billing JSON-RPC service consumer")
}

func startSubscriptionJsonRpcConsuemr() {
	resp, err := demo.RunSubscriptionJsonRpcServiceConsumer(demoConf.ApiKey, demoConf.ServerPort)
	if err != nil {
		logrus.WithField("demoConfig", demoConf).
			WithError(err).
			Info("Failed to run demo VIP subscription JSON-RPC service consumer")
		return
	}

	logrus.WithField("response", resp).Info("Run demo VIP subscription JSON-RPC service consumer")
}
