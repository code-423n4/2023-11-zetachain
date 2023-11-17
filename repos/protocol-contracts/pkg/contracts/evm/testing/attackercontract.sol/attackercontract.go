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

// AttackerContractMetaData contains all meta data concerning the AttackerContract contract.
var AttackerContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"victimContractAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wzeta\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"victimMethod\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"victimContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620009ae380380620009ae8339818101604052810190620000379190620001a0565b826000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff1663095ea7b360008054906101000a900473ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6040518363ffffffff1660e01b8152600401620000f492919062000250565b602060405180830381600087803b1580156200010f57600080fd5b505af115801562000124573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200014a9190620001fc565b50806001819055505050506200031a565b6000815190506200016c81620002cc565b92915050565b6000815190506200018381620002e6565b92915050565b6000815190506200019a8162000300565b92915050565b600080600060608486031215620001bc57620001bb620002c7565b5b6000620001cc868287016200015b565b9350506020620001df868287016200015b565b9250506040620001f28682870162000189565b9150509250925092565b600060208284031215620002155762000214620002c7565b5b6000620002258482850162000172565b91505092915050565b62000239816200027d565b82525050565b6200024a81620002bd565b82525050565b60006040820190506200026760008301856200022e565b6200027660208301846200023f565b9392505050565b60006200028a826200029d565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600080fd5b620002d7816200027d565b8114620002e357600080fd5b50565b620002f18162000291565b8114620002fd57600080fd5b50565b6200030b81620002bd565b81146200031757600080fd5b50565b610684806200032a6000396000f3fe6080604052600436106100435760003560e01c806323b872dd1461004f57806370a082311461008c57806389bc0bb7146100c9578063a9059cbb146100f45761004a565b3661004a57005b600080fd5b34801561005b57600080fd5b506100766004803603810190610071919061034c565b610131565b60405161008391906104b3565b60405180910390f35b34801561009857600080fd5b506100b360048036038101906100ae919061031f565b610146565b6040516100c0919061051d565b60405180910390f35b3480156100d557600080fd5b506100de610159565b6040516100eb9190610461565b60405180910390f35b34801561010057600080fd5b5061011b6004803603810190610116919061039f565b61017d565b60405161012891906104b3565b60405180910390f35b600061013b610191565b600190509392505050565b6000610150610191565b60009050919050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000610187610191565b6001905092915050565b6001805414156101a8576101a36101bf565b6101bd565b600260015414156101bc576101bb61024f565b5b5b565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e609055e3060006040518363ffffffff1660e01b815260040161021b9291906104ce565b600060405180830381600087803b15801561023557600080fd5b505af1158015610249573d6000803e3d6000fd5b50505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d9caed1273f39fd6e51aad88f6f4ce6ab8827279cfffb922663060006040518463ffffffff1660e01b81526004016102c19392919061047c565b600060405180830381600087803b1580156102db57600080fd5b505af11580156102ef573d6000803e3d6000fd5b50505050565b60008135905061030481610620565b92915050565b60008135905061031981610637565b92915050565b600060208284031215610335576103346105a3565b5b6000610343848285016102f5565b91505092915050565b600080600060608486031215610365576103646105a3565b5b6000610373868287016102f5565b9350506020610384868287016102f5565b92505060406103958682870161030a565b9150509250925092565b600080604083850312156103b6576103b56105a3565b5b60006103c4858286016102f5565b92505060206103d58582860161030a565b9150509250929050565b6103e881610549565b82525050565b6103f78161055b565b82525050565b61040681610591565b82525050565b6000610419600483610538565b9150610424826105a8565b602082019050919050565b600061043c602a83610538565b9150610447826105d1565b604082019050919050565b61045b81610587565b82525050565b600060208201905061047660008301846103df565b92915050565b600060608201905061049160008301866103df565b61049e60208301856103df565b6104ab60408301846103fd565b949350505050565b60006020820190506104c860008301846103ee565b92915050565b600060808201905081810360008301526104e78161042f565b90506104f660208301856103df565b61050360408301846103fd565b81810360608301526105148161040c565b90509392505050565b60006020820190506105326000830184610452565b92915050565b600082825260208201905092915050565b600061055482610567565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061059c82610587565b9050919050565b600080fd5b7f3078303000000000000000000000000000000000000000000000000000000000600082015250565b7f307866333946643665353161616438384636463463653661423838323732373960008201527f6366664662393232363600000000000000000000000000000000000000000000602082015250565b61062981610549565b811461063457600080fd5b50565b61064081610587565b811461064b57600080fd5b5056fea264697066735822122075954d652b0eb8d814b91524c47a04c537d638cec0eda3ca0e4ee67767e1a37064736f6c63430008070033",
}

// AttackerContractABI is the input ABI used to generate the binding from.
// Deprecated: Use AttackerContractMetaData.ABI instead.
var AttackerContractABI = AttackerContractMetaData.ABI

// AttackerContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AttackerContractMetaData.Bin instead.
var AttackerContractBin = AttackerContractMetaData.Bin

// DeployAttackerContract deploys a new Ethereum contract, binding an instance of AttackerContract to it.
func DeployAttackerContract(auth *bind.TransactOpts, backend bind.ContractBackend, victimContractAddress_ common.Address, wzeta common.Address, victimMethod *big.Int) (common.Address, *types.Transaction, *AttackerContract, error) {
	parsed, err := AttackerContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AttackerContractBin), backend, victimContractAddress_, wzeta, victimMethod)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AttackerContract{AttackerContractCaller: AttackerContractCaller{contract: contract}, AttackerContractTransactor: AttackerContractTransactor{contract: contract}, AttackerContractFilterer: AttackerContractFilterer{contract: contract}}, nil
}

