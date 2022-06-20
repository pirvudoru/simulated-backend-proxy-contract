// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "./Target.sol";

contract Proxy {
    Target _target;

    constructor(address targetAddress) {
        _target = Target(targetAddress);
    }

    function mint(address destinationAddress) external {
        _target.mint(destinationAddress, 42);
    }
}
