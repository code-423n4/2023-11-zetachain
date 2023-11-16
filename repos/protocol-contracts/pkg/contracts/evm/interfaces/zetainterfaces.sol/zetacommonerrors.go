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

// ZetaCommonErrorsMetaData contains all meta data concerning the ZetaCommonErrors contract.
var ZetaCommonErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"}]",
}

// ZetaCommonErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaCommonErrorsMetaData.ABI instead.
var ZetaCommonErrorsABI = ZetaCommonErrorsMetaData.ABI

// ZetaCommonErrors is an auto generated Go binding around an Ethereum contract.
type ZetaCommonErrors struct {
	ZetaCommonErrorsCaller     // Read-only binding to the contract
	ZetaCommonErrorsTransactor // Write-only binding to the contract
	ZetaCommonErrorsFilterer   // Log filterer for contract events
}

// ZetaCommonErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaCommonErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaCommonErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaCommonErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaCommonErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaCommonErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaCommonErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaCommonErrorsSession struct {
	Contract     *ZetaCommonErrors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZetaCommonErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaCommonErrorsCallerSession struct {
	Contract *ZetaCommonErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ZetaCommonErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaCommonErrorsTransactorSession struct {
	Contract     *ZetaCommonErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ZetaCommonErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaCommonErrorsRaw struct {
	Contract *ZetaCommonErrors // Generic contract binding to access the raw methods on
}

// ZetaCommonErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaCommonErrorsCallerRaw struct {
	Contract *ZetaCommonErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaCommonErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaCommonErrorsTransactorRaw struct {
	Contract *ZetaCommonErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaCommonErrors creates a new instance of ZetaCommonErrors, bound to a specific deployed contract.
func NewZetaCommonErrors(address common.Address, backend bind.ContractBackend) (*ZetaCommonErrors, error) {
	contract, err := bindZetaCommonErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaCommonErrors{ZetaCommonErrorsCaller: ZetaCommonErrorsCaller{contract: contract}, ZetaCommonErrorsTransactor: ZetaCommonErrorsTransactor{contract: contract}, ZetaCommonErrorsFilterer: ZetaCommonErrorsFilterer{contract: contract}}, nil
}

// NewZetaCommonErrorsCaller creates a new read-only instance of ZetaCommonErrors, bound to a specific deployed contract.
func NewZetaCommonErrorsCaller(address common.Address, caller bind.ContractCaller) (*ZetaCommonErrorsCaller, error) {
	contract, err := bindZetaCommonErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaCommonErrorsCaller{contract: contract}, nil
}

// NewZetaCommonErrorsTransactor creates a new write-only instance of ZetaCommonErrors, bound to a specific deployed contract.
func NewZetaCommonErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaCommonErrorsTransactor, error) {
	contract, err := bindZetaCommonErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaCommonErrorsTransactor{contract: contract}, nil
}

// NewZetaCommonErrorsFilterer creates a new log filterer instance of ZetaCommonErrors, bound to a specific deployed contract.
func NewZetaCommonErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaCommonErrorsFilterer, error) {
	contract, err := bindZetaCommonErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaCommonErrorsFilterer{contract: contract}, nil
}

// bindZetaCommonErrors binds a generic wrapper to an already deployed contract.
func bindZetaCommonErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaCommonErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaCommonErrors *ZetaCommonErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaCommonErrors.Contract.ZetaCommonErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaCommonErrors *ZetaCommonErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaCommonErrors.Contract.ZetaCommonErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaCommonErrors *ZetaCommonErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaCommonErrors.Contract.ZetaCommonErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaCommonErrors *ZetaCommonErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaCommonErrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaCommonErrors *ZetaCommonErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaCommonErrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaCommonErrors *ZetaCommonErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaCommonErrors.Contract.contract.Transact(opts, method, params...)
}
