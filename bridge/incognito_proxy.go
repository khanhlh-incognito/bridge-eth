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
const IncognitoProxyABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"numVals\",\"type\":\"uint256\"}],\"name\":\"extractCommitteeFromInst\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"instPaths\",\"type\":\"bytes32[][2]\"},{\"internalType\":\"bool[][2]\",\"name\":\"instPathIsLefts\",\"type\":\"bool[][2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"instRoots\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"blkData\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256[][2]\",\"name\":\"sigIdxs\",\"type\":\"uint256[][2]\"},{\"internalType\":\"uint8[][2]\",\"name\":\"sigVs\",\"type\":\"uint8[][2]\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"sigRs\",\"type\":\"bytes32[][2]\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"sigSs\",\"type\":\"bytes32[][2]\"}],\"name\":\"swapBridgeCommittee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"committee\",\"type\":\"address[]\"},{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"verifySig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"left\",\"type\":\"bool[]\"}],\"name\":\"instructionInMerkleTree\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"}],\"name\":\"extractMetaFromInst\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bridgeCommittees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"beaconCommittees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isBeacon\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"instHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blkHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"instPath\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"instPathIsLeft\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"instRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdx\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigS\",\"type\":\"bytes32[]\"}],\"name\":\"instructionApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"beaconCommittee\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"bridgeCommittee\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"LogUint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"val\",\"type\":\"string\"}],\"name\":\"LogString\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"val\",\"type\":\"bytes32\"}],\"name\":\"LogBytes32\",\"type\":\"event\"}]"

// IncognitoProxyFuncSigs maps the 4-byte function signature to its string representation.
var IncognitoProxyFuncSigs = map[string]string{
	"f203a5ed": "beaconCommittees(uint256)",
	"9b30b637": "bridgeCommittees(uint256)",
	"116fcab9": "extractCommitteeFromInst(bytes,uint256)",
	"8a66a7d4": "extractMetaFromInst(bytes)",
	"f65d2116": "instructionApproved(bool,bytes32,uint256,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])",
	"47c4b328": "instructionInMerkleTree(bytes32,bytes32,bytes32[],bool[])",
	"262f7220": "swapBridgeCommittee(bytes,bytes32[][2],bool[][2],bytes32[2],bytes32[2],uint256[][2],uint8[][2],bytes32[][2],bytes32[][2])",
	"3aacfdad": "verifySig(address[],bytes32,uint8[],bytes32[],bytes32[])",
}

