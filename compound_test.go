package main

import (
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/incognitochain/bridge-eth/bridge/compoundAgent"
	"github.com/incognitochain/bridge-eth/bridge/vault"

	"github.com/incognitochain/bridge-eth/bridge/compound"

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
type CompoundTradingTestSuite struct {
	*TradingTestSuite

	CompoundDeployedAddr common.Address
	CompoundController   common.Address

	InccBATTokenIDStr string
	InccDAITokenIDStr string
	InccETHTokenIDStr string
	InccSAITokenIDStr string

	cBATAddressStr string
	cDAIAddressStr string
	cETHAddressStr string
	cREPAddressStr string
	cSAIAddressStr string

	// token amounts for tests
	DepositingEther       float64
	KBNBalanceAfterStep1  *big.Int
	SALTBalanceAfterStep2 *big.Int
}

func NewCompoundTradingTestSuite(tradingTestSuite *TradingTestSuite) *CompoundTradingTestSuite {
	return &CompoundTradingTestSuite{
		TradingTestSuite: tradingTestSuite,
	}
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (tradingSuite *CompoundTradingTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")
	// Kovan env

	tradingSuite.InccBATTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000062"
	tradingSuite.InccDAITokenIDStr = "0000000000000000000000000000000000000000000000000000000000000061"
	tradingSuite.InccETHTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000052"
	tradingSuite.InccSAITokenIDStr = "0000000000000000000000000000000000000000000000000000000000000051"
	tradingSuite.cBATAddressStr = "0xd5ff020f970462816fdd31a603cb7d120e48376e"                            // kovan
	tradingSuite.cDAIAddressStr = "0xe7bc397dbd069fc7d0109c0636d06888bb50668c"                            // kovan
	tradingSuite.cETHAddressStr = "0xf92fbe0d3c0dcdae407923b2ac17ec223b1084e4"                            // kovan
	tradingSuite.cREPAddressStr = "0xfd874be7e6733bdc6dca9c7cdd97c225ec235d39"                            // kovan
	tradingSuite.CompoundDeployedAddr = common.HexToAddress("0xA57854f84C2979027BE534112Ce9005F92f06c36") //kovan
	tradingSuite.CompoundController = common.HexToAddress("0x1f5D7F3CaAC149fE41b8bd62A3673FE6eC0AB73b")
	tradingSuite.DepositingEther = float64(0.2)
	tradingSuite.KyberContractAddr = common.HexToAddress("0x692f391bCc85cefCe8C237C01e1f636BbD70EA4D") // kovan
}

func (tradingSuite *CompoundTradingTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	tradingSuite.ETHClient.Close()
}

func (tradingSuite *CompoundTradingTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (tradingSuite *CompoundTradingTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestCompoundTradingTestSuite(t *testing.T) {
	fmt.Println("Starting entry point for compound test suite...")

	tradingSuite := new(TradingTestSuite)
	suite.Run(t, tradingSuite)

	kyberTradingSuite := NewCompoundTradingTestSuite(tradingSuite)
	suite.Run(t, kyberTradingSuite)

	fmt.Println("Finishing entry point for compound test suite...")
}

func (tradingSuite *CompoundTradingTestSuite) mintCoin(
	srcQty *big.Int,
	srcTokenIDStr string,
	destTokenIDStr string,
) {
	compounProxyAbi, _ := abi.JSON(strings.NewReader(compound.CompoundABI))
	compounAgentAbi, _ := abi.JSON(strings.NewReader(compoundAgent.CompoundAgentABI))

	srcToken := common.HexToAddress(srcTokenIDStr)
	destToken := common.HexToAddress(destTokenIDStr)
	inputAgent, _ := compounAgentAbi.Pack("mint", destToken, srcQty)
	timestamp := []byte(randomizeTimestamp())
	tempData := append(srcToken.Bytes(), inputAgent...)
	tempData1 := append(tempData, timestamp...)
	data := rawsha3(tempData1)
	// sign for proxy verification
	signBytes, _ := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)
	inputProxy, _ := compounProxyAbi.Pack("execute", srcToken, srcQty, inputAgent, timestamp, signBytes)
	tradingSuite.callVault(srcToken, srcQty, destToken, inputProxy, timestamp, "mint")
}

func (tradingSuite *CompoundTradingTestSuite) borrowCoin(
	srcQty *big.Int,
	srcTokenIDStr string,
	destTokenIDStr string,
	addCollaterals []common.Address,
) {
	compounProxyAbi, _ := abi.JSON(strings.NewReader(compound.CompoundABI))
	compounAgentAbi, _ := abi.JSON(strings.NewReader(compoundAgent.CompoundAgentABI))

	srcToken := common.HexToAddress(srcTokenIDStr)
	destToken := common.HexToAddress(destTokenIDStr)
	inputAgent, _ := compounAgentAbi.Pack("borrow", destToken, srcQty, addCollaterals)
	timestamp := []byte(randomizeTimestamp())
	tempData := append(srcToken.Bytes(), inputAgent...)
	tempData1 := append(tempData, timestamp...)
	data := rawsha3(tempData1)
	// sign for proxy verification
	signBytes, _ := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)
	inputProxy, _ := compounProxyAbi.Pack("execute", srcToken, srcQty, inputAgent, timestamp, signBytes)
	tradingSuite.callVault(srcToken, srcQty, destToken, inputProxy, timestamp, "borrow")
}

func (tradingSuite *CompoundTradingTestSuite) redeemCoin(
	srcQty *big.Int,
	srcTokenIDStr string,
	destTokenIDStr string,
	isUnderlyingToken bool,
	removecollateral common.Address,
) {
	compounProxyAbi, _ := abi.JSON(strings.NewReader(compound.CompoundABI))
	compounAgentAbi, _ := abi.JSON(strings.NewReader(compoundAgent.CompoundAgentABI))

	srcToken := common.HexToAddress(srcTokenIDStr)
	destToken := common.HexToAddress(destTokenIDStr)
	inputAgent, _ := compounAgentAbi.Pack("redeem", destToken, srcQty, isUnderlyingToken, removecollateral)
	timestamp := []byte(randomizeTimestamp())
	tempData := append(srcToken.Bytes(), inputAgent...)
	tempData1 := append(tempData, timestamp...)
	data := rawsha3(tempData1)
	// sign for proxy verification
	signBytes, _ := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)
	inputProxy, _ := compounProxyAbi.Pack("execute", srcToken, srcQty, inputAgent, timestamp, signBytes)
	tradingSuite.callVault(srcToken, srcQty, destToken, inputProxy, timestamp, "redeem")
}

func (tradingSuite *CompoundTradingTestSuite) repayBorrow(
	srcQty *big.Int,
	srcTokenIDStr string,
	destTokenIDStr string,
) {
	compounProxyAbi, _ := abi.JSON(strings.NewReader(compound.CompoundABI))
	compounAgentAbi, _ := abi.JSON(strings.NewReader(compoundAgent.CompoundAgentABI))

	srcToken := common.HexToAddress(srcTokenIDStr)
	destToken := common.HexToAddress(destTokenIDStr)
	inputAgent, _ := compounAgentAbi.Pack("repayBorrow", destToken, srcQty)
	timestamp := []byte(randomizeTimestamp())
	tempData := append(srcToken.Bytes(), inputAgent...)
	tempData1 := append(tempData, timestamp...)
	data := rawsha3(tempData1)
	// sign for proxy verification
	signBytes, _ := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)
	inputProxy, _ := compounProxyAbi.Pack("execute", srcToken, srcQty, inputAgent, timestamp, signBytes)
	tradingSuite.callVault(srcToken, srcQty, destToken, inputProxy, timestamp, "repayBorrow")
}

