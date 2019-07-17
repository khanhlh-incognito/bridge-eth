package bridge

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
	"github.com/incognitochain/bridge-eth/checkMulSig"
	"github.com/incognitochain/bridge-eth/common/base58"
	"github.com/incognitochain/bridge-eth/incognito_proxy"
	"github.com/incognitochain/bridge-eth/vault"
)

var auth *bind.TransactOpts
var genesisAcc *account

type Platform struct {
	*contracts
	sim *backends.SimulatedBackend
}

func init() {
	fmt.Println("Initializing genesis account...")
	genesisAcc = newAccount()
	auth = bind.NewKeyedTransactor(genesisAcc.PrivateKey)
}

func TestSimulatedSwapBeacon(t *testing.T) {
	// body := getBeaconSwapProof()
	body := getBridgeSwapProof()
	if len(body) < 1 {
		t.Fatal(fmt.Errorf("empty beacon swap proof"))
	}

	r := getProofResult{}
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		t.Fatal(err)
	}
	proof, err := decodeProof(&r)
	if err != nil {
		t.Fatal(err)
	}
	_ = proof

	p, err := setupWithCommittee()
	if err != nil {
		t.Fatal(err)
	}
	_ = p

	auth.GasLimit = 7000000
	fmt.Printf("inst len: %d\n", len(proof.instruction))
	numPk := big.NewInt(int64((len(proof.instruction) - 35) / pubkey_size))
	fmt.Printf("numPk: %d\n", numPk)
	tx, err := p.inc.SwapCommittee(
		auth,
		proof.instruction,
		numPk,

		proof.beaconInstPath,
		proof.beaconInstPathIsLeft,
		proof.beaconInstPathLen,
		proof.beaconInstRoot,
		proof.beaconBlkData,
		proof.beaconBlkHash,
		proof.beaconSignerSig,
		proof.beaconNumR,
		proof.beaconXs,
		proof.beaconYs,
		proof.beaconRIdxs,
		proof.beaconNumSig,
		proof.beaconSigIdxs,
		proof.beaconRx,
		proof.beaconRy,
		proof.beaconR,

		proof.bridgeInstPath,
		proof.bridgeInstPathIsLeft,
		proof.bridgeInstPathLen,
		proof.bridgeInstRoot,
		proof.bridgeBlkData,
		proof.bridgeBlkHash,
		proof.bridgeSignerSig,
		proof.bridgeNumR,
		proof.bridgeXs,
		proof.bridgeYs,
		proof.bridgeRIdxs,
		proof.bridgeNumSig,
		proof.bridgeSigIdxs,
		proof.bridgeRx,
		proof.bridgeRy,
		proof.bridgeR,
	)
	if err != nil {
		fmt.Println("err:", err)
	}
	p.sim.Commit()
	printReceipt(p.sim, tx)
}

func TestSimulatedBurn(t *testing.T) {
	proof, err := getAndDecodeBurnProof("")
	if err != nil {
		t.Fatal(err)
	}

	p, err := setupWithCommittee()
	if err != nil {
		t.Fatalf("Fail to deloy contract: %v\n", err)
	}

	oldBalance, newBalance, err := deposit(p, int64(5e18))
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("deposit to vault: %d -> %d\n", oldBalance, newBalance)

	withdrawer := common.HexToAddress("0x0FFBd68F130809BcA7b32D9536c8339E9A844620")
	fmt.Printf("withdrawer init balance: %d\n", p.getBalance(withdrawer))

	tx, err := withdraw(p.v, auth, proof)
	if err != nil {
		fmt.Println("err:", err)
	}
	p.sim.Commit()
	printReceipt(p.sim, tx)

	fmt.Printf("withdrawer new balance: %d\n", p.getBalance(withdrawer))
}

func (p *Platform) getBalance(addr common.Address) *big.Int {
	bal, _ := p.sim.BalanceAt(context.Background(), addr, nil)
	return bal
}

