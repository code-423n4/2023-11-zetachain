// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package systemcontract

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

// ZContext is an auto generated low-level Go binding around an user-defined struct.
type ZContext struct {
	Origin  []byte
	Sender  common.Address
	ChainID *big.Int
}

// SystemContractMetaData contains all meta data concerning the SystemContract contract.
var SystemContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wzeta_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uniswapv2Factory_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uniswapv2Router02_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerIsNotFungibleModule\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CantBeIdenticalAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CantBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"SetConnectorZEVM\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"SetGasCoin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"SetGasPrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"SetGasZetaPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"SetWZeta\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SystemContractDeployed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FUNGIBLE_MODULE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"origin\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"internalType\":\"structzContext\",\"name\":\"context\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"zrc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"depositAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gasCoinZRC20ByChainId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gasPriceByChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gasZetaPoolByChainId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setConnectorZEVMAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"zrc20\",\"type\":\"address\"}],\"name\":\"setGasCoinZRC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setGasPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"}],\"name\":\"setGasZetaPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setWZETAContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapv2FactoryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"uniswapv2PairFor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapv2Router02Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wZetaContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zetaConnectorZEVMAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b50604051620018aa380380620018aa8339818101604052810190620000379190620001ac565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614620000b1576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1660601b815250508073ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1660601b815250507f80699e81136d69cb8367ad52a994e25c722a86da654b561d0c14b61a777e7ac560405160405180910390a15050506200025b565b600081519050620001a68162000241565b92915050565b600080600060608486031215620001c857620001c76200023c565b5b6000620001d88682870162000195565b9350506020620001eb8682870162000195565b9250506040620001fe8682870162000195565b9150509250925092565b600062000215826200021c565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6200024c8162000208565b81146200025857600080fd5b50565b60805160601c60a05160601c61161c6200028e600039600061051b0152600081816105bd0152610bc5015261161c6000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806397770dff11610097578063c63585cc11610066578063c63585cc1461025e578063d7fd7afb1461028e578063d936a012146102be578063ee2815ba146102dc576100f5565b806397770dff146101ec578063a7cb050714610208578063c39aca3714610224578063c62178ac14610240576100f5565b8063513a9c05116100d3578063513a9c0514610164578063569541b914610194578063842da36d146101b257806391dd645f146101d0576100f5565b80630be15547146100fa5780631f0e251b1461012a5780633ce4a5bc14610146575b600080fd5b610114600480360381019061010f9190611022565b6102f8565b60405161012191906112b1565b60405180910390f35b610144600480360381019061013f9190610ebf565b61032b565b005b61014e6104a8565b60405161015b91906112b1565b60405180910390f35b61017e60048036038101906101799190611022565b6104c0565b60405161018b91906112b1565b60405180910390f35b61019c6104f3565b6040516101a991906112b1565b60405180910390f35b6101ba610519565b6040516101c791906112b1565b60405180910390f35b6101ea60048036038101906101e5919061104f565b61053d565b005b61020660048036038101906102019190610ebf565b610697565b005b610222600480360381019061021d919061108f565b610814565b005b61023e60048036038101906102399190610f6c565b6108e1565b005b610248610b13565b60405161025591906112b1565b60405180910390f35b61027860048036038101906102739190610eec565b610b39565b60405161028591906112b1565b60405180910390f35b6102a860048036038101906102a39190611022565b610bab565b6040516102b5919061134a565b60405180910390f35b6102c6610bc3565b6040516102d391906112b1565b60405180910390f35b6102f660048036038101906102f1919061104f565b610be7565b005b60016020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103a4576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561040b576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f3ade88e3922d64780e1bf4460d364c2970b69da813f9c0c07a1c187b5647636c600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660405161049d91906112b1565b60405180910390a150565b73735b14bb79463307aacbed86daf3322b1e6226ab81565b60026020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105b6576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006106057f0000000000000000000000000000000000000000000000000000000000000000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684610b39565b9050806002600085815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f0ecec485166da6139b13bb7e033e9446e2d35348e80ebf1180d4afe2dba1704e838260405161068a929190611365565b60405180910390a1505050565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610710576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610777576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fdba79d534382d1a8ae108e4c8ecb27c6ae42ab8b91d44eedf88bd329f3868d5e600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660405161080991906112b1565b60405180910390a150565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461088d576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600080848152602001908152602001600020819055507f49f492222906ac486c3c1401fa545626df1f0c0e5a77a05597ea2ed66af9850d82826040516108d592919061138e565b60405180910390a15050565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461095a576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614806109d357503073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b15610a0a576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8473ffffffffffffffffffffffffffffffffffffffff166347e7ef2484866040518363ffffffff1660e01b8152600401610a459291906112cc565b602060405180830381600087803b158015610a5f57600080fd5b505af1158015610a73573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a979190610f3f565b508273ffffffffffffffffffffffffffffffffffffffff1663de43156e87878786866040518663ffffffff1660e01b8152600401610ad99594939291906112f5565b600060405180830381600087803b158015610af357600080fd5b505af1158015610b07573d6000803e3d6000fd5b50505050505050505050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000806000610b488585610cef565b91509150858282604051602001610b60929190611243565b60405160208183030381529060405280519060200120604051602001610b8792919061126f565b6040516020818303038152906040528051906020012060001c925050509392505050565b60006020528060005260406000206000915090505481565b7f000000000000000000000000000000000000000000000000000000000000000081565b73735b14bb79463307aacbed86daf3322b1e6226ab73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c60576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806001600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fd1b36d30f6248e97c473b4d1348ca164a4ef6759022f54a58ec200326c39c45d8282604051610ce3929190611365565b60405180910390a15050565b6000808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415610d58576040517fcb1e7cfe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1610610d92578284610d95565b83835b8092508193505050600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610e04576040517f78b507da00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9250929050565b600081359050610e1a816115a1565b92915050565b600081519050610e2f816115b8565b92915050565b60008083601f840112610e4b57610e4a61150e565b5b8235905067ffffffffffffffff811115610e6857610e67611509565b5b602083019150836001820283011115610e8457610e8361151d565b5b9250929050565b600060608284031215610ea157610ea0611513565b5b81905092915050565b600081359050610eb9816115cf565b92915050565b600060208284031215610ed557610ed461152c565b5b6000610ee384828501610e0b565b91505092915050565b600080600060608486031215610f0557610f0461152c565b5b6000610f1386828701610e0b565b9350506020610f2486828701610e0b565b9250506040610f3586828701610e0b565b9150509250925092565b600060208284031215610f5557610f5461152c565b5b6000610f6384828501610e20565b91505092915050565b60008060008060008060a08789031215610f8957610f8861152c565b5b600087013567ffffffffffffffff811115610fa757610fa6611522565b5b610fb389828a01610e8b565b9650506020610fc489828a01610e0b565b9550506040610fd589828a01610eaa565b9450506060610fe689828a01610e0b565b935050608087013567ffffffffffffffff81111561100757611006611522565b5b61101389828a01610e35565b92509250509295509295509295565b6000602082840312156110385761103761152c565b5b600061104684828501610eaa565b91505092915050565b600080604083850312156110665761106561152c565b5b600061107485828601610eaa565b925050602061108585828601610e0b565b9150509250929050565b600080604083850312156110a6576110a561152c565b5b60006110b485828601610eaa565b92505060206110c585828601610eaa565b9150509250929050565b6110d881611475565b82525050565b6110e781611475565b82525050565b6110fe6110f982611475565b6114d6565b82525050565b61111561111082611493565b6114e8565b82525050565b600061112783856113b7565b93506111348385846114c7565b61113d83611531565b840190509392505050565b600061115483856113c8565b93506111618385846114c7565b61116a83611531565b840190509392505050565b60006111826020836113d9565b915061118d8261154f565b602082019050919050565b60006111a56001836113d9565b91506111b082611578565b600182019050919050565b6000606083016111ce60008401846113fb565b85830360008701526111e183828461111b565b925050506111f260208401846113e4565b6111ff60208601826110cf565b5061120d604084018461145e565b61121a6040860182611225565b508091505092915050565b61122e816114bd565b82525050565b61123d816114bd565b82525050565b600061124f82856110ed565b60148201915061125f82846110ed565b6014820191508190509392505050565b600061127a82611198565b915061128682856110ed565b6014820191506112968284611104565b6020820191506112a582611175565b91508190509392505050565b60006020820190506112c660008301846110de565b92915050565b60006040820190506112e160008301856110de565b6112ee6020830184611234565b9392505050565b6000608082019050818103600083015261130f81886111bb565b905061131e60208301876110de565b61132b6040830186611234565b818103606083015261133e818486611148565b90509695505050505050565b600060208201905061135f6000830184611234565b92915050565b600060408201905061137a6000830185611234565b61138760208301846110de565b9392505050565b60006040820190506113a36000830185611234565b6113b06020830184611234565b9392505050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b60006113f36020840184610e0b565b905092915050565b6000808335600160200384360303811261141857611417611527565b5b83810192508235915060208301925067ffffffffffffffff8211156114405761143f611504565b5b60018202360384131561145657611455611518565b5b509250929050565b600061146d6020840184610eaa565b905092915050565b60006114808261149d565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60006114e1826114f2565b9050919050565b6000819050919050565b60006114fd82611542565b9050919050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b7f96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f600082015250565b7fff00000000000000000000000000000000000000000000000000000000000000600082015250565b6115aa81611475565b81146115b557600080fd5b50565b6115c181611487565b81146115cc57600080fd5b50565b6115d8816114bd565b81146115e357600080fd5b5056fea2646970667358221220d4153f74dccdcf8c4e404c06ec70932b2acc46ffc5784c13399fb6b547739b2064736f6c63430008070033",
}

