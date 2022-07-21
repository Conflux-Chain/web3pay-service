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

// AppConfigConfigEntry is an auto generated low-level Go binding around an user-defined struct.
type AppConfigConfigEntry struct {
	ResourceId string
	Weight     uint32
	Index      uint32
}

// AppConfigConfigRequest is an auto generated low-level Go binding around an user-defined struct.
type AppConfigConfigRequest struct {
	Id         uint32
	ResourceId string
	Weight     uint32
	Op         uint8
}

// APPCoinMetaData contains all meta data concerning the APPCoin contract.
var APPCoinMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"AuthorizedOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"Frozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"weight\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"enumAppConfig.OP\",\"name\":\"op\",\"type\":\"uint8\"}],\"name\":\"ResourceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"RevokedOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"Sent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"apiCoin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"appOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"authorizeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"charge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"resourceId\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"weight\",\"type\":\"uint32\"},{\"internalType\":\"enumAppConfig.OP\",\"name\":\"op\",\"type\":\"uint8\"}],\"internalType\":\"structAppConfig.ConfigRequest\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"configResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"resourceId\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"weight\",\"type\":\"uint32\"},{\"internalType\":\"enumAppConfig.OP\",\"name\":\"op\",\"type\":\"uint8\"}],\"internalType\":\"structAppConfig.ConfigRequest[]\",\"name\":\"entries\",\"type\":\"tuple[]\"}],\"name\":\"configResourceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultOperators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceWithdrawAfterBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"f\",\"type\":\"bool\"}],\"name\":\"freeze\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"frozenMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"granularity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"indexArray\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"apiCoin_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"appOwner_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"initOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"isOperatorFor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"listResources\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"resourceId\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"weight\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAppConfig.ConfigEntry[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextConfigId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"operatorBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"operatorSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"resourceConfigures\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"resourceId\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"weight\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"revokeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"diff\",\"type\":\"uint256\"}],\"name\":\"setForceWithdrawAfterBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"takeProfit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"tokensReceived\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalCharged\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalTakenProfit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_APPCoin *APPCoinCaller) Allowance(opts *bind.CallOpts, holder common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "allowance", holder, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_APPCoin *APPCoinSession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _APPCoin.Contract.Allowance(&_APPCoin.CallOpts, holder, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_APPCoin *APPCoinCallerSession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _APPCoin.Contract.Allowance(&_APPCoin.CallOpts, holder, spender)
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

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenHolder) view returns(uint256)
func (_APPCoin *APPCoinCaller) BalanceOf(opts *bind.CallOpts, tokenHolder common.Address) (*big.Int, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "balanceOf", tokenHolder)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenHolder) view returns(uint256)
func (_APPCoin *APPCoinSession) BalanceOf(tokenHolder common.Address) (*big.Int, error) {
	return _APPCoin.Contract.BalanceOf(&_APPCoin.CallOpts, tokenHolder)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenHolder) view returns(uint256)
