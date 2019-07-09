MAX_RANGE: constant(uint256) = 2 ** 255
INST_MAX_PATH: constant(uint256) = 8 # support up to 2 ** INST_MAX_PATH instructions per block

INST_LENGTH: constant(int128) = 300

COMM_SIZE: constant(uint256) = 8
PUBKEY_SIZE: constant(int128) = 33 # each pubkey is 33 bytes
PUBKEY_LENGTH: constant(int128) = INST_LENGTH # length of the array storing all pubkeys

# TODO: update to 2/3+1
MIN_SIGN: constant(uint256) = 2

struct Committee:
    Pubkeys: bytes[PUBKEY_LENGTH]
    PrevBlk: uint256

SwapBeaconCommittee: event({newCommitteeRoot: bytes32})
SwapBridgeCommittee: event({newCommitteeRoot: bytes32})

# Debug events
NotifyString: event({content: string[120]})
NotifyBytes32: event({content: bytes32})
NotifyBool: event({content: bool})
NotifyUint256: event({content: uint256})
NotifyPubkey: event({content: bytes[33]})

beaconComm: public(map(uint256, Committee))
bridgeComm: public(map(uint256, Committee))
latestBeaconBlk: public(uint256)
latestBridgeBlk: public(uint256)

@public
def __init__(_beaconComm: bytes[PUBKEY_LENGTH], _bridgeComm: bytes[PUBKEY_LENGTH]):
    self.beaconComm[0] = Committee({Pubkeys: _beaconComm, PrevBlk: 0})
    self.bridgeComm[0] = Committee({Pubkeys: _bridgeComm, PrevBlk: 0})

@constant
@public
def parseSwapInst(inst: bytes[INST_LENGTH], numPk: int128) -> (uint256, uint256, bytes[PUBKEY_LENGTH]):
    type: uint256 = convert(slice(inst, start=0, len=3), uint256)
    height: uint256 = extract32(inst, 3, type=uint256)
    pubkeys: bytes[PUBKEY_LENGTH] = slice(inst, start=35, len=numPk*PUBKEY_SIZE)
    return type, height, pubkeys

@constant
@public
def findComm(beaconHeight: uint256, bridgeHeight: uint256) -> (bytes[PUBKEY_LENGTH], bytes[PUBKEY_LENGTH]):
    beacon: bytes[PUBKEY_LENGTH]
    height: uint256 = self.latestBeaconBlk
    for i in range(MAX_RANGE):
        if height <= beaconHeight:
            beacon = self.beaconComm[height].Pubkeys
            break
        height = self.beaconComm[height].PrevBlk

    bridge: bytes[PUBKEY_LENGTH]
    height = self.latestBridgeBlk
    for i in range(MAX_RANGE):
        if height <= bridgeHeight:
            bridge = self.bridgeComm[height].Pubkeys
            break
        height = self.bridgeComm[height].PrevBlk
    return beacon, bridge

# TODO: remove test func
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
def verifyInst(
    pubkey: bytes[PUBKEY_LENGTH],
    instHash: bytes32,
    instPath: bytes32[INST_MAX_PATH],
    instPathIsLeft: bool[INST_MAX_PATH],
    instPathLen: int128,
    instRoot: bytes32,
    blkHash: bytes32,
    signerSig: bytes32,
) -> bool:
    # Check if inst is in merkle tree with root instRoot
    if not self.instructionInMerkleTree(instHash, instRoot, instPath, instPathIsLeft, instPathLen):
        # log.NotifyString("instruction is not in merkle tree")
        # log.NotifyBytes32(instHash)
        # log.NotifyBytes32(instRoot)
        return False

    # TODO: Check if signerSig is valid
    # Check if enough validators signed this block

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
    beaconSignerSig: bytes32,
    bridgeInstHash: bytes32,
    bridgeHeight: uint256,
    bridgeInstPath: bytes32[INST_MAX_PATH],
    bridgeInstPathIsLeft: bool[INST_MAX_PATH],
    bridgeInstPathLen: int128,
    bridgeInstRoot: bytes32,
    bridgeBlkData: bytes32,
    bridgeBlkHash: bytes32,
    bridgeSignerSig: bytes32,
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
    beacon: bytes[PUBKEY_LENGTH]
    bridge: bytes[PUBKEY_LENGTH]
    beacon, bridge = self.findComm(beaconHeight, bridgeHeight)
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
        beaconSignerSig,
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
        bridgeSignerSig,
    ):
        # log.NotifyString("failed verify bridge instruction")
        return False

    return True

@public
def swapCommittee(
    inst: bytes[INST_LENGTH], # content of swap instruction
    numPk: int128,
    beaconInstPath: bytes32[INST_MAX_PATH],
    beaconInstPathIsLeft: bool[INST_MAX_PATH],
    beaconInstPathLen: int128,
    beaconInstRoot: bytes32,
    beaconBlkData: bytes32, # hash of the rest of the beacon block
    beaconBlkHash: bytes32,
    beaconSignerSig: bytes32, # aggregated signature of some committee members
    bridgeInstPath: bytes32[INST_MAX_PATH],
    bridgeInstPathIsLeft: bool[INST_MAX_PATH],
    bridgeInstPathLen: int128,
    bridgeInstRoot: bytes32,
    bridgeBlkData: bytes32, # hash of the rest of the bridge block
    bridgeBlkHash: bytes32,
    bridgeSignerSig: bytes32,
) -> bool:
    # Check if beaconInstRoot is in block with hash beaconBlkHash
    instHash: bytes32 = keccak256(inst)

    # Parse instruction and check metadata
    type: uint256
    startHeight: uint256
    pubkeys: bytes[PUBKEY_LENGTH]
    type, startHeight, pubkeys = self.parseSwapInst(inst, numPk)
    # log.NotifyBytes32(extract32(pubkeys, 0, type=bytes32))

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
        beaconSignerSig,
        instHash,
        self.latestBridgeBlk,
        bridgeInstPath,
        bridgeInstPathIsLeft,
        bridgeInstPathLen,
        bridgeInstRoot,
        bridgeBlkData,
        bridgeBlkHash,
        bridgeSignerSig,
    ):
        return False

    # Swap committee
    if type == 3616817: # Metadata type and shardID of swap beacon
        self.beaconComm[startHeight] = Committee({Pubkeys: pubkeys, PrevBlk: self.latestBeaconBlk})
        self.latestBeaconBlk = startHeight
        log.NotifyString("updated beacon committee")
        # log.SwapBeaconCommittee(newCommRoot)
    elif type == 3617073:
        self.bridgeComm[startHeight] = Committee({Pubkeys: pubkeys, PrevBlk: self.latestBridgeBlk})
        self.latestBridgeBlk = startHeight
        log.NotifyString("updated bridge committee")
        # log.SwapBridgeCommittee(newCommRoot)

    # # # log.NotifyBytes32(newCommRoot)
    log.NotifyString("no exeception...")
    return True

