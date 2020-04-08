// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package compoundAgent

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

// CompoundAgentABI is the input ABI used to generate the binding from.
const CompoundAgentABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_incognitoSmartContract\",\"type\":\"address\"},{\"internalType\":\"contractComptroller\",\"name\":\"_comptroller\",\"type\":\"address\"},{\"internalType\":\"contractCEther\",\"name\":\"_cEther\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proxyCompound\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"ETH_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"addToMarkets\",\"type\":\"address[]\"}],\"name\":\"borrow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"addToMarkets\",\"type\":\"address[]\"}],\"name\":\"borrowByMultiCollateral\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"incognitoSmartContract\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"}],\"name\":\"liquidateBorrow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxyCompound\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isredeemUnderlying\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"exitToMarkets\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"repayBorrow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxyCompound\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CompoundAgent is an auto generated Go binding around an Ethereum contract.
type CompoundAgent struct {
	CompoundAgentCaller     // Read-only binding to the contract
	CompoundAgentTransactor // Write-only binding to the contract
	CompoundAgentFilterer   // Log filterer for contract events
}

// CompoundAgentCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompoundAgentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundAgentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompoundAgentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundAgentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompoundAgentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundAgentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompoundAgentSession struct {
	Contract     *CompoundAgent    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CompoundAgentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompoundAgentCallerSession struct {
	Contract *CompoundAgentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CompoundAgentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompoundAgentTransactorSession struct {
	Contract     *CompoundAgentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CompoundAgentRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompoundAgentRaw struct {
	Contract *CompoundAgent // Generic contract binding to access the raw methods on
}

// CompoundAgentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompoundAgentCallerRaw struct {
	Contract *CompoundAgentCaller // Generic read-only contract binding to access the raw methods on
}

// CompoundAgentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompoundAgentTransactorRaw struct {
	Contract *CompoundAgentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompoundAgent creates a new instance of CompoundAgent, bound to a specific deployed contract.
func NewCompoundAgent(address common.Address, backend bind.ContractBackend) (*CompoundAgent, error) {
	contract, err := bindCompoundAgent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CompoundAgent{CompoundAgentCaller: CompoundAgentCaller{contract: contract}, CompoundAgentTransactor: CompoundAgentTransactor{contract: contract}, CompoundAgentFilterer: CompoundAgentFilterer{contract: contract}}, nil
}

// NewCompoundAgentCaller creates a new read-only instance of CompoundAgent, bound to a specific deployed contract.
func NewCompoundAgentCaller(address common.Address, caller bind.ContractCaller) (*CompoundAgentCaller, error) {
	contract, err := bindCompoundAgent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompoundAgentCaller{contract: contract}, nil
}

// NewCompoundAgentTransactor creates a new write-only instance of CompoundAgent, bound to a specific deployed contract.
func NewCompoundAgentTransactor(address common.Address, transactor bind.ContractTransactor) (*CompoundAgentTransactor, error) {
	contract, err := bindCompoundAgent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompoundAgentTransactor{contract: contract}, nil
}

// NewCompoundAgentFilterer creates a new log filterer instance of CompoundAgent, bound to a specific deployed contract.
func NewCompoundAgentFilterer(address common.Address, filterer bind.ContractFilterer) (*CompoundAgentFilterer, error) {
	contract, err := bindCompoundAgent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompoundAgentFilterer{contract: contract}, nil
}

// bindCompoundAgent binds a generic wrapper to an already deployed contract.
func bindCompoundAgent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CompoundAgentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompoundAgent *CompoundAgentRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CompoundAgent.Contract.CompoundAgentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompoundAgent *CompoundAgentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompoundAgent.Contract.CompoundAgentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompoundAgent *CompoundAgentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompoundAgent.Contract.CompoundAgentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompoundAgent *CompoundAgentCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CompoundAgent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompoundAgent *CompoundAgentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompoundAgent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompoundAgent *CompoundAgentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompoundAgent.Contract.contract.Transact(opts, method, params...)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_CompoundAgent *CompoundAgentCaller) ETHCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CompoundAgent.contract.Call(opts, out, "ETH_CONTRACT_ADDRESS")
	return *ret0, err
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_CompoundAgent *CompoundAgentSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _CompoundAgent.Contract.ETHCONTRACTADDRESS(&_CompoundAgent.CallOpts)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_CompoundAgent *CompoundAgentCallerSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _CompoundAgent.Contract.ETHCONTRACTADDRESS(&_CompoundAgent.CallOpts)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_CompoundAgent *CompoundAgentCaller) IncognitoSmartContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CompoundAgent.contract.Call(opts, out, "incognitoSmartContract")
	return *ret0, err
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_CompoundAgent *CompoundAgentSession) IncognitoSmartContract() (common.Address, error) {
	return _CompoundAgent.Contract.IncognitoSmartContract(&_CompoundAgent.CallOpts)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_CompoundAgent *CompoundAgentCallerSession) IncognitoSmartContract() (common.Address, error) {
	return _CompoundAgent.Contract.IncognitoSmartContract(&_CompoundAgent.CallOpts)
}

