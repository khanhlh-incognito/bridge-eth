package main

// Basic imports
import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"math/rand"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/incognitochain/bridge-eth/bridge/zrxtrade"
	"github.com/incognitochain/bridge-eth/common/base58"
	"github.com/incognitochain/bridge-eth/consensus/signatureschemes/bridgesig"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/sha3"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/incognitochain/bridge-eth/erc20"
	"github.com/incognitochain/bridge-eth/rpccaller"
	"github.com/stretchr/testify/require"
)

type IssuingETHRes struct {
	rpccaller.RPCBaseRes
	Result interface{} `json:"Result"`
}

type BurningForDepositToSCRes struct {
	rpccaller.RPCBaseRes
	Result interface{} `json:"Result"`
}

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type TradingTestSuite struct {
	suite.Suite
	IncBurningAddrStr string
	IncPrivKeyStr     string
	IncPaymentAddrStr string

	GeneratedPrivKeyForSC ecdsa.PrivateKey
	GeneratedPubKeyForSC  ecdsa.PublicKey

	IncEtherTokenIDStr string
	IncDAITokenIDStr   string
	IncSAITokenIDStr   string

	IncBridgeHost string
	IncRPCHost    string

	EtherAddressStr string
	DAIAddressStr   string
	SAIAddressStr   string

	ETHPrivKeyStr   string
	ETHOwnerAddrStr string

	ETHHost    string
	ETHPrivKey *ecdsa.PrivateKey
	ETHClient  *ethclient.Client

	VaultAddr            common.Address
	KBNTradeDeployedAddr common.Address
	ZRXTradeDeployedAddr common.Address

	KyberContractAddr common.Address
	ZRXContractAddr   common.Address
	WETHAddr          common.Address

	Quote0xUrl string

	// token amounts for tests
	DepositingEther float64
	DAIBalanceAfterStep1 *big.Int
	SAIBalanceAfterStep2 *big.Int
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (tradingSuite *TradingTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")

	// 0x kovan env
	tradingSuite.IncBurningAddrStr = "15pABFiJVeh9D5uiQEhQX4SVibGGbdAVipQxBdxkmDqAJaoG1EdFKHBrNfs"
	tradingSuite.IncPrivKeyStr = "112t8roafGgHL1rhAP9632Yef3sx5k8xgp8cwK4MCJsCL1UWcxXvpzg97N4dwvcD735iKf31Q2ZgrAvKfVjeSUEvnzKJyyJD3GqqSZdxN4or"
	tradingSuite.IncPaymentAddrStr = "12S5Lrs1XeQLbqN4ySyKtjAjd2d7sBP2tjFijzmp6avrrkQCNFMpkXm3FPzj2Wcu2ZNqJEmh9JriVuRErVwhuQnLmWSaggobEWsBEci"
	// tradingSuite.GeneratedPubKeyForSCStr = "8224890Cd5A537792d1B8B56c95FAb8a1A5E98B1"

	tradingSuite.IncEtherTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000099"
	tradingSuite.IncDAITokenIDStr = "0000000000000000000000000000000000000000000000000000000000000098"
	tradingSuite.IncSAITokenIDStr = "0000000000000000000000000000000000000000000000000000000000000097"

	tradingSuite.EtherAddressStr = "0x0000000000000000000000000000000000000000"
	tradingSuite.DAIAddressStr = "0x4f96fe3b7a6cf9725f59d353f723c1bdb64ca6aa"
	tradingSuite.SAIAddressStr = "0xc4375b7de8af5a38a93548eb8453a498222c4ff2"

	tradingSuite.ETHPrivKeyStr = "B8DB29A7A43FB88AD520F762C5FDF6F1B0155637FA1E5CB2C796AFE9E5C04E31"
	tradingSuite.ETHOwnerAddrStr = "FD94c46ab8dCF0928d5113a6fEaa925793504e16"

	tradingSuite.ETHHost = "https://kovan.infura.io/v3/93fe721349134964aa71071a713c5cef"
	tradingSuite.IncBridgeHost = "http://127.0.0.1:9338"
	tradingSuite.IncRPCHost = "http://127.0.0.1:9334"
	tradingSuite.KyberContractAddr = common.HexToAddress("0xdd974d5c2e2928dea5f71b9825b8b646686bd200")
	tradingSuite.ZRXContractAddr = common.HexToAddress("0xf1ec01d6236d3cd881a0bf0130ea25fe4234003e")
	tradingSuite.WETHAddr = common.HexToAddress("0xd0a1e359811322d97991e03f863a0c30c2cf029c")

	tradingSuite.VaultAddr = common.HexToAddress("0x36E59EA6f27C7d33432adD1C017Ee70Def5eE7d9")
	tradingSuite.KBNTradeDeployedAddr = common.HexToAddress("0x6dE453d6f7EB17Eb51Ee1A20544a1cD9200a1a55")
	tradingSuite.ZRXTradeDeployedAddr = common.HexToAddress("0x0F248fA4c3F93e495bff1FfD4Eed4BA1D9911654")

	tradingSuite.Quote0xUrl = "https://kovan.api.0x.org/swap/v0/quote?sellToken=%v&buyToken=%v&sellAmount=%v"

	tradingSuite.DepositingEther = float64(0.05)

	// generate a new keys pair for SC
	genKeysPairForSC(tradingSuite)

	// connect to ethereum network
	connectToETH(tradingSuite)
}

func (tradingSuite *TradingTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	tradingSuite.ETHClient.Close()
}

func (tradingSuite *TradingTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (tradingSuite *TradingTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestTradingTestSuite(t *testing.T) {
	fmt.Println("Starting entry point...")
	tradingSuite := new(TradingTestSuite)
	suite.Run(t, tradingSuite)
	fmt.Println("Finishing entry point...")
}

func getBalanceOnETHNet(
	tradingSuite *TradingTestSuite,
	tokenAddr common.Address,
	ownerAddr common.Address,
) *big.Int {
	if tokenAddr.Hex() == tradingSuite.EtherAddressStr {
		balance, err := tradingSuite.ETHClient.BalanceAt(context.Background(), ownerAddr, nil)
		require.Equal(tradingSuite.T(), nil, err)
		return balance
	}
	// erc20 token
	instance, err := erc20.NewErc20(tokenAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)

	balance, err := instance.BalanceOf(&bind.CallOpts{}, ownerAddr)
	require.Equal(tradingSuite.T(), nil, err)
	return balance
}

func connectToETH(tradingSuite *TradingTestSuite) {
	privKeyHex := tradingSuite.ETHPrivKeyStr
	privKey, err := crypto.HexToECDSA(privKeyHex)
	require.Equal(tradingSuite.T(), nil, err)

	fmt.Printf("Sign Txs with address: %s\n", crypto.PubkeyToAddress(privKey.PublicKey).Hex())

	network := "development"
	fmt.Printf("Connecting to network %s\n", network)
	client, err := ethclient.Dial(tradingSuite.ETHHost)
	require.Equal(tradingSuite.T(), nil, err)

	tradingSuite.ETHClient = client
	tradingSuite.ETHPrivKey = privKey
}

func depositETH(
	tradingSuite *TradingTestSuite,
	amt float64,
	incPaymentAddrStr string,
) common.Hash {
	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)

	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
	auth.Value = big.NewInt(int64(amt * params.Ether))
	tx, err := c.Deposit(auth, incPaymentAddrStr)
	require.Equal(tradingSuite.T(), nil, err)
	txHash := tx.Hash()

	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("deposited, txHash: %x\n", txHash[:])
	return txHash
}

func depositERC20ToBridge(
	tradingSuite *TradingTestSuite,
	amt *big.Int,
	tokenAddr common.Address,
	incPaymentAddrStr string,
) common.Hash {
	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)

	erc20Token, _ := erc20.NewErc20(tokenAddr, tradingSuite.ETHClient)
	auth.GasPrice = big.NewInt(50000000000)
	auth.GasLimit = 1000000
	tx2, apprErr := erc20Token.Approve(auth, tradingSuite.VaultAddr, amt)
	tx2Hash := tx2.Hash()
	fmt.Printf("Approve tx, txHash: %x\n", tx2Hash[:])
	require.Equal(tradingSuite.T(), nil, apprErr)
	time.Sleep(15 * time.Second)
	auth.GasPrice = big.NewInt(50000000000)
	auth.GasLimit = 1000000

	fmt.Println("Starting deposit erc20 to vault contract")
	tx, err := c.DepositERC20(auth, tokenAddr, amt, incPaymentAddrStr)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("Finished deposit erc20 to vault contract")
	txHash := tx.Hash()

	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("deposited erc20 token to bridge, txHash: %x\n", txHash[:])
	return txHash
}

func callIssuingETHReq(
	tradingSuite *TradingTestSuite,
	incTokenIDStr string,
	ethDepositProof []string,
	ethBlockHash string,
	ethTxIdx uint,
) (map[string]interface{}, error) {
	rpcClient := rpccaller.NewRPCClient()
	meta := map[string]interface{}{
		"IncTokenID": incTokenIDStr,
		"BlockHash":  ethBlockHash,
		"ProofStrs":  ethDepositProof,
		"TxIndex":    ethTxIdx,
	}
	params := []interface{}{
		tradingSuite.IncPrivKeyStr,
		nil,
		5,
		-1,
		meta,
	}
	var res IssuingETHRes
	err := rpcClient.RPCCall(
		"",
		tradingSuite.IncRPCHost,
		"",
		"createandsendtxwithissuingethreq",
		params,
		&res,
	)
	if err != nil {
		return nil, err
	}

	response, _ := json.Marshal(res)
	fmt.Println("get response", string(response))

	return res.Result.(map[string]interface{}), nil
}

func callBurningPToken(
	tradingSuite *TradingTestSuite,
	incTokenIDStr string,
	amount *big.Int,
	remoteAddrStr string,
	burningMethod string,
) (map[string]interface{}, error) {
	rpcClient := rpccaller.NewRPCClient()
	meta := map[string]interface{}{
		"TokenID":     incTokenIDStr,
		"TokenTxType": 1,
		"TokenName":   "",
		"TokenSymbol": "",
		"TokenAmount": amount.Uint64(),
		"TokenReceivers": map[string]uint64{
			tradingSuite.IncBurningAddrStr: amount.Uint64(),
		},
		"RemoteAddress": remoteAddrStr,
		"Privacy":       true,
		"TokenFee":      0,
	}
	params := []interface{}{
		tradingSuite.IncPrivKeyStr,
		nil,
		5,
		-1,
		meta,
		"",
		0,
	}
	var res BurningForDepositToSCRes
	err := rpcClient.RPCCall(
		"",
		tradingSuite.IncRPCHost,
		"",
		burningMethod,
		params,
		&res,
	)
	if err != nil {
		fmt.Println("calling burning ptokens err: ", err)
		return nil, err
	}
	bb, _ := json.Marshal(res)
	fmt.Println("calling burning ptokens res: ", string(bb))
	if res.RPCError != nil {
		return nil, errors.New(res.RPCError.Message)
	}
	return res.Result.(map[string]interface{}), nil
}

func submitBurnProofForDepositToSC(
	tradingSuite *TradingTestSuite,
	burningTxIDStr string,
) {
	proof, err := getAndDecodeBurnProofV2(tradingSuite.IncBridgeHost, burningTxIDStr, "getburnprooffordeposittosc")
	require.Equal(tradingSuite.T(), nil, err)

	// Get contract instance
	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)

	// Burn
	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
	// auth.GasPrice = big.NewInt(1000000)
	// auth.GasLimit = 4000000
	tx, err := SubmitBurnProof(c, auth, proof)
	require.Equal(tradingSuite.T(), nil, err)

	txHash := tx.Hash()
	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("burned, txHash: %x\n", txHash[:])
}

