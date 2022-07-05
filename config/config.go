package config

import (
	utilcfg "github.com/Conflux-Chain/go-conflux-util/config"
)

// this can be overwriten when go build
var EnvPrefix string = "web3pay"

func init() {
	utilcfg.MustInit(EnvPrefix)
}
