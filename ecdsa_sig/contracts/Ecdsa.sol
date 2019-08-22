pragma solidity >=0.4.21 <0.6.0;

contract ECDSA {
    uint constant COMM_SIZE = 10;

    function verify(int128 numSig, address[COMM_SIZE] memory committee, bytes32 msgHash, uint256[COMM_SIZE] memory v, bytes32[COMM_SIZE] memory r, bytes32[COMM_SIZE] memory s)
    public pure returns (bool) {
        for (uint i = 0; i<uint256(numSig); i++){
            if (ecrecover(msgHash, uint8(v[i]), r[i], s[i]) != committee[i]) {
                return false;
            }
        }
        return true;
    }
}
