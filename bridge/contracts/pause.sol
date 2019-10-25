pragma solidity >=0.5.0 <0.6.0;

contract AdminPausable {
    address public admin;
    bool public paused;
    uint public expire;

    constructor(address _admin) internal {
        admin = _admin;
        paused = false;
        expire = block.timestamp + 365 * 1 days;
    }

    event Paused(address pauser);
    event Unpaused(address pauser);

    modifier onlyAdmin() {
        require(msg.sender == admin, "not admin");
        _;
    }

    modifier isPaused() {
        require(paused, "not paused right now");
        _;
    }

    modifier isNotPaused() {
        require(!paused, "paused right now");
        _;
    }

    modifier isNotExpired() {
        require(block.timestamp < expire, "expired");
        _;
    }

    function pause() public onlyAdmin isNotPaused isNotExpired {
        paused = true;
        emit Paused(msg.sender);
    }

    function unpause() public onlyAdmin isPaused {
        paused = false;
        emit Unpaused(msg.sender);
    }
}
