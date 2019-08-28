// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridge

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IncognitoProxyABI is the input ABI used to generate the binding from.
const IncognitoProxyABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"}],\"name\":\"parseBurnInstruction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"instPaths\",\"type\":\"bytes32[][2]\"},{\"internalType\":\"bool[][2]\",\"name\":\"instPathIsLefts\",\"type\":\"bool[][2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"instRoots\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"blkData\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256[][2]\",\"name\":\"sigIdxs\",\"type\":\"uint256[][2]\"},{\"internalType\":\"uint8[][2]\",\"name\":\"sigVs\",\"type\":\"uint8[][2]\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"sigRs\",\"type\":\"bytes32[][2]\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"sigSs\",\"type\":\"bytes32[][2]\"}],\"name\":\"swapBridgeCommittee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"committee\",\"type\":\"address[]\"},{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"verifySig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"left\",\"type\":\"bool[]\"}],\"name\":\"instructionInMerkleTree\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"numVals\",\"type\":\"uint256\"}],\"name\":\"extractCommitteeFromInstruction\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"}],\"name\":\"extractMetaFromInstruction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bridgeCommittees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"instPath\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"instPathIsLeft\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"instRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdx\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigS\",\"type\":\"bytes32[]\"}],\"name\":\"swapBeaconCommittee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"beaconCommittees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isBeacon\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"instHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blkHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"instPath\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"instPathIsLeft\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"instRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdx\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigS\",\"type\":\"bytes32[]\"}],\"name\":\"instructionApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blkHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isBeacon\",\"type\":\"bool\"}],\"name\":\"findCommitteeFromHeight\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"committee\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"beaconCommittee\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"bridgeCommittee\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"LogUint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"val\",\"type\":\"string\"}],\"name\":\"LogString\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"val\",\"type\":\"bytes32\"}],\"name\":\"LogBytes32\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"LogAddress\",\"type\":\"event\"}]"

// IncognitoProxyFuncSigs maps the 4-byte function signature to its string representation.
var IncognitoProxyFuncSigs = map[string]string{
	"f203a5ed": "beaconCommittees(uint256)",
	"9b30b637": "bridgeCommittees(uint256)",
	"8eb60066": "extractCommitteeFromInstruction(bytes,uint256)",
	"90500bae": "extractMetaFromInstruction(bytes)",
	"ff2293f6": "findCommitteeFromHeight(uint256,bool)",
	"f65d2116": "instructionApproved(bool,bytes32,uint256,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])",
	"47c4b328": "instructionInMerkleTree(bytes32,bytes32,bytes32[],bool[])",
	"036010bd": "parseBurnInstruction(bytes)",
	"e41be775": "swapBeaconCommittee(bytes,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])",
	"262f7220": "swapBridgeCommittee(bytes,bytes32[][2],bool[][2],bytes32[2],bytes32[2],uint256[][2],uint8[][2],bytes32[][2],bytes32[][2])",
	"3aacfdad": "verifySig(address[],bytes32,uint8[],bytes32[],bytes32[])",
}