// SystemContractABI is the input ABI used to generate the binding from.
// Deprecated: Use SystemContractMetaData.ABI instead.
var SystemContractABI = SystemContractMetaData.ABI

// SystemContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SystemContractMetaData.Bin instead.
var SystemContractBin = SystemContractMetaData.Bin

// DeploySystemContract deploys a new Ethereum contract, binding an instance of SystemContract to it.
func DeploySystemContract(auth *bind.TransactOpts, backend bind.ContractBackend, wzeta_ common.Address, uniswapv2Factory_ common.Address, uniswapv2Router02_ common.Address) (common.Address, *types.Transaction, *SystemContract, error) {
	parsed, err := SystemContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SystemContractBin), backend, wzeta_, uniswapv2Factory_, uniswapv2Router02_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SystemContract{SystemContractCaller: SystemContractCaller{contract: contract}, SystemContractTransactor: SystemContractTransactor{contract: contract}, SystemContractFilterer: SystemContractFilterer{contract: contract}}, nil
}

// SystemContract is an auto generated Go binding around an Ethereum contract.
type SystemContract struct {
	SystemContractCaller     // Read-only binding to the contract
	SystemContractTransactor // Write-only binding to the contract
	SystemContractFilterer   // Log filterer for contract events
}

// SystemContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemContractSession struct {
	Contract     *SystemContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemContractCallerSession struct {
	Contract *SystemContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SystemContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemContractTransactorSession struct {
	Contract     *SystemContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SystemContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemContractRaw struct {
	Contract *SystemContract // Generic contract binding to access the raw methods on
}

// SystemContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemContractCallerRaw struct {
	Contract *SystemContractCaller // Generic read-only contract binding to access the raw methods on
}

// SystemContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemContractTransactorRaw struct {
	Contract *SystemContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemContract creates a new instance of SystemContract, bound to a specific deployed contract.
func NewSystemContract(address common.Address, backend bind.ContractBackend) (*SystemContract, error) {
	contract, err := bindSystemContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemContract{SystemContractCaller: SystemContractCaller{contract: contract}, SystemContractTransactor: SystemContractTransactor{contract: contract}, SystemContractFilterer: SystemContractFilterer{contract: contract}}, nil
}

// NewSystemContractCaller creates a new read-only instance of SystemContract, bound to a specific deployed contract.
func NewSystemContractCaller(address common.Address, caller bind.ContractCaller) (*SystemContractCaller, error) {
	contract, err := bindSystemContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemContractCaller{contract: contract}, nil
}

// NewSystemContractTransactor creates a new write-only instance of SystemContract, bound to a specific deployed contract.
func NewSystemContractTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemContractTransactor, error) {
	contract, err := bindSystemContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemContractTransactor{contract: contract}, nil
}

