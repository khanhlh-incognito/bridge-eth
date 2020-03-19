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

	kbnmultiTrade "github.com/incognitochain/bridge-eth/bridge/kbnmultitrade"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type KovanKyberTradingTestSuite struct {
	*TradingTestSuite

	KyberTradeDeployedAddr common.Address

	KyberContractAddr    common.Address
	WETHAddr             common.Address
	EtherAddressStrKyber string

	IncKBNTokenIDStr  string
	IncSALTTokenIDStr string
	IncOMGTokenIDStr  string
	IncSNTTokenIDStr  string
	IncPOLYTokenIDStr string
	IncZILTokenIDStr  string

	KBNAddressStr  string
	SALTAddressStr string
	OMGAddressStr  string
	SNTAddressStr  string
	POLYAddressStr string
	ZILAddressStr  string

	IncPrivKeyStrKb     string
	IncPaymentAddrStrKb string

	// token amounts for tests
	DepositingEther             float64
	OMGBalanceAfterStep1        *big.Int
	POLYBalanceAfterStep2       *big.Int
	KyberMultiTradeDeployedAddr common.Address
}

func NewKovanKyberTradingTestSuite(tradingTestSuite *TradingTestSuite) *KovanKyberTradingTestSuite {
	return &KovanKyberTradingTestSuite{
		TradingTestSuite: tradingTestSuite,
	}
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (tradingSuite *KovanKyberTradingTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")
	// Kovan env
	tradingSuite.IncPrivKeyStrKb = "112t8ro4yu78UE82jpto12rp3Cd8Z2Mse7fcavSyXXP82oApE1cg9y8hWq69Z74fFHGJrQyHz54vU8Mv1kx5qZ54cRQJPkF5Cb7DhNqL5Tgt"
	tradingSuite.IncPaymentAddrStrKb = "12RyGbTyktYkXe7mNwmZeD4rktqxtHMe3Tsyf4XiZdKVGFssEHaF1ZUTpXZmpFACuDotVr7a6FEw8v6FTn8DEMqpHNxZ8fJW3KNN3i1"

	tradingSuite.EtherAddressStrKyber = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	tradingSuite.IncPOLYTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000100"
	tradingSuite.IncOMGTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000101"
	tradingSuite.IncZILTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000102"
	tradingSuite.IncKBNTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000082"
	tradingSuite.IncSALTTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000081"
	tradingSuite.IncSNTTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000071"

	tradingSuite.ZILAddressStr = "0xAb74653cac23301066ABa8eba62b9Abd8a8c51d6"
	tradingSuite.POLYAddressStr = "0xd92266fd053161115163a7311067F0A4745070b5"
	tradingSuite.OMGAddressStr = "0xdB7ec4E4784118D9733710e46F7C83fE7889596a"
	tradingSuite.KBNAddressStr = "0xad67cB4d63C9da94AcA37fDF2761AaDF780ff4a2" // kovan
	tradingSuite.SALTAddressStr = "0x6fEE5727EE4CdCBD91f3A873ef2966dF31713A04"
	tradingSuite.SNTAddressStr = "0x4c99B04682fbF9020Fcb31677F8D8d66832d3322"

	tradingSuite.KyberTradeDeployedAddr = common.HexToAddress("0x2646d88aa309AD221045d8823f57b4456a4D4BD4")      //kovan
	tradingSuite.KyberMultiTradeDeployedAddr = common.HexToAddress("0x7e8b944741644f4Be04F13081D7676C3a6B8f39d") //kovan
	tradingSuite.DepositingEther = float64(0.01)
	tradingSuite.KyberContractAddr = common.HexToAddress("0x692f391bCc85cefCe8C237C01e1f636BbD70EA4D") // kovan
}

func (tradingSuite *KovanKyberTradingTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	tradingSuite.ETHClient.Close()
}

func (tradingSuite *KovanKyberTradingTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (tradingSuite *KovanKyberTradingTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestKovanKyberTradingTestSuite(t *testing.T) {
	fmt.Println("Starting entry point for Kyber test suite...")

	tradingSuite := new(TradingTestSuite)
	suite.Run(t, tradingSuite)

	kyberTradingSuite := NewKovanKyberTradingTestSuite(tradingSuite)
	suite.Run(t, kyberTradingSuite)

	fmt.Println("Finishing entry point for 0x test suite...")
}

func (tradingSuite *KovanKyberTradingTestSuite) executeWithKyber(
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

func (tradingSuite *KovanKyberTradingTestSuite) executeMultiTradeWithKyber(
	srcQties []*big.Int,
	srcTokenIDStrs []string,
	destTokenIDStrs []string,
) {
	tradeAbi, _ := abi.JSON(strings.NewReader(kbnmultiTrade.KbnmultiTradeABI))
	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
	auth.GasPrice = big.NewInt(50000000000)
	auth.GasLimit = 2000000
	// Deploy kbnMultitrade
	// kbnMultiTradeAddr, tx, _, err := kbnmultiTrade.DeployKbnmultiTrade(auth, tradingSuite.ETHClient, tradingSuite.KyberContractAddr, tradingSuite.VaultAddr)
	// require.Equal(tradingSuite.T(), nil, err)
	// fmt.Println("deployed kbnMultitrade")
	// fmt.Printf("addr: %s\n", kbnMultiTradeAddr.Hex())
	// tradingSuite.KyberMultiTradeDeployedAddr = kbnMultiTradeAddr

	// Get contract instance
	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)
	auth.GasPrice = big.NewInt(50000000000)
	auth.GasLimit = 5000000
	sourceAddresses := make([]common.Address, 0)
	for _, p := range srcTokenIDStrs {
		sourceAddresses = append(sourceAddresses, common.HexToAddress(p))
	}
	destAddresses := make([]common.Address, 0)
	for _, p := range destTokenIDStrs {
		destAddresses = append(destAddresses, common.HexToAddress(p))
	}
	input, _ := tradeAbi.Pack("trade", sourceAddresses, srcQties, destAddresses)
	timestamp := []byte(randomizeTimestamp())
	tempData := append(tradingSuite.KyberMultiTradeDeployedAddr[:], input...)
	tempData1 := append(tempData, timestamp...)
	data := rawsha3(tempData1)
	signBytes, _ := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)

	tx, err := c.ExecuteMulti(
		auth,
		sourceAddresses,
		srcQties,
		destAddresses,
		tradingSuite.KyberMultiTradeDeployedAddr,
		input,
		timestamp,
		signBytes,
	)
	require.Equal(tradingSuite.T(), nil, err)
	txHash := tx.Hash()
	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("Kyber multi trade executed , txHash: %x\n", txHash[:])
}

func (tradingSuite *KovanKyberTradingTestSuite) Test1TradeEthForKNCWithKyber() {
	return
	fmt.Println("============ TEST 1 TRADE ETHER FOR KNC WITH Kyber AGGREGATOR ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()

	// get info balance initialization
	balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)

	balpEthInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance initialization : ", balpEthInit)

	balpKNCInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncKNCTokenIDStr)
	fmt.Println("[INFO] pKNC balance initialization : ", balpKNCInit)

	balEthInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance initialization : ", balEthInit)

	balKNCInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.KNCAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] KNC balance initialization : ", balKNCInit)

	balEthScInit := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance initialization on SC : ", balEthScInit)

	balKNCScInit := tradingSuite.getDepositedBalance(
		tradingSuite.KNCAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] KNC balance initialization on SC : ", balKNCScInit)

	fmt.Println("------------ STEP 1: porting ETH to pETH --------------")
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
	)
	time.Sleep(15 * time.Second)
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	// get ETH remain after depsit
	balEthAfDep := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain after deposit : ", balEthAfDep)
	// TODO : assert ETH balance

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(10 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(70 * time.Second)
	// check PRV and token balance after issuing
	balPrvAfIssS1, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 1: ", balPrvAfIssS1)
	// TODO assert PRV balance remain
	balpEthAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance after issuing step 1 : ", balpEthAfIssS1)
	// TODO assert pETH balance issuing

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
	time.Sleep(70 * time.Second)

	// check PRV and token balance after burning
	balPrvAfBurnS2, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after burning step 2: ", balPrvAfBurnS2)
	// TODO assert PRV balance remain
	balpEthAfBurnS2, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance after burning step 2 : ", balpEthAfBurnS2)
	// TODO assert pETH balance issuing

	tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))
	balEthScDepS2 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance after deposit on SC at step 2: ", balEthScDepS2)
	// TODO assert ETH balane on SC
	balKNCScS2 := tradingSuite.getDepositedBalance(
		tradingSuite.KNCAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] KNC balance on SC at step 2 : ", balKNCScS2)

	//require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPETH, big.NewInt(1000000000)), deposited)
	fmt.Println("------------ step 3: execute trade ETH for KNC through Kyber aggregator --------------")
	tradingSuite.executeWithKyber(
		tradeAmount,
		tradingSuite.EtherAddressStr,
		tradingSuite.KNCAddressStr,
	)
	balEthScTradeS3 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after trade at step 3 : ", balEthScTradeS3)
	// TODO assert ETH balane on SC
	balKNCScTradeS3 := tradingSuite.getDepositedBalance(
		tradingSuite.KNCAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] KNC balance on SC after trade at step 3 : ", balKNCScTradeS3)
	// TODO assert KNC balane on SC
	require.NotEqual(tradingSuite.T(), balKNCScTradeS3, balKNCScS2, "trade failed")

	fmt.Println("------------ step 4: withdrawing KNC from SC to pKNC on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.KNCAddressStr,
		balKNCScTradeS3,
	)
	time.Sleep(15 * time.Second)

	balKNCScDepS4 := tradingSuite.getDepositedBalance(
		tradingSuite.KNCAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] KNC balance on SC after withdraw at step 4 : ", balKNCScDepS4)
	// TODO assert KNC balane on SC
	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(10 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncKNCTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(70 * time.Second)

	balpEthAfIssS4, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncKNCTokenIDStr,
	)
	fmt.Println("[INFO] pKNC balance after issuing step 4 : ", balpEthAfIssS4)

	balPrvAfIssS4, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 4: ", balPrvAfIssS4)
	// TODO assert PRV balance remain

	fmt.Println("------------ step 5: withdrawing pKNC from Incognito to KNC --------------")
	withdrawingPKNC := big.NewInt(0).Div(balKNCScTradeS3, big.NewInt(1000000000))
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncKNCTokenIDStr,
		withdrawingPKNC,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(70 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))

	balKNC := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.KNCAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	//tradingSuite.KNCBalanceAfterStep1 = balKNC
	fmt.Println("KNC balance after trade: ", balKNC)
	// require.Equal(tradingSuite.T(), withdrawingPKNC.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())
	balEth := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("ETH balance after trade: ", balEth)
}

