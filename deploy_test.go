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
	auth.GasLimit = 6000000
	tx, err := c.Withdraw(
		auth,
		proof.instruction,

		proof.beaconHeight,
		proof.beaconInstPath,
		proof.beaconInstPathIsLeft,
		proof.beaconInstPathLen,
		proof.beaconInstRoot,
		proof.beaconBlkData,
		proof.beaconBlkHash,
		proof.beaconSignerPubkeys,
		proof.beaconSignerCount,
		proof.beaconSignerSig,
		proof.beaconSignerPaths,
		proof.beaconSignerPathIsLeft,
		proof.beaconSignerPathLen,

		proof.bridgeHeight,
		proof.bridgeInstPath,
		proof.bridgeInstPathIsLeft,
		proof.bridgeInstPathLen,
		proof.bridgeInstRoot,
		proof.bridgeBlkData,
		proof.bridgeBlkHash,
		proof.bridgeSignerPubkeys,
		proof.bridgeSignerCount,
		proof.bridgeSignerSig,
		proof.bridgeSignerPaths,
		proof.bridgeSignerPathIsLeft,
		proof.bridgeSignerPathLen,
	)
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

func TestDeployProxyAndVault(t *testing.T) {
	privKey, client, err := connect()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	// Genesis committee
	beaconCommRoot, bridgeCommRoot, err := getCommittee()
	if err != nil {
		t.Fatal(err)
	}

	// Deploy incognito_proxy
	auth := bind.NewKeyedTransactor(privKey)
	incAddr, _, _, err := incognito_proxy.DeployIncognitoProxy(auth, client, beaconCommRoot, bridgeCommRoot)
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

func getCommittee() ([32]byte, [32]byte, error) {
	emptyArr := [32]byte{}
	r, err := hex.DecodeString("72207d624ba3da378cd08466e8658a575a3706b70e20ceaa71611d9ff89e77bc")
	if err != nil {
		return emptyArr, emptyArr, err
	}
	beaconCommRoot := [32]byte{}
	copy(beaconCommRoot[:], r)
	fmt.Printf("beaconCommRoot: %x\n", beaconCommRoot[:])

	r, err = hex.DecodeString("bf7420ad8e1cf089575a77a6c04788ece2268d43adcdbd37c1e28b9628f8ab1a")
	if err != nil {
		return emptyArr, emptyArr, err
	}
	bridgeCommRoot := [32]byte{}
	copy(bridgeCommRoot[:], r)
	fmt.Printf("bridgeCommRoot: %x\n", bridgeCommRoot[:])
	return beaconCommRoot, bridgeCommRoot, nil
}
