// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetatokenconsumeruniv3

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

// ZetaTokenConsumerUniV3MetaData contains all meta data concerning the ZetaTokenConsumerUniV3 contract.
var ZetaTokenConsumerUniV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"zetaToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uniswapV3Router_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uniswapV3Factory_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WETH9Address_\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"zetaPoolFee_\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"tokenPoolFee_\",\"type\":\"uint24\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ErrorSendingETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputCantBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"EthExchangedForZeta\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"TokenExchangedForZeta\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"ZetaExchangedForEth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"ZetaExchangedForToken\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WETH9Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zetaTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getEthFromZeta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"zetaTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getTokenFromZeta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"name\":\"getZetaFromEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getZetaFromToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasZetaLiquidity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenPoolFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV3Factory\",\"outputs\":[{\"internalType\":\"contractIUniswapV3Factory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV3Router\",\"outputs\":[{\"internalType\":\"contractISwapRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zetaPoolFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zetaToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6101406040523480156200001257600080fd5b50604051620029833803806200298383398181016040528101906200003891906200028a565b600073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff161480620000a05750600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16145b80620000d85750600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16145b80620001105750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b1562000148576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8573ffffffffffffffffffffffffffffffffffffffff1660e08173ffffffffffffffffffffffffffffffffffffffff1660601b815250508473ffffffffffffffffffffffffffffffffffffffff166101008173ffffffffffffffffffffffffffffffffffffffff1660601b815250508373ffffffffffffffffffffffffffffffffffffffff166101208173ffffffffffffffffffffffffffffffffffffffff1660601b815250508273ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff1660601b815250508162ffffff1660808162ffffff1660e81b815250508062ffffff1660a08162ffffff1660e81b81525050505050505050620003a2565b6000815190506200026d816200036e565b92915050565b600081519050620002848162000388565b92915050565b60008060008060008060c08789031215620002aa57620002a962000369565b5b6000620002ba89828a016200025c565b9650506020620002cd89828a016200025c565b9550506040620002e089828a016200025c565b9450506060620002f389828a016200025c565b93505060806200030689828a0162000273565b92505060a06200031989828a0162000273565b9150509295509295509295565b600062000333826200033a565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600062ffffff82169050919050565b600080fd5b620003798162000326565b81146200038557600080fd5b50565b62000393816200035a565b81146200039f57600080fd5b50565b60805160e81c60a05160e81c60c05160601c60e05160601c6101005160601c6101205160601c6124ad620004d660003960008181610d900152610ddb01526000818161046f0152818161068f015281816107cd015281816108c2015281816109fd01528181610b6f0152818161115401526112b20152600081816103af0152818161056101528181610748015281816109b301528181610a1f01528181610a7301528181610e380152818161110a0152818161117601526111c90152600081816103730152818161070601528181610aaf01528181610c1c01528181610e170152818161120b01526113c10152600081816106e501528181610db4015261122c0152600081816103eb01528181610727015281816108e601528181610aeb01528181610e5901526111ea01526124ad6000f3fe6080604052600436106100a05760003560e01c806354c49a2a1161006457806354c49a2a1461019a5780635b549182146101d75780635d9dfdde1461020257806380801f841461022d578063a53fb10b14610258578063c469cf1414610295576100a7565b8063013b2ff8146100ac57806321e093b1146100dc5780632405620a146101075780632c76d7a6146101445780633cbd70051461016f576100a7565b366100a757005b600080fd5b6100c660048036038101906100c1919061190a565b6102c0565b6040516100d391906120a2565b60405180910390f35b3480156100e857600080fd5b506100f161055f565b6040516100fe9190611e44565b60405180910390f35b34801561011357600080fd5b5061012e6004803603810190610129919061194a565b610583565b60405161013b91906120a2565b60405180910390f35b34801561015057600080fd5b506101596108c0565b6040516101669190611f71565b60405180910390f35b34801561017b57600080fd5b506101846108e4565b6040516101919190612087565b60405180910390f35b3480156101a657600080fd5b506101c160048036038101906101bc91906119b1565b610908565b6040516101ce91906120a2565b60405180910390f35b3480156101e357600080fd5b506101ec610d8e565b6040516101f99190611f8c565b60405180910390f35b34801561020e57600080fd5b50610217610db2565b6040516102249190612087565b60405180910390f35b34801561023957600080fd5b50610242610dd6565b60405161024f9190611f56565b60405180910390f35b34801561026457600080fd5b5061027f600480360381019061027a919061194a565b610fc7565b60405161028c91906120a2565b60405180910390f35b3480156102a157600080fd5b506102aa6113bf565b6040516102b79190611e44565b60405180910390f35b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610328576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000341415610363576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006040518061010001604052807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000062ffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff16815260200160c84261043d9190612129565b8152602001348152602001848152602001600073ffffffffffffffffffffffffffffffffffffffff16815250905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663414bf38934846040518363ffffffff1660e01b81526004016104c7919061206b565b6020604051808303818588803b1580156104e057600080fd5b505af11580156104f4573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906105199190611a5e565b90507f87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da11348260405161054c9291906120bd565b60405180910390a1809250505092915050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614806105eb5750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b15610622576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600082141561065d576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61068a3330848673ffffffffffffffffffffffffffffffffffffffff166113e3909392919063ffffffff16565b6106d57f0000000000000000000000000000000000000000000000000000000000000000838573ffffffffffffffffffffffffffffffffffffffff1661146c9092919063ffffffff16565b60006040518060a00160405280857f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000060405160200161077b959493929190611db9565b60405160208183030381529060405281526020018773ffffffffffffffffffffffffffffffffffffffff16815260200160c8426107b89190612129565b815260200184815260200186815250905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663c04b8d59836040518263ffffffff1660e01b81526004016108249190612049565b602060405180830381600087803b15801561083e57600080fd5b505af1158015610852573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108769190611a5e565b90507f017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f8585836040516108ab93929190611f1f565b60405180910390a18092505050949350505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415610970576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008214156109ab576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109f83330847f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166113e3909392919063ffffffff16565b610a637f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661146c9092919063ffffffff16565b60006040518061010001604052807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000062ffffff1681526020013073ffffffffffffffffffffffffffffffffffffffff16815260200160c842610b3d9190612129565b8152602001848152602001858152602001600073ffffffffffffffffffffffffffffffffffffffff16815250905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663414bf389836040518263ffffffff1660e01b8152600401610bc6919061206b565b602060405180830381600087803b158015610be057600080fd5b505af1158015610bf4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c189190611a5e565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16632e1a7d4d826040518263ffffffff1660e01b8152600401610c7391906120a2565b600060405180830381600087803b158015610c8d57600080fd5b505af1158015610ca1573d6000803e3d6000fd5b505050507f74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae8482604051610cd69291906120bd565b60405180910390a160008673ffffffffffffffffffffffffffffffffffffffff1682604051610d0490611e2f565b60006040518083038185875af1925050503d8060008114610d41576040519150601f19603f3d011682016040523d82523d6000602084013e610d46565b606091505b5050905080610d81576040517f3794aeaf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8193505050509392505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b6000807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16631698ee827f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006040518463ffffffff1660e01b8152600401610e9693929190611e88565b60206040518083038186803b158015610eae57600080fd5b505afa158015610ec2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ee691906118dd565b9050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610f27576000915050610fc4565b600081905060008173ffffffffffffffffffffffffffffffffffffffff16631a6865026040518163ffffffff1660e01b815260040160206040518083038186803b158015610f7457600080fd5b505afa158015610f88573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fac9190611a31565b6fffffffffffffffffffffffffffffffff1611925050505b90565b60008060009054906101000a900460ff161561100f576040517f29f745a700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60016000806101000a81548160ff021916908315150217905550600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614806110905750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b156110c7576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000821415611102576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61114f3330847f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166113e3909392919063ffffffff16565b6111ba7f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661146c9092919063ffffffff16565b60006040518060a001604052807f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000089604051602001611260959493929190611db9565b60405160208183030381529060405281526020018773ffffffffffffffffffffffffffffffffffffffff16815260200160c84261129d9190612129565b815260200184815260200186815250905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663c04b8d59836040518263ffffffff1660e01b81526004016113099190612049565b602060405180830381600087803b15801561132357600080fd5b505af1158015611337573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061135b9190611a5e565b90507f0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b85858360405161139093929190611f1f565b60405180910390a1809250505060008060006101000a81548160ff021916908315150217905550949350505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b611466846323b872dd60e01b85858560405160240161140493929190611ebf565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506115ca565b50505050565b6000811480611505575060008373ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e30856040518363ffffffff1660e01b81526004016114b3929190611e5f565b60206040518083038186803b1580156114cb57600080fd5b505afa1580156114df573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115039190611a5e565b145b611544576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161153b90612029565b60405180910390fd5b6115c58363095ea7b360e01b8484604051602401611563929190611ef6565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506115ca565b505050565b600061162c826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166116919092919063ffffffff16565b905060008151111561168c578080602001905181019061164c9190611a04565b61168b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161168290612009565b60405180910390fd5b5b505050565b60606116a084846000856116a9565b90509392505050565b6060824710156116ee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116e590611fc9565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516117179190611e18565b60006040518083038185875af1925050503d8060008114611754576040519150601f19603f3d011682016040523d82523d6000602084013e611759565b606091505b509150915061176a87838387611776565b92505050949350505050565b606083156117d9576000835114156117d157611791856117ec565b6117d0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117c790611fe9565b60405180910390fd5b5b8290506117e4565b6117e3838361180f565b5b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6000825111156118225781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118569190611fa7565b60405180910390fd5b60008135905061186e8161241b565b92915050565b6000815190506118838161241b565b92915050565b60008151905061189881612432565b92915050565b6000815190506118ad81612449565b92915050565b6000813590506118c281612460565b92915050565b6000815190506118d781612460565b92915050565b6000602082840312156118f3576118f26122d2565b5b600061190184828501611874565b91505092915050565b60008060408385031215611921576119206122d2565b5b600061192f8582860161185f565b9250506020611940858286016118b3565b9150509250929050565b60008060008060808587031215611964576119636122d2565b5b60006119728782880161185f565b9450506020611983878288016118b3565b93505060406119948782880161185f565b92505060606119a5878288016118b3565b91505092959194509250565b6000806000606084860312156119ca576119c96122d2565b5b60006119d88682870161185f565b93505060206119e9868287016118b3565b92505060406119fa868287016118b3565b9150509250925092565b600060208284031215611a1a57611a196122d2565b5b6000611a2884828501611889565b91505092915050565b600060208284031215611a4757611a466122d2565b5b6000611a558482850161189e565b91505092915050565b600060208284031215611a7457611a736122d2565b5b6000611a82848285016118c8565b91505092915050565b611a948161217f565b82525050565b611aa38161217f565b82525050565b611aba611ab58261217f565b61226d565b82525050565b611ac981612191565b82525050565b6000611ada826120e6565b611ae481856120fc565b9350611af481856020860161223a565b611afd816122d7565b840191505092915050565b6000611b13826120e6565b611b1d818561210d565b9350611b2d81856020860161223a565b80840191505092915050565b611b42816121f2565b82525050565b611b5181612204565b82525050565b6000611b62826120f1565b611b6c8185612118565b9350611b7c81856020860161223a565b611b85816122d7565b840191505092915050565b6000611b9d602683612118565b9150611ba882612302565b604082019050919050565b6000611bc060008361210d565b9150611bcb82612351565b600082019050919050565b6000611be3601d83612118565b9150611bee82612354565b602082019050919050565b6000611c06602a83612118565b9150611c118261237d565b604082019050919050565b6000611c29603683612118565b9150611c34826123cc565b604082019050919050565b600060a0830160008301518482036000860152611c5c8282611acf565b9150506020830151611c716020860182611a8b565b506040830151611c846040860182611d9b565b506060830151611c976060860182611d9b565b506080830151611caa6080860182611d9b565b508091505092915050565b61010082016000820151611ccc6000850182611a8b565b506020820151611cdf6020850182611a8b565b506040820151611cf26040850182611d66565b506060820151611d056060850182611a8b565b506080820151611d186080850182611d9b565b5060a0820151611d2b60a0850182611d9b565b5060c0820151611d3e60c0850182611d9b565b5060e0820151611d5160e0850182611d57565b50505050565b611d60816121b9565b82525050565b611d6f816121d9565b82525050565b611d7e816121d9565b82525050565b611d95611d90826121d9565b612291565b82525050565b611da4816121e8565b82525050565b611db3816121e8565b82525050565b6000611dc58288611aa9565b601482019150611dd58287611d84565b600382019150611de58286611aa9565b601482019150611df58285611d84565b600382019150611e058284611aa9565b6014820191508190509695505050505050565b6000611e248284611b08565b915081905092915050565b6000611e3a82611bb3565b9150819050919050565b6000602082019050611e596000830184611a9a565b92915050565b6000604082019050611e746000830185611a9a565b611e816020830184611a9a565b9392505050565b6000606082019050611e9d6000830186611a9a565b611eaa6020830185611a9a565b611eb76040830184611d75565b949350505050565b6000606082019050611ed46000830186611a9a565b611ee16020830185611a9a565b611eee6040830184611daa565b949350505050565b6000604082019050611f0b6000830185611a9a565b611f186020830184611daa565b9392505050565b6000606082019050611f346000830186611a9a565b611f416020830185611daa565b611f4e6040830184611daa565b949350505050565b6000602082019050611f6b6000830184611ac0565b92915050565b6000602082019050611f866000830184611b39565b92915050565b6000602082019050611fa16000830184611b48565b92915050565b60006020820190508181036000830152611fc18184611b57565b905092915050565b60006020820190508181036000830152611fe281611b90565b9050919050565b6000602082019050818103600083015261200281611bd6565b9050919050565b6000602082019050818103600083015261202281611bf9565b9050919050565b6000602082019050818103600083015261204281611c1c565b9050919050565b600060208201905081810360008301526120638184611c3f565b905092915050565b6000610100820190506120816000830184611cb5565b92915050565b600060208201905061209c6000830184611d75565b92915050565b60006020820190506120b76000830184611daa565b92915050565b60006040820190506120d26000830185611daa565b6120df6020830184611daa565b9392505050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000612134826121e8565b915061213f836121e8565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115612174576121736122a3565b5b828201905092915050565b600061218a826121b9565b9050919050565b60008115159050919050565b60006fffffffffffffffffffffffffffffffff82169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600062ffffff82169050919050565b6000819050919050565b60006121fd82612216565b9050919050565b600061220f82612216565b9050919050565b600061222182612228565b9050919050565b6000612233826121b9565b9050919050565b60005b8381101561225857808201518184015260208101905061223d565b83811115612267576000848401525b50505050565b60006122788261227f565b9050919050565b600061228a826122f5565b9050919050565b600061229c826122e8565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600080fd5b6000601f19601f8301169050919050565b60008160e81b9050919050565b60008160601b9050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b50565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b7f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60008201527f20746f206e6f6e2d7a65726f20616c6c6f77616e636500000000000000000000602082015250565b6124248161217f565b811461242f57600080fd5b50565b61243b81612191565b811461244657600080fd5b50565b6124528161219d565b811461245d57600080fd5b50565b612469816121e8565b811461247457600080fd5b5056fea26469706673582212205b01f0993d1aca09a588ccda9fc1b26d546e37338ba85cdc92932152534aeee064736f6c63430008070033",
}

