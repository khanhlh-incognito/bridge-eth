// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kbnmultiTrade

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

// KbnmultiTradeABI is the input ABI used to generate the binding from.
const KbnmultiTradeABI = "[{\"inputs\":[{\"internalType\":\"contractKyberNetwork\",\"name\":\"_kyberNetworkProxyContract\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_incognitoSmartContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"ETH_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcQty\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"getConversionRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"incognitoSmartContract\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kyberNetworkProxyContract\",\"outputs\":[{\"internalType\":\"contractKyberNetwork\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"srcTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"srcQties\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"destTokens\",\"type\":\"address[]\"}],\"name\":\"trade\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// KbnmultiTradeBin is the compiled bytecode used for deploying new contracts.
var KbnmultiTradeBin = "0x608060405234801561001057600080fd5b50604051610db5380380610db58339818101604052604081101561003357600080fd5b508051602090910151600180546001600160a01b039384166001600160a01b03199182161790915560008054939092169216919091179055610d3b8061007a6000396000f3fe60806040526004361061004a5760003560e01c80630aea81881461004c5780630f721819146100a857806372e94bf6146102ec578063785250da1461031d578063b42a644b14610332575b005b34801561005857600080fd5b5061008f6004803603606081101561006f57600080fd5b506001600160a01b03813581169160208101359160409091013516610347565b6040805192835260208301919091528051918290030190f35b610253600480360360608110156100be57600080fd5b8101906020810181356401000000008111156100d957600080fd5b8201836020820111156100eb57600080fd5b8035906020019184602083028401116401000000008311171561010d57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929594936020810193503591505064010000000081111561015d57600080fd5b82018360208201111561016f57600080fd5b8035906020019184602083028401116401000000008311171561019157600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092959493602081019350359150506401000000008111156101e157600080fd5b8201836020820111156101f357600080fd5b8035906020019184602083028401116401000000008311171561021557600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506103e2945050505050565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b8381101561029757818101518382015260200161027f565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156102d65781810151838201526020016102be565b5050505090500194505050505060405180910390f35b3480156102f857600080fd5b506103016106b3565b604080516001600160a01b039092168252519081900360200190f35b34801561032957600080fd5b506103016106b8565b34801561033e57600080fd5b506103016106c7565b6001546040805163809a9e5560e01b81526001600160a01b0386811660048301528481166024830152604482018690528251600094859492169263809a9e55926064808301939192829003018186803b1580156103a357600080fd5b505afa1580156103b7573d6000803e3d6000fd5b505050506040513d60408110156103cd57600080fd5b50805160209091015190969095509350505050565b60005460609081906001600160a01b031633146103fe57600080fd5b83518551148015610410575084518351145b61041957600080fd5b60608351604051908082528060200260200182016040528015610446578160200160208202803883390190505b50905060005b86518110156106a65785818151811061046157fe5b602002602001015161048588838151811061047857fe5b60200260200101516106d6565b101561049057600080fd5b84818151811061049c57fe5b60200260200101516001600160a01b03168782815181106104b957fe5b60200260200101516001600160a01b031614156104d557600080fd5b60006001600160a01b03168782815181106104ec57fe5b60200260200101516001600160a01b03161461060b5761054887828151811061051157fe5b6020026020010151600160009054906101000a90046001600160a01b031688848151811061053b57fe5b6020026020010151610768565b60006001600160a01b031685828151811061055f57fe5b60200260200101516001600160a01b0316146105ca5760006105bb88838151811061058657fe5b602002602001015188848151811061059a57fe5b60200260200101518885815181106105ae57fe5b60200260200101516108a7565b116105c557600080fd5b610606565b60006105fc8883815181106105db57fe5b60200260200101518884815181106105ef57fe5b60200260200101516109cb565b1161060657600080fd5b610647565b600061063d86838151811061061c57fe5b602002602001015188848151811061063057fe5b6020026020010151610af8565b1161064757600080fd5b61065685828151811061047857fe5b82828151811061066257fe5b60200260200101818152505061069e85828151811061067d57fe5b602002602001015183838151811061069157fe5b6020026020010151610c23565b60010161044c565b5092959294509192505050565b600081565b6001546001600160a01b031681565b6000546001600160a01b031681565b60006001600160a01b0382166106ee57503031610763565b604080516370a0823160e01b815230600482015290516001600160a01b038416916370a08231916024808301926020929190829003018186803b15801561073457600080fd5b505afa158015610748573d6000803e3d6000fd5b505050506040513d602081101561075e57600080fd5b505190505b919050565b6001600160a01b038316156108a257826001600160a01b031663095ea7b38360006040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050602060405180830381600087803b1580156107d857600080fd5b505af11580156107ec573d6000803e3d6000fd5b505050506040513d602081101561080257600080fd5b505161080d57600080fd5b826001600160a01b031663095ea7b383836040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050602060405180830381600087803b15801561086d57600080fd5b505af1158015610881573d6000803e3d6000fd5b505050506040513d602081101561089757600080fd5b50516108a257600080fd5b505050565b6001546040805163809a9e5560e01b81526001600160a01b0386811660048301528481166024830152604482018690528251600094859492169263809a9e55926064808301939192829003018186803b15801561090357600080fd5b505afa158015610917573d6000803e3d6000fd5b505050506040513d604081101561092d57600080fd5b505160015460408051637409e2eb60e01b81526001600160a01b038981166004830152602482018990528781166044830152606482018590529151939450911691637409e2eb916084808201926020929091908290030181600087803b15801561099657600080fd5b505af11580156109aa573d6000803e3d6000fd5b505050506040513d60208110156109c057600080fd5b505195945050505050565b6001546040805163809a9e5560e01b81526001600160a01b03858116600483015273eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee6024830152604482018590528251600094859492169263809a9e55926064808301939192829003018186803b158015610a3957600080fd5b505afa158015610a4d573d6000803e3d6000fd5b505050506040513d6040811015610a6357600080fd5b505160015460408051630eee887760e21b81526001600160a01b03888116600483015260248201889052604482018590529151939450911691633bba21dc916064808201926020929091908290030181600087803b158015610ac457600080fd5b505af1158015610ad8573d6000803e3d6000fd5b505050506040513d6020811015610aee57600080fd5b5051949350505050565b60003031821115610b0857600080fd5b6001546040805163809a9e5560e01b815273eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee60048201526001600160a01b0386811660248301523460448301528251600094919091169263809a9e559260648082019391829003018186803b158015610b7457600080fd5b505afa158015610b88573d6000803e3d6000fd5b505050506040513d6040811015610b9e57600080fd5b505160015460408051633d15022b60e11b81526001600160a01b038881166004830152602482018590529151939450911691637a2a0456918691604480830192602092919082900301818588803b158015610bf857600080fd5b505af1158015610c0c573d6000803e3d6000fd5b50505050506040513d6020811015610aee57600080fd5b6001600160a01b038216610c7e573031811115610c3f57600080fd5b600080546040516001600160a01b039091169183156108fc02918491818181858888f19350505050158015610c78573d6000803e3d6000fd5b50610d02565b600080546040805163a9059cbb60e01b81526001600160a01b0392831660048201526024810185905290519185169263a9059cbb926044808401936020939083900390910190829087803b158015610cd557600080fd5b505af1158015610ce9573d6000803e3d6000fd5b505050506040513d6020811015610cff57600080fd5b50505b505056fea265627a7a7231582048dd87d10618075320adfcacf6538acc9feaf0b61d65d4a85e097d6911658cc964736f6c634300050d0032"

