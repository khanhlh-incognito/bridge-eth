pragma solidity >= 0.5.12;

import './utils.sol';
import './IERC20.sol';

interface CTokenInterface {
    function redeem(uint redeemTokens) external returns (uint);
    function redeemUnderlying(uint redeemAmount) external returns (uint);
    function borrow(uint borrowAmount) external returns (uint);
}

contract CEther is CTokenInterface {
    function mint() external payable;
    function repayBorrow() external payable;
    function repayBorrowBehalf(address borrower) external;
    function liquidateBorrow(address borrower, CTokenInterface cTokenCollateral) external payable;
}

contract CErc20 is CTokenInterface {
    address public underlying;
    function mint(uint mintAmount) external returns (uint);
    function repayBorrow(uint repayAmount) external returns (uint);
    function repayBorrowBehalf(address borrower, uint repayAmount) external returns (uint);
    function liquidateBorrow(address borrower, uint repayAmount, CTokenInterface cTokenCollateral) external returns (uint);
}

interface Comptroller {
    function enterMarkets(address[] calldata cTokens) external returns (uint[] memory);
    function exitMarket(address cTokenAddress) external returns (uint);
}

contract CompoundAgent is TradeUtils {

    Comptroller comptroller;
    CEther cEther;
    address public proxyCompound;

    constructor(address payable _incognitoSmartContract, Comptroller _comptroller, CEther _cEther, address _proxyCompound) public {
        incognitoSmartContract = _incognitoSmartContract;
        comptroller = _comptroller;
        cEther = _cEther;
        proxyCompound = _proxyCompound;
    }

    modifier onlyProxyCompound () {
        require(msg.sender == proxyCompound, "Only proxy compound call this agent");
        _;
    }

    // fallback function which allows transfer eth.
    function() external payable {}

    /**
     * @dev Call mint func to compound 
     * @param cToken: cToken of compound
     * @param amount: total to mint
     * @return bool: token address, amount recieved
     */
    function mint(address cToken, uint amount) external payable onlyProxyCompound returns (address, uint) {
        if(cToken == address(cEther)) {
            CEther(cToken).mint.value(msg.value)();
        } else {
            approve(IERC20(CErc20(cToken).underlying()), cToken, amount);
            require(CErc20(cToken).mint(amount) == 0);
        }
        uint amountAfter = balanceOf(IERC20(cToken));
        transfer(IERC20(cToken), amountAfter);
        return (cToken, amountAfter);
    }

    /**
     * @dev Call borrow func to compound 
     * @param cToken: cToken of compound
     * @param amount: total to borrow
     * @param addToMarkets: add tokens to market as collateral
     * @return bool: token address, amount recieved
     */
    function borrow(address cToken, uint amount, address[] calldata addToMarkets) external onlyProxyCompound returns (address, uint) {
        if(addToMarkets.length > 0) {
            comptroller.enterMarkets(addToMarkets);
        }
        require(CTokenInterface(cToken).borrow(amount) == 0);
        uint amountAfter = 0;
        if (cToken == address(cEther)) {
            amountAfter = balanceOf(ETH_CONTRACT_ADDRESS);
            transfer(ETH_CONTRACT_ADDRESS, amountAfter);
            return (address(ETH_CONTRACT_ADDRESS), amountAfter); 
        } else {
            amountAfter = balanceOf(IERC20(CErc20(cToken).underlying()));
            transfer(IERC20(CErc20(cToken).underlying()), amountAfter);
            return (CErc20(cToken).underlying(), amountAfter);
        }
    }

    /**
     * @dev Call borrow func to compound 
     * @param cToken: cToken of compound
     * @param amount: total to borrow
     * @param addToMarkets: add token to market as collateral
     * @return bool: token address, amount recieved
     */
    function borrowByMultiCollateral(address cToken, uint amount, address[] calldata addToMarkets) external onlyProxyCompound returns (address[] memory, uint[] memory) {
        if(addToMarkets.length > 0) {
            comptroller.enterMarkets(addToMarkets);
        }
        require(CTokenInterface(cToken).borrow(amount) == 0);
        uint amountAfter = 0;
        if (cToken == address(cEther)) {
            amountAfter = balanceOf(ETH_CONTRACT_ADDRESS);
            transfer(ETH_CONTRACT_ADDRESS, amountAfter);
        } else {
            amountAfter = balanceOf(IERC20(CErc20(cToken).underlying()));
            transfer(IERC20(CErc20(cToken).underlying()), amountAfter);
        }
        address[] memory addressAfter = new address[](1);
        addressAfter[0] = cToken;
        uint[] memory balAfter = new uint[](1);
        balAfter[0] = amountAfter;
        return (addressAfter,  balAfter);
    }

    /**
     * @dev Call mint func to compound 
     * @param cToken: cToken of compound
     * @param amount: total to redeem
     * @param isredeemUnderlying: whether redeem cToken or native coin
     * @param exitToMarkets: remove coin from market
     * @return bool: token address, amount recieved
     */
    function redeem(address cToken, uint amount, bool isredeemUnderlying, address exitToMarkets) external onlyProxyCompound returns (address, uint) {
        if(exitToMarkets != address(0x0)) {
            comptroller.exitMarket(exitToMarkets);   
        }
        if(!isredeemUnderlying) {
            require(CTokenInterface(cToken).redeem(amount) == 0);
        } else {
            require(CTokenInterface(cToken).redeemUnderlying(amount) == 0);
        }
        uint amountAfter = 0;
        if (cToken == address(cEther)) {
            amountAfter = balanceOf(ETH_CONTRACT_ADDRESS);
            transfer(ETH_CONTRACT_ADDRESS, amountAfter);
            return (address(ETH_CONTRACT_ADDRESS), amountAfter); 
        } else {
            amountAfter = balanceOf(IERC20(CErc20(cToken).underlying()));
            transfer(IERC20(CErc20(cToken).underlying()), amountAfter);
            return (CErc20(cToken).underlying(), amountAfter);
        }
    }

    /**
     * @dev Call repayBorrow func to compound 
     * @param cToken: cToken of compound
     * @param amount: total to repay
     * @return bool: token address, amount (default 0)
     */
    function repayBorrow(address cToken, uint amount) external payable onlyProxyCompound returns (address, uint) {
        if(cToken == address(cEther)) {
            CEther(cToken).repayBorrow.value(msg.value)();
        } else {
            approve(IERC20(CErc20(cToken).underlying()), cToken, amount);
            require(CErc20(cToken).repayBorrow(amount) == 0);
        }

        return (cToken, 0);
    }

    // TODO:
    // function repayBorrowBehalf() external payable onlyProxyCompound return (address, uint) {}

    /**
     * @dev Call liquidateBorrow func to compound 
     * @param cToken: cToken of compound
     * @param borrower: the address can't pay this borrow
     * @param repayAmount: total to repay
     * @param cTokenCollateral: the address of ctoken 
     * @return bool: token address, amount recieved
     */
    function liquidateBorrow(address cToken, address borrower, uint repayAmount, address cTokenCollateral) external payable onlyProxyCompound returns (address, uint) {
        if(cToken == address(cEther)) {
            CEther(cToken).liquidateBorrow.value(msg.value)(borrower, CTokenInterface(cTokenCollateral));
        } else {
            approve(IERC20(CErc20(cToken).underlying()), cToken, repayAmount);
            require(CErc20(cToken).liquidateBorrow(borrower, repayAmount, CTokenInterface(cTokenCollateral)) == 0);
        }
        uint amountAfter = balanceOf(IERC20(cTokenCollateral));
        transfer(IERC20(cTokenCollateral), amountAfter);
        return (cTokenCollateral, amountAfter);
    }

    /**
    * @dev upgrade proxy compound  
    */
    function upgrade(address _proxyCompound) external onlyProxyCompound returns (address, uint) {
        proxyCompound = _proxyCompound;
        return (address(0x0), 0);
    }
}