// ProxyCompound is a free data retrieval call binding the contract method 0x370f81fe.
//
// Solidity: function proxyCompound() constant returns(address)
func (_CompoundAgent *CompoundAgentCaller) ProxyCompound(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CompoundAgent.contract.Call(opts, out, "proxyCompound")
	return *ret0, err
}

// ProxyCompound is a free data retrieval call binding the contract method 0x370f81fe.
//
// Solidity: function proxyCompound() constant returns(address)
func (_CompoundAgent *CompoundAgentSession) ProxyCompound() (common.Address, error) {
	return _CompoundAgent.Contract.ProxyCompound(&_CompoundAgent.CallOpts)
}

// ProxyCompound is a free data retrieval call binding the contract method 0x370f81fe.
//
// Solidity: function proxyCompound() constant returns(address)
func (_CompoundAgent *CompoundAgentCallerSession) ProxyCompound() (common.Address, error) {
	return _CompoundAgent.Contract.ProxyCompound(&_CompoundAgent.CallOpts)
}

// Borrow is a paid mutator transaction binding the contract method 0x26d5e251.
//
// Solidity: function borrow(address cToken, uint256 amount, address[] addToMarkets) returns(address, uint256)
func (_CompoundAgent *CompoundAgentTransactor) Borrow(opts *bind.TransactOpts, cToken common.Address, amount *big.Int, addToMarkets []common.Address) (*types.Transaction, error) {
	return _CompoundAgent.contract.Transact(opts, "borrow", cToken, amount, addToMarkets)
}

// Borrow is a paid mutator transaction binding the contract method 0x26d5e251.
//
// Solidity: function borrow(address cToken, uint256 amount, address[] addToMarkets) returns(address, uint256)
func (_CompoundAgent *CompoundAgentSession) Borrow(cToken common.Address, amount *big.Int, addToMarkets []common.Address) (*types.Transaction, error) {
	return _CompoundAgent.Contract.Borrow(&_CompoundAgent.TransactOpts, cToken, amount, addToMarkets)
}

// Borrow is a paid mutator transaction binding the contract method 0x26d5e251.
//
// Solidity: function borrow(address cToken, uint256 amount, address[] addToMarkets) returns(address, uint256)
func (_CompoundAgent *CompoundAgentTransactorSession) Borrow(cToken common.Address, amount *big.Int, addToMarkets []common.Address) (*types.Transaction, error) {
	return _CompoundAgent.Contract.Borrow(&_CompoundAgent.TransactOpts, cToken, amount, addToMarkets)
}

// BorrowByMultiCollateral is a paid mutator transaction binding the contract method 0xb6dbc8ed.
//
// Solidity: function borrowByMultiCollateral(address cToken, uint256 amount, address[] addToMarkets) returns(address[], uint256[])
func (_CompoundAgent *CompoundAgentTransactor) BorrowByMultiCollateral(opts *bind.TransactOpts, cToken common.Address, amount *big.Int, addToMarkets []common.Address) (*types.Transaction, error) {
	return _CompoundAgent.contract.Transact(opts, "borrowByMultiCollateral", cToken, amount, addToMarkets)
}

// BorrowByMultiCollateral is a paid mutator transaction binding the contract method 0xb6dbc8ed.
//
// Solidity: function borrowByMultiCollateral(address cToken, uint256 amount, address[] addToMarkets) returns(address[], uint256[])
func (_CompoundAgent *CompoundAgentSession) BorrowByMultiCollateral(cToken common.Address, amount *big.Int, addToMarkets []common.Address) (*types.Transaction, error) {
	return _CompoundAgent.Contract.BorrowByMultiCollateral(&_CompoundAgent.TransactOpts, cToken, amount, addToMarkets)
}

