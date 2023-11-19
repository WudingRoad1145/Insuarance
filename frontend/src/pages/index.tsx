import Head from "next/head";
import Image from "next/image";
import styles from "@/styles/Home.module.css";
import { useState } from "react";
import earthImage from 'src/earth.jpeg'; 
import logo from 'src/logo.png';
import { ethers } from 'ethers';

export default function Home() {
	const [provider, setProvider] = useState(null);
	const [signer, setSigner] = useState(null);
	const [contract, setContract] = useState(null);
	const [userAddress, setUserAddress] = useState(null);
	const [city, setCity] = useState('');
	const [amount, setAmount] = useState('');
  
	const contractAddress = "0x9a216460B793de4eF62F3556dADae37504ce525e";
	const contractABI = [
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": true,
					"internalType": "address",
					"name": "user",
					"type": "address"
				},
				{
					"indexed": false,
					"internalType": "uint256",
					"name": "amount",
					"type": "uint256"
				},
				{
					"indexed": false,
					"internalType": "string",
					"name": "city",
					"type": "string"
				}
			],
			"name": "Deposited",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": false,
					"internalType": "string",
					"name": "city",
					"type": "string"
				},
				{
					"indexed": false,
					"internalType": "uint256",
					"name": "amount",
					"type": "uint256"
				}
			],
			"name": "Donated",
			"type": "event"
		},
		{
			"anonymous": false,
			"inputs": [
				{
					"indexed": true,
					"internalType": "address",
					"name": "user",
					"type": "address"
				},
				{
					"indexed": false,
					"internalType": "uint256",
					"name": "amount",
					"type": "uint256"
				}
			],
			"name": "Withdrawn",
			"type": "event"
		},
		{
			"inputs": [
				{
					"internalType": "string",
					"name": "city",
					"type": "string"
				}
			],
			"name": "deposit",
			"outputs": [],
			"stateMutability": "payable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "string",
					"name": "city",
					"type": "string"
				}
			],
			"name": "distributeDonations",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "string",
					"name": "city",
					"type": "string"
				}
			],
			"name": "donate",
			"outputs": [],
			"stateMutability": "payable",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "address",
					"name": "user",
					"type": "address"
				}
			],
			"name": "getBalance",
			"outputs": [
				{
					"internalType": "uint256",
					"name": "",
					"type": "uint256"
				}
			],
			"stateMutability": "view",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "string",
					"name": "city",
					"type": "string"
				}
			],
			"name": "getCityDonations",
			"outputs": [
				{
					"internalType": "uint256",
					"name": "",
					"type": "uint256"
				}
			],
			"stateMutability": "view",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "address",
					"name": "user",
					"type": "address"
				}
			],
			"name": "getUserCity",
			"outputs": [
				{
					"internalType": "string",
					"name": "",
					"type": "string"
				}
			],
			"stateMutability": "view",
			"type": "function"
		},
		{
			"inputs": [
				{
					"internalType": "uint256",
					"name": "amount",
					"type": "uint256"
				},
				{
					"internalType": "address",
					"name": "targetAccount",
					"type": "address"
				}
			],
			"name": "withdraw",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		}
	];
	const [isNetworkSwitchHighlighted, setIsNetworkSwitchHighlighted] =
		useState(false);
	const [isConnectHighlighted, setIsConnectHighlighted] = useState(false);

	const closeAll = () => {
		setIsNetworkSwitchHighlighted(false);
		setIsConnectHighlighted(false);
	};

	// Function to connect to the user's wallet
	const connectWalletHandler = async () => {
		if (window.ethereum) {
		  try {
			
			const provider = new ethers.BrowserProvider(window.ethereum);
			// It will prompt user for account connections if it isnt connected
			const signer = await provider.getSigner();
			console.log("Account:", await signer.getAddress());
			const newContract = new ethers.Contract(contractAddress, contractABI, signer);
			setProvider(provider);
			setSigner(signer);
			setContract(newContract);
			const address = await newSigner.getAddress();
			setUserAddress(address);
	
			// Request access to the user's wallet
			await provider.send("eth_requestAccounts", []);
		  } catch (err) {
			console.error(err);
		  }
		} else {
		  console.error("Please install MetaMask!");
		}
	  };
	
	  // Function to handle depositing
	  const handleDeposit = async () => {
		if (!contract) return;
	
		try {
		  const transaction = await contract.deposit(city, {
			value: ethers.utils.parseEther(amount)
		  });
		  await transaction.wait();
		  console.log('Deposit successful');
		} catch (err) {
		  console.error(err);
		}
	  };

		// Function to handle donating
		const handleDonate = async () => {
			if (!contract) return;
		
			try {
				const transaction = await contract.donate(city, {
				value: ethers.utils.parseEther(amount)
				});
				await transaction.wait();
				console.log('Deposit successful');
			} catch (err) {
				console.error(err);
			}
			};
	
	  // Function to handle withdrawal
	  const handleWithdraw = async () => {
		if (!contract) return;
	
		try {
		  const transaction = await contract.withdraw(ethers.utils.parseEther(amount), userAddress);
		  await transaction.wait();
		  console.log('Withdrawal successful');
		} catch (err) {
		  console.error(err);
		}
	  };

	return (
		<>
			<Head>
				<title>In-Sua-Rance</title>
                <meta name="description" content="Protect Yourself From the Unknown" />
                <link rel="icon" href="/favicon.ico" />
			</Head>
			<Image src={earthImage} alt="Earth" layout="fill" objectFit="cover" className={styles.earthImage} />
			<header>
				<div
					className={styles.backdrop}
					style={{
						opacity:
							isConnectHighlighted || isNetworkSwitchHighlighted
								? 1
								: 0,
					}}
				/>
				<div className={styles.header}>
					<div className={styles.logo}>
						<Image
							src={logo}
							alt="Project Logo"
							height="80"
							width="220"
						/>
					</div>
					<div className={styles.buttons}>
						<div
							onClick={closeAll}
							className={`${styles.highlight} ${
								isNetworkSwitchHighlighted
									? styles.highlightSelected
									: ``
							}`}
						>
							<w3m-network-button />
						</div>
						<div
							onClick={closeAll}
							className={`${styles.highlight} ${
								isConnectHighlighted
									? styles.highlightSelected
									: ``
							}`}
						>
							<w3m-button />
						</div>
					</div>
				</div>
			</header>
			<main className={styles.main}>
			<div className={styles.container}>
                <Image src={earthImage} alt="Earth" layout="fill" objectFit="cover" className={styles.earthImage} />
                <div className={styles.overlay}>
                    <h1 className={styles.title}>Disaster-triggered Auto-Aid-Distribution</h1>
                    <p className={styles.description}>Credible Web2 API calls using SUAVE kettles</p>
                    <p className={styles.description}>Safe AA vault with auto-distribution modules on Arbitrum Sepolia</p>
					<p className={styles.description}>Connected using WalletConnect</p>
                    <p className={styles.description}>Arx NFC Cards for off-line Signing and Spending</p>
                    <div className={styles.buttonGroup}>
					<input
						type="text"
						value={city}
						onChange={(e) => setCity(e.target.value)}
						placeholder="Enter your city"
					/>
					<input
						type="text"
						value={amount}
						onChange={(e) => setAmount(e.target.value)}
						placeholder="Enter amount to deposit/withdraw"
					/>
					<button className={styles.button} onClick={handleDeposit}>Enter Plan</button>
					<button className={styles.button} onClick={handleDonate}>DONATE</button>
					<button className={styles.button} onClick={handleWithdraw}>File Claim</button>
                    </div>
                </div>
            </div>
			</main>
		</>
	);
}