func (_APPCoin *APPCoinCallerSession) BalanceOf(tokenHolder common.Address) (*big.Int, error) {
	return _APPCoin.Contract.BalanceOf(&_APPCoin.CallOpts, tokenHolder)
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

// DefaultOperators is a free data retrieval call binding the contract method 0x06e48538.
//
// Solidity: function defaultOperators() view returns(address[])
func (_APPCoin *APPCoinCaller) DefaultOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "defaultOperators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// DefaultOperators is a free data retrieval call binding the contract method 0x06e48538.
//
// Solidity: function defaultOperators() view returns(address[])
func (_APPCoin *APPCoinSession) DefaultOperators() ([]common.Address, error) {
	return _APPCoin.Contract.DefaultOperators(&_APPCoin.CallOpts)
}

// DefaultOperators is a free data retrieval call binding the contract method 0x06e48538.
//
// Solidity: function defaultOperators() view returns(address[])
func (_APPCoin *APPCoinCallerSession) DefaultOperators() ([]common.Address, error) {
	return _APPCoin.Contract.DefaultOperators(&_APPCoin.CallOpts)
}

// ForceWithdrawAfterBlock is a free data retrieval call binding the contract method 0x82ad41a1.
//
// Solidity: function forceWithdrawAfterBlock() view returns(uint256)
func (_APPCoin *APPCoinCaller) ForceWithdrawAfterBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "forceWithdrawAfterBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ForceWithdrawAfterBlock is a free data retrieval call binding the contract method 0x82ad41a1.
//
// Solidity: function forceWithdrawAfterBlock() view returns(uint256)
func (_APPCoin *APPCoinSession) ForceWithdrawAfterBlock() (*big.Int, error) {
	return _APPCoin.Contract.ForceWithdrawAfterBlock(&_APPCoin.CallOpts)
}

// ForceWithdrawAfterBlock is a free data retrieval call binding the contract method 0x82ad41a1.
//
// Solidity: function forceWithdrawAfterBlock() view returns(uint256)
func (_APPCoin *APPCoinCallerSession) ForceWithdrawAfterBlock() (*big.Int, error) {
	return _APPCoin.Contract.ForceWithdrawAfterBlock(&_APPCoin.CallOpts)
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

// Granularity is a free data retrieval call binding the contract method 0x556f0dc7.
//
// Solidity: function granularity() view returns(uint256)
func (_APPCoin *APPCoinCaller) Granularity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "granularity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Granularity is a free data retrieval call binding the contract method 0x556f0dc7.
//
// Solidity: function granularity() view returns(uint256)
func (_APPCoin *APPCoinSession) Granularity() (*big.Int, error) {
	return _APPCoin.Contract.Granularity(&_APPCoin.CallOpts)
}

// Granularity is a free data retrieval call binding the contract method 0x556f0dc7.
//
// Solidity: function granularity() view returns(uint256)
func (_APPCoin *APPCoinCallerSession) Granularity() (*big.Int, error) {
	return _APPCoin.Contract.Granularity(&_APPCoin.CallOpts)
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

// IsOperatorFor is a free data retrieval call binding the contract method 0xd95b6371.
//
// Solidity: function isOperatorFor(address operator, address tokenHolder) view returns(bool)
func (_APPCoin *APPCoinCaller) IsOperatorFor(opts *bind.CallOpts, operator common.Address, tokenHolder common.Address) (bool, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "isOperatorFor", operator, tokenHolder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorFor is a free data retrieval call binding the contract method 0xd95b6371.
//
// Solidity: function isOperatorFor(address operator, address tokenHolder) view returns(bool)
func (_APPCoin *APPCoinSession) IsOperatorFor(operator common.Address, tokenHolder common.Address) (bool, error) {
	return _APPCoin.Contract.IsOperatorFor(&_APPCoin.CallOpts, operator, tokenHolder)
}

// IsOperatorFor is a free data retrieval call binding the contract method 0xd95b6371.
//
// Solidity: function isOperatorFor(address operator, address tokenHolder) view returns(bool)
func (_APPCoin *APPCoinCallerSession) IsOperatorFor(operator common.Address, tokenHolder common.Address) (bool, error) {
	return _APPCoin.Contract.IsOperatorFor(&_APPCoin.CallOpts, operator, tokenHolder)
}

// ListResources is a free data retrieval call binding the contract method 0x49459cd2.
//
// Solidity: function listResources(uint256 offset, uint256 limit) view returns((string,uint32,uint32)[], uint256 total)
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
// Solidity: function listResources(uint256 offset, uint256 limit) view returns((string,uint32,uint32)[], uint256 total)
func (_APPCoin *APPCoinSession) ListResources(offset *big.Int, limit *big.Int) ([]AppConfigConfigEntry, *big.Int, error) {
	return _APPCoin.Contract.ListResources(&_APPCoin.CallOpts, offset, limit)
}

// ListResources is a free data retrieval call binding the contract method 0x49459cd2.
//
// Solidity: function listResources(uint256 offset, uint256 limit) view returns((string,uint32,uint32)[], uint256 total)
func (_APPCoin *APPCoinCallerSession) ListResources(offset *big.Int, limit *big.Int) ([]AppConfigConfigEntry, *big.Int, error) {
	return _APPCoin.Contract.ListResources(&_APPCoin.CallOpts, offset, limit)
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

// ResourceConfigures is a free data retrieval call binding the contract method 0xc86b6ee2.
//
// Solidity: function resourceConfigures(uint32 ) view returns(string resourceId, uint32 weight, uint32 index)
func (_APPCoin *APPCoinCaller) ResourceConfigures(opts *bind.CallOpts, arg0 uint32) (struct {
	ResourceId string
	Weight     uint32
	Index      uint32
}, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "resourceConfigures", arg0)

	outstruct := new(struct {
		ResourceId string
		Weight     uint32
		Index      uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ResourceId = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Weight = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Index = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// ResourceConfigures is a free data retrieval call binding the contract method 0xc86b6ee2.
//
// Solidity: function resourceConfigures(uint32 ) view returns(string resourceId, uint32 weight, uint32 index)
func (_APPCoin *APPCoinSession) ResourceConfigures(arg0 uint32) (struct {
	ResourceId string
	Weight     uint32
	Index      uint32
}, error) {
	return _APPCoin.Contract.ResourceConfigures(&_APPCoin.CallOpts, arg0)
}

// ResourceConfigures is a free data retrieval call binding the contract method 0xc86b6ee2.
//
// Solidity: function resourceConfigures(uint32 ) view returns(string resourceId, uint32 weight, uint32 index)
func (_APPCoin *APPCoinCallerSession) ResourceConfigures(arg0 uint32) (struct {
	ResourceId string
	Weight     uint32
	Index      uint32
}, error) {
	return _APPCoin.Contract.ResourceConfigures(&_APPCoin.CallOpts, arg0)
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

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_APPCoin *APPCoinCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _APPCoin.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_APPCoin *APPCoinSession) TotalSupply() (*big.Int, error) {
	return _APPCoin.Contract.TotalSupply(&_APPCoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_APPCoin *APPCoinCallerSession) TotalSupply() (*big.Int, error) {
	return _APPCoin.Contract.TotalSupply(&_APPCoin.CallOpts)
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

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_APPCoin *APPCoinTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_APPCoin *APPCoinSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _APPCoin.Contract.Approve(&_APPCoin.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_APPCoin *APPCoinTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _APPCoin.Contract.Approve(&_APPCoin.TransactOpts, spender, value)
}

// AuthorizeOperator is a paid mutator transaction binding the contract method 0x959b8c3f.
//
// Solidity: function authorizeOperator(address operator) returns()
func (_APPCoin *APPCoinTransactor) AuthorizeOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "authorizeOperator", operator)
}

// AuthorizeOperator is a paid mutator transaction binding the contract method 0x959b8c3f.
//
// Solidity: function authorizeOperator(address operator) returns()
func (_APPCoin *APPCoinSession) AuthorizeOperator(operator common.Address) (*types.Transaction, error) {
	return _APPCoin.Contract.AuthorizeOperator(&_APPCoin.TransactOpts, operator)
}

// AuthorizeOperator is a paid mutator transaction binding the contract method 0x959b8c3f.
//
// Solidity: function authorizeOperator(address operator) returns()
func (_APPCoin *APPCoinTransactorSession) AuthorizeOperator(operator common.Address) (*types.Transaction, error) {
	return _APPCoin.Contract.AuthorizeOperator(&_APPCoin.TransactOpts, operator)
}

// Burn is a paid mutator transaction binding the contract method 0xfe9d9303.
//
// Solidity: function burn(uint256 amount, bytes ) returns()
func (_APPCoin *APPCoinTransactor) Burn(opts *bind.TransactOpts, amount *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "burn", amount, arg1)
}

// Burn is a paid mutator transaction binding the contract method 0xfe9d9303.
//
// Solidity: function burn(uint256 amount, bytes ) returns()
func (_APPCoin *APPCoinSession) Burn(amount *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _APPCoin.Contract.Burn(&_APPCoin.TransactOpts, amount, arg1)
}

// Burn is a paid mutator transaction binding the contract method 0xfe9d9303.
//
// Solidity: function burn(uint256 amount, bytes ) returns()
func (_APPCoin *APPCoinTransactorSession) Burn(amount *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _APPCoin.Contract.Burn(&_APPCoin.TransactOpts, amount, arg1)
}

// Charge is a paid mutator transaction binding the contract method 0xba77cae6.
//
// Solidity: function charge(address account, uint256 amount, bytes data) returns()
func (_APPCoin *APPCoinTransactor) Charge(opts *bind.TransactOpts, account common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "charge", account, amount, data)
}

// Charge is a paid mutator transaction binding the contract method 0xba77cae6.
//
// Solidity: function charge(address account, uint256 amount, bytes data) returns()
func (_APPCoin *APPCoinSession) Charge(account common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _APPCoin.Contract.Charge(&_APPCoin.TransactOpts, account, amount, data)
}

// Charge is a paid mutator transaction binding the contract method 0xba77cae6.
//
// Solidity: function charge(address account, uint256 amount, bytes data) returns()
func (_APPCoin *APPCoinTransactorSession) Charge(account common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _APPCoin.Contract.Charge(&_APPCoin.TransactOpts, account, amount, data)
}

// ConfigResource is a paid mutator transaction binding the contract method 0xabeef6f7.
//
// Solidity: function configResource((uint32,string,uint32,uint8) entry) returns()
func (_APPCoin *APPCoinTransactor) ConfigResource(opts *bind.TransactOpts, entry AppConfigConfigRequest) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "configResource", entry)
}

// ConfigResource is a paid mutator transaction binding the contract method 0xabeef6f7.
//
// Solidity: function configResource((uint32,string,uint32,uint8) entry) returns()
func (_APPCoin *APPCoinSession) ConfigResource(entry AppConfigConfigRequest) (*types.Transaction, error) {
	return _APPCoin.Contract.ConfigResource(&_APPCoin.TransactOpts, entry)
}

// ConfigResource is a paid mutator transaction binding the contract method 0xabeef6f7.
//
// Solidity: function configResource((uint32,string,uint32,uint8) entry) returns()
func (_APPCoin *APPCoinTransactorSession) ConfigResource(entry AppConfigConfigRequest) (*types.Transaction, error) {
	return _APPCoin.Contract.ConfigResource(&_APPCoin.TransactOpts, entry)
}

// ConfigResourceBatch is a paid mutator transaction binding the contract method 0x58c550d3.
//
// Solidity: function configResourceBatch((uint32,string,uint32,uint8)[] entries) returns()
func (_APPCoin *APPCoinTransactor) ConfigResourceBatch(opts *bind.TransactOpts, entries []AppConfigConfigRequest) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "configResourceBatch", entries)
}

// ConfigResourceBatch is a paid mutator transaction binding the contract method 0x58c550d3.
//
// Solidity: function configResourceBatch((uint32,string,uint32,uint8)[] entries) returns()
func (_APPCoin *APPCoinSession) ConfigResourceBatch(entries []AppConfigConfigRequest) (*types.Transaction, error) {
	return _APPCoin.Contract.ConfigResourceBatch(&_APPCoin.TransactOpts, entries)
}

// ConfigResourceBatch is a paid mutator transaction binding the contract method 0x58c550d3.
//
// Solidity: function configResourceBatch((uint32,string,uint32,uint8)[] entries) returns()
func (_APPCoin *APPCoinTransactorSession) ConfigResourceBatch(entries []AppConfigConfigRequest) (*types.Transaction, error) {
	return _APPCoin.Contract.ConfigResourceBatch(&_APPCoin.TransactOpts, entries)
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

// Init is a paid mutator transaction binding the contract method 0x4d91d7d9.
//
// Solidity: function init(address apiCoin_, address appOwner_, string name_, string symbol_) returns()
func (_APPCoin *APPCoinTransactor) Init(opts *bind.TransactOpts, apiCoin_ common.Address, appOwner_ common.Address, name_ string, symbol_ string) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "init", apiCoin_, appOwner_, name_, symbol_)
}

