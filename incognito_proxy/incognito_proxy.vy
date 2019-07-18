MAX_RANGE: constant(uint256) = 2 ** 255
INST_MAX_PATH: constant(uint256) = 8 # support up to 2 ** INST_MAX_PATH instructions per block

INST_LENGTH: constant(int128) = 300

COMM_SIZE: constant(uint256) = 8
PUBKEY_SIZE: constant(int128) = 33 # each pubkey is 33 bytes
PUBKEY_LENGTH: constant(int128) = INST_LENGTH # length of the array storing all pubkeys

MIN_SIGN: constant(uint256) = 3

struct Committee:
    Pubkeys: bytes[PUBKEY_LENGTH]
    PrevBlk: uint256

# External contract
contract MulSigP256:
    def checkMulSig(xPks: uint256[8], yPks: uint256[8], idxSigs: uint256[8], numSig: int128, xR: uint256, yR: uint256, bytesR: bytes[33], sig: uint256, mess: bytes32) -> bool: constant

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
mulsig: public(MulSigP256)

@public
def __init__(_beaconComm: bytes[PUBKEY_LENGTH], _bridgeComm: bytes[PUBKEY_LENGTH], _mulsig: address):
    self.beaconComm[0] = Committee({Pubkeys: _beaconComm, PrevBlk: 0})
    self.bridgeComm[0] = Committee({Pubkeys: _bridgeComm, PrevBlk: 0})
    self.mulsig = MulSigP256(_mulsig)

@constant
@public
def parseSwapInst(inst: bytes[INST_LENGTH], numPk: int128) -> (uint256, uint256, bytes[PUBKEY_LENGTH]):
    type: uint256 = convert(slice(inst, start=0, len=3), uint256)
    height: uint256 = extract32(inst, 3, type=uint256)
    pubkeys: bytes[PUBKEY_LENGTH] = slice(inst, start=35, len=numPk*PUBKEY_SIZE)
    return type, height, pubkeys

@constant
@public
def findComm(blkHeight: uint256, isBeacon: bool) -> bytes[PUBKEY_LENGTH]:
    comm: bytes[PUBKEY_LENGTH]
    if isBeacon:
        height: uint256 = self.latestBeaconBlk
        for i in range(MAX_RANGE):
            if height <= blkHeight:
                comm = self.beaconComm[height].Pubkeys
                break
            height = self.beaconComm[height].PrevBlk
    else:
        height: uint256 = self.latestBridgeBlk
        for i in range(MAX_RANGE):
            if height <= blkHeight:
                comm= self.bridgeComm[height].Pubkeys
                break
            height = self.bridgeComm[height].PrevBlk
    return comm

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
def verifyCompressPoint(
    pk: bytes[PUBKEY_SIZE],
    x: uint256,
    y: uint256,
) -> bool:
    pkx: uint256 = convert(slice(pk, start=1, len=32), uint256)
    if pkx != x:
        return False
    pkFormat: uint256 = convert(slice(pk, start=0, len=1), uint256)
    format: uint256 = 2
    if y % 2 == 1:
        format = 3
    if pkFormat != format:
        return False
    return True

@constant
@public
def verifySig(
    pubkey: bytes[PUBKEY_LENGTH],
    signerSig: uint256,
    numR: int128,
    xs: uint256[COMM_SIZE],
    ys: uint256[COMM_SIZE],
    rIdxs: int128[COMM_SIZE],
    numSig: int128,
    sigIdxs: uint256[COMM_SIZE],
    rx: uint256,
    ry: uint256,
    r: bytes[PUBKEY_SIZE],
    blk: bytes32,
) -> bool:
    # Check if enough validators signed this block
    if numSig < MIN_SIGN:
        return False

    # Check if decompressed xs, ys are correct
    for i in range(COMM_SIZE):
        if i >= numR:
            break
        idx: int128 = rIdxs[i]
        pk: bytes[PUBKEY_SIZE] = slice(pubkey, start=idx * PUBKEY_SIZE, len=PUBKEY_SIZE)
        if not self.verifyCompressPoint(pk, xs[i], ys[i]):
            return False

    # Check if signerSig is valid
    if not self.mulsig.checkMulSig(
        xs,
        ys,
        sigIdxs,
        numSig,
        rx,
        ry,
        r,
        signerSig,
        blk,
    ):
        return False

    return True

