// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kbntrade

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

// KbntradeABI is the input ABI used to generate the binding from.
const KbntradeABI = "[{\"inputs\":[{\"internalType\":\"contractKyberNetwork\",\"name\":\"_kyberNetworkProxyContract\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_incognitoSmartContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"ETH_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"incognitoSmartContract\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kyberNetworkProxyContract\",\"outputs\":[{\"internalType\":\"contractKyberNetwork\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcQty\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"getConversionRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcQty\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"trade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// KbntradeBin is the compiled bytecode used for deploying new contracts.
const KbntradeBin = `0x608060405234801561001057600080fd5b50604051610a7f380380610a7f8339818101604052604081101561003357600080fd5b50805160209091015160018054600160a060020a03938416600160a060020a03199182161790915560008054939092169216919091179055610a058061007a6000396000f3fe60806040526004361061004d5760e060020a60003504630aea8188811461004f5780630e32db52146100ab57806372e94bf614610104578063785250da14610135578063b42a644b1461014a575b005b34801561005b57600080fd5b506100926004803603606081101561007257600080fd5b50600160a060020a0381358116916020810135916040909101351661015f565b6040805192835260208301919091528051918290030190f35b6100e1600480360360608110156100c157600080fd5b50600160a060020a038135811691602081013591604090910135166101fd565b60408051600160a060020a03909316835260208301919091528051918290030190f35b34801561011057600080fd5b506101196102f1565b60408051600160a060020a039092168252519081900360200190f35b34801561014157600080fd5b506101196102f6565b34801561015657600080fd5b50610119610305565b6001546040805160e060020a63809a9e55028152600160a060020a0386811660048301528481166024830152604482018690528251600094859492169263809a9e55926064808301939192829003018186803b1580156101be57600080fd5b505afa1580156101d2573d6000803e3d6000fd5b505050506040513d60408110156101e857600080fd5b50805160209091015190969095509350505050565b600080548190600160a060020a0316331461021757600080fd5b8361022186610314565b101561022c57600080fd5b82600160a060020a031685600160a060020a0316141561024b57600080fd5b6000600160a060020a038616156102ba57600154610274908790600160a060020a0316876103bf565b600160a060020a0384161561029f576000610290878787610504565b1161029a57600080fd5b6102b5565b60006102ab8787610644565b116102b557600080fd5b6102d0565b60006102c6858761078d565b116102d057600080fd5b6102d984610314565b90506102e584826108d4565b92959294509192505050565b600081565b600154600160a060020a031681565b600054600160a060020a031681565b6000600160a060020a03821661032c575030316103ba565b604080517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529051600160a060020a038416916370a08231916024808301926020929190829003018186803b15801561038b57600080fd5b505afa15801561039f573d6000803e3d6000fd5b505050506040513d60208110156103b557600080fd5b505190505b919050565b600160a060020a038316156104ff5782600160a060020a031663095ea7b38360006040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b15801561043257600080fd5b505af1158015610446573d6000803e3d6000fd5b505050506040513d602081101561045c57600080fd5b505161046757600080fd5b82600160a060020a031663095ea7b383836040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b1580156104ca57600080fd5b505af11580156104de573d6000803e3d6000fd5b505050506040513d60208110156104f457600080fd5b50516104ff57600080fd5b505050565b6001546040805160e060020a63809a9e55028152600160a060020a0386811660048301528481166024830152604482018690528251600094859492169263809a9e55926064808301939192829003018186803b15801561056357600080fd5b505afa158015610577573d6000803e3d6000fd5b505050506040513d604081101561058d57600080fd5b5051600154604080517f7409e2eb000000000000000000000000000000000000000000000000000000008152600160a060020a038981166004830152602482018990528781166044830152606482018590529151939450911691637409e2eb916084808201926020929091908290030181600087803b15801561060f57600080fd5b505af1158015610623573d6000803e3d6000fd5b505050506040513d602081101561063957600080fd5b505195945050505050565b6001546040805160e060020a63809a9e55028152600160a060020a03858116600483015273eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee6024830152604482018590528251600094859492169263809a9e55926064808301939192829003018186803b1580156106b557600080fd5b505afa1580156106c9573d6000803e3d6000fd5b505050506040513d60408110156106df57600080fd5b5051600154604080517f3bba21dc000000000000000000000000000000000000000000000000000000008152600160a060020a03888116600483015260248201889052604482018590529151939450911691633bba21dc916064808201926020929091908290030181600087803b15801561075957600080fd5b505af115801561076d573d6000803e3d6000fd5b505050506040513d602081101561078357600080fd5b5051949350505050565b6000303182111561079d57600080fd5b6001546040805160e060020a63809a9e5502815273eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee6004820152600160a060020a0386811660248301523460448301528251600094919091169263809a9e559260648082019391829003018186803b15801561080c57600080fd5b505afa158015610820573d6000803e3d6000fd5b505050506040513d604081101561083657600080fd5b5051600154604080517f7a2a0456000000000000000000000000000000000000000000000000000000008152600160a060020a038881166004830152602482018590529151939450911691637a2a0456918691604480830192602092919082900301818588803b1580156108a957600080fd5b505af11580156108bd573d6000803e3d6000fd5b50505050506040513d602081101561078357600080fd5b600160a060020a03821661092f5730318111156108f057600080fd5b60008054604051600160a060020a039091169183156108fc02918491818181858888f19350505050158015610929573d6000803e3d6000fd5b506109cc565b60008054604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a0392831660048201526024810185905290519185169263a9059cbb926044808401936020939083900390910190829087803b15801561099f57600080fd5b505af11580156109b3573d6000803e3d6000fd5b505050506040513d60208110156109c957600080fd5b50505b505056fea265627a7a723158201430a7c4188c3ebc9c07978124866665ceae8829f7eb81a9fb1dd78ff587f4af64736f6c634300050c0032`

