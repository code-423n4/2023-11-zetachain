// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zeta

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

// ZetaEthMetaData contains all meta data concerning the ZetaEth contract.
var ZetaEthMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001a5238038062001a5283398181016040528101906200003791906200037d565b6040518060400160405280600481526020017f5a657461000000000000000000000000000000000000000000000000000000008152506040518060400160405280600481526020017f5a455441000000000000000000000000000000000000000000000000000000008152508160039080519060200190620000bb9291906200029f565b508060049080519060200190620000d49291906200029f565b5050506200011682620000ec6200011e60201b60201c565b60ff16600a620000fd919062000504565b836200010a919062000641565b6200012760201b60201c565b5050620007e3565b60006012905090565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156200019a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200019190620003fc565b60405180910390fd5b620001ae600083836200029560201b60201c565b8060026000828254620001c291906200044c565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516200027591906200041e565b60405180910390a362000291600083836200029a60201b60201c565b5050565b505050565b505050565b828054620002ad90620006e0565b90600052602060002090601f016020900481019282620002d157600085556200031d565b82601f10620002ec57805160ff19168380011785556200031d565b828001600101855582156200031d579182015b828111156200031c578251825591602001919060010190620002ff565b5b5090506200032c919062000330565b5090565b5b808211156200034b57600081600090555060010162000331565b5090565b6000815190506200036081620007af565b92915050565b6000815190506200037781620007c9565b92915050565b6000806040838503121562000397576200039662000774565b5b6000620003a7858286016200034f565b9250506020620003ba8582860162000366565b9150509250929050565b6000620003d3601f836200043b565b9150620003e08262000786565b602082019050919050565b620003f681620006d6565b82525050565b600060208201905081810360008301526200041781620003c4565b9050919050565b6000602082019050620004356000830184620003eb565b92915050565b600082825260208201905092915050565b60006200045982620006d6565b91506200046683620006d6565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156200049e576200049d62000716565b5b828201905092915050565b6000808291508390505b6001851115620004fb57808604811115620004d357620004d262000716565b5b6001851615620004e35780820291505b8081029050620004f38562000779565b9450620004b3565b94509492505050565b60006200051182620006d6565b91506200051e83620006d6565b92506200054d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff848462000555565b905092915050565b6000826200056757600190506200063a565b816200057757600090506200063a565b81600181146200059057600281146200059b57620005d1565b60019150506200063a565b60ff841115620005b057620005af62000716565b5b8360020a915084821115620005ca57620005c962000716565b5b506200063a565b5060208310610133831016604e8410600b84101617156200060b5782820a90508381111562000605576200060462000716565b5b6200063a565b6200061a8484846001620004a9565b9250905081840481111562000634576200063362000716565b5b81810290505b9392505050565b60006200064e82620006d6565b91506200065b83620006d6565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161562000697576200069662000716565b5b828202905092915050565b6000620006af82620006b6565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006002820490506001821680620006f957607f821691505b6020821081141562000710576200070f62000745565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600080fd5b60008160011c9050919050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b620007ba81620006a2565b8114620007c657600080fd5b50565b620007d481620006d6565b8114620007e057600080fd5b50565b61125f80620007f36000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80633950935111610071578063395093511461016857806370a082311461019857806395d89b41146101c8578063a457c2d7146101e6578063a9059cbb14610216578063dd62ed3e14610246576100a9565b806306fdde03146100ae578063095ea7b3146100cc57806318160ddd146100fc57806323b872dd1461011a578063313ce5671461014a575b600080fd5b6100b6610276565b6040516100c39190610d20565b60405180910390f35b6100e660048036038101906100e19190610b6a565b610308565b6040516100f39190610d05565b60405180910390f35b61010461032b565b6040516101119190610e22565b60405180910390f35b610134600480360381019061012f9190610b17565b610335565b6040516101419190610d05565b60405180910390f35b610152610364565b60405161015f9190610e3d565b60405180910390f35b610182600480360381019061017d9190610b6a565b61036d565b60405161018f9190610d05565b60405180910390f35b6101b260048036038101906101ad9190610aaa565b6103a4565b6040516101bf9190610e22565b60405180910390f35b6101d06103ec565b6040516101dd9190610d20565b60405180910390f35b61020060048036038101906101fb9190610b6a565b61047e565b60405161020d9190610d05565b60405180910390f35b610230600480360381019061022b9190610b6a565b6104f5565b60405161023d9190610d05565b60405180910390f35b610260600480360381019061025b9190610ad7565b610518565b60405161026d9190610e22565b60405180910390f35b60606003805461028590610f52565b80601f01602080910402602001604051908101604052809291908181526020018280546102b190610f52565b80156102fe5780601f106102d3576101008083540402835291602001916102fe565b820191906000526020600020905b8154815290600101906020018083116102e157829003601f168201915b5050505050905090565b60008061031361059f565b90506103208185856105a7565b600191505092915050565b6000600254905090565b60008061034061059f565b905061034d858285610772565b6103588585856107fe565b60019150509392505050565b60006012905090565b60008061037861059f565b905061039981858561038a8589610518565b6103949190610e74565b6105a7565b600191505092915050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6060600480546103fb90610f52565b80601f016020809104026020016040519081016040528092919081815260200182805461042790610f52565b80156104745780601f1061044957610100808354040283529160200191610474565b820191906000526020600020905b81548152906001019060200180831161045757829003601f168201915b5050505050905090565b60008061048961059f565b905060006104978286610518565b9050838110156104dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104d390610e02565b60405180910390fd5b6104e982868684036105a7565b60019250505092915050565b60008061050061059f565b905061050d8185856107fe565b600191505092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610617576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060e90610de2565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610687576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161067e90610d62565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040516107659190610e22565b60405180910390a3505050565b600061077e8484610518565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146107f857818110156107ea576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107e190610d82565b60405180910390fd5b6107f784848484036105a7565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561086e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161086590610dc2565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156108de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d590610d42565b60405180910390fd5b6108e9838383610a76565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508181101561096f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161096690610da2565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610a5d9190610e22565b60405180910390a3610a70848484610a7b565b50505050565b505050565b505050565b600081359050610a8f816111fb565b92915050565b600081359050610aa481611212565b92915050565b600060208284031215610ac057610abf610fe2565b5b6000610ace84828501610a80565b91505092915050565b60008060408385031215610aee57610aed610fe2565b5b6000610afc85828601610a80565b9250506020610b0d85828601610a80565b9150509250929050565b600080600060608486031215610b3057610b2f610fe2565b5b6000610b3e86828701610a80565b9350506020610b4f86828701610a80565b9250506040610b6086828701610a95565b9150509250925092565b60008060408385031215610b8157610b80610fe2565b5b6000610b8f85828601610a80565b9250506020610ba085828601610a95565b9150509250929050565b610bb381610edc565b82525050565b6000610bc482610e58565b610bce8185610e63565b9350610bde818560208601610f1f565b610be781610fe7565b840191505092915050565b6000610bff602383610e63565b9150610c0a82610ff8565b604082019050919050565b6000610c22602283610e63565b9150610c2d82611047565b604082019050919050565b6000610c45601d83610e63565b9150610c5082611096565b602082019050919050565b6000610c68602683610e63565b9150610c73826110bf565b604082019050919050565b6000610c8b602583610e63565b9150610c968261110e565b604082019050919050565b6000610cae602483610e63565b9150610cb98261115d565b604082019050919050565b6000610cd1602583610e63565b9150610cdc826111ac565b604082019050919050565b610cf081610f08565b82525050565b610cff81610f12565b82525050565b6000602082019050610d1a6000830184610baa565b92915050565b60006020820190508181036000830152610d3a8184610bb9565b905092915050565b60006020820190508181036000830152610d5b81610bf2565b9050919050565b60006020820190508181036000830152610d7b81610c15565b9050919050565b60006020820190508181036000830152610d9b81610c38565b9050919050565b60006020820190508181036000830152610dbb81610c5b565b9050919050565b60006020820190508181036000830152610ddb81610c7e565b9050919050565b60006020820190508181036000830152610dfb81610ca1565b9050919050565b60006020820190508181036000830152610e1b81610cc4565b9050919050565b6000602082019050610e376000830184610ce7565b92915050565b6000602082019050610e526000830184610cf6565b92915050565b600081519050919050565b600082825260208201905092915050565b6000610e7f82610f08565b9150610e8a83610f08565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115610ebf57610ebe610f84565b5b828201905092915050565b6000610ed582610ee8565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b83811015610f3d578082015181840152602081019050610f22565b83811115610f4c576000848401525b50505050565b60006002820490506001821680610f6a57607f821691505b60208210811415610f7e57610f7d610fb3565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600080fd5b6000601f19601f8301169050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000600082015250565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b61120481610eca565b811461120f57600080fd5b50565b61121b81610f08565b811461122657600080fd5b5056fea2646970667358221220e33277034a5236435f4dc6a93d4c4dc71fb8a6be9f4a752ea3f374446caf920b64736f6c63430008070033",
}

// ZetaEthABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaEthMetaData.ABI instead.
var ZetaEthABI = ZetaEthMetaData.ABI

// ZetaEthBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZetaEthMetaData.Bin instead.
var ZetaEthBin = ZetaEthMetaData.Bin

// DeployZetaEth deploys a new Ethereum contract, binding an instance of ZetaEth to it.
func DeployZetaEth(auth *bind.TransactOpts, backend bind.ContractBackend, creator common.Address, initialSupply *big.Int) (common.Address, *types.Transaction, *ZetaEth, error) {
	parsed, err := ZetaEthMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZetaEthBin), backend, creator, initialSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZetaEth{ZetaEthCaller: ZetaEthCaller{contract: contract}, ZetaEthTransactor: ZetaEthTransactor{contract: contract}, ZetaEthFilterer: ZetaEthFilterer{contract: contract}}, nil
}

// ZetaEth is an auto generated Go binding around an Ethereum contract.
type ZetaEth struct {
	ZetaEthCaller     // Read-only binding to the contract
	ZetaEthTransactor // Write-only binding to the contract
	ZetaEthFilterer   // Log filterer for contract events
}

// ZetaEthCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaEthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaEthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaEthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaEthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaEthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaEthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaEthSession struct {
	Contract     *ZetaEth          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZetaEthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaEthCallerSession struct {
	Contract *ZetaEthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ZetaEthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaEthTransactorSession struct {
	Contract     *ZetaEthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ZetaEthRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaEthRaw struct {
	Contract *ZetaEth // Generic contract binding to access the raw methods on
}

// ZetaEthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaEthCallerRaw struct {
	Contract *ZetaEthCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaEthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaEthTransactorRaw struct {
	Contract *ZetaEthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaEth creates a new instance of ZetaEth, bound to a specific deployed contract.
func NewZetaEth(address common.Address, backend bind.ContractBackend) (*ZetaEth, error) {
	contract, err := bindZetaEth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaEth{ZetaEthCaller: ZetaEthCaller{contract: contract}, ZetaEthTransactor: ZetaEthTransactor{contract: contract}, ZetaEthFilterer: ZetaEthFilterer{contract: contract}}, nil
}

// NewZetaEthCaller creates a new read-only instance of ZetaEth, bound to a specific deployed contract.
func NewZetaEthCaller(address common.Address, caller bind.ContractCaller) (*ZetaEthCaller, error) {
	contract, err := bindZetaEth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaEthCaller{contract: contract}, nil
}

// NewZetaEthTransactor creates a new write-only instance of ZetaEth, bound to a specific deployed contract.
func NewZetaEthTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaEthTransactor, error) {
	contract, err := bindZetaEth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaEthTransactor{contract: contract}, nil
}

// NewZetaEthFilterer creates a new log filterer instance of ZetaEth, bound to a specific deployed contract.
func NewZetaEthFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaEthFilterer, error) {
	contract, err := bindZetaEth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaEthFilterer{contract: contract}, nil
}

// bindZetaEth binds a generic wrapper to an already deployed contract.
func bindZetaEth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaEthMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaEth *ZetaEthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaEth.Contract.ZetaEthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaEth *ZetaEthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaEth.Contract.ZetaEthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaEth *ZetaEthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaEth.Contract.ZetaEthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaEth *ZetaEthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaEth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaEth *ZetaEthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaEth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaEth *ZetaEthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaEth.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZetaEth *ZetaEthCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZetaEth.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZetaEth *ZetaEthSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ZetaEth.Contract.Allowance(&_ZetaEth.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZetaEth *ZetaEthCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ZetaEth.Contract.Allowance(&_ZetaEth.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZetaEth *ZetaEthCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZetaEth.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZetaEth *ZetaEthSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ZetaEth.Contract.BalanceOf(&_ZetaEth.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZetaEth *ZetaEthCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ZetaEth.Contract.BalanceOf(&_ZetaEth.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZetaEth *ZetaEthCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZetaEth.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZetaEth *ZetaEthSession) Decimals() (uint8, error) {
	return _ZetaEth.Contract.Decimals(&_ZetaEth.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZetaEth *ZetaEthCallerSession) Decimals() (uint8, error) {
	return _ZetaEth.Contract.Decimals(&_ZetaEth.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZetaEth *ZetaEthCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZetaEth.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZetaEth *ZetaEthSession) Name() (string, error) {
	return _ZetaEth.Contract.Name(&_ZetaEth.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZetaEth *ZetaEthCallerSession) Name() (string, error) {
	return _ZetaEth.Contract.Name(&_ZetaEth.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZetaEth *ZetaEthCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZetaEth.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZetaEth *ZetaEthSession) Symbol() (string, error) {
	return _ZetaEth.Contract.Symbol(&_ZetaEth.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZetaEth *ZetaEthCallerSession) Symbol() (string, error) {
	return _ZetaEth.Contract.Symbol(&_ZetaEth.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZetaEth *ZetaEthCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZetaEth.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZetaEth *ZetaEthSession) TotalSupply() (*big.Int, error) {
	return _ZetaEth.Contract.TotalSupply(&_ZetaEth.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZetaEth *ZetaEthCallerSession) TotalSupply() (*big.Int, error) {
	return _ZetaEth.Contract.TotalSupply(&_ZetaEth.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ZetaEth *ZetaEthTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZetaEth.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ZetaEth *ZetaEthSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZetaEth.Contract.Approve(&_ZetaEth.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ZetaEth *ZetaEthTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZetaEth.Contract.Approve(&_ZetaEth.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ZetaEth *ZetaEthTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ZetaEth.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ZetaEth *ZetaEthSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ZetaEth.Contract.DecreaseAllowance(&_ZetaEth.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ZetaEth *ZetaEthTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ZetaEth.Contract.DecreaseAllowance(&_ZetaEth.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ZetaEth *ZetaEthTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ZetaEth.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ZetaEth *ZetaEthSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ZetaEth.Contract.IncreaseAllowance(&_ZetaEth.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ZetaEth *ZetaEthTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ZetaEth.Contract.IncreaseAllowance(&_ZetaEth.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ZetaEth *ZetaEthTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZetaEth.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ZetaEth *ZetaEthSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZetaEth.Contract.Transfer(&_ZetaEth.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ZetaEth *ZetaEthTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZetaEth.Contract.Transfer(&_ZetaEth.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ZetaEth *ZetaEthTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZetaEth.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ZetaEth *ZetaEthSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZetaEth.Contract.TransferFrom(&_ZetaEth.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ZetaEth *ZetaEthTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZetaEth.Contract.TransferFrom(&_ZetaEth.TransactOpts, from, to, amount)
}

// ZetaEthApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ZetaEth contract.
type ZetaEthApprovalIterator struct {
	Event *ZetaEthApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaEthApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaEthApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaEthApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaEthApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaEthApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaEthApproval represents a Approval event raised by the ZetaEth contract.
type ZetaEthApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZetaEth *ZetaEthFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ZetaEthApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ZetaEth.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ZetaEthApprovalIterator{contract: _ZetaEth.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZetaEth *ZetaEthFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ZetaEthApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ZetaEth.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaEthApproval)
				if err := _ZetaEth.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZetaEth *ZetaEthFilterer) ParseApproval(log types.Log) (*ZetaEthApproval, error) {
	event := new(ZetaEthApproval)
	if err := _ZetaEth.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaEthTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ZetaEth contract.
type ZetaEthTransferIterator struct {
	Event *ZetaEthTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaEthTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaEthTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaEthTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaEthTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaEthTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaEthTransfer represents a Transfer event raised by the ZetaEth contract.
type ZetaEthTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZetaEth *ZetaEthFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ZetaEthTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaEth.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaEthTransferIterator{contract: _ZetaEth.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZetaEth *ZetaEthFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ZetaEthTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaEth.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaEthTransfer)
				if err := _ZetaEth.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZetaEth *ZetaEthFilterer) ParseTransfer(log types.Log) (*ZetaEthTransfer, error) {
	event := new(ZetaEthTransfer)
	if err := _ZetaEth.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