func submitBurnProofForWithdrawal(
	tradingSuite *TradingTestSuite,
	burningTxIDStr string,
) {
	proof, err := getAndDecodeBurnProofV2(tradingSuite.IncBridgeHost, burningTxIDStr, "getburnproof")
	require.Equal(tradingSuite.T(), nil, err)

	// Get contract instance
	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)

	// Burn
	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
	// auth.GasPrice = big.NewInt(1000000)
	// auth.GasLimit = 4000000
	tx, err := Withdraw(c, auth, proof)
	require.Equal(tradingSuite.T(), nil, err)

	txHash := tx.Hash()
	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("burned, txHash: %x\n", txHash[:])
}

func genKeysPairForSC(tradingSuite *TradingTestSuite) {
	incPriKeyBytes, _, err := base58.Base58Check{}.Decode(tradingSuite.IncPrivKeyStr)
	require.Equal(tradingSuite.T(), nil, err)

	tradingSuite.GeneratedPrivKeyForSC, tradingSuite.GeneratedPubKeyForSC = bridgesig.KeyGen(incPriKeyBytes)
}

func randomizeTimestamp() string {
	randomTime := rand.Int63n(time.Now().Unix()-94608000) + 94608000
	randomNow := time.Unix(randomTime, 0)
	return randomNow.String()
}

