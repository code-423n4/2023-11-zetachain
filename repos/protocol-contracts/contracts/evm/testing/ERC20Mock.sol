// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

/**
 * @dev ZetaEth is an implementation of OpenZeppelin's ERC20
 */
contract ERC20Mock is ERC20 {
    constructor(string memory name, string memory symbol, address creator, uint256 initialSupply) ERC20(name, symbol) {
        _mint(creator, initialSupply * (10 ** uint256(decimals())));
    }
}
