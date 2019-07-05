
# External Contracts
contract Incognito_proxy:
    def parseSwapInst(inst: bytes[150]) -> (uint256, bytes32): constant
    def notifyPls(v: bytes32): modifying
    def instructionInMerkleTree(leaf: bytes32, root: bytes32, path: bytes32[8], left: bool[8], length: int128) -> bool: constant
    def pubkeyInMerkleTree(leaf: bytes32, root: bytes32, path: bytes32[3], left: bool[3], length: int128) -> bool: constant
    def verifyInst(commRoot: bytes32, instHash: bytes32, instPath: bytes32[8], instPathIsLeft: bool[8], instPathLen: int128, instRoot: bytes32, blkHash: bytes32, signerPubkeys: bytes[264], signerCount: int128, signerSig: bytes32, signerPaths: bytes32[24], signerPathIsLeft: bool[24], signerPathLen: int128) -> bool: constant
    def instructionApproved(instHash: bytes32, beaconInstPath: bytes32[8], beaconInstPathIsLeft: bool[8], beaconInstPathLen: int128, beaconInstRoot: bytes32, beaconBlkData: bytes32, beaconBlkHash: bytes32, beaconSignerPubkeys: bytes[264], beaconSignerCount: int128, beaconSignerSig: bytes32, beaconSignerPaths: bytes32[24], beaconSignerPathIsLeft: bool[24], beaconSignerPathLen: int128, bridgeInstPath: bytes32[8], bridgeInstPathIsLeft: bool[8], bridgeInstPathLen: int128, bridgeInstRoot: bytes32, bridgeBlkData: bytes32, bridgeBlkHash: bytes32, bridgeSignerPubkeys: bytes[264], bridgeSignerCount: int128, bridgeSignerSig: bytes32, bridgeSignerPaths: bytes32[24], bridgeSignerPathIsLeft: bool[24], bridgeSignerPathLen: int128) -> bool: constant
    def swapCommittee(inst: bytes[150], beaconInstPath: bytes32[8], beaconInstPathIsLeft: bool[8], beaconInstPathLen: int128, beaconInstRoot: bytes32, beaconBlkData: bytes32, beaconBlkHash: bytes32, beaconSignerPubkeys: bytes[264], beaconSignerCount: int128, beaconSignerSig: bytes32, beaconSignerPaths: bytes32[24], beaconSignerPathIsLeft: bool[24], beaconSignerPathLen: int128, bridgeInstPath: bytes32[8], bridgeInstPathIsLeft: bool[8], bridgeInstPathLen: int128, bridgeInstRoot: bytes32, bridgeBlkData: bytes32, bridgeBlkHash: bytes32, bridgeSignerPubkeys: bytes[264], bridgeSignerCount: int128, bridgeSignerSig: bytes32, bridgeSignerPaths: bytes32[24], bridgeSignerPathIsLeft: bool[24], bridgeSignerPathLen: int128) -> bool: modifying
    def beaconCommRoot() -> bytes32: constant
    def bridgeCommRoot() -> bytes32: constant


