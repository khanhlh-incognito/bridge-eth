package main

import (
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/incognitochain/bridge-eth/bridge/kbntrade"
	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/stretchr/testify/suite"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"

	"github.com/stretchr/testify/require"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type KyberTradingTestSuite struct {
	*TradingTestSuite

	KyberTradeDeployedAddr common.Address

	KyberContractAddr    common.Address
	WETHAddr             common.Address
	EtherAddressStrKyber common.Address

	IncOMGTokenIDStr  string
	IncPOLYTokenIDStr string

	OMGAddressStr  string
	POLYAddressStr string

	// token amounts for tests
	DepositingEther       float64
	OMGBalanceAfterStep1  *big.Int
	POLYBalanceAfterStep2 *big.Int
}

func NewKyberTradingTestSuite(tradingTestSuite *TradingTestSuite) *KyberTradingTestSuite {
	return &KyberTradingTestSuite{
		TradingTestSuite: tradingTestSuite,
	}
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (tradingSuite *KyberTradingTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")
	// Kovan env
	tradingSuite.EtherAddressStrKyber = common.HexToAddress("0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee")

	tradingSuite.IncOMGTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000092"
	tradingSuite.IncPOLYTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000091"
	tradingSuite.OMGAddressStr = "0xC4375B7De8af5a38a93548eb8453a498222C4fF2"
	tradingSuite.POLYAddressStr = "0xdB7ec4E4784118D9733710e46F7C83fE7889596a"
	tradingSuite.KyberTradeDeployedAddr = common.HexToAddress("0xCbB310ef7e7bcA0f5Ec82b23AEAA3bf2157DD39a")
	tradingSuite.DepositingEther = float64(0.05)
	tradingSuite.KyberContractAddr = common.HexToAddress("0x692f391bCc85cefCe8C237C01e1f636BbD70EA4D")
}

func (tradingSuite *KyberTradingTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	tradingSuite.ETHClient.Close()
}

func (tradingSuite *KyberTradingTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (tradingSuite *KyberTradingTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestKyberTradingTestSuite(t *testing.T) {
	fmt.Println("Starting entry point for Kyber test suite...")

	tradingSuite := new(TradingTestSuite)
	suite.Run(t, tradingSuite)

	kyberTradingSuite := NewKyberTradingTestSuite(tradingSuite)
	suite.Run(t, kyberTradingSuite)

	fmt.Println("Finishing entry point for 0x test suite...")
}

func (tradingSuite *KyberTradingTestSuite) executeWithKyber(
	srcQty *big.Int,
	srcTokenIDStr string,
	destTokenIDStr string,
) {
	tradeAbi, _ := abi.JSON(strings.NewReader(kbntrade.KbntradeABI))

	// Get contract instance
	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)
	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
	auth.GasPrice = big.NewInt(50000000000)
	auth.GasLimit = 2000000
	srcToken := common.HexToAddress(srcTokenIDStr)
	destToken := common.HexToAddress(destTokenIDStr)
	auth.GasPrice = big.NewInt(50000000000)
	auth.GasLimit = 2000000
	input, _ := tradeAbi.Pack("trade", srcToken, srcQty, destToken)
	timestamp := []byte(randomizeTimestamp())
	tempData := append(tradingSuite.KyberTradeDeployedAddr[:], input...)
	tempData1 := append(tempData, timestamp...)
	data := rawsha3(tempData1)
	signBytes, _ := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)

	tx, err := c.Execute(
		auth,
		srcToken,
		srcQty,
		destToken,
		tradingSuite.KyberTradeDeployedAddr,
		input,
		timestamp,
		signBytes,
	)
	require.Equal(tradingSuite.T(), nil, err)
	txHash := tx.Hash()
	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("Kyber trade executed , txHash: %x\n", txHash[:])
}

func (tradingSuite *KyberTradingTestSuite) Test1TradeEthForOMGWithKyber() {
	fmt.Println("============ TEST TRADE ETHER FOR OMG WITH Kyber AGGREGATOR ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("------------ STEP 1: porting ETH to pETH --------------")
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
	)
	time.Sleep(15 * time.Second)
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(90 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	fmt.Println("------------ STEP 2: burning pETH to deposit ETH to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncEtherTokenIDStr,
		burningPETH,
		pubKeyToAddrStr[2:],
		"createandsendburningfordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))
	deposited := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("deposited EHT: ", deposited)
	//require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPETH, big.NewInt(1000000000)), deposited)

	fmt.Println("------------ step 3: execute trade ETH for OMG through Kyber aggregator --------------")
	tradingSuite.executeWithKyber(
		tradeAmount,
		tradingSuite.EtherAddressStr,
		tradingSuite.OMGAddressStr,
	)
	time.Sleep(15 * time.Second)
	omgTraded := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("omgTraded: ", omgTraded)

	fmt.Println("------------ step 4: withdrawing OMG from SC to pOMG on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.OMGAddressStr,
		omgTraded,
	)
	time.Sleep(15 * time.Second)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(90 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncOMGTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	fmt.Println("------------ step 5: withdrawing pOMG from Incognito to OMG --------------")
	withdrawingPOMG := big.NewInt(0).Div(omgTraded, big.NewInt(1000000000))
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncOMGTokenIDStr,
		withdrawingPOMG,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))

	bal := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.OMGAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	tradingSuite.OMGBalanceAfterStep1 = bal
	fmt.Println("OMG balance after step 1: ", tradingSuite.OMGBalanceAfterStep1)
	// require.Equal(tradingSuite.T(), withdrawingPOMG.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())
}

func (tradingSuite *KyberTradingTestSuite) Test2TradeOMGForPOLYWithKyber() {
	fmt.Println("============ TEST TRADE OMG FOR POLY WITH KYBER AGGREGATOR ===========")
	fmt.Println("------------ step 0: declaration & initialization --------------")
	depositingOMG := tradingSuite.OMGBalanceAfterStep1
	burningPOMG := big.NewInt(0).Div(depositingOMG, big.NewInt(1000000000))
	tradeAmount := depositingOMG

	omgbal := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.OMGAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("omg balance of owner: ", omgbal)

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("pubKeyToAddrStr: ", pubKeyToAddrStr)

	fmt.Println("------------ step 1: porting OMG to pOMG --------------")
	txHash := tradingSuite.depositERC20ToBridge(
		depositingOMG,
		common.HexToAddress(tradingSuite.OMGAddressStr),
		tradingSuite.IncPaymentAddrStr,
	)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(90 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncOMGTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)
	// TODO: check the new balance on peth of the incognito account

	fmt.Println("------------ step 2: burning pOMG to deposit OMG to SC --------------")

	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncOMGTokenIDStr,
		burningPOMG,
		pubKeyToAddrStr[2:],
		"createandsendburningfordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))
	deposited := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("deposited omg: ", deposited)
	// require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPOMG, big.NewInt(1000000000)), deposited)

	fmt.Println("------------ step 3: execute trade OMG for POLY through Kyber aggregator --------------")
	tradingSuite.executeWithKyber(
		tradeAmount,
		tradingSuite.OMGAddressStr,
		tradingSuite.POLYAddressStr,
	)
	time.Sleep(15 * time.Second)
	polyTraded := tradingSuite.getDepositedBalance(
		tradingSuite.POLYAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("polyTraded: ", polyTraded)

	fmt.Println("------------ step 4: withdrawing POLY from SC to pPOLY on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.POLYAddressStr,
		polyTraded,
	)
	time.Sleep(15 * time.Second)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(90 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncPOLYTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	fmt.Println("------------ step 5: withdrawing pPOLY from Incognito to POLY --------------")
	withdrawingPPOLY := big.NewInt(0).Div(polyTraded, big.NewInt(1000000000))
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncPOLYTokenIDStr,
		withdrawingPPOLY,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))

	bal := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.POLYAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	tradingSuite.POLYBalanceAfterStep2 = bal
	fmt.Println("POLY balance after step 2: ", tradingSuite.POLYBalanceAfterStep2)
	// require.Equal(tradingSuite.T(), withdrawingPPOLY.Uint64(), bal.Uint64())
	// require.Equal(tradingSuite.T(), withdrawingPPOLY.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())
}

func (tradingSuite *KyberTradingTestSuite) Test3TradePOLYForEthWithKyber() {
	fmt.Println("============ TEST TRADE POLY FOR ETH WITH KYBER AGGREGATOR ===========")
	fmt.Println("------------ step 0: declaration & initialization --------------")
	depositingPOLY := tradingSuite.POLYBalanceAfterStep2
	burningPPOLY := big.NewInt(0).Div(depositingPOLY, big.NewInt(1000000000))
	tradeAmount := depositingPOLY

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("pubKeyToAddrStr: ", pubKeyToAddrStr)

	fmt.Println("------------ step 1: porting POLY to pPOLY --------------")
	txHash := tradingSuite.depositERC20ToBridge(
		depositingPOLY,
		common.HexToAddress(tradingSuite.POLYAddressStr),
		tradingSuite.IncPaymentAddrStr,
	)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(90 * time.Second)

	issuingRes, err := tradingSuite.callIssuingETHReq(
		tradingSuite.IncPOLYTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("issuingRes: ", issuingRes)
	time.Sleep(120 * time.Second)

	fmt.Println("------------ step 2: burning pPOLY to deposit POLY to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncPOLYTokenIDStr,
		burningPPOLY,
		pubKeyToAddrStr[2:],
		"createandsendburningfordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(140 * time.Second)

	tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))
	deposited := tradingSuite.getDepositedBalance(
		tradingSuite.POLYAddressStr,
		pubKeyToAddrStr,
	)
	require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPPOLY, big.NewInt(1000000000)), deposited)

	fmt.Println("------------ step 3: execute trade POLY for ETH through Kyber aggregator --------------")
	tradingSuite.executeWithKyber(
		tradeAmount,
		tradingSuite.POLYAddressStr,
		tradingSuite.EtherAddressStr,
	)
	etherTraded := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("etherTraded: ", etherTraded)

	fmt.Println("------------ step 4: withdrawing ETH from SC to pETH on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.EtherAddressStr,
		etherTraded,
	)
	time.Sleep(15 * time.Second)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(90 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(140 * time.Second)

	fmt.Println("------------ step 5: withdrawing pETH from Incognito to ETH --------------")
	withdrawingPETH := big.NewInt(0).Div(etherTraded, big.NewInt(1000000000))
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncEtherTokenIDStr,
		withdrawingPETH,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(140 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))

	bal := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("Ether balance after step 3: ", bal)
	// require.Equal(tradingSuite.T(), withdrawingPETH.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())
}
