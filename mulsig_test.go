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
	beaconHeight := proof.beaconHeight.Bytes()
	h := [32]byte{}
	copy(h[32-len(beaconHeight):], beaconHeight)
	beaconInstHash := common.Keccak256(proof.instruction, beaconHeight)
	res, err := c.inc.InstructionInMerkleTree(
		nil,
		beaconInstHash,
		proof.beaconInstRoot,
		proof.beaconInstPath,
		proof.beaconInstPathIsLeft,
		proof.beaconInstPathLen,
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
	comm, err := c.inc.FindComm(nil, proof.bridgeHeight, false)
	fmt.Printf("comm: %x %v\n", comm, err)
	res, err := c.inc.VerifySig(
		nil,
		comm,
		proof.bridgeSignerSig,
		proof.bridgeNumR,
		proof.bridgeXs,
		proof.bridgeYs,
		proof.bridgeRIdxs,
		proof.bridgeNumSig,
		proof.bridgeSigIdxs,
		proof.bridgeRp,
		proof.bridgeRpx,
		proof.bridgeRpy,
		proof.bridgeR,
		proof.bridgeBlkHash,
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
	beaconHeight := proof.beaconHeight.Bytes()
	h := [32]byte{}
	copy(h[32-len(beaconHeight):], beaconHeight)
	inst := [300]byte{}
	copy(inst[:], proof.instruction)
	beaconInstHash := common.Keccak256(inst[:], beaconHeight)
	res, err := c.inc.InstructionApproved(
		nil,
		true,
		beaconInstHash,
		proof.beaconHeight,
		proof.beaconInstPath,
		proof.beaconInstPathIsLeft,
		proof.beaconInstPathLen,
		proof.beaconInstRoot,
		proof.beaconBlkData,
		proof.beaconBlkHash,
		proof.beaconSignerSig,
		proof.beaconNumR,
		proof.beaconXs,
		proof.beaconYs,
		proof.beaconRIdxs,
		proof.beaconNumSig,
		proof.beaconSigIdxs,
		proof.beaconRp,
		proof.beaconRpx,
		proof.beaconRpy,
		proof.beaconR,
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
		proof.beaconXs,
		proof.beaconYs,
		proof.beaconSigIdxs,
		proof.beaconNumSig,
		proof.beaconRpx,
		proof.beaconRpy,
		proof.beaconR,
		proof.beaconSignerSig,
		proof.beaconBlkHash,
	)
	fmt.Println(res, err)

	res, err = c.ms.CheckMulSig(
		nil,
		proof.bridgeXs,
		proof.bridgeYs,
		proof.bridgeSigIdxs,
		proof.bridgeNumSig,
		proof.bridgeRpx,
		proof.bridgeRpy,
		proof.bridgeR,
		proof.bridgeSignerSig,
		proof.bridgeBlkHash,
	)
	fmt.Println(res, err)
}