// IncognitoProxyBin is the compiled bytecode used for deploying new contracts.
var IncognitoProxyBin = "0x60806040523480156200001157600080fd5b5060405162001c1938038062001c1983398101604081905262000034916200024b565b60408051808201909152828152600060208083018290528154600181018084559280528351805193949360029092027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563019262000097928492909101906200011b565b5060209182015160019182015560408051808201909152848152600081840181905282548084018085559390915281518051939550919360029091027fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf601926200010592849201906200011b565b506020820151816001015550505050506200032d565b82805482825590600052602060002090810192821562000173579160200282015b828111156200017357825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906200013c565b506200018192915062000185565b5090565b620001ac91905b80821115620001815780546001600160a01b03191681556001016200018c565b90565b8051620001bc8162000313565b92915050565b600082601f830112620001d457600080fd5b8151620001eb620001e582620002e0565b620002b9565b915081818352602084019350602081019050838560208402820111156200021157600080fd5b60005b838110156200024157816200022a8882620001af565b845250602092830192919091019060010162000214565b5050505092915050565b600080604083850312156200025f57600080fd5b82516001600160401b038111156200027657600080fd5b6200028485828601620001c2565b92505060208301516001600160401b03811115620002a157600080fd5b620002af85828601620001c2565b9150509250929050565b6040518181016001600160401b0381118282101715620002d857600080fd5b604052919050565b60006001600160401b03821115620002f757600080fd5b5060209081020190565b60006001600160a01b038216620001bc565b6200031e8162000301565b81146200032a57600080fd5b50565b6118dc806200033d6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806390500bae1161007157806390500bae146101425780639b30b63714610164578063e41be77514610184578063f203a5ed14610197578063f65d2116146101aa578063ff2293f6146101bd576100a9565b8063036010bd146100ae578063262f7220146100da5780633aacfdad146100ef57806347c4b3281461010f5780638eb6006614610122575b600080fd5b6100c16100bc366004611256565b6101d0565b6040516100d19493929190611749565b60405180910390f35b6100ed6100e836600461128a565b610265565b005b6101026100fd366004610f7f565b610309565b6040516100d191906116e8565b61010261011d3660046111c7565b6103ea565b610135610130366004611544565b6104da565b6040516100d191906116d0565b610155610150366004611256565b61055b565b6040516100d193929190611771565b610177610172366004611594565b6105e5565b6040516100d191906116f6565b6100ed6101923660046113f0565b61060b565b6101776101a5366004611594565b6106ae565b6101026101b836600461104f565b6106bb565b6101356101cb3660046115b2565b610843565b6000806000806000856000815181106101e557fe5b602001015160f81c60f81b60f81c60ff1662010000028660018151811061020857fe5b602001015160f81c60f81b60f81c60ff16610100028760028151811061022a57fe5b602091010151602389015160438a01516063909a015160f89290921c9290920161ffff169290920162ffffff16989097965090945092505050565b885160208a0120600080546102bf91600191849190600019810190811061028857fe5b9060005260206000209060020201600101548c6000600281106102a757fe5b60200201518c518c518c518c518c518c518c516106bb565b6102c857600080fd5b7fa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f6040516102f590611739565b60405180910390a150505050505050505050565b6000805b84518110156103db5786818151811061032257fe5b60200260200101516001600160a01b031660018787848151811061034257fe5b602002602001015187858151811061035657fe5b602002602001015187868151811061036a57fe5b60200260200101516040516000815260200160405260405161038f9493929190611704565b6020604051602081039080840390855afa1580156103b1573d6000803e3d6000fd5b505050602060405103516001600160a01b0316146103d35760009150506103e1565b60010161030d565b50600190505b95945050505050565b600084815b84518110156104cc5783818151811061040457fe5b6020026020010151156104545784818151811061041d57fe5b6020026020010151826040516020016104379291906116aa565b6040516020818303038152906040528051906020012091506104c4565b84818151811061046057fe5b60200260200101516000801b14156104855781826040516020016104379291906116aa565b8185828151811061049257fe5b60200260200101516040516020016104ab9291906116aa565b6040516020818303038152906040528051906020012091505b6001016103ef565b50841490505b949350505050565b60608082604051908082528060200260200182016040528015610507578160200160208202803883390190505b5090506000805b8481101561054f576020810260638701015191508183828151811061052f57fe5b6001600160a01b039092166020928302919091019091015260010161050e565b50909150505b92915050565b6000806000808460008151811061056e57fe5b602001015160f81c60f81b60f81c60ff1662010000028560018151811061059157fe5b602001015160f81c60f81b60f81c60ff1661010002866002815181106105b357fe5b602091010151602388015160439098015160f89190911c9190910161ffff169190910162ffffff169690945092505050565b600181815481106105f257fe5b6000918252602090912060016002909202010154905081565b885160208a01206000805461064d91600191849190600019810190811061062e57fe5b9060005260206000209060020201600101548c8c8c8c8c8c8c8c6106bb565b61065657600080fd5b60008060006106648d61055b565b9250925092507fa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f60405161069790611739565b60405180910390a150505050505050505050505050565b600081815481106105f257fe5b600060606106c98b8e610843565b905060005b86518110156107925760008111801561071057508660018203815181106106f157fe5b602002602001015187828151811061070557fe5b602002602001015111155b8061072f5750815187828151811061072457fe5b602002602001015110155b1561073f57600092505050610834565b8187828151811061074c57fe5b60200260200101518151811061075e57fe5b602002602001015182828151811061077257fe5b6001600160a01b03909216602092830291909101909101526001016106ce565b50600087896040516020016107a89291906116aa565b6040516020818303038152906040528051906020012090507e9fd52f05c0ded31d6fb0ee580b923f85e99cf1a5a1da342f25e73c45829c83816040516107ee91906116f6565b60405180910390a1600382516002028161080457fe5b0487511161081757600092505050610834565b6108248282888888610309565b61082d57600080fd5b6001925050505b9b9a5050505050505050505050565b6060811561090a57600054805b801561090357846000600183038154811061086757fe5b906000526020600020906002020160010154116108fa576000600182038154811061088e57fe5b600091825260209182902060029091020180546040805182850281018501909152818152928301828280156108ec57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116108ce575b505050505092505050610555565b60001901610850565b5050610555565b600154805b8015610954578460018083038154811061092557fe5b9060005260206000209060020201600101541161094b5760018082038154811061088e57fe5b6000190161090f565b505092915050565b803561055581611867565b600082601f83011261097857600080fd5b813561098b610986826117bf565b611799565b915081818352602084019350602081019050838560208402820111156109b057600080fd5b60005b838110156109dc57816109c6888261095c565b84525060209283019291909101906001016109b3565b5050505092915050565b600082601f8301126109f757600080fd5b6002610a05610986826117df565b9150818360005b838110156109dc5781358601610a228882610b2e565b8452506020928301929190910190600101610a0c565b600082601f830112610a4957600080fd5b6002610a57610986826117df565b9150818360005b838110156109dc5781358601610a748882610c6f565b8452506020928301929190910190600101610a5e565b600082601f830112610a9b57600080fd5b6002610aa9610986826117df565b9150818360005b838110156109dc5781358601610ac68882610d4f565b8452506020928301929190910190600101610ab0565b600082601f830112610aed57600080fd5b6002610afb610986826117df565b9150818360005b838110156109dc5781358601610b188882610e2f565b8452506020928301929190910190600101610b02565b600082601f830112610b3f57600080fd5b8135610b4d610986826117bf565b91508181835260208401935060208101905083856020840282011115610b7257600080fd5b60005b838110156109dc5781610b888882610f0f565b8452506020928301929190910190600101610b75565b600082601f830112610baf57600080fd5b8135610bbd610986826117bf565b91508181835260208401935060208101905083856020840282011115610be257600080fd5b60005b838110156109dc5781610bf88882610f0f565b8452506020928301929190910190600101610be5565b600082601f830112610c1f57600080fd5b6002610c2d610986826117df565b91508183856020840282011115610c4357600080fd5b60005b838110156109dc5781610c598882610f1a565b8452506020928301929190910190600101610c46565b600082601f830112610c8057600080fd5b8135610c8e610986826117bf565b91508181835260208401935060208101905083856020840282011115610cb357600080fd5b60005b838110156109dc5781610cc98882610f1a565b8452506020928301929190910190600101610cb6565b600082601f830112610cf057600080fd5b8135610cfe610986826117bf565b91508181835260208401935060208101905083856020840282011115610d2357600080fd5b60005b838110156109dc5781610d398882610f1a565b8452506020928301929190910190600101610d26565b600082601f830112610d6057600080fd5b8135610d6e610986826117bf565b91508181835260208401935060208101905083856020840282011115610d9357600080fd5b60005b838110156109dc5781610da98882610f1a565b8452506020928301929190910190600101610d96565b600082601f830112610dd057600080fd5b8135610dde610986826117bf565b91508181835260208401935060208101905083856020840282011115610e0357600080fd5b60005b838110156109dc5781610e198882610f1a565b8452506020928301929190910190600101610e06565b600082601f830112610e4057600080fd5b8135610e4e610986826117bf565b91508181835260208401935060208101905083856020840282011115610e7357600080fd5b60005b838110156109dc5781610e898882610f74565b8452506020928301929190910190600101610e76565b600082601f830112610eb057600080fd5b8135610ebe610986826117bf565b91508181835260208401935060208101905083856020840282011115610ee357600080fd5b60005b838110156109dc5781610ef98882610f74565b8452506020928301929190910190600101610ee6565b80356105558161187e565b803561055581611887565b600082601f830112610f3657600080fd5b8135610f44610986826117fc565b91508082526020830160208301858383011115610f6057600080fd5b610f6b83828461185b565b50505092915050565b803561055581611890565b600080600080600060a08688031215610f9757600080fd5b85356001600160401b03811115610fad57600080fd5b610fb988828901610967565b9550506020610fca88828901610f1a565b94505060408601356001600160401b03811115610fe657600080fd5b610ff288828901610e9f565b93505060608601356001600160401b0381111561100e57600080fd5b61101a88828901610cdf565b92505060808601356001600160401b0381111561103657600080fd5b61104288828901610cdf565b9150509295509295909350565b60008060008060008060008060008060006101608c8e03121561107157600080fd5b600061107d8e8e610f0f565b9b5050602061108e8e828f01610f1a565b9a5050604061109f8e828f01610f1a565b99505060608c01356001600160401b038111156110bb57600080fd5b6110c78e828f01610cdf565b98505060808c01356001600160401b038111156110e357600080fd5b6110ef8e828f01610b9e565b97505060a06111008e828f01610f1a565b96505060c06111118e828f01610f1a565b95505060e08c01356001600160401b0381111561112d57600080fd5b6111398e828f01610dbf565b9450506101008c01356001600160401b0381111561115657600080fd5b6111628e828f01610e9f565b9350506101208c01356001600160401b0381111561117f57600080fd5b61118b8e828f01610cdf565b9250506101408c01356001600160401b038111156111a857600080fd5b6111b48e828f01610cdf565b9150509295989b509295989b9093969950565b600080600080608085870312156111dd57600080fd5b60006111e98787610f1a565b94505060206111fa87828801610f1a565b93505060408501356001600160401b0381111561121657600080fd5b61122287828801610cdf565b92505060608501356001600160401b0381111561123e57600080fd5b61124a87828801610b9e565b91505092959194509250565b60006020828403121561126857600080fd5b81356001600160401b0381111561127e57600080fd5b6104d284828501610f25565b60008060008060008060008060006101608a8c0312156112a957600080fd5b89356001600160401b038111156112bf57600080fd5b6112cb8c828d01610f25565b99505060208a01356001600160401b038111156112e757600080fd5b6112f38c828d01610a38565b98505060408a01356001600160401b0381111561130f57600080fd5b61131b8c828d016109e6565b975050606061132c8c828d01610c0e565b96505060a061133d8c828d01610c0e565b95505060e08a01356001600160401b0381111561135957600080fd5b6113658c828d01610a8a565b9450506101008a01356001600160401b0381111561138257600080fd5b61138e8c828d01610adc565b9350506101208a01356001600160401b038111156113ab57600080fd5b6113b78c828d01610a38565b9250506101408a01356001600160401b038111156113d457600080fd5b6113e08c828d01610a38565b9150509295985092959850929598565b60008060008060008060008060006101208a8c03121561140f57600080fd5b89356001600160401b0381111561142557600080fd5b6114318c828d01610f25565b99505060208a01356001600160401b0381111561144d57600080fd5b6114598c828d01610cdf565b98505060408a01356001600160401b0381111561147557600080fd5b6114818c828d01610b9e565b97505060606114928c828d01610f1a565b96505060806114a38c828d01610f1a565b95505060a08a01356001600160401b038111156114bf57600080fd5b6114cb8c828d01610dbf565b94505060c08a01356001600160401b038111156114e757600080fd5b6114f38c828d01610e9f565b93505060e08a01356001600160401b0381111561150f57600080fd5b61151b8c828d01610cdf565b9250506101008a01356001600160401b0381111561153857600080fd5b6113e08c828d01610cdf565b6000806040838503121561155757600080fd5b82356001600160401b0381111561156d57600080fd5b61157985828601610f25565b925050602061158a85828601610f1a565b9150509250929050565b6000602082840312156115a657600080fd5b60006104d28484610f1a565b600080604083850312156115c557600080fd5b60006115d18585610f1a565b925050602061158a85828601610f0f565b60006115ee83836115f6565b505060200190565b6115ff81611836565b82525050565b600061161082611829565b61161a818561182d565b935061162583611823565b8060005b8381101561165357815161163d88826115e2565b975061164883611823565b925050600101611629565b509495945050505050565b6115ff81611841565b6115ff81611846565b6115ff61167c82611846565b611846565b600061168e60048361182d565b63446f6e6560e01b815260200192915050565b6115ff81611855565b60006116b68285611670565b6020820191506116c68284611670565b5060200192915050565b602080825281016116e18184611605565b9392505050565b60208101610555828461165e565b602081016105558284611667565b608081016117128287611667565b61171f60208301866116a1565b61172c6040830185611667565b6103e16060830184611667565b6020808252810161055581611681565b608081016117578287611667565b61176460208301866115f6565b61172c60408301856115f6565b6060810161177f8286611667565b61178c6020830185611667565b6104d26040830184611667565b6040518181016001600160401b03811182821017156117b757600080fd5b604052919050565b60006001600160401b038211156117d557600080fd5b5060209081020190565b60006001600160401b038211156117f557600080fd5b5060200290565b60006001600160401b0382111561181257600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b600061055582611849565b151590565b90565b6001600160a01b031690565b60ff1690565b82818337506000910152565b61187081611836565b811461187b57600080fd5b50565b61187081611841565b61187081611846565b6118708161185556fea365627a7a72315820a618f7f2c6b8dadc5e56bd965783c51fd8945d7ee8dca5cc671d01de5c11f1436c6578706572696d656e74616cf564736f6c634300050b0040"

