// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

contract Target {
    event Mint(address indexed destination, uint256 amount);

    modifier onlyMinter() {
        require(
            minters[msg.sender],
            "MinterRole: caller does not have the Minter role"
        );
        _;
    }

    mapping(address => bool) minters;
    mapping(address => uint256) amounts;

    function addMinter(address minter) external {
        minters[minter] = true;
    }

    function mint(address destinationAddress, uint256 amount) external onlyMinter {
        amounts[destinationAddress] += amount;
        emit Mint(destinationAddress, amount);
    }

    function balanceOf(address destinationAddress) external returns (uint256) {
        return amounts[destinationAddress];
    }
}
