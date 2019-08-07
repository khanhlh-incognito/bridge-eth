// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vault

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

// VaultABI is the input ABI used to generate the binding from.
const VaultABI = "[{\"name\":\"Deposit\",\"inputs\":[{\"type\":\"address\",\"name\":\"_token\",\"indexed\":false},{\"type\":\"string\",\"name\":\"_incognito_address\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"_amount\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Withdraw\",\"inputs\":[{\"type\":\"address\",\"name\":\"_token\",\"indexed\":false},{\"type\":\"address\",\"name\":\"_to\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"_amount\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NotifyString\",\"inputs\":[{\"type\":\"string\",\"name\":\"content\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NotifyBytes32\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"content\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NotifyBool\",\"inputs\":[{\"type\":\"bool\",\"name\":\"content\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NotifyUint256\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"content\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NotifyAddress\",\"inputs\":[{\"type\":\"address\",\"name\":\"content\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"incognitoProxyAddress\"}],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"deposit\",\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"incognito_address\"}],\"constant\":false,\"payable\":true,\"type\":\"function\",\"gas\":26610},{\"name\":\"depositERC20\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"token\"},{\"type\":\"uint256\",\"name\":\"amount\"},{\"type\":\"string\",\"name\":\"incognito_address\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":29066},{\"name\":\"parseBurnInst\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"},{\"type\":\"address\",\"name\":\"out\"},{\"type\":\"address\",\"name\":\"out\"},{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes\",\"name\":\"inst\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2693},{\"name\":\"withdraw\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes\",\"name\":\"inst\"},{\"type\":\"uint256\",\"name\":\"beaconHeight\"},{\"type\":\"bytes32[8]\",\"name\":\"beaconInstPath\"},{\"type\":\"bool[8]\",\"name\":\"beaconInstPathIsLeft\"},{\"type\":\"int128\",\"name\":\"beaconInstPathLen\"},{\"type\":\"bytes32\",\"name\":\"beaconInstRoot\"},{\"type\":\"bytes32\",\"name\":\"beaconBlkData\"},{\"type\":\"bytes32\",\"name\":\"beaconBlkHash\"},{\"type\":\"uint256\",\"name\":\"beaconSignerSig\"},{\"type\":\"int128\",\"name\":\"beaconNumR\"},{\"type\":\"uint256[8]\",\"name\":\"beaconXs\"},{\"type\":\"uint256[8]\",\"name\":\"beaconYs\"},{\"type\":\"int128[8]\",\"name\":\"beaconRIdxs\"},{\"type\":\"int128\",\"name\":\"beaconNumSig\"},{\"type\":\"uint256[8]\",\"name\":\"beaconSigIdxs\"},{\"type\":\"bytes\",\"name\":\"beaconRp\"},{\"type\":\"uint256\",\"name\":\"beaconRpx\"},{\"type\":\"uint256\",\"name\":\"beaconRpy\"},{\"type\":\"bytes\",\"name\":\"beaconR\"},{\"type\":\"uint256\",\"name\":\"bridgeHeight\"},{\"type\":\"bytes32[8]\",\"name\":\"bridgeInstPath\"},{\"type\":\"bool[8]\",\"name\":\"bridgeInstPathIsLeft\"},{\"type\":\"int128\",\"name\":\"bridgeInstPathLen\"},{\"type\":\"bytes32\",\"name\":\"bridgeInstRoot\"},{\"type\":\"bytes32\",\"name\":\"bridgeBlkData\"},{\"type\":\"bytes32\",\"name\":\"bridgeBlkHash\"},{\"type\":\"uint256\",\"name\":\"bridgeSignerSig\"},{\"type\":\"int128\",\"name\":\"bridgeNumR\"},{\"type\":\"uint256[8]\",\"name\":\"bridgeXs\"},{\"type\":\"uint256[8]\",\"name\":\"bridgeYs\"},{\"type\":\"int128[8]\",\"name\":\"bridgeRIdxs\"},{\"type\":\"int128\",\"name\":\"bridgeNumSig\"},{\"type\":\"uint256[8]\",\"name\":\"bridgeSigIdxs\"},{\"type\":\"bytes\",\"name\":\"bridgeRp\"},{\"type\":\"uint256\",\"name\":\"bridgeRpx\"},{\"type\":\"uint256\",\"name\":\"bridgeRpy\"},{\"type\":\"bytes\",\"name\":\"bridgeR\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":95970},{\"name\":\"withdrawed\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"arg0\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":748},{\"name\":\"incognito\",\"outputs\":[{\"type\":\"address\",\"unit\":\"Incognito_proxy\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":663}]"

