package bridge

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/incognitochain/bridge-eth/erc20"
	"github.com/incognitochain/bridge-eth/vault"
)

func TestERC20Burn(t *testing.T) {
	// Get proof
	txID := ""
	proof, err := getAndDecodeBurnProof(txID)
	if err != nil {
		t.Fatal(err)
	}

	// Get contract instance
	privKey, v, _, _, _ := setupERC20Test(t)

	// Burn
	auth := bind.NewKeyedTransactor(privKey)
	auth.GasLimit = 6000000
	tx, err := v.Withdraw(
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
	fmt.Printf("burned erc20, txHash: %x\n", txHash[:])
}

func TestERC20Lock(t *testing.T) {
	// Get contract instance
	privKey, v, vaultAddr, token, tokenAddr := setupERC20Test(t)

	// Approve
	amount := int64(1e9)
	err := approveERC20(privKey, vaultAddr, token, amount)
	if err != nil {
		t.Fatal(err)
	}

	// Deposit
	if err := depositERC20(privKey, v, tokenAddr, amount); err != nil {
		t.Fatal(err)
	}
}

func TestERC20Deposit(t *testing.T) {
	privKey, v, _, _, tokenAddr := setupERC20Test(t)

	// Deposit
	amount := int64(1e9)
	if err := depositERC20(privKey, v, tokenAddr, amount); err != nil {
		t.Fatal(err)
	}
}

func TestERC20Approve(t *testing.T) {
	// Get contract instances
	privKey, _, vaultAddr, token, _ := setupERC20Test(t)

	// Approve
	amount := int64(1e9)
	err := approveERC20(privKey, vaultAddr, token, amount)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetAllowance(t *testing.T) {
	// Get contract instances
	privKey, _, vaultAddr, token, _ := setupERC20Test(t)

	// Allowance
	userAddr := crypto.PubkeyToAddress(privKey.PublicKey)
	allow, err := token.Allowance(nil, userAddr, vaultAddr)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("allowance: %d\n", allow)

	bal, _ := token.BalanceOf(nil, userAddr)
	fmt.Printf("balanceOf owner: %d\n", bal)
	bal, _ = token.BalanceOf(nil, vaultAddr)
	fmt.Printf("balanceOf spender: %d\n", bal)
}

func TestERC20Deploy(t *testing.T) {
	privKey, client, err := connect()
	if err != nil {
		t.Error(err)
	}
	defer client.Close()

	// Deploy incognito_proxy
	auth := bind.NewKeyedTransactor(privKey)
	name := "erc20 v3"
	symbol := "?"
	decimals := big.NewInt(0)
	supply := big.NewInt(1000000000000000000)
	addr, _, _, err := erc20.DeployErc20(auth, client, name, symbol, decimals, supply)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("deployed erc20")
	fmt.Printf("addr: %x\n", addr[:])
}

func setupERC20Test(t *testing.T) (*ecdsa.PrivateKey, *vault.Vault, common.Address, *erc20.Erc20, common.Address) {
	privKey, client, err := connect()
	if err != nil {
		t.Fatal(err)
	}

	// Get contract instance
	v, vaultAddr, token, tokenAddr, err := instantiate(client)
	if err != nil {
		t.Fatal(err)
	}
	return privKey, v, vaultAddr, token, tokenAddr
}

func depositERC20(
	privKey *ecdsa.PrivateKey,
	v *vault.Vault,
	tokenAddr common.Address,
	amount int64,
) error {
	auth := bind.NewKeyedTransactor(privKey)
	tx, err := v.DepositERC20(auth, tokenAddr, big.NewInt(amount), IncPaymentAddr)
	if err != nil {
		return err
	}
	txHash := tx.Hash()
	fmt.Printf("erc20 deposited, txHash: %x\n", txHash[:])
	return nil
}

func approveERC20(privKey *ecdsa.PrivateKey, spender common.Address, token *erc20.Erc20, amount int64) error {
	// Check balance
	userAddr := crypto.PubkeyToAddress(privKey.PublicKey)
	bal, _ := token.BalanceOf(nil, userAddr)
	fmt.Printf("erc20 balance: %d\n", bal)

	// Approve
	auth := bind.NewKeyedTransactor(privKey)
	tx, err := token.Approve(auth, spender, big.NewInt(amount))
	if err != nil {
		return err
	}
	txHash := tx.Hash()
	fmt.Printf("erc20 approved, txHash: %x\n", txHash[:])
	return nil
}

func instantiate(client *ethclient.Client) (*vault.Vault, common.Address, *erc20.Erc20, common.Address, error) {
	// Get contract instance
	vaultAddr := common.HexToAddress(VaultAddress)
	c, err := vault.NewVault(vaultAddr, client)
	if err != nil {
		return nil, common.Address{}, nil, common.Address{}, err
	}

	// ERC20 token
	tokenAddr := common.HexToAddress(TokenAddress)
	token, err := erc20.NewErc20(tokenAddr, client)
	if err != nil {
		return nil, common.Address{}, nil, common.Address{}, err
	}
	return c, vaultAddr, token, tokenAddr, nil
}
