// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ownable2step

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
	_ = abi.ConvertType
)

// Ownable2StepMetaData contains all meta data concerning the Ownable2Step contract.
var Ownable2StepMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Ownable2StepABI is the input ABI used to generate the binding from.
// Deprecated: Use Ownable2StepMetaData.ABI instead.
var Ownable2StepABI = Ownable2StepMetaData.ABI

// Ownable2Step is an auto generated Go binding around an Ethereum contract.
type Ownable2Step struct {
	Ownable2StepCaller     // Read-only binding to the contract
	Ownable2StepTransactor // Write-only binding to the contract
	Ownable2StepFilterer   // Log filterer for contract events
}

// Ownable2StepCaller is an auto generated read-only Go binding around an Ethereum contract.
type Ownable2StepCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ownable2StepTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Ownable2StepTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ownable2StepFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Ownable2StepFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ownable2StepSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Ownable2StepSession struct {
	Contract     *Ownable2Step     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Ownable2StepCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Ownable2StepCallerSession struct {
	Contract *Ownable2StepCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// Ownable2StepTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Ownable2StepTransactorSession struct {
	Contract     *Ownable2StepTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// Ownable2StepRaw is an auto generated low-level Go binding around an Ethereum contract.
type Ownable2StepRaw struct {
	Contract *Ownable2Step // Generic contract binding to access the raw methods on
}

// Ownable2StepCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Ownable2StepCallerRaw struct {
	Contract *Ownable2StepCaller // Generic read-only contract binding to access the raw methods on
}

// Ownable2StepTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Ownable2StepTransactorRaw struct {
	Contract *Ownable2StepTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable2Step creates a new instance of Ownable2Step, bound to a specific deployed contract.
func NewOwnable2Step(address common.Address, backend bind.ContractBackend) (*Ownable2Step, error) {
	contract, err := bindOwnable2Step(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable2Step{Ownable2StepCaller: Ownable2StepCaller{contract: contract}, Ownable2StepTransactor: Ownable2StepTransactor{contract: contract}, Ownable2StepFilterer: Ownable2StepFilterer{contract: contract}}, nil
}

// NewOwnable2StepCaller creates a new read-only instance of Ownable2Step, bound to a specific deployed contract.
func NewOwnable2StepCaller(address common.Address, caller bind.ContractCaller) (*Ownable2StepCaller, error) {
	contract, err := bindOwnable2Step(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Ownable2StepCaller{contract: contract}, nil
}

// NewOwnable2StepTransactor creates a new write-only instance of Ownable2Step, bound to a specific deployed contract.
func NewOwnable2StepTransactor(address common.Address, transactor bind.ContractTransactor) (*Ownable2StepTransactor, error) {
	contract, err := bindOwnable2Step(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Ownable2StepTransactor{contract: contract}, nil
}

// NewOwnable2StepFilterer creates a new log filterer instance of Ownable2Step, bound to a specific deployed contract.
func NewOwnable2StepFilterer(address common.Address, filterer bind.ContractFilterer) (*Ownable2StepFilterer, error) {
	contract, err := bindOwnable2Step(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Ownable2StepFilterer{contract: contract}, nil
}

// bindOwnable2Step binds a generic wrapper to an already deployed contract.
func bindOwnable2Step(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Ownable2StepMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable2Step *Ownable2StepRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable2Step.Contract.Ownable2StepCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable2Step *Ownable2StepRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable2Step.Contract.Ownable2StepTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable2Step *Ownable2StepRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable2Step.Contract.Ownable2StepTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable2Step *Ownable2StepCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable2Step.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable2Step *Ownable2StepTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable2Step.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable2Step *Ownable2StepTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable2Step.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable2Step *Ownable2StepCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable2Step.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable2Step *Ownable2StepSession) Owner() (common.Address, error) {
	return _Ownable2Step.Contract.Owner(&_Ownable2Step.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable2Step *Ownable2StepCallerSession) Owner() (common.Address, error) {
	return _Ownable2Step.Contract.Owner(&_Ownable2Step.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Ownable2Step *Ownable2StepCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable2Step.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Ownable2Step *Ownable2StepSession) PendingOwner() (common.Address, error) {
	return _Ownable2Step.Contract.PendingOwner(&_Ownable2Step.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Ownable2Step *Ownable2StepCallerSession) PendingOwner() (common.Address, error) {
	return _Ownable2Step.Contract.PendingOwner(&_Ownable2Step.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Ownable2Step *Ownable2StepTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable2Step.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Ownable2Step *Ownable2StepSession) AcceptOwnership() (*types.Transaction, error) {
	return _Ownable2Step.Contract.AcceptOwnership(&_Ownable2Step.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Ownable2Step *Ownable2StepTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Ownable2Step.Contract.AcceptOwnership(&_Ownable2Step.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable2Step *Ownable2StepTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable2Step.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable2Step *Ownable2StepSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable2Step.Contract.RenounceOwnership(&_Ownable2Step.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable2Step *Ownable2StepTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable2Step.Contract.RenounceOwnership(&_Ownable2Step.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable2Step *Ownable2StepTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable2Step.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable2Step *Ownable2StepSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable2Step.Contract.TransferOwnership(&_Ownable2Step.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable2Step *Ownable2StepTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable2Step.Contract.TransferOwnership(&_Ownable2Step.TransactOpts, newOwner)
}

// Ownable2StepOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Ownable2Step contract.
type Ownable2StepOwnershipTransferStartedIterator struct {
	Event *Ownable2StepOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *Ownable2StepOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ownable2StepOwnershipTransferStarted)
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
		it.Event = new(Ownable2StepOwnershipTransferStarted)
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
func (it *Ownable2StepOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ownable2StepOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ownable2StepOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Ownable2Step contract.
type Ownable2StepOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Ownable2Step *Ownable2StepFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Ownable2StepOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable2Step.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Ownable2StepOwnershipTransferStartedIterator{contract: _Ownable2Step.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Ownable2Step *Ownable2StepFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *Ownable2StepOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable2Step.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ownable2StepOwnershipTransferStarted)
				if err := _Ownable2Step.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Ownable2Step *Ownable2StepFilterer) ParseOwnershipTransferStarted(log types.Log) (*Ownable2StepOwnershipTransferStarted, error) {
	event := new(Ownable2StepOwnershipTransferStarted)
	if err := _Ownable2Step.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Ownable2StepOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable2Step contract.
type Ownable2StepOwnershipTransferredIterator struct {
	Event *Ownable2StepOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Ownable2StepOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ownable2StepOwnershipTransferred)
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
		it.Event = new(Ownable2StepOwnershipTransferred)
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
func (it *Ownable2StepOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ownable2StepOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ownable2StepOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable2Step contract.
type Ownable2StepOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable2Step *Ownable2StepFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Ownable2StepOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable2Step.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Ownable2StepOwnershipTransferredIterator{contract: _Ownable2Step.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable2Step *Ownable2StepFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Ownable2StepOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable2Step.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ownable2StepOwnershipTransferred)
				if err := _Ownable2Step.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ownable2Step *Ownable2StepFilterer) ParseOwnershipTransferred(log types.Log) (*Ownable2StepOwnershipTransferred, error) {
	event := new(Ownable2StepOwnershipTransferred)
	if err := _Ownable2Step.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
