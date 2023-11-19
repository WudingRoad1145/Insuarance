// SPDX-License-Identifier: MIT
pragma solidity ^0.8.8;

contract newVault {
    // Mapping to keep track of each user's balance
    mapping(address => uint256) private balances;
    // Mapping to keep track of the city each user registered
    mapping(address => string) private userCities;
    // Mapping to track total donations intended for each city
    mapping(string => uint256) private cityDonations;
    // Array to keep track of all users who have deposited
    address[] private users;

    // Event for logging deposits
    event Deposited(address indexed user, uint256 amount, string city);
    // Event for logging donations
    event Donated(string city, uint256 amount);
    // Event for logging withdrawals
    event Withdrawn(address indexed user, uint256 amount);

    // Function to deposit Ether into the vault and register a city
    function deposit(string memory city) public payable {
        require(msg.value > 0, "Deposit amount must be greater than 0");
        if (balances[msg.sender] == 0) {
            // Add user to the list only if they're depositing for the first time
            users.push(msg.sender);
        }
        balances[msg.sender] += msg.value;
        userCities[msg.sender] = city;
        emit Deposited(msg.sender, msg.value, city);
    }

    // Function to distribute donations to users from a specified city
    function distributeDonations(string memory city) public {
        uint256 totalDonation = cityDonations[city];
        require(totalDonation > 0, "No donations available for this city");

        uint256 cityBalanceTotal = 0;
        uint256[] memory userShares = new uint256[](users.length);

        // First, calculate the total balance of all users in the city
        for (uint256 i = 0; i < users.length; i++) {
            if (keccak256(bytes(userCities[users[i]])) == keccak256(bytes(city))) {
                cityBalanceTotal += balances[users[i]];
                userShares[i] = balances[users[i]];
            }
        }

        require(cityBalanceTotal > 0, "No users to distribute donations to");

        // Distribute donations proportionally based on user balance
        for (uint256 i = 0; i < users.length; i++) {
            if (userShares[i] > 0) {
                uint256 donationShare = (userShares[i] * totalDonation) / cityBalanceTotal;
                balances[users[i]] += donationShare;
                // Emit an event or log for each user's donation share (optional)
            }
        }

        // Reset the donations for the city after distribution
        cityDonations[city] = 0;
    }

    // Function to donate Ether to a city
    function donate(string memory city) public payable {
        require(msg.value > 0, "Donation amount must be greater than 0");
        cityDonations[city] += msg.value;
        emit Donated(city, msg.value);
    }

    // Function to withdraw Ether from the vault
    function withdraw(uint256 amount, address targetAccount) public {
        require(balances[msg.sender] >= amount, "Insufficient balance");
        balances[msg.sender] -= amount;
        payable(targetAccount).transfer(amount);
        emit Withdrawn(msg.sender, amount);
    }

    // Function to check the balance of a user
    function getBalance(address user) public view returns (uint256) {
        return balances[user];
    }

    // Function to check the total donations for a city
    function getCityDonations(string memory city) public view returns (uint256) {
        return cityDonations[city];
    }

    // Function to get the city registered by a user
    function getUserCity(address user) public view returns (string memory) {
        return userCities[user];
    }
}
