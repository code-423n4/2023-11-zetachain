// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20custody

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

// ERC20CustodyMetaData contains all meta data concerning the ERC20Custody contract.
var ERC20CustodyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"TSSAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"TSSAddressUpdater_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"zetaFee_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zetaMaxFee_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"zeta_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTSSUpdater\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IsPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZetaMaxFeeExceeded\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"TSSAddressUpdater_\",\"type\":\"address\"}],\"name\":\"RenouncedTSSUpdater\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"Unwhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"TSSAddress_\",\"type\":\"address\"}],\"name\":\"UpdatedTSSAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"zetaFee_\",\"type\":\"uint256\"}],\"name\":\"UpdatedZetaFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"Whitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"TSSAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TSSAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceTSSAddressUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"unwhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"TSSAddress_\",\"type\":\"address\"}],\"name\":\"updateTSSAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zetaFee_\",\"type\":\"uint256\"}],\"name\":\"updateZetaFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zeta\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zetaFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zetaMaxFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b50604051620021a0380380620021a0833981810160405281019062000037919062000156565b6001600081905550846001806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550826003819055508073ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1660601b815250508160808181525050505050505062000283565b600081519050620001228162000235565b92915050565b60008151905062000139816200024f565b92915050565b600081519050620001508162000269565b92915050565b600080600080600060a0868803121562000175576200017462000230565b5b6000620001858882890162000111565b9550506020620001988882890162000111565b9450506040620001ab888289016200013f565b9350506060620001be888289016200013f565b9250506080620001d18882890162000128565b9150509295509295909350565b6000620001eb8262000206565b9050919050565b6000620001ff82620001de565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600080fd5b6200024081620001de565b81146200024c57600080fd5b50565b6200025a81620001f2565b81146200026657600080fd5b50565b620002748162000226565b81146200028057600080fd5b50565b60805160a05160601c611edf620002c160003960008181610dca01528181610e31015261102d01526000818161042d0152610c310152611edf6000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80639b19251a11610097578063e5408cfa11610066578063e5408cfa1461024d578063e609055e1461026b578063e8f9cb3a14610287578063ed11692b146102a557610100565b80639b19251a146101c9578063d936547e146101e5578063d9caed1214610215578063de2f6c5e1461023157610100565b80637bdaded3116100d35780637bdaded3146101695780638456cb5914610187578063950837aa146101915780639a590427146101ad57610100565b80633f4ba83a1461010557806353ee30a31461010f57806354b61e811461012d5780635c975abb1461014b575b600080fd5b61010d6102af565b005b6101176103ce565b6040516101249190611a04565b60405180910390f35b6101356103f2565b6040516101429190611a04565b60405180910390f35b610153610418565b6040516101609190611a7f565b60405180910390f35b61017161042b565b60405161017e9190611ba0565b60405180910390f35b61018f61044f565b005b6101ab60048036038101906101a691906116b3565b6105f5565b005b6101c760048036038101906101c29190611807565b61075c565b005b6101e360048036038101906101de9190611807565b61087f565b005b6101ff60048036038101906101fa9190611807565b6109a2565b60405161020c9190611a7f565b60405180910390f35b61022f600480360381019061022a91906116e0565b6109c2565b005b61024b60048036038101906102469190611834565b610b6f565b005b610255610cca565b6040516102629190611ba0565b60405180910390f35b61028560048036038101906102809190611760565b610cd0565b005b61028f61102b565b60405161029c9190611ae3565b60405180910390f35b6102ad61104f565b005b60018054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610334576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600160009054906101000a900460ff1661037a576040517f6cd6020100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600160006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa336040516103c49190611a04565b60405180910390a1565b60018054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900460ff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b60018054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104d4576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600160009054906101000a900460ff161561051b576040517f1309a56300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff1660018054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156105a2576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018060006101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258336040516105eb9190611a04565b60405180910390a1565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461067b576040517e611fa600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156106e2576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806001806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fcd2958db8285a532edf298cbe1aa28ea3fb5ec82461253f9a8c1477924e29947816040516107519190611a04565b60405180910390a150565b60018054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107e1576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508073ffffffffffffffffffffffffffffffffffffffff167f51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da4679160405160405180910390a250565b60018054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610904576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508073ffffffffffffffffffffffffffffffffffffffff167faab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a5460405160405180910390a250565b60046020528060005260406000206000915054906101000a900460ff1681565b6109ca6111f6565b60018054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a4f576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610ad2576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610afd83828473ffffffffffffffffffffffffffffffffffffffff166112469092919063ffffffff16565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb83604051610b5a9190611ba0565b60405180910390a3610b6a6112cc565b505050565b60018054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610bf4576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000811415610c2f576040517faf13986d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f0000000000000000000000000000000000000000000000000000000000000000811115610c89576040517fc1be451300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806003819055507f6d2d8e313fbaf76898bb9fa55e4b52525e49c7d7182d0874f97bd9076e81d52381604051610cbf9190611ba0565b60405180910390a150565b60035481565b610cd86111f6565b600160009054906101000a900460ff1615610d1f576040517f1309a56300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610da2576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060035414158015610e025750600073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1614155b15610e7757610e763360018054906101000a900473ffffffffffffffffffffffffffffffffffffffff166003547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166112d6909392919063ffffffff16565b5b60008473ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401610eb29190611a04565b60206040518083038186803b158015610eca57600080fd5b505afa158015610ede573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f029190611861565b9050610f313330868873ffffffffffffffffffffffffffffffffffffffff166112d6909392919063ffffffff16565b8473ffffffffffffffffffffffffffffffffffffffff167f1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae8888848973ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401610fa59190611a04565b60206040518083038186803b158015610fbd57600080fd5b505afa158015610fd1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ff59190611861565b610fff9190611bfe565b8787604051611012959493929190611a9a565b60405180910390a2506110236112cc565b505050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146110d5576040517e611fa600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff1660018054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561115c576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f39ac266bfc011581be62c138d96e4e8782812013bb66fffb4cd207f4bfee3777336040516111ec9190611a04565b60405180910390a1565b6002600054141561123c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161123390611b80565b60405180910390fd5b6002600081905550565b6112c78363a9059cbb60e01b8484604051602401611265929190611a56565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061135f565b505050565b6001600081905550565b611359846323b872dd60e01b8585856040516024016112f793929190611a1f565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061135f565b50505050565b60006113c1826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166114269092919063ffffffff16565b905060008151111561142157808060200190518101906113e19190611733565b611420576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161141790611b60565b60405180910390fd5b5b505050565b6060611435848460008561143e565b90509392505050565b606082471015611483576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161147a90611b20565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516114ac91906119ed565b60006040518083038185875af1925050503d80600081146114e9576040519150601f19603f3d011682016040523d82523d6000602084013e6114ee565b606091505b50915091506114ff8783838761150b565b92505050949350505050565b6060831561156e576000835114156115665761152685611581565b611565576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161155c90611b40565b60405180910390fd5b5b829050611579565b61157883836115a4565b5b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6000825111156115b75781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115eb9190611afe565b60405180910390fd5b60008135905061160381611e4d565b92915050565b60008151905061161881611e64565b92915050565b60008083601f84011261163457611633611d38565b5b8235905067ffffffffffffffff81111561165157611650611d33565b5b60208301915083600182028301111561166d5761166c611d3d565b5b9250929050565b60008135905061168381611e7b565b92915050565b60008135905061169881611e92565b92915050565b6000815190506116ad81611e92565b92915050565b6000602082840312156116c9576116c8611d47565b5b60006116d7848285016115f4565b91505092915050565b6000806000606084860312156116f9576116f8611d47565b5b6000611707868287016115f4565b935050602061171886828701611674565b925050604061172986828701611689565b9150509250925092565b60006020828403121561174957611748611d47565b5b600061175784828501611609565b91505092915050565b6000806000806000806080878903121561177d5761177c611d47565b5b600087013567ffffffffffffffff81111561179b5761179a611d42565b5b6117a789828a0161161e565b965096505060206117ba89828a01611674565b94505060406117cb89828a01611689565b935050606087013567ffffffffffffffff8111156117ec576117eb611d42565b5b6117f889828a0161161e565b92509250509295509295509295565b60006020828403121561181d5761181c611d47565b5b600061182b84828501611674565b91505092915050565b60006020828403121561184a57611849611d47565b5b600061185884828501611689565b91505092915050565b60006020828403121561187757611876611d47565b5b60006118858482850161169e565b91505092915050565b61189781611c32565b82525050565b6118a681611c44565b82525050565b60006118b88385611bd1565b93506118c5838584611cc2565b6118ce83611d4c565b840190509392505050565b60006118e482611bbb565b6118ee8185611be2565b93506118fe818560208601611cd1565b80840191505092915050565b61191381611c8c565b82525050565b600061192482611bc6565b61192e8185611bed565b935061193e818560208601611cd1565b61194781611d4c565b840191505092915050565b600061195f602683611bed565b915061196a82611d5d565b604082019050919050565b6000611982601d83611bed565b915061198d82611dac565b602082019050919050565b60006119a5602a83611bed565b91506119b082611dd5565b604082019050919050565b60006119c8601f83611bed565b91506119d382611e24565b602082019050919050565b6119e781611c82565b82525050565b60006119f982846118d9565b915081905092915050565b6000602082019050611a19600083018461188e565b92915050565b6000606082019050611a34600083018661188e565b611a41602083018561188e565b611a4e60408301846119de565b949350505050565b6000604082019050611a6b600083018561188e565b611a7860208301846119de565b9392505050565b6000602082019050611a94600083018461189d565b92915050565b60006060820190508181036000830152611ab58187896118ac565b9050611ac460208301866119de565b8181036040830152611ad78184866118ac565b90509695505050505050565b6000602082019050611af8600083018461190a565b92915050565b60006020820190508181036000830152611b188184611919565b905092915050565b60006020820190508181036000830152611b3981611952565b9050919050565b60006020820190508181036000830152611b5981611975565b9050919050565b60006020820190508181036000830152611b7981611998565b9050919050565b60006020820190508181036000830152611b99816119bb565b9050919050565b6000602082019050611bb560008301846119de565b92915050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000611c0982611c82565b9150611c1483611c82565b925082821015611c2757611c26611d04565b5b828203905092915050565b6000611c3d82611c62565b9050919050565b60008115159050919050565b6000611c5b82611c32565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000611c9782611c9e565b9050919050565b6000611ca982611cb0565b9050919050565b6000611cbb82611c62565b9050919050565b82818337600083830152505050565b60005b83811015611cef578082015181840152602081019050611cd4565b83811115611cfe576000848401525b50505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b611e5681611c32565b8114611e6157600080fd5b50565b611e6d81611c44565b8114611e7857600080fd5b50565b611e8481611c50565b8114611e8f57600080fd5b50565b611e9b81611c82565b8114611ea657600080fd5b5056fea26469706673582212205768544075d85a56b1d380f8fcc0c8829b8a656c656277c25c81c31608d9e4ef64736f6c63430008070033",
}

