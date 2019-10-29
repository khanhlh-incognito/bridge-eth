package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/incognitochain/bridge-eth/bridge/incognito_proxy"
	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/incognitochain/bridge-eth/common/base58"
	"github.com/incognitochain/bridge-eth/consensus/signatureschemes/bridgesig"
	"github.com/incognitochain/bridge-eth/erc20"
	"github.com/incognitochain/bridge-eth/jsonresult"
	"github.com/pkg/errors"
)

type contracts struct {
	v         *vault.Vault
	vAddr     common.Address
	inc       *incognito_proxy.IncognitoProxy
	incAddr   common.Address
	token     *erc20.Erc20
	tokenAddr common.Address
}

type getProofResult struct {
	Result jsonresult.GetInstructionProof
	Error  struct {
		Code       int
		Message    string
		StackTrace string
	}
}

type decodedProof struct {
	Instruction []byte
	Heights     [2]*big.Int

	InstPaths       [2][][32]byte
	InstPathIsLefts [2][]bool
	InstRoots       [2][32]byte
	BlkData         [2][32]byte
	SigIdxs         [2][]*big.Int
	SigVs           [2][]uint8
	SigRs           [2][][32]byte
	SigSs           [2][][32]byte
}

func getAndDecodeBurnProof(txID string) (*decodedProof, error) {
	body := getBurnProof(txID)
	if len(body) < 1 {
		return nil, fmt.Errorf("burn proof not found")
	}

	r := getProofResult{}
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		return nil, err
	}
	return decodeProof(&r)
}

func getCommittee(url string) ([]common.Address, []common.Address, error) {
	payload := strings.NewReader("{\n    \"id\": 1,\n    \"jsonrpc\": \"1.0\",\n    \"method\": \"getbeaconbeststate\",\n    \"params\": []\n}")

	req, _ := http.NewRequest("POST", url, payload)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	type beaconBestStateResult struct {
		BeaconCommittee []string
		ShardCommittee  map[string][]string
	}

	type getBeaconBestStateResult struct {
		Result beaconBestStateResult
		Error  string
		Id     int
	}

	r := getBeaconBestStateResult{}
	err = json.Unmarshal([]byte(body), &r)
	if err != nil {
		return nil, nil, err
	}

	// Genesis committee
	beaconOld := make([]common.Address, len(r.Result.BeaconCommittee))
	for i, pk := range r.Result.BeaconCommittee {
		cpk := &CommitteePublicKey{}
		cpk.FromString(pk)
		addr, err := convertPubkeyToAddress(*cpk)
		if err != nil {
			return nil, nil, err
		}
		beaconOld[i] = addr
		fmt.Printf("beaconOld: %s\n", addr.Hex())
	}

	bridgeOld := make([]common.Address, len(r.Result.ShardCommittee["1"]))
	for i, pk := range r.Result.ShardCommittee["1"] {
		cpk := &CommitteePublicKey{}
		cpk.FromString(pk)
		addr, err := convertPubkeyToAddress(*cpk)
		if err != nil {
			return nil, nil, err
		}
		bridgeOld[i] = addr
		fmt.Printf("bridgeOld: %s\n", addr.Hex())
	}

	return beaconOld, bridgeOld, nil
}

func getBurnProof(txID string) string {
	url := "http://127.0.0.1:9344"
	// url := "https://dev-test-node.incognito.org/"

	if len(txID) == 0 {
		txID = "87c89c1c19cec3061eff9cfefdcc531d9456ac48de568b3974c5b0a88d5f3834"
	}
	payload := strings.NewReader(fmt.Sprintf("{\n    \"id\": 1,\n    \"jsonrpc\": \"1.0\",\n    \"method\": \"getburnproof\",\n    \"params\": [\n    \t\"%s\"\n    ]\n}", txID))

	req, _ := http.NewRequest("POST", url, payload)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("err:", err)
		return ""
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(string(body))
	return string(body)
}

