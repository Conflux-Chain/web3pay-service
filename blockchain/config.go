package blockchain

import (
	"github.com/Conflux-Chain/go-conflux-util/viper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
)

type Config struct {
	ControllerContractAddr common.Address
	CreatorAddr            *common.Address
	OperatorPrivateKey     string
}

func MustNewConfigFromViper() *Config {
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

	conf := &Config{
		ControllerContractAddr: common.HexToAddress(rawConfig.ControllerContractAddr),
		OperatorPrivateKey:     rawConfig.OperatorPrivateKey,
	}

	if len(rawConfig.CreatorAddr) > 0 {
		if !common.IsHexAddress(rawConfig.CreatorAddr) {
			logrus.WithField(
				"creatorAddr", rawConfig.CreatorAddr,
			).Fatal("Invalid creator address configured")
		}

		creator := common.HexToAddress(rawConfig.CreatorAddr)
		conf.CreatorAddr = &creator
	}

	return conf
}
