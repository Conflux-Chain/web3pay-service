package config

import (
	utilcfg "github.com/Conflux-Chain/go-conflux-util/config"
	"github.com/subosito/gotenv"
)

// this can be overwriten when go build
var EnvPrefix string = "web3pay"

func init() {
	// Initialize logging, metric, alert, and configuration systems.
	utilcfg.MustInit(EnvPrefix)

	// load environment variables from .env file from current directory
	gotenv.Load()
}
