package bridge

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/incognitochain/bridge-eth/erc20"
)

const TokenAddress = "1809edaddbd1a039b2c549129bb51ce1303cabee"

func TestDeployERC20(t *testing.T) {
	privKey, client, err := connect()
	if err != nil {
		t.Error(err)
	}
	defer client.Close()

	// Deploy incognito_proxy
	auth := bind.NewKeyedTransactor(privKey)
	name := "Super duper erc20"
	symbol := "="
	decimals := big.NewInt(0)
	supply := big.NewInt(1000000000000000000)
	addr, _, _, err := erc20.DeployErc20(auth, client, name, symbol, decimals, supply)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("deployed erc20")
	fmt.Printf("addr: %x\n", addr[:])
}
