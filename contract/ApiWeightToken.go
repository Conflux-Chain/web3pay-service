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

// AppConfigConfigRequest is an auto generated low-level Go binding around an user-defined struct.
type AppConfigConfigRequest struct {
	Id         uint32
	ResourceId string
	Weight     *big.Int
	Op         uint8
}

// IAppConfigConfigEntry is an auto generated low-level Go binding around an user-defined struct.
type IAppConfigConfigEntry struct {
	ResourceId    string
	Weight        *big.Int
	Index         uint32
	PendingOP     uint8
	PendingWeight *big.Int
	SubmitSeconds *big.Int
	RequestTimes  *big.Int
}

// IAppConfigResourceUseDetail is an auto generated low-level Go binding around an user-defined struct.
type IAppConfigResourceUseDetail struct {
	Id    uint32
	Times *big.Int
}

// ApiWeightTokenMetaData contains all meta data concerning the ApiWeightToken contract.
var ApiWeightTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIApp\",\"name\":\"belongsTo\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"enumIAppConfig.OP\",\"name\":\"op\",\"type\":\"uint8\"}],\"name\":\"ResourceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newWeight\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"enumIAppConfig.OP\",\"name\":\"op\",\"type\":\"uint8\"}],\"name\":\"ResourcePending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AIRDROP_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BILLING_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FIRST_CONFIG_ID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FT_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TAKE_PROFIT_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_AIRDROP_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"times\",\"type\":\"uint256\"}],\"internalType\":\"structIAppConfig.ResourceUseDetail[]\",\"name\":\"useDetail\",\"type\":\"tuple[]\"}],\"name\":\"addRequestTimes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"belongsToApp\",\"outputs\":[{\"internalType\":\"contractIApp\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"resourceId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"enumIAppConfig.OP\",\"name\":\"op\",\"type\":\"uint8\"}],\"internalType\":\"structAppConfig.ConfigRequest\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"configResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"resourceId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"enumIAppConfig.OP\",\"name\":\"op\",\"type\":\"uint8\"}],\"internalType\":\"structAppConfig.ConfigRequest[]\",\"name\":\"entries\",\"type\":\"tuple[]\"}],\"name\":\"configResourceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flushPendingConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"indexArray\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIApp\",\"name\":\"belongsTo\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"defaultWeight\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"listResources\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"resourceId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"enumIAppConfig.OP\",\"name\":\"pendingOP\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"pendingWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"submitSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestTimes\",\"type\":\"uint256\"}],\"internalType\":\"structIAppConfig.ConfigEntry[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint32[]\",\"name\":\"ids\",\"type\":\"uint32[]\"}],\"name\":\"listUserRequestCounter\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"times\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextConfigId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingIdArray\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"resourceConfigures\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"resourceId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"enumIAppConfig.OP\",\"name\":\"pendingOP\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"pendingWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"submitSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestTimes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"seconds_\",\"type\":\"uint256\"}],\"name\":\"setPendingSeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ApiWeightTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiWeightTokenMetaData.ABI instead.
var ApiWeightTokenABI = ApiWeightTokenMetaData.ABI

// ApiWeightToken is an auto generated Go binding around an Ethereum contract.
type ApiWeightToken struct {
	ApiWeightTokenCaller     // Read-only binding to the contract
	ApiWeightTokenTransactor // Write-only binding to the contract
	ApiWeightTokenFilterer   // Log filterer for contract events
}

// ApiWeightTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiWeightTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiWeightTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiWeightTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiWeightTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiWeightTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiWeightTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiWeightTokenSession struct {
	Contract     *ApiWeightToken   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiWeightTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiWeightTokenCallerSession struct {
	Contract *ApiWeightTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ApiWeightTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiWeightTokenTransactorSession struct {
	Contract     *ApiWeightTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ApiWeightTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiWeightTokenRaw struct {
	Contract *ApiWeightToken // Generic contract binding to access the raw methods on
}

// ApiWeightTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiWeightTokenCallerRaw struct {
	Contract *ApiWeightTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ApiWeightTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiWeightTokenTransactorRaw struct {
	Contract *ApiWeightTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApiWeightToken creates a new instance of ApiWeightToken, bound to a specific deployed contract.
func NewApiWeightToken(address common.Address, backend bind.ContractBackend) (*ApiWeightToken, error) {
	contract, err := bindApiWeightToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ApiWeightToken{ApiWeightTokenCaller: ApiWeightTokenCaller{contract: contract}, ApiWeightTokenTransactor: ApiWeightTokenTransactor{contract: contract}, ApiWeightTokenFilterer: ApiWeightTokenFilterer{contract: contract}}, nil
}

// NewApiWeightTokenCaller creates a new read-only instance of ApiWeightToken, bound to a specific deployed contract.
func NewApiWeightTokenCaller(address common.Address, caller bind.ContractCaller) (*ApiWeightTokenCaller, error) {
	contract, err := bindApiWeightToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiWeightTokenCaller{contract: contract}, nil
}

// NewApiWeightTokenTransactor creates a new write-only instance of ApiWeightToken, bound to a specific deployed contract.
func NewApiWeightTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiWeightTokenTransactor, error) {
	contract, err := bindApiWeightToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiWeightTokenTransactor{contract: contract}, nil
}

// NewApiWeightTokenFilterer creates a new log filterer instance of ApiWeightToken, bound to a specific deployed contract.
func NewApiWeightTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiWeightTokenFilterer, error) {
	contract, err := bindApiWeightToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiWeightTokenFilterer{contract: contract}, nil
}

// bindApiWeightToken binds a generic wrapper to an already deployed contract.
func bindApiWeightToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiWeightTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ApiWeightToken *ApiWeightTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ApiWeightToken.Contract.ApiWeightTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ApiWeightToken *ApiWeightTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ApiWeightToken.Contract.ApiWeightTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ApiWeightToken *ApiWeightTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ApiWeightToken.Contract.ApiWeightTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ApiWeightToken *ApiWeightTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ApiWeightToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ApiWeightToken *ApiWeightTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ApiWeightToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ApiWeightToken *ApiWeightTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ApiWeightToken.Contract.contract.Transact(opts, method, params...)
}

