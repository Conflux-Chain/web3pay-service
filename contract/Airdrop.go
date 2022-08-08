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

// AirdropMetaData contains all meta data concerning the Airdrop contract.
var AirdropMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"AppOwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"Drop\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"Frozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"enumAppConfig.OP\",\"name\":\"op\",\"type\":\"uint8\"}],\"name\":\"ResourceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newWeight\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"enumAppConfig.OP\",\"name\":\"op\",\"type\":\"uint8\"}],\"name\":\"ResourcePending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Spend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FIRST_CONFIG_ID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FT_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"airdrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"to\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"reason\",\"type\":\"string[]\"}],\"name\":\"airdropBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"apiCoin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"appOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOfWithAirdrop\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"airdrop_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"times\",\"type\":\"uint256\"}],\"internalType\":\"structAPPCoin.ResourceUseDetail[]\",\"name\":\"useDetail\",\"type\":\"tuple[]\"}],\"name\":\"charge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"times\",\"type\":\"uint256\"}],\"internalType\":\"structAPPCoin.ResourceUseDetail[]\",\"name\":\"useDetail\",\"type\":\"tuple[]\"}],\"internalType\":\"structAPPCoin.ChargeRequest[]\",\"name\":\"requestArray\",\"type\":\"tuple[]\"}],\"name\":\"chargeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"chargedMapping\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"resourceId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"enumAppConfig.OP\",\"name\":\"op\",\"type\":\"uint8\"}],\"internalType\":\"structAppConfig.ConfigRequest\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"configResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"resourceId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"enumAppConfig.OP\",\"name\":\"op\",\"type\":\"uint8\"}],\"internalType\":\"structAppConfig.ConfigRequest[]\",\"name\":\"entries\",\"type\":\"tuple[]\"}],\"name\":\"configResourceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flushPendingConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceWithdrawDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"f\",\"type\":\"bool\"}],\"name\":\"freeze\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"frozenMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"indexArray\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"apiCoin_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"appOwner_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"defaultWeight\",\"type\":\"uint256\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"initOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"listResources\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"resourceId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"enumAppConfig.OP\",\"name\":\"pendingOP\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"pendingWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"submitSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestTimes\",\"type\":\"uint256\"}],\"internalType\":\"structAppConfig.ConfigEntry[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"listUser\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structAPPCoin.UserCharged[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32[]\",\"name\":\"ids\",\"type\":\"uint32[]\"}],\"name\":\"listUserRequestCounter\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"times\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextConfigId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingIdArray\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"resourceConfigures\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"resourceId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"enumAppConfig.OP\",\"name\":\"pendingOP\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"pendingWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"submitSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestTimes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setForceWithdrawDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"seconds_\",\"type\":\"uint256\"}],\"name\":\"setPendingSeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"takeProfit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"tokensReceived\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalCharged\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalTakenProfit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"transferAppOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"users\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AirdropABI is the input ABI used to generate the binding from.
// Deprecated: Use AirdropMetaData.ABI instead.
var AirdropABI = AirdropMetaData.ABI

// Airdrop is an auto generated Go binding around an Ethereum contract.
type Airdrop struct {
	AirdropCaller     // Read-only binding to the contract
	AirdropTransactor // Write-only binding to the contract
	AirdropFilterer   // Log filterer for contract events
}

// AirdropCaller is an auto generated read-only Go binding around an Ethereum contract.
type AirdropCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirdropTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AirdropTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirdropFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AirdropFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirdropSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AirdropSession struct {
	Contract     *Airdrop          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AirdropCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AirdropCallerSession struct {
	Contract *AirdropCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AirdropTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AirdropTransactorSession struct {
	Contract     *AirdropTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AirdropRaw is an auto generated low-level Go binding around an Ethereum contract.
type AirdropRaw struct {
	Contract *Airdrop // Generic contract binding to access the raw methods on
}

// AirdropCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AirdropCallerRaw struct {
	Contract *AirdropCaller // Generic read-only contract binding to access the raw methods on
}

// AirdropTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AirdropTransactorRaw struct {
	Contract *AirdropTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAirdrop creates a new instance of Airdrop, bound to a specific deployed contract.
func NewAirdrop(address common.Address, backend bind.ContractBackend) (*Airdrop, error) {
	contract, err := bindAirdrop(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Airdrop{AirdropCaller: AirdropCaller{contract: contract}, AirdropTransactor: AirdropTransactor{contract: contract}, AirdropFilterer: AirdropFilterer{contract: contract}}, nil
}

// NewAirdropCaller creates a new read-only instance of Airdrop, bound to a specific deployed contract.
func NewAirdropCaller(address common.Address, caller bind.ContractCaller) (*AirdropCaller, error) {
	contract, err := bindAirdrop(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AirdropCaller{contract: contract}, nil
}

// NewAirdropTransactor creates a new write-only instance of Airdrop, bound to a specific deployed contract.
func NewAirdropTransactor(address common.Address, transactor bind.ContractTransactor) (*AirdropTransactor, error) {
	contract, err := bindAirdrop(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AirdropTransactor{contract: contract}, nil
}

// NewAirdropFilterer creates a new log filterer instance of Airdrop, bound to a specific deployed contract.
func NewAirdropFilterer(address common.Address, filterer bind.ContractFilterer) (*AirdropFilterer, error) {
	contract, err := bindAirdrop(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AirdropFilterer{contract: contract}, nil
}

// bindAirdrop binds a generic wrapper to an already deployed contract.
func bindAirdrop(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AirdropABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Airdrop *AirdropRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Airdrop.Contract.AirdropCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Airdrop *AirdropRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.Contract.AirdropTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Airdrop *AirdropRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Airdrop.Contract.AirdropTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Airdrop *AirdropCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Airdrop.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Airdrop *AirdropTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Airdrop *AirdropTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Airdrop.Contract.contract.Transact(opts, method, params...)
}

// FIRSTCONFIGID is a free data retrieval call binding the contract method 0x0fe58201.
//
// Solidity: function FIRST_CONFIG_ID() view returns(uint32)
func (_Airdrop *AirdropCaller) FIRSTCONFIGID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "FIRST_CONFIG_ID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// FIRSTCONFIGID is a free data retrieval call binding the contract method 0x0fe58201.
//
// Solidity: function FIRST_CONFIG_ID() view returns(uint32)
func (_Airdrop *AirdropSession) FIRSTCONFIGID() (uint32, error) {
	return _Airdrop.Contract.FIRSTCONFIGID(&_Airdrop.CallOpts)
}

// FIRSTCONFIGID is a free data retrieval call binding the contract method 0x0fe58201.
//
// Solidity: function FIRST_CONFIG_ID() view returns(uint32)
func (_Airdrop *AirdropCallerSession) FIRSTCONFIGID() (uint32, error) {
	return _Airdrop.Contract.FIRSTCONFIGID(&_Airdrop.CallOpts)
}

// FTID is a free data retrieval call binding the contract method 0x01052d57.
//
// Solidity: function FT_ID() view returns(uint256)
func (_Airdrop *AirdropCaller) FTID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "FT_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FTID is a free data retrieval call binding the contract method 0x01052d57.
//
// Solidity: function FT_ID() view returns(uint256)
func (_Airdrop *AirdropSession) FTID() (*big.Int, error) {
	return _Airdrop.Contract.FTID(&_Airdrop.CallOpts)
}

// FTID is a free data retrieval call binding the contract method 0x01052d57.
//
// Solidity: function FT_ID() view returns(uint256)
func (_Airdrop *AirdropCallerSession) FTID() (*big.Int, error) {
	return _Airdrop.Contract.FTID(&_Airdrop.CallOpts)
}

// ApiCoin is a free data retrieval call binding the contract method 0x4ebf365a.
//
// Solidity: function apiCoin() view returns(address)
func (_Airdrop *AirdropCaller) ApiCoin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "apiCoin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ApiCoin is a free data retrieval call binding the contract method 0x4ebf365a.
//
// Solidity: function apiCoin() view returns(address)
func (_Airdrop *AirdropSession) ApiCoin() (common.Address, error) {
	return _Airdrop.Contract.ApiCoin(&_Airdrop.CallOpts)
}

// ApiCoin is a free data retrieval call binding the contract method 0x4ebf365a.
//
// Solidity: function apiCoin() view returns(address)
func (_Airdrop *AirdropCallerSession) ApiCoin() (common.Address, error) {
	return _Airdrop.Contract.ApiCoin(&_Airdrop.CallOpts)
}

// AppOwner is a free data retrieval call binding the contract method 0x32b93a86.
//
// Solidity: function appOwner() view returns(address)
func (_Airdrop *AirdropCaller) AppOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "appOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AppOwner is a free data retrieval call binding the contract method 0x32b93a86.
//
// Solidity: function appOwner() view returns(address)
func (_Airdrop *AirdropSession) AppOwner() (common.Address, error) {
	return _Airdrop.Contract.AppOwner(&_Airdrop.CallOpts)
}

// AppOwner is a free data retrieval call binding the contract method 0x32b93a86.
//
// Solidity: function appOwner() view returns(address)
func (_Airdrop *AirdropCallerSession) AppOwner() (common.Address, error) {
	return _Airdrop.Contract.AppOwner(&_Airdrop.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Airdrop *AirdropCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Airdrop *AirdropSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Airdrop.Contract.BalanceOf(&_Airdrop.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Airdrop *AirdropCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Airdrop.Contract.BalanceOf(&_Airdrop.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Airdrop *AirdropCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Airdrop *AirdropSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Airdrop.Contract.BalanceOfBatch(&_Airdrop.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Airdrop *AirdropCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Airdrop.Contract.BalanceOfBatch(&_Airdrop.CallOpts, accounts, ids)
}

// BalanceOfWithAirdrop is a free data retrieval call binding the contract method 0x1215193f.
//
// Solidity: function balanceOfWithAirdrop(address owner) view returns(uint256 total, uint256 airdrop_)
func (_Airdrop *AirdropCaller) BalanceOfWithAirdrop(opts *bind.CallOpts, owner common.Address) (struct {
	Total   *big.Int
	Airdrop *big.Int
}, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "balanceOfWithAirdrop", owner)

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
// Solidity: function balanceOfWithAirdrop(address owner) view returns(uint256 total, uint256 airdrop_)
func (_Airdrop *AirdropSession) BalanceOfWithAirdrop(owner common.Address) (struct {
	Total   *big.Int
	Airdrop *big.Int
}, error) {
	return _Airdrop.Contract.BalanceOfWithAirdrop(&_Airdrop.CallOpts, owner)
}

// BalanceOfWithAirdrop is a free data retrieval call binding the contract method 0x1215193f.
//
// Solidity: function balanceOfWithAirdrop(address owner) view returns(uint256 total, uint256 airdrop_)
func (_Airdrop *AirdropCallerSession) BalanceOfWithAirdrop(owner common.Address) (struct {
	Total   *big.Int
	Airdrop *big.Int
}, error) {
	return _Airdrop.Contract.BalanceOfWithAirdrop(&_Airdrop.CallOpts, owner)
}

// ChargedMapping is a free data retrieval call binding the contract method 0x8ade7975.
//
// Solidity: function chargedMapping(address ) view returns(uint256)
func (_Airdrop *AirdropCaller) ChargedMapping(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "chargedMapping", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChargedMapping is a free data retrieval call binding the contract method 0x8ade7975.
//
// Solidity: function chargedMapping(address ) view returns(uint256)
func (_Airdrop *AirdropSession) ChargedMapping(arg0 common.Address) (*big.Int, error) {
	return _Airdrop.Contract.ChargedMapping(&_Airdrop.CallOpts, arg0)
}

// ChargedMapping is a free data retrieval call binding the contract method 0x8ade7975.
//
// Solidity: function chargedMapping(address ) view returns(uint256)
func (_Airdrop *AirdropCallerSession) ChargedMapping(arg0 common.Address) (*big.Int, error) {
	return _Airdrop.Contract.ChargedMapping(&_Airdrop.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Airdrop *AirdropCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Airdrop *AirdropSession) Decimals() (uint8, error) {
	return _Airdrop.Contract.Decimals(&_Airdrop.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Airdrop *AirdropCallerSession) Decimals() (uint8, error) {
	return _Airdrop.Contract.Decimals(&_Airdrop.CallOpts)
}

// ForceWithdrawDelay is a free data retrieval call binding the contract method 0x6d27d1af.
//
// Solidity: function forceWithdrawDelay() view returns(uint256)
func (_Airdrop *AirdropCaller) ForceWithdrawDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "forceWithdrawDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ForceWithdrawDelay is a free data retrieval call binding the contract method 0x6d27d1af.
//
// Solidity: function forceWithdrawDelay() view returns(uint256)
func (_Airdrop *AirdropSession) ForceWithdrawDelay() (*big.Int, error) {
	return _Airdrop.Contract.ForceWithdrawDelay(&_Airdrop.CallOpts)
}

// ForceWithdrawDelay is a free data retrieval call binding the contract method 0x6d27d1af.
//
// Solidity: function forceWithdrawDelay() view returns(uint256)
func (_Airdrop *AirdropCallerSession) ForceWithdrawDelay() (*big.Int, error) {
	return _Airdrop.Contract.ForceWithdrawDelay(&_Airdrop.CallOpts)
}

// FrozenMap is a free data retrieval call binding the contract method 0xe316468b.
//
// Solidity: function frozenMap(address ) view returns(uint256)
func (_Airdrop *AirdropCaller) FrozenMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "frozenMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FrozenMap is a free data retrieval call binding the contract method 0xe316468b.
//
// Solidity: function frozenMap(address ) view returns(uint256)
func (_Airdrop *AirdropSession) FrozenMap(arg0 common.Address) (*big.Int, error) {
	return _Airdrop.Contract.FrozenMap(&_Airdrop.CallOpts, arg0)
}

// FrozenMap is a free data retrieval call binding the contract method 0xe316468b.
//
// Solidity: function frozenMap(address ) view returns(uint256)
func (_Airdrop *AirdropCallerSession) FrozenMap(arg0 common.Address) (*big.Int, error) {
	return _Airdrop.Contract.FrozenMap(&_Airdrop.CallOpts, arg0)
}

// IndexArray is a free data retrieval call binding the contract method 0xebed9b2c.
//
// Solidity: function indexArray(uint256 ) view returns(uint32)
func (_Airdrop *AirdropCaller) IndexArray(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "indexArray", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// IndexArray is a free data retrieval call binding the contract method 0xebed9b2c.
//
// Solidity: function indexArray(uint256 ) view returns(uint32)
func (_Airdrop *AirdropSession) IndexArray(arg0 *big.Int) (uint32, error) {
	return _Airdrop.Contract.IndexArray(&_Airdrop.CallOpts, arg0)
}

// IndexArray is a free data retrieval call binding the contract method 0xebed9b2c.
//
// Solidity: function indexArray(uint256 ) view returns(uint32)
func (_Airdrop *AirdropCallerSession) IndexArray(arg0 *big.Int) (uint32, error) {
	return _Airdrop.Contract.IndexArray(&_Airdrop.CallOpts, arg0)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Airdrop *AirdropCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Airdrop *AirdropSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Airdrop.Contract.IsApprovedForAll(&_Airdrop.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Airdrop *AirdropCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Airdrop.Contract.IsApprovedForAll(&_Airdrop.CallOpts, account, operator)
}

// ListResources is a free data retrieval call binding the contract method 0x49459cd2.
//
// Solidity: function listResources(uint256 offset, uint256 limit) view returns((string,uint256,uint32,uint8,uint256,uint256,uint256)[], uint256 total)
func (_Airdrop *AirdropCaller) ListResources(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]AppConfigConfigEntry, *big.Int, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "listResources", offset, limit)

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
func (_Airdrop *AirdropSession) ListResources(offset *big.Int, limit *big.Int) ([]AppConfigConfigEntry, *big.Int, error) {
	return _Airdrop.Contract.ListResources(&_Airdrop.CallOpts, offset, limit)
}

// ListResources is a free data retrieval call binding the contract method 0x49459cd2.
//
// Solidity: function listResources(uint256 offset, uint256 limit) view returns((string,uint256,uint32,uint8,uint256,uint256,uint256)[], uint256 total)
func (_Airdrop *AirdropCallerSession) ListResources(offset *big.Int, limit *big.Int) ([]AppConfigConfigEntry, *big.Int, error) {
	return _Airdrop.Contract.ListResources(&_Airdrop.CallOpts, offset, limit)
}

// ListUser is a free data retrieval call binding the contract method 0x3ec36183.
//
// Solidity: function listUser(uint256 offset, uint256 limit) view returns((address,uint256)[], uint256 total)
func (_Airdrop *AirdropCaller) ListUser(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]APPCoinUserCharged, *big.Int, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "listUser", offset, limit)

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
func (_Airdrop *AirdropSession) ListUser(offset *big.Int, limit *big.Int) ([]APPCoinUserCharged, *big.Int, error) {
	return _Airdrop.Contract.ListUser(&_Airdrop.CallOpts, offset, limit)
}

// ListUser is a free data retrieval call binding the contract method 0x3ec36183.
//
// Solidity: function listUser(uint256 offset, uint256 limit) view returns((address,uint256)[], uint256 total)
func (_Airdrop *AirdropCallerSession) ListUser(offset *big.Int, limit *big.Int) ([]APPCoinUserCharged, *big.Int, error) {
	return _Airdrop.Contract.ListUser(&_Airdrop.CallOpts, offset, limit)
}

// ListUserRequestCounter is a free data retrieval call binding the contract method 0xa5e34731.
//
// Solidity: function listUserRequestCounter(address user, uint32[] ids) view returns(uint256[] times)
func (_Airdrop *AirdropCaller) ListUserRequestCounter(opts *bind.CallOpts, user common.Address, ids []uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "listUserRequestCounter", user, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ListUserRequestCounter is a free data retrieval call binding the contract method 0xa5e34731.
//
// Solidity: function listUserRequestCounter(address user, uint32[] ids) view returns(uint256[] times)
func (_Airdrop *AirdropSession) ListUserRequestCounter(user common.Address, ids []uint32) ([]*big.Int, error) {
	return _Airdrop.Contract.ListUserRequestCounter(&_Airdrop.CallOpts, user, ids)
}

// ListUserRequestCounter is a free data retrieval call binding the contract method 0xa5e34731.
//
// Solidity: function listUserRequestCounter(address user, uint32[] ids) view returns(uint256[] times)
func (_Airdrop *AirdropCallerSession) ListUserRequestCounter(user common.Address, ids []uint32) ([]*big.Int, error) {
	return _Airdrop.Contract.ListUserRequestCounter(&_Airdrop.CallOpts, user, ids)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Airdrop *AirdropCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Airdrop *AirdropSession) Name() (string, error) {
	return _Airdrop.Contract.Name(&_Airdrop.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Airdrop *AirdropCallerSession) Name() (string, error) {
	return _Airdrop.Contract.Name(&_Airdrop.CallOpts)
}

// NextConfigId is a free data retrieval call binding the contract method 0x99d726c7.
//
// Solidity: function nextConfigId() view returns(uint32)
func (_Airdrop *AirdropCaller) NextConfigId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "nextConfigId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NextConfigId is a free data retrieval call binding the contract method 0x99d726c7.
//
// Solidity: function nextConfigId() view returns(uint32)
func (_Airdrop *AirdropSession) NextConfigId() (uint32, error) {
	return _Airdrop.Contract.NextConfigId(&_Airdrop.CallOpts)
}

// NextConfigId is a free data retrieval call binding the contract method 0x99d726c7.
//
// Solidity: function nextConfigId() view returns(uint32)
func (_Airdrop *AirdropCallerSession) NextConfigId() (uint32, error) {
	return _Airdrop.Contract.NextConfigId(&_Airdrop.CallOpts)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_Airdrop *AirdropCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_Airdrop *AirdropSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _Airdrop.Contract.OnERC1155BatchReceived(&_Airdrop.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_Airdrop *AirdropCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _Airdrop.Contract.OnERC1155BatchReceived(&_Airdrop.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_Airdrop *AirdropCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_Airdrop *AirdropSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _Airdrop.Contract.OnERC1155Received(&_Airdrop.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_Airdrop *AirdropCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _Airdrop.Contract.OnERC1155Received(&_Airdrop.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Airdrop *AirdropCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Airdrop *AirdropSession) Owner() (common.Address, error) {
	return _Airdrop.Contract.Owner(&_Airdrop.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Airdrop *AirdropCallerSession) Owner() (common.Address, error) {
	return _Airdrop.Contract.Owner(&_Airdrop.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Airdrop *AirdropCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Airdrop *AirdropSession) Paused() (bool, error) {
	return _Airdrop.Contract.Paused(&_Airdrop.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Airdrop *AirdropCallerSession) Paused() (bool, error) {
	return _Airdrop.Contract.Paused(&_Airdrop.CallOpts)
}

// PendingIdArray is a free data retrieval call binding the contract method 0xac8ef2b5.
//
// Solidity: function pendingIdArray(uint256 ) view returns(uint32)
func (_Airdrop *AirdropCaller) PendingIdArray(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "pendingIdArray", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// PendingIdArray is a free data retrieval call binding the contract method 0xac8ef2b5.
//
// Solidity: function pendingIdArray(uint256 ) view returns(uint32)
func (_Airdrop *AirdropSession) PendingIdArray(arg0 *big.Int) (uint32, error) {
	return _Airdrop.Contract.PendingIdArray(&_Airdrop.CallOpts, arg0)
}

// PendingIdArray is a free data retrieval call binding the contract method 0xac8ef2b5.
//
// Solidity: function pendingIdArray(uint256 ) view returns(uint32)
func (_Airdrop *AirdropCallerSession) PendingIdArray(arg0 *big.Int) (uint32, error) {
	return _Airdrop.Contract.PendingIdArray(&_Airdrop.CallOpts, arg0)
}

// PendingSeconds is a free data retrieval call binding the contract method 0xc819c70e.
//
// Solidity: function pendingSeconds() view returns(uint256)
func (_Airdrop *AirdropCaller) PendingSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "pendingSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingSeconds is a free data retrieval call binding the contract method 0xc819c70e.
//
// Solidity: function pendingSeconds() view returns(uint256)
func (_Airdrop *AirdropSession) PendingSeconds() (*big.Int, error) {
	return _Airdrop.Contract.PendingSeconds(&_Airdrop.CallOpts)
}

// PendingSeconds is a free data retrieval call binding the contract method 0xc819c70e.
//
// Solidity: function pendingSeconds() view returns(uint256)
func (_Airdrop *AirdropCallerSession) PendingSeconds() (*big.Int, error) {
	return _Airdrop.Contract.PendingSeconds(&_Airdrop.CallOpts)
}

// ResourceConfigures is a free data retrieval call binding the contract method 0xc86b6ee2.
//
// Solidity: function resourceConfigures(uint32 ) view returns(string resourceId, uint256 weight, uint32 index, uint8 pendingOP, uint256 pendingWeight, uint256 submitSeconds, uint256 requestTimes)
func (_Airdrop *AirdropCaller) ResourceConfigures(opts *bind.CallOpts, arg0 uint32) (struct {
	ResourceId    string
	Weight        *big.Int
	Index         uint32
	PendingOP     uint8
	PendingWeight *big.Int
	SubmitSeconds *big.Int
	RequestTimes  *big.Int
}, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "resourceConfigures", arg0)

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
func (_Airdrop *AirdropSession) ResourceConfigures(arg0 uint32) (struct {
	ResourceId    string
	Weight        *big.Int
	Index         uint32
	PendingOP     uint8
	PendingWeight *big.Int
	SubmitSeconds *big.Int
	RequestTimes  *big.Int
}, error) {
	return _Airdrop.Contract.ResourceConfigures(&_Airdrop.CallOpts, arg0)
}

// ResourceConfigures is a free data retrieval call binding the contract method 0xc86b6ee2.
//
// Solidity: function resourceConfigures(uint32 ) view returns(string resourceId, uint256 weight, uint32 index, uint8 pendingOP, uint256 pendingWeight, uint256 submitSeconds, uint256 requestTimes)
func (_Airdrop *AirdropCallerSession) ResourceConfigures(arg0 uint32) (struct {
	ResourceId    string
	Weight        *big.Int
	Index         uint32
	PendingOP     uint8
	PendingWeight *big.Int
	SubmitSeconds *big.Int
	RequestTimes  *big.Int
}, error) {
	return _Airdrop.Contract.ResourceConfigures(&_Airdrop.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Airdrop *AirdropCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Airdrop *AirdropSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Airdrop.Contract.SupportsInterface(&_Airdrop.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Airdrop *AirdropCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Airdrop.Contract.SupportsInterface(&_Airdrop.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Airdrop *AirdropCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Airdrop *AirdropSession) Symbol() (string, error) {
	return _Airdrop.Contract.Symbol(&_Airdrop.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Airdrop *AirdropCallerSession) Symbol() (string, error) {
	return _Airdrop.Contract.Symbol(&_Airdrop.CallOpts)
}

// TotalCharged is a free data retrieval call binding the contract method 0xa6f55d77.
//
// Solidity: function totalCharged() view returns(uint256)
func (_Airdrop *AirdropCaller) TotalCharged(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "totalCharged")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCharged is a free data retrieval call binding the contract method 0xa6f55d77.
//
// Solidity: function totalCharged() view returns(uint256)
func (_Airdrop *AirdropSession) TotalCharged() (*big.Int, error) {
	return _Airdrop.Contract.TotalCharged(&_Airdrop.CallOpts)
}

// TotalCharged is a free data retrieval call binding the contract method 0xa6f55d77.
//
// Solidity: function totalCharged() view returns(uint256)
func (_Airdrop *AirdropCallerSession) TotalCharged() (*big.Int, error) {
	return _Airdrop.Contract.TotalCharged(&_Airdrop.CallOpts)
}

// TotalTakenProfit is a free data retrieval call binding the contract method 0x55297bcf.
//
// Solidity: function totalTakenProfit() view returns(uint256)
func (_Airdrop *AirdropCaller) TotalTakenProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "totalTakenProfit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalTakenProfit is a free data retrieval call binding the contract method 0x55297bcf.
//
// Solidity: function totalTakenProfit() view returns(uint256)
func (_Airdrop *AirdropSession) TotalTakenProfit() (*big.Int, error) {
	return _Airdrop.Contract.TotalTakenProfit(&_Airdrop.CallOpts)
}

// TotalTakenProfit is a free data retrieval call binding the contract method 0x55297bcf.
//
// Solidity: function totalTakenProfit() view returns(uint256)
func (_Airdrop *AirdropCallerSession) TotalTakenProfit() (*big.Int, error) {
	return _Airdrop.Contract.TotalTakenProfit(&_Airdrop.CallOpts)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_Airdrop *AirdropCaller) Uri(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "uri", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_Airdrop *AirdropSession) Uri(tokenId *big.Int) (string, error) {
	return _Airdrop.Contract.Uri(&_Airdrop.CallOpts, tokenId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_Airdrop *AirdropCallerSession) Uri(tokenId *big.Int) (string, error) {
	return _Airdrop.Contract.Uri(&_Airdrop.CallOpts, tokenId)
}

// Users is a free data retrieval call binding the contract method 0x365b98b2.
//
// Solidity: function users(uint256 ) view returns(address)
func (_Airdrop *AirdropCaller) Users(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "users", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Users is a free data retrieval call binding the contract method 0x365b98b2.
//
// Solidity: function users(uint256 ) view returns(address)
func (_Airdrop *AirdropSession) Users(arg0 *big.Int) (common.Address, error) {
	return _Airdrop.Contract.Users(&_Airdrop.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0x365b98b2.
//
// Solidity: function users(uint256 ) view returns(address)
func (_Airdrop *AirdropCallerSession) Users(arg0 *big.Int) (common.Address, error) {
	return _Airdrop.Contract.Users(&_Airdrop.CallOpts, arg0)
}

// Airdrop is a paid mutator transaction binding the contract method 0x65978e0f.
//
// Solidity: function airdrop(address to, uint256 amount, string reason) returns()
func (_Airdrop *AirdropTransactor) Airdrop(opts *bind.TransactOpts, to common.Address, amount *big.Int, reason string) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "airdrop", to, amount, reason)
}

// Airdrop is a paid mutator transaction binding the contract method 0x65978e0f.
//
// Solidity: function airdrop(address to, uint256 amount, string reason) returns()
func (_Airdrop *AirdropSession) Airdrop(to common.Address, amount *big.Int, reason string) (*types.Transaction, error) {
	return _Airdrop.Contract.Airdrop(&_Airdrop.TransactOpts, to, amount, reason)
}

// Airdrop is a paid mutator transaction binding the contract method 0x65978e0f.
//
// Solidity: function airdrop(address to, uint256 amount, string reason) returns()
func (_Airdrop *AirdropTransactorSession) Airdrop(to common.Address, amount *big.Int, reason string) (*types.Transaction, error) {
	return _Airdrop.Contract.Airdrop(&_Airdrop.TransactOpts, to, amount, reason)
}

// AirdropBatch is a paid mutator transaction binding the contract method 0x9c5c2c32.
//
// Solidity: function airdropBatch(address[] to, uint256[] amount, string[] reason) returns()
func (_Airdrop *AirdropTransactor) AirdropBatch(opts *bind.TransactOpts, to []common.Address, amount []*big.Int, reason []string) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "airdropBatch", to, amount, reason)
}

// AirdropBatch is a paid mutator transaction binding the contract method 0x9c5c2c32.
//
// Solidity: function airdropBatch(address[] to, uint256[] amount, string[] reason) returns()
func (_Airdrop *AirdropSession) AirdropBatch(to []common.Address, amount []*big.Int, reason []string) (*types.Transaction, error) {
	return _Airdrop.Contract.AirdropBatch(&_Airdrop.TransactOpts, to, amount, reason)
}

// AirdropBatch is a paid mutator transaction binding the contract method 0x9c5c2c32.
//
// Solidity: function airdropBatch(address[] to, uint256[] amount, string[] reason) returns()
func (_Airdrop *AirdropTransactorSession) AirdropBatch(to []common.Address, amount []*big.Int, reason []string) (*types.Transaction, error) {
	return _Airdrop.Contract.AirdropBatch(&_Airdrop.TransactOpts, to, amount, reason)
}

// Charge is a paid mutator transaction binding the contract method 0x509435c4.
//
// Solidity: function charge(address account, uint256 amount, bytes data, (uint32,uint256)[] useDetail) returns()
func (_Airdrop *AirdropTransactor) Charge(opts *bind.TransactOpts, account common.Address, amount *big.Int, data []byte, useDetail []APPCoinResourceUseDetail) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "charge", account, amount, data, useDetail)
}

// Charge is a paid mutator transaction binding the contract method 0x509435c4.
//
// Solidity: function charge(address account, uint256 amount, bytes data, (uint32,uint256)[] useDetail) returns()
func (_Airdrop *AirdropSession) Charge(account common.Address, amount *big.Int, data []byte, useDetail []APPCoinResourceUseDetail) (*types.Transaction, error) {
	return _Airdrop.Contract.Charge(&_Airdrop.TransactOpts, account, amount, data, useDetail)
}

// Charge is a paid mutator transaction binding the contract method 0x509435c4.
//
// Solidity: function charge(address account, uint256 amount, bytes data, (uint32,uint256)[] useDetail) returns()
func (_Airdrop *AirdropTransactorSession) Charge(account common.Address, amount *big.Int, data []byte, useDetail []APPCoinResourceUseDetail) (*types.Transaction, error) {
	return _Airdrop.Contract.Charge(&_Airdrop.TransactOpts, account, amount, data, useDetail)
}

// ChargeBatch is a paid mutator transaction binding the contract method 0x0664ec3b.
//
// Solidity: function chargeBatch((address,uint256,bytes,(uint32,uint256)[])[] requestArray) returns()
func (_Airdrop *AirdropTransactor) ChargeBatch(opts *bind.TransactOpts, requestArray []APPCoinChargeRequest) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "chargeBatch", requestArray)
}

// ChargeBatch is a paid mutator transaction binding the contract method 0x0664ec3b.
//
// Solidity: function chargeBatch((address,uint256,bytes,(uint32,uint256)[])[] requestArray) returns()
func (_Airdrop *AirdropSession) ChargeBatch(requestArray []APPCoinChargeRequest) (*types.Transaction, error) {
	return _Airdrop.Contract.ChargeBatch(&_Airdrop.TransactOpts, requestArray)
}

// ChargeBatch is a paid mutator transaction binding the contract method 0x0664ec3b.
//
// Solidity: function chargeBatch((address,uint256,bytes,(uint32,uint256)[])[] requestArray) returns()
func (_Airdrop *AirdropTransactorSession) ChargeBatch(requestArray []APPCoinChargeRequest) (*types.Transaction, error) {
	return _Airdrop.Contract.ChargeBatch(&_Airdrop.TransactOpts, requestArray)
}

// ConfigResource is a paid mutator transaction binding the contract method 0xf2721b7a.
//
// Solidity: function configResource((uint32,string,uint256,uint8) entry) returns()
func (_Airdrop *AirdropTransactor) ConfigResource(opts *bind.TransactOpts, entry AppConfigConfigRequest) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "configResource", entry)
}

// ConfigResource is a paid mutator transaction binding the contract method 0xf2721b7a.
//
// Solidity: function configResource((uint32,string,uint256,uint8) entry) returns()
func (_Airdrop *AirdropSession) ConfigResource(entry AppConfigConfigRequest) (*types.Transaction, error) {
	return _Airdrop.Contract.ConfigResource(&_Airdrop.TransactOpts, entry)
}

// ConfigResource is a paid mutator transaction binding the contract method 0xf2721b7a.
//
// Solidity: function configResource((uint32,string,uint256,uint8) entry) returns()
func (_Airdrop *AirdropTransactorSession) ConfigResource(entry AppConfigConfigRequest) (*types.Transaction, error) {
	return _Airdrop.Contract.ConfigResource(&_Airdrop.TransactOpts, entry)
}

// ConfigResourceBatch is a paid mutator transaction binding the contract method 0x10b270aa.
//
// Solidity: function configResourceBatch((uint32,string,uint256,uint8)[] entries) returns()
func (_Airdrop *AirdropTransactor) ConfigResourceBatch(opts *bind.TransactOpts, entries []AppConfigConfigRequest) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "configResourceBatch", entries)
}

// ConfigResourceBatch is a paid mutator transaction binding the contract method 0x10b270aa.
//
// Solidity: function configResourceBatch((uint32,string,uint256,uint8)[] entries) returns()
func (_Airdrop *AirdropSession) ConfigResourceBatch(entries []AppConfigConfigRequest) (*types.Transaction, error) {
	return _Airdrop.Contract.ConfigResourceBatch(&_Airdrop.TransactOpts, entries)
}

// ConfigResourceBatch is a paid mutator transaction binding the contract method 0x10b270aa.
//
// Solidity: function configResourceBatch((uint32,string,uint256,uint8)[] entries) returns()
func (_Airdrop *AirdropTransactorSession) ConfigResourceBatch(entries []AppConfigConfigRequest) (*types.Transaction, error) {
	return _Airdrop.Contract.ConfigResourceBatch(&_Airdrop.TransactOpts, entries)
}

// FlushPendingConfig is a paid mutator transaction binding the contract method 0xb8f2272a.
//
// Solidity: function flushPendingConfig() returns()
func (_Airdrop *AirdropTransactor) FlushPendingConfig(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "flushPendingConfig")
}

// FlushPendingConfig is a paid mutator transaction binding the contract method 0xb8f2272a.
//
// Solidity: function flushPendingConfig() returns()
func (_Airdrop *AirdropSession) FlushPendingConfig() (*types.Transaction, error) {
	return _Airdrop.Contract.FlushPendingConfig(&_Airdrop.TransactOpts)
}

// FlushPendingConfig is a paid mutator transaction binding the contract method 0xb8f2272a.
//
// Solidity: function flushPendingConfig() returns()
func (_Airdrop *AirdropTransactorSession) FlushPendingConfig() (*types.Transaction, error) {
	return _Airdrop.Contract.FlushPendingConfig(&_Airdrop.TransactOpts)
}

// ForceWithdraw is a paid mutator transaction binding the contract method 0x7be80b39.
//
// Solidity: function forceWithdraw() returns()
func (_Airdrop *AirdropTransactor) ForceWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "forceWithdraw")
}

// ForceWithdraw is a paid mutator transaction binding the contract method 0x7be80b39.
//
// Solidity: function forceWithdraw() returns()
func (_Airdrop *AirdropSession) ForceWithdraw() (*types.Transaction, error) {
	return _Airdrop.Contract.ForceWithdraw(&_Airdrop.TransactOpts)
}

// ForceWithdraw is a paid mutator transaction binding the contract method 0x7be80b39.
//
// Solidity: function forceWithdraw() returns()
func (_Airdrop *AirdropTransactorSession) ForceWithdraw() (*types.Transaction, error) {
	return _Airdrop.Contract.ForceWithdraw(&_Airdrop.TransactOpts)
}

// Freeze is a paid mutator transaction binding the contract method 0xbf120ae5.
//
// Solidity: function freeze(address acc, bool f) returns()
func (_Airdrop *AirdropTransactor) Freeze(opts *bind.TransactOpts, acc common.Address, f bool) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "freeze", acc, f)
}

// Freeze is a paid mutator transaction binding the contract method 0xbf120ae5.
//
// Solidity: function freeze(address acc, bool f) returns()
func (_Airdrop *AirdropSession) Freeze(acc common.Address, f bool) (*types.Transaction, error) {
	return _Airdrop.Contract.Freeze(&_Airdrop.TransactOpts, acc, f)
}

// Freeze is a paid mutator transaction binding the contract method 0xbf120ae5.
//
// Solidity: function freeze(address acc, bool f) returns()
func (_Airdrop *AirdropTransactorSession) Freeze(acc common.Address, f bool) (*types.Transaction, error) {
	return _Airdrop.Contract.Freeze(&_Airdrop.TransactOpts, acc, f)
}

// Init is a paid mutator transaction binding the contract method 0x9c26dd41.
//
// Solidity: function init(address apiCoin_, address appOwner_, string name_, string symbol_, string uri_, uint256 defaultWeight) returns()
func (_Airdrop *AirdropTransactor) Init(opts *bind.TransactOpts, apiCoin_ common.Address, appOwner_ common.Address, name_ string, symbol_ string, uri_ string, defaultWeight *big.Int) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "init", apiCoin_, appOwner_, name_, symbol_, uri_, defaultWeight)
}

// Init is a paid mutator transaction binding the contract method 0x9c26dd41.
//
// Solidity: function init(address apiCoin_, address appOwner_, string name_, string symbol_, string uri_, uint256 defaultWeight) returns()
func (_Airdrop *AirdropSession) Init(apiCoin_ common.Address, appOwner_ common.Address, name_ string, symbol_ string, uri_ string, defaultWeight *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.Init(&_Airdrop.TransactOpts, apiCoin_, appOwner_, name_, symbol_, uri_, defaultWeight)
}

// Init is a paid mutator transaction binding the contract method 0x9c26dd41.
//
// Solidity: function init(address apiCoin_, address appOwner_, string name_, string symbol_, string uri_, uint256 defaultWeight) returns()
func (_Airdrop *AirdropTransactorSession) Init(apiCoin_ common.Address, appOwner_ common.Address, name_ string, symbol_ string, uri_ string, defaultWeight *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.Init(&_Airdrop.TransactOpts, apiCoin_, appOwner_, name_, symbol_, uri_, defaultWeight)
}

// InitOwner is a paid mutator transaction binding the contract method 0x0d009297.
//
// Solidity: function initOwner(address owner_) returns()
func (_Airdrop *AirdropTransactor) InitOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "initOwner", owner_)
}

// InitOwner is a paid mutator transaction binding the contract method 0x0d009297.
//
// Solidity: function initOwner(address owner_) returns()
func (_Airdrop *AirdropSession) InitOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.InitOwner(&_Airdrop.TransactOpts, owner_)
}

// InitOwner is a paid mutator transaction binding the contract method 0x0d009297.
//
// Solidity: function initOwner(address owner_) returns()
func (_Airdrop *AirdropTransactorSession) InitOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.InitOwner(&_Airdrop.TransactOpts, owner_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Airdrop *AirdropTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Airdrop *AirdropSession) Pause() (*types.Transaction, error) {
	return _Airdrop.Contract.Pause(&_Airdrop.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Airdrop *AirdropTransactorSession) Pause() (*types.Transaction, error) {
	return _Airdrop.Contract.Pause(&_Airdrop.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0xfa89401a.
//
// Solidity: function refund(address account) returns()
func (_Airdrop *AirdropTransactor) Refund(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "refund", account)
}

// Refund is a paid mutator transaction binding the contract method 0xfa89401a.
//
// Solidity: function refund(address account) returns()
func (_Airdrop *AirdropSession) Refund(account common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.Refund(&_Airdrop.TransactOpts, account)
}

// Refund is a paid mutator transaction binding the contract method 0xfa89401a.
//
// Solidity: function refund(address account) returns()
func (_Airdrop *AirdropTransactorSession) Refund(account common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.Refund(&_Airdrop.TransactOpts, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Airdrop *AirdropTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Airdrop *AirdropSession) RenounceOwnership() (*types.Transaction, error) {
	return _Airdrop.Contract.RenounceOwnership(&_Airdrop.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Airdrop *AirdropTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Airdrop.Contract.RenounceOwnership(&_Airdrop.TransactOpts)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Airdrop *AirdropTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Airdrop *AirdropSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Airdrop.Contract.SafeBatchTransferFrom(&_Airdrop.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Airdrop *AirdropTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Airdrop.Contract.SafeBatchTransferFrom(&_Airdrop.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Airdrop *AirdropTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Airdrop *AirdropSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Airdrop.Contract.SafeTransferFrom(&_Airdrop.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Airdrop *AirdropTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Airdrop.Contract.SafeTransferFrom(&_Airdrop.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Airdrop *AirdropTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Airdrop *AirdropSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Airdrop.Contract.SetApprovalForAll(&_Airdrop.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Airdrop *AirdropTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Airdrop.Contract.SetApprovalForAll(&_Airdrop.TransactOpts, operator, approved)
}

// SetForceWithdrawDelay is a paid mutator transaction binding the contract method 0xcb5f869e.
//
// Solidity: function setForceWithdrawDelay(uint256 delay) returns()
func (_Airdrop *AirdropTransactor) SetForceWithdrawDelay(opts *bind.TransactOpts, delay *big.Int) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "setForceWithdrawDelay", delay)
}

// SetForceWithdrawDelay is a paid mutator transaction binding the contract method 0xcb5f869e.
//
// Solidity: function setForceWithdrawDelay(uint256 delay) returns()
func (_Airdrop *AirdropSession) SetForceWithdrawDelay(delay *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.SetForceWithdrawDelay(&_Airdrop.TransactOpts, delay)
}

// SetForceWithdrawDelay is a paid mutator transaction binding the contract method 0xcb5f869e.
//
// Solidity: function setForceWithdrawDelay(uint256 delay) returns()
func (_Airdrop *AirdropTransactorSession) SetForceWithdrawDelay(delay *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.SetForceWithdrawDelay(&_Airdrop.TransactOpts, delay)
}

// SetPendingSeconds is a paid mutator transaction binding the contract method 0xacb00a11.
//
// Solidity: function setPendingSeconds(uint256 seconds_) returns()
func (_Airdrop *AirdropTransactor) SetPendingSeconds(opts *bind.TransactOpts, seconds_ *big.Int) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "setPendingSeconds", seconds_)
}

// SetPendingSeconds is a paid mutator transaction binding the contract method 0xacb00a11.
//
// Solidity: function setPendingSeconds(uint256 seconds_) returns()
func (_Airdrop *AirdropSession) SetPendingSeconds(seconds_ *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.SetPendingSeconds(&_Airdrop.TransactOpts, seconds_)
}

// SetPendingSeconds is a paid mutator transaction binding the contract method 0xacb00a11.
//
// Solidity: function setPendingSeconds(uint256 seconds_) returns()
func (_Airdrop *AirdropTransactorSession) SetPendingSeconds(seconds_ *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.SetPendingSeconds(&_Airdrop.TransactOpts, seconds_)
}

// TakeProfit is a paid mutator transaction binding the contract method 0x0cec2f17.
//
// Solidity: function takeProfit(address to, uint256 amount) returns()
func (_Airdrop *AirdropTransactor) TakeProfit(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "takeProfit", to, amount)
}

// TakeProfit is a paid mutator transaction binding the contract method 0x0cec2f17.
//
// Solidity: function takeProfit(address to, uint256 amount) returns()
func (_Airdrop *AirdropSession) TakeProfit(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.TakeProfit(&_Airdrop.TransactOpts, to, amount)
}

// TakeProfit is a paid mutator transaction binding the contract method 0x0cec2f17.
//
// Solidity: function takeProfit(address to, uint256 amount) returns()
func (_Airdrop *AirdropTransactorSession) TakeProfit(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.TakeProfit(&_Airdrop.TransactOpts, to, amount)
}

// TokensReceived is a paid mutator transaction binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address from, address , uint256 amount, bytes , bytes ) returns()
func (_Airdrop *AirdropTransactor) TokensReceived(opts *bind.TransactOpts, arg0 common.Address, from common.Address, arg2 common.Address, amount *big.Int, arg4 []byte, arg5 []byte) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "tokensReceived", arg0, from, arg2, amount, arg4, arg5)
}

// TokensReceived is a paid mutator transaction binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address from, address , uint256 amount, bytes , bytes ) returns()
func (_Airdrop *AirdropSession) TokensReceived(arg0 common.Address, from common.Address, arg2 common.Address, amount *big.Int, arg4 []byte, arg5 []byte) (*types.Transaction, error) {
	return _Airdrop.Contract.TokensReceived(&_Airdrop.TransactOpts, arg0, from, arg2, amount, arg4, arg5)
}

// TokensReceived is a paid mutator transaction binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address from, address , uint256 amount, bytes , bytes ) returns()
func (_Airdrop *AirdropTransactorSession) TokensReceived(arg0 common.Address, from common.Address, arg2 common.Address, amount *big.Int, arg4 []byte, arg5 []byte) (*types.Transaction, error) {
	return _Airdrop.Contract.TokensReceived(&_Airdrop.TransactOpts, arg0, from, arg2, amount, arg4, arg5)
}

// TransferAppOwner is a paid mutator transaction binding the contract method 0xd3fb43a3.
//
// Solidity: function transferAppOwner(address to, address controller) returns()
func (_Airdrop *AirdropTransactor) TransferAppOwner(opts *bind.TransactOpts, to common.Address, controller common.Address) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "transferAppOwner", to, controller)
}

// TransferAppOwner is a paid mutator transaction binding the contract method 0xd3fb43a3.
//
// Solidity: function transferAppOwner(address to, address controller) returns()
func (_Airdrop *AirdropSession) TransferAppOwner(to common.Address, controller common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.TransferAppOwner(&_Airdrop.TransactOpts, to, controller)
}

// TransferAppOwner is a paid mutator transaction binding the contract method 0xd3fb43a3.
//
// Solidity: function transferAppOwner(address to, address controller) returns()
func (_Airdrop *AirdropTransactorSession) TransferAppOwner(to common.Address, controller common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.TransferAppOwner(&_Airdrop.TransactOpts, to, controller)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Airdrop *AirdropTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Airdrop *AirdropSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.TransferOwnership(&_Airdrop.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Airdrop *AirdropTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.TransferOwnership(&_Airdrop.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Airdrop *AirdropTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Airdrop *AirdropSession) Unpause() (*types.Transaction, error) {
	return _Airdrop.Contract.Unpause(&_Airdrop.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Airdrop *AirdropTransactorSession) Unpause() (*types.Transaction, error) {
	return _Airdrop.Contract.Unpause(&_Airdrop.TransactOpts)
}

// WithdrawRequest is a paid mutator transaction binding the contract method 0x380e687a.
//
// Solidity: function withdrawRequest() returns()
func (_Airdrop *AirdropTransactor) WithdrawRequest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "withdrawRequest")
}

// WithdrawRequest is a paid mutator transaction binding the contract method 0x380e687a.
//
// Solidity: function withdrawRequest() returns()
func (_Airdrop *AirdropSession) WithdrawRequest() (*types.Transaction, error) {
	return _Airdrop.Contract.WithdrawRequest(&_Airdrop.TransactOpts)
}

// WithdrawRequest is a paid mutator transaction binding the contract method 0x380e687a.
//
// Solidity: function withdrawRequest() returns()
func (_Airdrop *AirdropTransactorSession) WithdrawRequest() (*types.Transaction, error) {
	return _Airdrop.Contract.WithdrawRequest(&_Airdrop.TransactOpts)
}

// AirdropAppOwnerChangedIterator is returned from FilterAppOwnerChanged and is used to iterate over the raw logs and unpacked data for AppOwnerChanged events raised by the Airdrop contract.
type AirdropAppOwnerChangedIterator struct {
	Event *AirdropAppOwnerChanged // Event containing the contract specifics and raw log

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
func (it *AirdropAppOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropAppOwnerChanged)
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
		it.Event = new(AirdropAppOwnerChanged)
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
func (it *AirdropAppOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropAppOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropAppOwnerChanged represents a AppOwnerChanged event raised by the Airdrop contract.
type AirdropAppOwnerChanged struct {
	To  common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAppOwnerChanged is a free log retrieval operation binding the contract event 0xf90c9f5ec06d0b2a7c2a8bde8a2a10b1813a6eb0eaffe5de9f5bf28e5c092fd1.
//
// Solidity: event AppOwnerChanged(address indexed to)
func (_Airdrop *AirdropFilterer) FilterAppOwnerChanged(opts *bind.FilterOpts, to []common.Address) (*AirdropAppOwnerChangedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "AppOwnerChanged", toRule)
	if err != nil {
		return nil, err
	}
	return &AirdropAppOwnerChangedIterator{contract: _Airdrop.contract, event: "AppOwnerChanged", logs: logs, sub: sub}, nil
}

// WatchAppOwnerChanged is a free log subscription operation binding the contract event 0xf90c9f5ec06d0b2a7c2a8bde8a2a10b1813a6eb0eaffe5de9f5bf28e5c092fd1.
//
// Solidity: event AppOwnerChanged(address indexed to)
func (_Airdrop *AirdropFilterer) WatchAppOwnerChanged(opts *bind.WatchOpts, sink chan<- *AirdropAppOwnerChanged, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "AppOwnerChanged", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropAppOwnerChanged)
				if err := _Airdrop.contract.UnpackLog(event, "AppOwnerChanged", log); err != nil {
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
func (_Airdrop *AirdropFilterer) ParseAppOwnerChanged(log types.Log) (*AirdropAppOwnerChanged, error) {
	event := new(AirdropAppOwnerChanged)
	if err := _Airdrop.contract.UnpackLog(event, "AppOwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Airdrop contract.
type AirdropApprovalForAllIterator struct {
	Event *AirdropApprovalForAll // Event containing the contract specifics and raw log

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
func (it *AirdropApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropApprovalForAll)
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
		it.Event = new(AirdropApprovalForAll)
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
func (it *AirdropApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropApprovalForAll represents a ApprovalForAll event raised by the Airdrop contract.
type AirdropApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Airdrop *AirdropFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*AirdropApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &AirdropApprovalForAllIterator{contract: _Airdrop.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Airdrop *AirdropFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *AirdropApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropApprovalForAll)
				if err := _Airdrop.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Airdrop *AirdropFilterer) ParseApprovalForAll(log types.Log) (*AirdropApprovalForAll, error) {
	event := new(AirdropApprovalForAll)
	if err := _Airdrop.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropDropIterator is returned from FilterDrop and is used to iterate over the raw logs and unpacked data for Drop events raised by the Airdrop contract.
type AirdropDropIterator struct {
	Event *AirdropDrop // Event containing the contract specifics and raw log

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
func (it *AirdropDropIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropDrop)
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
		it.Event = new(AirdropDrop)
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
func (it *AirdropDropIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropDropIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropDrop represents a Drop event raised by the Airdrop contract.
type AirdropDrop struct {
	To     common.Address
	Amount *big.Int
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDrop is a free log retrieval operation binding the contract event 0x31adfe67d2df2c31426e3a5935de0d88d990beee41a012bad350132a880e4c54.
//
// Solidity: event Drop(address indexed to, uint256 amount, string reason)
func (_Airdrop *AirdropFilterer) FilterDrop(opts *bind.FilterOpts, to []common.Address) (*AirdropDropIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "Drop", toRule)
	if err != nil {
		return nil, err
	}
	return &AirdropDropIterator{contract: _Airdrop.contract, event: "Drop", logs: logs, sub: sub}, nil
}

// WatchDrop is a free log subscription operation binding the contract event 0x31adfe67d2df2c31426e3a5935de0d88d990beee41a012bad350132a880e4c54.
//
// Solidity: event Drop(address indexed to, uint256 amount, string reason)
func (_Airdrop *AirdropFilterer) WatchDrop(opts *bind.WatchOpts, sink chan<- *AirdropDrop, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "Drop", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropDrop)
				if err := _Airdrop.contract.UnpackLog(event, "Drop", log); err != nil {
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

// ParseDrop is a log parse operation binding the contract event 0x31adfe67d2df2c31426e3a5935de0d88d990beee41a012bad350132a880e4c54.
//
// Solidity: event Drop(address indexed to, uint256 amount, string reason)
func (_Airdrop *AirdropFilterer) ParseDrop(log types.Log) (*AirdropDrop, error) {
	event := new(AirdropDrop)
	if err := _Airdrop.contract.UnpackLog(event, "Drop", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropFrozenIterator is returned from FilterFrozen and is used to iterate over the raw logs and unpacked data for Frozen events raised by the Airdrop contract.
type AirdropFrozenIterator struct {
	Event *AirdropFrozen // Event containing the contract specifics and raw log

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
func (it *AirdropFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropFrozen)
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
		it.Event = new(AirdropFrozen)
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
func (it *AirdropFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropFrozen represents a Frozen event raised by the Airdrop contract.
type AirdropFrozen struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFrozen is a free log retrieval operation binding the contract event 0x8a5c4736a33c7b7f29a2c34ea9ff9608afc5718d56f6fd6dcbd2d3711a1a4913.
//
// Solidity: event Frozen(address indexed addr)
func (_Airdrop *AirdropFilterer) FilterFrozen(opts *bind.FilterOpts, addr []common.Address) (*AirdropFrozenIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "Frozen", addrRule)
	if err != nil {
		return nil, err
	}
	return &AirdropFrozenIterator{contract: _Airdrop.contract, event: "Frozen", logs: logs, sub: sub}, nil
}

// WatchFrozen is a free log subscription operation binding the contract event 0x8a5c4736a33c7b7f29a2c34ea9ff9608afc5718d56f6fd6dcbd2d3711a1a4913.
//
// Solidity: event Frozen(address indexed addr)
func (_Airdrop *AirdropFilterer) WatchFrozen(opts *bind.WatchOpts, sink chan<- *AirdropFrozen, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "Frozen", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropFrozen)
				if err := _Airdrop.contract.UnpackLog(event, "Frozen", log); err != nil {
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
func (_Airdrop *AirdropFilterer) ParseFrozen(log types.Log) (*AirdropFrozen, error) {
	event := new(AirdropFrozen)
	if err := _Airdrop.contract.UnpackLog(event, "Frozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Airdrop contract.
type AirdropOwnershipTransferredIterator struct {
	Event *AirdropOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AirdropOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropOwnershipTransferred)
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
		it.Event = new(AirdropOwnershipTransferred)
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
func (it *AirdropOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropOwnershipTransferred represents a OwnershipTransferred event raised by the Airdrop contract.
type AirdropOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Airdrop *AirdropFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AirdropOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AirdropOwnershipTransferredIterator{contract: _Airdrop.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Airdrop *AirdropFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AirdropOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropOwnershipTransferred)
				if err := _Airdrop.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Airdrop *AirdropFilterer) ParseOwnershipTransferred(log types.Log) (*AirdropOwnershipTransferred, error) {
	event := new(AirdropOwnershipTransferred)
	if err := _Airdrop.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Airdrop contract.
type AirdropPausedIterator struct {
	Event *AirdropPaused // Event containing the contract specifics and raw log

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
func (it *AirdropPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropPaused)
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
		it.Event = new(AirdropPaused)
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
func (it *AirdropPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropPaused represents a Paused event raised by the Airdrop contract.
type AirdropPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Airdrop *AirdropFilterer) FilterPaused(opts *bind.FilterOpts) (*AirdropPausedIterator, error) {

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AirdropPausedIterator{contract: _Airdrop.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Airdrop *AirdropFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AirdropPaused) (event.Subscription, error) {

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropPaused)
				if err := _Airdrop.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Airdrop *AirdropFilterer) ParsePaused(log types.Log) (*AirdropPaused, error) {
	event := new(AirdropPaused)
	if err := _Airdrop.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropResourceChangedIterator is returned from FilterResourceChanged and is used to iterate over the raw logs and unpacked data for ResourceChanged events raised by the Airdrop contract.
type AirdropResourceChangedIterator struct {
	Event *AirdropResourceChanged // Event containing the contract specifics and raw log

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
func (it *AirdropResourceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropResourceChanged)
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
		it.Event = new(AirdropResourceChanged)
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
func (it *AirdropResourceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropResourceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropResourceChanged represents a ResourceChanged event raised by the Airdrop contract.
type AirdropResourceChanged struct {
	Id     uint32
	Weight *big.Int
	Op     uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterResourceChanged is a free log retrieval operation binding the contract event 0x754de9a177534372342453272825c15351df3e7886699c2fece028d09379c400.
//
// Solidity: event ResourceChanged(uint32 indexed id, uint256 indexed weight, uint8 indexed op)
func (_Airdrop *AirdropFilterer) FilterResourceChanged(opts *bind.FilterOpts, id []uint32, weight []*big.Int, op []uint8) (*AirdropResourceChangedIterator, error) {

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

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "ResourceChanged", idRule, weightRule, opRule)
	if err != nil {
		return nil, err
	}
	return &AirdropResourceChangedIterator{contract: _Airdrop.contract, event: "ResourceChanged", logs: logs, sub: sub}, nil
}

// WatchResourceChanged is a free log subscription operation binding the contract event 0x754de9a177534372342453272825c15351df3e7886699c2fece028d09379c400.
//
// Solidity: event ResourceChanged(uint32 indexed id, uint256 indexed weight, uint8 indexed op)
func (_Airdrop *AirdropFilterer) WatchResourceChanged(opts *bind.WatchOpts, sink chan<- *AirdropResourceChanged, id []uint32, weight []*big.Int, op []uint8) (event.Subscription, error) {

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

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "ResourceChanged", idRule, weightRule, opRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropResourceChanged)
				if err := _Airdrop.contract.UnpackLog(event, "ResourceChanged", log); err != nil {
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
func (_Airdrop *AirdropFilterer) ParseResourceChanged(log types.Log) (*AirdropResourceChanged, error) {
	event := new(AirdropResourceChanged)
	if err := _Airdrop.contract.UnpackLog(event, "ResourceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropResourcePendingIterator is returned from FilterResourcePending and is used to iterate over the raw logs and unpacked data for ResourcePending events raised by the Airdrop contract.
type AirdropResourcePendingIterator struct {
	Event *AirdropResourcePending // Event containing the contract specifics and raw log

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
func (it *AirdropResourcePendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropResourcePending)
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
		it.Event = new(AirdropResourcePending)
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
func (it *AirdropResourcePendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropResourcePendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropResourcePending represents a ResourcePending event raised by the Airdrop contract.
type AirdropResourcePending struct {
	Id        uint32
	NewWeight *big.Int
	Op        uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterResourcePending is a free log retrieval operation binding the contract event 0xc2a05ca9bcad8635a6049c1240c2a1d2ed0c20c4f4b091ef84f7af523faf93d5.
//
// Solidity: event ResourcePending(uint32 indexed id, uint256 indexed newWeight, uint8 indexed op)
func (_Airdrop *AirdropFilterer) FilterResourcePending(opts *bind.FilterOpts, id []uint32, newWeight []*big.Int, op []uint8) (*AirdropResourcePendingIterator, error) {

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

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "ResourcePending", idRule, newWeightRule, opRule)
	if err != nil {
		return nil, err
	}
	return &AirdropResourcePendingIterator{contract: _Airdrop.contract, event: "ResourcePending", logs: logs, sub: sub}, nil
}

// WatchResourcePending is a free log subscription operation binding the contract event 0xc2a05ca9bcad8635a6049c1240c2a1d2ed0c20c4f4b091ef84f7af523faf93d5.
//
// Solidity: event ResourcePending(uint32 indexed id, uint256 indexed newWeight, uint8 indexed op)
func (_Airdrop *AirdropFilterer) WatchResourcePending(opts *bind.WatchOpts, sink chan<- *AirdropResourcePending, id []uint32, newWeight []*big.Int, op []uint8) (event.Subscription, error) {

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

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "ResourcePending", idRule, newWeightRule, opRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropResourcePending)
				if err := _Airdrop.contract.UnpackLog(event, "ResourcePending", log); err != nil {
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
func (_Airdrop *AirdropFilterer) ParseResourcePending(log types.Log) (*AirdropResourcePending, error) {
	event := new(AirdropResourcePending)
	if err := _Airdrop.contract.UnpackLog(event, "ResourcePending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropSpendIterator is returned from FilterSpend and is used to iterate over the raw logs and unpacked data for Spend events raised by the Airdrop contract.
type AirdropSpendIterator struct {
	Event *AirdropSpend // Event containing the contract specifics and raw log

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
func (it *AirdropSpendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropSpend)
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
		it.Event = new(AirdropSpend)
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
func (it *AirdropSpendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropSpendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropSpend represents a Spend event raised by the Airdrop contract.
type AirdropSpend struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSpend is a free log retrieval operation binding the contract event 0xaeba90871f7da8a443096c396877004da901c92fcab3ec900a99cecddb19ec4d.
//
// Solidity: event Spend(address indexed from, uint256 amount)
func (_Airdrop *AirdropFilterer) FilterSpend(opts *bind.FilterOpts, from []common.Address) (*AirdropSpendIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "Spend", fromRule)
	if err != nil {
		return nil, err
	}
	return &AirdropSpendIterator{contract: _Airdrop.contract, event: "Spend", logs: logs, sub: sub}, nil
}

// WatchSpend is a free log subscription operation binding the contract event 0xaeba90871f7da8a443096c396877004da901c92fcab3ec900a99cecddb19ec4d.
//
// Solidity: event Spend(address indexed from, uint256 amount)
func (_Airdrop *AirdropFilterer) WatchSpend(opts *bind.WatchOpts, sink chan<- *AirdropSpend, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "Spend", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropSpend)
				if err := _Airdrop.contract.UnpackLog(event, "Spend", log); err != nil {
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

// ParseSpend is a log parse operation binding the contract event 0xaeba90871f7da8a443096c396877004da901c92fcab3ec900a99cecddb19ec4d.
//
// Solidity: event Spend(address indexed from, uint256 amount)
func (_Airdrop *AirdropFilterer) ParseSpend(log types.Log) (*AirdropSpend, error) {
	event := new(AirdropSpend)
	if err := _Airdrop.contract.UnpackLog(event, "Spend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the Airdrop contract.
type AirdropTransferBatchIterator struct {
	Event *AirdropTransferBatch // Event containing the contract specifics and raw log

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
func (it *AirdropTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropTransferBatch)
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
		it.Event = new(AirdropTransferBatch)
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
func (it *AirdropTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropTransferBatch represents a TransferBatch event raised by the Airdrop contract.
type AirdropTransferBatch struct {
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
func (_Airdrop *AirdropFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*AirdropTransferBatchIterator, error) {

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

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AirdropTransferBatchIterator{contract: _Airdrop.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Airdrop *AirdropFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *AirdropTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropTransferBatch)
				if err := _Airdrop.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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
func (_Airdrop *AirdropFilterer) ParseTransferBatch(log types.Log) (*AirdropTransferBatch, error) {
	event := new(AirdropTransferBatch)
	if err := _Airdrop.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the Airdrop contract.
type AirdropTransferSingleIterator struct {
	Event *AirdropTransferSingle // Event containing the contract specifics and raw log

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
func (it *AirdropTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropTransferSingle)
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
		it.Event = new(AirdropTransferSingle)
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
func (it *AirdropTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropTransferSingle represents a TransferSingle event raised by the Airdrop contract.
type AirdropTransferSingle struct {
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
func (_Airdrop *AirdropFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*AirdropTransferSingleIterator, error) {

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

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AirdropTransferSingleIterator{contract: _Airdrop.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Airdrop *AirdropFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *AirdropTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropTransferSingle)
				if err := _Airdrop.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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
func (_Airdrop *AirdropFilterer) ParseTransferSingle(log types.Log) (*AirdropTransferSingle, error) {
	event := new(AirdropTransferSingle)
	if err := _Airdrop.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the Airdrop contract.
type AirdropURIIterator struct {
	Event *AirdropURI // Event containing the contract specifics and raw log

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
func (it *AirdropURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropURI)
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
		it.Event = new(AirdropURI)
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
func (it *AirdropURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropURI represents a URI event raised by the Airdrop contract.
type AirdropURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Airdrop *AirdropFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*AirdropURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &AirdropURIIterator{contract: _Airdrop.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Airdrop *AirdropFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *AirdropURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropURI)
				if err := _Airdrop.contract.UnpackLog(event, "URI", log); err != nil {
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
func (_Airdrop *AirdropFilterer) ParseURI(log types.Log) (*AirdropURI, error) {
	event := new(AirdropURI)
	if err := _Airdrop.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Airdrop contract.
type AirdropUnpausedIterator struct {
	Event *AirdropUnpaused // Event containing the contract specifics and raw log

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
func (it *AirdropUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropUnpaused)
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
		it.Event = new(AirdropUnpaused)
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
func (it *AirdropUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropUnpaused represents a Unpaused event raised by the Airdrop contract.
type AirdropUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Airdrop *AirdropFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AirdropUnpausedIterator, error) {

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AirdropUnpausedIterator{contract: _Airdrop.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Airdrop *AirdropFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AirdropUnpaused) (event.Subscription, error) {

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropUnpaused)
				if err := _Airdrop.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Airdrop *AirdropFilterer) ParseUnpaused(log types.Log) (*AirdropUnpaused, error) {
	event := new(AirdropUnpaused)
	if err := _Airdrop.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Airdrop contract.
type AirdropWithdrawIterator struct {
	Event *AirdropWithdraw // Event containing the contract specifics and raw log

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
func (it *AirdropWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropWithdraw)
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
		it.Event = new(AirdropWithdraw)
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
func (it *AirdropWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropWithdraw represents a Withdraw event raised by the Airdrop contract.
type AirdropWithdraw struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address account, uint256 amount)
func (_Airdrop *AirdropFilterer) FilterWithdraw(opts *bind.FilterOpts) (*AirdropWithdrawIterator, error) {

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &AirdropWithdrawIterator{contract: _Airdrop.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address account, uint256 amount)
func (_Airdrop *AirdropFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *AirdropWithdraw) (event.Subscription, error) {

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropWithdraw)
				if err := _Airdrop.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_Airdrop *AirdropFilterer) ParseWithdraw(log types.Log) (*AirdropWithdraw, error) {
	event := new(AirdropWithdraw)
	if err := _Airdrop.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
