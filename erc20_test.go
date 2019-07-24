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
	"github.com/incognitochain/bridge-eth/incognito_proxy"
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
	privKey, c := connectAndInstantiate(t)

	// Burn
	auth := bind.NewKeyedTransactor(privKey)
	tx, err := withdraw(c.v, auth, proof)
	if err != nil {
		t.Fatal(err)
	}
	txHash := tx.Hash()
	fmt.Printf("burned erc20, txHash: %x\n", txHash[:])
}

func TestERC20Lock(t *testing.T) {
	// Get contract instance
	privKey, c := connectAndInstantiate(t)

	// Approve
	amount := int64(1e9)
	err := approveERC20(privKey, c.vAddr, c.token, amount)
	if err != nil {
		t.Fatal(err)
	}

	// Deposit
	if err := depositERC20(privKey, c.v, c.tokenAddr, amount); err != nil {
		t.Fatal(err)
	}
}

func TestERC20Deposit(t *testing.T) {
	privKey, c := connectAndInstantiate(t)

	// Deposit
	amount := int64(1e9)
	if err := depositERC20(privKey, c.v, c.tokenAddr, amount); err != nil {
		t.Fatal(err)
	}
}

func TestERC20Approve(t *testing.T) {
	// Get contract instances
	privKey, c := connectAndInstantiate(t)

	// Approve
	amount := int64(1e9)
	err := approveERC20(privKey, c.vAddr, c.token, amount)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetAllowance(t *testing.T) {
	// Get contract instances
	privKey, c := connectAndInstantiate(t)

	// Allowance
	userAddr := crypto.PubkeyToAddress(privKey.PublicKey)
	allow, err := c.token.Allowance(nil, userAddr, c.vAddr)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("allowance: %d\n", allow)

	bal, _ := c.token.BalanceOf(nil, userAddr)
	fmt.Printf("balanceOf owner: %d\n", bal)
	bal, _ = c.token.BalanceOf(nil, c.vAddr)
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
	name := "v3"
	symbol := "$"
	decimals := big.NewInt(0)
	supply := big.NewInt(1000000000000000000)
	addr, _, _, err := erc20.DeployErc20(auth, client, name, symbol, decimals, supply)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("deployed erc20")
	fmt.Printf("addr: %x\n", addr[:])
}

func connectAndInstantiate(t *testing.T) (*ecdsa.PrivateKey, *contracts) {
	privKey, client, err := connect()
	if err != nil {
		t.Fatal(err)
	}

	// Get contract instance
	c, err := instantiate(client)
	if err != nil {
		t.Fatal(err)
	}
	return privKey, c
}

func depositERC20(
	privKey *ecdsa.PrivateKey,
	v *vault.Vault,
	tokenAddr common.Address,
	amount int64,
) error {
	auth := bind.NewKeyedTransactor(privKey)
	auth.GasLimit = 1000000
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

func instantiate(client *ethclient.Client) (*contracts, error) {
	// Get contract instance
	var err error
	c := &contracts{}
	c.incAddr = common.HexToAddress(IncognitoProxyAddress)
	c.inc, err = incognito_proxy.NewIncognitoProxy(c.incAddr, client)
	if err != nil {
		return nil, err
	}

	// Vault
	c.vAddr = common.HexToAddress(VaultAddress)
	c.v, err = vault.NewVault(c.vAddr, client)
	if err != nil {
		return nil, err
	}

	// ERC20 token
	c.tokenAddr = common.HexToAddress(TokenAddress)
	c.token, err = erc20.NewErc20(c.tokenAddr, client)
	if err != nil {
		return nil, err
	}
	return c, nil
}