func rawsha3(b []byte) []byte {
	hashF := sha3.NewLegacyKeccak256()
	hashF.Write(b)
	buf := hashF.Sum(nil)
	return buf
}

func rlpHash(x interface{}) (h common.Hash) {
	hw := sha3.NewLegacyKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}

func executeWith0x(
	tradingSuite *TradingTestSuite,
	srcQty *big.Int,
	srcTokenName string,
	srcTokenIDStr string,
	destTokenName string,
	destTokenIDStr string,
) {
	tradeAbi, _ := abi.JSON(strings.NewReader(zrxtrade.ZrxtradeABI))

	// Get contract instance
	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)
	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
	auth.GasPrice = big.NewInt(50000000000)
	auth.GasLimit = 2000000
	// quote
	srcToken := common.HexToAddress(srcTokenIDStr)
	destToken := common.HexToAddress(destTokenIDStr)

	quoteData, _ := quote0x(tradingSuite.Quote0xUrl, srcTokenName, destTokenName, srcQty)
	forwarder := common.HexToAddress(quoteData["to"].(string))
	dt := common.Hex2Bytes(quoteData["data"].(string)[2:])
	auth.Value, _ = big.NewInt(0).SetString(quoteData["protocolFee"].(string), 10)
	auth.GasPrice, _ = big.NewInt(0).SetString(quoteData["gasPrice"].(string), 10)
	input, _ := tradeAbi.Pack("trade", srcToken, srcQty, destToken, dt, forwarder)
	timestamp := []byte(randomizeTimestamp())
	tempData := append(tradingSuite.ZRXTradeDeployedAddr[:], input...)
	tempData1 := append(tempData, timestamp...)
	data := rawsha3(tempData1)
	signBytes, _ := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)

	tx, err := c.Execute(
		auth,
		srcToken,
		srcQty,
		destToken,
		tradingSuite.ZRXTradeDeployedAddr,
		input,
		timestamp,
		signBytes,
	)
	require.Equal(tradingSuite.T(), nil, err)
	txHash := tx.Hash()
	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("0x trade executed , txHash: %x\n", txHash[:])
}

