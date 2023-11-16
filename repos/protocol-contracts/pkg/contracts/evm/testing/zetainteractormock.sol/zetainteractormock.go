// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetainteractormock

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

// ZetaInterfacesZetaMessage is an auto generated low-level Go binding around an user-defined struct.
type ZetaInterfacesZetaMessage struct {
	ZetaTxSenderAddress []byte
	SourceChainId       *big.Int
	DestinationAddress  common.Address
	ZetaValue           *big.Int
	Message             []byte
}

// ZetaInterfacesZetaRevert is an auto generated low-level Go binding around an user-defined struct.
type ZetaInterfacesZetaRevert struct {
	ZetaTxSenderAddress common.Address
	SourceChainId       *big.Int
	DestinationAddress  []byte
	DestinationChainId  *big.Int
	RemainingZetaValue  *big.Int
	Message             []byte
}

// ZetaInteractorMockMetaData contains all meta data concerning the ZetaInteractorMock contract.
var ZetaInteractorMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"zetaConnectorAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"InvalidCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDestinationChainId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZetaMessageCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZetaRevertCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"connector\",\"outputs\":[{\"internalType\":\"contractZetaConnector\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"interactorsByChainId\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"zetaTxSenderAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"zetaValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structZetaInterfaces.ZetaMessage\",\"name\":\"zetaMessage\",\"type\":\"tuple\"}],\"name\":\"onZetaMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"zetaTxSenderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"destinationAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingZetaValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structZetaInterfaces.ZetaRevert\",\"name\":\"zetaRevert\",\"type\":\"tuple\"}],\"name\":\"onZetaRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"contractAddress\",\"type\":\"bytes\"}],\"name\":\"setInteractorByChainId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b506040516200123c3803806200123c833981810160405281019062000037919062000228565b80620000586200004c6200010760201b60201c565b6200010f60201b60201c565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415620000c0576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b46608081815250508073ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1660601b815250505050620002ad565b600033905090565b600160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556200014a816200014d60201b620005bd1760201c565b50565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600081519050620002228162000293565b92915050565b6000602082840312156200024157620002406200028e565b5b6000620002518482850162000211565b91505092915050565b600062000267826200026e565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6200029e816200025a565b8114620002aa57600080fd5b50565b60805160a05160601c610f5f620002dd6000396000818161049b0152610683015260006103690152610f5f6000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c806379ba50971161006657806379ba50971461013157806383f3084f1461013b5780638da5cb5b14610159578063e30c397814610177578063f2fde38b146101955761009e565b80632618143f146100a35780633749c51a146100d35780633ff0693c146100ef5780634fd3f7d71461010b578063715018a614610127575b600080fd5b6100bd60048036038101906100b891906109ea565b6101b1565b6040516100ca9190610c03565b60405180910390f35b6100ed60048036038101906100e89190610958565b610251565b005b610109600480360381019061010491906109a1565b6102e7565b005b61012560048036038101906101209190610a17565b6103c8565b005b61012f6103f8565b005b61013961040c565b005b610143610499565b6040516101509190610c25565b60405180910390f35b6101616104bd565b60405161016e9190610be8565b60405180910390f35b61017f6104e6565b60405161018c9190610be8565b60405180910390f35b6101af60048036038101906101aa919061092b565b610510565b005b600260205280600052604060002060009150905080546101d090610de4565b80601f01602080910402602001604051908101604052809291908181526020018280546101fc90610de4565b80156102495780601f1061021e57610100808354040283529160200191610249565b820191906000526020600020905b81548152906001019060200180831161022c57829003601f168201915b505050505081565b8061025a610681565b600260008260200135815260200190815260200160002060405161027e9190610bd1565b60405180910390208180600001906102969190610c80565b6040516102a4929190610bb8565b6040518091039020146102e3576040517fb473eb0a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b806102f0610681565b3073ffffffffffffffffffffffffffffffffffffffff1681600001602081019061031a919061092b565b73ffffffffffffffffffffffffffffffffffffffff1614610367576040517fc03e9c3f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000008160200135146103c4576040517fc03e9c3f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b6103d0610713565b81816002600086815260200190815260200160002091906103f29291906107ca565b50505050565b610400610713565b61040a6000610791565b565b60006104166107c2565b90508073ffffffffffffffffffffffffffffffffffffffff166104376104e6565b73ffffffffffffffffffffffffffffffffffffffff161461048d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161048490610c40565b60405180910390fd5b61049681610791565b50565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610518610713565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff166105786104bd565b73ffffffffffffffffffffffffffffffffffffffff167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461071157336040517fcbd9d2e00000000000000000000000000000000000000000000000000000000081526004016107089190610be8565b60405180910390fd5b565b61071b6107c2565b73ffffffffffffffffffffffffffffffffffffffff166107396104bd565b73ffffffffffffffffffffffffffffffffffffffff161461078f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161078690610c60565b60405180910390fd5b565b600160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556107bf816105bd565b50565b600033905090565b8280546107d690610de4565b90600052602060002090601f0160209004810192826107f8576000855561083f565b82601f1061081157803560ff191683800117855561083f565b8280016001018555821561083f579182015b8281111561083e578235825591602001919060010190610823565b5b50905061084c9190610850565b5090565b5b80821115610869576000816000905550600101610851565b5090565b60008135905061087c81610efb565b92915050565b60008083601f84011261089857610897610e4a565b5b8235905067ffffffffffffffff8111156108b5576108b4610e45565b5b6020830191508360018202830111156108d1576108d0610e5e565b5b9250929050565b600060a082840312156108ee576108ed610e54565b5b81905092915050565b600060c0828403121561090d5761090c610e54565b5b81905092915050565b60008135905061092581610f12565b92915050565b60006020828403121561094157610940610e6d565b5b600061094f8482850161086d565b91505092915050565b60006020828403121561096e5761096d610e6d565b5b600082013567ffffffffffffffff81111561098c5761098b610e68565b5b610998848285016108d8565b91505092915050565b6000602082840312156109b7576109b6610e6d565b5b600082013567ffffffffffffffff8111156109d5576109d4610e68565b5b6109e1848285016108f7565b91505092915050565b600060208284031215610a00576109ff610e6d565b5b6000610a0e84828501610916565b91505092915050565b600080600060408486031215610a3057610a2f610e6d565b5b6000610a3e86828701610916565b935050602084013567ffffffffffffffff811115610a5f57610a5e610e68565b5b610a6b86828701610882565b92509250509250925092565b610a8081610d30565b82525050565b6000610a928385610d14565b9350610a9f838584610da2565b82840190509392505050565b6000610ab682610cf8565b610ac08185610d03565b9350610ad0818560208601610db1565b610ad981610e72565b840191505092915050565b60008154610af181610de4565b610afb8186610d14565b94506001821660008114610b165760018114610b2757610b5a565b60ff19831686528186019350610b5a565b610b3085610ce3565b60005b83811015610b5257815481890152600182019150602081019050610b33565b838801955050505b50505092915050565b610b6c81610d6c565b82525050565b6000610b7f602983610d1f565b9150610b8a82610e83565b604082019050919050565b6000610ba2602083610d1f565b9150610bad82610ed2565b602082019050919050565b6000610bc5828486610a86565b91508190509392505050565b6000610bdd8284610ae4565b915081905092915050565b6000602082019050610bfd6000830184610a77565b92915050565b60006020820190508181036000830152610c1d8184610aab565b905092915050565b6000602082019050610c3a6000830184610b63565b92915050565b60006020820190508181036000830152610c5981610b72565b9050919050565b60006020820190508181036000830152610c7981610b95565b9050919050565b60008083356001602003843603038112610c9d57610c9c610e59565b5b80840192508235915067ffffffffffffffff821115610cbf57610cbe610e4f565b5b602083019250600182023603831315610cdb57610cda610e63565b5b509250929050565b60008190508160005260206000209050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000610d3b82610d42565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000610d7782610d7e565b9050919050565b6000610d8982610d90565b9050919050565b6000610d9b82610d42565b9050919050565b82818337600083830152505050565b60005b83811015610dcf578082015181840152602081019050610db4565b83811115610dde576000848401525b50505050565b60006002820490506001821680610dfc57607f821691505b60208210811415610e1057610e0f610e16565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4f776e61626c6532537465703a2063616c6c6572206973206e6f74207468652060008201527f6e6577206f776e65720000000000000000000000000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b610f0481610d30565b8114610f0f57600080fd5b50565b610f1b81610d62565b8114610f2657600080fd5b5056fea26469706673582212206999e36457fead1d5e7b5a125ef2e0a49dcff212174326676d6fd53cff4e990364736f6c63430008070033",
}

// ZetaInteractorMockABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaInteractorMockMetaData.ABI instead.
var ZetaInteractorMockABI = ZetaInteractorMockMetaData.ABI

// ZetaInteractorMockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZetaInteractorMockMetaData.Bin instead.
var ZetaInteractorMockBin = ZetaInteractorMockMetaData.Bin

// DeployZetaInteractorMock deploys a new Ethereum contract, binding an instance of ZetaInteractorMock to it.
func DeployZetaInteractorMock(auth *bind.TransactOpts, backend bind.ContractBackend, zetaConnectorAddress common.Address) (common.Address, *types.Transaction, *ZetaInteractorMock, error) {
	parsed, err := ZetaInteractorMockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZetaInteractorMockBin), backend, zetaConnectorAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZetaInteractorMock{ZetaInteractorMockCaller: ZetaInteractorMockCaller{contract: contract}, ZetaInteractorMockTransactor: ZetaInteractorMockTransactor{contract: contract}, ZetaInteractorMockFilterer: ZetaInteractorMockFilterer{contract: contract}}, nil
}

// ZetaInteractorMock is an auto generated Go binding around an Ethereum contract.
type ZetaInteractorMock struct {
	ZetaInteractorMockCaller     // Read-only binding to the contract
	ZetaInteractorMockTransactor // Write-only binding to the contract
	ZetaInteractorMockFilterer   // Log filterer for contract events
}

// ZetaInteractorMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaInteractorMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaInteractorMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaInteractorMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaInteractorMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaInteractorMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaInteractorMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaInteractorMockSession struct {
	Contract     *ZetaInteractorMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ZetaInteractorMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaInteractorMockCallerSession struct {
	Contract *ZetaInteractorMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ZetaInteractorMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaInteractorMockTransactorSession struct {
	Contract     *ZetaInteractorMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ZetaInteractorMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaInteractorMockRaw struct {
	Contract *ZetaInteractorMock // Generic contract binding to access the raw methods on
}

// ZetaInteractorMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaInteractorMockCallerRaw struct {
	Contract *ZetaInteractorMockCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaInteractorMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaInteractorMockTransactorRaw struct {
	Contract *ZetaInteractorMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaInteractorMock creates a new instance of ZetaInteractorMock, bound to a specific deployed contract.
func NewZetaInteractorMock(address common.Address, backend bind.ContractBackend) (*ZetaInteractorMock, error) {
	contract, err := bindZetaInteractorMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaInteractorMock{ZetaInteractorMockCaller: ZetaInteractorMockCaller{contract: contract}, ZetaInteractorMockTransactor: ZetaInteractorMockTransactor{contract: contract}, ZetaInteractorMockFilterer: ZetaInteractorMockFilterer{contract: contract}}, nil
}

// NewZetaInteractorMockCaller creates a new read-only instance of ZetaInteractorMock, bound to a specific deployed contract.
func NewZetaInteractorMockCaller(address common.Address, caller bind.ContractCaller) (*ZetaInteractorMockCaller, error) {
	contract, err := bindZetaInteractorMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaInteractorMockCaller{contract: contract}, nil
}

// NewZetaInteractorMockTransactor creates a new write-only instance of ZetaInteractorMock, bound to a specific deployed contract.
func NewZetaInteractorMockTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaInteractorMockTransactor, error) {
	contract, err := bindZetaInteractorMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaInteractorMockTransactor{contract: contract}, nil
}

// NewZetaInteractorMockFilterer creates a new log filterer instance of ZetaInteractorMock, bound to a specific deployed contract.
func NewZetaInteractorMockFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaInteractorMockFilterer, error) {
	contract, err := bindZetaInteractorMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaInteractorMockFilterer{contract: contract}, nil
}

// bindZetaInteractorMock binds a generic wrapper to an already deployed contract.
func bindZetaInteractorMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaInteractorMockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaInteractorMock *ZetaInteractorMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaInteractorMock.Contract.ZetaInteractorMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaInteractorMock *ZetaInteractorMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaInteractorMock.Contract.ZetaInteractorMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaInteractorMock *ZetaInteractorMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaInteractorMock.Contract.ZetaInteractorMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaInteractorMock *ZetaInteractorMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaInteractorMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaInteractorMock *ZetaInteractorMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaInteractorMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaInteractorMock *ZetaInteractorMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaInteractorMock.Contract.contract.Transact(opts, method, params...)
}

