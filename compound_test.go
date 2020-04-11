package main

import (
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/incognitochain/bridge-eth/bridge/compoundAgent"

	"github.com/incognitochain/bridge-eth/bridge/compoundLogic"

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
	tradingSuite.CompoundDeployedAddr = common.HexToAddress("0xaB58Ae78640C33CaD1FAf9Ca6902983622C926C6") //kovan
	tradingSuite.DepositingEther = float64(0.12)
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
	compounAgentLogicAbi, _ := abi.JSON(strings.NewReader(compoundLogic.CompoundLogicABI))

	srcToken := common.HexToAddress(srcTokenIDStr)
	destToken := common.HexToAddress(destTokenIDStr)
	inputAgent, _ := compounAgentLogicAbi.Pack("mint", destToken, srcQty)
	tradingSuite.callVault(srcToken, srcQty, destToken, inputAgent, "mint")
}

func (tradingSuite *CompoundTradingTestSuite) borrowCoin(
	collateralAmount *big.Int,
	borrowAmount *big.Int,
	cTokenBorrow string,
	srcTokenIDStr string,
	destTokenIDStr string,
	addCollaterals []common.Address,
) {
	compounAgentLogicAbi, _ := abi.JSON(strings.NewReader(compoundLogic.CompoundLogicABI))

	cBorrow := common.HexToAddress(cTokenBorrow)
	srcToken := common.HexToAddress(srcTokenIDStr)
	destToken := common.HexToAddress(destTokenIDStr)
	inputAgent, _ := compounAgentLogicAbi.Pack("borrow", cBorrow, borrowAmount, addCollaterals)
	tradingSuite.callVault(srcToken, collateralAmount, destToken, inputAgent, "borrow")
}

func (tradingSuite *CompoundTradingTestSuite) redeemCoin(
	amountSentFromVault *big.Int,
	amountRedeem *big.Int,
	srcTokenIDStr string,
	destTokenIDStr string,
	isUnderlyingToken bool,
	removecollateral common.Address,
) {
	compounAgentLogicAbi, _ := abi.JSON(strings.NewReader(compoundLogic.CompoundLogicABI))

	srcToken := common.HexToAddress(srcTokenIDStr)
	destToken := common.HexToAddress(destTokenIDStr)
	inputAgent, _ := compounAgentLogicAbi.Pack("redeem", srcToken, amountRedeem, isUnderlyingToken, removecollateral)
	tradingSuite.callVault(srcToken, amountSentFromVault, destToken, inputAgent, "redeem")
}

func (tradingSuite *CompoundTradingTestSuite) repayBorrow(
	srcQty *big.Int,
	srcTokenIDStr string,
	destTokenIDStr string,
) {
	compounAgentLogicAbi, _ := abi.JSON(strings.NewReader(compoundLogic.CompoundLogicABI))

	srcToken := common.HexToAddress(srcTokenIDStr)
	destToken := common.HexToAddress(destTokenIDStr)
	inputAgent, _ := compounAgentLogicAbi.Pack("repayBorrow", destToken, srcQty)
	tradingSuite.callVault(srcToken, srcQty, destToken, inputAgent, "repayBorrow")
}

func (tradingSuite *CompoundTradingTestSuite) liquidateBorrow(
	srcQty *big.Int,
	srcTokenIDStr string,
	destTokenIDStr string,
	borrower common.Address,
) {
	compounAgentLogicAbi, _ := abi.JSON(strings.NewReader(compoundLogic.CompoundLogicABI))

	srcToken := common.HexToAddress(srcTokenIDStr)
	destToken := common.HexToAddress(destTokenIDStr)
	inputAgent, _ := compounAgentLogicAbi.Pack("liquidateBorrow", destToken, borrower, srcQty, destToken)
	tradingSuite.callVault(srcToken, srcQty, destToken, inputAgent, "liquidateBorrow")
}

