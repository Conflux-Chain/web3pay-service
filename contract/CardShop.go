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

// ICardTemplateProps is an auto generated low-level Go binding around an user-defined struct.
type ICardTemplateProps struct {
	Keys   []string
	Values []string
}

// ICardTemplateTemplate is an auto generated low-level Go binding around an user-defined struct.
type ICardTemplateTemplate struct {
	Id               *big.Int
	Name             string
	Description      string
	Price            *big.Int
	Duration         *big.Int
	GiveawayDuration *big.Int
	Props            ICardTemplateProps
}

// ICardsCard is an auto generated low-level Go binding around an user-defined struct.
type ICardsCard struct {
	Id       *big.Int
	Duration *big.Int
	Owner    common.Address
	Count    *big.Int
	Template ICardTemplateTemplate
}

// CardShopMetaData contains all meta data concerning the CardShop contract.
var CardShopMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cardId\",\"type\":\"uint256\"}],\"name\":\"GAVEN_CARD\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"belongsToApp\",\"outputs\":[{\"internalType\":\"contractIApp\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"templateId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"buyWithAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"templateId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"buyWithEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getCard\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"giveawayDuration\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"}],\"internalType\":\"structICardTemplate.Props\",\"name\":\"props\",\"type\":\"tuple\"}],\"internalType\":\"structICardTemplate.Template\",\"name\":\"template\",\"type\":\"tuple\"}],\"internalType\":\"structICards.Card\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"receiverArr\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"countArr\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"templateId\",\"type\":\"uint256\"}],\"name\":\"giveCardBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIApp\",\"name\":\"belongsTo_\",\"type\":\"address\"},{\"internalType\":\"contractICardTemplate\",\"name\":\"template_\",\"type\":\"address\"},{\"internalType\":\"contractICards\",\"name\":\"instance_\",\"type\":\"address\"},{\"internalType\":\"contractICardTracker\",\"name\":\"tracker_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"instance\",\"outputs\":[{\"internalType\":\"contractICards\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextCardId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"template\",\"outputs\":[{\"internalType\":\"contractICardTemplate\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tracker\",\"outputs\":[{\"internalType\":\"contractICardTracker\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CardShopABI is the input ABI used to generate the binding from.
// Deprecated: Use CardShopMetaData.ABI instead.
var CardShopABI = CardShopMetaData.ABI

// CardShop is an auto generated Go binding around an Ethereum contract.
type CardShop struct {
	CardShopCaller     // Read-only binding to the contract
	CardShopTransactor // Write-only binding to the contract
	CardShopFilterer   // Log filterer for contract events
}

