// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetareceivermock

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

// ZetaInterfacesZetaMessage is an auto generated low-level Go binding around an user-defined struct.
type ZetaInterfacesZetaMessage struct {
	ZetaTxSenderAddress []byte
	SourceChainId       *big.Int
	DestinationAddress  common.Address
	ZetaValue           *big.Int
	Message             []byte
}

// ZetaInterfacesZetaRevert is an auto generated low-level Go binding around an user-defined struct.
type ZetaInterfacesZetaRevert struct {
	ZetaTxSenderAddress common.Address
	SourceChainId       *big.Int
	DestinationAddress  []byte
	DestinationChainId  *big.Int
	RemainingZetaValue  *big.Int
	Message             []byte
}

// ZetaReceiverMockMetaData contains all meta data concerning the ZetaReceiverMock contract.
var ZetaReceiverMockMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"}],\"name\":\"MockOnZetaMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"zetaTxSenderAddress\",\"type\":\"address\"}],\"name\":\"MockOnZetaRevert\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"zetaTxSenderAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"zetaValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structZetaInterfaces.ZetaMessage\",\"name\":\"zetaMessage\",\"type\":\"tuple\"}],\"name\":\"onZetaMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"zetaTxSenderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingZetaValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structZetaInterfaces.ZetaRevert\",\"name\":\"zetaRevert\",\"type\":\"tuple\"}],\"name\":\"onZetaRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506102d5806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80633749c51a1461003b5780633ff0693c14610057575b600080fd5b6100556004803603810190610050919061018b565b610073565b005b610071600480360381019061006c91906101d4565b6100bf565b005b7f72a301dee3abcbe15615f3e253229bf4b4f508460603d674991c9a777b833c6e8160400160208101906100a7919061015e565b6040516100b4919061022c565b60405180910390a150565b7f53bd04e26f94f13ff43da96839541821041c309c6f624712192cbe3a2d133cc48160000160208101906100f3919061015e565b604051610100919061022c565b60405180910390a150565b60008135905061011a81610288565b92915050565b600060a0828403121561013657610135610279565b5b81905092915050565b600060c0828403121561015557610154610279565b5b81905092915050565b60006020828403121561017457610173610283565b5b60006101828482850161010b565b91505092915050565b6000602082840312156101a1576101a0610283565b5b600082013567ffffffffffffffff8111156101bf576101be61027e565b5b6101cb84828501610120565b91505092915050565b6000602082840312156101ea576101e9610283565b5b600082013567ffffffffffffffff8111156102085761020761027e565b5b6102148482850161013f565b91505092915050565b61022681610247565b82525050565b6000602082019050610241600083018461021d565b92915050565b600061025282610259565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b600080fd5b600080fd5b61029181610247565b811461029c57600080fd5b5056fea264697066735822122008485e037d9a20d8e9f8d1e9456b89006367d84f7e0966e1d820fe73c0d706ea64736f6c63430008070033",
}

// ZetaReceiverMockABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaReceiverMockMetaData.ABI instead.
var ZetaReceiverMockABI = ZetaReceiverMockMetaData.ABI

// ZetaReceiverMockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZetaReceiverMockMetaData.Bin instead.
var ZetaReceiverMockBin = ZetaReceiverMockMetaData.Bin

// DeployZetaReceiverMock deploys a new Ethereum contract, binding an instance of ZetaReceiverMock to it.
func DeployZetaReceiverMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZetaReceiverMock, error) {
	parsed, err := ZetaReceiverMockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZetaReceiverMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZetaReceiverMock{ZetaReceiverMockCaller: ZetaReceiverMockCaller{contract: contract}, ZetaReceiverMockTransactor: ZetaReceiverMockTransactor{contract: contract}, ZetaReceiverMockFilterer: ZetaReceiverMockFilterer{contract: contract}}, nil
}

// ZetaReceiverMock is an auto generated Go binding around an Ethereum contract.
type ZetaReceiverMock struct {
	ZetaReceiverMockCaller     // Read-only binding to the contract
	ZetaReceiverMockTransactor // Write-only binding to the contract
	ZetaReceiverMockFilterer   // Log filterer for contract events
}

// ZetaReceiverMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaReceiverMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaReceiverMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaReceiverMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaReceiverMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaReceiverMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaReceiverMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaReceiverMockSession struct {
	Contract     *ZetaReceiverMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZetaReceiverMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaReceiverMockCallerSession struct {
	Contract *ZetaReceiverMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ZetaReceiverMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaReceiverMockTransactorSession struct {
	Contract     *ZetaReceiverMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ZetaReceiverMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaReceiverMockRaw struct {
	Contract *ZetaReceiverMock // Generic contract binding to access the raw methods on
}

// ZetaReceiverMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaReceiverMockCallerRaw struct {
	Contract *ZetaReceiverMockCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaReceiverMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaReceiverMockTransactorRaw struct {
	Contract *ZetaReceiverMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaReceiverMock creates a new instance of ZetaReceiverMock, bound to a specific deployed contract.
func NewZetaReceiverMock(address common.Address, backend bind.ContractBackend) (*ZetaReceiverMock, error) {
	contract, err := bindZetaReceiverMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaReceiverMock{ZetaReceiverMockCaller: ZetaReceiverMockCaller{contract: contract}, ZetaReceiverMockTransactor: ZetaReceiverMockTransactor{contract: contract}, ZetaReceiverMockFilterer: ZetaReceiverMockFilterer{contract: contract}}, nil
}

// NewZetaReceiverMockCaller creates a new read-only instance of ZetaReceiverMock, bound to a specific deployed contract.
func NewZetaReceiverMockCaller(address common.Address, caller bind.ContractCaller) (*ZetaReceiverMockCaller, error) {
	contract, err := bindZetaReceiverMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaReceiverMockCaller{contract: contract}, nil
}

// NewZetaReceiverMockTransactor creates a new write-only instance of ZetaReceiverMock, bound to a specific deployed contract.
func NewZetaReceiverMockTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaReceiverMockTransactor, error) {
	contract, err := bindZetaReceiverMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaReceiverMockTransactor{contract: contract}, nil
}

// NewZetaReceiverMockFilterer creates a new log filterer instance of ZetaReceiverMock, bound to a specific deployed contract.
func NewZetaReceiverMockFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaReceiverMockFilterer, error) {
	contract, err := bindZetaReceiverMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaReceiverMockFilterer{contract: contract}, nil
}

// bindZetaReceiverMock binds a generic wrapper to an already deployed contract.
func bindZetaReceiverMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaReceiverMockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaReceiverMock *ZetaReceiverMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaReceiverMock.Contract.ZetaReceiverMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaReceiverMock *ZetaReceiverMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaReceiverMock.Contract.ZetaReceiverMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaReceiverMock *ZetaReceiverMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaReceiverMock.Contract.ZetaReceiverMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaReceiverMock *ZetaReceiverMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaReceiverMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaReceiverMock *ZetaReceiverMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaReceiverMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaReceiverMock *ZetaReceiverMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaReceiverMock.Contract.contract.Transact(opts, method, params...)
}

// OnZetaMessage is a paid mutator transaction binding the contract method 0x3749c51a.
//
// Solidity: function onZetaMessage((bytes,uint256,address,uint256,bytes) zetaMessage) returns()
func (_ZetaReceiverMock *ZetaReceiverMockTransactor) OnZetaMessage(opts *bind.TransactOpts, zetaMessage ZetaInterfacesZetaMessage) (*types.Transaction, error) {
	return _ZetaReceiverMock.contract.Transact(opts, "onZetaMessage", zetaMessage)
}

// OnZetaMessage is a paid mutator transaction binding the contract method 0x3749c51a.
//
// Solidity: function onZetaMessage((bytes,uint256,address,uint256,bytes) zetaMessage) returns()
func (_ZetaReceiverMock *ZetaReceiverMockSession) OnZetaMessage(zetaMessage ZetaInterfacesZetaMessage) (*types.Transaction, error) {
	return _ZetaReceiverMock.Contract.OnZetaMessage(&_ZetaReceiverMock.TransactOpts, zetaMessage)
}

