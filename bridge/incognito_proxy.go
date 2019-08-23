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
const IncognitoProxyABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"numVals\",\"type\":\"uint256\"}],\"name\":\"extractCommitteeFromInst\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"latestBridgeBlk\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"instPaths\",\"type\":\"bytes32[][2]\"},{\"internalType\":\"bool[][2]\",\"name\":\"instPathIsLefts\",\"type\":\"bool[][2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"instRoots\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"blkData\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256[][2]\",\"name\":\"sigIdxs\",\"type\":\"uint256[][2]\"},{\"internalType\":\"uint8[][2]\",\"name\":\"sigVs\",\"type\":\"uint8[][2]\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"sigRs\",\"type\":\"bytes32[][2]\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"sigSs\",\"type\":\"bytes32[][2]\"}],\"name\":\"swapBridgeCommittee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"committee\",\"type\":\"address[]\"},{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"verifySig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"left\",\"type\":\"bool[]\"}],\"name\":\"instructionInMerkleTree\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"latestBeaconBlk\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"}],\"name\":\"extractMetaFromInst\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isBeacon\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"instHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blkHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"instPath\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"instPathIsLeft\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"instRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdx\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigS\",\"type\":\"bytes32[]\"}],\"name\":\"instructionApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"beaconCommittee\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"bridgeCommittee\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// IncognitoProxyFuncSigs maps the 4-byte function signature to its string representation.
var IncognitoProxyFuncSigs = map[string]string{
	"116fcab9": "extractCommitteeFromInst(bytes,uint256)",
	"8a66a7d4": "extractMetaFromInst(bytes)",
	"f65d2116": "instructionApproved(bool,bytes32,uint256,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])",
	"47c4b328": "instructionInMerkleTree(bytes32,bytes32,bytes32[],bool[])",
	"82dc4fb6": "latestBeaconBlk()",
	"1f140f24": "latestBridgeBlk()",
	"262f7220": "swapBridgeCommittee(bytes,bytes32[][2],bool[][2],bytes32[2],bytes32[2],uint256[][2],uint8[][2],bytes32[][2],bytes32[][2])",
	"3aacfdad": "verifySig(address[],bytes32,uint8[],bytes32[],bytes32[])",
}