// ERC20CustodyABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20CustodyMetaData.ABI instead.
var ERC20CustodyABI = ERC20CustodyMetaData.ABI

// ERC20CustodyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20CustodyMetaData.Bin instead.
var ERC20CustodyBin = ERC20CustodyMetaData.Bin

// DeployERC20Custody deploys a new Ethereum contract, binding an instance of ERC20Custody to it.
func DeployERC20Custody(auth *bind.TransactOpts, backend bind.ContractBackend, TSSAddress_ common.Address, TSSAddressUpdater_ common.Address, zetaFee_ *big.Int, zetaMaxFee_ *big.Int, zeta_ common.Address) (common.Address, *types.Transaction, *ERC20Custody, error) {
	parsed, err := ERC20CustodyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20CustodyBin), backend, TSSAddress_, TSSAddressUpdater_, zetaFee_, zetaMaxFee_, zeta_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Custody{ERC20CustodyCaller: ERC20CustodyCaller{contract: contract}, ERC20CustodyTransactor: ERC20CustodyTransactor{contract: contract}, ERC20CustodyFilterer: ERC20CustodyFilterer{contract: contract}}, nil
}

// ERC20Custody is an auto generated Go binding around an Ethereum contract.
type ERC20Custody struct {
	ERC20CustodyCaller     // Read-only binding to the contract
	ERC20CustodyTransactor // Write-only binding to the contract
	ERC20CustodyFilterer   // Log filterer for contract events
}

