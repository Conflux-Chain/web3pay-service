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

// IAppConfigChargeRequest is an auto generated low-level Go binding around an user-defined struct.
type IAppConfigChargeRequest struct {
	Account   common.Address
	Amount    *big.Int
	Data      []byte
	UseDetail []IAppConfigResourceUseDetail
}

// AppMetaData contains all meta data concerning the App contract.
var AppMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Frozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHARGE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TAKE_PROFIT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_ID_AIRDROP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_ID_COIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_ID_VIP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAW_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"airdrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"receivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"reasons\",\"type\":\"string[]\"}],\"name\":\"airdropBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"appRegistry\",\"outputs\":[{\"internalType\":\"contractIAppRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cardShop\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"times\",\"type\":\"uint256\"}],\"internalType\":\"structIAppConfig.ResourceUseDetail[]\",\"name\":\"useDetail\",\"type\":\"tuple[]\"}],\"internalType\":\"structIAppConfig.ChargeRequest[]\",\"name\":\"requestArray\",\"type\":\"tuple[]\"}],\"name\":\"chargeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deferTimeSecs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"depositAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"toAssets\",\"type\":\"bool\"}],\"name\":\"forceWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"contractIWithdrawHook\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ethMin\",\"type\":\"uint256\"}],\"name\":\"forceWithdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getApiWeightToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAppCoin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVipCoin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAppCoinV2\",\"name\":\"appCoin_\",\"type\":\"address\"},{\"internalType\":\"contractIVipCoin\",\"name\":\"vipCoin_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"apiWeightToken_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deferTimeSecs_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"contractIAppRegistry\",\"name\":\"appRegistry_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"link\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"makeCard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paymentType\",\"outputs\":[{\"internalType\":\"enumIApp.PaymentType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestForceWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cardShop_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"link_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description_\",\"type\":\"string\"},{\"internalType\":\"enumIApp.PaymentType\",\"name\":\"paymentType_\",\"type\":\"uint8\"}],\"name\":\"setProps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"takeProfit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalCharged\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalTakenProfit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"toAssets\",\"type\":\"bool\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawSchedules\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AppABI is the input ABI used to generate the binding from.
// Deprecated: Use AppMetaData.ABI instead.
var AppABI = AppMetaData.ABI

// App is an auto generated Go binding around an Ethereum contract.
type App struct {
	AppCaller     // Read-only binding to the contract
	AppTransactor // Write-only binding to the contract
	AppFilterer   // Log filterer for contract events
}

