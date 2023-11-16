// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetainterfaces

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

// ZetaTokenConsumerMetaData contains all meta data concerning the ZetaTokenConsumer contract.
var ZetaTokenConsumerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"EthExchangedForZeta\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"TokenExchangedForZeta\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"ZetaExchangedForEth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"ZetaExchangedForToken\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zetaTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getEthFromZeta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"zetaTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getTokenFromZeta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"name\":\"getZetaFromEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getZetaFromToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasZetaLiquidity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ZetaTokenConsumerABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaTokenConsumerMetaData.ABI instead.
var ZetaTokenConsumerABI = ZetaTokenConsumerMetaData.ABI

// ZetaTokenConsumer is an auto generated Go binding around an Ethereum contract.
type ZetaTokenConsumer struct {
	ZetaTokenConsumerCaller     // Read-only binding to the contract
	ZetaTokenConsumerTransactor // Write-only binding to the contract
	ZetaTokenConsumerFilterer   // Log filterer for contract events
}

// ZetaTokenConsumerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaTokenConsumerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaTokenConsumerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaTokenConsumerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaTokenConsumerSession struct {
	Contract     *ZetaTokenConsumer // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ZetaTokenConsumerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaTokenConsumerCallerSession struct {
	Contract *ZetaTokenConsumerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ZetaTokenConsumerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaTokenConsumerTransactorSession struct {
	Contract     *ZetaTokenConsumerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ZetaTokenConsumerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaTokenConsumerRaw struct {
	Contract *ZetaTokenConsumer // Generic contract binding to access the raw methods on
}

// ZetaTokenConsumerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaTokenConsumerCallerRaw struct {
	Contract *ZetaTokenConsumerCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaTokenConsumerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaTokenConsumerTransactorRaw struct {
	Contract *ZetaTokenConsumerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaTokenConsumer creates a new instance of ZetaTokenConsumer, bound to a specific deployed contract.
func NewZetaTokenConsumer(address common.Address, backend bind.ContractBackend) (*ZetaTokenConsumer, error) {
	contract, err := bindZetaTokenConsumer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumer{ZetaTokenConsumerCaller: ZetaTokenConsumerCaller{contract: contract}, ZetaTokenConsumerTransactor: ZetaTokenConsumerTransactor{contract: contract}, ZetaTokenConsumerFilterer: ZetaTokenConsumerFilterer{contract: contract}}, nil
}

// NewZetaTokenConsumerCaller creates a new read-only instance of ZetaTokenConsumer, bound to a specific deployed contract.
func NewZetaTokenConsumerCaller(address common.Address, caller bind.ContractCaller) (*ZetaTokenConsumerCaller, error) {
	contract, err := bindZetaTokenConsumer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerCaller{contract: contract}, nil
}

// NewZetaTokenConsumerTransactor creates a new write-only instance of ZetaTokenConsumer, bound to a specific deployed contract.
func NewZetaTokenConsumerTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaTokenConsumerTransactor, error) {
	contract, err := bindZetaTokenConsumer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerTransactor{contract: contract}, nil
}

// NewZetaTokenConsumerFilterer creates a new log filterer instance of ZetaTokenConsumer, bound to a specific deployed contract.
func NewZetaTokenConsumerFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaTokenConsumerFilterer, error) {
	contract, err := bindZetaTokenConsumer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerFilterer{contract: contract}, nil
}

// bindZetaTokenConsumer binds a generic wrapper to an already deployed contract.
func bindZetaTokenConsumer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaTokenConsumerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaTokenConsumer *ZetaTokenConsumerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaTokenConsumer.Contract.ZetaTokenConsumerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaTokenConsumer *ZetaTokenConsumerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaTokenConsumer.Contract.ZetaTokenConsumerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaTokenConsumer *ZetaTokenConsumerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaTokenConsumer.Contract.ZetaTokenConsumerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaTokenConsumer *ZetaTokenConsumerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaTokenConsumer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaTokenConsumer *ZetaTokenConsumerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaTokenConsumer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaTokenConsumer *ZetaTokenConsumerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaTokenConsumer.Contract.contract.Transact(opts, method, params...)
}