func setup(beaconComm, bridgeComm []byte) (*Platform, error) {
	alloc := make(core.GenesisAlloc)
	balance, _ := big.NewInt(1).SetString("100000000000000000000", 10) // 100 eth
	alloc[auth.From] = core.GenesisAccount{Balance: balance}
	sim := backends.NewSimulatedBackend(alloc, 8000000)
	p := &Platform{sim: sim, contracts: &contracts{}}

	// MulSigP256
	var err error
	var tx *types.Transaction
	_ = tx
	p.msAddr, tx, p.ms, err = checkMulSig.DeployMulSigP256(auth, sim)
	if err != nil {
		return nil, fmt.Errorf("failed to deploy mulsig contract: %v", err)
	}
	sim.Commit()
	// printReceipt(sim, tx)

	// IncognitoProxy
	p.incAddr, tx, p.inc, err = incognito_proxy.DeployIncognitoProxy(auth, sim, beaconComm, bridgeComm, p.msAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to deploy IncognitoProxy contract: %v", err)
	}
	sim.Commit()
	fmt.Printf("deployed bridge, addr: %x\n", p.incAddr)
	// printReceipt(sim, tx)

	// Vault
	p.vAddr, tx, p.v, err = vault.DeployVault(auth, sim, p.incAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to deploy Vault contract: %v", err)
	}
	sim.Commit()
	fmt.Printf("deployed vault, addr: %x\n", p.vAddr)
	// printReceipt(sim, tx)

	return p, nil
}

func setupWithCommittee() (*Platform, error) {
	url := "http://127.0.0.1:9334"

	payload := strings.NewReader("{\n    \"id\": 1,\n    \"jsonrpc\": \"1.0\",\n    \"method\": \"getbeaconbeststate\",\n    \"params\": []\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "127.0.0.1:9334")
	req.Header.Add("accept-encoding", "gzip, deflate")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	type beaconBestStateResult struct {
		BeaconCommittee []string
		ShardCommittee  map[string][]string
	}

	type getBeaconBestStateResult struct {
		Result beaconBestStateResult
		Error  string
		Id     int
	}

	r := getBeaconBestStateResult{}
	err = json.Unmarshal([]byte(body), &r)
	if err != nil {
		return nil, err
	}

	// Genesis committee
	beaconOld := []byte{}
	for i, val := range r.Result.BeaconCommittee {
		pk, _, _ := base58.Base58Check{}.Decode(val)
		fmt.Printf("pk[%d]: %x %d\n", i, pk, len(pk))
		beaconOld = append(beaconOld, pk...)
	}

	bridgeOld := []byte{}
	for i, val := range r.Result.ShardCommittee["1"] {
		pk, _, _ := base58.Base58Check{}.Decode(val)
		fmt.Printf("pk[%d]: %x %d\n", i, pk, len(pk))
		bridgeOld = append(bridgeOld, pk...)
	}

	return setup(beaconOld, bridgeOld)
}

type account struct {
	PrivateKey *ecdsa.PrivateKey
	Address    common.Address
}

func newAccount() *account {
	// key, _ := crypto.GenerateKey()
	// crypto.SaveECDSA("genesisKey.hex", key)
	key, _ := crypto.LoadECDSA("genesisKey.hex")
	return &account{
		PrivateKey: key,
		Address:    crypto.PubkeyToAddress(key.PublicKey),
	}
}

func printReceipt(sim *backends.SimulatedBackend, tx *types.Transaction) {
	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
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
		case "0x8b1126c8e4087477c3efd9e3785935b29c778491c70e249de774345f7ca9b7f9":
			format = "%s"
		case "0xb42152598f9b870207037767fd41b627a327c9434c796b2ee501d68acec68d1b":
			format = "%x"
		case "0x6c8f06ff564112a969115be5f33d4a0f87ba918c9c9bc3090fe631968e818be4":
			format = "%t"
			data = log.Data[len(log.Data)-1] > 0
		case "0x8e2fc7b10a4f77a18c553db9a8f8c24d9e379da2557cb61ad4cc513a2f992cbd":
			format = "%s"
			data = big.NewInt(int64(0)).SetBytes(log.Data)
		case "0x0ac6e167e94338a282ec23bdd86f338fc787bd67f48b3ade098144aac3fcd86e":
			format = "%x"
			data = log.Data[12:]
		}

		fmt.Printf(fmt.Sprintf("logs[%%d]: %s\n", format), i, data)
		// for _, topic := range log.Topics {
		// 	fmt.Printf("topic: %x\n", topic)
		// }
	}
}

func getBridgeSwapProof() string {
	url := "http://127.0.0.1:9338"

	block := 14
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

func getBeaconSwapProof() string {
	url := "http://127.0.0.1:9338"

	block := 16
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

func deposit(p *Platform, amount int64) (*big.Int, *big.Int, error) {
	initBalance := p.getBalance(p.vAddr)
	auth := bind.NewKeyedTransactor(genesisAcc.PrivateKey)
	auth.Value = big.NewInt(amount)
	_, err := p.v.Deposit(auth, "")
	if err != nil {
		return nil, nil, err
	}
	p.sim.Commit()
	newBalance := p.getBalance(p.vAddr)
	return initBalance, newBalance, nil
}
