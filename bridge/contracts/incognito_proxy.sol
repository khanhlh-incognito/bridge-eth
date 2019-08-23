pragma solidity >=0.5.0 <0.6.0;
pragma experimental ABIEncoderV2;

contract IncognitoProxy {
    uint public latestBeaconBlk;
    uint public latestBridgeBlk;

    constructor(address[] memory beaconCommittee, address[] memory bridgeCommittee) public {
        latestBridgeBlk = beaconCommittee.length;
        latestBridgeBlk = bridgeCommittee.length;
    }

    function swapBridgeCommittee(
        bytes memory inst,
        bytes32[][2] memory instPaths,
        bool[][2] memory instPathIsLefts,
        bytes32[2] memory instRoots,
        bytes32[2] memory blkData,
        uint[][2] memory sigIdxs,
        uint8[][2] memory sigVs,
        bytes32[][2] memory sigRs,
        bytes32[][2] memory sigSs
    ) public {
        // TODO: Check if beaconInstRoot is in block
        bytes32 instHash;

        // TODO: Metadata type and shardID of swap bridge instruction

        // Verify instruction on beacon
        require(instructionApproved(
            true,
            instHash,
            latestBeaconBlk,
            instPaths[0],
            instPathIsLefts[0],
            instRoots[0],
            blkData[0],
            sigIdxs[0],
            sigVs[0],
            sigRs[0],
            sigSs[0]
        ));

        // Verify instruction on bridge
        require(instructionApproved(
            false,
            instHash,
            latestBridgeBlk,
            instPaths[1],
            instPathIsLefts[1],
            instRoots[1],
            blkData[1],
            sigIdxs[1],
            sigVs[1],
            sigRs[1],
            sigSs[1]
        ));

        // Parse instruction and check metadata
        // (uint meta, uint startHeight, uint numVals) = extractMetaFromInst(inst);
        // address[] memory pubkeys = extractCommitteeFromInst(inst, numVals);

        // TODO: Swap committee
        latestBridgeBlk = 0;
    }

    function instructionApproved(
        bool isBeacon,
        bytes32 instHash,
        uint blkHeight,
        bytes32[] memory instPath,
        bool[] memory instPathIsLeft,
        bytes32 instRoot,
        bytes32 blkData,
        uint[] memory sigIdx,
        uint8[] memory sigV,
        bytes32[] memory sigR,
        bytes32[] memory sigS
    ) public view returns (bool) {
        // TODO: Find committee in charge of this block
        uint a = latestBridgeBlk;

        // TODO: Get pubkey of signers
        address[] memory signers = new address[](sigIdx.length);

        // TODO: Get block hash from instRoot and other data

        // TODO: Check if enough validators signed this block

        // Check that signature is correct
        bytes32 blk = 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8;
        require(verifySig(signers, blk, sigV, sigR, sigS));

        // Check that inst is in block
        require(instructionInMerkleTree(
            instHash,
            instRoot,
            instPath,
            instPathIsLeft
        ));
    }

    function instructionInMerkleTree(
        bytes32 leaf,
        bytes32 root,
        bytes32[] memory path,
        bool[] memory left
    ) public pure returns (bool) {
        // TODO
        return true;
    }

    function extractMetaFromInst(bytes memory inst) public pure returns(uint, uint, uint) {
        // TODO
        return (0, 0, 0);
    }

    function extractCommitteeFromInst(bytes memory inst, uint numVals) public pure returns (address[] memory) {
        // TODO
        address[] memory addr = new address[](numVals);
        return addr;
    }

    function verifySig(
        address[] memory committee,
        bytes32 msgHash,
        uint8[] memory v,
        bytes32[] memory r,
        bytes32[] memory s
    ) public pure returns (bool) {
        for (uint i = 0; i < committee.length; i++){
            if (ecrecover(msgHash, v[i], r[i], s[i]) != committee[i]) {
                return false;
            }
        }
        return true;
    }
}
