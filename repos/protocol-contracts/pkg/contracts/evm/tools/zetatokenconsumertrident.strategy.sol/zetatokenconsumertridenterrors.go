// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetatokenconsumertrident

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

// ZetaTokenConsumerTridentErrorsMetaData contains all meta data concerning the ZetaTokenConsumerTridentErrors contract.
var ZetaTokenConsumerTridentErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ErrorSendingETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputCantBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyError\",\"type\":\"error\"}]",
}

// ZetaTokenConsumerTridentErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaTokenConsumerTridentErrorsMetaData.ABI instead.
var ZetaTokenConsumerTridentErrorsABI = ZetaTokenConsumerTridentErrorsMetaData.ABI

// ZetaTokenConsumerTridentErrors is an auto generated Go binding around an Ethereum contract.
type ZetaTokenConsumerTridentErrors struct {
	ZetaTokenConsumerTridentErrorsCaller     // Read-only binding to the contract
	ZetaTokenConsumerTridentErrorsTransactor // Write-only binding to the contract
	ZetaTokenConsumerTridentErrorsFilterer   // Log filterer for contract events
}

// ZetaTokenConsumerTridentErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaTokenConsumerTridentErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerTridentErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaTokenConsumerTridentErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerTridentErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaTokenConsumerTridentErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerTridentErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaTokenConsumerTridentErrorsSession struct {
	Contract     *ZetaTokenConsumerTridentErrors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                   // Call options to use throughout this session
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ZetaTokenConsumerTridentErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaTokenConsumerTridentErrorsCallerSession struct {
	Contract *ZetaTokenConsumerTridentErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                         // Call options to use throughout this session
}

// ZetaTokenConsumerTridentErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaTokenConsumerTridentErrorsTransactorSession struct {
	Contract     *ZetaTokenConsumerTridentErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// ZetaTokenConsumerTridentErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaTokenConsumerTridentErrorsRaw struct {
	Contract *ZetaTokenConsumerTridentErrors // Generic contract binding to access the raw methods on
}

// ZetaTokenConsumerTridentErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaTokenConsumerTridentErrorsCallerRaw struct {
	Contract *ZetaTokenConsumerTridentErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaTokenConsumerTridentErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaTokenConsumerTridentErrorsTransactorRaw struct {
	Contract *ZetaTokenConsumerTridentErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaTokenConsumerTridentErrors creates a new instance of ZetaTokenConsumerTridentErrors, bound to a specific deployed contract.
func NewZetaTokenConsumerTridentErrors(address common.Address, backend bind.ContractBackend) (*ZetaTokenConsumerTridentErrors, error) {
	contract, err := bindZetaTokenConsumerTridentErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerTridentErrors{ZetaTokenConsumerTridentErrorsCaller: ZetaTokenConsumerTridentErrorsCaller{contract: contract}, ZetaTokenConsumerTridentErrorsTransactor: ZetaTokenConsumerTridentErrorsTransactor{contract: contract}, ZetaTokenConsumerTridentErrorsFilterer: ZetaTokenConsumerTridentErrorsFilterer{contract: contract}}, nil
}

// NewZetaTokenConsumerTridentErrorsCaller creates a new read-only instance of ZetaTokenConsumerTridentErrors, bound to a specific deployed contract.
func NewZetaTokenConsumerTridentErrorsCaller(address common.Address, caller bind.ContractCaller) (*ZetaTokenConsumerTridentErrorsCaller, error) {
	contract, err := bindZetaTokenConsumerTridentErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerTridentErrorsCaller{contract: contract}, nil
}

// NewZetaTokenConsumerTridentErrorsTransactor creates a new write-only instance of ZetaTokenConsumerTridentErrors, bound to a specific deployed contract.
func NewZetaTokenConsumerTridentErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaTokenConsumerTridentErrorsTransactor, error) {
	contract, err := bindZetaTokenConsumerTridentErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerTridentErrorsTransactor{contract: contract}, nil
}

// NewZetaTokenConsumerTridentErrorsFilterer creates a new log filterer instance of ZetaTokenConsumerTridentErrors, bound to a specific deployed contract.
func NewZetaTokenConsumerTridentErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaTokenConsumerTridentErrorsFilterer, error) {
	contract, err := bindZetaTokenConsumerTridentErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerTridentErrorsFilterer{contract: contract}, nil
}

// bindZetaTokenConsumerTridentErrors binds a generic wrapper to an already deployed contract.
func bindZetaTokenConsumerTridentErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaTokenConsumerTridentErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaTokenConsumerTridentErrors *ZetaTokenConsumerTridentErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaTokenConsumerTridentErrors.Contract.ZetaTokenConsumerTridentErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaTokenConsumerTridentErrors *ZetaTokenConsumerTridentErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaTokenConsumerTridentErrors.Contract.ZetaTokenConsumerTridentErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaTokenConsumerTridentErrors *ZetaTokenConsumerTridentErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaTokenConsumerTridentErrors.Contract.ZetaTokenConsumerTridentErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaTokenConsumerTridentErrors *ZetaTokenConsumerTridentErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaTokenConsumerTridentErrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaTokenConsumerTridentErrors *ZetaTokenConsumerTridentErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaTokenConsumerTridentErrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaTokenConsumerTridentErrors *ZetaTokenConsumerTridentErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaTokenConsumerTridentErrors.Contract.contract.Transact(opts, method, params...)
}
