// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface Victim {
    function deposit(bytes calldata recipient, address asset, uint256 amount, bytes calldata message) external;

    function withdraw(address recipient, address asset, uint256 amount) external;
}

contract AttackerContract {
    address public victimContractAddress;
    uint256 private _victimMethod;

    constructor(address victimContractAddress_, address wzeta, uint256 victimMethod) {
        victimContractAddress = victimContractAddress_;
        IERC20(wzeta).approve(victimContractAddress, type(uint256).max);
        _victimMethod = victimMethod;
    }

    // Fallback function to receive ETH
    receive() external payable {}

    function attackDeposit() internal {
        Victim(victimContractAddress).deposit("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266", address(this), 0, "0x00");
    }

    function attackWidrawal() internal {
        Victim(victimContractAddress).withdraw(0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266, address(this), 0);
    }

    function attack() internal {
        if (_victimMethod == 1) {
            attackDeposit();
        } else if (_victimMethod == 2) {
            attackWidrawal();
        }
    }

    function balanceOf(address account) external returns (uint256) {
        attack();
        return 0;
    }

    function transferFrom(address from, address to, uint256 amount) public returns (bool) {
        attack();
        return true;
    }

    function transfer(address to, uint256 amount) public returns (bool) {
        attack();
        return true;
    }
}