@constant
@public
def instructionApproved(
    isBeacon: bool,
    instHash: bytes32,
    height: uint256,
    instPath: bytes32[INST_MAX_PATH],
    instPathIsLeft: bool[INST_MAX_PATH],
    instPathLen: int128,
    instRoot: bytes32,
    blkData: bytes32,
    blkHash: bytes32,
    signerSig: uint256,
    numR: int128,
    xs: uint256[COMM_SIZE],
    ys: uint256[COMM_SIZE],
    rIdxs: int128[COMM_SIZE],
    numSig: int128,
    sigIdxs: uint256[COMM_SIZE],
    rx: uint256,
    ry: uint256,
    r: bytes[PUBKEY_SIZE],
) -> bool:
    # Find committees signed this block
    comm: bytes[PUBKEY_LENGTH] = self.findComm(height, isBeacon)

    # Check if instRoot is in block with hash blkHash
    blk: bytes32 = keccak256(concat(blkData, instRoot))
    if not blk == blkHash:
        return False

    # Check that beacon signature is correct
    if not self.verifySig(
        comm,
        signerSig,
        numR,
        xs,
        ys,
        rIdxs,
        numSig,
        sigIdxs,
        rx,
        ry,
        r,
        blk,
    ):
        return False

    # Check that inst is in beacon block
    if not self.instructionInMerkleTree(
        instHash,
        instRoot,
        instPath,
        instPathIsLeft,
        instPathLen,
    ):
        # log.NotifyString("failed verifying instruction")
        return False

    return True

@public
def swapBridgeCommittee(
    inst: bytes[INST_LENGTH], # content of swap instruction
    numPk: int128,
    beaconInstPath: bytes32[INST_MAX_PATH],
    beaconInstPathIsLeft: bool[INST_MAX_PATH],
    beaconInstPathLen: int128,
    beaconInstRoot: bytes32,
    beaconBlkData: bytes32, # hash of the rest of the beacon block
    beaconBlkHash: bytes32,
    beaconSignerSig: uint256, # aggregated signature
    beaconNumR: int128,
    beaconXs: uint256[COMM_SIZE], # decompressed x of pks who aggregated R
    beaconYs: uint256[COMM_SIZE],
    beaconRIdxs: int128[COMM_SIZE], # indices of members who aggregated R
    beaconNumSig: int128,
    beaconSigIdxs: uint256[COMM_SIZE], # indices of members who signed
    beaconRx: uint256,
    beaconRy: uint256,
    beaconR: bytes[PUBKEY_SIZE],
    bridgeInstPath: bytes32[INST_MAX_PATH],
    bridgeInstPathIsLeft: bool[INST_MAX_PATH],
    bridgeInstPathLen: int128,
    bridgeInstRoot: bytes32,
    bridgeBlkData: bytes32,
    bridgeBlkHash: bytes32,
    bridgeSignerSig: uint256,
    bridgeNumR: int128,
    bridgeXs: uint256[COMM_SIZE],
    bridgeYs: uint256[COMM_SIZE],
    bridgeRIdxs: int128[COMM_SIZE],
    bridgeNumSig: int128,
    bridgeSigIdxs: uint256[COMM_SIZE],
    bridgeRx: uint256,
    bridgeRy: uint256,
    bridgeR: bytes[PUBKEY_SIZE],
) -> bool:
    # Check if beaconInstRoot is in block with hash beaconBlkHash
    instHash: bytes32 = keccak256(inst)

    # Parse instruction and check metadata
    type: uint256
    startHeight: uint256
    pubkeys: bytes[PUBKEY_LENGTH]
    type, startHeight, pubkeys = self.parseSwapInst(inst, numPk)
    # log.NotifyBytes32(extract32(pubkeys, 0, type=bytes32))

    # Metadata type and shardID of swap bridge instruction
    assert type == 3617073

    # Verify instruction on beacon
    if not self.instructionApproved(
        True,
        instHash,
        self.latestBeaconBlk,
        beaconInstPath,
        beaconInstPathIsLeft,
        beaconInstPathLen,
        beaconInstRoot,
        beaconBlkData,
        beaconBlkHash,
        beaconSignerSig,
        beaconNumR,
        beaconXs,
        beaconYs,
        beaconRIdxs,
        beaconNumSig,
        beaconSigIdxs,
        beaconRx,
        beaconRy,
        beaconR,
    ):
        return False

    # Verify instruction on bridge
    if not self.instructionApproved(
        False,
        instHash,
        self.latestBridgeBlk,
        bridgeInstPath,
        bridgeInstPathIsLeft,
        bridgeInstPathLen,
        bridgeInstRoot,
        bridgeBlkData,
        bridgeBlkHash,
        bridgeSignerSig,
        bridgeNumR,
        bridgeXs,
        bridgeYs,
        bridgeRIdxs,
        bridgeNumSig,
        bridgeSigIdxs,
        bridgeRx,
        bridgeRy,
        bridgeR,
    ):
        return False

    # Swap committee
    self.bridgeComm[startHeight] = Committee({Pubkeys: pubkeys, PrevBlk: self.latestBridgeBlk})
    self.latestBridgeBlk = startHeight
    log.NotifyString("updated bridge committee")
    return True