// ZetaTokenConsumerUniV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaTokenConsumerUniV3MetaData.ABI instead.
var ZetaTokenConsumerUniV3ABI = ZetaTokenConsumerUniV3MetaData.ABI

// ZetaTokenConsumerUniV3Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZetaTokenConsumerUniV3MetaData.Bin instead.
var ZetaTokenConsumerUniV3Bin = ZetaTokenConsumerUniV3MetaData.Bin

// DeployZetaTokenConsumerUniV3 deploys a new Ethereum contract, binding an instance of ZetaTokenConsumerUniV3 to it.
func DeployZetaTokenConsumerUniV3(auth *bind.TransactOpts, backend bind.ContractBackend, zetaToken_ common.Address, uniswapV3Router_ common.Address, uniswapV3Factory_ common.Address, WETH9Address_ common.Address, zetaPoolFee_ *big.Int, tokenPoolFee_ *big.Int) (common.Address, *types.Transaction, *ZetaTokenConsumerUniV3, error) {
	parsed, err := ZetaTokenConsumerUniV3MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZetaTokenConsumerUniV3Bin), backend, zetaToken_, uniswapV3Router_, uniswapV3Factory_, WETH9Address_, zetaPoolFee_, tokenPoolFee_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZetaTokenConsumerUniV3{ZetaTokenConsumerUniV3Caller: ZetaTokenConsumerUniV3Caller{contract: contract}, ZetaTokenConsumerUniV3Transactor: ZetaTokenConsumerUniV3Transactor{contract: contract}, ZetaTokenConsumerUniV3Filterer: ZetaTokenConsumerUniV3Filterer{contract: contract}}, nil
}