// VaultBin is the compiled bytecode used for deploying new contracts.
const VaultBin = `0x740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a05260206114876101403934156100a157600080fd5b602061148760c03960c05160205181106100ba57600080fd5b506101405160015561146f56600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a05263a26e118660005114156101db5760a06004356004016101403760806004356004013511156100c157600080fd5b6b033b2e3c9fd0803ce8000000343031340110156100de57600080fd5b3031340111156100ed57600080fd5b600061026052346102a05260606102205261022051610280526101408051602001806102205161026001828460006004600a8704601201f161012e57600080fd5b505061022051610260015160206001820306601f8201039050610220516102600161020081516080818352015b836102005110151561016c57610189565b6000610200516020850101535b815160010180835281141561015b575b50505050602061022051610260015160206001820306601f8201039050610220510101610220527f2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e61022051610260a1005b635a67cb8760005114156103fc5734156101f457600080fd5b600435602051811061020557600080fd5b5060a060443560040161014037608060443560040135111561022657600080fd5b6004353b61023357600080fd5b600435301861024157600080fd5b60206102a060246370a0823161022052306102405261023c6004355afa61026757600080fd5b6000506102a05161020052670de0b6b3a76400006024356102005160243501101561029157600080fd5b610200516024350111156102a457600080fd5b6004353b6102b157600080fd5b60043530186102bf57600080fd5b60206103a060646323b872dd6102e05233610300523061032052602435610340526102fc60006004355af16102f357600080fd5b6000506103a0516102c0526102c05161030b57600080fd5b600435610420526024356104605260606103e0526103e051610440526101408051602001806103e05161042001828460006004600a8704601201f161034f57600080fd5b50506103e051610420015160206001820306601f82010390506103e051610420016103c081516080818352015b836103c05110151561038d576103aa565b60006103c0516020850101535b815160010180835281141561037c575b5050505060206103e051610420015160206001820306601f82010390506103e05101016103e0527f2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e6103e051610420a1005b637e16e6e1600051141561064757341561041557600080fd5b61014c6004356004016101403761012c60043560040135111561043757600080fd5b60006003602082066104800161014051828401111561045557600080fd5b61012c806104a08260206020880688030161014001600060046030f150508181528090509050905080602001516000825180602090131561049557600080fd5b80919012156104a357600080fd5b806020036101000a82049050905090506102c05261014080516003602082038113156104ce57600080fd5b6020810660208204811561050857816020036101000a6001820160200260200186015104826101000a826020026020018701510201610513565b806020026020018501515b90509050905090509050602051811061052b57600080fd5b61062052610140805160236020820381131561054657600080fd5b6020810660208204811561058057816020036101000a6001820160200260200186015104826101000a82602002602001870151020161058b565b806020026020018501515b9050905090509050905060205181106105a357600080fd5b6106405261014080516043602082038113156105be57600080fd5b602081066020820481156105f857816020036101000a6001820160200260200186015104826101000a826020026020018701510201610603565b806020026020018501515b90509050905090509050610660526080610680526106a06102c051815261062051816020015261064051816040015261066051816060015250610680516106a0f350005b6306181142600051141561134657341561066057600080fd5b61014c6004356004016101403761012c60043560040135111561068257600080fd5b6000610120525b602061012051016101205260e06101205110156106a557610689565b6000610120525b610120516101440135600281106106c257600080fd5b50602061012051016101205260e06101205110156106df576106ac565b60605161024435806040519013156106f657600080fd5b809190121561070457600080fd5b506060516102e4358060405190131561071c57600080fd5b809190121561072a57600080fd5b506000610120525b602061012051016101205260e061012051101561074e57610732565b6000610120525b602061012051016101205260e061012051101561077157610755565b6000610120525b6060516101205161050401358060405190131561079457600080fd5b80919012156107a257600080fd5b50602061012051016101205260e06101205110156107bf57610778565b60605161060435806040519013156107d657600080fd5b80919012156107e457600080fd5b506000610120525b602061012051016101205260e0610120511015610808576107ec565b6041610724356004016102c03760216107243560040135111561082a57600080fd5b6041610784356004016103403760216107843560040135111561084c57600080fd5b6000610120525b602061012051016101205260e061012051101561086f57610853565b6000610120525b610120516108c401356002811061088c57600080fd5b50602061012051016101205260e06101205110156108a957610876565b6060516109c435806040519013156108c057600080fd5b80919012156108ce57600080fd5b50606051610a6435806040519013156108e657600080fd5b80919012156108f457600080fd5b506000610120525b602061012051016101205260e0610120511015610918576108fc565b6000610120525b602061012051016101205260e061012051101561093b5761091f565b6000610120525b60605161012051610c8401358060405190131561095e57600080fd5b809190121561096c57600080fd5b50602061012051016101205260e061012051101561098957610942565b606051610d8435806040519013156109a057600080fd5b80919012156109ae57600080fd5b506000610120525b602061012051016101205260e06101205110156109d2576109b6565b6041610ea4356004016103c0376021610ea4356004013511156109f457600080fd5b6041610f0435600401610440376021610f0435600401351115610a1657600080fd5b60006104c05260006105205260806107206101846020637e16e6e1610540528061056052610140808051602001808461056001828460006004600a8704601201f1610a6057600080fd5b50508051820160206001820306601f820103905060200191505061055c90506000305af1610a8d57600080fd5b61072080516104c05260208101516104e05260408101516105005260608101516105205250623732316104c05114610ac457600080fd5b6104e0511515610ae4576105205130311015610adf57600080fd5b610b47565b610520516104e0513b610af657600080fd5b6104e0513018610b0557600080fd5b602061084060246370a082316107c052306107e0526107dc6104e0515afa610b2c57600080fd5b6000506108405110156107a0526107a051610b4657600080fd5b5b61014080516020820120905061086052600061014061012c806020846108a001018260208501600060046030f15050805182019150506024356020826108a0010152602081019050806108a0526108a0905080516020820120905061088052600061014061012c80602084610a6001018260208501600060046030f15050805182019150506107a435602082610a6001015260208101905080610a6052610a609050805160208201209050610a405260006108605160e05260c052604060c0205415610c1257600080fd5b6001543b610c1f57600080fd5b6001543018610c2d57600080fd5b60206114e06108846107c06394558fb7610c00526001610c205261088051610c4052602435610c6052610c80604480358252806020013582602001528060400135826040015280606001358260600152806080013582608001528060a001358260a001528060c001358260c001528060e001358260e001525050610d8061014480358252806020013582602001528060400135826040015280606001358260600152806080013582608001528060a001358260a001528060c001358260c001528060e001358260e00152505061024435610e805261026435610ea05261028435610ec0526102a435610ee0526102c435610f00526102e435610f2052610f4061030480358252806020013582602001528060400135826040015280606001358260600152806080013582608001528060a001358260a001528060c001358260c001528060e001358260e00152505061104061040480358252806020013582602001528060400135826040015280606001358260600152806080013582608001528060a001358260a001528060c001358260c001528060e001358260e00152505061114061050480358252806020013582602001528060400135826040015280606001358260600152806080013582608001528060a001358260a001528060c001358260c001528060e001358260e001525050610604356112405261126061062480358252806020013582602001528060400135826040015280606001358260600152806080013582608001528060a001358260a001528060c001358260c001528060e001358260e00152505080611360526102c08080516020018084610c2001828460006004600a8704601201f1610ea457600080fd5b50508051820160206001820306601f82010390506020019150506107443561138052610764356113a052806113c0526103408080516020018084610c2001828460006004600a8704601201f1610ef957600080fd5b50508051820160206001820306601f8201039050602001915050610c1c90506001545afa610f2657600080fd5b6000506114e051610f3657600080fd5b6001543b610f4357600080fd5b6001543018610f5157600080fd5b6020611de06108846107c06394558fb761150052600061152052610a4051611540526107a435611560526115806107c480358252806020013582602001528060400135826040015280606001358260600152806080013582608001528060a001358260a001528060c001358260c001528060e001358260e0015250506116806108c480358252806020013582602001528060400135826040015280606001358260600152806080013582608001528060a001358260a001528060c001358260c001528060e001358260e0015250506109c435611780526109e4356117a052610a04356117c052610a24356117e052610a443561180052610a643561182052611840610a8480358252806020013582602001528060400135826040015280606001358260600152806080013582608001528060a001358260a001528060c001358260c001528060e001358260e001525050611940610b8480358252806020013582602001528060400135826040015280606001358260600152806080013582608001528060a001358260a001528060c001358260c001528060e001358260e001525050611a40610c8480358252806020013582602001528060400135826040015280606001358260600152806080013582608001528060a001358260a001528060c001358260c001528060e001358260e001525050610d8435611b4052611b60610da480358252806020013582602001528060400135826040015280606001358260600152806080013582608001528060a001358260a001528060c001358260c001528060e001358260e00152505080611c60526103c0808051602001808461152001828460006004600a8704601201f16111ca57600080fd5b50508051820160206001820306601f8201039050602001915050610ec435611c8052610ee435611ca05280611cc052610440808051602001808461152001828460006004600a8704601201f161121f57600080fd5b50508051820160206001820306601f820103905060200191505061151c90506001545afa61124c57600080fd5b600050611de05161125c57600080fd5b600160006108605160e05260c052604060c020556104e051151561129b57600060006000600061052051610500516000f161129657600080fd5b611305565b6104e0513b6112a957600080fd5b6104e05130186112b857600080fd5b6020611ec0604463a9059cbb611e205261050051611e405261052051611e6052611e3c60006104e0515af16112ec57600080fd5b600050611ec051611e0052611e005161130457600080fd5b5b6104e051611ee05261050051611f005261052051611f20527f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb6060611ee0a1005b63dca40d9e600051141561137b57341561135f57600080fd5b600060043560e05260c052604060c0205460005260206000f350005b638a98453860005114156113a257341561139457600080fd5b60015460005260206000f350005b60006000fd5b6100c761146f036100c76000396100c761146f036000f3`