// ERC20CustodyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20CustodyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20CustodyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20CustodyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20CustodyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20CustodyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20CustodySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20CustodySession struct {
	Contract     *ERC20Custody     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CustodyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CustodyCallerSession struct {
	Contract *ERC20CustodyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ERC20CustodyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20CustodyTransactorSession struct {
	Contract     *ERC20CustodyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ERC20CustodyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20CustodyRaw struct {
	Contract *ERC20Custody // Generic contract binding to access the raw methods on
}

// ERC20CustodyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CustodyCallerRaw struct {
	Contract *ERC20CustodyCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20CustodyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20CustodyTransactorRaw struct {
	Contract *ERC20CustodyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Custody creates a new instance of ERC20Custody, bound to a specific deployed contract.
func NewERC20Custody(address common.Address, backend bind.ContractBackend) (*ERC20Custody, error) {
	contract, err := bindERC20Custody(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Custody{ERC20CustodyCaller: ERC20CustodyCaller{contract: contract}, ERC20CustodyTransactor: ERC20CustodyTransactor{contract: contract}, ERC20CustodyFilterer: ERC20CustodyFilterer{contract: contract}}, nil
}

// NewERC20CustodyCaller creates a new read-only instance of ERC20Custody, bound to a specific deployed contract.
func NewERC20CustodyCaller(address common.Address, caller bind.ContractCaller) (*ERC20CustodyCaller, error) {
	contract, err := bindERC20Custody(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyCaller{contract: contract}, nil
}

// NewERC20CustodyTransactor creates a new write-only instance of ERC20Custody, bound to a specific deployed contract.
func NewERC20CustodyTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20CustodyTransactor, error) {
	contract, err := bindERC20Custody(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTransactor{contract: contract}, nil
}

// NewERC20CustodyFilterer creates a new log filterer instance of ERC20Custody, bound to a specific deployed contract.
func NewERC20CustodyFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20CustodyFilterer, error) {
	contract, err := bindERC20Custody(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyFilterer{contract: contract}, nil
}

// bindERC20Custody binds a generic wrapper to an already deployed contract.
func bindERC20Custody(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20CustodyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Custody *ERC20CustodyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Custody.Contract.ERC20CustodyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Custody *ERC20CustodyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Custody.Contract.ERC20CustodyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Custody *ERC20CustodyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Custody.Contract.ERC20CustodyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Custody *ERC20CustodyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Custody.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Custody *ERC20CustodyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Custody.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Custody *ERC20CustodyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Custody.Contract.contract.Transact(opts, method, params...)
}

// TSSAddress is a free data retrieval call binding the contract method 0x53ee30a3.
//
// Solidity: function TSSAddress() view returns(address)
func (_ERC20Custody *ERC20CustodyCaller) TSSAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Custody.contract.Call(opts, &out, "TSSAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TSSAddress is a free data retrieval call binding the contract method 0x53ee30a3.
//
// Solidity: function TSSAddress() view returns(address)
func (_ERC20Custody *ERC20CustodySession) TSSAddress() (common.Address, error) {
	return _ERC20Custody.Contract.TSSAddress(&_ERC20Custody.CallOpts)
}

// TSSAddress is a free data retrieval call binding the contract method 0x53ee30a3.
//
// Solidity: function TSSAddress() view returns(address)
func (_ERC20Custody *ERC20CustodyCallerSession) TSSAddress() (common.Address, error) {
	return _ERC20Custody.Contract.TSSAddress(&_ERC20Custody.CallOpts)
}

// TSSAddressUpdater is a free data retrieval call binding the contract method 0x54b61e81.
//
// Solidity: function TSSAddressUpdater() view returns(address)
func (_ERC20Custody *ERC20CustodyCaller) TSSAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Custody.contract.Call(opts, &out, "TSSAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TSSAddressUpdater is a free data retrieval call binding the contract method 0x54b61e81.
//
// Solidity: function TSSAddressUpdater() view returns(address)
func (_ERC20Custody *ERC20CustodySession) TSSAddressUpdater() (common.Address, error) {
	return _ERC20Custody.Contract.TSSAddressUpdater(&_ERC20Custody.CallOpts)
}

// TSSAddressUpdater is a free data retrieval call binding the contract method 0x54b61e81.
//
// Solidity: function TSSAddressUpdater() view returns(address)
func (_ERC20Custody *ERC20CustodyCallerSession) TSSAddressUpdater() (common.Address, error) {
	return _ERC20Custody.Contract.TSSAddressUpdater(&_ERC20Custody.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20Custody *ERC20CustodyCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20Custody.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20Custody *ERC20CustodySession) Paused() (bool, error) {
	return _ERC20Custody.Contract.Paused(&_ERC20Custody.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20Custody *ERC20CustodyCallerSession) Paused() (bool, error) {
	return _ERC20Custody.Contract.Paused(&_ERC20Custody.CallOpts)
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_ERC20Custody *ERC20CustodyCaller) Whitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20Custody.contract.Call(opts, &out, "whitelisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_ERC20Custody *ERC20CustodySession) Whitelisted(arg0 common.Address) (bool, error) {
	return _ERC20Custody.Contract.Whitelisted(&_ERC20Custody.CallOpts, arg0)
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_ERC20Custody *ERC20CustodyCallerSession) Whitelisted(arg0 common.Address) (bool, error) {
	return _ERC20Custody.Contract.Whitelisted(&_ERC20Custody.CallOpts, arg0)
}

// Zeta is a free data retrieval call binding the contract method 0xe8f9cb3a.
//
// Solidity: function zeta() view returns(address)
func (_ERC20Custody *ERC20CustodyCaller) Zeta(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Custody.contract.Call(opts, &out, "zeta")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Zeta is a free data retrieval call binding the contract method 0xe8f9cb3a.
//
// Solidity: function zeta() view returns(address)
func (_ERC20Custody *ERC20CustodySession) Zeta() (common.Address, error) {
	return _ERC20Custody.Contract.Zeta(&_ERC20Custody.CallOpts)
}

// Zeta is a free data retrieval call binding the contract method 0xe8f9cb3a.
//
// Solidity: function zeta() view returns(address)
func (_ERC20Custody *ERC20CustodyCallerSession) Zeta() (common.Address, error) {
	return _ERC20Custody.Contract.Zeta(&_ERC20Custody.CallOpts)
}

// ZetaFee is a free data retrieval call binding the contract method 0xe5408cfa.
//
// Solidity: function zetaFee() view returns(uint256)
func (_ERC20Custody *ERC20CustodyCaller) ZetaFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Custody.contract.Call(opts, &out, "zetaFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZetaFee is a free data retrieval call binding the contract method 0xe5408cfa.
//
// Solidity: function zetaFee() view returns(uint256)
func (_ERC20Custody *ERC20CustodySession) ZetaFee() (*big.Int, error) {
	return _ERC20Custody.Contract.ZetaFee(&_ERC20Custody.CallOpts)
}

// ZetaFee is a free data retrieval call binding the contract method 0xe5408cfa.
//
// Solidity: function zetaFee() view returns(uint256)
func (_ERC20Custody *ERC20CustodyCallerSession) ZetaFee() (*big.Int, error) {
	return _ERC20Custody.Contract.ZetaFee(&_ERC20Custody.CallOpts)
}

// ZetaMaxFee is a free data retrieval call binding the contract method 0x7bdaded3.
//
// Solidity: function zetaMaxFee() view returns(uint256)
func (_ERC20Custody *ERC20CustodyCaller) ZetaMaxFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Custody.contract.Call(opts, &out, "zetaMaxFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZetaMaxFee is a free data retrieval call binding the contract method 0x7bdaded3.
//
// Solidity: function zetaMaxFee() view returns(uint256)
func (_ERC20Custody *ERC20CustodySession) ZetaMaxFee() (*big.Int, error) {
	return _ERC20Custody.Contract.ZetaMaxFee(&_ERC20Custody.CallOpts)
}

// ZetaMaxFee is a free data retrieval call binding the contract method 0x7bdaded3.
//
// Solidity: function zetaMaxFee() view returns(uint256)
func (_ERC20Custody *ERC20CustodyCallerSession) ZetaMaxFee() (*big.Int, error) {
	return _ERC20Custody.Contract.ZetaMaxFee(&_ERC20Custody.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xe609055e.
//
// Solidity: function deposit(bytes recipient, address asset, uint256 amount, bytes message) returns()
func (_ERC20Custody *ERC20CustodyTransactor) Deposit(opts *bind.TransactOpts, recipient []byte, asset common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _ERC20Custody.contract.Transact(opts, "deposit", recipient, asset, amount, message)
}

// Deposit is a paid mutator transaction binding the contract method 0xe609055e.
//
// Solidity: function deposit(bytes recipient, address asset, uint256 amount, bytes message) returns()
func (_ERC20Custody *ERC20CustodySession) Deposit(recipient []byte, asset common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _ERC20Custody.Contract.Deposit(&_ERC20Custody.TransactOpts, recipient, asset, amount, message)
}

// Deposit is a paid mutator transaction binding the contract method 0xe609055e.
//
// Solidity: function deposit(bytes recipient, address asset, uint256 amount, bytes message) returns()
func (_ERC20Custody *ERC20CustodyTransactorSession) Deposit(recipient []byte, asset common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _ERC20Custody.Contract.Deposit(&_ERC20Custody.TransactOpts, recipient, asset, amount, message)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20Custody *ERC20CustodyTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Custody.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20Custody *ERC20CustodySession) Pause() (*types.Transaction, error) {
	return _ERC20Custody.Contract.Pause(&_ERC20Custody.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20Custody *ERC20CustodyTransactorSession) Pause() (*types.Transaction, error) {
	return _ERC20Custody.Contract.Pause(&_ERC20Custody.TransactOpts)
}

// RenounceTSSAddressUpdater is a paid mutator transaction binding the contract method 0xed11692b.
//
// Solidity: function renounceTSSAddressUpdater() returns()
func (_ERC20Custody *ERC20CustodyTransactor) RenounceTSSAddressUpdater(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Custody.contract.Transact(opts, "renounceTSSAddressUpdater")
}

// RenounceTSSAddressUpdater is a paid mutator transaction binding the contract method 0xed11692b.
//
// Solidity: function renounceTSSAddressUpdater() returns()
func (_ERC20Custody *ERC20CustodySession) RenounceTSSAddressUpdater() (*types.Transaction, error) {
	return _ERC20Custody.Contract.RenounceTSSAddressUpdater(&_ERC20Custody.TransactOpts)
}

// RenounceTSSAddressUpdater is a paid mutator transaction binding the contract method 0xed11692b.
//
// Solidity: function renounceTSSAddressUpdater() returns()
func (_ERC20Custody *ERC20CustodyTransactorSession) RenounceTSSAddressUpdater() (*types.Transaction, error) {
	return _ERC20Custody.Contract.RenounceTSSAddressUpdater(&_ERC20Custody.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC20Custody *ERC20CustodyTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Custody.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC20Custody *ERC20CustodySession) Unpause() (*types.Transaction, error) {
	return _ERC20Custody.Contract.Unpause(&_ERC20Custody.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC20Custody *ERC20CustodyTransactorSession) Unpause() (*types.Transaction, error) {
	return _ERC20Custody.Contract.Unpause(&_ERC20Custody.TransactOpts)
}

// Unwhitelist is a paid mutator transaction binding the contract method 0x9a590427.
//
// Solidity: function unwhitelist(address asset) returns()
func (_ERC20Custody *ERC20CustodyTransactor) Unwhitelist(opts *bind.TransactOpts, asset common.Address) (*types.Transaction, error) {
	return _ERC20Custody.contract.Transact(opts, "unwhitelist", asset)
}

// Unwhitelist is a paid mutator transaction binding the contract method 0x9a590427.
//
// Solidity: function unwhitelist(address asset) returns()
func (_ERC20Custody *ERC20CustodySession) Unwhitelist(asset common.Address) (*types.Transaction, error) {
	return _ERC20Custody.Contract.Unwhitelist(&_ERC20Custody.TransactOpts, asset)
}

// Unwhitelist is a paid mutator transaction binding the contract method 0x9a590427.
//
// Solidity: function unwhitelist(address asset) returns()
func (_ERC20Custody *ERC20CustodyTransactorSession) Unwhitelist(asset common.Address) (*types.Transaction, error) {
	return _ERC20Custody.Contract.Unwhitelist(&_ERC20Custody.TransactOpts, asset)
}

// UpdateTSSAddress is a paid mutator transaction binding the contract method 0x950837aa.
//
// Solidity: function updateTSSAddress(address TSSAddress_) returns()
func (_ERC20Custody *ERC20CustodyTransactor) UpdateTSSAddress(opts *bind.TransactOpts, TSSAddress_ common.Address) (*types.Transaction, error) {
	return _ERC20Custody.contract.Transact(opts, "updateTSSAddress", TSSAddress_)
}

// UpdateTSSAddress is a paid mutator transaction binding the contract method 0x950837aa.
//
// Solidity: function updateTSSAddress(address TSSAddress_) returns()
func (_ERC20Custody *ERC20CustodySession) UpdateTSSAddress(TSSAddress_ common.Address) (*types.Transaction, error) {
	return _ERC20Custody.Contract.UpdateTSSAddress(&_ERC20Custody.TransactOpts, TSSAddress_)
}

// UpdateTSSAddress is a paid mutator transaction binding the contract method 0x950837aa.
//
// Solidity: function updateTSSAddress(address TSSAddress_) returns()
func (_ERC20Custody *ERC20CustodyTransactorSession) UpdateTSSAddress(TSSAddress_ common.Address) (*types.Transaction, error) {
	return _ERC20Custody.Contract.UpdateTSSAddress(&_ERC20Custody.TransactOpts, TSSAddress_)
}

// UpdateZetaFee is a paid mutator transaction binding the contract method 0xde2f6c5e.
//
// Solidity: function updateZetaFee(uint256 zetaFee_) returns()
func (_ERC20Custody *ERC20CustodyTransactor) UpdateZetaFee(opts *bind.TransactOpts, zetaFee_ *big.Int) (*types.Transaction, error) {
	return _ERC20Custody.contract.Transact(opts, "updateZetaFee", zetaFee_)
}

// UpdateZetaFee is a paid mutator transaction binding the contract method 0xde2f6c5e.
//
// Solidity: function updateZetaFee(uint256 zetaFee_) returns()
func (_ERC20Custody *ERC20CustodySession) UpdateZetaFee(zetaFee_ *big.Int) (*types.Transaction, error) {
	return _ERC20Custody.Contract.UpdateZetaFee(&_ERC20Custody.TransactOpts, zetaFee_)
}

// UpdateZetaFee is a paid mutator transaction binding the contract method 0xde2f6c5e.
//
// Solidity: function updateZetaFee(uint256 zetaFee_) returns()
func (_ERC20Custody *ERC20CustodyTransactorSession) UpdateZetaFee(zetaFee_ *big.Int) (*types.Transaction, error) {
	return _ERC20Custody.Contract.UpdateZetaFee(&_ERC20Custody.TransactOpts, zetaFee_)
}

// Whitelist is a paid mutator transaction binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address asset) returns()
func (_ERC20Custody *ERC20CustodyTransactor) Whitelist(opts *bind.TransactOpts, asset common.Address) (*types.Transaction, error) {
	return _ERC20Custody.contract.Transact(opts, "whitelist", asset)
}

// Whitelist is a paid mutator transaction binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address asset) returns()
func (_ERC20Custody *ERC20CustodySession) Whitelist(asset common.Address) (*types.Transaction, error) {
	return _ERC20Custody.Contract.Whitelist(&_ERC20Custody.TransactOpts, asset)
}

// Whitelist is a paid mutator transaction binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address asset) returns()
func (_ERC20Custody *ERC20CustodyTransactorSession) Whitelist(asset common.Address) (*types.Transaction, error) {
	return _ERC20Custody.Contract.Whitelist(&_ERC20Custody.TransactOpts, asset)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address asset, uint256 amount) returns()
func (_ERC20Custody *ERC20CustodyTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Custody.contract.Transact(opts, "withdraw", recipient, asset, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address asset, uint256 amount) returns()
func (_ERC20Custody *ERC20CustodySession) Withdraw(recipient common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Custody.Contract.Withdraw(&_ERC20Custody.TransactOpts, recipient, asset, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address asset, uint256 amount) returns()
func (_ERC20Custody *ERC20CustodyTransactorSession) Withdraw(recipient common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Custody.Contract.Withdraw(&_ERC20Custody.TransactOpts, recipient, asset, amount)
}

// ERC20CustodyDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the ERC20Custody contract.
type ERC20CustodyDepositedIterator struct {
	Event *ERC20CustodyDeposited // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyDeposited)
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
		it.Event = new(ERC20CustodyDeposited)
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
func (it *ERC20CustodyDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyDeposited represents a Deposited event raised by the ERC20Custody contract.
type ERC20CustodyDeposited struct {
	Recipient []byte
	Asset     common.Address
	Amount    *big.Int
	Message   []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae.
//
// Solidity: event Deposited(bytes recipient, address indexed asset, uint256 amount, bytes message)
func (_ERC20Custody *ERC20CustodyFilterer) FilterDeposited(opts *bind.FilterOpts, asset []common.Address) (*ERC20CustodyDepositedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _ERC20Custody.contract.FilterLogs(opts, "Deposited", assetRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyDepositedIterator{contract: _ERC20Custody.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae.
//
// Solidity: event Deposited(bytes recipient, address indexed asset, uint256 amount, bytes message)
func (_ERC20Custody *ERC20CustodyFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *ERC20CustodyDeposited, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _ERC20Custody.contract.WatchLogs(opts, "Deposited", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyDeposited)
				if err := _ERC20Custody.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae.
//
// Solidity: event Deposited(bytes recipient, address indexed asset, uint256 amount, bytes message)
func (_ERC20Custody *ERC20CustodyFilterer) ParseDeposited(log types.Log) (*ERC20CustodyDeposited, error) {
	event := new(ERC20CustodyDeposited)
	if err := _ERC20Custody.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ERC20Custody contract.
type ERC20CustodyPausedIterator struct {
	Event *ERC20CustodyPaused // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyPaused)
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
		it.Event = new(ERC20CustodyPaused)
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
func (it *ERC20CustodyPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyPaused represents a Paused event raised by the ERC20Custody contract.
type ERC20CustodyPaused struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address sender)
func (_ERC20Custody *ERC20CustodyFilterer) FilterPaused(opts *bind.FilterOpts) (*ERC20CustodyPausedIterator, error) {

	logs, sub, err := _ERC20Custody.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyPausedIterator{contract: _ERC20Custody.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address sender)
func (_ERC20Custody *ERC20CustodyFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ERC20CustodyPaused) (event.Subscription, error) {

	logs, sub, err := _ERC20Custody.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyPaused)
				if err := _ERC20Custody.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address sender)
func (_ERC20Custody *ERC20CustodyFilterer) ParsePaused(log types.Log) (*ERC20CustodyPaused, error) {
	event := new(ERC20CustodyPaused)
	if err := _ERC20Custody.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyRenouncedTSSUpdaterIterator is returned from FilterRenouncedTSSUpdater and is used to iterate over the raw logs and unpacked data for RenouncedTSSUpdater events raised by the ERC20Custody contract.
type ERC20CustodyRenouncedTSSUpdaterIterator struct {
	Event *ERC20CustodyRenouncedTSSUpdater // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyRenouncedTSSUpdaterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyRenouncedTSSUpdater)
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
		it.Event = new(ERC20CustodyRenouncedTSSUpdater)
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
func (it *ERC20CustodyRenouncedTSSUpdaterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyRenouncedTSSUpdaterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyRenouncedTSSUpdater represents a RenouncedTSSUpdater event raised by the ERC20Custody contract.
type ERC20CustodyRenouncedTSSUpdater struct {
	TSSAddressUpdater common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRenouncedTSSUpdater is a free log retrieval operation binding the contract event 0x39ac266bfc011581be62c138d96e4e8782812013bb66fffb4cd207f4bfee3777.
//
// Solidity: event RenouncedTSSUpdater(address TSSAddressUpdater_)
func (_ERC20Custody *ERC20CustodyFilterer) FilterRenouncedTSSUpdater(opts *bind.FilterOpts) (*ERC20CustodyRenouncedTSSUpdaterIterator, error) {

	logs, sub, err := _ERC20Custody.contract.FilterLogs(opts, "RenouncedTSSUpdater")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyRenouncedTSSUpdaterIterator{contract: _ERC20Custody.contract, event: "RenouncedTSSUpdater", logs: logs, sub: sub}, nil
}

// WatchRenouncedTSSUpdater is a free log subscription operation binding the contract event 0x39ac266bfc011581be62c138d96e4e8782812013bb66fffb4cd207f4bfee3777.
//
// Solidity: event RenouncedTSSUpdater(address TSSAddressUpdater_)
func (_ERC20Custody *ERC20CustodyFilterer) WatchRenouncedTSSUpdater(opts *bind.WatchOpts, sink chan<- *ERC20CustodyRenouncedTSSUpdater) (event.Subscription, error) {

	logs, sub, err := _ERC20Custody.contract.WatchLogs(opts, "RenouncedTSSUpdater")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyRenouncedTSSUpdater)
				if err := _ERC20Custody.contract.UnpackLog(event, "RenouncedTSSUpdater", log); err != nil {
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

// ParseRenouncedTSSUpdater is a log parse operation binding the contract event 0x39ac266bfc011581be62c138d96e4e8782812013bb66fffb4cd207f4bfee3777.
//
// Solidity: event RenouncedTSSUpdater(address TSSAddressUpdater_)
func (_ERC20Custody *ERC20CustodyFilterer) ParseRenouncedTSSUpdater(log types.Log) (*ERC20CustodyRenouncedTSSUpdater, error) {
	event := new(ERC20CustodyRenouncedTSSUpdater)
	if err := _ERC20Custody.contract.UnpackLog(event, "RenouncedTSSUpdater", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ERC20Custody contract.
type ERC20CustodyUnpausedIterator struct {
	Event *ERC20CustodyUnpaused // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyUnpaused)
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
		it.Event = new(ERC20CustodyUnpaused)
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
func (it *ERC20CustodyUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyUnpaused represents a Unpaused event raised by the ERC20Custody contract.
type ERC20CustodyUnpaused struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address sender)
func (_ERC20Custody *ERC20CustodyFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ERC20CustodyUnpausedIterator, error) {

	logs, sub, err := _ERC20Custody.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyUnpausedIterator{contract: _ERC20Custody.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address sender)
func (_ERC20Custody *ERC20CustodyFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ERC20CustodyUnpaused) (event.Subscription, error) {

	logs, sub, err := _ERC20Custody.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyUnpaused)
				if err := _ERC20Custody.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address sender)
func (_ERC20Custody *ERC20CustodyFilterer) ParseUnpaused(log types.Log) (*ERC20CustodyUnpaused, error) {
	event := new(ERC20CustodyUnpaused)
	if err := _ERC20Custody.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyUnwhitelistedIterator is returned from FilterUnwhitelisted and is used to iterate over the raw logs and unpacked data for Unwhitelisted events raised by the ERC20Custody contract.
type ERC20CustodyUnwhitelistedIterator struct {
	Event *ERC20CustodyUnwhitelisted // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyUnwhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyUnwhitelisted)
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
		it.Event = new(ERC20CustodyUnwhitelisted)
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
func (it *ERC20CustodyUnwhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyUnwhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyUnwhitelisted represents a Unwhitelisted event raised by the ERC20Custody contract.
type ERC20CustodyUnwhitelisted struct {
	Asset common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnwhitelisted is a free log retrieval operation binding the contract event 0x51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da46791.
//
// Solidity: event Unwhitelisted(address indexed asset)
func (_ERC20Custody *ERC20CustodyFilterer) FilterUnwhitelisted(opts *bind.FilterOpts, asset []common.Address) (*ERC20CustodyUnwhitelistedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _ERC20Custody.contract.FilterLogs(opts, "Unwhitelisted", assetRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyUnwhitelistedIterator{contract: _ERC20Custody.contract, event: "Unwhitelisted", logs: logs, sub: sub}, nil
}

// WatchUnwhitelisted is a free log subscription operation binding the contract event 0x51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da46791.
//
// Solidity: event Unwhitelisted(address indexed asset)
func (_ERC20Custody *ERC20CustodyFilterer) WatchUnwhitelisted(opts *bind.WatchOpts, sink chan<- *ERC20CustodyUnwhitelisted, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _ERC20Custody.contract.WatchLogs(opts, "Unwhitelisted", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyUnwhitelisted)
				if err := _ERC20Custody.contract.UnpackLog(event, "Unwhitelisted", log); err != nil {
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

// ParseUnwhitelisted is a log parse operation binding the contract event 0x51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da46791.
//
// Solidity: event Unwhitelisted(address indexed asset)
func (_ERC20Custody *ERC20CustodyFilterer) ParseUnwhitelisted(log types.Log) (*ERC20CustodyUnwhitelisted, error) {
	event := new(ERC20CustodyUnwhitelisted)
	if err := _ERC20Custody.contract.UnpackLog(event, "Unwhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyUpdatedTSSAddressIterator is returned from FilterUpdatedTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedTSSAddress events raised by the ERC20Custody contract.
type ERC20CustodyUpdatedTSSAddressIterator struct {
	Event *ERC20CustodyUpdatedTSSAddress // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyUpdatedTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyUpdatedTSSAddress)
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
		it.Event = new(ERC20CustodyUpdatedTSSAddress)
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
func (it *ERC20CustodyUpdatedTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyUpdatedTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyUpdatedTSSAddress represents a UpdatedTSSAddress event raised by the ERC20Custody contract.
type ERC20CustodyUpdatedTSSAddress struct {
	TSSAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdatedTSSAddress is a free log retrieval operation binding the contract event 0xcd2958db8285a532edf298cbe1aa28ea3fb5ec82461253f9a8c1477924e29947.
//
// Solidity: event UpdatedTSSAddress(address TSSAddress_)
func (_ERC20Custody *ERC20CustodyFilterer) FilterUpdatedTSSAddress(opts *bind.FilterOpts) (*ERC20CustodyUpdatedTSSAddressIterator, error) {

	logs, sub, err := _ERC20Custody.contract.FilterLogs(opts, "UpdatedTSSAddress")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyUpdatedTSSAddressIterator{contract: _ERC20Custody.contract, event: "UpdatedTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedTSSAddress is a free log subscription operation binding the contract event 0xcd2958db8285a532edf298cbe1aa28ea3fb5ec82461253f9a8c1477924e29947.
//
// Solidity: event UpdatedTSSAddress(address TSSAddress_)
func (_ERC20Custody *ERC20CustodyFilterer) WatchUpdatedTSSAddress(opts *bind.WatchOpts, sink chan<- *ERC20CustodyUpdatedTSSAddress) (event.Subscription, error) {

	logs, sub, err := _ERC20Custody.contract.WatchLogs(opts, "UpdatedTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyUpdatedTSSAddress)
				if err := _ERC20Custody.contract.UnpackLog(event, "UpdatedTSSAddress", log); err != nil {
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

// ParseUpdatedTSSAddress is a log parse operation binding the contract event 0xcd2958db8285a532edf298cbe1aa28ea3fb5ec82461253f9a8c1477924e29947.
//
// Solidity: event UpdatedTSSAddress(address TSSAddress_)
func (_ERC20Custody *ERC20CustodyFilterer) ParseUpdatedTSSAddress(log types.Log) (*ERC20CustodyUpdatedTSSAddress, error) {
	event := new(ERC20CustodyUpdatedTSSAddress)
	if err := _ERC20Custody.contract.UnpackLog(event, "UpdatedTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyUpdatedZetaFeeIterator is returned from FilterUpdatedZetaFee and is used to iterate over the raw logs and unpacked data for UpdatedZetaFee events raised by the ERC20Custody contract.
type ERC20CustodyUpdatedZetaFeeIterator struct {
	Event *ERC20CustodyUpdatedZetaFee // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyUpdatedZetaFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyUpdatedZetaFee)
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
		it.Event = new(ERC20CustodyUpdatedZetaFee)
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
func (it *ERC20CustodyUpdatedZetaFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyUpdatedZetaFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyUpdatedZetaFee represents a UpdatedZetaFee event raised by the ERC20Custody contract.
type ERC20CustodyUpdatedZetaFee struct {
	ZetaFee *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdatedZetaFee is a free log retrieval operation binding the contract event 0x6d2d8e313fbaf76898bb9fa55e4b52525e49c7d7182d0874f97bd9076e81d523.
//
// Solidity: event UpdatedZetaFee(uint256 zetaFee_)
func (_ERC20Custody *ERC20CustodyFilterer) FilterUpdatedZetaFee(opts *bind.FilterOpts) (*ERC20CustodyUpdatedZetaFeeIterator, error) {

	logs, sub, err := _ERC20Custody.contract.FilterLogs(opts, "UpdatedZetaFee")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyUpdatedZetaFeeIterator{contract: _ERC20Custody.contract, event: "UpdatedZetaFee", logs: logs, sub: sub}, nil
}

// WatchUpdatedZetaFee is a free log subscription operation binding the contract event 0x6d2d8e313fbaf76898bb9fa55e4b52525e49c7d7182d0874f97bd9076e81d523.
//
// Solidity: event UpdatedZetaFee(uint256 zetaFee_)
func (_ERC20Custody *ERC20CustodyFilterer) WatchUpdatedZetaFee(opts *bind.WatchOpts, sink chan<- *ERC20CustodyUpdatedZetaFee) (event.Subscription, error) {

	logs, sub, err := _ERC20Custody.contract.WatchLogs(opts, "UpdatedZetaFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyUpdatedZetaFee)
				if err := _ERC20Custody.contract.UnpackLog(event, "UpdatedZetaFee", log); err != nil {
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

// ParseUpdatedZetaFee is a log parse operation binding the contract event 0x6d2d8e313fbaf76898bb9fa55e4b52525e49c7d7182d0874f97bd9076e81d523.
//
// Solidity: event UpdatedZetaFee(uint256 zetaFee_)
func (_ERC20Custody *ERC20CustodyFilterer) ParseUpdatedZetaFee(log types.Log) (*ERC20CustodyUpdatedZetaFee, error) {
	event := new(ERC20CustodyUpdatedZetaFee)
	if err := _ERC20Custody.contract.UnpackLog(event, "UpdatedZetaFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyWhitelistedIterator is returned from FilterWhitelisted and is used to iterate over the raw logs and unpacked data for Whitelisted events raised by the ERC20Custody contract.
type ERC20CustodyWhitelistedIterator struct {
	Event *ERC20CustodyWhitelisted // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyWhitelisted)
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
		it.Event = new(ERC20CustodyWhitelisted)
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
func (it *ERC20CustodyWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyWhitelisted represents a Whitelisted event raised by the ERC20Custody contract.
type ERC20CustodyWhitelisted struct {
	Asset common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWhitelisted is a free log retrieval operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed asset)
func (_ERC20Custody *ERC20CustodyFilterer) FilterWhitelisted(opts *bind.FilterOpts, asset []common.Address) (*ERC20CustodyWhitelistedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _ERC20Custody.contract.FilterLogs(opts, "Whitelisted", assetRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyWhitelistedIterator{contract: _ERC20Custody.contract, event: "Whitelisted", logs: logs, sub: sub}, nil
}

// WatchWhitelisted is a free log subscription operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed asset)
func (_ERC20Custody *ERC20CustodyFilterer) WatchWhitelisted(opts *bind.WatchOpts, sink chan<- *ERC20CustodyWhitelisted, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _ERC20Custody.contract.WatchLogs(opts, "Whitelisted", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyWhitelisted)
				if err := _ERC20Custody.contract.UnpackLog(event, "Whitelisted", log); err != nil {
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

// ParseWhitelisted is a log parse operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed asset)
func (_ERC20Custody *ERC20CustodyFilterer) ParseWhitelisted(log types.Log) (*ERC20CustodyWhitelisted, error) {
	event := new(ERC20CustodyWhitelisted)
	if err := _ERC20Custody.contract.UnpackLog(event, "Whitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the ERC20Custody contract.
type ERC20CustodyWithdrawnIterator struct {
	Event *ERC20CustodyWithdrawn // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyWithdrawn)
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
		it.Event = new(ERC20CustodyWithdrawn)
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
func (it *ERC20CustodyWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyWithdrawn represents a Withdrawn event raised by the ERC20Custody contract.
type ERC20CustodyWithdrawn struct {
	Recipient common.Address
	Asset     common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed recipient, address indexed asset, uint256 amount)
func (_ERC20Custody *ERC20CustodyFilterer) FilterWithdrawn(opts *bind.FilterOpts, recipient []common.Address, asset []common.Address) (*ERC20CustodyWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _ERC20Custody.contract.FilterLogs(opts, "Withdrawn", recipientRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyWithdrawnIterator{contract: _ERC20Custody.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed recipient, address indexed asset, uint256 amount)
func (_ERC20Custody *ERC20CustodyFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *ERC20CustodyWithdrawn, recipient []common.Address, asset []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _ERC20Custody.contract.WatchLogs(opts, "Withdrawn", recipientRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyWithdrawn)
				if err := _ERC20Custody.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed recipient, address indexed asset, uint256 amount)
func (_ERC20Custody *ERC20CustodyFilterer) ParseWithdrawn(log types.Log) (*ERC20CustodyWithdrawn, error) {
	event := new(ERC20CustodyWithdrawn)
	if err := _ERC20Custody.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
