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
const IncognitoProxyABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"}],\"name\":\"parseBurnInstruction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"instPaths\",\"type\":\"bytes32[][2]\"},{\"internalType\":\"bool[][2]\",\"name\":\"instPathIsLefts\",\"type\":\"bool[][2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"instRoots\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"blkData\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256[][2]\",\"name\":\"sigIdxs\",\"type\":\"uint256[][2]\"},{\"internalType\":\"uint8[][2]\",\"name\":\"sigVs\",\"type\":\"uint8[][2]\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"sigRs\",\"type\":\"bytes32[][2]\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"sigSs\",\"type\":\"bytes32[][2]\"}],\"name\":\"swapBridgeCommittee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"committee\",\"type\":\"address[]\"},{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"verifySig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"left\",\"type\":\"bool[]\"}],\"name\":\"instructionInMerkleTree\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"numVals\",\"type\":\"uint256\"}],\"name\":\"extractCommitteeFromInstruction\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"}],\"name\":\"extractMetaFromInstruction\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bridgeCommittees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blkHeight\",\"type\":\"uint256\"}],\"name\":\"findBeaconCommitteeFromHeight\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"committee\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"instPath\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"instPathIsLeft\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"instRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdx\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigS\",\"type\":\"bytes32[]\"}],\"name\":\"swapBeaconCommittee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"beaconCommittees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blkHeight\",\"type\":\"uint256\"}],\"name\":\"findBridgeCommitteeFromHeight\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"committee\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isBeacon\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"instHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blkHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"instPath\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"instPathIsLeft\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"instRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdx\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigS\",\"type\":\"bytes32[]\"}],\"name\":\"instructionApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"beaconCommittee\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"bridgeCommittee\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"LogUint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"val\",\"type\":\"string\"}],\"name\":\"LogString\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"val\",\"type\":\"bytes32\"}],\"name\":\"LogBytes32\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"LogAddress\",\"type\":\"event\"}]"

// IncognitoProxyFuncSigs maps the 4-byte function signature to its string representation.
var IncognitoProxyFuncSigs = map[string]string{
	"f203a5ed": "beaconCommittees(uint256)",
	"9b30b637": "bridgeCommittees(uint256)",
	"8eb60066": "extractCommitteeFromInstruction(bytes,uint256)",
	"90500bae": "extractMetaFromInstruction(bytes)",
	"b600ffdb": "findBeaconCommitteeFromHeight(uint256)",
	"f5205fde": "findBridgeCommitteeFromHeight(uint256)",
	"f65d2116": "instructionApproved(bool,bytes32,uint256,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])",
	"47c4b328": "instructionInMerkleTree(bytes32,bytes32,bytes32[],bool[])",
	"036010bd": "parseBurnInstruction(bytes)",
	"e41be775": "swapBeaconCommittee(bytes,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])",
	"262f7220": "swapBridgeCommittee(bytes,bytes32[][2],bool[][2],bytes32[2],bytes32[2],uint256[][2],uint8[][2],bytes32[][2],bytes32[][2])",
	"3aacfdad": "verifySig(address[],bytes32,uint8[],bytes32[],bytes32[])",
}

