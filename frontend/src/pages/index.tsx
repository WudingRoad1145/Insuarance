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
	  // ... Your Contract ABI here ...
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
			const newProvider = new ethers.providers.Web3Provider(window.ethereum);
			const newSigner = newProvider.getSigner();
			const newContract = new ethers.Contract(contractAddress, contractABI, newSigner);
			setProvider(newProvider);
			setSigner(newSigner);
			setContract(newContract);
			const address = await newSigner.getAddress();
			setUserAddress(address);
	
			// Request access to the user's wallet
			await newProvider.send("eth_requestAccounts", []);
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
                        <button className={styles.button}>Enter Plan</button>
                        <button className={styles.button}>DONATE</button>
						<button className={styles.button}>File Claim</button>
                    </div>
                </div>
            </div>
			</main>
		</>
	);
}
