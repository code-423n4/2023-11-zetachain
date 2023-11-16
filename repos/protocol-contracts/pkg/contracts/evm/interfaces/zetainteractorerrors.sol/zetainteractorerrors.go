// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetainteractorerrors

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

// ZetaInteractorErrorsMetaData contains all meta data concerning the ZetaInteractorErrors contract.
var ZetaInteractorErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"InvalidCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDestinationChainId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZetaMessageCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZetaRevertCall\",\"type\":\"error\"}]",
}

// ZetaInteractorErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaInteractorErrorsMetaData.ABI instead.
var ZetaInteractorErrorsABI = ZetaInteractorErrorsMetaData.ABI

// ZetaInteractorErrors is an auto generated Go binding around an Ethereum contract.
type ZetaInteractorErrors struct {
	ZetaInteractorErrorsCaller     // Read-only binding to the contract
	ZetaInteractorErrorsTransactor // Write-only binding to the contract
	ZetaInteractorErrorsFilterer   // Log filterer for contract events
}

// ZetaInteractorErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaInteractorErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaInteractorErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaInteractorErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaInteractorErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaInteractorErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaInteractorErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaInteractorErrorsSession struct {
	Contract     *ZetaInteractorErrors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ZetaInteractorErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaInteractorErrorsCallerSession struct {
	Contract *ZetaInteractorErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ZetaInteractorErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaInteractorErrorsTransactorSession struct {
	Contract     *ZetaInteractorErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ZetaInteractorErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaInteractorErrorsRaw struct {
	Contract *ZetaInteractorErrors // Generic contract binding to access the raw methods on
}

// ZetaInteractorErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaInteractorErrorsCallerRaw struct {
	Contract *ZetaInteractorErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaInteractorErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaInteractorErrorsTransactorRaw struct {
	Contract *ZetaInteractorErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaInteractorErrors creates a new instance of ZetaInteractorErrors, bound to a specific deployed contract.
func NewZetaInteractorErrors(address common.Address, backend bind.ContractBackend) (*ZetaInteractorErrors, error) {
	contract, err := bindZetaInteractorErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaInteractorErrors{ZetaInteractorErrorsCaller: ZetaInteractorErrorsCaller{contract: contract}, ZetaInteractorErrorsTransactor: ZetaInteractorErrorsTransactor{contract: contract}, ZetaInteractorErrorsFilterer: ZetaInteractorErrorsFilterer{contract: contract}}, nil
}

// NewZetaInteractorErrorsCaller creates a new read-only instance of ZetaInteractorErrors, bound to a specific deployed contract.
func NewZetaInteractorErrorsCaller(address common.Address, caller bind.ContractCaller) (*ZetaInteractorErrorsCaller, error) {
	contract, err := bindZetaInteractorErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaInteractorErrorsCaller{contract: contract}, nil
}

// NewZetaInteractorErrorsTransactor creates a new write-only instance of ZetaInteractorErrors, bound to a specific deployed contract.
func NewZetaInteractorErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaInteractorErrorsTransactor, error) {
	contract, err := bindZetaInteractorErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaInteractorErrorsTransactor{contract: contract}, nil
}

// NewZetaInteractorErrorsFilterer creates a new log filterer instance of ZetaInteractorErrors, bound to a specific deployed contract.
func NewZetaInteractorErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaInteractorErrorsFilterer, error) {
	contract, err := bindZetaInteractorErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaInteractorErrorsFilterer{contract: contract}, nil
}

// bindZetaInteractorErrors binds a generic wrapper to an already deployed contract.
func bindZetaInteractorErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaInteractorErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaInteractorErrors *ZetaInteractorErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaInteractorErrors.Contract.ZetaInteractorErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaInteractorErrors *ZetaInteractorErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaInteractorErrors.Contract.ZetaInteractorErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaInteractorErrors *ZetaInteractorErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaInteractorErrors.Contract.ZetaInteractorErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaInteractorErrors *ZetaInteractorErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaInteractorErrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaInteractorErrors *ZetaInteractorErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaInteractorErrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaInteractorErrors *ZetaInteractorErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaInteractorErrors.Contract.contract.Transact(opts, method, params...)
}