// IncognitoProxyBin is the compiled bytecode used for deploying new contracts.
var IncognitoProxyBin = "0x60806040523480156200001157600080fd5b50604051620017523803806200175283398101604081905262000034916200024b565b60408051808201909152828152600060208083018290528154600181018084559280528351805193949360029092027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563019262000097928492909101906200011b565b5060209182015160019182015560408051808201909152848152600081840181905282548084018085559390915281518051939550919360029091027fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf601926200010592849201906200011b565b506020820151816001015550505050506200032d565b82805482825590600052602060002090810192821562000173579160200282015b828111156200017357825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906200013c565b506200018192915062000185565b5090565b620001ac91905b80821115620001815780546001600160a01b03191681556001016200018c565b90565b8051620001bc8162000313565b92915050565b600082601f830112620001d457600080fd5b8151620001eb620001e582620002e0565b620002b9565b915081818352602084019350602081019050838560208402820111156200021157600080fd5b60005b838110156200024157816200022a8882620001af565b845250602092830192919091019060010162000214565b5050505092915050565b600080604083850312156200025f57600080fd5b82516001600160401b038111156200027657600080fd5b6200028485828601620001c2565b92505060208301516001600160401b03811115620002a157600080fd5b620002af85828601620001c2565b9150509250929050565b6040518181016001600160401b0381118282101715620002d857600080fd5b604052919050565b60006001600160401b03821115620002f757600080fd5b5060209081020190565b60006001600160a01b038216620001bc565b6200031e8162000301565b81146200032a57600080fd5b50565b611415806200033d6000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80638a66a7d41161005b5780638a66a7d4146100fe5780639b30b63714610120578063f203a5ed14610140578063f65d21161461015357610088565b8063116fcab91461008d578063262f7220146100b65780633aacfdad146100cb57806347c4b328146100eb575b600080fd5b6100a061009b3660046110d5565b610166565b6040516100ad9190611231565b60405180910390f35b6100c96100c4366004610f6f565b61019e565b005b6100de6100d9366004610c64565b6102d5565b6040516100ad9190611249565b6100de6100f9366004610eac565b6103b6565b61011161010c366004610f3b565b6104a6565b6040516100ad939291906112aa565b61013361012e366004611125565b6104b0565b6040516100ad9190611257565b61013361014e366004611125565b6104d6565b6100de610161366004610d34565b6104e3565b60608082604051908082528060200260200182016040528015610193578160200160208202803883390190505b509150505b92915050565b885160208a01206040517e9fd52f05c0ded31d6fb0ee580b923f85e99cf1a5a1da342f25e73c45829c83906101d4908390611257565b60405180910390a186516040517e9fd52f05c0ded31d6fb0ee580b923f85e99cf1a5a1da342f25e73c45829c839161020b91611257565b60405180910390a160208701516040517e9fd52f05c0ded31d6fb0ee580b923f85e99cf1a5a1da342f25e73c45829c839161024591611257565b60405180910390a1600080548190600019810190811061026157fe5b60009182526020822060016002909202018101548154909350600019810190811061028857fe5b5060009081526040519091507fa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f906102bf9061129a565b60405180910390a1505050505050505050505050565b6000805b84518110156103a7578681815181106102ee57fe5b60200260200101516001600160a01b031660018787848151811061030e57fe5b602002602001015187858151811061032257fe5b602002602001015187868151811061033657fe5b60200260200101516040516000815260200160405260405161035b9493929190611265565b6020604051602081039080840390855afa15801561037d573d6000803e3d6000fd5b505050602060405103516001600160a01b03161461039f5760009150506103ad565b6001016102d9565b50600190505b95945050505050565b600084815b8451811015610498578381815181106103d057fe5b602002602001015115610420578481815181106103e957fe5b60200260200101518260405160200161040392919061120b565b604051602081830303815290604052805190602001209150610490565b84818151811061042c57fe5b60200260200101516000801b141561045157818260405160200161040392919061120b565b8185828151811061045e57fe5b602002602001015160405160200161047792919061120b565b6040516020818303038152906040528051906020012091505b6001016103bb565b50841490505b949350505050565b5060009081908190565b600181815481106104bd57fe5b6000918252602090912060016002909202010154905081565b600081815481106104bd57fe5b600060608c15610569576000805460001981019081106104ff57fe5b6000918252602091829020600290910201805460408051828502810185019091528181529283018282801561055d57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161053f575b505050505090506105e1565b60018054600019810190811061057b57fe5b600091825260209182902060029091020180546040805182850281018501909152818152928301828280156105d957602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116105bb575b505050505090505b7f1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac861060f82828888886102d5565b61061857600080fd5b6106248d8a8d8d6103b6565b61062d57600080fd5b5060019d9c50505050505050505050505050565b8035610198816113a0565b600082601f83011261065d57600080fd5b813561067061066b826112f8565b6112d2565b9150818183526020840193506020810190508385602084028201111561069557600080fd5b60005b838110156106c157816106ab8882610641565b8452506020928301929190910190600101610698565b5050505092915050565b600082601f8301126106dc57600080fd5b60026106ea61066b82611318565b9150818360005b838110156106c157813586016107078882610813565b84525060209283019291909101906001016106f1565b600082601f83011261072e57600080fd5b600261073c61066b82611318565b9150818360005b838110156106c157813586016107598882610954565b8452506020928301929190910190600101610743565b600082601f83011261078057600080fd5b600261078e61066b82611318565b9150818360005b838110156106c157813586016107ab8882610a34565b8452506020928301929190910190600101610795565b600082601f8301126107d257600080fd5b60026107e061066b82611318565b9150818360005b838110156106c157813586016107fd8882610b14565b84525060209283019291909101906001016107e7565b600082601f83011261082457600080fd5b813561083261066b826112f8565b9150818183526020840193506020810190508385602084028201111561085757600080fd5b60005b838110156106c1578161086d8882610bf4565b845250602092830192919091019060010161085a565b600082601f83011261089457600080fd5b81356108a261066b826112f8565b915081818352602084019350602081019050838560208402820111156108c757600080fd5b60005b838110156106c157816108dd8882610bf4565b84525060209283019291909101906001016108ca565b600082601f83011261090457600080fd5b600261091261066b82611318565b9150818385602084028201111561092857600080fd5b60005b838110156106c1578161093e8882610bff565b845250602092830192919091019060010161092b565b600082601f83011261096557600080fd5b813561097361066b826112f8565b9150818183526020840193506020810190508385602084028201111561099857600080fd5b60005b838110156106c157816109ae8882610bff565b845250602092830192919091019060010161099b565b600082601f8301126109d557600080fd5b81356109e361066b826112f8565b91508181835260208401935060208101905083856020840282011115610a0857600080fd5b60005b838110156106c15781610a1e8882610bff565b8452506020928301929190910190600101610a0b565b600082601f830112610a4557600080fd5b8135610a5361066b826112f8565b91508181835260208401935060208101905083856020840282011115610a7857600080fd5b60005b838110156106c15781610a8e8882610bff565b8452506020928301929190910190600101610a7b565b600082601f830112610ab557600080fd5b8135610ac361066b826112f8565b91508181835260208401935060208101905083856020840282011115610ae857600080fd5b60005b838110156106c15781610afe8882610bff565b8452506020928301929190910190600101610aeb565b600082601f830112610b2557600080fd5b8135610b3361066b826112f8565b91508181835260208401935060208101905083856020840282011115610b5857600080fd5b60005b838110156106c15781610b6e8882610c59565b8452506020928301929190910190600101610b5b565b600082601f830112610b9557600080fd5b8135610ba361066b826112f8565b91508181835260208401935060208101905083856020840282011115610bc857600080fd5b60005b838110156106c15781610bde8882610c59565b8452506020928301929190910190600101610bcb565b8035610198816113b7565b8035610198816113c0565b600082601f830112610c1b57600080fd5b8135610c2961066b82611335565b91508082526020830160208301858383011115610c4557600080fd5b610c50838284611394565b50505092915050565b8035610198816113c9565b600080600080600060a08688031215610c7c57600080fd5b85356001600160401b03811115610c9257600080fd5b610c9e8882890161064c565b9550506020610caf88828901610bff565b94505060408601356001600160401b03811115610ccb57600080fd5b610cd788828901610b84565b93505060608601356001600160401b03811115610cf357600080fd5b610cff888289016109c4565b92505060808601356001600160401b03811115610d1b57600080fd5b610d27888289016109c4565b9150509295509295909350565b60008060008060008060008060008060006101608c8e031215610d5657600080fd5b6000610d628e8e610bf4565b9b50506020610d738e828f01610bff565b9a50506040610d848e828f01610bff565b99505060608c01356001600160401b03811115610da057600080fd5b610dac8e828f016109c4565b98505060808c01356001600160401b03811115610dc857600080fd5b610dd48e828f01610883565b97505060a0610de58e828f01610bff565b96505060c0610df68e828f01610bff565b95505060e08c01356001600160401b03811115610e1257600080fd5b610e1e8e828f01610aa4565b9450506101008c01356001600160401b03811115610e3b57600080fd5b610e478e828f01610b84565b9350506101208c01356001600160401b03811115610e6457600080fd5b610e708e828f016109c4565b9250506101408c01356001600160401b03811115610e8d57600080fd5b610e998e828f016109c4565b9150509295989b509295989b9093969950565b60008060008060808587031215610ec257600080fd5b6000610ece8787610bff565b9450506020610edf87828801610bff565b93505060408501356001600160401b03811115610efb57600080fd5b610f07878288016109c4565b92505060608501356001600160401b03811115610f2357600080fd5b610f2f87828801610883565b91505092959194509250565b600060208284031215610f4d57600080fd5b81356001600160401b03811115610f6357600080fd5b61049e84828501610c0a565b60008060008060008060008060006101608a8c031215610f8e57600080fd5b89356001600160401b03811115610fa457600080fd5b610fb08c828d01610c0a565b99505060208a01356001600160401b03811115610fcc57600080fd5b610fd88c828d0161071d565b98505060408a01356001600160401b03811115610ff457600080fd5b6110008c828d016106cb565b97505060606110118c828d016108f3565b96505060a06110228c828d016108f3565b95505060e08a01356001600160401b0381111561103e57600080fd5b61104a8c828d0161076f565b9450506101008a01356001600160401b0381111561106757600080fd5b6110738c828d016107c1565b9350506101208a01356001600160401b0381111561109057600080fd5b61109c8c828d0161071d565b9250506101408a01356001600160401b038111156110b957600080fd5b6110c58c828d0161071d565b9150509295985092959850929598565b600080604083850312156110e857600080fd5b82356001600160401b038111156110fe57600080fd5b61110a85828601610c0a565b925050602061111b85828601610bff565b9150509250929050565b60006020828403121561113757600080fd5b600061049e8484610bff565b600061114f8383611157565b505060200190565b6111608161136f565b82525050565b600061117182611362565b61117b8185611366565b93506111868361135c565b8060005b838110156111b457815161119e8882611143565b97506111a98361135c565b92505060010161118a565b509495945050505050565b6111608161137a565b6111608161137f565b6111606111dd8261137f565b61137f565b60006111ef600483611366565b63446f6e6560e01b815260200192915050565b6111608161138e565b600061121782856111d1565b60208201915061122782846111d1565b5060200192915050565b602080825281016112428184611166565b9392505050565b6020810161019882846111bf565b6020810161019882846111c8565b6080810161127382876111c8565b6112806020830186611202565b61128d60408301856111c8565b6103ad60608301846111c8565b60208082528101610198816111e2565b606081016112b882866111c8565b6112c560208301856111c8565b61049e60408301846111c8565b6040518181016001600160401b03811182821017156112f057600080fd5b604052919050565b60006001600160401b0382111561130e57600080fd5b5060209081020190565b60006001600160401b0382111561132e57600080fd5b5060200290565b60006001600160401b0382111561134b57600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b600061019882611382565b151590565b90565b6001600160a01b031690565b60ff1690565b82818337506000910152565b6113a98161136f565b81146113b457600080fd5b50565b6113a98161137a565b6113a98161137f565b6113a98161138e56fea365627a7a72315820676b7a2f3543b821c47740ebbdee39f41e5bd215b7ce6daf8d6d9f2c8b1338436c6578706572696d656e74616cf564736f6c634300050b0040"

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

