package main

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
	token := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3}
	to := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 5, 6}
	amount := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 7, 8, 9}
	in := &burnInst{
		meta:   72,
		shard:  1,
		token:  common.BytesToAddress(token),
		to:     common.BytesToAddress(to),
		amount: big.NewInt(0).SetBytes(amount),
	}
	data := []byte{in.meta, in.shard}
	data = append(data, token[:]...)
	data = append(data, to[:]...)
	data = append(data, amount[:]...)

	p, _, err := setupFixedCommittee()
	if err != nil {
		t.Error(err)
	}
	resMeta, resShard, resToken, resTo, resAmount, err := p.v.ParseBurnInst(nil, data)
	out := &burnInst{
		meta:   resMeta,
		shard:  resShard,
		token:  resToken,
		to:     resTo,
		amount: resAmount,
	}
	if err != nil {
		t.Error(err)
	}
	checkBurnInst(t, in, out)
}

func checkBurnInst(t *testing.T, in, out *burnInst) {
	if in.meta != out.meta {
		t.Error(errors.Errorf("incorrect meta: expect %x, got %x", out.meta, in.meta))
	}
	if in.shard != out.shard {
		t.Error(errors.Errorf("incorrect shard: expect %x, got %x", out.shard, in.shard))
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
	meta   uint8
	shard  uint8
	token  common.Address
	to     common.Address
	amount *big.Int
}

func TestFixedVaultBurn(t *testing.T) {
	proof := getFixedBurnProof()

	p, _, err := setupFixedCommittee()
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

func setupFixedCommittee() (*Platform, *committees, error) {
	c := getFixedCommittee()
	p, err := setup(c.beacons, c.bridges)
	return p, c, err
}

func getFixedCommittee() *committees {
	beaconCommPrivs := []string{
		"5a417f54357fff96fe4c2a9cafd322ed72b52bf046beb69a9730a26181088489",
		"b9cd32581922f447acb1cfd148069fc40cbbce1e8badb84c4b509486e6f713ce",
		"22e23970b853407e16ccb174443f27c37dbbea05729aba546ee649e0aef2d9cb",
		"4d16dadc89656fbda140e2fe467631ddac3ed9cc326cef3a8f1b1bd5f3cfd155",
	}
	beaconComm := []string{
		"0xA5301a0d25103967bf0e29db1576cba3408fD9bB",
		"0x9BC0faE7BB432828759B6e391e0cC99995057791",
		"0x6cbc2937FEe477bbda360A842EeEbF92c2FAb613",
		"0xcabF3DB93eB48a61d41486AcC9281B6240411403",
	}
	beacons := make([]common.Address, len(beaconComm))
	beaconPrivs := make([][]byte, len(beaconCommPrivs))
	for i, p := range beaconComm {
		beacons[i] = common.HexToAddress(p)
		priv, _ := hex.DecodeString(beaconCommPrivs[i])
		beaconPrivs[i] = priv
	}

	bridgeComm := []string{
		"0x3c78124783E8e39D1E084FdDD0E097334ba2D945",
		"0x76E34d8a527961286E55532620Af5b84F3C6538F",
		"0x68686dB6874588D2404155D00A73F82a50FDd190",
		"0x1533ac4d2922C150551f2F5dc2b0c1eDE382b890",
	}
	bridgeCommPrivs := []string{
		"3560e649ce326a2eb9fbb59fba4b29e10fb064627f61487aecc8b92afbb127dd",
		"b71af1a7e2ca74400187cbf2333ab1f20e9b39517347fb655ffa309d1b51b2b0",
		"07f91f98513c203103f8d44683ce47920d1aea0eaf1cb86a373be835374d1490",
		"7412e24d4ac1796866c44a0d5b966f8db1c3022bba8afd370a09dc49a14efeb4",
	}

	bridges := make([]common.Address, len(bridgeComm))
	bridgePrivs := make([][]byte, len(bridgeCommPrivs))
	for i, p := range bridgeComm {
		bridges[i] = common.HexToAddress(p)
		priv, _ := hex.DecodeString(bridgeCommPrivs[i])
		bridgePrivs[i] = priv
	}
	return &committees{
		beacons:     beacons,
		beaconPrivs: beaconPrivs,
		bridges:     bridges,
		bridgePrivs: bridgePrivs,
	}
}

type committees struct {
	beacons     []common.Address
	bridges     []common.Address
	beaconPrivs [][]byte
	bridgePrivs [][]byte
}

func getFixedBurnProof() *decodedProof {
	proofMarshalled := `{"Instruction":"SAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAOci2LcdzAFS1H0kOFVqRdM1fWMfAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA6NSlEAAQph4uU2jzYJA/wwowaQNkV1YDSZxF3b7zhinToad2CgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","Heights":[13,15],"InstPaths":[[[85,90,80,85,102,77,35,181,79,15,218,226,21,248,182,95,46,161,151,187,156,132,212,157,64,83,229,28,33,225,60,39],[90,120,63,154,71,105,58,60,79,100,14,227,126,162,217,227,120,234,73,47,138,250,88,120,245,244,79,86,210,74,95,107]],[]],"InstPathIsLefts":[[false,false],[]],"InstRoots":[[232,220,37,58,255,80,189,228,5,6,217,133,81,89,207,211,88,46,172,168,207,11,240,31,39,121,213,208,161,87,191,205],[189,164,48,17,53,82,33,194,160,26,125,127,160,224,216,173,239,49,214,141,169,214,192,76,85,66,30,230,107,198,183,59]],"BlkData":[[85,105,22,139,166,247,192,177,35,127,254,184,33,212,88,184,142,32,141,79,0,92,62,215,250,238,97,155,250,103,212,226],[43,69,182,160,171,28,79,48,143,171,86,32,244,96,9,226,26,232,37,254,92,251,3,56,247,81,147,139,97,238,181,249]],"SigIdxs":[[0,2,3],[0,1,2,3]],"SigVs":["HBsb","GxwcHA=="],"SigRs":[[[64,129,127,189,84,17,9,178,133,57,35,78,163,181,179,201,225,45,218,234,194,20,180,7,17,207,92,110,26,73,205,24],[195,54,155,40,185,122,123,141,83,10,38,226,70,217,107,128,172,33,9,9,218,212,254,215,89,20,84,138,200,144,225,255],[46,36,148,21,139,126,174,64,59,164,229,133,213,79,228,67,35,172,234,237,175,125,35,138,2,114,186,151,184,76,242,134]],[[78,44,39,129,74,56,101,155,93,226,145,64,149,247,158,112,186,70,242,82,115,255,142,140,101,38,183,206,197,186,152,138],[119,42,151,249,112,98,99,187,116,121,228,46,126,106,38,166,45,217,156,48,83,103,34,168,246,123,178,82,208,197,9,8],[190,186,249,188,241,169,26,152,54,224,240,95,81,179,223,224,224,81,209,193,57,97,3,165,197,40,229,2,208,20,113,131],[80,12,25,116,195,24,37,40,247,144,182,31,57,127,236,239,190,90,236,59,218,97,205,16,88,102,235,156,220,144,198,234]]],"SigSs":[[[118,52,192,85,239,241,154,66,156,98,190,204,10,253,210,148,162,165,195,201,212,206,135,182,178,209,193,172,240,180,99,110],[61,45,9,108,235,184,112,73,21,154,9,116,162,95,175,207,11,22,215,160,46,209,83,59,118,230,83,209,241,17,168,127],[59,15,150,127,223,9,234,160,9,170,195,109,79,118,137,45,250,156,249,145,226,175,235,60,65,86,182,85,224,162,235,145]],[[101,235,22,155,225,149,55,249,33,205,164,198,58,228,75,33,48,196,171,76,14,181,90,194,239,204,40,159,26,179,46,3],[121,162,14,137,255,11,150,184,16,237,220,66,226,26,240,69,154,30,136,46,68,254,217,190,122,100,41,254,221,122,158,66],[84,26,159,168,124,109,132,10,52,124,177,99,8,53,85,52,194,213,247,52,147,255,17,146,245,114,252,139,176,126,82,203],[39,215,56,10,60,107,131,198,45,193,242,191,111,195,106,238,112,77,177,198,3,88,157,118,76,170,148,65,187,106,118,172]]]}`
	proof := &decodedProof{}
	json.Unmarshal([]byte(proofMarshalled), proof)

	// fmt.Printf("proof: %+v\n", proof)
	// s := []string{"0x4f8daa779f713e539bbe8942e390325fc360a4811b0a568fbcfd9c2547f612d9", "0x48c8fb9529c8de1d9097366747c555847f29e73838c36c33eec1fcc266441fc1", "0x1e81121869b2d39c39e0d5bce27a233e7488c254c4119b8455241d89b17090dc", "0x1c8060da0d3650501b6c7c7d4940d224dbc64eab4d8beae4b6358f4b9b538e2a", "0x06a086905048b493cbefc4d64fe1ee468836ffec67a81fcd205a92c42667830b", "0x5fbcab3fbe8e8e5a978dcfbae54f68917f8056e802cebc2479846413ec054e97", "0x280704f0e0f70f1aa43ee9e0fc21fc09edafa00ec29cebff696c504e7fa7c43b", "0x7fdc885e553ae521b0879cb86c3891e5cdaf4bad6cd1fd29f2482fa282f93978", "0x6ee340fa29ed5b41285297a7bc0451881a4aad138fa6fa43a7b507d58561106d", "0x7c5b3abc8224dc35982a1f9dfd34d9a18862aad8f54ed2d905b6e81c7fbcf8a7"}
	// r := []string{"0x012305d4c970ee1872d87b9c8de53bfa698782222bb318861690f9863f4c5a2a", "0x03e6a2f1af4613c257eea0b11977d482f22969335504ed55ff07e3225e677fcc", "0x86e2b5df59171a87b94e38af69474eb97c3a741fc6e3d488e3e27d63da44612f", "0xd7deea8a087d4344bcb39f897de886be7020a22518b177b3541e2cb81d10bf80", "0x778bf6835438261ba8485774aec7c44b6239c30a9e68f74aa417af673df31d35", "0x98f4a9d71f60ea86321f84e4ca6cce1ef5827e2ae8e22ad1944d70125820a626", "0xcba676c98e3489e27ee9d37d4ec39239f60523390f30c30ad6b4feb60eacc5e3", "0xbf5e08d182998ec31f19a8d3744be78d5025b7fb33f87b4e71ec1d797cf92ddc", "0xd6f28bc51138223f0e6620f42e4fba1159130d4dfa136e7ab35ff4081e4cee32", "0x62e9fd38a63b0180a41f777a30717b198782dc8e43ebec3281a3f0e5e6aca069"}
	// v := []int{27, 28, 28, 27, 27, 27, 27, 28, 28, 28}
	// for i, a := range s {
	// 	d, _ := hexutil.Decode(a)
	// 	copy(proof.BeaconSigSs[i][:], toBytes32BigEndian(d))
	// 	copy(proof.BridgeSigSs[i][:], toBytes32BigEndian(d))
	// }
	// for i, a := range r {
	// 	d, _ := hexutil.Decode(a)
	// 	copy(proof.BeaconSigRs[i][:], toBytes32BigEndian(d))
	// 	copy(proof.BridgeSigRs[i][:], toBytes32BigEndian(d))
	// }
	// for i, a := range v {
	// 	proof.BeaconSigVs[i] = big.NewInt(int64(a))
	// 	proof.BridgeSigVs[i] = big.NewInt(int64(a))
	// }
	// fmt.Printf("new proof: %+v\n", proof)
	// pp, err := json.Marshal(proof)
	// fmt.Printf("marshalled: %s\n%v\n", string(pp), err)
	return proof
}

func toBytes32BigEndian(b []byte) []byte {
	a := copyToBytes32(b)
	return a[:]
}

func copyToBytes32(b []byte) [32]byte {
	a := [32]byte{}
	copy(a[32-len(b):], b)
	return a
}