// DeployIncognitoProxy deploys a new Ethereum contract, binding an instance of IncognitoProxy to it.
func DeployIncognitoProxy(auth *bind.TransactOpts, backend bind.ContractBackend, beaconCommittee []common.Address, bridgeCommittee []common.Address) (common.Address, *types.Transaction, *IncognitoProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(IncognitoProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IncognitoProxyBin), backend, beaconCommittee, bridgeCommittee)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IncognitoProxy{IncognitoProxyCaller: IncognitoProxyCaller{contract: contract}, IncognitoProxyTransactor: IncognitoProxyTransactor{contract: contract}, IncognitoProxyFilterer: IncognitoProxyFilterer{contract: contract}}, nil
}

// IncognitoProxy is an auto generated Go binding around an Ethereum contract.
type IncognitoProxy struct {
	IncognitoProxyCaller     // Read-only binding to the contract
	IncognitoProxyTransactor // Write-only binding to the contract
	IncognitoProxyFilterer   // Log filterer for contract events
}

// IncognitoProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type IncognitoProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncognitoProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IncognitoProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncognitoProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IncognitoProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncognitoProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IncognitoProxySession struct {
	Contract     *IncognitoProxy   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IncognitoProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IncognitoProxyCallerSession struct {
	Contract *IncognitoProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IncognitoProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IncognitoProxyTransactorSession struct {
	Contract     *IncognitoProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IncognitoProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type IncognitoProxyRaw struct {
	Contract *IncognitoProxy // Generic contract binding to access the raw methods on
}

// IncognitoProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IncognitoProxyCallerRaw struct {
	Contract *IncognitoProxyCaller // Generic read-only contract binding to access the raw methods on
}

// IncognitoProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IncognitoProxyTransactorRaw struct {
	Contract *IncognitoProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIncognitoProxy creates a new instance of IncognitoProxy, bound to a specific deployed contract.
func NewIncognitoProxy(address common.Address, backend bind.ContractBackend) (*IncognitoProxy, error) {
	contract, err := bindIncognitoProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IncognitoProxy{IncognitoProxyCaller: IncognitoProxyCaller{contract: contract}, IncognitoProxyTransactor: IncognitoProxyTransactor{contract: contract}, IncognitoProxyFilterer: IncognitoProxyFilterer{contract: contract}}, nil
}

// NewIncognitoProxyCaller creates a new read-only instance of IncognitoProxy, bound to a specific deployed contract.
func NewIncognitoProxyCaller(address common.Address, caller bind.ContractCaller) (*IncognitoProxyCaller, error) {
	contract, err := bindIncognitoProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyCaller{contract: contract}, nil
}

// NewIncognitoProxyTransactor creates a new write-only instance of IncognitoProxy, bound to a specific deployed contract.
func NewIncognitoProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*IncognitoProxyTransactor, error) {
	contract, err := bindIncognitoProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyTransactor{contract: contract}, nil
}

