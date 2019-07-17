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
	beaconComm := []string{"035e324cb260cf972a0b3cbe4d2bf2b5b01e5746712474101839beeaba8a15cd00", "02f302e420f5947118dce4d7f4bc6f018da2c04f4ac199d3ed9e15dbdbdf2f2a00", "02688b24164f8ee7b0df258dccb4776fe846c2ad41d32c054469887331e4aa2500"}
	beacons := []byte{}
	for _, p := range beaconComm {
		pk, _ := hex.DecodeString(p)
		beacons = append(beacons, pk...)
	}

	bridgeComm := []string{"03b2c0633aa510738c7847611f6ce6aa2ea4f2106476c2df4343b3e15e608105d9", "02b70062bbb139729ede0ef1273205e1f12c43af9a79d9da97f09cea11a865d301", "021c0d087fe263012e41f3319c6e40c2484d78d000596d05d3207f21bb13ef4601"}
	bridges := []byte{}
	for _, p := range bridgeComm {
		pk, _ := hex.DecodeString(p)
		bridges = append(bridges, pk...)
	}
	return beacons, bridges, nil
}
