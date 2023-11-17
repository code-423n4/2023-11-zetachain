// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tridentconcentratedliquiditypoolfactory

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

// ConcentratedLiquidityPoolFactoryMetaData contains all meta data concerning the ConcentratedLiquidityPoolFactory contract.
var ConcentratedLiquidityPoolFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"getPools\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"pairPools\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ConcentratedLiquidityPoolFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ConcentratedLiquidityPoolFactoryMetaData.ABI instead.
var ConcentratedLiquidityPoolFactoryABI = ConcentratedLiquidityPoolFactoryMetaData.ABI

// ConcentratedLiquidityPoolFactory is an auto generated Go binding around an Ethereum contract.
type ConcentratedLiquidityPoolFactory struct {
	ConcentratedLiquidityPoolFactoryCaller     // Read-only binding to the contract
	ConcentratedLiquidityPoolFactoryTransactor // Write-only binding to the contract
	ConcentratedLiquidityPoolFactoryFilterer   // Log filterer for contract events
}

// ConcentratedLiquidityPoolFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConcentratedLiquidityPoolFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConcentratedLiquidityPoolFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConcentratedLiquidityPoolFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConcentratedLiquidityPoolFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConcentratedLiquidityPoolFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConcentratedLiquidityPoolFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConcentratedLiquidityPoolFactorySession struct {
	Contract     *ConcentratedLiquidityPoolFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                     // Call options to use throughout this session
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ConcentratedLiquidityPoolFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConcentratedLiquidityPoolFactoryCallerSession struct {
	Contract *ConcentratedLiquidityPoolFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                           // Call options to use throughout this session
}

// ConcentratedLiquidityPoolFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConcentratedLiquidityPoolFactoryTransactorSession struct {
	Contract     *ConcentratedLiquidityPoolFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                           // Transaction auth options to use throughout this session
}

// ConcentratedLiquidityPoolFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConcentratedLiquidityPoolFactoryRaw struct {
	Contract *ConcentratedLiquidityPoolFactory // Generic contract binding to access the raw methods on
}

// ConcentratedLiquidityPoolFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConcentratedLiquidityPoolFactoryCallerRaw struct {
	Contract *ConcentratedLiquidityPoolFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ConcentratedLiquidityPoolFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConcentratedLiquidityPoolFactoryTransactorRaw struct {
	Contract *ConcentratedLiquidityPoolFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConcentratedLiquidityPoolFactory creates a new instance of ConcentratedLiquidityPoolFactory, bound to a specific deployed contract.
func NewConcentratedLiquidityPoolFactory(address common.Address, backend bind.ContractBackend) (*ConcentratedLiquidityPoolFactory, error) {
	contract, err := bindConcentratedLiquidityPoolFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConcentratedLiquidityPoolFactory{ConcentratedLiquidityPoolFactoryCaller: ConcentratedLiquidityPoolFactoryCaller{contract: contract}, ConcentratedLiquidityPoolFactoryTransactor: ConcentratedLiquidityPoolFactoryTransactor{contract: contract}, ConcentratedLiquidityPoolFactoryFilterer: ConcentratedLiquidityPoolFactoryFilterer{contract: contract}}, nil
}

// NewConcentratedLiquidityPoolFactoryCaller creates a new read-only instance of ConcentratedLiquidityPoolFactory, bound to a specific deployed contract.
func NewConcentratedLiquidityPoolFactoryCaller(address common.Address, caller bind.ContractCaller) (*ConcentratedLiquidityPoolFactoryCaller, error) {
	contract, err := bindConcentratedLiquidityPoolFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConcentratedLiquidityPoolFactoryCaller{contract: contract}, nil
}

// NewConcentratedLiquidityPoolFactoryTransactor creates a new write-only instance of ConcentratedLiquidityPoolFactory, bound to a specific deployed contract.
func NewConcentratedLiquidityPoolFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ConcentratedLiquidityPoolFactoryTransactor, error) {
	contract, err := bindConcentratedLiquidityPoolFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConcentratedLiquidityPoolFactoryTransactor{contract: contract}, nil
}

// NewConcentratedLiquidityPoolFactoryFilterer creates a new log filterer instance of ConcentratedLiquidityPoolFactory, bound to a specific deployed contract.
func NewConcentratedLiquidityPoolFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ConcentratedLiquidityPoolFactoryFilterer, error) {
	contract, err := bindConcentratedLiquidityPoolFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConcentratedLiquidityPoolFactoryFilterer{contract: contract}, nil
}

// bindConcentratedLiquidityPoolFactory binds a generic wrapper to an already deployed contract.
func bindConcentratedLiquidityPoolFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ConcentratedLiquidityPoolFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConcentratedLiquidityPoolFactory *ConcentratedLiquidityPoolFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConcentratedLiquidityPoolFactory.Contract.ConcentratedLiquidityPoolFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConcentratedLiquidityPoolFactory *ConcentratedLiquidityPoolFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConcentratedLiquidityPoolFactory.Contract.ConcentratedLiquidityPoolFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConcentratedLiquidityPoolFactory *ConcentratedLiquidityPoolFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConcentratedLiquidityPoolFactory.Contract.ConcentratedLiquidityPoolFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConcentratedLiquidityPoolFactory *ConcentratedLiquidityPoolFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConcentratedLiquidityPoolFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConcentratedLiquidityPoolFactory *ConcentratedLiquidityPoolFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConcentratedLiquidityPoolFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConcentratedLiquidityPoolFactory *ConcentratedLiquidityPoolFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConcentratedLiquidityPoolFactory.Contract.contract.Transact(opts, method, params...)
}

// GetPools is a free data retrieval call binding the contract method 0x71a25812.
//
// Solidity: function getPools(address token0, address token1, uint256 startIndex, uint256 count) view returns(address[] pairPools)
func (_ConcentratedLiquidityPoolFactory *ConcentratedLiquidityPoolFactoryCaller) GetPools(opts *bind.CallOpts, token0 common.Address, token1 common.Address, startIndex *big.Int, count *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _ConcentratedLiquidityPoolFactory.contract.Call(opts, &out, "getPools", token0, token1, startIndex, count)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPools is a free data retrieval call binding the contract method 0x71a25812.
//
// Solidity: function getPools(address token0, address token1, uint256 startIndex, uint256 count) view returns(address[] pairPools)
func (_ConcentratedLiquidityPoolFactory *ConcentratedLiquidityPoolFactorySession) GetPools(token0 common.Address, token1 common.Address, startIndex *big.Int, count *big.Int) ([]common.Address, error) {
	return _ConcentratedLiquidityPoolFactory.Contract.GetPools(&_ConcentratedLiquidityPoolFactory.CallOpts, token0, token1, startIndex, count)
}

// GetPools is a free data retrieval call binding the contract method 0x71a25812.
//
// Solidity: function getPools(address token0, address token1, uint256 startIndex, uint256 count) view returns(address[] pairPools)
func (_ConcentratedLiquidityPoolFactory *ConcentratedLiquidityPoolFactoryCallerSession) GetPools(token0 common.Address, token1 common.Address, startIndex *big.Int, count *big.Int) ([]common.Address, error) {
	return _ConcentratedLiquidityPoolFactory.Contract.GetPools(&_ConcentratedLiquidityPoolFactory.CallOpts, token0, token1, startIndex, count)
}
