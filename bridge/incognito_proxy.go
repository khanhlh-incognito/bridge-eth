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
const IncognitoProxyABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"numVals\",\"type\":\"uint256\"}],\"name\":\"extractCommitteeFromInst\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"instPaths\",\"type\":\"bytes32[][2]\"},{\"internalType\":\"bool[][2]\",\"name\":\"instPathIsLefts\",\"type\":\"bool[][2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"instRoots\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"blkData\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256[][2]\",\"name\":\"sigIdxs\",\"type\":\"uint256[][2]\"},{\"internalType\":\"uint8[][2]\",\"name\":\"sigVs\",\"type\":\"uint8[][2]\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"sigRs\",\"type\":\"bytes32[][2]\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"sigSs\",\"type\":\"bytes32[][2]\"}],\"name\":\"swapBridgeCommittee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"committee\",\"type\":\"address[]\"},{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"verifySig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"left\",\"type\":\"bool[]\"}],\"name\":\"instructionInMerkleTree\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"}],\"name\":\"extractMetaFromInst\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bridgeCommittees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"beaconCommittees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isBeacon\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"instHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blkHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"instPath\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"instPathIsLeft\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"instRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdx\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigS\",\"type\":\"bytes32[]\"}],\"name\":\"instructionApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"beaconCommittee\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"bridgeCommittee\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"LogUint\",\"type\":\"event\"}]"

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
var IncognitoProxyBin = "0x60806040523480156200001157600080fd5b50604051620015c7380380620015c783398101604081905262000034916200024b565b60408051808201909152828152600060208083018290528154600181018084559280528351805193949360029092027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563019262000097928492909101906200011b565b5060209182015160019182015560408051808201909152848152600081840181905282548084018085559390915281518051939550919360029091027fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf601926200010592849201906200011b565b506020820151816001015550505050506200032d565b82805482825590600052602060002090810192821562000173579160200282015b828111156200017357825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906200013c565b506200018192915062000185565b5090565b620001ac91905b80821115620001815780546001600160a01b03191681556001016200018c565b90565b8051620001bc8162000313565b92915050565b600082601f830112620001d457600080fd5b8151620001eb620001e582620002e0565b620002b9565b915081818352602084019350602081019050838560208402820111156200021157600080fd5b60005b838110156200024157816200022a8882620001af565b845250602092830192919091019060010162000214565b5050505092915050565b600080604083850312156200025f57600080fd5b82516001600160401b038111156200027657600080fd5b6200028485828601620001c2565b92505060208301516001600160401b03811115620002a157600080fd5b620002af85828601620001c2565b9150509250929050565b6040518181016001600160401b0381118282101715620002d857600080fd5b604052919050565b60006001600160401b03821115620002f757600080fd5b5060209081020190565b60006001600160a01b038216620001bc565b6200031e8162000301565b81146200032a57600080fd5b50565b61128a806200033d6000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80638a66a7d41161005b5780638a66a7d4146100fe5780639b30b63714610120578063f203a5ed14610140578063f65d21161461015357610088565b8063116fcab91461008d578063262f7220146100b65780633aacfdad146100cb57806347c4b328146100eb575b600080fd5b6100a061009b366004610fb1565b610166565b6040516100ad91906110b6565b60405180910390f35b6100c96100c4366004610e4b565b61019e565b005b6100de6100d9366004610b40565b610296565b6040516100ad91906110ce565b6100de6100f9366004610d88565b610377565b61011161010c366004610e17565b610382565b6040516100ad9392919061111f565b61013361012e366004611001565b61038c565b6040516100ad9190611111565b61013361014e366004611001565b6103b2565b6100de610161366004610c10565b6103bf565b60608082604051908082528060200260200182016040528015610193578160200160208202803883390190505b509150505b92915050565b600080548190819060001981019081106101b457fe5b90600052602060002090600202016001015490506101fb600183838d6000600281106101dc57fe5b60200201518d518d518d518d518d518d518d60005b60200201516103bf565b61020457600080fd5b6001805460009190600019810190811061021a57fe5b906000526020600020906002020160010154905061027f600084838e60016002811061024257fe5b60200201518e600160200201518e600160200201518e600160200201518e600160200201518e600160200201518e600160200201518e60016101f1565b61028857600080fd5b505050505050505050505050565b6000805b8651811015610368578681815181106102af57fe5b60200260200101516001600160a01b03166001878784815181106102cf57fe5b60200260200101518785815181106102e357fe5b60200260200101518786815181106102f757fe5b60200260200101516040516000815260200160405260405161031c94939291906110dc565b6020604051602081039080840390855afa15801561033e573d6000803e3d6000fd5b505050602060405103516001600160a01b03161461036057600091505061036e565b60010161029a565b50600190505b95945050505050565b60015b949350505050565b5060009081908190565b6001818154811061039957fe5b6000918252602090912060016002909202010154905081565b6000818154811061039957fe5b600060608c15610445576000805460001981019081106103db57fe5b6000918252602091829020600290910201805460408051828502810185019091528181529283018282801561043957602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161041b575b505050505090506104bd565b60018054600019810190811061045757fe5b600091825260209182902060029091020180546040805182850281018501909152818152928301828280156104b557602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610497575b505050505090505b7f1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac86104eb8282888888610296565b6104f457600080fd5b6105008d8a8d8d610377565b61050957600080fd5b5060019d9c50505050505050505050505050565b803561019881611215565b600082601f83011261053957600080fd5b813561054c6105478261116d565b611147565b9150818183526020840193506020810190508385602084028201111561057157600080fd5b60005b8381101561059d5781610587888261051d565b8452506020928301929190910190600101610574565b5050505092915050565b600082601f8301126105b857600080fd5b60026105c66105478261118d565b9150818360005b8381101561059d57813586016105e388826106ef565b84525060209283019291909101906001016105cd565b600082601f83011261060a57600080fd5b60026106186105478261118d565b9150818360005b8381101561059d57813586016106358882610830565b845250602092830192919091019060010161061f565b600082601f83011261065c57600080fd5b600261066a6105478261118d565b9150818360005b8381101561059d57813586016106878882610910565b8452506020928301929190910190600101610671565b600082601f8301126106ae57600080fd5b60026106bc6105478261118d565b9150818360005b8381101561059d57813586016106d988826109f0565b84525060209283019291909101906001016106c3565b600082601f83011261070057600080fd5b813561070e6105478261116d565b9150818183526020840193506020810190508385602084028201111561073357600080fd5b60005b8381101561059d57816107498882610ad0565b8452506020928301929190910190600101610736565b600082601f83011261077057600080fd5b813561077e6105478261116d565b915081818352602084019350602081019050838560208402820111156107a357600080fd5b60005b8381101561059d57816107b98882610ad0565b84525060209283019291909101906001016107a6565b600082601f8301126107e057600080fd5b60026107ee6105478261118d565b9150818385602084028201111561080457600080fd5b60005b8381101561059d578161081a8882610adb565b8452506020928301929190910190600101610807565b600082601f83011261084157600080fd5b813561084f6105478261116d565b9150818183526020840193506020810190508385602084028201111561087457600080fd5b60005b8381101561059d578161088a8882610adb565b8452506020928301929190910190600101610877565b600082601f8301126108b157600080fd5b81356108bf6105478261116d565b915081818352602084019350602081019050838560208402820111156108e457600080fd5b60005b8381101561059d57816108fa8882610adb565b84525060209283019291909101906001016108e7565b600082601f83011261092157600080fd5b813561092f6105478261116d565b9150818183526020840193506020810190508385602084028201111561095457600080fd5b60005b8381101561059d578161096a8882610adb565b8452506020928301929190910190600101610957565b600082601f83011261099157600080fd5b813561099f6105478261116d565b915081818352602084019350602081019050838560208402820111156109c457600080fd5b60005b8381101561059d57816109da8882610adb565b84525060209283019291909101906001016109c7565b600082601f830112610a0157600080fd5b8135610a0f6105478261116d565b91508181835260208401935060208101905083856020840282011115610a3457600080fd5b60005b8381101561059d5781610a4a8882610b35565b8452506020928301929190910190600101610a37565b600082601f830112610a7157600080fd5b8135610a7f6105478261116d565b91508181835260208401935060208101905083856020840282011115610aa457600080fd5b60005b8381101561059d5781610aba8882610b35565b8452506020928301929190910190600101610aa7565b80356101988161122c565b803561019881611235565b600082601f830112610af757600080fd5b8135610b05610547826111aa565b91508082526020830160208301858383011115610b2157600080fd5b610b2c838284611209565b50505092915050565b80356101988161123e565b600080600080600060a08688031215610b5857600080fd5b85356001600160401b03811115610b6e57600080fd5b610b7a88828901610528565b9550506020610b8b88828901610adb565b94505060408601356001600160401b03811115610ba757600080fd5b610bb388828901610a60565b93505060608601356001600160401b03811115610bcf57600080fd5b610bdb888289016108a0565b92505060808601356001600160401b03811115610bf757600080fd5b610c03888289016108a0565b9150509295509295909350565b60008060008060008060008060008060006101608c8e031215610c3257600080fd5b6000610c3e8e8e610ad0565b9b50506020610c4f8e828f01610adb565b9a50506040610c608e828f01610adb565b99505060608c01356001600160401b03811115610c7c57600080fd5b610c888e828f016108a0565b98505060808c01356001600160401b03811115610ca457600080fd5b610cb08e828f0161075f565b97505060a0610cc18e828f01610adb565b96505060c0610cd28e828f01610adb565b95505060e08c01356001600160401b03811115610cee57600080fd5b610cfa8e828f01610980565b9450506101008c01356001600160401b03811115610d1757600080fd5b610d238e828f01610a60565b9350506101208c01356001600160401b03811115610d4057600080fd5b610d4c8e828f016108a0565b9250506101408c01356001600160401b03811115610d6957600080fd5b610d758e828f016108a0565b9150509295989b509295989b9093969950565b60008060008060808587031215610d9e57600080fd5b6000610daa8787610adb565b9450506020610dbb87828801610adb565b93505060408501356001600160401b03811115610dd757600080fd5b610de3878288016108a0565b92505060608501356001600160401b03811115610dff57600080fd5b610e0b8782880161075f565b91505092959194509250565b600060208284031215610e2957600080fd5b81356001600160401b03811115610e3f57600080fd5b61037a84828501610ae6565b60008060008060008060008060006101608a8c031215610e6a57600080fd5b89356001600160401b03811115610e8057600080fd5b610e8c8c828d01610ae6565b99505060208a01356001600160401b03811115610ea857600080fd5b610eb48c828d016105f9565b98505060408a01356001600160401b03811115610ed057600080fd5b610edc8c828d016105a7565b9750506060610eed8c828d016107cf565b96505060a0610efe8c828d016107cf565b95505060e08a01356001600160401b03811115610f1a57600080fd5b610f268c828d0161064b565b9450506101008a01356001600160401b03811115610f4357600080fd5b610f4f8c828d0161069d565b9350506101208a01356001600160401b03811115610f6c57600080fd5b610f788c828d016105f9565b9250506101408a01356001600160401b03811115610f9557600080fd5b610fa18c828d016105f9565b9150509295985092959850929598565b60008060408385031215610fc457600080fd5b82356001600160401b03811115610fda57600080fd5b610fe685828601610ae6565b9250506020610ff785828601610adb565b9150509250929050565b60006020828403121561101357600080fd5b600061037a8484610adb565b600061102b8383611033565b505060200190565b61103c816111e4565b82525050565b600061104d826111d7565b61105781856111db565b9350611062836111d1565b8060005b8381101561109057815161107a888261101f565b9750611085836111d1565b925050600101611066565b509495945050505050565b61103c816111ef565b61103c816111f4565b61103c81611203565b602080825281016110c78184611042565b9392505050565b60208101610198828461109b565b608081016110ea82876110a4565b6110f760208301866110ad565b61110460408301856110a4565b61036e60608301846110a4565b6020810161019882846110a4565b6060810161112d82866110a4565b61113a60208301856110a4565b61037a60408301846110a4565b6040518181016001600160401b038111828210171561116557600080fd5b604052919050565b60006001600160401b0382111561118357600080fd5b5060209081020190565b60006001600160401b038211156111a357600080fd5b5060200290565b60006001600160401b038211156111c057600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b6000610198826111f7565b151590565b90565b6001600160a01b031690565b60ff1690565b82818337506000910152565b61121e816111e4565b811461122957600080fd5b50565b61121e816111ef565b61121e816111f4565b61121e8161120356fea365627a7a723158201ba0dc653c96ce86284e48581f61119db6ac300884dd6dbb6f904caa484933796c6578706572696d656e74616cf564736f6c634300050b0040"

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
