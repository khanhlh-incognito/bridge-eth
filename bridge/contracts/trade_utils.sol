pragma solidity >= 0.5.12;

import './IERC20.sol';

contract TradeUtils {
	IERC20 constant public ETH_CONTRACT_ADDRESS = IERC20(0x0000000000000000000000000000000000000000);
	address payable public incognitoSmartContract;

	modifier isIncognitoSmartContract {
	    require(msg.sender == incognitoSmartContract);
	    _;
	}

	// fallback function is used to receive eth.
	function() external payable {}

	function balanceOf(IERC20 token) internal view returns (uint256) {
		if (token == ETH_CONTRACT_ADDRESS) {
			return address(this).balance;
		}
        return token.balanceOf(address(this));
    }

	function transfer(IERC20 token, uint amount) internal {
		if (token == ETH_CONTRACT_ADDRESS) {
			require(address(this).balance >= amount);
			incognitoSmartContract.transfer(amount);
		} else {
			token.transfer(incognitoSmartContract, amount);
		}
	}

	function approve(IERC20 token, address proxy, uint amount) internal {
		if (token != ETH_CONTRACT_ADDRESS) {
			require(token.approve(proxy, 0));
			require(token.approve(proxy, amount));
		}
	}
}