@public
def swapBeaconCommittee(
    inst: bytes[INST_LENGTH], # content of swap instruction
    numPk: int128,
    beaconInstPath: bytes32[INST_MAX_PATH],
    beaconInstPathIsLeft: bool[INST_MAX_PATH],
    beaconInstPathLen: int128,
    beaconInstRoot: bytes32,
    beaconBlkData: bytes32, # hash of the rest of the beacon block
    beaconBlkHash: bytes32,
    beaconSignerSig: uint256, # aggregated signature
    beaconNumR: int128,
    beaconXs: uint256[COMM_SIZE], # decompressed x of pks who aggregated R
    beaconYs: uint256[COMM_SIZE],
    beaconRIdxs: int128[COMM_SIZE], # indices of members who aggregated R
    beaconNumSig: int128,
    beaconSigIdxs: uint256[COMM_SIZE], # indices of members who signed
    beaconRx: uint256,
    beaconRy: uint256,
    beaconR: bytes[PUBKEY_SIZE],
) -> bool:
    # Check if beaconInstRoot is in block with hash beaconBlkHash
    instHash: bytes32 = keccak256(inst)

    # Parse instruction and check metadata
    type: uint256
    startHeight: uint256
    pubkeys: bytes[PUBKEY_LENGTH]
    type, startHeight, pubkeys = self.parseSwapInst(inst, numPk)
    # log.NotifyBytes32(extract32(pubkeys, 0, type=bytes32))

    # Metadata type and shardID of swap bridge instruction
    assert type == 3616817

    # Verify instruction on beacon
    if not self.instructionApproved(
        True,
        instHash,
        self.latestBeaconBlk,
        beaconInstPath,
        beaconInstPathIsLeft,
        beaconInstPathLen,
        beaconInstRoot,
        beaconBlkData,
        beaconBlkHash,
        beaconSignerSig,
        beaconNumR,
        beaconXs,
        beaconYs,
        beaconRIdxs,
        beaconNumSig,
        beaconSigIdxs,
        beaconRx,
        beaconRy,
        beaconR,
    ):
        return False

    # Swap committee
    self.beaconComm[startHeight] = Committee({Pubkeys: pubkeys, PrevBlk: self.latestBeaconBlk})
    self.latestBeaconBlk = startHeight
    log.NotifyString("updated beacon committee")

    log.NotifyString("no exeception...")
    return True