func (tradingSuite *CompoundTradingTestSuite) liquidateBorrow(
	srcQty *big.Int,
	srcTokenIDStr string,
	destTokenIDStr string,
	borrower common.Address,
	CCollateral common.Address,
) {
	compounProxyAbi, _ := abi.JSON(strings.NewReader(compound.CompoundABI))
	compounAgentAbi, _ := abi.JSON(strings.NewReader(compoundAgent.CompoundAgentABI))

	srcToken := common.HexToAddress(srcTokenIDStr)
	destToken := common.HexToAddress(destTokenIDStr)
	inputAgent, _ := compounAgentAbi.Pack("liquidateBorrow", destToken, borrower, srcQty, CCollateral)
	timestamp := []byte(randomizeTimestamp())
	tempData := append(srcToken.Bytes(), inputAgent...)
	tempData1 := append(tempData, timestamp...)
	data := rawsha3(tempData1)
	// sign for proxy verification
	signBytes, _ := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)
	inputProxy, _ := compounProxyAbi.Pack("execute", srcToken, srcQty, inputAgent, timestamp, signBytes)
	tradingSuite.callVault(srcToken, srcQty, destToken, inputProxy, timestamp, "liquidateBorrow")
}

func (tradingSuite *CompoundTradingTestSuite) callVault(
	srcToken common.Address,
	srcQty *big.Int,
	destToken common.Address,
	inputProxy []byte,
	timestamp []byte,
	whocall string,
) {

	tempData := append(tradingSuite.CompoundDeployedAddr[:], inputProxy...)
	tempData1 := append(tempData, timestamp...)
	data := rawsha3(tempData1)
	// sign for vault verification
	signBytes, _ := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)
	// // Get contract instance
	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)
	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
	auth.GasPrice = big.NewInt(10000000000)
	tx, err := c.Execute(
		auth,
		srcToken,
		srcQty,
		destToken,
		tradingSuite.CompoundDeployedAddr,
		inputProxy,
		timestamp,
		signBytes,
	)
	require.Equal(tradingSuite.T(), nil, err)
	txHash := tx.Hash()
	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("%v token executed , txHash: %x\n", whocall, txHash[:])
}

