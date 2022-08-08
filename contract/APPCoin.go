// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// APPCoinChargeRequest is an auto generated low-level Go binding around an user-defined struct.
type APPCoinChargeRequest struct {
	Account   common.Address
	Amount    *big.Int
	Data      []byte
	UseDetail []APPCoinResourceUseDetail
}

// APPCoinResourceUseDetail is an auto generated low-level Go binding around an user-defined struct.
type APPCoinResourceUseDetail struct {
	Id    uint32
	Times *big.Int
}

// APPCoinUserCharged is an auto generated low-level Go binding around an user-defined struct.
type APPCoinUserCharged struct {
	User   common.Address
	Amount *big.Int
}

// AppConfigConfigEntry is an auto generated low-level Go binding around an user-defined struct.
type AppConfigConfigEntry struct {
	ResourceId    string
	Weight        *big.Int
	Index         uint32
	PendingOP     uint8
	PendingWeight *big.Int
	SubmitSeconds *big.Int
	RequestTimes  *big.Int
}

// AppConfigConfigRequest is an auto generated low-level Go binding around an user-defined struct.
type AppConfigConfigRequest struct {
	Id         uint32
	ResourceId string
	Weight     *big.Int
	Op         uint8
}

// APPCoinMetaData contains all meta data concerning the APPCoin contract.
var APPCoinMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"AppOwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"Frozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"enumAppConfig.OP\",\"name\":\"op\",\"type\":\"uint8\"}],\"name\":\"ResourceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newWeight\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"enumAppConfig.OP\",\"name\":\"op\",\"type\":\"uint8\"}],\"name\":\"ResourcePending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FIRST_CONFIG_ID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FT_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"apiCoin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"appOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOfWithAirdrop\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"airdrop\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"times\",\"type\":\"uint256\"}],\"internalType\":\"structAPPCoin.ResourceUseDetail[]\",\"name\":\"useDetail\",\"type\":\"tuple[]\"}],\"name\":\"charge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"times\",\"type\":\"uint256\"}],\"internalType\":\"structAPPCoin.ResourceUseDetail[]\",\"name\":\"useDetail\",\"type\":\"tuple[]\"}],\"internalType\":\"structAPPCoin.ChargeRequest[]\",\"name\":\"requestArray\",\"type\":\"tuple[]\"}],\"name\":\"chargeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"chargedMapping\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"resourceId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"enumAppConfig.OP\",\"name\":\"op\",\"type\":\"uint8\"}],\"internalType\":\"structAppConfig.ConfigRequest\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"configResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"resourceId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"enumAppConfig.OP\",\"name\":\"op\",\"type\":\"uint8\"}],\"internalType\":\"structAppConfig.ConfigRequest[]\",\"name\":\"entries\",\"type\":\"tuple[]\"}],\"name\":\"configResourceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flushPendingConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceWithdrawDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"f\",\"type\":\"bool\"}],\"name\":\"freeze\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"frozenMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"indexArray\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"apiCoin_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"appOwner_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"defaultWeight\",\"type\":\"uint256\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"initOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"listResources\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"resourceId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"enumAppConfig.OP\",\"name\":\"pendingOP\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"pendingWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"submitSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestTimes\",\"type\":\"uint256\"}],\"internalType\":\"structAppConfig.ConfigEntry[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"listUser\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structAPPCoin.UserCharged[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32[]\",\"name\":\"ids\",\"type\":\"uint32[]\"}],\"name\":\"listUserRequestCounter\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"times\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextConfigId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingIdArray\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"resourceConfigures\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"resourceId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"enumAppConfig.OP\",\"name\":\"pendingOP\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"pendingWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"submitSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestTimes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setForceWithdrawDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"seconds_\",\"type\":\"uint256\"}],\"name\":\"setPendingSeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"takeProfit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"tokensReceived\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalCharged\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalTakenProfit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"transferAppOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"users\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// APPCoinABI is the input ABI used to generate the binding from.
// Deprecated: Use APPCoinMetaData.ABI instead.
var APPCoinABI = APPCoinMetaData.ABI

// APPCoin is an auto generated Go binding around an Ethereum contract.
type APPCoin struct {
	APPCoinCaller     // Read-only binding to the contract
	APPCoinTransactor // Write-only binding to the contract
	APPCoinFilterer   // Log filterer for contract events
}

// APPCoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type APPCoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// APPCoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type APPCoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// APPCoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type APPCoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// APPCoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type APPCoinSession struct {
	Contract     *APPCoin          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// APPCoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type APPCoinCallerSession struct {
	Contract *APPCoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// APPCoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type APPCoinTransactorSession struct {
	Contract     *APPCoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// APPCoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type APPCoinRaw struct {
	Contract *APPCoin // Generic contract binding to access the raw methods on
}

// APPCoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type APPCoinCallerRaw struct {
	Contract *APPCoinCaller // Generic read-only contract binding to access the raw methods on
}

// APPCoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type APPCoinTransactorRaw struct {
	Contract *APPCoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAPPCoin creates a new instance of APPCoin, bound to a specific deployed contract.
func NewAPPCoin(address common.Address, backend bind.ContractBackend) (*APPCoin, error) {
	contract, err := bindAPPCoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &APPCoin{APPCoinCaller: APPCoinCaller{contract: contract}, APPCoinTransactor: APPCoinTransactor{contract: contract}, APPCoinFilterer: APPCoinFilterer{contract: contract}}, nil
}

// NewAPPCoinCaller creates a new read-only instance of APPCoin, bound to a specific deployed contract.
func NewAPPCoinCaller(address common.Address, caller bind.ContractCaller) (*APPCoinCaller, error) {
	contract, err := bindAPPCoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &APPCoinCaller{contract: contract}, nil
}

// NewAPPCoinTransactor creates a new write-only instance of APPCoin, bound to a specific deployed contract.
func NewAPPCoinTransactor(address common.Address, transactor bind.ContractTransactor) (*APPCoinTransactor, error) {
	contract, err := bindAPPCoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &APPCoinTransactor{contract: contract}, nil
}

// NewAPPCoinFilterer creates a new log filterer instance of APPCoin, bound to a specific deployed contract.
func NewAPPCoinFilterer(address common.Address, filterer bind.ContractFilterer) (*APPCoinFilterer, error) {
	contract, err := bindAPPCoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &APPCoinFilterer{contract: contract}, nil
}

// bindAPPCoin binds a generic wrapper to an already deployed contract.
func bindAPPCoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(APPCoinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_APPCoin *APPCoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _APPCoin.Contract.APPCoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_APPCoin *APPCoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _APPCoin.Contract.APPCoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_APPCoin *APPCoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _APPCoin.Contract.APPCoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_APPCoin *APPCoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _APPCoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_APPCoin *APPCoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _APPCoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_APPCoin *APPCoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _APPCoin.Contract.contract.Transact(opts, method, params...)
}

// FIRSTCONFIGID is a free data retrieval call binding the contract method 0x0fe58201.
//
// Solidity: function FIRST_CONFIG_ID() view returns(uint32)
func (_APPCoin *APPCoinCaller) FIRSTCONFIGID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "FIRST_CONFIG_ID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// FIRSTCONFIGID is a free data retrieval call binding the contract method 0x0fe58201.
//
// Solidity: function FIRST_CONFIG_ID() view returns(uint32)
func (_APPCoin *APPCoinSession) FIRSTCONFIGID() (uint32, error) {
	return _APPCoin.Contract.FIRSTCONFIGID(&_APPCoin.CallOpts)
}

// FIRSTCONFIGID is a free data retrieval call binding the contract method 0x0fe58201.
//
// Solidity: function FIRST_CONFIG_ID() view returns(uint32)
func (_APPCoin *APPCoinCallerSession) FIRSTCONFIGID() (uint32, error) {
	return _APPCoin.Contract.FIRSTCONFIGID(&_APPCoin.CallOpts)
}

// FTID is a free data retrieval call binding the contract method 0x01052d57.
//
// Solidity: function FT_ID() view returns(uint256)
func (_APPCoin *APPCoinCaller) FTID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "FT_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FTID is a free data retrieval call binding the contract method 0x01052d57.
//
// Solidity: function FT_ID() view returns(uint256)
func (_APPCoin *APPCoinSession) FTID() (*big.Int, error) {
	return _APPCoin.Contract.FTID(&_APPCoin.CallOpts)
}

// FTID is a free data retrieval call binding the contract method 0x01052d57.
//
// Solidity: function FT_ID() view returns(uint256)
func (_APPCoin *APPCoinCallerSession) FTID() (*big.Int, error) {
	return _APPCoin.Contract.FTID(&_APPCoin.CallOpts)
}

