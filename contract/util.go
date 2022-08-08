package contract

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
)

var (
	errEventSigMismatch = errors.New("event signature mismatch")

	// event names
	EventControllerAppCreated   = "APP_CREATED"
	EventAirdropDrop            = "Drop"
	EventAppCoinTransferSingle  = "TransferSingle"
	EventAppCoinFrozen          = "Frozen"
	EventAppCointWithdraw       = "Withdraw"
	EventAppCoinResourceChanged = "ResourceChanged"
)

func UnpackControllerAPPCREATED(ctrlAbi *abi.ABI, log *types.Log) (*ControllerAPPCREATED, error) {
	eventObj := new(ControllerAPPCREATED)
	err := unpackLogEventData(eventObj, ctrlAbi, EventControllerAppCreated, log)

	return eventObj, err
}

func UnpackAPPCoinTransferSingle(appCoinAbi *abi.ABI, log *types.Log) (*APPCoinTransferSingle, error) {
	eventObj := new(APPCoinTransferSingle)
	err := unpackLogEventData(eventObj, appCoinAbi, EventAppCoinTransferSingle, log)

	return eventObj, err
}

func UnpackAirdropDrop(airdropAbi *abi.ABI, log *types.Log) (*AirdropDrop, error) {
	eventObj := new(AirdropDrop)
	err := unpackLogEventData(eventObj, airdropAbi, EventAirdropDrop, log)

	return eventObj, err
}

func UnpackAppCoinFrozen(appCoinAbi *abi.ABI, log *types.Log) (*APPCoinFrozen, error) {
	eventObj := new(APPCoinFrozen)
	err := unpackLogEventData(eventObj, appCoinAbi, EventAppCoinFrozen, log)

	return eventObj, err
}

func UnpackAppCoinWithdraw(appCoinAbi *abi.ABI, log *types.Log) (*APPCoinWithdraw, error) {
	eventObj := new(APPCoinWithdraw)
	err := unpackLogEventData(eventObj, appCoinAbi, EventAppCointWithdraw, log)

	return eventObj, err
}

func UnpackAppCoinResourceChanged(appCoinAbi *abi.ABI, log *types.Log) (*APPCoinResourceChanged, error) {
	eventObj := new(APPCoinResourceChanged)
	err := unpackLogEventData(eventObj, appCoinAbi, EventAppCoinResourceChanged, log)

	return eventObj, err
}

func unpackLogEventData(outPtr interface{}, contractAbi *abi.ABI, event string, log *types.Log) error {
	if log.Topics[0] != contractAbi.Events[event].ID {
		return errEventSigMismatch
	}

	if len(log.Data) > 0 {
		if err := contractAbi.UnpackIntoInterface(outPtr, event, log.Data); err != nil {
			return errors.WithMessage(err, "failed to unpack log data")
		}
	}

	var indexed abi.Arguments
	for _, arg := range contractAbi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}

	if err := abi.ParseTopics(outPtr, indexed, log.Topics[1:]); err != nil {
		return errors.WithMessage(err, "failed to parse log topics")
	}

	return nil
}