func quote0x(
	quote0xUrl string,
	srcToken, destToken string,
	srcQty *big.Int,
) (map[string]interface{}, error) {
	var (
		err       error
		resp      *http.Response
		bodyBytes []byte
		result    interface{}
	)
	url := fmt.Sprintf(quote0xUrl, srcToken, destToken, srcQty.String())
	if resp, err = http.Get(url); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Request returns with fucking error!!!")
	}
	if bodyBytes, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(bodyBytes, &result); err != nil {
		return nil, err
	}
	return result.(map[string]interface{}), nil
}

func getDepositedBalance(
	tradingSuite *TradingTestSuite,
	ethTokenAddrStr string,
	ownerAddrStr string,
) *big.Int {
	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)
	token := common.HexToAddress(ethTokenAddrStr)
	owner := common.HexToAddress(ownerAddrStr)
	bal, err := c.GetDepositedBalance(nil, token, owner)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Printf("deposited balance: %d\n", bal)
	return bal
}

func requestWithdraw(
	tradingSuite *TradingTestSuite,
	withdrawalETHTokenIDStr string,
	amount *big.Int,
) common.Hash {
	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)
	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)

	token := common.HexToAddress(withdrawalETHTokenIDStr)
	// amount := big.NewInt(0.1 * params.Ether)
	timestamp := []byte(randomizeTimestamp())
	tempData := append([]byte(tradingSuite.IncPaymentAddrStr), token[:]...)
	tempData1 := append(tempData, timestamp...)
	data := rawsha3(tempData1)
	signBytes, _ := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)
	auth.GasPrice = big.NewInt(50000000000)
	auth.GasLimit = 2000000
	tx, err := c.RequestWithdraw(auth, tradingSuite.IncPaymentAddrStr, token, amount, signBytes, timestamp)
	require.Equal(tradingSuite.T(), nil, err)

	txHash := tx.Hash()
	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("request withdrawal, txHash: %x\n", txHash[:])
	return txHash
}

