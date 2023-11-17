// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testuniswapv3contracts

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

// IPoolInitializerMetaData contains all meta data concerning the IPoolInitializer contract.
var IPoolInitializerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"}],\"name\":\"createAndInitializePoolIfNecessary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IPoolInitializerABI is the input ABI used to generate the binding from.
// Deprecated: Use IPoolInitializerMetaData.ABI instead.
var IPoolInitializerABI = IPoolInitializerMetaData.ABI

// IPoolInitializer is an auto generated Go binding around an Ethereum contract.
type IPoolInitializer struct {
	IPoolInitializerCaller     // Read-only binding to the contract
	IPoolInitializerTransactor // Write-only binding to the contract
	IPoolInitializerFilterer   // Log filterer for contract events
}

// IPoolInitializerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPoolInitializerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolInitializerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPoolInitializerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolInitializerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPoolInitializerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolInitializerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPoolInitializerSession struct {
	Contract     *IPoolInitializer // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPoolInitializerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPoolInitializerCallerSession struct {
	Contract *IPoolInitializerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IPoolInitializerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPoolInitializerTransactorSession struct {
	Contract     *IPoolInitializerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IPoolInitializerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPoolInitializerRaw struct {
	Contract *IPoolInitializer // Generic contract binding to access the raw methods on
}

// IPoolInitializerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPoolInitializerCallerRaw struct {
	Contract *IPoolInitializerCaller // Generic read-only contract binding to access the raw methods on
}

// IPoolInitializerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPoolInitializerTransactorRaw struct {
	Contract *IPoolInitializerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPoolInitializer creates a new instance of IPoolInitializer, bound to a specific deployed contract.
func NewIPoolInitializer(address common.Address, backend bind.ContractBackend) (*IPoolInitializer, error) {
	contract, err := bindIPoolInitializer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPoolInitializer{IPoolInitializerCaller: IPoolInitializerCaller{contract: contract}, IPoolInitializerTransactor: IPoolInitializerTransactor{contract: contract}, IPoolInitializerFilterer: IPoolInitializerFilterer{contract: contract}}, nil
}

// NewIPoolInitializerCaller creates a new read-only instance of IPoolInitializer, bound to a specific deployed contract.
func NewIPoolInitializerCaller(address common.Address, caller bind.ContractCaller) (*IPoolInitializerCaller, error) {
	contract, err := bindIPoolInitializer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPoolInitializerCaller{contract: contract}, nil
}

// NewIPoolInitializerTransactor creates a new write-only instance of IPoolInitializer, bound to a specific deployed contract.
func NewIPoolInitializerTransactor(address common.Address, transactor bind.ContractTransactor) (*IPoolInitializerTransactor, error) {
	contract, err := bindIPoolInitializer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPoolInitializerTransactor{contract: contract}, nil
}

// NewIPoolInitializerFilterer creates a new log filterer instance of IPoolInitializer, bound to a specific deployed contract.
func NewIPoolInitializerFilterer(address common.Address, filterer bind.ContractFilterer) (*IPoolInitializerFilterer, error) {
	contract, err := bindIPoolInitializer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPoolInitializerFilterer{contract: contract}, nil
}

// bindIPoolInitializer binds a generic wrapper to an already deployed contract.
func bindIPoolInitializer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPoolInitializerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPoolInitializer *IPoolInitializerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPoolInitializer.Contract.IPoolInitializerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPoolInitializer *IPoolInitializerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPoolInitializer.Contract.IPoolInitializerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPoolInitializer *IPoolInitializerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPoolInitializer.Contract.IPoolInitializerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPoolInitializer *IPoolInitializerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPoolInitializer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPoolInitializer *IPoolInitializerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPoolInitializer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPoolInitializer *IPoolInitializerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPoolInitializer.Contract.contract.Transact(opts, method, params...)
}

// CreateAndInitializePoolIfNecessary is a paid mutator transaction binding the contract method 0x13ead562.
//
// Solidity: function createAndInitializePoolIfNecessary(address token0, address token1, uint24 fee, uint160 sqrtPriceX96) payable returns(address pool)
func (_IPoolInitializer *IPoolInitializerTransactor) CreateAndInitializePoolIfNecessary(opts *bind.TransactOpts, token0 common.Address, token1 common.Address, fee *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _IPoolInitializer.contract.Transact(opts, "createAndInitializePoolIfNecessary", token0, token1, fee, sqrtPriceX96)
}

// CreateAndInitializePoolIfNecessary is a paid mutator transaction binding the contract method 0x13ead562.
//
// Solidity: function createAndInitializePoolIfNecessary(address token0, address token1, uint24 fee, uint160 sqrtPriceX96) payable returns(address pool)
func (_IPoolInitializer *IPoolInitializerSession) CreateAndInitializePoolIfNecessary(token0 common.Address, token1 common.Address, fee *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _IPoolInitializer.Contract.CreateAndInitializePoolIfNecessary(&_IPoolInitializer.TransactOpts, token0, token1, fee, sqrtPriceX96)
}

// CreateAndInitializePoolIfNecessary is a paid mutator transaction binding the contract method 0x13ead562.
//
// Solidity: function createAndInitializePoolIfNecessary(address token0, address token1, uint24 fee, uint160 sqrtPriceX96) payable returns(address pool)
func (_IPoolInitializer *IPoolInitializerTransactorSession) CreateAndInitializePoolIfNecessary(token0 common.Address, token1 common.Address, fee *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _IPoolInitializer.Contract.CreateAndInitializePoolIfNecessary(&_IPoolInitializer.TransactOpts, token0, token1, fee, sqrtPriceX96)
}
