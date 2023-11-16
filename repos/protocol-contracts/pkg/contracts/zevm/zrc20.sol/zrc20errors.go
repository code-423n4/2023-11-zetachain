// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zrc20

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

// ZRC20ErrorsMetaData contains all meta data concerning the ZRC20Errors contract.
var ZRC20ErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CallerIsNotFungibleModule\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasFeeTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LowAllowance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LowBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroGasCoin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroGasPrice\",\"type\":\"error\"}]",
}

// ZRC20ErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use ZRC20ErrorsMetaData.ABI instead.
var ZRC20ErrorsABI = ZRC20ErrorsMetaData.ABI

// ZRC20Errors is an auto generated Go binding around an Ethereum contract.
type ZRC20Errors struct {
	ZRC20ErrorsCaller     // Read-only binding to the contract
	ZRC20ErrorsTransactor // Write-only binding to the contract
	ZRC20ErrorsFilterer   // Log filterer for contract events
}

// ZRC20ErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZRC20ErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZRC20ErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZRC20ErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZRC20ErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZRC20ErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZRC20ErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZRC20ErrorsSession struct {
	Contract     *ZRC20Errors      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZRC20ErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZRC20ErrorsCallerSession struct {
	Contract *ZRC20ErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ZRC20ErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZRC20ErrorsTransactorSession struct {
	Contract     *ZRC20ErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ZRC20ErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZRC20ErrorsRaw struct {
	Contract *ZRC20Errors // Generic contract binding to access the raw methods on
}

// ZRC20ErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZRC20ErrorsCallerRaw struct {
	Contract *ZRC20ErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// ZRC20ErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZRC20ErrorsTransactorRaw struct {
	Contract *ZRC20ErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZRC20Errors creates a new instance of ZRC20Errors, bound to a specific deployed contract.
func NewZRC20Errors(address common.Address, backend bind.ContractBackend) (*ZRC20Errors, error) {
	contract, err := bindZRC20Errors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZRC20Errors{ZRC20ErrorsCaller: ZRC20ErrorsCaller{contract: contract}, ZRC20ErrorsTransactor: ZRC20ErrorsTransactor{contract: contract}, ZRC20ErrorsFilterer: ZRC20ErrorsFilterer{contract: contract}}, nil
}

// NewZRC20ErrorsCaller creates a new read-only instance of ZRC20Errors, bound to a specific deployed contract.
func NewZRC20ErrorsCaller(address common.Address, caller bind.ContractCaller) (*ZRC20ErrorsCaller, error) {
	contract, err := bindZRC20Errors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZRC20ErrorsCaller{contract: contract}, nil
}

// NewZRC20ErrorsTransactor creates a new write-only instance of ZRC20Errors, bound to a specific deployed contract.
func NewZRC20ErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*ZRC20ErrorsTransactor, error) {
	contract, err := bindZRC20Errors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZRC20ErrorsTransactor{contract: contract}, nil
}

// NewZRC20ErrorsFilterer creates a new log filterer instance of ZRC20Errors, bound to a specific deployed contract.
func NewZRC20ErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*ZRC20ErrorsFilterer, error) {
	contract, err := bindZRC20Errors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZRC20ErrorsFilterer{contract: contract}, nil
}

// bindZRC20Errors binds a generic wrapper to an already deployed contract.
func bindZRC20Errors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZRC20ErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZRC20Errors *ZRC20ErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZRC20Errors.Contract.ZRC20ErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZRC20Errors *ZRC20ErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZRC20Errors.Contract.ZRC20ErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZRC20Errors *ZRC20ErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZRC20Errors.Contract.ZRC20ErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZRC20Errors *ZRC20ErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZRC20Errors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZRC20Errors *ZRC20ErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZRC20Errors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZRC20Errors *ZRC20ErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZRC20Errors.Contract.contract.Transact(opts, method, params...)
}