// ApiCoin is a free data retrieval call binding the contract method 0x4ebf365a.
//
// Solidity: function apiCoin() view returns(address)
func (_APPCoin *APPCoinCaller) ApiCoin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "apiCoin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ApiCoin is a free data retrieval call binding the contract method 0x4ebf365a.
//
// Solidity: function apiCoin() view returns(address)
func (_APPCoin *APPCoinSession) ApiCoin() (common.Address, error) {
	return _APPCoin.Contract.ApiCoin(&_APPCoin.CallOpts)
}

// ApiCoin is a free data retrieval call binding the contract method 0x4ebf365a.
//
// Solidity: function apiCoin() view returns(address)
func (_APPCoin *APPCoinCallerSession) ApiCoin() (common.Address, error) {
	return _APPCoin.Contract.ApiCoin(&_APPCoin.CallOpts)
}

// AppOwner is a free data retrieval call binding the contract method 0x32b93a86.
//
// Solidity: function appOwner() view returns(address)
func (_APPCoin *APPCoinCaller) AppOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "appOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AppOwner is a free data retrieval call binding the contract method 0x32b93a86.
//
// Solidity: function appOwner() view returns(address)
func (_APPCoin *APPCoinSession) AppOwner() (common.Address, error) {
	return _APPCoin.Contract.AppOwner(&_APPCoin.CallOpts)
}

