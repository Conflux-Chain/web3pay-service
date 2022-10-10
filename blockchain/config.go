package blockchain

import (
	"github.com/Conflux-Chain/go-conflux-util/viper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
)

type Config struct {
	AppRegistryContractAddr common.Address
	OwnerAddr               *common.Address
	OperatorPrivateKey      string
}

func MustNewConfigFromViper() *Config {
	var rawConfig struct {
		AppRegistryContractAddr string
		OwnerAddr               string
		OperatorPrivateKey      string
	}

	viper.MustUnmarshalKey("blockchain", &rawConfig)

	if !common.IsHexAddress(rawConfig.AppRegistryContractAddr) {
		logrus.WithField(
			"appRegistryContractAddr", rawConfig.AppRegistryContractAddr,
		).Fatal("Invalid `AppRegistry` contract address configured")
	}

	if len(rawConfig.OperatorPrivateKey) == 0 {
		logrus.Fatal("Operator private key not configured")
	}

	conf := &Config{
		AppRegistryContractAddr: common.HexToAddress(rawConfig.AppRegistryContractAddr),
		OperatorPrivateKey:      rawConfig.OperatorPrivateKey,
	}

	if len(rawConfig.OwnerAddr) > 0 {
		if !common.IsHexAddress(rawConfig.OwnerAddr) {
			logrus.WithField(
				"ownerAddr", rawConfig.OwnerAddr,
			).Fatal("Invalid owner address configured")
		}

		owner := common.HexToAddress(rawConfig.OwnerAddr)
		conf.OwnerAddr = &owner
	}

	return conf
}
