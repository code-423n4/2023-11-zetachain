// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package connectorerrors

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

// ConnectorErrorsMetaData contains all meta data concerning the ConnectorErrors contract.
var ConnectorErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotPauser\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotTss\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotTssOrUpdater\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotTssUpdater\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"name\":\"ExceedsMaxSupply\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZetaTransferError\",\"type\":\"error\"}]",
}

// ConnectorErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use ConnectorErrorsMetaData.ABI instead.
var ConnectorErrorsABI = ConnectorErrorsMetaData.ABI

// ConnectorErrors is an auto generated Go binding around an Ethereum contract.
type ConnectorErrors struct {
	ConnectorErrorsCaller     // Read-only binding to the contract
	ConnectorErrorsTransactor // Write-only binding to the contract
	ConnectorErrorsFilterer   // Log filterer for contract events
}

// ConnectorErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConnectorErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConnectorErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConnectorErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConnectorErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConnectorErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConnectorErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConnectorErrorsSession struct {
	Contract     *ConnectorErrors  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConnectorErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConnectorErrorsCallerSession struct {
	Contract *ConnectorErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ConnectorErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConnectorErrorsTransactorSession struct {
	Contract     *ConnectorErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ConnectorErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConnectorErrorsRaw struct {
	Contract *ConnectorErrors // Generic contract binding to access the raw methods on
}

// ConnectorErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConnectorErrorsCallerRaw struct {
	Contract *ConnectorErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// ConnectorErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConnectorErrorsTransactorRaw struct {
	Contract *ConnectorErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConnectorErrors creates a new instance of ConnectorErrors, bound to a specific deployed contract.
func NewConnectorErrors(address common.Address, backend bind.ContractBackend) (*ConnectorErrors, error) {
	contract, err := bindConnectorErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConnectorErrors{ConnectorErrorsCaller: ConnectorErrorsCaller{contract: contract}, ConnectorErrorsTransactor: ConnectorErrorsTransactor{contract: contract}, ConnectorErrorsFilterer: ConnectorErrorsFilterer{contract: contract}}, nil
}

// NewConnectorErrorsCaller creates a new read-only instance of ConnectorErrors, bound to a specific deployed contract.
func NewConnectorErrorsCaller(address common.Address, caller bind.ContractCaller) (*ConnectorErrorsCaller, error) {
	contract, err := bindConnectorErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConnectorErrorsCaller{contract: contract}, nil
}

// NewConnectorErrorsTransactor creates a new write-only instance of ConnectorErrors, bound to a specific deployed contract.
func NewConnectorErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*ConnectorErrorsTransactor, error) {
	contract, err := bindConnectorErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConnectorErrorsTransactor{contract: contract}, nil
}

// NewConnectorErrorsFilterer creates a new log filterer instance of ConnectorErrors, bound to a specific deployed contract.
func NewConnectorErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*ConnectorErrorsFilterer, error) {
	contract, err := bindConnectorErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConnectorErrorsFilterer{contract: contract}, nil
}

// bindConnectorErrors binds a generic wrapper to an already deployed contract.
func bindConnectorErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ConnectorErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConnectorErrors *ConnectorErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConnectorErrors.Contract.ConnectorErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConnectorErrors *ConnectorErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConnectorErrors.Contract.ConnectorErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConnectorErrors *ConnectorErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConnectorErrors.Contract.ConnectorErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConnectorErrors *ConnectorErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConnectorErrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConnectorErrors *ConnectorErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConnectorErrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConnectorErrors *ConnectorErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConnectorErrors.Contract.contract.Transact(opts, method, params...)
}
