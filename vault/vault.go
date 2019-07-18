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
const VaultABI = "[{\"name\":\"Deposit\",\"inputs\":[{\"type\":\"address\",\"name\":\"_token\",\"indexed\":false},{\"type\":\"string\",\"name\":\"_incognito_address\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"_amount\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Withdraw\",\"inputs\":[{\"type\":\"address\",\"name\":\"_token\",\"indexed\":false},{\"type\":\"address\",\"name\":\"_to\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"_amount\",\"indexed\":false,\"unit\":\"wei\"}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NotifyString\",\"inputs\":[{\"type\":\"string\",\"name\":\"content\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NotifyBytes32\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"content\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NotifyBool\",\"inputs\":[{\"type\":\"bool\",\"name\":\"content\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NotifyUint256\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"content\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NotifyAddress\",\"inputs\":[{\"type\":\"address\",\"name\":\"content\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"incognitoProxyAddress\"}],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"deposit\",\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"incognito_address\"}],\"constant\":false,\"payable\":true,\"type\":\"function\",\"gas\":25690},{\"name\":\"depositERC20\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"token\"},{\"type\":\"uint256\",\"name\":\"amount\"},{\"type\":\"string\",\"name\":\"incognito_address\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":27797},{\"name\":\"parseBurnInst\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"},{\"type\":\"address\",\"name\":\"out\"},{\"type\":\"address\",\"name\":\"out\"},{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes\",\"name\":\"inst\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":2681},{\"name\":\"withdraw\",\"outputs\":[],\"inputs\":[{\"type\":\"bytes\",\"name\":\"inst\"},{\"type\":\"uint256\",\"name\":\"beaconHeight\"},{\"type\":\"bytes32[8]\",\"name\":\"beaconInstPath\"},{\"type\":\"bool[8]\",\"name\":\"beaconInstPathIsLeft\"},{\"type\":\"int128\",\"name\":\"beaconInstPathLen\"},{\"type\":\"bytes32\",\"name\":\"beaconInstRoot\"},{\"type\":\"bytes32\",\"name\":\"beaconBlkData\"},{\"type\":\"bytes32\",\"name\":\"beaconBlkHash\"},{\"type\":\"uint256\",\"name\":\"beaconSignerSig\"},{\"type\":\"int128\",\"name\":\"beaconNumR\"},{\"type\":\"uint256[8]\",\"name\":\"beaconXs\"},{\"type\":\"uint256[8]\",\"name\":\"beaconYs\"},{\"type\":\"int128[8]\",\"name\":\"beaconRIdxs\"},{\"type\":\"int128\",\"name\":\"beaconNumSig\"},{\"type\":\"uint256[8]\",\"name\":\"beaconSigIdxs\"},{\"type\":\"uint256\",\"name\":\"beaconRx\"},{\"type\":\"uint256\",\"name\":\"beaconRy\"},{\"type\":\"bytes\",\"name\":\"beaconR\"},{\"type\":\"uint256\",\"name\":\"bridgeHeight\"},{\"type\":\"bytes32[8]\",\"name\":\"bridgeInstPath\"},{\"type\":\"bool[8]\",\"name\":\"bridgeInstPathIsLeft\"},{\"type\":\"int128\",\"name\":\"bridgeInstPathLen\"},{\"type\":\"bytes32\",\"name\":\"bridgeInstRoot\"},{\"type\":\"bytes32\",\"name\":\"bridgeBlkData\"},{\"type\":\"bytes32\",\"name\":\"bridgeBlkHash\"},{\"type\":\"uint256\",\"name\":\"bridgeSignerSig\"},{\"type\":\"int128\",\"name\":\"bridgeNumR\"},{\"type\":\"uint256[8]\",\"name\":\"bridgeXs\"},{\"type\":\"uint256[8]\",\"name\":\"bridgeYs\"},{\"type\":\"int128[8]\",\"name\":\"bridgeRIdxs\"},{\"type\":\"int128\",\"name\":\"bridgeNumSig\"},{\"type\":\"uint256[8]\",\"name\":\"bridgeSigIdxs\"},{\"type\":\"uint256\",\"name\":\"bridgeRx\"},{\"type\":\"uint256\",\"name\":\"bridgeRy\"},{\"type\":\"bytes\",\"name\":\"bridgeR\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":102377},{\"name\":\"withdrawed\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"bytes32\",\"name\":\"arg0\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":736},{\"name\":\"incognito\",\"outputs\":[{\"type\":\"address\",\"unit\":\"Incognito_proxy\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":633}]"

