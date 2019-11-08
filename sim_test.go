package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/incognitochain/bridge-eth/bridge/incognito_proxy"
	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/incognitochain/bridge-eth/erc20"
	"github.com/pkg/errors"
)

var auth *bind.TransactOpts
var genesisAcc *account

type Platform struct {
	*contracts
	sim *backends.SimulatedBackend
}

func init() {
	fmt.Println("Initializing genesis account...")
	genesisAcc = loadAccount()
	auth = bind.NewKeyedTransactor(genesisAcc.PrivateKey)
}

func TestSimulatedSwapBridge(t *testing.T) {
	p, err := setupWithHardcodedCommittee()
	// _, err := setupWithLocalCommittee()
	if err != nil {
		t.Fatalf("%+v", err)
	}

	blocks := []int{10, 20, 30}
	for _, b := range blocks {
		url := "http://54.39.158.106:19032"
		proof, err := getAndDecodeBridgeSwapProof(url, b)
		if err != nil {
			t.Fatal(err)
		}

		auth.GasLimit = 7000000
		fmt.Printf("inst len: %d\n", len(proof.Instruction))
		tx, err := SwapBridge(p.inc, auth, proof)
		if err != nil {
			fmt.Println("err:", err)
		}
		p.sim.Commit()
		printReceipt(p.sim, tx)
	}
}

func TestSimulatedSwapBeacon(t *testing.T) {
	body := getBeaconSwapProof(20)
	if len(body) < 1 {
		t.Fatal(fmt.Errorf("empty beacon swap proof"))
	}

	r := getProofResult{}
	if err := json.Unmarshal([]byte(body), &r); err != nil {
		t.Fatalf("%+v", err)
	}
	if len(r.Result.Instruction) == 0 {
		t.Fatal("invalid swap proof")
	}
	proof, err := decodeProof(&r)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	// a, _ := json.Marshal(proof)
	// fmt.Printf("proof: %s\n", string(a))

	p, err := setupWithHardcodedCommittee()
	// p, err := setupWithLocalCommittee()
	if err != nil {
		t.Fatalf("%+v", err)
	}

	auth.GasLimit = 7000000
	fmt.Printf("inst len: %d\n", len(proof.Instruction))
	tx, err := SwapBeacon(p.inc, auth, proof)
	if err != nil {
		fmt.Println("err:", err)
	}
	p.sim.Commit()
	printReceipt(p.sim, tx)
}

func TestSimulatedBurnETH(t *testing.T) {
	proof, err := getAndDecodeBurnProof("3056832abff4fae1ed18163ded4d24cd94c1a6f1dc2ee0819170c85d508b7266")
	if err != nil {
		t.Fatal(err)
	}
	// a, _ := json.Marshal(proof)
	// fmt.Println(string(a))

	p, err := setupWithHardcodedCommittee()
	// p, err := setupWithLocalCommittee()
	if err != nil {
		t.Fatalf("%+v", err)
	}

	oldBalance, newBalance, err := deposit(p, big.NewInt(int64(5e18)))
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("deposit to vault: %d -> %d\n", oldBalance, newBalance)

	withdrawer := common.HexToAddress("0xe722D8b71DCC0152D47D2438556a45D3357d631f")
	fmt.Printf("withdrawer init balance: %d\n", p.getBalance(withdrawer))

	auth.GasLimit = 8000000
	tx, err := Withdraw(p.v, auth, proof)
	if err != nil {
		fmt.Println("err:", err)
	}
	p.sim.Commit()
	printReceipt(p.sim, tx)

	fmt.Printf("withdrawer new balance: %d\n", p.getBalance(withdrawer))
}

func TestSimulatedBurnERC20(t *testing.T) {
	proof, err := getAndDecodeBurnProof("5da2ee413a5d78f25f016fd41994875054afd64d776b95f76515489c9e0f5a13")
	if err != nil {
		t.Fatal(err)
	}
	// a, _ := json.Marshal(proof)
	// fmt.Println(string(a))

	p, err := setupWithHardcodedCommittee()
	// p, err := setupWithLocalCommittee()
	if err != nil {
		t.Fatalf("%+v", err)
	}

	oldBalance, newBalance, err := lockSimERC20WithBalance(p, p.token, p.tokenAddr, big.NewInt(int64(1e9)))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("deposit erc20 to vault: %d -> %d\n", oldBalance, newBalance)

	withdrawer := common.HexToAddress("0xe722D8b71DCC0152D47D2438556a45D3357d631f")
	fmt.Printf("withdrawer init balance: %d\n", getBalanceERC20(p.token, withdrawer))

	auth.GasLimit = 8000000
	tx, err := Withdraw(p.v, auth, proof)
	if err != nil {
		fmt.Println("err:", err)
	}
	p.sim.Commit()
	printReceipt(p.sim, tx)

	fmt.Printf("withdrawer new balance: %d\n", getBalanceERC20(p.token, withdrawer))
}