// IncognitoProxyBin is the compiled bytecode used for deploying new contracts.
var IncognitoProxyBin = "0x60806040523480156200001157600080fd5b50604051620012e9380380620012e98339810160408190526200003491620000db565b5160015550620001bd565b80516200004c81620001a3565b92915050565b600082601f8301126200006457600080fd5b81516200007b620000758262000170565b62000149565b91508181835260208401935060208101905083856020840282011115620000a157600080fd5b60005b83811015620000d15781620000ba88826200003f565b8452506020928301929190910190600101620000a4565b5050505092915050565b60008060408385031215620000ef57600080fd5b82516001600160401b038111156200010657600080fd5b620001148582860162000052565b92505060208301516001600160401b038111156200013157600080fd5b6200013f8582860162000052565b9150509250929050565b6040518181016001600160401b03811182821017156200016857600080fd5b604052919050565b60006001600160401b038211156200018757600080fd5b5060209081020190565b60006001600160a01b0382166200004c565b620001ae8162000191565b8114620001ba57600080fd5b50565b61111c80620001cd6000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806347c4b3281161005b57806347c4b3281461010057806382dc4fb6146101135780638a66a7d41461011b578063f65d21161461013d57610088565b8063116fcab91461008d5780631f140f24146100b6578063262f7220146100cb5780633aacfdad146100e0575b600080fd5b6100a061009b366004610e61565b610150565b6040516100ad9190610f48565b60405180910390f35b6100be610188565b6040516100ad9190610fa3565b6100de6100d9366004610cfb565b61018e565b005b6100f36100ee3660046109f0565b61023b565b6040516100ad9190610f60565b6100f361010e366004610c38565b61031c565b6100be610327565b61012e610129366004610cc7565b61032d565b6040516100ad93929190610fb1565b6100f361014b366004610ac0565b610337565b6060808260405190808252806020026020018201604052801561017d578160200160208202803883390190505b509150505b92915050565b60015481565b60006101c56001826000548c6000600281106101a657fe5b60200201518c518c518c518c518c518c518c60005b6020020151610337565b6101ce57600080fd5b6102216000826001548c6001600281106101e457fe5b60200201518c600160200201518c600160200201518c600160200201518c600160200201518c600160200201518c600160200201518c60016101bb565b61022a57600080fd5b505060006001555050505050505050565b6000805b865181101561030d5786818151811061025457fe5b60200260200101516001600160a01b031660018787848151811061027457fe5b602002602001015187858151811061028857fe5b602002602001015187868151811061029c57fe5b6020026020010151604051600081526020016040526040516102c19493929190610f6e565b6020604051602081039080840390855afa1580156102e3573d6000803e3d6000fd5b505050602060405103516001600160a01b031614610305576000915050610313565b60010161023f565b50600190505b95945050505050565b60015b949350505050565b60005481565b5060009081908190565b60008060015490506060865160405190808252806020026020018201604052801561036c578160200160208202803883390190505b5090507f1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac861039d828289898961023b565b6103a657600080fd5b6103b28e8b8e8e61031c565b6103bb57600080fd5b5050509b9a5050505050505050505050565b8035610182816110a7565b600082601f8301126103e957600080fd5b81356103fc6103f782610fff565b610fd9565b9150818183526020840193506020810190508385602084028201111561042157600080fd5b60005b8381101561044d578161043788826103cd565b8452506020928301929190910190600101610424565b5050505092915050565b600082601f83011261046857600080fd5b60026104766103f78261101f565b9150818360005b8381101561044d5781358601610493888261059f565b845250602092830192919091019060010161047d565b600082601f8301126104ba57600080fd5b60026104c86103f78261101f565b9150818360005b8381101561044d57813586016104e588826106e0565b84525060209283019291909101906001016104cf565b600082601f83011261050c57600080fd5b600261051a6103f78261101f565b9150818360005b8381101561044d578135860161053788826107c0565b8452506020928301929190910190600101610521565b600082601f83011261055e57600080fd5b600261056c6103f78261101f565b9150818360005b8381101561044d578135860161058988826108a0565b8452506020928301929190910190600101610573565b600082601f8301126105b057600080fd5b81356105be6103f782610fff565b915081818352602084019350602081019050838560208402820111156105e357600080fd5b60005b8381101561044d57816105f98882610980565b84525060209283019291909101906001016105e6565b600082601f83011261062057600080fd5b813561062e6103f782610fff565b9150818183526020840193506020810190508385602084028201111561065357600080fd5b60005b8381101561044d57816106698882610980565b8452506020928301929190910190600101610656565b600082601f83011261069057600080fd5b600261069e6103f78261101f565b915081838560208402820111156106b457600080fd5b60005b8381101561044d57816106ca888261098b565b84525060209283019291909101906001016106b7565b600082601f8301126106f157600080fd5b81356106ff6103f782610fff565b9150818183526020840193506020810190508385602084028201111561072457600080fd5b60005b8381101561044d578161073a888261098b565b8452506020928301929190910190600101610727565b600082601f83011261076157600080fd5b813561076f6103f782610fff565b9150818183526020840193506020810190508385602084028201111561079457600080fd5b60005b8381101561044d57816107aa888261098b565b8452506020928301929190910190600101610797565b600082601f8301126107d157600080fd5b81356107df6103f782610fff565b9150818183526020840193506020810190508385602084028201111561080457600080fd5b60005b8381101561044d578161081a888261098b565b8452506020928301929190910190600101610807565b600082601f83011261084157600080fd5b813561084f6103f782610fff565b9150818183526020840193506020810190508385602084028201111561087457600080fd5b60005b8381101561044d578161088a888261098b565b8452506020928301929190910190600101610877565b600082601f8301126108b157600080fd5b81356108bf6103f782610fff565b915081818352602084019350602081019050838560208402820111156108e457600080fd5b60005b8381101561044d57816108fa88826109e5565b84525060209283019291909101906001016108e7565b600082601f83011261092157600080fd5b813561092f6103f782610fff565b9150818183526020840193506020810190508385602084028201111561095457600080fd5b60005b8381101561044d578161096a88826109e5565b8452506020928301929190910190600101610957565b8035610182816110be565b8035610182816110c7565b600082601f8301126109a757600080fd5b81356109b56103f78261103c565b915080825260208301602083018583830111156109d157600080fd5b6109dc83828461109b565b50505092915050565b8035610182816110d0565b600080600080600060a08688031215610a0857600080fd5b85356001600160401b03811115610a1e57600080fd5b610a2a888289016103d8565b9550506020610a3b8882890161098b565b94505060408601356001600160401b03811115610a5757600080fd5b610a6388828901610910565b93505060608601356001600160401b03811115610a7f57600080fd5b610a8b88828901610750565b92505060808601356001600160401b03811115610aa757600080fd5b610ab388828901610750565b9150509295509295909350565b60008060008060008060008060008060006101608c8e031215610ae257600080fd5b6000610aee8e8e610980565b9b50506020610aff8e828f0161098b565b9a50506040610b108e828f0161098b565b99505060608c01356001600160401b03811115610b2c57600080fd5b610b388e828f01610750565b98505060808c01356001600160401b03811115610b5457600080fd5b610b608e828f0161060f565b97505060a0610b718e828f0161098b565b96505060c0610b828e828f0161098b565b95505060e08c01356001600160401b03811115610b9e57600080fd5b610baa8e828f01610830565b9450506101008c01356001600160401b03811115610bc757600080fd5b610bd38e828f01610910565b9350506101208c01356001600160401b03811115610bf057600080fd5b610bfc8e828f01610750565b9250506101408c01356001600160401b03811115610c1957600080fd5b610c258e828f01610750565b9150509295989b509295989b9093969950565b60008060008060808587031215610c4e57600080fd5b6000610c5a878761098b565b9450506020610c6b8782880161098b565b93505060408501356001600160401b03811115610c8757600080fd5b610c9387828801610750565b92505060608501356001600160401b03811115610caf57600080fd5b610cbb8782880161060f565b91505092959194509250565b600060208284031215610cd957600080fd5b81356001600160401b03811115610cef57600080fd5b61031f84828501610996565b60008060008060008060008060006101608a8c031215610d1a57600080fd5b89356001600160401b03811115610d3057600080fd5b610d3c8c828d01610996565b99505060208a01356001600160401b03811115610d5857600080fd5b610d648c828d016104a9565b98505060408a01356001600160401b03811115610d8057600080fd5b610d8c8c828d01610457565b9750506060610d9d8c828d0161067f565b96505060a0610dae8c828d0161067f565b95505060e08a01356001600160401b03811115610dca57600080fd5b610dd68c828d016104fb565b9450506101008a01356001600160401b03811115610df357600080fd5b610dff8c828d0161054d565b9350506101208a01356001600160401b03811115610e1c57600080fd5b610e288c828d016104a9565b9250506101408a01356001600160401b03811115610e4557600080fd5b610e518c828d016104a9565b9150509295985092959850929598565b60008060408385031215610e7457600080fd5b82356001600160401b03811115610e8a57600080fd5b610e9685828601610996565b9250506020610ea78582860161098b565b9150509250929050565b6000610ebd8383610ec5565b505060200190565b610ece81611076565b82525050565b6000610edf82611069565b610ee9818561106d565b9350610ef483611063565b8060005b83811015610f22578151610f0c8882610eb1565b9750610f1783611063565b925050600101610ef8565b509495945050505050565b610ece81611081565b610ece81611086565b610ece81611095565b60208082528101610f598184610ed4565b9392505050565b602081016101828284610f2d565b60808101610f7c8287610f36565b610f896020830186610f3f565b610f966040830185610f36565b6103136060830184610f36565b602081016101828284610f36565b60608101610fbf8286610f36565b610fcc6020830185610f36565b61031f6040830184610f36565b6040518181016001600160401b0381118282101715610ff757600080fd5b604052919050565b60006001600160401b0382111561101557600080fd5b5060209081020190565b60006001600160401b0382111561103557600080fd5b5060200290565b60006001600160401b0382111561105257600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b600061018282611089565b151590565b90565b6001600160a01b031690565b60ff1690565b82818337506000910152565b6110b081611076565b81146110bb57600080fd5b50565b6110b081611081565b6110b081611086565b6110b08161109556fea365627a7a7231582028419a2318be17b5b9396157e4239610ead975cadd0ee40e39763bc774dd8d256c6578706572696d656e74616cf564736f6c634300050b0040"

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

