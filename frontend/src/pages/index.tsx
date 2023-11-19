import Head from "next/head";
import Image from "next/image";
import styles from "@/styles/Home.module.css";
import { useState } from "react";
import earthImage from 'src/earth.jpeg'; 
import logo from 'src/logo.png';


export default function Home() {
	
	const [isNetworkSwitchHighlighted, setIsNetworkSwitchHighlighted] =
		useState(false);
	const [isConnectHighlighted, setIsConnectHighlighted] = useState(false);

	const closeAll = () => {
		setIsNetworkSwitchHighlighted(false);
		setIsConnectHighlighted(false);
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