func (tradingSuite *KovanKyberTradingTestSuite) Test2TradeEthForOMGWithKyber() {
	return
	fmt.Println("============ TEST 2 TRADE ETHER FOR OMG WITH Kyber AGGREGATOR ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()

	// get info balance initialization
	balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)

	balpEthInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance initialization : ", balpEthInit)

	balpOMGInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncOMGTokenIDStr)
	fmt.Println("[INFO] pOMG balance initialization : ", balpOMGInit)

	balEthInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance initialization : ", balEthInit)

	balOMGInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.OMGAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] OMG balance initialization : ", balOMGInit)

	balEthScInit := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance initialization on SC : ", balEthScInit)

	balOMGScInit := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance initialization on SC : ", balOMGScInit)

	fmt.Println("------------ STEP 1: porting ETH to pETH --------------")
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
	)
	time.Sleep(15 * time.Second)
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	// get ETH remain after depsit
	balEthAfDep := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain after deposit : ", balEthAfDep)
	// TODO : assert ETH balance

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(10 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(70 * time.Second)
	// check PRV and token balance after issuing
	balPrvAfIssS1, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 1: ", balPrvAfIssS1)
	// TODO assert PRV balance remain
	balpEthAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance after issuing step 1 : ", balpEthAfIssS1)
	// TODO assert pETH balance issuing

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
	time.Sleep(70 * time.Second)

	// check PRV and token balance after burning
	balPrvAfBurnS2, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after burning step 2: ", balPrvAfBurnS2)
	// TODO assert PRV balance remain
	balpEthAfBurnS2, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance after burning step 2 : ", balpEthAfBurnS2)
	// TODO assert pETH balance issuing

	tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))
	balEthScDepS2 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance after deposit on SC at step 2: ", balEthScDepS2)
	// TODO assert ETH balane on SC
	balOMGScS2 := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance on SC at step 2 : ", balOMGScS2)

	//require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPETH, big.NewInt(1000000000)), deposited)
	fmt.Println("------------ step 3: execute trade ETH for OMG through Kyber aggregator --------------")
	tradingSuite.executeWithKyber(
		tradeAmount,
		tradingSuite.EtherAddressStr,
		tradingSuite.OMGAddressStr,
	)
	balEthScTradeS3 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after trade at step 3 : ", balEthScTradeS3)
	// TODO assert ETH balane on SC
	balOMGScTradeS3 := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance on SC after trade at step 3 : ", balOMGScTradeS3)
	// TODO assert OMG balane on SC

	fmt.Println("------------ step 4: withdrawing OMG from SC to pOMG on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.OMGAddressStr,
		balOMGScTradeS3,
	)
	time.Sleep(15 * time.Second)

	balOMGScDepS4 := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance on SC after withdraw at step 4 : ", balOMGScDepS4)
	// TODO assert OMG balane on SC
	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(10 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncOMGTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(70 * time.Second)

	balpEthAfIssS4, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncOMGTokenIDStr,
	)
	fmt.Println("[INFO] pOMG balance after issuing step 4 : ", balpEthAfIssS4)

	balPrvAfIssS4, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 4: ", balPrvAfIssS4)
	// TODO assert PRV balance remain

	fmt.Println("------------ step 5: withdrawing pOMG from Incognito to OMG --------------")
	withdrawingPOMG := big.NewInt(0).Div(balOMGScTradeS3, big.NewInt(1000000000))
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncOMGTokenIDStr,
		withdrawingPOMG,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(70 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))

	balOMG := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.OMGAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	tradingSuite.OMGBalanceAfterStep1 = balOMG
	fmt.Println("OMG balance after trade: ", balOMG)
	// require.Equal(tradingSuite.T(), withdrawingPOMG.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())
	balEth := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("ETH balance after trade: ", balEth)
}