// CardShopCaller is an auto generated read-only Go binding around an Ethereum contract.
type CardShopCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardShopTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CardShopTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardShopFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CardShopFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardShopSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CardShopSession struct {
	Contract     *CardShop         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CardShopCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CardShopCallerSession struct {
	Contract *CardShopCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CardShopTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CardShopTransactorSession struct {
	Contract     *CardShopTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CardShopRaw is an auto generated low-level Go binding around an Ethereum contract.
type CardShopRaw struct {
	Contract *CardShop // Generic contract binding to access the raw methods on
}

// CardShopCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CardShopCallerRaw struct {
	Contract *CardShopCaller // Generic read-only contract binding to access the raw methods on
}

// CardShopTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CardShopTransactorRaw struct {
	Contract *CardShopTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCardShop creates a new instance of CardShop, bound to a specific deployed contract.
func NewCardShop(address common.Address, backend bind.ContractBackend) (*CardShop, error) {
	contract, err := bindCardShop(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CardShop{CardShopCaller: CardShopCaller{contract: contract}, CardShopTransactor: CardShopTransactor{contract: contract}, CardShopFilterer: CardShopFilterer{contract: contract}}, nil
}

// NewCardShopCaller creates a new read-only instance of CardShop, bound to a specific deployed contract.
func NewCardShopCaller(address common.Address, caller bind.ContractCaller) (*CardShopCaller, error) {
	contract, err := bindCardShop(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CardShopCaller{contract: contract}, nil
}

// NewCardShopTransactor creates a new write-only instance of CardShop, bound to a specific deployed contract.
func NewCardShopTransactor(address common.Address, transactor bind.ContractTransactor) (*CardShopTransactor, error) {
	contract, err := bindCardShop(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CardShopTransactor{contract: contract}, nil
}

// NewCardShopFilterer creates a new log filterer instance of CardShop, bound to a specific deployed contract.
func NewCardShopFilterer(address common.Address, filterer bind.ContractFilterer) (*CardShopFilterer, error) {
	contract, err := bindCardShop(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CardShopFilterer{contract: contract}, nil
}

// bindCardShop binds a generic wrapper to an already deployed contract.
func bindCardShop(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CardShopABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CardShop *CardShopRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CardShop.Contract.CardShopCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CardShop *CardShopRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CardShop.Contract.CardShopTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CardShop *CardShopRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CardShop.Contract.CardShopTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CardShop *CardShopCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CardShop.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CardShop *CardShopTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CardShop.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CardShop *CardShopTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CardShop.Contract.contract.Transact(opts, method, params...)
}

// BelongsToApp is a free data retrieval call binding the contract method 0x53f63ca6.
//
// Solidity: function belongsToApp() view returns(address)
func (_CardShop *CardShopCaller) BelongsToApp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CardShop.contract.Call(opts, &out, "belongsToApp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BelongsToApp is a free data retrieval call binding the contract method 0x53f63ca6.
//
// Solidity: function belongsToApp() view returns(address)
func (_CardShop *CardShopSession) BelongsToApp() (common.Address, error) {
	return _CardShop.Contract.BelongsToApp(&_CardShop.CallOpts)
}

// BelongsToApp is a free data retrieval call binding the contract method 0x53f63ca6.
//
// Solidity: function belongsToApp() view returns(address)
func (_CardShop *CardShopCallerSession) BelongsToApp() (common.Address, error) {
	return _CardShop.Contract.BelongsToApp(&_CardShop.CallOpts)
}

// GetCard is a free data retrieval call binding the contract method 0x9188d312.
//
// Solidity: function getCard(uint256 id) view returns((uint256,uint256,address,uint256,(uint256,string,string,uint256,uint256,uint256,(string[],string[]))))
func (_CardShop *CardShopCaller) GetCard(opts *bind.CallOpts, id *big.Int) (ICardsCard, error) {
	var out []interface{}
	err := _CardShop.contract.Call(opts, &out, "getCard", id)

	if err != nil {
		return *new(ICardsCard), err
	}

	out0 := *abi.ConvertType(out[0], new(ICardsCard)).(*ICardsCard)

	return out0, err

}

// GetCard is a free data retrieval call binding the contract method 0x9188d312.
//
// Solidity: function getCard(uint256 id) view returns((uint256,uint256,address,uint256,(uint256,string,string,uint256,uint256,uint256,(string[],string[]))))
func (_CardShop *CardShopSession) GetCard(id *big.Int) (ICardsCard, error) {
	return _CardShop.Contract.GetCard(&_CardShop.CallOpts, id)
}

// GetCard is a free data retrieval call binding the contract method 0x9188d312.
//
// Solidity: function getCard(uint256 id) view returns((uint256,uint256,address,uint256,(uint256,string,string,uint256,uint256,uint256,(string[],string[]))))
func (_CardShop *CardShopCallerSession) GetCard(id *big.Int) (ICardsCard, error) {
	return _CardShop.Contract.GetCard(&_CardShop.CallOpts, id)
}

// Instance is a free data retrieval call binding the contract method 0x022ec095.
//
// Solidity: function instance() view returns(address)
func (_CardShop *CardShopCaller) Instance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CardShop.contract.Call(opts, &out, "instance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Instance is a free data retrieval call binding the contract method 0x022ec095.
//
// Solidity: function instance() view returns(address)
func (_CardShop *CardShopSession) Instance() (common.Address, error) {
	return _CardShop.Contract.Instance(&_CardShop.CallOpts)
}

// Instance is a free data retrieval call binding the contract method 0x022ec095.
//
// Solidity: function instance() view returns(address)
func (_CardShop *CardShopCallerSession) Instance() (common.Address, error) {
	return _CardShop.Contract.Instance(&_CardShop.CallOpts)
}

// NextCardId is a free data retrieval call binding the contract method 0xc5263e8c.
//
// Solidity: function nextCardId() view returns(uint256)
func (_CardShop *CardShopCaller) NextCardId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CardShop.contract.Call(opts, &out, "nextCardId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextCardId is a free data retrieval call binding the contract method 0xc5263e8c.
//
// Solidity: function nextCardId() view returns(uint256)
func (_CardShop *CardShopSession) NextCardId() (*big.Int, error) {
	return _CardShop.Contract.NextCardId(&_CardShop.CallOpts)
}

// NextCardId is a free data retrieval call binding the contract method 0xc5263e8c.
//
// Solidity: function nextCardId() view returns(uint256)
func (_CardShop *CardShopCallerSession) NextCardId() (*big.Int, error) {
	return _CardShop.Contract.NextCardId(&_CardShop.CallOpts)
}

// Template is a free data retrieval call binding the contract method 0x6f2ddd93.
//
// Solidity: function template() view returns(address)
func (_CardShop *CardShopCaller) Template(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CardShop.contract.Call(opts, &out, "template")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Template is a free data retrieval call binding the contract method 0x6f2ddd93.
//
// Solidity: function template() view returns(address)
func (_CardShop *CardShopSession) Template() (common.Address, error) {
	return _CardShop.Contract.Template(&_CardShop.CallOpts)
}

// Template is a free data retrieval call binding the contract method 0x6f2ddd93.
//
// Solidity: function template() view returns(address)
func (_CardShop *CardShopCallerSession) Template() (common.Address, error) {
	return _CardShop.Contract.Template(&_CardShop.CallOpts)
}

// Tracker is a free data retrieval call binding the contract method 0xf52bccad.
//
// Solidity: function tracker() view returns(address)
func (_CardShop *CardShopCaller) Tracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CardShop.contract.Call(opts, &out, "tracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tracker is a free data retrieval call binding the contract method 0xf52bccad.
//
// Solidity: function tracker() view returns(address)
func (_CardShop *CardShopSession) Tracker() (common.Address, error) {
	return _CardShop.Contract.Tracker(&_CardShop.CallOpts)
}

// Tracker is a free data retrieval call binding the contract method 0xf52bccad.
//
// Solidity: function tracker() view returns(address)
func (_CardShop *CardShopCallerSession) Tracker() (common.Address, error) {
	return _CardShop.Contract.Tracker(&_CardShop.CallOpts)
}

// BuyWithAsset is a paid mutator transaction binding the contract method 0x9ffd313b.
//
// Solidity: function buyWithAsset(address receiver, uint256 templateId, uint256 count) returns()
func (_CardShop *CardShopTransactor) BuyWithAsset(opts *bind.TransactOpts, receiver common.Address, templateId *big.Int, count *big.Int) (*types.Transaction, error) {
	return _CardShop.contract.Transact(opts, "buyWithAsset", receiver, templateId, count)
}

// BuyWithAsset is a paid mutator transaction binding the contract method 0x9ffd313b.
//
// Solidity: function buyWithAsset(address receiver, uint256 templateId, uint256 count) returns()
func (_CardShop *CardShopSession) BuyWithAsset(receiver common.Address, templateId *big.Int, count *big.Int) (*types.Transaction, error) {
	return _CardShop.Contract.BuyWithAsset(&_CardShop.TransactOpts, receiver, templateId, count)
}

// BuyWithAsset is a paid mutator transaction binding the contract method 0x9ffd313b.
//
// Solidity: function buyWithAsset(address receiver, uint256 templateId, uint256 count) returns()
func (_CardShop *CardShopTransactorSession) BuyWithAsset(receiver common.Address, templateId *big.Int, count *big.Int) (*types.Transaction, error) {
	return _CardShop.Contract.BuyWithAsset(&_CardShop.TransactOpts, receiver, templateId, count)
}

// BuyWithEth is a paid mutator transaction binding the contract method 0x8085d06d.
//
// Solidity: function buyWithEth(address receiver, uint256 templateId, uint256 count) payable returns()
func (_CardShop *CardShopTransactor) BuyWithEth(opts *bind.TransactOpts, receiver common.Address, templateId *big.Int, count *big.Int) (*types.Transaction, error) {
	return _CardShop.contract.Transact(opts, "buyWithEth", receiver, templateId, count)
}

// BuyWithEth is a paid mutator transaction binding the contract method 0x8085d06d.
//
// Solidity: function buyWithEth(address receiver, uint256 templateId, uint256 count) payable returns()
func (_CardShop *CardShopSession) BuyWithEth(receiver common.Address, templateId *big.Int, count *big.Int) (*types.Transaction, error) {
	return _CardShop.Contract.BuyWithEth(&_CardShop.TransactOpts, receiver, templateId, count)
}

// BuyWithEth is a paid mutator transaction binding the contract method 0x8085d06d.
//
// Solidity: function buyWithEth(address receiver, uint256 templateId, uint256 count) payable returns()
func (_CardShop *CardShopTransactorSession) BuyWithEth(receiver common.Address, templateId *big.Int, count *big.Int) (*types.Transaction, error) {
	return _CardShop.Contract.BuyWithEth(&_CardShop.TransactOpts, receiver, templateId, count)
}

// GiveCardBatch is a paid mutator transaction binding the contract method 0x66f85a09.
//
// Solidity: function giveCardBatch(address[] receiverArr, uint256[] countArr, uint256 templateId) returns()
func (_CardShop *CardShopTransactor) GiveCardBatch(opts *bind.TransactOpts, receiverArr []common.Address, countArr []*big.Int, templateId *big.Int) (*types.Transaction, error) {
	return _CardShop.contract.Transact(opts, "giveCardBatch", receiverArr, countArr, templateId)
}

// GiveCardBatch is a paid mutator transaction binding the contract method 0x66f85a09.
//
// Solidity: function giveCardBatch(address[] receiverArr, uint256[] countArr, uint256 templateId) returns()
func (_CardShop *CardShopSession) GiveCardBatch(receiverArr []common.Address, countArr []*big.Int, templateId *big.Int) (*types.Transaction, error) {
	return _CardShop.Contract.GiveCardBatch(&_CardShop.TransactOpts, receiverArr, countArr, templateId)
}

// GiveCardBatch is a paid mutator transaction binding the contract method 0x66f85a09.
//
// Solidity: function giveCardBatch(address[] receiverArr, uint256[] countArr, uint256 templateId) returns()
func (_CardShop *CardShopTransactorSession) GiveCardBatch(receiverArr []common.Address, countArr []*big.Int, templateId *big.Int) (*types.Transaction, error) {
	return _CardShop.Contract.GiveCardBatch(&_CardShop.TransactOpts, receiverArr, countArr, templateId)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address belongsTo_, address template_, address instance_, address tracker_) returns()
func (_CardShop *CardShopTransactor) Initialize(opts *bind.TransactOpts, belongsTo_ common.Address, template_ common.Address, instance_ common.Address, tracker_ common.Address) (*types.Transaction, error) {
	return _CardShop.contract.Transact(opts, "initialize", belongsTo_, template_, instance_, tracker_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address belongsTo_, address template_, address instance_, address tracker_) returns()
func (_CardShop *CardShopSession) Initialize(belongsTo_ common.Address, template_ common.Address, instance_ common.Address, tracker_ common.Address) (*types.Transaction, error) {
	return _CardShop.Contract.Initialize(&_CardShop.TransactOpts, belongsTo_, template_, instance_, tracker_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address belongsTo_, address template_, address instance_, address tracker_) returns()
func (_CardShop *CardShopTransactorSession) Initialize(belongsTo_ common.Address, template_ common.Address, instance_ common.Address, tracker_ common.Address) (*types.Transaction, error) {
	return _CardShop.Contract.Initialize(&_CardShop.TransactOpts, belongsTo_, template_, instance_, tracker_)
}

// CardShopGAVENCARDIterator is returned from FilterGAVENCARD and is used to iterate over the raw logs and unpacked data for GAVENCARD events raised by the CardShop contract.
type CardShopGAVENCARDIterator struct {
	Event *CardShopGAVENCARD // Event containing the contract specifics and raw log

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
func (it *CardShopGAVENCARDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardShopGAVENCARD)
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
		it.Event = new(CardShopGAVENCARD)
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
func (it *CardShopGAVENCARDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardShopGAVENCARDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardShopGAVENCARD represents a GAVENCARD event raised by the CardShop contract.
type CardShopGAVENCARD struct {
	Operator common.Address
	To       common.Address
	CardId   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGAVENCARD is a free log retrieval operation binding the contract event 0x5e8eec90a2cf75235f88a955dbe5c00291cd2837e4a1dc81624872515a02bdd0.
//
// Solidity: event GAVEN_CARD(address indexed operator, address indexed to, uint256 cardId)
func (_CardShop *CardShopFilterer) FilterGAVENCARD(opts *bind.FilterOpts, operator []common.Address, to []common.Address) (*CardShopGAVENCARDIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CardShop.contract.FilterLogs(opts, "GAVEN_CARD", operatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CardShopGAVENCARDIterator{contract: _CardShop.contract, event: "GAVEN_CARD", logs: logs, sub: sub}, nil
}

// WatchGAVENCARD is a free log subscription operation binding the contract event 0x5e8eec90a2cf75235f88a955dbe5c00291cd2837e4a1dc81624872515a02bdd0.
//
// Solidity: event GAVEN_CARD(address indexed operator, address indexed to, uint256 cardId)
func (_CardShop *CardShopFilterer) WatchGAVENCARD(opts *bind.WatchOpts, sink chan<- *CardShopGAVENCARD, operator []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CardShop.contract.WatchLogs(opts, "GAVEN_CARD", operatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardShopGAVENCARD)
				if err := _CardShop.contract.UnpackLog(event, "GAVEN_CARD", log); err != nil {
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

// ParseGAVENCARD is a log parse operation binding the contract event 0x5e8eec90a2cf75235f88a955dbe5c00291cd2837e4a1dc81624872515a02bdd0.
//
// Solidity: event GAVEN_CARD(address indexed operator, address indexed to, uint256 cardId)
func (_CardShop *CardShopFilterer) ParseGAVENCARD(log types.Log) (*CardShopGAVENCARD, error) {
	event := new(CardShopGAVENCARD)
	if err := _CardShop.contract.UnpackLog(event, "GAVEN_CARD", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
