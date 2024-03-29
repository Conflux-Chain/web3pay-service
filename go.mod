module github.com/Conflux-Chain/web3pay-service

go 1.15

require (
	github.com/Conflux-Chain/go-conflux-util v0.0.0-20220907035343-2d1233bccd70
	github.com/MoeYang/go-queue v0.0.0-20210407055646-c5a229ee466c
	github.com/ethereum/go-ethereum v1.10.15
	github.com/gammazero/workerpool v1.1.2
	github.com/go-resty/resty/v2 v2.7.0
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/rpc v1.2.0
	github.com/hashicorp/go-memdb v1.3.3
	github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d
	github.com/mcuadros/go-defaults v1.2.0
	github.com/openweb3/go-rpc-provider v0.3.1
	github.com/openweb3/web3go v0.2.3
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/shopspring/decimal v1.3.1
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.5.0
	github.com/spf13/viper v1.10.0
	github.com/stretchr/testify v1.7.0
	github.com/subosito/gotenv v1.2.0
	github.com/urfave/negroni v1.0.0
	go.uber.org/multierr v1.6.0
	golang.org/x/net v0.0.0-20220111093109-d55c255bac03
	gorm.io/driver/sqlite v1.3.6
	gorm.io/gorm v1.23.8
)

require (
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce
	github.com/google/uuid v1.3.0 // indirect
	github.com/schollz/jsonstore v1.1.0
	golang.org/x/sys v0.0.0-20220405052023-b1e9470b6e64 // indirect
	gopkg.in/yaml.v3 v3.0.0 // indirect
)

// for debugging development
// replace github.com/Conflux-Chain/go-conflux-util => ../go-conflux-util
