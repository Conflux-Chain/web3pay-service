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

// AppRegistryAppInfo is an auto generated low-level Go binding around an user-defined struct.
type AppRegistryAppInfo struct {
	Addr       common.Address
	CreateTime *big.Int
}

// AppRegistryMetaData contains all meta data concerning the AppRegistry contract.
var AppRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"apiWeightToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vipCoin\",\"type\":\"address\"}],\"name\":\"Created\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"Removed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CREATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"addUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"appFactory\",\"outputs\":[{\"internalType\":\"contractAppFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"link\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"enumIApp.PaymentType\",\"name\":\"paymentType_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"deferTimeSecs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"defaultApiWeight\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creatorRoleDisabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createTime\",\"type\":\"uint256\"}],\"internalType\":\"structAppRegistry.AppInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createTime\",\"type\":\"uint256\"}],\"internalType\":\"structAppRegistry.AppInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExchanger\",\"outputs\":[{\"internalType\":\"contractISwapExchange\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAppFactory\",\"name\":\"appFactory_\",\"type\":\"address\"},{\"internalType\":\"contractISwapExchange\",\"name\":\"exchanger_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"list\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createTime\",\"type\":\"uint256\"}],\"internalType\":\"structAppRegistry.AppInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"listByOwner\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createTime\",\"type\":\"uint256\"}],\"internalType\":\"structAppRegistry.AppInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"listByUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createTime\",\"type\":\"uint256\"}],\"internalType\":\"structAppRegistry.AppInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"}],\"name\":\"setCreatorRoleDisabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISwapExchange\",\"name\":\"exchanger_\",\"type\":\"address\"}],\"name\":\"setExchanger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AppRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use AppRegistryMetaData.ABI instead.
var AppRegistryABI = AppRegistryMetaData.ABI

// AppRegistry is an auto generated Go binding around an Ethereum contract.
type AppRegistry struct {
	AppRegistryCaller     // Read-only binding to the contract
	AppRegistryTransactor // Write-only binding to the contract
	AppRegistryFilterer   // Log filterer for contract events
}

// AppRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AppRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AppRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AppRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AppRegistrySession struct {
	Contract     *AppRegistry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AppRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AppRegistryCallerSession struct {
	Contract *AppRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AppRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AppRegistryTransactorSession struct {
	Contract     *AppRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AppRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AppRegistryRaw struct {
	Contract *AppRegistry // Generic contract binding to access the raw methods on
}

// AppRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AppRegistryCallerRaw struct {
	Contract *AppRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// AppRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AppRegistryTransactorRaw struct {
	Contract *AppRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAppRegistry creates a new instance of AppRegistry, bound to a specific deployed contract.
func NewAppRegistry(address common.Address, backend bind.ContractBackend) (*AppRegistry, error) {
	contract, err := bindAppRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AppRegistry{AppRegistryCaller: AppRegistryCaller{contract: contract}, AppRegistryTransactor: AppRegistryTransactor{contract: contract}, AppRegistryFilterer: AppRegistryFilterer{contract: contract}}, nil
}

// NewAppRegistryCaller creates a new read-only instance of AppRegistry, bound to a specific deployed contract.
func NewAppRegistryCaller(address common.Address, caller bind.ContractCaller) (*AppRegistryCaller, error) {
	contract, err := bindAppRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AppRegistryCaller{contract: contract}, nil
}

// NewAppRegistryTransactor creates a new write-only instance of AppRegistry, bound to a specific deployed contract.
func NewAppRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*AppRegistryTransactor, error) {
	contract, err := bindAppRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AppRegistryTransactor{contract: contract}, nil
}

// NewAppRegistryFilterer creates a new log filterer instance of AppRegistry, bound to a specific deployed contract.
func NewAppRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*AppRegistryFilterer, error) {
	contract, err := bindAppRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AppRegistryFilterer{contract: contract}, nil
}

// bindAppRegistry binds a generic wrapper to an already deployed contract.
func bindAppRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AppRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AppRegistry *AppRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AppRegistry.Contract.AppRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AppRegistry *AppRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AppRegistry.Contract.AppRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AppRegistry *AppRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AppRegistry.Contract.AppRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AppRegistry *AppRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AppRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AppRegistry *AppRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AppRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AppRegistry *AppRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AppRegistry.Contract.contract.Transact(opts, method, params...)
}

