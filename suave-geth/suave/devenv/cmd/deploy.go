package main

import (
	"crypto/ecdsa"
	"fmt"

	_ "embed"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/suave/e2e"
	"github.com/ethereum/go-ethereum/suave/sdk"
)

var (
	exNodeEthAddr = common.HexToAddress("b5feafbdd752ad52afb7e1bd2e40432a485bbb7f")
	exNodeNetAddr = "http://localhost:8545"

	// This account is funded in both devnev networks
	// address: 0xBE69d72ca5f88aCba033a063dF5DBe43a4148De0
	fundedAccount = newPrivKeyFromHex("91ab9a7e53c220e6210460b65a7a3bb2ca181412a8a7b43ff336b3df1737ce12")
)

var (
	simpleClaimArtifact = e2e.SimpleClaimContract
	//bundleBidContract = e2e.BundleBidContract
	//mevShareArtifact  = e2e.MevShareBidContract
)

func main() {
	rpc, _ := rpc.Dial(exNodeNetAddr)
	mevmClt := sdk.NewClient(rpc, fundedAccount.priv, exNodeEthAddr)

	var simpleClaimContract *sdk.Contract
	_ = simpleClaimContract

	txnResult, err := sdk.DeployContract(simpleClaimArtifact.Code, mevmClt)
	if err != nil {
		fmt.Errorf("Failed to deploy contract: %v", err)
	}
	receipt, err := txnResult.Wait()
	if err != nil {
		fmt.Errorf("Failed to wait for transaction result: %v", err)
	}
	if receipt.Status == 0 {
		fmt.Errorf("Failed to deploy contract: %v", err)
	}

	fmt.Printf("- SimpleClaim contract deployed: %s\n", receipt.ContractAddress)

	simpleClaimContract = sdk.GetContract(receipt.ContractAddress, simpleClaimArtifact.Abi, mevmClt)
	
	transactionResult, err := simpleClaimContract.SendTransaction("createClaim", []interface{}{1, "Istanbul", 100}, []byte{})
	if err != nil {
		fmt.Printf("Failed to send transaction: %v\n", err)
		return
	}
	
	receipt, err = transactionResult.Wait()
	if err != nil {
		fmt.Printf("Failed to wait for transaction result: %v\n", err)
		return
	}
	
	if receipt.Status == 0 {
		fmt.Println("Transaction failed")
		return
	}
	
	fmt.Printf("- createClaim succeeded: %s\n", receipt.ContractAddress)	
}

// Helpers, not unique to SUAVE

type privKey struct {
	priv *ecdsa.PrivateKey
}

func newPrivKeyFromHex(hex string) *privKey {
	key, err := crypto.HexToECDSA(hex)
	if err != nil {
		panic(fmt.Sprintf("failed to parse private key: %v", err))
	}
	return &privKey{priv: key}
}


