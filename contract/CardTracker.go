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

// ICardTrackerVipInfo is an auto generated low-level Go binding around an user-defined struct.
type ICardTrackerVipInfo struct {
	ExpireAt *big.Int
	Props    ICardTemplateProps
}

// CardTrackerMetaData contains all meta data concerning the CardTracker contract.
var CardTrackerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"eventSource\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expireAt\",\"type\":\"uint256\"}],\"name\":\"VipChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"giveawayDuration\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"}],\"internalType\":\"structICardTemplate.Props\",\"name\":\"props\",\"type\":\"tuple\"}],\"internalType\":\"structICardTemplate.Template\",\"name\":\"template\",\"type\":\"tuple\"}],\"internalType\":\"structICards.Card\",\"name\":\"card\",\"type\":\"tuple\"}],\"name\":\"applyCard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getVipInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"expireAt\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"}],\"internalType\":\"structICardTemplate.Props\",\"name\":\"props\",\"type\":\"tuple\"}],\"internalType\":\"structICardTracker.VipInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"eventSource\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CardTrackerABI is the input ABI used to generate the binding from.
// Deprecated: Use CardTrackerMetaData.ABI instead.
var CardTrackerABI = CardTrackerMetaData.ABI

// CardTracker is an auto generated Go binding around an Ethereum contract.
type CardTracker struct {
	CardTrackerCaller     // Read-only binding to the contract
	CardTrackerTransactor // Write-only binding to the contract
	CardTrackerFilterer   // Log filterer for contract events
}