contract CompoundProxy is TradeUtils {
     
    mapping(bytes32 => bool) public sigDataUsed;
    mapping(address => address) public agents;
    
    Comptroller comptroller;
    CEther cEther;

    constructor(address payable _incognitoSmartContract, Comptroller _comptroller, CEther _cEther) public {
        incognitoSmartContract = _incognitoSmartContract;
        comptroller = _comptroller;
        cEther = _cEther;
    }
    
    /**
     * @dev Checks if a caller already has agent or not
     * @notice First, we check inside the storage of this contract itself. If the
     * hash has been used before, we return the result. Otherwise, we query
     * previous vault recursively until the first compound proxy
     * @param caller: address of caller
     * @return address: agent contract address of this caller
     */
    function isAgentExist(address caller) internal returns(address) {
        if(agents[caller] == address(0x0)) {
            agents[caller] = address(new CompoundAgent(incognitoSmartContract, comptroller, cEther, address(this)));
        }
        return agents[caller];
    }
    
    /**
     * @dev Checks if a sig data has been used before
     * @notice First, we check inside the storage of this contract itself. If the
     * hash has been used before, we return the result. Otherwise, we query
     * previous vault recursively until the first Vault (prevVault address is 0x0)
     * @param hash: of the sig data
     * @return bool: whether the sig data has been used or not
     */
    function isSigDataUsed(bytes32 hash) public view returns(bool) {
        if (sigDataUsed[hash]) {
            return true;
        }
        return false;
    }
    
    /**
     * @dev Call agent contract to interact with compound
     * @param srcToken: The address of token
     * @param callData: The function and param to foward to agent
     * @param timestamp: Random data for signing 
     * @param signData: The signature of owner
     * @return bool: whether the sig data has been used or not
     */
    function execute(
        IERC20 srcToken,
        uint amount,
        bytes memory callData,
        bytes memory timestamp,
        bytes memory signData
    ) public payable isIncognitoSmartContract returns(address, uint)  {
        //verify ower signs data from input
        address verifier = verifySignData(abi.encodePacked(srcToken, callData, timestamp), signData);
        address agent = isAgentExist(verifier);
        if(srcToken != ETH_CONTRACT_ADDRESS) {
            srcToken.transfer(agent, amount);
            checkSuccess();
        }
        (bool success, bytes memory result) = agent.call.value(msg.value)(callData);
        require(success);

        return abi.decode(result, (address, uint));
    }

    /**
     * @dev : borrow by sending many coins to 
     */
    function executeMulti(
        IERC20[] memory srcTokens,
        uint[] memory amounts,
        bytes memory callData,
        bytes memory timestamp,
        bytes memory signData
    ) public payable isIncognitoSmartContract returns(address[] memory, uint[] memory) {
        require(srcTokens.length == amounts.length);
        //verify ower signs data from input
        address verifier = verifySignData(abi.encodePacked(callData, timestamp), signData);
        address agent = isAgentExist(verifier);
        for(uint i = 0; i < srcTokens.length; i++) {
            if(srcTokens[i] != ETH_CONTRACT_ADDRESS) {
                srcTokens[i].transfer(agent, amounts[i]);
                checkSuccess();
            }
        }
        (bool success, bytes memory result) = agent.call.value(msg.value)(callData);
        require(success);

        return abi.decode(result, (address[], uint[]));
    }

    /**
     * @dev generate address from signature data and hash.
     */
    function sigToAddress(bytes memory signData, bytes32 hash) public pure returns (address) {
        bytes32 s;
        bytes32 r;
        uint8 v;
        assembly {
            r := mload(add(signData, 0x20))
            s := mload(add(signData, 0x40))
        }
        v = uint8(signData[64]) + 27;
        return ecrecover(hash, v, r, s);
    }

    /**
     * @dev return agent contract address.
     */
    function getAgentAddress(address caller) public view returns (address) {
        return agents[caller];
    }

    /**
     * @dev verify sign data
     */
    function verifySignData(bytes memory data, bytes memory signData) internal returns(address){
        bytes32 hash = keccak256(data);
        require(!isSigDataUsed(hash));
        address verifier = sigToAddress(signData, hash);
        // mark data hash of sig as used
        sigDataUsed[hash] = true;
        
        return verifier;
    }
}

