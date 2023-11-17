// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetatokenconsumerpancakev3

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

// WETH9MetaData contains all meta data concerning the WETH9 contract.
var WETH9MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WETH9ABI is the input ABI used to generate the binding from.
// Deprecated: Use WETH9MetaData.ABI instead.
var WETH9ABI = WETH9MetaData.ABI

// WETH9 is an auto generated Go binding around an Ethereum contract.
type WETH9 struct {
	WETH9Caller     // Read-only binding to the contract
	WETH9Transactor // Write-only binding to the contract
	WETH9Filterer   // Log filterer for contract events
}

// WETH9Caller is an auto generated read-only Go binding around an Ethereum contract.
type WETH9Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WETH9Transactor is an auto generated write-only Go binding around an Ethereum contract.
type WETH9Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WETH9Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WETH9Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WETH9Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WETH9Session struct {
	Contract     *WETH9            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WETH9CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WETH9CallerSession struct {
	Contract *WETH9Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WETH9TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WETH9TransactorSession struct {
	Contract     *WETH9Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WETH9Raw is an auto generated low-level Go binding around an Ethereum contract.
type WETH9Raw struct {
	Contract *WETH9 // Generic contract binding to access the raw methods on
}

// WETH9CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WETH9CallerRaw struct {
	Contract *WETH9Caller // Generic read-only contract binding to access the raw methods on
}

// WETH9TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WETH9TransactorRaw struct {
	Contract *WETH9Transactor // Generic write-only contract binding to access the raw methods on
}

// NewWETH9 creates a new instance of WETH9, bound to a specific deployed contract.
func NewWETH9(address common.Address, backend bind.ContractBackend) (*WETH9, error) {
	contract, err := bindWETH9(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WETH9{WETH9Caller: WETH9Caller{contract: contract}, WETH9Transactor: WETH9Transactor{contract: contract}, WETH9Filterer: WETH9Filterer{contract: contract}}, nil
}

// NewWETH9Caller creates a new read-only instance of WETH9, bound to a specific deployed contract.
func NewWETH9Caller(address common.Address, caller bind.ContractCaller) (*WETH9Caller, error) {
	contract, err := bindWETH9(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WETH9Caller{contract: contract}, nil
}

// NewWETH9Transactor creates a new write-only instance of WETH9, bound to a specific deployed contract.
func NewWETH9Transactor(address common.Address, transactor bind.ContractTransactor) (*WETH9Transactor, error) {
	contract, err := bindWETH9(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WETH9Transactor{contract: contract}, nil
}

// NewWETH9Filterer creates a new log filterer instance of WETH9, bound to a specific deployed contract.
func NewWETH9Filterer(address common.Address, filterer bind.ContractFilterer) (*WETH9Filterer, error) {
	contract, err := bindWETH9(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WETH9Filterer{contract: contract}, nil
}

// bindWETH9 binds a generic wrapper to an already deployed contract.
func bindWETH9(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WETH9MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WETH9 *WETH9Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WETH9.Contract.WETH9Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WETH9 *WETH9Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETH9.Contract.WETH9Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WETH9 *WETH9Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WETH9.Contract.WETH9Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WETH9 *WETH9CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WETH9.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WETH9 *WETH9TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETH9.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WETH9 *WETH9TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WETH9.Contract.contract.Transact(opts, method, params...)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WETH9 *WETH9Transactor) Withdraw(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _WETH9.contract.Transact(opts, "withdraw", wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WETH9 *WETH9Session) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _WETH9.Contract.Withdraw(&_WETH9.TransactOpts, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WETH9 *WETH9TransactorSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _WETH9.Contract.Withdraw(&_WETH9.TransactOpts, wad)
}