// IncognitoProxyBin is the compiled bytecode used for deploying new contracts.
var IncognitoProxyBin = "0x60806040523480156200001157600080fd5b5060405162001ebe38038062001ebe83398101604081905262000034916200024b565b60408051808201909152828152600060208083018290528154600181018084559280528351805193949360029092027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563019262000097928492909101906200011b565b5060209182015160019182015560408051808201909152848152600081840181905282548084018085559390915281518051939550919360029091027fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf601926200010592849201906200011b565b506020820151816001015550505050506200032d565b82805482825590600052602060002090810192821562000173579160200282015b828111156200017357825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906200013c565b506200018192915062000185565b5090565b620001ac91905b80821115620001815780546001600160a01b03191681556001016200018c565b90565b8051620001bc8162000313565b92915050565b600082601f830112620001d457600080fd5b8151620001eb620001e582620002e0565b620002b9565b915081818352602084019350602081019050838560208402820111156200021157600080fd5b60005b838110156200024157816200022a8882620001af565b845250602092830192919091019060010162000214565b5050505092915050565b600080604083850312156200025f57600080fd5b82516001600160401b038111156200027657600080fd5b6200028485828601620001c2565b92505060208301516001600160401b03811115620002a157600080fd5b620002af85828601620001c2565b9150509250929050565b6040518181016001600160401b0381118282101715620002d857600080fd5b604052919050565b60006001600160401b03821115620002f757600080fd5b5060209081020190565b60006001600160a01b038216620001bc565b6200031e8162000301565b81146200032a57600080fd5b50565b611b81806200033d6000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80639b30b637116100715780639b30b63714610170578063b600ffdb14610190578063e41be775146101a3578063f203a5ed146101b6578063f5205fde146101c9578063f65d2116146101dc576100b4565b8063036010bd146100b9578063262f7220146100e55780633aacfdad146100fa57806347c4b3281461011a5780638eb600661461012d57806390500bae1461014d575b600080fd5b6100cc6100c736600461152b565b6101ef565b6040516100dc9493929190611a0b565b60405180910390f35b6100f86100f3366004611567565b610284565b005b61010d610108366004611254565b61049d565b6040516100dc91906119aa565b61010d61012836600461149c565b61059b565b61014061013b366004611821565b610698565b6040516100dc9190611992565b61016061015b36600461152b565b61072d565b6040516100dc9493929190611a33565b61018361017e366004611871565b610795565b6040516100dc91906119fd565b61014061019e366004611871565b6107bb565b6100f86101b13660046116cd565b61087e565b6101836101c4366004611871565b6109a8565b6101406101d7366004611871565b6109b5565b61010d6101ea366004611324565b610a02565b60008060008060008560008151811061020457fe5b602001015160f81c60f81b60f81c60ff1662010000028660018151811061022757fe5b602001015160f81c60f81b60f81c60ff16610100028760028151811061024957fe5b602091010151602389015160438a01516063909a015160f89290921c9290920161ffff169290920162ffffff16989097965090945092505050565b885160208a0120600080546102e59160019184919060001981019081106102a757fe5b9060005260206000209060020201600101548c6000600281106102c657fe5b60200201518c518c518c518c518c518c518c60005b6020020151610a02565b6102ee57600080fd5b6103636000826001808080549050038154811061030757fe5b9060005260206000209060020201600101548c60016002811061032657fe5b60200201518c600160200201518c600160200201518c600160200201518c600160200201518c600160200201518c600160200201518c60016102db565b61036c57600080fd5b60008060008061037b8e61072d565b93509350935093508360ff16604714801561039957508260ff166001145b6103a257600080fd5b60606103ae8f83610698565b6040805180820190915281815260208082018690526001805480820180835560009290925283518051959650919460029091027fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60192610412928492910190610ba5565b50602082015181600101555050507f0ac68d08c5119b8cdb4058edbf0d4168f208ec3935d26a8f1f0d92eb9d4de8bf8360405161044f91906119fd565b60405180910390a17fa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f604051610484906119ed565b60405180910390a1505050505050505050505050505050565b600082518451146104ad57600080fd5b81518451146104bb57600080fd5b60005b845181101561058c578681815181106104d357fe5b60200260200101516001600160a01b03166001878784815181106104f357fe5b602002602001015187858151811061050757fe5b602002602001015187868151811061051b57fe5b60200260200101516040516000815260200160405260405161054094939291906119b8565b6020604051602081039080840390855afa158015610562573d6000803e3d6000fd5b505050602060405103516001600160a01b031614610584576000915050610592565b6001016104be565b50600190505b95945050505050565b600082518251146105ab57600080fd5b8460005b845181101561068c578381815181106105c457fe5b602002602001015115610614578481815181106105dd57fe5b6020026020010151826040516020016105f792919061196c565b604051602081830303815290604052805190602001209150610684565b84818151811061062057fe5b60200260200101516000801b14156106455781826040516020016105f792919061196c565b8185828151811061065257fe5b602002602001015160405160200161066b92919061196c565b6040516020818303038152906040528051906020012091505b6001016105af565b50909314949350505050565b6060816020026042018351146106ad57600080fd5b6060826040519080825280602002602001820160405280156106d9578160200160208202803883390190505b5090506000805b84811015610721576020810260628701015191508183828151811061070157fe5b6001600160a01b03909216602092830291909101909101526001016106e0565b50909150505b92915050565b60008060008060428551101561074257600080fd5b60008560008151811061075157fe5b602001015160f81c60f81b60f81c905060008660018151811061077057fe5b01602001516022880151604290980151929860f89190911c9796509194509092505050565b600181815481106107a257fe5b6000918252602090912060016002909202010154905081565b600054606090805b80156108765783600060018303815481106107da57fe5b9060005260206000209060020201600101541161086d576000600182038154811061080157fe5b6000918252602091829020600290910201805460408051828502810185019091528181529283018282801561085f57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610841575b505050505092505050610879565b600019016107c3565b50505b919050565b885160208a0120600080546108c09160019184919060001981019081106108a157fe5b9060005260206000209060020201600101548c8c8c8c8c8c8c8c610a02565b6108c957600080fd5b6000806000806108d88e61072d565b93509350935093508360ff1660461480156108f657508260ff166001145b6108ff57600080fd5b606061090b8f83610698565b604080518082019091528181526020808201869052600080546001810180835591805283518051959650919460029091027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563019261096d928492910190610ba5565b50602082015181600101555050507fa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f604051610484906119ed565b600081815481106107a257fe5b600154606090805b801561087657836001808303815481106109d357fe5b906000526020600020906002020160010154116109f95760018082038154811061080157fe5b600019016109bd565b600060608c15610a1c57610a158b6107bb565b9050610a28565b610a258b6109b5565b90505b60005b8651811015610aef57600081118015610a6d5750866001820381518110610a4e57fe5b6020026020010151878281518110610a6257fe5b602002602001015111155b80610a8c57508151878281518110610a8157fe5b602002602001015110155b15610a9c57600092505050610b96565b81878281518110610aa957fe5b602002602001015181518110610abb57fe5b6020026020010151828281518110610acf57fe5b6001600160a01b0390921660209283029190910190910152600101610a2b565b5060008789604051602001610b0592919061196c565b60405160208183030381529060405280519060200120604051602001610b2b9190611957565b6040516020818303038152906040528051906020012090506003825160020281610b5157fe5b04875111610b6457600092505050610b96565b610b71828288888861049d565b610b7a57600080fd5b610b868d8a8d8d61059b565b610b8f57600080fd5b6001925050505b9b9a5050505050505050505050565b828054828255906000526020600020908101928215610bfa579160200282015b82811115610bfa57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190610bc5565b50610c06929150610c0a565b5090565b610c2e91905b80821115610c065780546001600160a01b0319168155600101610c10565b90565b803561072781611b0c565b600082601f830112610c4d57600080fd5b8135610c60610c5b82611a67565b611a41565b91508181835260208401935060208101905083856020840282011115610c8557600080fd5b60005b83811015610cb15781610c9b8882610c31565b8452506020928301929190910190600101610c88565b5050505092915050565b600082601f830112610ccc57600080fd5b6002610cda610c5b82611a87565b9150818360005b83811015610cb15781358601610cf78882610e03565b8452506020928301929190910190600101610ce1565b600082601f830112610d1e57600080fd5b6002610d2c610c5b82611a87565b9150818360005b83811015610cb15781358601610d498882610f44565b8452506020928301929190910190600101610d33565b600082601f830112610d7057600080fd5b6002610d7e610c5b82611a87565b9150818360005b83811015610cb15781358601610d9b8882611024565b8452506020928301929190910190600101610d85565b600082601f830112610dc257600080fd5b6002610dd0610c5b82611a87565b9150818360005b83811015610cb15781358601610ded8882611104565b8452506020928301929190910190600101610dd7565b600082601f830112610e1457600080fd5b8135610e22610c5b82611a67565b91508181835260208401935060208101905083856020840282011115610e4757600080fd5b60005b83811015610cb15781610e5d88826111e4565b8452506020928301929190910190600101610e4a565b600082601f830112610e8457600080fd5b8135610e92610c5b82611a67565b91508181835260208401935060208101905083856020840282011115610eb757600080fd5b60005b83811015610cb15781610ecd88826111e4565b8452506020928301929190910190600101610eba565b600082601f830112610ef457600080fd5b6002610f02610c5b82611a87565b91508183856020840282011115610f1857600080fd5b60005b83811015610cb15781610f2e88826111ef565b8452506020928301929190910190600101610f1b565b600082601f830112610f5557600080fd5b8135610f63610c5b82611a67565b91508181835260208401935060208101905083856020840282011115610f8857600080fd5b60005b83811015610cb15781610f9e88826111ef565b8452506020928301929190910190600101610f8b565b600082601f830112610fc557600080fd5b8135610fd3610c5b82611a67565b91508181835260208401935060208101905083856020840282011115610ff857600080fd5b60005b83811015610cb1578161100e88826111ef565b8452506020928301929190910190600101610ffb565b600082601f83011261103557600080fd5b8135611043610c5b82611a67565b9150818183526020840193506020810190508385602084028201111561106857600080fd5b60005b83811015610cb1578161107e88826111ef565b845250602092830192919091019060010161106b565b600082601f8301126110a557600080fd5b81356110b3610c5b82611a67565b915081818352602084019350602081019050838560208402820111156110d857600080fd5b60005b83811015610cb157816110ee88826111ef565b84525060209283019291909101906001016110db565b600082601f83011261111557600080fd5b8135611123610c5b82611a67565b9150818183526020840193506020810190508385602084028201111561114857600080fd5b60005b83811015610cb1578161115e8882611249565b845250602092830192919091019060010161114b565b600082601f83011261118557600080fd5b8135611193610c5b82611a67565b915081818352602084019350602081019050838560208402820111156111b857600080fd5b60005b83811015610cb157816111ce8882611249565b84525060209283019291909101906001016111bb565b803561072781611b23565b803561072781611b2c565b600082601f83011261120b57600080fd5b8135611219610c5b82611aa4565b9150808252602083016020830185838301111561123557600080fd5b611240838284611b00565b50505092915050565b803561072781611b35565b600080600080600060a0868803121561126c57600080fd5b85356001600160401b0381111561128257600080fd5b61128e88828901610c3c565b955050602061129f888289016111ef565b94505060408601356001600160401b038111156112bb57600080fd5b6112c788828901611174565b93505060608601356001600160401b038111156112e357600080fd5b6112ef88828901610fb4565b92505060808601356001600160401b0381111561130b57600080fd5b61131788828901610fb4565b9150509295509295909350565b60008060008060008060008060008060006101608c8e03121561134657600080fd5b60006113528e8e6111e4565b9b505060206113638e828f016111ef565b9a505060406113748e828f016111ef565b99505060608c01356001600160401b0381111561139057600080fd5b61139c8e828f01610fb4565b98505060808c01356001600160401b038111156113b857600080fd5b6113c48e828f01610e73565b97505060a06113d58e828f016111ef565b96505060c06113e68e828f016111ef565b95505060e08c01356001600160401b0381111561140257600080fd5b61140e8e828f01611094565b9450506101008c01356001600160401b0381111561142b57600080fd5b6114378e828f01611174565b9350506101208c01356001600160401b0381111561145457600080fd5b6114608e828f01610fb4565b9250506101408c01356001600160401b0381111561147d57600080fd5b6114898e828f01610fb4565b9150509295989b509295989b9093969950565b600080600080608085870312156114b257600080fd5b60006114be87876111ef565b94505060206114cf878288016111ef565b93505060408501356001600160401b038111156114eb57600080fd5b6114f787828801610fb4565b92505060608501356001600160401b0381111561151357600080fd5b61151f87828801610e73565b91505092959194509250565b60006020828403121561153d57600080fd5b81356001600160401b0381111561155357600080fd5b61155f848285016111fa565b949350505050565b60008060008060008060008060006101608a8c03121561158657600080fd5b89356001600160401b0381111561159c57600080fd5b6115a88c828d016111fa565b99505060208a01356001600160401b038111156115c457600080fd5b6115d08c828d01610d0d565b98505060408a01356001600160401b038111156115ec57600080fd5b6115f88c828d01610cbb565b97505060606116098c828d01610ee3565b96505060a061161a8c828d01610ee3565b95505060e08a01356001600160401b0381111561163657600080fd5b6116428c828d01610d5f565b9450506101008a01356001600160401b0381111561165f57600080fd5b61166b8c828d01610db1565b9350506101208a01356001600160401b0381111561168857600080fd5b6116948c828d01610d0d565b9250506101408a01356001600160401b038111156116b157600080fd5b6116bd8c828d01610d0d565b9150509295985092959850929598565b60008060008060008060008060006101208a8c0312156116ec57600080fd5b89356001600160401b0381111561170257600080fd5b61170e8c828d016111fa565b99505060208a01356001600160401b0381111561172a57600080fd5b6117368c828d01610fb4565b98505060408a01356001600160401b0381111561175257600080fd5b61175e8c828d01610e73565b975050606061176f8c828d016111ef565b96505060806117808c828d016111ef565b95505060a08a01356001600160401b0381111561179c57600080fd5b6117a88c828d01611094565b94505060c08a01356001600160401b038111156117c457600080fd5b6117d08c828d01611174565b93505060e08a01356001600160401b038111156117ec57600080fd5b6117f88c828d01610fb4565b9250506101008a01356001600160401b0381111561181557600080fd5b6116bd8c828d01610fb4565b6000806040838503121561183457600080fd5b82356001600160401b0381111561184a57600080fd5b611856858286016111fa565b9250506020611867858286016111ef565b9150509250929050565b60006020828403121561188357600080fd5b600061155f84846111ef565b600061189b83836118a3565b505060200190565b6118ac81611ade565b82525050565b60006118bd82611ad1565b6118c78185611ad5565b93506118d283611acb565b8060005b838110156119005781516118ea888261188f565b97506118f583611acb565b9250506001016118d6565b509495945050505050565b6118ac81611ae9565b6118ac81610c2e565b6118ac61192982610c2e565b610c2e565b600061193b600483611ad5565b63446f6e6560e01b815260200192915050565b6118ac81611afa565b6000611963828461191d565b50602001919050565b6000611978828561191d565b602082019150611988828461191d565b5060200192915050565b602080825281016119a381846118b2565b9392505050565b60208101610727828461190b565b608081016119c68287611914565b6119d3602083018661194e565b6119e06040830185611914565b6105926060830184611914565b602080825281016107278161192e565b602081016107278284611914565b60808101611a198287611914565b611a2660208301866118a3565b6119e060408301856118a3565b608081016119c6828761194e565b6040518181016001600160401b0381118282101715611a5f57600080fd5b604052919050565b60006001600160401b03821115611a7d57600080fd5b5060209081020190565b60006001600160401b03821115611a9d57600080fd5b5060200290565b60006001600160401b03821115611aba57600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b600061072782611aee565b151590565b6001600160a01b031690565b60ff1690565b82818337506000910152565b611b1581611ade565b8114611b2057600080fd5b50565b611b1581611ae9565b611b1581610c2e565b611b1581611afa56fea365627a7a7231582091903543ff9fd5576c15eb819f1fb6c96c9f34b8ac94a91fb5f8b28e0a4d5aac6c6578706572696d656e74616cf564736f6c634300050b0040"

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
// Solidity: function extractMetaFromInstruction(bytes inst) constant returns(uint8, uint8, uint256, uint256)
func (_IncognitoProxy *IncognitoProxyCaller) ExtractMetaFromInstruction(opts *bind.CallOpts, inst []byte) (uint8, uint8, *big.Int, *big.Int, error) {
	var (
		ret0 = new(uint8)
		ret1 = new(uint8)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _IncognitoProxy.contract.Call(opts, out, "extractMetaFromInstruction", inst)
	return *ret0, *ret1, *ret2, *ret3, err
}

// ExtractMetaFromInstruction is a free data retrieval call binding the contract method 0x90500bae.
//
// Solidity: function extractMetaFromInstruction(bytes inst) constant returns(uint8, uint8, uint256, uint256)
func (_IncognitoProxy *IncognitoProxySession) ExtractMetaFromInstruction(inst []byte) (uint8, uint8, *big.Int, *big.Int, error) {
	return _IncognitoProxy.Contract.ExtractMetaFromInstruction(&_IncognitoProxy.CallOpts, inst)
}

// ExtractMetaFromInstruction is a free data retrieval call binding the contract method 0x90500bae.
//
// Solidity: function extractMetaFromInstruction(bytes inst) constant returns(uint8, uint8, uint256, uint256)
func (_IncognitoProxy *IncognitoProxyCallerSession) ExtractMetaFromInstruction(inst []byte) (uint8, uint8, *big.Int, *big.Int, error) {
	return _IncognitoProxy.Contract.ExtractMetaFromInstruction(&_IncognitoProxy.CallOpts, inst)
}

// FindBeaconCommitteeFromHeight is a free data retrieval call binding the contract method 0xb600ffdb.
//
// Solidity: function findBeaconCommitteeFromHeight(uint256 blkHeight) constant returns(address[] committee)
func (_IncognitoProxy *IncognitoProxyCaller) FindBeaconCommitteeFromHeight(opts *bind.CallOpts, blkHeight *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "findBeaconCommitteeFromHeight", blkHeight)
	return *ret0, err
}