// OnZetaMessage is a paid mutator transaction binding the contract method 0x3749c51a.
//
// Solidity: function onZetaMessage((bytes,uint256,address,uint256,bytes) zetaMessage) returns()
func (_ZetaReceiverMock *ZetaReceiverMockTransactorSession) OnZetaMessage(zetaMessage ZetaInterfacesZetaMessage) (*types.Transaction, error) {
	return _ZetaReceiverMock.Contract.OnZetaMessage(&_ZetaReceiverMock.TransactOpts, zetaMessage)
}

// OnZetaRevert is a paid mutator transaction binding the contract method 0x3ff0693c.
//
// Solidity: function onZetaRevert((address,uint256,bytes,uint256,uint256,bytes) zetaRevert) returns()
func (_ZetaReceiverMock *ZetaReceiverMockTransactor) OnZetaRevert(opts *bind.TransactOpts, zetaRevert ZetaInterfacesZetaRevert) (*types.Transaction, error) {
	return _ZetaReceiverMock.contract.Transact(opts, "onZetaRevert", zetaRevert)
}

// OnZetaRevert is a paid mutator transaction binding the contract method 0x3ff0693c.
//
// Solidity: function onZetaRevert((address,uint256,bytes,uint256,uint256,bytes) zetaRevert) returns()
func (_ZetaReceiverMock *ZetaReceiverMockSession) OnZetaRevert(zetaRevert ZetaInterfacesZetaRevert) (*types.Transaction, error) {
	return _ZetaReceiverMock.Contract.OnZetaRevert(&_ZetaReceiverMock.TransactOpts, zetaRevert)
}

// OnZetaRevert is a paid mutator transaction binding the contract method 0x3ff0693c.
//
// Solidity: function onZetaRevert((address,uint256,bytes,uint256,uint256,bytes) zetaRevert) returns()
func (_ZetaReceiverMock *ZetaReceiverMockTransactorSession) OnZetaRevert(zetaRevert ZetaInterfacesZetaRevert) (*types.Transaction, error) {
	return _ZetaReceiverMock.Contract.OnZetaRevert(&_ZetaReceiverMock.TransactOpts, zetaRevert)
}

