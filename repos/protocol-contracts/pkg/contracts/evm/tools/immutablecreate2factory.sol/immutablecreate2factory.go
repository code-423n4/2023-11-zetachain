// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package immutablecreate2factory

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

// ImmutableCreate2FactoryMetaData contains all meta data concerning the ImmutableCreate2Factory contract.
var ImmutableCreate2FactoryMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"deploymentAddress\",\"type\":\"address\"}],\"name\":\"hasBeenDeployed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\"},{\"name\":\"initializationCode\",\"type\":\"bytes\"}],\"name\":\"safeCreate2AndTransfer\",\"outputs\":[{\"name\":\"deploymentAddress\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\"},{\"name\":\"initializationCode\",\"type\":\"bytes\"}],\"name\":\"safeCreate2\",\"outputs\":[{\"name\":\"deploymentAddress\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\"},{\"name\":\"initCode\",\"type\":\"bytes\"}],\"name\":\"findCreate2Address\",\"outputs\":[{\"name\":\"deploymentAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\"},{\"name\":\"initCodeHash\",\"type\":\"bytes32\"}],\"name\":\"findCreate2AddressViaHash\",\"outputs\":[{\"name\":\"deploymentAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610c54806100206000396000f3fe60806040526004361061004a5760003560e01c806308508b8f1461004f57806329346003146100b857806364e030871461017b57806385cf97ab14610280578063a49a7c9014610350575b600080fd5b34801561005b57600080fd5b5061009e6004803603602081101561007257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506103d5565b604051808215151515815260200191505060405180910390f35b610139600480360360408110156100ce57600080fd5b8101908080359060200190929190803590602001906401000000008111156100f557600080fd5b82018360208201111561010757600080fd5b8035906020019184600183028401116401000000008311171561012957600080fd5b909192939192939050505061042a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61023e6004803603604081101561019157600080fd5b8101908080359060200190929190803590602001906401000000008111156101b857600080fd5b8201836020820111156101ca57600080fd5b803590602001918460018302840111640100000000831117156101ec57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506105cf565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561028c57600080fd5b5061030e600480360360408110156102a357600080fd5b8101908080359060200190929190803590602001906401000000008111156102ca57600080fd5b8201836020820111156102dc57600080fd5b803590602001918460018302840111640100000000831117156102fe57600080fd5b9091929391929390505050610698565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561035c57600080fd5b506103936004803603604081101561037357600080fd5b8101908080359060200190929190803590602001909291905050506107be565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b6000833373ffffffffffffffffffffffffffffffffffffffff168160601c73ffffffffffffffffffffffffffffffffffffffff16148061048b5750600060601b6bffffffffffffffffffffffff1916816bffffffffffffffffffffffff1916145b6104e0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526045815260200180610b956045913960600191505060405180910390fd5b61052e8585858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506108b4565b91508173ffffffffffffffffffffffffffffffffffffffff1663f2fde38b336040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b1580156105af57600080fd5b505af11580156105c3573d6000803e3d6000fd5b50505050509392505050565b6000823373ffffffffffffffffffffffffffffffffffffffff168160601c73ffffffffffffffffffffffffffffffffffffffff1614806106305750600060601b6bffffffffffffffffffffffff1916816bffffffffffffffffffffffff1916145b610685576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526045815260200180610b956045913960600191505060405180910390fd5b61068f84846108b4565b91505092915050565b6000308484846040516020018083838082843780830192505050925050506040516020818303038152906040528051906020012060405160200180807fff000000000000000000000000000000000000000000000000000000000000008152506001018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b815260140183815260200182815260200193505050506040516020818303038152906040528051906020012060001c90506000808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156107b657600090506107b7565b5b9392505050565b600030838360405160200180807fff000000000000000000000000000000000000000000000000000000000000008152506001018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b815260140183815260200182815260200193505050506040516020818303038152906040528051906020012060001c90506000808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156108ad57600090506108ae565b5b92915050565b6000606082905060003085836040516020018082805190602001908083835b602083106108f657805182526020820191506020810190506020830392506108d3565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012060405160200180807fff000000000000000000000000000000000000000000000000000000000000008152506001018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b815260140183815260200182815260200193505050506040516020818303038152906040528051906020012060001c90506000808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610a63576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603f815260200180610b56603f913960400191505060405180910390fd5b81602001825186818334f5945050508073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614610af6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526046815260200180610bda6046913960600191505060405180910390fd5b60016000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050509291505056fe496e76616c696420636f6e7472616374206372656174696f6e202d20636f6e74726163742068617320616c7265616479206265656e206465706c6f7965642e496e76616c69642073616c74202d206669727374203230206279746573206f66207468652073616c74206d757374206d617463682063616c6c696e6720616464726573732e4661696c656420746f206465706c6f7920636f6e7472616374207573696e672070726f76696465642073616c7420616e6420696e697469616c697a6174696f6e20636f64652ea265627a7a72305820df6f8cf8f81946679a805d41e12170a28006ccfb64bce758c74108440cc9337b64736f6c634300050a0032",
}

