MAX_RANGE: constant(uint256) = 2 ** 255

INST_MAX_PATH: constant(uint256) = 8 # support up to 2 ** INST_MAX_PATH instructions per block

PUBKEY_MAX_PATH: constant(uint256) = 3 # up to 2 ** PUBKEY_MAX_PATH signers
COMM_SIZE: constant(uint256) = 2 ** PUBKEY_MAX_PATH

PUBKEY_NODE: constant(uint256) = COMM_SIZE * PUBKEY_MAX_PATH # number of merkle proof nodes to prove all pubkeys
PUBKEY_SIZE: constant(int128) = 33 # each pubkey is 33 bytes
PUBKEY_LENGTH: constant(int128) = PUBKEY_SIZE * COMM_SIZE # length of the array storing all pubkeys

INST_LENGTH: constant(uint256) = 150

# TODO: update to 2/3+1
MIN_SIGN: constant(uint256) = 2

struct Committee:
    Pubkeys: bytes32
    PrevBlk: uint256

SwapBeaconCommittee: event({newCommitteeRoot: bytes32})
SwapBridgeCommittee: event({newCommitteeRoot: bytes32})

# Debug events
NotifyString: event({content: string[120]})
NotifyBytes32: event({content: bytes32})
NotifyBool: event({content: bool})
NotifyUint256: event({content: uint256})

beaconCommRoot: public(map(uint256, Committee))
bridgeCommRoot: public(map(uint256, Committee))
latestBeaconBlk: public(uint256)
latestBridgeBlk: public(uint256)

@public
def __init__(_beaconCommRoot: bytes32, _bridgeCommRoot: bytes32):
    self.beaconCommRoot[0] = Committee({Pubkeys: _beaconCommRoot, PrevBlk: 0})
    self.bridgeCommRoot[0] = Committee({Pubkeys: _bridgeCommRoot, PrevBlk: 0})

@constant
@public
def parseSwapInst(inst: bytes[INST_LENGTH]) -> (uint256, uint256, bytes32):
    type: uint256 = convert(slice(inst, start=0, len=3), uint256)
    height: uint256 = extract32(inst, 3, type=uint256)
    newCommRoot: bytes32 = convert(slice(inst, start=35, len=32), bytes32)
    return type, height, newCommRoot

@constant
@public
def findCommRoot(beaconHeight: uint256, bridgeHeight: uint256) -> (bytes32, bytes32):
    beacon: bytes32
    height: uint256 = self.latestBeaconBlk
    for i in range(MAX_RANGE):
        if height <= beaconHeight:
            beacon = self.beaconCommRoot[height].Pubkeys
            break
        height = self.beaconCommRoot[height].PrevBlk

    bridge: bytes32
    height = self.latestBridgeBlk
    for i in range(MAX_RANGE):
        if height <= bridgeHeight:
            bridge = self.bridgeCommRoot[height].Pubkeys
            break
        height = self.bridgeCommRoot[height].PrevBlk
    return beacon, bridge

@public
def notifyPls(v: bytes32):
    a: uint256 = 135790246810123
    b: uint256 = convert(v, uint256)
    log.NotifyBytes32(convert(a, bytes32))
    log.NotifyBytes32(v)
    log.NotifyBool(b == a)

@constant
@public
def instructionInMerkleTree(
    leaf: bytes32,
    root: bytes32,
    path: bytes32[INST_MAX_PATH],
    left: bool[INST_MAX_PATH],
    length: int128
) -> bool:
    hash: bytes32 = leaf
    for i in range(INST_MAX_PATH):
        if i >= length:
            break
        if left[i]:
            hash = keccak256(concat(path[i], hash))
        elif convert(path[i], uint256) == 0:
            hash = keccak256(concat(hash, hash))
        else:
            hash = keccak256(concat(hash, path[i]))
    return hash == root

@constant
@public
def pubkeyInMerkleTree(
    leaf: bytes32,
    root: bytes32,
    path: bytes32[PUBKEY_MAX_PATH],
    left: bool[PUBKEY_MAX_PATH],
    length: int128
) -> bool:
    hash: bytes32 = leaf
    for i in range(PUBKEY_MAX_PATH):
        if i >= length:
            break
        if left[i]:
            hash = keccak256(concat(path[i], hash))
        elif convert(path[i], uint256) == 0:
            hash = keccak256(concat(hash, hash))
        else:
            hash = keccak256(concat(hash, path[i]))
    return hash == root

