# External Contracts
contract Incognito_proxy:
    def instructionApproved(isBeacon: bool, instHash: bytes32, blkHeight: uint256, instPath: bytes32[8], instPathIsLeft: bool[8], instPathLen: int128, instRoot: bytes32, blkData: bytes32, numSig: int128, sigIdxs: uint256[8], v: uint256[8], r: bytes32[8], s: bytes32[8]) -> bool: constant

contract Erc20:
    def transfer(_to: address, _value: uint256) -> bool: modifying
    def transferFrom(_from: address, _to: address, _value: uint256) -> bool: modifying
    def balanceOf(arg0: address) -> uint256: constant


# All these constants must mimic incognity_proxy
INST_LENGTH: constant(int128) = 300
INST_MAX_PATH: constant(uint256) = 8
COMM_SIZE: constant(uint256) = 8
INC_ADDRESS_LENGTH: constant(uint256) = 128

ETH_TOKEN: constant(address) = 0x0000000000000000000000000000000000000000

Deposit: event({_token: address, _incognito_address: string[INC_ADDRESS_LENGTH], _amount: wei_value})
Withdraw: event({_token: address, _to: address, _amount: wei_value})


NotifyString: event({content: string[128]})
NotifyBytes32: event({content: bytes32})
NotifyBool: event({content: bool})
NotifyUint256: event({content: uint256})
NotifyAddress: event({content: address})


withdrawed: public(map(bytes32, bool))
incognito: public(Incognito_proxy)

@public
def __init__(incognitoProxyAddress: address):
    self.incognito = Incognito_proxy(incognitoProxyAddress)

@public
@payable
def deposit(incognito_address: string[INC_ADDRESS_LENGTH]):
    assert msg.value + self.balance <= 10 ** 27
    log.Deposit(ETH_TOKEN, incognito_address, msg.value)

@public
def depositERC20(token: address, amount: uint256, incognito_address: string[INC_ADDRESS_LENGTH]):
    tokenBalance: uint256 = Erc20(token).balanceOf(self)
    assert amount + tokenBalance <= 10 ** 18
    success: bool = Erc20(token).transferFrom(msg.sender, self, amount)
    assert success
    log.Deposit(token, incognito_address, amount)

@constant
@public
def parseBurnInst(inst: bytes[INST_LENGTH]) -> (uint256, address, address, uint256):
    type: uint256 = convert(slice(inst, start=0, len=3), uint256)
    token: address = extract32(inst, 3, type=address)
    to: address = extract32(inst, 35, type=address)
    amount: uint256 = extract32(inst, 67, type=uint256)
    return type, token, to, amount

@public
def withdraw(
    inst: bytes[INST_LENGTH],
    beaconHeight: uint256,
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
    bridgeHeight: uint256,
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
):
    type: uint256 = 0
    token: address
    to: address
    burned: uint256 = 0
    type, token, to, burned = self.parseBurnInst(inst)
    # log.NotifyUint256(type)
    # log.NotifyAddress(token)
    # log.NotifyAddress(to)
    # log.NotifyUint256(burned)

    # Check instruction type
    assert type == 3617329 # Burn metadata and shardID of bridge

    # Check if balance is enough
    if token == ETH_TOKEN:
        assert self.balance >= burned
    else:
        enough: bool = Erc20(token).balanceOf(self) >= burned
        assert enough

    # Each instruction can only by redeemed once
    instHash: bytes32 = keccak256(inst)
    beaconInstHash: bytes32 = keccak256(concat(inst, convert(beaconHeight, bytes32)))
    bridgeInstHash: bytes32 = keccak256(concat(inst, convert(bridgeHeight, bytes32)))
    assert self.withdrawed[instHash] == False

    # Check if instruction is approved on beacon
    assert self.incognito.instructionApproved(
        True,
        beaconInstHash,
        beaconHeight,
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
    )

    # Check if instruction is approved on bridge
    assert self.incognito.instructionApproved(
        False,
        bridgeInstHash,
        bridgeHeight,
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
    )

    # Send and notify
    self.withdrawed[instHash] = True
    if token == ETH_TOKEN:
        send(to, burned)
    else:
        success: bool = Erc20(token).transfer(to, burned)
        assert success

    log.Withdraw(token, to, burned)