// AppOwner is a free data retrieval call binding the contract method 0x32b93a86.
//
// Solidity: function appOwner() view returns(address)
func (_APPCoin *APPCoinCallerSession) AppOwner() (common.Address, error) {
	return _APPCoin.Contract.AppOwner(&_APPCoin.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_APPCoin *APPCoinCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_APPCoin *APPCoinSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _APPCoin.Contract.BalanceOf(&_APPCoin.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_APPCoin *APPCoinCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _APPCoin.Contract.BalanceOf(&_APPCoin.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_APPCoin *APPCoinCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_APPCoin *APPCoinSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _APPCoin.Contract.BalanceOfBatch(&_APPCoin.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_APPCoin *APPCoinCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _APPCoin.Contract.BalanceOfBatch(&_APPCoin.CallOpts, accounts, ids)
}

// BalanceOfWithAirdrop is a free data retrieval call binding the contract method 0x1215193f.
//
// Solidity: function balanceOfWithAirdrop(address owner) view returns(uint256 total, uint256 airdrop)
func (_APPCoin *APPCoinCaller) BalanceOfWithAirdrop(opts *bind.CallOpts, owner common.Address) (struct {
	Total   *big.Int
	Airdrop *big.Int
}, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "balanceOfWithAirdrop", owner)

	outstruct := new(struct {
		Total   *big.Int
		Airdrop *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Total = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Airdrop = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// BalanceOfWithAirdrop is a free data retrieval call binding the contract method 0x1215193f.
//
// Solidity: function balanceOfWithAirdrop(address owner) view returns(uint256 total, uint256 airdrop)
func (_APPCoin *APPCoinSession) BalanceOfWithAirdrop(owner common.Address) (struct {
	Total   *big.Int
	Airdrop *big.Int
}, error) {
	return _APPCoin.Contract.BalanceOfWithAirdrop(&_APPCoin.CallOpts, owner)
}

// BalanceOfWithAirdrop is a free data retrieval call binding the contract method 0x1215193f.
//
// Solidity: function balanceOfWithAirdrop(address owner) view returns(uint256 total, uint256 airdrop)
func (_APPCoin *APPCoinCallerSession) BalanceOfWithAirdrop(owner common.Address) (struct {
	Total   *big.Int
	Airdrop *big.Int
}, error) {
	return _APPCoin.Contract.BalanceOfWithAirdrop(&_APPCoin.CallOpts, owner)
}

// ChargedMapping is a free data retrieval call binding the contract method 0x8ade7975.
//
// Solidity: function chargedMapping(address ) view returns(uint256)
func (_APPCoin *APPCoinCaller) ChargedMapping(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "chargedMapping", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChargedMapping is a free data retrieval call binding the contract method 0x8ade7975.
//
// Solidity: function chargedMapping(address ) view returns(uint256)
func (_APPCoin *APPCoinSession) ChargedMapping(arg0 common.Address) (*big.Int, error) {
	return _APPCoin.Contract.ChargedMapping(&_APPCoin.CallOpts, arg0)
}

// ChargedMapping is a free data retrieval call binding the contract method 0x8ade7975.
//
// Solidity: function chargedMapping(address ) view returns(uint256)
func (_APPCoin *APPCoinCallerSession) ChargedMapping(arg0 common.Address) (*big.Int, error) {
	return _APPCoin.Contract.ChargedMapping(&_APPCoin.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_APPCoin *APPCoinCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_APPCoin *APPCoinSession) Decimals() (uint8, error) {
	return _APPCoin.Contract.Decimals(&_APPCoin.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_APPCoin *APPCoinCallerSession) Decimals() (uint8, error) {
	return _APPCoin.Contract.Decimals(&_APPCoin.CallOpts)
}

// ForceWithdrawDelay is a free data retrieval call binding the contract method 0x6d27d1af.
//
// Solidity: function forceWithdrawDelay() view returns(uint256)
func (_APPCoin *APPCoinCaller) ForceWithdrawDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "forceWithdrawDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ForceWithdrawDelay is a free data retrieval call binding the contract method 0x6d27d1af.
//
// Solidity: function forceWithdrawDelay() view returns(uint256)
func (_APPCoin *APPCoinSession) ForceWithdrawDelay() (*big.Int, error) {
	return _APPCoin.Contract.ForceWithdrawDelay(&_APPCoin.CallOpts)
}

// ForceWithdrawDelay is a free data retrieval call binding the contract method 0x6d27d1af.
//
// Solidity: function forceWithdrawDelay() view returns(uint256)
func (_APPCoin *APPCoinCallerSession) ForceWithdrawDelay() (*big.Int, error) {
	return _APPCoin.Contract.ForceWithdrawDelay(&_APPCoin.CallOpts)
}

// FrozenMap is a free data retrieval call binding the contract method 0xe316468b.
//
// Solidity: function frozenMap(address ) view returns(uint256)
func (_APPCoin *APPCoinCaller) FrozenMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "frozenMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FrozenMap is a free data retrieval call binding the contract method 0xe316468b.
//
// Solidity: function frozenMap(address ) view returns(uint256)
func (_APPCoin *APPCoinSession) FrozenMap(arg0 common.Address) (*big.Int, error) {
	return _APPCoin.Contract.FrozenMap(&_APPCoin.CallOpts, arg0)
}

// FrozenMap is a free data retrieval call binding the contract method 0xe316468b.
//
// Solidity: function frozenMap(address ) view returns(uint256)
func (_APPCoin *APPCoinCallerSession) FrozenMap(arg0 common.Address) (*big.Int, error) {
	return _APPCoin.Contract.FrozenMap(&_APPCoin.CallOpts, arg0)
}

// IndexArray is a free data retrieval call binding the contract method 0xebed9b2c.
//
// Solidity: function indexArray(uint256 ) view returns(uint32)
func (_APPCoin *APPCoinCaller) IndexArray(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "indexArray", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// IndexArray is a free data retrieval call binding the contract method 0xebed9b2c.
//
// Solidity: function indexArray(uint256 ) view returns(uint32)
func (_APPCoin *APPCoinSession) IndexArray(arg0 *big.Int) (uint32, error) {
	return _APPCoin.Contract.IndexArray(&_APPCoin.CallOpts, arg0)
}

// IndexArray is a free data retrieval call binding the contract method 0xebed9b2c.
//
// Solidity: function indexArray(uint256 ) view returns(uint32)
func (_APPCoin *APPCoinCallerSession) IndexArray(arg0 *big.Int) (uint32, error) {
	return _APPCoin.Contract.IndexArray(&_APPCoin.CallOpts, arg0)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_APPCoin *APPCoinCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_APPCoin *APPCoinSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _APPCoin.Contract.IsApprovedForAll(&_APPCoin.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_APPCoin *APPCoinCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _APPCoin.Contract.IsApprovedForAll(&_APPCoin.CallOpts, account, operator)
}

// ListResources is a free data retrieval call binding the contract method 0x49459cd2.
//
// Solidity: function listResources(uint256 offset, uint256 limit) view returns((string,uint256,uint32,uint8,uint256,uint256,uint256)[], uint256 total)
func (_APPCoin *APPCoinCaller) ListResources(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]AppConfigConfigEntry, *big.Int, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "listResources", offset, limit)

	if err != nil {
		return *new([]AppConfigConfigEntry), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]AppConfigConfigEntry)).(*[]AppConfigConfigEntry)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// ListResources is a free data retrieval call binding the contract method 0x49459cd2.
//
// Solidity: function listResources(uint256 offset, uint256 limit) view returns((string,uint256,uint32,uint8,uint256,uint256,uint256)[], uint256 total)
func (_APPCoin *APPCoinSession) ListResources(offset *big.Int, limit *big.Int) ([]AppConfigConfigEntry, *big.Int, error) {
	return _APPCoin.Contract.ListResources(&_APPCoin.CallOpts, offset, limit)
}

// ListResources is a free data retrieval call binding the contract method 0x49459cd2.
//
// Solidity: function listResources(uint256 offset, uint256 limit) view returns((string,uint256,uint32,uint8,uint256,uint256,uint256)[], uint256 total)
func (_APPCoin *APPCoinCallerSession) ListResources(offset *big.Int, limit *big.Int) ([]AppConfigConfigEntry, *big.Int, error) {
	return _APPCoin.Contract.ListResources(&_APPCoin.CallOpts, offset, limit)
}

// ListUser is a free data retrieval call binding the contract method 0x3ec36183.
//
// Solidity: function listUser(uint256 offset, uint256 limit) view returns((address,uint256)[], uint256 total)
func (_APPCoin *APPCoinCaller) ListUser(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]APPCoinUserCharged, *big.Int, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "listUser", offset, limit)

	if err != nil {
		return *new([]APPCoinUserCharged), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]APPCoinUserCharged)).(*[]APPCoinUserCharged)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// ListUser is a free data retrieval call binding the contract method 0x3ec36183.
//
// Solidity: function listUser(uint256 offset, uint256 limit) view returns((address,uint256)[], uint256 total)
func (_APPCoin *APPCoinSession) ListUser(offset *big.Int, limit *big.Int) ([]APPCoinUserCharged, *big.Int, error) {
	return _APPCoin.Contract.ListUser(&_APPCoin.CallOpts, offset, limit)
}

// ListUser is a free data retrieval call binding the contract method 0x3ec36183.
//
// Solidity: function listUser(uint256 offset, uint256 limit) view returns((address,uint256)[], uint256 total)
func (_APPCoin *APPCoinCallerSession) ListUser(offset *big.Int, limit *big.Int) ([]APPCoinUserCharged, *big.Int, error) {
	return _APPCoin.Contract.ListUser(&_APPCoin.CallOpts, offset, limit)
}

// ListUserRequestCounter is a free data retrieval call binding the contract method 0xa5e34731.
//
// Solidity: function listUserRequestCounter(address user, uint32[] ids) view returns(uint256[] times)
func (_APPCoin *APPCoinCaller) ListUserRequestCounter(opts *bind.CallOpts, user common.Address, ids []uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "listUserRequestCounter", user, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ListUserRequestCounter is a free data retrieval call binding the contract method 0xa5e34731.
//
// Solidity: function listUserRequestCounter(address user, uint32[] ids) view returns(uint256[] times)
func (_APPCoin *APPCoinSession) ListUserRequestCounter(user common.Address, ids []uint32) ([]*big.Int, error) {
	return _APPCoin.Contract.ListUserRequestCounter(&_APPCoin.CallOpts, user, ids)
}

// ListUserRequestCounter is a free data retrieval call binding the contract method 0xa5e34731.
//
// Solidity: function listUserRequestCounter(address user, uint32[] ids) view returns(uint256[] times)
func (_APPCoin *APPCoinCallerSession) ListUserRequestCounter(user common.Address, ids []uint32) ([]*big.Int, error) {
	return _APPCoin.Contract.ListUserRequestCounter(&_APPCoin.CallOpts, user, ids)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_APPCoin *APPCoinCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_APPCoin *APPCoinSession) Name() (string, error) {
	return _APPCoin.Contract.Name(&_APPCoin.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_APPCoin *APPCoinCallerSession) Name() (string, error) {
	return _APPCoin.Contract.Name(&_APPCoin.CallOpts)
}

// NextConfigId is a free data retrieval call binding the contract method 0x99d726c7.
//
// Solidity: function nextConfigId() view returns(uint32)
func (_APPCoin *APPCoinCaller) NextConfigId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "nextConfigId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NextConfigId is a free data retrieval call binding the contract method 0x99d726c7.
//
// Solidity: function nextConfigId() view returns(uint32)
func (_APPCoin *APPCoinSession) NextConfigId() (uint32, error) {
	return _APPCoin.Contract.NextConfigId(&_APPCoin.CallOpts)
}

// NextConfigId is a free data retrieval call binding the contract method 0x99d726c7.
//
// Solidity: function nextConfigId() view returns(uint32)
func (_APPCoin *APPCoinCallerSession) NextConfigId() (uint32, error) {
	return _APPCoin.Contract.NextConfigId(&_APPCoin.CallOpts)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_APPCoin *APPCoinCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_APPCoin *APPCoinSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _APPCoin.Contract.OnERC1155BatchReceived(&_APPCoin.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_APPCoin *APPCoinCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _APPCoin.Contract.OnERC1155BatchReceived(&_APPCoin.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_APPCoin *APPCoinCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_APPCoin *APPCoinSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _APPCoin.Contract.OnERC1155Received(&_APPCoin.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_APPCoin *APPCoinCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _APPCoin.Contract.OnERC1155Received(&_APPCoin.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_APPCoin *APPCoinCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_APPCoin *APPCoinSession) Owner() (common.Address, error) {
	return _APPCoin.Contract.Owner(&_APPCoin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_APPCoin *APPCoinCallerSession) Owner() (common.Address, error) {
	return _APPCoin.Contract.Owner(&_APPCoin.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_APPCoin *APPCoinCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_APPCoin *APPCoinSession) Paused() (bool, error) {
	return _APPCoin.Contract.Paused(&_APPCoin.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_APPCoin *APPCoinCallerSession) Paused() (bool, error) {
	return _APPCoin.Contract.Paused(&_APPCoin.CallOpts)
}

// PendingIdArray is a free data retrieval call binding the contract method 0xac8ef2b5.
//
// Solidity: function pendingIdArray(uint256 ) view returns(uint32)
func (_APPCoin *APPCoinCaller) PendingIdArray(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "pendingIdArray", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// PendingIdArray is a free data retrieval call binding the contract method 0xac8ef2b5.
//
// Solidity: function pendingIdArray(uint256 ) view returns(uint32)
func (_APPCoin *APPCoinSession) PendingIdArray(arg0 *big.Int) (uint32, error) {
	return _APPCoin.Contract.PendingIdArray(&_APPCoin.CallOpts, arg0)
}

// PendingIdArray is a free data retrieval call binding the contract method 0xac8ef2b5.
//
// Solidity: function pendingIdArray(uint256 ) view returns(uint32)
func (_APPCoin *APPCoinCallerSession) PendingIdArray(arg0 *big.Int) (uint32, error) {
	return _APPCoin.Contract.PendingIdArray(&_APPCoin.CallOpts, arg0)
}

// PendingSeconds is a free data retrieval call binding the contract method 0xc819c70e.
//
// Solidity: function pendingSeconds() view returns(uint256)
func (_APPCoin *APPCoinCaller) PendingSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "pendingSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingSeconds is a free data retrieval call binding the contract method 0xc819c70e.
//
// Solidity: function pendingSeconds() view returns(uint256)
func (_APPCoin *APPCoinSession) PendingSeconds() (*big.Int, error) {
	return _APPCoin.Contract.PendingSeconds(&_APPCoin.CallOpts)
}

// PendingSeconds is a free data retrieval call binding the contract method 0xc819c70e.
//
// Solidity: function pendingSeconds() view returns(uint256)
func (_APPCoin *APPCoinCallerSession) PendingSeconds() (*big.Int, error) {
	return _APPCoin.Contract.PendingSeconds(&_APPCoin.CallOpts)
}

// ResourceConfigures is a free data retrieval call binding the contract method 0xc86b6ee2.
//
// Solidity: function resourceConfigures(uint32 ) view returns(string resourceId, uint256 weight, uint32 index, uint8 pendingOP, uint256 pendingWeight, uint256 submitSeconds, uint256 requestTimes)
func (_APPCoin *APPCoinCaller) ResourceConfigures(opts *bind.CallOpts, arg0 uint32) (struct {
	ResourceId    string
	Weight        *big.Int
	Index         uint32
	PendingOP     uint8
	PendingWeight *big.Int
	SubmitSeconds *big.Int
	RequestTimes  *big.Int
}, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "resourceConfigures", arg0)

	outstruct := new(struct {
		ResourceId    string
		Weight        *big.Int
		Index         uint32
		PendingOP     uint8
		PendingWeight *big.Int
		SubmitSeconds *big.Int
		RequestTimes  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ResourceId = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Weight = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Index = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.PendingOP = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.PendingWeight = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.SubmitSeconds = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.RequestTimes = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ResourceConfigures is a free data retrieval call binding the contract method 0xc86b6ee2.
//
// Solidity: function resourceConfigures(uint32 ) view returns(string resourceId, uint256 weight, uint32 index, uint8 pendingOP, uint256 pendingWeight, uint256 submitSeconds, uint256 requestTimes)
func (_APPCoin *APPCoinSession) ResourceConfigures(arg0 uint32) (struct {
	ResourceId    string
	Weight        *big.Int
	Index         uint32
	PendingOP     uint8
	PendingWeight *big.Int
	SubmitSeconds *big.Int
	RequestTimes  *big.Int
}, error) {
	return _APPCoin.Contract.ResourceConfigures(&_APPCoin.CallOpts, arg0)
}

// ResourceConfigures is a free data retrieval call binding the contract method 0xc86b6ee2.
//
// Solidity: function resourceConfigures(uint32 ) view returns(string resourceId, uint256 weight, uint32 index, uint8 pendingOP, uint256 pendingWeight, uint256 submitSeconds, uint256 requestTimes)
func (_APPCoin *APPCoinCallerSession) ResourceConfigures(arg0 uint32) (struct {
	ResourceId    string
	Weight        *big.Int
	Index         uint32
	PendingOP     uint8
	PendingWeight *big.Int
	SubmitSeconds *big.Int
	RequestTimes  *big.Int
}, error) {
	return _APPCoin.Contract.ResourceConfigures(&_APPCoin.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_APPCoin *APPCoinCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_APPCoin *APPCoinSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _APPCoin.Contract.SupportsInterface(&_APPCoin.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_APPCoin *APPCoinCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _APPCoin.Contract.SupportsInterface(&_APPCoin.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_APPCoin *APPCoinCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_APPCoin *APPCoinSession) Symbol() (string, error) {
	return _APPCoin.Contract.Symbol(&_APPCoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_APPCoin *APPCoinCallerSession) Symbol() (string, error) {
	return _APPCoin.Contract.Symbol(&_APPCoin.CallOpts)
}

// TotalCharged is a free data retrieval call binding the contract method 0xa6f55d77.
//
// Solidity: function totalCharged() view returns(uint256)
func (_APPCoin *APPCoinCaller) TotalCharged(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "totalCharged")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCharged is a free data retrieval call binding the contract method 0xa6f55d77.
//
// Solidity: function totalCharged() view returns(uint256)
func (_APPCoin *APPCoinSession) TotalCharged() (*big.Int, error) {
	return _APPCoin.Contract.TotalCharged(&_APPCoin.CallOpts)
}

// TotalCharged is a free data retrieval call binding the contract method 0xa6f55d77.
//
// Solidity: function totalCharged() view returns(uint256)
func (_APPCoin *APPCoinCallerSession) TotalCharged() (*big.Int, error) {
	return _APPCoin.Contract.TotalCharged(&_APPCoin.CallOpts)
}

// TotalTakenProfit is a free data retrieval call binding the contract method 0x55297bcf.
//
// Solidity: function totalTakenProfit() view returns(uint256)
func (_APPCoin *APPCoinCaller) TotalTakenProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "totalTakenProfit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalTakenProfit is a free data retrieval call binding the contract method 0x55297bcf.
//
// Solidity: function totalTakenProfit() view returns(uint256)
func (_APPCoin *APPCoinSession) TotalTakenProfit() (*big.Int, error) {
	return _APPCoin.Contract.TotalTakenProfit(&_APPCoin.CallOpts)
}

// TotalTakenProfit is a free data retrieval call binding the contract method 0x55297bcf.
//
// Solidity: function totalTakenProfit() view returns(uint256)
func (_APPCoin *APPCoinCallerSession) TotalTakenProfit() (*big.Int, error) {
	return _APPCoin.Contract.TotalTakenProfit(&_APPCoin.CallOpts)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_APPCoin *APPCoinCaller) Uri(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "uri", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_APPCoin *APPCoinSession) Uri(tokenId *big.Int) (string, error) {
	return _APPCoin.Contract.Uri(&_APPCoin.CallOpts, tokenId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_APPCoin *APPCoinCallerSession) Uri(tokenId *big.Int) (string, error) {
	return _APPCoin.Contract.Uri(&_APPCoin.CallOpts, tokenId)
}

// Users is a free data retrieval call binding the contract method 0x365b98b2.
//
// Solidity: function users(uint256 ) view returns(address)
func (_APPCoin *APPCoinCaller) Users(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "users", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Users is a free data retrieval call binding the contract method 0x365b98b2.
//
// Solidity: function users(uint256 ) view returns(address)
func (_APPCoin *APPCoinSession) Users(arg0 *big.Int) (common.Address, error) {
	return _APPCoin.Contract.Users(&_APPCoin.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0x365b98b2.
//
// Solidity: function users(uint256 ) view returns(address)
func (_APPCoin *APPCoinCallerSession) Users(arg0 *big.Int) (common.Address, error) {
	return _APPCoin.Contract.Users(&_APPCoin.CallOpts, arg0)
}

// Charge is a paid mutator transaction binding the contract method 0x509435c4.
//
// Solidity: function charge(address account, uint256 amount, bytes , (uint32,uint256)[] useDetail) returns()
func (_APPCoin *APPCoinTransactor) Charge(opts *bind.TransactOpts, account common.Address, amount *big.Int, arg2 []byte, useDetail []APPCoinResourceUseDetail) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "charge", account, amount, arg2, useDetail)
}

// Charge is a paid mutator transaction binding the contract method 0x509435c4.
//
// Solidity: function charge(address account, uint256 amount, bytes , (uint32,uint256)[] useDetail) returns()
func (_APPCoin *APPCoinSession) Charge(account common.Address, amount *big.Int, arg2 []byte, useDetail []APPCoinResourceUseDetail) (*types.Transaction, error) {
	return _APPCoin.Contract.Charge(&_APPCoin.TransactOpts, account, amount, arg2, useDetail)
}

// Charge is a paid mutator transaction binding the contract method 0x509435c4.
//
// Solidity: function charge(address account, uint256 amount, bytes , (uint32,uint256)[] useDetail) returns()
func (_APPCoin *APPCoinTransactorSession) Charge(account common.Address, amount *big.Int, arg2 []byte, useDetail []APPCoinResourceUseDetail) (*types.Transaction, error) {
	return _APPCoin.Contract.Charge(&_APPCoin.TransactOpts, account, amount, arg2, useDetail)
}

// ChargeBatch is a paid mutator transaction binding the contract method 0x0664ec3b.
//
// Solidity: function chargeBatch((address,uint256,bytes,(uint32,uint256)[])[] requestArray) returns()
func (_APPCoin *APPCoinTransactor) ChargeBatch(opts *bind.TransactOpts, requestArray []APPCoinChargeRequest) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "chargeBatch", requestArray)
}

// ChargeBatch is a paid mutator transaction binding the contract method 0x0664ec3b.
//
// Solidity: function chargeBatch((address,uint256,bytes,(uint32,uint256)[])[] requestArray) returns()
func (_APPCoin *APPCoinSession) ChargeBatch(requestArray []APPCoinChargeRequest) (*types.Transaction, error) {
	return _APPCoin.Contract.ChargeBatch(&_APPCoin.TransactOpts, requestArray)
}

// ChargeBatch is a paid mutator transaction binding the contract method 0x0664ec3b.
//
// Solidity: function chargeBatch((address,uint256,bytes,(uint32,uint256)[])[] requestArray) returns()
func (_APPCoin *APPCoinTransactorSession) ChargeBatch(requestArray []APPCoinChargeRequest) (*types.Transaction, error) {
	return _APPCoin.Contract.ChargeBatch(&_APPCoin.TransactOpts, requestArray)
}

// ConfigResource is a paid mutator transaction binding the contract method 0xf2721b7a.
//
// Solidity: function configResource((uint32,string,uint256,uint8) entry) returns()
func (_APPCoin *APPCoinTransactor) ConfigResource(opts *bind.TransactOpts, entry AppConfigConfigRequest) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "configResource", entry)
}

// ConfigResource is a paid mutator transaction binding the contract method 0xf2721b7a.
//
// Solidity: function configResource((uint32,string,uint256,uint8) entry) returns()
func (_APPCoin *APPCoinSession) ConfigResource(entry AppConfigConfigRequest) (*types.Transaction, error) {
	return _APPCoin.Contract.ConfigResource(&_APPCoin.TransactOpts, entry)
}

// ConfigResource is a paid mutator transaction binding the contract method 0xf2721b7a.
//
// Solidity: function configResource((uint32,string,uint256,uint8) entry) returns()
func (_APPCoin *APPCoinTransactorSession) ConfigResource(entry AppConfigConfigRequest) (*types.Transaction, error) {
	return _APPCoin.Contract.ConfigResource(&_APPCoin.TransactOpts, entry)
}

// ConfigResourceBatch is a paid mutator transaction binding the contract method 0x10b270aa.
//
// Solidity: function configResourceBatch((uint32,string,uint256,uint8)[] entries) returns()
func (_APPCoin *APPCoinTransactor) ConfigResourceBatch(opts *bind.TransactOpts, entries []AppConfigConfigRequest) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "configResourceBatch", entries)
}

// ConfigResourceBatch is a paid mutator transaction binding the contract method 0x10b270aa.
//
// Solidity: function configResourceBatch((uint32,string,uint256,uint8)[] entries) returns()
func (_APPCoin *APPCoinSession) ConfigResourceBatch(entries []AppConfigConfigRequest) (*types.Transaction, error) {
	return _APPCoin.Contract.ConfigResourceBatch(&_APPCoin.TransactOpts, entries)
}

// ConfigResourceBatch is a paid mutator transaction binding the contract method 0x10b270aa.
//
// Solidity: function configResourceBatch((uint32,string,uint256,uint8)[] entries) returns()
func (_APPCoin *APPCoinTransactorSession) ConfigResourceBatch(entries []AppConfigConfigRequest) (*types.Transaction, error) {
	return _APPCoin.Contract.ConfigResourceBatch(&_APPCoin.TransactOpts, entries)
}

// FlushPendingConfig is a paid mutator transaction binding the contract method 0xb8f2272a.
//
// Solidity: function flushPendingConfig() returns()
func (_APPCoin *APPCoinTransactor) FlushPendingConfig(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "flushPendingConfig")
}

// FlushPendingConfig is a paid mutator transaction binding the contract method 0xb8f2272a.
//
// Solidity: function flushPendingConfig() returns()
func (_APPCoin *APPCoinSession) FlushPendingConfig() (*types.Transaction, error) {
	return _APPCoin.Contract.FlushPendingConfig(&_APPCoin.TransactOpts)
}

// FlushPendingConfig is a paid mutator transaction binding the contract method 0xb8f2272a.
//
// Solidity: function flushPendingConfig() returns()
func (_APPCoin *APPCoinTransactorSession) FlushPendingConfig() (*types.Transaction, error) {
	return _APPCoin.Contract.FlushPendingConfig(&_APPCoin.TransactOpts)
}

// ForceWithdraw is a paid mutator transaction binding the contract method 0x7be80b39.
//
// Solidity: function forceWithdraw() returns()
func (_APPCoin *APPCoinTransactor) ForceWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "forceWithdraw")
}

// ForceWithdraw is a paid mutator transaction binding the contract method 0x7be80b39.
//
// Solidity: function forceWithdraw() returns()
func (_APPCoin *APPCoinSession) ForceWithdraw() (*types.Transaction, error) {
	return _APPCoin.Contract.ForceWithdraw(&_APPCoin.TransactOpts)
}

// ForceWithdraw is a paid mutator transaction binding the contract method 0x7be80b39.
//
// Solidity: function forceWithdraw() returns()
func (_APPCoin *APPCoinTransactorSession) ForceWithdraw() (*types.Transaction, error) {
	return _APPCoin.Contract.ForceWithdraw(&_APPCoin.TransactOpts)
}

// Freeze is a paid mutator transaction binding the contract method 0xbf120ae5.
//
// Solidity: function freeze(address acc, bool f) returns()
func (_APPCoin *APPCoinTransactor) Freeze(opts *bind.TransactOpts, acc common.Address, f bool) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "freeze", acc, f)
}

// Freeze is a paid mutator transaction binding the contract method 0xbf120ae5.
//
// Solidity: function freeze(address acc, bool f) returns()
func (_APPCoin *APPCoinSession) Freeze(acc common.Address, f bool) (*types.Transaction, error) {
	return _APPCoin.Contract.Freeze(&_APPCoin.TransactOpts, acc, f)
}

// Freeze is a paid mutator transaction binding the contract method 0xbf120ae5.
//
// Solidity: function freeze(address acc, bool f) returns()
func (_APPCoin *APPCoinTransactorSession) Freeze(acc common.Address, f bool) (*types.Transaction, error) {
	return _APPCoin.Contract.Freeze(&_APPCoin.TransactOpts, acc, f)
}

// Init is a paid mutator transaction binding the contract method 0x9c26dd41.
//
// Solidity: function init(address apiCoin_, address appOwner_, string name_, string symbol_, string uri_, uint256 defaultWeight) returns()
func (_APPCoin *APPCoinTransactor) Init(opts *bind.TransactOpts, apiCoin_ common.Address, appOwner_ common.Address, name_ string, symbol_ string, uri_ string, defaultWeight *big.Int) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "init", apiCoin_, appOwner_, name_, symbol_, uri_, defaultWeight)
}

// Init is a paid mutator transaction binding the contract method 0x9c26dd41.
//
// Solidity: function init(address apiCoin_, address appOwner_, string name_, string symbol_, string uri_, uint256 defaultWeight) returns()
func (_APPCoin *APPCoinSession) Init(apiCoin_ common.Address, appOwner_ common.Address, name_ string, symbol_ string, uri_ string, defaultWeight *big.Int) (*types.Transaction, error) {
	return _APPCoin.Contract.Init(&_APPCoin.TransactOpts, apiCoin_, appOwner_, name_, symbol_, uri_, defaultWeight)
}

// Init is a paid mutator transaction binding the contract method 0x9c26dd41.
//
// Solidity: function init(address apiCoin_, address appOwner_, string name_, string symbol_, string uri_, uint256 defaultWeight) returns()
func (_APPCoin *APPCoinTransactorSession) Init(apiCoin_ common.Address, appOwner_ common.Address, name_ string, symbol_ string, uri_ string, defaultWeight *big.Int) (*types.Transaction, error) {
	return _APPCoin.Contract.Init(&_APPCoin.TransactOpts, apiCoin_, appOwner_, name_, symbol_, uri_, defaultWeight)
}

// InitOwner is a paid mutator transaction binding the contract method 0x0d009297.
//
// Solidity: function initOwner(address owner_) returns()
func (_APPCoin *APPCoinTransactor) InitOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "initOwner", owner_)
}

// InitOwner is a paid mutator transaction binding the contract method 0x0d009297.
//
// Solidity: function initOwner(address owner_) returns()
func (_APPCoin *APPCoinSession) InitOwner(owner_ common.Address) (*types.Transaction, error) {
	return _APPCoin.Contract.InitOwner(&_APPCoin.TransactOpts, owner_)
}

// InitOwner is a paid mutator transaction binding the contract method 0x0d009297.
//
// Solidity: function initOwner(address owner_) returns()
func (_APPCoin *APPCoinTransactorSession) InitOwner(owner_ common.Address) (*types.Transaction, error) {
	return _APPCoin.Contract.InitOwner(&_APPCoin.TransactOpts, owner_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_APPCoin *APPCoinTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_APPCoin *APPCoinSession) Pause() (*types.Transaction, error) {
	return _APPCoin.Contract.Pause(&_APPCoin.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_APPCoin *APPCoinTransactorSession) Pause() (*types.Transaction, error) {
	return _APPCoin.Contract.Pause(&_APPCoin.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0xfa89401a.
//
// Solidity: function refund(address account) returns()
func (_APPCoin *APPCoinTransactor) Refund(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "refund", account)
}

// Refund is a paid mutator transaction binding the contract method 0xfa89401a.
//
// Solidity: function refund(address account) returns()
func (_APPCoin *APPCoinSession) Refund(account common.Address) (*types.Transaction, error) {
	return _APPCoin.Contract.Refund(&_APPCoin.TransactOpts, account)
}

// Refund is a paid mutator transaction binding the contract method 0xfa89401a.
//
// Solidity: function refund(address account) returns()
func (_APPCoin *APPCoinTransactorSession) Refund(account common.Address) (*types.Transaction, error) {
	return _APPCoin.Contract.Refund(&_APPCoin.TransactOpts, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_APPCoin *APPCoinTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_APPCoin *APPCoinSession) RenounceOwnership() (*types.Transaction, error) {
	return _APPCoin.Contract.RenounceOwnership(&_APPCoin.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_APPCoin *APPCoinTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _APPCoin.Contract.RenounceOwnership(&_APPCoin.TransactOpts)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_APPCoin *APPCoinTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_APPCoin *APPCoinSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _APPCoin.Contract.SafeBatchTransferFrom(&_APPCoin.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_APPCoin *APPCoinTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _APPCoin.Contract.SafeBatchTransferFrom(&_APPCoin.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_APPCoin *APPCoinTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_APPCoin *APPCoinSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _APPCoin.Contract.SafeTransferFrom(&_APPCoin.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_APPCoin *APPCoinTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _APPCoin.Contract.SafeTransferFrom(&_APPCoin.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_APPCoin *APPCoinTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_APPCoin *APPCoinSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _APPCoin.Contract.SetApprovalForAll(&_APPCoin.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_APPCoin *APPCoinTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _APPCoin.Contract.SetApprovalForAll(&_APPCoin.TransactOpts, operator, approved)
}

// SetForceWithdrawDelay is a paid mutator transaction binding the contract method 0xcb5f869e.
//
// Solidity: function setForceWithdrawDelay(uint256 delay) returns()
func (_APPCoin *APPCoinTransactor) SetForceWithdrawDelay(opts *bind.TransactOpts, delay *big.Int) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "setForceWithdrawDelay", delay)
}

// SetForceWithdrawDelay is a paid mutator transaction binding the contract method 0xcb5f869e.
//
// Solidity: function setForceWithdrawDelay(uint256 delay) returns()
func (_APPCoin *APPCoinSession) SetForceWithdrawDelay(delay *big.Int) (*types.Transaction, error) {
	return _APPCoin.Contract.SetForceWithdrawDelay(&_APPCoin.TransactOpts, delay)
}

// SetForceWithdrawDelay is a paid mutator transaction binding the contract method 0xcb5f869e.
//
// Solidity: function setForceWithdrawDelay(uint256 delay) returns()
func (_APPCoin *APPCoinTransactorSession) SetForceWithdrawDelay(delay *big.Int) (*types.Transaction, error) {
	return _APPCoin.Contract.SetForceWithdrawDelay(&_APPCoin.TransactOpts, delay)
}

// SetPendingSeconds is a paid mutator transaction binding the contract method 0xacb00a11.
//
// Solidity: function setPendingSeconds(uint256 seconds_) returns()
func (_APPCoin *APPCoinTransactor) SetPendingSeconds(opts *bind.TransactOpts, seconds_ *big.Int) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "setPendingSeconds", seconds_)
}

// SetPendingSeconds is a paid mutator transaction binding the contract method 0xacb00a11.
//
// Solidity: function setPendingSeconds(uint256 seconds_) returns()
func (_APPCoin *APPCoinSession) SetPendingSeconds(seconds_ *big.Int) (*types.Transaction, error) {
	return _APPCoin.Contract.SetPendingSeconds(&_APPCoin.TransactOpts, seconds_)
}

// SetPendingSeconds is a paid mutator transaction binding the contract method 0xacb00a11.
//
// Solidity: function setPendingSeconds(uint256 seconds_) returns()
func (_APPCoin *APPCoinTransactorSession) SetPendingSeconds(seconds_ *big.Int) (*types.Transaction, error) {
	return _APPCoin.Contract.SetPendingSeconds(&_APPCoin.TransactOpts, seconds_)
}

// TakeProfit is a paid mutator transaction binding the contract method 0x0cec2f17.
//
// Solidity: function takeProfit(address to, uint256 amount) returns()
func (_APPCoin *APPCoinTransactor) TakeProfit(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "takeProfit", to, amount)
}

// TakeProfit is a paid mutator transaction binding the contract method 0x0cec2f17.
//
// Solidity: function takeProfit(address to, uint256 amount) returns()
func (_APPCoin *APPCoinSession) TakeProfit(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _APPCoin.Contract.TakeProfit(&_APPCoin.TransactOpts, to, amount)
}

// TakeProfit is a paid mutator transaction binding the contract method 0x0cec2f17.
//
// Solidity: function takeProfit(address to, uint256 amount) returns()
func (_APPCoin *APPCoinTransactorSession) TakeProfit(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _APPCoin.Contract.TakeProfit(&_APPCoin.TransactOpts, to, amount)
}

// TokensReceived is a paid mutator transaction binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address from, address , uint256 amount, bytes , bytes ) returns()
func (_APPCoin *APPCoinTransactor) TokensReceived(opts *bind.TransactOpts, arg0 common.Address, from common.Address, arg2 common.Address, amount *big.Int, arg4 []byte, arg5 []byte) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "tokensReceived", arg0, from, arg2, amount, arg4, arg5)
}

// TokensReceived is a paid mutator transaction binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address from, address , uint256 amount, bytes , bytes ) returns()
func (_APPCoin *APPCoinSession) TokensReceived(arg0 common.Address, from common.Address, arg2 common.Address, amount *big.Int, arg4 []byte, arg5 []byte) (*types.Transaction, error) {
	return _APPCoin.Contract.TokensReceived(&_APPCoin.TransactOpts, arg0, from, arg2, amount, arg4, arg5)
}

// TokensReceived is a paid mutator transaction binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address from, address , uint256 amount, bytes , bytes ) returns()
func (_APPCoin *APPCoinTransactorSession) TokensReceived(arg0 common.Address, from common.Address, arg2 common.Address, amount *big.Int, arg4 []byte, arg5 []byte) (*types.Transaction, error) {
	return _APPCoin.Contract.TokensReceived(&_APPCoin.TransactOpts, arg0, from, arg2, amount, arg4, arg5)
}

// TransferAppOwner is a paid mutator transaction binding the contract method 0xd3fb43a3.
//
// Solidity: function transferAppOwner(address to, address controller) returns()
func (_APPCoin *APPCoinTransactor) TransferAppOwner(opts *bind.TransactOpts, to common.Address, controller common.Address) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "transferAppOwner", to, controller)
}

// TransferAppOwner is a paid mutator transaction binding the contract method 0xd3fb43a3.
//
// Solidity: function transferAppOwner(address to, address controller) returns()
func (_APPCoin *APPCoinSession) TransferAppOwner(to common.Address, controller common.Address) (*types.Transaction, error) {
	return _APPCoin.Contract.TransferAppOwner(&_APPCoin.TransactOpts, to, controller)
}

// TransferAppOwner is a paid mutator transaction binding the contract method 0xd3fb43a3.
//
// Solidity: function transferAppOwner(address to, address controller) returns()
func (_APPCoin *APPCoinTransactorSession) TransferAppOwner(to common.Address, controller common.Address) (*types.Transaction, error) {
	return _APPCoin.Contract.TransferAppOwner(&_APPCoin.TransactOpts, to, controller)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_APPCoin *APPCoinTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_APPCoin *APPCoinSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _APPCoin.Contract.TransferOwnership(&_APPCoin.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_APPCoin *APPCoinTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _APPCoin.Contract.TransferOwnership(&_APPCoin.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_APPCoin *APPCoinTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_APPCoin *APPCoinSession) Unpause() (*types.Transaction, error) {
	return _APPCoin.Contract.Unpause(&_APPCoin.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_APPCoin *APPCoinTransactorSession) Unpause() (*types.Transaction, error) {
	return _APPCoin.Contract.Unpause(&_APPCoin.TransactOpts)
}

// WithdrawRequest is a paid mutator transaction binding the contract method 0x380e687a.
//
// Solidity: function withdrawRequest() returns()
func (_APPCoin *APPCoinTransactor) WithdrawRequest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "withdrawRequest")
}

// WithdrawRequest is a paid mutator transaction binding the contract method 0x380e687a.
//
// Solidity: function withdrawRequest() returns()
func (_APPCoin *APPCoinSession) WithdrawRequest() (*types.Transaction, error) {
	return _APPCoin.Contract.WithdrawRequest(&_APPCoin.TransactOpts)
}

// WithdrawRequest is a paid mutator transaction binding the contract method 0x380e687a.
//
// Solidity: function withdrawRequest() returns()
func (_APPCoin *APPCoinTransactorSession) WithdrawRequest() (*types.Transaction, error) {
	return _APPCoin.Contract.WithdrawRequest(&_APPCoin.TransactOpts)
}

// APPCoinAppOwnerChangedIterator is returned from FilterAppOwnerChanged and is used to iterate over the raw logs and unpacked data for AppOwnerChanged events raised by the APPCoin contract.
type APPCoinAppOwnerChangedIterator struct {
	Event *APPCoinAppOwnerChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *APPCoinAppOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(APPCoinAppOwnerChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(APPCoinAppOwnerChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *APPCoinAppOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *APPCoinAppOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// APPCoinAppOwnerChanged represents a AppOwnerChanged event raised by the APPCoin contract.
type APPCoinAppOwnerChanged struct {
	To  common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAppOwnerChanged is a free log retrieval operation binding the contract event 0xf90c9f5ec06d0b2a7c2a8bde8a2a10b1813a6eb0eaffe5de9f5bf28e5c092fd1.
//
// Solidity: event AppOwnerChanged(address indexed to)
func (_APPCoin *APPCoinFilterer) FilterAppOwnerChanged(opts *bind.FilterOpts, to []common.Address) (*APPCoinAppOwnerChangedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _APPCoin.contract.FilterLogs(opts, "AppOwnerChanged", toRule)
	if err != nil {
		return nil, err
	}
	return &APPCoinAppOwnerChangedIterator{contract: _APPCoin.contract, event: "AppOwnerChanged", logs: logs, sub: sub}, nil
}

// WatchAppOwnerChanged is a free log subscription operation binding the contract event 0xf90c9f5ec06d0b2a7c2a8bde8a2a10b1813a6eb0eaffe5de9f5bf28e5c092fd1.
//
// Solidity: event AppOwnerChanged(address indexed to)
func (_APPCoin *APPCoinFilterer) WatchAppOwnerChanged(opts *bind.WatchOpts, sink chan<- *APPCoinAppOwnerChanged, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _APPCoin.contract.WatchLogs(opts, "AppOwnerChanged", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(APPCoinAppOwnerChanged)
				if err := _APPCoin.contract.UnpackLog(event, "AppOwnerChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAppOwnerChanged is a log parse operation binding the contract event 0xf90c9f5ec06d0b2a7c2a8bde8a2a10b1813a6eb0eaffe5de9f5bf28e5c092fd1.
//
// Solidity: event AppOwnerChanged(address indexed to)
func (_APPCoin *APPCoinFilterer) ParseAppOwnerChanged(log types.Log) (*APPCoinAppOwnerChanged, error) {
	event := new(APPCoinAppOwnerChanged)
	if err := _APPCoin.contract.UnpackLog(event, "AppOwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// APPCoinApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the APPCoin contract.
type APPCoinApprovalForAllIterator struct {
	Event *APPCoinApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *APPCoinApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(APPCoinApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(APPCoinApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *APPCoinApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *APPCoinApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// APPCoinApprovalForAll represents a ApprovalForAll event raised by the APPCoin contract.
type APPCoinApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_APPCoin *APPCoinFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*APPCoinApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _APPCoin.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &APPCoinApprovalForAllIterator{contract: _APPCoin.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_APPCoin *APPCoinFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *APPCoinApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _APPCoin.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(APPCoinApprovalForAll)
				if err := _APPCoin.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_APPCoin *APPCoinFilterer) ParseApprovalForAll(log types.Log) (*APPCoinApprovalForAll, error) {
	event := new(APPCoinApprovalForAll)
	if err := _APPCoin.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// APPCoinFrozenIterator is returned from FilterFrozen and is used to iterate over the raw logs and unpacked data for Frozen events raised by the APPCoin contract.
type APPCoinFrozenIterator struct {
	Event *APPCoinFrozen // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *APPCoinFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(APPCoinFrozen)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(APPCoinFrozen)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *APPCoinFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *APPCoinFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// APPCoinFrozen represents a Frozen event raised by the APPCoin contract.
type APPCoinFrozen struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFrozen is a free log retrieval operation binding the contract event 0x8a5c4736a33c7b7f29a2c34ea9ff9608afc5718d56f6fd6dcbd2d3711a1a4913.
//
// Solidity: event Frozen(address indexed addr)
func (_APPCoin *APPCoinFilterer) FilterFrozen(opts *bind.FilterOpts, addr []common.Address) (*APPCoinFrozenIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _APPCoin.contract.FilterLogs(opts, "Frozen", addrRule)
	if err != nil {
		return nil, err
	}
	return &APPCoinFrozenIterator{contract: _APPCoin.contract, event: "Frozen", logs: logs, sub: sub}, nil
}

// WatchFrozen is a free log subscription operation binding the contract event 0x8a5c4736a33c7b7f29a2c34ea9ff9608afc5718d56f6fd6dcbd2d3711a1a4913.
//
// Solidity: event Frozen(address indexed addr)
func (_APPCoin *APPCoinFilterer) WatchFrozen(opts *bind.WatchOpts, sink chan<- *APPCoinFrozen, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _APPCoin.contract.WatchLogs(opts, "Frozen", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(APPCoinFrozen)
				if err := _APPCoin.contract.UnpackLog(event, "Frozen", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFrozen is a log parse operation binding the contract event 0x8a5c4736a33c7b7f29a2c34ea9ff9608afc5718d56f6fd6dcbd2d3711a1a4913.
//
// Solidity: event Frozen(address indexed addr)
func (_APPCoin *APPCoinFilterer) ParseFrozen(log types.Log) (*APPCoinFrozen, error) {
	event := new(APPCoinFrozen)
	if err := _APPCoin.contract.UnpackLog(event, "Frozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// APPCoinOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the APPCoin contract.
type APPCoinOwnershipTransferredIterator struct {
	Event *APPCoinOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *APPCoinOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(APPCoinOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(APPCoinOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *APPCoinOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *APPCoinOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// APPCoinOwnershipTransferred represents a OwnershipTransferred event raised by the APPCoin contract.
type APPCoinOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_APPCoin *APPCoinFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*APPCoinOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _APPCoin.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &APPCoinOwnershipTransferredIterator{contract: _APPCoin.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_APPCoin *APPCoinFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *APPCoinOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _APPCoin.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(APPCoinOwnershipTransferred)
				if err := _APPCoin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_APPCoin *APPCoinFilterer) ParseOwnershipTransferred(log types.Log) (*APPCoinOwnershipTransferred, error) {
	event := new(APPCoinOwnershipTransferred)
	if err := _APPCoin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// APPCoinPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the APPCoin contract.
type APPCoinPausedIterator struct {
	Event *APPCoinPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *APPCoinPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(APPCoinPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(APPCoinPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *APPCoinPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *APPCoinPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// APPCoinPaused represents a Paused event raised by the APPCoin contract.
type APPCoinPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_APPCoin *APPCoinFilterer) FilterPaused(opts *bind.FilterOpts) (*APPCoinPausedIterator, error) {

	logs, sub, err := _APPCoin.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &APPCoinPausedIterator{contract: _APPCoin.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_APPCoin *APPCoinFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *APPCoinPaused) (event.Subscription, error) {

	logs, sub, err := _APPCoin.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(APPCoinPaused)
				if err := _APPCoin.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_APPCoin *APPCoinFilterer) ParsePaused(log types.Log) (*APPCoinPaused, error) {
	event := new(APPCoinPaused)
	if err := _APPCoin.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// APPCoinResourceChangedIterator is returned from FilterResourceChanged and is used to iterate over the raw logs and unpacked data for ResourceChanged events raised by the APPCoin contract.
type APPCoinResourceChangedIterator struct {
	Event *APPCoinResourceChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *APPCoinResourceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(APPCoinResourceChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(APPCoinResourceChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *APPCoinResourceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *APPCoinResourceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// APPCoinResourceChanged represents a ResourceChanged event raised by the APPCoin contract.
type APPCoinResourceChanged struct {
	Id     uint32
	Weight *big.Int
	Op     uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterResourceChanged is a free log retrieval operation binding the contract event 0x754de9a177534372342453272825c15351df3e7886699c2fece028d09379c400.
//
// Solidity: event ResourceChanged(uint32 indexed id, uint256 indexed weight, uint8 indexed op)
func (_APPCoin *APPCoinFilterer) FilterResourceChanged(opts *bind.FilterOpts, id []uint32, weight []*big.Int, op []uint8) (*APPCoinResourceChangedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var weightRule []interface{}
	for _, weightItem := range weight {
		weightRule = append(weightRule, weightItem)
	}
	var opRule []interface{}
	for _, opItem := range op {
		opRule = append(opRule, opItem)
	}

	logs, sub, err := _APPCoin.contract.FilterLogs(opts, "ResourceChanged", idRule, weightRule, opRule)
	if err != nil {
		return nil, err
	}
	return &APPCoinResourceChangedIterator{contract: _APPCoin.contract, event: "ResourceChanged", logs: logs, sub: sub}, nil
}

// WatchResourceChanged is a free log subscription operation binding the contract event 0x754de9a177534372342453272825c15351df3e7886699c2fece028d09379c400.
//
// Solidity: event ResourceChanged(uint32 indexed id, uint256 indexed weight, uint8 indexed op)
func (_APPCoin *APPCoinFilterer) WatchResourceChanged(opts *bind.WatchOpts, sink chan<- *APPCoinResourceChanged, id []uint32, weight []*big.Int, op []uint8) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var weightRule []interface{}
	for _, weightItem := range weight {
		weightRule = append(weightRule, weightItem)
	}
	var opRule []interface{}
	for _, opItem := range op {
		opRule = append(opRule, opItem)
	}

	logs, sub, err := _APPCoin.contract.WatchLogs(opts, "ResourceChanged", idRule, weightRule, opRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(APPCoinResourceChanged)
				if err := _APPCoin.contract.UnpackLog(event, "ResourceChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseResourceChanged is a log parse operation binding the contract event 0x754de9a177534372342453272825c15351df3e7886699c2fece028d09379c400.
//
// Solidity: event ResourceChanged(uint32 indexed id, uint256 indexed weight, uint8 indexed op)
func (_APPCoin *APPCoinFilterer) ParseResourceChanged(log types.Log) (*APPCoinResourceChanged, error) {
	event := new(APPCoinResourceChanged)
	if err := _APPCoin.contract.UnpackLog(event, "ResourceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// APPCoinResourcePendingIterator is returned from FilterResourcePending and is used to iterate over the raw logs and unpacked data for ResourcePending events raised by the APPCoin contract.
type APPCoinResourcePendingIterator struct {
	Event *APPCoinResourcePending // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *APPCoinResourcePendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(APPCoinResourcePending)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(APPCoinResourcePending)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *APPCoinResourcePendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *APPCoinResourcePendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// APPCoinResourcePending represents a ResourcePending event raised by the APPCoin contract.
type APPCoinResourcePending struct {
	Id        uint32
	NewWeight *big.Int
	Op        uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterResourcePending is a free log retrieval operation binding the contract event 0xc2a05ca9bcad8635a6049c1240c2a1d2ed0c20c4f4b091ef84f7af523faf93d5.
//
// Solidity: event ResourcePending(uint32 indexed id, uint256 indexed newWeight, uint8 indexed op)
func (_APPCoin *APPCoinFilterer) FilterResourcePending(opts *bind.FilterOpts, id []uint32, newWeight []*big.Int, op []uint8) (*APPCoinResourcePendingIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var newWeightRule []interface{}
	for _, newWeightItem := range newWeight {
		newWeightRule = append(newWeightRule, newWeightItem)
	}
	var opRule []interface{}
	for _, opItem := range op {
		opRule = append(opRule, opItem)
	}

	logs, sub, err := _APPCoin.contract.FilterLogs(opts, "ResourcePending", idRule, newWeightRule, opRule)
	if err != nil {
		return nil, err
	}
	return &APPCoinResourcePendingIterator{contract: _APPCoin.contract, event: "ResourcePending", logs: logs, sub: sub}, nil
}

// WatchResourcePending is a free log subscription operation binding the contract event 0xc2a05ca9bcad8635a6049c1240c2a1d2ed0c20c4f4b091ef84f7af523faf93d5.
//
// Solidity: event ResourcePending(uint32 indexed id, uint256 indexed newWeight, uint8 indexed op)
func (_APPCoin *APPCoinFilterer) WatchResourcePending(opts *bind.WatchOpts, sink chan<- *APPCoinResourcePending, id []uint32, newWeight []*big.Int, op []uint8) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var newWeightRule []interface{}
	for _, newWeightItem := range newWeight {
		newWeightRule = append(newWeightRule, newWeightItem)
	}
	var opRule []interface{}
	for _, opItem := range op {
		opRule = append(opRule, opItem)
	}

	logs, sub, err := _APPCoin.contract.WatchLogs(opts, "ResourcePending", idRule, newWeightRule, opRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(APPCoinResourcePending)
				if err := _APPCoin.contract.UnpackLog(event, "ResourcePending", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseResourcePending is a log parse operation binding the contract event 0xc2a05ca9bcad8635a6049c1240c2a1d2ed0c20c4f4b091ef84f7af523faf93d5.
//
// Solidity: event ResourcePending(uint32 indexed id, uint256 indexed newWeight, uint8 indexed op)
func (_APPCoin *APPCoinFilterer) ParseResourcePending(log types.Log) (*APPCoinResourcePending, error) {
	event := new(APPCoinResourcePending)
	if err := _APPCoin.contract.UnpackLog(event, "ResourcePending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// APPCoinTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the APPCoin contract.
type APPCoinTransferBatchIterator struct {
	Event *APPCoinTransferBatch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *APPCoinTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(APPCoinTransferBatch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(APPCoinTransferBatch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *APPCoinTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *APPCoinTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// APPCoinTransferBatch represents a TransferBatch event raised by the APPCoin contract.
type APPCoinTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_APPCoin *APPCoinFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*APPCoinTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _APPCoin.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &APPCoinTransferBatchIterator{contract: _APPCoin.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_APPCoin *APPCoinFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *APPCoinTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _APPCoin.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(APPCoinTransferBatch)
				if err := _APPCoin.contract.UnpackLog(event, "TransferBatch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_APPCoin *APPCoinFilterer) ParseTransferBatch(log types.Log) (*APPCoinTransferBatch, error) {
	event := new(APPCoinTransferBatch)
	if err := _APPCoin.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// APPCoinTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the APPCoin contract.
type APPCoinTransferSingleIterator struct {
	Event *APPCoinTransferSingle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *APPCoinTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(APPCoinTransferSingle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(APPCoinTransferSingle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *APPCoinTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *APPCoinTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// APPCoinTransferSingle represents a TransferSingle event raised by the APPCoin contract.
type APPCoinTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_APPCoin *APPCoinFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*APPCoinTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _APPCoin.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &APPCoinTransferSingleIterator{contract: _APPCoin.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_APPCoin *APPCoinFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *APPCoinTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _APPCoin.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(APPCoinTransferSingle)
				if err := _APPCoin.contract.UnpackLog(event, "TransferSingle", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_APPCoin *APPCoinFilterer) ParseTransferSingle(log types.Log) (*APPCoinTransferSingle, error) {
	event := new(APPCoinTransferSingle)
	if err := _APPCoin.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// APPCoinURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the APPCoin contract.
type APPCoinURIIterator struct {
	Event *APPCoinURI // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *APPCoinURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(APPCoinURI)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(APPCoinURI)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *APPCoinURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *APPCoinURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// APPCoinURI represents a URI event raised by the APPCoin contract.
type APPCoinURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_APPCoin *APPCoinFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*APPCoinURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _APPCoin.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &APPCoinURIIterator{contract: _APPCoin.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_APPCoin *APPCoinFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *APPCoinURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _APPCoin.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(APPCoinURI)
				if err := _APPCoin.contract.UnpackLog(event, "URI", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_APPCoin *APPCoinFilterer) ParseURI(log types.Log) (*APPCoinURI, error) {
	event := new(APPCoinURI)
	if err := _APPCoin.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// APPCoinUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the APPCoin contract.
type APPCoinUnpausedIterator struct {
	Event *APPCoinUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *APPCoinUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(APPCoinUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(APPCoinUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *APPCoinUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *APPCoinUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// APPCoinUnpaused represents a Unpaused event raised by the APPCoin contract.
type APPCoinUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_APPCoin *APPCoinFilterer) FilterUnpaused(opts *bind.FilterOpts) (*APPCoinUnpausedIterator, error) {

	logs, sub, err := _APPCoin.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &APPCoinUnpausedIterator{contract: _APPCoin.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_APPCoin *APPCoinFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *APPCoinUnpaused) (event.Subscription, error) {

	logs, sub, err := _APPCoin.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(APPCoinUnpaused)
				if err := _APPCoin.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_APPCoin *APPCoinFilterer) ParseUnpaused(log types.Log) (*APPCoinUnpaused, error) {
	event := new(APPCoinUnpaused)
	if err := _APPCoin.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// APPCoinWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the APPCoin contract.
type APPCoinWithdrawIterator struct {
	Event *APPCoinWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *APPCoinWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(APPCoinWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(APPCoinWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *APPCoinWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *APPCoinWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// APPCoinWithdraw represents a Withdraw event raised by the APPCoin contract.
type APPCoinWithdraw struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address account, uint256 amount)
func (_APPCoin *APPCoinFilterer) FilterWithdraw(opts *bind.FilterOpts) (*APPCoinWithdrawIterator, error) {

	logs, sub, err := _APPCoin.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &APPCoinWithdrawIterator{contract: _APPCoin.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address account, uint256 amount)
func (_APPCoin *APPCoinFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *APPCoinWithdraw) (event.Subscription, error) {

	logs, sub, err := _APPCoin.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(APPCoinWithdraw)
				if err := _APPCoin.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address account, uint256 amount)
func (_APPCoin *APPCoinFilterer) ParseWithdraw(log types.Log) (*APPCoinWithdraw, error) {
	event := new(APPCoinWithdraw)
	if err := _APPCoin.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
