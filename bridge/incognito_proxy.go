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
const IncognitoProxyABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"}],\"name\":\"parseBurnInstruction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"instPaths\",\"type\":\"bytes32[][2]\"},{\"internalType\":\"bool[][2]\",\"name\":\"instPathIsLefts\",\"type\":\"bool[][2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"instRoots\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"blkData\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256[][2]\",\"name\":\"sigIdxs\",\"type\":\"uint256[][2]\"},{\"internalType\":\"uint8[][2]\",\"name\":\"sigVs\",\"type\":\"uint8[][2]\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"sigRs\",\"type\":\"bytes32[][2]\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"sigSs\",\"type\":\"bytes32[][2]\"}],\"name\":\"swapBridgeCommittee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"committee\",\"type\":\"address[]\"},{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"verifySig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"left\",\"type\":\"bool[]\"}],\"name\":\"instructionInMerkleTree\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"numVals\",\"type\":\"uint256\"}],\"name\":\"extractCommitteeFromInstruction\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"}],\"name\":\"extractMetaFromInstruction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bridgeCommittees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"beaconCommittees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isBeacon\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"instHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blkHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"instPath\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"instPathIsLeft\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"instRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdx\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigS\",\"type\":\"bytes32[]\"}],\"name\":\"instructionApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blkHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isBeacon\",\"type\":\"bool\"}],\"name\":\"findCommitteeFromHeight\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"committee\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"beaconCommittee\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"bridgeCommittee\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"LogUint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"val\",\"type\":\"string\"}],\"name\":\"LogString\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"val\",\"type\":\"bytes32\"}],\"name\":\"LogBytes32\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"LogAddress\",\"type\":\"event\"}]"

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
	"262f7220": "swapBridgeCommittee(bytes,bytes32[][2],bool[][2],bytes32[2],bytes32[2],uint256[][2],uint8[][2],bytes32[][2],bytes32[][2])",
	"3aacfdad": "verifySig(address[],bytes32,uint8[],bytes32[],bytes32[])",
}

