package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"math/big"
	"strings"
	"testing"

	ec "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/incognitochain/bridge-eth/blockchain"
	"github.com/incognitochain/bridge-eth/common"
	"github.com/incognitochain/bridge-eth/consensus/bridgesig"
	"github.com/pkg/errors"
)

func TestFixedVerifySig(t *testing.T) {
	p, _ := setupFixedCommittee()

	testCases := []struct {
		desc string
		in   *committeeSig
		out  bool
		err  bool
	}{
		{
			desc: "Valid sig",
			in:   getFixedCommitteeSig(),
			out:  true,
		},
		{
			desc: "Invalid committee",
			in: func() *committeeSig {
				sig := getFixedCommitteeSig()
				sig.addrs[1][2] = 123
				return sig
			}(),
			out: false,
		},
		{
			desc: "Invalid msgHash",
			in: func() *committeeSig {
				sig := getFixedCommitteeSig()
				sig.msgHash[0] = 123
				return sig
			}(),
			out: false,
		},
		{
			desc: "Invalid v",
			in: func() *committeeSig {
				sig := getFixedCommitteeSig()
				sig.v[1] = 123
				return sig
			}(),
			out: false,
		},
		{
			desc: "Invalid r",
			in: func() *committeeSig {
				sig := getFixedCommitteeSig()
				sig.r[2][3] = 123
				return sig
			}(),
			out: false,
		},
		{
			desc: "Invalid s",
			in: func() *committeeSig {
				sig := getFixedCommitteeSig()
				sig.s[3][4] = 123
				return sig
			}(),
			out: false,
		},
		{
			desc: "Not enough committee members",
			in: func() *committeeSig {
				sig := getFixedCommitteeSig()
				sig.addrs = sig.addrs[:2]
				return sig
			}(),
			err: true,
		},
		{
			desc: "Not enough v",
			in: func() *committeeSig {
				sig := getFixedCommitteeSig()
				sig.v = sig.v[:2]
				return sig
			}(),
			err: true,
		},
		{
			desc: "Not enough r",
			in: func() *committeeSig {
				sig := getFixedCommitteeSig()
				sig.r = sig.r[:2]
				return sig
			}(),
			err: true,
		},
		{
			desc: "Not enough s",
			in: func() *committeeSig {
				sig := getFixedCommitteeSig()
				sig.s = sig.s[:2]
				return sig
			}(),
			err: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			res, err := p.inc.VerifySig(nil, tc.in.addrs, tc.in.msgHash, tc.in.v, tc.in.r, tc.in.s)
			isErr := err != nil
			if isErr != tc.err {
				t.Error(errors.Errorf("expect error = %t, got %v", tc.err, err))
			}
			if tc.err {
				return
			}
			if res != tc.out {
				t.Errorf("verifySig failed, expect %v, got %v", tc.out, res)
			}
		})
	}
}

func getFixedCommitteeSig() *committeeSig {
	validationData := "{\"ProducerBLSSig\":\"D4sg/eVi8yI+rX9WOwCWBEG+4mWXjGNorl2m3ppRCvE=\",\"ProducerBriSig\":null,\"ValidatiorsIdx\":[0,1,2,3],\"AggSig\":\"AriRXDvXcPDqkMNAQjHR61f3xis6YLNskuYF7vQJNzE=\",\"BridgeSig\":[\"/tSXMa9s1PKAxDC9H6etSMPcnAOEqqQYum3TfWtOKQpyvHxA1jllDkLmB68M6pp54bTUWenqXMQVWW+2GAcBjgA=\",\"MJyhaCCm8B6uwK/w6/OMqr7AW1Szo1etRfTcru0ZenZUwea0LVXhPo2QRKeO+Q1n12J2yRv4sUkhRLLL9zw1SwE=\",\"DOpccVDrw6SbGqs4+YP/Ti1nx4gg/xpsuHB7DBuhO2RMl8hAaUz2TVZ6hv+r8z0YLiUw/k6FEFY+5dg/EjMRAQA=\",\"qPEXt4KgFR8ZMw7JelEeEwsWQ7gW/IrzWMpx++zjQ6dLdeXwKcGwxoaBWhWnEpma+MVVQw1LvzzuvtzIBGZDKgE=\"]}"
	d, _ := DecodeValidationData(validationData)
	vs := []byte{}
	rs := [][32]byte{}
	ss := [][32]byte{}
	for _, sig := range d.BridgeSig {
		v, r, s, _ := bridgesig.DecodeECDSASig(sig)
		vs = append(vs, v)
		rs = append(rs, toByte32(r))
		ss = append(ss, toByte32(s))
	}

	hash, _ := common.Hash{}.NewHashFromStr("cb53ba7574335ecfa0fddcb136b387330af322784fb759c80ca7bb790a1c0f9d")
	addrs, _ := getFixedCommittee()
	msgHash := toByte32(crypto.Keccak256Hash(hash.GetBytes()).Bytes())
	return &committeeSig{
		addrs:   addrs,
		msgHash: msgHash,
		v:       vs,
		r:       rs,
		s:       ss,
	}
}

