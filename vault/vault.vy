# External Contracts
contract Incognito_proxy:
    def instructionApproved(beaconInstHash: bytes32, beaconHeight: uint256, beaconInstPath: bytes32[8], beaconInstPathIsLeft: bool[8], beaconInstPathLen: int128, beaconInstRoot: bytes32, beaconBlkData: bytes32, beaconBlkHash: bytes32, beaconSignerSig: bytes32, bridgeInstHash: bytes32, bridgeHeight: uint256, bridgeInstPath: bytes32[8], bridgeInstPathIsLeft: bool[8], bridgeInstPathLen: int128, bridgeInstRoot: bytes32, bridgeBlkData: bytes32, bridgeBlkHash: bytes32, bridgeSignerSig: bytes32) -> bool: constant

contract Erc20:
    def transfer(_to: address, _value: uint256) -> bool: modifying
    def transferFrom(_from: address, _to: address, _value: uint256) -> bool: modifying
    def balanceOf(arg0: address) -> uint256: constant


# All these constants must mimic incognity_proxy
INST_LENGTH: constant(int128) = 300
INST_MAX_PATH: constant(uint256) = 8
COMM_SIZE: constant(uint256) = 8
PUBKEY_SIZE: constant(int128) = 33
PUBKEY_LENGTH: constant(int128) = INST_LENGTH
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
    assert msg.value <= 10 ** 27
    log.Deposit(ETH_TOKEN, incognito_address, msg.value)

@public
def depositERC20(token: address, amount: uint256, incognito_address: string[INC_ADDRESS_LENGTH]):
    assert amount <= 10 ** 18
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

# TODO: remove test function
@constant
@public
def testExtract(a: bytes[INST_LENGTH]) -> (address, wei_value):
    x: address = extract32(a, 0, type=address)
    s: uint256 = 12345
    t: wei_value = as_wei_value(s, "gwei")
    return x, t

@public
def withdraw(
    inst: bytes[INST_LENGTH],
    beaconHeight: uint256,
    beaconInstPath: bytes32[INST_MAX_PATH],
    beaconInstPathIsLeft: bool[INST_MAX_PATH],
    beaconInstPathLen: int128,
    beaconInstRoot: bytes32,
    beaconBlkData: bytes32,
    beaconBlkHash: bytes32,
    beaconSignerSig: uint256,
    bridgeHeight: uint256,
    bridgeInstPath: bytes32[INST_MAX_PATH],
    bridgeInstPathIsLeft: bool[INST_MAX_PATH],
    bridgeInstPathLen: int128,
    bridgeInstRoot: bytes32,
    bridgeBlkData: bytes32,
    bridgeBlkHash: bytes32,
    bridgeSignerSig: uint256,
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

    # # Check if instruction is approved on Incognito
    # assert self.incognito.instructionApproved(
    #     beaconInstHash,
    #     beaconHeight,
    #     beaconInstPath,
    #     beaconInstPathIsLeft,
    #     beaconInstPathLen,
    #     beaconInstRoot,
    #     beaconBlkData,
    #     beaconBlkHash,
    #     beaconSignerSig,
    #     bridgeInstHash,
    #     bridgeHeight,
    #     bridgeInstPath,
    #     bridgeInstPathIsLeft,
    #     bridgeInstPathLen,
    #     bridgeInstRoot,
    #     bridgeBlkData,
    #     bridgeBlkHash,
    #     bridgeSignerSig,
    # )

    # Send and notify
    self.withdrawed[instHash] = True
    if token == ETH_TOKEN:
        send(to, burned)
    else:
        success: bool = Erc20(token).transfer(to, burned)
        assert success

    log.Withdraw(token, to, burned)

