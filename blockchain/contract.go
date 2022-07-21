package blockchain

import (
	"github.com/Conflux-Chain/web3pay-service/contract"
	"github.com/ethereum/go-ethereum/common"
)

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
	stub *contract.APPCoin
}

func newAppCoinContractObj(
	contractAddr, owner *common.Address, stub *contract.APPCoin,
) *appCoinContractObj {
	return &appCoinContractObj{
		contractObj: &contractObj{
			addr: contractAddr, owner: owner,
		},
		stub: stub,
	}
}