// NewSystemContractFilterer creates a new log filterer instance of SystemContract, bound to a specific deployed contract.
func NewSystemContractFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemContractFilterer, error) {
	contract, err := bindSystemContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemContractFilterer{contract: contract}, nil
}

// bindSystemContract binds a generic wrapper to an already deployed contract.
func bindSystemContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SystemContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemContract *SystemContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemContract.Contract.SystemContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemContract *SystemContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemContract.Contract.SystemContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemContract *SystemContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemContract.Contract.SystemContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemContract *SystemContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemContract *SystemContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemContract *SystemContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemContract.Contract.contract.Transact(opts, method, params...)
}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_SystemContract *SystemContractCaller) FUNGIBLEMODULEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemContract.contract.Call(opts, &out, "FUNGIBLE_MODULE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_SystemContract *SystemContractSession) FUNGIBLEMODULEADDRESS() (common.Address, error) {
	return _SystemContract.Contract.FUNGIBLEMODULEADDRESS(&_SystemContract.CallOpts)
}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_SystemContract *SystemContractCallerSession) FUNGIBLEMODULEADDRESS() (common.Address, error) {
	return _SystemContract.Contract.FUNGIBLEMODULEADDRESS(&_SystemContract.CallOpts)
}