// LatestBeaconBlk is a free data retrieval call binding the contract method 0x82dc4fb6.
//
// Solidity: function latestBeaconBlk() constant returns(uint256)
func (_IncognitoProxy *IncognitoProxyCaller) LatestBeaconBlk(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "latestBeaconBlk")
	return *ret0, err
}

// LatestBeaconBlk is a free data retrieval call binding the contract method 0x82dc4fb6.
//
// Solidity: function latestBeaconBlk() constant returns(uint256)
func (_IncognitoProxy *IncognitoProxySession) LatestBeaconBlk() (*big.Int, error) {
	return _IncognitoProxy.Contract.LatestBeaconBlk(&_IncognitoProxy.CallOpts)
}

// LatestBeaconBlk is a free data retrieval call binding the contract method 0x82dc4fb6.
//
// Solidity: function latestBeaconBlk() constant returns(uint256)
func (_IncognitoProxy *IncognitoProxyCallerSession) LatestBeaconBlk() (*big.Int, error) {
	return _IncognitoProxy.Contract.LatestBeaconBlk(&_IncognitoProxy.CallOpts)
}

// LatestBridgeBlk is a free data retrieval call binding the contract method 0x1f140f24.
//
// Solidity: function latestBridgeBlk() constant returns(uint256)
func (_IncognitoProxy *IncognitoProxyCaller) LatestBridgeBlk(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "latestBridgeBlk")
	return *ret0, err
}

// LatestBridgeBlk is a free data retrieval call binding the contract method 0x1f140f24.
//
// Solidity: function latestBridgeBlk() constant returns(uint256)
func (_IncognitoProxy *IncognitoProxySession) LatestBridgeBlk() (*big.Int, error) {
	return _IncognitoProxy.Contract.LatestBridgeBlk(&_IncognitoProxy.CallOpts)
}

// LatestBridgeBlk is a free data retrieval call binding the contract method 0x1f140f24.
//
// Solidity: function latestBridgeBlk() constant returns(uint256)
func (_IncognitoProxy *IncognitoProxyCallerSession) LatestBridgeBlk() (*big.Int, error) {
	return _IncognitoProxy.Contract.LatestBridgeBlk(&_IncognitoProxy.CallOpts)
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