// ImmutableCreate2FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ImmutableCreate2FactoryMetaData.ABI instead.
var ImmutableCreate2FactoryABI = ImmutableCreate2FactoryMetaData.ABI

// ImmutableCreate2FactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ImmutableCreate2FactoryMetaData.Bin instead.
var ImmutableCreate2FactoryBin = ImmutableCreate2FactoryMetaData.Bin

// DeployImmutableCreate2Factory deploys a new Ethereum contract, binding an instance of ImmutableCreate2Factory to it.
func DeployImmutableCreate2Factory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ImmutableCreate2Factory, error) {
	parsed, err := ImmutableCreate2FactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ImmutableCreate2FactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ImmutableCreate2Factory{ImmutableCreate2FactoryCaller: ImmutableCreate2FactoryCaller{contract: contract}, ImmutableCreate2FactoryTransactor: ImmutableCreate2FactoryTransactor{contract: contract}, ImmutableCreate2FactoryFilterer: ImmutableCreate2FactoryFilterer{contract: contract}}, nil
}

// ImmutableCreate2Factory is an auto generated Go binding around an Ethereum contract.
type ImmutableCreate2Factory struct {
	ImmutableCreate2FactoryCaller     // Read-only binding to the contract
	ImmutableCreate2FactoryTransactor // Write-only binding to the contract
	ImmutableCreate2FactoryFilterer   // Log filterer for contract events
}

// ImmutableCreate2FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ImmutableCreate2FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImmutableCreate2FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ImmutableCreate2FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImmutableCreate2FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ImmutableCreate2FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImmutableCreate2FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ImmutableCreate2FactorySession struct {
	Contract     *ImmutableCreate2Factory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ImmutableCreate2FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ImmutableCreate2FactoryCallerSession struct {
	Contract *ImmutableCreate2FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// ImmutableCreate2FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ImmutableCreate2FactoryTransactorSession struct {
	Contract     *ImmutableCreate2FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ImmutableCreate2FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ImmutableCreate2FactoryRaw struct {
	Contract *ImmutableCreate2Factory // Generic contract binding to access the raw methods on
}

// ImmutableCreate2FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ImmutableCreate2FactoryCallerRaw struct {
	Contract *ImmutableCreate2FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ImmutableCreate2FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ImmutableCreate2FactoryTransactorRaw struct {
	Contract *ImmutableCreate2FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewImmutableCreate2Factory creates a new instance of ImmutableCreate2Factory, bound to a specific deployed contract.
func NewImmutableCreate2Factory(address common.Address, backend bind.ContractBackend) (*ImmutableCreate2Factory, error) {
	contract, err := bindImmutableCreate2Factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ImmutableCreate2Factory{ImmutableCreate2FactoryCaller: ImmutableCreate2FactoryCaller{contract: contract}, ImmutableCreate2FactoryTransactor: ImmutableCreate2FactoryTransactor{contract: contract}, ImmutableCreate2FactoryFilterer: ImmutableCreate2FactoryFilterer{contract: contract}}, nil
}

// NewImmutableCreate2FactoryCaller creates a new read-only instance of ImmutableCreate2Factory, bound to a specific deployed contract.
func NewImmutableCreate2FactoryCaller(address common.Address, caller bind.ContractCaller) (*ImmutableCreate2FactoryCaller, error) {
	contract, err := bindImmutableCreate2Factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ImmutableCreate2FactoryCaller{contract: contract}, nil
}

// NewImmutableCreate2FactoryTransactor creates a new write-only instance of ImmutableCreate2Factory, bound to a specific deployed contract.
func NewImmutableCreate2FactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ImmutableCreate2FactoryTransactor, error) {
	contract, err := bindImmutableCreate2Factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ImmutableCreate2FactoryTransactor{contract: contract}, nil
}

// NewImmutableCreate2FactoryFilterer creates a new log filterer instance of ImmutableCreate2Factory, bound to a specific deployed contract.
func NewImmutableCreate2FactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ImmutableCreate2FactoryFilterer, error) {
	contract, err := bindImmutableCreate2Factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ImmutableCreate2FactoryFilterer{contract: contract}, nil
}