// BorrowByMultiCollateral is a paid mutator transaction binding the contract method 0xb6dbc8ed.
//
// Solidity: function borrowByMultiCollateral(address cToken, uint256 amount, address[] addToMarkets) returns(address[], uint256[])
func (_CompoundAgent *CompoundAgentTransactorSession) BorrowByMultiCollateral(cToken common.Address, amount *big.Int, addToMarkets []common.Address) (*types.Transaction, error) {
	return _CompoundAgent.Contract.BorrowByMultiCollateral(&_CompoundAgent.TransactOpts, cToken, amount, addToMarkets)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0x64fd7078.
//
// Solidity: function liquidateBorrow(address cToken, address borrower, uint256 repayAmount, address cTokenCollateral) returns(address, uint256)
func (_CompoundAgent *CompoundAgentTransactor) LiquidateBorrow(opts *bind.TransactOpts, cToken common.Address, borrower common.Address, repayAmount *big.Int, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _CompoundAgent.contract.Transact(opts, "liquidateBorrow", cToken, borrower, repayAmount, cTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0x64fd7078.
//
// Solidity: function liquidateBorrow(address cToken, address borrower, uint256 repayAmount, address cTokenCollateral) returns(address, uint256)
func (_CompoundAgent *CompoundAgentSession) LiquidateBorrow(cToken common.Address, borrower common.Address, repayAmount *big.Int, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _CompoundAgent.Contract.LiquidateBorrow(&_CompoundAgent.TransactOpts, cToken, borrower, repayAmount, cTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0x64fd7078.
//
// Solidity: function liquidateBorrow(address cToken, address borrower, uint256 repayAmount, address cTokenCollateral) returns(address, uint256)
func (_CompoundAgent *CompoundAgentTransactorSession) LiquidateBorrow(cToken common.Address, borrower common.Address, repayAmount *big.Int, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _CompoundAgent.Contract.LiquidateBorrow(&_CompoundAgent.TransactOpts, cToken, borrower, repayAmount, cTokenCollateral)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address cToken, uint256 amount) returns(address, uint256)
func (_CompoundAgent *CompoundAgentTransactor) Mint(opts *bind.TransactOpts, cToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CompoundAgent.contract.Transact(opts, "mint", cToken, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address cToken, uint256 amount) returns(address, uint256)
func (_CompoundAgent *CompoundAgentSession) Mint(cToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CompoundAgent.Contract.Mint(&_CompoundAgent.TransactOpts, cToken, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address cToken, uint256 amount) returns(address, uint256)
func (_CompoundAgent *CompoundAgentTransactorSession) Mint(cToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CompoundAgent.Contract.Mint(&_CompoundAgent.TransactOpts, cToken, amount)
}

// Redeem is a paid mutator transaction binding the contract method 0x7520f7ed.
//
// Solidity: function redeem(address cToken, uint256 amount, bool isredeemUnderlying, address exitToMarkets) returns(address, uint256)
func (_CompoundAgent *CompoundAgentTransactor) Redeem(opts *bind.TransactOpts, cToken common.Address, amount *big.Int, isredeemUnderlying bool, exitToMarkets common.Address) (*types.Transaction, error) {
	return _CompoundAgent.contract.Transact(opts, "redeem", cToken, amount, isredeemUnderlying, exitToMarkets)
}

// Redeem is a paid mutator transaction binding the contract method 0x7520f7ed.
//
// Solidity: function redeem(address cToken, uint256 amount, bool isredeemUnderlying, address exitToMarkets) returns(address, uint256)
func (_CompoundAgent *CompoundAgentSession) Redeem(cToken common.Address, amount *big.Int, isredeemUnderlying bool, exitToMarkets common.Address) (*types.Transaction, error) {
	return _CompoundAgent.Contract.Redeem(&_CompoundAgent.TransactOpts, cToken, amount, isredeemUnderlying, exitToMarkets)
}

// Redeem is a paid mutator transaction binding the contract method 0x7520f7ed.
//
// Solidity: function redeem(address cToken, uint256 amount, bool isredeemUnderlying, address exitToMarkets) returns(address, uint256)
func (_CompoundAgent *CompoundAgentTransactorSession) Redeem(cToken common.Address, amount *big.Int, isredeemUnderlying bool, exitToMarkets common.Address) (*types.Transaction, error) {
	return _CompoundAgent.Contract.Redeem(&_CompoundAgent.TransactOpts, cToken, amount, isredeemUnderlying, exitToMarkets)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0xabdb5ea8.
//
// Solidity: function repayBorrow(address cToken, uint256 amount) returns(address, uint256)
func (_CompoundAgent *CompoundAgentTransactor) RepayBorrow(opts *bind.TransactOpts, cToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CompoundAgent.contract.Transact(opts, "repayBorrow", cToken, amount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0xabdb5ea8.
//
// Solidity: function repayBorrow(address cToken, uint256 amount) returns(address, uint256)
func (_CompoundAgent *CompoundAgentSession) RepayBorrow(cToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CompoundAgent.Contract.RepayBorrow(&_CompoundAgent.TransactOpts, cToken, amount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0xabdb5ea8.
//
// Solidity: function repayBorrow(address cToken, uint256 amount) returns(address, uint256)
func (_CompoundAgent *CompoundAgentTransactorSession) RepayBorrow(cToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CompoundAgent.Contract.RepayBorrow(&_CompoundAgent.TransactOpts, cToken, amount)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address _proxyCompound) returns(address, uint256)
func (_CompoundAgent *CompoundAgentTransactor) Upgrade(opts *bind.TransactOpts, _proxyCompound common.Address) (*types.Transaction, error) {
	return _CompoundAgent.contract.Transact(opts, "upgrade", _proxyCompound)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address _proxyCompound) returns(address, uint256)
func (_CompoundAgent *CompoundAgentSession) Upgrade(_proxyCompound common.Address) (*types.Transaction, error) {
	return _CompoundAgent.Contract.Upgrade(&_CompoundAgent.TransactOpts, _proxyCompound)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address _proxyCompound) returns(address, uint256)
func (_CompoundAgent *CompoundAgentTransactorSession) Upgrade(_proxyCompound common.Address) (*types.Transaction, error) {
	return _CompoundAgent.Contract.Upgrade(&_CompoundAgent.TransactOpts, _proxyCompound)
}