// HasZetaLiquidity is a free data retrieval call binding the contract method 0x80801f84.
//
// Solidity: function hasZetaLiquidity() view returns(bool)
func (_ZetaTokenConsumer *ZetaTokenConsumerCaller) HasZetaLiquidity(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZetaTokenConsumer.contract.Call(opts, &out, "hasZetaLiquidity")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasZetaLiquidity is a free data retrieval call binding the contract method 0x80801f84.
//
// Solidity: function hasZetaLiquidity() view returns(bool)
func (_ZetaTokenConsumer *ZetaTokenConsumerSession) HasZetaLiquidity() (bool, error) {
	return _ZetaTokenConsumer.Contract.HasZetaLiquidity(&_ZetaTokenConsumer.CallOpts)
}

// HasZetaLiquidity is a free data retrieval call binding the contract method 0x80801f84.
//
// Solidity: function hasZetaLiquidity() view returns(bool)
func (_ZetaTokenConsumer *ZetaTokenConsumerCallerSession) HasZetaLiquidity() (bool, error) {
	return _ZetaTokenConsumer.Contract.HasZetaLiquidity(&_ZetaTokenConsumer.CallOpts)
}

// GetEthFromZeta is a paid mutator transaction binding the contract method 0x54c49a2a.
//
// Solidity: function getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumer *ZetaTokenConsumerTransactor) GetEthFromZeta(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumer.contract.Transact(opts, "getEthFromZeta", destinationAddress, minAmountOut, zetaTokenAmount)
}

// GetEthFromZeta is a paid mutator transaction binding the contract method 0x54c49a2a.
//
// Solidity: function getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumer *ZetaTokenConsumerSession) GetEthFromZeta(destinationAddress common.Address, minAmountOut *big.Int, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumer.Contract.GetEthFromZeta(&_ZetaTokenConsumer.TransactOpts, destinationAddress, minAmountOut, zetaTokenAmount)
}

// GetEthFromZeta is a paid mutator transaction binding the contract method 0x54c49a2a.
//
// Solidity: function getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumer *ZetaTokenConsumerTransactorSession) GetEthFromZeta(destinationAddress common.Address, minAmountOut *big.Int, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumer.Contract.GetEthFromZeta(&_ZetaTokenConsumer.TransactOpts, destinationAddress, minAmountOut, zetaTokenAmount)
}

// GetTokenFromZeta is a paid mutator transaction binding the contract method 0xa53fb10b.
//
// Solidity: function getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumer *ZetaTokenConsumerTransactor) GetTokenFromZeta(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumer.contract.Transact(opts, "getTokenFromZeta", destinationAddress, minAmountOut, outputToken, zetaTokenAmount)
}

// GetTokenFromZeta is a paid mutator transaction binding the contract method 0xa53fb10b.
//
// Solidity: function getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumer *ZetaTokenConsumerSession) GetTokenFromZeta(destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumer.Contract.GetTokenFromZeta(&_ZetaTokenConsumer.TransactOpts, destinationAddress, minAmountOut, outputToken, zetaTokenAmount)
}

// GetTokenFromZeta is a paid mutator transaction binding the contract method 0xa53fb10b.
//
// Solidity: function getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumer *ZetaTokenConsumerTransactorSession) GetTokenFromZeta(destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumer.Contract.GetTokenFromZeta(&_ZetaTokenConsumer.TransactOpts, destinationAddress, minAmountOut, outputToken, zetaTokenAmount)
}

// GetZetaFromEth is a paid mutator transaction binding the contract method 0x013b2ff8.
//
// Solidity: function getZetaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_ZetaTokenConsumer *ZetaTokenConsumerTransactor) GetZetaFromEth(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumer.contract.Transact(opts, "getZetaFromEth", destinationAddress, minAmountOut)
}

