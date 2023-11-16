// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetatokenconsumeruniv2

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

// ZetaTokenConsumerUniV2ErrorsMetaData contains all meta data concerning the ZetaTokenConsumerUniV2Errors contract.
var ZetaTokenConsumerUniV2ErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InputCantBeZero\",\"type\":\"error\"}]",
}

// ZetaTokenConsumerUniV2ErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaTokenConsumerUniV2ErrorsMetaData.ABI instead.
var ZetaTokenConsumerUniV2ErrorsABI = ZetaTokenConsumerUniV2ErrorsMetaData.ABI

// ZetaTokenConsumerUniV2Errors is an auto generated Go binding around an Ethereum contract.
type ZetaTokenConsumerUniV2Errors struct {
	ZetaTokenConsumerUniV2ErrorsCaller     // Read-only binding to the contract
	ZetaTokenConsumerUniV2ErrorsTransactor // Write-only binding to the contract
	ZetaTokenConsumerUniV2ErrorsFilterer   // Log filterer for contract events
}

// ZetaTokenConsumerUniV2ErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaTokenConsumerUniV2ErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerUniV2ErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaTokenConsumerUniV2ErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerUniV2ErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaTokenConsumerUniV2ErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerUniV2ErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaTokenConsumerUniV2ErrorsSession struct {
	Contract     *ZetaTokenConsumerUniV2Errors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ZetaTokenConsumerUniV2ErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaTokenConsumerUniV2ErrorsCallerSession struct {
	Contract *ZetaTokenConsumerUniV2ErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// ZetaTokenConsumerUniV2ErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaTokenConsumerUniV2ErrorsTransactorSession struct {
	Contract     *ZetaTokenConsumerUniV2ErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// ZetaTokenConsumerUniV2ErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaTokenConsumerUniV2ErrorsRaw struct {
	Contract *ZetaTokenConsumerUniV2Errors // Generic contract binding to access the raw methods on
}

// ZetaTokenConsumerUniV2ErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaTokenConsumerUniV2ErrorsCallerRaw struct {
	Contract *ZetaTokenConsumerUniV2ErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaTokenConsumerUniV2ErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaTokenConsumerUniV2ErrorsTransactorRaw struct {
	Contract *ZetaTokenConsumerUniV2ErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaTokenConsumerUniV2Errors creates a new instance of ZetaTokenConsumerUniV2Errors, bound to a specific deployed contract.
func NewZetaTokenConsumerUniV2Errors(address common.Address, backend bind.ContractBackend) (*ZetaTokenConsumerUniV2Errors, error) {
	contract, err := bindZetaTokenConsumerUniV2Errors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerUniV2Errors{ZetaTokenConsumerUniV2ErrorsCaller: ZetaTokenConsumerUniV2ErrorsCaller{contract: contract}, ZetaTokenConsumerUniV2ErrorsTransactor: ZetaTokenConsumerUniV2ErrorsTransactor{contract: contract}, ZetaTokenConsumerUniV2ErrorsFilterer: ZetaTokenConsumerUniV2ErrorsFilterer{contract: contract}}, nil
}

// NewZetaTokenConsumerUniV2ErrorsCaller creates a new read-only instance of ZetaTokenConsumerUniV2Errors, bound to a specific deployed contract.
func NewZetaTokenConsumerUniV2ErrorsCaller(address common.Address, caller bind.ContractCaller) (*ZetaTokenConsumerUniV2ErrorsCaller, error) {
	contract, err := bindZetaTokenConsumerUniV2Errors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerUniV2ErrorsCaller{contract: contract}, nil
}

// NewZetaTokenConsumerUniV2ErrorsTransactor creates a new write-only instance of ZetaTokenConsumerUniV2Errors, bound to a specific deployed contract.
func NewZetaTokenConsumerUniV2ErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaTokenConsumerUniV2ErrorsTransactor, error) {
	contract, err := bindZetaTokenConsumerUniV2Errors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerUniV2ErrorsTransactor{contract: contract}, nil
}

// NewZetaTokenConsumerUniV2ErrorsFilterer creates a new log filterer instance of ZetaTokenConsumerUniV2Errors, bound to a specific deployed contract.
func NewZetaTokenConsumerUniV2ErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaTokenConsumerUniV2ErrorsFilterer, error) {
	contract, err := bindZetaTokenConsumerUniV2Errors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerUniV2ErrorsFilterer{contract: contract}, nil
}

// bindZetaTokenConsumerUniV2Errors binds a generic wrapper to an already deployed contract.
func bindZetaTokenConsumerUniV2Errors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaTokenConsumerUniV2ErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaTokenConsumerUniV2Errors *ZetaTokenConsumerUniV2ErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaTokenConsumerUniV2Errors.Contract.ZetaTokenConsumerUniV2ErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaTokenConsumerUniV2Errors *ZetaTokenConsumerUniV2ErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV2Errors.Contract.ZetaTokenConsumerUniV2ErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaTokenConsumerUniV2Errors *ZetaTokenConsumerUniV2ErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV2Errors.Contract.ZetaTokenConsumerUniV2ErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaTokenConsumerUniV2Errors *ZetaTokenConsumerUniV2ErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaTokenConsumerUniV2Errors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaTokenConsumerUniV2Errors *ZetaTokenConsumerUniV2ErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV2Errors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaTokenConsumerUniV2Errors *ZetaTokenConsumerUniV2ErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV2Errors.Contract.contract.Transact(opts, method, params...)
}
