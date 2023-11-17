// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetaconnectorzevm

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

// ZetaInterfacesMetaData contains all meta data concerning the ZetaInterfaces contract.
var ZetaInterfacesMetaData = &bind.MetaData{
	ABI: "[]",
}

// ZetaInterfacesABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaInterfacesMetaData.ABI instead.
var ZetaInterfacesABI = ZetaInterfacesMetaData.ABI

// ZetaInterfaces is an auto generated Go binding around an Ethereum contract.
type ZetaInterfaces struct {
	ZetaInterfacesCaller     // Read-only binding to the contract
	ZetaInterfacesTransactor // Write-only binding to the contract
	ZetaInterfacesFilterer   // Log filterer for contract events
}

// ZetaInterfacesCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaInterfacesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaInterfacesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaInterfacesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaInterfacesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaInterfacesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaInterfacesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaInterfacesSession struct {
	Contract     *ZetaInterfaces   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZetaInterfacesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaInterfacesCallerSession struct {
	Contract *ZetaInterfacesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ZetaInterfacesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaInterfacesTransactorSession struct {
	Contract     *ZetaInterfacesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ZetaInterfacesRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaInterfacesRaw struct {
	Contract *ZetaInterfaces // Generic contract binding to access the raw methods on
}

// ZetaInterfacesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaInterfacesCallerRaw struct {
	Contract *ZetaInterfacesCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaInterfacesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaInterfacesTransactorRaw struct {
	Contract *ZetaInterfacesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaInterfaces creates a new instance of ZetaInterfaces, bound to a specific deployed contract.
func NewZetaInterfaces(address common.Address, backend bind.ContractBackend) (*ZetaInterfaces, error) {
	contract, err := bindZetaInterfaces(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaInterfaces{ZetaInterfacesCaller: ZetaInterfacesCaller{contract: contract}, ZetaInterfacesTransactor: ZetaInterfacesTransactor{contract: contract}, ZetaInterfacesFilterer: ZetaInterfacesFilterer{contract: contract}}, nil
}

// NewZetaInterfacesCaller creates a new read-only instance of ZetaInterfaces, bound to a specific deployed contract.
func NewZetaInterfacesCaller(address common.Address, caller bind.ContractCaller) (*ZetaInterfacesCaller, error) {
	contract, err := bindZetaInterfaces(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaInterfacesCaller{contract: contract}, nil
}

// NewZetaInterfacesTransactor creates a new write-only instance of ZetaInterfaces, bound to a specific deployed contract.
func NewZetaInterfacesTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaInterfacesTransactor, error) {
	contract, err := bindZetaInterfaces(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaInterfacesTransactor{contract: contract}, nil
}

// NewZetaInterfacesFilterer creates a new log filterer instance of ZetaInterfaces, bound to a specific deployed contract.
func NewZetaInterfacesFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaInterfacesFilterer, error) {
	contract, err := bindZetaInterfaces(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaInterfacesFilterer{contract: contract}, nil
}

// bindZetaInterfaces binds a generic wrapper to an already deployed contract.
func bindZetaInterfaces(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaInterfacesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaInterfaces *ZetaInterfacesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaInterfaces.Contract.ZetaInterfacesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaInterfaces *ZetaInterfacesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaInterfaces.Contract.ZetaInterfacesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaInterfaces *ZetaInterfacesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaInterfaces.Contract.ZetaInterfacesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaInterfaces *ZetaInterfacesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaInterfaces.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaInterfaces *ZetaInterfacesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaInterfaces.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaInterfaces *ZetaInterfacesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaInterfaces.Contract.contract.Transact(opts, method, params...)
}
