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

// IAPPCoinMetaData contains all meta data concerning the IAPPCoin contract.
var IAPPCoinMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"apiCoin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IAPPCoinABI is the input ABI used to generate the binding from.
// Deprecated: Use IAPPCoinMetaData.ABI instead.
var IAPPCoinABI = IAPPCoinMetaData.ABI

// IAPPCoin is an auto generated Go binding around an Ethereum contract.
type IAPPCoin struct {
	IAPPCoinCaller     // Read-only binding to the contract
	IAPPCoinTransactor // Write-only binding to the contract
	IAPPCoinFilterer   // Log filterer for contract events
}

// IAPPCoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAPPCoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAPPCoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAPPCoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAPPCoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAPPCoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAPPCoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAPPCoinSession struct {
	Contract     *IAPPCoin         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAPPCoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAPPCoinCallerSession struct {
	Contract *IAPPCoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IAPPCoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAPPCoinTransactorSession struct {
	Contract     *IAPPCoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IAPPCoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAPPCoinRaw struct {
	Contract *IAPPCoin // Generic contract binding to access the raw methods on
}

// IAPPCoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAPPCoinCallerRaw struct {
	Contract *IAPPCoinCaller // Generic read-only contract binding to access the raw methods on
}

// IAPPCoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAPPCoinTransactorRaw struct {
	Contract *IAPPCoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAPPCoin creates a new instance of IAPPCoin, bound to a specific deployed contract.
func NewIAPPCoin(address common.Address, backend bind.ContractBackend) (*IAPPCoin, error) {
	contract, err := bindIAPPCoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAPPCoin{IAPPCoinCaller: IAPPCoinCaller{contract: contract}, IAPPCoinTransactor: IAPPCoinTransactor{contract: contract}, IAPPCoinFilterer: IAPPCoinFilterer{contract: contract}}, nil
}

// NewIAPPCoinCaller creates a new read-only instance of IAPPCoin, bound to a specific deployed contract.
func NewIAPPCoinCaller(address common.Address, caller bind.ContractCaller) (*IAPPCoinCaller, error) {
	contract, err := bindIAPPCoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAPPCoinCaller{contract: contract}, nil
}

// NewIAPPCoinTransactor creates a new write-only instance of IAPPCoin, bound to a specific deployed contract.
func NewIAPPCoinTransactor(address common.Address, transactor bind.ContractTransactor) (*IAPPCoinTransactor, error) {
	contract, err := bindIAPPCoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAPPCoinTransactor{contract: contract}, nil
}

// NewIAPPCoinFilterer creates a new log filterer instance of IAPPCoin, bound to a specific deployed contract.
func NewIAPPCoinFilterer(address common.Address, filterer bind.ContractFilterer) (*IAPPCoinFilterer, error) {
	contract, err := bindIAPPCoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAPPCoinFilterer{contract: contract}, nil
}

// bindIAPPCoin binds a generic wrapper to an already deployed contract.
func bindIAPPCoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAPPCoinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAPPCoin *IAPPCoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAPPCoin.Contract.IAPPCoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAPPCoin *IAPPCoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAPPCoin.Contract.IAPPCoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAPPCoin *IAPPCoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAPPCoin.Contract.IAPPCoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAPPCoin *IAPPCoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAPPCoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAPPCoin *IAPPCoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAPPCoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAPPCoin *IAPPCoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAPPCoin.Contract.contract.Transact(opts, method, params...)
}

// ApiCoin is a paid mutator transaction binding the contract method 0x4ebf365a.
//
// Solidity: function apiCoin() returns(address)
func (_IAPPCoin *IAPPCoinTransactor) ApiCoin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAPPCoin.contract.Transact(opts, "apiCoin")
}

// ApiCoin is a paid mutator transaction binding the contract method 0x4ebf365a.
//
// Solidity: function apiCoin() returns(address)
func (_IAPPCoin *IAPPCoinSession) ApiCoin() (*types.Transaction, error) {
	return _IAPPCoin.Contract.ApiCoin(&_IAPPCoin.TransactOpts)
}

// ApiCoin is a paid mutator transaction binding the contract method 0x4ebf365a.
//
// Solidity: function apiCoin() returns(address)
func (_IAPPCoin *IAPPCoinTransactorSession) ApiCoin() (*types.Transaction, error) {
	return _IAPPCoin.Contract.ApiCoin(&_IAPPCoin.TransactOpts)
}