// GetZetaFromEth is a paid mutator transaction binding the contract method 0x013b2ff8.
//
// Solidity: function getZetaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_ZetaTokenConsumer *ZetaTokenConsumerSession) GetZetaFromEth(destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumer.Contract.GetZetaFromEth(&_ZetaTokenConsumer.TransactOpts, destinationAddress, minAmountOut)
}

// GetZetaFromEth is a paid mutator transaction binding the contract method 0x013b2ff8.
//
// Solidity: function getZetaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_ZetaTokenConsumer *ZetaTokenConsumerTransactorSession) GetZetaFromEth(destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumer.Contract.GetZetaFromEth(&_ZetaTokenConsumer.TransactOpts, destinationAddress, minAmountOut)
}

// GetZetaFromToken is a paid mutator transaction binding the contract method 0x2405620a.
//
// Solidity: function getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_ZetaTokenConsumer *ZetaTokenConsumerTransactor) GetZetaFromToken(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumer.contract.Transact(opts, "getZetaFromToken", destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// GetZetaFromToken is a paid mutator transaction binding the contract method 0x2405620a.
//
// Solidity: function getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_ZetaTokenConsumer *ZetaTokenConsumerSession) GetZetaFromToken(destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumer.Contract.GetZetaFromToken(&_ZetaTokenConsumer.TransactOpts, destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// GetZetaFromToken is a paid mutator transaction binding the contract method 0x2405620a.
//
// Solidity: function getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_ZetaTokenConsumer *ZetaTokenConsumerTransactorSession) GetZetaFromToken(destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumer.Contract.GetZetaFromToken(&_ZetaTokenConsumer.TransactOpts, destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// ZetaTokenConsumerEthExchangedForZetaIterator is returned from FilterEthExchangedForZeta and is used to iterate over the raw logs and unpacked data for EthExchangedForZeta events raised by the ZetaTokenConsumer contract.
type ZetaTokenConsumerEthExchangedForZetaIterator struct {
	Event *ZetaTokenConsumerEthExchangedForZeta // Event containing the contract specifics and raw log

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
func (it *ZetaTokenConsumerEthExchangedForZetaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerEthExchangedForZeta)
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
		it.Event = new(ZetaTokenConsumerEthExchangedForZeta)
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
func (it *ZetaTokenConsumerEthExchangedForZetaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerEthExchangedForZetaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerEthExchangedForZeta represents a EthExchangedForZeta event raised by the ZetaTokenConsumer contract.
type ZetaTokenConsumerEthExchangedForZeta struct {
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEthExchangedForZeta is a free log retrieval operation binding the contract event 0x87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da11.
//
// Solidity: event EthExchangedForZeta(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumer *ZetaTokenConsumerFilterer) FilterEthExchangedForZeta(opts *bind.FilterOpts) (*ZetaTokenConsumerEthExchangedForZetaIterator, error) {

	logs, sub, err := _ZetaTokenConsumer.contract.FilterLogs(opts, "EthExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerEthExchangedForZetaIterator{contract: _ZetaTokenConsumer.contract, event: "EthExchangedForZeta", logs: logs, sub: sub}, nil
}

// WatchEthExchangedForZeta is a free log subscription operation binding the contract event 0x87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da11.
//
// Solidity: event EthExchangedForZeta(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumer *ZetaTokenConsumerFilterer) WatchEthExchangedForZeta(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerEthExchangedForZeta) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumer.contract.WatchLogs(opts, "EthExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerEthExchangedForZeta)
				if err := _ZetaTokenConsumer.contract.UnpackLog(event, "EthExchangedForZeta", log); err != nil {
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

// ParseEthExchangedForZeta is a log parse operation binding the contract event 0x87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da11.
//
// Solidity: event EthExchangedForZeta(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumer *ZetaTokenConsumerFilterer) ParseEthExchangedForZeta(log types.Log) (*ZetaTokenConsumerEthExchangedForZeta, error) {
	event := new(ZetaTokenConsumerEthExchangedForZeta)
	if err := _ZetaTokenConsumer.contract.UnpackLog(event, "EthExchangedForZeta", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaTokenConsumerTokenExchangedForZetaIterator is returned from FilterTokenExchangedForZeta and is used to iterate over the raw logs and unpacked data for TokenExchangedForZeta events raised by the ZetaTokenConsumer contract.
type ZetaTokenConsumerTokenExchangedForZetaIterator struct {
	Event *ZetaTokenConsumerTokenExchangedForZeta // Event containing the contract specifics and raw log

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
func (it *ZetaTokenConsumerTokenExchangedForZetaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerTokenExchangedForZeta)
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
		it.Event = new(ZetaTokenConsumerTokenExchangedForZeta)
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
func (it *ZetaTokenConsumerTokenExchangedForZetaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerTokenExchangedForZetaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerTokenExchangedForZeta represents a TokenExchangedForZeta event raised by the ZetaTokenConsumer contract.
type ZetaTokenConsumerTokenExchangedForZeta struct {
	Token     common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenExchangedForZeta is a free log retrieval operation binding the contract event 0x017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f.
//
// Solidity: event TokenExchangedForZeta(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumer *ZetaTokenConsumerFilterer) FilterTokenExchangedForZeta(opts *bind.FilterOpts) (*ZetaTokenConsumerTokenExchangedForZetaIterator, error) {

	logs, sub, err := _ZetaTokenConsumer.contract.FilterLogs(opts, "TokenExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerTokenExchangedForZetaIterator{contract: _ZetaTokenConsumer.contract, event: "TokenExchangedForZeta", logs: logs, sub: sub}, nil
}

// WatchTokenExchangedForZeta is a free log subscription operation binding the contract event 0x017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f.
//
// Solidity: event TokenExchangedForZeta(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumer *ZetaTokenConsumerFilterer) WatchTokenExchangedForZeta(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerTokenExchangedForZeta) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumer.contract.WatchLogs(opts, "TokenExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerTokenExchangedForZeta)
				if err := _ZetaTokenConsumer.contract.UnpackLog(event, "TokenExchangedForZeta", log); err != nil {
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

// ParseTokenExchangedForZeta is a log parse operation binding the contract event 0x017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f.
//
// Solidity: event TokenExchangedForZeta(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumer *ZetaTokenConsumerFilterer) ParseTokenExchangedForZeta(log types.Log) (*ZetaTokenConsumerTokenExchangedForZeta, error) {
	event := new(ZetaTokenConsumerTokenExchangedForZeta)
	if err := _ZetaTokenConsumer.contract.UnpackLog(event, "TokenExchangedForZeta", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaTokenConsumerZetaExchangedForEthIterator is returned from FilterZetaExchangedForEth and is used to iterate over the raw logs and unpacked data for ZetaExchangedForEth events raised by the ZetaTokenConsumer contract.
type ZetaTokenConsumerZetaExchangedForEthIterator struct {
	Event *ZetaTokenConsumerZetaExchangedForEth // Event containing the contract specifics and raw log

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
func (it *ZetaTokenConsumerZetaExchangedForEthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerZetaExchangedForEth)
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
		it.Event = new(ZetaTokenConsumerZetaExchangedForEth)
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
func (it *ZetaTokenConsumerZetaExchangedForEthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerZetaExchangedForEthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerZetaExchangedForEth represents a ZetaExchangedForEth event raised by the ZetaTokenConsumer contract.
type ZetaTokenConsumerZetaExchangedForEth struct {
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterZetaExchangedForEth is a free log retrieval operation binding the contract event 0x74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae.
//
// Solidity: event ZetaExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumer *ZetaTokenConsumerFilterer) FilterZetaExchangedForEth(opts *bind.FilterOpts) (*ZetaTokenConsumerZetaExchangedForEthIterator, error) {

	logs, sub, err := _ZetaTokenConsumer.contract.FilterLogs(opts, "ZetaExchangedForEth")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerZetaExchangedForEthIterator{contract: _ZetaTokenConsumer.contract, event: "ZetaExchangedForEth", logs: logs, sub: sub}, nil
}

// WatchZetaExchangedForEth is a free log subscription operation binding the contract event 0x74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae.
//
// Solidity: event ZetaExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumer *ZetaTokenConsumerFilterer) WatchZetaExchangedForEth(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerZetaExchangedForEth) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumer.contract.WatchLogs(opts, "ZetaExchangedForEth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerZetaExchangedForEth)
				if err := _ZetaTokenConsumer.contract.UnpackLog(event, "ZetaExchangedForEth", log); err != nil {
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

// ParseZetaExchangedForEth is a log parse operation binding the contract event 0x74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae.
//
// Solidity: event ZetaExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumer *ZetaTokenConsumerFilterer) ParseZetaExchangedForEth(log types.Log) (*ZetaTokenConsumerZetaExchangedForEth, error) {
	event := new(ZetaTokenConsumerZetaExchangedForEth)
	if err := _ZetaTokenConsumer.contract.UnpackLog(event, "ZetaExchangedForEth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaTokenConsumerZetaExchangedForTokenIterator is returned from FilterZetaExchangedForToken and is used to iterate over the raw logs and unpacked data for ZetaExchangedForToken events raised by the ZetaTokenConsumer contract.
type ZetaTokenConsumerZetaExchangedForTokenIterator struct {
	Event *ZetaTokenConsumerZetaExchangedForToken // Event containing the contract specifics and raw log

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
func (it *ZetaTokenConsumerZetaExchangedForTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerZetaExchangedForToken)
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
		it.Event = new(ZetaTokenConsumerZetaExchangedForToken)
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
func (it *ZetaTokenConsumerZetaExchangedForTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerZetaExchangedForTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerZetaExchangedForToken represents a ZetaExchangedForToken event raised by the ZetaTokenConsumer contract.
type ZetaTokenConsumerZetaExchangedForToken struct {
	Token     common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterZetaExchangedForToken is a free log retrieval operation binding the contract event 0x0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b.
//
// Solidity: event ZetaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumer *ZetaTokenConsumerFilterer) FilterZetaExchangedForToken(opts *bind.FilterOpts) (*ZetaTokenConsumerZetaExchangedForTokenIterator, error) {

	logs, sub, err := _ZetaTokenConsumer.contract.FilterLogs(opts, "ZetaExchangedForToken")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerZetaExchangedForTokenIterator{contract: _ZetaTokenConsumer.contract, event: "ZetaExchangedForToken", logs: logs, sub: sub}, nil
}

// WatchZetaExchangedForToken is a free log subscription operation binding the contract event 0x0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b.
//
// Solidity: event ZetaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumer *ZetaTokenConsumerFilterer) WatchZetaExchangedForToken(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerZetaExchangedForToken) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumer.contract.WatchLogs(opts, "ZetaExchangedForToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerZetaExchangedForToken)
				if err := _ZetaTokenConsumer.contract.UnpackLog(event, "ZetaExchangedForToken", log); err != nil {
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

// ParseZetaExchangedForToken is a log parse operation binding the contract event 0x0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b.
//
// Solidity: event ZetaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumer *ZetaTokenConsumerFilterer) ParseZetaExchangedForToken(log types.Log) (*ZetaTokenConsumerZetaExchangedForToken, error) {
	event := new(ZetaTokenConsumerZetaExchangedForToken)
	if err := _ZetaTokenConsumer.contract.UnpackLog(event, "ZetaExchangedForToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
