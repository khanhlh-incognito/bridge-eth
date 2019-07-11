
# External Contracts
contract Incognito_proxy:
    def parseSwapInst(inst: bytes[300], numPk: int128) -> (uint256, uint256, bytes[300]): constant
    def findComm(beaconHeight: uint256, bridgeHeight: uint256) -> (bytes[300], bytes[300]): constant
    def notifyPls(v: bytes32): modifying
    def instructionInMerkleTree(leaf: bytes32, root: bytes32, path: bytes32[8], left: bool[8], length: int128) -> bool: constant
    def verifyCompressPoint(pk: bytes[33], x: uint256, y: uint256) -> bool: constant
    def verifySig(pubkey: bytes[300], signerSig: uint256, numR: int128, xs: uint256[8], ys: uint256[8], rIdxs: int128[8], numSig: int128, sigIdxs: uint256[8], rx: uint256, ry: uint256, r: bytes[33], blk: bytes32) -> bool: constant
    def instructionApproved(beaconInstHash: bytes32, beaconHeight: uint256, beaconInstPath: bytes32[8], beaconInstPathIsLeft: bool[8], beaconInstPathLen: int128, beaconInstRoot: bytes32, beaconBlkData: bytes32, beaconBlkHash: bytes32, beaconSignerSig: uint256, beaconNumR: int128, beaconXs: uint256[8], beaconYs: uint256[8], beaconRIdxs: int128[8], beaconNumSig: int128, beaconSigIdxs: uint256[8], beaconRx: uint256, beaconRy: uint256, beaconR: bytes[33], bridgeInstHash: bytes32, bridgeHeight: uint256, bridgeInstPath: bytes32[8], bridgeInstPathIsLeft: bool[8], bridgeInstPathLen: int128, bridgeInstRoot: bytes32, bridgeBlkData: bytes32, bridgeBlkHash: bytes32, bridgeSignerSig: uint256, bridgeNumR: int128, bridgeXs: uint256[8], bridgeYs: uint256[8], bridgeRIdxs: int128[8], bridgeNumSig: int128, bridgeSigIdxs: uint256[8], bridgeRx: uint256, bridgeRy: uint256, bridgeR: bytes[33]) -> bool: constant
    def swapCommittee(inst: bytes[300], numPk: int128, beaconInstPath: bytes32[8], beaconInstPathIsLeft: bool[8], beaconInstPathLen: int128, beaconInstRoot: bytes32, beaconBlkData: bytes32, beaconBlkHash: bytes32, beaconSignerSig: uint256, beaconNumR: int128, beaconXs: uint256[8], beaconYs: uint256[8], beaconRIdxs: int128[8], beaconNumSig: int128, beaconSigIdxs: uint256[8], beaconRx: uint256, beaconRy: uint256, beaconR: bytes[33], bridgeInstPath: bytes32[8], bridgeInstPathIsLeft: bool[8], bridgeInstPathLen: int128, bridgeInstRoot: bytes32, bridgeBlkData: bytes32, bridgeBlkHash: bytes32, bridgeSignerSig: uint256, bridgeNumR: int128, bridgeXs: uint256[8], bridgeYs: uint256[8], bridgeRIdxs: int128[8], bridgeNumSig: int128, bridgeSigIdxs: uint256[8], bridgeRx: uint256, bridgeRy: uint256, bridgeR: bytes[33]) -> bool: modifying
    def beaconComm__Pubkeys(arg0: uint256) -> bytes[300]: constant
    def beaconComm__PrevBlk(arg0: uint256) -> uint256: constant
    def bridgeComm__Pubkeys(arg0: uint256) -> bytes[300]: constant
    def bridgeComm__PrevBlk(arg0: uint256) -> uint256: constant
    def latestBeaconBlk() -> uint256: constant
    def latestBridgeBlk() -> uint256: constant
    def mulsig() -> address(MulSigP256): constant


