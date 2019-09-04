pragma solidity >=0.5.0 <0.6.0;
pragma experimental ABIEncoderV2;

import "openzeppelin-solidity/contracts/token/ERC20/ERC20.sol";

contract IncognitoProxy {
  function instructionApproved(
    bool,
    bytes32,
    uint,
    bytes32[] memory,
    bool[] memory,
    bytes32,
    bytes32,
    uint[] memory,
    uint8[] memory,
    bytes32[] memory,
    bytes32[] memory
  ) public returns (bool) {}
}

contract Vault {
  address constant ETH_TOKEN = 0x0000000000000000000000000000000000000000;
  address public owner;
  mapping(bytes32 => bool) withdrawed;
  IncognitoProxy incognito;

  event Deposit(address token, string incognitoAddress, uint amount);
  event Withdraw(address token, address to, uint amount);

  constructor(address incognitoProxyAddress) public payable {
    /* Set the owner to the creator of this contract */
    owner = msg.sender;
    incognito = IncognitoProxy(incognitoProxyAddress);
  }

  function deposit(string memory incognitoAddress) public payable {
    // require((msg.value + address(this).balance) <= 10 ** 27, "Balance of this contract has been reaching to its uint's maximum.");
    assert(msg.value + address(this).balance <= 10 ** 27);
    emit Deposit(ETH_TOKEN, incognitoAddress, msg.value);
  }

  function depositERC20(address token, uint amount, string memory incognitoAddress) public payable {
    ERC20 erc20Interface = ERC20(token);
    uint tokenBalance = erc20Interface.balanceOf(address(this));
    assert(amount + tokenBalance <= 10 ** 18);
    bool success = erc20Interface.transferFrom(msg.sender, address(this), amount);
    assert(success);
    emit Deposit(token, incognitoAddress, amount);
  }

  function parseBurnInst(bytes memory inst) public pure returns (uint, address, address payable, uint) {
    uint meta = uint8(inst[2]) + uint8(inst[1]) * 2 ** 8 + uint8(inst[0]) * 2 ** 16;
    address token;
    address payable to;
    uint amount;
    assembly {
      // skip first 0x20 bytes (stored length of inst)
      token := mload(add(inst, 0x23)) // [3:35]
      to := mload(add(inst, 0x43)) // [35:67]
      amount := mload(add(inst, 0x63)) // [67:99]
    }
    return (meta, token, to, amount);
  }

  function verifyInst(
    bytes memory inst,
    uint[2] memory heights,
    bytes32[][2] memory instPaths,
    bool[][2] memory instPathIsLefts,
    bytes32[2] memory instRoots,
    bytes32[2] memory blkData,
    uint[][2] memory sigIdxs,
    uint8[][2] memory sigVs,
    bytes32[][2] memory sigRs,
    bytes32[][2] memory sigSs
  ) internal {
    // Each instruction can only by redeemed once
    bytes32 instHash = keccak256(inst);
    // bytes32 beaconInstHash = keccak256(concat(inst, Memory.toBytes(heights[0], 32)));
    bytes32 beaconInstHash = keccak256(abi.encodePacked(inst, bytes32(heights[0])));
    bytes32 bridgeInstHash = keccak256(abi.encodePacked(inst, bytes32(heights[1])));
    assert(withdrawed[instHash] == false);

    // Verify instruction on beacon
    assert(incognito.instructionApproved(
      true,
      beaconInstHash,
      heights[0],
      instPaths[0],
      instPathIsLefts[0],
      instRoots[0],
      blkData[0],
      sigIdxs[0],
      sigVs[0],
      sigRs[0],
      sigSs[0]
    ));

    // Verify instruction on bridge
    assert(incognito.instructionApproved(
      false,
      bridgeInstHash,
      heights[0],
      instPaths[1],
      instPathIsLefts[1],
      instRoots[1],
      blkData[1],
      sigIdxs[1],
      sigVs[1],
      sigRs[1],
      sigSs[1]
    ));
    withdrawed[instHash] = true;
  }

  function withdraw(
    bytes memory inst,
    uint[2] memory heights,
    bytes32[][2] memory instPaths,
    bool[][2] memory instPathIsLefts,
    bytes32[2] memory instRoots,
    bytes32[2] memory blkData,
    uint[][2] memory sigIdxs,
    uint8[][2] memory sigVs,
    bytes32[][2] memory sigRs,
    bytes32[][2] memory sigSs
  ) public payable {
    (uint meta, address token, address payable to, uint burned) = parseBurnInst(inst);
    // Check instruction type
    assert(meta == 3617329); // Burn metadata and shardID of bridge

    ERC20 erc20Interface = ERC20(token);
    // Check if balance is enough
    if (token == ETH_TOKEN) {
      assert(address(this).balance >= burned);
    } else {
      assert(erc20Interface.balanceOf(address(this)) >= burned);
    }

    verifyInst(
      inst,
      heights,
      instPaths,
      instPathIsLefts,
      instRoots,
      blkData,
      sigIdxs,
      sigVs,
      sigRs,
      sigSs
    );

    // Send and notify
    if (token == ETH_TOKEN) {
      to.transfer(burned);
    } else {
      bool success = erc20Interface.transfer(to, burned);
      assert(success);
    }
    emit Withdraw(token, to, burned);
  }
}