// ExtractCommitteeFromInst is a free data retrieval call binding the contract method 0x116fcab9.
//
// Solidity: function extractCommitteeFromInst(bytes inst, uint256 numVals) constant returns(address[])
func (_IncognitoProxy *IncognitoProxyCaller) ExtractCommitteeFromInst(opts *bind.CallOpts, inst []byte, numVals *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "extractCommitteeFromInst", inst, numVals)
	return *ret0, err
}

// ExtractCommitteeFromInst is a free data retrieval call binding the contract method 0x116fcab9.
//
// Solidity: function extractCommitteeFromInst(bytes inst, uint256 numVals) constant returns(address[])
func (_IncognitoProxy *IncognitoProxySession) ExtractCommitteeFromInst(inst []byte, numVals *big.Int) ([]common.Address, error) {
	return _IncognitoProxy.Contract.ExtractCommitteeFromInst(&_IncognitoProxy.CallOpts, inst, numVals)
}

// ExtractCommitteeFromInst is a free data retrieval call binding the contract method 0x116fcab9.
//
// Solidity: function extractCommitteeFromInst(bytes inst, uint256 numVals) constant returns(address[])
func (_IncognitoProxy *IncognitoProxyCallerSession) ExtractCommitteeFromInst(inst []byte, numVals *big.Int) ([]common.Address, error) {
	return _IncognitoProxy.Contract.ExtractCommitteeFromInst(&_IncognitoProxy.CallOpts, inst, numVals)
}

