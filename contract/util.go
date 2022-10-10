package contract

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
)

const (
	// operation codes for resource pending configurations
	OpCodeResourceConfigAdd = iota
	OpCodeResourceConfigUpdate
	OpCodeResourceConfigDelete
	OpCodeResourceConfigNoPending
	OpCodeResourceConfigPendingInitDefault

	// token id
	TokenIdCoin    = 0
	TokenIdAirdrop = 1
)

var (
	errEventSigMismatch = errors.New("event signature mismatch")

	// events
	EventAppRegistryCreated = "Created"

	EventAppDeposit  = "Deposit"
	EventAppWithdraw = "Withdraw"
	EventAppFrozen   = "Frozen"

	EventApiTokenWeightResourceChanged = "ResourceChanged"

	EventVipCoinTransferSingle = "TransferSingle"

	// methods
	MethodAppChargeBatch                   = "chargeBatch"
	MethodApiWeightTokenFlushPendingConfig = "flushPendingConfig"
)

func UnpackAppRegistryCreated(abi *abi.ABI, log *types.Log) (*AppRegistryCreated, error) {
	eventObj := new(AppRegistryCreated)
	err := unpackLogEventData(eventObj, abi, EventAppRegistryCreated, log)

	return eventObj, err
}

func UnpackAppDeposit(appAbi *abi.ABI, log *types.Log) (*AppDeposit, error) {
	eventObj := new(AppDeposit)
	err := unpackLogEventData(eventObj, appAbi, EventAppDeposit, log)

	return eventObj, err
}

func UnpackAppFrozen(appAbi *abi.ABI, log *types.Log) (*AppFrozen, error) {
	eventObj := new(AppFrozen)
	err := unpackLogEventData(eventObj, appAbi, EventAppWithdraw, log)

	return eventObj, err
}

func UnpackAppWithdraw(appAbi *abi.ABI, log *types.Log) (*AppWithdraw, error) {
	eventObj := new(AppWithdraw)
	err := unpackLogEventData(eventObj, appAbi, EventAppWithdraw, log)

	return eventObj, err
}

func UnpackApiWeightTokenResourceChanged(awtAbi *abi.ABI, log *types.Log) (*ApiWeightTokenResourceChanged, error) {
	eventObj := new(ApiWeightTokenResourceChanged)
	err := unpackLogEventData(eventObj, awtAbi, EventApiTokenWeightResourceChanged, log)

	return eventObj, err
}

func UnpackVipCoinTransferSingle(vcAbi *abi.ABI, log *types.Log) (*VipCoinTransferSingle, error) {
	eventObj := new(VipCoinTransferSingle)
	err := unpackLogEventData(eventObj, vcAbi, EventVipCoinTransferSingle, log)

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