// func (tradingSuite *CompoundTradingTestSuite) executeMultiTradeWithKyber(
// 	srcQties []*big.Int,
// 	srcTokenIDStrs []string,
// 	destTokenIDStrs []string,
// ) {
// 	tradeAbi, _ := abi.JSON(strings.NewReader(kbnmultiTrade.KbnmultiTradeABI))
// 	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
// 	auth.GasPrice = big.NewInt(50000000000)
// 	auth.GasLimit = 2000000
// 	// Deploy kbnMultitrade
// 	// kbnMultiTradeAddr, tx, _, err := kbnmultiTrade.DeployKbnmultiTrade(auth, tradingSuite.ETHClient, tradingSuite.KyberContractAddr, tradingSuite.VaultAddr)
// 	// require.Equal(tradingSuite.T(), nil, err)
// 	// fmt.Println("deployed kbnMultitrade")
// 	// fmt.Printf("addr: %s\n", kbnMultiTradeAddr.Hex())
// 	// tradingSuite.KyberMultiTradeDeployedAddr = kbnMultiTradeAddr

// 	// Get contract instance
// 	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	auth.GasPrice = big.NewInt(50000000000)
// 	auth.GasLimit = 5000000
// 	sourceAddresses := make([]common.Address, 0)
// 	for _, p := range srcTokenIDStrs {
// 		sourceAddresses = append(sourceAddresses, common.HexToAddress(p))
// 	}
// 	destAddresses := make([]common.Address, 0)
// 	for _, p := range destTokenIDStrs {
// 		destAddresses = append(destAddresses, common.HexToAddress(p))
// 	}
// 	expectRates := make([]*big.Int, 0)
// 	for i := range destTokenIDStrs {
// 		expectRates = append(expectRates, tradingSuite.getExpectedRate(srcTokenIDStrs[i], destTokenIDStrs[i], srcQties[i]))
// 	}

