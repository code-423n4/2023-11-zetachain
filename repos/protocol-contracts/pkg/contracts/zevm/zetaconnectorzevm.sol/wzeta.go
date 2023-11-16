// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetaconnectorzevm

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

// WZETAMetaData contains all meta data concerning the WZETA contract.
var WZETAMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WZETAABI is the input ABI used to generate the binding from.
// Deprecated: Use WZETAMetaData.ABI instead.
var WZETAABI = WZETAMetaData.ABI

// WZETA is an auto generated Go binding around an Ethereum contract.
type WZETA struct {
	WZETACaller     // Read-only binding to the contract
	WZETATransactor // Write-only binding to the contract
	WZETAFilterer   // Log filterer for contract events
}

// WZETACaller is an auto generated read-only Go binding around an Ethereum contract.
type WZETACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WZETATransactor is an auto generated write-only Go binding around an Ethereum contract.
type WZETATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WZETAFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WZETAFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WZETASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WZETASession struct {
	Contract     *WZETA            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WZETACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WZETACallerSession struct {
	Contract *WZETACaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WZETATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WZETATransactorSession struct {
	Contract     *WZETATransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WZETARaw is an auto generated low-level Go binding around an Ethereum contract.
type WZETARaw struct {
	Contract *WZETA // Generic contract binding to access the raw methods on
}

// WZETACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WZETACallerRaw struct {
	Contract *WZETACaller // Generic read-only contract binding to access the raw methods on
}

// WZETATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WZETATransactorRaw struct {
	Contract *WZETATransactor // Generic write-only contract binding to access the raw methods on
}

// NewWZETA creates a new instance of WZETA, bound to a specific deployed contract.
func NewWZETA(address common.Address, backend bind.ContractBackend) (*WZETA, error) {
	contract, err := bindWZETA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WZETA{WZETACaller: WZETACaller{contract: contract}, WZETATransactor: WZETATransactor{contract: contract}, WZETAFilterer: WZETAFilterer{contract: contract}}, nil
}

// NewWZETACaller creates a new read-only instance of WZETA, bound to a specific deployed contract.
func NewWZETACaller(address common.Address, caller bind.ContractCaller) (*WZETACaller, error) {
	contract, err := bindWZETA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WZETACaller{contract: contract}, nil
}

// NewWZETATransactor creates a new write-only instance of WZETA, bound to a specific deployed contract.
func NewWZETATransactor(address common.Address, transactor bind.ContractTransactor) (*WZETATransactor, error) {
	contract, err := bindWZETA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WZETATransactor{contract: contract}, nil
}

// NewWZETAFilterer creates a new log filterer instance of WZETA, bound to a specific deployed contract.
func NewWZETAFilterer(address common.Address, filterer bind.ContractFilterer) (*WZETAFilterer, error) {
	contract, err := bindWZETA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WZETAFilterer{contract: contract}, nil
}

// bindWZETA binds a generic wrapper to an already deployed contract.
func bindWZETA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WZETAMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WZETA *WZETARaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WZETA.Contract.WZETACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WZETA *WZETARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WZETA.Contract.WZETATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WZETA *WZETARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WZETA.Contract.WZETATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WZETA *WZETACallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WZETA.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WZETA *WZETATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WZETA.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WZETA *WZETATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WZETA.Contract.contract.Transact(opts, method, params...)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_WZETA *WZETATransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WZETA.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_WZETA *WZETASession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WZETA.Contract.TransferFrom(&_WZETA.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_WZETA *WZETATransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WZETA.Contract.TransferFrom(&_WZETA.TransactOpts, src, dst, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WZETA *WZETATransactor) Withdraw(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _WZETA.contract.Transact(opts, "withdraw", wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WZETA *WZETASession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _WZETA.Contract.Withdraw(&_WZETA.TransactOpts, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WZETA *WZETATransactorSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _WZETA.Contract.Withdraw(&_WZETA.TransactOpts, wad)
}
