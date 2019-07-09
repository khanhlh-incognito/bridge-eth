
# External Contracts
contract Incognito_proxy:
    def parseSwapInst(inst: bytes[300], numPk: int128) -> (uint256, uint256, bytes[300]): constant
    def findComm(beaconHeight: uint256, bridgeHeight: uint256) -> (bytes[300], bytes[300]): constant
    def notifyPls(v: bytes32): modifying
    def instructionInMerkleTree(leaf: bytes32, root: bytes32, path: bytes32[8], left: bool[8], length: int128) -> bool: constant
    def verifyInst(pubkey: bytes[300], instHash: bytes32, instPath: bytes32[8], instPathIsLeft: bool[8], instPathLen: int128, instRoot: bytes32, blkHash: bytes32, signerSig: bytes32) -> bool: constant
    def instructionApproved(beaconInstHash: bytes32, beaconHeight: uint256, beaconInstPath: bytes32[8], beaconInstPathIsLeft: bool[8], beaconInstPathLen: int128, beaconInstRoot: bytes32, beaconBlkData: bytes32, beaconBlkHash: bytes32, beaconSignerSig: bytes32, bridgeInstHash: bytes32, bridgeHeight: uint256, bridgeInstPath: bytes32[8], bridgeInstPathIsLeft: bool[8], bridgeInstPathLen: int128, bridgeInstRoot: bytes32, bridgeBlkData: bytes32, bridgeBlkHash: bytes32, bridgeSignerSig: bytes32) -> bool: constant
    def swapCommittee(inst: bytes[300], numPk: int128, beaconInstPath: bytes32[8], beaconInstPathIsLeft: bool[8], beaconInstPathLen: int128, beaconInstRoot: bytes32, beaconBlkData: bytes32, beaconBlkHash: bytes32, beaconSignerSig: bytes32, bridgeInstPath: bytes32[8], bridgeInstPathIsLeft: bool[8], bridgeInstPathLen: int128, bridgeInstRoot: bytes32, bridgeBlkData: bytes32, bridgeBlkHash: bytes32, bridgeSignerSig: bytes32) -> bool: modifying
    def beaconComm__Pubkeys(arg0: uint256) -> bytes[300]: constant
    def beaconComm__PrevBlk(arg0: uint256) -> uint256: constant
    def bridgeComm__Pubkeys(arg0: uint256) -> bytes[300]: constant
    def bridgeComm__PrevBlk(arg0: uint256) -> uint256: constant
    def latestBeaconBlk() -> uint256: constant
    def latestBridgeBlk() -> uint256: constant