// ZetaTokenConsumerUniV3 is an auto generated Go binding around an Ethereum contract.
type ZetaTokenConsumerUniV3 struct {
	ZetaTokenConsumerUniV3Caller     // Read-only binding to the contract
	ZetaTokenConsumerUniV3Transactor // Write-only binding to the contract
	ZetaTokenConsumerUniV3Filterer   // Log filterer for contract events
}

// ZetaTokenConsumerUniV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaTokenConsumerUniV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerUniV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaTokenConsumerUniV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerUniV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaTokenConsumerUniV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerUniV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaTokenConsumerUniV3Session struct {
	Contract     *ZetaTokenConsumerUniV3 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ZetaTokenConsumerUniV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaTokenConsumerUniV3CallerSession struct {
	Contract *ZetaTokenConsumerUniV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ZetaTokenConsumerUniV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaTokenConsumerUniV3TransactorSession struct {
	Contract     *ZetaTokenConsumerUniV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ZetaTokenConsumerUniV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaTokenConsumerUniV3Raw struct {
	Contract *ZetaTokenConsumerUniV3 // Generic contract binding to access the raw methods on
}

// ZetaTokenConsumerUniV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaTokenConsumerUniV3CallerRaw struct {
	Contract *ZetaTokenConsumerUniV3Caller // Generic read-only contract binding to access the raw methods on
}

// ZetaTokenConsumerUniV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaTokenConsumerUniV3TransactorRaw struct {
	Contract *ZetaTokenConsumerUniV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaTokenConsumerUniV3 creates a new instance of ZetaTokenConsumerUniV3, bound to a specific deployed contract.
func NewZetaTokenConsumerUniV3(address common.Address, backend bind.ContractBackend) (*ZetaTokenConsumerUniV3, error) {
	contract, err := bindZetaTokenConsumerUniV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerUniV3{ZetaTokenConsumerUniV3Caller: ZetaTokenConsumerUniV3Caller{contract: contract}, ZetaTokenConsumerUniV3Transactor: ZetaTokenConsumerUniV3Transactor{contract: contract}, ZetaTokenConsumerUniV3Filterer: ZetaTokenConsumerUniV3Filterer{contract: contract}}, nil
}

// NewZetaTokenConsumerUniV3Caller creates a new read-only instance of ZetaTokenConsumerUniV3, bound to a specific deployed contract.
func NewZetaTokenConsumerUniV3Caller(address common.Address, caller bind.ContractCaller) (*ZetaTokenConsumerUniV3Caller, error) {
	contract, err := bindZetaTokenConsumerUniV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerUniV3Caller{contract: contract}, nil
}

// NewZetaTokenConsumerUniV3Transactor creates a new write-only instance of ZetaTokenConsumerUniV3, bound to a specific deployed contract.
func NewZetaTokenConsumerUniV3Transactor(address common.Address, transactor bind.ContractTransactor) (*ZetaTokenConsumerUniV3Transactor, error) {
	contract, err := bindZetaTokenConsumerUniV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerUniV3Transactor{contract: contract}, nil
}

// NewZetaTokenConsumerUniV3Filterer creates a new log filterer instance of ZetaTokenConsumerUniV3, bound to a specific deployed contract.
func NewZetaTokenConsumerUniV3Filterer(address common.Address, filterer bind.ContractFilterer) (*ZetaTokenConsumerUniV3Filterer, error) {
	contract, err := bindZetaTokenConsumerUniV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerUniV3Filterer{contract: contract}, nil
}

// bindZetaTokenConsumerUniV3 binds a generic wrapper to an already deployed contract.
func bindZetaTokenConsumerUniV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaTokenConsumerUniV3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaTokenConsumerUniV3.Contract.ZetaTokenConsumerUniV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.Contract.ZetaTokenConsumerUniV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.Contract.ZetaTokenConsumerUniV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaTokenConsumerUniV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.Contract.contract.Transact(opts, method, params...)
}