// Init is a paid mutator transaction binding the contract method 0x4d91d7d9.
//
// Solidity: function init(address apiCoin_, address appOwner_, string name_, string symbol_) returns()
func (_APPCoin *APPCoinSession) Init(apiCoin_ common.Address, appOwner_ common.Address, name_ string, symbol_ string) (*types.Transaction, error) {
	return _APPCoin.Contract.Init(&_APPCoin.TransactOpts, apiCoin_, appOwner_, name_, symbol_)
}

// Init is a paid mutator transaction binding the contract method 0x4d91d7d9.
//
// Solidity: function init(address apiCoin_, address appOwner_, string name_, string symbol_) returns()
func (_APPCoin *APPCoinTransactorSession) Init(apiCoin_ common.Address, appOwner_ common.Address, name_ string, symbol_ string) (*types.Transaction, error) {
	return _APPCoin.Contract.Init(&_APPCoin.TransactOpts, apiCoin_, appOwner_, name_, symbol_)
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

// OperatorBurn is a paid mutator transaction binding the contract method 0xfc673c4f.
//
// Solidity: function operatorBurn(address account, uint256 amount, bytes data, bytes operatorData) returns()
func (_APPCoin *APPCoinTransactor) OperatorBurn(opts *bind.TransactOpts, account common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "operatorBurn", account, amount, data, operatorData)
}