// IncognitoProxyBin is the compiled bytecode used for deploying new contracts.
var IncognitoProxyBin = "0x60806040523480156200001157600080fd5b5060405162001a4f38038062001a4f83398101604081905262000034916200024b565b60408051808201909152828152600060208083018290528154600181018084559280528351805193949360029092027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563019262000097928492909101906200011b565b5060209182015160019182015560408051808201909152848152600081840181905282548084018085559390915281518051939550919360029091027fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf601926200010592849201906200011b565b506020820151816001015550505050506200032d565b82805482825590600052602060002090810192821562000173579160200282015b828111156200017357825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906200013c565b506200018192915062000185565b5090565b620001ac91905b80821115620001815780546001600160a01b03191681556001016200018c565b90565b8051620001bc8162000313565b92915050565b600082601f830112620001d457600080fd5b8151620001eb620001e582620002e0565b620002b9565b915081818352602084019350602081019050838560208402820111156200021157600080fd5b60005b838110156200024157816200022a8882620001af565b845250602092830192919091019060010162000214565b5050505092915050565b600080604083850312156200025f57600080fd5b82516001600160401b038111156200027657600080fd5b6200028485828601620001c2565b92505060208301516001600160401b03811115620002a157600080fd5b620002af85828601620001c2565b9150509250929050565b6040518181016001600160401b0381118282101715620002d857600080fd5b604052919050565b60006001600160401b03821115620002f757600080fd5b5060209081020190565b60006001600160a01b038216620001bc565b6200031e8162000301565b81146200032a57600080fd5b50565b611712806200033d6000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c806390500bae1161006657806390500bae146101375780639b30b63714610159578063f203a5ed14610179578063f65d21161461018c578063ff2293f61461019f5761009e565b8063036010bd146100a3578063262f7220146100cf5780633aacfdad146100e457806347c4b328146101045780638eb6006614610117575b600080fd5b6100b66100b13660046111e0565b6101b2565b6040516100c6949392919061157f565b60405180910390f35b6100e26100dd366004611214565b610247565b005b6100f76100f2366004610f09565b610379565b6040516100c6919061151e565b6100f7610112366004611151565b61045a565b61012a61012536600461137a565b61054a565b6040516100c69190611506565b61014a6101453660046111e0565b6105cb565b6040516100c6939291906115a7565b61016c6101673660046113ca565b610655565b6040516100c69190611571565b61016c6101873660046113ca565b61067b565b6100f761019a366004610fd9565b610688565b61012a6101ad3660046113e8565b6107cd565b6000806000806000856000815181106101c757fe5b602001015160f81c60f81b60f81c60ff166201000002866001815181106101ea57fe5b602001015160f81c60f81b60f81c60ff16610100028760028151811061020c57fe5b602091010151602389015160438a01516063909a015160f89290921c9290920161ffff169290920162ffffff16989097965090945092505050565b885160208a0120600080548190600019810190811061026257fe5b90600052602060002090600202016001015490506102a9600183838d60006002811061028a57fe5b60200201518d518d518d518d518d518d518d60005b6020020151610688565b6102b257600080fd5b600180546000919060001981019081106102c857fe5b906000526020600020906002020160010154905061032d600084838e6001600281106102f057fe5b60200201518e600160200201518e600160200201518e600160200201518e600160200201518e600160200201518e600160200201518e600161029f565b61033657600080fd5b7fa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f60405161036390611561565b60405180910390a1505050505050505050505050565b6000805b845181101561044b5786818151811061039257fe5b60200260200101516001600160a01b03166001878784815181106103b257fe5b60200260200101518785815181106103c657fe5b60200260200101518786815181106103da57fe5b6020026020010151604051600081526020016040526040516103ff949392919061152c565b6020604051602081039080840390855afa158015610421573d6000803e3d6000fd5b505050602060405103516001600160a01b031614610443576000915050610451565b60010161037d565b50600190505b95945050505050565b600084815b845181101561053c5783818151811061047457fe5b6020026020010151156104c45784818151811061048d57fe5b6020026020010151826040516020016104a79291906114e0565b604051602081830303815290604052805190602001209150610534565b8481815181106104d057fe5b60200260200101516000801b14156104f55781826040516020016104a79291906114e0565b8185828151811061050257fe5b602002602001015160405160200161051b9291906114e0565b6040516020818303038152906040528051906020012091505b60010161045f565b50841490505b949350505050565b60608082604051908082528060200260200182016040528015610577578160200160208202803883390190505b5090506000805b848110156105bf576020810260638701015191508183828151811061059f57fe5b6001600160a01b039092166020928302919091019091015260010161057e565b50909150505b92915050565b600080600080846000815181106105de57fe5b602001015160f81c60f81b60f81c60ff1662010000028560018151811061060157fe5b602001015160f81c60f81b60f81c60ff16610100028660028151811061062357fe5b602091010151602388015160439098015160f89190911c9190910161ffff169190910162ffffff169690945092505050565b6001818154811061066257fe5b6000918252602090912060016002909202010154905081565b6000818154811061066257fe5b600060606106968b8e6107cd565b905060005b865181101561075f576000811180156106dd57508660018203815181106106be57fe5b60200260200101518782815181106106d257fe5b602002602001015111155b806106fc575081518782815181106106f157fe5b602002602001015110155b1561070c576000925050506107be565b8187828151811061071957fe5b60200260200101518151811061072b57fe5b602002602001015182828151811061073f57fe5b6001600160a01b039092166020928302919091019091015260010161069b565b50600381516002028161076e57fe5b048651116107805760009150506107be565b7f1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac86107ae8282888888610379565b6107b757600080fd5b6001925050505b9b9a5050505050505050505050565b6060811561089457600054805b801561088d5784600060018303815481106107f157fe5b90600052602060002090600202016001015411610884576000600182038154811061081857fe5b6000918252602091829020600290910201805460408051828502810185019091528181529283018282801561087657602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610858575b5050505050925050506105c5565b600019016107da565b50506105c5565b600154805b80156108de57846001808303815481106108af57fe5b906000526020600020906002020160010154116108d55760018082038154811061081857fe5b60001901610899565b505092915050565b80356105c58161169d565b600082601f83011261090257600080fd5b8135610915610910826115f5565b6115cf565b9150818183526020840193506020810190508385602084028201111561093a57600080fd5b60005b83811015610966578161095088826108e6565b845250602092830192919091019060010161093d565b5050505092915050565b600082601f83011261098157600080fd5b600261098f61091082611615565b9150818360005b8381101561096657813586016109ac8882610ab8565b8452506020928301929190910190600101610996565b600082601f8301126109d357600080fd5b60026109e161091082611615565b9150818360005b8381101561096657813586016109fe8882610bf9565b84525060209283019291909101906001016109e8565b600082601f830112610a2557600080fd5b6002610a3361091082611615565b9150818360005b838110156109665781358601610a508882610cd9565b8452506020928301929190910190600101610a3a565b600082601f830112610a7757600080fd5b6002610a8561091082611615565b9150818360005b838110156109665781358601610aa28882610db9565b8452506020928301929190910190600101610a8c565b600082601f830112610ac957600080fd5b8135610ad7610910826115f5565b91508181835260208401935060208101905083856020840282011115610afc57600080fd5b60005b838110156109665781610b128882610e99565b8452506020928301929190910190600101610aff565b600082601f830112610b3957600080fd5b8135610b47610910826115f5565b91508181835260208401935060208101905083856020840282011115610b6c57600080fd5b60005b838110156109665781610b828882610e99565b8452506020928301929190910190600101610b6f565b600082601f830112610ba957600080fd5b6002610bb761091082611615565b91508183856020840282011115610bcd57600080fd5b60005b838110156109665781610be38882610ea4565b8452506020928301929190910190600101610bd0565b600082601f830112610c0a57600080fd5b8135610c18610910826115f5565b91508181835260208401935060208101905083856020840282011115610c3d57600080fd5b60005b838110156109665781610c538882610ea4565b8452506020928301929190910190600101610c40565b600082601f830112610c7a57600080fd5b8135610c88610910826115f5565b91508181835260208401935060208101905083856020840282011115610cad57600080fd5b60005b838110156109665781610cc38882610ea4565b8452506020928301929190910190600101610cb0565b600082601f830112610cea57600080fd5b8135610cf8610910826115f5565b91508181835260208401935060208101905083856020840282011115610d1d57600080fd5b60005b838110156109665781610d338882610ea4565b8452506020928301929190910190600101610d20565b600082601f830112610d5a57600080fd5b8135610d68610910826115f5565b91508181835260208401935060208101905083856020840282011115610d8d57600080fd5b60005b838110156109665781610da38882610ea4565b8452506020928301929190910190600101610d90565b600082601f830112610dca57600080fd5b8135610dd8610910826115f5565b91508181835260208401935060208101905083856020840282011115610dfd57600080fd5b60005b838110156109665781610e138882610efe565b8452506020928301929190910190600101610e00565b600082601f830112610e3a57600080fd5b8135610e48610910826115f5565b91508181835260208401935060208101905083856020840282011115610e6d57600080fd5b60005b838110156109665781610e838882610efe565b8452506020928301929190910190600101610e70565b80356105c5816116b4565b80356105c5816116bd565b600082601f830112610ec057600080fd5b8135610ece61091082611632565b91508082526020830160208301858383011115610eea57600080fd5b610ef5838284611691565b50505092915050565b80356105c5816116c6565b600080600080600060a08688031215610f2157600080fd5b85356001600160401b03811115610f3757600080fd5b610f43888289016108f1565b9550506020610f5488828901610ea4565b94505060408601356001600160401b03811115610f7057600080fd5b610f7c88828901610e29565b93505060608601356001600160401b03811115610f9857600080fd5b610fa488828901610c69565b92505060808601356001600160401b03811115610fc057600080fd5b610fcc88828901610c69565b9150509295509295909350565b60008060008060008060008060008060006101608c8e031215610ffb57600080fd5b60006110078e8e610e99565b9b505060206110188e828f01610ea4565b9a505060406110298e828f01610ea4565b99505060608c01356001600160401b0381111561104557600080fd5b6110518e828f01610c69565b98505060808c01356001600160401b0381111561106d57600080fd5b6110798e828f01610b28565b97505060a061108a8e828f01610ea4565b96505060c061109b8e828f01610ea4565b95505060e08c01356001600160401b038111156110b757600080fd5b6110c38e828f01610d49565b9450506101008c01356001600160401b038111156110e057600080fd5b6110ec8e828f01610e29565b9350506101208c01356001600160401b0381111561110957600080fd5b6111158e828f01610c69565b9250506101408c01356001600160401b0381111561113257600080fd5b61113e8e828f01610c69565b9150509295989b509295989b9093969950565b6000806000806080858703121561116757600080fd5b60006111738787610ea4565b945050602061118487828801610ea4565b93505060408501356001600160401b038111156111a057600080fd5b6111ac87828801610c69565b92505060608501356001600160401b038111156111c857600080fd5b6111d487828801610b28565b91505092959194509250565b6000602082840312156111f257600080fd5b81356001600160401b0381111561120857600080fd5b61054284828501610eaf565b60008060008060008060008060006101608a8c03121561123357600080fd5b89356001600160401b0381111561124957600080fd5b6112558c828d01610eaf565b99505060208a01356001600160401b0381111561127157600080fd5b61127d8c828d016109c2565b98505060408a01356001600160401b0381111561129957600080fd5b6112a58c828d01610970565b97505060606112b68c828d01610b98565b96505060a06112c78c828d01610b98565b95505060e08a01356001600160401b038111156112e357600080fd5b6112ef8c828d01610a14565b9450506101008a01356001600160401b0381111561130c57600080fd5b6113188c828d01610a66565b9350506101208a01356001600160401b0381111561133557600080fd5b6113418c828d016109c2565b9250506101408a01356001600160401b0381111561135e57600080fd5b61136a8c828d016109c2565b9150509295985092959850929598565b6000806040838503121561138d57600080fd5b82356001600160401b038111156113a357600080fd5b6113af85828601610eaf565b92505060206113c085828601610ea4565b9150509250929050565b6000602082840312156113dc57600080fd5b60006105428484610ea4565b600080604083850312156113fb57600080fd5b60006114078585610ea4565b92505060206113c085828601610e99565b6000611424838361142c565b505060200190565b6114358161166c565b82525050565b60006114468261165f565b6114508185611663565b935061145b83611659565b8060005b838110156114895781516114738882611418565b975061147e83611659565b92505060010161145f565b509495945050505050565b61143581611677565b6114358161167c565b6114356114b28261167c565b61167c565b60006114c4600483611663565b63446f6e6560e01b815260200192915050565b6114358161168b565b60006114ec82856114a6565b6020820191506114fc82846114a6565b5060200192915050565b60208082528101611517818461143b565b9392505050565b602081016105c58284611494565b6080810161153a828761149d565b61154760208301866114d7565b611554604083018561149d565b610451606083018461149d565b602080825281016105c5816114b7565b602081016105c5828461149d565b6080810161158d828761149d565b61159a602083018661142c565b611554604083018561142c565b606081016115b5828661149d565b6115c2602083018561149d565b610542604083018461149d565b6040518181016001600160401b03811182821017156115ed57600080fd5b604052919050565b60006001600160401b0382111561160b57600080fd5b5060209081020190565b60006001600160401b0382111561162b57600080fd5b5060200290565b60006001600160401b0382111561164857600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b60006105c58261167f565b151590565b90565b6001600160a01b031690565b60ff1690565b82818337506000910152565b6116a68161166c565b81146116b157600080fd5b50565b6116a681611677565b6116a68161167c565b6116a68161168b56fea365627a7a72315820a65592924d8d4cbd4846966f2645a3caf5dcf5ed8898ca04636a0d0840ba49596c6578706572696d656e74616cf564736f6c634300050b0040"

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

