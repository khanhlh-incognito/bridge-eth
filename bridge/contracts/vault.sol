pragma solidity ^0.5.12;
pragma experimental ABIEncoderV2;

import "./IERC20.sol";
import "./pause.sol";

contract Incognito {
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
    ) public view returns (bool);
}

contract Withdrawable {
    function isWithdrawed(bytes32) public view returns (bool);
}

contract Vault is AdminPausable {
    address constant ETH_TOKEN = 0x0000000000000000000000000000000000000000;
    mapping(bytes32 => bool) public withdrawed;
    Incognito public incognito;
    Withdrawable public prevVault;
    address payable public newVault;

    event Deposit(address token, string incognitoAddress, uint amount);
    event Withdraw(address token, address to, uint amount);
    event Migrate(address newVault);
    event MoveAssets(address[] assets);
    event UpdateIncognitoProxy(address newIncognitoProxy);

    constructor(address admin, address incognitoProxyAddress, address _prevVault) public AdminPausable(admin) {
        incognito = Incognito(incognitoProxyAddress);
        prevVault = Withdrawable(_prevVault);
        newVault = address(0);
    }

    function deposit(string memory incognitoAddress) public payable isNotPaused {
        // require((msg.value + address(this).balance) <= 10 ** 27, "Balance of this contract has been reaching to its uint's maximum.");
        require(msg.value + address(this).balance <= 10 ** 27);
        emit Deposit(ETH_TOKEN, incognitoAddress, msg.value);
    }

    function depositERC20(address token, uint amount, string memory incognitoAddress) public payable isNotPaused {
        IERC20 erc20Interface = IERC20(token);
        uint tokenBalance = erc20Interface.balanceOf(address(this));
        require(amount + tokenBalance <= 10 ** 18);
        require(erc20Interface.transferFrom(msg.sender, address(this), amount));
        emit Deposit(token, incognitoAddress, amount);
    }

    function isWithdrawed(bytes32 hash) public view returns(bool) {
        if (withdrawed[hash]) {
            return true;
        } else if (address(prevVault) == address(0)) {
            return false;
        }
        return prevVault.isWithdrawed(hash);
    }

    function parseBurnInst(bytes memory inst) public pure returns (uint8, uint8, address, address payable, uint) {
        uint8 meta = uint8(inst[0]);
        uint8 shard = uint8(inst[1]);
        address token;
        address payable to;
        uint amount;
        assembly {
            // skip first 0x20 bytes (stored length of inst)
            token := mload(add(inst, 0x22)) // [2:34]
            to := mload(add(inst, 0x42)) // [34:66]
            amount := mload(add(inst, 0x62)) // [66:98]
        }
        return (meta, shard, token, to, amount);
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
        bytes32 beaconInstHash = keccak256(abi.encodePacked(inst, heights[0]));
        bytes32 bridgeInstHash = keccak256(abi.encodePacked(inst, heights[1]));
        require(!isWithdrawed(instHash));

        // Verify instruction on beacon
        require(incognito.instructionApproved(
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
        require(incognito.instructionApproved(
            false,
            bridgeInstHash,
            heights[1],
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
    ) public payable isNotPaused {
        (uint8 meta, uint8 shard, address token, address payable to, uint burned) = parseBurnInst(inst);
        require(meta == 72 && shard == 1); // Check instruction type

        // Check if balance is enough
        if (token == ETH_TOKEN) {
            require(address(this).balance >= burned);
        } else {
            require(IERC20(token).balanceOf(address(this)) >= burned);
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
            require(IERC20(token).transfer(to, burned));
        }
        emit Withdraw(token, to, burned);
    }

    function migrate(address payable _newVault) public onlyAdmin isPaused {
        newVault = _newVault;
        emit Migrate(_newVault);
    }

    function moveAssets(address[] memory assets) public onlyAdmin isPaused {
        require(newVault != address(0));
        for (uint i = 0; i < assets.length; i++) {
            if (assets[i] == ETH_TOKEN) {
                newVault.transfer(address(this).balance);
            } else {
                uint bal = IERC20(assets[i]).balanceOf(address(this));
                if (bal > 0) {
                    require(IERC20(assets[i]).transfer(newVault, bal));
                }
            }
        }
        emit MoveAssets(assets);
    }

    function updateIncognitoProxy(address newIncognitoProxy) public onlyAdmin isPaused {
        require(newIncognitoProxy != address(0));
        incognito = Incognito(newIncognitoProxy);
        emit UpdateIncognitoProxy(newIncognitoProxy);
    }
}
