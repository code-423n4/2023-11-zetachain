// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswap

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

// UniswapImportsMetaData contains all meta data concerning the UniswapImports contract.
var UniswapImportsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080604052348015600f57600080fd5b50603e80601d6000396000f3fe6080604052600080fdfea265627a7a72315820e9232a161dc235f259a171daf4cc9c53677aabc950914989e6b4f90c1688088664736f6c63430005100032",
}

// UniswapImportsABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapImportsMetaData.ABI instead.
var UniswapImportsABI = UniswapImportsMetaData.ABI

// UniswapImportsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UniswapImportsMetaData.Bin instead.
var UniswapImportsBin = UniswapImportsMetaData.Bin

// DeployUniswapImports deploys a new Ethereum contract, binding an instance of UniswapImports to it.
func DeployUniswapImports(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UniswapImports, error) {
	parsed, err := UniswapImportsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UniswapImportsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UniswapImports{UniswapImportsCaller: UniswapImportsCaller{contract: contract}, UniswapImportsTransactor: UniswapImportsTransactor{contract: contract}, UniswapImportsFilterer: UniswapImportsFilterer{contract: contract}}, nil
}

// UniswapImports is an auto generated Go binding around an Ethereum contract.
type UniswapImports struct {
	UniswapImportsCaller     // Read-only binding to the contract
	UniswapImportsTransactor // Write-only binding to the contract
	UniswapImportsFilterer   // Log filterer for contract events
}

// UniswapImportsCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapImportsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapImportsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapImportsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapImportsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapImportsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapImportsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapImportsSession struct {
	Contract     *UniswapImports   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapImportsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapImportsCallerSession struct {
	Contract *UniswapImportsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// UniswapImportsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapImportsTransactorSession struct {
	Contract     *UniswapImportsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// UniswapImportsRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapImportsRaw struct {
	Contract *UniswapImports // Generic contract binding to access the raw methods on
}

// UniswapImportsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapImportsCallerRaw struct {
	Contract *UniswapImportsCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapImportsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapImportsTransactorRaw struct {
	Contract *UniswapImportsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapImports creates a new instance of UniswapImports, bound to a specific deployed contract.
func NewUniswapImports(address common.Address, backend bind.ContractBackend) (*UniswapImports, error) {
	contract, err := bindUniswapImports(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapImports{UniswapImportsCaller: UniswapImportsCaller{contract: contract}, UniswapImportsTransactor: UniswapImportsTransactor{contract: contract}, UniswapImportsFilterer: UniswapImportsFilterer{contract: contract}}, nil
}

// NewUniswapImportsCaller creates a new read-only instance of UniswapImports, bound to a specific deployed contract.
func NewUniswapImportsCaller(address common.Address, caller bind.ContractCaller) (*UniswapImportsCaller, error) {
	contract, err := bindUniswapImports(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapImportsCaller{contract: contract}, nil
}

// NewUniswapImportsTransactor creates a new write-only instance of UniswapImports, bound to a specific deployed contract.
func NewUniswapImportsTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapImportsTransactor, error) {
	contract, err := bindUniswapImports(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapImportsTransactor{contract: contract}, nil
}

// NewUniswapImportsFilterer creates a new log filterer instance of UniswapImports, bound to a specific deployed contract.
func NewUniswapImportsFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapImportsFilterer, error) {
	contract, err := bindUniswapImports(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapImportsFilterer{contract: contract}, nil
}

// bindUniswapImports binds a generic wrapper to an already deployed contract.
func bindUniswapImports(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniswapImportsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapImports *UniswapImportsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapImports.Contract.UniswapImportsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapImports *UniswapImportsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapImports.Contract.UniswapImportsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapImports *UniswapImportsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapImports.Contract.UniswapImportsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapImports *UniswapImportsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapImports.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapImports *UniswapImportsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapImports.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapImports *UniswapImportsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapImports.Contract.contract.Transact(opts, method, params...)
}