// NewIncognitoProxyFilterer creates a new log filterer instance of IncognitoProxy, bound to a specific deployed contract.
func NewIncognitoProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*IncognitoProxyFilterer, error) {
	contract, err := bindIncognitoProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyFilterer{contract: contract}, nil
}

// bindIncognitoProxy binds a generic wrapper to an already deployed contract.
func bindIncognitoProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IncognitoProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IncognitoProxy *IncognitoProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IncognitoProxy.Contract.IncognitoProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IncognitoProxy *IncognitoProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.IncognitoProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IncognitoProxy *IncognitoProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.IncognitoProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IncognitoProxy *IncognitoProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IncognitoProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IncognitoProxy *IncognitoProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IncognitoProxy *IncognitoProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.contract.Transact(opts, method, params...)
}

// BeaconCommittees is a free data retrieval call binding the contract method 0xf203a5ed.
//
// Solidity: function beaconCommittees(uint256 ) constant returns(uint256 startBlock)
func (_IncognitoProxy *IncognitoProxyCaller) BeaconCommittees(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "beaconCommittees", arg0)
	return *ret0, err
}

// BeaconCommittees is a free data retrieval call binding the contract method 0xf203a5ed.
//
// Solidity: function beaconCommittees(uint256 ) constant returns(uint256 startBlock)
func (_IncognitoProxy *IncognitoProxySession) BeaconCommittees(arg0 *big.Int) (*big.Int, error) {
	return _IncognitoProxy.Contract.BeaconCommittees(&_IncognitoProxy.CallOpts, arg0)
}

// BeaconCommittees is a free data retrieval call binding the contract method 0xf203a5ed.
//
// Solidity: function beaconCommittees(uint256 ) constant returns(uint256 startBlock)
func (_IncognitoProxy *IncognitoProxyCallerSession) BeaconCommittees(arg0 *big.Int) (*big.Int, error) {
	return _IncognitoProxy.Contract.BeaconCommittees(&_IncognitoProxy.CallOpts, arg0)
}

// BridgeCommittees is a free data retrieval call binding the contract method 0x9b30b637.
//
// Solidity: function bridgeCommittees(uint256 ) constant returns(uint256 startBlock)
func (_IncognitoProxy *IncognitoProxyCaller) BridgeCommittees(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "bridgeCommittees", arg0)
	return *ret0, err
}

// BridgeCommittees is a free data retrieval call binding the contract method 0x9b30b637.
//
// Solidity: function bridgeCommittees(uint256 ) constant returns(uint256 startBlock)
func (_IncognitoProxy *IncognitoProxySession) BridgeCommittees(arg0 *big.Int) (*big.Int, error) {
	return _IncognitoProxy.Contract.BridgeCommittees(&_IncognitoProxy.CallOpts, arg0)
}

// BridgeCommittees is a free data retrieval call binding the contract method 0x9b30b637.
//
// Solidity: function bridgeCommittees(uint256 ) constant returns(uint256 startBlock)
func (_IncognitoProxy *IncognitoProxyCallerSession) BridgeCommittees(arg0 *big.Int) (*big.Int, error) {
	return _IncognitoProxy.Contract.BridgeCommittees(&_IncognitoProxy.CallOpts, arg0)
}

// ExtractCommitteeFromInstruction is a free data retrieval call binding the contract method 0x8eb60066.
//
// Solidity: function extractCommitteeFromInstruction(bytes inst, uint256 numVals) constant returns(address[])
func (_IncognitoProxy *IncognitoProxyCaller) ExtractCommitteeFromInstruction(opts *bind.CallOpts, inst []byte, numVals *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "extractCommitteeFromInstruction", inst, numVals)
	return *ret0, err
}