// DeployKbntrade deploys a new Ethereum contract, binding an instance of Kbntrade to it.
func DeployKbntrade(auth *bind.TransactOpts, backend bind.ContractBackend, _kyberNetworkProxyContract common.Address, _incognitoSmartContract common.Address) (common.Address, *types.Transaction, *Kbntrade, error) {
	parsed, err := abi.JSON(strings.NewReader(KbntradeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KbntradeBin), backend, _kyberNetworkProxyContract, _incognitoSmartContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Kbntrade{KbntradeCaller: KbntradeCaller{contract: contract}, KbntradeTransactor: KbntradeTransactor{contract: contract}, KbntradeFilterer: KbntradeFilterer{contract: contract}}, nil
}

// Kbntrade is an auto generated Go binding around an Ethereum contract.
type Kbntrade struct {
	KbntradeCaller     // Read-only binding to the contract
	KbntradeTransactor // Write-only binding to the contract
	KbntradeFilterer   // Log filterer for contract events
}

// KbntradeCaller is an auto generated read-only Go binding around an Ethereum contract.
type KbntradeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KbntradeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KbntradeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KbntradeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KbntradeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KbntradeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KbntradeSession struct {
	Contract     *Kbntrade         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KbntradeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KbntradeCallerSession struct {
	Contract *KbntradeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// KbntradeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KbntradeTransactorSession struct {
	Contract     *KbntradeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// KbntradeRaw is an auto generated low-level Go binding around an Ethereum contract.
type KbntradeRaw struct {
	Contract *Kbntrade // Generic contract binding to access the raw methods on
}

// KbntradeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KbntradeCallerRaw struct {
	Contract *KbntradeCaller // Generic read-only contract binding to access the raw methods on
}

// KbntradeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KbntradeTransactorRaw struct {
	Contract *KbntradeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKbntrade creates a new instance of Kbntrade, bound to a specific deployed contract.
func NewKbntrade(address common.Address, backend bind.ContractBackend) (*Kbntrade, error) {
	contract, err := bindKbntrade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Kbntrade{KbntradeCaller: KbntradeCaller{contract: contract}, KbntradeTransactor: KbntradeTransactor{contract: contract}, KbntradeFilterer: KbntradeFilterer{contract: contract}}, nil
}

// NewKbntradeCaller creates a new read-only instance of Kbntrade, bound to a specific deployed contract.
func NewKbntradeCaller(address common.Address, caller bind.ContractCaller) (*KbntradeCaller, error) {
	contract, err := bindKbntrade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KbntradeCaller{contract: contract}, nil
}

// NewKbntradeTransactor creates a new write-only instance of Kbntrade, bound to a specific deployed contract.
func NewKbntradeTransactor(address common.Address, transactor bind.ContractTransactor) (*KbntradeTransactor, error) {
	contract, err := bindKbntrade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KbntradeTransactor{contract: contract}, nil
}

// NewKbntradeFilterer creates a new log filterer instance of Kbntrade, bound to a specific deployed contract.
func NewKbntradeFilterer(address common.Address, filterer bind.ContractFilterer) (*KbntradeFilterer, error) {
	contract, err := bindKbntrade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KbntradeFilterer{contract: contract}, nil
}

// bindKbntrade binds a generic wrapper to an already deployed contract.
func bindKbntrade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KbntradeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kbntrade *KbntradeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Kbntrade.Contract.KbntradeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kbntrade *KbntradeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kbntrade.Contract.KbntradeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kbntrade *KbntradeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kbntrade.Contract.KbntradeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kbntrade *KbntradeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Kbntrade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kbntrade *KbntradeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kbntrade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kbntrade *KbntradeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kbntrade.Contract.contract.Transact(opts, method, params...)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_Kbntrade *KbntradeCaller) ETHCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Kbntrade.contract.Call(opts, out, "ETH_CONTRACT_ADDRESS")
	return *ret0, err
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_Kbntrade *KbntradeSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _Kbntrade.Contract.ETHCONTRACTADDRESS(&_Kbntrade.CallOpts)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_Kbntrade *KbntradeCallerSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _Kbntrade.Contract.ETHCONTRACTADDRESS(&_Kbntrade.CallOpts)
}