func (tradingSuite *KovanKyberTradingTestSuite) Test3TradeOMGForZILWithKyber() {
	return
	fmt.Println("============ TEST 3 TRADE OMG FOR ZIL WITH KYBER AGGREGATOR ===========")
	fmt.Println("------------ step 0: declaration & initialization --------------")
	depositingOMG := tradingSuite.OMGBalanceAfterStep1
	burningPOMG := big.NewInt(0).Div(depositingOMG, big.NewInt(1000000000))
	tradeAmount := depositingOMG

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("pubKeyToAddrStr: ", pubKeyToAddrStr)

	// get info balance initialization
	balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)

	balpZILInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncZILTokenIDStr)
	fmt.Println("[INFO] pZIL balance initialization : ", balpZILInit)

	balpOMGInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncOMGTokenIDStr)
	fmt.Println("[INFO] pOMG balance initialization : ", balpOMGInit)

	balZILInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.ZILAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ZIL balance initialization : ", balZILInit)

	balOMGInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.OMGAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] OMG balance initialization : ", balOMGInit)

	balZILScInit := tradingSuite.getDepositedBalance(
		tradingSuite.ZILAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ZIL balance initialization on SC : ", balZILScInit)

	balOMGScInit := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance initialization on SC : ", balOMGScInit)

	fmt.Println("------------ step 1: porting OMG to pOMG --------------")
	txHash := tradingSuite.depositERC20ToBridge(
		depositingOMG,
		common.HexToAddress(tradingSuite.OMGAddressStr),
		tradingSuite.IncPaymentAddrStr,
	)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	// get OMG remain after depsit
	balOMGAfDep := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.OMGAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] OMG balance remain after deposit : ", balOMGAfDep)
	// TODO : assert OMG balance

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(10 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncOMGTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(70 * time.Second)
	// check PRV and token balance after issuing
	balPrvAfIssS1, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 1: ", balPrvAfIssS1)
	// TODO assert PRV balance remain
	balpOMGAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncOMGTokenIDStr)
	fmt.Println("[INFO] pOMG balance after issuing step 1 : ", balpOMGAfIssS1)
	// TODO assert pOMG balance issuing

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
	time.Sleep(70 * time.Second)

	// check PRV and token balance after burning
	balPrvAfBurnS2, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after burning step 2: ", balPrvAfBurnS2)
	// TODO assert PRV balance remain
	balpOMGAfBurnS2, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncOMGTokenIDStr)
	fmt.Println("[INFO] pOMG balance after burning step 2 : ", balpOMGAfBurnS2)
	// TODO assert pOMG balance issuing

	tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))

	balOMGScDepS2 := tradingSuite.getDepositedBalance(
		tradingSuite.IncOMGTokenIDStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance after deposit on SC before trade at step 2: ", balOMGScDepS2)
	// TODO assert OMG balane on SC
	// require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPOMG, big.NewInt(1000000000)), deposited)
	balZILScS2 := tradingSuite.getDepositedBalance(
		tradingSuite.IncZILTokenIDStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ZIL balance on SC before trade at step 2 : ", balZILScS2)

	fmt.Println("------------ step 3: execute trade OMG for ZIL through Kyber aggregator --------------")
	tradingSuite.executeWithKyber(
		tradeAmount,
		tradingSuite.OMGAddressStr,
		tradingSuite.ZILAddressStr,
	)
	time.Sleep(15 * time.Second)
	balZILScTradeS3 := tradingSuite.getDepositedBalance(
		tradingSuite.ZILAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ZIL balance on SC after trade at step 3 : ", balZILScTradeS3)
	// TODO assert ZIL balane on SC
	balOMGScTradeS3 := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance on SC after trade at step 3 : ", balOMGScTradeS3)
	// TODO assert OMG balane on SC

	fmt.Println("------------ step 4: withdrawing ZIL from SC to pZIL on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.ZILAddressStr,
		balZILScTradeS3,
	)
	time.Sleep(15 * time.Second)

	balZILScDepS4 := tradingSuite.getDepositedBalance(
		tradingSuite.ZILAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] pZIL balance on SC after withdraw at step 4 : ", balZILScDepS4)
	// TODO assert ZIL balane on SC

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(10 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncZILTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(70 * time.Second)

	balpZILAfIssS4, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncZILTokenIDStr,
	)
	fmt.Println("[INFO] pZILbalance after issuing step 4 : ", balpZILAfIssS4)
	// TODO assert pZIL balance issuing
	balPrvAfIssS4, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 4: ", balPrvAfIssS4)
	// TODO assert PRV balance remain

	fmt.Println("------------ step 5: withdrawing pZIL from Incognito to ZIL --------------")
	withdrawingPZIL := big.NewInt(0).Div(balZILScTradeS3, big.NewInt(1000000000))
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncZILTokenIDStr,
		withdrawingPZIL,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(70 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(burningTxID.(string))

	balZIL := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.ZILAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	tradingSuite.POLYBalanceAfterStep2 = balZIL
	fmt.Println("ZIL balance after step 2: ", balZIL)
	// require.Equal(tradingSuite.T(), withdrawingPZIL.Uint64(), bal.Uint64())
	// require.Equal(tradingSuite.T(), withdrawingPZIL.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())
	balOMG := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.OMGAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("OMG balance after trade: ", balOMG)
}

func (tradingSuite *KovanKyberTradingTestSuite) Test4TradeZILForEthWithKyber() {
	return
	fmt.Println("============ TEST TRADE ZIL FOR ETH WITH KYBER AGGREGATOR ===========")
	fmt.Println("------------ step 0: declaration & initialization --------------")
	depositingZIL := tradingSuite.POLYBalanceAfterStep2
	burningPZIL := big.NewInt(0).Div(depositingZIL, big.NewInt(1000000000))
	tradeAmount := depositingZIL

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("pubKeyToAddrStr: ", pubKeyToAddrStr)

	// get info balance initialization
	balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)

	balpZILInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncZILTokenIDStr)
	fmt.Println("[INFO] pZIL balance initialization : ", balpZILInit)

	balpEthInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance initialization : ", balpEthInit)

	balZILInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.ZILAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ZIL balance initialization : ", balZILInit)

	balEthInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance initialization : ", balEthInit)

	balEthScInit := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance initialization on SC : ", balEthScInit)

	balZILScInit := tradingSuite.getDepositedBalance(
		tradingSuite.ZILAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ZIL balance initialization on SC : ", balZILScInit)

	fmt.Println("------------ step 1: porting ZIL to pZIL --------------")
	txHash := tradingSuite.depositERC20ToBridge(
		depositingZIL,
		common.HexToAddress(tradingSuite.ZILAddressStr),
		tradingSuite.IncPaymentAddrStr,
	)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)
	// get ZIL remain after depsit
	balZILAfDep := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.ZILAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ZIL balance remain after deposit : ", balZILAfDep)
	// TODO : assert ZIL balance
	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(10 * time.Second)

	issuingRes, err := tradingSuite.callIssuingETHReq(
		tradingSuite.IncZILTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("issuingRes: ", issuingRes)
	time.Sleep(70 * time.Second)

	// check PRV and token balance after issuing
	balPrvAfIssS1, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 1: ", balPrvAfIssS1)
	// TODO assert PRV balance remain
	balpZILAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncZILTokenIDStr)
	fmt.Println("[INFO] pZIL balance after issuing step 1 : ", balpZILAfIssS1)
	// TODO assert pZIL balance issuing

	fmt.Println("------------ step 2: burning pZIL to deposit ZIL to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncZILTokenIDStr,
		burningPZIL,
		pubKeyToAddrStr[2:],
		"createandsendburningfordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(140 * time.Second)

	// check PRV and token balance after burning
	balPrvAfBurnS2, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after burning step 2: ", balPrvAfBurnS2)
	// TODO assert PRV balance remain
	balpZILAfBurnS2, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncZILTokenIDStr)
	fmt.Println("[INFO] pZILbalance after burning step 2 : ", balpZILAfBurnS2)
	// TODO assert pZIL balance issuing

	tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))

	balZILScDepS2 := tradingSuite.getDepositedBalance(
		tradingSuite.ZILAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ZIL balance after deposit on SC at step 2: ", balZILScDepS2)
	//require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPZIL, big.NewInt(1000000000)), balZILScDepS2)

	balEthScS2 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC at step 2 : ", balEthScS2)

	fmt.Println("------------ step 3: execute trade ZIL for ETH through Kyber aggregator --------------")
	tradingSuite.executeWithKyber(
		tradeAmount,
		tradingSuite.ZILAddressStr,
		tradingSuite.EtherAddressStr,
	)
	balEthScTradeS3 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after trade at step 3 : ", balEthScTradeS3)
	// TODO assert ETH balane on SC
	balZILScTradeS3 := tradingSuite.getDepositedBalance(
		tradingSuite.ZILAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ZIL balance on SC after trade at step 3 : ", balZILScTradeS3)
	// TODO assert ZIL balane on SC
	fmt.Println("------------ step 4: withdrawing ETH from SC to pETH on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.EtherAddressStr,
		balEthScTradeS3,
	)
	time.Sleep(15 * time.Second)
	balEthScDepS4 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after withdraw at step 4 : ", balEthScDepS4)
	// TODO assert ETH balane on SC

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(10 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(140 * time.Second)

	balpEthAfIssS4, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncEtherTokenIDStr,
	)
	fmt.Println("[INFO] pETH balance after issuing step 4 : ", balpEthAfIssS4)
	// TODO assert pETH balance issuing
	balPrvAfIssS4, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 4: ", balPrvAfIssS4)
	// TODO assert PRV balance remain

	fmt.Println("------------ step 5: withdrawing pETH from Incognito to ETH --------------")
	withdrawingPETH := big.NewInt(0).Div(balEthScTradeS3, big.NewInt(1000000000))
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

	balETH := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("Ether balance after trade: ", balETH)
	// require.Equal(tradingSuite.T(), withdrawingPETH.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())
	balZIL := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.ZILAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("ZIL balance after trade: ", balZIL)
}

func (tradingSuite *KovanKyberTradingTestSuite) Test5TradeEthithKyber() {
return
	fmt.Println("============ TEST 5 TRADE ETHER Kyber AGGREGATOR ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()

	// get info balance initialization
	balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)

	balpEthInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance initialization : ", balpEthInit)

	balpOMGInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncOMGTokenIDStr)
	fmt.Println("[INFO] pOMG balance initialization : ", balpOMGInit)

	balEthInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance initialization : ", balEthInit)

	balOMGInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.OMGAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] OMG balance initialization : ", balOMGInit)

	balEthScInit := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance initialization on SC : ", balEthScInit)

	balOMGScInit := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance initialization on SC : ", balOMGScInit)

	fmt.Println("------------ STEP 1: porting ETH to pETH --------------")
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
	)
	time.Sleep(15 * time.Second)
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	// get ETH remain after depsit
	balEthAfDep := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain after deposit : ", balEthAfDep)
	// TODO : assert ETH balance

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(10 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(70 * time.Second)
	// check PRV and token balance after issuing
	balPrvAfIssS1, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 1: ", balPrvAfIssS1)
	// TODO assert PRV balance remain
	balpEthAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance after issuing step 1 : ", balpEthAfIssS1)
	// TODO assert pETH balance issuing

	fmt.Println("------------ STEP 2: burning pETH to deposit ZIL to SC --------------")
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
	time.Sleep(70 * time.Second)

	// check PRV and token balance after burning
	balPrvAfBurnS2, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after burning step 2: ", balPrvAfBurnS2)
	// TODO assert PRV balance remain
	balpEthAfBurnS2, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance after burning step 2 : ", balpEthAfBurnS2)
	// TODO assert pETH balance issuing

	tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))
	balEthScDepS2 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance after deposit on SC at step 2: ", balEthScDepS2)
	// TODO assert ETH balane on SC
	balOMGScS2 := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance on SC at step 2 : ", balOMGScS2)

	//require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPETH, big.NewInt(1000000000)), deposited)
	fmt.Println("------------ step 3: execute trade through Kyber aggregator --------------")
	fmt.Println("------------ step 3.1: execute trade ETH for KBN through Kyber aggregator --------------")
	tradingSuite.executeWithKyber(
		tradeAmount,
		tradingSuite.EtherAddressStr,
		tradingSuite.KBNAddressStr,
	)
	balEthScTradeS31 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after trade at step 3.1 : ", balEthScTradeS31)
	// TODO assert ETH balane on SC
	balKBNScTradeS31 := tradingSuite.getDepositedBalance(
		tradingSuite.KBNAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] KBN balance on SC after trade at step 3.1 : ", balKBNScTradeS31)
	// TODO assert OMG balane on SC

	fmt.Println("------------ step 3.2: execute trade KBN for SALT through Kyber aggregator --------------")
	tradeAmount = balKBNScTradeS31
	tradingSuite.executeWithKyber(
		tradeAmount,
		tradingSuite.KBNAddressStr,
		tradingSuite.SALTAddressStr,
	)
	balKBNScTradeS32 := tradingSuite.getDepositedBalance(
		tradingSuite.KBNAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance on SC after trade at step 3.2 : ", balKBNScTradeS32)
	// TODO assert OMG balane on SC

	balSALTScTradeS32 := tradingSuite.getDepositedBalance(
		tradingSuite.SALTAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SALT balance on SC after trade at step 3.2 : ", balSALTScTradeS32)
	// TODO assert ETH balane on SC

	fmt.Println("------------ step 3.2: execute trade SALT for ETH through Kyber aggregator --------------")
	tradeAmount =balSALTScTradeS32
	tradingSuite.executeWithKyber(
		tradeAmount,
		tradingSuite.SALTAddressStr,
		tradingSuite.EtherAddressStr,
	)

	balSALTScTradeS33 := tradingSuite.getDepositedBalance(
		tradingSuite.SALTAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SALT balance on SC after trade at step 3.3 : ", balSALTScTradeS33)
	// TODO assert ETH balane on SC

	balEthScTradeS33 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after trade at step 3.3 : ", balEthScTradeS33)
	// TODO assert ETH balane on SC

	fmt.Println("------------ step 4: withdrawing ETH from SC to pETH on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.EtherAddressStr,
		balEthScTradeS33,
	)
	time.Sleep(15 * time.Second)
	balEthScDepS4 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after withdraw at step 4 : ", balEthScDepS4)
	// TODO assert ETH balane on SC

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(10 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(140 * time.Second)

	balpEthAfIssS4, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncEtherTokenIDStr,
	)
	fmt.Println("[INFO] pETH balance after issuing step 4 : ", balpEthAfIssS4)
	// TODO assert pETH balance issuing
	balPrvAfIssS4, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 4: ", balPrvAfIssS4)
	// TODO assert PRV balance remain

	fmt.Println("------------ step 5: withdrawing pETH from Incognito to ETH --------------")
	withdrawingPETH := big.NewInt(0).Div(balEthScTradeS33, big.NewInt(1000000000))
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

	balETH := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("Ether balance after trade: ", balETH)
	// require.Equal(tradingSuite.T(), withdrawingPETH.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())

}

func (tradingSuite *KovanKyberTradingTestSuite) Test6MultiForKNCWithKyber() {

	fmt.Println("============ TEST 6 MULTI WITH Kyber AGGREGATOR ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
  tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
  AmountDeposit := big.NewInt(int64(tradingSuite.DepositingEther * 2 * params.Ether))
	burningPETH := big.NewInt(0).Div(AmountDeposit, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()

	// get info balance initialization
	balPrvInit, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance initialization : ", balPrvInit)

	balpEthInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance initialization : ", balpEthInit)

	balpKBNInit, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncKBNTokenIDStr)
	fmt.Println("[INFO] pKBN balance initialization : ", balpKBNInit)

	balEthInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance initialization : ", balEthInit)

	balKBNInit := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.KBNAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] KBN balance initialization : ", balKBNInit)

	balEthScInit := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance initialization on SC : ", balEthScInit)

	balKBNScInit := tradingSuite.getDepositedBalance(
		tradingSuite.KBNAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] KBN balance initialization on SC : ", balKBNScInit)

  fmt.Println("------------ STEP 1: porting ETH to pETH --------------")
  fmt.Println("ETH deposit amount : ",tradingSuite.DepositingEther*2)
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther*2,
		tradingSuite.IncPaymentAddrStr,
	)
	time.Sleep(15 * time.Second)
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	// get ETH remain after depsit
	balEthAfDep := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("[INFO] ETH balance remain after deposit : ", balEthAfDep)
	// TODO : assert ETH balance

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(10 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(70 * time.Second)
	// check PRV and token balance after issuing
	balPrvAfIssS1, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 1: ", balPrvAfIssS1)
	// TODO assert PRV balance remain
	balpEthAfIssS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance after issuing step 1 : ", balpEthAfIssS1)
	// TODO assert pETH balance issuing

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
	time.Sleep(70 * time.Second)

	// check PRV and token balance after burning
	balPrvAfBurnS2, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after burning step 2: ", balPrvAfBurnS2)
	// TODO assert PRV balance remain
	balpEthAfBurnS2, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr, tradingSuite.IncEtherTokenIDStr)
	fmt.Println("[INFO] pETH balance after burning step 2 : ", balpEthAfBurnS2)
	// TODO assert pETH balance issuing

	tradingSuite.submitBurnProofForDepositToSC(burningTxID.(string))
	balEthScDepS2 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance after deposit on SC at step 2: ", balEthScDepS2)
	// TODO assert ETH balane on SC
	balKNCScS2 := tradingSuite.getDepositedBalance(
		tradingSuite.KNCAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] KNC balance on SC at step 2 : ", balKNCScS2)

	//require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPETH, big.NewInt(1000000000)), deposited)
	fmt.Println("------------ step 3: execute trade ETH for KBN through Kyber aggregator --------------")
	tradingSuite.executeWithKyber(
		tradeAmount,
		tradingSuite.EtherAddressStr,
		tradingSuite.KBNAddressStr,
	)
	balEthScTradeS3 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after trade at step 3 : ", balEthScTradeS3)
	// TODO assert ETH balane on SC
	balKBNScTradeS3 := tradingSuite.getDepositedBalance(
		tradingSuite.KBNAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] KBN balance on SC after trade at step 3 : ", balKBNScTradeS3)
	// TODO assert KBN balane on SC
	//require.NotEqual(tradingSuite.T(),balKNCScTradeS3,balKBNScS2,"trade failed" )

	fmt.Println("------------ step 4: execute trade ETH and KBN for OMG and SALT and through Kyber aggregator --------------")

	balOMGScTradeS4 := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance on SC before trade at step 4 : ", balOMGScTradeS4)
	// TODO assert ETH balane on SC
	balSALScTradeS4 := tradingSuite.getDepositedBalance(
		tradingSuite.SALTAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SALT balance on SC before trade at step 4 : ", balSALScTradeS4)

	tradingSuite.executeMultiTradeWithKyber(
		[]*big.Int{balKBNScTradeS3, balEthScTradeS3},
		[]string{tradingSuite.KBNAddressStr, tradingSuite.EtherAddressStr},
		[]string{tradingSuite.OMGAddressStr, tradingSuite.SALTAddressStr},
	)
	time.Sleep(15 * time.Second)
	balEthScTradeS4 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after trade at step 4 : ", balEthScTradeS4)
	// TODO assert ETH balane on SC
	balKBNScTradeS4 := tradingSuite.getDepositedBalance(
		tradingSuite.KBNAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] KBN balance on SC after trade at step 4 : ", balKBNScTradeS4)
	// TODO assert KBN balane on SC
	balOMGScTradeS41 := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance on SC after trade at step 4 : ", balOMGScTradeS41)
	// TODO assert ETH balane on SC
	balSALScTradeS41 := tradingSuite.getDepositedBalance(
		tradingSuite.SALTAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SALT balance on SC after trade at step 4 : ", balSALScTradeS41)

	fmt.Println("------------ step 5: execute trade ETH and KBN for OMG and SALT and through Kyber aggregator --------------")

	tradingSuite.executeMultiTradeWithKyber(
		[]*big.Int{balOMGScTradeS41, balSALScTradeS41},
		[]string{tradingSuite.OMGAddressStr, tradingSuite.SALTAddressStr},
		[]string{tradingSuite.KBNAddressStr, tradingSuite.EtherAddressStr},
	)
	time.Sleep(15 * time.Second)
	balEthScTradeS5 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after trade at step 5 : ", balEthScTradeS5)
	// TODO assert ETH balane on SC
	balKBNScTradeS5 := tradingSuite.getDepositedBalance(
		tradingSuite.KBNAddressStr,
		pubKeyToAddrStr,
  )
  
	fmt.Println("[INFO] KBN balance on SC after trade at step 5 : ", balKBNScTradeS5)
	// TODO assert KBN balane on SC
	balOMGScTradeS5 := tradingSuite.getDepositedBalance(
		tradingSuite.OMGAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] OMG balance on SC after trade at step 5 : ", balOMGScTradeS5)
	// TODO assert ETH balane on SC
	balSALScTradeS5 := tradingSuite.getDepositedBalance(
		tradingSuite.SALTAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] SALT balance on SC after trade at step 5 : ", balSALScTradeS5)

	fmt.Println("------------ step 6: withdrawing ETH from SC to pETH on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.EtherAddressStr,
		balEthScTradeS5,
	)
	time.Sleep(15 * time.Second)
	balEthScDepS4 := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("[INFO] ETH balance on SC after withdraw at step 6 : ", balEthScDepS4)
	// TODO assert ETH balane on SC

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(10 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(140 * time.Second)

	balpEthAfIssS4, _ := tradingSuite.getBalanceTokenIncAccount(
		tradingSuite.IncPrivKeyStr,
		tradingSuite.IncEtherTokenIDStr,
	)
	fmt.Println("[INFO] pETH balance after issuing step 6 : ", balpEthAfIssS4)
	// TODO assert pETH balance issuing
	balPrvAfIssS4, _ := tradingSuite.getBalancePrvIncAccount(tradingSuite.IncPrivKeyStr)
	fmt.Println("[INFO] PRV balance after issuing step 6: ", balPrvAfIssS4)
	// TODO assert PRV balance remain

	fmt.Println("------------ step 7: withdrawing pETH from Incognito to ETH --------------")
	withdrawingPETH := big.NewInt(0).Div(balEthScTradeS5, big.NewInt(1000000000))
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

	balETH := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
	)
	fmt.Println("Ether balance after trade: ", balETH)

}
