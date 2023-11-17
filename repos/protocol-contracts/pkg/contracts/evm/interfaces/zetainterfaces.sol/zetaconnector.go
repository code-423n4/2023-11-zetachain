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

// ZetaInterfacesSendInput is an auto generated low-level Go binding around an user-defined struct.
type ZetaInterfacesSendInput struct {
	DestinationChainId  *big.Int
	DestinationAddress  []byte
	DestinationGasLimit *big.Int
	Message             []byte
	ZetaValueAndGas     *big.Int
	ZetaParams          []byte
}

// ZetaConnectorMetaData contains all meta data concerning the ZetaConnector contract.
var ZetaConnectorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"destinationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"zetaValueAndGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"zetaParams\",\"type\":\"bytes\"}],\"internalType\":\"structZetaInterfaces.SendInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ZetaConnectorABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaConnectorMetaData.ABI instead.
var ZetaConnectorABI = ZetaConnectorMetaData.ABI

// ZetaConnector is an auto generated Go binding around an Ethereum contract.
type ZetaConnector struct {
	ZetaConnectorCaller     // Read-only binding to the contract
	ZetaConnectorTransactor // Write-only binding to the contract
	ZetaConnectorFilterer   // Log filterer for contract events
}

// ZetaConnectorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaConnectorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaConnectorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaConnectorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaConnectorSession struct {
	Contract     *ZetaConnector    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZetaConnectorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaConnectorCallerSession struct {
	Contract *ZetaConnectorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ZetaConnectorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaConnectorTransactorSession struct {
	Contract     *ZetaConnectorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ZetaConnectorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaConnectorRaw struct {
	Contract *ZetaConnector // Generic contract binding to access the raw methods on
}

// ZetaConnectorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaConnectorCallerRaw struct {
	Contract *ZetaConnectorCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaConnectorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaConnectorTransactorRaw struct {
	Contract *ZetaConnectorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaConnector creates a new instance of ZetaConnector, bound to a specific deployed contract.
func NewZetaConnector(address common.Address, backend bind.ContractBackend) (*ZetaConnector, error) {
	contract, err := bindZetaConnector(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaConnector{ZetaConnectorCaller: ZetaConnectorCaller{contract: contract}, ZetaConnectorTransactor: ZetaConnectorTransactor{contract: contract}, ZetaConnectorFilterer: ZetaConnectorFilterer{contract: contract}}, nil
}

// NewZetaConnectorCaller creates a new read-only instance of ZetaConnector, bound to a specific deployed contract.
func NewZetaConnectorCaller(address common.Address, caller bind.ContractCaller) (*ZetaConnectorCaller, error) {
	contract, err := bindZetaConnector(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorCaller{contract: contract}, nil
}

// NewZetaConnectorTransactor creates a new write-only instance of ZetaConnector, bound to a specific deployed contract.
func NewZetaConnectorTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaConnectorTransactor, error) {
	contract, err := bindZetaConnector(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorTransactor{contract: contract}, nil
}

// NewZetaConnectorFilterer creates a new log filterer instance of ZetaConnector, bound to a specific deployed contract.
func NewZetaConnectorFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaConnectorFilterer, error) {
	contract, err := bindZetaConnector(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorFilterer{contract: contract}, nil
}

// bindZetaConnector binds a generic wrapper to an already deployed contract.
func bindZetaConnector(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaConnectorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnector *ZetaConnectorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnector.Contract.ZetaConnectorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnector *ZetaConnectorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnector.Contract.ZetaConnectorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnector *ZetaConnectorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnector.Contract.ZetaConnectorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnector *ZetaConnectorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnector.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnector *ZetaConnectorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnector.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnector *ZetaConnectorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnector.Contract.contract.Transact(opts, method, params...)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_ZetaConnector *ZetaConnectorTransactor) Send(opts *bind.TransactOpts, input ZetaInterfacesSendInput) (*types.Transaction, error) {
	return _ZetaConnector.contract.Transact(opts, "send", input)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_ZetaConnector *ZetaConnectorSession) Send(input ZetaInterfacesSendInput) (*types.Transaction, error) {
	return _ZetaConnector.Contract.Send(&_ZetaConnector.TransactOpts, input)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_ZetaConnector *ZetaConnectorTransactorSession) Send(input ZetaInterfacesSendInput) (*types.Transaction, error) {
	return _ZetaConnector.Contract.Send(&_ZetaConnector.TransactOpts, input)
}