// CardTrackerCaller is an auto generated read-only Go binding around an Ethereum contract.
type CardTrackerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardTrackerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CardTrackerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardTrackerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CardTrackerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardTrackerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CardTrackerSession struct {
	Contract     *CardTracker      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CardTrackerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CardTrackerCallerSession struct {
	Contract *CardTrackerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CardTrackerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CardTrackerTransactorSession struct {
	Contract     *CardTrackerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CardTrackerRaw is an auto generated low-level Go binding around an Ethereum contract.
type CardTrackerRaw struct {
	Contract *CardTracker // Generic contract binding to access the raw methods on
}

// CardTrackerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CardTrackerCallerRaw struct {
	Contract *CardTrackerCaller // Generic read-only contract binding to access the raw methods on
}

// CardTrackerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CardTrackerTransactorRaw struct {
	Contract *CardTrackerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCardTracker creates a new instance of CardTracker, bound to a specific deployed contract.
func NewCardTracker(address common.Address, backend bind.ContractBackend) (*CardTracker, error) {
	contract, err := bindCardTracker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CardTracker{CardTrackerCaller: CardTrackerCaller{contract: contract}, CardTrackerTransactor: CardTrackerTransactor{contract: contract}, CardTrackerFilterer: CardTrackerFilterer{contract: contract}}, nil
}

// NewCardTrackerCaller creates a new read-only instance of CardTracker, bound to a specific deployed contract.
func NewCardTrackerCaller(address common.Address, caller bind.ContractCaller) (*CardTrackerCaller, error) {
	contract, err := bindCardTracker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CardTrackerCaller{contract: contract}, nil
}

// NewCardTrackerTransactor creates a new write-only instance of CardTracker, bound to a specific deployed contract.
func NewCardTrackerTransactor(address common.Address, transactor bind.ContractTransactor) (*CardTrackerTransactor, error) {
	contract, err := bindCardTracker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CardTrackerTransactor{contract: contract}, nil
}

// NewCardTrackerFilterer creates a new log filterer instance of CardTracker, bound to a specific deployed contract.
func NewCardTrackerFilterer(address common.Address, filterer bind.ContractFilterer) (*CardTrackerFilterer, error) {
	contract, err := bindCardTracker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CardTrackerFilterer{contract: contract}, nil
}

// bindCardTracker binds a generic wrapper to an already deployed contract.
func bindCardTracker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CardTrackerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CardTracker *CardTrackerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CardTracker.Contract.CardTrackerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CardTracker *CardTrackerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CardTracker.Contract.CardTrackerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CardTracker *CardTrackerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CardTracker.Contract.CardTrackerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CardTracker *CardTrackerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CardTracker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CardTracker *CardTrackerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CardTracker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CardTracker *CardTrackerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CardTracker.Contract.contract.Transact(opts, method, params...)
}

// GetVipInfo is a free data retrieval call binding the contract method 0x05e1aaa1.
//
// Solidity: function getVipInfo(address account) view returns((uint256,(string[],string[])))
func (_CardTracker *CardTrackerCaller) GetVipInfo(opts *bind.CallOpts, account common.Address) (ICardTrackerVipInfo, error) {
	var out []interface{}
	err := _CardTracker.contract.Call(opts, &out, "getVipInfo", account)

	if err != nil {
		return *new(ICardTrackerVipInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ICardTrackerVipInfo)).(*ICardTrackerVipInfo)

	return out0, err

}

// GetVipInfo is a free data retrieval call binding the contract method 0x05e1aaa1.
//
// Solidity: function getVipInfo(address account) view returns((uint256,(string[],string[])))
func (_CardTracker *CardTrackerSession) GetVipInfo(account common.Address) (ICardTrackerVipInfo, error) {
	return _CardTracker.Contract.GetVipInfo(&_CardTracker.CallOpts, account)
}

// GetVipInfo is a free data retrieval call binding the contract method 0x05e1aaa1.
//
// Solidity: function getVipInfo(address account) view returns((uint256,(string[],string[])))
func (_CardTracker *CardTrackerCallerSession) GetVipInfo(account common.Address) (ICardTrackerVipInfo, error) {
	return _CardTracker.Contract.GetVipInfo(&_CardTracker.CallOpts, account)
}

// ApplyCard is a paid mutator transaction binding the contract method 0x7ed41358.
//
// Solidity: function applyCard(address from, address to, (uint256,uint256,address,uint256,(uint256,string,string,uint256,uint256,uint256,(string[],string[]))) card) returns()
func (_CardTracker *CardTrackerTransactor) ApplyCard(opts *bind.TransactOpts, from common.Address, to common.Address, card ICardsCard) (*types.Transaction, error) {
	return _CardTracker.contract.Transact(opts, "applyCard", from, to, card)
}

// ApplyCard is a paid mutator transaction binding the contract method 0x7ed41358.
//
// Solidity: function applyCard(address from, address to, (uint256,uint256,address,uint256,(uint256,string,string,uint256,uint256,uint256,(string[],string[]))) card) returns()
func (_CardTracker *CardTrackerSession) ApplyCard(from common.Address, to common.Address, card ICardsCard) (*types.Transaction, error) {
	return _CardTracker.Contract.ApplyCard(&_CardTracker.TransactOpts, from, to, card)
}

// ApplyCard is a paid mutator transaction binding the contract method 0x7ed41358.
//
// Solidity: function applyCard(address from, address to, (uint256,uint256,address,uint256,(uint256,string,string,uint256,uint256,uint256,(string[],string[]))) card) returns()
func (_CardTracker *CardTrackerTransactorSession) ApplyCard(from common.Address, to common.Address, card ICardsCard) (*types.Transaction, error) {
	return _CardTracker.Contract.ApplyCard(&_CardTracker.TransactOpts, from, to, card)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address eventSource) returns()
func (_CardTracker *CardTrackerTransactor) Initialize(opts *bind.TransactOpts, eventSource common.Address) (*types.Transaction, error) {
	return _CardTracker.contract.Transact(opts, "initialize", eventSource)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address eventSource) returns()
func (_CardTracker *CardTrackerSession) Initialize(eventSource common.Address) (*types.Transaction, error) {
	return _CardTracker.Contract.Initialize(&_CardTracker.TransactOpts, eventSource)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address eventSource) returns()
func (_CardTracker *CardTrackerTransactorSession) Initialize(eventSource common.Address) (*types.Transaction, error) {
	return _CardTracker.Contract.Initialize(&_CardTracker.TransactOpts, eventSource)
}

// CardTrackerVipChangedIterator is returned from FilterVipChanged and is used to iterate over the raw logs and unpacked data for VipChanged events raised by the CardTracker contract.
type CardTrackerVipChangedIterator struct {
	Event *CardTrackerVipChanged // Event containing the contract specifics and raw log

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
func (it *CardTrackerVipChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardTrackerVipChanged)
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
		it.Event = new(CardTrackerVipChanged)
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
func (it *CardTrackerVipChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardTrackerVipChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardTrackerVipChanged represents a VipChanged event raised by the CardTracker contract.
type CardTrackerVipChanged struct {
	Account  common.Address
	ExpireAt *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVipChanged is a free log retrieval operation binding the contract event 0x20f6d1c339decb8b3e4b6e6e792d7d35f3ed985b25407cd7246d3e2784c6cefb.
//
// Solidity: event VipChanged(address indexed account, uint256 expireAt)
func (_CardTracker *CardTrackerFilterer) FilterVipChanged(opts *bind.FilterOpts, account []common.Address) (*CardTrackerVipChangedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CardTracker.contract.FilterLogs(opts, "VipChanged", accountRule)
	if err != nil {
		return nil, err
	}
	return &CardTrackerVipChangedIterator{contract: _CardTracker.contract, event: "VipChanged", logs: logs, sub: sub}, nil
}

// WatchVipChanged is a free log subscription operation binding the contract event 0x20f6d1c339decb8b3e4b6e6e792d7d35f3ed985b25407cd7246d3e2784c6cefb.
//
// Solidity: event VipChanged(address indexed account, uint256 expireAt)
func (_CardTracker *CardTrackerFilterer) WatchVipChanged(opts *bind.WatchOpts, sink chan<- *CardTrackerVipChanged, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _CardTracker.contract.WatchLogs(opts, "VipChanged", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardTrackerVipChanged)
				if err := _CardTracker.contract.UnpackLog(event, "VipChanged", log); err != nil {
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

// ParseVipChanged is a log parse operation binding the contract event 0x20f6d1c339decb8b3e4b6e6e792d7d35f3ed985b25407cd7246d3e2784c6cefb.
//
// Solidity: event VipChanged(address indexed account, uint256 expireAt)
func (_CardTracker *CardTrackerFilterer) ParseVipChanged(log types.Log) (*CardTrackerVipChanged, error) {
	event := new(CardTrackerVipChanged)
	if err := _CardTracker.contract.UnpackLog(event, "VipChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
