// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tridentipoolrouter

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

// IPoolRouterExactInputParams is an auto generated low-level Go binding around an user-defined struct.
type IPoolRouterExactInputParams struct {
	TokenIn          common.Address
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
	Path             []common.Address
	To               common.Address
	Unwrap           bool
}

// IPoolRouterExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IPoolRouterExactInputSingleParams struct {
	TokenIn          common.Address
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
	Pool             common.Address
	To               common.Address
	Unwrap           bool
}

// IPoolRouterExactOutputParams is an auto generated low-level Go binding around an user-defined struct.
type IPoolRouterExactOutputParams struct {
	TokenIn         common.Address
	AmountOut       *big.Int
	AmountInMaximum *big.Int
	Path            []common.Address
	To              common.Address
	Unwrap          bool
}

// IPoolRouterExactOutputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IPoolRouterExactOutputSingleParams struct {
	TokenIn         common.Address
	AmountOut       *big.Int
	AmountInMaximum *big.Int
	Pool            common.Address
	To              common.Address
	Unwrap          bool
}

// IPoolRouterMetaData contains all meta data concerning the IPoolRouter contract.
var IPoolRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"unwrap\",\"type\":\"bool\"}],\"internalType\":\"structIPoolRouter.ExactInputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"unwrap\",\"type\":\"bool\"}],\"internalType\":\"structIPoolRouter.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"unwrap\",\"type\":\"bool\"}],\"internalType\":\"structIPoolRouter.ExactOutputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"unwrap\",\"type\":\"bool\"}],\"internalType\":\"structIPoolRouter.ExactOutputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweep\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IPoolRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use IPoolRouterMetaData.ABI instead.
var IPoolRouterABI = IPoolRouterMetaData.ABI

// IPoolRouter is an auto generated Go binding around an Ethereum contract.
type IPoolRouter struct {
	IPoolRouterCaller     // Read-only binding to the contract
	IPoolRouterTransactor // Write-only binding to the contract
	IPoolRouterFilterer   // Log filterer for contract events
}

// IPoolRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPoolRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPoolRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPoolRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPoolRouterSession struct {
	Contract     *IPoolRouter      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPoolRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPoolRouterCallerSession struct {
	Contract *IPoolRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IPoolRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPoolRouterTransactorSession struct {
	Contract     *IPoolRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IPoolRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPoolRouterRaw struct {
	Contract *IPoolRouter // Generic contract binding to access the raw methods on
}

// IPoolRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPoolRouterCallerRaw struct {
	Contract *IPoolRouterCaller // Generic read-only contract binding to access the raw methods on
}

// IPoolRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPoolRouterTransactorRaw struct {
	Contract *IPoolRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPoolRouter creates a new instance of IPoolRouter, bound to a specific deployed contract.
func NewIPoolRouter(address common.Address, backend bind.ContractBackend) (*IPoolRouter, error) {
	contract, err := bindIPoolRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPoolRouter{IPoolRouterCaller: IPoolRouterCaller{contract: contract}, IPoolRouterTransactor: IPoolRouterTransactor{contract: contract}, IPoolRouterFilterer: IPoolRouterFilterer{contract: contract}}, nil
}

// NewIPoolRouterCaller creates a new read-only instance of IPoolRouter, bound to a specific deployed contract.
func NewIPoolRouterCaller(address common.Address, caller bind.ContractCaller) (*IPoolRouterCaller, error) {
	contract, err := bindIPoolRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPoolRouterCaller{contract: contract}, nil
}

// NewIPoolRouterTransactor creates a new write-only instance of IPoolRouter, bound to a specific deployed contract.
func NewIPoolRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*IPoolRouterTransactor, error) {
	contract, err := bindIPoolRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPoolRouterTransactor{contract: contract}, nil
}

// NewIPoolRouterFilterer creates a new log filterer instance of IPoolRouter, bound to a specific deployed contract.
func NewIPoolRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*IPoolRouterFilterer, error) {
	contract, err := bindIPoolRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPoolRouterFilterer{contract: contract}, nil
}

// bindIPoolRouter binds a generic wrapper to an already deployed contract.
func bindIPoolRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPoolRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPoolRouter *IPoolRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPoolRouter.Contract.IPoolRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPoolRouter *IPoolRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPoolRouter.Contract.IPoolRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPoolRouter *IPoolRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPoolRouter.Contract.IPoolRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPoolRouter *IPoolRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPoolRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPoolRouter *IPoolRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPoolRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPoolRouter *IPoolRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPoolRouter.Contract.contract.Transact(opts, method, params...)
}

// ExactInput is a paid mutator transaction binding the contract method 0x363a9dba.
//
// Solidity: function exactInput((address,uint256,uint256,address[],address,bool) params) payable returns(uint256 amountOut)
func (_IPoolRouter *IPoolRouterTransactor) ExactInput(opts *bind.TransactOpts, params IPoolRouterExactInputParams) (*types.Transaction, error) {
	return _IPoolRouter.contract.Transact(opts, "exactInput", params)
}

