// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetainteractor

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

// ZetaInteractorMetaData contains all meta data concerning the ZetaInteractor contract.
var ZetaInteractorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"InvalidCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDestinationChainId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZetaMessageCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZetaRevertCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"connector\",\"outputs\":[{\"internalType\":\"contractZetaConnector\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"interactorsByChainId\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"contractAddress\",\"type\":\"bytes\"}],\"name\":\"setInteractorByChainId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ZetaInteractorABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaInteractorMetaData.ABI instead.
var ZetaInteractorABI = ZetaInteractorMetaData.ABI

// ZetaInteractor is an auto generated Go binding around an Ethereum contract.
type ZetaInteractor struct {
	ZetaInteractorCaller     // Read-only binding to the contract
	ZetaInteractorTransactor // Write-only binding to the contract
	ZetaInteractorFilterer   // Log filterer for contract events
}

// ZetaInteractorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaInteractorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaInteractorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaInteractorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaInteractorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaInteractorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaInteractorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaInteractorSession struct {
	Contract     *ZetaInteractor   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZetaInteractorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaInteractorCallerSession struct {
	Contract *ZetaInteractorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ZetaInteractorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaInteractorTransactorSession struct {
	Contract     *ZetaInteractorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ZetaInteractorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaInteractorRaw struct {
	Contract *ZetaInteractor // Generic contract binding to access the raw methods on
}

// ZetaInteractorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaInteractorCallerRaw struct {
	Contract *ZetaInteractorCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaInteractorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaInteractorTransactorRaw struct {
	Contract *ZetaInteractorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaInteractor creates a new instance of ZetaInteractor, bound to a specific deployed contract.
func NewZetaInteractor(address common.Address, backend bind.ContractBackend) (*ZetaInteractor, error) {
	contract, err := bindZetaInteractor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaInteractor{ZetaInteractorCaller: ZetaInteractorCaller{contract: contract}, ZetaInteractorTransactor: ZetaInteractorTransactor{contract: contract}, ZetaInteractorFilterer: ZetaInteractorFilterer{contract: contract}}, nil
}

// NewZetaInteractorCaller creates a new read-only instance of ZetaInteractor, bound to a specific deployed contract.
func NewZetaInteractorCaller(address common.Address, caller bind.ContractCaller) (*ZetaInteractorCaller, error) {
	contract, err := bindZetaInteractor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaInteractorCaller{contract: contract}, nil
}

// NewZetaInteractorTransactor creates a new write-only instance of ZetaInteractor, bound to a specific deployed contract.
func NewZetaInteractorTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaInteractorTransactor, error) {
	contract, err := bindZetaInteractor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaInteractorTransactor{contract: contract}, nil
}

// NewZetaInteractorFilterer creates a new log filterer instance of ZetaInteractor, bound to a specific deployed contract.
func NewZetaInteractorFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaInteractorFilterer, error) {
	contract, err := bindZetaInteractor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaInteractorFilterer{contract: contract}, nil
}

// bindZetaInteractor binds a generic wrapper to an already deployed contract.
func bindZetaInteractor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaInteractorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaInteractor *ZetaInteractorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaInteractor.Contract.ZetaInteractorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaInteractor *ZetaInteractorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaInteractor.Contract.ZetaInteractorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaInteractor *ZetaInteractorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaInteractor.Contract.ZetaInteractorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaInteractor *ZetaInteractorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaInteractor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaInteractor *ZetaInteractorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaInteractor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaInteractor *ZetaInteractorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaInteractor.Contract.contract.Transact(opts, method, params...)
}

