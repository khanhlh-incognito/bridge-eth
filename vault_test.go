package bridge

import (
	"bytes"
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

func getFixedCommittee() (*big.Int, [comm_size]common.Address, *big.Int, [comm_size]common.Address) {
	beaconComm := []string{
		"0xcb859a5fC20EEeCc4Cec191d8CCe5e31a2CC1dAF",
		"0x6294A3d1caE192f06dEe33152559531D447A076f",
		"0xA815f542096c8De2ECe5aB18d4cf9D2aBc5202EC",
		"0x4B895A89606aD73d2Fd7b887583858d6f2Cd229c",
		"0xdd0523326fD818a16783D392324003D1df163758",
		"0xf5b0A7D1270642e55a92A99D1AF9bb2B59982C71",
		"0xbEBE7795d8297c4A59542a81541878e2dBA95253",
		"0x894a0bEbb56cE3099A34f26b259D4038AE0E6088",
		"0x6D6abB339E215a92c190f045D88E8aF79A32Dd16",
		"0x81a3B54a6216585C6A262AAF2c4340Ac53F7c10c",
	}
	numBeaconVals := big.NewInt(int64(len(beaconComm)))
	beacons := [comm_size]common.Address{}
	for i, p := range beaconComm {
		beacons[i] = common.HexToAddress(p)
	}

	bridgeComm := []string{
		"0xcb859a5fC20EEeCc4Cec191d8CCe5e31a2CC1dAF",
		"0x6294A3d1caE192f06dEe33152559531D447A076f",
		"0xA815f542096c8De2ECe5aB18d4cf9D2aBc5202EC",
		"0x4B895A89606aD73d2Fd7b887583858d6f2Cd229c",
		"0xdd0523326fD818a16783D392324003D1df163758",
		"0xf5b0A7D1270642e55a92A99D1AF9bb2B59982C71",
		"0xbEBE7795d8297c4A59542a81541878e2dBA95253",
		"0x894a0bEbb56cE3099A34f26b259D4038AE0E6088",
		"0x6D6abB339E215a92c190f045D88E8aF79A32Dd16",
		"0x81a3B54a6216585C6A262AAF2c4340Ac53F7c10c",
	}
	numBridgeVals := big.NewInt(int64(len(bridgeComm)))
	bridges := [comm_size]common.Address{}
	for i, p := range bridgeComm {
		bridges[i] = common.HexToAddress(p)
	}
	return numBeaconVals, beacons, numBridgeVals, bridges
}

func getFixedBurnProof() *decodedProof {
	proofMarshalled := `{"Instruction":"NzIxAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADnIti3HcwBUtR9JDhVakXTNX1jHwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAOjUpRAANDhfjaiwxXQ5i1beSKxWlB1TzP3+nP8eBsPOGRycyIcAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==","BeaconHeight":90,"BridgeHeight":86,"BeaconInstPath":[[127,145,240,157,94,42,115,23,1,92,233,155,174,164,6,134,146,159,39,127,48,214,128,183,173,208,206,50,145,33,234,192],[107,7,138,64,74,51,162,247,79,71,175,231,172,143,192,26,126,6,198,136,51,9,51,151,45,103,226,12,155,162,89,95],[175,198,81,148,19,191,203,195,121,37,156,114,205,37,95,27,106,225,198,194,157,207,67,86,215,135,228,157,186,58,52,94],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]],"BeaconInstPathIsLeft":[false,false,false,false,false,false,false,false],"BeaconInstPathLen":3,"BeaconInstRoot":[182,177,199,187,216,156,248,250,179,99,27,99,38,145,186,64,23,134,167,124,69,61,44,219,136,169,13,218,214,99,185,160],"BeaconBlkData":[166,5,103,245,35,148,218,105,162,15,59,154,128,175,144,30,64,28,172,214,198,17,81,111,72,199,255,159,46,80,190,150],"BeaconNumSig":10,"BeaconSigVs":[27,28,28,27,27,27,27,28,28,28],"BeaconSigRs":[[1,35,5,212,201,112,238,24,114,216,123,156,141,229,59,250,105,135,130,34,43,179,24,134,22,144,249,134,63,76,90,42],[3,230,162,241,175,70,19,194,87,238,160,177,25,119,212,130,242,41,105,51,85,4,237,85,255,7,227,34,94,103,127,204],[134,226,181,223,89,23,26,135,185,78,56,175,105,71,78,185,124,58,116,31,198,227,212,136,227,226,125,99,218,68,97,47],[215,222,234,138,8,125,67,68,188,179,159,137,125,232,134,190,112,32,162,37,24,177,119,179,84,30,44,184,29,16,191,128],[119,139,246,131,84,56,38,27,168,72,87,116,174,199,196,75,98,57,195,10,158,104,247,74,164,23,175,103,61,243,29,53],[152,244,169,215,31,96,234,134,50,31,132,228,202,108,206,30,245,130,126,42,232,226,42,209,148,77,112,18,88,32,166,38],[203,166,118,201,142,52,137,226,126,233,211,125,78,195,146,57,246,5,35,57,15,48,195,10,214,180,254,182,14,172,197,227],[191,94,8,209,130,153,142,195,31,25,168,211,116,75,231,141,80,37,183,251,51,248,123,78,113,236,29,121,124,249,45,220],[214,242,139,197,17,56,34,63,14,102,32,244,46,79,186,17,89,19,13,77,250,19,110,122,179,95,244,8,30,76,238,50],[98,233,253,56,166,59,1,128,164,31,119,122,48,113,123,25,135,130,220,142,67,235,236,50,129,163,240,229,230,172,160,105]],"BeaconSigSs":[[79,141,170,119,159,113,62,83,155,190,137,66,227,144,50,95,195,96,164,129,27,10,86,143,188,253,156,37,71,246,18,217],[72,200,251,149,41,200,222,29,144,151,54,103,71,197,85,132,127,41,231,56,56,195,108,51,238,193,252,194,102,68,31,193],[30,129,18,24,105,178,211,156,57,224,213,188,226,122,35,62,116,136,194,84,196,17,155,132,85,36,29,137,177,112,144,220],[28,128,96,218,13,54,80,80,27,108,124,125,73,64,210,36,219,198,78,171,77,139,234,228,182,53,143,75,155,83,142,42],[6,160,134,144,80,72,180,147,203,239,196,214,79,225,238,70,136,54,255,236,103,168,31,205,32,90,146,196,38,103,131,11],[95,188,171,63,190,142,142,90,151,141,207,186,229,79,104,145,127,128,86,232,2,206,188,36,121,132,100,19,236,5,78,151],[40,7,4,240,224,247,15,26,164,62,233,224,252,33,252,9,237,175,160,14,194,156,235,255,105,108,80,78,127,167,196,59],[127,220,136,94,85,58,229,33,176,135,156,184,108,56,145,229,205,175,75,173,108,209,253,41,242,72,47,162,130,249,57,120],[110,227,64,250,41,237,91,65,40,82,151,167,188,4,81,136,26,74,173,19,143,166,250,67,167,181,7,213,133,97,16,109],[124,91,58,188,130,36,220,53,152,42,31,157,253,52,217,161,136,98,170,216,245,78,210,217,5,182,232,28,127,188,248,167]],"BeaconSigIdxs":[0,1,2,3,4,5,6,7,8,9],"BridgeInstPath":[[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]],"BridgeInstPathIsLeft":[false,false,false,false,false,false,false,false],"BridgeInstPathLen":0,"BridgeInstRoot":[126,106,9,14,53,221,12,29,140,88,180,61,168,214,103,126,139,62,61,154,231,254,161,140,104,196,112,142,111,226,4,59],"BridgeBlkData":[111,71,179,205,55,92,27,19,249,236,151,129,89,130,173,111,36,68,63,114,139,213,62,40,109,183,170,182,88,109,26,243],"BridgeNumSig":10,"BridgeSigVs":[27,28,28,27,27,27,27,28,28,28],"BridgeSigRs":[[1,35,5,212,201,112,238,24,114,216,123,156,141,229,59,250,105,135,130,34,43,179,24,134,22,144,249,134,63,76,90,42],[3,230,162,241,175,70,19,194,87,238,160,177,25,119,212,130,242,41,105,51,85,4,237,85,255,7,227,34,94,103,127,204],[134,226,181,223,89,23,26,135,185,78,56,175,105,71,78,185,124,58,116,31,198,227,212,136,227,226,125,99,218,68,97,47],[215,222,234,138,8,125,67,68,188,179,159,137,125,232,134,190,112,32,162,37,24,177,119,179,84,30,44,184,29,16,191,128],[119,139,246,131,84,56,38,27,168,72,87,116,174,199,196,75,98,57,195,10,158,104,247,74,164,23,175,103,61,243,29,53],[152,244,169,215,31,96,234,134,50,31,132,228,202,108,206,30,245,130,126,42,232,226,42,209,148,77,112,18,88,32,166,38],[203,166,118,201,142,52,137,226,126,233,211,125,78,195,146,57,246,5,35,57,15,48,195,10,214,180,254,182,14,172,197,227],[191,94,8,209,130,153,142,195,31,25,168,211,116,75,231,141,80,37,183,251,51,248,123,78,113,236,29,121,124,249,45,220],[214,242,139,197,17,56,34,63,14,102,32,244,46,79,186,17,89,19,13,77,250,19,110,122,179,95,244,8,30,76,238,50],[98,233,253,56,166,59,1,128,164,31,119,122,48,113,123,25,135,130,220,142,67,235,236,50,129,163,240,229,230,172,160,105]],"BridgeSigSs":[[79,141,170,119,159,113,62,83,155,190,137,66,227,144,50,95,195,96,164,129,27,10,86,143,188,253,156,37,71,246,18,217],[72,200,251,149,41,200,222,29,144,151,54,103,71,197,85,132,127,41,231,56,56,195,108,51,238,193,252,194,102,68,31,193],[30,129,18,24,105,178,211,156,57,224,213,188,226,122,35,62,116,136,194,84,196,17,155,132,85,36,29,137,177,112,144,220],[28,128,96,218,13,54,80,80,27,108,124,125,73,64,210,36,219,198,78,171,77,139,234,228,182,53,143,75,155,83,142,42],[6,160,134,144,80,72,180,147,203,239,196,214,79,225,238,70,136,54,255,236,103,168,31,205,32,90,146,196,38,103,131,11],[95,188,171,63,190,142,142,90,151,141,207,186,229,79,104,145,127,128,86,232,2,206,188,36,121,132,100,19,236,5,78,151],[40,7,4,240,224,247,15,26,164,62,233,224,252,33,252,9,237,175,160,14,194,156,235,255,105,108,80,78,127,167,196,59],[127,220,136,94,85,58,229,33,176,135,156,184,108,56,145,229,205,175,75,173,108,209,253,41,242,72,47,162,130,249,57,120],[110,227,64,250,41,237,91,65,40,82,151,167,188,4,81,136,26,74,173,19,143,166,250,67,167,181,7,213,133,97,16,109],[124,91,58,188,130,36,220,53,152,42,31,157,253,52,217,161,136,98,170,216,245,78,210,217,5,182,232,28,127,188,248,167]],"BridgeSigIdxs":[0,1,2,3,4,5,6,7,8,9]}`
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
