// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package attackercontract

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

// VictimMetaData contains all meta data concerning the Victim contract.
var VictimMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VictimABI is the input ABI used to generate the binding from.
// Deprecated: Use VictimMetaData.ABI instead.
var VictimABI = VictimMetaData.ABI

// Victim is an auto generated Go binding around an Ethereum contract.
type Victim struct {
	VictimCaller     // Read-only binding to the contract
	VictimTransactor // Write-only binding to the contract
	VictimFilterer   // Log filterer for contract events
}

// VictimCaller is an auto generated read-only Go binding around an Ethereum contract.
type VictimCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VictimTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VictimTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VictimFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VictimFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VictimSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VictimSession struct {
	Contract     *Victim           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VictimCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VictimCallerSession struct {
	Contract *VictimCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VictimTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VictimTransactorSession struct {
	Contract     *VictimTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VictimRaw is an auto generated low-level Go binding around an Ethereum contract.
type VictimRaw struct {
	Contract *Victim // Generic contract binding to access the raw methods on
}

// VictimCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VictimCallerRaw struct {
	Contract *VictimCaller // Generic read-only contract binding to access the raw methods on
}

// VictimTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VictimTransactorRaw struct {
	Contract *VictimTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVictim creates a new instance of Victim, bound to a specific deployed contract.
func NewVictim(address common.Address, backend bind.ContractBackend) (*Victim, error) {
	contract, err := bindVictim(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Victim{VictimCaller: VictimCaller{contract: contract}, VictimTransactor: VictimTransactor{contract: contract}, VictimFilterer: VictimFilterer{contract: contract}}, nil
}

// NewVictimCaller creates a new read-only instance of Victim, bound to a specific deployed contract.
func NewVictimCaller(address common.Address, caller bind.ContractCaller) (*VictimCaller, error) {
	contract, err := bindVictim(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VictimCaller{contract: contract}, nil
}

// NewVictimTransactor creates a new write-only instance of Victim, bound to a specific deployed contract.
func NewVictimTransactor(address common.Address, transactor bind.ContractTransactor) (*VictimTransactor, error) {
	contract, err := bindVictim(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VictimTransactor{contract: contract}, nil
}

// NewVictimFilterer creates a new log filterer instance of Victim, bound to a specific deployed contract.
func NewVictimFilterer(address common.Address, filterer bind.ContractFilterer) (*VictimFilterer, error) {
	contract, err := bindVictim(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VictimFilterer{contract: contract}, nil
}

// bindVictim binds a generic wrapper to an already deployed contract.
func bindVictim(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VictimMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Victim *VictimRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Victim.Contract.VictimCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Victim *VictimRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Victim.Contract.VictimTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Victim *VictimRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Victim.Contract.VictimTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Victim *VictimCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Victim.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Victim *VictimTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Victim.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Victim *VictimTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Victim.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0xe609055e.
//
// Solidity: function deposit(bytes recipient, address asset, uint256 amount, bytes message) returns()
func (_Victim *VictimTransactor) Deposit(opts *bind.TransactOpts, recipient []byte, asset common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _Victim.contract.Transact(opts, "deposit", recipient, asset, amount, message)
}

// Deposit is a paid mutator transaction binding the contract method 0xe609055e.
//
// Solidity: function deposit(bytes recipient, address asset, uint256 amount, bytes message) returns()
func (_Victim *VictimSession) Deposit(recipient []byte, asset common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _Victim.Contract.Deposit(&_Victim.TransactOpts, recipient, asset, amount, message)
}

// Deposit is a paid mutator transaction binding the contract method 0xe609055e.
//
// Solidity: function deposit(bytes recipient, address asset, uint256 amount, bytes message) returns()
func (_Victim *VictimTransactorSession) Deposit(recipient []byte, asset common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _Victim.Contract.Deposit(&_Victim.TransactOpts, recipient, asset, amount, message)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address asset, uint256 amount) returns()
func (_Victim *VictimTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Victim.contract.Transact(opts, "withdraw", recipient, asset, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address asset, uint256 amount) returns()
func (_Victim *VictimSession) Withdraw(recipient common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Victim.Contract.Withdraw(&_Victim.TransactOpts, recipient, asset, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address asset, uint256 amount) returns()
func (_Victim *VictimTransactorSession) Withdraw(recipient common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Victim.Contract.Withdraw(&_Victim.TransactOpts, recipient, asset, amount)
}
