// Code generated by suave/gen. DO NOT EDIT.
// Hash: 57e33b744fe719c804c0e3c1335a8652572947e3d147830608d8089b3d6adb37
package artifacts

import (
	"github.com/ethereum/go-ethereum/common"
)

// List of suave precompile addresses
var (
	buildEthBlockAddr            = common.HexToAddress("0x0000000000000000000000000000000042100001")
	confidentialInputsAddr       = common.HexToAddress("0x0000000000000000000000000000000042010001")
	confidentialRetrieveAddr     = common.HexToAddress("0x0000000000000000000000000000000042020001")
	confidentialStoreAddr        = common.HexToAddress("0x0000000000000000000000000000000042020000")
	earthCallAddr                = common.HexToAddress("0x0000000000000000000000000000000000004444")
	ethcallAddr                  = common.HexToAddress("0x0000000000000000000000000000000042100003")
	extractHintAddr              = common.HexToAddress("0x0000000000000000000000000000000042100037")
	fetchBidsAddr                = common.HexToAddress("0x0000000000000000000000000000000042030001")
	fillMevShareBundleAddr       = common.HexToAddress("0x0000000000000000000000000000000043200001")
	newBidAddr                   = common.HexToAddress("0x0000000000000000000000000000000042030000")
	signEthTransactionAddr       = common.HexToAddress("0x0000000000000000000000000000000040100001")
	simulateBundleAddr           = common.HexToAddress("0x0000000000000000000000000000000042100000")
	submitBundleJsonRPCAddr      = common.HexToAddress("0x0000000000000000000000000000000043000001")
	submitEthBlockBidToRelayAddr = common.HexToAddress("0x0000000000000000000000000000000042100002")
)

var SuaveMethods = map[string]common.Address{
	"buildEthBlock":            buildEthBlockAddr,
	"confidentialInputs":       confidentialInputsAddr,
	"confidentialRetrieve":     confidentialRetrieveAddr,
	"confidentialStore":        confidentialStoreAddr,
	"earthCall":                earthCallAddr,
	"ethcall":                  ethcallAddr,
	"extractHint":              extractHintAddr,
	"fetchBids":                fetchBidsAddr,
	"fillMevShareBundle":       fillMevShareBundleAddr,
	"newBid":                   newBidAddr,
	"signEthTransaction":       signEthTransactionAddr,
	"simulateBundle":           simulateBundleAddr,
	"submitBundleJsonRPC":      submitBundleJsonRPCAddr,
	"submitEthBlockBidToRelay": submitEthBlockBidToRelayAddr,
}

func PrecompileAddressToName(addr common.Address) string {
	switch addr {
	case buildEthBlockAddr:
		return "buildEthBlock"
	case confidentialInputsAddr:
		return "confidentialInputs"
	case confidentialRetrieveAddr:
		return "confidentialRetrieve"
	case confidentialStoreAddr:
		return "confidentialStore"
	case earthCallAddr:
		return "earthCall"
	case ethcallAddr:
		return "ethcall"
	case extractHintAddr:
		return "extractHint"
	case fetchBidsAddr:
		return "fetchBids"
	case fillMevShareBundleAddr:
		return "fillMevShareBundle"
	case newBidAddr:
		return "newBid"
	case signEthTransactionAddr:
		return "signEthTransaction"
	case simulateBundleAddr:
		return "simulateBundle"
	case submitBundleJsonRPCAddr:
		return "submitBundleJsonRPC"
	case submitEthBlockBidToRelayAddr:
		return "submitEthBlockBidToRelay"
	}
	return ""
}
