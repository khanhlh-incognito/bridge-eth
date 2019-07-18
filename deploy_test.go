package bridge

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/incognitochain/bridge-eth/incognito_proxy"
	"github.com/incognitochain/bridge-eth/vault"
)

func TestBurn(t *testing.T) {
	txID := ""

	// Get proof
	proof, err := getAndDecodeBurnProof(txID)
	if err != nil {
		t.Fatal(err)
	}

	// Connect to ETH
	privKey, client, err := connect()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	// Get contract instance
	vaultAddr := common.HexToAddress(VaultAddress)
	c, err := vault.NewVault(vaultAddr, client)
	if err != nil {
		t.Fatal(err)
	}

	// Burn
	auth := bind.NewKeyedTransactor(privKey)
	tx, err := withdraw(c, auth, proof)
	if err != nil {
		t.Fatal(err)
	}
	txHash := tx.Hash()
	fmt.Printf("burned, txHash: %x\n", txHash[:])
}

func TestDeposit(t *testing.T) {
	privKey, client, err := connect()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	// Get contract instance
	vaultAddr := common.HexToAddress(VaultAddress)
	c, err := vault.NewVault(vaultAddr, client)
	if err != nil {
		t.Fatal(err)
	}

	// Deposit
	auth := bind.NewKeyedTransactor(privKey)
	auth.Value = big.NewInt(0.1 * params.Ether)
	tx, err := c.Deposit(auth, IncPaymentAddr)
	if err != nil {
		t.Fatal(err)
	}
	txHash := tx.Hash()
	fmt.Printf("deposited, txHash: %x\n", txHash[:])
}

func TestGetCommittee(t *testing.T) {
	_, c := connectAndInstantiate(t)
	beaconBlk, _ := c.inc.LatestBeaconBlk(nil)
	for {
		pubkeys, err := c.inc.BeaconCommPubkeys(nil, beaconBlk)
		if err != nil {
			t.Fatal(err)
		}
		prevBeaconBlk, err := c.inc.BeaconCommPrevBlk(nil, beaconBlk)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("beaconBlk: %d %x\n", beaconBlk, pubkeys)

		if beaconBlk.Uint64() == 0 {
			break
		}
		beaconBlk = prevBeaconBlk
	}
	bridgeBlk, _ := c.inc.LatestBridgeBlk(nil)
	for {
		pubkeys, err := c.inc.BridgeCommPubkeys(nil, beaconBlk)
		if err != nil {
			t.Fatal(err)
		}
		prevBridgeBlk, err := c.inc.BridgeCommPrevBlk(nil, bridgeBlk)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("bridgeBlk: %d %x\n", bridgeBlk, pubkeys)

		if bridgeBlk.Uint64() == 0 {
			break
		}
		bridgeBlk = prevBridgeBlk
	}
}

func TestDeployProxyAndVault(t *testing.T) {
	privKey, client, err := connect()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	// Genesis committee
	beaconComm, bridgeComm, err := getCommittee()
	if err != nil {
		t.Fatal(err)
	}

	// Deploy MulSigP256
	// msAddr, _, _, err := checkMulSig.DeployMulSigP256(auth, client)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// fmt.Println("deployed MulSigP256")
	msAddr := common.HexToAddress("0x1383AA86AC1f7BF638baAD2311448a40493f691F")
	fmt.Printf("addr: %x\n", msAddr[:])

	// Deploy incognito_proxy
	auth := bind.NewKeyedTransactor(privKey)
	incAddr, _, _, err := incognito_proxy.DeployIncognitoProxy(auth, client, beaconComm, bridgeComm, msAddr)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("deployed incognito_proxy")
	fmt.Printf("addr: %x\n", incAddr[:])

	// Deploy vault
	vaultAddr, _, _, err := vault.DeployVault(auth, client, incAddr)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("deployed vault")
	fmt.Printf("addr: %x\n", vaultAddr[:])
}

func connect() (*ecdsa.PrivateKey, *ethclient.Client, error) {
	privKeyHex := os.Getenv("PRIVKEY")
	privKey, err := crypto.HexToECDSA(privKeyHex)
	if err != nil {
		return nil, nil, err
	}

	client, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil {
		return nil, nil, err
	}

	return privKey, client, nil
}

func getCommittee() ([]byte, []byte, error) {
	beaconComm := []string{"02a96a04ad76a0034efc8819e93308823ce7a3b76fd694f961ee909124096baf00", "0242653de0e9af9dd3725008519157314eb5a845dec2cd646ce9e03f780175b700", "028c49fc5f3e001c36095335c53b0b7320f6a1c932424e92c9de344b55e80ddf00"}
	beacons := []byte{}
	for _, p := range beaconComm {
		pk, _ := hex.DecodeString(p)
		beacons = append(beacons, pk...)
	}

	bridgeComm := []string{"0253d262c2b6a55606ff9d32e195231ec57e4d23a6efd1c02143a58fd0c2591d01", "02dee56cbbde5ef6d03a9e69bf3784ae4a8460d0058a6082eee4be2ed5c4fd3301", "02ec388db662801da0fe3c41f39085369ed4df71d42ec96924012243dc9c67d201"}
	bridges := []byte{}
	for _, p := range bridgeComm {
		pk, _ := hex.DecodeString(p)
		bridges = append(bridges, pk...)
	}
	return beacons, bridges, nil
}