// DeployKbnmultiTrade deploys a new Ethereum contract, binding an instance of KbnmultiTrade to it.
func DeployKbnmultiTrade(auth *bind.TransactOpts, backend bind.ContractBackend, _kyberNetworkProxyContract common.Address, _incognitoSmartContract common.Address) (common.Address, *types.Transaction, *KbnmultiTrade, error) {
	parsed, err := abi.JSON(strings.NewReader(KbnmultiTradeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KbnmultiTradeBin), backend, _kyberNetworkProxyContract, _incognitoSmartContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KbnmultiTrade{KbnmultiTradeCaller: KbnmultiTradeCaller{contract: contract}, KbnmultiTradeTransactor: KbnmultiTradeTransactor{contract: contract}, KbnmultiTradeFilterer: KbnmultiTradeFilterer{contract: contract}}, nil
}

// KbnmultiTrade is an auto generated Go binding around an Ethereum contract.
type KbnmultiTrade struct {
	KbnmultiTradeCaller     // Read-only binding to the contract
	KbnmultiTradeTransactor // Write-only binding to the contract
	KbnmultiTradeFilterer   // Log filterer for contract events
}

// KbnmultiTradeCaller is an auto generated read-only Go binding around an Ethereum contract.
type KbnmultiTradeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KbnmultiTradeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KbnmultiTradeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KbnmultiTradeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KbnmultiTradeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KbnmultiTradeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KbnmultiTradeSession struct {
	Contract     *KbnmultiTrade    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KbnmultiTradeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KbnmultiTradeCallerSession struct {
	Contract *KbnmultiTradeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// KbnmultiTradeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KbnmultiTradeTransactorSession struct {
	Contract     *KbnmultiTradeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// KbnmultiTradeRaw is an auto generated low-level Go binding around an Ethereum contract.
type KbnmultiTradeRaw struct {
	Contract *KbnmultiTrade // Generic contract binding to access the raw methods on
}

// KbnmultiTradeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KbnmultiTradeCallerRaw struct {
	Contract *KbnmultiTradeCaller // Generic read-only contract binding to access the raw methods on
}

// KbnmultiTradeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KbnmultiTradeTransactorRaw struct {
	Contract *KbnmultiTradeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKbnmultiTrade creates a new instance of KbnmultiTrade, bound to a specific deployed contract.
func NewKbnmultiTrade(address common.Address, backend bind.ContractBackend) (*KbnmultiTrade, error) {
	contract, err := bindKbnmultiTrade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KbnmultiTrade{KbnmultiTradeCaller: KbnmultiTradeCaller{contract: contract}, KbnmultiTradeTransactor: KbnmultiTradeTransactor{contract: contract}, KbnmultiTradeFilterer: KbnmultiTradeFilterer{contract: contract}}, nil
}

// NewKbnmultiTradeCaller creates a new read-only instance of KbnmultiTrade, bound to a specific deployed contract.
func NewKbnmultiTradeCaller(address common.Address, caller bind.ContractCaller) (*KbnmultiTradeCaller, error) {
	contract, err := bindKbnmultiTrade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KbnmultiTradeCaller{contract: contract}, nil
}

// NewKbnmultiTradeTransactor creates a new write-only instance of KbnmultiTrade, bound to a specific deployed contract.
func NewKbnmultiTradeTransactor(address common.Address, transactor bind.ContractTransactor) (*KbnmultiTradeTransactor, error) {
	contract, err := bindKbnmultiTrade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KbnmultiTradeTransactor{contract: contract}, nil
}

// NewKbnmultiTradeFilterer creates a new log filterer instance of KbnmultiTrade, bound to a specific deployed contract.
func NewKbnmultiTradeFilterer(address common.Address, filterer bind.ContractFilterer) (*KbnmultiTradeFilterer, error) {
	contract, err := bindKbnmultiTrade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KbnmultiTradeFilterer{contract: contract}, nil
}

// bindKbnmultiTrade binds a generic wrapper to an already deployed contract.
func bindKbnmultiTrade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KbnmultiTradeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KbnmultiTrade *KbnmultiTradeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KbnmultiTrade.Contract.KbnmultiTradeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KbnmultiTrade *KbnmultiTradeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KbnmultiTrade.Contract.KbnmultiTradeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KbnmultiTrade *KbnmultiTradeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KbnmultiTrade.Contract.KbnmultiTradeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KbnmultiTrade *KbnmultiTradeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KbnmultiTrade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KbnmultiTrade *KbnmultiTradeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KbnmultiTrade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KbnmultiTrade *KbnmultiTradeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KbnmultiTrade.Contract.contract.Transact(opts, method, params...)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_KbnmultiTrade *KbnmultiTradeCaller) ETHCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KbnmultiTrade.contract.Call(opts, out, "ETH_CONTRACT_ADDRESS")
	return *ret0, err
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_KbnmultiTrade *KbnmultiTradeSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _KbnmultiTrade.Contract.ETHCONTRACTADDRESS(&_KbnmultiTrade.CallOpts)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_KbnmultiTrade *KbnmultiTradeCallerSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _KbnmultiTrade.Contract.ETHCONTRACTADDRESS(&_KbnmultiTrade.CallOpts)
}