func decodeProof(r *getProofResult) (*decodedProof, error) {
	inst := decode(r.Result.Instruction)
	fmt.Printf("inst: %d %x\n", len(inst), inst)

	// Block heights
	beaconHeight := big.NewInt(0).SetBytes(decode(r.Result.BeaconHeight))
	bridgeHeight := big.NewInt(0).SetBytes(decode(r.Result.BridgeHeight))
	heights := [2]*big.Int{beaconHeight, bridgeHeight}
	fmt.Printf("beaconHeight: %d\n", beaconHeight)
	fmt.Printf("bridgeHeight: %d\n", bridgeHeight)

	beaconInstRoot := decode32(r.Result.BeaconInstRoot)
	beaconInstPath := make([][32]byte, len(r.Result.BeaconInstPath))
	beaconInstPathIsLeft := make([]bool, len(r.Result.BeaconInstPath))
	for i, path := range r.Result.BeaconInstPath {
		beaconInstPath[i] = decode32(path)
		beaconInstPathIsLeft[i] = r.Result.BeaconInstPathIsLeft[i]
	}
	// fmt.Printf("beaconInstRoot: %x\n", beaconInstRoot)

	beaconBlkData := toByte32(decode(r.Result.BeaconBlkData))
	fmt.Printf("data: %s %s\n", r.Result.BeaconBlkData, r.Result.BeaconInstRoot)
	fmt.Printf("expected beaconBlkHash: %x\n", keccak256(beaconBlkData[:], beaconInstRoot[:]))

	beaconSigVs, beaconSigRs, beaconSigSs, err := decodeSigs(r.Result.BeaconSigs)
	if err != nil {
		return nil, err
	}

	beaconSigIdxs := []*big.Int{}
	for _, sIdx := range r.Result.BeaconSigIdxs {
		beaconSigIdxs = append(beaconSigIdxs, big.NewInt(int64(sIdx)))
	}

	// For bridge
	bridgeInstRoot := decode32(r.Result.BridgeInstRoot)
	bridgeInstPath := make([][32]byte, len(r.Result.BridgeInstPath))
	bridgeInstPathIsLeft := make([]bool, len(r.Result.BridgeInstPath))
	for i, path := range r.Result.BridgeInstPath {
		bridgeInstPath[i] = decode32(path)
		bridgeInstPathIsLeft[i] = r.Result.BridgeInstPathIsLeft[i]
	}
	// fmt.Printf("bridgeInstRoot: %x\n", bridgeInstRoot)
	bridgeBlkData := toByte32(decode(r.Result.BridgeBlkData))

	bridgeSigVs, bridgeSigRs, bridgeSigSs, err := decodeSigs(r.Result.BridgeSigs)
	if err != nil {
		return nil, err
	}

	bridgeSigIdxs := []*big.Int{}
	for _, sIdx := range r.Result.BridgeSigIdxs {
		bridgeSigIdxs = append(bridgeSigIdxs, big.NewInt(int64(sIdx)))
		// fmt.Printf("bridgeSigIdxs[%d]: %d\n", i, j)
	}

	// Merge beacon and bridge proof
	instPaths := [2][][32]byte{beaconInstPath, bridgeInstPath}
	instPathIsLefts := [2][]bool{beaconInstPathIsLeft, bridgeInstPathIsLeft}
	instRoots := [2][32]byte{beaconInstRoot, bridgeInstRoot}
	blkData := [2][32]byte{beaconBlkData, bridgeBlkData}
	sigIdxs := [2][]*big.Int{beaconSigIdxs, bridgeSigIdxs}
	sigVs := [2][]uint8{beaconSigVs, bridgeSigVs}
	sigRs := [2][][32]byte{beaconSigRs, bridgeSigRs}
	sigSs := [2][][32]byte{beaconSigSs, bridgeSigSs}

	return &decodedProof{
		Instruction:     inst,
		Heights:         heights,
		InstPaths:       instPaths,
		InstPathIsLefts: instPathIsLefts,
		InstRoots:       instRoots,
		BlkData:         blkData,
		SigIdxs:         sigIdxs,
		SigVs:           sigVs,
		SigRs:           sigRs,
		SigSs:           sigSs,
	}, nil
}

func decodeSigs(sigs []string) (
	sigVs []uint8,
	sigRs [][32]byte,
	sigSs [][32]byte,
	err error,
) {
	sigVs = make([]uint8, len(sigs))
	sigRs = make([][32]byte, len(sigs))
	sigSs = make([][32]byte, len(sigs))
	for i, sig := range sigs {
		v, r, s, e := bridgesig.DecodeECDSASig(decode(sig))
		if e != nil {
			err = e
			return
		}
		sigVs[i] = uint8(v)
		copy(sigRs[i][:], r)
		copy(sigSs[i][:], s)
	}
	return
}

