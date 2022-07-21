package blockchain

import (
	"github.com/Conflux-Chain/web3pay-service/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go"
)

type contractBindCallContext struct {
	contractClient *web3go.ClientForContract
	signer         bind.SignerFn
}

type contractObj struct {
	addr  *common.Address
	owner *common.Address
}

type controllerContractObj struct {
	*contractObj
	stub *contract.Controller
}

func newControllerContractObj(
	contractAddr, owner *common.Address, stub *contract.Controller,
) *controllerContractObj {
	return &controllerContractObj{
		contractObj: &contractObj{
			addr: contractAddr, owner: owner,
		},
		stub: stub,
	}
}

type appCoinContractObj struct {
	*contractObj
	stub      *contract.APPCoin
	resources map[string]contract.AppConfigConfigEntry
}

func newAppCoinContractObj(
	contractAddr, owner *common.Address, stub *contract.APPCoin,
) *appCoinContractObj {
	return &appCoinContractObj{
		contractObj: &contractObj{
			addr: contractAddr, owner: owner,
		},
		stub:      stub,
		resources: make(map[string]contract.AppConfigConfigEntry),
	}
}