// AttackerContract is an auto generated Go binding around an Ethereum contract.
type AttackerContract struct {
	AttackerContractCaller     // Read-only binding to the contract
	AttackerContractTransactor // Write-only binding to the contract
	AttackerContractFilterer   // Log filterer for contract events
}

// AttackerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type AttackerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttackerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AttackerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttackerContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AttackerContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttackerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AttackerContractSession struct {
	Contract     *AttackerContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AttackerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AttackerContractCallerSession struct {
	Contract *AttackerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// AttackerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AttackerContractTransactorSession struct {
	Contract     *AttackerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// AttackerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type AttackerContractRaw struct {
	Contract *AttackerContract // Generic contract binding to access the raw methods on
}

// AttackerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AttackerContractCallerRaw struct {
	Contract *AttackerContractCaller // Generic read-only contract binding to access the raw methods on
}

// AttackerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AttackerContractTransactorRaw struct {
	Contract *AttackerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAttackerContract creates a new instance of AttackerContract, bound to a specific deployed contract.
func NewAttackerContract(address common.Address, backend bind.ContractBackend) (*AttackerContract, error) {
	contract, err := bindAttackerContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AttackerContract{AttackerContractCaller: AttackerContractCaller{contract: contract}, AttackerContractTransactor: AttackerContractTransactor{contract: contract}, AttackerContractFilterer: AttackerContractFilterer{contract: contract}}, nil
}

// NewAttackerContractCaller creates a new read-only instance of AttackerContract, bound to a specific deployed contract.
func NewAttackerContractCaller(address common.Address, caller bind.ContractCaller) (*AttackerContractCaller, error) {
	contract, err := bindAttackerContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AttackerContractCaller{contract: contract}, nil
}

// NewAttackerContractTransactor creates a new write-only instance of AttackerContract, bound to a specific deployed contract.
func NewAttackerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*AttackerContractTransactor, error) {
	contract, err := bindAttackerContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AttackerContractTransactor{contract: contract}, nil
}

// NewAttackerContractFilterer creates a new log filterer instance of AttackerContract, bound to a specific deployed contract.
func NewAttackerContractFilterer(address common.Address, filterer bind.ContractFilterer) (*AttackerContractFilterer, error) {
	contract, err := bindAttackerContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AttackerContractFilterer{contract: contract}, nil
}

// bindAttackerContract binds a generic wrapper to an already deployed contract.
func bindAttackerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AttackerContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AttackerContract *AttackerContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AttackerContract.Contract.AttackerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AttackerContract *AttackerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AttackerContract.Contract.AttackerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AttackerContract *AttackerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AttackerContract.Contract.AttackerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AttackerContract *AttackerContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AttackerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AttackerContract *AttackerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AttackerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AttackerContract *AttackerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AttackerContract.Contract.contract.Transact(opts, method, params...)
}

// VictimContractAddress is a free data retrieval call binding the contract method 0x89bc0bb7.
//
// Solidity: function victimContractAddress() view returns(address)
func (_AttackerContract *AttackerContractCaller) VictimContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AttackerContract.contract.Call(opts, &out, "victimContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VictimContractAddress is a free data retrieval call binding the contract method 0x89bc0bb7.
//
// Solidity: function victimContractAddress() view returns(address)
func (_AttackerContract *AttackerContractSession) VictimContractAddress() (common.Address, error) {
	return _AttackerContract.Contract.VictimContractAddress(&_AttackerContract.CallOpts)
}

// VictimContractAddress is a free data retrieval call binding the contract method 0x89bc0bb7.
//
// Solidity: function victimContractAddress() view returns(address)
func (_AttackerContract *AttackerContractCallerSession) VictimContractAddress() (common.Address, error) {
	return _AttackerContract.Contract.VictimContractAddress(&_AttackerContract.CallOpts)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) returns(uint256)
func (_AttackerContract *AttackerContractTransactor) BalanceOf(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _AttackerContract.contract.Transact(opts, "balanceOf", account)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) returns(uint256)
func (_AttackerContract *AttackerContractSession) BalanceOf(account common.Address) (*types.Transaction, error) {
	return _AttackerContract.Contract.BalanceOf(&_AttackerContract.TransactOpts, account)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) returns(uint256)
func (_AttackerContract *AttackerContractTransactorSession) BalanceOf(account common.Address) (*types.Transaction, error) {
	return _AttackerContract.Contract.BalanceOf(&_AttackerContract.TransactOpts, account)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_AttackerContract *AttackerContractTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AttackerContract.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_AttackerContract *AttackerContractSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AttackerContract.Contract.Transfer(&_AttackerContract.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_AttackerContract *AttackerContractTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AttackerContract.Contract.Transfer(&_AttackerContract.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_AttackerContract *AttackerContractTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AttackerContract.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_AttackerContract *AttackerContractSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AttackerContract.Contract.TransferFrom(&_AttackerContract.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_AttackerContract *AttackerContractTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AttackerContract.Contract.TransferFrom(&_AttackerContract.TransactOpts, from, to, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AttackerContract *AttackerContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AttackerContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AttackerContract *AttackerContractSession) Receive() (*types.Transaction, error) {
	return _AttackerContract.Contract.Receive(&_AttackerContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AttackerContract *AttackerContractTransactorSession) Receive() (*types.Transaction, error) {
	return _AttackerContract.Contract.Receive(&_AttackerContract.TransactOpts)
}
