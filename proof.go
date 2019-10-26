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
		"0x3cD69B1A595B7A9589391538d29ee7663326e4d3",
		"0xc687470342f4E80ECEf6bBd25e276266d40b8429",
		"0x2A40c96b41AdEc5641F28eF923e270B73e29bb53",
		"0x131B772A9ADe1793F000024eAb23b77bEd3BFe64",
	}
	beacons := make([]common.Address, len(beaconComm))
	for i, p := range beaconComm {
		beacons[i] = common.HexToAddress(p)
	}

	bridgeComm := []string{
		"0x28655822DAf6c4B32303B06e875F92dC6e242cE4",
		"0xD2902ab2F5dF2b17C5A5aa380f511F04a2542E10",
		"0xB67376ad63EAdC22f05efE428e93f09D4f13B4fD",
		"0x40bAA64EAFbD355f5427d127979f377cfA48cc10",
		"0x3A4bEd76c180Ec9A30188612A1e6286d066CAb8E",
		"0x36EB03fcb2D726AF268D08f8B88d898C38cB106b",
		"0x951eE2C001A7e85fd3C05824C65118dd8cCCAcE2",
		"0xbB3C4E1F96aF584D7f6bF09B8787A6F9d4Ec8086",
		"0x8EE26caC71869c9319a51E2E75F2EbeBc123Ca41",
		"0xA1f8B6D06A342430C085f781219D9E422355448C",
		"0x2A1c3ac8Fcd510d1858bFFd4f7725834730a8B99",
		"0x2B63D31edd7f0F6d79552cc4D588a50e718fBc60",
		"0xab72FBD0E97012D735fdc0852Fe49497bdb39670",
		"0x859f8C5f27639cAc508838Fc37BA0E7D9d575040",
		"0x1353c08E0A996BA7BB0cdcB06B3a145C68C63197",
		"0xB6275a14D14FC57F884c546A1d181814e5769401",
		"0x641d2bDc33f17e5626E3E3CA51e21099BCF40956",
		"0x9a0dc54A0894401fC02d088A42c26FaE74AE7b1c",
		"0x7d209B6F57D852e08ceff7222696C8a0E60b3Be6",
		"0xb57165795D96cE1b1010623e24E9fB56c1ff1Ae9",
		"0x6E6b5F936421E528048C034CFd80Fb4de94156A2",
		"0x7030AAA1b347c35c60067221E127750746d61c7E",
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
