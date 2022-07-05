package main

import (
	// ensure viper based configuration initialized at the very beginning
	_ "github.com/Conflux-Chain/web3pay-service/config"

	"github.com/Conflux-Chain/web3pay-service/cmd"
)

func main() {
	cmd.Execute()
}