// ExtractCommitteeFromInstruction is a free data retrieval call binding the contract method 0x8eb60066.
//
// Solidity: function extractCommitteeFromInstruction(bytes inst, uint256 numVals) constant returns(address[])
func (_IncognitoProxy *IncognitoProxySession) ExtractCommitteeFromInstruction(inst []byte, numVals *big.Int) ([]common.Address, error) {
	return _IncognitoProxy.Contract.ExtractCommitteeFromInstruction(&_IncognitoProxy.CallOpts, inst, numVals)
}

// ExtractCommitteeFromInstruction is a free data retrieval call binding the contract method 0x8eb60066.
//
// Solidity: function extractCommitteeFromInstruction(bytes inst, uint256 numVals) constant returns(address[])
func (_IncognitoProxy *IncognitoProxyCallerSession) ExtractCommitteeFromInstruction(inst []byte, numVals *big.Int) ([]common.Address, error) {
	return _IncognitoProxy.Contract.ExtractCommitteeFromInstruction(&_IncognitoProxy.CallOpts, inst, numVals)
}

// ExtractMetaFromInstruction is a free data retrieval call binding the contract method 0x90500bae.
//
// Solidity: function extractMetaFromInstruction(bytes inst) constant returns(uint256, uint256, uint256)
func (_IncognitoProxy *IncognitoProxyCaller) ExtractMetaFromInstruction(opts *bind.CallOpts, inst []byte) (*big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _IncognitoProxy.contract.Call(opts, out, "extractMetaFromInstruction", inst)
	return *ret0, *ret1, *ret2, err
}

// ExtractMetaFromInstruction is a free data retrieval call binding the contract method 0x90500bae.
//
// Solidity: function extractMetaFromInstruction(bytes inst) constant returns(uint256, uint256, uint256)
func (_IncognitoProxy *IncognitoProxySession) ExtractMetaFromInstruction(inst []byte) (*big.Int, *big.Int, *big.Int, error) {
	return _IncognitoProxy.Contract.ExtractMetaFromInstruction(&_IncognitoProxy.CallOpts, inst)
}

// ExtractMetaFromInstruction is a free data retrieval call binding the contract method 0x90500bae.
//
// Solidity: function extractMetaFromInstruction(bytes inst) constant returns(uint256, uint256, uint256)
func (_IncognitoProxy *IncognitoProxyCallerSession) ExtractMetaFromInstruction(inst []byte) (*big.Int, *big.Int, *big.Int, error) {
	return _IncognitoProxy.Contract.ExtractMetaFromInstruction(&_IncognitoProxy.CallOpts, inst)
}

// FindCommitteeFromHeight is a free data retrieval call binding the contract method 0xff2293f6.
//
// Solidity: function findCommitteeFromHeight(uint256 blkHeight, bool isBeacon) constant returns(address[] committee)
func (_IncognitoProxy *IncognitoProxyCaller) FindCommitteeFromHeight(opts *bind.CallOpts, blkHeight *big.Int, isBeacon bool) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "findCommitteeFromHeight", blkHeight, isBeacon)
	return *ret0, err
}

// FindCommitteeFromHeight is a free data retrieval call binding the contract method 0xff2293f6.
//
// Solidity: function findCommitteeFromHeight(uint256 blkHeight, bool isBeacon) constant returns(address[] committee)
func (_IncognitoProxy *IncognitoProxySession) FindCommitteeFromHeight(blkHeight *big.Int, isBeacon bool) ([]common.Address, error) {
	return _IncognitoProxy.Contract.FindCommitteeFromHeight(&_IncognitoProxy.CallOpts, blkHeight, isBeacon)
}

// FindCommitteeFromHeight is a free data retrieval call binding the contract method 0xff2293f6.
//
// Solidity: function findCommitteeFromHeight(uint256 blkHeight, bool isBeacon) constant returns(address[] committee)
func (_IncognitoProxy *IncognitoProxyCallerSession) FindCommitteeFromHeight(blkHeight *big.Int, isBeacon bool) ([]common.Address, error) {
	return _IncognitoProxy.Contract.FindCommitteeFromHeight(&_IncognitoProxy.CallOpts, blkHeight, isBeacon)
}

// ParseBurnInstruction is a free data retrieval call binding the contract method 0x036010bd.
//
// Solidity: function parseBurnInstruction(bytes inst) constant returns(uint256, address, address, uint256)
func (_IncognitoProxy *IncognitoProxyCaller) ParseBurnInstruction(opts *bind.CallOpts, inst []byte) (*big.Int, common.Address, common.Address, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(common.Address)
		ret2 = new(common.Address)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _IncognitoProxy.contract.Call(opts, out, "parseBurnInstruction", inst)
	return *ret0, *ret1, *ret2, *ret3, err
}

// ParseBurnInstruction is a free data retrieval call binding the contract method 0x036010bd.
//
// Solidity: function parseBurnInstruction(bytes inst) constant returns(uint256, address, address, uint256)
func (_IncognitoProxy *IncognitoProxySession) ParseBurnInstruction(inst []byte) (*big.Int, common.Address, common.Address, *big.Int, error) {
	return _IncognitoProxy.Contract.ParseBurnInstruction(&_IncognitoProxy.CallOpts, inst)
}

// ParseBurnInstruction is a free data retrieval call binding the contract method 0x036010bd.
//
// Solidity: function parseBurnInstruction(bytes inst) constant returns(uint256, address, address, uint256)
func (_IncognitoProxy *IncognitoProxyCallerSession) ParseBurnInstruction(inst []byte) (*big.Int, common.Address, common.Address, *big.Int, error) {
	return _IncognitoProxy.Contract.ParseBurnInstruction(&_IncognitoProxy.CallOpts, inst)
}

// VerifySig is a free data retrieval call binding the contract method 0x3aacfdad.
//
// Solidity: function verifySig(address[] committee, bytes32 msgHash, uint8[] v, bytes32[] r, bytes32[] s) constant returns(bool)
func (_IncognitoProxy *IncognitoProxyCaller) VerifySig(opts *bind.CallOpts, committee []common.Address, msgHash [32]byte, v []uint8, r [][32]byte, s [][32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "verifySig", committee, msgHash, v, r, s)
	return *ret0, err
}

