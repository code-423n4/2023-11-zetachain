// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package interfaces

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

// ISystemMetaData contains all meta data concerning the ISystem contract.
var ISystemMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"FUNGIBLE_MODULE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"gasCoinZRC20ByChainId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"gasPriceByChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"gasZetaPoolByChainId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapv2FactoryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wZetaContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ISystemABI is the input ABI used to generate the binding from.
// Deprecated: Use ISystemMetaData.ABI instead.
var ISystemABI = ISystemMetaData.ABI

// ISystem is an auto generated Go binding around an Ethereum contract.
type ISystem struct {
	ISystemCaller     // Read-only binding to the contract
	ISystemTransactor // Write-only binding to the contract
	ISystemFilterer   // Log filterer for contract events
}

// ISystemCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISystemCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISystemTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISystemTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISystemFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISystemFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISystemSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISystemSession struct {
	Contract     *ISystem          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISystemCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISystemCallerSession struct {
	Contract *ISystemCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ISystemTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISystemTransactorSession struct {
	Contract     *ISystemTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ISystemRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISystemRaw struct {
	Contract *ISystem // Generic contract binding to access the raw methods on
}

// ISystemCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISystemCallerRaw struct {
	Contract *ISystemCaller // Generic read-only contract binding to access the raw methods on
}

// ISystemTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISystemTransactorRaw struct {
	Contract *ISystemTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISystem creates a new instance of ISystem, bound to a specific deployed contract.
func NewISystem(address common.Address, backend bind.ContractBackend) (*ISystem, error) {
	contract, err := bindISystem(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISystem{ISystemCaller: ISystemCaller{contract: contract}, ISystemTransactor: ISystemTransactor{contract: contract}, ISystemFilterer: ISystemFilterer{contract: contract}}, nil
}

// NewISystemCaller creates a new read-only instance of ISystem, bound to a specific deployed contract.
func NewISystemCaller(address common.Address, caller bind.ContractCaller) (*ISystemCaller, error) {
	contract, err := bindISystem(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISystemCaller{contract: contract}, nil
}

// NewISystemTransactor creates a new write-only instance of ISystem, bound to a specific deployed contract.
func NewISystemTransactor(address common.Address, transactor bind.ContractTransactor) (*ISystemTransactor, error) {
	contract, err := bindISystem(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISystemTransactor{contract: contract}, nil
}

// NewISystemFilterer creates a new log filterer instance of ISystem, bound to a specific deployed contract.
func NewISystemFilterer(address common.Address, filterer bind.ContractFilterer) (*ISystemFilterer, error) {
	contract, err := bindISystem(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISystemFilterer{contract: contract}, nil
}

// bindISystem binds a generic wrapper to an already deployed contract.
func bindISystem(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISystemMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISystem *ISystemRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISystem.Contract.ISystemCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISystem *ISystemRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISystem.Contract.ISystemTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISystem *ISystemRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISystem.Contract.ISystemTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISystem *ISystemCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISystem.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISystem *ISystemTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISystem.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISystem *ISystemTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISystem.Contract.contract.Transact(opts, method, params...)
}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_ISystem *ISystemCaller) FUNGIBLEMODULEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISystem.contract.Call(opts, &out, "FUNGIBLE_MODULE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_ISystem *ISystemSession) FUNGIBLEMODULEADDRESS() (common.Address, error) {
	return _ISystem.Contract.FUNGIBLEMODULEADDRESS(&_ISystem.CallOpts)
}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_ISystem *ISystemCallerSession) FUNGIBLEMODULEADDRESS() (common.Address, error) {
	return _ISystem.Contract.FUNGIBLEMODULEADDRESS(&_ISystem.CallOpts)
}

// GasCoinZRC20ByChainId is a free data retrieval call binding the contract method 0x0be15547.
//
// Solidity: function gasCoinZRC20ByChainId(uint256 chainID) view returns(address)
func (_ISystem *ISystemCaller) GasCoinZRC20ByChainId(opts *bind.CallOpts, chainID *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ISystem.contract.Call(opts, &out, "gasCoinZRC20ByChainId", chainID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasCoinZRC20ByChainId is a free data retrieval call binding the contract method 0x0be15547.
//
// Solidity: function gasCoinZRC20ByChainId(uint256 chainID) view returns(address)
func (_ISystem *ISystemSession) GasCoinZRC20ByChainId(chainID *big.Int) (common.Address, error) {
	return _ISystem.Contract.GasCoinZRC20ByChainId(&_ISystem.CallOpts, chainID)
}

// GasCoinZRC20ByChainId is a free data retrieval call binding the contract method 0x0be15547.
//
// Solidity: function gasCoinZRC20ByChainId(uint256 chainID) view returns(address)
func (_ISystem *ISystemCallerSession) GasCoinZRC20ByChainId(chainID *big.Int) (common.Address, error) {
	return _ISystem.Contract.GasCoinZRC20ByChainId(&_ISystem.CallOpts, chainID)
}

// GasPriceByChainId is a free data retrieval call binding the contract method 0xd7fd7afb.
//
// Solidity: function gasPriceByChainId(uint256 chainID) view returns(uint256)
func (_ISystem *ISystemCaller) GasPriceByChainId(opts *bind.CallOpts, chainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ISystem.contract.Call(opts, &out, "gasPriceByChainId", chainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasPriceByChainId is a free data retrieval call binding the contract method 0xd7fd7afb.
//
// Solidity: function gasPriceByChainId(uint256 chainID) view returns(uint256)
func (_ISystem *ISystemSession) GasPriceByChainId(chainID *big.Int) (*big.Int, error) {
	return _ISystem.Contract.GasPriceByChainId(&_ISystem.CallOpts, chainID)
}

// GasPriceByChainId is a free data retrieval call binding the contract method 0xd7fd7afb.
//
// Solidity: function gasPriceByChainId(uint256 chainID) view returns(uint256)
func (_ISystem *ISystemCallerSession) GasPriceByChainId(chainID *big.Int) (*big.Int, error) {
	return _ISystem.Contract.GasPriceByChainId(&_ISystem.CallOpts, chainID)
}

// GasZetaPoolByChainId is a free data retrieval call binding the contract method 0x513a9c05.
//
// Solidity: function gasZetaPoolByChainId(uint256 chainID) view returns(address)
func (_ISystem *ISystemCaller) GasZetaPoolByChainId(opts *bind.CallOpts, chainID *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ISystem.contract.Call(opts, &out, "gasZetaPoolByChainId", chainID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasZetaPoolByChainId is a free data retrieval call binding the contract method 0x513a9c05.
//
// Solidity: function gasZetaPoolByChainId(uint256 chainID) view returns(address)
func (_ISystem *ISystemSession) GasZetaPoolByChainId(chainID *big.Int) (common.Address, error) {
	return _ISystem.Contract.GasZetaPoolByChainId(&_ISystem.CallOpts, chainID)
}

// GasZetaPoolByChainId is a free data retrieval call binding the contract method 0x513a9c05.
//
// Solidity: function gasZetaPoolByChainId(uint256 chainID) view returns(address)
func (_ISystem *ISystemCallerSession) GasZetaPoolByChainId(chainID *big.Int) (common.Address, error) {
	return _ISystem.Contract.GasZetaPoolByChainId(&_ISystem.CallOpts, chainID)
}

// Uniswapv2FactoryAddress is a free data retrieval call binding the contract method 0xd936a012.
//
// Solidity: function uniswapv2FactoryAddress() view returns(address)
func (_ISystem *ISystemCaller) Uniswapv2FactoryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISystem.contract.Call(opts, &out, "uniswapv2FactoryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Uniswapv2FactoryAddress is a free data retrieval call binding the contract method 0xd936a012.
//
// Solidity: function uniswapv2FactoryAddress() view returns(address)
func (_ISystem *ISystemSession) Uniswapv2FactoryAddress() (common.Address, error) {
	return _ISystem.Contract.Uniswapv2FactoryAddress(&_ISystem.CallOpts)
}

// Uniswapv2FactoryAddress is a free data retrieval call binding the contract method 0xd936a012.
//
// Solidity: function uniswapv2FactoryAddress() view returns(address)
func (_ISystem *ISystemCallerSession) Uniswapv2FactoryAddress() (common.Address, error) {
	return _ISystem.Contract.Uniswapv2FactoryAddress(&_ISystem.CallOpts)
}

// WZetaContractAddress is a free data retrieval call binding the contract method 0x569541b9.
//
// Solidity: function wZetaContractAddress() view returns(address)
func (_ISystem *ISystemCaller) WZetaContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISystem.contract.Call(opts, &out, "wZetaContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WZetaContractAddress is a free data retrieval call binding the contract method 0x569541b9.
//
// Solidity: function wZetaContractAddress() view returns(address)
func (_ISystem *ISystemSession) WZetaContractAddress() (common.Address, error) {
	return _ISystem.Contract.WZetaContractAddress(&_ISystem.CallOpts)
}

// WZetaContractAddress is a free data retrieval call binding the contract method 0x569541b9.
//
// Solidity: function wZetaContractAddress() view returns(address)
func (_ISystem *ISystemCallerSession) WZetaContractAddress() (common.Address, error) {
	return _ISystem.Contract.WZetaContractAddress(&_ISystem.CallOpts)
}