// OperatorBurn is a paid mutator transaction binding the contract method 0xfc673c4f.
//
// Solidity: function operatorBurn(address account, uint256 amount, bytes data, bytes operatorData) returns()
func (_APPCoin *APPCoinSession) OperatorBurn(account common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _APPCoin.Contract.OperatorBurn(&_APPCoin.TransactOpts, account, amount, data, operatorData)
}

// OperatorBurn is a paid mutator transaction binding the contract method 0xfc673c4f.
//
// Solidity: function operatorBurn(address account, uint256 amount, bytes data, bytes operatorData) returns()
func (_APPCoin *APPCoinTransactorSession) OperatorBurn(account common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _APPCoin.Contract.OperatorBurn(&_APPCoin.TransactOpts, account, amount, data, operatorData)
}

// OperatorSend is a paid mutator transaction binding the contract method 0x62ad1b83.
//
// Solidity: function operatorSend(address sender, address recipient, uint256 amount, bytes data, bytes operatorData) returns()
func (_APPCoin *APPCoinTransactor) OperatorSend(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "operatorSend", sender, recipient, amount, data, operatorData)
}

// OperatorSend is a paid mutator transaction binding the contract method 0x62ad1b83.
//
// Solidity: function operatorSend(address sender, address recipient, uint256 amount, bytes data, bytes operatorData) returns()
func (_APPCoin *APPCoinSession) OperatorSend(sender common.Address, recipient common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _APPCoin.Contract.OperatorSend(&_APPCoin.TransactOpts, sender, recipient, amount, data, operatorData)
}

// OperatorSend is a paid mutator transaction binding the contract method 0x62ad1b83.
//
// Solidity: function operatorSend(address sender, address recipient, uint256 amount, bytes data, bytes operatorData) returns()
func (_APPCoin *APPCoinTransactorSession) OperatorSend(sender common.Address, recipient common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _APPCoin.Contract.OperatorSend(&_APPCoin.TransactOpts, sender, recipient, amount, data, operatorData)
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

// RevokeOperator is a paid mutator transaction binding the contract method 0xfad8b32a.
//
// Solidity: function revokeOperator(address operator) returns()
func (_APPCoin *APPCoinTransactor) RevokeOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "revokeOperator", operator)
}