func lockSimERC20WithTxs(
	p *Platform,
	token *erc20.Erc20,
	tokenAddr common.Address,
	amount *big.Int,
) (*types.Transaction, *types.Transaction, error) {
	txApprove, err := approveERC20(genesisAcc.PrivateKey, p.vAddr, token, amount)
	if err != nil {
		return nil, nil, err
	}
	p.sim.Commit()

	txDeposit, err := depositERC20(genesisAcc.PrivateKey, p.v, tokenAddr, amount)
	if err != nil {
		return txApprove, nil, err
	}
	p.sim.Commit()
	return txApprove, txDeposit, nil
}

func lockSimERC20WithBalance(
	p *Platform,
	token *erc20.Erc20,
	tokenAddr common.Address,
	amount *big.Int,
) (*big.Int, *big.Int, error) {
	initBalance := getBalanceERC20(token, p.vAddr)
	fmt.Printf("bal: %d\n", getBalanceERC20(token, genesisAcc.Address))
	if _, _, err := lockSimERC20WithTxs(p, token, tokenAddr, amount); err != nil {
		return nil, nil, err
	}
	newBalance := getBalanceERC20(token, p.vAddr)
	return initBalance, newBalance, nil
}

func getBalanceERC20(token *erc20.Erc20, addr common.Address) *big.Int {
	bal, err := token.BalanceOf(nil, addr)
	if err != nil {
		return big.NewInt(-1)
	}
	return bal
}

func (p *Platform) getBalance(addr common.Address) *big.Int {
	bal, _ := p.sim.BalanceAt(context.Background(), addr, nil)
	return bal
}

func setup(
	beaconComm []common.Address,
	bridgeComm []common.Address,
	decimals []int,
	accs ...common.Address,
) (*Platform, error) {
	alloc := make(core.GenesisAlloc)
	balance, _ := big.NewInt(1).SetString("1000000000000000000000000000000", 10) // 1E30 wei
	alloc[auth.From] = core.GenesisAccount{Balance: balance}
	for _, acc := range accs {
		alloc[acc] = core.GenesisAccount{Balance: balance}
	}
	sim := backends.NewSimulatedBackend(alloc, 8000000)
	p := &Platform{sim: sim, contracts: &contracts{tokens: map[int]tokenInfo{}}}

	var err error
	var tx *types.Transaction
	_ = tx

	// ERC20: always deploy first so its address is fixed
	p.tokenAddr, tx, p.token, err = erc20.DeployErc20(auth, sim, "MyErc20", "ERC", big.NewInt(0), big.NewInt(int64(1e18)))
	if err != nil {
		return nil, fmt.Errorf("failed to deploy ERC20 contract: %v", err)
	}
	// fmt.Printf("token addr: %s\n", p.tokenAddr.Hex())
	sim.Commit()

	// Deploy erc20s with different decimals to test
	ercBal := big.NewInt(20)
	ercBal = ercBal.Mul(ercBal, big.NewInt(int64(1e18)))
	ercBal = ercBal.Mul(ercBal, big.NewInt(int64(1e18)))
	for _, d := range decimals {
		tokenAddr, _, token, err := erc20.DeployErc20(auth, sim, "MyErc20", "ERC", big.NewInt(int64(d)), ercBal)
		if err != nil {
			return nil, fmt.Errorf("failed to deploy ERC20 contract: %v", err)
		}
		p.tokens[d] = tokenInfo{c: token, addr: tokenAddr}
	}
	sim.Commit()

	// IncognitoProxy
	admin := auth.From
	p.incAddr, tx, p.inc, err = incognito_proxy.DeployIncognitoProxy(auth, sim, admin, beaconComm, bridgeComm)
	if err != nil {
		return nil, fmt.Errorf("failed to deploy IncognitoProxy contract: %v", err)
	}
	sim.Commit()
	// fmt.Printf("deployed bridge, addr: %x ", p.incAddr)
	// printReceipt(sim, tx)

	// Vault
	prevVault := common.Address{}
	p.vAddr, tx, p.v, err = vault.DeployVault(auth, sim, admin, p.incAddr, prevVault)
	if err != nil {
		return nil, fmt.Errorf("failed to deploy Vault contract: %v", err)
	}
	sim.Commit()
	// fmt.Printf("deployed vault, addr: %x ", p.vAddr)
	// printReceipt(sim, tx)

	return p, nil
}

func setupWithLocalCommittee() (*Platform, error) {
	url := "http://127.0.0.1:9334"
	beaconOld, bridgeOld, err := getCommittee(url)
	if err != nil {
		return nil, err
	}
	return setup(beaconOld, bridgeOld, []int{})
}

