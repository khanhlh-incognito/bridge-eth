pragma solidity >=0.5.0 <0.6.0;
pragma experimental ABIEncoderV2;

contract IncognitoProxy {
    struct Committee {
        address[] pubkeys;
        uint startBlock;
    }

    Committee[] public beaconCommittees;
    Committee[] public bridgeCommittees;

    event LogUint(uint val);
    event LogString(string val);
    event LogBytes32(bytes32 val);
    event LogAddress(address val);

    constructor(address[] memory beaconCommittee, address[] memory bridgeCommittee) public {
        beaconCommittees.push(Committee({
            pubkeys: beaconCommittee,
            startBlock: 0
        }));

        bridgeCommittees.push(Committee({
            pubkeys: bridgeCommittee,
            startBlock: 0
        }));
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
        bytes32 instHash = keccak256(inst);
        // emit LogBytes32(instHash);
        // emit LogBytes32(instRoots[0]);
        // emit LogBytes32(instRoots[1]);

        // Verify instruction on beacon
        uint latestBeaconBlk = beaconCommittees[beaconCommittees.length-1].startBlock;
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
        uint latestBridgeBlk = bridgeCommittees[bridgeCommittees.length-1].startBlock;
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
        // (uint meta, uint startHeight, uint numVals) = extractMetaFromInstruction(inst);
        // emit LogUint(meta);
        // emit LogUint(startHeight);
        // emit LogUint(numVals);
        // (uint meta, address token, address to, uint amount) = parseBurnInstruction(inst);
        // emit LogUint(meta);
        // emit LogAddress(token);
        // emit LogAddress(to);
        // emit LogUint(amount);
        // address[] memory pubkeys = extractCommitteeFromInstruction(inst, numVals);

        // TODO: Swap committee
        // latestBridgeBlk = 0;
        emit LogString("Done");
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
    ) public returns (bool) {
        // Find committee in charge of this block
        address[] memory signers = findCommitteeFromHeight(blkHeight, isBeacon);
        emit LogAddress(signers[0]);
        emit LogAddress(signers[1]);

        // TODO: Get block hash from instRoot and other data

        // TODO: Check if enough validators signed this block

        // Check that signature is correct
        bytes32 blk = 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8;
        require(verifySig(signers, blk, sigV, sigR, sigS));

        // // Check that inst is in block
        // require(instructionInMerkleTree(
        //     instHash,
        //     instRoot,
        //     instPath,
        //     instPathIsLeft
        // ));
        return true;
    }

    function findCommitteeFromHeight(uint blkHeight, bool isBeacon) public view returns (address[] memory committee) {
        // TODO: optmized with binary search
        if (isBeacon) {
            uint n = beaconCommittees.length;
            for (uint i = n; i > 0; i--) {
                if (beaconCommittees[i-1].startBlock <= blkHeight) {
                    return beaconCommittees[i-1].pubkeys;
                }
            }
        } else {
            uint n = bridgeCommittees.length;
            for (uint i = n; i > 0; i--) {
                if (bridgeCommittees[i-1].startBlock <= blkHeight) {
                    return bridgeCommittees[i-1].pubkeys;
                }
            }
        }
    }

    function instructionInMerkleTree(
        bytes32 leaf,
        bytes32 root,
        bytes32[] memory path,
        bool[] memory left
    ) public returns (bool) {
        bytes32 hash = leaf;
        for (uint i = 0; i < path.length; i++) {
            if (left[i]) {
                hash = keccak256(abi.encodePacked(path[i], hash));
            } else if (path[i] == 0x0) {
                hash = keccak256(abi.encodePacked(hash, hash));
            } else {
                hash = keccak256(abi.encodePacked(hash, path[i]));
            }
        }
        return hash == root;
    }

    function extractMetaFromInstruction(bytes memory inst) public pure returns(uint, uint, uint) {
        uint meta = uint8(inst[2]) + uint8(inst[1]) * 2 ** 8 + uint8(inst[0]) * 2 ** 16;
        uint height;
        uint numVals;
        assembly {
            // skip first 0x20 bytes (stored length of inst)
            height := mload(add(inst, 0x23)) // [3:35]
            numVals := mload(add(inst, 0x43)) // [35:67]
        }
        return (meta, height, numVals);
    }

    function parseBurnInstruction(bytes memory inst) public pure returns (uint, address, address, uint) {
        uint meta = uint8(inst[2]) + uint8(inst[1]) * 2 ** 8 + uint8(inst[0]) * 2 ** 16;
        address token;
        address to;
        uint amount;
        assembly {
            // skip first 0x20 bytes (stored length of inst)
            token := mload(add(inst, 0x23)) // [3:35]
            to := mload(add(inst, 0x43)) // [35:67]
            amount := mload(add(inst, 0x63)) // [67:99]
        }
        return (meta, token, to, amount);
    }

    function extractCommitteeFromInstruction(bytes memory inst, uint numVals) public pure returns (address[] memory) {
        address[] memory addr = new address[](numVals);
        address tmp;
        for (uint i = 0; i < numVals; i++) {
            assembly {
                // skip first 0x20 bytes (stored length of inst)
                tmp := mload(add(add(inst, 0x63), mul(i, 0x20))) // 67+i*32
            }
            addr[i] = tmp;
        }
        return addr;
    }

    function verifySig(
        address[] memory committee,
        bytes32 msgHash,
        uint8[] memory v,
        bytes32[] memory r,
        bytes32[] memory s
    ) public returns (bool) {
        // emit LogUint(committee.length);
        for (uint i = 0; i < v.length; i++){
            if (ecrecover(msgHash, v[i], r[i], s[i]) != committee[i]) {
                return false;
            }
        }
        return true;
    }
}
