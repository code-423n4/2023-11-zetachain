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

// ZetaInterfacesZetaMessage is an auto generated low-level Go binding around an user-defined struct.
type ZetaInterfacesZetaMessage struct {
	ZetaTxSenderAddress []byte
	SourceChainId       *big.Int
	DestinationAddress  common.Address
	ZetaValue           *big.Int
	Message             []byte
}

// ZetaInterfacesZetaRevert is an auto generated low-level Go binding around an user-defined struct.
type ZetaInterfacesZetaRevert struct {
	ZetaTxSenderAddress common.Address
	SourceChainId       *big.Int
	DestinationAddress  []byte
	DestinationChainId  *big.Int
	RemainingZetaValue  *big.Int
	Message             []byte
}

// ZetaReceiverMetaData contains all meta data concerning the ZetaReceiver contract.
var ZetaReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"zetaTxSenderAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"zetaValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structZetaInterfaces.ZetaMessage\",\"name\":\"zetaMessage\",\"type\":\"tuple\"}],\"name\":\"onZetaMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"zetaTxSenderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingZetaValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structZetaInterfaces.ZetaRevert\",\"name\":\"zetaRevert\",\"type\":\"tuple\"}],\"name\":\"onZetaRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ZetaReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaReceiverMetaData.ABI instead.
var ZetaReceiverABI = ZetaReceiverMetaData.ABI

// ZetaReceiver is an auto generated Go binding around an Ethereum contract.
type ZetaReceiver struct {
	ZetaReceiverCaller     // Read-only binding to the contract
	ZetaReceiverTransactor // Write-only binding to the contract
	ZetaReceiverFilterer   // Log filterer for contract events
}

// ZetaReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaReceiverSession struct {
	Contract     *ZetaReceiver     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZetaReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaReceiverCallerSession struct {
	Contract *ZetaReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ZetaReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaReceiverTransactorSession struct {
	Contract     *ZetaReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ZetaReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaReceiverRaw struct {
	Contract *ZetaReceiver // Generic contract binding to access the raw methods on
}

// ZetaReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaReceiverCallerRaw struct {
	Contract *ZetaReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaReceiverTransactorRaw struct {
	Contract *ZetaReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaReceiver creates a new instance of ZetaReceiver, bound to a specific deployed contract.
func NewZetaReceiver(address common.Address, backend bind.ContractBackend) (*ZetaReceiver, error) {
	contract, err := bindZetaReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaReceiver{ZetaReceiverCaller: ZetaReceiverCaller{contract: contract}, ZetaReceiverTransactor: ZetaReceiverTransactor{contract: contract}, ZetaReceiverFilterer: ZetaReceiverFilterer{contract: contract}}, nil
}

// NewZetaReceiverCaller creates a new read-only instance of ZetaReceiver, bound to a specific deployed contract.
func NewZetaReceiverCaller(address common.Address, caller bind.ContractCaller) (*ZetaReceiverCaller, error) {
	contract, err := bindZetaReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaReceiverCaller{contract: contract}, nil
}

// NewZetaReceiverTransactor creates a new write-only instance of ZetaReceiver, bound to a specific deployed contract.
func NewZetaReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaReceiverTransactor, error) {
	contract, err := bindZetaReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaReceiverTransactor{contract: contract}, nil
}

// NewZetaReceiverFilterer creates a new log filterer instance of ZetaReceiver, bound to a specific deployed contract.
func NewZetaReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaReceiverFilterer, error) {
	contract, err := bindZetaReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaReceiverFilterer{contract: contract}, nil
}

// bindZetaReceiver binds a generic wrapper to an already deployed contract.
func bindZetaReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaReceiver *ZetaReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaReceiver.Contract.ZetaReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaReceiver *ZetaReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaReceiver.Contract.ZetaReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaReceiver *ZetaReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaReceiver.Contract.ZetaReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaReceiver *ZetaReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaReceiver *ZetaReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaReceiver *ZetaReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaReceiver.Contract.contract.Transact(opts, method, params...)
}

// OnZetaMessage is a paid mutator transaction binding the contract method 0x3749c51a.
//
// Solidity: function onZetaMessage((bytes,uint256,address,uint256,bytes) zetaMessage) returns()
func (_ZetaReceiver *ZetaReceiverTransactor) OnZetaMessage(opts *bind.TransactOpts, zetaMessage ZetaInterfacesZetaMessage) (*types.Transaction, error) {
	return _ZetaReceiver.contract.Transact(opts, "onZetaMessage", zetaMessage)
}

// OnZetaMessage is a paid mutator transaction binding the contract method 0x3749c51a.
//
// Solidity: function onZetaMessage((bytes,uint256,address,uint256,bytes) zetaMessage) returns()
func (_ZetaReceiver *ZetaReceiverSession) OnZetaMessage(zetaMessage ZetaInterfacesZetaMessage) (*types.Transaction, error) {
	return _ZetaReceiver.Contract.OnZetaMessage(&_ZetaReceiver.TransactOpts, zetaMessage)
}

// OnZetaMessage is a paid mutator transaction binding the contract method 0x3749c51a.
//
// Solidity: function onZetaMessage((bytes,uint256,address,uint256,bytes) zetaMessage) returns()
func (_ZetaReceiver *ZetaReceiverTransactorSession) OnZetaMessage(zetaMessage ZetaInterfacesZetaMessage) (*types.Transaction, error) {
	return _ZetaReceiver.Contract.OnZetaMessage(&_ZetaReceiver.TransactOpts, zetaMessage)
}

// OnZetaRevert is a paid mutator transaction binding the contract method 0x3ff0693c.
//
// Solidity: function onZetaRevert((address,uint256,bytes,uint256,uint256,bytes) zetaRevert) returns()
func (_ZetaReceiver *ZetaReceiverTransactor) OnZetaRevert(opts *bind.TransactOpts, zetaRevert ZetaInterfacesZetaRevert) (*types.Transaction, error) {
	return _ZetaReceiver.contract.Transact(opts, "onZetaRevert", zetaRevert)
}

// OnZetaRevert is a paid mutator transaction binding the contract method 0x3ff0693c.
//
// Solidity: function onZetaRevert((address,uint256,bytes,uint256,uint256,bytes) zetaRevert) returns()
func (_ZetaReceiver *ZetaReceiverSession) OnZetaRevert(zetaRevert ZetaInterfacesZetaRevert) (*types.Transaction, error) {
	return _ZetaReceiver.Contract.OnZetaRevert(&_ZetaReceiver.TransactOpts, zetaRevert)
}

// OnZetaRevert is a paid mutator transaction binding the contract method 0x3ff0693c.
//
// Solidity: function onZetaRevert((address,uint256,bytes,uint256,uint256,bytes) zetaRevert) returns()
func (_ZetaReceiver *ZetaReceiverTransactorSession) OnZetaRevert(zetaRevert ZetaInterfacesZetaRevert) (*types.Transaction, error) {
	return _ZetaReceiver.Contract.OnZetaRevert(&_ZetaReceiver.TransactOpts, zetaRevert)
}
