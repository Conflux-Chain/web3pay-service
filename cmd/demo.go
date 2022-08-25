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
	client.ClientConfig `mapstructure:",squash"`
	CustomerKey         string
	ServerPort          int
	RpcStyle            string
}

var (
	demoConf = demoConfig{}

	demoCmd = &cobra.Command{
		Use:   "demo",
		Short: "Run billed service provider and consumer demo",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	demoProviderCmd = &cobra.Command{
		Use:   "provider",
		Short: "Run demo billed service provider",
		Run:   startProvider,
	}

	demoConsumerCmd = &cobra.Command{
		Use:   "consumer",
		Short: "Run demo billed service consumer",
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
		&demoConf.Gateway, "gateway", "g", "", "billing gateway URL",
	)
	viper.BindPFlag("demo.gateway", flags.Lookup("gateway"))

	// billing key
	flags.StringVarP(
		&demoConf.BillingKey, "billing-key", "b", "",
		"billing key (use 'ecrecover' command to generate one)",
	)
	viper.BindPFlag("demo.billingKey", flags.Lookup("billing-key"))

	// customer key
	flags.StringVarP(
		&demoConf.CustomerKey, "customer-key", "c", "",
		"customer key (use 'ecrecover' command to generate one)",
	)
	viper.BindPFlag("demo.customerKey", flags.Lookup("customer-key"))

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
}

func startProvider(cmd *cobra.Command, args []string) {
	// load config from env vars or flags
	viperutil.MustUnmarshalKey("demo", &demoConf)

	if len(demoConf.Gateway) == 0 {
		logrus.Info("Billing gateway must be provided")
		return
	}

	if len(demoConf.BillingKey) == 0 {
		logrus.Info("Billing key must be provided")
		return
	}

	if strings.EqualFold(demoConf.RpcStyle, "restful") {
		startRestfulProvider()
	} else {
		startJsonRpcProvider()
	}
}

func startRestfulProvider() {
	logrus.WithField("listenPort", demoConf.ServerPort).
		Info("Starting demo RESTful billed service provider...")

	err := demo.RunRestfulServiceProvider(demoConf.ClientConfig, demoConf.ServerPort)
	if err != nil {
		logrus.WithField("demoConfig", demoConf).
			Info("Failed to serve RESTful billed service provider")
	}
}

func startJsonRpcProvider() {
	logrus.WithField("listenPort", demoConf.ServerPort).
		Info("Starting demo JSON-RPC billed service provider...")

	err := demo.RunJsonRpcServiceProvider(demoConf.ClientConfig, demoConf.ServerPort)
	if err != nil {
		logrus.WithField("demoConfig", demoConf).
			Info("Failed to serve JSON-RPC billed service provider")
	}
}

func startConsumer(cmd *cobra.Command, args []string) {
	// load config from env vars or flags
	viperutil.MustUnmarshalKey("demo", &demoConf)

	if len(demoConf.CustomerKey) == 0 {
		logrus.Info("Customer key must be provided")
		return
	}

	if strings.EqualFold(demoConf.RpcStyle, "restful") {
		startRestfulConsumer()
	} else {
		startJsonRpcConsuemr()
	}
}

func startRestfulConsumer() {
	resp, err := demo.RunRestfulServiceConsumer(demoConf.CustomerKey, demoConf.ServerPort)
	if err != nil {
		logrus.WithField("demoConfig", demoConf).
			WithError(err).
			Info("Failed to run demo billed RESTful service consumer")
		return
	}

	logrus.WithField("response", resp).Info("Run demo billed RESTful service consumer")
}

func startJsonRpcConsuemr() {
	resp, err := demo.RunJsonRpcServiceConsumer(demoConf.CustomerKey, demoConf.ServerPort)
	if err != nil {
		logrus.WithField("demoConfig", demoConf).
			WithError(err).
			Info("Failed to run demo billed JSON-RPC service consumer")
		return
	}

	logrus.WithField("response", resp).Info("Run demo billed JSON-RPC service consumer")
}