// VerifySig is a free data retrieval call binding the contract method 0x3aacfdad.
//
// Solidity: function verifySig(address[] committee, bytes32 msgHash, uint8[] v, bytes32[] r, bytes32[] s) constant returns(bool)
func (_IncognitoProxy *IncognitoProxySession) VerifySig(committee []common.Address, msgHash [32]byte, v []uint8, r [][32]byte, s [][32]byte) (bool, error) {
	return _IncognitoProxy.Contract.VerifySig(&_IncognitoProxy.CallOpts, committee, msgHash, v, r, s)
}

// VerifySig is a free data retrieval call binding the contract method 0x3aacfdad.
//
// Solidity: function verifySig(address[] committee, bytes32 msgHash, uint8[] v, bytes32[] r, bytes32[] s) constant returns(bool)
func (_IncognitoProxy *IncognitoProxyCallerSession) VerifySig(committee []common.Address, msgHash [32]byte, v []uint8, r [][32]byte, s [][32]byte) (bool, error) {
	return _IncognitoProxy.Contract.VerifySig(&_IncognitoProxy.CallOpts, committee, msgHash, v, r, s)
}

// InstructionApproved is a paid mutator transaction binding the contract method 0xf65d2116.
//
// Solidity: function instructionApproved(bool isBeacon, bytes32 instHash, uint256 blkHeight, bytes32[] instPath, bool[] instPathIsLeft, bytes32 instRoot, bytes32 blkData, uint256[] sigIdx, uint8[] sigV, bytes32[] sigR, bytes32[] sigS) returns(bool)
func (_IncognitoProxy *IncognitoProxyTransactor) InstructionApproved(opts *bind.TransactOpts, isBeacon bool, instHash [32]byte, blkHeight *big.Int, instPath [][32]byte, instPathIsLeft []bool, instRoot [32]byte, blkData [32]byte, sigIdx []*big.Int, sigV []uint8, sigR [][32]byte, sigS [][32]byte) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "instructionApproved", isBeacon, instHash, blkHeight, instPath, instPathIsLeft, instRoot, blkData, sigIdx, sigV, sigR, sigS)
}

// InstructionApproved is a paid mutator transaction binding the contract method 0xf65d2116.
//
// Solidity: function instructionApproved(bool isBeacon, bytes32 instHash, uint256 blkHeight, bytes32[] instPath, bool[] instPathIsLeft, bytes32 instRoot, bytes32 blkData, uint256[] sigIdx, uint8[] sigV, bytes32[] sigR, bytes32[] sigS) returns(bool)
func (_IncognitoProxy *IncognitoProxySession) InstructionApproved(isBeacon bool, instHash [32]byte, blkHeight *big.Int, instPath [][32]byte, instPathIsLeft []bool, instRoot [32]byte, blkData [32]byte, sigIdx []*big.Int, sigV []uint8, sigR [][32]byte, sigS [][32]byte) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.InstructionApproved(&_IncognitoProxy.TransactOpts, isBeacon, instHash, blkHeight, instPath, instPathIsLeft, instRoot, blkData, sigIdx, sigV, sigR, sigS)
}

// InstructionApproved is a paid mutator transaction binding the contract method 0xf65d2116.
//
// Solidity: function instructionApproved(bool isBeacon, bytes32 instHash, uint256 blkHeight, bytes32[] instPath, bool[] instPathIsLeft, bytes32 instRoot, bytes32 blkData, uint256[] sigIdx, uint8[] sigV, bytes32[] sigR, bytes32[] sigS) returns(bool)
func (_IncognitoProxy *IncognitoProxyTransactorSession) InstructionApproved(isBeacon bool, instHash [32]byte, blkHeight *big.Int, instPath [][32]byte, instPathIsLeft []bool, instRoot [32]byte, blkData [32]byte, sigIdx []*big.Int, sigV []uint8, sigR [][32]byte, sigS [][32]byte) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.InstructionApproved(&_IncognitoProxy.TransactOpts, isBeacon, instHash, blkHeight, instPath, instPathIsLeft, instRoot, blkData, sigIdx, sigV, sigR, sigS)
}

// InstructionInMerkleTree is a paid mutator transaction binding the contract method 0x47c4b328.
//
// Solidity: function instructionInMerkleTree(bytes32 leaf, bytes32 root, bytes32[] path, bool[] left) returns(bool)
func (_IncognitoProxy *IncognitoProxyTransactor) InstructionInMerkleTree(opts *bind.TransactOpts, leaf [32]byte, root [32]byte, path [][32]byte, left []bool) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "instructionInMerkleTree", leaf, root, path, left)
}

// InstructionInMerkleTree is a paid mutator transaction binding the contract method 0x47c4b328.
//
// Solidity: function instructionInMerkleTree(bytes32 leaf, bytes32 root, bytes32[] path, bool[] left) returns(bool)
func (_IncognitoProxy *IncognitoProxySession) InstructionInMerkleTree(leaf [32]byte, root [32]byte, path [][32]byte, left []bool) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.InstructionInMerkleTree(&_IncognitoProxy.TransactOpts, leaf, root, path, left)
}

// InstructionInMerkleTree is a paid mutator transaction binding the contract method 0x47c4b328.
//
// Solidity: function instructionInMerkleTree(bytes32 leaf, bytes32 root, bytes32[] path, bool[] left) returns(bool)
func (_IncognitoProxy *IncognitoProxyTransactorSession) InstructionInMerkleTree(leaf [32]byte, root [32]byte, path [][32]byte, left []bool) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.InstructionInMerkleTree(&_IncognitoProxy.TransactOpts, leaf, root, path, left)
}

// SwapBeaconCommittee is a paid mutator transaction binding the contract method 0xe41be775.
//
// Solidity: function swapBeaconCommittee(bytes inst, bytes32[] instPath, bool[] instPathIsLeft, bytes32 instRoot, bytes32 blkData, uint256[] sigIdx, uint8[] sigV, bytes32[] sigR, bytes32[] sigS) returns()
func (_IncognitoProxy *IncognitoProxyTransactor) SwapBeaconCommittee(opts *bind.TransactOpts, inst []byte, instPath [][32]byte, instPathIsLeft []bool, instRoot [32]byte, blkData [32]byte, sigIdx []*big.Int, sigV []uint8, sigR [][32]byte, sigS [][32]byte) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "swapBeaconCommittee", inst, instPath, instPathIsLeft, instRoot, blkData, sigIdx, sigV, sigR, sigS)
}