func (tradingSuite *CompoundTradingTestSuite) callVault(
	srcToken common.Address,
	srcQty *big.Int,
	destToken common.Address,
	inputAgent []byte,
	whocall string,
) {
	compounAgent, _ := abi.JSON(strings.NewReader(compoundAgent.CompoundAgentABI))
	inputAgent, _ = compounAgent.Pack("delegateCall", inputAgent)

	compounProxyAbi, _ := abi.JSON(strings.NewReader(compound.CompoundABI))
	timestamp := []byte(randomizeTimestamp())
	tempData := append(srcToken.Bytes(), inputAgent...)
	tempData1 := append(tempData, timestamp...)
	data := rawsha3(tempData1)
	// sign for proxy verification
	signBytes, _ := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)
	inputProxy, _ := compounProxyAbi.Pack("execute", srcToken, srcQty, inputAgent, timestamp, signBytes)
	tempData = append(tradingSuite.CompoundDeployedAddr[:], inputProxy...)
	tempData1 = append(tempData, timestamp...)
	data = rawsha3(tempData1)
	// sign for vault verification
	signBytes, _ = crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)
	// // Get contract instance
	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)
	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
	auth.GasPrice = big.NewInt(10000000000)
	// auth.GasLimit = 2500000
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

func (tradingSuite *CompoundTradingTestSuite) Test1TradeEthForKBNWithKyber() {
	fmt.Println("============ TEST COMPOUND ETHER FOR CETH WITH COMPOUND PROXY AGGREGATOR ===========")
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

	fmt.Println("Waiting 40s for 15 blocks confirmation")
	time.Sleep(40 * time.Second)
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
	require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPETH, big.NewInt(1000000000)), deposited)

	deposited := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("deposited EHT: ", deposited)

	fmt.Println("------------ step 3: mint cETH through Compound aggregator --------------")
	tradingSuite.mintCoin(
		big.NewInt(0).Div(tradeAmount, big.NewInt(int64(3))),
		tradingSuite.EtherAddressStr,
		tradingSuite.cETHAddressStr,
	)

	deposited = tradingSuite.getDepositedBalance(
		tradingSuite.cETHAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("After: deposited cEHT: ", deposited)

	fmt.Println("------------ step 4: collateral cETH and borrow DAI through Compound aggregator --------------")
	tradingSuite.borrowCoin(
		big.NewInt(0).Div(deposited, big.NewInt(int64(3))),
		big.NewInt(10000000),
		tradingSuite.cDAIAddressStr,
		tradingSuite.cETHAddressStr,
		tradingSuite.DAIAddressStr,
		[]common.Address{common.HexToAddress(tradingSuite.cETHAddressStr)},
	)

	DAIBorrowed := tradingSuite.getDepositedBalance(
		tradingSuite.DAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("DAI borrowed: ", DAIBorrowed)

	fmt.Println("------------ step 5: Redeem cETH through Compound aggregator --------------")
	tradingSuite.redeemCoin(
		big.NewInt(0).Div(deposited, big.NewInt(int64(3))),
		big.NewInt(0).Div(deposited, big.NewInt(int64(3))),
		tradingSuite.cETHAddressStr,
		tradingSuite.EtherAddressStr,
		false,
		common.Address{},
	)

	ETHredeem := tradingSuite.getDepositedBalance(
		tradingSuite.DAIAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("ETH after call redeem: ", ETHredeem)

	fmt.Println("------------ step 6: Borrow ETH and Repay ETH through Compound aggregator --------------")

	tradingSuite.borrowCoin(
		big.NewInt(0).Div(deposited, big.NewInt(int64(3))),
		big.NewInt(10000000),
		tradingSuite.cETHAddressStr,
		tradingSuite.cETHAddressStr,
		tradingSuite.EtherAddressStr,
		[]common.Address{},
	)

	tradingSuite.repayBorrow(
		big.NewInt(10000000),
		tradingSuite.EtherAddressStr,
		tradingSuite.cETHAddressStr,
	)
}