// FindBeaconCommitteeFromHeight is a free data retrieval call binding the contract method 0xb600ffdb.
//
// Solidity: function findBeaconCommitteeFromHeight(uint256 blkHeight) constant returns(address[] committee)
func (_IncognitoProxy *IncognitoProxySession) FindBeaconCommitteeFromHeight(blkHeight *big.Int) ([]common.Address, error) {
	return _IncognitoProxy.Contract.FindBeaconCommitteeFromHeight(&_IncognitoProxy.CallOpts, blkHeight)
}

// FindBeaconCommitteeFromHeight is a free data retrieval call binding the contract method 0xb600ffdb.
//
// Solidity: function findBeaconCommitteeFromHeight(uint256 blkHeight) constant returns(address[] committee)
func (_IncognitoProxy *IncognitoProxyCallerSession) FindBeaconCommitteeFromHeight(blkHeight *big.Int) ([]common.Address, error) {
	return _IncognitoProxy.Contract.FindBeaconCommitteeFromHeight(&_IncognitoProxy.CallOpts, blkHeight)
}

// FindBridgeCommitteeFromHeight is a free data retrieval call binding the contract method 0xf5205fde.
//
// Solidity: function findBridgeCommitteeFromHeight(uint256 blkHeight) constant returns(address[] committee)
func (_IncognitoProxy *IncognitoProxyCaller) FindBridgeCommitteeFromHeight(opts *bind.CallOpts, blkHeight *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "findBridgeCommitteeFromHeight", blkHeight)
	return *ret0, err
}