// Connector is a free data retrieval call binding the contract method 0x83f3084f.
//
// Solidity: function connector() view returns(address)
func (_ZetaInteractorMock *ZetaInteractorMockCaller) Connector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaInteractorMock.contract.Call(opts, &out, "connector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Connector is a free data retrieval call binding the contract method 0x83f3084f.
//
// Solidity: function connector() view returns(address)
func (_ZetaInteractorMock *ZetaInteractorMockSession) Connector() (common.Address, error) {
	return _ZetaInteractorMock.Contract.Connector(&_ZetaInteractorMock.CallOpts)
}

// Connector is a free data retrieval call binding the contract method 0x83f3084f.
//
// Solidity: function connector() view returns(address)
func (_ZetaInteractorMock *ZetaInteractorMockCallerSession) Connector() (common.Address, error) {
	return _ZetaInteractorMock.Contract.Connector(&_ZetaInteractorMock.CallOpts)
}

// InteractorsByChainId is a free data retrieval call binding the contract method 0x2618143f.
//
// Solidity: function interactorsByChainId(uint256 ) view returns(bytes)
func (_ZetaInteractorMock *ZetaInteractorMockCaller) InteractorsByChainId(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _ZetaInteractorMock.contract.Call(opts, &out, "interactorsByChainId", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// InteractorsByChainId is a free data retrieval call binding the contract method 0x2618143f.
//
// Solidity: function interactorsByChainId(uint256 ) view returns(bytes)
func (_ZetaInteractorMock *ZetaInteractorMockSession) InteractorsByChainId(arg0 *big.Int) ([]byte, error) {
	return _ZetaInteractorMock.Contract.InteractorsByChainId(&_ZetaInteractorMock.CallOpts, arg0)
}

// InteractorsByChainId is a free data retrieval call binding the contract method 0x2618143f.
//
// Solidity: function interactorsByChainId(uint256 ) view returns(bytes)
func (_ZetaInteractorMock *ZetaInteractorMockCallerSession) InteractorsByChainId(arg0 *big.Int) ([]byte, error) {
	return _ZetaInteractorMock.Contract.InteractorsByChainId(&_ZetaInteractorMock.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZetaInteractorMock *ZetaInteractorMockCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaInteractorMock.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZetaInteractorMock *ZetaInteractorMockSession) Owner() (common.Address, error) {
	return _ZetaInteractorMock.Contract.Owner(&_ZetaInteractorMock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZetaInteractorMock *ZetaInteractorMockCallerSession) Owner() (common.Address, error) {
	return _ZetaInteractorMock.Contract.Owner(&_ZetaInteractorMock.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_ZetaInteractorMock *ZetaInteractorMockCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaInteractorMock.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_ZetaInteractorMock *ZetaInteractorMockSession) PendingOwner() (common.Address, error) {
	return _ZetaInteractorMock.Contract.PendingOwner(&_ZetaInteractorMock.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_ZetaInteractorMock *ZetaInteractorMockCallerSession) PendingOwner() (common.Address, error) {
	return _ZetaInteractorMock.Contract.PendingOwner(&_ZetaInteractorMock.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ZetaInteractorMock *ZetaInteractorMockTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaInteractorMock.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ZetaInteractorMock *ZetaInteractorMockSession) AcceptOwnership() (*types.Transaction, error) {
	return _ZetaInteractorMock.Contract.AcceptOwnership(&_ZetaInteractorMock.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ZetaInteractorMock *ZetaInteractorMockTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ZetaInteractorMock.Contract.AcceptOwnership(&_ZetaInteractorMock.TransactOpts)
}

// OnZetaMessage is a paid mutator transaction binding the contract method 0x3749c51a.
//
// Solidity: function onZetaMessage((bytes,uint256,address,uint256,bytes) zetaMessage) returns()
func (_ZetaInteractorMock *ZetaInteractorMockTransactor) OnZetaMessage(opts *bind.TransactOpts, zetaMessage ZetaInterfacesZetaMessage) (*types.Transaction, error) {
	return _ZetaInteractorMock.contract.Transact(opts, "onZetaMessage", zetaMessage)
}

// OnZetaMessage is a paid mutator transaction binding the contract method 0x3749c51a.
//
// Solidity: function onZetaMessage((bytes,uint256,address,uint256,bytes) zetaMessage) returns()
func (_ZetaInteractorMock *ZetaInteractorMockSession) OnZetaMessage(zetaMessage ZetaInterfacesZetaMessage) (*types.Transaction, error) {
	return _ZetaInteractorMock.Contract.OnZetaMessage(&_ZetaInteractorMock.TransactOpts, zetaMessage)
}

// OnZetaMessage is a paid mutator transaction binding the contract method 0x3749c51a.
//
// Solidity: function onZetaMessage((bytes,uint256,address,uint256,bytes) zetaMessage) returns()
func (_ZetaInteractorMock *ZetaInteractorMockTransactorSession) OnZetaMessage(zetaMessage ZetaInterfacesZetaMessage) (*types.Transaction, error) {
	return _ZetaInteractorMock.Contract.OnZetaMessage(&_ZetaInteractorMock.TransactOpts, zetaMessage)
}

// OnZetaRevert is a paid mutator transaction binding the contract method 0x3ff0693c.
//
// Solidity: function onZetaRevert((address,uint256,bytes,uint256,uint256,bytes) zetaRevert) returns()
func (_ZetaInteractorMock *ZetaInteractorMockTransactor) OnZetaRevert(opts *bind.TransactOpts, zetaRevert ZetaInterfacesZetaRevert) (*types.Transaction, error) {
	return _ZetaInteractorMock.contract.Transact(opts, "onZetaRevert", zetaRevert)
}

// OnZetaRevert is a paid mutator transaction binding the contract method 0x3ff0693c.
//
// Solidity: function onZetaRevert((address,uint256,bytes,uint256,uint256,bytes) zetaRevert) returns()
func (_ZetaInteractorMock *ZetaInteractorMockSession) OnZetaRevert(zetaRevert ZetaInterfacesZetaRevert) (*types.Transaction, error) {
	return _ZetaInteractorMock.Contract.OnZetaRevert(&_ZetaInteractorMock.TransactOpts, zetaRevert)
}

// OnZetaRevert is a paid mutator transaction binding the contract method 0x3ff0693c.
//
// Solidity: function onZetaRevert((address,uint256,bytes,uint256,uint256,bytes) zetaRevert) returns()
func (_ZetaInteractorMock *ZetaInteractorMockTransactorSession) OnZetaRevert(zetaRevert ZetaInterfacesZetaRevert) (*types.Transaction, error) {
	return _ZetaInteractorMock.Contract.OnZetaRevert(&_ZetaInteractorMock.TransactOpts, zetaRevert)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZetaInteractorMock *ZetaInteractorMockTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaInteractorMock.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZetaInteractorMock *ZetaInteractorMockSession) RenounceOwnership() (*types.Transaction, error) {
	return _ZetaInteractorMock.Contract.RenounceOwnership(&_ZetaInteractorMock.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZetaInteractorMock *ZetaInteractorMockTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ZetaInteractorMock.Contract.RenounceOwnership(&_ZetaInteractorMock.TransactOpts)
}

// SetInteractorByChainId is a paid mutator transaction binding the contract method 0x4fd3f7d7.
//
// Solidity: function setInteractorByChainId(uint256 destinationChainId, bytes contractAddress) returns()
func (_ZetaInteractorMock *ZetaInteractorMockTransactor) SetInteractorByChainId(opts *bind.TransactOpts, destinationChainId *big.Int, contractAddress []byte) (*types.Transaction, error) {
	return _ZetaInteractorMock.contract.Transact(opts, "setInteractorByChainId", destinationChainId, contractAddress)
}

// SetInteractorByChainId is a paid mutator transaction binding the contract method 0x4fd3f7d7.
//
// Solidity: function setInteractorByChainId(uint256 destinationChainId, bytes contractAddress) returns()
func (_ZetaInteractorMock *ZetaInteractorMockSession) SetInteractorByChainId(destinationChainId *big.Int, contractAddress []byte) (*types.Transaction, error) {
	return _ZetaInteractorMock.Contract.SetInteractorByChainId(&_ZetaInteractorMock.TransactOpts, destinationChainId, contractAddress)
}

// SetInteractorByChainId is a paid mutator transaction binding the contract method 0x4fd3f7d7.
//
// Solidity: function setInteractorByChainId(uint256 destinationChainId, bytes contractAddress) returns()
func (_ZetaInteractorMock *ZetaInteractorMockTransactorSession) SetInteractorByChainId(destinationChainId *big.Int, contractAddress []byte) (*types.Transaction, error) {
	return _ZetaInteractorMock.Contract.SetInteractorByChainId(&_ZetaInteractorMock.TransactOpts, destinationChainId, contractAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZetaInteractorMock *ZetaInteractorMockTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ZetaInteractorMock.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZetaInteractorMock *ZetaInteractorMockSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZetaInteractorMock.Contract.TransferOwnership(&_ZetaInteractorMock.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZetaInteractorMock *ZetaInteractorMockTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZetaInteractorMock.Contract.TransferOwnership(&_ZetaInteractorMock.TransactOpts, newOwner)
}

// ZetaInteractorMockOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the ZetaInteractorMock contract.
type ZetaInteractorMockOwnershipTransferStartedIterator struct {
	Event *ZetaInteractorMockOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *ZetaInteractorMockOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaInteractorMockOwnershipTransferStarted)
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
		it.Event = new(ZetaInteractorMockOwnershipTransferStarted)
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
func (it *ZetaInteractorMockOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaInteractorMockOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaInteractorMockOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the ZetaInteractorMock contract.
type ZetaInteractorMockOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_ZetaInteractorMock *ZetaInteractorMockFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ZetaInteractorMockOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZetaInteractorMock.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ZetaInteractorMockOwnershipTransferStartedIterator{contract: _ZetaInteractorMock.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_ZetaInteractorMock *ZetaInteractorMockFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *ZetaInteractorMockOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZetaInteractorMock.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaInteractorMockOwnershipTransferStarted)
				if err := _ZetaInteractorMock.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_ZetaInteractorMock *ZetaInteractorMockFilterer) ParseOwnershipTransferStarted(log types.Log) (*ZetaInteractorMockOwnershipTransferStarted, error) {
	event := new(ZetaInteractorMockOwnershipTransferStarted)
	if err := _ZetaInteractorMock.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaInteractorMockOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ZetaInteractorMock contract.
type ZetaInteractorMockOwnershipTransferredIterator struct {
	Event *ZetaInteractorMockOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ZetaInteractorMockOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaInteractorMockOwnershipTransferred)
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
		it.Event = new(ZetaInteractorMockOwnershipTransferred)
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
func (it *ZetaInteractorMockOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaInteractorMockOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaInteractorMockOwnershipTransferred represents a OwnershipTransferred event raised by the ZetaInteractorMock contract.
type ZetaInteractorMockOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZetaInteractorMock *ZetaInteractorMockFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ZetaInteractorMockOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZetaInteractorMock.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ZetaInteractorMockOwnershipTransferredIterator{contract: _ZetaInteractorMock.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZetaInteractorMock *ZetaInteractorMockFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ZetaInteractorMockOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZetaInteractorMock.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaInteractorMockOwnershipTransferred)
				if err := _ZetaInteractorMock.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZetaInteractorMock *ZetaInteractorMockFilterer) ParseOwnershipTransferred(log types.Log) (*ZetaInteractorMockOwnershipTransferred, error) {
	event := new(ZetaInteractorMockOwnershipTransferred)
	if err := _ZetaInteractorMock.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
