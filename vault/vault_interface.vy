
# External Contracts
contract Vault:
    def deposit(incognito_address: string[128]): modifying
    def depositERC20(token: address, amount: uint256, incognito_address: string[128]): modifying
    def parseBurnInst(inst: bytes[300]) -> (uint256, address, address, uint256): constant
    def withdraw(inst: bytes[300], beaconHeight: uint256, beaconInstPath: bytes32[8], beaconInstPathIsLeft: bool[8], beaconInstPathLen: int128, beaconInstRoot: bytes32, beaconBlkData: bytes32, beaconNumSig: int128, beaconSigIdxs: uint256[10], beaconSigVs: uint256[10], beaconSigRs: bytes32[10], beaconSigSs: bytes32[10], bridgeHeight: uint256, bridgeInstPath: bytes32[8], bridgeInstPathIsLeft: bool[8], bridgeInstPathLen: int128, bridgeInstRoot: bytes32, bridgeBlkData: bytes32, bridgeNumSig: int128, bridgeSigIdxs: uint256[10], bridgeSigVs: uint256[10], bridgeSigRs: bytes32[10], bridgeSigSs: bytes32[10]): modifying
    def withdrawed(arg0: bytes32) -> bool: constant
    def incognito() -> address(Incognito_proxy): constant


