package cmd

import (
	"context"
	"os"
	"sync"

	"github.com/Conflux-Chain/web3pay-service/api"
	"github.com/Conflux-Chain/web3pay-service/blockchain"
	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/service"
	"github.com/Conflux-Chain/web3pay-service/store/memdb"
	"github.com/Conflux-Chain/web3pay-service/store/sqlite"
	"github.com/Conflux-Chain/web3pay-service/util"
	"github.com/Conflux-Chain/web3pay-service/worker"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "web3pay-service",
	Short: "Decentralized Web3 Payment Service for Conflux Network",
	Long:  `Backend service for web3 payment system served for fee billing and blockchain submitting etc.`,
	Run:   start,
}

func start(cmd *cobra.Command, args []string) {
	// sqlite store
	config := sqlite.MustNewConfigFromViper()
	sqliteStore := config.MustOpenOrCreate(model.All...)
	defer sqliteStore.Close()

	// memory store
	memStore := memdb.MustNewStoreFromViper()
	defer memStore.Close()

	// blockchain config
	chainConfig := blockchain.MustNewConfigFromViper()

	// blockchain ops provider
	chainOpsProvider := blockchain.MustNewProvider(chainConfig)

	// service factory
	serviceFactory := service.MustNewFactory(sqliteStore, memStore, chainOpsProvider)

	// monitor
	chainMonitor := blockchain.MustNewMonitor(chainConfig, chainOpsProvider, serviceFactory.Blockchain)

	// worker
	chainWorker := worker.MustNewBlockchainWorkerFromViper(
		chainOpsProvider, sqliteStore, serviceFactory.Billing, serviceFactory.Blockchain,
	)

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	// start monitor server
	wg.Add(1)
	go func() {
		defer wg.Done()
		chainMonitor.Sync(ctx)
	}()

	// start blockchain worker
	wg.Add(1)
	go func() {
		defer wg.Done()
		chainWorker.Run(ctx)
	}()

	// start RPC server
	wg.Add(1)
	go func() {
		defer wg.Done()

		go api.MustServe(serviceFactory)
		<-ctx.Done()

		err := api.Shutdown()
		logrus.WithError(err).Info("RPC server shut down")
	}()

	util.GracefulShutdown(&wg, cancel)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.web3pay-service.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