// SwapBeaconCommittee is a paid mutator transaction binding the contract method 0xe41be775.
//
// Solidity: function swapBeaconCommittee(bytes inst, bytes32[] instPath, bool[] instPathIsLeft, bytes32 instRoot, bytes32 blkData, uint256[] sigIdx, uint8[] sigV, bytes32[] sigR, bytes32[] sigS) returns()
func (_IncognitoProxy *IncognitoProxySession) SwapBeaconCommittee(inst []byte, instPath [][32]byte, instPathIsLeft []bool, instRoot [32]byte, blkData [32]byte, sigIdx []*big.Int, sigV []uint8, sigR [][32]byte, sigS [][32]byte) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.SwapBeaconCommittee(&_IncognitoProxy.TransactOpts, inst, instPath, instPathIsLeft, instRoot, blkData, sigIdx, sigV, sigR, sigS)
}

// SwapBeaconCommittee is a paid mutator transaction binding the contract method 0xe41be775.
//
// Solidity: function swapBeaconCommittee(bytes inst, bytes32[] instPath, bool[] instPathIsLeft, bytes32 instRoot, bytes32 blkData, uint256[] sigIdx, uint8[] sigV, bytes32[] sigR, bytes32[] sigS) returns()
func (_IncognitoProxy *IncognitoProxyTransactorSession) SwapBeaconCommittee(inst []byte, instPath [][32]byte, instPathIsLeft []bool, instRoot [32]byte, blkData [32]byte, sigIdx []*big.Int, sigV []uint8, sigR [][32]byte, sigS [][32]byte) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.SwapBeaconCommittee(&_IncognitoProxy.TransactOpts, inst, instPath, instPathIsLeft, instRoot, blkData, sigIdx, sigV, sigR, sigS)
}

