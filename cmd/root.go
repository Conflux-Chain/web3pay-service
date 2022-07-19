package cmd

import (
	"os"

	"github.com/Conflux-Chain/web3pay-service/model"
	"github.com/Conflux-Chain/web3pay-service/store/sqlite"
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
	// prepare sqlite store
	config := sqlite.MustNewConfigFromViper()
	_ = config.MustOpenOrCreate(model.All...)
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
