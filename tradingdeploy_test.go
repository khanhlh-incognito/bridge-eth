package main

// Basic imports
import (
	"fmt"
	"math/big"

	"testing"

	"github.com/incognitochain/bridge-eth/common/base58"
	"github.com/incognitochain/bridge-eth/bridge/incognito_proxy"
	"github.com/incognitochain/bridge-eth/bridge/kbntrade"
	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/incognitochain/bridge-eth/bridge/zrxtrade"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/incognitochain/bridge-eth/consensus/signatureschemes/bridgesig"
)

func initParams() *TradingTestSuite {

	tradingSuite := &TradingTestSuite{}

	// kovan params
	tradingSuite.IncBurningAddrStr = "15pABFiJVeh9D5uiQEhQX4SVibGGbdAVipQxBdxkmDqAJaoG1EdFKHBrNfs"
	tradingSuite.IncPrivKeyStr = "112t8roafGgHL1rhAP9632Yef3sx5k8xgp8cwK4MCJsCL1UWcxXvpzg97N4dwvcD735iKf31Q2ZgrAvKfVjeSUEvnzKJyyJD3GqqSZdxN4or"
	tradingSuite.IncPaymentAddrStr = "12S5Lrs1XeQLbqN4ySyKtjAjd2d7sBP2tjFijzmp6avrrkQCNFMpkXm3FPzj2Wcu2ZNqJEmh9JriVuRErVwhuQnLmWSaggobEWsBEci"

	tradingSuite.IncEtherTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000099"
	tradingSuite.IncDAITokenIDStr = "0000000000000000000000000000000000000000000000000000000000000098"
	tradingSuite.IncSAITokenIDStr = "0000000000000000000000000000000000000000000000000000000000000097"

	tradingSuite.EtherAddressStr = "0x0000000000000000000000000000000000000000"
	tradingSuite.DAIAddressStr = "0x4f96fe3b7a6cf9725f59d353f723c1bdb64ca6aa"
	tradingSuite.SAIAddressStr = "0xc4375b7de8af5a38a93548eb8453a498222c4ff2"

	tradingSuite.ETHPrivKeyStr = "B8DB29A7A43FB88AD520F762C5FDF6F1B0155637FA1E5CB2C796AFE9E5C04E31"
	tradingSuite.ETHOwnerAddrStr = "FD94c46ab8dCF0928d5113a6fEaa925793504e16"

	tradingSuite.ETHHost = "https://kovan.infura.io/v3/93fe721349134964aa71071a713c5cef"
	tradingSuite.IncBridgeHost = "http://127.0.0.1:9338"
	tradingSuite.IncRPCHost = "http://127.0.0.1:9334"
	tradingSuite.KyberContractAddr = common.HexToAddress("0xdd974d5c2e2928dea5f71b9825b8b646686bd200")
	tradingSuite.ZRXContractAddr = common.HexToAddress("0xf1ec01d6236d3cd881a0bf0130ea25fe4234003e")
	tradingSuite.WETHAddr = common.HexToAddress("0xd0a1e359811322d97991e03f863a0c30c2cf029c")

	incPriKeyBytes, _, _ := base58.Base58Check{}.Decode(tradingSuite.IncPrivKeyStr)

	tradingSuite.GeneratedPrivKeyForSC, tradingSuite.GeneratedPubKeyForSC = bridgesig.KeyGen(incPriKeyBytes)

	privKeyHex := tradingSuite.ETHPrivKeyStr
	privKey, _ := crypto.HexToECDSA(privKeyHex)

	fmt.Printf("Sign Txs with address: %s\n", crypto.PubkeyToAddress(privKey.PublicKey).Hex())

	client, _ := ethclient.Dial(tradingSuite.ETHHost)

	tradingSuite.ETHClient = client
	tradingSuite.ETHPrivKey = privKey

	return tradingSuite
}


func deployAllTradingContracts(t *testing.T) {
	tradingSuite := initParams()

	admin := common.HexToAddress(Admin)
	fmt.Println("Admin address:", admin.Hex())

	// Genesis committee
	beaconComm, bridgeComm, err := getCommittee(tradingSuite.IncBridgeHost)
	if err != nil {
		t.Error(err)
		return
	}

	// Deploy incognito_proxy
	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
	auth.Value = big.NewInt(0)
	// auth.GasPrice = big.NewInt(10000000000)
	// auth.GasLimit = 4000000
	incAddr, tx, _, err := incognito_proxy.DeployIncognitoProxy(auth, tradingSuite.ETHClient, admin, beaconComm, bridgeComm)
	if err != nil {
		t.Error(err)
		return
	}

	// incAddr := common.HexToAddress(IncognitoProxyAddress)
	fmt.Println("deployed incognito_proxy")
	fmt.Printf("addr: %s\n", incAddr.Hex())

	// Wait until tx is confirmed
	err = wait(tradingSuite.ETHClient, tx.Hash())
	if err != nil {
		t.Error(err)
		return
	}

	// Deploy vault
	prevVault := common.Address{}
	vaultAddr, tx, _, err := vault.DeployVault(auth, tradingSuite.ETHClient, admin, incAddr, prevVault)
	if err != nil {
		t.Error(err)
		return
	}
	tradingSuite.VaultAddr = vaultAddr
	fmt.Println("deployed vault")
	fmt.Printf("addr: %s\n", vaultAddr.Hex())

	// Wait until tx is confirmed
	err = wait(tradingSuite.ETHClient, tx.Hash())
	if err != nil {
		t.Error(err)
		return
	}

	// Deploy kbntrade
	kbnTradeAddr, tx, _, err := kbntrade.DeployKbntrade(auth, tradingSuite.ETHClient, tradingSuite.KyberContractAddr, tradingSuite.VaultAddr)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("deployed kbntrade")
	fmt.Printf("addr: %s\n", kbnTradeAddr.Hex())
	tradingSuite.KBNTradeDeployedAddr = kbnTradeAddr

	// Wait until tx is confirmed
	err = wait(tradingSuite.ETHClient, tx.Hash())
	if err != nil {
		t.Error(err)
		return
	}

	// Deploy 0xTrade
	zrxTradeAddr, tx, _, err := zrxtrade.DeployZrxtrade(auth, tradingSuite.ETHClient, tradingSuite.WETHAddr, tradingSuite.ZRXContractAddr, tradingSuite.VaultAddr)
	if err != nil {
		t.Error(err)
		return
	}

	// Wait until tx is confirmed
	err = wait(tradingSuite.ETHClient, tx.Hash())
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("deployed zrxTrade")
	fmt.Printf("addr: %s\n", zrxTradeAddr.Hex())
	tradingSuite.ZRXTradeDeployedAddr = zrxTradeAddr
}

func TestDeployTradingContracts(t *testing.T) {
	deployAllTradingContracts(t)
}
