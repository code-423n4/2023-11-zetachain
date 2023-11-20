# Report


## Gas Optimizations


| |Issue|Instances|
|-|:-|:-:|
| [GAS-1](#GAS-1) | Using bools for storage incurs overhead | 6 |
| [GAS-2](#GAS-2) | For Operations that will not overflow, you could use unchecked | 344 |
| [GAS-3](#GAS-3) | Use Custom Errors | 1 |
| [GAS-4](#GAS-4) | Long revert strings | 1 |
| [GAS-5](#GAS-5) | Functions guaranteed to revert when called by normal users can be marked `payable` | 17 |
| [GAS-6](#GAS-6) | Using `private` rather than `public` for constants, saves gas | 3 |
| [GAS-7](#GAS-7) | Use shift Right/Left instead of division/multiplication if possible | 4 |
| [GAS-8](#GAS-8) | Use != 0 instead of > 0 for unsigned integer comparison | 10 |
### <a name="GAS-1"></a>[GAS-1] Using bools for storage incurs overhead
Use uint256(1) and uint256(2) for true/false to avoid a Gwarmaccess (100 gas), and to avoid Gsset (20000 gas) when changing from ‘false’ to ‘true’, after having been ‘true’ in the past. See [source](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/58f635312aa21f947cae5f8578638a85aa2519f5/contracts/security/ReentrancyGuard.sol#L23-L27).

*Instances (6)*:
```solidity
File: contracts/evm/ERC20Custody.sol

24:     bool public paused;

36:     mapping(IERC20 => bool) public whitelisted;

```

```solidity
File: contracts/evm/tools/ImmutableCreate2Factory.sol

24:     mapping(address => bool) private _deployed;

```

```solidity
File: contracts/evm/tools/ZetaTokenConsumerPancakeV3.strategy.sol

69:     bool internal _locked;

```

```solidity
File: contracts/evm/tools/ZetaTokenConsumerTrident.strategy.sol

43:     bool internal _locked;

```

```solidity
File: contracts/evm/tools/ZetaTokenConsumerUniV3.strategy.sol

40:     bool internal _locked;

```

### <a name="GAS-2"></a>[GAS-2] For Operations that will not overflow, you could use unchecked

*Instances (344)*:
```solidity
File: contracts/evm/ERC20Custody.sol

5: import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

5: import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

5: import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

5: import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

6: import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

6: import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

6: import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

6: import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

6: import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

7: import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

7: import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

7: import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

184:         emit Deposited(recipient, asset, asset.balanceOf(address(this)) - oldBalance, message);

```

```solidity
File: contracts/evm/Zeta.eth.sol

4: import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

4: import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

4: import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

4: import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

11:         _mint(creator, initialSupply * (10 ** uint256(decimals())));

11:         _mint(creator, initialSupply * (10 ** uint256(decimals())));

11:         _mint(creator, initialSupply * (10 ** uint256(decimals())));

```

```solidity
File: contracts/evm/Zeta.non-eth.sol

4: import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";

4: import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";

4: import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";

4: import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";

4: import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";

6: import "./interfaces/ZetaErrors.sol";

6: import "./interfaces/ZetaErrors.sol";

8: import "./interfaces/ZetaNonEthInterface.sol";

8: import "./interfaces/ZetaNonEthInterface.sol";

```

```solidity
File: contracts/evm/ZetaConnector.base.sol

4: import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

4: import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

4: import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

4: import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

5: import "@openzeppelin/contracts/security/Pausable.sol";

5: import "@openzeppelin/contracts/security/Pausable.sol";

5: import "@openzeppelin/contracts/security/Pausable.sol";

7: import "./interfaces/ConnectorErrors.sol";

7: import "./interfaces/ConnectorErrors.sol";

8: import "./interfaces/ZetaInterfaces.sol";

8: import "./interfaces/ZetaInterfaces.sol";

```

```solidity
File: contracts/evm/ZetaConnector.eth.sol

4: import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

4: import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

4: import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

4: import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

6: import "./interfaces/ConnectorErrors.sol";

6: import "./interfaces/ConnectorErrors.sol";

7: import "./ZetaConnector.base.sol";

8: import "./interfaces/ZetaInterfaces.sol";

8: import "./interfaces/ZetaInterfaces.sol";

```

```solidity
File: contracts/evm/ZetaConnector.non-eth.sol

4: import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

4: import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

4: import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

4: import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

6: import "./ZetaConnector.base.sol";

7: import "./interfaces/ZetaInterfaces.sol";

7: import "./interfaces/ZetaInterfaces.sol";

8: import "./interfaces/ZetaNonEthInterface.sol";

8: import "./interfaces/ZetaNonEthInterface.sol";

16:     uint256 public maxSupply = 2 ** 256 - 1;

16:     uint256 public maxSupply = 2 ** 256 - 1;

16:     uint256 public maxSupply = 2 ** 256 - 1;

69:         if (zetaValue + ZetaNonEthInterface(zetaToken).totalSupply() > maxSupply) revert ExceedsMaxSupply(maxSupply);

96:         if (remainingZetaValue + ZetaNonEthInterface(zetaToken).totalSupply() > maxSupply)

```

```solidity
File: contracts/evm/interfaces/ZetaNonEthInterface.sol

4: import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

4: import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

4: import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

4: import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

```

```solidity
File: contracts/evm/tools/ImmutableCreate2Factory.sol

1: pragma solidity 0.5.10; // optimization enabled, 99999 runs, evm: petersburg

1: pragma solidity 0.5.10; // optimization enabled, 99999 runs, evm: petersburg

35:             uint160( // downcast to match the address type.

35:             uint160( // downcast to match the address type.

36:                 uint256( // convert to uint to truncate upper digits.

36:                 uint256( // convert to uint to truncate upper digits.

37:                     keccak256( // compute the CREATE2 hash using 4 inputs.

37:                     keccak256( // compute the CREATE2 hash using 4 inputs.

38:                         abi.encodePacked( // pack all inputs to the hash together.

38:                         abi.encodePacked( // pack all inputs to the hash together.

39:                             hex"ff", // start with 0xff to distinguish from RLP.

39:                             hex"ff", // start with 0xff to distinguish from RLP.

40:                             address(this), // this contract will be the caller.

40:                             address(this), // this contract will be the caller.

41:                             salt, // pass in the supplied salt value.

41:                             salt, // pass in the supplied salt value.

42:                             keccak256(abi.encodePacked(initCode)) // pass in the hash of initialization code.

42:                             keccak256(abi.encodePacked(initCode)) // pass in the hash of initialization code.

50:         require(!_deployed[targetDeploymentAddress], "Invalid contract creation - contract has already been deployed.");

55:             let encoded_data := add(0x20, initCode) // load initialization code.

55:             let encoded_data := add(0x20, initCode) // load initialization code.

56:             let encoded_size := mload(initCode) // load the init code's length.

56:             let encoded_size := mload(initCode) // load the init code's length.

59:                 callvalue, // forward any attached value.

59:                 callvalue, // forward any attached value.

60:                 encoded_data, // pass in initialization code.

60:                 encoded_data, // pass in initialization code.

61:                 encoded_size, // pass in init code's length.

61:                 encoded_size, // pass in init code's length.

62:                 salt // pass in the salt value.

62:                 salt // pass in the salt value.

114:             uint160( // downcast to match the address type.

114:             uint160( // downcast to match the address type.

115:                 uint256( // convert to uint to truncate upper digits.

115:                 uint256( // convert to uint to truncate upper digits.

116:                     keccak256( // compute the CREATE2 hash using 4 inputs.

116:                     keccak256( // compute the CREATE2 hash using 4 inputs.

117:                         abi.encodePacked( // pack all inputs to the hash together.

117:                         abi.encodePacked( // pack all inputs to the hash together.

118:                             hex"ff", // start with 0xff to distinguish from RLP.

118:                             hex"ff", // start with 0xff to distinguish from RLP.

119:                             address(this), // this contract will be the caller.

119:                             address(this), // this contract will be the caller.

120:                             salt, // pass in the supplied salt value.

120:                             salt, // pass in the supplied salt value.

121:                             keccak256(abi.encodePacked(initCode)) // pass in the hash of initialization code.

121:                             keccak256(abi.encodePacked(initCode)) // pass in the hash of initialization code.

154:             uint160( // downcast to match the address type.

154:             uint160( // downcast to match the address type.

155:                 uint256( // convert to uint to truncate upper digits.

155:                 uint256( // convert to uint to truncate upper digits.

156:                     keccak256( // compute the CREATE2 hash using 4 inputs.

156:                     keccak256( // compute the CREATE2 hash using 4 inputs.

157:                         abi.encodePacked( // pack all inputs to the hash together.

157:                         abi.encodePacked( // pack all inputs to the hash together.

158:                             hex"ff", // start with 0xff to distinguish from RLP.

158:                             hex"ff", // start with 0xff to distinguish from RLP.

159:                             address(this), // this contract will be the caller.

159:                             address(this), // this contract will be the caller.

160:                             salt, // pass in the supplied salt value.

160:                             salt, // pass in the supplied salt value.

161:                             initCodeHash // pass in the hash of initialization code.

161:                             initCodeHash // pass in the hash of initialization code.

197:             "Invalid salt - first 20 bytes of the salt must match calling address."

```

```solidity
File: contracts/evm/tools/ZetaInteractor.sol

4: import "@openzeppelin/contracts/access/Ownable2Step.sol";

4: import "@openzeppelin/contracts/access/Ownable2Step.sol";

4: import "@openzeppelin/contracts/access/Ownable2Step.sol";

6: import "../interfaces/ZetaInterfaces.sol";

6: import "../interfaces/ZetaInterfaces.sol";

7: import "../interfaces/ZetaInteractorErrors.sol";

7: import "../interfaces/ZetaInteractorErrors.sol";

```

```solidity
File: contracts/evm/tools/ZetaTokenConsumerPancakeV3.strategy.sol

4: import "@openzeppelin/contracts/interfaces/IERC20.sol";

4: import "@openzeppelin/contracts/interfaces/IERC20.sol";

4: import "@openzeppelin/contracts/interfaces/IERC20.sol";

5: import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

5: import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

5: import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

5: import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

5: import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

6: import "@uniswap/v3-periphery/contracts/interfaces/ISwapRouter.sol";

6: import "@uniswap/v3-periphery/contracts/interfaces/ISwapRouter.sol";

6: import "@uniswap/v3-periphery/contracts/interfaces/ISwapRouter.sol";

6: import "@uniswap/v3-periphery/contracts/interfaces/ISwapRouter.sol";

6: import "@uniswap/v3-periphery/contracts/interfaces/ISwapRouter.sol";

7: import "@uniswap/v3-core/contracts/interfaces/IUniswapV3Factory.sol";

7: import "@uniswap/v3-core/contracts/interfaces/IUniswapV3Factory.sol";

7: import "@uniswap/v3-core/contracts/interfaces/IUniswapV3Factory.sol";

7: import "@uniswap/v3-core/contracts/interfaces/IUniswapV3Factory.sol";

7: import "@uniswap/v3-core/contracts/interfaces/IUniswapV3Factory.sol";

8: import "@uniswap/v3-core/contracts/interfaces/IUniswapV3Pool.sol";

8: import "@uniswap/v3-core/contracts/interfaces/IUniswapV3Pool.sol";

8: import "@uniswap/v3-core/contracts/interfaces/IUniswapV3Pool.sol";

8: import "@uniswap/v3-core/contracts/interfaces/IUniswapV3Pool.sol";

8: import "@uniswap/v3-core/contracts/interfaces/IUniswapV3Pool.sol";

10: import "../interfaces/ZetaInterfaces.sol";

10: import "../interfaces/ZetaInterfaces.sol";

```

```solidity
File: contracts/evm/tools/ZetaTokenConsumerTrident.strategy.sol

4: import "@openzeppelin/contracts/interfaces/IERC20.sol";

4: import "@openzeppelin/contracts/interfaces/IERC20.sol";

4: import "@openzeppelin/contracts/interfaces/IERC20.sol";

5: import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

5: import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

5: import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

5: import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

5: import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

6: import "@uniswap/v3-periphery/contracts/interfaces/IQuoter.sol";

6: import "@uniswap/v3-periphery/contracts/interfaces/IQuoter.sol";

6: import "@uniswap/v3-periphery/contracts/interfaces/IQuoter.sol";

6: import "@uniswap/v3-periphery/contracts/interfaces/IQuoter.sol";

6: import "@uniswap/v3-periphery/contracts/interfaces/IQuoter.sol";

8: import "../interfaces/ZetaInterfaces.sol";

8: import "../interfaces/ZetaInterfaces.sol";

9: import "./interfaces/TridentConcentratedLiquidityPoolFactory.sol";

9: import "./interfaces/TridentConcentratedLiquidityPoolFactory.sol";

10: import "./interfaces/TridentIPoolRouter.sol";

10: import "./interfaces/TridentIPoolRouter.sol";

```

```solidity
File: contracts/evm/tools/ZetaTokenConsumerUniV2.strategy.sol

4: import "@openzeppelin/contracts/interfaces/IERC20.sol";

4: import "@openzeppelin/contracts/interfaces/IERC20.sol";

4: import "@openzeppelin/contracts/interfaces/IERC20.sol";

5: import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

5: import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

5: import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

5: import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

5: import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

6: import "@uniswap/v2-periphery/contracts/interfaces/IUniswapV2Router02.sol";

6: import "@uniswap/v2-periphery/contracts/interfaces/IUniswapV2Router02.sol";

6: import "@uniswap/v2-periphery/contracts/interfaces/IUniswapV2Router02.sol";

6: import "@uniswap/v2-periphery/contracts/interfaces/IUniswapV2Router02.sol";

6: import "@uniswap/v2-periphery/contracts/interfaces/IUniswapV2Router02.sol";

8: import "../interfaces/ZetaInterfaces.sol";

8: import "../interfaces/ZetaInterfaces.sol";

49:             block.timestamp + MAX_DEADLINE

52:         uint256 amountOut = amounts[path.length - 1];

87:             block.timestamp + MAX_DEADLINE

89:         uint256 amountOut = amounts[path.length - 1];

115:             block.timestamp + MAX_DEADLINE

118:         uint256 amountOut = amounts[path.length - 1];

153:             block.timestamp + MAX_DEADLINE

156:         uint256 amountOut = amounts[path.length - 1];

168:             return amounts[path.length - 1] > 0;

```

```solidity
File: contracts/evm/tools/ZetaTokenConsumerUniV3.strategy.sol

4: import "@openzeppelin/contracts/interfaces/IERC20.sol";

4: import "@openzeppelin/contracts/interfaces/IERC20.sol";

4: import "@openzeppelin/contracts/interfaces/IERC20.sol";

5: import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

5: import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

5: import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

5: import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

5: import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

6: import "@uniswap/v3-periphery/contracts/interfaces/ISwapRouter.sol";

6: import "@uniswap/v3-periphery/contracts/interfaces/ISwapRouter.sol";

6: import "@uniswap/v3-periphery/contracts/interfaces/ISwapRouter.sol";

6: import "@uniswap/v3-periphery/contracts/interfaces/ISwapRouter.sol";

6: import "@uniswap/v3-periphery/contracts/interfaces/ISwapRouter.sol";

7: import "@uniswap/v3-core/contracts/interfaces/IUniswapV3Factory.sol";

7: import "@uniswap/v3-core/contracts/interfaces/IUniswapV3Factory.sol";

7: import "@uniswap/v3-core/contracts/interfaces/IUniswapV3Factory.sol";

7: import "@uniswap/v3-core/contracts/interfaces/IUniswapV3Factory.sol";

7: import "@uniswap/v3-core/contracts/interfaces/IUniswapV3Factory.sol";

8: import "@uniswap/v3-core/contracts/interfaces/IUniswapV3Pool.sol";

8: import "@uniswap/v3-core/contracts/interfaces/IUniswapV3Pool.sol";

8: import "@uniswap/v3-core/contracts/interfaces/IUniswapV3Pool.sol";

8: import "@uniswap/v3-core/contracts/interfaces/IUniswapV3Pool.sol";

8: import "@uniswap/v3-core/contracts/interfaces/IUniswapV3Pool.sol";

10: import "../interfaces/ZetaInterfaces.sol";

10: import "../interfaces/ZetaInterfaces.sol";

82:             deadline: block.timestamp + MAX_DEADLINE,

111:             deadline: block.timestamp + MAX_DEADLINE,

136:             deadline: block.timestamp + MAX_DEADLINE,

171:             deadline: block.timestamp + MAX_DEADLINE,

```

```solidity
File: contracts/evm/tools/interfaces/TridentIPoolRouter.sol

18:         address tokenIn; /// @dev the input token address. If tokenIn is address(0), msg.value will be wrapped and used as input token

18:         address tokenIn; /// @dev the input token address. If tokenIn is address(0), msg.value will be wrapped and used as input token

18:         address tokenIn; /// @dev the input token address. If tokenIn is address(0), msg.value will be wrapped and used as input token

19:         uint256 amountIn; /// @dev The amount of input tokens to send

19:         uint256 amountIn; /// @dev The amount of input tokens to send

19:         uint256 amountIn; /// @dev The amount of input tokens to send

20:         uint256 amountOutMinimum; /// @dev minimum required amount of output token after swap

20:         uint256 amountOutMinimum; /// @dev minimum required amount of output token after swap

20:         uint256 amountOutMinimum; /// @dev minimum required amount of output token after swap

21:         address pool; /// @dev pool address to swap

21:         address pool; /// @dev pool address to swap

21:         address pool; /// @dev pool address to swap

22:         address to; /// @dev address to receive

22:         address to; /// @dev address to receive

22:         address to; /// @dev address to receive

23:         bool unwrap; /// @dev unwrap if output token is wrapped klay

23:         bool unwrap; /// @dev unwrap if output token is wrapped klay

23:         bool unwrap; /// @dev unwrap if output token is wrapped klay

27:         address tokenIn; /// @dev the token address to swap-in. If tokenIn is address(0), msg.value will be wrapped and used as input token

27:         address tokenIn; /// @dev the token address to swap-in. If tokenIn is address(0), msg.value will be wrapped and used as input token

27:         address tokenIn; /// @dev the token address to swap-in. If tokenIn is address(0), msg.value will be wrapped and used as input token

27:         address tokenIn; /// @dev the token address to swap-in. If tokenIn is address(0), msg.value will be wrapped and used as input token

28:         uint256 amountIn; /// @dev The amount of input tokens to send.

28:         uint256 amountIn; /// @dev The amount of input tokens to send.

28:         uint256 amountIn; /// @dev The amount of input tokens to send.

29:         uint256 amountOutMinimum; /// @dev minimum required amount of output token after swap

29:         uint256 amountOutMinimum; /// @dev minimum required amount of output token after swap

29:         uint256 amountOutMinimum; /// @dev minimum required amount of output token after swap

30:         address[] path; /// @dev An array of pool addresses to pass through

30:         address[] path; /// @dev An array of pool addresses to pass through

30:         address[] path; /// @dev An array of pool addresses to pass through

31:         address to; /// @dev recipient of the output tokens

31:         address to; /// @dev recipient of the output tokens

31:         address to; /// @dev recipient of the output tokens

32:         bool unwrap; /// @dev unwrap if output token is wrapped klay

32:         bool unwrap; /// @dev unwrap if output token is wrapped klay

32:         bool unwrap; /// @dev unwrap if output token is wrapped klay

36:         address tokenIn; /// @dev the input token address. If tokenIn is address(0), msg.value will be wrapped and used as input token

36:         address tokenIn; /// @dev the input token address. If tokenIn is address(0), msg.value will be wrapped and used as input token

36:         address tokenIn; /// @dev the input token address. If tokenIn is address(0), msg.value will be wrapped and used as input token

37:         uint256 amountOut; /// @dev The amount of output tokens to receive

37:         uint256 amountOut; /// @dev The amount of output tokens to receive

37:         uint256 amountOut; /// @dev The amount of output tokens to receive

38:         uint256 amountInMaximum; /// @dev maximum available amount of input token after swap

38:         uint256 amountInMaximum; /// @dev maximum available amount of input token after swap

38:         uint256 amountInMaximum; /// @dev maximum available amount of input token after swap

39:         address pool; /// @dev pool address to swap

39:         address pool; /// @dev pool address to swap

39:         address pool; /// @dev pool address to swap

40:         address to; /// @dev address to receive

40:         address to; /// @dev address to receive

40:         address to; /// @dev address to receive

41:         bool unwrap; /// @dev unwrap if output token is wrapped klay

41:         bool unwrap; /// @dev unwrap if output token is wrapped klay

41:         bool unwrap; /// @dev unwrap if output token is wrapped klay

45:         address tokenIn; /// @dev the token address to swap-in. If tokenIn is address(0), msg.value will be wrapped and used as input token

45:         address tokenIn; /// @dev the token address to swap-in. If tokenIn is address(0), msg.value will be wrapped and used as input token

45:         address tokenIn; /// @dev the token address to swap-in. If tokenIn is address(0), msg.value will be wrapped and used as input token

45:         address tokenIn; /// @dev the token address to swap-in. If tokenIn is address(0), msg.value will be wrapped and used as input token

46:         uint256 amountOut; /// @dev The amount of output tokens to receive

46:         uint256 amountOut; /// @dev The amount of output tokens to receive

46:         uint256 amountOut; /// @dev The amount of output tokens to receive

47:         uint256 amountInMaximum; /// @dev  maximum available amount of input token after swap

47:         uint256 amountInMaximum; /// @dev  maximum available amount of input token after swap

47:         uint256 amountInMaximum; /// @dev  maximum available amount of input token after swap

48:         address[] path; /// @dev An array of pool addresses to pass through

48:         address[] path; /// @dev An array of pool addresses to pass through

48:         address[] path; /// @dev An array of pool addresses to pass through

49:         address to; /// @dev recipient of the output tokens

49:         address to; /// @dev recipient of the output tokens

49:         address to; /// @dev recipient of the output tokens

50:         bool unwrap; /// @dev unwrap if output token is wrapped klay

50:         bool unwrap; /// @dev unwrap if output token is wrapped klay

50:         bool unwrap; /// @dev unwrap if output token is wrapped klay

```

```solidity
File: contracts/zevm/SystemContract.sol

4: import "./interfaces/zContract.sol";

4: import "./interfaces/zContract.sol";

5: import "./interfaces/IZRC20.sol";

5: import "./interfaces/IZRC20.sol";

110:                             hex"96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f" // init code hash

110:                             hex"96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f" // init code hash

```

```solidity
File: contracts/zevm/Uniswap.sol

4: import "@uniswap/v2-core/contracts/UniswapV2Pair.sol";

4: import "@uniswap/v2-core/contracts/UniswapV2Pair.sol";

4: import "@uniswap/v2-core/contracts/UniswapV2Pair.sol";

4: import "@uniswap/v2-core/contracts/UniswapV2Pair.sol";

5: import "@uniswap/v2-core/contracts/UniswapV2Factory.sol";

5: import "@uniswap/v2-core/contracts/UniswapV2Factory.sol";

5: import "@uniswap/v2-core/contracts/UniswapV2Factory.sol";

5: import "@uniswap/v2-core/contracts/UniswapV2Factory.sol";

```

```solidity
File: contracts/zevm/UniswapPeriphery.sol

4: import "@uniswap/v2-periphery/contracts/UniswapV2Router02.sol";

4: import "@uniswap/v2-periphery/contracts/UniswapV2Router02.sol";

4: import "@uniswap/v2-periphery/contracts/UniswapV2Router02.sol";

4: import "@uniswap/v2-periphery/contracts/UniswapV2Router02.sol";

```

```solidity
File: contracts/zevm/WZETA.sol

21:         balanceOf[msg.sender] += msg.value;

27:         balanceOf[msg.sender] -= wad;

49:         if (src != msg.sender && allowance[src][msg.sender] != uint(-1)) {

51:             allowance[src][msg.sender] -= wad;

54:         balanceOf[src] -= wad;

55:         balanceOf[dst] += wad;

```

```solidity
File: contracts/zevm/ZRC20.sol

3: import "./Interfaces.sol";

163:         _approve(sender, _msgSender(), currentAllowance - amount);

185:         _balances[sender] = senderBalance - amount;

186:         _balances[recipient] += amount;

194:         _totalSupply += amount;

195:         _balances[account] += amount;

205:         _balances[account] = accountBalance - amount;

206:         _totalSupply -= amount;

245:         uint256 gasFee = gasPrice * GAS_LIMIT + PROTOCOL_FLAT_FEE;

245:         uint256 gasFee = gasPrice * GAS_LIMIT + PROTOCOL_FLAT_FEE;

```

```solidity
File: contracts/zevm/interfaces/IUniswapV2Router02.sol

4: import "./IUniswapV2Router01.sol";

```

### <a name="GAS-3"></a>[GAS-3] Use Custom Errors
[Source](https://blog.soliditylang.org/2021/04/21/custom-errors/)
Instead of using error strings, to reduce deployment and runtime cost, you should use Custom Errors. This would save both deployment and runtime cost.

*Instances (1)*:
```solidity
File: contracts/evm/tools/ImmutableCreate2Factory.sol

50:         require(!_deployed[targetDeploymentAddress], "Invalid contract creation - contract has already been deployed.");

```

### <a name="GAS-4"></a>[GAS-4] Long revert strings

*Instances (1)*:
```solidity
File: contracts/evm/tools/ImmutableCreate2Factory.sol

50:         require(!_deployed[targetDeploymentAddress], "Invalid contract creation - contract has already been deployed.");

```

### <a name="GAS-5"></a>[GAS-5] Functions guaranteed to revert when called by normal users can be marked `payable`
If a function modifier such as `onlyOwner` is used, the function will revert if a normal user tries to pay the function. Marking the function as `payable` will lower the gas cost for legitimate callers because the compiler will not include checks for whether a payment was provided.

*Instances (17)*:
```solidity
File: contracts/evm/ERC20Custody.sol

80:     function updateTSSAddress(address TSSAddress_) external onlyTSSUpdater {

92:     function updateZetaFee(uint256 zetaFee_) external onlyTSS {

107:     function renounceTSSAddressUpdater() external onlyTSSUpdater {

118:     function pause() external onlyTSS {

132:     function unpause() external onlyTSS {

144:     function whitelist(IERC20 asset) external onlyTSS {

153:     function unwhitelist(IERC20 asset) external onlyTSS {

193:     function withdraw(address recipient, IERC20 asset, uint256 amount) external nonReentrant onlyTSS {

```

```solidity
File: contracts/evm/ZetaConnector.base.sol

117:     function updatePauserAddress(address pauserAddress_) external onlyPauser {

140:     function renounceTssAddressUpdater() external onlyTssUpdater {

151:     function pause() external onlyPauser {

159:     function unpause() external onlyPauser {

```

```solidity
File: contracts/evm/ZetaConnector.non-eth.sol

31:     function setMaxSupply(uint256 maxSupply_) external onlyTssAddress {

```

```solidity
File: contracts/evm/tools/ZetaInteractor.sol

54:     function setInteractorByChainId(uint256 destinationChainId, bytes calldata contractAddress) external onlyOwner {

```

```solidity
File: contracts/zevm/ZRC20.sol

270:     function updateSystemContractAddress(address addr) external onlyFungible {

279:     function updateGasLimit(uint256 gasLimit) external onlyFungible {

288:     function updateProtocolFlatFee(uint256 protocolFlatFee) external onlyFungible {

```

### <a name="GAS-6"></a>[GAS-6] Using `private` rather than `public` for constants, saves gas
If needed, the values can be read from the verified contract source code, or if there are multiple values there can be a single getter function that [returns a tuple](https://github.com/code-423n4/2022-08-frax/blob/90f55a9ce4e25bceed3a74290b854341d8de6afa/src/contracts/FraxlendPair.sol#L156-L178) of the values of all currently-public constants. Saves **3406-3606 gas** in deployment gas due to the compiler not having to create non-payable getter functions for deployment calldata, not having to store the bytes of the value outside of where it's used, and not adding another entry to the method ID table

*Instances (3)*:
```solidity
File: contracts/zevm/SystemContract.sol

31:     address public constant FUNGIBLE_MODULE_ADDRESS = 0x735b14BB79463307AAcBED86DAf3322B1e6226aB;

```

```solidity
File: contracts/zevm/ZRC20.sol

22:     address public constant FUNGIBLE_MODULE_ADDRESS = 0x735b14BB79463307AAcBED86DAf3322B1e6226aB;

```

```solidity
File: contracts/zevm/ZetaConnectorZEVM.sol

63:     address public constant FUNGIBLE_MODULE_ADDRESS = payable(0x735b14BB79463307AAcBED86DAf3322B1e6226aB);

```

### <a name="GAS-7"></a>[GAS-7] Use shift Right/Left instead of division/multiplication if possible

*Instances (4)*:
```solidity
File: contracts/evm/tools/ZetaTokenConsumerUniV2.strategy.sol

6: import "@uniswap/v2-periphery/contracts/interfaces/IUniswapV2Router02.sol";

```

```solidity
File: contracts/zevm/Uniswap.sol

4: import "@uniswap/v2-core/contracts/UniswapV2Pair.sol";

5: import "@uniswap/v2-core/contracts/UniswapV2Factory.sol";

```

```solidity
File: contracts/zevm/UniswapPeriphery.sol

4: import "@uniswap/v2-periphery/contracts/UniswapV2Router02.sol";

```

### <a name="GAS-8"></a>[GAS-8] Use != 0 instead of > 0 for unsigned integer comparison

*Instances (10)*:
```solidity
File: contracts/evm/ZetaConnector.eth.sol

63:         if (message.length > 0) {

89:         if (message.length > 0) {

```

```solidity
File: contracts/evm/ZetaConnector.non-eth.sol

72:         if (message.length > 0) {

100:         if (message.length > 0) {

```

```solidity
File: contracts/evm/tools/ZetaTokenConsumerPancakeV3.strategy.sol

218:         return pool.liquidity() > 0;

```

```solidity
File: contracts/evm/tools/ZetaTokenConsumerTrident.strategy.sol

69:         if (token0 < token1) return (token0, token1);

```

```solidity
File: contracts/evm/tools/ZetaTokenConsumerUniV2.strategy.sol

168:             return amounts[path.length - 1] > 0;

```

```solidity
File: contracts/evm/tools/ZetaTokenConsumerUniV3.strategy.sol

193:         return pool.liquidity() > 0;

```

```solidity
File: contracts/evm/tools/interfaces/TridentConcentratedLiquidityPoolFactory.sol

14: pragma solidity >=0.8.0;

```

```solidity
File: contracts/evm/tools/interfaces/TridentIPoolRouter.sol

14: pragma solidity >=0.8.0;

```


## Non Critical Issues


| |Issue|Instances|
|-|:-|:-:|
| [NC-1](#NC-1) | Return values of `approve()` not checked | 2 |
### <a name="NC-1"></a>[NC-1] Return values of `approve()` not checked
Not all IERC20 implementations `revert()` when there's a failure in `approve()`. The function signature has a boolean return value and they indicate errors that way instead. By not checking the return value, operations that should have marked as failed, may potentially go through without actually approving anything

*Instances (2)*:
```solidity
File: contracts/zevm/ZRC20.sol

146:         _approve(_msgSender(), spender, amount);

163:         _approve(sender, _msgSender(), currentAllowance - amount);

```


## Low Issues


| |Issue|Instances|
|-|:-|:-:|
| [L-1](#L-1) |  `abi.encodePacked()` should not be used with dynamic types when passing the result to a hash function such as `keccak256()` | 3 |
| [L-2](#L-2) | Use of `tx.origin` is unsafe in almost every context | 3 |
| [L-3](#L-3) | Do not use deprecated library functions | 12 |
| [L-4](#L-4) | Empty Function Body - Consider commenting why | 10 |
| [L-5](#L-5) | Unsafe ERC20 operation(s) | 6 |
| [L-6](#L-6) | Unspecific compiler version pragma | 2 |
### <a name="L-1"></a>[L-1]  `abi.encodePacked()` should not be used with dynamic types when passing the result to a hash function such as `keccak256()`
Use `abi.encode()` instead which will pad items to 32 bytes, which will [prevent hash collisions](https://docs.soliditylang.org/en/v0.8.13/abi-spec.html#non-standard-packed-mode) (e.g. `abi.encodePacked(0x123,0x456)` => `0x123456` => `abi.encodePacked(0x1,0x23456)`, but `abi.encode(0x123,0x456)` => `0x0...1230...456`). "Unless there is a compelling reason, `abi.encode` should be preferred". If there is only one argument to `abi.encodePacked()` it can often be cast to `bytes()` or `bytes32()` [instead](https://ethereum.stackexchange.com/questions/30912/how-to-compare-strings-in-solidity#answer-82739).
If all arguments are strings and or bytes, `bytes.concat()` should be used instead

*Instances (3)*:
```solidity
File: contracts/evm/tools/ImmutableCreate2Factory.sol

42:                             keccak256(abi.encodePacked(initCode)) // pass in the hash of initialization code.

121:                             keccak256(abi.encodePacked(initCode)) // pass in the hash of initialization code.

```

```solidity
File: contracts/zevm/SystemContract.sol

109:                             keccak256(abi.encodePacked(token0, token1)),

```

### <a name="L-2"></a>[L-2] Use of `tx.origin` is unsafe in almost every context
According to [Vitalik Buterin](https://ethereum.stackexchange.com/questions/196/how-do-i-make-my-dapp-serenity-proof), contracts should _not_ `assume that tx.origin will continue to be usable or meaningful`. An example of this is [EIP-3074](https://eips.ethereum.org/EIPS/eip-3074#allowing-txorigin-as-signer-1) which explicitly mentions the intention to change its semantics when it's used with new op codes. There have also been calls to [remove](https://github.com/ethereum/solidity/issues/683) `tx.origin`, and there are [security issues](solidity.readthedocs.io/en/v0.4.24/security-considerations.html#tx-origin) associated with using it for authorization. For these reasons, it's best to completely avoid the feature.

*Instances (3)*:
```solidity
File: contracts/evm/ZetaConnector.eth.sol

36:             tx.origin,

```

```solidity
File: contracts/evm/ZetaConnector.non-eth.sol

44:             tx.origin,

```

```solidity
File: contracts/zevm/ZetaConnectorZEVM.sol

99:             tx.origin,

```

### <a name="L-3"></a>[L-3] Do not use deprecated library functions

*Instances (12)*:
```solidity
File: contracts/evm/tools/ZetaTokenConsumerPancakeV3.strategy.sol

136:         IERC20(inputToken).safeApprove(address(pancakeV3Router), inputTokenAmount);

160:         IERC20(zetaToken).safeApprove(address(pancakeV3Router), zetaTokenAmount);

194:         IERC20(zetaToken).safeApprove(address(pancakeV3Router), zetaTokenAmount);

```

```solidity
File: contracts/evm/tools/ZetaTokenConsumerTrident.strategy.sol

109:         IERC20(inputToken).safeApprove(address(tridentRouter), inputTokenAmount);

145:         IERC20(zetaToken).safeApprove(address(tridentRouter), zetaTokenAmount);

176:         IERC20(zetaToken).safeApprove(address(tridentRouter), zetaTokenAmount);

```

```solidity
File: contracts/evm/tools/ZetaTokenConsumerUniV2.strategy.sol

68:         IERC20(inputToken).safeApprove(address(uniswapV2Router), inputTokenAmount);

104:         IERC20(zetaToken).safeApprove(address(uniswapV2Router), zetaTokenAmount);

134:         IERC20(zetaToken).safeApprove(address(uniswapV2Router), zetaTokenAmount);

```

```solidity
File: contracts/evm/tools/ZetaTokenConsumerUniV3.strategy.sol

108:         IERC20(inputToken).safeApprove(address(uniswapV3Router), inputTokenAmount);

133:         IERC20(zetaToken).safeApprove(address(uniswapV3Router), zetaTokenAmount);

168:         IERC20(zetaToken).safeApprove(address(uniswapV3Router), zetaTokenAmount);

```

### <a name="L-4"></a>[L-4] Empty Function Body - Consider commenting why

*Instances (10)*:
```solidity
File: contracts/evm/ZetaConnector.base.sol

166:     function send(ZetaInterfaces.SendInput calldata input) external virtual {}

179:     ) external virtual {}

193:     ) external virtual {}

```

```solidity
File: contracts/evm/ZetaConnector.eth.sol

21:     ) ZetaConnectorBase(zetaToken_, tssAddress_, tssAddressUpdater_, pauserAddress_) {}

```

```solidity
File: contracts/evm/ZetaConnector.non-eth.sol

25:     ) ZetaConnectorBase(zetaTokenAddress_, tssAddress_, tssAddressUpdater_, pauserAddress_) {}

```

```solidity
File: contracts/evm/tools/ZetaTokenConsumerPancakeV3.strategy.sol

101:     receive() external payable {}

```

```solidity
File: contracts/evm/tools/ZetaTokenConsumerTrident.strategy.sol

66:     receive() external payable {}

```

```solidity
File: contracts/evm/tools/ZetaTokenConsumerUniV3.strategy.sol

72:     receive() external payable {}

```

```solidity
File: contracts/zevm/Uniswap.sol

7: contract UniswapImports {}

```

```solidity
File: contracts/zevm/UniswapPeriphery.sol

6: contract UniswapImports {}

```

### <a name="L-5"></a>[L-5] Unsafe ERC20 operation(s)

*Instances (6)*:
```solidity
File: contracts/evm/ZetaConnector.eth.sol

32:         bool success = IERC20(zetaToken).transferFrom(msg.sender, address(this), input.zetaValueAndGas);

60:         bool success = IERC20(zetaToken).transfer(destinationAddress, zetaValue);

86:         bool success = IERC20(zetaToken).transfer(zetaTxSenderAddress, remainingZetaValue);

```

```solidity
File: contracts/zevm/WZETA.sol

28:         msg.sender.transfer(wad);

```

```solidity
File: contracts/zevm/ZRC20.sol

258:         if (!IZRC20(gasZRC20).transferFrom(msg.sender, FUNGIBLE_MODULE_ADDRESS, gasFee)) {

```

```solidity
File: contracts/zevm/ZetaConnectorZEVM.sol

94:         if (!WZETA(wzeta).transferFrom(msg.sender, address(this), input.zetaValueAndGas)) revert WZETATransferFailed();

```

### <a name="L-6"></a>[L-6] Unspecific compiler version pragma

*Instances (2)*:
```solidity
File: contracts/evm/tools/interfaces/TridentConcentratedLiquidityPoolFactory.sol

14: pragma solidity >=0.8.0;

```

```solidity
File: contracts/evm/tools/interfaces/TridentIPoolRouter.sol

14: pragma solidity >=0.8.0;

```


## Medium Issues


| |Issue|Instances|
|-|:-|:-:|
| [M-1](#M-1) | Centralization Risk for trusted owners | 4 |
### <a name="M-1"></a>[M-1] Centralization Risk for trusted owners

#### Impact:
Contracts have owners with privileged rights to perform admin tasks and need to be trusted to not perform malicious updates or drain funds.

*Instances (4)*:
```solidity
File: contracts/evm/tools/ImmutableCreate2Factory.sol

3: interface Ownable {

207:         Ownable(deploymentAddress).transferOwnership(msg.sender);

```

```solidity
File: contracts/evm/tools/ZetaInteractor.sol

9: abstract contract ZetaInteractor is Ownable2Step, ZetaInteractorErrors {

54:     function setInteractorByChainId(uint256 destinationChainId, bytes calldata contractAddress) external onlyOwner {

```

