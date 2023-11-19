// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.8;

import "../libraries/Suave.sol";

contract SimpleClaimContract {

    event SuccessfulClaimEvent(
        address indexed claimer,
		string location
	);

    function createClaim(uint64 threshold, string calldata location, uint256 amount) public {
        string memory realLifeData = Suave.earthCall(location);
        require(bytes(realLifeData).length > 0, "Real life data is empty :(");
        // check the vault/contract on L1 is ready to claim
        uint256 balance = SimpleVault.getBalance(msg.sender);
        require(balance > amount);
        emit SimpleClaimContract.SuccessfulClaimEvent(msg.sender, location);
    }

    // Function to settle
    function settle(string calldata rpc, uint256 amount, address payable targetL1Address) public {
        // Interact with the SimpleVault contract
        SimpleVault.withdraw(rpc, amount, targetL1Address, SimpleVault.vault);
    }

    // Fallback function to receive Ether
    receive() external payable {}
}

library SimpleVault {
    address constant vault = 0xF64385857Aa608Fa0b025b9Ac5d25720ccD1e683;
    string constant getBalanceSig = "getBalance((address))";
    string constant getWithdrawSig = "withdraw((uint256, address))";

    function getBalance(address claimer)
        internal
        view
        returns (uint256 balance)
    {
        bytes memory output = Suave.ethcall(vault, abi.encodeWithSignature(getBalanceSig, claimer));
        balance = abi.decode(output, (uint64));
        return balance;
    }

    function withdraw(string memory builderUrl, uint256 amount, address targetL1Address, address vaultContractAddress) public {
        // Encode the function call to the SimpleVault contract
        bytes memory callData = abi.encodeWithSignature("withdraw(uint256,address)", amount, targetL1Address);

        // Construct the JSON-RPC request
        string memory jsonRpcRequest = string(abi.encodePacked(
            '{"jsonrpc":"2.0","method":"eth_sendTransaction","params":[{"from": "', 
            toHexString(msg.sender),
            '", "to": "', 
            toHexString(vaultContractAddress),
            '", "data": "', 
            toHexString(callData), 
            '"}], "id":1}'
        ));

        // Convert the JSON string to bytes
        bytes memory requestData = bytes(jsonRpcRequest);

        // Call submitBundleJsonRPC with the encoded request
        Suave.submitBundleJsonRPC(builderUrl, "eth_sendBundle", requestData);
    }
    
    // Utility function to convert an address to a hex string
    function toHexString(address account) internal pure returns (string memory) {
        return toHexString(abi.encodePacked(account), 20);
    }

    // Utility function to convert bytes to a hex string
    function toHexString(bytes memory data) internal pure returns (string memory) {
        return toHexString(data, data.length);
    }

    function toHexString(bytes memory data, uint256 length) internal pure returns (string memory) {
        bytes memory buffer = new bytes(2 * length + 2);
        buffer[0] = '0';
        buffer[1] = 'x';
        for (uint256 i = 0; i < length; ++i) {
            buffer[2 * i + 2] = byteToHexChar(uint8(data[i] >> 4));
            buffer[2 * i + 3] = byteToHexChar(uint8(data[i] & 0x0F));
        }
        return string(buffer);
    }

    // Utility function to convert a byte to a hex character
    function byteToHexChar(uint8 b) internal pure returns (bytes1) {
        if (b < 10) return bytes1(b + 0x30);
        else return bytes1(b + 0x61 - 10);
    }
}