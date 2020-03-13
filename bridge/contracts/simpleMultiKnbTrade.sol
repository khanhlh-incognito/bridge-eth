pragma solidity >=0.5.12;

import './trade_utils.sol';
import './IERC20.sol';

contract KyberNetwork {
    function trade(IERC20 src, uint srcAmount, IERC20 dest, address destAddress, uint maxDestAmount, uint minConversionRate, address walletId) public payable returns(uint);
    function swapTokenToToken(IERC20 src, uint srcAmount, IERC20 dest, uint minConversionRate) public returns(uint);
    function swapEtherToToken(IERC20 token, uint minConversionRate) public payable returns(uint);
    function swapTokenToEther(IERC20 token, uint srcAmount, uint minConversionRate) public returns(uint);
    function getExpectedRate(IERC20 src, IERC20 dest, uint srcQty) public view returns(uint expectedRate, uint slippageRate);
}

contract KBNMultiTrade is TradeUtils {
    // Variables
    KyberNetwork public kyberNetworkProxyContract;
    IERC20 constant KYBER_ETH_TOKEN_ADDRESS = IERC20(0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE);

    // Functions
    /**
     * @dev Contract constructor
     * @param _kyberNetworkProxyContract KyberNetworkProxy contract address
     */
    constructor(KyberNetwork _kyberNetworkProxyContract, address payable _incognitoSmartContract) public {
        kyberNetworkProxyContract = _kyberNetworkProxyContract;
        incognitoSmartContract = _incognitoSmartContract;
    }

    // fallback function which allows transfer eth.
    function() external payable {}

    /**
     * @dev Gets the conversion rate for the destToken given the srcQty.
     * @param srcToken source token contract address
     * @param srcQty amount of source tokens
     * @param destToken destination token contract address
     */
    function getConversionRates(IERC20 srcToken, uint srcQty, IERC20 destToken) public view returns (uint, uint) {
        return kyberNetworkProxyContract.getExpectedRate(srcToken, destToken, srcQty);
    }

    function trade(address[] memory srcTokens, uint[] memory srcQties, address[] memory destTokens) public payable isIncognitoSmartContract returns (address[] memory, uint[] memory) {
        require(srcTokens.length == srcQties.length && destTokens.length == srcTokens.length);
        uint[] memory amounts = new uint[](destTokens.length);
        for(uint i = 0; i < srcTokens.length; i++) {
            require(balanceOf(IERC20(srcTokens[i])) >= srcQties[i]);
            require(srcTokens[i] != destTokens[i]);
            if (IERC20(srcTokens[i]) != ETH_CONTRACT_ADDRESS) {
                // approve
                approve(IERC20(srcTokens[i]), address(kyberNetworkProxyContract), srcQties[i]);
                if (IERC20(destTokens[i]) != ETH_CONTRACT_ADDRESS) { // token to token.
                    require(tokenToToken(IERC20(srcTokens[i]), srcQties[i], IERC20(destTokens[i])) > 0);
                } else {
                    require(tokenToEth(IERC20(srcTokens[i]), srcQties[i]) > 0);
                }
            } else {
                require(ethToToken(IERC20(destTokens[i]), srcQties[i]) > 0);
            }
            // transfer back to incognito smart contract
            amounts[i] = balanceOf(IERC20(destTokens[i]));
            transfer(IERC20(destTokens[i]), amounts[i]);
        }
        return (destTokens, amounts);
    }

    function ethToToken(IERC20 token, uint srcQty) internal returns (uint) {
        // Get the minimum conversion rate
        require(address(this).balance >= srcQty);
        (uint minConversionRate,) = kyberNetworkProxyContract.getExpectedRate(KYBER_ETH_TOKEN_ADDRESS, token, msg.value);
        return kyberNetworkProxyContract.swapEtherToToken.value(srcQty)(token, minConversionRate);
    }

    function tokenToEth(IERC20 token, uint amount) internal returns (uint) {
        (uint minConversionRate,) = kyberNetworkProxyContract.getExpectedRate(token, KYBER_ETH_TOKEN_ADDRESS, amount);
        return kyberNetworkProxyContract.swapTokenToEther(token, amount, minConversionRate);
    }

    function tokenToToken(IERC20 srcToken, uint srcQty, IERC20 destToken) internal returns (uint) {
        (uint minConversionRate,) = kyberNetworkProxyContract.getExpectedRate(srcToken, destToken, srcQty);
        return kyberNetworkProxyContract.swapTokenToToken(srcToken, srcQty, destToken, minConversionRate);
    }
}