// WETH9Address is a free data retrieval call binding the contract method 0xc469cf14.
//
// Solidity: function WETH9Address() view returns(address)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Caller) WETH9Address(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaTokenConsumerUniV3.contract.Call(opts, &out, "WETH9Address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9Address is a free data retrieval call binding the contract method 0xc469cf14.
//
// Solidity: function WETH9Address() view returns(address)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Session) WETH9Address() (common.Address, error) {
	return _ZetaTokenConsumerUniV3.Contract.WETH9Address(&_ZetaTokenConsumerUniV3.CallOpts)
}

// WETH9Address is a free data retrieval call binding the contract method 0xc469cf14.
//
// Solidity: function WETH9Address() view returns(address)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3CallerSession) WETH9Address() (common.Address, error) {
	return _ZetaTokenConsumerUniV3.Contract.WETH9Address(&_ZetaTokenConsumerUniV3.CallOpts)
}

// HasZetaLiquidity is a free data retrieval call binding the contract method 0x80801f84.
//
// Solidity: function hasZetaLiquidity() view returns(bool)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Caller) HasZetaLiquidity(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZetaTokenConsumerUniV3.contract.Call(opts, &out, "hasZetaLiquidity")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasZetaLiquidity is a free data retrieval call binding the contract method 0x80801f84.
//
// Solidity: function hasZetaLiquidity() view returns(bool)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Session) HasZetaLiquidity() (bool, error) {
	return _ZetaTokenConsumerUniV3.Contract.HasZetaLiquidity(&_ZetaTokenConsumerUniV3.CallOpts)
}