func setupWithHardcodedCommittee() (*Platform, error) {
	beaconOld, bridgeOld := getCommitteeHardcoded()
	return setup(beaconOld, bridgeOld, []int{})
}

type account struct {
	PrivateKey *ecdsa.PrivateKey
	Address    common.Address
}

func loadAccount() *account {
	key, _ := crypto.LoadECDSA("genesisKey.hex")
	return &account{
		PrivateKey: key,
		Address:    crypto.PubkeyToAddress(key.PublicKey),
	}
}

func newAccount() *account {
	key, _ := crypto.GenerateKey()
	// crypto.SaveECDSA("genesisKey.hex", key)
	return &account{
		PrivateKey: key,
		Address:    crypto.PubkeyToAddress(key.PublicKey),
	}
}

func printReceipt(sim *backends.SimulatedBackend, tx *types.Transaction) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	receipt, err := sim.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		fmt.Println("receipt err:", err)
	}
	fmt.Printf("tx gas used: %v\n", receipt.CumulativeGasUsed)

	if len(receipt.Logs) == 0 {
		fmt.Println("empty log")
		return
	}

	for i, log := range receipt.Logs {
		var data interface{}
		data = log.Data

		format := "%+v"
		switch log.Topics[0].Hex() {
		case "0x8b1126c8e4087477c3efd9e3785935b29c778491c70e249de774345f7ca9b7f9", "0xa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f": // bytes32
			format = "%s"
		case "0xb42152598f9b870207037767fd41b627a327c9434c796b2ee501d68acec68d1b", "0x009fd52f05c0ded31d6fb0ee580b923f85e99cf1a5a1da342f25e73c45829c83":
			format = "%x"
		case "0x6c8f06ff564112a969115be5f33d4a0f87ba918c9c9bc3090fe631968e818be4": // bool
			format = "%t"
			data = log.Data[len(log.Data)-1] > 0
		case "0x8e2fc7b10a4f77a18c553db9a8f8c24d9e379da2557cb61ad4cc513a2f992cbd", "0x0ac68d08c5119b8cdb4058edbf0d4168f208ec3935d26a8f1f0d92eb9d4de8bf": // uint
			format = "%s"
			data = big.NewInt(int64(0)).SetBytes(log.Data)
		case "0x0ac6e167e94338a282ec23bdd86f338fc787bd67f48b3ade098144aac3fcd86e", "0xb123f68b8ba02b447d91a6629e121111b7dd6061ff418a60139c8bf00522a284": // address
			format = "%x"
			data = log.Data[12:]
		}

		fmt.Printf(fmt.Sprintf("logs[%%d]: %s\n", format), i, data)
		// for _, topic := range log.Topics {
		// 	fmt.Printf("topic: %x\n", topic)
		// }
	}
}

func getAndDecodeBridgeSwapProof(url string, block int) (*decodedProof, error) {
	body := getBridgeSwapProof(url, block)
	if len(body) < 1 {
		return nil, fmt.Errorf("no bridge swap proof found")
	}
	r := getProofResult{}
	if err := json.Unmarshal([]byte(body), &r); err != nil {
		return nil, err
	}
	if len(r.Result.Instruction) == 0 {
		return nil, fmt.Errorf("invalid swap proof")
	}
	proof, err := decodeProof(&r)
	if err != nil {
		return nil, err
	}
	return proof, nil
}

func getBridgeSwapProof(url string, block int) string {
	payload := strings.NewReader(fmt.Sprintf("{\n    \"id\": 1,\n    \"jsonrpc\": \"1.0\",\n    \"method\": \"getbridgeswapproof\",\n    \"params\": [\n    \t%d\n    ]\n}", block))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "127.0.0.1:9338")
	req.Header.Add("accept-encoding", "gzip, deflate")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("err:", err)
		return ""
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(string(body))
	return string(body)
}

func getBeaconSwapProof(block int) string {
	url := "http://127.0.0.1:9344"

	payload := strings.NewReader(fmt.Sprintf("{\n    \"id\": 1,\n    \"jsonrpc\": \"1.0\",\n    \"method\": \"getbeaconswapproof\",\n    \"params\": [\n    \t%d\n    ]\n}", block))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "127.0.0.1:9338")
	req.Header.Add("accept-encoding", "gzip, deflate")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("err:", err)
		return ""
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(string(body))
	return string(body)
}

func deposit(p *Platform, amount *big.Int) (*big.Int, *big.Int, error) {
	initBalance := p.getBalance(p.vAddr)
	auth := bind.NewKeyedTransactor(genesisAcc.PrivateKey)
	auth.GasLimit = 0
	auth.Value = amount
	_, err := p.v.Deposit(auth, "")
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	p.sim.Commit()
	newBalance := p.getBalance(p.vAddr)
	return initBalance, newBalance, nil
}
