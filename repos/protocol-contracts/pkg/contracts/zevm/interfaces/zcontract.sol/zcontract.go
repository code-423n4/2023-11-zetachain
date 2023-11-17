// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zcontract

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

// ZContext is an auto generated low-level Go binding around an user-defined struct.
type ZContext struct {
	Origin  []byte
	Sender  common.Address
	ChainID *big.Int
}

// ZContractMetaData contains all meta data concerning the ZContract contract.
var ZContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"origin\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"internalType\":\"structzContext\",\"name\":\"context\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"zrc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"onCrossChainCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ZContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ZContractMetaData.ABI instead.
var ZContractABI = ZContractMetaData.ABI

// ZContract is an auto generated Go binding around an Ethereum contract.
type ZContract struct {
	ZContractCaller     // Read-only binding to the contract
	ZContractTransactor // Write-only binding to the contract
	ZContractFilterer   // Log filterer for contract events
}

// ZContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZContractSession struct {
	Contract     *ZContract        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZContractCallerSession struct {
	Contract *ZContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ZContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZContractTransactorSession struct {
	Contract     *ZContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ZContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZContractRaw struct {
	Contract *ZContract // Generic contract binding to access the raw methods on
}

// ZContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZContractCallerRaw struct {
	Contract *ZContractCaller // Generic read-only contract binding to access the raw methods on
}

// ZContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZContractTransactorRaw struct {
	Contract *ZContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZContract creates a new instance of ZContract, bound to a specific deployed contract.
func NewZContract(address common.Address, backend bind.ContractBackend) (*ZContract, error) {
	contract, err := bindZContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZContract{ZContractCaller: ZContractCaller{contract: contract}, ZContractTransactor: ZContractTransactor{contract: contract}, ZContractFilterer: ZContractFilterer{contract: contract}}, nil
}

// NewZContractCaller creates a new read-only instance of ZContract, bound to a specific deployed contract.
func NewZContractCaller(address common.Address, caller bind.ContractCaller) (*ZContractCaller, error) {
	contract, err := bindZContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZContractCaller{contract: contract}, nil
}

// NewZContractTransactor creates a new write-only instance of ZContract, bound to a specific deployed contract.
func NewZContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ZContractTransactor, error) {
	contract, err := bindZContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZContractTransactor{contract: contract}, nil
}

// NewZContractFilterer creates a new log filterer instance of ZContract, bound to a specific deployed contract.
func NewZContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ZContractFilterer, error) {
	contract, err := bindZContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZContractFilterer{contract: contract}, nil
}

// bindZContract binds a generic wrapper to an already deployed contract.
func bindZContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZContract *ZContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZContract.Contract.ZContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZContract *ZContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZContract.Contract.ZContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZContract *ZContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZContract.Contract.ZContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZContract *ZContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZContract *ZContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZContract *ZContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZContract.Contract.contract.Transact(opts, method, params...)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0xde43156e.
//
// Solidity: function onCrossChainCall((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_ZContract *ZContractTransactor) OnCrossChainCall(opts *bind.TransactOpts, context ZContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _ZContract.contract.Transact(opts, "onCrossChainCall", context, zrc20, amount, message)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0xde43156e.
//
// Solidity: function onCrossChainCall((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_ZContract *ZContractSession) OnCrossChainCall(context ZContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _ZContract.Contract.OnCrossChainCall(&_ZContract.TransactOpts, context, zrc20, amount, message)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0xde43156e.
//
// Solidity: function onCrossChainCall((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_ZContract *ZContractTransactorSession) OnCrossChainCall(context ZContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _ZContract.Contract.OnCrossChainCall(&_ZContract.TransactOpts, context, zrc20, amount, message)
}