func toByte32(s []byte) [32]byte {
	a := [32]byte{}
	copy(a[:], s)
	return a
}

func decode(s string) []byte {
	d, _ := hex.DecodeString(s)
	return d
}

func decode32(s string) [32]byte {
	return toByte32(decode(s))
}

func keccak256(b ...[]byte) [32]byte {
	h := crypto.Keccak256(b...)
	r := [32]byte{}
	copy(r[:], h)
	return r
}

func getCommitteeHardcoded() ([]common.Address, []common.Address) {
	beaconComm := []string{
		"0x77b9F481f979e16Cf4234866C0803c2e65106862",
		"0x8F4aEd5Adb0347eF2Db3bDce9f0DF2747D0107E8",
		"0xcf836142D459B1257ed52aef34E62F6F4e7eF894",
		"0x4C265F5eb8Eb68A1eA99626cD838558836438f80",
		"0xc7CC1EE53eF28CC551e97476F0dB01596D945fE0",
		"0xFd119B9DBEb478154E3650F14f55db3787E1bd38",
		"0x780329f064F0BE00FdbDeA4bA8A5b04C7AB2866c",
	}
	beacons := make([]common.Address, len(beaconComm))
	for i, p := range beaconComm {
		beacons[i] = common.HexToAddress(p)
	}

	bridgeComm := []string{
		"0xaced3cf99897a55057B7513b8505c22DaF9378D9",
		"0x8CbDC490c4188b721e210622027a54E14f27CA7F",
		"0x74d0C0A0A5f89527b4e2850EA09d4F6cE9BBb3bB",
		"0x6964D5c5A7C1503E2228852d1EC115c0e7a20593",
		"0x803c90C23a8a34a676B57CaF0372026C988B416d",
		"0x17E21A7a018046ab3cAE7Aab4215087a2497a7D7",
		"0x5bA8281b5BE1F864E52B3ef8FcEF80560e41005C",
		"0x9Ee3002A85701ae62B16e92e0d8F2044D79a35B6",
		"0xb7eF123cc555cA977aE2fbB5A3690ce57628C664",
		"0x211880118421A814Da0151A4bd06be703DB3654e",
		"0x6A25Ed4Ef6Fa034c895D5721D73dBEC5163Ad4f9",
		"0x1A7232f56F4D71e794D8Bbfc5fa5991d544e1C9f",
		"0xcFcFc3A2CC9DFdF98aC075441E45818C7A70a29e",
		"0x99118446796dFa58d8327834347806711f67Cb79",
		"0x8986acdde31E4519acFcabb139Fd2A2B1da274b2",
		"0xE59C59D87f52D39B1BB8136966e0E1817D7a845A",
		"0x604589220D909878ebDC906d0b33b433Fc3cc0a3",
		"0xf069494c92A85DD31FE6850D8EfE6F2398Ea072c",
		"0x93E1b517726d05c235AE3AF53fa84C34d400Cae9",
		"0x6284C7FD0F623E705d0e0a2D4621299B98eA3895",
		"0xf57Ac7832b1C8F7f5C3E228eF7811D58647A70BF",
		"0x8fa98CBa06b199922E9Acc5749F25FF549e5eEbd",
	}
	bridges := make([]common.Address, len(bridgeComm))
	for i, p := range bridgeComm {
		bridges[i] = common.HexToAddress(p)
	}
	return beacons, bridges
}

func convertPubkeyToAddress(cKey CommitteePublicKey) (common.Address, error) {
	pk, err := crypto.DecompressPubkey(cKey.MiningPubKey[BRI_CONSENSUS])
	if err != nil {
		return common.Address{}, errors.Wrapf(err, "cKey: %+v", cKey)
	}
	address := crypto.PubkeyToAddress(*pk)
	return address, nil
}

var BRI_CONSENSUS = "dsa"

type CommitteePublicKey struct {
	IncPubKey    []byte
	MiningPubKey map[string][]byte
}

func (pubKey *CommitteePublicKey) FromString(keyString string) error {
	keyBytes, ver, err := base58.Base58Check{}.Decode(keyString)
	if (ver != 0) || (err != nil) {
		return errors.New("Wrong input")
	}
	return json.Unmarshal(keyBytes, pubKey)
}
