package main

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/incognitochain/bridge-eth/erc20/usdt"
)

func TestUSDTDeploy(t *testing.T) {
	privKey, client, err := connect()
	if err != nil {
		t.Error(err)
	}
	defer client.Close()

	// Deploy USDT
	auth := bind.NewKeyedTransactor(privKey)
	bal, _ := big.NewInt(1).SetString("100000000000", 10)
	addr, _, _, err := usdt.DeployUsdt(auth, client, bal, "Tether USD", "USDT", big.NewInt(6))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("deployed usdt")
	fmt.Printf("addr: %s\n", addr.Hex())
}