// HasZetaLiquidity is a free data retrieval call binding the contract method 0x80801f84.
//
// Solidity: function hasZetaLiquidity() view returns(bool)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3CallerSession) HasZetaLiquidity() (bool, error) {
	return _ZetaTokenConsumerUniV3.Contract.HasZetaLiquidity(&_ZetaTokenConsumerUniV3.CallOpts)
}

// TokenPoolFee is a free data retrieval call binding the contract method 0x5d9dfdde.
//
// Solidity: function tokenPoolFee() view returns(uint24)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Caller) TokenPoolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZetaTokenConsumerUniV3.contract.Call(opts, &out, "tokenPoolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenPoolFee is a free data retrieval call binding the contract method 0x5d9dfdde.
//
// Solidity: function tokenPoolFee() view returns(uint24)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Session) TokenPoolFee() (*big.Int, error) {
	return _ZetaTokenConsumerUniV3.Contract.TokenPoolFee(&_ZetaTokenConsumerUniV3.CallOpts)
}

// TokenPoolFee is a free data retrieval call binding the contract method 0x5d9dfdde.
//
// Solidity: function tokenPoolFee() view returns(uint24)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3CallerSession) TokenPoolFee() (*big.Int, error) {
	return _ZetaTokenConsumerUniV3.Contract.TokenPoolFee(&_ZetaTokenConsumerUniV3.CallOpts)
}

// UniswapV3Factory is a free data retrieval call binding the contract method 0x5b549182.
//
// Solidity: function uniswapV3Factory() view returns(address)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Caller) UniswapV3Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaTokenConsumerUniV3.contract.Call(opts, &out, "uniswapV3Factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV3Factory is a free data retrieval call binding the contract method 0x5b549182.
//
// Solidity: function uniswapV3Factory() view returns(address)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Session) UniswapV3Factory() (common.Address, error) {
	return _ZetaTokenConsumerUniV3.Contract.UniswapV3Factory(&_ZetaTokenConsumerUniV3.CallOpts)
}

// UniswapV3Factory is a free data retrieval call binding the contract method 0x5b549182.
//
// Solidity: function uniswapV3Factory() view returns(address)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3CallerSession) UniswapV3Factory() (common.Address, error) {
	return _ZetaTokenConsumerUniV3.Contract.UniswapV3Factory(&_ZetaTokenConsumerUniV3.CallOpts)
}

// UniswapV3Router is a free data retrieval call binding the contract method 0x2c76d7a6.
//
// Solidity: function uniswapV3Router() view returns(address)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Caller) UniswapV3Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaTokenConsumerUniV3.contract.Call(opts, &out, "uniswapV3Router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV3Router is a free data retrieval call binding the contract method 0x2c76d7a6.
//
// Solidity: function uniswapV3Router() view returns(address)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Session) UniswapV3Router() (common.Address, error) {
	return _ZetaTokenConsumerUniV3.Contract.UniswapV3Router(&_ZetaTokenConsumerUniV3.CallOpts)
}

// UniswapV3Router is a free data retrieval call binding the contract method 0x2c76d7a6.
//
// Solidity: function uniswapV3Router() view returns(address)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3CallerSession) UniswapV3Router() (common.Address, error) {
	return _ZetaTokenConsumerUniV3.Contract.UniswapV3Router(&_ZetaTokenConsumerUniV3.CallOpts)
}