type committeeSig struct {
	addrs   []ec.Address
	msgHash [32]byte
	v       []uint8
	r       [][32]byte
	s       [][32]byte
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
		meta    uint8
		shard   uint8
		height  int
		numVals int
		err     bool
	}{
		{
			desc:    "Extract swap beacon instruction",
			inst:    buildDecodedSwapConfirmInst(70, 1, 123, addrs),
			meta:    70,
			shard:   1,
			height:  123,
			numVals: len(addrs),
		},
		{
			desc:    "Extract swap bridge instruction",
			inst:    buildDecodedSwapConfirmInst(71, 2, 19827312, addrs[:2]),
			meta:    71,
			shard:   2,
			height:  19827312,
			numVals: 2,
		},
		{
			desc: "Instruction too short",
			inst: make([]byte, 65),
			err:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			meta, shard, height, numVals, err := p.inc.ExtractMetaFromInstruction(nil, tc.inst)
			isErr := err != nil
			if isErr != tc.err {
				t.Error(errors.Errorf("expect error = %t, got %v", tc.err, err))
			}
			if tc.err {
				return
			}
			if meta != tc.meta {
				t.Errorf("invalid meta, expect %v, got %v", tc.meta, meta)
			}
			if shard != tc.shard {
				t.Errorf("invalid shard, expect %v, got %v", tc.shard, shard)
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

func TestExtractCommitteeFromInstruction(t *testing.T) {
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
		numVals int
		out     []string
		err     bool
	}{
		{
			desc:    "Extract beacon committee",
			inst:    buildDecodedSwapConfirmInst(70, 1, 789, addrs),
			numVals: len(addrs),
			out:     addrs,
		},
		{
			desc:    "Extract bridge committee",
			inst:    buildDecodedSwapConfirmInst(71, 1, 19827312, addrs[:2]),
			numVals: 2,
			out:     addrs[:2],
		},
		{
			desc:    "Instruction too short",
			inst:    make([]byte, 97),
			numVals: 1,
			err:     true,
		},
		{
			desc:    "numVals too big",
			inst:    buildDecodedSwapConfirmInst(70, 1, 789, addrs),
			numVals: 8,
			err:     true,
		},
		{
			desc:    "numVals too low",
			inst:    buildDecodedSwapConfirmInst(70, 1, 789, addrs),
			numVals: 2,
			err:     true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			comm, err := p.inc.ExtractCommitteeFromInstruction(nil, tc.inst, big.NewInt(int64(tc.numVals)))
			isErr := err != nil
			if isErr != tc.err {
				t.Error(errors.Errorf("expect error = %t, got %v", tc.err, err))
			}
			if tc.err {
				return
			}

			for i, c := range comm {
				addr := c.Hex()
				addr = addr[2:] // ignore 0x
				if strings.ToLower(addr) != tc.out[i] {
					t.Errorf("invalid committee[%d], expect %v, got %v", i, tc.out[i], addr)
				}
			}
		})
	}
}

func TestInstructionInMerkleTree(t *testing.T) {
	p, _ := setupFixedCommittee()
	testCases := []struct {
		desc string
		in   *merklePath
		out  bool
		err  bool
	}{
		{
			desc: "In merkle tree",
			in:   buildMerklePathTestcase(8, 6, 6),
			out:  true,
		},
		{
			desc: "Not in merkle tree",
			in:   buildMerklePathTestcase(8, 6, 4),
			out:  false,
		},
		{
			desc: "Random leaf",
			in:   buildMerklePathTestcase(8, 6, -1),
			out:  false,
		},
		{
			desc: "Big tree",
			in:   buildMerklePathTestcase(100000, 12345, 12345),
			out:  true,
		},
		{
			desc: "Single node",
			in:   buildMerklePathTestcase(1, 0, 0),
			out:  true,
		},
		{
			desc: "Invalid left.length",
			in: func() *merklePath {
				mp := buildMerklePathTestcase(10, 9, 9)
				mp.left = mp.left[:len(mp.left)-2]
				return mp
			}(),
			err: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			isIn, err := p.inc.InstructionInMerkleTree(nil, tc.in.leaf, tc.in.root, tc.in.path, tc.in.left)
			isErr := err != nil
			if isErr != tc.err {
				t.Error(errors.Errorf("expect error = %t, got %v", tc.err, err))
			}
			if tc.err {
				return
			}

			if tc.out != isIn {
				t.Errorf("check inst in merkle tree error, expect %v, got %v", tc.out, isIn)
			}
		})
	}
}

