pragma solidity ^0.5.0;

contract EcRec {
    function go(uint8 _v) public returns (address) {
        address signer = address(0x0);
        signer = ecrecover(bytes32(0), _v, bytes32(0x0), bytes32(0x0));
        // for (int i = 0; i < 10; i++) {
        //     signer = ecrecover(_message, _v, _r, _s);
        // }
        return signer;
    }
}