@constant
@public
def verifyInst(
    commRoot: bytes32,
    instHash: bytes32,
    instPath: bytes32[INST_MAX_PATH],
    instPathIsLeft: bool[INST_MAX_PATH],
    instPathLen: int128,
    instRoot: bytes32,
    blkHash: bytes32,
    signerPubkeys: bytes[PUBKEY_LENGTH],
    signerCount: int128,
    signerSig: bytes32,
    signerPaths: bytes32[PUBKEY_NODE],
    signerPathIsLeft: bool[PUBKEY_NODE],
    signerPathLen: int128
) -> bool:
    # Check if enough validators signed this block
    if signerCount < MIN_SIGN:
        # log.NotifyString("not enough sig")
        return False

    # Check if inst is in merkle tree with root instRoot
    if not self.instructionInMerkleTree(instHash, instRoot, instPath, instPathIsLeft, instPathLen):
        # log.NotifyString("instruction is not in merkle tree")
        # log.NotifyBytes32(instHash)
        # log.NotifyBytes32(instRoot)
        return False

    # TODO: Check if signerSig is valid

    # Check if signerPubkeys are in merkle tree with root commRoot
    for i in range(COMM_SIZE):
        if i >= signerCount:
            break

        # Get hash of the pubkey
        signerPubkeyHash: bytes32 = keccak256(slice(signerPubkeys, start=i * PUBKEY_SIZE, len=PUBKEY_SIZE))

        path: bytes32[PUBKEY_MAX_PATH]
        left: bool[PUBKEY_MAX_PATH]
        for j in range(PUBKEY_MAX_PATH):
            if j >= signerPathLen:
                break
            path[j] = signerPaths[i * signerPathLen + j]
            left[j] = signerPathIsLeft[i * signerPathLen + j]

        # log.NotifyBytes32(signerPubkeyHash)
        if not self.pubkeyInMerkleTree(signerPubkeyHash, commRoot, path, left, signerPathLen):
            # log.NotifyString("pubkey not in merkle tree")
            return False

    return True

@constant
@public
def instructionApproved(
    beaconInstHash: bytes32,
    beaconHeight: uint256,
    beaconInstPath: bytes32[INST_MAX_PATH],
    beaconInstPathIsLeft: bool[INST_MAX_PATH],
    beaconInstPathLen: int128,
    beaconInstRoot: bytes32,
    beaconBlkData: bytes32,
    beaconBlkHash: bytes32,
    beaconSignerPubkeys: bytes[PUBKEY_LENGTH],
    beaconSignerCount: int128,
    beaconSignerSig: bytes32,
    beaconSignerPaths: bytes32[PUBKEY_NODE],
    beaconSignerPathIsLeft: bool[PUBKEY_NODE],
    beaconSignerPathLen: int128,
    bridgeInstHash: bytes32,
    bridgeHeight: uint256,
    bridgeInstPath: bytes32[INST_MAX_PATH],
    bridgeInstPathIsLeft: bool[INST_MAX_PATH],
    bridgeInstPathLen: int128,
    bridgeInstRoot: bytes32,
    bridgeBlkData: bytes32,
    bridgeBlkHash: bytes32,
    bridgeSignerPubkeys: bytes[PUBKEY_LENGTH],
    bridgeSignerCount: int128,
    bridgeSignerSig: bytes32,
    bridgeSignerPaths: bytes32[PUBKEY_NODE],
    bridgeSignerPathIsLeft: bool[PUBKEY_NODE],
    bridgeSignerPathLen: int128
) -> bool:
    blk: bytes32 = keccak256(concat(beaconBlkData, beaconInstRoot))
    if not blk == beaconBlkHash:
        # log.NotifyString("instruction merkle root is not in beacon block")
        # log.NotifyBytes32(beaconInstRoot)
        # log.NotifyBytes32(beaconBlkData)
        # log.NotifyBytes32(beaconInstHash)
        # log.NotifyBytes32(blk)
        return False

    # Find committees signed this instruction
    beacon: bytes32
    bridge: bytes32
    beacon, bridge = self.findCommRoot(beaconHeight, bridgeHeight)
    # log.NotifyBytes32(beacon)
    # log.NotifyBytes32(bridge)

    # Check that inst is in beacon block
    if not self.verifyInst(
        beacon,
        beaconInstHash,
        beaconInstPath,
        beaconInstPathIsLeft,
        beaconInstPathLen,
        beaconInstRoot,
        beaconBlkHash,
        beaconSignerPubkeys,
        beaconSignerCount,
        beaconSignerSig,
        beaconSignerPaths,
        beaconSignerPathIsLeft,
        beaconSignerPathLen
    ):
        # log.NotifyString("failed verifying beacon instruction")
        return False

    # Check if bridgeInstRoot is in block with hash bridgeBlkHash
    blk = keccak256(concat(bridgeBlkData, bridgeInstRoot))
    if not blk == bridgeBlkHash:
        # log.NotifyString("instruction merkle root is not in bridge block")
        return False

    # Check that inst is in bridge block
    if not self.verifyInst(
        bridge,
        bridgeInstHash,
        bridgeInstPath,
        bridgeInstPathIsLeft,
        bridgeInstPathLen,
        bridgeInstRoot,
        bridgeBlkHash,
        bridgeSignerPubkeys,
        bridgeSignerCount,
        bridgeSignerSig,
        bridgeSignerPaths,
        bridgeSignerPathIsLeft,
        bridgeSignerPathLen
    ):
        # log.NotifyString("failed verify bridge instruction")
        return False

    return True

