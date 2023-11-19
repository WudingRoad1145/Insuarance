pragma solidity ^0.8.8;

import "../libraries/Suave.sol";

contract ExampleEthCallSource {
    function callTarget(address target, uint256 expected) public view{
        bytes memory output = Suave.ethcall(target, abi.encodeWithSignature("get()"));
        (uint256 num) = abi.decode(output, (uint64));
        require(num == expected);
    }
    function example(string calldata location) public {
        Suave.earthCall(location);
    }
}

contract X {
    function createClaim(uint64 earth) public {
        // Function implementation
    }

    function settle(string calldata location) public {
        Suave.earthCall(location);
    }
}

contract ExampleEthCallTarget {
    function get() public view returns (uint256) {
        return 101;
    }
}