// FindBridgeCommitteeFromHeight is a free data retrieval call binding the contract method 0xf5205fde.
//
// Solidity: function findBridgeCommitteeFromHeight(uint256 blkHeight) constant returns(address[] committee)
func (_IncognitoProxy *IncognitoProxySession) FindBridgeCommitteeFromHeight(blkHeight *big.Int) ([]common.Address, error) {
	return _IncognitoProxy.Contract.FindBridgeCommitteeFromHeight(&_IncognitoProxy.CallOpts, blkHeight)
}

// FindBridgeCommitteeFromHeight is a free data retrieval call binding the contract method 0xf5205fde.
//
// Solidity: function findBridgeCommitteeFromHeight(uint256 blkHeight) constant returns(address[] committee)
func (_IncognitoProxy *IncognitoProxyCallerSession) FindBridgeCommitteeFromHeight(blkHeight *big.Int) ([]common.Address, error) {
	return _IncognitoProxy.Contract.FindBridgeCommitteeFromHeight(&_IncognitoProxy.CallOpts, blkHeight)
}

// InstructionApproved is a free data retrieval call binding the contract method 0xf65d2116.
//
// Solidity: function instructionApproved(bool isBeacon, bytes32 instHash, uint256 blkHeight, bytes32[] instPath, bool[] instPathIsLeft, bytes32 instRoot, bytes32 blkData, uint256[] sigIdx, uint8[] sigV, bytes32[] sigR, bytes32[] sigS) constant returns(bool)
func (_IncognitoProxy *IncognitoProxyCaller) InstructionApproved(opts *bind.CallOpts, isBeacon bool, instHash [32]byte, blkHeight *big.Int, instPath [][32]byte, instPathIsLeft []bool, instRoot [32]byte, blkData [32]byte, sigIdx []*big.Int, sigV []uint8, sigR [][32]byte, sigS [][32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "instructionApproved", isBeacon, instHash, blkHeight, instPath, instPathIsLeft, instRoot, blkData, sigIdx, sigV, sigR, sigS)
	return *ret0, err
}

