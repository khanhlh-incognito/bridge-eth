package main

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"testing"

	ec "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/incognitochain/bridge-eth/common"
	"github.com/incognitochain/bridge-eth/consensus/bridgesig"
)

// {
// 	"Hash": "951a678d297c56971163532b0cc11118b844ec7326f227ebc21f0171a6097635",
// 	"Height": 45,
// 	"ValidationData": "{\"ProducerBLSSig\":\"FVj54LstEsxs7W7jOceiezCSBPJ/F2NGowSqaPFpLzg=\",\"ProducerBriSig\":null,\"ValidatiorsIdx\":[0,1,2,3],\"AggSig\":\"KwrsVXog3ZN2jkjG4MMggsTDnrX3PBRE4dlv67++KlY=\",\"BridgeSig\":[\"GOAcW/psvsZspExMqrtWYOVj+JttJVZmixH9j0mbRXhW/UmmlZvPk4fWRtQCXkZzrAgaIxoMvzj0g0aE2amVQwE=\",\"udWBfZvFO3AUMmMhvl5/2WpJ7hDmWxuxbPskRC1hwPpb1y+mcnzKye45854mfgCyzCY1mzCR7qveF7bVYHRNbwE=\",\"nxfEKeJslQWcEKhR20w8xt/T4JLB3E+ECpkcx+fjFYs4m5q6bABtiKbfFiY9yOAjma42EOtFyKw0fF1dedkTNwA=\",\"POaLQ2r5VjGlH5EqsT5cyr4RnSQbUt+wfRxNroQ21qZ1Iix/eJkDeW3W/rROmo66Os5WBef/huTX+qGx9MN0ngE=\"]}",
// 	"ConsensusType": "bls",
// 	"Version": 1,
// 	"Epoch": 3,
// 	"Round": 1,
// 	"Time": 1566966083,
// 	"PreviousBlockHash": "e7f9031fb8cab29ca19d11011799327568e5cf280cf08979b05cd1c88d8ce9d0",
// 	"NextBlockHash": "",
// 	"Instructions": [
// 		[
// 			"72",
// 			"1",
// 			"1111111111111111111116onJqC",
// 			"e722D8b71DCC0152D47D2438556a45D3357d631f",
// 			"13xteieoTFDLUs",
// 			"7e4a9d8da4fac015af7738a45e56c644120b9473d72508e47a9e057285bf9785",
// 			"1111111111111111111111111111111116dTTiD",
// 			"169tMnQY"
// 		],
// 		[
// 			"37",
// 			"-2",
// 			"{\"ShardID\":0,\"TxsFee\":{},\"ShardBlockHeight\":33}"
// 		],
// 		[
// 			"37",
// 			"-2",
// 			"{\"ShardID\":0,\"TxsFee\":{\"0000000000000000000000000000000000000000000000000000000000000000\":0},\"ShardBlockHeight\":34}"
// 		]
// 	],
// 	"Size": 2293
// },

func TestFixedVerifySig(t *testing.T) {
	p, err := setupWithLocalCommittee()
	// p, err := setupFixedCommittee()
	if err != nil {
		t.Error(err)
	}

	validationData := "{\"ProducerBLSSig\":\"FVj54LstEsxs7W7jOceiezCSBPJ/F2NGowSqaPFpLzg=\",\"ProducerBriSig\":null,\"ValidatiorsIdx\":[0,1,2,3],\"AggSig\":\"KwrsVXog3ZN2jkjG4MMggsTDnrX3PBRE4dlv67++KlY=\",\"BridgeSig\":[\"GOAcW/psvsZspExMqrtWYOVj+JttJVZmixH9j0mbRXhW/UmmlZvPk4fWRtQCXkZzrAgaIxoMvzj0g0aE2amVQwE=\",\"udWBfZvFO3AUMmMhvl5/2WpJ7hDmWxuxbPskRC1hwPpb1y+mcnzKye45854mfgCyzCY1mzCR7qveF7bVYHRNbwE=\",\"nxfEKeJslQWcEKhR20w8xt/T4JLB3E+ECpkcx+fjFYs4m5q6bABtiKbfFiY9yOAjma42EOtFyKw0fF1dedkTNwA=\",\"POaLQ2r5VjGlH5EqsT5cyr4RnSQbUt+wfRxNroQ21qZ1Iix/eJkDeW3W/rROmo66Os5WBef/huTX+qGx9MN0ngE=\"]}"
	d, _ := DecodeValidationData(validationData)

	hash, _ := (common.Hash{}).NewHashFromStr("951a678d297c56971163532b0cc11118b844ec7326f227ebc21f0171a6097635")

	beacons := []string{
		"AhixK1RFobDelCNigs4Hy+KBG9NC/dyV1fu9Hen2D++r",
		"A52NI9M1oXn7gphEPiJarspTXOI1WE4XiU0a/rCcxY9p",
		"AhIFfgHHQiKma4aznJ33g1jwzi5Sh/vn2OWhi11E80yp",
		"A57kXi9TfbU0v9dekOCMjVRF4yO322mod4ili3YP3EGO",
	}
	committee := []ec.Address{}
	for _, b := range beacons {
		pubkey, _ := base64.StdEncoding.DecodeString(b)
		pk, _ := crypto.DecompressPubkey(pubkey)
		addr := crypto.PubkeyToAddress(*pk)
		fmt.Printf("%x\n", addr[:])
		committee = append(committee, addr)
	}

	vs := []byte{}
	rs := [][32]byte{}
	ss := [][32]byte{}
	for _, sig := range d.BridgeSig {
		v, r, s, err := bridgesig.DecodeECDSASig(sig)
		if err != nil {
			t.Fatal(err)
		}
		vs = append(vs, v)
		rs = append(rs, toByte32(r))
		ss = append(ss, toByte32(s))
	}
	res, err := p.inc.VerifySig(nil, committee, toByte32(crypto.Keccak256Hash(hash.GetBytes()).Bytes()), vs, rs, ss)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("verify:", res)
}