// VaultBin is the compiled bytecode used for deploying new contracts.
const VaultBin = `0x740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a05260206117506101403934156100a157600080fd5b602061175060c03960c05160205181106100ba57600080fd5b506101405160015561173856600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a05263a26e118660005114156101d057602060046101403760a06004356004016101603760806004356004013511156100c957600080fd5b6b033b2e3c9fd0803ce80000003411156100e257600080fd5b600061028052346102c052606061024052610240516102a0526101608051602001806102405161028001828460006004600a8704601201f161012357600080fd5b505061024051610280015160206001820306601f8201039050610240516102800161022081516080818352015b83610220511015156101615761017e565b6000610220516020850101535b8151600101808352811415610150575b50505050602061024051610280015160206001820306601f8201039050610240510101610240527f2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e61024051610280a1005b635a67cb87600051141561039957606060046101403734156101f157600080fd5b600435602051811061020257600080fd5b5060a06044356004016101a037608060443560040135111561022357600080fd5b670de0b6b3a764000061016051111561023b57600080fd5b610140513b61024957600080fd5b61014051301861025857600080fd5b602061034060646323b872dd61028052336102a052306102c052610160516102e05261029c6000610140515af161028e57600080fd5b6000506103405161026052610260516102a657600080fd5b610140516103c0526101605161040052606061038052610380516103e0526101a0805160200180610380516103c001828460006004600a8704601201f16102ec57600080fd5b5050610380516103c0015160206001820306601f8201039050610380516103c00161036081516080818352015b836103605110151561032a57610347565b6000610360516020850101535b8151600101808352811415610319575b505050506020610380516103c0015160206001820306601f8201039050610380510101610380527f2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e610380516103c0a1005b637e16e6e160005114156105eb57602060046101403734156103ba57600080fd5b61014c6004356004016101603761012c6004356004013511156103dc57600080fd5b60006003602082066104a0016101605182840111156103fa57600080fd5b61012c806104c08260206020880688030161016001600060046030f150508181528090509050905080602001516000825180602090131561043a57600080fd5b809190121561044857600080fd5b806020036101000a82049050905090506102e052610160805160036020820381131561047357600080fd5b602081066020820481156104ad57816020036101000a6001820160200260200186015104826101000a8260200260200187015102016104b8565b806020026020018501515b9050905090509050905060205181106104d057600080fd5b6106405261016080516023602082038113156104eb57600080fd5b6020810660208204811561052557816020036101000a6001820160200260200186015104826101000a826020026020018701510201610530565b806020026020018501515b90509050905090509050602051811061054857600080fd5b61066052610160805160436020820381131561056357600080fd5b6020810660208204811561059d57816020036101000a6001820160200260200186015104826101000a8260200260200187015102016105a8565b806020026020018501515b905090509050905090506106805260806106a0526106c06102e0518152610640518160200152610660518160400152610680518160600152506106a0516106c0f3005b63d52fd484600051141561160857610ee0600461014037341561060d57600080fd5b61014c6004356004016110203761012c60043560040135111561062f57600080fd5b610144356002811061064057600080fd5b50610164356002811061065257600080fd5b50610184356002811061066457600080fd5b506101a4356002811061067657600080fd5b506101c4356002811061068857600080fd5b506101e4356002811061069a57600080fd5b5061020435600281106106ac57600080fd5b5061022435600281106106be57600080fd5b5060605161024435806040519013156106d657600080fd5b80919012156106e457600080fd5b506060516102e435806040519013156106fc57600080fd5b809190121561070a57600080fd5b50606051610504358060405190131561072257600080fd5b809190121561073057600080fd5b50606051610524358060405190131561074857600080fd5b809190121561075657600080fd5b50606051610544358060405190131561076e57600080fd5b809190121561077c57600080fd5b50606051610564358060405190131561079457600080fd5b80919012156107a257600080fd5b5060605161058435806040519013156107ba57600080fd5b80919012156107c857600080fd5b506060516105a435806040519013156107e057600080fd5b80919012156107ee57600080fd5b506060516105c4358060405190131561080657600080fd5b809190121561081457600080fd5b506060516105e4358060405190131561082c57600080fd5b809190121561083a57600080fd5b50606051610604358060405190131561085257600080fd5b809190121561086057600080fd5b506041610764356004016111a03760216107643560040135111561088357600080fd5b6108a4356002811061089457600080fd5b506108c435600281106108a657600080fd5b506108e435600281106108b857600080fd5b5061090435600281106108ca57600080fd5b5061092435600281106108dc57600080fd5b5061094435600281106108ee57600080fd5b50610964356002811061090057600080fd5b50610984356002811061091257600080fd5b506060516109a4358060405190131561092a57600080fd5b809190121561093857600080fd5b50606051610a44358060405190131561095057600080fd5b809190121561095e57600080fd5b50606051610c64358060405190131561097657600080fd5b809190121561098457600080fd5b50606051610c84358060405190131561099c57600080fd5b80919012156109aa57600080fd5b50606051610ca435806040519013156109c257600080fd5b80919012156109d057600080fd5b50606051610cc435806040519013156109e857600080fd5b80919012156109f657600080fd5b50606051610ce43580604051901315610a0e57600080fd5b8091901215610a1c57600080fd5b50606051610d043580604051901315610a3457600080fd5b8091901215610a4257600080fd5b50606051610d243580604051901315610a5a57600080fd5b8091901215610a6857600080fd5b50606051610d443580604051901315610a8057600080fd5b8091901215610a8e57600080fd5b50606051610d643580604051901315610aa657600080fd5b8091901215610ab457600080fd5b506041610ec435600401611220376021610ec435600401351115610ad757600080fd5b60006112a05260006113005260806115006101846020637e16e6e1611320528061134052611020808051602001808461134001828460006004600a8704601201f1610b2157600080fd5b50508051820160206001820306601f820103905060200191505061133c90506000305af1610b4e57600080fd5b61150080516112a05260208101516112c05260408101516112e05260608101516113005250623732316112a05114610b8557600080fd5b6112c0511515610ba5576113005130311015610ba057600080fd5b610c08565b611300516112c0513b610bb757600080fd5b6112c0513018610bc657600080fd5b602061162060246370a082316115a052306115c0526115bc6112c0515afa610bed57600080fd5b6000506116205110156115805261158051610c0757600080fd5b5b61102080516020820120905061164052600061102061012c8060208461168001018260208501600060046030f1505080518201915050610160516020826116800101526020810190508061168052611680905080516020820120905061166052600061102061012c8060208461184001018260208501600060046030f15050805182019150506108c051602082611840010152602081019050806118405261184090508051602082012090506118205260006116405160e05260c052604060c0205415610cd457600080fd5b6001543b610ce157600080fd5b6001543018610cef57600080fd5b60206122406108046107a063175bc8d76119e0526001611a005261166051611a205261016051611a4052611a6061018080600060200201518260006020020152806001602002015182600160200201528060026020020151826002602002015280600360200201518260036020020152806004602002015182600460200201528060056020020151826005602002015280600660200201518260066020020152806007602002015182600760200201525050611b606102808060006020020151826000602002015280600160200201518260016020020152806002602002015182600260200201528060036020020151826003602002015280600460200201518260046020020152806005602002015182600560200201528060066020020151826006602002015280600760200201518260076020020152505061038051611c60526103a051611c80526103c051611ca0526103e051611cc05261040051611ce05261042051611d0052611d2061044080600060200201518260006020020152806001602002015182600160200201528060026020020151826002602002015280600360200201518260036020020152806004602002015182600460200201528060056020020151826005602002015280600660200201518260066020020152806007602002015182600760200201525050611e2061054080600060200201518260006020020152806001602002015182600160200201528060026020020151826002602002015280600360200201518260036020020152806004602002015182600460200201528060056020020151826005602002015280600660200201518260066020020152806007602002015182600760200201525050611f20610640806000602002015182600060200201528060016020020151826001602002015280600260200201518260026020020152806003602002015182600360200201528060046020020151826004602002015280600560200201518260056020020152806006602002015182600660200201528060076020020151826007602002015250506107405161202052612040610760806000602002015182600060200201528060016020020151826001602002015280600260200201518260026020020152806003602002015182600360200201528060046020020151826004602002015280600560200201518260056020020152806006602002015182600660200201528060076020020151826007602002015250506108605161214052610880516121605280612180526111a08080516020018084611a0001828460006004600a8704601201f16110bc57600080fd5b50508051820160206001820306601f82010390506020019150506119fc90506001545afa6110e957600080fd5b600050612240516110f957600080fd5b6001543b61110657600080fd5b600154301861111457600080fd5b6020612ac06108046107a063175bc8d761226052600061228052611820516122a0526108c0516122c0526122e06108e0806000602002015182600060200201528060016020020151826001602002015280600260200201518260026020020152806003602002015182600360200201528060046020020151826004602002015280600560200201518260056020020152806006602002015182600660200201528060076020020151826007602002015250506123e06109e080600060200201518260006020020152806001602002015182600160200201528060026020020151826002602002015280600360200201518260036020020152806004602002015182600460200201528060056020020151826005602002015280600660200201518260066020020152806007602002015182600760200201525050610ae0516124e052610b005161250052610b205161252052610b405161254052610b605161256052610b8051612580526125a0610ba0806000602002015182600060200201528060016020020151826001602002015280600260200201518260026020020152806003602002015182600360200201528060046020020151826004602002015280600560200201518260056020020152806006602002015182600660200201528060076020020151826007602002015250506126a0610ca0806000602002015182600060200201528060016020020151826001602002015280600260200201518260026020020152806003602002015182600360200201528060046020020151826004602002015280600560200201518260056020020152806006602002015182600660200201528060076020020151826007602002015250506127a0610da080600060200201518260006020020152806001602002015182600160200201528060026020020151826002602002015280600360200201518260036020020152806004602002015182600460200201528060056020020151826005602002015280600660200201518260066020020152806007602002015182600760200201525050610ea0516128a0526128c0610ec080600060200201518260006020020152806001602002015182600160200201528060026020020151826002602002015280600360200201518260036020020152806004602002015182600460200201528060056020020151826005602002015280600660200201518260066020020152806007602002015182600760200201525050610fc0516129c052610fe0516129e05280612a0052611220808051602001808461228001828460006004600a8704601201f16114e157600080fd5b50508051820160206001820306601f820103905060200191505061227c90506001545afa61150e57600080fd5b600050612ac05161151e57600080fd5b600160006116405160e05260c052604060c020556112c051151561155d576000600060006000611300516112e0516000f161155857600080fd5b6115c7565b6112c0513b61156b57600080fd5b6112c051301861157a57600080fd5b6020612ba0604463a9059cbb612b00526112e051612b205261130051612b4052612b1c60006112c0515af16115ae57600080fd5b600050612ba051612ae052612ae0516115c657600080fd5b5b6112c051612bc0526112e051612be05261130051612c00527f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb6060612bc0a1005b63dca40d9e6000511415611645576020600461014037341561162957600080fd5b60006101405160e05260c052604060c0205460005260206000f3005b638a984538600051141561166b57341561165e57600080fd5b60015460005260206000f3005b60006000fd5b6100c7611738036100c76000396100c7611738036000f3`

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