// GetConversionRates is a free data retrieval call binding the contract method 0x0aea8188.
//
// Solidity: function getConversionRates(address srcToken, uint256 srcQty, address destToken) constant returns(uint256, uint256)
func (_KbnmultiTrade *KbnmultiTradeCaller) GetConversionRates(opts *bind.CallOpts, srcToken common.Address, srcQty *big.Int, destToken common.Address) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _KbnmultiTrade.contract.Call(opts, out, "getConversionRates", srcToken, srcQty, destToken)
	return *ret0, *ret1, err
}

// GetConversionRates is a free data retrieval call binding the contract method 0x0aea8188.
//
// Solidity: function getConversionRates(address srcToken, uint256 srcQty, address destToken) constant returns(uint256, uint256)
func (_KbnmultiTrade *KbnmultiTradeSession) GetConversionRates(srcToken common.Address, srcQty *big.Int, destToken common.Address) (*big.Int, *big.Int, error) {
	return _KbnmultiTrade.Contract.GetConversionRates(&_KbnmultiTrade.CallOpts, srcToken, srcQty, destToken)
}

// GetConversionRates is a free data retrieval call binding the contract method 0x0aea8188.
//
// Solidity: function getConversionRates(address srcToken, uint256 srcQty, address destToken) constant returns(uint256, uint256)
func (_KbnmultiTrade *KbnmultiTradeCallerSession) GetConversionRates(srcToken common.Address, srcQty *big.Int, destToken common.Address) (*big.Int, *big.Int, error) {
	return _KbnmultiTrade.Contract.GetConversionRates(&_KbnmultiTrade.CallOpts, srcToken, srcQty, destToken)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_KbnmultiTrade *KbnmultiTradeCaller) IncognitoSmartContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KbnmultiTrade.contract.Call(opts, out, "incognitoSmartContract")
	return *ret0, err
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_KbnmultiTrade *KbnmultiTradeSession) IncognitoSmartContract() (common.Address, error) {
	return _KbnmultiTrade.Contract.IncognitoSmartContract(&_KbnmultiTrade.CallOpts)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_KbnmultiTrade *KbnmultiTradeCallerSession) IncognitoSmartContract() (common.Address, error) {
	return _KbnmultiTrade.Contract.IncognitoSmartContract(&_KbnmultiTrade.CallOpts)
}