// ZetaPoolFee is a free data retrieval call binding the contract method 0x3cbd7005.
//
// Solidity: function zetaPoolFee() view returns(uint24)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Caller) ZetaPoolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZetaTokenConsumerUniV3.contract.Call(opts, &out, "zetaPoolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZetaPoolFee is a free data retrieval call binding the contract method 0x3cbd7005.
//
// Solidity: function zetaPoolFee() view returns(uint24)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Session) ZetaPoolFee() (*big.Int, error) {
	return _ZetaTokenConsumerUniV3.Contract.ZetaPoolFee(&_ZetaTokenConsumerUniV3.CallOpts)
}

// ZetaPoolFee is a free data retrieval call binding the contract method 0x3cbd7005.
//
// Solidity: function zetaPoolFee() view returns(uint24)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3CallerSession) ZetaPoolFee() (*big.Int, error) {
	return _ZetaTokenConsumerUniV3.Contract.ZetaPoolFee(&_ZetaTokenConsumerUniV3.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Caller) ZetaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaTokenConsumerUniV3.contract.Call(opts, &out, "zetaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Session) ZetaToken() (common.Address, error) {
	return _ZetaTokenConsumerUniV3.Contract.ZetaToken(&_ZetaTokenConsumerUniV3.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3CallerSession) ZetaToken() (common.Address, error) {
	return _ZetaTokenConsumerUniV3.Contract.ZetaToken(&_ZetaTokenConsumerUniV3.CallOpts)
}

// GetEthFromZeta is a paid mutator transaction binding the contract method 0x54c49a2a.
//
// Solidity: function getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Transactor) GetEthFromZeta(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.contract.Transact(opts, "getEthFromZeta", destinationAddress, minAmountOut, zetaTokenAmount)
}

// GetEthFromZeta is a paid mutator transaction binding the contract method 0x54c49a2a.
//
// Solidity: function getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Session) GetEthFromZeta(destinationAddress common.Address, minAmountOut *big.Int, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.Contract.GetEthFromZeta(&_ZetaTokenConsumerUniV3.TransactOpts, destinationAddress, minAmountOut, zetaTokenAmount)
}

// GetEthFromZeta is a paid mutator transaction binding the contract method 0x54c49a2a.
//
// Solidity: function getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3TransactorSession) GetEthFromZeta(destinationAddress common.Address, minAmountOut *big.Int, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.Contract.GetEthFromZeta(&_ZetaTokenConsumerUniV3.TransactOpts, destinationAddress, minAmountOut, zetaTokenAmount)
}

// GetTokenFromZeta is a paid mutator transaction binding the contract method 0xa53fb10b.
//
// Solidity: function getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Transactor) GetTokenFromZeta(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.contract.Transact(opts, "getTokenFromZeta", destinationAddress, minAmountOut, outputToken, zetaTokenAmount)
}

// GetTokenFromZeta is a paid mutator transaction binding the contract method 0xa53fb10b.
//
// Solidity: function getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Session) GetTokenFromZeta(destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.Contract.GetTokenFromZeta(&_ZetaTokenConsumerUniV3.TransactOpts, destinationAddress, minAmountOut, outputToken, zetaTokenAmount)
}

// GetTokenFromZeta is a paid mutator transaction binding the contract method 0xa53fb10b.
//
// Solidity: function getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3TransactorSession) GetTokenFromZeta(destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.Contract.GetTokenFromZeta(&_ZetaTokenConsumerUniV3.TransactOpts, destinationAddress, minAmountOut, outputToken, zetaTokenAmount)
}

// GetZetaFromEth is a paid mutator transaction binding the contract method 0x013b2ff8.
//
// Solidity: function getZetaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Transactor) GetZetaFromEth(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.contract.Transact(opts, "getZetaFromEth", destinationAddress, minAmountOut)
}

// GetZetaFromEth is a paid mutator transaction binding the contract method 0x013b2ff8.
//
// Solidity: function getZetaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Session) GetZetaFromEth(destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.Contract.GetZetaFromEth(&_ZetaTokenConsumerUniV3.TransactOpts, destinationAddress, minAmountOut)
}

// GetZetaFromEth is a paid mutator transaction binding the contract method 0x013b2ff8.
//
// Solidity: function getZetaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3TransactorSession) GetZetaFromEth(destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.Contract.GetZetaFromEth(&_ZetaTokenConsumerUniV3.TransactOpts, destinationAddress, minAmountOut)
}

// GetZetaFromToken is a paid mutator transaction binding the contract method 0x2405620a.
//
// Solidity: function getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Transactor) GetZetaFromToken(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.contract.Transact(opts, "getZetaFromToken", destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// GetZetaFromToken is a paid mutator transaction binding the contract method 0x2405620a.
//
// Solidity: function getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Session) GetZetaFromToken(destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.Contract.GetZetaFromToken(&_ZetaTokenConsumerUniV3.TransactOpts, destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// GetZetaFromToken is a paid mutator transaction binding the contract method 0x2405620a.
//
// Solidity: function getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3TransactorSession) GetZetaFromToken(destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.Contract.GetZetaFromToken(&_ZetaTokenConsumerUniV3.TransactOpts, destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Session) Receive() (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.Contract.Receive(&_ZetaTokenConsumerUniV3.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3TransactorSession) Receive() (*types.Transaction, error) {
	return _ZetaTokenConsumerUniV3.Contract.Receive(&_ZetaTokenConsumerUniV3.TransactOpts)
}