// SwapBridgeCommittee is a paid mutator transaction binding the contract method 0x262f7220.
//
// Solidity: function swapBridgeCommittee(bytes inst, bytes32[][2] instPaths, bool[][2] instPathIsLefts, bytes32[2] instRoots, bytes32[2] blkData, uint256[][2] sigIdxs, uint8[][2] sigVs, bytes32[][2] sigRs, bytes32[][2] sigSs) returns()
func (_IncognitoProxy *IncognitoProxyTransactor) SwapBridgeCommittee(opts *bind.TransactOpts, inst []byte, instPaths [2][][32]byte, instPathIsLefts [2][]bool, instRoots [2][32]byte, blkData [2][32]byte, sigIdxs [2][]*big.Int, sigVs [2][]uint8, sigRs [2][][32]byte, sigSs [2][][32]byte) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "swapBridgeCommittee", inst, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// SwapBridgeCommittee is a paid mutator transaction binding the contract method 0x262f7220.
//
// Solidity: function swapBridgeCommittee(bytes inst, bytes32[][2] instPaths, bool[][2] instPathIsLefts, bytes32[2] instRoots, bytes32[2] blkData, uint256[][2] sigIdxs, uint8[][2] sigVs, bytes32[][2] sigRs, bytes32[][2] sigSs) returns()
func (_IncognitoProxy *IncognitoProxySession) SwapBridgeCommittee(inst []byte, instPaths [2][][32]byte, instPathIsLefts [2][]bool, instRoots [2][32]byte, blkData [2][32]byte, sigIdxs [2][]*big.Int, sigVs [2][]uint8, sigRs [2][][32]byte, sigSs [2][][32]byte) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.SwapBridgeCommittee(&_IncognitoProxy.TransactOpts, inst, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// SwapBridgeCommittee is a paid mutator transaction binding the contract method 0x262f7220.
//
// Solidity: function swapBridgeCommittee(bytes inst, bytes32[][2] instPaths, bool[][2] instPathIsLefts, bytes32[2] instRoots, bytes32[2] blkData, uint256[][2] sigIdxs, uint8[][2] sigVs, bytes32[][2] sigRs, bytes32[][2] sigSs) returns()
func (_IncognitoProxy *IncognitoProxyTransactorSession) SwapBridgeCommittee(inst []byte, instPaths [2][][32]byte, instPathIsLefts [2][]bool, instRoots [2][32]byte, blkData [2][32]byte, sigIdxs [2][]*big.Int, sigVs [2][]uint8, sigRs [2][][32]byte, sigSs [2][][32]byte) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.SwapBridgeCommittee(&_IncognitoProxy.TransactOpts, inst, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// IncognitoProxyLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the IncognitoProxy contract.
type IncognitoProxyLogAddressIterator struct {
	Event *IncognitoProxyLogAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncognitoProxyLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoProxyLogAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncognitoProxyLogAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncognitoProxyLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoProxyLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoProxyLogAddress represents a LogAddress event raised by the IncognitoProxy contract.
type IncognitoProxyLogAddress struct {
	Val common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0xb123f68b8ba02b447d91a6629e121111b7dd6061ff418a60139c8bf00522a284.
//
// Solidity: event LogAddress(address val)
func (_IncognitoProxy *IncognitoProxyFilterer) FilterLogAddress(opts *bind.FilterOpts) (*IncognitoProxyLogAddressIterator, error) {

	logs, sub, err := _IncognitoProxy.contract.FilterLogs(opts, "LogAddress")
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyLogAddressIterator{contract: _IncognitoProxy.contract, event: "LogAddress", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0xb123f68b8ba02b447d91a6629e121111b7dd6061ff418a60139c8bf00522a284.
//
// Solidity: event LogAddress(address val)
func (_IncognitoProxy *IncognitoProxyFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *IncognitoProxyLogAddress) (event.Subscription, error) {

	logs, sub, err := _IncognitoProxy.contract.WatchLogs(opts, "LogAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoProxyLogAddress)
				if err := _IncognitoProxy.contract.UnpackLog(event, "LogAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogAddress is a log parse operation binding the contract event 0xb123f68b8ba02b447d91a6629e121111b7dd6061ff418a60139c8bf00522a284.
//
// Solidity: event LogAddress(address val)
func (_IncognitoProxy *IncognitoProxyFilterer) ParseLogAddress(log types.Log) (*IncognitoProxyLogAddress, error) {
	event := new(IncognitoProxyLogAddress)
	if err := _IncognitoProxy.contract.UnpackLog(event, "LogAddress", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoProxyLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the IncognitoProxy contract.
type IncognitoProxyLogBytes32Iterator struct {
	Event *IncognitoProxyLogBytes32 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncognitoProxyLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoProxyLogBytes32)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncognitoProxyLogBytes32)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncognitoProxyLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoProxyLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoProxyLogBytes32 represents a LogBytes32 event raised by the IncognitoProxy contract.
type IncognitoProxyLogBytes32 struct {
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0x009fd52f05c0ded31d6fb0ee580b923f85e99cf1a5a1da342f25e73c45829c83.
//
// Solidity: event LogBytes32(bytes32 val)
func (_IncognitoProxy *IncognitoProxyFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*IncognitoProxyLogBytes32Iterator, error) {

	logs, sub, err := _IncognitoProxy.contract.FilterLogs(opts, "LogBytes32")
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyLogBytes32Iterator{contract: _IncognitoProxy.contract, event: "LogBytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0x009fd52f05c0ded31d6fb0ee580b923f85e99cf1a5a1da342f25e73c45829c83.
//
// Solidity: event LogBytes32(bytes32 val)
func (_IncognitoProxy *IncognitoProxyFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *IncognitoProxyLogBytes32) (event.Subscription, error) {

	logs, sub, err := _IncognitoProxy.contract.WatchLogs(opts, "LogBytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoProxyLogBytes32)
				if err := _IncognitoProxy.contract.UnpackLog(event, "LogBytes32", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogBytes32 is a log parse operation binding the contract event 0x009fd52f05c0ded31d6fb0ee580b923f85e99cf1a5a1da342f25e73c45829c83.
//
// Solidity: event LogBytes32(bytes32 val)
func (_IncognitoProxy *IncognitoProxyFilterer) ParseLogBytes32(log types.Log) (*IncognitoProxyLogBytes32, error) {
	event := new(IncognitoProxyLogBytes32)
	if err := _IncognitoProxy.contract.UnpackLog(event, "LogBytes32", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoProxyLogStringIterator is returned from FilterLogString and is used to iterate over the raw logs and unpacked data for LogString events raised by the IncognitoProxy contract.
type IncognitoProxyLogStringIterator struct {
	Event *IncognitoProxyLogString // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncognitoProxyLogStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoProxyLogString)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncognitoProxyLogString)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncognitoProxyLogStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoProxyLogStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoProxyLogString represents a LogString event raised by the IncognitoProxy contract.
type IncognitoProxyLogString struct {
	Val string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogString is a free log retrieval operation binding the contract event 0xa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f.
//
// Solidity: event LogString(string val)
func (_IncognitoProxy *IncognitoProxyFilterer) FilterLogString(opts *bind.FilterOpts) (*IncognitoProxyLogStringIterator, error) {

	logs, sub, err := _IncognitoProxy.contract.FilterLogs(opts, "LogString")
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyLogStringIterator{contract: _IncognitoProxy.contract, event: "LogString", logs: logs, sub: sub}, nil
}

// WatchLogString is a free log subscription operation binding the contract event 0xa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f.
//
// Solidity: event LogString(string val)
func (_IncognitoProxy *IncognitoProxyFilterer) WatchLogString(opts *bind.WatchOpts, sink chan<- *IncognitoProxyLogString) (event.Subscription, error) {

	logs, sub, err := _IncognitoProxy.contract.WatchLogs(opts, "LogString")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoProxyLogString)
				if err := _IncognitoProxy.contract.UnpackLog(event, "LogString", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogString is a log parse operation binding the contract event 0xa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f.
//
// Solidity: event LogString(string val)
func (_IncognitoProxy *IncognitoProxyFilterer) ParseLogString(log types.Log) (*IncognitoProxyLogString, error) {
	event := new(IncognitoProxyLogString)
	if err := _IncognitoProxy.contract.UnpackLog(event, "LogString", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoProxyLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the IncognitoProxy contract.
type IncognitoProxyLogUintIterator struct {
	Event *IncognitoProxyLogUint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IncognitoProxyLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoProxyLogUint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IncognitoProxyLogUint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IncognitoProxyLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoProxyLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoProxyLogUint represents a LogUint event raised by the IncognitoProxy contract.
type IncognitoProxyLogUint struct {
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x0ac68d08c5119b8cdb4058edbf0d4168f208ec3935d26a8f1f0d92eb9d4de8bf.
//
// Solidity: event LogUint(uint256 val)
func (_IncognitoProxy *IncognitoProxyFilterer) FilterLogUint(opts *bind.FilterOpts) (*IncognitoProxyLogUintIterator, error) {

	logs, sub, err := _IncognitoProxy.contract.FilterLogs(opts, "LogUint")
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyLogUintIterator{contract: _IncognitoProxy.contract, event: "LogUint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x0ac68d08c5119b8cdb4058edbf0d4168f208ec3935d26a8f1f0d92eb9d4de8bf.
//
// Solidity: event LogUint(uint256 val)
func (_IncognitoProxy *IncognitoProxyFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *IncognitoProxyLogUint) (event.Subscription, error) {

	logs, sub, err := _IncognitoProxy.contract.WatchLogs(opts, "LogUint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoProxyLogUint)
				if err := _IncognitoProxy.contract.UnpackLog(event, "LogUint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogUint is a log parse operation binding the contract event 0x0ac68d08c5119b8cdb4058edbf0d4168f208ec3935d26a8f1f0d92eb9d4de8bf.
//
// Solidity: event LogUint(uint256 val)
func (_IncognitoProxy *IncognitoProxyFilterer) ParseLogUint(log types.Log) (*IncognitoProxyLogUint, error) {
	event := new(IncognitoProxyLogUint)
	if err := _IncognitoProxy.contract.UnpackLog(event, "LogUint", log); err != nil {
		return nil, err
	}
	return event, nil
}