type ValidationData struct {
	ProducerBLSSig []byte
	ProducerBriSig []byte
	ValidatiorsIdx []int
	AggSig         []byte
	BridgeSig      [][]byte
}

func DecodeValidationData(data string) (*ValidationData, error) {
	var valData ValidationData
	err := json.Unmarshal([]byte(data), &valData)
	if err != nil {
		return nil, err
	}
	return &valData, nil
}

func TestFixedSwapBridge(t *testing.T) {
	proof := getFixedSwapProof()

	p, err := setupWithLocalCommittee()
	// p, err := setupFixedCommittee()
	if err != nil {
		t.Error(err)
	}

	tx, err := SwapBridge(p.inc, auth, proof)
	if err != nil {
		t.Error(err)
	}
	p.sim.Commit()
	printReceipt(p.sim, tx)
}

func TestFixedSwapBeacon(t *testing.T) {
	proof := getFixedSwapProof()

	p, err := setupFixedCommittee()
	if err != nil {
		t.Error(err)
	}

	tx, err := SwapBeacon(p.inc, auth, proof)
	if err != nil {
		t.Error(err)
	}
	p.sim.Commit()
	printReceipt(p.sim, tx)
}

func TestExtractMetaFromInstruction(t *testing.T) {
	p, _ := setupFixedCommittee()
	addrs := []string{
		"834f98e1b7324450b798359c9febba74fb1fd888",
		"1250ba2c592ac5d883a0b20112022f541898e65b",
		"2464c00eab37be5a679d6e5f7c8f87864b03bfce",
		"6d4850ab610be9849566c09da24b37c5cfa93e50",
	}
	testCases := []struct {
		desc    string
		inst    []byte
		meta    int
		height  int
		numVals int
		err     bool
	}{
		{
			desc:    "Extract swap beacon instruction",
			inst:    buildDecodedSwapConfirmInst(70, 1, 123, addrs),
			meta:    3616817,
			height:  123,
			numVals: len(addrs),
		},
		{
			desc:    "Extract swap bridge instruction",
			inst:    buildDecodedSwapConfirmInst(71, 1, 19827312, addrs[:2]),
			meta:    3617073,
			height:  19827312,
			numVals: 2,
		},
		{
			desc: "Instruction too short",
			inst: make([]byte, 66),
			err:  true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			meta, height, numVals, err := p.inc.ExtractMetaFromInstruction(nil, tc.inst)
			isErr := err != nil
			if isErr != tc.err {
				t.Error(err)
			}
			if tc.err {
				return
			}
			if meta.Int64() != int64(tc.meta) {
				t.Errorf("invalid meta, expect %v, got %v", tc.meta, meta)
			}
			if height.Int64() != int64(tc.height) {
				t.Errorf("invalid height, expect %v, got %v", tc.height, height)
			}
			if numVals.Int64() != int64(tc.numVals) {
				t.Errorf("invalid numVals, expect %v, got %v", tc.numVals, numVals)
			}
		})
	}
}