// CREATORROLE is a free data retrieval call binding the contract method 0x8aeda25a.
//
// Solidity: function CREATOR_ROLE() view returns(bytes32)
func (_AppRegistry *AppRegistryCaller) CREATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AppRegistry.contract.Call(opts, &out, "CREATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CREATORROLE is a free data retrieval call binding the contract method 0x8aeda25a.
//
// Solidity: function CREATOR_ROLE() view returns(bytes32)
func (_AppRegistry *AppRegistrySession) CREATORROLE() ([32]byte, error) {
	return _AppRegistry.Contract.CREATORROLE(&_AppRegistry.CallOpts)
}

// CREATORROLE is a free data retrieval call binding the contract method 0x8aeda25a.
//
// Solidity: function CREATOR_ROLE() view returns(bytes32)
func (_AppRegistry *AppRegistryCallerSession) CREATORROLE() ([32]byte, error) {
	return _AppRegistry.Contract.CREATORROLE(&_AppRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AppRegistry *AppRegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AppRegistry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AppRegistry *AppRegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AppRegistry.Contract.DEFAULTADMINROLE(&_AppRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AppRegistry *AppRegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AppRegistry.Contract.DEFAULTADMINROLE(&_AppRegistry.CallOpts)
}

// AppFactory is a free data retrieval call binding the contract method 0x09817fa1.
//
// Solidity: function appFactory() view returns(address)
func (_AppRegistry *AppRegistryCaller) AppFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AppRegistry.contract.Call(opts, &out, "appFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AppFactory is a free data retrieval call binding the contract method 0x09817fa1.
//
// Solidity: function appFactory() view returns(address)
func (_AppRegistry *AppRegistrySession) AppFactory() (common.Address, error) {
	return _AppRegistry.Contract.AppFactory(&_AppRegistry.CallOpts)
}

// AppFactory is a free data retrieval call binding the contract method 0x09817fa1.
//
// Solidity: function appFactory() view returns(address)
func (_AppRegistry *AppRegistryCallerSession) AppFactory() (common.Address, error) {
	return _AppRegistry.Contract.AppFactory(&_AppRegistry.CallOpts)
}

// CreatorRoleDisabled is a free data retrieval call binding the contract method 0x32277071.
//
// Solidity: function creatorRoleDisabled() view returns(bool)
func (_AppRegistry *AppRegistryCaller) CreatorRoleDisabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AppRegistry.contract.Call(opts, &out, "creatorRoleDisabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CreatorRoleDisabled is a free data retrieval call binding the contract method 0x32277071.
//
// Solidity: function creatorRoleDisabled() view returns(bool)
func (_AppRegistry *AppRegistrySession) CreatorRoleDisabled() (bool, error) {
	return _AppRegistry.Contract.CreatorRoleDisabled(&_AppRegistry.CallOpts)
}

// CreatorRoleDisabled is a free data retrieval call binding the contract method 0x32277071.
//
// Solidity: function creatorRoleDisabled() view returns(bool)
func (_AppRegistry *AppRegistryCallerSession) CreatorRoleDisabled() (bool, error) {
	return _AppRegistry.Contract.CreatorRoleDisabled(&_AppRegistry.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address app) view returns((address,uint256))
func (_AppRegistry *AppRegistryCaller) Get(opts *bind.CallOpts, app common.Address) (AppRegistryAppInfo, error) {
	var out []interface{}
	err := _AppRegistry.contract.Call(opts, &out, "get", app)

	if err != nil {
		return *new(AppRegistryAppInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(AppRegistryAppInfo)).(*AppRegistryAppInfo)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address app) view returns((address,uint256))
func (_AppRegistry *AppRegistrySession) Get(app common.Address) (AppRegistryAppInfo, error) {
	return _AppRegistry.Contract.Get(&_AppRegistry.CallOpts, app)
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address app) view returns((address,uint256))
func (_AppRegistry *AppRegistryCallerSession) Get(app common.Address) (AppRegistryAppInfo, error) {
	return _AppRegistry.Contract.Get(&_AppRegistry.CallOpts, app)
}

// Get0 is a free data retrieval call binding the contract method 0xd81e8423.
//
// Solidity: function get(address owner, address app) view returns((address,uint256))
func (_AppRegistry *AppRegistryCaller) Get0(opts *bind.CallOpts, owner common.Address, app common.Address) (AppRegistryAppInfo, error) {
	var out []interface{}
	err := _AppRegistry.contract.Call(opts, &out, "get0", owner, app)

	if err != nil {
		return *new(AppRegistryAppInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(AppRegistryAppInfo)).(*AppRegistryAppInfo)

	return out0, err

}

// Get0 is a free data retrieval call binding the contract method 0xd81e8423.
//
// Solidity: function get(address owner, address app) view returns((address,uint256))
func (_AppRegistry *AppRegistrySession) Get0(owner common.Address, app common.Address) (AppRegistryAppInfo, error) {
	return _AppRegistry.Contract.Get0(&_AppRegistry.CallOpts, owner, app)
}

// Get0 is a free data retrieval call binding the contract method 0xd81e8423.
//
// Solidity: function get(address owner, address app) view returns((address,uint256))
func (_AppRegistry *AppRegistryCallerSession) Get0(owner common.Address, app common.Address) (AppRegistryAppInfo, error) {
	return _AppRegistry.Contract.Get0(&_AppRegistry.CallOpts, owner, app)
}

// GetExchanger is a free data retrieval call binding the contract method 0x4b1acf39.
//
// Solidity: function getExchanger() view returns(address)
func (_AppRegistry *AppRegistryCaller) GetExchanger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AppRegistry.contract.Call(opts, &out, "getExchanger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetExchanger is a free data retrieval call binding the contract method 0x4b1acf39.
//
// Solidity: function getExchanger() view returns(address)
func (_AppRegistry *AppRegistrySession) GetExchanger() (common.Address, error) {
	return _AppRegistry.Contract.GetExchanger(&_AppRegistry.CallOpts)
}

// GetExchanger is a free data retrieval call binding the contract method 0x4b1acf39.
//
// Solidity: function getExchanger() view returns(address)
func (_AppRegistry *AppRegistryCallerSession) GetExchanger() (common.Address, error) {
	return _AppRegistry.Contract.GetExchanger(&_AppRegistry.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AppRegistry *AppRegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AppRegistry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AppRegistry *AppRegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AppRegistry.Contract.GetRoleAdmin(&_AppRegistry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AppRegistry *AppRegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AppRegistry.Contract.GetRoleAdmin(&_AppRegistry.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AppRegistry *AppRegistryCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AppRegistry.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AppRegistry *AppRegistrySession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _AppRegistry.Contract.GetRoleMember(&_AppRegistry.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AppRegistry *AppRegistryCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _AppRegistry.Contract.GetRoleMember(&_AppRegistry.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AppRegistry *AppRegistryCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _AppRegistry.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AppRegistry *AppRegistrySession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AppRegistry.Contract.GetRoleMemberCount(&_AppRegistry.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AppRegistry *AppRegistryCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AppRegistry.Contract.GetRoleMemberCount(&_AppRegistry.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AppRegistry *AppRegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AppRegistry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AppRegistry *AppRegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AppRegistry.Contract.HasRole(&_AppRegistry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AppRegistry *AppRegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AppRegistry.Contract.HasRole(&_AppRegistry.CallOpts, role, account)
}

// List is a free data retrieval call binding the contract method 0x50fd7367.
//
// Solidity: function list(uint256 offset, uint256 limit) view returns(uint256, (address,uint256)[])
func (_AppRegistry *AppRegistryCaller) List(opts *bind.CallOpts, offset *big.Int, limit *big.Int) (*big.Int, []AppRegistryAppInfo, error) {
	var out []interface{}
	err := _AppRegistry.contract.Call(opts, &out, "list", offset, limit)

	if err != nil {
		return *new(*big.Int), *new([]AppRegistryAppInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]AppRegistryAppInfo)).(*[]AppRegistryAppInfo)

	return out0, out1, err

}

// List is a free data retrieval call binding the contract method 0x50fd7367.
//
// Solidity: function list(uint256 offset, uint256 limit) view returns(uint256, (address,uint256)[])
func (_AppRegistry *AppRegistrySession) List(offset *big.Int, limit *big.Int) (*big.Int, []AppRegistryAppInfo, error) {
	return _AppRegistry.Contract.List(&_AppRegistry.CallOpts, offset, limit)
}

// List is a free data retrieval call binding the contract method 0x50fd7367.
//
// Solidity: function list(uint256 offset, uint256 limit) view returns(uint256, (address,uint256)[])
func (_AppRegistry *AppRegistryCallerSession) List(offset *big.Int, limit *big.Int) (*big.Int, []AppRegistryAppInfo, error) {
	return _AppRegistry.Contract.List(&_AppRegistry.CallOpts, offset, limit)
}

// ListByOwner is a free data retrieval call binding the contract method 0x7ed7943c.
//
// Solidity: function listByOwner(address owner, uint256 offset, uint256 limit) view returns(uint256, (address,uint256)[])
func (_AppRegistry *AppRegistryCaller) ListByOwner(opts *bind.CallOpts, owner common.Address, offset *big.Int, limit *big.Int) (*big.Int, []AppRegistryAppInfo, error) {
	var out []interface{}
	err := _AppRegistry.contract.Call(opts, &out, "listByOwner", owner, offset, limit)

	if err != nil {
		return *new(*big.Int), *new([]AppRegistryAppInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]AppRegistryAppInfo)).(*[]AppRegistryAppInfo)

	return out0, out1, err

}

// ListByOwner is a free data retrieval call binding the contract method 0x7ed7943c.
//
// Solidity: function listByOwner(address owner, uint256 offset, uint256 limit) view returns(uint256, (address,uint256)[])
func (_AppRegistry *AppRegistrySession) ListByOwner(owner common.Address, offset *big.Int, limit *big.Int) (*big.Int, []AppRegistryAppInfo, error) {
	return _AppRegistry.Contract.ListByOwner(&_AppRegistry.CallOpts, owner, offset, limit)
}

// ListByOwner is a free data retrieval call binding the contract method 0x7ed7943c.
//
// Solidity: function listByOwner(address owner, uint256 offset, uint256 limit) view returns(uint256, (address,uint256)[])
func (_AppRegistry *AppRegistryCallerSession) ListByOwner(owner common.Address, offset *big.Int, limit *big.Int) (*big.Int, []AppRegistryAppInfo, error) {
	return _AppRegistry.Contract.ListByOwner(&_AppRegistry.CallOpts, owner, offset, limit)
}

// ListByUser is a free data retrieval call binding the contract method 0x7888ad21.
//
// Solidity: function listByUser(address user, uint256 offset, uint256 limit) view returns(uint256, (address,uint256)[])
func (_AppRegistry *AppRegistryCaller) ListByUser(opts *bind.CallOpts, user common.Address, offset *big.Int, limit *big.Int) (*big.Int, []AppRegistryAppInfo, error) {
	var out []interface{}
	err := _AppRegistry.contract.Call(opts, &out, "listByUser", user, offset, limit)

	if err != nil {
		return *new(*big.Int), *new([]AppRegistryAppInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]AppRegistryAppInfo)).(*[]AppRegistryAppInfo)

	return out0, out1, err

}

// ListByUser is a free data retrieval call binding the contract method 0x7888ad21.
//
// Solidity: function listByUser(address user, uint256 offset, uint256 limit) view returns(uint256, (address,uint256)[])
func (_AppRegistry *AppRegistrySession) ListByUser(user common.Address, offset *big.Int, limit *big.Int) (*big.Int, []AppRegistryAppInfo, error) {
	return _AppRegistry.Contract.ListByUser(&_AppRegistry.CallOpts, user, offset, limit)
}

// ListByUser is a free data retrieval call binding the contract method 0x7888ad21.
//
// Solidity: function listByUser(address user, uint256 offset, uint256 limit) view returns(uint256, (address,uint256)[])
func (_AppRegistry *AppRegistryCallerSession) ListByUser(user common.Address, offset *big.Int, limit *big.Int) (*big.Int, []AppRegistryAppInfo, error) {
	return _AppRegistry.Contract.ListByUser(&_AppRegistry.CallOpts, user, offset, limit)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AppRegistry *AppRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AppRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AppRegistry *AppRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AppRegistry.Contract.SupportsInterface(&_AppRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AppRegistry *AppRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AppRegistry.Contract.SupportsInterface(&_AppRegistry.CallOpts, interfaceId)
}

// AddUser is a paid mutator transaction binding the contract method 0x421b2d8b.
//
// Solidity: function addUser(address user) returns(bool)
func (_AppRegistry *AppRegistryTransactor) AddUser(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _AppRegistry.contract.Transact(opts, "addUser", user)
}

// AddUser is a paid mutator transaction binding the contract method 0x421b2d8b.
//
// Solidity: function addUser(address user) returns(bool)
func (_AppRegistry *AppRegistrySession) AddUser(user common.Address) (*types.Transaction, error) {
	return _AppRegistry.Contract.AddUser(&_AppRegistry.TransactOpts, user)
}

// AddUser is a paid mutator transaction binding the contract method 0x421b2d8b.
//
// Solidity: function addUser(address user) returns(bool)
func (_AppRegistry *AppRegistryTransactorSession) AddUser(user common.Address) (*types.Transaction, error) {
	return _AppRegistry.Contract.AddUser(&_AppRegistry.TransactOpts, user)
}

// Create is a paid mutator transaction binding the contract method 0x63835563.
//
// Solidity: function create(string name, string symbol, string link, string description, uint8 paymentType_, uint256 deferTimeSecs, uint256 defaultApiWeight, address owner) returns(address)
func (_AppRegistry *AppRegistryTransactor) Create(opts *bind.TransactOpts, name string, symbol string, link string, description string, paymentType_ uint8, deferTimeSecs *big.Int, defaultApiWeight *big.Int, owner common.Address) (*types.Transaction, error) {
	return _AppRegistry.contract.Transact(opts, "create", name, symbol, link, description, paymentType_, deferTimeSecs, defaultApiWeight, owner)
}

// Create is a paid mutator transaction binding the contract method 0x63835563.
//
// Solidity: function create(string name, string symbol, string link, string description, uint8 paymentType_, uint256 deferTimeSecs, uint256 defaultApiWeight, address owner) returns(address)
func (_AppRegistry *AppRegistrySession) Create(name string, symbol string, link string, description string, paymentType_ uint8, deferTimeSecs *big.Int, defaultApiWeight *big.Int, owner common.Address) (*types.Transaction, error) {
	return _AppRegistry.Contract.Create(&_AppRegistry.TransactOpts, name, symbol, link, description, paymentType_, deferTimeSecs, defaultApiWeight, owner)
}

// Create is a paid mutator transaction binding the contract method 0x63835563.
//
// Solidity: function create(string name, string symbol, string link, string description, uint8 paymentType_, uint256 deferTimeSecs, uint256 defaultApiWeight, address owner) returns(address)
func (_AppRegistry *AppRegistryTransactorSession) Create(name string, symbol string, link string, description string, paymentType_ uint8, deferTimeSecs *big.Int, defaultApiWeight *big.Int, owner common.Address) (*types.Transaction, error) {
	return _AppRegistry.Contract.Create(&_AppRegistry.TransactOpts, name, symbol, link, description, paymentType_, deferTimeSecs, defaultApiWeight, owner)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AppRegistry *AppRegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AppRegistry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AppRegistry *AppRegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AppRegistry.Contract.GrantRole(&_AppRegistry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AppRegistry *AppRegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AppRegistry.Contract.GrantRole(&_AppRegistry.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address appFactory_, address exchanger_) returns()
func (_AppRegistry *AppRegistryTransactor) Initialize(opts *bind.TransactOpts, appFactory_ common.Address, exchanger_ common.Address) (*types.Transaction, error) {
	return _AppRegistry.contract.Transact(opts, "initialize", appFactory_, exchanger_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address appFactory_, address exchanger_) returns()
func (_AppRegistry *AppRegistrySession) Initialize(appFactory_ common.Address, exchanger_ common.Address) (*types.Transaction, error) {
	return _AppRegistry.Contract.Initialize(&_AppRegistry.TransactOpts, appFactory_, exchanger_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address appFactory_, address exchanger_) returns()
func (_AppRegistry *AppRegistryTransactorSession) Initialize(appFactory_ common.Address, exchanger_ common.Address) (*types.Transaction, error) {
	return _AppRegistry.Contract.Initialize(&_AppRegistry.TransactOpts, appFactory_, exchanger_)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(address app) returns()
func (_AppRegistry *AppRegistryTransactor) Remove(opts *bind.TransactOpts, app common.Address) (*types.Transaction, error) {
	return _AppRegistry.contract.Transact(opts, "remove", app)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(address app) returns()
func (_AppRegistry *AppRegistrySession) Remove(app common.Address) (*types.Transaction, error) {
	return _AppRegistry.Contract.Remove(&_AppRegistry.TransactOpts, app)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(address app) returns()
func (_AppRegistry *AppRegistryTransactorSession) Remove(app common.Address) (*types.Transaction, error) {
	return _AppRegistry.Contract.Remove(&_AppRegistry.TransactOpts, app)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AppRegistry *AppRegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AppRegistry.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AppRegistry *AppRegistrySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AppRegistry.Contract.RenounceRole(&_AppRegistry.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AppRegistry *AppRegistryTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AppRegistry.Contract.RenounceRole(&_AppRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AppRegistry *AppRegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AppRegistry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AppRegistry *AppRegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AppRegistry.Contract.RevokeRole(&_AppRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AppRegistry *AppRegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AppRegistry.Contract.RevokeRole(&_AppRegistry.TransactOpts, role, account)
}

// SetCreatorRoleDisabled is a paid mutator transaction binding the contract method 0xd86193bb.
//
// Solidity: function setCreatorRoleDisabled(bool disabled) returns()
func (_AppRegistry *AppRegistryTransactor) SetCreatorRoleDisabled(opts *bind.TransactOpts, disabled bool) (*types.Transaction, error) {
	return _AppRegistry.contract.Transact(opts, "setCreatorRoleDisabled", disabled)
}

// SetCreatorRoleDisabled is a paid mutator transaction binding the contract method 0xd86193bb.
//
// Solidity: function setCreatorRoleDisabled(bool disabled) returns()
func (_AppRegistry *AppRegistrySession) SetCreatorRoleDisabled(disabled bool) (*types.Transaction, error) {
	return _AppRegistry.Contract.SetCreatorRoleDisabled(&_AppRegistry.TransactOpts, disabled)
}

// SetCreatorRoleDisabled is a paid mutator transaction binding the contract method 0xd86193bb.
//
// Solidity: function setCreatorRoleDisabled(bool disabled) returns()
func (_AppRegistry *AppRegistryTransactorSession) SetCreatorRoleDisabled(disabled bool) (*types.Transaction, error) {
	return _AppRegistry.Contract.SetCreatorRoleDisabled(&_AppRegistry.TransactOpts, disabled)
}

// SetExchanger is a paid mutator transaction binding the contract method 0xda909b09.
//
// Solidity: function setExchanger(address exchanger_) returns()
func (_AppRegistry *AppRegistryTransactor) SetExchanger(opts *bind.TransactOpts, exchanger_ common.Address) (*types.Transaction, error) {
	return _AppRegistry.contract.Transact(opts, "setExchanger", exchanger_)
}

// SetExchanger is a paid mutator transaction binding the contract method 0xda909b09.
//
// Solidity: function setExchanger(address exchanger_) returns()
func (_AppRegistry *AppRegistrySession) SetExchanger(exchanger_ common.Address) (*types.Transaction, error) {
	return _AppRegistry.Contract.SetExchanger(&_AppRegistry.TransactOpts, exchanger_)
}

// SetExchanger is a paid mutator transaction binding the contract method 0xda909b09.
//
// Solidity: function setExchanger(address exchanger_) returns()
func (_AppRegistry *AppRegistryTransactorSession) SetExchanger(exchanger_ common.Address) (*types.Transaction, error) {
	return _AppRegistry.Contract.SetExchanger(&_AppRegistry.TransactOpts, exchanger_)
}

// AppRegistryCreatedIterator is returned from FilterCreated and is used to iterate over the raw logs and unpacked data for Created events raised by the AppRegistry contract.
type AppRegistryCreatedIterator struct {
	Event *AppRegistryCreated // Event containing the contract specifics and raw log

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
func (it *AppRegistryCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppRegistryCreated)
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
		it.Event = new(AppRegistryCreated)
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
func (it *AppRegistryCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppRegistryCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppRegistryCreated represents a Created event raised by the AppRegistry contract.
type AppRegistryCreated struct {
	App            common.Address
	Operator       common.Address
	Owner          common.Address
	ApiWeightToken common.Address
	VipCoin        common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCreated is a free log retrieval operation binding the contract event 0x55353cd176c454ba526e09653a5a00a4cc4d1468d129208df7f84e0a359f688d.
//
// Solidity: event Created(address indexed app, address indexed operator, address indexed owner, address apiWeightToken, address vipCoin)
func (_AppRegistry *AppRegistryFilterer) FilterCreated(opts *bind.FilterOpts, app []common.Address, operator []common.Address, owner []common.Address) (*AppRegistryCreatedIterator, error) {

	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _AppRegistry.contract.FilterLogs(opts, "Created", appRule, operatorRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &AppRegistryCreatedIterator{contract: _AppRegistry.contract, event: "Created", logs: logs, sub: sub}, nil
}

// WatchCreated is a free log subscription operation binding the contract event 0x55353cd176c454ba526e09653a5a00a4cc4d1468d129208df7f84e0a359f688d.
//
// Solidity: event Created(address indexed app, address indexed operator, address indexed owner, address apiWeightToken, address vipCoin)
func (_AppRegistry *AppRegistryFilterer) WatchCreated(opts *bind.WatchOpts, sink chan<- *AppRegistryCreated, app []common.Address, operator []common.Address, owner []common.Address) (event.Subscription, error) {

	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _AppRegistry.contract.WatchLogs(opts, "Created", appRule, operatorRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppRegistryCreated)
				if err := _AppRegistry.contract.UnpackLog(event, "Created", log); err != nil {
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

// ParseCreated is a log parse operation binding the contract event 0x55353cd176c454ba526e09653a5a00a4cc4d1468d129208df7f84e0a359f688d.
//
// Solidity: event Created(address indexed app, address indexed operator, address indexed owner, address apiWeightToken, address vipCoin)
func (_AppRegistry *AppRegistryFilterer) ParseCreated(log types.Log) (*AppRegistryCreated, error) {
	event := new(AppRegistryCreated)
	if err := _AppRegistry.contract.UnpackLog(event, "Created", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AppRegistry contract.
type AppRegistryInitializedIterator struct {
	Event *AppRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *AppRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppRegistryInitialized)
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
		it.Event = new(AppRegistryInitialized)
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
func (it *AppRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppRegistryInitialized represents a Initialized event raised by the AppRegistry contract.
type AppRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AppRegistry *AppRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*AppRegistryInitializedIterator, error) {

	logs, sub, err := _AppRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AppRegistryInitializedIterator{contract: _AppRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AppRegistry *AppRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AppRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _AppRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppRegistryInitialized)
				if err := _AppRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AppRegistry *AppRegistryFilterer) ParseInitialized(log types.Log) (*AppRegistryInitialized, error) {
	event := new(AppRegistryInitialized)
	if err := _AppRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppRegistryRemovedIterator is returned from FilterRemoved and is used to iterate over the raw logs and unpacked data for Removed events raised by the AppRegistry contract.
type AppRegistryRemovedIterator struct {
	Event *AppRegistryRemoved // Event containing the contract specifics and raw log

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
func (it *AppRegistryRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppRegistryRemoved)
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
		it.Event = new(AppRegistryRemoved)
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
func (it *AppRegistryRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppRegistryRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppRegistryRemoved represents a Removed event raised by the AppRegistry contract.
type AppRegistryRemoved struct {
	App      common.Address
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRemoved is a free log retrieval operation binding the contract event 0x40e634d0e26d9ec2e860e4dd9b7b2cfbb569b6058362a1a54d3a94718bc49587.
//
// Solidity: event Removed(address indexed app, address indexed operator)
func (_AppRegistry *AppRegistryFilterer) FilterRemoved(opts *bind.FilterOpts, app []common.Address, operator []common.Address) (*AppRegistryRemovedIterator, error) {

	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AppRegistry.contract.FilterLogs(opts, "Removed", appRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &AppRegistryRemovedIterator{contract: _AppRegistry.contract, event: "Removed", logs: logs, sub: sub}, nil
}

// WatchRemoved is a free log subscription operation binding the contract event 0x40e634d0e26d9ec2e860e4dd9b7b2cfbb569b6058362a1a54d3a94718bc49587.
//
// Solidity: event Removed(address indexed app, address indexed operator)
func (_AppRegistry *AppRegistryFilterer) WatchRemoved(opts *bind.WatchOpts, sink chan<- *AppRegistryRemoved, app []common.Address, operator []common.Address) (event.Subscription, error) {

	var appRule []interface{}
	for _, appItem := range app {
		appRule = append(appRule, appItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AppRegistry.contract.WatchLogs(opts, "Removed", appRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppRegistryRemoved)
				if err := _AppRegistry.contract.UnpackLog(event, "Removed", log); err != nil {
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

// ParseRemoved is a log parse operation binding the contract event 0x40e634d0e26d9ec2e860e4dd9b7b2cfbb569b6058362a1a54d3a94718bc49587.
//
// Solidity: event Removed(address indexed app, address indexed operator)
func (_AppRegistry *AppRegistryFilterer) ParseRemoved(log types.Log) (*AppRegistryRemoved, error) {
	event := new(AppRegistryRemoved)
	if err := _AppRegistry.contract.UnpackLog(event, "Removed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppRegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AppRegistry contract.
type AppRegistryRoleAdminChangedIterator struct {
	Event *AppRegistryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AppRegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppRegistryRoleAdminChanged)
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
		it.Event = new(AppRegistryRoleAdminChanged)
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
func (it *AppRegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppRegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppRegistryRoleAdminChanged represents a RoleAdminChanged event raised by the AppRegistry contract.
type AppRegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AppRegistry *AppRegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AppRegistryRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AppRegistry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AppRegistryRoleAdminChangedIterator{contract: _AppRegistry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AppRegistry *AppRegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AppRegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AppRegistry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppRegistryRoleAdminChanged)
				if err := _AppRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AppRegistry *AppRegistryFilterer) ParseRoleAdminChanged(log types.Log) (*AppRegistryRoleAdminChanged, error) {
	event := new(AppRegistryRoleAdminChanged)
	if err := _AppRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppRegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AppRegistry contract.
type AppRegistryRoleGrantedIterator struct {
	Event *AppRegistryRoleGranted // Event containing the contract specifics and raw log

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
func (it *AppRegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppRegistryRoleGranted)
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
		it.Event = new(AppRegistryRoleGranted)
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
func (it *AppRegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppRegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppRegistryRoleGranted represents a RoleGranted event raised by the AppRegistry contract.
type AppRegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AppRegistry *AppRegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AppRegistryRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AppRegistry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AppRegistryRoleGrantedIterator{contract: _AppRegistry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AppRegistry *AppRegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AppRegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AppRegistry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppRegistryRoleGranted)
				if err := _AppRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AppRegistry *AppRegistryFilterer) ParseRoleGranted(log types.Log) (*AppRegistryRoleGranted, error) {
	event := new(AppRegistryRoleGranted)
	if err := _AppRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppRegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AppRegistry contract.
type AppRegistryRoleRevokedIterator struct {
	Event *AppRegistryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AppRegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppRegistryRoleRevoked)
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
		it.Event = new(AppRegistryRoleRevoked)
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
func (it *AppRegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppRegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppRegistryRoleRevoked represents a RoleRevoked event raised by the AppRegistry contract.
type AppRegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AppRegistry *AppRegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AppRegistryRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AppRegistry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AppRegistryRoleRevokedIterator{contract: _AppRegistry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AppRegistry *AppRegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AppRegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AppRegistry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppRegistryRoleRevoked)
				if err := _AppRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AppRegistry *AppRegistryFilterer) ParseRoleRevoked(log types.Log) (*AppRegistryRoleRevoked, error) {
	event := new(AppRegistryRoleRevoked)
	if err := _AppRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
