package bridge

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/incognitochain/bridge-eth/incognito_proxy"
	"github.com/incognitochain/bridge-eth/vault"
	"github.com/pkg/errors"
)

func TestMassSend(t *testing.T) {
	addrs := []string{
		"0xDF1A9BE4dA9Ed6CDC28bea3c23E81B8a3e4480Ae",
		"0x354e2c1ee8f254f379A17679Dd14e3afa61c0346",
		"0x9a6A22438307C68A794485b86Faa6b72Aa67Ded7",
		"0x7A279AEe9cc310B64F0F159904271c0a68014082",
	}

	privKey, client, err := connect()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	// Deposit
	nonce, err := client.NonceAt(context.Background(), crypto.PubkeyToAddress(privKey.PublicKey), nil)
	if err != nil {
		t.Fatal(err)
	}

	for i, addr := range addrs {
		txHash, err := transfer(client, privKey, addr, nonce+uint64(i))
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("sent, txHash: %s\n", txHash)
	}
}

func transfer(
	client *ethclient.Client,
	privKey *ecdsa.PrivateKey,
	to string,
	nonce uint64,
) (string, error) {
	toAddress := common.HexToAddress(to)
	value := big.NewInt(0.1 * params.Ether)
	gasLimit := uint64(30000)
	gasPrice := big.NewInt(20000000000)
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", errors.WithStack(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privKey)
	if err != nil {
		return "", errors.WithStack(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return signedTx.Hash().String(), nil
}

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
	auth.GasPrice = big.NewInt(20000000000)
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
	// url := "http://test-node.incognito.org:9334"
	// url := "http://0.0.0.0:9334"
	// numBeaconVals, beaconComm, numBridgeVals, bridgeComm, err := getCommittee(url)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	numBeaconVals, beaconComm, numBridgeVals, bridgeComm := getCommitteeHardcoded()

	// Deploy MulSigP256
	// msAddr, _, _, err := checkMulSig.DeployMulSigP256(auth, client)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// fmt.Println("deployed MulSigP256")
	msAddr := common.HexToAddress(MulSigP256)
	fmt.Printf("addr: %x\n", msAddr[:])

	// // Deploy incognito_proxy
	auth := bind.NewKeyedTransactor(privKey)
	auth.GasPrice = big.NewInt(20000000000)
	incAddr, _, _, err := incognito_proxy.DeployIncognitoProxy(auth, client, numBeaconVals, beaconComm, numBridgeVals, bridgeComm, msAddr)
	if err != nil {
		t.Fatal(err)
	}
	// incAddr := common.HexToAddress(IncognitoProxyAddress)
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

	client, err := ethclient.Dial("https://kovan.infura.io/v3/29fead42346b4bfa88dd5fd7e56b6406")
	if err != nil {
		return nil, nil, err
	}

	return privKey, client, nil
}