// RevokeOperator is a paid mutator transaction binding the contract method 0xfad8b32a.
//
// Solidity: function revokeOperator(address operator) returns()
func (_APPCoin *APPCoinSession) RevokeOperator(operator common.Address) (*types.Transaction, error) {
	return _APPCoin.Contract.RevokeOperator(&_APPCoin.TransactOpts, operator)
}

// RevokeOperator is a paid mutator transaction binding the contract method 0xfad8b32a.
//
// Solidity: function revokeOperator(address operator) returns()
func (_APPCoin *APPCoinTransactorSession) RevokeOperator(operator common.Address) (*types.Transaction, error) {
	return _APPCoin.Contract.RevokeOperator(&_APPCoin.TransactOpts, operator)
}

// Send is a paid mutator transaction binding the contract method 0x9bd9bbc6.
//
// Solidity: function send(address recipient, uint256 amount, bytes data) returns()
func (_APPCoin *APPCoinTransactor) Send(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "send", recipient, amount, data)
}

// Send is a paid mutator transaction binding the contract method 0x9bd9bbc6.
//
// Solidity: function send(address recipient, uint256 amount, bytes data) returns()
func (_APPCoin *APPCoinSession) Send(recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _APPCoin.Contract.Send(&_APPCoin.TransactOpts, recipient, amount, data)
}

// Send is a paid mutator transaction binding the contract method 0x9bd9bbc6.
//
// Solidity: function send(address recipient, uint256 amount, bytes data) returns()
func (_APPCoin *APPCoinTransactorSession) Send(recipient common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _APPCoin.Contract.Send(&_APPCoin.TransactOpts, recipient, amount, data)
}

// SetForceWithdrawAfterBlock is a paid mutator transaction binding the contract method 0x7ce7b295.
//
// Solidity: function setForceWithdrawAfterBlock(uint256 diff) returns()
func (_APPCoin *APPCoinTransactor) SetForceWithdrawAfterBlock(opts *bind.TransactOpts, diff *big.Int) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "setForceWithdrawAfterBlock", diff)
}

// SetForceWithdrawAfterBlock is a paid mutator transaction binding the contract method 0x7ce7b295.
//
// Solidity: function setForceWithdrawAfterBlock(uint256 diff) returns()
func (_APPCoin *APPCoinSession) SetForceWithdrawAfterBlock(diff *big.Int) (*types.Transaction, error) {
	return _APPCoin.Contract.SetForceWithdrawAfterBlock(&_APPCoin.TransactOpts, diff)
}