// Connector is a free data retrieval call binding the contract method 0x83f3084f.
//
// Solidity: function connector() view returns(address)
func (_ZetaInteractor *ZetaInteractorCaller) Connector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaInteractor.contract.Call(opts, &out, "connector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Connector is a free data retrieval call binding the contract method 0x83f3084f.
//
// Solidity: function connector() view returns(address)
func (_ZetaInteractor *ZetaInteractorSession) Connector() (common.Address, error) {
	return _ZetaInteractor.Contract.Connector(&_ZetaInteractor.CallOpts)
}

// Connector is a free data retrieval call binding the contract method 0x83f3084f.
//
// Solidity: function connector() view returns(address)
func (_ZetaInteractor *ZetaInteractorCallerSession) Connector() (common.Address, error) {
	return _ZetaInteractor.Contract.Connector(&_ZetaInteractor.CallOpts)
}

// InteractorsByChainId is a free data retrieval call binding the contract method 0x2618143f.
//
// Solidity: function interactorsByChainId(uint256 ) view returns(bytes)
func (_ZetaInteractor *ZetaInteractorCaller) InteractorsByChainId(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _ZetaInteractor.contract.Call(opts, &out, "interactorsByChainId", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// InteractorsByChainId is a free data retrieval call binding the contract method 0x2618143f.
//
// Solidity: function interactorsByChainId(uint256 ) view returns(bytes)
func (_ZetaInteractor *ZetaInteractorSession) InteractorsByChainId(arg0 *big.Int) ([]byte, error) {
	return _ZetaInteractor.Contract.InteractorsByChainId(&_ZetaInteractor.CallOpts, arg0)
}

// InteractorsByChainId is a free data retrieval call binding the contract method 0x2618143f.
//
// Solidity: function interactorsByChainId(uint256 ) view returns(bytes)
func (_ZetaInteractor *ZetaInteractorCallerSession) InteractorsByChainId(arg0 *big.Int) ([]byte, error) {
	return _ZetaInteractor.Contract.InteractorsByChainId(&_ZetaInteractor.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZetaInteractor *ZetaInteractorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaInteractor.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZetaInteractor *ZetaInteractorSession) Owner() (common.Address, error) {
	return _ZetaInteractor.Contract.Owner(&_ZetaInteractor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZetaInteractor *ZetaInteractorCallerSession) Owner() (common.Address, error) {
	return _ZetaInteractor.Contract.Owner(&_ZetaInteractor.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_ZetaInteractor *ZetaInteractorCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaInteractor.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_ZetaInteractor *ZetaInteractorSession) PendingOwner() (common.Address, error) {
	return _ZetaInteractor.Contract.PendingOwner(&_ZetaInteractor.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_ZetaInteractor *ZetaInteractorCallerSession) PendingOwner() (common.Address, error) {
	return _ZetaInteractor.Contract.PendingOwner(&_ZetaInteractor.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ZetaInteractor *ZetaInteractorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaInteractor.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ZetaInteractor *ZetaInteractorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ZetaInteractor.Contract.AcceptOwnership(&_ZetaInteractor.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ZetaInteractor *ZetaInteractorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ZetaInteractor.Contract.AcceptOwnership(&_ZetaInteractor.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZetaInteractor *ZetaInteractorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaInteractor.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZetaInteractor *ZetaInteractorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ZetaInteractor.Contract.RenounceOwnership(&_ZetaInteractor.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZetaInteractor *ZetaInteractorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ZetaInteractor.Contract.RenounceOwnership(&_ZetaInteractor.TransactOpts)
}

// SetInteractorByChainId is a paid mutator transaction binding the contract method 0x4fd3f7d7.
//
// Solidity: function setInteractorByChainId(uint256 destinationChainId, bytes contractAddress) returns()
func (_ZetaInteractor *ZetaInteractorTransactor) SetInteractorByChainId(opts *bind.TransactOpts, destinationChainId *big.Int, contractAddress []byte) (*types.Transaction, error) {
	return _ZetaInteractor.contract.Transact(opts, "setInteractorByChainId", destinationChainId, contractAddress)
}

// SetInteractorByChainId is a paid mutator transaction binding the contract method 0x4fd3f7d7.
//
// Solidity: function setInteractorByChainId(uint256 destinationChainId, bytes contractAddress) returns()
func (_ZetaInteractor *ZetaInteractorSession) SetInteractorByChainId(destinationChainId *big.Int, contractAddress []byte) (*types.Transaction, error) {
	return _ZetaInteractor.Contract.SetInteractorByChainId(&_ZetaInteractor.TransactOpts, destinationChainId, contractAddress)
}

// SetInteractorByChainId is a paid mutator transaction binding the contract method 0x4fd3f7d7.
//
// Solidity: function setInteractorByChainId(uint256 destinationChainId, bytes contractAddress) returns()
func (_ZetaInteractor *ZetaInteractorTransactorSession) SetInteractorByChainId(destinationChainId *big.Int, contractAddress []byte) (*types.Transaction, error) {
	return _ZetaInteractor.Contract.SetInteractorByChainId(&_ZetaInteractor.TransactOpts, destinationChainId, contractAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZetaInteractor *ZetaInteractorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ZetaInteractor.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZetaInteractor *ZetaInteractorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZetaInteractor.Contract.TransferOwnership(&_ZetaInteractor.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZetaInteractor *ZetaInteractorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZetaInteractor.Contract.TransferOwnership(&_ZetaInteractor.TransactOpts, newOwner)
}

// ZetaInteractorOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the ZetaInteractor contract.
type ZetaInteractorOwnershipTransferStartedIterator struct {
	Event *ZetaInteractorOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *ZetaInteractorOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaInteractorOwnershipTransferStarted)
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
		it.Event = new(ZetaInteractorOwnershipTransferStarted)
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
func (it *ZetaInteractorOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaInteractorOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaInteractorOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the ZetaInteractor contract.
type ZetaInteractorOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_ZetaInteractor *ZetaInteractorFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ZetaInteractorOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZetaInteractor.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ZetaInteractorOwnershipTransferStartedIterator{contract: _ZetaInteractor.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_ZetaInteractor *ZetaInteractorFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *ZetaInteractorOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZetaInteractor.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaInteractorOwnershipTransferStarted)
				if err := _ZetaInteractor.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_ZetaInteractor *ZetaInteractorFilterer) ParseOwnershipTransferStarted(log types.Log) (*ZetaInteractorOwnershipTransferStarted, error) {
	event := new(ZetaInteractorOwnershipTransferStarted)
	if err := _ZetaInteractor.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaInteractorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ZetaInteractor contract.
type ZetaInteractorOwnershipTransferredIterator struct {
	Event *ZetaInteractorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ZetaInteractorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaInteractorOwnershipTransferred)
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
		it.Event = new(ZetaInteractorOwnershipTransferred)
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
func (it *ZetaInteractorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaInteractorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaInteractorOwnershipTransferred represents a OwnershipTransferred event raised by the ZetaInteractor contract.
type ZetaInteractorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZetaInteractor *ZetaInteractorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ZetaInteractorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZetaInteractor.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ZetaInteractorOwnershipTransferredIterator{contract: _ZetaInteractor.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZetaInteractor *ZetaInteractorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ZetaInteractorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZetaInteractor.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaInteractorOwnershipTransferred)
				if err := _ZetaInteractor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ZetaInteractor *ZetaInteractorFilterer) ParseOwnershipTransferred(log types.Log) (*ZetaInteractorOwnershipTransferred, error) {
	event := new(ZetaInteractorOwnershipTransferred)
	if err := _ZetaInteractor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