func (tradingSuite *TradingTestSuite) Test1TradeEthForDaiWith0x() {
	fmt.Println("============ TEST TRADE ETHER FOR DAI WITH 0X AGGREGATOR ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()

	fmt.Println("------------ STEP 1: porting ETH to pETH --------------")
	txHash := depositETH(tradingSuite, tradingSuite.DepositingEther, tradingSuite.IncPaymentAddrStr)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(90 * time.Second)
	_, err = callIssuingETHReq(
		tradingSuite,
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	fmt.Println("------------ STEP 2: burning pETH to deposit ETH to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := callBurningPToken(
		tradingSuite,
		tradingSuite.IncEtherTokenIDStr,
		burningPETH,
		pubKeyToAddrStr[2:],
		"createandsendburningfordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	submitBurnProofForDepositToSC(tradingSuite, burningTxID.(string))
	deposited := getDepositedBalance(
		tradingSuite,
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("deposited EHT: ", deposited)
	//require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPETH, big.NewInt(1000000000)), deposited)

	fmt.Println("------------ step 3: execute trade ETH for DAI through 0x aggregator --------------")
	executeWith0x(
		tradingSuite,
		tradeAmount,
		"ETH",
		tradingSuite.EtherAddressStr,
		"DAI",
		tradingSuite.DAIAddressStr,
	)
	daiTraded := getDepositedBalance(
		tradingSuite,
		tradingSuite.DAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("daiTraded: ", daiTraded)

	fmt.Println("------------ step 4: withdrawing DAI from SC to pDAI on Incognito --------------")
	txHashByEmittingWithdrawalReq := requestWithdraw(tradingSuite, tradingSuite.DAIAddressStr, daiTraded)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(90 * time.Second)
	_, err = callIssuingETHReq(
		tradingSuite,
		tradingSuite.IncDAITokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	fmt.Println("------------ step 5: withdrawing pDAI from Incognito to DAI --------------")
	withdrawingPDAI := big.NewInt(0).Div(daiTraded, big.NewInt(1000000000))
	burningRes, err = callBurningPToken(
		tradingSuite,
		tradingSuite.IncDAITokenIDStr,
		withdrawingPDAI,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	submitBurnProofForWithdrawal(tradingSuite, burningTxID.(string))

	bal := getBalanceOnETHNet(
		tradingSuite,
		common.HexToAddress(tradingSuite.DAIAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	tradingSuite.DAIBalanceAfterStep1 = bal
	fmt.Println("DAI balance after step 1: ", tradingSuite.DAIBalanceAfterStep1)
	// require.Equal(tradingSuite.T(), withdrawingPDAI.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())
}

func (tradingSuite *TradingTestSuite) Test2TradeDaiForSaiWith0x() {
	fmt.Println("============ TEST TRADE DAI FOR SAI WITH 0X AGGREGATOR ===========")
	fmt.Println("------------ step 0: declaration & initialization --------------")
	depositingDAI := tradingSuite.DAIBalanceAfterStep1
	burningPDAI := big.NewInt(0).Div(depositingDAI, big.NewInt(1000000000))
	tradeAmount := depositingDAI

	daibal := getBalanceOnETHNet(
		tradingSuite,
		common.HexToAddress(tradingSuite.DAIAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("dai balance of owner: ", daibal)

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("pubKeyToAddrStr: ", pubKeyToAddrStr)

	fmt.Println("------------ step 1: porting DAI to pDAI --------------")
	txHash := depositERC20ToBridge(
		tradingSuite,
		depositingDAI,
		common.HexToAddress(tradingSuite.DAIAddressStr),
		tradingSuite.IncPaymentAddrStr,
	)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(90 * time.Second)

	_, err = callIssuingETHReq(
		tradingSuite,
		tradingSuite.IncDAITokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)
	// TODO: check the new balance on peth of the incognito account

	fmt.Println("------------ step 2: burning pDAI to deposit DAI to SC --------------")

	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := callBurningPToken(
		tradingSuite,
		tradingSuite.IncDAITokenIDStr,
		burningPDAI,
		pubKeyToAddrStr[2:],
		"createandsendburningfordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	submitBurnProofForDepositToSC(tradingSuite, burningTxID.(string))
	deposited := getDepositedBalance(
		tradingSuite,
		tradingSuite.DAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("deposited dai: ", deposited)
	// require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPDAI, big.NewInt(1000000000)), deposited)

	fmt.Println("------------ step 3: execute trade DAI for SAI through 0x aggregator --------------")
	executeWith0x(
		tradingSuite,
		tradeAmount,
		"DAI",
		tradingSuite.DAIAddressStr,
		"SAI",
		tradingSuite.SAIAddressStr,
	)
	saiTraded := getDepositedBalance(
		tradingSuite,
		tradingSuite.SAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("saiTraded: ", saiTraded)

	fmt.Println("------------ step 4: withdrawing SAI from SC to pSAI on Incognito --------------")
	txHashByEmittingWithdrawalReq := requestWithdraw(tradingSuite, tradingSuite.SAIAddressStr, saiTraded)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(90 * time.Second)

	_, err = callIssuingETHReq(
		tradingSuite,
		tradingSuite.IncSAITokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	fmt.Println("------------ step 5: withdrawing pSAI from Incognito to SAI --------------")
	withdrawingPSAI := big.NewInt(0).Div(saiTraded, big.NewInt(1000000000))
	burningRes, err = callBurningPToken(
		tradingSuite,
		tradingSuite.IncSAITokenIDStr,
		withdrawingPSAI,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	submitBurnProofForWithdrawal(tradingSuite, burningTxID.(string))

	bal := getBalanceOnETHNet(
		tradingSuite,
		common.HexToAddress(tradingSuite.SAIAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	tradingSuite.SAIBalanceAfterStep2 = bal
	fmt.Println("SAI balance after step 2: ", tradingSuite.SAIBalanceAfterStep2)
	// require.Equal(tradingSuite.T(), withdrawingPSAI.Uint64(), bal.Uint64())
	// require.Equal(tradingSuite.T(), withdrawingPSAI.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())
}

func (tradingSuite *TradingTestSuite) Test3TradeSaiForEthWith0x() {
	fmt.Println("============ TEST TRADE SAI FOR ETH WITH 0X AGGREGATOR ===========")
	fmt.Println("------------ step 0: declaration & initialization --------------")
	depositingSAI := tradingSuite.SAIBalanceAfterStep2
	burningPSAI := big.NewInt(0).Div(depositingSAI, big.NewInt(1000000000))
	tradeAmount := depositingSAI

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("pubKeyToAddrStr: ", pubKeyToAddrStr)

	fmt.Println("------------ step 1: porting SAI to pSAI --------------")
	txHash := depositERC20ToBridge(
		tradingSuite,
		depositingSAI,
		common.HexToAddress(tradingSuite.SAIAddressStr),
		tradingSuite.IncPaymentAddrStr,
	)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(90 * time.Second)

	issuingRes, err := callIssuingETHReq(
		tradingSuite,
		tradingSuite.IncSAITokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("issuingRes: ", issuingRes)
	time.Sleep(120 * time.Second)

	fmt.Println("------------ step 2: burning pSAI to deposit SAI to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := callBurningPToken(
		tradingSuite,
		tradingSuite.IncSAITokenIDStr,
		burningPSAI,
		pubKeyToAddrStr[2:],
		"createandsendburningfordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(140 * time.Second)

	submitBurnProofForDepositToSC(tradingSuite, burningTxID.(string))
	deposited := getDepositedBalance(
		tradingSuite,
		tradingSuite.SAIAddressStr,
		pubKeyToAddrStr,
	)
	require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPSAI, big.NewInt(1000000000)), deposited)

	fmt.Println("------------ step 3: execute trade SAI for ETH through 0x aggregator --------------")
	executeWith0x(
		tradingSuite,
		tradeAmount,
		"SAI",
		tradingSuite.SAIAddressStr,
		"ETH",
		tradingSuite.EtherAddressStr,
	)
	etherTraded := getDepositedBalance(
		tradingSuite,
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("etherTraded: ", etherTraded)

	fmt.Println("------------ step 4: withdrawing ETH from SC to pETH on Incognito --------------")
	txHashByEmittingWithdrawalReq := requestWithdraw(tradingSuite, tradingSuite.EtherAddressStr, etherTraded)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(90 * time.Second)

	_, err = callIssuingETHReq(
		tradingSuite,
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(140 * time.Second)

	fmt.Println("------------ step 5: withdrawing pETH from Incognito to ETH --------------")
	withdrawingPETH := big.NewInt(0).Div(etherTraded, big.NewInt(1000000000))
	burningRes, err = callBurningPToken(
		tradingSuite,
		tradingSuite.IncEtherTokenIDStr,
		withdrawingPETH,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(140 * time.Second)

	submitBurnProofForWithdrawal(tradingSuite, burningTxID.(string))

	bal := getBalanceOnETHNet(
		tradingSuite,
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("Ether balance after step 3: ", bal)
	// require.Equal(tradingSuite.T(), withdrawingPETH.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())
}