// AIRDROPID is a free data retrieval call binding the contract method 0xe0f194f1.
//
// Solidity: function AIRDROP_ID() view returns(uint256)
func (_ApiWeightToken *ApiWeightTokenCaller) AIRDROPID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ApiWeightToken.contract.Call(opts, &out, "AIRDROP_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AIRDROPID is a free data retrieval call binding the contract method 0xe0f194f1.
//
// Solidity: function AIRDROP_ID() view returns(uint256)
func (_ApiWeightToken *ApiWeightTokenSession) AIRDROPID() (*big.Int, error) {
	return _ApiWeightToken.Contract.AIRDROPID(&_ApiWeightToken.CallOpts)
}

// AIRDROPID is a free data retrieval call binding the contract method 0xe0f194f1.
//
// Solidity: function AIRDROP_ID() view returns(uint256)
func (_ApiWeightToken *ApiWeightTokenCallerSession) AIRDROPID() (*big.Int, error) {
	return _ApiWeightToken.Contract.AIRDROPID(&_ApiWeightToken.CallOpts)
}

// BILLINGID is a free data retrieval call binding the contract method 0x1c91d1ba.
//
// Solidity: function BILLING_ID() view returns(uint256)
func (_ApiWeightToken *ApiWeightTokenCaller) BILLINGID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ApiWeightToken.contract.Call(opts, &out, "BILLING_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BILLINGID is a free data retrieval call binding the contract method 0x1c91d1ba.
//
// Solidity: function BILLING_ID() view returns(uint256)
func (_ApiWeightToken *ApiWeightTokenSession) BILLINGID() (*big.Int, error) {
	return _ApiWeightToken.Contract.BILLINGID(&_ApiWeightToken.CallOpts)
}

// BILLINGID is a free data retrieval call binding the contract method 0x1c91d1ba.
//
// Solidity: function BILLING_ID() view returns(uint256)
func (_ApiWeightToken *ApiWeightTokenCallerSession) BILLINGID() (*big.Int, error) {
	return _ApiWeightToken.Contract.BILLINGID(&_ApiWeightToken.CallOpts)
}

// FIRSTCONFIGID is a free data retrieval call binding the contract method 0x0fe58201.
//
// Solidity: function FIRST_CONFIG_ID() view returns(uint32)
func (_ApiWeightToken *ApiWeightTokenCaller) FIRSTCONFIGID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ApiWeightToken.contract.Call(opts, &out, "FIRST_CONFIG_ID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// FIRSTCONFIGID is a free data retrieval call binding the contract method 0x0fe58201.
//
// Solidity: function FIRST_CONFIG_ID() view returns(uint32)
func (_ApiWeightToken *ApiWeightTokenSession) FIRSTCONFIGID() (uint32, error) {
	return _ApiWeightToken.Contract.FIRSTCONFIGID(&_ApiWeightToken.CallOpts)
}

// FIRSTCONFIGID is a free data retrieval call binding the contract method 0x0fe58201.
//
// Solidity: function FIRST_CONFIG_ID() view returns(uint32)
func (_ApiWeightToken *ApiWeightTokenCallerSession) FIRSTCONFIGID() (uint32, error) {
	return _ApiWeightToken.Contract.FIRSTCONFIGID(&_ApiWeightToken.CallOpts)
}

// FTID is a free data retrieval call binding the contract method 0x01052d57.
//
// Solidity: function FT_ID() view returns(uint256)
func (_ApiWeightToken *ApiWeightTokenCaller) FTID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ApiWeightToken.contract.Call(opts, &out, "FT_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FTID is a free data retrieval call binding the contract method 0x01052d57.
//
// Solidity: function FT_ID() view returns(uint256)
func (_ApiWeightToken *ApiWeightTokenSession) FTID() (*big.Int, error) {
	return _ApiWeightToken.Contract.FTID(&_ApiWeightToken.CallOpts)
}

// FTID is a free data retrieval call binding the contract method 0x01052d57.
//
// Solidity: function FT_ID() view returns(uint256)
func (_ApiWeightToken *ApiWeightTokenCallerSession) FTID() (*big.Int, error) {
	return _ApiWeightToken.Contract.FTID(&_ApiWeightToken.CallOpts)
}

// TAKEPROFITID is a free data retrieval call binding the contract method 0xfdd62e1c.
//
// Solidity: function TAKE_PROFIT_ID() view returns(uint256)
func (_ApiWeightToken *ApiWeightTokenCaller) TAKEPROFITID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ApiWeightToken.contract.Call(opts, &out, "TAKE_PROFIT_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TAKEPROFITID is a free data retrieval call binding the contract method 0xfdd62e1c.
//
// Solidity: function TAKE_PROFIT_ID() view returns(uint256)
func (_ApiWeightToken *ApiWeightTokenSession) TAKEPROFITID() (*big.Int, error) {
	return _ApiWeightToken.Contract.TAKEPROFITID(&_ApiWeightToken.CallOpts)
}

// TAKEPROFITID is a free data retrieval call binding the contract method 0xfdd62e1c.
//
// Solidity: function TAKE_PROFIT_ID() view returns(uint256)
func (_ApiWeightToken *ApiWeightTokenCallerSession) TAKEPROFITID() (*big.Int, error) {
	return _ApiWeightToken.Contract.TAKEPROFITID(&_ApiWeightToken.CallOpts)
}

// TOKENAIRDROPID is a free data retrieval call binding the contract method 0x25e0ba46.
//
// Solidity: function TOKEN_AIRDROP_ID() view returns(uint256)
func (_ApiWeightToken *ApiWeightTokenCaller) TOKENAIRDROPID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ApiWeightToken.contract.Call(opts, &out, "TOKEN_AIRDROP_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOKENAIRDROPID is a free data retrieval call binding the contract method 0x25e0ba46.
//
// Solidity: function TOKEN_AIRDROP_ID() view returns(uint256)
func (_ApiWeightToken *ApiWeightTokenSession) TOKENAIRDROPID() (*big.Int, error) {
	return _ApiWeightToken.Contract.TOKENAIRDROPID(&_ApiWeightToken.CallOpts)
}

// TOKENAIRDROPID is a free data retrieval call binding the contract method 0x25e0ba46.
//
// Solidity: function TOKEN_AIRDROP_ID() view returns(uint256)
func (_ApiWeightToken *ApiWeightTokenCallerSession) TOKENAIRDROPID() (*big.Int, error) {
	return _ApiWeightToken.Contract.TOKENAIRDROPID(&_ApiWeightToken.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ApiWeightToken *ApiWeightTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ApiWeightToken.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ApiWeightToken *ApiWeightTokenSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _ApiWeightToken.Contract.BalanceOf(&_ApiWeightToken.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ApiWeightToken *ApiWeightTokenCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _ApiWeightToken.Contract.BalanceOf(&_ApiWeightToken.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ApiWeightToken *ApiWeightTokenCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ApiWeightToken.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ApiWeightToken *ApiWeightTokenSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ApiWeightToken.Contract.BalanceOfBatch(&_ApiWeightToken.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ApiWeightToken *ApiWeightTokenCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ApiWeightToken.Contract.BalanceOfBatch(&_ApiWeightToken.CallOpts, accounts, ids)
}

// BelongsToApp is a free data retrieval call binding the contract method 0x53f63ca6.
//
// Solidity: function belongsToApp() view returns(address)
func (_ApiWeightToken *ApiWeightTokenCaller) BelongsToApp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ApiWeightToken.contract.Call(opts, &out, "belongsToApp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BelongsToApp is a free data retrieval call binding the contract method 0x53f63ca6.
//
// Solidity: function belongsToApp() view returns(address)
func (_ApiWeightToken *ApiWeightTokenSession) BelongsToApp() (common.Address, error) {
	return _ApiWeightToken.Contract.BelongsToApp(&_ApiWeightToken.CallOpts)
}

// BelongsToApp is a free data retrieval call binding the contract method 0x53f63ca6.
//
// Solidity: function belongsToApp() view returns(address)
func (_ApiWeightToken *ApiWeightTokenCallerSession) BelongsToApp() (common.Address, error) {
	return _ApiWeightToken.Contract.BelongsToApp(&_ApiWeightToken.CallOpts)
}

// IndexArray is a free data retrieval call binding the contract method 0xebed9b2c.
//
// Solidity: function indexArray(uint256 ) view returns(uint32)
func (_ApiWeightToken *ApiWeightTokenCaller) IndexArray(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var out []interface{}
	err := _ApiWeightToken.contract.Call(opts, &out, "indexArray", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// IndexArray is a free data retrieval call binding the contract method 0xebed9b2c.
//
// Solidity: function indexArray(uint256 ) view returns(uint32)
func (_ApiWeightToken *ApiWeightTokenSession) IndexArray(arg0 *big.Int) (uint32, error) {
	return _ApiWeightToken.Contract.IndexArray(&_ApiWeightToken.CallOpts, arg0)
}

// IndexArray is a free data retrieval call binding the contract method 0xebed9b2c.
//
// Solidity: function indexArray(uint256 ) view returns(uint32)
func (_ApiWeightToken *ApiWeightTokenCallerSession) IndexArray(arg0 *big.Int) (uint32, error) {
	return _ApiWeightToken.Contract.IndexArray(&_ApiWeightToken.CallOpts, arg0)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ApiWeightToken *ApiWeightTokenCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ApiWeightToken.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ApiWeightToken *ApiWeightTokenSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _ApiWeightToken.Contract.IsApprovedForAll(&_ApiWeightToken.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ApiWeightToken *ApiWeightTokenCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _ApiWeightToken.Contract.IsApprovedForAll(&_ApiWeightToken.CallOpts, account, operator)
}

// ListResources is a free data retrieval call binding the contract method 0x49459cd2.
//
// Solidity: function listResources(uint256 offset, uint256 limit) view returns((string,uint256,uint32,uint8,uint256,uint256,uint256)[], uint256 total)
func (_ApiWeightToken *ApiWeightTokenCaller) ListResources(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]IAppConfigConfigEntry, *big.Int, error) {
	var out []interface{}
	err := _ApiWeightToken.contract.Call(opts, &out, "listResources", offset, limit)

	if err != nil {
		return *new([]IAppConfigConfigEntry), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]IAppConfigConfigEntry)).(*[]IAppConfigConfigEntry)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// ListResources is a free data retrieval call binding the contract method 0x49459cd2.
//
// Solidity: function listResources(uint256 offset, uint256 limit) view returns((string,uint256,uint32,uint8,uint256,uint256,uint256)[], uint256 total)
func (_ApiWeightToken *ApiWeightTokenSession) ListResources(offset *big.Int, limit *big.Int) ([]IAppConfigConfigEntry, *big.Int, error) {
	return _ApiWeightToken.Contract.ListResources(&_ApiWeightToken.CallOpts, offset, limit)
}

// ListResources is a free data retrieval call binding the contract method 0x49459cd2.
//
// Solidity: function listResources(uint256 offset, uint256 limit) view returns((string,uint256,uint32,uint8,uint256,uint256,uint256)[], uint256 total)
func (_ApiWeightToken *ApiWeightTokenCallerSession) ListResources(offset *big.Int, limit *big.Int) ([]IAppConfigConfigEntry, *big.Int, error) {
	return _ApiWeightToken.Contract.ListResources(&_ApiWeightToken.CallOpts, offset, limit)
}

// ListUserRequestCounter is a free data retrieval call binding the contract method 0xa5e34731.
//
// Solidity: function listUserRequestCounter(address user, uint32[] ids) view returns(uint256[] times)
func (_ApiWeightToken *ApiWeightTokenCaller) ListUserRequestCounter(opts *bind.CallOpts, user common.Address, ids []uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ApiWeightToken.contract.Call(opts, &out, "listUserRequestCounter", user, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ListUserRequestCounter is a free data retrieval call binding the contract method 0xa5e34731.
//
// Solidity: function listUserRequestCounter(address user, uint32[] ids) view returns(uint256[] times)
func (_ApiWeightToken *ApiWeightTokenSession) ListUserRequestCounter(user common.Address, ids []uint32) ([]*big.Int, error) {
	return _ApiWeightToken.Contract.ListUserRequestCounter(&_ApiWeightToken.CallOpts, user, ids)
}

// ListUserRequestCounter is a free data retrieval call binding the contract method 0xa5e34731.
//
// Solidity: function listUserRequestCounter(address user, uint32[] ids) view returns(uint256[] times)
func (_ApiWeightToken *ApiWeightTokenCallerSession) ListUserRequestCounter(user common.Address, ids []uint32) ([]*big.Int, error) {
	return _ApiWeightToken.Contract.ListUserRequestCounter(&_ApiWeightToken.CallOpts, user, ids)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ApiWeightToken *ApiWeightTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ApiWeightToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ApiWeightToken *ApiWeightTokenSession) Name() (string, error) {
	return _ApiWeightToken.Contract.Name(&_ApiWeightToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ApiWeightToken *ApiWeightTokenCallerSession) Name() (string, error) {
	return _ApiWeightToken.Contract.Name(&_ApiWeightToken.CallOpts)
}

// NextConfigId is a free data retrieval call binding the contract method 0x99d726c7.
//
// Solidity: function nextConfigId() view returns(uint32)
func (_ApiWeightToken *ApiWeightTokenCaller) NextConfigId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ApiWeightToken.contract.Call(opts, &out, "nextConfigId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NextConfigId is a free data retrieval call binding the contract method 0x99d726c7.
//
// Solidity: function nextConfigId() view returns(uint32)
func (_ApiWeightToken *ApiWeightTokenSession) NextConfigId() (uint32, error) {
	return _ApiWeightToken.Contract.NextConfigId(&_ApiWeightToken.CallOpts)
}

// NextConfigId is a free data retrieval call binding the contract method 0x99d726c7.
//
// Solidity: function nextConfigId() view returns(uint32)
func (_ApiWeightToken *ApiWeightTokenCallerSession) NextConfigId() (uint32, error) {
	return _ApiWeightToken.Contract.NextConfigId(&_ApiWeightToken.CallOpts)
}

// PendingIdArray is a free data retrieval call binding the contract method 0xac8ef2b5.
//
// Solidity: function pendingIdArray(uint256 ) view returns(uint32)
func (_ApiWeightToken *ApiWeightTokenCaller) PendingIdArray(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var out []interface{}
	err := _ApiWeightToken.contract.Call(opts, &out, "pendingIdArray", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// PendingIdArray is a free data retrieval call binding the contract method 0xac8ef2b5.
//
// Solidity: function pendingIdArray(uint256 ) view returns(uint32)
func (_ApiWeightToken *ApiWeightTokenSession) PendingIdArray(arg0 *big.Int) (uint32, error) {
	return _ApiWeightToken.Contract.PendingIdArray(&_ApiWeightToken.CallOpts, arg0)
}

// PendingIdArray is a free data retrieval call binding the contract method 0xac8ef2b5.
//
// Solidity: function pendingIdArray(uint256 ) view returns(uint32)
func (_ApiWeightToken *ApiWeightTokenCallerSession) PendingIdArray(arg0 *big.Int) (uint32, error) {
	return _ApiWeightToken.Contract.PendingIdArray(&_ApiWeightToken.CallOpts, arg0)
}

// PendingSeconds is a free data retrieval call binding the contract method 0xc819c70e.
//
// Solidity: function pendingSeconds() view returns(uint256)
func (_ApiWeightToken *ApiWeightTokenCaller) PendingSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ApiWeightToken.contract.Call(opts, &out, "pendingSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingSeconds is a free data retrieval call binding the contract method 0xc819c70e.
//
// Solidity: function pendingSeconds() view returns(uint256)
func (_ApiWeightToken *ApiWeightTokenSession) PendingSeconds() (*big.Int, error) {
	return _ApiWeightToken.Contract.PendingSeconds(&_ApiWeightToken.CallOpts)
}

// PendingSeconds is a free data retrieval call binding the contract method 0xc819c70e.
//
// Solidity: function pendingSeconds() view returns(uint256)
func (_ApiWeightToken *ApiWeightTokenCallerSession) PendingSeconds() (*big.Int, error) {
	return _ApiWeightToken.Contract.PendingSeconds(&_ApiWeightToken.CallOpts)
}

// ResourceConfigures is a free data retrieval call binding the contract method 0xc86b6ee2.
//
// Solidity: function resourceConfigures(uint32 ) view returns(string resourceId, uint256 weight, uint32 index, uint8 pendingOP, uint256 pendingWeight, uint256 submitSeconds, uint256 requestTimes)
func (_ApiWeightToken *ApiWeightTokenCaller) ResourceConfigures(opts *bind.CallOpts, arg0 uint32) (struct {
	ResourceId    string
	Weight        *big.Int
	Index         uint32
	PendingOP     uint8
	PendingWeight *big.Int
	SubmitSeconds *big.Int
	RequestTimes  *big.Int
}, error) {
	var out []interface{}
	err := _ApiWeightToken.contract.Call(opts, &out, "resourceConfigures", arg0)

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
func (_ApiWeightToken *ApiWeightTokenSession) ResourceConfigures(arg0 uint32) (struct {
	ResourceId    string
	Weight        *big.Int
	Index         uint32
	PendingOP     uint8
	PendingWeight *big.Int
	SubmitSeconds *big.Int
	RequestTimes  *big.Int
}, error) {
	return _ApiWeightToken.Contract.ResourceConfigures(&_ApiWeightToken.CallOpts, arg0)
}

// ResourceConfigures is a free data retrieval call binding the contract method 0xc86b6ee2.
//
// Solidity: function resourceConfigures(uint32 ) view returns(string resourceId, uint256 weight, uint32 index, uint8 pendingOP, uint256 pendingWeight, uint256 submitSeconds, uint256 requestTimes)
func (_ApiWeightToken *ApiWeightTokenCallerSession) ResourceConfigures(arg0 uint32) (struct {
	ResourceId    string
	Weight        *big.Int
	Index         uint32
	PendingOP     uint8
	PendingWeight *big.Int
	SubmitSeconds *big.Int
	RequestTimes  *big.Int
}, error) {
	return _ApiWeightToken.Contract.ResourceConfigures(&_ApiWeightToken.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ApiWeightToken *ApiWeightTokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ApiWeightToken.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ApiWeightToken *ApiWeightTokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ApiWeightToken.Contract.SupportsInterface(&_ApiWeightToken.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ApiWeightToken *ApiWeightTokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ApiWeightToken.Contract.SupportsInterface(&_ApiWeightToken.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ApiWeightToken *ApiWeightTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ApiWeightToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ApiWeightToken *ApiWeightTokenSession) Symbol() (string, error) {
	return _ApiWeightToken.Contract.Symbol(&_ApiWeightToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ApiWeightToken *ApiWeightTokenCallerSession) Symbol() (string, error) {
	return _ApiWeightToken.Contract.Symbol(&_ApiWeightToken.CallOpts)
}

// TotalRequests is a free data retrieval call binding the contract method 0x8aea61dc.
//
// Solidity: function totalRequests() view returns(uint256)
func (_ApiWeightToken *ApiWeightTokenCaller) TotalRequests(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ApiWeightToken.contract.Call(opts, &out, "totalRequests")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRequests is a free data retrieval call binding the contract method 0x8aea61dc.
//
// Solidity: function totalRequests() view returns(uint256)
func (_ApiWeightToken *ApiWeightTokenSession) TotalRequests() (*big.Int, error) {
	return _ApiWeightToken.Contract.TotalRequests(&_ApiWeightToken.CallOpts)
}

// TotalRequests is a free data retrieval call binding the contract method 0x8aea61dc.
//
// Solidity: function totalRequests() view returns(uint256)
func (_ApiWeightToken *ApiWeightTokenCallerSession) TotalRequests() (*big.Int, error) {
	return _ApiWeightToken.Contract.TotalRequests(&_ApiWeightToken.CallOpts)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_ApiWeightToken *ApiWeightTokenCaller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _ApiWeightToken.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_ApiWeightToken *ApiWeightTokenSession) Uri(arg0 *big.Int) (string, error) {
	return _ApiWeightToken.Contract.Uri(&_ApiWeightToken.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_ApiWeightToken *ApiWeightTokenCallerSession) Uri(arg0 *big.Int) (string, error) {
	return _ApiWeightToken.Contract.Uri(&_ApiWeightToken.CallOpts, arg0)
}

// AddRequestTimes is a paid mutator transaction binding the contract method 0xcd9b856d.
//
// Solidity: function addRequestTimes(address account, (uint32,uint256)[] useDetail) returns()
func (_ApiWeightToken *ApiWeightTokenTransactor) AddRequestTimes(opts *bind.TransactOpts, account common.Address, useDetail []IAppConfigResourceUseDetail) (*types.Transaction, error) {
	return _ApiWeightToken.contract.Transact(opts, "addRequestTimes", account, useDetail)
}

// AddRequestTimes is a paid mutator transaction binding the contract method 0xcd9b856d.
//
// Solidity: function addRequestTimes(address account, (uint32,uint256)[] useDetail) returns()
func (_ApiWeightToken *ApiWeightTokenSession) AddRequestTimes(account common.Address, useDetail []IAppConfigResourceUseDetail) (*types.Transaction, error) {
	return _ApiWeightToken.Contract.AddRequestTimes(&_ApiWeightToken.TransactOpts, account, useDetail)
}

// AddRequestTimes is a paid mutator transaction binding the contract method 0xcd9b856d.
//
// Solidity: function addRequestTimes(address account, (uint32,uint256)[] useDetail) returns()
func (_ApiWeightToken *ApiWeightTokenTransactorSession) AddRequestTimes(account common.Address, useDetail []IAppConfigResourceUseDetail) (*types.Transaction, error) {
	return _ApiWeightToken.Contract.AddRequestTimes(&_ApiWeightToken.TransactOpts, account, useDetail)
}

// ConfigResource is a paid mutator transaction binding the contract method 0xf2721b7a.
//
// Solidity: function configResource((uint32,string,uint256,uint8) entry) returns()
func (_ApiWeightToken *ApiWeightTokenTransactor) ConfigResource(opts *bind.TransactOpts, entry AppConfigConfigRequest) (*types.Transaction, error) {
	return _ApiWeightToken.contract.Transact(opts, "configResource", entry)
}

// ConfigResource is a paid mutator transaction binding the contract method 0xf2721b7a.
//
// Solidity: function configResource((uint32,string,uint256,uint8) entry) returns()
func (_ApiWeightToken *ApiWeightTokenSession) ConfigResource(entry AppConfigConfigRequest) (*types.Transaction, error) {
	return _ApiWeightToken.Contract.ConfigResource(&_ApiWeightToken.TransactOpts, entry)
}

// ConfigResource is a paid mutator transaction binding the contract method 0xf2721b7a.
//
// Solidity: function configResource((uint32,string,uint256,uint8) entry) returns()
func (_ApiWeightToken *ApiWeightTokenTransactorSession) ConfigResource(entry AppConfigConfigRequest) (*types.Transaction, error) {
	return _ApiWeightToken.Contract.ConfigResource(&_ApiWeightToken.TransactOpts, entry)
}

// ConfigResourceBatch is a paid mutator transaction binding the contract method 0x10b270aa.
//
// Solidity: function configResourceBatch((uint32,string,uint256,uint8)[] entries) returns()
func (_ApiWeightToken *ApiWeightTokenTransactor) ConfigResourceBatch(opts *bind.TransactOpts, entries []AppConfigConfigRequest) (*types.Transaction, error) {
	return _ApiWeightToken.contract.Transact(opts, "configResourceBatch", entries)
}

// ConfigResourceBatch is a paid mutator transaction binding the contract method 0x10b270aa.
//
// Solidity: function configResourceBatch((uint32,string,uint256,uint8)[] entries) returns()
func (_ApiWeightToken *ApiWeightTokenSession) ConfigResourceBatch(entries []AppConfigConfigRequest) (*types.Transaction, error) {
	return _ApiWeightToken.Contract.ConfigResourceBatch(&_ApiWeightToken.TransactOpts, entries)
}

// ConfigResourceBatch is a paid mutator transaction binding the contract method 0x10b270aa.
//
// Solidity: function configResourceBatch((uint32,string,uint256,uint8)[] entries) returns()
func (_ApiWeightToken *ApiWeightTokenTransactorSession) ConfigResourceBatch(entries []AppConfigConfigRequest) (*types.Transaction, error) {
	return _ApiWeightToken.Contract.ConfigResourceBatch(&_ApiWeightToken.TransactOpts, entries)
}

// FlushPendingConfig is a paid mutator transaction binding the contract method 0xb8f2272a.
//
// Solidity: function flushPendingConfig() returns()
func (_ApiWeightToken *ApiWeightTokenTransactor) FlushPendingConfig(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ApiWeightToken.contract.Transact(opts, "flushPendingConfig")
}

// FlushPendingConfig is a paid mutator transaction binding the contract method 0xb8f2272a.
//
// Solidity: function flushPendingConfig() returns()
func (_ApiWeightToken *ApiWeightTokenSession) FlushPendingConfig() (*types.Transaction, error) {
	return _ApiWeightToken.Contract.FlushPendingConfig(&_ApiWeightToken.TransactOpts)
}

// FlushPendingConfig is a paid mutator transaction binding the contract method 0xb8f2272a.
//
// Solidity: function flushPendingConfig() returns()
func (_ApiWeightToken *ApiWeightTokenTransactorSession) FlushPendingConfig() (*types.Transaction, error) {
	return _ApiWeightToken.Contract.FlushPendingConfig(&_ApiWeightToken.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x9047cd7a.
//
// Solidity: function initialize(address belongsTo, string name, string symbol, string uri, uint256 defaultWeight) returns()
func (_ApiWeightToken *ApiWeightTokenTransactor) Initialize(opts *bind.TransactOpts, belongsTo common.Address, name string, symbol string, uri string, defaultWeight *big.Int) (*types.Transaction, error) {
	return _ApiWeightToken.contract.Transact(opts, "initialize", belongsTo, name, symbol, uri, defaultWeight)
}

// Initialize is a paid mutator transaction binding the contract method 0x9047cd7a.
//
// Solidity: function initialize(address belongsTo, string name, string symbol, string uri, uint256 defaultWeight) returns()
func (_ApiWeightToken *ApiWeightTokenSession) Initialize(belongsTo common.Address, name string, symbol string, uri string, defaultWeight *big.Int) (*types.Transaction, error) {
	return _ApiWeightToken.Contract.Initialize(&_ApiWeightToken.TransactOpts, belongsTo, name, symbol, uri, defaultWeight)
}

// Initialize is a paid mutator transaction binding the contract method 0x9047cd7a.
//
// Solidity: function initialize(address belongsTo, string name, string symbol, string uri, uint256 defaultWeight) returns()
func (_ApiWeightToken *ApiWeightTokenTransactorSession) Initialize(belongsTo common.Address, name string, symbol string, uri string, defaultWeight *big.Int) (*types.Transaction, error) {
	return _ApiWeightToken.Contract.Initialize(&_ApiWeightToken.TransactOpts, belongsTo, name, symbol, uri, defaultWeight)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_ApiWeightToken *ApiWeightTokenTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ApiWeightToken.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_ApiWeightToken *ApiWeightTokenSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ApiWeightToken.Contract.OnERC1155BatchReceived(&_ApiWeightToken.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_ApiWeightToken *ApiWeightTokenTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ApiWeightToken.Contract.OnERC1155BatchReceived(&_ApiWeightToken.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_ApiWeightToken *ApiWeightTokenTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ApiWeightToken.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_ApiWeightToken *ApiWeightTokenSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ApiWeightToken.Contract.OnERC1155Received(&_ApiWeightToken.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_ApiWeightToken *ApiWeightTokenTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ApiWeightToken.Contract.OnERC1155Received(&_ApiWeightToken.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ApiWeightToken *ApiWeightTokenTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _ApiWeightToken.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ApiWeightToken *ApiWeightTokenSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _ApiWeightToken.Contract.SafeBatchTransferFrom(&_ApiWeightToken.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ApiWeightToken *ApiWeightTokenTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _ApiWeightToken.Contract.SafeBatchTransferFrom(&_ApiWeightToken.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ApiWeightToken *ApiWeightTokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ApiWeightToken.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ApiWeightToken *ApiWeightTokenSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ApiWeightToken.Contract.SafeTransferFrom(&_ApiWeightToken.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ApiWeightToken *ApiWeightTokenTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ApiWeightToken.Contract.SafeTransferFrom(&_ApiWeightToken.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ApiWeightToken *ApiWeightTokenTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ApiWeightToken.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ApiWeightToken *ApiWeightTokenSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ApiWeightToken.Contract.SetApprovalForAll(&_ApiWeightToken.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ApiWeightToken *ApiWeightTokenTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ApiWeightToken.Contract.SetApprovalForAll(&_ApiWeightToken.TransactOpts, operator, approved)
}

// SetPendingSeconds is a paid mutator transaction binding the contract method 0xacb00a11.
//
// Solidity: function setPendingSeconds(uint256 seconds_) returns()
func (_ApiWeightToken *ApiWeightTokenTransactor) SetPendingSeconds(opts *bind.TransactOpts, seconds_ *big.Int) (*types.Transaction, error) {
	return _ApiWeightToken.contract.Transact(opts, "setPendingSeconds", seconds_)
}

// SetPendingSeconds is a paid mutator transaction binding the contract method 0xacb00a11.
//
// Solidity: function setPendingSeconds(uint256 seconds_) returns()
func (_ApiWeightToken *ApiWeightTokenSession) SetPendingSeconds(seconds_ *big.Int) (*types.Transaction, error) {
	return _ApiWeightToken.Contract.SetPendingSeconds(&_ApiWeightToken.TransactOpts, seconds_)
}

// SetPendingSeconds is a paid mutator transaction binding the contract method 0xacb00a11.
//
// Solidity: function setPendingSeconds(uint256 seconds_) returns()
func (_ApiWeightToken *ApiWeightTokenTransactorSession) SetPendingSeconds(seconds_ *big.Int) (*types.Transaction, error) {
	return _ApiWeightToken.Contract.SetPendingSeconds(&_ApiWeightToken.TransactOpts, seconds_)
}

// ApiWeightTokenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ApiWeightToken contract.
type ApiWeightTokenApprovalForAllIterator struct {
	Event *ApiWeightTokenApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ApiWeightTokenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiWeightTokenApprovalForAll)
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
		it.Event = new(ApiWeightTokenApprovalForAll)
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
func (it *ApiWeightTokenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiWeightTokenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiWeightTokenApprovalForAll represents a ApprovalForAll event raised by the ApiWeightToken contract.
type ApiWeightTokenApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ApiWeightToken *ApiWeightTokenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*ApiWeightTokenApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ApiWeightToken.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ApiWeightTokenApprovalForAllIterator{contract: _ApiWeightToken.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ApiWeightToken *ApiWeightTokenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ApiWeightTokenApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ApiWeightToken.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiWeightTokenApprovalForAll)
				if err := _ApiWeightToken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_ApiWeightToken *ApiWeightTokenFilterer) ParseApprovalForAll(log types.Log) (*ApiWeightTokenApprovalForAll, error) {
	event := new(ApiWeightTokenApprovalForAll)
	if err := _ApiWeightToken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiWeightTokenResourceChangedIterator is returned from FilterResourceChanged and is used to iterate over the raw logs and unpacked data for ResourceChanged events raised by the ApiWeightToken contract.
type ApiWeightTokenResourceChangedIterator struct {
	Event *ApiWeightTokenResourceChanged // Event containing the contract specifics and raw log

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
func (it *ApiWeightTokenResourceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiWeightTokenResourceChanged)
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
		it.Event = new(ApiWeightTokenResourceChanged)
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
func (it *ApiWeightTokenResourceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiWeightTokenResourceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiWeightTokenResourceChanged represents a ResourceChanged event raised by the ApiWeightToken contract.
type ApiWeightTokenResourceChanged struct {
	Id     uint32
	Weight *big.Int
	Op     uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterResourceChanged is a free log retrieval operation binding the contract event 0x754de9a177534372342453272825c15351df3e7886699c2fece028d09379c400.
//
// Solidity: event ResourceChanged(uint32 indexed id, uint256 indexed weight, uint8 indexed op)
func (_ApiWeightToken *ApiWeightTokenFilterer) FilterResourceChanged(opts *bind.FilterOpts, id []uint32, weight []*big.Int, op []uint8) (*ApiWeightTokenResourceChangedIterator, error) {

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

	logs, sub, err := _ApiWeightToken.contract.FilterLogs(opts, "ResourceChanged", idRule, weightRule, opRule)
	if err != nil {
		return nil, err
	}
	return &ApiWeightTokenResourceChangedIterator{contract: _ApiWeightToken.contract, event: "ResourceChanged", logs: logs, sub: sub}, nil
}

// WatchResourceChanged is a free log subscription operation binding the contract event 0x754de9a177534372342453272825c15351df3e7886699c2fece028d09379c400.
//
// Solidity: event ResourceChanged(uint32 indexed id, uint256 indexed weight, uint8 indexed op)
func (_ApiWeightToken *ApiWeightTokenFilterer) WatchResourceChanged(opts *bind.WatchOpts, sink chan<- *ApiWeightTokenResourceChanged, id []uint32, weight []*big.Int, op []uint8) (event.Subscription, error) {

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

	logs, sub, err := _ApiWeightToken.contract.WatchLogs(opts, "ResourceChanged", idRule, weightRule, opRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiWeightTokenResourceChanged)
				if err := _ApiWeightToken.contract.UnpackLog(event, "ResourceChanged", log); err != nil {
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
func (_ApiWeightToken *ApiWeightTokenFilterer) ParseResourceChanged(log types.Log) (*ApiWeightTokenResourceChanged, error) {
	event := new(ApiWeightTokenResourceChanged)
	if err := _ApiWeightToken.contract.UnpackLog(event, "ResourceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiWeightTokenResourcePendingIterator is returned from FilterResourcePending and is used to iterate over the raw logs and unpacked data for ResourcePending events raised by the ApiWeightToken contract.
type ApiWeightTokenResourcePendingIterator struct {
	Event *ApiWeightTokenResourcePending // Event containing the contract specifics and raw log

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
func (it *ApiWeightTokenResourcePendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiWeightTokenResourcePending)
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
		it.Event = new(ApiWeightTokenResourcePending)
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
func (it *ApiWeightTokenResourcePendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiWeightTokenResourcePendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiWeightTokenResourcePending represents a ResourcePending event raised by the ApiWeightToken contract.
type ApiWeightTokenResourcePending struct {
	Id        uint32
	NewWeight *big.Int
	Op        uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterResourcePending is a free log retrieval operation binding the contract event 0xc2a05ca9bcad8635a6049c1240c2a1d2ed0c20c4f4b091ef84f7af523faf93d5.
//
// Solidity: event ResourcePending(uint32 indexed id, uint256 indexed newWeight, uint8 indexed op)
func (_ApiWeightToken *ApiWeightTokenFilterer) FilterResourcePending(opts *bind.FilterOpts, id []uint32, newWeight []*big.Int, op []uint8) (*ApiWeightTokenResourcePendingIterator, error) {

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

	logs, sub, err := _ApiWeightToken.contract.FilterLogs(opts, "ResourcePending", idRule, newWeightRule, opRule)
	if err != nil {
		return nil, err
	}
	return &ApiWeightTokenResourcePendingIterator{contract: _ApiWeightToken.contract, event: "ResourcePending", logs: logs, sub: sub}, nil
}

// WatchResourcePending is a free log subscription operation binding the contract event 0xc2a05ca9bcad8635a6049c1240c2a1d2ed0c20c4f4b091ef84f7af523faf93d5.
//
// Solidity: event ResourcePending(uint32 indexed id, uint256 indexed newWeight, uint8 indexed op)
func (_ApiWeightToken *ApiWeightTokenFilterer) WatchResourcePending(opts *bind.WatchOpts, sink chan<- *ApiWeightTokenResourcePending, id []uint32, newWeight []*big.Int, op []uint8) (event.Subscription, error) {

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

	logs, sub, err := _ApiWeightToken.contract.WatchLogs(opts, "ResourcePending", idRule, newWeightRule, opRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiWeightTokenResourcePending)
				if err := _ApiWeightToken.contract.UnpackLog(event, "ResourcePending", log); err != nil {
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
func (_ApiWeightToken *ApiWeightTokenFilterer) ParseResourcePending(log types.Log) (*ApiWeightTokenResourcePending, error) {
	event := new(ApiWeightTokenResourcePending)
	if err := _ApiWeightToken.contract.UnpackLog(event, "ResourcePending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiWeightTokenTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the ApiWeightToken contract.
type ApiWeightTokenTransferBatchIterator struct {
	Event *ApiWeightTokenTransferBatch // Event containing the contract specifics and raw log

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
func (it *ApiWeightTokenTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiWeightTokenTransferBatch)
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
		it.Event = new(ApiWeightTokenTransferBatch)
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
func (it *ApiWeightTokenTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiWeightTokenTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiWeightTokenTransferBatch represents a TransferBatch event raised by the ApiWeightToken contract.
type ApiWeightTokenTransferBatch struct {
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
func (_ApiWeightToken *ApiWeightTokenFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ApiWeightTokenTransferBatchIterator, error) {

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

	logs, sub, err := _ApiWeightToken.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ApiWeightTokenTransferBatchIterator{contract: _ApiWeightToken.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ApiWeightToken *ApiWeightTokenFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *ApiWeightTokenTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ApiWeightToken.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiWeightTokenTransferBatch)
				if err := _ApiWeightToken.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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
func (_ApiWeightToken *ApiWeightTokenFilterer) ParseTransferBatch(log types.Log) (*ApiWeightTokenTransferBatch, error) {
	event := new(ApiWeightTokenTransferBatch)
	if err := _ApiWeightToken.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiWeightTokenTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the ApiWeightToken contract.
type ApiWeightTokenTransferSingleIterator struct {
	Event *ApiWeightTokenTransferSingle // Event containing the contract specifics and raw log

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
func (it *ApiWeightTokenTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiWeightTokenTransferSingle)
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
		it.Event = new(ApiWeightTokenTransferSingle)
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
func (it *ApiWeightTokenTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiWeightTokenTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiWeightTokenTransferSingle represents a TransferSingle event raised by the ApiWeightToken contract.
type ApiWeightTokenTransferSingle struct {
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
func (_ApiWeightToken *ApiWeightTokenFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ApiWeightTokenTransferSingleIterator, error) {

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

	logs, sub, err := _ApiWeightToken.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ApiWeightTokenTransferSingleIterator{contract: _ApiWeightToken.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ApiWeightToken *ApiWeightTokenFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *ApiWeightTokenTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ApiWeightToken.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiWeightTokenTransferSingle)
				if err := _ApiWeightToken.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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
func (_ApiWeightToken *ApiWeightTokenFilterer) ParseTransferSingle(log types.Log) (*ApiWeightTokenTransferSingle, error) {
	event := new(ApiWeightTokenTransferSingle)
	if err := _ApiWeightToken.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiWeightTokenURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the ApiWeightToken contract.
type ApiWeightTokenURIIterator struct {
	Event *ApiWeightTokenURI // Event containing the contract specifics and raw log

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
func (it *ApiWeightTokenURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiWeightTokenURI)
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
		it.Event = new(ApiWeightTokenURI)
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
func (it *ApiWeightTokenURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiWeightTokenURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiWeightTokenURI represents a URI event raised by the ApiWeightToken contract.
type ApiWeightTokenURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ApiWeightToken *ApiWeightTokenFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*ApiWeightTokenURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ApiWeightToken.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &ApiWeightTokenURIIterator{contract: _ApiWeightToken.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ApiWeightToken *ApiWeightTokenFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *ApiWeightTokenURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ApiWeightToken.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiWeightTokenURI)
				if err := _ApiWeightToken.contract.UnpackLog(event, "URI", log); err != nil {
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
func (_ApiWeightToken *ApiWeightTokenFilterer) ParseURI(log types.Log) (*ApiWeightTokenURI, error) {
	event := new(ApiWeightTokenURI)
	if err := _ApiWeightToken.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
