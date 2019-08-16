package bridge

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

func TestFixedParseBurnInst(t *testing.T) {
	metaType := []byte{7, 2, 1}
	token := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3}
	to := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 5, 6}
	amount := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 7, 8, 9}
	in := &burnInst{
		metaType: big.NewInt(0).SetBytes(metaType),
		token:    common.BytesToAddress(token),
		to:       common.BytesToAddress(to),
		amount:   big.NewInt(0).SetBytes(amount),
	}
	data := []byte{}
	data = append(data, metaType...)
	data = append(data, token[:]...)
	data = append(data, to[:]...)
	data = append(data, amount[:]...)

	p, err := setupFixedCommittee()
	if err != nil {
		t.Error(err)
	}
	resMeta, resToken, resTo, resAmount, err := p.v.ParseBurnInst(nil, data)
	out := &burnInst{
		metaType: resMeta,
		token:    resToken,
		to:       resTo,
		amount:   resAmount,
	}
	if err != nil {
		t.Error(err)
	}
	checkBurnInst(t, in, out)
}

func checkBurnInst(t *testing.T, in, out *burnInst) {
	if in.metaType.Cmp(out.metaType) != 0 {
		t.Error(errors.Errorf("incorrect metaType: expect %x, got %x", out.metaType, in.metaType))
	}
	if !bytes.Equal(in.token[:], out.token[:]) {
		t.Error(errors.Errorf("incorrect token: expect %x, got %x", out.token, in.token))
	}
	if !bytes.Equal(in.to[:], out.to[:]) {
		t.Error(errors.Errorf("incorrect to: expect %x, got %x", out.to, in.to))
	}
	if in.amount.Cmp(out.amount) != 0 {
		t.Error(errors.Errorf("incorrect amount: expect %x, got %x", out.amount, in.amount))
	}
}

type burnInst struct {
	metaType *big.Int
	token    common.Address
	to       common.Address
	amount   *big.Int
}

func TestFixedVaultBurn(t *testing.T) {
	proof := getFixedBurnProof()

	p, err := setupFixedCommittee()
	if err != nil {
		t.Error(err)
	}

	oldBalance, newBalance, err := deposit(p, int64(5e18))
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("deposit to vault: %d -> %d\n", oldBalance, newBalance)

	withdrawer := common.HexToAddress("0xe722D8b71DCC0152D47D2438556a45D3357d631f")
	fmt.Printf("withdrawer init balance: %d\n", p.getBalance(withdrawer))

	tx, err := withdraw(p.v, auth, proof)
	if err != nil {
		t.Error(err)
	}
	p.sim.Commit()
	printReceipt(p.sim, tx)

	fmt.Printf("withdrawer new balance: %d\n", p.getBalance(withdrawer))
}

func setupFixedCommittee() (*Platform, error) {
	numBeaconVals, beaconOld, numBridgeVals, bridgeOld := getFixedCommittee()
	return setup(numBeaconVals, beaconOld, numBridgeVals, bridgeOld)
}

func getFixedCommittee() (*big.Int, []byte, *big.Int, []byte) {
	beaconComm := []string{
		"02a96a04ad76a0034efc8819e93308823ce7a3b76fd694f961ee909124096baf00",
		"0242653de0e9af9dd3725008519157314eb5a845dec2cd646ce9e03f780175b700",
		"028c49fc5f3e001c36095335c53b0b7320f6a1c932424e92c9de344b55e80ddf00",
		"0205aae74cb0128a1863c970cbe87e827e28f92a91c2d4768fdb30a279dd081c00",
	}
	numBeaconVals := big.NewInt(int64(len(beaconComm)))
	beacons := []byte{}
	for _, p := range beaconComm {
		pk, _ := hex.DecodeString(p)
		beacons = append(beacons, pk...)
	}

	bridgeComm := []string{
		"0253d262c2b6a55606ff9d32e195231ec57e4d23a6efd1c02143a58fd0c2591d01",
		"02dee56cbbde5ef6d03a9e69bf3784ae4a8460d0058a6082eee4be2ed5c4fd3301",
		"02ec388db662801da0fe3c41f39085369ed4df71d42ec96924012243dc9c67d201",
		"039cc81f72a88a7436eb74bf10c7693af165324ba4d15baeb4e8d2f1c2ce25a101",
	}
	numBridgeVals := big.NewInt(int64(len(bridgeComm)))
	bridges := []byte{}
	for _, p := range bridgeComm {
		pk, _ := hex.DecodeString(p)
		bridges = append(bridges, pk...)
	}
	return numBeaconVals, beacons, numBridgeVals, bridges
}