// InstructionApproved is a free data retrieval call binding the contract method 0xf65d2116.
//
// Solidity: function instructionApproved(bool isBeacon, bytes32 instHash, uint256 blkHeight, bytes32[] instPath, bool[] instPathIsLeft, bytes32 instRoot, bytes32 blkData, uint256[] sigIdx, uint8[] sigV, bytes32[] sigR, bytes32[] sigS) constant returns(bool)
func (_IncognitoProxy *IncognitoProxySession) InstructionApproved(isBeacon bool, instHash [32]byte, blkHeight *big.Int, instPath [][32]byte, instPathIsLeft []bool, instRoot [32]byte, blkData [32]byte, sigIdx []*big.Int, sigV []uint8, sigR [][32]byte, sigS [][32]byte) (bool, error) {
	return _IncognitoProxy.Contract.InstructionApproved(&_IncognitoProxy.CallOpts, isBeacon, instHash, blkHeight, instPath, instPathIsLeft, instRoot, blkData, sigIdx, sigV, sigR, sigS)
}

// InstructionApproved is a free data retrieval call binding the contract method 0xf65d2116.
//
// Solidity: function instructionApproved(bool isBeacon, bytes32 instHash, uint256 blkHeight, bytes32[] instPath, bool[] instPathIsLeft, bytes32 instRoot, bytes32 blkData, uint256[] sigIdx, uint8[] sigV, bytes32[] sigR, bytes32[] sigS) constant returns(bool)
func (_IncognitoProxy *IncognitoProxyCallerSession) InstructionApproved(isBeacon bool, instHash [32]byte, blkHeight *big.Int, instPath [][32]byte, instPathIsLeft []bool, instRoot [32]byte, blkData [32]byte, sigIdx []*big.Int, sigV []uint8, sigR [][32]byte, sigS [][32]byte) (bool, error) {
	return _IncognitoProxy.Contract.InstructionApproved(&_IncognitoProxy.CallOpts, isBeacon, instHash, blkHeight, instPath, instPathIsLeft, instRoot, blkData, sigIdx, sigV, sigR, sigS)
}