// KyberNetworkProxyContract is a free data retrieval call binding the contract method 0x785250da.
//
// Solidity: function kyberNetworkProxyContract() constant returns(address)
func (_KbnmultiTrade *KbnmultiTradeCaller) KyberNetworkProxyContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KbnmultiTrade.contract.Call(opts, out, "kyberNetworkProxyContract")
	return *ret0, err
}

// KyberNetworkProxyContract is a free data retrieval call binding the contract method 0x785250da.
//
// Solidity: function kyberNetworkProxyContract() constant returns(address)
func (_KbnmultiTrade *KbnmultiTradeSession) KyberNetworkProxyContract() (common.Address, error) {
	return _KbnmultiTrade.Contract.KyberNetworkProxyContract(&_KbnmultiTrade.CallOpts)
}

// KyberNetworkProxyContract is a free data retrieval call binding the contract method 0x785250da.
//
// Solidity: function kyberNetworkProxyContract() constant returns(address)
func (_KbnmultiTrade *KbnmultiTradeCallerSession) KyberNetworkProxyContract() (common.Address, error) {
	return _KbnmultiTrade.Contract.KyberNetworkProxyContract(&_KbnmultiTrade.CallOpts)
}

// Trade is a paid mutator transaction binding the contract method 0x0f721819.
//
// Solidity: function trade(address[] srcTokens, uint256[] srcQties, address[] destTokens) returns(address[], uint256[])
func (_KbnmultiTrade *KbnmultiTradeTransactor) Trade(opts *bind.TransactOpts, srcTokens []common.Address, srcQties []*big.Int, destTokens []common.Address) (*types.Transaction, error) {
	return _KbnmultiTrade.contract.Transact(opts, "trade", srcTokens, srcQties, destTokens)
}

// Trade is a paid mutator transaction binding the contract method 0x0f721819.
//
// Solidity: function trade(address[] srcTokens, uint256[] srcQties, address[] destTokens) returns(address[], uint256[])
func (_KbnmultiTrade *KbnmultiTradeSession) Trade(srcTokens []common.Address, srcQties []*big.Int, destTokens []common.Address) (*types.Transaction, error) {
	return _KbnmultiTrade.Contract.Trade(&_KbnmultiTrade.TransactOpts, srcTokens, srcQties, destTokens)
}

// Trade is a paid mutator transaction binding the contract method 0x0f721819.
//
// Solidity: function trade(address[] srcTokens, uint256[] srcQties, address[] destTokens) returns(address[], uint256[])
func (_KbnmultiTrade *KbnmultiTradeTransactorSession) Trade(srcTokens []common.Address, srcQties []*big.Int, destTokens []common.Address) (*types.Transaction, error) {
	return _KbnmultiTrade.Contract.Trade(&_KbnmultiTrade.TransactOpts, srcTokens, srcQties, destTokens)
}
