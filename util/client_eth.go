package util

import (
	"time"

	"github.com/Conflux-Chain/go-conflux-util/viper"
	providers "github.com/openweb3/go-rpc-provider/provider_wrapper"
	"github.com/openweb3/web3go"
	"github.com/sirupsen/logrus"
)

type ethClientConfig struct {
	Http            string
	Retry           int
	RetryInterval   time.Duration `default:"1s"`
	RequestTimeout  time.Duration `default:"3s"`
	MaxConnsPerHost int           `default:"1024"`
}

func MustNewEthClientFromViper(customOpt ...func(*web3go.ClientOption)) *web3go.Client {
	var conf ethClientConfig
	viper.MustUnmarshalKey("eth", &conf)

	option := web3go.ClientOption{
		Option: providers.Option{
			RetryCount:           conf.Retry,
			RetryInterval:        conf.RetryInterval,
			RequestTimeout:       conf.RequestTimeout,
			MaxConnectionPerHost: conf.MaxConnsPerHost,
		},
	}

	if len(customOpt) > 0 {
		customOpt[0](&option)
	}

	eth, err := web3go.NewClientWithOption(conf.Http, option)

	if err != nil {
		logrus.WithError(err).Fatal("Failed to create eth client from viper")
	}

	return eth
}