// 	input, _ := tradeAbi.Pack("trade", sourceAddresses, srcQties, destAddresses, expectRates)
// 	timestamp := []byte(randomizeTimestamp())
// 	tempData := append(tradingSuite.KyberMultiTradeDeployedAddr[:], input...)
// 	tempData1 := append(tempData, timestamp...)
// 	data := rawsha3(tempData1)
// 	signBytes, _ := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)

// 	tx, err := c.ExecuteMulti(
// 		auth,
// 		sourceAddresses,
// 		srcQties,
// 		destAddresses,
// 		tradingSuite.KyberMultiTradeDeployedAddr,
// 		input,
// 		timestamp,
// 		signBytes,
// 	)
// 	require.Equal(tradingSuite.T(), nil, err)
// 	txHash := tx.Hash()
// 	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
// 		require.Equal(tradingSuite.T(), nil, err)
// 	}
// 	fmt.Printf("Kyber multi trade executed , txHash: %x\n", txHash[:])
// }

func (tradingSuite *CompoundTradingTestSuite) Test1TradeEthForKBNWithKyber() {
	fmt.Println("============ TEST TRADE ETHER FOR KBN WITH Kyber AGGREGATOR ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("------------ STEP 1: porting ETH to pETH --------------")
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
	)
	// time.Sleep(15 * time.Second)
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
	// require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPETH, big.NewInt(1000000000)), deposited)

	fmt.Println("------------ step 3: mint cETH through Compound aggregator --------------")
	tradingSuite.mintCoin(
		tradeAmount,
		tradingSuite.EtherAddressStr,
		tradingSuite.cETHAddressStr,
	)
	// time.Sleep(15 * time.Second)
	// kbnTraded := tradingSuite.getDepositedBalance(
	// 	tradingSuite.KBNAddressStr,
	// 	pubKeyToAddrStr,
	// )
	// fmt.Println("kbnTraded: ", kbnTraded)

	// fmt.Println("------------ step 4: withdrawing KBN from SC to pKBN on Incognito --------------")
	// txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
	// 	tradingSuite.KBNAddressStr,
	// 	kbnTraded,
	// )
	// time.Sleep(15 * time.Second)

	// _, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
	// require.Equal(tradingSuite.T(), nil, err)
	// fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	// fmt.Println("Waiting 90s for 15 blocks confirmation")
	// time.Sleep(90 * time.Second)
	// _, err = tradingSuite.callIssuingETHReq(
	// 	tradingSuite.IncKBNTokenIDStr,
	// 	ethDepositProof,
	// 	ethBlockHash,
	// 	ethTxIdx,
	// )
	// require.Equal(tradingSuite.T(), nil, err)
	// time.Sleep(120 * time.Second)

	// fmt.Println("------------ step 5: withdrawing pKBN from Incognito to KBN --------------")
	// withdrawingPKBN := big.NewInt(0).Div(kbnTraded, big.NewInt(1000000000))
	// burningRes, err = tradingSuite.callBurningPToken(
	// 	tradingSuite.IncKBNTokenIDStr,
	// 	withdrawingPKBN,
	// 	tradingSuite.ETHOwnerAddrStr,
	// 	"createandsendburningrequest",
	// )
	// require.Equal(tradingSuite.T(), nil, err)
	// burningTxID, found = burningRes["TxID"]
	// require.Equal(tradingSuite.T(), true, found)
	// time.Sleep(120 * time.Second)

	// tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))

	// bal := tradingSuite.getBalanceOnETHNet(
	// 	common.HexToAddress(tradingSuite.KBNAddressStr),
	// 	common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	// )
	// tradingSuite.KBNBalanceAfterStep1 = bal
	// fmt.Println("KBN balance after step 1: ", tradingSuite.KBNBalanceAfterStep1)
	// require.Equal(tradingSuite.T(), withdrawingPKBN.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())
}