// AppCaller is an auto generated read-only Go binding around an Ethereum contract.
type AppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AppFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AppSession struct {
	Contract     *App              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AppCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AppCallerSession struct {
	Contract *AppCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AppTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AppTransactorSession struct {
	Contract     *AppTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AppRaw is an auto generated low-level Go binding around an Ethereum contract.
type AppRaw struct {
	Contract *App // Generic contract binding to access the raw methods on
}

// AppCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AppCallerRaw struct {
	Contract *AppCaller // Generic read-only contract binding to access the raw methods on
}

// AppTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AppTransactorRaw struct {
	Contract *AppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApp creates a new instance of App, bound to a specific deployed contract.
func NewApp(address common.Address, backend bind.ContractBackend) (*App, error) {
	contract, err := bindApp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &App{AppCaller: AppCaller{contract: contract}, AppTransactor: AppTransactor{contract: contract}, AppFilterer: AppFilterer{contract: contract}}, nil
}

// NewAppCaller creates a new read-only instance of App, bound to a specific deployed contract.
func NewAppCaller(address common.Address, caller bind.ContractCaller) (*AppCaller, error) {
	contract, err := bindApp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AppCaller{contract: contract}, nil
}

// NewAppTransactor creates a new write-only instance of App, bound to a specific deployed contract.
func NewAppTransactor(address common.Address, transactor bind.ContractTransactor) (*AppTransactor, error) {
	contract, err := bindApp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AppTransactor{contract: contract}, nil
}

// NewAppFilterer creates a new log filterer instance of App, bound to a specific deployed contract.
func NewAppFilterer(address common.Address, filterer bind.ContractFilterer) (*AppFilterer, error) {
	contract, err := bindApp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AppFilterer{contract: contract}, nil
}

// bindApp binds a generic wrapper to an already deployed contract.
func bindApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AppABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_App *AppRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _App.Contract.AppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_App *AppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _App.Contract.AppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_App *AppRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _App.Contract.AppTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_App *AppCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _App.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_App *AppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _App.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_App *AppTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _App.Contract.contract.Transact(opts, method, params...)
}

// CHARGEROLE is a free data retrieval call binding the contract method 0xa2c74ad9.
//
// Solidity: function CHARGE_ROLE() view returns(bytes32)
func (_App *AppCaller) CHARGEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "CHARGE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CHARGEROLE is a free data retrieval call binding the contract method 0xa2c74ad9.
//
// Solidity: function CHARGE_ROLE() view returns(bytes32)
func (_App *AppSession) CHARGEROLE() ([32]byte, error) {
	return _App.Contract.CHARGEROLE(&_App.CallOpts)
}

// CHARGEROLE is a free data retrieval call binding the contract method 0xa2c74ad9.
//
// Solidity: function CHARGE_ROLE() view returns(bytes32)
func (_App *AppCallerSession) CHARGEROLE() ([32]byte, error) {
	return _App.Contract.CHARGEROLE(&_App.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_App *AppCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_App *AppSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _App.Contract.DEFAULTADMINROLE(&_App.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_App *AppCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _App.Contract.DEFAULTADMINROLE(&_App.CallOpts)
}

// TAKEPROFITROLE is a free data retrieval call binding the contract method 0x0da699d6.
//
// Solidity: function TAKE_PROFIT_ROLE() view returns(bytes32)
func (_App *AppCaller) TAKEPROFITROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "TAKE_PROFIT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TAKEPROFITROLE is a free data retrieval call binding the contract method 0x0da699d6.
//
// Solidity: function TAKE_PROFIT_ROLE() view returns(bytes32)
func (_App *AppSession) TAKEPROFITROLE() ([32]byte, error) {
	return _App.Contract.TAKEPROFITROLE(&_App.CallOpts)
}

// TAKEPROFITROLE is a free data retrieval call binding the contract method 0x0da699d6.
//
// Solidity: function TAKE_PROFIT_ROLE() view returns(bytes32)
func (_App *AppCallerSession) TAKEPROFITROLE() ([32]byte, error) {
	return _App.Contract.TAKEPROFITROLE(&_App.CallOpts)
}

// TOKENIDAIRDROP is a free data retrieval call binding the contract method 0xc9e3b913.
//
// Solidity: function TOKEN_ID_AIRDROP() view returns(uint256)
func (_App *AppCaller) TOKENIDAIRDROP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "TOKEN_ID_AIRDROP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOKENIDAIRDROP is a free data retrieval call binding the contract method 0xc9e3b913.
//
// Solidity: function TOKEN_ID_AIRDROP() view returns(uint256)
func (_App *AppSession) TOKENIDAIRDROP() (*big.Int, error) {
	return _App.Contract.TOKENIDAIRDROP(&_App.CallOpts)
}

// TOKENIDAIRDROP is a free data retrieval call binding the contract method 0xc9e3b913.
//
// Solidity: function TOKEN_ID_AIRDROP() view returns(uint256)
func (_App *AppCallerSession) TOKENIDAIRDROP() (*big.Int, error) {
	return _App.Contract.TOKENIDAIRDROP(&_App.CallOpts)
}

// TOKENIDCOIN is a free data retrieval call binding the contract method 0x45e2efd2.
//
// Solidity: function TOKEN_ID_COIN() view returns(uint256)
func (_App *AppCaller) TOKENIDCOIN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "TOKEN_ID_COIN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOKENIDCOIN is a free data retrieval call binding the contract method 0x45e2efd2.
//
// Solidity: function TOKEN_ID_COIN() view returns(uint256)
func (_App *AppSession) TOKENIDCOIN() (*big.Int, error) {
	return _App.Contract.TOKENIDCOIN(&_App.CallOpts)
}

// TOKENIDCOIN is a free data retrieval call binding the contract method 0x45e2efd2.
//
// Solidity: function TOKEN_ID_COIN() view returns(uint256)
func (_App *AppCallerSession) TOKENIDCOIN() (*big.Int, error) {
	return _App.Contract.TOKENIDCOIN(&_App.CallOpts)
}

// TOKENIDVIP is a free data retrieval call binding the contract method 0xdbd53229.
//
// Solidity: function TOKEN_ID_VIP() view returns(uint256)
func (_App *AppCaller) TOKENIDVIP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "TOKEN_ID_VIP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOKENIDVIP is a free data retrieval call binding the contract method 0xdbd53229.
//
// Solidity: function TOKEN_ID_VIP() view returns(uint256)
func (_App *AppSession) TOKENIDVIP() (*big.Int, error) {
	return _App.Contract.TOKENIDVIP(&_App.CallOpts)
}

// TOKENIDVIP is a free data retrieval call binding the contract method 0xdbd53229.
//
// Solidity: function TOKEN_ID_VIP() view returns(uint256)
func (_App *AppCallerSession) TOKENIDVIP() (*big.Int, error) {
	return _App.Contract.TOKENIDVIP(&_App.CallOpts)
}

// WITHDRAWROLE is a free data retrieval call binding the contract method 0xe02023a1.
//
// Solidity: function WITHDRAW_ROLE() view returns(bytes32)
func (_App *AppCaller) WITHDRAWROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "WITHDRAW_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWROLE is a free data retrieval call binding the contract method 0xe02023a1.
//
// Solidity: function WITHDRAW_ROLE() view returns(bytes32)
func (_App *AppSession) WITHDRAWROLE() ([32]byte, error) {
	return _App.Contract.WITHDRAWROLE(&_App.CallOpts)
}

// WITHDRAWROLE is a free data retrieval call binding the contract method 0xe02023a1.
//
// Solidity: function WITHDRAW_ROLE() view returns(bytes32)
func (_App *AppCallerSession) WITHDRAWROLE() ([32]byte, error) {
	return _App.Contract.WITHDRAWROLE(&_App.CallOpts)
}

// AppRegistry is a free data retrieval call binding the contract method 0xbb4fceb9.
//
// Solidity: function appRegistry() view returns(address)
func (_App *AppCaller) AppRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "appRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AppRegistry is a free data retrieval call binding the contract method 0xbb4fceb9.
//
// Solidity: function appRegistry() view returns(address)
func (_App *AppSession) AppRegistry() (common.Address, error) {
	return _App.Contract.AppRegistry(&_App.CallOpts)
}

// AppRegistry is a free data retrieval call binding the contract method 0xbb4fceb9.
//
// Solidity: function appRegistry() view returns(address)
func (_App *AppCallerSession) AppRegistry() (common.Address, error) {
	return _App.Contract.AppRegistry(&_App.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256, uint256)
func (_App *AppCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256, uint256)
func (_App *AppSession) BalanceOf(account common.Address) (*big.Int, *big.Int, error) {
	return _App.Contract.BalanceOf(&_App.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256, uint256)
func (_App *AppCallerSession) BalanceOf(account common.Address) (*big.Int, *big.Int, error) {
	return _App.Contract.BalanceOf(&_App.CallOpts, account)
}

// CardShop is a free data retrieval call binding the contract method 0xd6773e6b.
//
// Solidity: function cardShop() view returns(address)
func (_App *AppCaller) CardShop(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "cardShop")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CardShop is a free data retrieval call binding the contract method 0xd6773e6b.
//
// Solidity: function cardShop() view returns(address)
func (_App *AppSession) CardShop() (common.Address, error) {
	return _App.Contract.CardShop(&_App.CallOpts)
}

// CardShop is a free data retrieval call binding the contract method 0xd6773e6b.
//
// Solidity: function cardShop() view returns(address)
func (_App *AppCallerSession) CardShop() (common.Address, error) {
	return _App.Contract.CardShop(&_App.CallOpts)
}

// DeferTimeSecs is a free data retrieval call binding the contract method 0x3e0510f8.
//
// Solidity: function deferTimeSecs() view returns(uint256)
func (_App *AppCaller) DeferTimeSecs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "deferTimeSecs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeferTimeSecs is a free data retrieval call binding the contract method 0x3e0510f8.
//
// Solidity: function deferTimeSecs() view returns(uint256)
func (_App *AppSession) DeferTimeSecs() (*big.Int, error) {
	return _App.Contract.DeferTimeSecs(&_App.CallOpts)
}

// DeferTimeSecs is a free data retrieval call binding the contract method 0x3e0510f8.
//
// Solidity: function deferTimeSecs() view returns(uint256)
func (_App *AppCallerSession) DeferTimeSecs() (*big.Int, error) {
	return _App.Contract.DeferTimeSecs(&_App.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_App *AppCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_App *AppSession) Description() (string, error) {
	return _App.Contract.Description(&_App.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_App *AppCallerSession) Description() (string, error) {
	return _App.Contract.Description(&_App.CallOpts)
}

// GetApiWeightToken is a free data retrieval call binding the contract method 0xcf46544e.
//
// Solidity: function getApiWeightToken() view returns(address)
func (_App *AppCaller) GetApiWeightToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getApiWeightToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApiWeightToken is a free data retrieval call binding the contract method 0xcf46544e.
//
// Solidity: function getApiWeightToken() view returns(address)
func (_App *AppSession) GetApiWeightToken() (common.Address, error) {
	return _App.Contract.GetApiWeightToken(&_App.CallOpts)
}

// GetApiWeightToken is a free data retrieval call binding the contract method 0xcf46544e.
//
// Solidity: function getApiWeightToken() view returns(address)
func (_App *AppCallerSession) GetApiWeightToken() (common.Address, error) {
	return _App.Contract.GetApiWeightToken(&_App.CallOpts)
}

// GetAppCoin is a free data retrieval call binding the contract method 0xe7337217.
//
// Solidity: function getAppCoin() view returns(address)
func (_App *AppCaller) GetAppCoin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getAppCoin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAppCoin is a free data retrieval call binding the contract method 0xe7337217.
//
// Solidity: function getAppCoin() view returns(address)
func (_App *AppSession) GetAppCoin() (common.Address, error) {
	return _App.Contract.GetAppCoin(&_App.CallOpts)
}

// GetAppCoin is a free data retrieval call binding the contract method 0xe7337217.
//
// Solidity: function getAppCoin() view returns(address)
func (_App *AppCallerSession) GetAppCoin() (common.Address, error) {
	return _App.Contract.GetAppCoin(&_App.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_App *AppCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_App *AppSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _App.Contract.GetRoleAdmin(&_App.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_App *AppCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _App.Contract.GetRoleAdmin(&_App.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_App *AppCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_App *AppSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _App.Contract.GetRoleMember(&_App.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_App *AppCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _App.Contract.GetRoleMember(&_App.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_App *AppCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_App *AppSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _App.Contract.GetRoleMemberCount(&_App.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_App *AppCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _App.Contract.GetRoleMemberCount(&_App.CallOpts, role)
}

// GetVipCoin is a free data retrieval call binding the contract method 0x07a75b24.
//
// Solidity: function getVipCoin() view returns(address)
func (_App *AppCaller) GetVipCoin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "getVipCoin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVipCoin is a free data retrieval call binding the contract method 0x07a75b24.
//
// Solidity: function getVipCoin() view returns(address)
func (_App *AppSession) GetVipCoin() (common.Address, error) {
	return _App.Contract.GetVipCoin(&_App.CallOpts)
}

// GetVipCoin is a free data retrieval call binding the contract method 0x07a75b24.
//
// Solidity: function getVipCoin() view returns(address)
func (_App *AppCallerSession) GetVipCoin() (common.Address, error) {
	return _App.Contract.GetVipCoin(&_App.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_App *AppCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_App *AppSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _App.Contract.HasRole(&_App.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_App *AppCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _App.Contract.HasRole(&_App.CallOpts, role, account)
}

// Link is a free data retrieval call binding the contract method 0x1c4695f4.
//
// Solidity: function link() view returns(string)
func (_App *AppCaller) Link(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "link")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Link is a free data retrieval call binding the contract method 0x1c4695f4.
//
// Solidity: function link() view returns(string)
func (_App *AppSession) Link() (string, error) {
	return _App.Contract.Link(&_App.CallOpts)
}

// Link is a free data retrieval call binding the contract method 0x1c4695f4.
//
// Solidity: function link() view returns(string)
func (_App *AppCallerSession) Link() (string, error) {
	return _App.Contract.Link(&_App.CallOpts)
}

// PaymentType is a free data retrieval call binding the contract method 0x2763b8da.
//
// Solidity: function paymentType() view returns(uint8)
func (_App *AppCaller) PaymentType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "paymentType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PaymentType is a free data retrieval call binding the contract method 0x2763b8da.
//
// Solidity: function paymentType() view returns(uint8)
func (_App *AppSession) PaymentType() (uint8, error) {
	return _App.Contract.PaymentType(&_App.CallOpts)
}

// PaymentType is a free data retrieval call binding the contract method 0x2763b8da.
//
// Solidity: function paymentType() view returns(uint8)
func (_App *AppCallerSession) PaymentType() (uint8, error) {
	return _App.Contract.PaymentType(&_App.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_App *AppCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_App *AppSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _App.Contract.SupportsInterface(&_App.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_App *AppCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _App.Contract.SupportsInterface(&_App.CallOpts, interfaceId)
}

// TotalCharged is a free data retrieval call binding the contract method 0xa6f55d77.
//
// Solidity: function totalCharged() view returns(uint256)
func (_App *AppCaller) TotalCharged(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "totalCharged")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCharged is a free data retrieval call binding the contract method 0xa6f55d77.
//
// Solidity: function totalCharged() view returns(uint256)
func (_App *AppSession) TotalCharged() (*big.Int, error) {
	return _App.Contract.TotalCharged(&_App.CallOpts)
}

// TotalCharged is a free data retrieval call binding the contract method 0xa6f55d77.
//
// Solidity: function totalCharged() view returns(uint256)
func (_App *AppCallerSession) TotalCharged() (*big.Int, error) {
	return _App.Contract.TotalCharged(&_App.CallOpts)
}

// TotalTakenProfit is a free data retrieval call binding the contract method 0x55297bcf.
//
// Solidity: function totalTakenProfit() view returns(uint256)
func (_App *AppCaller) TotalTakenProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "totalTakenProfit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalTakenProfit is a free data retrieval call binding the contract method 0x55297bcf.
//
// Solidity: function totalTakenProfit() view returns(uint256)
func (_App *AppSession) TotalTakenProfit() (*big.Int, error) {
	return _App.Contract.TotalTakenProfit(&_App.CallOpts)
}

// TotalTakenProfit is a free data retrieval call binding the contract method 0x55297bcf.
//
// Solidity: function totalTakenProfit() view returns(uint256)
func (_App *AppCallerSession) TotalTakenProfit() (*big.Int, error) {
	return _App.Contract.TotalTakenProfit(&_App.CallOpts)
}

// WithdrawSchedules is a free data retrieval call binding the contract method 0x0a9aa21c.
//
// Solidity: function withdrawSchedules(address ) view returns(uint256)
func (_App *AppCaller) WithdrawSchedules(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _App.contract.Call(opts, &out, "withdrawSchedules", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawSchedules is a free data retrieval call binding the contract method 0x0a9aa21c.
//
// Solidity: function withdrawSchedules(address ) view returns(uint256)
func (_App *AppSession) WithdrawSchedules(arg0 common.Address) (*big.Int, error) {
	return _App.Contract.WithdrawSchedules(&_App.CallOpts, arg0)
}

// WithdrawSchedules is a free data retrieval call binding the contract method 0x0a9aa21c.
//
// Solidity: function withdrawSchedules(address ) view returns(uint256)
func (_App *AppCallerSession) WithdrawSchedules(arg0 common.Address) (*big.Int, error) {
	return _App.Contract.WithdrawSchedules(&_App.CallOpts, arg0)
}

// Airdrop is a paid mutator transaction binding the contract method 0x8ba4cc3c.
//
// Solidity: function airdrop(address receiver, uint256 amount) returns()
func (_App *AppTransactor) Airdrop(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _App.contract.Transact(opts, "airdrop", receiver, amount)
}

// Airdrop is a paid mutator transaction binding the contract method 0x8ba4cc3c.
//
// Solidity: function airdrop(address receiver, uint256 amount) returns()
func (_App *AppSession) Airdrop(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _App.Contract.Airdrop(&_App.TransactOpts, receiver, amount)
}

// Airdrop is a paid mutator transaction binding the contract method 0x8ba4cc3c.
//
// Solidity: function airdrop(address receiver, uint256 amount) returns()
func (_App *AppTransactorSession) Airdrop(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _App.Contract.Airdrop(&_App.TransactOpts, receiver, amount)
}

// AirdropBatch is a paid mutator transaction binding the contract method 0x9c5c2c32.
//
// Solidity: function airdropBatch(address[] receivers, uint256[] amounts, string[] reasons) returns()
func (_App *AppTransactor) AirdropBatch(opts *bind.TransactOpts, receivers []common.Address, amounts []*big.Int, reasons []string) (*types.Transaction, error) {
	return _App.contract.Transact(opts, "airdropBatch", receivers, amounts, reasons)
}

// AirdropBatch is a paid mutator transaction binding the contract method 0x9c5c2c32.
//
// Solidity: function airdropBatch(address[] receivers, uint256[] amounts, string[] reasons) returns()
func (_App *AppSession) AirdropBatch(receivers []common.Address, amounts []*big.Int, reasons []string) (*types.Transaction, error) {
	return _App.Contract.AirdropBatch(&_App.TransactOpts, receivers, amounts, reasons)
}

// AirdropBatch is a paid mutator transaction binding the contract method 0x9c5c2c32.
//
// Solidity: function airdropBatch(address[] receivers, uint256[] amounts, string[] reasons) returns()
func (_App *AppTransactorSession) AirdropBatch(receivers []common.Address, amounts []*big.Int, reasons []string) (*types.Transaction, error) {
	return _App.Contract.AirdropBatch(&_App.TransactOpts, receivers, amounts, reasons)
}

// ChargeBatch is a paid mutator transaction binding the contract method 0x0664ec3b.
//
// Solidity: function chargeBatch((address,uint256,bytes,(uint32,uint256)[])[] requestArray) returns()
func (_App *AppTransactor) ChargeBatch(opts *bind.TransactOpts, requestArray []IAppConfigChargeRequest) (*types.Transaction, error) {
	return _App.contract.Transact(opts, "chargeBatch", requestArray)
}

// ChargeBatch is a paid mutator transaction binding the contract method 0x0664ec3b.
//
// Solidity: function chargeBatch((address,uint256,bytes,(uint32,uint256)[])[] requestArray) returns()
func (_App *AppSession) ChargeBatch(requestArray []IAppConfigChargeRequest) (*types.Transaction, error) {
	return _App.Contract.ChargeBatch(&_App.TransactOpts, requestArray)
}

// ChargeBatch is a paid mutator transaction binding the contract method 0x0664ec3b.
//
// Solidity: function chargeBatch((address,uint256,bytes,(uint32,uint256)[])[] requestArray) returns()
func (_App *AppTransactorSession) ChargeBatch(requestArray []IAppConfigChargeRequest) (*types.Transaction, error) {
	return _App.Contract.ChargeBatch(&_App.TransactOpts, requestArray)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address receiver) returns()
func (_App *AppTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _App.contract.Transact(opts, "deposit", amount, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address receiver) returns()
func (_App *AppSession) Deposit(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _App.Contract.Deposit(&_App.TransactOpts, amount, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address receiver) returns()
func (_App *AppTransactorSession) Deposit(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _App.Contract.Deposit(&_App.TransactOpts, amount, receiver)
}

// DepositAsset is a paid mutator transaction binding the contract method 0x95e213d2.
//
// Solidity: function depositAsset(uint256 amount, address receiver) returns()
func (_App *AppTransactor) DepositAsset(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _App.contract.Transact(opts, "depositAsset", amount, receiver)
}

// DepositAsset is a paid mutator transaction binding the contract method 0x95e213d2.
//
// Solidity: function depositAsset(uint256 amount, address receiver) returns()
func (_App *AppSession) DepositAsset(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _App.Contract.DepositAsset(&_App.TransactOpts, amount, receiver)
}

// DepositAsset is a paid mutator transaction binding the contract method 0x95e213d2.
//
// Solidity: function depositAsset(uint256 amount, address receiver) returns()
func (_App *AppTransactorSession) DepositAsset(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _App.Contract.DepositAsset(&_App.TransactOpts, amount, receiver)
}

// ForceWithdraw is a paid mutator transaction binding the contract method 0xc8896bf6.
//
// Solidity: function forceWithdraw(address receiver, bool toAssets) returns()
func (_App *AppTransactor) ForceWithdraw(opts *bind.TransactOpts, receiver common.Address, toAssets bool) (*types.Transaction, error) {
	return _App.contract.Transact(opts, "forceWithdraw", receiver, toAssets)
}

// ForceWithdraw is a paid mutator transaction binding the contract method 0xc8896bf6.
//
// Solidity: function forceWithdraw(address receiver, bool toAssets) returns()
func (_App *AppSession) ForceWithdraw(receiver common.Address, toAssets bool) (*types.Transaction, error) {
	return _App.Contract.ForceWithdraw(&_App.TransactOpts, receiver, toAssets)
}

// ForceWithdraw is a paid mutator transaction binding the contract method 0xc8896bf6.
//
// Solidity: function forceWithdraw(address receiver, bool toAssets) returns()
func (_App *AppTransactorSession) ForceWithdraw(receiver common.Address, toAssets bool) (*types.Transaction, error) {
	return _App.Contract.ForceWithdraw(&_App.TransactOpts, receiver, toAssets)
}

// ForceWithdrawEth is a paid mutator transaction binding the contract method 0x24516d37.
//
// Solidity: function forceWithdrawEth(address receiver, address hook, uint256 ethMin) returns()
func (_App *AppTransactor) ForceWithdrawEth(opts *bind.TransactOpts, receiver common.Address, hook common.Address, ethMin *big.Int) (*types.Transaction, error) {
	return _App.contract.Transact(opts, "forceWithdrawEth", receiver, hook, ethMin)
}

// ForceWithdrawEth is a paid mutator transaction binding the contract method 0x24516d37.
//
// Solidity: function forceWithdrawEth(address receiver, address hook, uint256 ethMin) returns()
func (_App *AppSession) ForceWithdrawEth(receiver common.Address, hook common.Address, ethMin *big.Int) (*types.Transaction, error) {
	return _App.Contract.ForceWithdrawEth(&_App.TransactOpts, receiver, hook, ethMin)
}

// ForceWithdrawEth is a paid mutator transaction binding the contract method 0x24516d37.
//
// Solidity: function forceWithdrawEth(address receiver, address hook, uint256 ethMin) returns()
func (_App *AppTransactorSession) ForceWithdrawEth(receiver common.Address, hook common.Address, ethMin *big.Int) (*types.Transaction, error) {
	return _App.Contract.ForceWithdrawEth(&_App.TransactOpts, receiver, hook, ethMin)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_App *AppTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _App.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_App *AppSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _App.Contract.GrantRole(&_App.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_App *AppTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _App.Contract.GrantRole(&_App.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x32b5a1b3.
//
// Solidity: function initialize(address appCoin_, address vipCoin_, address apiWeightToken_, uint256 deferTimeSecs_, address owner, address appRegistry_) returns()
func (_App *AppTransactor) Initialize(opts *bind.TransactOpts, appCoin_ common.Address, vipCoin_ common.Address, apiWeightToken_ common.Address, deferTimeSecs_ *big.Int, owner common.Address, appRegistry_ common.Address) (*types.Transaction, error) {
	return _App.contract.Transact(opts, "initialize", appCoin_, vipCoin_, apiWeightToken_, deferTimeSecs_, owner, appRegistry_)
}

// Initialize is a paid mutator transaction binding the contract method 0x32b5a1b3.
//
// Solidity: function initialize(address appCoin_, address vipCoin_, address apiWeightToken_, uint256 deferTimeSecs_, address owner, address appRegistry_) returns()
func (_App *AppSession) Initialize(appCoin_ common.Address, vipCoin_ common.Address, apiWeightToken_ common.Address, deferTimeSecs_ *big.Int, owner common.Address, appRegistry_ common.Address) (*types.Transaction, error) {
	return _App.Contract.Initialize(&_App.TransactOpts, appCoin_, vipCoin_, apiWeightToken_, deferTimeSecs_, owner, appRegistry_)
}

// Initialize is a paid mutator transaction binding the contract method 0x32b5a1b3.
//
// Solidity: function initialize(address appCoin_, address vipCoin_, address apiWeightToken_, uint256 deferTimeSecs_, address owner, address appRegistry_) returns()
func (_App *AppTransactorSession) Initialize(appCoin_ common.Address, vipCoin_ common.Address, apiWeightToken_ common.Address, deferTimeSecs_ *big.Int, owner common.Address, appRegistry_ common.Address) (*types.Transaction, error) {
	return _App.Contract.Initialize(&_App.TransactOpts, appCoin_, vipCoin_, apiWeightToken_, deferTimeSecs_, owner, appRegistry_)
}

// MakeCard is a paid mutator transaction binding the contract method 0xb33bbe6d.
//
// Solidity: function makeCard(address to, uint256 tokenId, uint256 amount) returns()
func (_App *AppTransactor) MakeCard(opts *bind.TransactOpts, to common.Address, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _App.contract.Transact(opts, "makeCard", to, tokenId, amount)
}

// MakeCard is a paid mutator transaction binding the contract method 0xb33bbe6d.
//
// Solidity: function makeCard(address to, uint256 tokenId, uint256 amount) returns()
func (_App *AppSession) MakeCard(to common.Address, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _App.Contract.MakeCard(&_App.TransactOpts, to, tokenId, amount)
}

// MakeCard is a paid mutator transaction binding the contract method 0xb33bbe6d.
//
// Solidity: function makeCard(address to, uint256 tokenId, uint256 amount) returns()
func (_App *AppTransactorSession) MakeCard(to common.Address, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _App.Contract.MakeCard(&_App.TransactOpts, to, tokenId, amount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_App *AppTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _App.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_App *AppSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _App.Contract.RenounceRole(&_App.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_App *AppTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _App.Contract.RenounceRole(&_App.TransactOpts, role, account)
}

// RequestForceWithdraw is a paid mutator transaction binding the contract method 0xda60d27e.
//
// Solidity: function requestForceWithdraw() returns()
func (_App *AppTransactor) RequestForceWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _App.contract.Transact(opts, "requestForceWithdraw")
}

// RequestForceWithdraw is a paid mutator transaction binding the contract method 0xda60d27e.
//
// Solidity: function requestForceWithdraw() returns()
func (_App *AppSession) RequestForceWithdraw() (*types.Transaction, error) {
	return _App.Contract.RequestForceWithdraw(&_App.TransactOpts)
}

// RequestForceWithdraw is a paid mutator transaction binding the contract method 0xda60d27e.
//
// Solidity: function requestForceWithdraw() returns()
func (_App *AppTransactorSession) RequestForceWithdraw() (*types.Transaction, error) {
	return _App.Contract.RequestForceWithdraw(&_App.TransactOpts)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_App *AppTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _App.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_App *AppSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _App.Contract.RevokeRole(&_App.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_App *AppTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _App.Contract.RevokeRole(&_App.TransactOpts, role, account)
}

// SetProps is a paid mutator transaction binding the contract method 0xeec44255.
//
// Solidity: function setProps(address cardShop_, string link_, string description_, uint8 paymentType_) returns()
func (_App *AppTransactor) SetProps(opts *bind.TransactOpts, cardShop_ common.Address, link_ string, description_ string, paymentType_ uint8) (*types.Transaction, error) {
	return _App.contract.Transact(opts, "setProps", cardShop_, link_, description_, paymentType_)
}

// SetProps is a paid mutator transaction binding the contract method 0xeec44255.
//
// Solidity: function setProps(address cardShop_, string link_, string description_, uint8 paymentType_) returns()
func (_App *AppSession) SetProps(cardShop_ common.Address, link_ string, description_ string, paymentType_ uint8) (*types.Transaction, error) {
	return _App.Contract.SetProps(&_App.TransactOpts, cardShop_, link_, description_, paymentType_)
}

// SetProps is a paid mutator transaction binding the contract method 0xeec44255.
//
// Solidity: function setProps(address cardShop_, string link_, string description_, uint8 paymentType_) returns()
func (_App *AppTransactorSession) SetProps(cardShop_ common.Address, link_ string, description_ string, paymentType_ uint8) (*types.Transaction, error) {
	return _App.Contract.SetProps(&_App.TransactOpts, cardShop_, link_, description_, paymentType_)
}

// TakeProfit is a paid mutator transaction binding the contract method 0x0cec2f17.
//
// Solidity: function takeProfit(address to, uint256 amount) returns()
func (_App *AppTransactor) TakeProfit(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _App.contract.Transact(opts, "takeProfit", to, amount)
}

// TakeProfit is a paid mutator transaction binding the contract method 0x0cec2f17.
//
// Solidity: function takeProfit(address to, uint256 amount) returns()
func (_App *AppSession) TakeProfit(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _App.Contract.TakeProfit(&_App.TransactOpts, to, amount)
}

// TakeProfit is a paid mutator transaction binding the contract method 0x0cec2f17.
//
// Solidity: function takeProfit(address to, uint256 amount) returns()
func (_App *AppTransactorSession) TakeProfit(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _App.Contract.TakeProfit(&_App.TransactOpts, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5437e401.
//
// Solidity: function withdraw(address account, bool toAssets) returns()
func (_App *AppTransactor) Withdraw(opts *bind.TransactOpts, account common.Address, toAssets bool) (*types.Transaction, error) {
	return _App.contract.Transact(opts, "withdraw", account, toAssets)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5437e401.
//
// Solidity: function withdraw(address account, bool toAssets) returns()
func (_App *AppSession) Withdraw(account common.Address, toAssets bool) (*types.Transaction, error) {
	return _App.Contract.Withdraw(&_App.TransactOpts, account, toAssets)
}

// Withdraw is a paid mutator transaction binding the contract method 0x5437e401.
//
// Solidity: function withdraw(address account, bool toAssets) returns()
func (_App *AppTransactorSession) Withdraw(account common.Address, toAssets bool) (*types.Transaction, error) {
	return _App.Contract.Withdraw(&_App.TransactOpts, account, toAssets)
}

// AppDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the App contract.
type AppDepositIterator struct {
	Event *AppDeposit // Event containing the contract specifics and raw log

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
func (it *AppDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppDeposit)
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
		it.Event = new(AppDeposit)
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
func (it *AppDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppDeposit represents a Deposit event raised by the App contract.
type AppDeposit struct {
	Operator common.Address
	Receiver common.Address
	TokenId  *big.Int
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed operator, address indexed receiver, uint256 indexed tokenId, uint256 amount)
func (_App *AppFilterer) FilterDeposit(opts *bind.FilterOpts, operator []common.Address, receiver []common.Address, tokenId []*big.Int) (*AppDepositIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _App.contract.FilterLogs(opts, "Deposit", operatorRule, receiverRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AppDepositIterator{contract: _App.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed operator, address indexed receiver, uint256 indexed tokenId, uint256 amount)
func (_App *AppFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *AppDeposit, operator []common.Address, receiver []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _App.contract.WatchLogs(opts, "Deposit", operatorRule, receiverRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppDeposit)
				if err := _App.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed operator, address indexed receiver, uint256 indexed tokenId, uint256 amount)
func (_App *AppFilterer) ParseDeposit(log types.Log) (*AppDeposit, error) {
	event := new(AppDeposit)
	if err := _App.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppFrozenIterator is returned from FilterFrozen and is used to iterate over the raw logs and unpacked data for Frozen events raised by the App contract.
type AppFrozenIterator struct {
	Event *AppFrozen // Event containing the contract specifics and raw log

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
func (it *AppFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppFrozen)
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
		it.Event = new(AppFrozen)
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
func (it *AppFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppFrozen represents a Frozen event raised by the App contract.
type AppFrozen struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFrozen is a free log retrieval operation binding the contract event 0x8a5c4736a33c7b7f29a2c34ea9ff9608afc5718d56f6fd6dcbd2d3711a1a4913.
//
// Solidity: event Frozen(address indexed account)
func (_App *AppFilterer) FilterFrozen(opts *bind.FilterOpts, account []common.Address) (*AppFrozenIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _App.contract.FilterLogs(opts, "Frozen", accountRule)
	if err != nil {
		return nil, err
	}
	return &AppFrozenIterator{contract: _App.contract, event: "Frozen", logs: logs, sub: sub}, nil
}

// WatchFrozen is a free log subscription operation binding the contract event 0x8a5c4736a33c7b7f29a2c34ea9ff9608afc5718d56f6fd6dcbd2d3711a1a4913.
//
// Solidity: event Frozen(address indexed account)
func (_App *AppFilterer) WatchFrozen(opts *bind.WatchOpts, sink chan<- *AppFrozen, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _App.contract.WatchLogs(opts, "Frozen", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppFrozen)
				if err := _App.contract.UnpackLog(event, "Frozen", log); err != nil {
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
// Solidity: event Frozen(address indexed account)
func (_App *AppFilterer) ParseFrozen(log types.Log) (*AppFrozen, error) {
	event := new(AppFrozen)
	if err := _App.contract.UnpackLog(event, "Frozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the App contract.
type AppInitializedIterator struct {
	Event *AppInitialized // Event containing the contract specifics and raw log

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
func (it *AppInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppInitialized)
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
		it.Event = new(AppInitialized)
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
func (it *AppInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppInitialized represents a Initialized event raised by the App contract.
type AppInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_App *AppFilterer) FilterInitialized(opts *bind.FilterOpts) (*AppInitializedIterator, error) {

	logs, sub, err := _App.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AppInitializedIterator{contract: _App.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_App *AppFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AppInitialized) (event.Subscription, error) {

	logs, sub, err := _App.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppInitialized)
				if err := _App.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_App *AppFilterer) ParseInitialized(log types.Log) (*AppInitialized, error) {
	event := new(AppInitialized)
	if err := _App.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the App contract.
type AppRoleAdminChangedIterator struct {
	Event *AppRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AppRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppRoleAdminChanged)
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
		it.Event = new(AppRoleAdminChanged)
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
func (it *AppRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppRoleAdminChanged represents a RoleAdminChanged event raised by the App contract.
type AppRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_App *AppFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AppRoleAdminChangedIterator, error) {

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

	logs, sub, err := _App.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AppRoleAdminChangedIterator{contract: _App.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_App *AppFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AppRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _App.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppRoleAdminChanged)
				if err := _App.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_App *AppFilterer) ParseRoleAdminChanged(log types.Log) (*AppRoleAdminChanged, error) {
	event := new(AppRoleAdminChanged)
	if err := _App.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the App contract.
type AppRoleGrantedIterator struct {
	Event *AppRoleGranted // Event containing the contract specifics and raw log

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
func (it *AppRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppRoleGranted)
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
		it.Event = new(AppRoleGranted)
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
func (it *AppRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppRoleGranted represents a RoleGranted event raised by the App contract.
type AppRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_App *AppFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AppRoleGrantedIterator, error) {

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

	logs, sub, err := _App.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AppRoleGrantedIterator{contract: _App.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_App *AppFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AppRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _App.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppRoleGranted)
				if err := _App.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_App *AppFilterer) ParseRoleGranted(log types.Log) (*AppRoleGranted, error) {
	event := new(AppRoleGranted)
	if err := _App.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the App contract.
type AppRoleRevokedIterator struct {
	Event *AppRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AppRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppRoleRevoked)
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
		it.Event = new(AppRoleRevoked)
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
func (it *AppRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppRoleRevoked represents a RoleRevoked event raised by the App contract.
type AppRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_App *AppFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AppRoleRevokedIterator, error) {

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

	logs, sub, err := _App.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AppRoleRevokedIterator{contract: _App.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_App *AppFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AppRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _App.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppRoleRevoked)
				if err := _App.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_App *AppFilterer) ParseRoleRevoked(log types.Log) (*AppRoleRevoked, error) {
	event := new(AppRoleRevoked)
	if err := _App.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the App contract.
type AppWithdrawIterator struct {
	Event *AppWithdraw // Event containing the contract specifics and raw log

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
func (it *AppWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppWithdraw)
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
		it.Event = new(AppWithdraw)
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
func (it *AppWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppWithdraw represents a Withdraw event raised by the App contract.
type AppWithdraw struct {
	Operator common.Address
	Account  common.Address
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x3115d1449a7b732c986cba18244e897a450f61e1bb8d589cd2e69e6c8924f9f7.
//
// Solidity: event Withdraw(address indexed operator, address account, address indexed receiver, uint256 amount)
func (_App *AppFilterer) FilterWithdraw(opts *bind.FilterOpts, operator []common.Address, receiver []common.Address) (*AppWithdrawIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _App.contract.FilterLogs(opts, "Withdraw", operatorRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &AppWithdrawIterator{contract: _App.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x3115d1449a7b732c986cba18244e897a450f61e1bb8d589cd2e69e6c8924f9f7.
//
// Solidity: event Withdraw(address indexed operator, address account, address indexed receiver, uint256 amount)
func (_App *AppFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *AppWithdraw, operator []common.Address, receiver []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _App.contract.WatchLogs(opts, "Withdraw", operatorRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppWithdraw)
				if err := _App.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x3115d1449a7b732c986cba18244e897a450f61e1bb8d589cd2e69e6c8924f9f7.
//
// Solidity: event Withdraw(address indexed operator, address account, address indexed receiver, uint256 amount)
func (_App *AppFilterer) ParseWithdraw(log types.Log) (*AppWithdraw, error) {
	event := new(AppWithdraw)
	if err := _App.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