// InstructionInMerkleTree is a free data retrieval call binding the contract method 0x47c4b328.
//
// Solidity: function instructionInMerkleTree(bytes32 leaf, bytes32 root, bytes32[] path, bool[] left) constant returns(bool)
func (_IncognitoProxy *IncognitoProxyCaller) InstructionInMerkleTree(opts *bind.CallOpts, leaf [32]byte, root [32]byte, path [][32]byte, left []bool) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "instructionInMerkleTree", leaf, root, path, left)
	return *ret0, err
}

// InstructionInMerkleTree is a free data retrieval call binding the contract method 0x47c4b328.
//
// Solidity: function instructionInMerkleTree(bytes32 leaf, bytes32 root, bytes32[] path, bool[] left) constant returns(bool)
func (_IncognitoProxy *IncognitoProxySession) InstructionInMerkleTree(leaf [32]byte, root [32]byte, path [][32]byte, left []bool) (bool, error) {
	return _IncognitoProxy.Contract.InstructionInMerkleTree(&_IncognitoProxy.CallOpts, leaf, root, path, left)
}

// InstructionInMerkleTree is a free data retrieval call binding the contract method 0x47c4b328.
//
// Solidity: function instructionInMerkleTree(bytes32 leaf, bytes32 root, bytes32[] path, bool[] left) constant returns(bool)
func (_IncognitoProxy *IncognitoProxyCallerSession) InstructionInMerkleTree(leaf [32]byte, root [32]byte, path [][32]byte, left []bool) (bool, error) {
	return _IncognitoProxy.Contract.InstructionInMerkleTree(&_IncognitoProxy.CallOpts, leaf, root, path, left)
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