// DeployVault deploys a new Ethereum contract, binding an instance of Vault to it.
func DeployVault(auth *bind.TransactOpts, backend bind.ContractBackend, incognitoProxyAddress common.Address) (common.Address, *types.Transaction, *Vault, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VaultBin), backend, incognitoProxyAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}, VaultFilterer: VaultFilterer{contract: contract}}, nil
}

// Vault is an auto generated Go binding around an Ethereum contract.
type Vault struct {
	VaultCaller     // Read-only binding to the contract
	VaultTransactor // Write-only binding to the contract
	VaultFilterer   // Log filterer for contract events
}

// VaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultSession struct {
	Contract     *Vault            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultCallerSession struct {
	Contract *VaultCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultTransactorSession struct {
	Contract     *VaultTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultRaw struct {
	Contract *Vault // Generic contract binding to access the raw methods on
}

// VaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultCallerRaw struct {
	Contract *VaultCaller // Generic read-only contract binding to access the raw methods on
}

// VaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultTransactorRaw struct {
	Contract *VaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVault creates a new instance of Vault, bound to a specific deployed contract.
func NewVault(address common.Address, backend bind.ContractBackend) (*Vault, error) {
	contract, err := bindVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}, VaultFilterer: VaultFilterer{contract: contract}}, nil
}