// bindImmutableCreate2Factory binds a generic wrapper to an already deployed contract.
func bindImmutableCreate2Factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ImmutableCreate2FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ImmutableCreate2Factory *ImmutableCreate2FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ImmutableCreate2Factory.Contract.ImmutableCreate2FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ImmutableCreate2Factory *ImmutableCreate2FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ImmutableCreate2Factory.Contract.ImmutableCreate2FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ImmutableCreate2Factory *ImmutableCreate2FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ImmutableCreate2Factory.Contract.ImmutableCreate2FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ImmutableCreate2Factory *ImmutableCreate2FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ImmutableCreate2Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ImmutableCreate2Factory *ImmutableCreate2FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ImmutableCreate2Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ImmutableCreate2Factory *ImmutableCreate2FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ImmutableCreate2Factory.Contract.contract.Transact(opts, method, params...)
}

// FindCreate2Address is a free data retrieval call binding the contract method 0x85cf97ab.
//
// Solidity: function findCreate2Address(bytes32 salt, bytes initCode) view returns(address deploymentAddress)
func (_ImmutableCreate2Factory *ImmutableCreate2FactoryCaller) FindCreate2Address(opts *bind.CallOpts, salt [32]byte, initCode []byte) (common.Address, error) {
	var out []interface{}
	err := _ImmutableCreate2Factory.contract.Call(opts, &out, "findCreate2Address", salt, initCode)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FindCreate2Address is a free data retrieval call binding the contract method 0x85cf97ab.
//
// Solidity: function findCreate2Address(bytes32 salt, bytes initCode) view returns(address deploymentAddress)
func (_ImmutableCreate2Factory *ImmutableCreate2FactorySession) FindCreate2Address(salt [32]byte, initCode []byte) (common.Address, error) {
	return _ImmutableCreate2Factory.Contract.FindCreate2Address(&_ImmutableCreate2Factory.CallOpts, salt, initCode)
}

// FindCreate2Address is a free data retrieval call binding the contract method 0x85cf97ab.
//
// Solidity: function findCreate2Address(bytes32 salt, bytes initCode) view returns(address deploymentAddress)
func (_ImmutableCreate2Factory *ImmutableCreate2FactoryCallerSession) FindCreate2Address(salt [32]byte, initCode []byte) (common.Address, error) {
	return _ImmutableCreate2Factory.Contract.FindCreate2Address(&_ImmutableCreate2Factory.CallOpts, salt, initCode)
}

// FindCreate2AddressViaHash is a free data retrieval call binding the contract method 0xa49a7c90.
//
// Solidity: function findCreate2AddressViaHash(bytes32 salt, bytes32 initCodeHash) view returns(address deploymentAddress)
func (_ImmutableCreate2Factory *ImmutableCreate2FactoryCaller) FindCreate2AddressViaHash(opts *bind.CallOpts, salt [32]byte, initCodeHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ImmutableCreate2Factory.contract.Call(opts, &out, "findCreate2AddressViaHash", salt, initCodeHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FindCreate2AddressViaHash is a free data retrieval call binding the contract method 0xa49a7c90.
//
// Solidity: function findCreate2AddressViaHash(bytes32 salt, bytes32 initCodeHash) view returns(address deploymentAddress)
func (_ImmutableCreate2Factory *ImmutableCreate2FactorySession) FindCreate2AddressViaHash(salt [32]byte, initCodeHash [32]byte) (common.Address, error) {
	return _ImmutableCreate2Factory.Contract.FindCreate2AddressViaHash(&_ImmutableCreate2Factory.CallOpts, salt, initCodeHash)
}

// FindCreate2AddressViaHash is a free data retrieval call binding the contract method 0xa49a7c90.
//
// Solidity: function findCreate2AddressViaHash(bytes32 salt, bytes32 initCodeHash) view returns(address deploymentAddress)
func (_ImmutableCreate2Factory *ImmutableCreate2FactoryCallerSession) FindCreate2AddressViaHash(salt [32]byte, initCodeHash [32]byte) (common.Address, error) {
	return _ImmutableCreate2Factory.Contract.FindCreate2AddressViaHash(&_ImmutableCreate2Factory.CallOpts, salt, initCodeHash)
}

// HasBeenDeployed is a free data retrieval call binding the contract method 0x08508b8f.
//
// Solidity: function hasBeenDeployed(address deploymentAddress) view returns(bool)
func (_ImmutableCreate2Factory *ImmutableCreate2FactoryCaller) HasBeenDeployed(opts *bind.CallOpts, deploymentAddress common.Address) (bool, error) {
	var out []interface{}
	err := _ImmutableCreate2Factory.contract.Call(opts, &out, "hasBeenDeployed", deploymentAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasBeenDeployed is a free data retrieval call binding the contract method 0x08508b8f.
//
// Solidity: function hasBeenDeployed(address deploymentAddress) view returns(bool)
func (_ImmutableCreate2Factory *ImmutableCreate2FactorySession) HasBeenDeployed(deploymentAddress common.Address) (bool, error) {
	return _ImmutableCreate2Factory.Contract.HasBeenDeployed(&_ImmutableCreate2Factory.CallOpts, deploymentAddress)
}

// HasBeenDeployed is a free data retrieval call binding the contract method 0x08508b8f.
//
// Solidity: function hasBeenDeployed(address deploymentAddress) view returns(bool)
func (_ImmutableCreate2Factory *ImmutableCreate2FactoryCallerSession) HasBeenDeployed(deploymentAddress common.Address) (bool, error) {
	return _ImmutableCreate2Factory.Contract.HasBeenDeployed(&_ImmutableCreate2Factory.CallOpts, deploymentAddress)
}

// SafeCreate2 is a paid mutator transaction binding the contract method 0x64e03087.
//
// Solidity: function safeCreate2(bytes32 salt, bytes initializationCode) payable returns(address deploymentAddress)
func (_ImmutableCreate2Factory *ImmutableCreate2FactoryTransactor) SafeCreate2(opts *bind.TransactOpts, salt [32]byte, initializationCode []byte) (*types.Transaction, error) {
	return _ImmutableCreate2Factory.contract.Transact(opts, "safeCreate2", salt, initializationCode)
}

// SafeCreate2 is a paid mutator transaction binding the contract method 0x64e03087.
//
// Solidity: function safeCreate2(bytes32 salt, bytes initializationCode) payable returns(address deploymentAddress)
func (_ImmutableCreate2Factory *ImmutableCreate2FactorySession) SafeCreate2(salt [32]byte, initializationCode []byte) (*types.Transaction, error) {
	return _ImmutableCreate2Factory.Contract.SafeCreate2(&_ImmutableCreate2Factory.TransactOpts, salt, initializationCode)
}

// SafeCreate2 is a paid mutator transaction binding the contract method 0x64e03087.
//
// Solidity: function safeCreate2(bytes32 salt, bytes initializationCode) payable returns(address deploymentAddress)
func (_ImmutableCreate2Factory *ImmutableCreate2FactoryTransactorSession) SafeCreate2(salt [32]byte, initializationCode []byte) (*types.Transaction, error) {
	return _ImmutableCreate2Factory.Contract.SafeCreate2(&_ImmutableCreate2Factory.TransactOpts, salt, initializationCode)
}

// SafeCreate2AndTransfer is a paid mutator transaction binding the contract method 0x29346003.
//
// Solidity: function safeCreate2AndTransfer(bytes32 salt, bytes initializationCode) payable returns(address deploymentAddress)
func (_ImmutableCreate2Factory *ImmutableCreate2FactoryTransactor) SafeCreate2AndTransfer(opts *bind.TransactOpts, salt [32]byte, initializationCode []byte) (*types.Transaction, error) {
	return _ImmutableCreate2Factory.contract.Transact(opts, "safeCreate2AndTransfer", salt, initializationCode)
}

// SafeCreate2AndTransfer is a paid mutator transaction binding the contract method 0x29346003.
//
// Solidity: function safeCreate2AndTransfer(bytes32 salt, bytes initializationCode) payable returns(address deploymentAddress)
func (_ImmutableCreate2Factory *ImmutableCreate2FactorySession) SafeCreate2AndTransfer(salt [32]byte, initializationCode []byte) (*types.Transaction, error) {
	return _ImmutableCreate2Factory.Contract.SafeCreate2AndTransfer(&_ImmutableCreate2Factory.TransactOpts, salt, initializationCode)
}

// SafeCreate2AndTransfer is a paid mutator transaction binding the contract method 0x29346003.
//
// Solidity: function safeCreate2AndTransfer(bytes32 salt, bytes initializationCode) payable returns(address deploymentAddress)
func (_ImmutableCreate2Factory *ImmutableCreate2FactoryTransactorSession) SafeCreate2AndTransfer(salt [32]byte, initializationCode []byte) (*types.Transaction, error) {
	return _ImmutableCreate2Factory.Contract.SafeCreate2AndTransfer(&_ImmutableCreate2Factory.TransactOpts, salt, initializationCode)
}