// SetForceWithdrawAfterBlock is a paid mutator transaction binding the contract method 0x7ce7b295.
//
// Solidity: function setForceWithdrawAfterBlock(uint256 diff) returns()
func (_APPCoin *APPCoinTransactorSession) SetForceWithdrawAfterBlock(diff *big.Int) (*types.Transaction, error) {
	return _APPCoin.Contract.SetForceWithdrawAfterBlock(&_APPCoin.TransactOpts, diff)
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

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_APPCoin *APPCoinTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_APPCoin *APPCoinSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _APPCoin.Contract.Transfer(&_APPCoin.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_APPCoin *APPCoinTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _APPCoin.Contract.Transfer(&_APPCoin.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address holder, address recipient, uint256 amount) returns(bool)
func (_APPCoin *APPCoinTransactor) TransferFrom(opts *bind.TransactOpts, holder common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _APPCoin.contract.Transact(opts, "transferFrom", holder, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address holder, address recipient, uint256 amount) returns(bool)
func (_APPCoin *APPCoinSession) TransferFrom(holder common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _APPCoin.Contract.TransferFrom(&_APPCoin.TransactOpts, holder, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address holder, address recipient, uint256 amount) returns(bool)
func (_APPCoin *APPCoinTransactorSession) TransferFrom(holder common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _APPCoin.Contract.TransferFrom(&_APPCoin.TransactOpts, holder, recipient, amount)
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

// APPCoinApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the APPCoin contract.
type APPCoinApprovalIterator struct {
	Event *APPCoinApproval // Event containing the contract specifics and raw log

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
func (it *APPCoinApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(APPCoinApproval)
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
		it.Event = new(APPCoinApproval)
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
func (it *APPCoinApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *APPCoinApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// APPCoinApproval represents a Approval event raised by the APPCoin contract.
type APPCoinApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_APPCoin *APPCoinFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*APPCoinApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _APPCoin.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &APPCoinApprovalIterator{contract: _APPCoin.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_APPCoin *APPCoinFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *APPCoinApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _APPCoin.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(APPCoinApproval)
				if err := _APPCoin.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_APPCoin *APPCoinFilterer) ParseApproval(log types.Log) (*APPCoinApproval, error) {
	event := new(APPCoinApproval)
	if err := _APPCoin.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// APPCoinAuthorizedOperatorIterator is returned from FilterAuthorizedOperator and is used to iterate over the raw logs and unpacked data for AuthorizedOperator events raised by the APPCoin contract.
type APPCoinAuthorizedOperatorIterator struct {
	Event *APPCoinAuthorizedOperator // Event containing the contract specifics and raw log

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
func (it *APPCoinAuthorizedOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(APPCoinAuthorizedOperator)
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
		it.Event = new(APPCoinAuthorizedOperator)
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
func (it *APPCoinAuthorizedOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *APPCoinAuthorizedOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// APPCoinAuthorizedOperator represents a AuthorizedOperator event raised by the APPCoin contract.
type APPCoinAuthorizedOperator struct {
	Operator    common.Address
	TokenHolder common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAuthorizedOperator is a free log retrieval operation binding the contract event 0xf4caeb2d6ca8932a215a353d0703c326ec2d81fc68170f320eb2ab49e9df61f9.
//
// Solidity: event AuthorizedOperator(address indexed operator, address indexed tokenHolder)
func (_APPCoin *APPCoinFilterer) FilterAuthorizedOperator(opts *bind.FilterOpts, operator []common.Address, tokenHolder []common.Address) (*APPCoinAuthorizedOperatorIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenHolderRule []interface{}
	for _, tokenHolderItem := range tokenHolder {
		tokenHolderRule = append(tokenHolderRule, tokenHolderItem)
	}

	logs, sub, err := _APPCoin.contract.FilterLogs(opts, "AuthorizedOperator", operatorRule, tokenHolderRule)
	if err != nil {
		return nil, err
	}
	return &APPCoinAuthorizedOperatorIterator{contract: _APPCoin.contract, event: "AuthorizedOperator", logs: logs, sub: sub}, nil
}

// WatchAuthorizedOperator is a free log subscription operation binding the contract event 0xf4caeb2d6ca8932a215a353d0703c326ec2d81fc68170f320eb2ab49e9df61f9.
//
// Solidity: event AuthorizedOperator(address indexed operator, address indexed tokenHolder)
func (_APPCoin *APPCoinFilterer) WatchAuthorizedOperator(opts *bind.WatchOpts, sink chan<- *APPCoinAuthorizedOperator, operator []common.Address, tokenHolder []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenHolderRule []interface{}
	for _, tokenHolderItem := range tokenHolder {
		tokenHolderRule = append(tokenHolderRule, tokenHolderItem)
	}

	logs, sub, err := _APPCoin.contract.WatchLogs(opts, "AuthorizedOperator", operatorRule, tokenHolderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(APPCoinAuthorizedOperator)
				if err := _APPCoin.contract.UnpackLog(event, "AuthorizedOperator", log); err != nil {
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

// ParseAuthorizedOperator is a log parse operation binding the contract event 0xf4caeb2d6ca8932a215a353d0703c326ec2d81fc68170f320eb2ab49e9df61f9.
//
// Solidity: event AuthorizedOperator(address indexed operator, address indexed tokenHolder)
func (_APPCoin *APPCoinFilterer) ParseAuthorizedOperator(log types.Log) (*APPCoinAuthorizedOperator, error) {
	event := new(APPCoinAuthorizedOperator)
	if err := _APPCoin.contract.UnpackLog(event, "AuthorizedOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// APPCoinBurnedIterator is returned from FilterBurned and is used to iterate over the raw logs and unpacked data for Burned events raised by the APPCoin contract.
type APPCoinBurnedIterator struct {
	Event *APPCoinBurned // Event containing the contract specifics and raw log

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
func (it *APPCoinBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(APPCoinBurned)
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
		it.Event = new(APPCoinBurned)
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
func (it *APPCoinBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *APPCoinBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// APPCoinBurned represents a Burned event raised by the APPCoin contract.
type APPCoinBurned struct {
	Operator     common.Address
	From         common.Address
	Amount       *big.Int
	Data         []byte
	OperatorData []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBurned is a free log retrieval operation binding the contract event 0xa78a9be3a7b862d26933ad85fb11d80ef66b8f972d7cbba06621d583943a4098.
//
// Solidity: event Burned(address indexed operator, address indexed from, uint256 amount, bytes data, bytes operatorData)
func (_APPCoin *APPCoinFilterer) FilterBurned(opts *bind.FilterOpts, operator []common.Address, from []common.Address) (*APPCoinBurnedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _APPCoin.contract.FilterLogs(opts, "Burned", operatorRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &APPCoinBurnedIterator{contract: _APPCoin.contract, event: "Burned", logs: logs, sub: sub}, nil
}

// WatchBurned is a free log subscription operation binding the contract event 0xa78a9be3a7b862d26933ad85fb11d80ef66b8f972d7cbba06621d583943a4098.
//
// Solidity: event Burned(address indexed operator, address indexed from, uint256 amount, bytes data, bytes operatorData)
func (_APPCoin *APPCoinFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *APPCoinBurned, operator []common.Address, from []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _APPCoin.contract.WatchLogs(opts, "Burned", operatorRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(APPCoinBurned)
				if err := _APPCoin.contract.UnpackLog(event, "Burned", log); err != nil {
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

// ParseBurned is a log parse operation binding the contract event 0xa78a9be3a7b862d26933ad85fb11d80ef66b8f972d7cbba06621d583943a4098.
//
// Solidity: event Burned(address indexed operator, address indexed from, uint256 amount, bytes data, bytes operatorData)
func (_APPCoin *APPCoinFilterer) ParseBurned(log types.Log) (*APPCoinBurned, error) {
	event := new(APPCoinBurned)
	if err := _APPCoin.contract.UnpackLog(event, "Burned", log); err != nil {
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

// APPCoinMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the APPCoin contract.
type APPCoinMintedIterator struct {
	Event *APPCoinMinted // Event containing the contract specifics and raw log

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
func (it *APPCoinMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(APPCoinMinted)
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
		it.Event = new(APPCoinMinted)
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
func (it *APPCoinMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *APPCoinMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// APPCoinMinted represents a Minted event raised by the APPCoin contract.
type APPCoinMinted struct {
	Operator     common.Address
	To           common.Address
	Amount       *big.Int
	Data         []byte
	OperatorData []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x2fe5be0146f74c5bce36c0b80911af6c7d86ff27e89d5cfa61fc681327954e5d.
//
// Solidity: event Minted(address indexed operator, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_APPCoin *APPCoinFilterer) FilterMinted(opts *bind.FilterOpts, operator []common.Address, to []common.Address) (*APPCoinMintedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _APPCoin.contract.FilterLogs(opts, "Minted", operatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return &APPCoinMintedIterator{contract: _APPCoin.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x2fe5be0146f74c5bce36c0b80911af6c7d86ff27e89d5cfa61fc681327954e5d.
//
// Solidity: event Minted(address indexed operator, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_APPCoin *APPCoinFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *APPCoinMinted, operator []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _APPCoin.contract.WatchLogs(opts, "Minted", operatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(APPCoinMinted)
				if err := _APPCoin.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0x2fe5be0146f74c5bce36c0b80911af6c7d86ff27e89d5cfa61fc681327954e5d.
//
// Solidity: event Minted(address indexed operator, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_APPCoin *APPCoinFilterer) ParseMinted(log types.Log) (*APPCoinMinted, error) {
	event := new(APPCoinMinted)
	if err := _APPCoin.contract.UnpackLog(event, "Minted", log); err != nil {
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
	Weight uint32
	Op     uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterResourceChanged is a free log retrieval operation binding the contract event 0xea282fdc831c820b26acb635031c48df7c1ad3ff8d8e0788cd8bbffa287609fb.
//
// Solidity: event ResourceChanged(uint32 indexed id, uint32 indexed weight, uint8 indexed op)
func (_APPCoin *APPCoinFilterer) FilterResourceChanged(opts *bind.FilterOpts, id []uint32, weight []uint32, op []uint8) (*APPCoinResourceChangedIterator, error) {

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

// WatchResourceChanged is a free log subscription operation binding the contract event 0xea282fdc831c820b26acb635031c48df7c1ad3ff8d8e0788cd8bbffa287609fb.
//
// Solidity: event ResourceChanged(uint32 indexed id, uint32 indexed weight, uint8 indexed op)
func (_APPCoin *APPCoinFilterer) WatchResourceChanged(opts *bind.WatchOpts, sink chan<- *APPCoinResourceChanged, id []uint32, weight []uint32, op []uint8) (event.Subscription, error) {

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

// ParseResourceChanged is a log parse operation binding the contract event 0xea282fdc831c820b26acb635031c48df7c1ad3ff8d8e0788cd8bbffa287609fb.
//
// Solidity: event ResourceChanged(uint32 indexed id, uint32 indexed weight, uint8 indexed op)
func (_APPCoin *APPCoinFilterer) ParseResourceChanged(log types.Log) (*APPCoinResourceChanged, error) {
	event := new(APPCoinResourceChanged)
	if err := _APPCoin.contract.UnpackLog(event, "ResourceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// APPCoinRevokedOperatorIterator is returned from FilterRevokedOperator and is used to iterate over the raw logs and unpacked data for RevokedOperator events raised by the APPCoin contract.
type APPCoinRevokedOperatorIterator struct {
	Event *APPCoinRevokedOperator // Event containing the contract specifics and raw log

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
func (it *APPCoinRevokedOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(APPCoinRevokedOperator)
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
		it.Event = new(APPCoinRevokedOperator)
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
func (it *APPCoinRevokedOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *APPCoinRevokedOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// APPCoinRevokedOperator represents a RevokedOperator event raised by the APPCoin contract.
type APPCoinRevokedOperator struct {
	Operator    common.Address
	TokenHolder common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRevokedOperator is a free log retrieval operation binding the contract event 0x50546e66e5f44d728365dc3908c63bc5cfeeab470722c1677e3073a6ac294aa1.
//
// Solidity: event RevokedOperator(address indexed operator, address indexed tokenHolder)
func (_APPCoin *APPCoinFilterer) FilterRevokedOperator(opts *bind.FilterOpts, operator []common.Address, tokenHolder []common.Address) (*APPCoinRevokedOperatorIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenHolderRule []interface{}
	for _, tokenHolderItem := range tokenHolder {
		tokenHolderRule = append(tokenHolderRule, tokenHolderItem)
	}

	logs, sub, err := _APPCoin.contract.FilterLogs(opts, "RevokedOperator", operatorRule, tokenHolderRule)
	if err != nil {
		return nil, err
	}
	return &APPCoinRevokedOperatorIterator{contract: _APPCoin.contract, event: "RevokedOperator", logs: logs, sub: sub}, nil
}

// WatchRevokedOperator is a free log subscription operation binding the contract event 0x50546e66e5f44d728365dc3908c63bc5cfeeab470722c1677e3073a6ac294aa1.
//
// Solidity: event RevokedOperator(address indexed operator, address indexed tokenHolder)
func (_APPCoin *APPCoinFilterer) WatchRevokedOperator(opts *bind.WatchOpts, sink chan<- *APPCoinRevokedOperator, operator []common.Address, tokenHolder []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenHolderRule []interface{}
	for _, tokenHolderItem := range tokenHolder {
		tokenHolderRule = append(tokenHolderRule, tokenHolderItem)
	}

	logs, sub, err := _APPCoin.contract.WatchLogs(opts, "RevokedOperator", operatorRule, tokenHolderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(APPCoinRevokedOperator)
				if err := _APPCoin.contract.UnpackLog(event, "RevokedOperator", log); err != nil {
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

// ParseRevokedOperator is a log parse operation binding the contract event 0x50546e66e5f44d728365dc3908c63bc5cfeeab470722c1677e3073a6ac294aa1.
//
// Solidity: event RevokedOperator(address indexed operator, address indexed tokenHolder)
func (_APPCoin *APPCoinFilterer) ParseRevokedOperator(log types.Log) (*APPCoinRevokedOperator, error) {
	event := new(APPCoinRevokedOperator)
	if err := _APPCoin.contract.UnpackLog(event, "RevokedOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// APPCoinSentIterator is returned from FilterSent and is used to iterate over the raw logs and unpacked data for Sent events raised by the APPCoin contract.
type APPCoinSentIterator struct {
	Event *APPCoinSent // Event containing the contract specifics and raw log

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
func (it *APPCoinSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(APPCoinSent)
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
		it.Event = new(APPCoinSent)
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
func (it *APPCoinSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *APPCoinSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// APPCoinSent represents a Sent event raised by the APPCoin contract.
type APPCoinSent struct {
	Operator     common.Address
	From         common.Address
	To           common.Address
	Amount       *big.Int
	Data         []byte
	OperatorData []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSent is a free log retrieval operation binding the contract event 0x06b541ddaa720db2b10a4d0cdac39b8d360425fc073085fac19bc82614677987.
//
// Solidity: event Sent(address indexed operator, address indexed from, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_APPCoin *APPCoinFilterer) FilterSent(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*APPCoinSentIterator, error) {

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

	logs, sub, err := _APPCoin.contract.FilterLogs(opts, "Sent", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &APPCoinSentIterator{contract: _APPCoin.contract, event: "Sent", logs: logs, sub: sub}, nil
}

// WatchSent is a free log subscription operation binding the contract event 0x06b541ddaa720db2b10a4d0cdac39b8d360425fc073085fac19bc82614677987.
//
// Solidity: event Sent(address indexed operator, address indexed from, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_APPCoin *APPCoinFilterer) WatchSent(opts *bind.WatchOpts, sink chan<- *APPCoinSent, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _APPCoin.contract.WatchLogs(opts, "Sent", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(APPCoinSent)
				if err := _APPCoin.contract.UnpackLog(event, "Sent", log); err != nil {
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

// ParseSent is a log parse operation binding the contract event 0x06b541ddaa720db2b10a4d0cdac39b8d360425fc073085fac19bc82614677987.
//
// Solidity: event Sent(address indexed operator, address indexed from, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_APPCoin *APPCoinFilterer) ParseSent(log types.Log) (*APPCoinSent, error) {
	event := new(APPCoinSent)
	if err := _APPCoin.contract.UnpackLog(event, "Sent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// APPCoinTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the APPCoin contract.
type APPCoinTransferIterator struct {
	Event *APPCoinTransfer // Event containing the contract specifics and raw log

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
func (it *APPCoinTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(APPCoinTransfer)
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
		it.Event = new(APPCoinTransfer)
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
func (it *APPCoinTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *APPCoinTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// APPCoinTransfer represents a Transfer event raised by the APPCoin contract.
type APPCoinTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_APPCoin *APPCoinFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*APPCoinTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _APPCoin.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &APPCoinTransferIterator{contract: _APPCoin.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_APPCoin *APPCoinFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *APPCoinTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _APPCoin.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(APPCoinTransfer)
				if err := _APPCoin.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_APPCoin *APPCoinFilterer) ParseTransfer(log types.Log) (*APPCoinTransfer, error) {
	event := new(APPCoinTransfer)
	if err := _APPCoin.contract.UnpackLog(event, "Transfer", log); err != nil {
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