// NewVaultCaller creates a new read-only instance of Vault, bound to a specific deployed contract.
func NewVaultCaller(address common.Address, caller bind.ContractCaller) (*VaultCaller, error) {
	contract, err := bindVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultCaller{contract: contract}, nil
}

// NewVaultTransactor creates a new write-only instance of Vault, bound to a specific deployed contract.
func NewVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultTransactor, error) {
	contract, err := bindVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultTransactor{contract: contract}, nil
}

// NewVaultFilterer creates a new log filterer instance of Vault, bound to a specific deployed contract.
func NewVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultFilterer, error) {
	contract, err := bindVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultFilterer{contract: contract}, nil
}

// bindVault binds a generic wrapper to an already deployed contract.
func bindVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.VaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transact(opts, method, params...)
}

// Incognito is a free data retrieval call binding the contract method 0x8a984538.
//
// Solidity: function incognito() constant returns(address out)
func (_Vault *VaultCaller) Incognito(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "incognito")
	return *ret0, err
}

// Incognito is a free data retrieval call binding the contract method 0x8a984538.
//
// Solidity: function incognito() constant returns(address out)
func (_Vault *VaultSession) Incognito() (common.Address, error) {
	return _Vault.Contract.Incognito(&_Vault.CallOpts)
}

// Incognito is a free data retrieval call binding the contract method 0x8a984538.
//
// Solidity: function incognito() constant returns(address out)
func (_Vault *VaultCallerSession) Incognito() (common.Address, error) {
	return _Vault.Contract.Incognito(&_Vault.CallOpts)
}

// ParseBurnInst is a free data retrieval call binding the contract method 0x7e16e6e1.
//
// Solidity: function parseBurnInst(bytes inst) constant returns(uint256 out, address out, address out, uint256 out)
func (_Vault *VaultCaller) ParseBurnInst(opts *bind.CallOpts, inst []byte) (*big.Int, common.Address, common.Address, *big.Int, error) {
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
	err := _Vault.contract.Call(opts, out, "parseBurnInst", inst)
	return *ret0, *ret1, *ret2, *ret3, err
}

// ParseBurnInst is a free data retrieval call binding the contract method 0x7e16e6e1.
//
// Solidity: function parseBurnInst(bytes inst) constant returns(uint256 out, address out, address out, uint256 out)
func (_Vault *VaultSession) ParseBurnInst(inst []byte) (*big.Int, common.Address, common.Address, *big.Int, error) {
	return _Vault.Contract.ParseBurnInst(&_Vault.CallOpts, inst)
}

// ParseBurnInst is a free data retrieval call binding the contract method 0x7e16e6e1.
//
// Solidity: function parseBurnInst(bytes inst) constant returns(uint256 out, address out, address out, uint256 out)
func (_Vault *VaultCallerSession) ParseBurnInst(inst []byte) (*big.Int, common.Address, common.Address, *big.Int, error) {
	return _Vault.Contract.ParseBurnInst(&_Vault.CallOpts, inst)
}

// Withdrawed is a free data retrieval call binding the contract method 0xdca40d9e.
//
// Solidity: function withdrawed(bytes32 arg0) constant returns(bool out)
func (_Vault *VaultCaller) Withdrawed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "withdrawed", arg0)
	return *ret0, err
}

// Withdrawed is a free data retrieval call binding the contract method 0xdca40d9e.
//
// Solidity: function withdrawed(bytes32 arg0) constant returns(bool out)
func (_Vault *VaultSession) Withdrawed(arg0 [32]byte) (bool, error) {
	return _Vault.Contract.Withdrawed(&_Vault.CallOpts, arg0)
}

// Withdrawed is a free data retrieval call binding the contract method 0xdca40d9e.
//
// Solidity: function withdrawed(bytes32 arg0) constant returns(bool out)
func (_Vault *VaultCallerSession) Withdrawed(arg0 [32]byte) (bool, error) {
	return _Vault.Contract.Withdrawed(&_Vault.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xa26e1186.
//
// Solidity: function deposit(string incognito_address) returns()
func (_Vault *VaultTransactor) Deposit(opts *bind.TransactOpts, incognito_address string) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "deposit", incognito_address)
}

// Deposit is a paid mutator transaction binding the contract method 0xa26e1186.
//
// Solidity: function deposit(string incognito_address) returns()
func (_Vault *VaultSession) Deposit(incognito_address string) (*types.Transaction, error) {
	return _Vault.Contract.Deposit(&_Vault.TransactOpts, incognito_address)
}