@public
def swapCommittee(
    inst: bytes[INST_LENGTH], # content of swap instruction
    beaconInstPath: bytes32[INST_MAX_PATH],
    beaconInstPathIsLeft: bool[INST_MAX_PATH],
    beaconInstPathLen: int128,
    beaconInstRoot: bytes32,
    beaconBlkData: bytes32, # hash of the rest of the beacon block
    beaconBlkHash: bytes32,
    beaconSignerPubkeys: bytes[PUBKEY_LENGTH],
    beaconSignerCount: int128,
    beaconSignerSig: bytes32, # aggregated signature of some committee members
    beaconSignerPaths: bytes32[PUBKEY_NODE],
    beaconSignerPathIsLeft: bool[PUBKEY_NODE],
    beaconSignerPathLen: int128,
    bridgeInstPath: bytes32[INST_MAX_PATH],
    bridgeInstPathIsLeft: bool[INST_MAX_PATH],
    bridgeInstPathLen: int128,
    bridgeInstRoot: bytes32,
    bridgeBlkData: bytes32, # hash of the rest of the bridge block
    bridgeBlkHash: bytes32,
    bridgeSignerPubkeys: bytes[PUBKEY_LENGTH],
    bridgeSignerCount: int128,
    bridgeSignerSig: bytes32,
    bridgeSignerPaths: bytes32[PUBKEY_NODE],
    bridgeSignerPathIsLeft: bool[PUBKEY_NODE],
    bridgeSignerPathLen: int128
) -> bool:
    # Check if beaconInstRoot is in block with hash beaconBlkHash
    instHash: bytes32 = keccak256(inst)

    # Parse instruction and check metadata
    type: uint256
    startHeight: uint256
    newCommRoot: bytes32
    type, startHeight, newCommRoot = self.parseSwapInst(inst)
    log.NotifyUint256(startHeight)
    if type == 3616817: # Metadata type and shardID of swap beacon
        assert startHeight > self.latestBeaconBlk
    elif type == 3617073:
        assert startHeight > self.latestBridgeBlk
    else:
        assert False

    # Only swap if instruction is approved
    if not self.instructionApproved(
        instHash,
        self.latestBeaconBlk,
        beaconInstPath,
        beaconInstPathIsLeft,
        beaconInstPathLen,
        beaconInstRoot,
        beaconBlkData,
        beaconBlkHash,
        beaconSignerPubkeys,
        beaconSignerCount,
        beaconSignerSig,
        beaconSignerPaths,
        beaconSignerPathIsLeft,
        beaconSignerPathLen,
        instHash,
        self.latestBridgeBlk,
        bridgeInstPath,
        bridgeInstPathIsLeft,
        bridgeInstPathLen,
        bridgeInstRoot,
        bridgeBlkData,
        bridgeBlkHash,
        bridgeSignerPubkeys,
        bridgeSignerCount,
        bridgeSignerSig,
        bridgeSignerPaths,
        bridgeSignerPathIsLeft,
        bridgeSignerPathLen
    ):
        return False

    # Swap committee
    if type == 3616817: # Metadata type and shardID of swap beacon
        self.beaconCommRoot[startHeight] = Committee({Pubkeys: newCommRoot, PrevBlk: self.latestBeaconBlk})
        self.latestBeaconBlk = startHeight
        log.NotifyString("updated beacon committee")
        log.SwapBeaconCommittee(newCommRoot)
    elif type == 3617073:
        self.bridgeCommRoot[startHeight] = Committee({Pubkeys: newCommRoot, PrevBlk: self.latestBridgeBlk})
        self.latestBridgeBlk = startHeight
        log.NotifyString("updated bridge committee")
        log.SwapBridgeCommittee(newCommRoot)

    # # log.NotifyBytes32(newCommRoot)
    log.NotifyString("no exeception...")
    return True