// Withdraw is a paid mutator transaction binding the contract method 0xd52fd484.
//
// Solidity: function withdraw(bytes inst, uint256 beaconHeight, bytes32[8] beaconInstPath, bool[8] beaconInstPathIsLeft, int128 beaconInstPathLen, bytes32 beaconInstRoot, bytes32 beaconBlkData, bytes32 beaconBlkHash, uint256 beaconSignerSig, int128 beaconNumR, uint256[8] beaconXs, uint256[8] beaconYs, int128[8] beaconRIdxs, int128 beaconNumSig, uint256[8] beaconSigIdxs, uint256 beaconRx, uint256 beaconRy, bytes beaconR, uint256 bridgeHeight, bytes32[8] bridgeInstPath, bool[8] bridgeInstPathIsLeft, int128 bridgeInstPathLen, bytes32 bridgeInstRoot, bytes32 bridgeBlkData, bytes32 bridgeBlkHash, uint256 bridgeSignerSig, int128 bridgeNumR, uint256[8] bridgeXs, uint256[8] bridgeYs, int128[8] bridgeRIdxs, int128 bridgeNumSig, uint256[8] bridgeSigIdxs, uint256 bridgeRx, uint256 bridgeRy, bytes bridgeR) returns()
func (_Vault *VaultTransactor) Withdraw(opts *bind.TransactOpts, inst []byte, beaconHeight *big.Int, beaconInstPath [8][32]byte, beaconInstPathIsLeft [8]bool, beaconInstPathLen *big.Int, beaconInstRoot [32]byte, beaconBlkData [32]byte, beaconBlkHash [32]byte, beaconSignerSig *big.Int, beaconNumR *big.Int, beaconXs [8]*big.Int, beaconYs [8]*big.Int, beaconRIdxs [8]*big.Int, beaconNumSig *big.Int, beaconSigIdxs [8]*big.Int, beaconRx *big.Int, beaconRy *big.Int, beaconR []byte, bridgeHeight *big.Int, bridgeInstPath [8][32]byte, bridgeInstPathIsLeft [8]bool, bridgeInstPathLen *big.Int, bridgeInstRoot [32]byte, bridgeBlkData [32]byte, bridgeBlkHash [32]byte, bridgeSignerSig *big.Int, bridgeNumR *big.Int, bridgeXs [8]*big.Int, bridgeYs [8]*big.Int, bridgeRIdxs [8]*big.Int, bridgeNumSig *big.Int, bridgeSigIdxs [8]*big.Int, bridgeRx *big.Int, bridgeRy *big.Int, bridgeR []byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "withdraw", inst, beaconHeight, beaconInstPath, beaconInstPathIsLeft, beaconInstPathLen, beaconInstRoot, beaconBlkData, beaconBlkHash, beaconSignerSig, beaconNumR, beaconXs, beaconYs, beaconRIdxs, beaconNumSig, beaconSigIdxs, beaconRx, beaconRy, beaconR, bridgeHeight, bridgeInstPath, bridgeInstPathIsLeft, bridgeInstPathLen, bridgeInstRoot, bridgeBlkData, bridgeBlkHash, bridgeSignerSig, bridgeNumR, bridgeXs, bridgeYs, bridgeRIdxs, bridgeNumSig, bridgeSigIdxs, bridgeRx, bridgeRy, bridgeR)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd52fd484.
//
// Solidity: function withdraw(bytes inst, uint256 beaconHeight, bytes32[8] beaconInstPath, bool[8] beaconInstPathIsLeft, int128 beaconInstPathLen, bytes32 beaconInstRoot, bytes32 beaconBlkData, bytes32 beaconBlkHash, uint256 beaconSignerSig, int128 beaconNumR, uint256[8] beaconXs, uint256[8] beaconYs, int128[8] beaconRIdxs, int128 beaconNumSig, uint256[8] beaconSigIdxs, uint256 beaconRx, uint256 beaconRy, bytes beaconR, uint256 bridgeHeight, bytes32[8] bridgeInstPath, bool[8] bridgeInstPathIsLeft, int128 bridgeInstPathLen, bytes32 bridgeInstRoot, bytes32 bridgeBlkData, bytes32 bridgeBlkHash, uint256 bridgeSignerSig, int128 bridgeNumR, uint256[8] bridgeXs, uint256[8] bridgeYs, int128[8] bridgeRIdxs, int128 bridgeNumSig, uint256[8] bridgeSigIdxs, uint256 bridgeRx, uint256 bridgeRy, bytes bridgeR) returns()
func (_Vault *VaultSession) Withdraw(inst []byte, beaconHeight *big.Int, beaconInstPath [8][32]byte, beaconInstPathIsLeft [8]bool, beaconInstPathLen *big.Int, beaconInstRoot [32]byte, beaconBlkData [32]byte, beaconBlkHash [32]byte, beaconSignerSig *big.Int, beaconNumR *big.Int, beaconXs [8]*big.Int, beaconYs [8]*big.Int, beaconRIdxs [8]*big.Int, beaconNumSig *big.Int, beaconSigIdxs [8]*big.Int, beaconRx *big.Int, beaconRy *big.Int, beaconR []byte, bridgeHeight *big.Int, bridgeInstPath [8][32]byte, bridgeInstPathIsLeft [8]bool, bridgeInstPathLen *big.Int, bridgeInstRoot [32]byte, bridgeBlkData [32]byte, bridgeBlkHash [32]byte, bridgeSignerSig *big.Int, bridgeNumR *big.Int, bridgeXs [8]*big.Int, bridgeYs [8]*big.Int, bridgeRIdxs [8]*big.Int, bridgeNumSig *big.Int, bridgeSigIdxs [8]*big.Int, bridgeRx *big.Int, bridgeRy *big.Int, bridgeR []byte) (*types.Transaction, error) {
	return _Vault.Contract.Withdraw(&_Vault.TransactOpts, inst, beaconHeight, beaconInstPath, beaconInstPathIsLeft, beaconInstPathLen, beaconInstRoot, beaconBlkData, beaconBlkHash, beaconSignerSig, beaconNumR, beaconXs, beaconYs, beaconRIdxs, beaconNumSig, beaconSigIdxs, beaconRx, beaconRy, beaconR, bridgeHeight, bridgeInstPath, bridgeInstPathIsLeft, bridgeInstPathLen, bridgeInstRoot, bridgeBlkData, bridgeBlkHash, bridgeSignerSig, bridgeNumR, bridgeXs, bridgeYs, bridgeRIdxs, bridgeNumSig, bridgeSigIdxs, bridgeRx, bridgeRy, bridgeR)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd52fd484.
//
// Solidity: function withdraw(bytes inst, uint256 beaconHeight, bytes32[8] beaconInstPath, bool[8] beaconInstPathIsLeft, int128 beaconInstPathLen, bytes32 beaconInstRoot, bytes32 beaconBlkData, bytes32 beaconBlkHash, uint256 beaconSignerSig, int128 beaconNumR, uint256[8] beaconXs, uint256[8] beaconYs, int128[8] beaconRIdxs, int128 beaconNumSig, uint256[8] beaconSigIdxs, uint256 beaconRx, uint256 beaconRy, bytes beaconR, uint256 bridgeHeight, bytes32[8] bridgeInstPath, bool[8] bridgeInstPathIsLeft, int128 bridgeInstPathLen, bytes32 bridgeInstRoot, bytes32 bridgeBlkData, bytes32 bridgeBlkHash, uint256 bridgeSignerSig, int128 bridgeNumR, uint256[8] bridgeXs, uint256[8] bridgeYs, int128[8] bridgeRIdxs, int128 bridgeNumSig, uint256[8] bridgeSigIdxs, uint256 bridgeRx, uint256 bridgeRy, bytes bridgeR) returns()
func (_Vault *VaultTransactorSession) Withdraw(inst []byte, beaconHeight *big.Int, beaconInstPath [8][32]byte, beaconInstPathIsLeft [8]bool, beaconInstPathLen *big.Int, beaconInstRoot [32]byte, beaconBlkData [32]byte, beaconBlkHash [32]byte, beaconSignerSig *big.Int, beaconNumR *big.Int, beaconXs [8]*big.Int, beaconYs [8]*big.Int, beaconRIdxs [8]*big.Int, beaconNumSig *big.Int, beaconSigIdxs [8]*big.Int, beaconRx *big.Int, beaconRy *big.Int, beaconR []byte, bridgeHeight *big.Int, bridgeInstPath [8][32]byte, bridgeInstPathIsLeft [8]bool, bridgeInstPathLen *big.Int, bridgeInstRoot [32]byte, bridgeBlkData [32]byte, bridgeBlkHash [32]byte, bridgeSignerSig *big.Int, bridgeNumR *big.Int, bridgeXs [8]*big.Int, bridgeYs [8]*big.Int, bridgeRIdxs [8]*big.Int, bridgeNumSig *big.Int, bridgeSigIdxs [8]*big.Int, bridgeRx *big.Int, bridgeRy *big.Int, bridgeR []byte) (*types.Transaction, error) {
	return _Vault.Contract.Withdraw(&_Vault.TransactOpts, inst, beaconHeight, beaconInstPath, beaconInstPathIsLeft, beaconInstPathLen, beaconInstRoot, beaconBlkData, beaconBlkHash, beaconSignerSig, beaconNumR, beaconXs, beaconYs, beaconRIdxs, beaconNumSig, beaconSigIdxs, beaconRx, beaconRy, beaconR, bridgeHeight, bridgeInstPath, bridgeInstPathIsLeft, bridgeInstPathLen, bridgeInstRoot, bridgeBlkData, bridgeBlkHash, bridgeSignerSig, bridgeNumR, bridgeXs, bridgeYs, bridgeRIdxs, bridgeNumSig, bridgeSigIdxs, bridgeRx, bridgeRy, bridgeR)
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