// Deposit is a paid mutator transaction binding the contract method 0xa26e1186.
//
// Solidity: function deposit(string incognito_address) returns()
func (_Vault *VaultTransactorSession) Deposit(incognito_address string) (*types.Transaction, error) {
	return _Vault.Contract.Deposit(&_Vault.TransactOpts, incognito_address)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x5a67cb87.
//
// Solidity: function depositERC20(address token, uint256 amount, string incognito_address) returns()
func (_Vault *VaultTransactor) DepositERC20(opts *bind.TransactOpts, token common.Address, amount *big.Int, incognito_address string) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "depositERC20", token, amount, incognito_address)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x5a67cb87.
//
// Solidity: function depositERC20(address token, uint256 amount, string incognito_address) returns()
func (_Vault *VaultSession) DepositERC20(token common.Address, amount *big.Int, incognito_address string) (*types.Transaction, error) {
	return _Vault.Contract.DepositERC20(&_Vault.TransactOpts, token, amount, incognito_address)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x5a67cb87.
//
// Solidity: function depositERC20(address token, uint256 amount, string incognito_address) returns()
func (_Vault *VaultTransactorSession) DepositERC20(token common.Address, amount *big.Int, incognito_address string) (*types.Transaction, error) {
	return _Vault.Contract.DepositERC20(&_Vault.TransactOpts, token, amount, incognito_address)
}

// Withdraw is a paid mutator transaction binding the contract method 0x06181142.
//
// Solidity: function withdraw(bytes inst, uint256 beaconHeight, bytes32[8] beaconInstPath, bool[8] beaconInstPathIsLeft, int128 beaconInstPathLen, bytes32 beaconInstRoot, bytes32 beaconBlkData, bytes32 beaconBlkHash, uint256 beaconSignerSig, int128 beaconNumR, uint256[8] beaconXs, uint256[8] beaconYs, int128[8] beaconRIdxs, int128 beaconNumSig, uint256[8] beaconSigIdxs, bytes beaconRp, uint256 beaconRpx, uint256 beaconRpy, bytes beaconR, uint256 bridgeHeight, bytes32[8] bridgeInstPath, bool[8] bridgeInstPathIsLeft, int128 bridgeInstPathLen, bytes32 bridgeInstRoot, bytes32 bridgeBlkData, bytes32 bridgeBlkHash, uint256 bridgeSignerSig, int128 bridgeNumR, uint256[8] bridgeXs, uint256[8] bridgeYs, int128[8] bridgeRIdxs, int128 bridgeNumSig, uint256[8] bridgeSigIdxs, bytes bridgeRp, uint256 bridgeRpx, uint256 bridgeRpy, bytes bridgeR) returns()
func (_Vault *VaultTransactor) Withdraw(opts *bind.TransactOpts, inst []byte, beaconHeight *big.Int, beaconInstPath [8][32]byte, beaconInstPathIsLeft [8]bool, beaconInstPathLen *big.Int, beaconInstRoot [32]byte, beaconBlkData [32]byte, beaconBlkHash [32]byte, beaconSignerSig *big.Int, beaconNumR *big.Int, beaconXs [8]*big.Int, beaconYs [8]*big.Int, beaconRIdxs [8]*big.Int, beaconNumSig *big.Int, beaconSigIdxs [8]*big.Int, beaconRp []byte, beaconRpx *big.Int, beaconRpy *big.Int, beaconR []byte, bridgeHeight *big.Int, bridgeInstPath [8][32]byte, bridgeInstPathIsLeft [8]bool, bridgeInstPathLen *big.Int, bridgeInstRoot [32]byte, bridgeBlkData [32]byte, bridgeBlkHash [32]byte, bridgeSignerSig *big.Int, bridgeNumR *big.Int, bridgeXs [8]*big.Int, bridgeYs [8]*big.Int, bridgeRIdxs [8]*big.Int, bridgeNumSig *big.Int, bridgeSigIdxs [8]*big.Int, bridgeRp []byte, bridgeRpx *big.Int, bridgeRpy *big.Int, bridgeR []byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "withdraw", inst, beaconHeight, beaconInstPath, beaconInstPathIsLeft, beaconInstPathLen, beaconInstRoot, beaconBlkData, beaconBlkHash, beaconSignerSig, beaconNumR, beaconXs, beaconYs, beaconRIdxs, beaconNumSig, beaconSigIdxs, beaconRp, beaconRpx, beaconRpy, beaconR, bridgeHeight, bridgeInstPath, bridgeInstPathIsLeft, bridgeInstPathLen, bridgeInstRoot, bridgeBlkData, bridgeBlkHash, bridgeSignerSig, bridgeNumR, bridgeXs, bridgeYs, bridgeRIdxs, bridgeNumSig, bridgeSigIdxs, bridgeRp, bridgeRpx, bridgeRpy, bridgeR)
}