// ZetaTokenConsumerUniV3EthExchangedForZetaIterator is returned from FilterEthExchangedForZeta and is used to iterate over the raw logs and unpacked data for EthExchangedForZeta events raised by the ZetaTokenConsumerUniV3 contract.
type ZetaTokenConsumerUniV3EthExchangedForZetaIterator struct {
	Event *ZetaTokenConsumerUniV3EthExchangedForZeta // Event containing the contract specifics and raw log

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
func (it *ZetaTokenConsumerUniV3EthExchangedForZetaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerUniV3EthExchangedForZeta)
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
		it.Event = new(ZetaTokenConsumerUniV3EthExchangedForZeta)
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
func (it *ZetaTokenConsumerUniV3EthExchangedForZetaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerUniV3EthExchangedForZetaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerUniV3EthExchangedForZeta represents a EthExchangedForZeta event raised by the ZetaTokenConsumerUniV3 contract.
type ZetaTokenConsumerUniV3EthExchangedForZeta struct {
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEthExchangedForZeta is a free log retrieval operation binding the contract event 0x87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da11.
//
// Solidity: event EthExchangedForZeta(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Filterer) FilterEthExchangedForZeta(opts *bind.FilterOpts) (*ZetaTokenConsumerUniV3EthExchangedForZetaIterator, error) {

	logs, sub, err := _ZetaTokenConsumerUniV3.contract.FilterLogs(opts, "EthExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerUniV3EthExchangedForZetaIterator{contract: _ZetaTokenConsumerUniV3.contract, event: "EthExchangedForZeta", logs: logs, sub: sub}, nil
}

// WatchEthExchangedForZeta is a free log subscription operation binding the contract event 0x87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da11.
//
// Solidity: event EthExchangedForZeta(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Filterer) WatchEthExchangedForZeta(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerUniV3EthExchangedForZeta) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumerUniV3.contract.WatchLogs(opts, "EthExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerUniV3EthExchangedForZeta)
				if err := _ZetaTokenConsumerUniV3.contract.UnpackLog(event, "EthExchangedForZeta", log); err != nil {
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

// ParseEthExchangedForZeta is a log parse operation binding the contract event 0x87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da11.
//
// Solidity: event EthExchangedForZeta(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Filterer) ParseEthExchangedForZeta(log types.Log) (*ZetaTokenConsumerUniV3EthExchangedForZeta, error) {
	event := new(ZetaTokenConsumerUniV3EthExchangedForZeta)
	if err := _ZetaTokenConsumerUniV3.contract.UnpackLog(event, "EthExchangedForZeta", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaTokenConsumerUniV3TokenExchangedForZetaIterator is returned from FilterTokenExchangedForZeta and is used to iterate over the raw logs and unpacked data for TokenExchangedForZeta events raised by the ZetaTokenConsumerUniV3 contract.
type ZetaTokenConsumerUniV3TokenExchangedForZetaIterator struct {
	Event *ZetaTokenConsumerUniV3TokenExchangedForZeta // Event containing the contract specifics and raw log

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
func (it *ZetaTokenConsumerUniV3TokenExchangedForZetaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerUniV3TokenExchangedForZeta)
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
		it.Event = new(ZetaTokenConsumerUniV3TokenExchangedForZeta)
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
func (it *ZetaTokenConsumerUniV3TokenExchangedForZetaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerUniV3TokenExchangedForZetaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerUniV3TokenExchangedForZeta represents a TokenExchangedForZeta event raised by the ZetaTokenConsumerUniV3 contract.
type ZetaTokenConsumerUniV3TokenExchangedForZeta struct {
	Token     common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenExchangedForZeta is a free log retrieval operation binding the contract event 0x017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f.
//
// Solidity: event TokenExchangedForZeta(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Filterer) FilterTokenExchangedForZeta(opts *bind.FilterOpts) (*ZetaTokenConsumerUniV3TokenExchangedForZetaIterator, error) {

	logs, sub, err := _ZetaTokenConsumerUniV3.contract.FilterLogs(opts, "TokenExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerUniV3TokenExchangedForZetaIterator{contract: _ZetaTokenConsumerUniV3.contract, event: "TokenExchangedForZeta", logs: logs, sub: sub}, nil
}

// WatchTokenExchangedForZeta is a free log subscription operation binding the contract event 0x017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f.
//
// Solidity: event TokenExchangedForZeta(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Filterer) WatchTokenExchangedForZeta(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerUniV3TokenExchangedForZeta) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumerUniV3.contract.WatchLogs(opts, "TokenExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerUniV3TokenExchangedForZeta)
				if err := _ZetaTokenConsumerUniV3.contract.UnpackLog(event, "TokenExchangedForZeta", log); err != nil {
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

// ParseTokenExchangedForZeta is a log parse operation binding the contract event 0x017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f.
//
// Solidity: event TokenExchangedForZeta(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Filterer) ParseTokenExchangedForZeta(log types.Log) (*ZetaTokenConsumerUniV3TokenExchangedForZeta, error) {
	event := new(ZetaTokenConsumerUniV3TokenExchangedForZeta)
	if err := _ZetaTokenConsumerUniV3.contract.UnpackLog(event, "TokenExchangedForZeta", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaTokenConsumerUniV3ZetaExchangedForEthIterator is returned from FilterZetaExchangedForEth and is used to iterate over the raw logs and unpacked data for ZetaExchangedForEth events raised by the ZetaTokenConsumerUniV3 contract.
type ZetaTokenConsumerUniV3ZetaExchangedForEthIterator struct {
	Event *ZetaTokenConsumerUniV3ZetaExchangedForEth // Event containing the contract specifics and raw log

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
func (it *ZetaTokenConsumerUniV3ZetaExchangedForEthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerUniV3ZetaExchangedForEth)
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
		it.Event = new(ZetaTokenConsumerUniV3ZetaExchangedForEth)
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
func (it *ZetaTokenConsumerUniV3ZetaExchangedForEthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerUniV3ZetaExchangedForEthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerUniV3ZetaExchangedForEth represents a ZetaExchangedForEth event raised by the ZetaTokenConsumerUniV3 contract.
type ZetaTokenConsumerUniV3ZetaExchangedForEth struct {
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterZetaExchangedForEth is a free log retrieval operation binding the contract event 0x74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae.
//
// Solidity: event ZetaExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Filterer) FilterZetaExchangedForEth(opts *bind.FilterOpts) (*ZetaTokenConsumerUniV3ZetaExchangedForEthIterator, error) {

	logs, sub, err := _ZetaTokenConsumerUniV3.contract.FilterLogs(opts, "ZetaExchangedForEth")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerUniV3ZetaExchangedForEthIterator{contract: _ZetaTokenConsumerUniV3.contract, event: "ZetaExchangedForEth", logs: logs, sub: sub}, nil
}

// WatchZetaExchangedForEth is a free log subscription operation binding the contract event 0x74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae.
//
// Solidity: event ZetaExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Filterer) WatchZetaExchangedForEth(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerUniV3ZetaExchangedForEth) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumerUniV3.contract.WatchLogs(opts, "ZetaExchangedForEth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerUniV3ZetaExchangedForEth)
				if err := _ZetaTokenConsumerUniV3.contract.UnpackLog(event, "ZetaExchangedForEth", log); err != nil {
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

// ParseZetaExchangedForEth is a log parse operation binding the contract event 0x74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae.
//
// Solidity: event ZetaExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Filterer) ParseZetaExchangedForEth(log types.Log) (*ZetaTokenConsumerUniV3ZetaExchangedForEth, error) {
	event := new(ZetaTokenConsumerUniV3ZetaExchangedForEth)
	if err := _ZetaTokenConsumerUniV3.contract.UnpackLog(event, "ZetaExchangedForEth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaTokenConsumerUniV3ZetaExchangedForTokenIterator is returned from FilterZetaExchangedForToken and is used to iterate over the raw logs and unpacked data for ZetaExchangedForToken events raised by the ZetaTokenConsumerUniV3 contract.
type ZetaTokenConsumerUniV3ZetaExchangedForTokenIterator struct {
	Event *ZetaTokenConsumerUniV3ZetaExchangedForToken // Event containing the contract specifics and raw log

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
func (it *ZetaTokenConsumerUniV3ZetaExchangedForTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerUniV3ZetaExchangedForToken)
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
		it.Event = new(ZetaTokenConsumerUniV3ZetaExchangedForToken)
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
func (it *ZetaTokenConsumerUniV3ZetaExchangedForTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerUniV3ZetaExchangedForTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerUniV3ZetaExchangedForToken represents a ZetaExchangedForToken event raised by the ZetaTokenConsumerUniV3 contract.
type ZetaTokenConsumerUniV3ZetaExchangedForToken struct {
	Token     common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterZetaExchangedForToken is a free log retrieval operation binding the contract event 0x0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b.
//
// Solidity: event ZetaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Filterer) FilterZetaExchangedForToken(opts *bind.FilterOpts) (*ZetaTokenConsumerUniV3ZetaExchangedForTokenIterator, error) {

	logs, sub, err := _ZetaTokenConsumerUniV3.contract.FilterLogs(opts, "ZetaExchangedForToken")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerUniV3ZetaExchangedForTokenIterator{contract: _ZetaTokenConsumerUniV3.contract, event: "ZetaExchangedForToken", logs: logs, sub: sub}, nil
}

// WatchZetaExchangedForToken is a free log subscription operation binding the contract event 0x0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b.
//
// Solidity: event ZetaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Filterer) WatchZetaExchangedForToken(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerUniV3ZetaExchangedForToken) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumerUniV3.contract.WatchLogs(opts, "ZetaExchangedForToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerUniV3ZetaExchangedForToken)
				if err := _ZetaTokenConsumerUniV3.contract.UnpackLog(event, "ZetaExchangedForToken", log); err != nil {
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

// ParseZetaExchangedForToken is a log parse operation binding the contract event 0x0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b.
//
// Solidity: event ZetaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerUniV3 *ZetaTokenConsumerUniV3Filterer) ParseZetaExchangedForToken(log types.Log) (*ZetaTokenConsumerUniV3ZetaExchangedForToken, error) {
	event := new(ZetaTokenConsumerUniV3ZetaExchangedForToken)
	if err := _ZetaTokenConsumerUniV3.contract.UnpackLog(event, "ZetaExchangedForToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