// ZetaReceiverMockMockOnZetaMessageIterator is returned from FilterMockOnZetaMessage and is used to iterate over the raw logs and unpacked data for MockOnZetaMessage events raised by the ZetaReceiverMock contract.
type ZetaReceiverMockMockOnZetaMessageIterator struct {
	Event *ZetaReceiverMockMockOnZetaMessage // Event containing the contract specifics and raw log

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
func (it *ZetaReceiverMockMockOnZetaMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaReceiverMockMockOnZetaMessage)
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
		it.Event = new(ZetaReceiverMockMockOnZetaMessage)
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
func (it *ZetaReceiverMockMockOnZetaMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaReceiverMockMockOnZetaMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaReceiverMockMockOnZetaMessage represents a MockOnZetaMessage event raised by the ZetaReceiverMock contract.
type ZetaReceiverMockMockOnZetaMessage struct {
	DestinationAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMockOnZetaMessage is a free log retrieval operation binding the contract event 0x72a301dee3abcbe15615f3e253229bf4b4f508460603d674991c9a777b833c6e.
//
// Solidity: event MockOnZetaMessage(address destinationAddress)
func (_ZetaReceiverMock *ZetaReceiverMockFilterer) FilterMockOnZetaMessage(opts *bind.FilterOpts) (*ZetaReceiverMockMockOnZetaMessageIterator, error) {

	logs, sub, err := _ZetaReceiverMock.contract.FilterLogs(opts, "MockOnZetaMessage")
	if err != nil {
		return nil, err
	}
	return &ZetaReceiverMockMockOnZetaMessageIterator{contract: _ZetaReceiverMock.contract, event: "MockOnZetaMessage", logs: logs, sub: sub}, nil
}

// WatchMockOnZetaMessage is a free log subscription operation binding the contract event 0x72a301dee3abcbe15615f3e253229bf4b4f508460603d674991c9a777b833c6e.
//
// Solidity: event MockOnZetaMessage(address destinationAddress)
func (_ZetaReceiverMock *ZetaReceiverMockFilterer) WatchMockOnZetaMessage(opts *bind.WatchOpts, sink chan<- *ZetaReceiverMockMockOnZetaMessage) (event.Subscription, error) {

	logs, sub, err := _ZetaReceiverMock.contract.WatchLogs(opts, "MockOnZetaMessage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaReceiverMockMockOnZetaMessage)
				if err := _ZetaReceiverMock.contract.UnpackLog(event, "MockOnZetaMessage", log); err != nil {
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

// ParseMockOnZetaMessage is a log parse operation binding the contract event 0x72a301dee3abcbe15615f3e253229bf4b4f508460603d674991c9a777b833c6e.
//
// Solidity: event MockOnZetaMessage(address destinationAddress)
func (_ZetaReceiverMock *ZetaReceiverMockFilterer) ParseMockOnZetaMessage(log types.Log) (*ZetaReceiverMockMockOnZetaMessage, error) {
	event := new(ZetaReceiverMockMockOnZetaMessage)
	if err := _ZetaReceiverMock.contract.UnpackLog(event, "MockOnZetaMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaReceiverMockMockOnZetaRevertIterator is returned from FilterMockOnZetaRevert and is used to iterate over the raw logs and unpacked data for MockOnZetaRevert events raised by the ZetaReceiverMock contract.
type ZetaReceiverMockMockOnZetaRevertIterator struct {
	Event *ZetaReceiverMockMockOnZetaRevert // Event containing the contract specifics and raw log

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
func (it *ZetaReceiverMockMockOnZetaRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaReceiverMockMockOnZetaRevert)
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
		it.Event = new(ZetaReceiverMockMockOnZetaRevert)
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
func (it *ZetaReceiverMockMockOnZetaRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaReceiverMockMockOnZetaRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaReceiverMockMockOnZetaRevert represents a MockOnZetaRevert event raised by the ZetaReceiverMock contract.
type ZetaReceiverMockMockOnZetaRevert struct {
	ZetaTxSenderAddress common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMockOnZetaRevert is a free log retrieval operation binding the contract event 0x53bd04e26f94f13ff43da96839541821041c309c6f624712192cbe3a2d133cc4.
//
// Solidity: event MockOnZetaRevert(address zetaTxSenderAddress)
func (_ZetaReceiverMock *ZetaReceiverMockFilterer) FilterMockOnZetaRevert(opts *bind.FilterOpts) (*ZetaReceiverMockMockOnZetaRevertIterator, error) {

	logs, sub, err := _ZetaReceiverMock.contract.FilterLogs(opts, "MockOnZetaRevert")
	if err != nil {
		return nil, err
	}
	return &ZetaReceiverMockMockOnZetaRevertIterator{contract: _ZetaReceiverMock.contract, event: "MockOnZetaRevert", logs: logs, sub: sub}, nil
}

// WatchMockOnZetaRevert is a free log subscription operation binding the contract event 0x53bd04e26f94f13ff43da96839541821041c309c6f624712192cbe3a2d133cc4.
//
// Solidity: event MockOnZetaRevert(address zetaTxSenderAddress)
func (_ZetaReceiverMock *ZetaReceiverMockFilterer) WatchMockOnZetaRevert(opts *bind.WatchOpts, sink chan<- *ZetaReceiverMockMockOnZetaRevert) (event.Subscription, error) {

	logs, sub, err := _ZetaReceiverMock.contract.WatchLogs(opts, "MockOnZetaRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaReceiverMockMockOnZetaRevert)
				if err := _ZetaReceiverMock.contract.UnpackLog(event, "MockOnZetaRevert", log); err != nil {
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

// ParseMockOnZetaRevert is a log parse operation binding the contract event 0x53bd04e26f94f13ff43da96839541821041c309c6f624712192cbe3a2d133cc4.
//
// Solidity: event MockOnZetaRevert(address zetaTxSenderAddress)
func (_ZetaReceiverMock *ZetaReceiverMockFilterer) ParseMockOnZetaRevert(log types.Log) (*ZetaReceiverMockMockOnZetaRevert, error) {
	event := new(ZetaReceiverMockMockOnZetaRevert)
	if err := _ZetaReceiverMock.contract.UnpackLog(event, "MockOnZetaRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
