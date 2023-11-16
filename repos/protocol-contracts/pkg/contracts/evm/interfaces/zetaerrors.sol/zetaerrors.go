// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetaerrors

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

// ZetaErrorsMetaData contains all meta data concerning the ZetaErrors contract.
var ZetaErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotConnector\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotTss\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotTssOrUpdater\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotTssUpdater\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZetaTransferError\",\"type\":\"error\"}]",
}

// ZetaErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaErrorsMetaData.ABI instead.
var ZetaErrorsABI = ZetaErrorsMetaData.ABI

// ZetaErrors is an auto generated Go binding around an Ethereum contract.
type ZetaErrors struct {
	ZetaErrorsCaller     // Read-only binding to the contract
	ZetaErrorsTransactor // Write-only binding to the contract
	ZetaErrorsFilterer   // Log filterer for contract events
}

// ZetaErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaErrorsSession struct {
	Contract     *ZetaErrors       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZetaErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaErrorsCallerSession struct {
	Contract *ZetaErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ZetaErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaErrorsTransactorSession struct {
	Contract     *ZetaErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ZetaErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaErrorsRaw struct {
	Contract *ZetaErrors // Generic contract binding to access the raw methods on
}

// ZetaErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaErrorsCallerRaw struct {
	Contract *ZetaErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaErrorsTransactorRaw struct {
	Contract *ZetaErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaErrors creates a new instance of ZetaErrors, bound to a specific deployed contract.
func NewZetaErrors(address common.Address, backend bind.ContractBackend) (*ZetaErrors, error) {
	contract, err := bindZetaErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaErrors{ZetaErrorsCaller: ZetaErrorsCaller{contract: contract}, ZetaErrorsTransactor: ZetaErrorsTransactor{contract: contract}, ZetaErrorsFilterer: ZetaErrorsFilterer{contract: contract}}, nil
}

// NewZetaErrorsCaller creates a new read-only instance of ZetaErrors, bound to a specific deployed contract.
func NewZetaErrorsCaller(address common.Address, caller bind.ContractCaller) (*ZetaErrorsCaller, error) {
	contract, err := bindZetaErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaErrorsCaller{contract: contract}, nil
}

// NewZetaErrorsTransactor creates a new write-only instance of ZetaErrors, bound to a specific deployed contract.
func NewZetaErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaErrorsTransactor, error) {
	contract, err := bindZetaErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaErrorsTransactor{contract: contract}, nil
}

// NewZetaErrorsFilterer creates a new log filterer instance of ZetaErrors, bound to a specific deployed contract.
func NewZetaErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaErrorsFilterer, error) {
	contract, err := bindZetaErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaErrorsFilterer{contract: contract}, nil
}

// bindZetaErrors binds a generic wrapper to an already deployed contract.
func bindZetaErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaErrors *ZetaErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaErrors.Contract.ZetaErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaErrors *ZetaErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaErrors.Contract.ZetaErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaErrors *ZetaErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaErrors.Contract.ZetaErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaErrors *ZetaErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaErrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaErrors *ZetaErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaErrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaErrors *ZetaErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaErrors.Contract.contract.Transact(opts, method, params...)
}
