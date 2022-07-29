package blockchain

import (
	"github.com/Conflux-Chain/go-conflux-util/viper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
)

var stdConf *config

type config struct {
	controllerContractAddr common.Address
	creatorAddr            *common.Address
	operatorPrivateKey     string
}

func init() {
	var rawConfig struct {
		ControllerContractAddr string
		CreatorAddr            string
		OperatorPrivateKey     string
	}

	viper.MustUnmarshalKey("blockchain", &rawConfig)

	if !common.IsHexAddress(rawConfig.ControllerContractAddr) {
		logrus.WithField(
			"controllerContractAddr", rawConfig.ControllerContractAddr,
		).Fatal("Invalid controller contract address configured")
	}

	if len(rawConfig.OperatorPrivateKey) == 0 {
		logrus.Fatal("Operator private key not configured")
	}

	stdConf = &config{
		controllerContractAddr: common.HexToAddress(rawConfig.ControllerContractAddr),
		operatorPrivateKey:     rawConfig.OperatorPrivateKey,
	}

	if len(rawConfig.CreatorAddr) > 0 {
		if !common.IsHexAddress(rawConfig.CreatorAddr) {
			logrus.WithField(
				"creatorAddr", rawConfig.CreatorAddr,
			).Fatal("Invalid creator address configured")
		}

		creator := common.HexToAddress(rawConfig.CreatorAddr)
		stdConf.creatorAddr = &creator
	}
}
