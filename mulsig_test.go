package bridge

import (
	"fmt"
	"testing"

	"github.com/incognitochain/bridge-eth/common"
)

func TestInstructionInMerkleTree(t *testing.T) {
	// Get proof
	proof, err := getAndDecodeBurnProof("")
	if err != nil {
		t.Fatal(err)
	}

	_, c := connectAndInstantiate(t)
	beaconHeight := proof.BeaconHeight.Bytes()
	h := [32]byte{}
	copy(h[32-len(beaconHeight):], beaconHeight)
	beaconInstHash := common.Keccak256(proof.Instruction, beaconHeight)
	res, err := c.inc.InstructionInMerkleTree(
		nil,
		beaconInstHash,
		proof.BeaconInstRoot,
		proof.BeaconInstPath,
		proof.BeaconInstPathIsLeft,
		proof.BeaconInstPathLen,
	)
	fmt.Println(res, err)
}

func TestVerifySig(t *testing.T) {
	// Get proof
	proof, err := getAndDecodeBurnProof("")
	if err != nil {
		t.Fatal(err)
	}

	_, c := connectAndInstantiate(t)
	comm, _, err := c.inc.FindComm(nil, proof.BridgeHeight, false)
	fmt.Printf("comm: %x %v\n", comm, err)
	res, err := c.inc.VerifySig(
		nil,
		comm,
		proof.BridgeSignerSig,
		proof.BridgeNumR,
		proof.BridgeXs,
		proof.BridgeYs,
		proof.BridgeRIdxs,
		proof.BridgeNumSig,
		proof.BridgeSigIdxs,
		proof.BridgeRp,
		proof.BridgeRpx,
		proof.BridgeRpy,
		proof.BridgeR,
		proof.BridgeBlkHash,
	)
	fmt.Println(res, err)
}

func TestInstructionApproved(t *testing.T) {
	// Get proof
	proof, err := getAndDecodeBurnProof("")
	if err != nil {
		t.Fatal(err)
	}

	_, c := connectAndInstantiate(t)
	beaconHeight := proof.BeaconHeight.Bytes()
	h := [32]byte{}
	copy(h[32-len(beaconHeight):], beaconHeight)
	inst := [300]byte{}
	copy(inst[:], proof.Instruction)
	beaconInstHash := common.Keccak256(inst[:], beaconHeight)
	res, err := c.inc.InstructionApproved(
		nil,
		true,
		beaconInstHash,
		proof.BeaconHeight,
		proof.BeaconInstPath,
		proof.BeaconInstPathIsLeft,
		proof.BeaconInstPathLen,
		proof.BeaconInstRoot,
		proof.BeaconBlkData,
		proof.BeaconBlkHash,
		proof.BeaconSignerSig,
		proof.BeaconNumR,
		proof.BeaconXs,
		proof.BeaconYs,
		proof.BeaconRIdxs,
		proof.BeaconNumSig,
		proof.BeaconSigIdxs,
		proof.BeaconRp,
		proof.BeaconRpx,
		proof.BeaconRpy,
		proof.BeaconR,
	)
	fmt.Println(res, err)
}

func TestMulSig(t *testing.T) {
	// Get proof
	proof, err := getAndDecodeBurnProof("")
	if err != nil {
		t.Fatal(err)
	}

	_, c := connectAndInstantiate(t)
	res, err := c.ms.CheckMulSig(
		nil,
		proof.BeaconXs,
		proof.BeaconYs,
		proof.BeaconSigIdxs,
		proof.BeaconNumSig,
		proof.BeaconRpx,
		proof.BeaconRpy,
		proof.BeaconR,
		proof.BeaconSignerSig,
		proof.BeaconBlkHash,
	)
	fmt.Println(res, err)

	res, err = c.ms.CheckMulSig(
		nil,
		proof.BridgeXs,
		proof.BridgeYs,
		proof.BridgeSigIdxs,
		proof.BridgeNumSig,
		proof.BridgeRpx,
		proof.BridgeRpy,
		proof.BridgeR,
		proof.BridgeSignerSig,
		proof.BridgeBlkHash,
	)
	fmt.Println(res, err)
}