// GasCoinZRC20ByChainId is a free data retrieval call binding the contract method 0x0be15547.
//
// Solidity: function gasCoinZRC20ByChainId(uint256 ) view returns(address)
func (_SystemContract *SystemContractCaller) GasCoinZRC20ByChainId(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SystemContract.contract.Call(opts, &out, "gasCoinZRC20ByChainId", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasCoinZRC20ByChainId is a free data retrieval call binding the contract method 0x0be15547.
//
// Solidity: function gasCoinZRC20ByChainId(uint256 ) view returns(address)
func (_SystemContract *SystemContractSession) GasCoinZRC20ByChainId(arg0 *big.Int) (common.Address, error) {
	return _SystemContract.Contract.GasCoinZRC20ByChainId(&_SystemContract.CallOpts, arg0)
}

// GasCoinZRC20ByChainId is a free data retrieval call binding the contract method 0x0be15547.
//
// Solidity: function gasCoinZRC20ByChainId(uint256 ) view returns(address)
func (_SystemContract *SystemContractCallerSession) GasCoinZRC20ByChainId(arg0 *big.Int) (common.Address, error) {
	return _SystemContract.Contract.GasCoinZRC20ByChainId(&_SystemContract.CallOpts, arg0)
}

// GasPriceByChainId is a free data retrieval call binding the contract method 0xd7fd7afb.
//
// Solidity: function gasPriceByChainId(uint256 ) view returns(uint256)
func (_SystemContract *SystemContractCaller) GasPriceByChainId(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SystemContract.contract.Call(opts, &out, "gasPriceByChainId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasPriceByChainId is a free data retrieval call binding the contract method 0xd7fd7afb.
//
// Solidity: function gasPriceByChainId(uint256 ) view returns(uint256)
func (_SystemContract *SystemContractSession) GasPriceByChainId(arg0 *big.Int) (*big.Int, error) {
	return _SystemContract.Contract.GasPriceByChainId(&_SystemContract.CallOpts, arg0)
}

// GasPriceByChainId is a free data retrieval call binding the contract method 0xd7fd7afb.
//
// Solidity: function gasPriceByChainId(uint256 ) view returns(uint256)
func (_SystemContract *SystemContractCallerSession) GasPriceByChainId(arg0 *big.Int) (*big.Int, error) {
	return _SystemContract.Contract.GasPriceByChainId(&_SystemContract.CallOpts, arg0)
}

// GasZetaPoolByChainId is a free data retrieval call binding the contract method 0x513a9c05.
//
// Solidity: function gasZetaPoolByChainId(uint256 ) view returns(address)
func (_SystemContract *SystemContractCaller) GasZetaPoolByChainId(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SystemContract.contract.Call(opts, &out, "gasZetaPoolByChainId", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasZetaPoolByChainId is a free data retrieval call binding the contract method 0x513a9c05.
//
// Solidity: function gasZetaPoolByChainId(uint256 ) view returns(address)
func (_SystemContract *SystemContractSession) GasZetaPoolByChainId(arg0 *big.Int) (common.Address, error) {
	return _SystemContract.Contract.GasZetaPoolByChainId(&_SystemContract.CallOpts, arg0)
}

// GasZetaPoolByChainId is a free data retrieval call binding the contract method 0x513a9c05.
//
// Solidity: function gasZetaPoolByChainId(uint256 ) view returns(address)
func (_SystemContract *SystemContractCallerSession) GasZetaPoolByChainId(arg0 *big.Int) (common.Address, error) {
	return _SystemContract.Contract.GasZetaPoolByChainId(&_SystemContract.CallOpts, arg0)
}

// Uniswapv2FactoryAddress is a free data retrieval call binding the contract method 0xd936a012.
//
// Solidity: function uniswapv2FactoryAddress() view returns(address)
func (_SystemContract *SystemContractCaller) Uniswapv2FactoryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemContract.contract.Call(opts, &out, "uniswapv2FactoryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Uniswapv2FactoryAddress is a free data retrieval call binding the contract method 0xd936a012.
//
// Solidity: function uniswapv2FactoryAddress() view returns(address)
func (_SystemContract *SystemContractSession) Uniswapv2FactoryAddress() (common.Address, error) {
	return _SystemContract.Contract.Uniswapv2FactoryAddress(&_SystemContract.CallOpts)
}

// Uniswapv2FactoryAddress is a free data retrieval call binding the contract method 0xd936a012.
//
// Solidity: function uniswapv2FactoryAddress() view returns(address)
func (_SystemContract *SystemContractCallerSession) Uniswapv2FactoryAddress() (common.Address, error) {
	return _SystemContract.Contract.Uniswapv2FactoryAddress(&_SystemContract.CallOpts)
}

// Uniswapv2PairFor is a free data retrieval call binding the contract method 0xc63585cc.
//
// Solidity: function uniswapv2PairFor(address factory, address tokenA, address tokenB) pure returns(address pair)
func (_SystemContract *SystemContractCaller) Uniswapv2PairFor(opts *bind.CallOpts, factory common.Address, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	var out []interface{}
	err := _SystemContract.contract.Call(opts, &out, "uniswapv2PairFor", factory, tokenA, tokenB)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Uniswapv2PairFor is a free data retrieval call binding the contract method 0xc63585cc.
//
// Solidity: function uniswapv2PairFor(address factory, address tokenA, address tokenB) pure returns(address pair)
func (_SystemContract *SystemContractSession) Uniswapv2PairFor(factory common.Address, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _SystemContract.Contract.Uniswapv2PairFor(&_SystemContract.CallOpts, factory, tokenA, tokenB)
}

// Uniswapv2PairFor is a free data retrieval call binding the contract method 0xc63585cc.
//
// Solidity: function uniswapv2PairFor(address factory, address tokenA, address tokenB) pure returns(address pair)
func (_SystemContract *SystemContractCallerSession) Uniswapv2PairFor(factory common.Address, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _SystemContract.Contract.Uniswapv2PairFor(&_SystemContract.CallOpts, factory, tokenA, tokenB)
}

// Uniswapv2Router02Address is a free data retrieval call binding the contract method 0x842da36d.
//
// Solidity: function uniswapv2Router02Address() view returns(address)
func (_SystemContract *SystemContractCaller) Uniswapv2Router02Address(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemContract.contract.Call(opts, &out, "uniswapv2Router02Address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Uniswapv2Router02Address is a free data retrieval call binding the contract method 0x842da36d.
//
// Solidity: function uniswapv2Router02Address() view returns(address)
func (_SystemContract *SystemContractSession) Uniswapv2Router02Address() (common.Address, error) {
	return _SystemContract.Contract.Uniswapv2Router02Address(&_SystemContract.CallOpts)
}

// Uniswapv2Router02Address is a free data retrieval call binding the contract method 0x842da36d.
//
// Solidity: function uniswapv2Router02Address() view returns(address)
func (_SystemContract *SystemContractCallerSession) Uniswapv2Router02Address() (common.Address, error) {
	return _SystemContract.Contract.Uniswapv2Router02Address(&_SystemContract.CallOpts)
}

// WZetaContractAddress is a free data retrieval call binding the contract method 0x569541b9.
//
// Solidity: function wZetaContractAddress() view returns(address)
func (_SystemContract *SystemContractCaller) WZetaContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemContract.contract.Call(opts, &out, "wZetaContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WZetaContractAddress is a free data retrieval call binding the contract method 0x569541b9.
//
// Solidity: function wZetaContractAddress() view returns(address)
func (_SystemContract *SystemContractSession) WZetaContractAddress() (common.Address, error) {
	return _SystemContract.Contract.WZetaContractAddress(&_SystemContract.CallOpts)
}

// WZetaContractAddress is a free data retrieval call binding the contract method 0x569541b9.
//
// Solidity: function wZetaContractAddress() view returns(address)
func (_SystemContract *SystemContractCallerSession) WZetaContractAddress() (common.Address, error) {
	return _SystemContract.Contract.WZetaContractAddress(&_SystemContract.CallOpts)
}

// ZetaConnectorZEVMAddress is a free data retrieval call binding the contract method 0xc62178ac.
//
// Solidity: function zetaConnectorZEVMAddress() view returns(address)
func (_SystemContract *SystemContractCaller) ZetaConnectorZEVMAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemContract.contract.Call(opts, &out, "zetaConnectorZEVMAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaConnectorZEVMAddress is a free data retrieval call binding the contract method 0xc62178ac.
//
// Solidity: function zetaConnectorZEVMAddress() view returns(address)
func (_SystemContract *SystemContractSession) ZetaConnectorZEVMAddress() (common.Address, error) {
	return _SystemContract.Contract.ZetaConnectorZEVMAddress(&_SystemContract.CallOpts)
}

// ZetaConnectorZEVMAddress is a free data retrieval call binding the contract method 0xc62178ac.
//
// Solidity: function zetaConnectorZEVMAddress() view returns(address)
func (_SystemContract *SystemContractCallerSession) ZetaConnectorZEVMAddress() (common.Address, error) {
	return _SystemContract.Contract.ZetaConnectorZEVMAddress(&_SystemContract.CallOpts)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0xc39aca37.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_SystemContract *SystemContractTransactor) DepositAndCall(opts *bind.TransactOpts, context ZContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _SystemContract.contract.Transact(opts, "depositAndCall", context, zrc20, amount, target, message)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0xc39aca37.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_SystemContract *SystemContractSession) DepositAndCall(context ZContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _SystemContract.Contract.DepositAndCall(&_SystemContract.TransactOpts, context, zrc20, amount, target, message)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0xc39aca37.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_SystemContract *SystemContractTransactorSession) DepositAndCall(context ZContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _SystemContract.Contract.DepositAndCall(&_SystemContract.TransactOpts, context, zrc20, amount, target, message)
}

// SetConnectorZEVMAddress is a paid mutator transaction binding the contract method 0x1f0e251b.
//
// Solidity: function setConnectorZEVMAddress(address addr) returns()
func (_SystemContract *SystemContractTransactor) SetConnectorZEVMAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _SystemContract.contract.Transact(opts, "setConnectorZEVMAddress", addr)
}

// SetConnectorZEVMAddress is a paid mutator transaction binding the contract method 0x1f0e251b.
//
// Solidity: function setConnectorZEVMAddress(address addr) returns()
func (_SystemContract *SystemContractSession) SetConnectorZEVMAddress(addr common.Address) (*types.Transaction, error) {
	return _SystemContract.Contract.SetConnectorZEVMAddress(&_SystemContract.TransactOpts, addr)
}

// SetConnectorZEVMAddress is a paid mutator transaction binding the contract method 0x1f0e251b.
//
// Solidity: function setConnectorZEVMAddress(address addr) returns()
func (_SystemContract *SystemContractTransactorSession) SetConnectorZEVMAddress(addr common.Address) (*types.Transaction, error) {
	return _SystemContract.Contract.SetConnectorZEVMAddress(&_SystemContract.TransactOpts, addr)
}

// SetGasCoinZRC20 is a paid mutator transaction binding the contract method 0xee2815ba.
//
// Solidity: function setGasCoinZRC20(uint256 chainID, address zrc20) returns()
func (_SystemContract *SystemContractTransactor) SetGasCoinZRC20(opts *bind.TransactOpts, chainID *big.Int, zrc20 common.Address) (*types.Transaction, error) {
	return _SystemContract.contract.Transact(opts, "setGasCoinZRC20", chainID, zrc20)
}

// SetGasCoinZRC20 is a paid mutator transaction binding the contract method 0xee2815ba.
//
// Solidity: function setGasCoinZRC20(uint256 chainID, address zrc20) returns()
func (_SystemContract *SystemContractSession) SetGasCoinZRC20(chainID *big.Int, zrc20 common.Address) (*types.Transaction, error) {
	return _SystemContract.Contract.SetGasCoinZRC20(&_SystemContract.TransactOpts, chainID, zrc20)
}

// SetGasCoinZRC20 is a paid mutator transaction binding the contract method 0xee2815ba.
//
// Solidity: function setGasCoinZRC20(uint256 chainID, address zrc20) returns()
func (_SystemContract *SystemContractTransactorSession) SetGasCoinZRC20(chainID *big.Int, zrc20 common.Address) (*types.Transaction, error) {
	return _SystemContract.Contract.SetGasCoinZRC20(&_SystemContract.TransactOpts, chainID, zrc20)
}

// SetGasPrice is a paid mutator transaction binding the contract method 0xa7cb0507.
//
// Solidity: function setGasPrice(uint256 chainID, uint256 price) returns()
func (_SystemContract *SystemContractTransactor) SetGasPrice(opts *bind.TransactOpts, chainID *big.Int, price *big.Int) (*types.Transaction, error) {
	return _SystemContract.contract.Transact(opts, "setGasPrice", chainID, price)
}

// SetGasPrice is a paid mutator transaction binding the contract method 0xa7cb0507.
//
// Solidity: function setGasPrice(uint256 chainID, uint256 price) returns()
func (_SystemContract *SystemContractSession) SetGasPrice(chainID *big.Int, price *big.Int) (*types.Transaction, error) {
	return _SystemContract.Contract.SetGasPrice(&_SystemContract.TransactOpts, chainID, price)
}

// SetGasPrice is a paid mutator transaction binding the contract method 0xa7cb0507.
//
// Solidity: function setGasPrice(uint256 chainID, uint256 price) returns()
func (_SystemContract *SystemContractTransactorSession) SetGasPrice(chainID *big.Int, price *big.Int) (*types.Transaction, error) {
	return _SystemContract.Contract.SetGasPrice(&_SystemContract.TransactOpts, chainID, price)
}

// SetGasZetaPool is a paid mutator transaction binding the contract method 0x91dd645f.
//
// Solidity: function setGasZetaPool(uint256 chainID, address erc20) returns()
func (_SystemContract *SystemContractTransactor) SetGasZetaPool(opts *bind.TransactOpts, chainID *big.Int, erc20 common.Address) (*types.Transaction, error) {
	return _SystemContract.contract.Transact(opts, "setGasZetaPool", chainID, erc20)
}

// SetGasZetaPool is a paid mutator transaction binding the contract method 0x91dd645f.
//
// Solidity: function setGasZetaPool(uint256 chainID, address erc20) returns()
func (_SystemContract *SystemContractSession) SetGasZetaPool(chainID *big.Int, erc20 common.Address) (*types.Transaction, error) {
	return _SystemContract.Contract.SetGasZetaPool(&_SystemContract.TransactOpts, chainID, erc20)
}

// SetGasZetaPool is a paid mutator transaction binding the contract method 0x91dd645f.
//
// Solidity: function setGasZetaPool(uint256 chainID, address erc20) returns()
func (_SystemContract *SystemContractTransactorSession) SetGasZetaPool(chainID *big.Int, erc20 common.Address) (*types.Transaction, error) {
	return _SystemContract.Contract.SetGasZetaPool(&_SystemContract.TransactOpts, chainID, erc20)
}

// SetWZETAContractAddress is a paid mutator transaction binding the contract method 0x97770dff.
//
// Solidity: function setWZETAContractAddress(address addr) returns()
func (_SystemContract *SystemContractTransactor) SetWZETAContractAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _SystemContract.contract.Transact(opts, "setWZETAContractAddress", addr)
}

// SetWZETAContractAddress is a paid mutator transaction binding the contract method 0x97770dff.
//
// Solidity: function setWZETAContractAddress(address addr) returns()
func (_SystemContract *SystemContractSession) SetWZETAContractAddress(addr common.Address) (*types.Transaction, error) {
	return _SystemContract.Contract.SetWZETAContractAddress(&_SystemContract.TransactOpts, addr)
}

// SetWZETAContractAddress is a paid mutator transaction binding the contract method 0x97770dff.
//
// Solidity: function setWZETAContractAddress(address addr) returns()
func (_SystemContract *SystemContractTransactorSession) SetWZETAContractAddress(addr common.Address) (*types.Transaction, error) {
	return _SystemContract.Contract.SetWZETAContractAddress(&_SystemContract.TransactOpts, addr)
}

// SystemContractSetConnectorZEVMIterator is returned from FilterSetConnectorZEVM and is used to iterate over the raw logs and unpacked data for SetConnectorZEVM events raised by the SystemContract contract.
type SystemContractSetConnectorZEVMIterator struct {
	Event *SystemContractSetConnectorZEVM // Event containing the contract specifics and raw log

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
func (it *SystemContractSetConnectorZEVMIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemContractSetConnectorZEVM)
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
		it.Event = new(SystemContractSetConnectorZEVM)
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
func (it *SystemContractSetConnectorZEVMIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemContractSetConnectorZEVMIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemContractSetConnectorZEVM represents a SetConnectorZEVM event raised by the SystemContract contract.
type SystemContractSetConnectorZEVM struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetConnectorZEVM is a free log retrieval operation binding the contract event 0x3ade88e3922d64780e1bf4460d364c2970b69da813f9c0c07a1c187b5647636c.
//
// Solidity: event SetConnectorZEVM(address arg0)
func (_SystemContract *SystemContractFilterer) FilterSetConnectorZEVM(opts *bind.FilterOpts) (*SystemContractSetConnectorZEVMIterator, error) {

	logs, sub, err := _SystemContract.contract.FilterLogs(opts, "SetConnectorZEVM")
	if err != nil {
		return nil, err
	}
	return &SystemContractSetConnectorZEVMIterator{contract: _SystemContract.contract, event: "SetConnectorZEVM", logs: logs, sub: sub}, nil
}

// WatchSetConnectorZEVM is a free log subscription operation binding the contract event 0x3ade88e3922d64780e1bf4460d364c2970b69da813f9c0c07a1c187b5647636c.
//
// Solidity: event SetConnectorZEVM(address arg0)
func (_SystemContract *SystemContractFilterer) WatchSetConnectorZEVM(opts *bind.WatchOpts, sink chan<- *SystemContractSetConnectorZEVM) (event.Subscription, error) {

	logs, sub, err := _SystemContract.contract.WatchLogs(opts, "SetConnectorZEVM")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemContractSetConnectorZEVM)
				if err := _SystemContract.contract.UnpackLog(event, "SetConnectorZEVM", log); err != nil {
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

// ParseSetConnectorZEVM is a log parse operation binding the contract event 0x3ade88e3922d64780e1bf4460d364c2970b69da813f9c0c07a1c187b5647636c.
//
// Solidity: event SetConnectorZEVM(address arg0)
func (_SystemContract *SystemContractFilterer) ParseSetConnectorZEVM(log types.Log) (*SystemContractSetConnectorZEVM, error) {
	event := new(SystemContractSetConnectorZEVM)
	if err := _SystemContract.contract.UnpackLog(event, "SetConnectorZEVM", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemContractSetGasCoinIterator is returned from FilterSetGasCoin and is used to iterate over the raw logs and unpacked data for SetGasCoin events raised by the SystemContract contract.
type SystemContractSetGasCoinIterator struct {
	Event *SystemContractSetGasCoin // Event containing the contract specifics and raw log

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
func (it *SystemContractSetGasCoinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemContractSetGasCoin)
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
		it.Event = new(SystemContractSetGasCoin)
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
func (it *SystemContractSetGasCoinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemContractSetGasCoinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemContractSetGasCoin represents a SetGasCoin event raised by the SystemContract contract.
type SystemContractSetGasCoin struct {
	Arg0 *big.Int
	Arg1 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetGasCoin is a free log retrieval operation binding the contract event 0xd1b36d30f6248e97c473b4d1348ca164a4ef6759022f54a58ec200326c39c45d.
//
// Solidity: event SetGasCoin(uint256 arg0, address arg1)
func (_SystemContract *SystemContractFilterer) FilterSetGasCoin(opts *bind.FilterOpts) (*SystemContractSetGasCoinIterator, error) {

	logs, sub, err := _SystemContract.contract.FilterLogs(opts, "SetGasCoin")
	if err != nil {
		return nil, err
	}
	return &SystemContractSetGasCoinIterator{contract: _SystemContract.contract, event: "SetGasCoin", logs: logs, sub: sub}, nil
}

// WatchSetGasCoin is a free log subscription operation binding the contract event 0xd1b36d30f6248e97c473b4d1348ca164a4ef6759022f54a58ec200326c39c45d.
//
// Solidity: event SetGasCoin(uint256 arg0, address arg1)
func (_SystemContract *SystemContractFilterer) WatchSetGasCoin(opts *bind.WatchOpts, sink chan<- *SystemContractSetGasCoin) (event.Subscription, error) {

	logs, sub, err := _SystemContract.contract.WatchLogs(opts, "SetGasCoin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemContractSetGasCoin)
				if err := _SystemContract.contract.UnpackLog(event, "SetGasCoin", log); err != nil {
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

// ParseSetGasCoin is a log parse operation binding the contract event 0xd1b36d30f6248e97c473b4d1348ca164a4ef6759022f54a58ec200326c39c45d.
//
// Solidity: event SetGasCoin(uint256 arg0, address arg1)
func (_SystemContract *SystemContractFilterer) ParseSetGasCoin(log types.Log) (*SystemContractSetGasCoin, error) {
	event := new(SystemContractSetGasCoin)
	if err := _SystemContract.contract.UnpackLog(event, "SetGasCoin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemContractSetGasPriceIterator is returned from FilterSetGasPrice and is used to iterate over the raw logs and unpacked data for SetGasPrice events raised by the SystemContract contract.
type SystemContractSetGasPriceIterator struct {
	Event *SystemContractSetGasPrice // Event containing the contract specifics and raw log

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
func (it *SystemContractSetGasPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemContractSetGasPrice)
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
		it.Event = new(SystemContractSetGasPrice)
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
func (it *SystemContractSetGasPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemContractSetGasPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemContractSetGasPrice represents a SetGasPrice event raised by the SystemContract contract.
type SystemContractSetGasPrice struct {
	Arg0 *big.Int
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetGasPrice is a free log retrieval operation binding the contract event 0x49f492222906ac486c3c1401fa545626df1f0c0e5a77a05597ea2ed66af9850d.
//
// Solidity: event SetGasPrice(uint256 arg0, uint256 arg1)
func (_SystemContract *SystemContractFilterer) FilterSetGasPrice(opts *bind.FilterOpts) (*SystemContractSetGasPriceIterator, error) {

	logs, sub, err := _SystemContract.contract.FilterLogs(opts, "SetGasPrice")
	if err != nil {
		return nil, err
	}
	return &SystemContractSetGasPriceIterator{contract: _SystemContract.contract, event: "SetGasPrice", logs: logs, sub: sub}, nil
}

// WatchSetGasPrice is a free log subscription operation binding the contract event 0x49f492222906ac486c3c1401fa545626df1f0c0e5a77a05597ea2ed66af9850d.
//
// Solidity: event SetGasPrice(uint256 arg0, uint256 arg1)
func (_SystemContract *SystemContractFilterer) WatchSetGasPrice(opts *bind.WatchOpts, sink chan<- *SystemContractSetGasPrice) (event.Subscription, error) {

	logs, sub, err := _SystemContract.contract.WatchLogs(opts, "SetGasPrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemContractSetGasPrice)
				if err := _SystemContract.contract.UnpackLog(event, "SetGasPrice", log); err != nil {
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

// ParseSetGasPrice is a log parse operation binding the contract event 0x49f492222906ac486c3c1401fa545626df1f0c0e5a77a05597ea2ed66af9850d.
//
// Solidity: event SetGasPrice(uint256 arg0, uint256 arg1)
func (_SystemContract *SystemContractFilterer) ParseSetGasPrice(log types.Log) (*SystemContractSetGasPrice, error) {
	event := new(SystemContractSetGasPrice)
	if err := _SystemContract.contract.UnpackLog(event, "SetGasPrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemContractSetGasZetaPoolIterator is returned from FilterSetGasZetaPool and is used to iterate over the raw logs and unpacked data for SetGasZetaPool events raised by the SystemContract contract.
type SystemContractSetGasZetaPoolIterator struct {
	Event *SystemContractSetGasZetaPool // Event containing the contract specifics and raw log

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
func (it *SystemContractSetGasZetaPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemContractSetGasZetaPool)
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
		it.Event = new(SystemContractSetGasZetaPool)
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
func (it *SystemContractSetGasZetaPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemContractSetGasZetaPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemContractSetGasZetaPool represents a SetGasZetaPool event raised by the SystemContract contract.
type SystemContractSetGasZetaPool struct {
	Arg0 *big.Int
	Arg1 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetGasZetaPool is a free log retrieval operation binding the contract event 0x0ecec485166da6139b13bb7e033e9446e2d35348e80ebf1180d4afe2dba1704e.
//
// Solidity: event SetGasZetaPool(uint256 arg0, address arg1)
func (_SystemContract *SystemContractFilterer) FilterSetGasZetaPool(opts *bind.FilterOpts) (*SystemContractSetGasZetaPoolIterator, error) {

	logs, sub, err := _SystemContract.contract.FilterLogs(opts, "SetGasZetaPool")
	if err != nil {
		return nil, err
	}
	return &SystemContractSetGasZetaPoolIterator{contract: _SystemContract.contract, event: "SetGasZetaPool", logs: logs, sub: sub}, nil
}

// WatchSetGasZetaPool is a free log subscription operation binding the contract event 0x0ecec485166da6139b13bb7e033e9446e2d35348e80ebf1180d4afe2dba1704e.
//
// Solidity: event SetGasZetaPool(uint256 arg0, address arg1)
func (_SystemContract *SystemContractFilterer) WatchSetGasZetaPool(opts *bind.WatchOpts, sink chan<- *SystemContractSetGasZetaPool) (event.Subscription, error) {

	logs, sub, err := _SystemContract.contract.WatchLogs(opts, "SetGasZetaPool")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemContractSetGasZetaPool)
				if err := _SystemContract.contract.UnpackLog(event, "SetGasZetaPool", log); err != nil {
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

// ParseSetGasZetaPool is a log parse operation binding the contract event 0x0ecec485166da6139b13bb7e033e9446e2d35348e80ebf1180d4afe2dba1704e.
//
// Solidity: event SetGasZetaPool(uint256 arg0, address arg1)
func (_SystemContract *SystemContractFilterer) ParseSetGasZetaPool(log types.Log) (*SystemContractSetGasZetaPool, error) {
	event := new(SystemContractSetGasZetaPool)
	if err := _SystemContract.contract.UnpackLog(event, "SetGasZetaPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemContractSetWZetaIterator is returned from FilterSetWZeta and is used to iterate over the raw logs and unpacked data for SetWZeta events raised by the SystemContract contract.
type SystemContractSetWZetaIterator struct {
	Event *SystemContractSetWZeta // Event containing the contract specifics and raw log

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
func (it *SystemContractSetWZetaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemContractSetWZeta)
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
		it.Event = new(SystemContractSetWZeta)
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
func (it *SystemContractSetWZetaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemContractSetWZetaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemContractSetWZeta represents a SetWZeta event raised by the SystemContract contract.
type SystemContractSetWZeta struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetWZeta is a free log retrieval operation binding the contract event 0xdba79d534382d1a8ae108e4c8ecb27c6ae42ab8b91d44eedf88bd329f3868d5e.
//
// Solidity: event SetWZeta(address arg0)
func (_SystemContract *SystemContractFilterer) FilterSetWZeta(opts *bind.FilterOpts) (*SystemContractSetWZetaIterator, error) {

	logs, sub, err := _SystemContract.contract.FilterLogs(opts, "SetWZeta")
	if err != nil {
		return nil, err
	}
	return &SystemContractSetWZetaIterator{contract: _SystemContract.contract, event: "SetWZeta", logs: logs, sub: sub}, nil
}

// WatchSetWZeta is a free log subscription operation binding the contract event 0xdba79d534382d1a8ae108e4c8ecb27c6ae42ab8b91d44eedf88bd329f3868d5e.
//
// Solidity: event SetWZeta(address arg0)
func (_SystemContract *SystemContractFilterer) WatchSetWZeta(opts *bind.WatchOpts, sink chan<- *SystemContractSetWZeta) (event.Subscription, error) {

	logs, sub, err := _SystemContract.contract.WatchLogs(opts, "SetWZeta")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemContractSetWZeta)
				if err := _SystemContract.contract.UnpackLog(event, "SetWZeta", log); err != nil {
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

// ParseSetWZeta is a log parse operation binding the contract event 0xdba79d534382d1a8ae108e4c8ecb27c6ae42ab8b91d44eedf88bd329f3868d5e.
//
// Solidity: event SetWZeta(address arg0)
func (_SystemContract *SystemContractFilterer) ParseSetWZeta(log types.Log) (*SystemContractSetWZeta, error) {
	event := new(SystemContractSetWZeta)
	if err := _SystemContract.contract.UnpackLog(event, "SetWZeta", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemContractSystemContractDeployedIterator is returned from FilterSystemContractDeployed and is used to iterate over the raw logs and unpacked data for SystemContractDeployed events raised by the SystemContract contract.
type SystemContractSystemContractDeployedIterator struct {
	Event *SystemContractSystemContractDeployed // Event containing the contract specifics and raw log

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
func (it *SystemContractSystemContractDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemContractSystemContractDeployed)
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
		it.Event = new(SystemContractSystemContractDeployed)
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
func (it *SystemContractSystemContractDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemContractSystemContractDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemContractSystemContractDeployed represents a SystemContractDeployed event raised by the SystemContract contract.
type SystemContractSystemContractDeployed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSystemContractDeployed is a free log retrieval operation binding the contract event 0x80699e81136d69cb8367ad52a994e25c722a86da654b561d0c14b61a777e7ac5.
//
// Solidity: event SystemContractDeployed()
func (_SystemContract *SystemContractFilterer) FilterSystemContractDeployed(opts *bind.FilterOpts) (*SystemContractSystemContractDeployedIterator, error) {

	logs, sub, err := _SystemContract.contract.FilterLogs(opts, "SystemContractDeployed")
	if err != nil {
		return nil, err
	}
	return &SystemContractSystemContractDeployedIterator{contract: _SystemContract.contract, event: "SystemContractDeployed", logs: logs, sub: sub}, nil
}

// WatchSystemContractDeployed is a free log subscription operation binding the contract event 0x80699e81136d69cb8367ad52a994e25c722a86da654b561d0c14b61a777e7ac5.
//
// Solidity: event SystemContractDeployed()
func (_SystemContract *SystemContractFilterer) WatchSystemContractDeployed(opts *bind.WatchOpts, sink chan<- *SystemContractSystemContractDeployed) (event.Subscription, error) {

	logs, sub, err := _SystemContract.contract.WatchLogs(opts, "SystemContractDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemContractSystemContractDeployed)
				if err := _SystemContract.contract.UnpackLog(event, "SystemContractDeployed", log); err != nil {
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

// ParseSystemContractDeployed is a log parse operation binding the contract event 0x80699e81136d69cb8367ad52a994e25c722a86da654b561d0c14b61a777e7ac5.
//
// Solidity: event SystemContractDeployed()
func (_SystemContract *SystemContractFilterer) ParseSystemContractDeployed(log types.Log) (*SystemContractSystemContractDeployed, error) {
	event := new(SystemContractSystemContractDeployed)
	if err := _SystemContract.contract.UnpackLog(event, "SystemContractDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