func getFixedBurnProof() *decodedProof {
	proofMarshalled := `{"Instruction":"NzIxAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADnIti3HcwBUtR9JDhVakXTNX1jHwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAOjUpRAANDhfjaiwxXQ5i1beSKxWlB1TzP3+nP8eBsPOGRycyIcAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==","BeaconHeight":90,"BridgeHeight":86,"BeaconInstPath":[[127,145,240,157,94,42,115,23,1,92,233,155,174,164,6,134,146,159,39,127,48,214,128,183,173,208,206,50,145,33,234,192],[107,7,138,64,74,51,162,247,79,71,175,231,172,143,192,26,126,6,198,136,51,9,51,151,45,103,226,12,155,162,89,95],[175,198,81,148,19,191,203,195,121,37,156,114,205,37,95,27,106,225,198,194,157,207,67,86,215,135,228,157,186,58,52,94],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]],"BeaconInstPathIsLeft":[false,false,false,false,false,false,false,false],"BeaconInstPathLen":3,"BeaconInstRoot":[182,177,199,187,216,156,248,250,179,99,27,99,38,145,186,64,23,134,167,124,69,61,44,219,136,169,13,218,214,99,185,160],"BeaconBlkData":[166,5,103,245,35,148,218,105,162,15,59,154,128,175,144,30,64,28,172,214,198,17,81,111,72,199,255,159,46,80,190,150],"BeaconSignerSig":108653766851384127000058368519586331265508118354654876655603893380561160074650,"BeaconNumR":4,"BeaconXs":[76628189482987364869146169160110220061660726563679851006904348516982792630016,30031526630203414515762634695278890735218732611664146584180656826754956375808,63454520445189848446618105030313607321390685434043153264799945723547105484544,2563524614833302616300629989486978541294361866987495309309020333370860116992,0,0,0,0],"BeaconYs":[9636992064743471482255653601005295963147419362613637810560330997782348639040,36629607870677315039431568064966848933680928918874345555215893802439865437646,27998033041715606256724099267641575019608365624764129368431228564081554798658,21266207138192514800206514642409627249467085679781700776536289494899814007326,0,0,0,0],"BeaconRIdxs":[0,1,2,3,0,0,0,0],"BeaconNumSig":3,"BeaconSigIdxs":[0,1,2,0,0,0,0,0],"BeaconRp":"An3MmviR93OXrE+nCYQWWnkw6r3Xhp0QYfFCnwpR51cU","BeaconRpx":56900612444499346098380332036540696803164364369164037985635157644787677878036,"BeaconRpy":18514822436875343477197831737170591568097261912543075746770606015982726635642,"BeaconR":"A01Ishk6V8Y2A0/IPZpEWERs0pqP+Za6bA7FSA1Eiqmi","BridgeInstPath":[[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]],"BridgeInstPathIsLeft":[false,false,false,false,false,false,false,false],"BridgeInstPathLen":0,"BridgeInstRoot":[126,106,9,14,53,221,12,29,140,88,180,61,168,214,103,126,139,62,61,154,231,254,161,140,104,196,112,142,111,226,4,59],"BridgeBlkData":[111,71,179,205,55,92,27,19,249,236,151,129,89,130,173,111,36,68,63,114,139,213,62,40,109,183,170,182,88,109,26,243],"BridgeSignerSig":23834169346739917051200768753647431329210340859271718824561737031758316402076,"BridgeNumR":4,"BridgeXs":[37913685936621062233942649412298552394934022531604109971145288653189322185985,100818810816853263089761344527239819816697567674074098611568731254340756255489,106845753764596956410597914868945361934191890042000742122478033889400013181441,70914390837265325671900757599180068068737973558198010961381520433542764798209,0,0,0,0],"BridgeYs":[61220416244066161862675065339680945577902515429105185136860145906020910676056,63407347531535486888692009740089405771718500233453842136668634622269019724914,113992982342209111422350952149806883524194480643239209681083181558399009244220,102255150656278975065826750760865118653834936497361527525164792990929259379201,0,0,0,0],"BridgeRIdxs":[0,1,2,3,0,0,0,0],"BridgeNumSig":3,"BridgeSigIdxs":[0,1,2,0,0,0,0,0],"BridgeRp":"AjiNljrdELvBQ+woA1AOsMJnMR2xU+itNetUNdSC32kT","BridgeRpx":25579681805706495406223410167195188483903432596929165811882052828964694485267,"BridgeRpy":94209519572664110476851563406995380813113054628712847488879332788808650222052,"BridgeR":"Ak3g+zUmE3/tk3SOi3m3UXTnm0+fkN4FPWHdkqkWaYO0"}`
	proof := &decodedProof{}
	json.Unmarshal([]byte(proofMarshalled), proof)
	return proof
}

func toBytes32BigEndian(b []byte) []byte {
	a := [32]byte{}
	copy(a[32-len(b):], b)
	return a[:]
}
