package main

import (
	"crypto/ecdsa"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/incognitochain/bridge-eth/bridge/vault"
)

func TestPauseVault(t *testing.T) {
	privKey, c := getVault(t)

	// Pause vault
	auth := bind.NewKeyedTransactor(privKey)
	_, err := c.Pause(auth)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMigrateVault(t *testing.T) {
	privKey, c := getVault(t)

	// Migrate vault
	newVault := common.HexToAddress("")
	auth := bind.NewKeyedTransactor(privKey)
	_, err := c.Migrate(auth, newVault)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMoveAssetsVault(t *testing.T) {
	privKey, c := getVault(t)

	// Migrate vault
	assets := []common.Address{
		common.Address{}, // ETH
	}
	auth := bind.NewKeyedTransactor(privKey)
	_, err := c.MoveAssets(auth, assets)
	if err != nil {
		t.Fatal(err)
	}
}

func getVault(t *testing.T) (*ecdsa.PrivateKey, *vault.Vault) {
	privKey, client, err := connect()
	if err != nil {
		t.Fatal(err)
	}

	// Get vault instance
	vaultAddr := common.HexToAddress(VaultAddress)
	c, err := vault.NewVault(vaultAddr, client)
	if err != nil {
		t.Fatal(err)
	}
	return privKey, c
}