// GetConversionRates is a free data retrieval call binding the contract method 0x0aea8188.
//
// Solidity: function getConversionRates(address srcToken, uint256 srcQty, address destToken) constant returns(uint256, uint256)
func (_Kbntrade *KbntradeCaller) GetConversionRates(opts *bind.CallOpts, srcToken common.Address, srcQty *big.Int, destToken common.Address) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Kbntrade.contract.Call(opts, out, "getConversionRates", srcToken, srcQty, destToken)
	return *ret0, *ret1, err
}

// GetConversionRates is a free data retrieval call binding the contract method 0x0aea8188.
//
// Solidity: function getConversionRates(address srcToken, uint256 srcQty, address destToken) constant returns(uint256, uint256)
func (_Kbntrade *KbntradeSession) GetConversionRates(srcToken common.Address, srcQty *big.Int, destToken common.Address) (*big.Int, *big.Int, error) {
	return _Kbntrade.Contract.GetConversionRates(&_Kbntrade.CallOpts, srcToken, srcQty, destToken)
}

// GetConversionRates is a free data retrieval call binding the contract method 0x0aea8188.
//
// Solidity: function getConversionRates(address srcToken, uint256 srcQty, address destToken) constant returns(uint256, uint256)
func (_Kbntrade *KbntradeCallerSession) GetConversionRates(srcToken common.Address, srcQty *big.Int, destToken common.Address) (*big.Int, *big.Int, error) {
	return _Kbntrade.Contract.GetConversionRates(&_Kbntrade.CallOpts, srcToken, srcQty, destToken)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_Kbntrade *KbntradeCaller) IncognitoSmartContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Kbntrade.contract.Call(opts, out, "incognitoSmartContract")
	return *ret0, err
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_Kbntrade *KbntradeSession) IncognitoSmartContract() (common.Address, error) {
	return _Kbntrade.Contract.IncognitoSmartContract(&_Kbntrade.CallOpts)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_Kbntrade *KbntradeCallerSession) IncognitoSmartContract() (common.Address, error) {
	return _Kbntrade.Contract.IncognitoSmartContract(&_Kbntrade.CallOpts)
}

// KyberNetworkProxyContract is a free data retrieval call binding the contract method 0x785250da.
//
// Solidity: function kyberNetworkProxyContract() constant returns(address)
func (_Kbntrade *KbntradeCaller) KyberNetworkProxyContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Kbntrade.contract.Call(opts, out, "kyberNetworkProxyContract")
	return *ret0, err
}

// KyberNetworkProxyContract is a free data retrieval call binding the contract method 0x785250da.
//
// Solidity: function kyberNetworkProxyContract() constant returns(address)
func (_Kbntrade *KbntradeSession) KyberNetworkProxyContract() (common.Address, error) {
	return _Kbntrade.Contract.KyberNetworkProxyContract(&_Kbntrade.CallOpts)
}

// KyberNetworkProxyContract is a free data retrieval call binding the contract method 0x785250da.
//
// Solidity: function kyberNetworkProxyContract() constant returns(address)
func (_Kbntrade *KbntradeCallerSession) KyberNetworkProxyContract() (common.Address, error) {
	return _Kbntrade.Contract.KyberNetworkProxyContract(&_Kbntrade.CallOpts)
}

// Trade is a paid mutator transaction binding the contract method 0x0e32db52.
//
// Solidity: function trade(address srcToken, uint256 srcQty, address destToken) returns(address, uint256)
func (_Kbntrade *KbntradeTransactor) Trade(opts *bind.TransactOpts, srcToken common.Address, srcQty *big.Int, destToken common.Address) (*types.Transaction, error) {
	return _Kbntrade.contract.Transact(opts, "trade", srcToken, srcQty, destToken)
}

// Trade is a paid mutator transaction binding the contract method 0x0e32db52.
//
// Solidity: function trade(address srcToken, uint256 srcQty, address destToken) returns(address, uint256)
func (_Kbntrade *KbntradeSession) Trade(srcToken common.Address, srcQty *big.Int, destToken common.Address) (*types.Transaction, error) {
	return _Kbntrade.Contract.Trade(&_Kbntrade.TransactOpts, srcToken, srcQty, destToken)
}

// Trade is a paid mutator transaction binding the contract method 0x0e32db52.
//
// Solidity: function trade(address srcToken, uint256 srcQty, address destToken) returns(address, uint256)
func (_Kbntrade *KbntradeTransactorSession) Trade(srcToken common.Address, srcQty *big.Int, destToken common.Address) (*types.Transaction, error) {
	return _Kbntrade.Contract.Trade(&_Kbntrade.TransactOpts, srcToken, srcQty, destToken)
}