// Withdraw is a paid mutator transaction binding the contract method 0x06181142.
//
// Solidity: function withdraw(bytes inst, uint256 beaconHeight, bytes32[8] beaconInstPath, bool[8] beaconInstPathIsLeft, int128 beaconInstPathLen, bytes32 beaconInstRoot, bytes32 beaconBlkData, bytes32 beaconBlkHash, uint256 beaconSignerSig, int128 beaconNumR, uint256[8] beaconXs, uint256[8] beaconYs, int128[8] beaconRIdxs, int128 beaconNumSig, uint256[8] beaconSigIdxs, bytes beaconRp, uint256 beaconRpx, uint256 beaconRpy, bytes beaconR, uint256 bridgeHeight, bytes32[8] bridgeInstPath, bool[8] bridgeInstPathIsLeft, int128 bridgeInstPathLen, bytes32 bridgeInstRoot, bytes32 bridgeBlkData, bytes32 bridgeBlkHash, uint256 bridgeSignerSig, int128 bridgeNumR, uint256[8] bridgeXs, uint256[8] bridgeYs, int128[8] bridgeRIdxs, int128 bridgeNumSig, uint256[8] bridgeSigIdxs, bytes bridgeRp, uint256 bridgeRpx, uint256 bridgeRpy, bytes bridgeR) returns()
func (_Vault *VaultSession) Withdraw(inst []byte, beaconHeight *big.Int, beaconInstPath [8][32]byte, beaconInstPathIsLeft [8]bool, beaconInstPathLen *big.Int, beaconInstRoot [32]byte, beaconBlkData [32]byte, beaconBlkHash [32]byte, beaconSignerSig *big.Int, beaconNumR *big.Int, beaconXs [8]*big.Int, beaconYs [8]*big.Int, beaconRIdxs [8]*big.Int, beaconNumSig *big.Int, beaconSigIdxs [8]*big.Int, beaconRp []byte, beaconRpx *big.Int, beaconRpy *big.Int, beaconR []byte, bridgeHeight *big.Int, bridgeInstPath [8][32]byte, bridgeInstPathIsLeft [8]bool, bridgeInstPathLen *big.Int, bridgeInstRoot [32]byte, bridgeBlkData [32]byte, bridgeBlkHash [32]byte, bridgeSignerSig *big.Int, bridgeNumR *big.Int, bridgeXs [8]*big.Int, bridgeYs [8]*big.Int, bridgeRIdxs [8]*big.Int, bridgeNumSig *big.Int, bridgeSigIdxs [8]*big.Int, bridgeRp []byte, bridgeRpx *big.Int, bridgeRpy *big.Int, bridgeR []byte) (*types.Transaction, error) {
	return _Vault.Contract.Withdraw(&_Vault.TransactOpts, inst, beaconHeight, beaconInstPath, beaconInstPathIsLeft, beaconInstPathLen, beaconInstRoot, beaconBlkData, beaconBlkHash, beaconSignerSig, beaconNumR, beaconXs, beaconYs, beaconRIdxs, beaconNumSig, beaconSigIdxs, beaconRp, beaconRpx, beaconRpy, beaconR, bridgeHeight, bridgeInstPath, bridgeInstPathIsLeft, bridgeInstPathLen, bridgeInstRoot, bridgeBlkData, bridgeBlkHash, bridgeSignerSig, bridgeNumR, bridgeXs, bridgeYs, bridgeRIdxs, bridgeNumSig, bridgeSigIdxs, bridgeRp, bridgeRpx, bridgeRpy, bridgeR)
}

// Withdraw is a paid mutator transaction binding the contract method 0x06181142.
//
// Solidity: function withdraw(bytes inst, uint256 beaconHeight, bytes32[8] beaconInstPath, bool[8] beaconInstPathIsLeft, int128 beaconInstPathLen, bytes32 beaconInstRoot, bytes32 beaconBlkData, bytes32 beaconBlkHash, uint256 beaconSignerSig, int128 beaconNumR, uint256[8] beaconXs, uint256[8] beaconYs, int128[8] beaconRIdxs, int128 beaconNumSig, uint256[8] beaconSigIdxs, bytes beaconRp, uint256 beaconRpx, uint256 beaconRpy, bytes beaconR, uint256 bridgeHeight, bytes32[8] bridgeInstPath, bool[8] bridgeInstPathIsLeft, int128 bridgeInstPathLen, bytes32 bridgeInstRoot, bytes32 bridgeBlkData, bytes32 bridgeBlkHash, uint256 bridgeSignerSig, int128 bridgeNumR, uint256[8] bridgeXs, uint256[8] bridgeYs, int128[8] bridgeRIdxs, int128 bridgeNumSig, uint256[8] bridgeSigIdxs, bytes bridgeRp, uint256 bridgeRpx, uint256 bridgeRpy, bytes bridgeR) returns()
func (_Vault *VaultTransactorSession) Withdraw(inst []byte, beaconHeight *big.Int, beaconInstPath [8][32]byte, beaconInstPathIsLeft [8]bool, beaconInstPathLen *big.Int, beaconInstRoot [32]byte, beaconBlkData [32]byte, beaconBlkHash [32]byte, beaconSignerSig *big.Int, beaconNumR *big.Int, beaconXs [8]*big.Int, beaconYs [8]*big.Int, beaconRIdxs [8]*big.Int, beaconNumSig *big.Int, beaconSigIdxs [8]*big.Int, beaconRp []byte, beaconRpx *big.Int, beaconRpy *big.Int, beaconR []byte, bridgeHeight *big.Int, bridgeInstPath [8][32]byte, bridgeInstPathIsLeft [8]bool, bridgeInstPathLen *big.Int, bridgeInstRoot [32]byte, bridgeBlkData [32]byte, bridgeBlkHash [32]byte, bridgeSignerSig *big.Int, bridgeNumR *big.Int, bridgeXs [8]*big.Int, bridgeYs [8]*big.Int, bridgeRIdxs [8]*big.Int, bridgeNumSig *big.Int, bridgeSigIdxs [8]*big.Int, bridgeRp []byte, bridgeRpx *big.Int, bridgeRpy *big.Int, bridgeR []byte) (*types.Transaction, error) {
	return _Vault.Contract.Withdraw(&_Vault.TransactOpts, inst, beaconHeight, beaconInstPath, beaconInstPathIsLeft, beaconInstPathLen, beaconInstRoot, beaconBlkData, beaconBlkHash, beaconSignerSig, beaconNumR, beaconXs, beaconYs, beaconRIdxs, beaconNumSig, beaconSigIdxs, beaconRp, beaconRpx, beaconRpy, beaconR, bridgeHeight, bridgeInstPath, bridgeInstPathIsLeft, bridgeInstPathLen, bridgeInstRoot, bridgeBlkData, bridgeBlkHash, bridgeSignerSig, bridgeNumR, bridgeXs, bridgeYs, bridgeRIdxs, bridgeNumSig, bridgeSigIdxs, bridgeRp, bridgeRpx, bridgeRpy, bridgeR)
}

// VaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Vault contract.
type VaultDepositIterator struct {
	Event *VaultDeposit // Event containing the contract specifics and raw log

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
func (it *VaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultDeposit)
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
		it.Event = new(VaultDeposit)
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
func (it *VaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultDeposit represents a Deposit event raised by the Vault contract.
type VaultDeposit struct {
	Token            common.Address
	IncognitoAddress string
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e.
//
// Solidity: event Deposit(address _token, string _incognito_address, uint256 _amount)
func (_Vault *VaultFilterer) FilterDeposit(opts *bind.FilterOpts) (*VaultDepositIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &VaultDepositIterator{contract: _Vault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e.
//
// Solidity: event Deposit(address _token, string _incognito_address, uint256 _amount)
func (_Vault *VaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *VaultDeposit) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultDeposit)
				if err := _Vault.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// VaultNotifyAddressIterator is returned from FilterNotifyAddress and is used to iterate over the raw logs and unpacked data for NotifyAddress events raised by the Vault contract.
type VaultNotifyAddressIterator struct {
	Event *VaultNotifyAddress // Event containing the contract specifics and raw log

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
func (it *VaultNotifyAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultNotifyAddress)
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
		it.Event = new(VaultNotifyAddress)
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
func (it *VaultNotifyAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultNotifyAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultNotifyAddress represents a NotifyAddress event raised by the Vault contract.
type VaultNotifyAddress struct {
	Content common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNotifyAddress is a free log retrieval operation binding the contract event 0x0ac6e167e94338a282ec23bdd86f338fc787bd67f48b3ade098144aac3fcd86e.
//
// Solidity: event NotifyAddress(address content)
func (_Vault *VaultFilterer) FilterNotifyAddress(opts *bind.FilterOpts) (*VaultNotifyAddressIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "NotifyAddress")
	if err != nil {
		return nil, err
	}
	return &VaultNotifyAddressIterator{contract: _Vault.contract, event: "NotifyAddress", logs: logs, sub: sub}, nil
}

// WatchNotifyAddress is a free log subscription operation binding the contract event 0x0ac6e167e94338a282ec23bdd86f338fc787bd67f48b3ade098144aac3fcd86e.
//
// Solidity: event NotifyAddress(address content)
func (_Vault *VaultFilterer) WatchNotifyAddress(opts *bind.WatchOpts, sink chan<- *VaultNotifyAddress) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "NotifyAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultNotifyAddress)
				if err := _Vault.contract.UnpackLog(event, "NotifyAddress", log); err != nil {
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

// VaultNotifyBoolIterator is returned from FilterNotifyBool and is used to iterate over the raw logs and unpacked data for NotifyBool events raised by the Vault contract.
type VaultNotifyBoolIterator struct {
	Event *VaultNotifyBool // Event containing the contract specifics and raw log

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
func (it *VaultNotifyBoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultNotifyBool)
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
		it.Event = new(VaultNotifyBool)
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
func (it *VaultNotifyBoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultNotifyBoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultNotifyBool represents a NotifyBool event raised by the Vault contract.
type VaultNotifyBool struct {
	Content bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNotifyBool is a free log retrieval operation binding the contract event 0x6c8f06ff564112a969115be5f33d4a0f87ba918c9c9bc3090fe631968e818be4.
//
// Solidity: event NotifyBool(bool content)
func (_Vault *VaultFilterer) FilterNotifyBool(opts *bind.FilterOpts) (*VaultNotifyBoolIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "NotifyBool")
	if err != nil {
		return nil, err
	}
	return &VaultNotifyBoolIterator{contract: _Vault.contract, event: "NotifyBool", logs: logs, sub: sub}, nil
}

// WatchNotifyBool is a free log subscription operation binding the contract event 0x6c8f06ff564112a969115be5f33d4a0f87ba918c9c9bc3090fe631968e818be4.
//
// Solidity: event NotifyBool(bool content)
func (_Vault *VaultFilterer) WatchNotifyBool(opts *bind.WatchOpts, sink chan<- *VaultNotifyBool) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "NotifyBool")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultNotifyBool)
				if err := _Vault.contract.UnpackLog(event, "NotifyBool", log); err != nil {
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

// VaultNotifyBytes32Iterator is returned from FilterNotifyBytes32 and is used to iterate over the raw logs and unpacked data for NotifyBytes32 events raised by the Vault contract.
type VaultNotifyBytes32Iterator struct {
	Event *VaultNotifyBytes32 // Event containing the contract specifics and raw log

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
func (it *VaultNotifyBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultNotifyBytes32)
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
		it.Event = new(VaultNotifyBytes32)
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
func (it *VaultNotifyBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultNotifyBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultNotifyBytes32 represents a NotifyBytes32 event raised by the Vault contract.
type VaultNotifyBytes32 struct {
	Content [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNotifyBytes32 is a free log retrieval operation binding the contract event 0xb42152598f9b870207037767fd41b627a327c9434c796b2ee501d68acec68d1b.
//
// Solidity: event NotifyBytes32(bytes32 content)
func (_Vault *VaultFilterer) FilterNotifyBytes32(opts *bind.FilterOpts) (*VaultNotifyBytes32Iterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "NotifyBytes32")
	if err != nil {
		return nil, err
	}
	return &VaultNotifyBytes32Iterator{contract: _Vault.contract, event: "NotifyBytes32", logs: logs, sub: sub}, nil
}

// WatchNotifyBytes32 is a free log subscription operation binding the contract event 0xb42152598f9b870207037767fd41b627a327c9434c796b2ee501d68acec68d1b.
//
// Solidity: event NotifyBytes32(bytes32 content)
func (_Vault *VaultFilterer) WatchNotifyBytes32(opts *bind.WatchOpts, sink chan<- *VaultNotifyBytes32) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "NotifyBytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultNotifyBytes32)
				if err := _Vault.contract.UnpackLog(event, "NotifyBytes32", log); err != nil {
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

// VaultNotifyStringIterator is returned from FilterNotifyString and is used to iterate over the raw logs and unpacked data for NotifyString events raised by the Vault contract.
type VaultNotifyStringIterator struct {
	Event *VaultNotifyString // Event containing the contract specifics and raw log

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
func (it *VaultNotifyStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultNotifyString)
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
		it.Event = new(VaultNotifyString)
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
func (it *VaultNotifyStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultNotifyStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultNotifyString represents a NotifyString event raised by the Vault contract.
type VaultNotifyString struct {
	Content string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNotifyString is a free log retrieval operation binding the contract event 0x8b1126c8e4087477c3efd9e3785935b29c778491c70e249de774345f7ca9b7f9.
//
// Solidity: event NotifyString(string content)
func (_Vault *VaultFilterer) FilterNotifyString(opts *bind.FilterOpts) (*VaultNotifyStringIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "NotifyString")
	if err != nil {
		return nil, err
	}
	return &VaultNotifyStringIterator{contract: _Vault.contract, event: "NotifyString", logs: logs, sub: sub}, nil
}

// WatchNotifyString is a free log subscription operation binding the contract event 0x8b1126c8e4087477c3efd9e3785935b29c778491c70e249de774345f7ca9b7f9.
//
// Solidity: event NotifyString(string content)
func (_Vault *VaultFilterer) WatchNotifyString(opts *bind.WatchOpts, sink chan<- *VaultNotifyString) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "NotifyString")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultNotifyString)
				if err := _Vault.contract.UnpackLog(event, "NotifyString", log); err != nil {
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

// VaultNotifyUint256Iterator is returned from FilterNotifyUint256 and is used to iterate over the raw logs and unpacked data for NotifyUint256 events raised by the Vault contract.
type VaultNotifyUint256Iterator struct {
	Event *VaultNotifyUint256 // Event containing the contract specifics and raw log

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
func (it *VaultNotifyUint256Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultNotifyUint256)
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
		it.Event = new(VaultNotifyUint256)
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
func (it *VaultNotifyUint256Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultNotifyUint256Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultNotifyUint256 represents a NotifyUint256 event raised by the Vault contract.
type VaultNotifyUint256 struct {
	Content *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNotifyUint256 is a free log retrieval operation binding the contract event 0x8e2fc7b10a4f77a18c553db9a8f8c24d9e379da2557cb61ad4cc513a2f992cbd.
//
// Solidity: event NotifyUint256(uint256 content)
func (_Vault *VaultFilterer) FilterNotifyUint256(opts *bind.FilterOpts) (*VaultNotifyUint256Iterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "NotifyUint256")
	if err != nil {
		return nil, err
	}
	return &VaultNotifyUint256Iterator{contract: _Vault.contract, event: "NotifyUint256", logs: logs, sub: sub}, nil
}

// WatchNotifyUint256 is a free log subscription operation binding the contract event 0x8e2fc7b10a4f77a18c553db9a8f8c24d9e379da2557cb61ad4cc513a2f992cbd.
//
// Solidity: event NotifyUint256(uint256 content)
func (_Vault *VaultFilterer) WatchNotifyUint256(opts *bind.WatchOpts, sink chan<- *VaultNotifyUint256) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "NotifyUint256")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultNotifyUint256)
				if err := _Vault.contract.UnpackLog(event, "NotifyUint256", log); err != nil {
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

// VaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Vault contract.
type VaultWithdrawIterator struct {
	Event *VaultWithdraw // Event containing the contract specifics and raw log

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
func (it *VaultWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultWithdraw)
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
		it.Event = new(VaultWithdraw)
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
func (it *VaultWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultWithdraw represents a Withdraw event raised by the Vault contract.
type VaultWithdraw struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address _token, address _to, uint256 _amount)
func (_Vault *VaultFilterer) FilterWithdraw(opts *bind.FilterOpts) (*VaultWithdrawIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &VaultWithdrawIterator{contract: _Vault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address _token, address _to, uint256 _amount)
func (_Vault *VaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *VaultWithdraw) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultWithdraw)
				if err := _Vault.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
