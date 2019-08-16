
# External Contracts
contract Vault:
    def deposit(incognito_address: string[128]): modifying
    def depositERC20(token: address, amount: uint256, incognito_address: string[128]): modifying
    def parseBurnInst(inst: bytes[300]) -> (uint256, address, address, uint256): constant
    def withdraw(inst: bytes[300], beaconHeight: uint256, beaconInstPath: bytes32[8], beaconInstPathIsLeft: bool[8], beaconInstPathLen: int128, beaconInstRoot: bytes32, beaconBlkData: bytes32, beaconSignerSig: uint256, beaconNumR: int128, beaconXs: uint256[8], beaconYs: uint256[8], beaconRIdxs: int128[8], beaconNumSig: int128, beaconSigIdxs: uint256[8], beaconRp: bytes[33], beaconRpx: uint256, beaconRpy: uint256, beaconR: bytes[33], bridgeHeight: uint256, bridgeInstPath: bytes32[8], bridgeInstPathIsLeft: bool[8], bridgeInstPathLen: int128, bridgeInstRoot: bytes32, bridgeBlkData: bytes32, bridgeSignerSig: uint256, bridgeNumR: int128, bridgeXs: uint256[8], bridgeYs: uint256[8], bridgeRIdxs: int128[8], bridgeNumSig: int128, bridgeSigIdxs: uint256[8], bridgeRp: bytes[33], bridgeRpx: uint256, bridgeRpy: uint256, bridgeR: bytes[33]): modifying
    def withdrawed(arg0: bytes32) -> bool: constant
    def incognito() -> address(Incognito_proxy): constant