func buildDecodedSwapConfirmInst(meta, shard, height int, addrs []string) []byte {
	a := []byte{}
	for _, addr := range addrs {
		d, _ := hex.DecodeString(addr)
		a = append(a, toBytes32BigEndian(d)...)
	}
	decoded := []byte(strconv.Itoa(meta))
	decoded = append(decoded, []byte(strconv.Itoa(shard))...)
	decoded = append(decoded, toBytes32BigEndian(big.NewInt(int64(height)).Bytes())...)
	decoded = append(decoded, toBytes32BigEndian(big.NewInt(int64(len(addrs))).Bytes())...)
	decoded = append(decoded, a...)
	return decoded
}

func getFixedSwapProof() *decodedProof {
	proofMarshalled := `{"Instruction":"NzIxAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADnIti3HcwBUtR9JDhVakXTNX1jHwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAOjUpRAAhZe/hXIFnnrkCCXXc5QLEkTGVl6kOHevFcD6pI2dSn4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==","BeaconHeight":null,"BridgeHeight":null,"InstPaths":[[[58,76,164,226,143,9,56,24,249,8,101,50,180,182,65,214,160,25,58,19,10,33,82,145,249,51,70,94,253,96,207,119],[174,225,54,126,22,189,67,9,66,120,9,83,111,18,173,218,137,190,75,123,3,136,22,229,206,200,195,97,75,196,117,26]],[]],"InstPathIsLefts":[[false,false],[]],"InstRoots":[[232,187,50,63,150,125,194,188,141,73,224,250,113,204,65,231,155,131,251,202,183,96,21,75,23,200,110,183,224,135,84,182],[94,225,86,247,145,34,10,49,207,19,182,114,127,158,248,70,13,30,37,157,124,0,40,168,109,58,227,250,22,10,145,186]],"BlkData":[[178,88,135,35,194,25,232,206,53,145,97,158,121,132,132,149,45,36,244,34,78,14,5,239,248,194,59,136,29,60,10,144],[23,196,85,167,176,125,122,12,251,231,84,78,247,1,32,90,211,242,229,89,87,5,241,66,115,130,55,84,87,202,81,28]],"SigIdxs":[[0,1,2,3],[0,2,3]],"SigVs":["HBwbHA==","Gxwb"],"SigRs":[[[24,224,28,91,250,108,190,198,108,164,76,76,170,187,86,96,229,99,248,155,109,37,86,102,139,17,253,143,73,155,69,120],[185,213,129,125,155,197,59,112,20,50,99,33,190,94,127,217,106,73,238,16,230,91,27,177,108,251,36,68,45,97,192,250],[159,23,196,41,226,108,149,5,156,16,168,81,219,76,60,198,223,211,224,146,193,220,79,132,10,153,28,199,231,227,21,139],[60,230,139,67,106,249,86,49,165,31,145,42,177,62,92,202,190,17,157,36,27,82,223,176,125,28,77,174,132,54,214,166]],[[197,169,45,86,21,102,63,91,58,74,45,171,141,205,21,238,223,25,227,120,51,96,173,242,69,55,239,167,144,205,209,67],[250,85,133,198,3,115,165,20,77,74,223,0,210,150,115,202,133,10,165,13,87,180,115,76,160,121,88,98,64,120,219,134],[177,140,130,135,167,210,186,230,95,201,243,135,43,72,238,110,222,247,124,165,140,100,59,24,19,215,240,83,5,24,99,28]]],"SigSs":[[[86,253,73,166,149,155,207,147,135,214,70,212,2,94,70,115,172,8,26,35,26,12,191,56,244,131,70,132,217,169,149,67],[91,215,47,166,114,124,202,201,238,57,243,158,38,126,0,178,204,38,53,155,48,145,238,171,222,23,182,213,96,116,77,111],[56,155,154,186,108,0,109,136,166,223,22,38,61,200,224,35,153,174,54,16,235,69,200,172,52,124,93,93,121,217,19,55],[117,34,44,127,120,153,3,121,109,214,254,180,78,154,142,186,58,206,86,5,231,255,134,228,215,250,161,177,244,195,116,158]],[[111,159,225,148,249,194,241,127,208,175,178,144,161,4,69,50,38,156,11,145,225,150,189,159,244,163,147,49,248,180,230,59],[65,117,229,199,4,212,145,43,140,199,192,229,132,142,181,169,62,112,139,235,241,150,218,15,134,97,13,203,126,179,127,71],[35,218,222,233,11,80,205,184,15,123,8,77,36,234,86,217,173,220,156,76,181,26,76,130,123,152,165,154,31,228,231,142]]]}`
	proof := &decodedProof{}
	json.Unmarshal([]byte(proofMarshalled), proof)
	return proof
}
