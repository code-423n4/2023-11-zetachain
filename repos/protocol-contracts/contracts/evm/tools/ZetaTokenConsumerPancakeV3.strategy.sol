// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "@openzeppelin/contracts/interfaces/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@uniswap/v3-periphery/contracts/interfaces/ISwapRouter.sol";
import "@uniswap/v3-core/contracts/interfaces/IUniswapV3Factory.sol";
import "@uniswap/v3-core/contracts/interfaces/IUniswapV3Pool.sol";

import "../interfaces/ZetaInterfaces.sol";

interface ZetaTokenConsumerUniV3Errors {
    error InputCantBeZero();

    error ErrorSendingETH();

    error ReentrancyError();
}

interface WETH9 {
    function withdraw(uint256 wad) external;
}

interface ISwapRouterPancake is IUniswapV3SwapCallback {
    struct ExactInputSingleParams {
        address tokenIn;
        address tokenOut;
        uint24 fee;
        address recipient;
        uint256 amountIn;
        uint256 amountOutMinimum;
        uint160 sqrtPriceLimitX96;
    }

    /// @notice Swaps `amountIn` of one token for as much as possible of another token
    /// @param params The parameters necessary for the swap, encoded as `ExactInputSingleParams` in calldata
    /// @return amountOut The amount of the received token
    function exactInputSingle(ExactInputSingleParams calldata params) external payable returns (uint256 amountOut);

    struct ExactInputParams {
        bytes path;
        address recipient;
        uint256 amountIn;
        uint256 amountOutMinimum;
    }

    /// @notice Swaps `amountIn` of one token for as much as possible of another along the specified path
    /// @param params The parameters necessary for the multi-hop swap, encoded as `ExactInputParams` in calldata
    /// @return amountOut The amount of the received token
    function exactInput(ExactInputParams calldata params) external payable returns (uint256 amountOut);
}

/**
 * @dev Uniswap V3 strategy for ZetaTokenConsumer
 */
