MAX_RANGE: constant(uint256) = 2 ** 255
INST_MAX_PATH: constant(uint256) = 8 # support up to 2 ** INST_MAX_PATH instructions per block

INST_LENGTH: constant(int128) = 300

COMM_SIZE: constant(uint256) = 10

struct Committee:
    Pubkeys: address[COMM_SIZE]
    PrevBlk: uint256
    NumVals: uint256

# External contract
contract ECDSA:
    def verify(numSig: int128, committee: address[COMM_SIZE], msgHash: bytes32, v: uint256[COMM_SIZE], r: bytes32[COMM_SIZE], s: bytes32[COMM_SIZE]) -> bool: constant

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
ecdsa: public(ECDSA)

@public
def __init__(
    numBeaconVals: uint256,
    _beaconComm: address[COMM_SIZE],
    numBridgeVals: uint256,
    _bridgeComm: address[COMM_SIZE],
    _ecdsa: address
):
    self.beaconComm[0] = Committee({Pubkeys: _beaconComm, PrevBlk: 0, NumVals: numBeaconVals})
    self.bridgeComm[0] = Committee({Pubkeys: _bridgeComm, PrevBlk: 0, NumVals: numBridgeVals})
    self.ecdsa = ECDSA(_ecdsa)

@constant
@public
def extractMetaFromInst(inst: bytes[INST_LENGTH], numPk: int128) -> (uint256, uint256, uint256):
    type: uint256 = convert(slice(inst, start=0, len=3), uint256)
    height: uint256 = extract32(inst, 3, type=uint256)
    numVals: uint256 = extract32(inst, 35, type=uint256)
    return type, height, numVals

@constant
@public
def extractCommitteeFromInst(inst: bytes[INST_LENGTH], numVals: uint256) -> address[COMM_SIZE]:
    pubkeys: address[COMM_SIZE]
    for i in range(COMM_SIZE):
        if i >= convert(numVals, int128):
            break
        pubkeys[i] = extract32(inst, 67+i*32, type=address)
    return pubkeys

@constant
@public
def findCommitteeFromHeight(blkHeight: uint256, isBeacon: bool) -> uint256:
    height: uint256 = self.latestBeaconBlk
    if isBeacon:
        for i in range(MAX_RANGE):
            if height <= blkHeight:
                break
            height = self.beaconComm[height].PrevBlk
    else:
        height = self.latestBridgeBlk
        for i in range(MAX_RANGE):
            if height <= blkHeight:
                break
            height = self.bridgeComm[height].PrevBlk
    return height

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
def verifySig(
    numSig: int128,
    signers: address[COMM_SIZE],
    v: uint256[COMM_SIZE],
    r: bytes32[COMM_SIZE],
    s: bytes32[COMM_SIZE],
    blk: bytes32,
) -> bool:
    # NOTE: comment out checkMulSig to increase #validators of testnet
    # Check if signerSig is valid
    msgHash: bytes32 = 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8

    if not self.ecdsa.verify(
        numSig,
        signers,
        msgHash,
        v,
        r,
        s
    ):
        return False

    return True

