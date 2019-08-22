pragma solidity >=0.4.21 <0.6.0;

contract ECDSA {
    function verify(address[] memory committee, bytes32 msgHash, uint256[] memory v, bytes32[] memory r, bytes32[] memory s)
    public pure returns (bool) {
        for (uint i = 0; i<committee.length; i++){
            if (ecrecover(msgHash, uint8(v[i]), r[i], s[i]) != committee[i]) {
                return false;
            }
        }
        return true;
    }
}