contract ZetaTokenConsumerPancakeV3 is ZetaTokenConsumer, ZetaTokenConsumerUniV3Errors {
    using SafeERC20 for IERC20;
    uint256 internal constant MAX_DEADLINE = 200;

    uint24 public immutable zetaPoolFee;
    uint24 public immutable tokenPoolFee;

    address public immutable WETH9Address;
    address public immutable zetaToken;

    ISwapRouterPancake public immutable pancakeV3Router;
    IUniswapV3Factory public immutable uniswapV3Factory;

    bool internal _locked;

    constructor(
        address zetaToken_,
        address pancakeV3Router_,
        address uniswapV3Factory_,
        address WETH9Address_,
        uint24 zetaPoolFee_,
        uint24 tokenPoolFee_
    ) {
        if (
            zetaToken_ == address(0) ||
            pancakeV3Router_ == address(0) ||
            uniswapV3Factory_ == address(0) ||
            WETH9Address_ == address(0)
        ) revert ZetaCommonErrors.InvalidAddress();

        zetaToken = zetaToken_;
        pancakeV3Router = ISwapRouterPancake(pancakeV3Router_);
        uniswapV3Factory = IUniswapV3Factory(uniswapV3Factory_);
        WETH9Address = WETH9Address_;
        zetaPoolFee = zetaPoolFee_;
        tokenPoolFee = tokenPoolFee_;
    }

    modifier nonReentrant() {
        if (_locked) revert ReentrancyError();
        _locked = true;
        _;
        _locked = false;
    }

    receive() external payable {}

    function getZetaFromEth(
        address destinationAddress,
        uint256 minAmountOut
    ) external payable override returns (uint256) {
        if (destinationAddress == address(0)) revert ZetaCommonErrors.InvalidAddress();
        if (msg.value == 0) revert InputCantBeZero();

        ISwapRouterPancake.ExactInputSingleParams memory params = ISwapRouterPancake.ExactInputSingleParams({
            tokenIn: WETH9Address,
            tokenOut: zetaToken,
            fee: zetaPoolFee,
            recipient: destinationAddress,
            amountIn: msg.value,
            amountOutMinimum: minAmountOut,
            sqrtPriceLimitX96: 0
        });

        uint256 amountOut = pancakeV3Router.exactInputSingle{value: msg.value}(params);

        emit EthExchangedForZeta(msg.value, amountOut);
        return amountOut;
    }

    function getZetaFromToken(
        address destinationAddress,
        uint256 minAmountOut,
        address inputToken,
        uint256 inputTokenAmount
    ) external override returns (uint256) {
        if (destinationAddress == address(0) || inputToken == address(0)) revert ZetaCommonErrors.InvalidAddress();
        if (inputTokenAmount == 0) revert InputCantBeZero();

        IERC20(inputToken).safeTransferFrom(msg.sender, address(this), inputTokenAmount);
        IERC20(inputToken).safeApprove(address(pancakeV3Router), inputTokenAmount);

        ISwapRouterPancake.ExactInputParams memory params = ISwapRouterPancake.ExactInputParams({
            path: abi.encodePacked(inputToken, tokenPoolFee, WETH9Address, zetaPoolFee, zetaToken),
            recipient: destinationAddress,
            amountIn: inputTokenAmount,
            amountOutMinimum: minAmountOut
        });

        uint256 amountOut = pancakeV3Router.exactInput(params);

        emit TokenExchangedForZeta(inputToken, inputTokenAmount, amountOut);
        return amountOut;
    }

    function getEthFromZeta(
        address destinationAddress,
        uint256 minAmountOut,
        uint256 zetaTokenAmount
    ) external override returns (uint256) {
        if (destinationAddress == address(0)) revert ZetaCommonErrors.InvalidAddress();
        if (zetaTokenAmount == 0) revert InputCantBeZero();

        IERC20(zetaToken).safeTransferFrom(msg.sender, address(this), zetaTokenAmount);
        IERC20(zetaToken).safeApprove(address(pancakeV3Router), zetaTokenAmount);

        ISwapRouterPancake.ExactInputSingleParams memory params = ISwapRouterPancake.ExactInputSingleParams({
            tokenIn: zetaToken,
            tokenOut: WETH9Address,
            fee: zetaPoolFee,
            recipient: address(this),
            amountIn: zetaTokenAmount,
            amountOutMinimum: minAmountOut,
            sqrtPriceLimitX96: 0
        });

        uint256 amountOut = pancakeV3Router.exactInputSingle(params);

        WETH9(WETH9Address).withdraw(amountOut);

        emit ZetaExchangedForEth(zetaTokenAmount, amountOut);

        (bool sent, ) = destinationAddress.call{value: amountOut}("");
        if (!sent) revert ErrorSendingETH();

        return amountOut;
    }

    function getTokenFromZeta(
        address destinationAddress,
        uint256 minAmountOut,
        address outputToken,
        uint256 zetaTokenAmount
    ) external override nonReentrant returns (uint256) {
        if (destinationAddress == address(0) || outputToken == address(0)) revert ZetaCommonErrors.InvalidAddress();
        if (zetaTokenAmount == 0) revert InputCantBeZero();

        IERC20(zetaToken).safeTransferFrom(msg.sender, address(this), zetaTokenAmount);
        IERC20(zetaToken).safeApprove(address(pancakeV3Router), zetaTokenAmount);

        ISwapRouterPancake.ExactInputParams memory params = ISwapRouterPancake.ExactInputParams({
            path: abi.encodePacked(zetaToken, zetaPoolFee, WETH9Address, tokenPoolFee, outputToken),
            recipient: destinationAddress,
            amountIn: zetaTokenAmount,
            amountOutMinimum: minAmountOut
        });

        uint256 amountOut = pancakeV3Router.exactInput(params);

        emit ZetaExchangedForToken(outputToken, zetaTokenAmount, amountOut);
        return amountOut;
    }

    function hasZetaLiquidity() external view override returns (bool) {
        address poolAddress = uniswapV3Factory.getPool(WETH9Address, zetaToken, zetaPoolFee);

        if (poolAddress == address(0)) {
            return false;
        }

        //@dev: if pool does exist, get its liquidity
        IUniswapV3Pool pool = IUniswapV3Pool(poolAddress);
        return pool.liquidity() > 0;
    }
}
