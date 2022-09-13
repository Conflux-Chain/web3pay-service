package config

import (
	utilcfg "github.com/Conflux-Chain/go-conflux-util/config"
	"github.com/Conflux-Chain/go-conflux-util/log"
	"github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"
)

// this can be overwriten when go build
var EnvPrefix string = "web3pay"

func init() {
	utilcfg.MustInit(EnvPrefix)

	// Add alert hook for logrus fatal/warn/error level
	hookLevels := []logrus.Level{logrus.FatalLevel, logrus.WarnLevel, logrus.ErrorLevel}
	dingTalkAlertHook := log.NewDingTalkAlertHook(hookLevels)
	logrus.AddHook(dingTalkAlertHook)

	// load environment variables from .env file from current directory
	gotenv.Load()
}
