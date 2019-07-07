
# External Contracts
contract Vault:
    def deposit(incognito_address: string[128]): modifying
    def depositERC20(token: address, amount: uint256, incognito_address: string[128]): modifying
    def parseBurnInst(inst: bytes[150]) -> (uint256, address, address, uint256): constant
    def testExtract(a: bytes[150]) -> (address, uint256(wei)): constant
    def withdraw(inst: bytes[150], beaconHeight: uint256, beaconInstPath: bytes32[8], beaconInstPathIsLeft: bool[8], beaconInstPathLen: int128, beaconInstRoot: bytes32, beaconBlkData: bytes32, beaconBlkHash: bytes32, beaconSignerPubkeys: bytes[264], beaconSignerCount: int128, beaconSignerSig: bytes32, beaconSignerPaths: bytes32[24], beaconSignerPathIsLeft: bool[24], beaconSignerPathLen: int128, bridgeHeight: uint256, bridgeInstPath: bytes32[8], bridgeInstPathIsLeft: bool[8], bridgeInstPathLen: int128, bridgeInstRoot: bytes32, bridgeBlkData: bytes32, bridgeBlkHash: bytes32, bridgeSignerPubkeys: bytes[264], bridgeSignerCount: int128, bridgeSignerSig: bytes32, bridgeSignerPaths: bytes32[24], bridgeSignerPathIsLeft: bool[24], bridgeSignerPathLen: int128): modifying
    def withdrawed(arg0: bytes32) -> bool: constant
    def incognito() -> address(Incognito_proxy): constant