// ExactInput is a paid mutator transaction binding the contract method 0x363a9dba.
//
// Solidity: function exactInput((address,uint256,uint256,address[],address,bool) params) payable returns(uint256 amountOut)
func (_IPoolRouter *IPoolRouterSession) ExactInput(params IPoolRouterExactInputParams) (*types.Transaction, error) {
	return _IPoolRouter.Contract.ExactInput(&_IPoolRouter.TransactOpts, params)
}

// ExactInput is a paid mutator transaction binding the contract method 0x363a9dba.
//
// Solidity: function exactInput((address,uint256,uint256,address[],address,bool) params) payable returns(uint256 amountOut)
func (_IPoolRouter *IPoolRouterTransactorSession) ExactInput(params IPoolRouterExactInputParams) (*types.Transaction, error) {
	return _IPoolRouter.Contract.ExactInput(&_IPoolRouter.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0xc07f5c32.
//
// Solidity: function exactInputSingle((address,uint256,uint256,address,address,bool) params) payable returns(uint256 amountOut)
func (_IPoolRouter *IPoolRouterTransactor) ExactInputSingle(opts *bind.TransactOpts, params IPoolRouterExactInputSingleParams) (*types.Transaction, error) {
	return _IPoolRouter.contract.Transact(opts, "exactInputSingle", params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0xc07f5c32.
//
// Solidity: function exactInputSingle((address,uint256,uint256,address,address,bool) params) payable returns(uint256 amountOut)
func (_IPoolRouter *IPoolRouterSession) ExactInputSingle(params IPoolRouterExactInputSingleParams) (*types.Transaction, error) {
	return _IPoolRouter.Contract.ExactInputSingle(&_IPoolRouter.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0xc07f5c32.
//
// Solidity: function exactInputSingle((address,uint256,uint256,address,address,bool) params) payable returns(uint256 amountOut)
func (_IPoolRouter *IPoolRouterTransactorSession) ExactInputSingle(params IPoolRouterExactInputSingleParams) (*types.Transaction, error) {
	return _IPoolRouter.Contract.ExactInputSingle(&_IPoolRouter.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xbee20e05.
//
// Solidity: function exactOutput((address,uint256,uint256,address[],address,bool) params) payable returns(uint256 amountIn)
func (_IPoolRouter *IPoolRouterTransactor) ExactOutput(opts *bind.TransactOpts, params IPoolRouterExactOutputParams) (*types.Transaction, error) {
	return _IPoolRouter.contract.Transact(opts, "exactOutput", params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xbee20e05.
//
// Solidity: function exactOutput((address,uint256,uint256,address[],address,bool) params) payable returns(uint256 amountIn)
func (_IPoolRouter *IPoolRouterSession) ExactOutput(params IPoolRouterExactOutputParams) (*types.Transaction, error) {
	return _IPoolRouter.Contract.ExactOutput(&_IPoolRouter.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xbee20e05.
//
// Solidity: function exactOutput((address,uint256,uint256,address[],address,bool) params) payable returns(uint256 amountIn)
func (_IPoolRouter *IPoolRouterTransactorSession) ExactOutput(params IPoolRouterExactOutputParams) (*types.Transaction, error) {
	return _IPoolRouter.Contract.ExactOutput(&_IPoolRouter.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0x54c1b650.
//
// Solidity: function exactOutputSingle((address,uint256,uint256,address,address,bool) params) payable returns(uint256 amountIn)
func (_IPoolRouter *IPoolRouterTransactor) ExactOutputSingle(opts *bind.TransactOpts, params IPoolRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _IPoolRouter.contract.Transact(opts, "exactOutputSingle", params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0x54c1b650.
//
// Solidity: function exactOutputSingle((address,uint256,uint256,address,address,bool) params) payable returns(uint256 amountIn)
func (_IPoolRouter *IPoolRouterSession) ExactOutputSingle(params IPoolRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _IPoolRouter.Contract.ExactOutputSingle(&_IPoolRouter.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0x54c1b650.
//
// Solidity: function exactOutputSingle((address,uint256,uint256,address,address,bool) params) payable returns(uint256 amountIn)
func (_IPoolRouter *IPoolRouterTransactorSession) ExactOutputSingle(params IPoolRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _IPoolRouter.Contract.ExactOutputSingle(&_IPoolRouter.TransactOpts, params)
}

// Sweep is a paid mutator transaction binding the contract method 0xdc2c256f.
//
// Solidity: function sweep(address token, uint256 amount, address recipient) payable returns()
func (_IPoolRouter *IPoolRouterTransactor) Sweep(opts *bind.TransactOpts, token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IPoolRouter.contract.Transact(opts, "sweep", token, amount, recipient)
}

// Sweep is a paid mutator transaction binding the contract method 0xdc2c256f.
//
// Solidity: function sweep(address token, uint256 amount, address recipient) payable returns()
func (_IPoolRouter *IPoolRouterSession) Sweep(token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IPoolRouter.Contract.Sweep(&_IPoolRouter.TransactOpts, token, amount, recipient)
}

// Sweep is a paid mutator transaction binding the contract method 0xdc2c256f.
//
// Solidity: function sweep(address token, uint256 amount, address recipient) payable returns()
func (_IPoolRouter *IPoolRouterTransactorSession) Sweep(token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IPoolRouter.Contract.Sweep(&_IPoolRouter.TransactOpts, token, amount, recipient)
}