// VerifySig is a paid mutator transaction binding the contract method 0x3aacfdad.
//
// Solidity: function verifySig(address[] committee, bytes32 msgHash, uint8[] v, bytes32[] r, bytes32[] s) returns(bool)
func (_IncognitoProxy *IncognitoProxyTransactor) VerifySig(opts *bind.TransactOpts, committee []common.Address, msgHash [32]byte, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "verifySig", committee, msgHash, v, r, s)
}

// VerifySig is a paid mutator transaction binding the contract method 0x3aacfdad.
//
// Solidity: function verifySig(address[] committee, bytes32 msgHash, uint8[] v, bytes32[] r, bytes32[] s) returns(bool)
func (_IncognitoProxy *IncognitoProxySession) VerifySig(committee []common.Address, msgHash [32]byte, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.VerifySig(&_IncognitoProxy.TransactOpts, committee, msgHash, v, r, s)
}

// VerifySig is a paid mutator transaction binding the contract method 0x3aacfdad.
//
// Solidity: function verifySig(address[] committee, bytes32 msgHash, uint8[] v, bytes32[] r, bytes32[] s) returns(bool)
func (_IncognitoProxy *IncognitoProxyTransactorSession) VerifySig(committee []common.Address, msgHash [32]byte, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.VerifySig(&_IncognitoProxy.TransactOpts, committee, msgHash, v, r, s)
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