@constant
@public
def instructionApproved(
    isBeacon: bool,
    instHash: bytes32,
    blkHeight: uint256,
    instPath: bytes32[INST_MAX_PATH],
    instPathIsLeft: bool[INST_MAX_PATH],
    instPathLen: int128,
    instRoot: bytes32,
    blkData: bytes32,
    numSig: int128,
    sigIdxs: uint256[COMM_SIZE],
    v: uint256[COMM_SIZE],
    r: bytes32[COMM_SIZE],
    s: bytes32[COMM_SIZE],
) -> bool:
    # Find committee in charge of this block
    commHeight: uint256 = self.findCommitteeFromHeight(blkHeight, isBeacon)
    comm: address[COMM_SIZE]
    numVals: uint256
    if isBeacon:
        comm = self.beaconComm[commHeight].Pubkeys
        numVals = self.beaconComm[commHeight].NumVals
    else:
        comm = self.bridgeComm[commHeight].Pubkeys
        numVals = self.bridgeComm[commHeight].NumVals

    # Get pubkey of signers
    signers: address[COMM_SIZE]
    for i in range(COMM_SIZE):
        if i >= numSig:
            break
        signers[i] = comm[sigIdxs[i]]

    # Get block hash from instRoot and other data
    blk: bytes32 = keccak256(concat(blkData, instRoot))

    # Check if enough validators signed this block
    if convert(numSig, uint256) < 1 + numVals * 2 / 3:
        return False

    # Check that beacon signature is correct
    if not self.verifySig(
        numSig,
        signers,
        v,
        r,
        s,
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
    beaconNumSig: int128,
    beaconSigIdxs: uint256[COMM_SIZE], # indices of members who signed
    beaconSigVs: uint256[COMM_SIZE],
    beaconSigRs: bytes32[COMM_SIZE],
    beaconSigSs: bytes32[COMM_SIZE],
    bridgeInstPath: bytes32[INST_MAX_PATH],
    bridgeInstPathIsLeft: bool[INST_MAX_PATH],
    bridgeInstPathLen: int128,
    bridgeInstRoot: bytes32,
    bridgeBlkData: bytes32, # hash of the rest of the beacon block
    bridgeNumSig: int128,
    bridgeSigIdxs: uint256[COMM_SIZE], # indices of members who signed
    bridgeSigVs: uint256[COMM_SIZE],
    bridgeSigRs: bytes32[COMM_SIZE],
    bridgeSigSs: bytes32[COMM_SIZE],
) -> bool:
    # Check if beaconInstRoot is in block
    instHash: bytes32 = keccak256(inst)

    # Parse instruction and check metadata
    type: uint256
    startHeight: uint256
    numVals: uint256
    type, startHeight, numVals = self.extractMetaFromInst(inst, numPk)
    pubkeys: address[COMM_SIZE] = self.extractCommitteeFromInst(inst, numVals)
    # log.NotifyBytes32(extract32(pubkeys, 0, type=bytes32))
    # log.NotifyUint256(numVals)

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
        beaconNumSig,
        beaconSigIdxs,
        beaconSigVs,
        beaconSigRs,
        beaconSigSs,
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
        bridgeNumSig,
        bridgeSigIdxs,
        bridgeSigVs,
        bridgeSigRs,
        bridgeSigSs,
    ):
        return False

    # Swap committee
    self.bridgeComm[startHeight] = Committee({Pubkeys: pubkeys, PrevBlk: self.latestBridgeBlk, NumVals: numVals})
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
    beaconNumSig: int128,
    beaconSigIdxs: uint256[COMM_SIZE], # indices of members who signed
    beaconSigVs: uint256[COMM_SIZE],
    beaconSigRs: bytes32[COMM_SIZE],
    beaconSigSs: bytes32[COMM_SIZE],
) -> bool:
    # Check if beaconInstRoot is in block
    instHash: bytes32 = keccak256(inst)

    # Parse instruction and check metadata
    type: uint256
    startHeight: uint256
    numVals: uint256
    type, startHeight, numVals = self.extractMetaFromInst(inst, numPk)
    pubkeys: address[COMM_SIZE] = self.extractCommitteeFromInst(inst, numVals)
    # log.NotifyBytes32(extract32(pubkeys, 0, type=bytes32))

    # Metadata type and shardID of swap beacon instruction
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
        beaconNumSig,
        beaconSigIdxs,
        beaconSigVs,
        beaconSigRs,
        beaconSigSs,
    ):
        return False

    # Swap committee
    self.beaconComm[startHeight] = Committee({Pubkeys: pubkeys, PrevBlk: self.latestBeaconBlk, NumVals: numVals})
    self.latestBeaconBlk = startHeight
    log.NotifyString("updated beacon committee")
    return True