func buildMerklePathTestcase(numInst, startNodeID, leafID int) *merklePath {
	mp := buildInstructionMerklePath(numInst, startNodeID)
	if leafID < 0 {
		// Randomize 32 bytes
		h := randomMerkleHashes(1)
		mp.leaf = toByte32(h[0])
	} else {
		mp.leaf = toByte32(mp.merkles[leafID])
	}
	return mp
}

func buildInstructionMerklePath(numInst, startNodeID int) *merklePath {
	data := randomMerkleHashes(numInst)
	merkles := blockchain.BuildKeccak256MerkleTree(data)
	p, l := blockchain.GetKeccak256MerkleProofFromTree(merkles, startNodeID)
	path := [][32]byte{}
	left := []bool{}
	for i, x := range p {
		path = append(path, toByte32(x))
		left = append(left, l[i])
	}

	return &merklePath{
		merkles: merkles,
		leaf:    toByte32(merkles[startNodeID]),
		root:    toByte32(merkles[len(merkles)-1]),
		path:    path,
		left:    left,
	}
}

func randomMerkleHashes(n int) [][]byte {
	h := [][]byte{}
	for i := 0; i < n; i++ {
		b := make([]byte, 32)
		rand.Read(b)
		h = append(h, b)
	}
	return h
}

type merklePath struct {
	merkles [][]byte
	leaf    [32]byte
	root    [32]byte
	path    [][32]byte
	left    []bool
}

func TestIncognitoProxyConstructor(t *testing.T) {
	p, _ := setupFixedCommittee()
	beaconStart, err := p.inc.BeaconCommittees(nil, big.NewInt(0))
	if err != nil {
		t.Error(err)
	}
	if beaconStart.Uint64() != uint64(0) {
		t.Errorf("incorrect startBlock, expect 0, got %d", beaconStart)
	}
	bridgeStart, err := p.inc.BridgeCommittees(nil, big.NewInt(0))
	if err != nil {
		t.Error(err)
	}
	if bridgeStart.Uint64() != uint64(0) {
		t.Errorf("incorrect startBlock, expect 0, got %d", bridgeStart)
	}
}

func buildDecodedSwapConfirmInst(meta, shard, height int, addrs []string) []byte {
	a := []byte{}
	for _, addr := range addrs {
		d, _ := hex.DecodeString(addr)
		a = append(a, toBytes32BigEndian(d)...)
	}
	decoded := []byte{byte(meta)}
	decoded = append(decoded, byte(shard))
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
