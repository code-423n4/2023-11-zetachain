// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetatokenconsumerpancakev3

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

// ISwapRouterPancakeExactInputParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterPancakeExactInputParams struct {
	Path             []byte
	Recipient        common.Address
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
}

// ISwapRouterPancakeExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterPancakeExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	AmountIn          *big.Int
	AmountOutMinimum  *big.Int
	SqrtPriceLimitX96 *big.Int
}

// ISwapRouterPancakeMetaData contains all meta data concerning the ISwapRouterPancake contract.
var ISwapRouterPancakeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouterPancake.ExactInputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouterPancake.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ISwapRouterPancakeABI is the input ABI used to generate the binding from.
// Deprecated: Use ISwapRouterPancakeMetaData.ABI instead.
var ISwapRouterPancakeABI = ISwapRouterPancakeMetaData.ABI

// ISwapRouterPancake is an auto generated Go binding around an Ethereum contract.
type ISwapRouterPancake struct {
	ISwapRouterPancakeCaller     // Read-only binding to the contract
	ISwapRouterPancakeTransactor // Write-only binding to the contract
	ISwapRouterPancakeFilterer   // Log filterer for contract events
}

// ISwapRouterPancakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISwapRouterPancakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapRouterPancakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISwapRouterPancakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapRouterPancakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISwapRouterPancakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapRouterPancakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISwapRouterPancakeSession struct {
	Contract     *ISwapRouterPancake // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ISwapRouterPancakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISwapRouterPancakeCallerSession struct {
	Contract *ISwapRouterPancakeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ISwapRouterPancakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISwapRouterPancakeTransactorSession struct {
	Contract     *ISwapRouterPancakeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ISwapRouterPancakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISwapRouterPancakeRaw struct {
	Contract *ISwapRouterPancake // Generic contract binding to access the raw methods on
}

// ISwapRouterPancakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISwapRouterPancakeCallerRaw struct {
	Contract *ISwapRouterPancakeCaller // Generic read-only contract binding to access the raw methods on
}

// ISwapRouterPancakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISwapRouterPancakeTransactorRaw struct {
	Contract *ISwapRouterPancakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISwapRouterPancake creates a new instance of ISwapRouterPancake, bound to a specific deployed contract.
func NewISwapRouterPancake(address common.Address, backend bind.ContractBackend) (*ISwapRouterPancake, error) {
	contract, err := bindISwapRouterPancake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISwapRouterPancake{ISwapRouterPancakeCaller: ISwapRouterPancakeCaller{contract: contract}, ISwapRouterPancakeTransactor: ISwapRouterPancakeTransactor{contract: contract}, ISwapRouterPancakeFilterer: ISwapRouterPancakeFilterer{contract: contract}}, nil
}

// NewISwapRouterPancakeCaller creates a new read-only instance of ISwapRouterPancake, bound to a specific deployed contract.
func NewISwapRouterPancakeCaller(address common.Address, caller bind.ContractCaller) (*ISwapRouterPancakeCaller, error) {
	contract, err := bindISwapRouterPancake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISwapRouterPancakeCaller{contract: contract}, nil
}

// NewISwapRouterPancakeTransactor creates a new write-only instance of ISwapRouterPancake, bound to a specific deployed contract.
func NewISwapRouterPancakeTransactor(address common.Address, transactor bind.ContractTransactor) (*ISwapRouterPancakeTransactor, error) {
	contract, err := bindISwapRouterPancake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISwapRouterPancakeTransactor{contract: contract}, nil
}

// NewISwapRouterPancakeFilterer creates a new log filterer instance of ISwapRouterPancake, bound to a specific deployed contract.
func NewISwapRouterPancakeFilterer(address common.Address, filterer bind.ContractFilterer) (*ISwapRouterPancakeFilterer, error) {
	contract, err := bindISwapRouterPancake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISwapRouterPancakeFilterer{contract: contract}, nil
}

// bindISwapRouterPancake binds a generic wrapper to an already deployed contract.
func bindISwapRouterPancake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISwapRouterPancakeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwapRouterPancake *ISwapRouterPancakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwapRouterPancake.Contract.ISwapRouterPancakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwapRouterPancake *ISwapRouterPancakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwapRouterPancake.Contract.ISwapRouterPancakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwapRouterPancake *ISwapRouterPancakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwapRouterPancake.Contract.ISwapRouterPancakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwapRouterPancake *ISwapRouterPancakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwapRouterPancake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwapRouterPancake *ISwapRouterPancakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwapRouterPancake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwapRouterPancake *ISwapRouterPancakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwapRouterPancake.Contract.contract.Transact(opts, method, params...)
}

// ExactInput is a paid mutator transaction binding the contract method 0xb858183f.
//
// Solidity: function exactInput((bytes,address,uint256,uint256) params) payable returns(uint256 amountOut)
func (_ISwapRouterPancake *ISwapRouterPancakeTransactor) ExactInput(opts *bind.TransactOpts, params ISwapRouterPancakeExactInputParams) (*types.Transaction, error) {
	return _ISwapRouterPancake.contract.Transact(opts, "exactInput", params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xb858183f.
//
// Solidity: function exactInput((bytes,address,uint256,uint256) params) payable returns(uint256 amountOut)
func (_ISwapRouterPancake *ISwapRouterPancakeSession) ExactInput(params ISwapRouterPancakeExactInputParams) (*types.Transaction, error) {
	return _ISwapRouterPancake.Contract.ExactInput(&_ISwapRouterPancake.TransactOpts, params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xb858183f.
//
// Solidity: function exactInput((bytes,address,uint256,uint256) params) payable returns(uint256 amountOut)
func (_ISwapRouterPancake *ISwapRouterPancakeTransactorSession) ExactInput(params ISwapRouterPancakeExactInputParams) (*types.Transaction, error) {
	return _ISwapRouterPancake.Contract.ExactInput(&_ISwapRouterPancake.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x04e45aaf.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_ISwapRouterPancake *ISwapRouterPancakeTransactor) ExactInputSingle(opts *bind.TransactOpts, params ISwapRouterPancakeExactInputSingleParams) (*types.Transaction, error) {
	return _ISwapRouterPancake.contract.Transact(opts, "exactInputSingle", params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x04e45aaf.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_ISwapRouterPancake *ISwapRouterPancakeSession) ExactInputSingle(params ISwapRouterPancakeExactInputSingleParams) (*types.Transaction, error) {
	return _ISwapRouterPancake.Contract.ExactInputSingle(&_ISwapRouterPancake.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x04e45aaf.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_ISwapRouterPancake *ISwapRouterPancakeTransactorSession) ExactInputSingle(params ISwapRouterPancakeExactInputSingleParams) (*types.Transaction, error) {
	return _ISwapRouterPancake.Contract.ExactInputSingle(&_ISwapRouterPancake.TransactOpts, params)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_ISwapRouterPancake *ISwapRouterPancakeTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _ISwapRouterPancake.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_ISwapRouterPancake *ISwapRouterPancakeSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _ISwapRouterPancake.Contract.UniswapV3SwapCallback(&_ISwapRouterPancake.TransactOpts, amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_ISwapRouterPancake *ISwapRouterPancakeTransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _ISwapRouterPancake.Contract.UniswapV3SwapCallback(&_ISwapRouterPancake.TransactOpts, amount0Delta, amount1Delta, data)
}