// ExtractMetaFromInst is a free data retrieval call binding the contract method 0x8a66a7d4.
//
// Solidity: function extractMetaFromInst(bytes inst) constant returns(uint256, uint256, uint256)
func (_IncognitoProxy *IncognitoProxyCaller) ExtractMetaFromInst(opts *bind.CallOpts, inst []byte) (*big.Int, *big.Int, *big.Int, error) {
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
	err := _IncognitoProxy.contract.Call(opts, out, "extractMetaFromInst", inst)
	return *ret0, *ret1, *ret2, err
}

// ExtractMetaFromInst is a free data retrieval call binding the contract method 0x8a66a7d4.
//
// Solidity: function extractMetaFromInst(bytes inst) constant returns(uint256, uint256, uint256)
func (_IncognitoProxy *IncognitoProxySession) ExtractMetaFromInst(inst []byte) (*big.Int, *big.Int, *big.Int, error) {
	return _IncognitoProxy.Contract.ExtractMetaFromInst(&_IncognitoProxy.CallOpts, inst)
}

// ExtractMetaFromInst is a free data retrieval call binding the contract method 0x8a66a7d4.
//
// Solidity: function extractMetaFromInst(bytes inst) constant returns(uint256, uint256, uint256)
func (_IncognitoProxy *IncognitoProxyCallerSession) ExtractMetaFromInst(inst []byte) (*big.Int, *big.Int, *big.Int, error) {
	return _IncognitoProxy.Contract.ExtractMetaFromInst(&_IncognitoProxy.CallOpts, inst)
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
