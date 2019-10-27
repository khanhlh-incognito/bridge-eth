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

// AdminPausableABI is the input ABI used to generate the binding from.
const AdminPausableABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ndays\",\"type\":\"uint256\"}],\"name\":\"Extend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"expire\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"extend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_successor\",\"type\":\"address\"}],\"name\":\"retire\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"successor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AdminPausableBin is the compiled bytecode used for deploying new contracts.
const AdminPausableBin = `0x608060405234801561001057600080fd5b506040516106fb3803806106fb8339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b03199092169190911790556001805460ff60a01b191690556301e13380420160025561067f8061007c6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c806379599f961161006657806379599f96146100ea5780638456cb59146101045780639714378c1461010c5780639e6371ba14610129578063f851a4401461014f57610093565b80633f4ba83a146100985780634e71d92d146100a25780635c975abb146100aa5780636ff968c3146100c6575b600080fd5b6100a0610157565b005b6100a0610239565b6100b2610320565b604080519115158252519081900360200190f35b6100ce610330565b604080516001600160a01b039092168252519081900360200190f35b6100f261033f565b60408051918252519081900360200190f35b6100a0610345565b6100a06004803603602081101561012257600080fd5b503561046a565b6100a06004803603602081101561013f57600080fd5b50356001600160a01b031661058e565b6100ce61063b565b6000546001600160a01b031633146101a2576040805162461bcd60e51b81526020600482015260096024820152683737ba1030b236b4b760b91b604482015290519081900360640190fd5b600154600160a01b900460ff166101f7576040805162461bcd60e51b81526020600482015260146024820152736e6f7420706175736564207269676874206e6f7760601b604482015290519081900360640190fd5b6001805460ff60a01b191690556040805133815290517f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa9181900360200190a1565b6002544210610279576040805162461bcd60e51b8152602060048201526007602482015266195e1c1a5c995960ca1b604482015290519081900360640190fd5b6001546001600160a01b031633146102c7576040805162461bcd60e51b815260206004820152600c60248201526b1d5b985d5d1a1bdc9a5e995960a21b604482015290519081900360640190fd5b600154600080546001600160a01b0319166001600160a01b0392831617908190556040805191909216815290517f0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc9181900360200190a1565b600154600160a01b900460ff1681565b6001546001600160a01b031681565b60025481565b6000546001600160a01b03163314610390576040805162461bcd60e51b81526020600482015260096024820152683737ba1030b236b4b760b91b604482015290519081900360640190fd5b600154600160a01b900460ff16156103e2576040805162461bcd60e51b815260206004820152601060248201526f706175736564207269676874206e6f7760801b604482015290519081900360640190fd5b6002544210610422576040805162461bcd60e51b8152602060048201526007602482015266195e1c1a5c995960ca1b604482015290519081900360640190fd5b6001805460ff60a01b1916600160a01b1790556040805133815290517f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2589181900360200190a1565b6000546001600160a01b031633146104b5576040805162461bcd60e51b81526020600482015260096024820152683737ba1030b236b4b760b91b604482015290519081900360640190fd5b60025442106104f5576040805162461bcd60e51b8152602060048201526007602482015266195e1c1a5c995960ca1b604482015290519081900360640190fd5b61016e811061054b576040805162461bcd60e51b815260206004820152601a60248201527f63616e6e6f7420657874656e6420666f7220746f6f206c6f6e67000000000000604482015290519081900360640190fd5b600280546201518083020190556040805182815290517f02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e89181900360200190a150565b6000546001600160a01b031633146105d9576040805162461bcd60e51b81526020600482015260096024820152683737ba1030b236b4b760b91b604482015290519081900360640190fd5b6002544210610619576040805162461bcd60e51b8152602060048201526007602482015266195e1c1a5c995960ca1b604482015290519081900360640190fd5b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b03168156fea265627a7a7231582010bbfdbfe6f11fa642302cfc8fd52669d48b65f0b18d0f9520fe6de262f4037664736f6c634300050c0032`

// DeployAdminPausable deploys a new Ethereum contract, binding an instance of AdminPausable to it.
func DeployAdminPausable(auth *bind.TransactOpts, backend bind.ContractBackend, _admin common.Address) (common.Address, *types.Transaction, *AdminPausable, error) {
	parsed, err := abi.JSON(strings.NewReader(AdminPausableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AdminPausableBin), backend, _admin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AdminPausable{AdminPausableCaller: AdminPausableCaller{contract: contract}, AdminPausableTransactor: AdminPausableTransactor{contract: contract}, AdminPausableFilterer: AdminPausableFilterer{contract: contract}}, nil
}

// AdminPausable is an auto generated Go binding around an Ethereum contract.
type AdminPausable struct {
	AdminPausableCaller     // Read-only binding to the contract
	AdminPausableTransactor // Write-only binding to the contract
	AdminPausableFilterer   // Log filterer for contract events
}

// AdminPausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdminPausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminPausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdminPausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminPausableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdminPausableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminPausableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdminPausableSession struct {
	Contract     *AdminPausable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdminPausableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdminPausableCallerSession struct {
	Contract *AdminPausableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AdminPausableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdminPausableTransactorSession struct {
	Contract     *AdminPausableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AdminPausableRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdminPausableRaw struct {
	Contract *AdminPausable // Generic contract binding to access the raw methods on
}

// AdminPausableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdminPausableCallerRaw struct {
	Contract *AdminPausableCaller // Generic read-only contract binding to access the raw methods on
}

// AdminPausableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdminPausableTransactorRaw struct {
	Contract *AdminPausableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdminPausable creates a new instance of AdminPausable, bound to a specific deployed contract.
func NewAdminPausable(address common.Address, backend bind.ContractBackend) (*AdminPausable, error) {
	contract, err := bindAdminPausable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AdminPausable{AdminPausableCaller: AdminPausableCaller{contract: contract}, AdminPausableTransactor: AdminPausableTransactor{contract: contract}, AdminPausableFilterer: AdminPausableFilterer{contract: contract}}, nil
}

// NewAdminPausableCaller creates a new read-only instance of AdminPausable, bound to a specific deployed contract.
func NewAdminPausableCaller(address common.Address, caller bind.ContractCaller) (*AdminPausableCaller, error) {
	contract, err := bindAdminPausable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdminPausableCaller{contract: contract}, nil
}

// NewAdminPausableTransactor creates a new write-only instance of AdminPausable, bound to a specific deployed contract.
func NewAdminPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*AdminPausableTransactor, error) {
	contract, err := bindAdminPausable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdminPausableTransactor{contract: contract}, nil
}

// NewAdminPausableFilterer creates a new log filterer instance of AdminPausable, bound to a specific deployed contract.
func NewAdminPausableFilterer(address common.Address, filterer bind.ContractFilterer) (*AdminPausableFilterer, error) {
	contract, err := bindAdminPausable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdminPausableFilterer{contract: contract}, nil
}

// bindAdminPausable binds a generic wrapper to an already deployed contract.
func bindAdminPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdminPausableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdminPausable *AdminPausableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AdminPausable.Contract.AdminPausableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdminPausable *AdminPausableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminPausable.Contract.AdminPausableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdminPausable *AdminPausableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdminPausable.Contract.AdminPausableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdminPausable *AdminPausableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AdminPausable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdminPausable *AdminPausableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminPausable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdminPausable *AdminPausableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdminPausable.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_AdminPausable *AdminPausableCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AdminPausable.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_AdminPausable *AdminPausableSession) Admin() (common.Address, error) {
	return _AdminPausable.Contract.Admin(&_AdminPausable.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_AdminPausable *AdminPausableCallerSession) Admin() (common.Address, error) {
	return _AdminPausable.Contract.Admin(&_AdminPausable.CallOpts)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() constant returns(uint256)
func (_AdminPausable *AdminPausableCaller) Expire(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AdminPausable.contract.Call(opts, out, "expire")
	return *ret0, err
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() constant returns(uint256)
func (_AdminPausable *AdminPausableSession) Expire() (*big.Int, error) {
	return _AdminPausable.Contract.Expire(&_AdminPausable.CallOpts)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() constant returns(uint256)
func (_AdminPausable *AdminPausableCallerSession) Expire() (*big.Int, error) {
	return _AdminPausable.Contract.Expire(&_AdminPausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_AdminPausable *AdminPausableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AdminPausable.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_AdminPausable *AdminPausableSession) Paused() (bool, error) {
	return _AdminPausable.Contract.Paused(&_AdminPausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_AdminPausable *AdminPausableCallerSession) Paused() (bool, error) {
	return _AdminPausable.Contract.Paused(&_AdminPausable.CallOpts)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() constant returns(address)
func (_AdminPausable *AdminPausableCaller) Successor(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AdminPausable.contract.Call(opts, out, "successor")
	return *ret0, err
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() constant returns(address)
func (_AdminPausable *AdminPausableSession) Successor() (common.Address, error) {
	return _AdminPausable.Contract.Successor(&_AdminPausable.CallOpts)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() constant returns(address)
func (_AdminPausable *AdminPausableCallerSession) Successor() (common.Address, error) {
	return _AdminPausable.Contract.Successor(&_AdminPausable.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_AdminPausable *AdminPausableTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminPausable.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_AdminPausable *AdminPausableSession) Claim() (*types.Transaction, error) {
	return _AdminPausable.Contract.Claim(&_AdminPausable.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_AdminPausable *AdminPausableTransactorSession) Claim() (*types.Transaction, error) {
	return _AdminPausable.Contract.Claim(&_AdminPausable.TransactOpts)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_AdminPausable *AdminPausableTransactor) Extend(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _AdminPausable.contract.Transact(opts, "extend", n)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_AdminPausable *AdminPausableSession) Extend(n *big.Int) (*types.Transaction, error) {
	return _AdminPausable.Contract.Extend(&_AdminPausable.TransactOpts, n)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_AdminPausable *AdminPausableTransactorSession) Extend(n *big.Int) (*types.Transaction, error) {
	return _AdminPausable.Contract.Extend(&_AdminPausable.TransactOpts, n)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AdminPausable *AdminPausableTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminPausable.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AdminPausable *AdminPausableSession) Pause() (*types.Transaction, error) {
	return _AdminPausable.Contract.Pause(&_AdminPausable.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AdminPausable *AdminPausableTransactorSession) Pause() (*types.Transaction, error) {
	return _AdminPausable.Contract.Pause(&_AdminPausable.TransactOpts)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_AdminPausable *AdminPausableTransactor) Retire(opts *bind.TransactOpts, _successor common.Address) (*types.Transaction, error) {
	return _AdminPausable.contract.Transact(opts, "retire", _successor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_AdminPausable *AdminPausableSession) Retire(_successor common.Address) (*types.Transaction, error) {
	return _AdminPausable.Contract.Retire(&_AdminPausable.TransactOpts, _successor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_AdminPausable *AdminPausableTransactorSession) Retire(_successor common.Address) (*types.Transaction, error) {
	return _AdminPausable.Contract.Retire(&_AdminPausable.TransactOpts, _successor)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AdminPausable *AdminPausableTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminPausable.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AdminPausable *AdminPausableSession) Unpause() (*types.Transaction, error) {
	return _AdminPausable.Contract.Unpause(&_AdminPausable.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AdminPausable *AdminPausableTransactorSession) Unpause() (*types.Transaction, error) {
	return _AdminPausable.Contract.Unpause(&_AdminPausable.TransactOpts)
}

// AdminPausableClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the AdminPausable contract.
type AdminPausableClaimIterator struct {
	Event *AdminPausableClaim // Event containing the contract specifics and raw log

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
func (it *AdminPausableClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminPausableClaim)
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
		it.Event = new(AdminPausableClaim)
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
func (it *AdminPausableClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminPausableClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminPausableClaim represents a Claim event raised by the AdminPausable contract.
type AdminPausableClaim struct {
	Claimer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_AdminPausable *AdminPausableFilterer) FilterClaim(opts *bind.FilterOpts) (*AdminPausableClaimIterator, error) {

	logs, sub, err := _AdminPausable.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &AdminPausableClaimIterator{contract: _AdminPausable.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_AdminPausable *AdminPausableFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *AdminPausableClaim) (event.Subscription, error) {

	logs, sub, err := _AdminPausable.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminPausableClaim)
				if err := _AdminPausable.contract.UnpackLog(event, "Claim", log); err != nil {
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

// AdminPausableExtendIterator is returned from FilterExtend and is used to iterate over the raw logs and unpacked data for Extend events raised by the AdminPausable contract.
type AdminPausableExtendIterator struct {
	Event *AdminPausableExtend // Event containing the contract specifics and raw log

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
func (it *AdminPausableExtendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminPausableExtend)
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
		it.Event = new(AdminPausableExtend)
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
func (it *AdminPausableExtendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminPausableExtendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminPausableExtend represents a Extend event raised by the AdminPausable contract.
type AdminPausableExtend struct {
	Ndays *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterExtend is a free log retrieval operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_AdminPausable *AdminPausableFilterer) FilterExtend(opts *bind.FilterOpts) (*AdminPausableExtendIterator, error) {

	logs, sub, err := _AdminPausable.contract.FilterLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return &AdminPausableExtendIterator{contract: _AdminPausable.contract, event: "Extend", logs: logs, sub: sub}, nil
}

// WatchExtend is a free log subscription operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_AdminPausable *AdminPausableFilterer) WatchExtend(opts *bind.WatchOpts, sink chan<- *AdminPausableExtend) (event.Subscription, error) {

	logs, sub, err := _AdminPausable.contract.WatchLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminPausableExtend)
				if err := _AdminPausable.contract.UnpackLog(event, "Extend", log); err != nil {
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

// AdminPausablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the AdminPausable contract.
type AdminPausablePausedIterator struct {
	Event *AdminPausablePaused // Event containing the contract specifics and raw log

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
func (it *AdminPausablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminPausablePaused)
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
		it.Event = new(AdminPausablePaused)
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
func (it *AdminPausablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminPausablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminPausablePaused represents a Paused event raised by the AdminPausable contract.
type AdminPausablePaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_AdminPausable *AdminPausableFilterer) FilterPaused(opts *bind.FilterOpts) (*AdminPausablePausedIterator, error) {

	logs, sub, err := _AdminPausable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AdminPausablePausedIterator{contract: _AdminPausable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_AdminPausable *AdminPausableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AdminPausablePaused) (event.Subscription, error) {

	logs, sub, err := _AdminPausable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminPausablePaused)
				if err := _AdminPausable.contract.UnpackLog(event, "Paused", log); err != nil {
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

// AdminPausableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the AdminPausable contract.
type AdminPausableUnpausedIterator struct {
	Event *AdminPausableUnpaused // Event containing the contract specifics and raw log

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
func (it *AdminPausableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminPausableUnpaused)
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
		it.Event = new(AdminPausableUnpaused)
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
func (it *AdminPausableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminPausableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminPausableUnpaused represents a Unpaused event raised by the AdminPausable contract.
type AdminPausableUnpaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_AdminPausable *AdminPausableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AdminPausableUnpausedIterator, error) {

	logs, sub, err := _AdminPausable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AdminPausableUnpausedIterator{contract: _AdminPausable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_AdminPausable *AdminPausableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AdminPausableUnpaused) (event.Subscription, error) {

	logs, sub, err := _AdminPausable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminPausableUnpaused)
				if err := _AdminPausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20Bin is the compiled bytecode used for deploying new contracts.
const IERC20Bin = `0x`

// DeployIERC20 deploys a new Ethereum contract, binding an instance of IERC20 to it.
func DeployIERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// IncognitoABI is the input ABI used to generate the binding from.
const IncognitoABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"name\":\"instructionApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IncognitoBin is the compiled bytecode used for deploying new contracts.
const IncognitoBin = `0x`

// DeployIncognito deploys a new Ethereum contract, binding an instance of Incognito to it.
func DeployIncognito(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Incognito, error) {
	parsed, err := abi.JSON(strings.NewReader(IncognitoABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IncognitoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Incognito{IncognitoCaller: IncognitoCaller{contract: contract}, IncognitoTransactor: IncognitoTransactor{contract: contract}, IncognitoFilterer: IncognitoFilterer{contract: contract}}, nil
}

// Incognito is an auto generated Go binding around an Ethereum contract.
type Incognito struct {
	IncognitoCaller     // Read-only binding to the contract
	IncognitoTransactor // Write-only binding to the contract
	IncognitoFilterer   // Log filterer for contract events
}

// IncognitoCaller is an auto generated read-only Go binding around an Ethereum contract.
type IncognitoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncognitoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IncognitoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncognitoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IncognitoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncognitoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IncognitoSession struct {
	Contract     *Incognito        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IncognitoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IncognitoCallerSession struct {
	Contract *IncognitoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IncognitoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IncognitoTransactorSession struct {
	Contract     *IncognitoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IncognitoRaw is an auto generated low-level Go binding around an Ethereum contract.
type IncognitoRaw struct {
	Contract *Incognito // Generic contract binding to access the raw methods on
}

// IncognitoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IncognitoCallerRaw struct {
	Contract *IncognitoCaller // Generic read-only contract binding to access the raw methods on
}

// IncognitoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IncognitoTransactorRaw struct {
	Contract *IncognitoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIncognito creates a new instance of Incognito, bound to a specific deployed contract.
func NewIncognito(address common.Address, backend bind.ContractBackend) (*Incognito, error) {
	contract, err := bindIncognito(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Incognito{IncognitoCaller: IncognitoCaller{contract: contract}, IncognitoTransactor: IncognitoTransactor{contract: contract}, IncognitoFilterer: IncognitoFilterer{contract: contract}}, nil
}

// NewIncognitoCaller creates a new read-only instance of Incognito, bound to a specific deployed contract.
func NewIncognitoCaller(address common.Address, caller bind.ContractCaller) (*IncognitoCaller, error) {
	contract, err := bindIncognito(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IncognitoCaller{contract: contract}, nil
}

// NewIncognitoTransactor creates a new write-only instance of Incognito, bound to a specific deployed contract.
func NewIncognitoTransactor(address common.Address, transactor bind.ContractTransactor) (*IncognitoTransactor, error) {
	contract, err := bindIncognito(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IncognitoTransactor{contract: contract}, nil
}

// NewIncognitoFilterer creates a new log filterer instance of Incognito, bound to a specific deployed contract.
func NewIncognitoFilterer(address common.Address, filterer bind.ContractFilterer) (*IncognitoFilterer, error) {
	contract, err := bindIncognito(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IncognitoFilterer{contract: contract}, nil
}

// bindIncognito binds a generic wrapper to an already deployed contract.
func bindIncognito(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IncognitoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Incognito *IncognitoRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Incognito.Contract.IncognitoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Incognito *IncognitoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Incognito.Contract.IncognitoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Incognito *IncognitoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Incognito.Contract.IncognitoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Incognito *IncognitoCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Incognito.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Incognito *IncognitoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Incognito.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Incognito *IncognitoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Incognito.Contract.contract.Transact(opts, method, params...)
}

// InstructionApproved is a free data retrieval call binding the contract method 0xf65d2116.
//
// Solidity: function instructionApproved(bool , bytes32 , uint256 , bytes32[] , bool[] , bytes32 , bytes32 , uint256[] , uint8[] , bytes32[] , bytes32[] ) constant returns(bool)
func (_Incognito *IncognitoCaller) InstructionApproved(opts *bind.CallOpts, arg0 bool, arg1 [32]byte, arg2 *big.Int, arg3 [][32]byte, arg4 []bool, arg5 [32]byte, arg6 [32]byte, arg7 []*big.Int, arg8 []uint8, arg9 [][32]byte, arg10 [][32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Incognito.contract.Call(opts, out, "instructionApproved", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
	return *ret0, err
}

// InstructionApproved is a free data retrieval call binding the contract method 0xf65d2116.
//
// Solidity: function instructionApproved(bool , bytes32 , uint256 , bytes32[] , bool[] , bytes32 , bytes32 , uint256[] , uint8[] , bytes32[] , bytes32[] ) constant returns(bool)
func (_Incognito *IncognitoSession) InstructionApproved(arg0 bool, arg1 [32]byte, arg2 *big.Int, arg3 [][32]byte, arg4 []bool, arg5 [32]byte, arg6 [32]byte, arg7 []*big.Int, arg8 []uint8, arg9 [][32]byte, arg10 [][32]byte) (bool, error) {
	return _Incognito.Contract.InstructionApproved(&_Incognito.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
}

// InstructionApproved is a free data retrieval call binding the contract method 0xf65d2116.
//
// Solidity: function instructionApproved(bool , bytes32 , uint256 , bytes32[] , bool[] , bytes32 , bytes32 , uint256[] , uint8[] , bytes32[] , bytes32[] ) constant returns(bool)
func (_Incognito *IncognitoCallerSession) InstructionApproved(arg0 bool, arg1 [32]byte, arg2 *big.Int, arg3 [][32]byte, arg4 []bool, arg5 [32]byte, arg6 [32]byte, arg7 []*big.Int, arg8 []uint8, arg9 [][32]byte, arg10 [][32]byte) (bool, error) {
	return _Incognito.Contract.InstructionApproved(&_Incognito.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
}

// VaultABI is the input ABI used to generate the binding from.
const VaultABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"incognitoProxyAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ndays\",\"type\":\"uint256\"}],\"name\":\"Extend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVault\",\"type\":\"address\"}],\"name\":\"Migrate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"}],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"}],\"name\":\"depositERC20\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"expire\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"extend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"incognito\",\"outputs\":[{\"internalType\":\"contractIncognito\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newVault\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"}],\"name\":\"parseBurnInst\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_successor\",\"type\":\"address\"}],\"name\":\"retire\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"successor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"uint256[2]\",\"name\":\"heights\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"instPaths\",\"type\":\"bytes32[][2]\"},{\"internalType\":\"bool[][2]\",\"name\":\"instPathIsLefts\",\"type\":\"bool[][2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"instRoots\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"blkData\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256[][2]\",\"name\":\"sigIdxs\",\"type\":\"uint256[][2]\"},{\"internalType\":\"uint8[][2]\",\"name\":\"sigVs\",\"type\":\"uint8[][2]\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"sigRs\",\"type\":\"bytes32[][2]\"},{\"internalType\":\"bytes32[][2]\",\"name\":\"sigSs\",\"type\":\"bytes32[][2]\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"withdrawed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VaultBin is the compiled bytecode used for deploying new contracts.
const VaultBin = `0x60806040523480156200001157600080fd5b50604051620020b5380380620020b583398101604081905262000034916200009d565b600080546001600160a01b03199081166001600160a01b03948516179091556001805460ff60a01b19169055426301e1338001600255600480548216929093169190911790915560058054909116905562000108565b80516200009781620000ee565b92915050565b60008060408385031215620000b157600080fd5b6000620000bf85856200008a565b9250506020620000d2858286016200008a565b9150509250929050565b60006001600160a01b03821662000097565b620000f981620000dc565b81146200010557600080fd5b50565b611f9d80620001186000396000f3fe6080604052600436106100fe5760003560e01c80638456cb5911610095578063a26e118611610064578063a26e118614610279578063ce5494bb1461028c578063dca40d9e146102ac578063f851a440146102cc578063f8b39745146102e1576100fe565b80638456cb59146102025780638a984538146102175780639714378c146102395780639e6371ba14610259576100fe565b80635c975abb116100d15780635c975abb146101625780636ff968c31461018d57806379599f96146101af5780637e16e6e1146101d1576100fe565b8063318d9e5d146101035780633f4ba83a146101255780634e71d92d1461013a5780635a67cb871461014f575b600080fd5b34801561010f57600080fd5b5061012361011e3660046115c4565b6102f4565b005b34801561013157600080fd5b50610123610562565b34801561014657600080fd5b506101236105fc565b61012361015d366004611560565b610698565b34801561016e57600080fd5b50610177610829565b6040516101849190611c66565b60405180910390f35b34801561019957600080fd5b506101a2610839565b6040516101849190611ba5565b3480156101bb57600080fd5b506101c4610848565b6040516101849190611daf565b3480156101dd57600080fd5b506101f16101ec366004611634565b61084e565b604051610184959493929190611dbd565b34801561020e57600080fd5b506101236108b0565b34801561022357600080fd5b5061022c610968565b6040516101849190611d41565b34801561024557600080fd5b50610123610254366004611616565b610977565b34801561026557600080fd5b5061012361027436600461153a565b610a20565b610123610287366004611634565b610a8d565b34801561029857600080fd5b506101236102a736600461153a565b610b07565b3480156102b857600080fd5b506101776102c7366004611616565b610ba5565b3480156102d857600080fd5b506101a2610bba565b6101236102ef366004611668565b610bc9565b6000546001600160a01b031633146103275760405162461bcd60e51b815260040161031e90611d9f565b60405180910390fd5b600154600160a01b900460ff166103505760405162461bcd60e51b815260040161031e90611d4f565b6005546001600160a01b031661036557600080fd5b60005b81518110156105275760006001600160a01b031682828151811061038857fe5b60200260200101516001600160a01b031614156103df576005546040516001600160a01b0390911690303180156108fc02916000818181858888f193505050501580156103d9573d6000803e3d6000fd5b5061051f565b60008282815181106103ed57fe5b60200260200101516001600160a01b03166370a08231306040518263ffffffff1660e01b81526004016104209190611ba5565b60206040518083038186803b15801561043857600080fd5b505afa15801561044c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061047091908101906117e3565b9050801561051d5782828151811061048457fe5b602090810291909101015160055460405163a9059cbb60e01b81526001600160a01b039283169263a9059cbb926104c2929116908590600401611be9565b602060405180830381600087803b1580156104dc57600080fd5b505af11580156104f0573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061051491908101906115f8565b61051d57600080fd5b505b600101610368565b507ff7262a15026689e5c048396ba0164e35c13a4ac7073569f46cfb8d1e22a83310816040516105579190611c55565b60405180910390a150565b6000546001600160a01b0316331461058c5760405162461bcd60e51b815260040161031e90611d9f565b600154600160a01b900460ff166105b55760405162461bcd60e51b815260040161031e90611d4f565b6001805460ff60a01b191690556040517f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa906105f2903390611bb3565b60405180910390a1565b600254421061061d5760405162461bcd60e51b815260040161031e90611d5f565b6001546001600160a01b031633146106475760405162461bcd60e51b815260040161031e90611d7f565b600154600080546001600160a01b0319166001600160a01b0392831617908190556040517f0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc926105f2921690611ba5565b600154600160a01b900460ff16156106c25760405162461bcd60e51b815260040161031e90611d8f565b6040516370a0823160e01b815283906000906001600160a01b038316906370a08231906106f3903090600401611ba5565b60206040518083038186803b15801561070b57600080fd5b505afa15801561071f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061074391908101906117e3565b9050670de0b6b3a7640000818501111561075c57600080fd5b6040516323b872dd60e01b81526001600160a01b038316906323b872dd9061078c90339030908990600401611bc1565b602060405180830381600087803b1580156107a657600080fd5b505af11580156107ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506107de91908101906115f8565b6107e757600080fd5b7f2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e85848660405161081a93929190611c26565b60405180910390a15050505050565b600154600160a01b900460ff1681565b6001546001600160a01b031681565b60025481565b6000806000806000808660008151811061086457fe5b602001015160f81c60f81b60f81c905060008760018151811061088357fe5b0160200151602289015160428a01516062909a0151939a60f89290921c9990985096509194509092505050565b6000546001600160a01b031633146108da5760405162461bcd60e51b815260040161031e90611d9f565b600154600160a01b900460ff16156109045760405162461bcd60e51b815260040161031e90611d8f565b60025442106109255760405162461bcd60e51b815260040161031e90611d5f565b6001805460ff60a01b1916600160a01b1790556040517f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258906105f2903390611bb3565b6004546001600160a01b031681565b6000546001600160a01b031633146109a15760405162461bcd60e51b815260040161031e90611d9f565b60025442106109c25760405162461bcd60e51b815260040161031e90611d5f565b61016e81106109e35760405162461bcd60e51b815260040161031e90611d6f565b600280546201518083020190556040517f02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e890610557908390611daf565b6000546001600160a01b03163314610a4a5760405162461bcd60e51b815260040161031e90611d9f565b6002544210610a6b5760405162461bcd60e51b815260040161031e90611d5f565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b600154600160a01b900460ff1615610ab75760405162461bcd60e51b815260040161031e90611d8f565b6b033b2e3c9fd0803ce8000000343031011115610ad357600080fd5b7f2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e6000823460405161055793929190611c26565b6000546001600160a01b03163314610b315760405162461bcd60e51b815260040161031e90611d9f565b600154600160a01b900460ff16610b5a5760405162461bcd60e51b815260040161031e90611d4f565b600580546001600160a01b0319166001600160a01b0383161790556040517fd58a618a39de682696ea37dd9a6bf9c793afa426fa1438e75c3966e3b541e45a90610557908390611bb3565b60036020526000908152604090205460ff1681565b6000546001600160a01b031681565b600154600160a01b900460ff1615610bf35760405162461bcd60e51b815260040161031e90611d8f565b6000806000806000610c048f61084e565b945094509450945094508460ff166048148015610c2457508360ff166001145b610c2d57600080fd5b6001600160a01b038316610c4e573031811115610c4957600080fd5b610cd7565b6040516370a0823160e01b815281906001600160a01b038516906370a0823190610c7c903090600401611ba5565b60206040518083038186803b158015610c9457600080fd5b505afa158015610ca8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610ccc91908101906117e3565b1015610cd757600080fd5b610ce98f8f8f8f8f8f8f8f8f8f610e08565b6001600160a01b038316610d33576040516001600160a01b0383169082156108fc029083906000818181858888f19350505050158015610d2d573d6000803e3d6000fd5b50610dbc565b60405163a9059cbb60e01b81526001600160a01b0384169063a9059cbb90610d619085908590600401611be9565b602060405180830381600087803b158015610d7b57600080fd5b505af1158015610d8f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610db391908101906115f8565b610dbc57600080fd5b7f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb838383604051610def93929190611c0b565b60405180910390a1505050505050505050505050505050565b895160208b012060008b8b826020020151604051602001610e2a929190611b83565b60408051601f198184030181529190528051602090910120905060008c8c60016020020151604051602001610e60929190611b83565b60408051601f1981840301815291815281516020928301206000868152600390935291205490915060ff1615610e9557600080fd5b6004546001600160a01b031663f65d21166001848f600060200201518f600060200201518f600060200201518f600060200201518f600060200201518f600060200201518f600060200201518f600060200201518f600060200201516040518c63ffffffff1660e01b8152600401610f179b9a99989796959493929190611c74565b60206040518083038186803b158015610f2f57600080fd5b505afa158015610f43573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610f6791908101906115f8565b610f7057600080fd5b600480546020808f01518e8201518e8301518e8401518e8501518e8601518e8701518e880151978e0151604051637b2e908b60e11b81526001600160a01b03909a169a63f65d21169a610fda9a60009a8f9a90999098909790969095909490939192909101611c74565b60206040518083038186803b158015610ff257600080fd5b505afa158015611006573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061102a91908101906115f8565b61103357600080fd5b50506000908152600360205260409020805460ff1916600117905550505050505050505050565b803561106581611f28565b92915050565b600082601f83011261107c57600080fd5b813561108f61108a82611e2f565b611e09565b915081818352602084019350602081019050838560208402820111156110b457600080fd5b60005b838110156110e057816110ca888261105a565b84525060209283019291909101906001016110b7565b5050505092915050565b600082601f8301126110fb57600080fd5b600261110961108a82611e4f565b9150818360005b838110156110e057813586016111268882611232565b8452506020928301929190910190600101611110565b600082601f83011261114d57600080fd5b600261115b61108a82611e4f565b9150818360005b838110156110e057813586016111788882611303565b8452506020928301929190910190600101611162565b600082601f83011261119f57600080fd5b60026111ad61108a82611e4f565b9150818360005b838110156110e057813586016111ca88826113d4565b84525060209283019291909101906001016111b4565b600082601f8301126111f157600080fd5b60026111ff61108a82611e4f565b9150818360005b838110156110e0578135860161121c8882611444565b8452506020928301929190910190600101611206565b600082601f83011261124357600080fd5b813561125161108a82611e2f565b9150818183526020840193506020810190508385602084028201111561127657600080fd5b60005b838110156110e0578161128c88826114b4565b8452506020928301929190910190600101611279565b600082601f8301126112b357600080fd5b60026112c161108a82611e4f565b915081838560208402820111156112d757600080fd5b60005b838110156110e057816112ed88826114ca565b84525060209283019291909101906001016112da565b600082601f83011261131457600080fd5b813561132261108a82611e2f565b9150818183526020840193506020810190508385602084028201111561134757600080fd5b60005b838110156110e0578161135d88826114ca565b845250602092830192919091019060010161134a565b600082601f83011261138457600080fd5b600261139261108a82611e4f565b915081838560208402820111156113a857600080fd5b60005b838110156110e057816113be88826114ca565b84525060209283019291909101906001016113ab565b600082601f8301126113e557600080fd5b81356113f361108a82611e2f565b9150818183526020840193506020810190508385602084028201111561141857600080fd5b60005b838110156110e0578161142e88826114ca565b845250602092830192919091019060010161141b565b600082601f83011261145557600080fd5b813561146361108a82611e2f565b9150818183526020840193506020810190508385602084028201111561148857600080fd5b60005b838110156110e0578161149e888261152f565b845250602092830192919091019060010161148b565b803561106581611f3f565b805161106581611f3f565b803561106581611f48565b600082601f8301126114e657600080fd5b81356114f461108a82611e6c565b9150808252602083016020830185838301111561151057600080fd5b61151b838284611ee2565b50505092915050565b805161106581611f48565b803561106581611f51565b60006020828403121561154c57600080fd5b6000611558848461105a565b949350505050565b60008060006060848603121561157557600080fd5b6000611581868661105a565b9350506020611592868287016114ca565b92505060408401356001600160401b038111156115ae57600080fd5b6115ba868287016114d5565b9150509250925092565b6000602082840312156115d657600080fd5b81356001600160401b038111156115ec57600080fd5b6115588482850161106b565b60006020828403121561160a57600080fd5b600061155884846114bf565b60006020828403121561162857600080fd5b600061155884846114ca565b60006020828403121561164657600080fd5b81356001600160401b0381111561165c57600080fd5b611558848285016114d5565b6000806000806000806000806000806101a08b8d03121561168857600080fd5b8a356001600160401b0381111561169e57600080fd5b6116aa8d828e016114d5565b9a505060206116bb8d828e01611373565b99505060608b01356001600160401b038111156116d757600080fd5b6116e38d828e0161113c565b98505060808b01356001600160401b038111156116ff57600080fd5b61170b8d828e016110ea565b97505060a061171c8d828e016112a2565b96505060e061172d8d828e016112a2565b9550506101208b01356001600160401b0381111561174a57600080fd5b6117568d828e0161118e565b9450506101408b01356001600160401b0381111561177357600080fd5b61177f8d828e016111e0565b9350506101608b01356001600160401b0381111561179c57600080fd5b6117a88d828e0161113c565b9250506101808b01356001600160401b038111156117c557600080fd5b6117d18d828e0161113c565b9150509295989b9194979a5092959850565b6000602082840312156117f557600080fd5b60006115588484611524565b600061180d8383611848565b505060200190565b600061180d83836119e2565b600061180d83836119eb565b600061180d8383611b7a565b61184281611ed0565b82525050565b61184281611eab565b600061185c82611e99565b6118668185611e9d565b935061187183611e93565b8060005b8381101561189f5781516118898882611801565b975061189483611e93565b925050600101611875565b509495945050505050565b60006118b582611e99565b6118bf8185611e9d565b93506118ca83611e93565b8060005b8381101561189f5781516118e28882611815565b97506118ed83611e93565b9250506001016118ce565b600061190382611e99565b61190d8185611e9d565b935061191883611e93565b8060005b8381101561189f5781516119308882611821565b975061193b83611e93565b92505060010161191c565b600061195182611e99565b61195b8185611e9d565b935061196683611e93565b8060005b8381101561189f57815161197e8882611821565b975061198983611e93565b92505060010161196a565b600061199f82611e99565b6119a98185611e9d565b93506119b483611e93565b8060005b8381101561189f5781516119cc888261182d565b97506119d783611e93565b9250506001016119b8565b61184281611eb6565b61184281611ebb565b60006119ff82611e99565b611a098185611ea6565b9350611a19818560208601611eee565b9290920192915050565b61184281611ed7565b6000611a3782611e99565b611a418185611e9d565b9350611a51818560208601611eee565b611a5a81611f1e565b9093019392505050565b6000611a71601483611e9d565b736e6f7420706175736564207269676874206e6f7760601b815260200192915050565b6000611aa1600783611e9d565b66195e1c1a5c995960ca1b815260200192915050565b6000611ac4601a83611e9d565b7f63616e6e6f7420657874656e6420666f7220746f6f206c6f6e67000000000000815260200192915050565b6000611afd600c83611e9d565b6b1d5b985d5d1a1bdc9a5e995960a21b815260200192915050565b6000611b25601083611e9d565b6f706175736564207269676874206e6f7760801b815260200192915050565b6000611b51600983611e9d565b683737ba1030b236b4b760b91b815260200192915050565b611842611b7582611ebb565b611ebb565b61184281611eca565b6000611b8f82856119f4565b9150611b9b8284611b69565b5060200192915050565b602081016110658284611848565b602081016110658284611839565b60608101611bcf8286611839565b611bdc6020830185611848565b61155860408301846119eb565b60408101611bf78285611839565b611c0460208301846119eb565b9392505050565b60608101611c198286611848565b611bdc6020830185611839565b60608101611c348286611848565b8181036020830152611c468185611a2c565b905061155860408301846119eb565b60208082528101611c048184611851565b6020810161106582846119e2565b6101608101611c83828e6119e2565b611c90602083018d6119eb565b611c9d604083018c6119eb565b8181036060830152611caf818b6118f8565b90508181036080830152611cc3818a6118aa565b9050611cd260a08301896119eb565b611cdf60c08301886119eb565b81810360e0830152611cf18187611946565b9050818103610100830152611d068186611994565b9050818103610120830152611d1b81856118f8565b9050818103610140830152611d3081846118f8565b9d9c50505050505050505050505050565b602081016110658284611a23565b6020808252810161106581611a64565b6020808252810161106581611a94565b6020808252810161106581611ab7565b6020808252810161106581611af0565b6020808252810161106581611b18565b6020808252810161106581611b44565b6020810161106582846119eb565b60a08101611dcb8288611b7a565b611dd86020830187611b7a565b611de56040830186611848565b611df26060830185611848565b611dff60808301846119eb565b9695505050505050565b6040518181016001600160401b0381118282101715611e2757600080fd5b604052919050565b60006001600160401b03821115611e4557600080fd5b5060209081020190565b60006001600160401b03821115611e6557600080fd5b5060200290565b60006001600160401b03821115611e8257600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b919050565b600061106582611ebe565b151590565b90565b6001600160a01b031690565b60ff1690565b6000611065825b600061106582611eab565b82818337506000910152565b60005b83811015611f09578181015183820152602001611ef1565b83811115611f18576000848401525b50505050565b601f01601f191690565b611f3181611eab565b8114611f3c57600080fd5b50565b611f3181611eb6565b611f3181611ebb565b611f3181611eca56fea365627a7a723158202b6742b616255c61630d8a7fc4f5bebc6eef35a2c00bd315884aef39d89b67666c6578706572696d656e74616cf564736f6c634300050c0040`

// DeployVault deploys a new Ethereum contract, binding an instance of Vault to it.
func DeployVault(auth *bind.TransactOpts, backend bind.ContractBackend, admin common.Address, incognitoProxyAddress common.Address) (common.Address, *types.Transaction, *Vault, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VaultBin), backend, admin, incognitoProxyAddress)
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

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Vault *VaultCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Vault *VaultSession) Admin() (common.Address, error) {
	return _Vault.Contract.Admin(&_Vault.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Vault *VaultCallerSession) Admin() (common.Address, error) {
	return _Vault.Contract.Admin(&_Vault.CallOpts)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() constant returns(uint256)
func (_Vault *VaultCaller) Expire(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "expire")
	return *ret0, err
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() constant returns(uint256)
func (_Vault *VaultSession) Expire() (*big.Int, error) {
	return _Vault.Contract.Expire(&_Vault.CallOpts)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() constant returns(uint256)
func (_Vault *VaultCallerSession) Expire() (*big.Int, error) {
	return _Vault.Contract.Expire(&_Vault.CallOpts)
}

// Incognito is a free data retrieval call binding the contract method 0x8a984538.
//
// Solidity: function incognito() constant returns(address)
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
// Solidity: function incognito() constant returns(address)
func (_Vault *VaultSession) Incognito() (common.Address, error) {
	return _Vault.Contract.Incognito(&_Vault.CallOpts)
}

// Incognito is a free data retrieval call binding the contract method 0x8a984538.
//
// Solidity: function incognito() constant returns(address)
func (_Vault *VaultCallerSession) Incognito() (common.Address, error) {
	return _Vault.Contract.Incognito(&_Vault.CallOpts)
}

// ParseBurnInst is a free data retrieval call binding the contract method 0x7e16e6e1.
//
// Solidity: function parseBurnInst(bytes inst) constant returns(uint8, uint8, address, address, uint256)
func (_Vault *VaultCaller) ParseBurnInst(opts *bind.CallOpts, inst []byte) (uint8, uint8, common.Address, common.Address, *big.Int, error) {
	var (
		ret0 = new(uint8)
		ret1 = new(uint8)
		ret2 = new(common.Address)
		ret3 = new(common.Address)
		ret4 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Vault.contract.Call(opts, out, "parseBurnInst", inst)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// ParseBurnInst is a free data retrieval call binding the contract method 0x7e16e6e1.
//
// Solidity: function parseBurnInst(bytes inst) constant returns(uint8, uint8, address, address, uint256)
func (_Vault *VaultSession) ParseBurnInst(inst []byte) (uint8, uint8, common.Address, common.Address, *big.Int, error) {
	return _Vault.Contract.ParseBurnInst(&_Vault.CallOpts, inst)
}

// ParseBurnInst is a free data retrieval call binding the contract method 0x7e16e6e1.
//
// Solidity: function parseBurnInst(bytes inst) constant returns(uint8, uint8, address, address, uint256)
func (_Vault *VaultCallerSession) ParseBurnInst(inst []byte) (uint8, uint8, common.Address, common.Address, *big.Int, error) {
	return _Vault.Contract.ParseBurnInst(&_Vault.CallOpts, inst)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Vault *VaultCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Vault *VaultSession) Paused() (bool, error) {
	return _Vault.Contract.Paused(&_Vault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Vault *VaultCallerSession) Paused() (bool, error) {
	return _Vault.Contract.Paused(&_Vault.CallOpts)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() constant returns(address)
func (_Vault *VaultCaller) Successor(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "successor")
	return *ret0, err
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() constant returns(address)
func (_Vault *VaultSession) Successor() (common.Address, error) {
	return _Vault.Contract.Successor(&_Vault.CallOpts)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() constant returns(address)
func (_Vault *VaultCallerSession) Successor() (common.Address, error) {
	return _Vault.Contract.Successor(&_Vault.CallOpts)
}

// Withdrawed is a free data retrieval call binding the contract method 0xdca40d9e.
//
// Solidity: function withdrawed(bytes32 ) constant returns(bool)
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
// Solidity: function withdrawed(bytes32 ) constant returns(bool)
func (_Vault *VaultSession) Withdrawed(arg0 [32]byte) (bool, error) {
	return _Vault.Contract.Withdrawed(&_Vault.CallOpts, arg0)
}

// Withdrawed is a free data retrieval call binding the contract method 0xdca40d9e.
//
// Solidity: function withdrawed(bytes32 ) constant returns(bool)
func (_Vault *VaultCallerSession) Withdrawed(arg0 [32]byte) (bool, error) {
	return _Vault.Contract.Withdrawed(&_Vault.CallOpts, arg0)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Vault *VaultTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Vault *VaultSession) Claim() (*types.Transaction, error) {
	return _Vault.Contract.Claim(&_Vault.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Vault *VaultTransactorSession) Claim() (*types.Transaction, error) {
	return _Vault.Contract.Claim(&_Vault.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xa26e1186.
//
// Solidity: function deposit(string incognitoAddress) returns()
func (_Vault *VaultTransactor) Deposit(opts *bind.TransactOpts, incognitoAddress string) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "deposit", incognitoAddress)
}

// Deposit is a paid mutator transaction binding the contract method 0xa26e1186.
//
// Solidity: function deposit(string incognitoAddress) returns()
func (_Vault *VaultSession) Deposit(incognitoAddress string) (*types.Transaction, error) {
	return _Vault.Contract.Deposit(&_Vault.TransactOpts, incognitoAddress)
}

// Deposit is a paid mutator transaction binding the contract method 0xa26e1186.
//
// Solidity: function deposit(string incognitoAddress) returns()
func (_Vault *VaultTransactorSession) Deposit(incognitoAddress string) (*types.Transaction, error) {
	return _Vault.Contract.Deposit(&_Vault.TransactOpts, incognitoAddress)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x5a67cb87.
//
// Solidity: function depositERC20(address token, uint256 amount, string incognitoAddress) returns()
func (_Vault *VaultTransactor) DepositERC20(opts *bind.TransactOpts, token common.Address, amount *big.Int, incognitoAddress string) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "depositERC20", token, amount, incognitoAddress)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x5a67cb87.
//
// Solidity: function depositERC20(address token, uint256 amount, string incognitoAddress) returns()
func (_Vault *VaultSession) DepositERC20(token common.Address, amount *big.Int, incognitoAddress string) (*types.Transaction, error) {
	return _Vault.Contract.DepositERC20(&_Vault.TransactOpts, token, amount, incognitoAddress)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x5a67cb87.
//
// Solidity: function depositERC20(address token, uint256 amount, string incognitoAddress) returns()
func (_Vault *VaultTransactorSession) DepositERC20(token common.Address, amount *big.Int, incognitoAddress string) (*types.Transaction, error) {
	return _Vault.Contract.DepositERC20(&_Vault.TransactOpts, token, amount, incognitoAddress)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_Vault *VaultTransactor) Extend(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "extend", n)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_Vault *VaultSession) Extend(n *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Extend(&_Vault.TransactOpts, n)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_Vault *VaultTransactorSession) Extend(n *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Extend(&_Vault.TransactOpts, n)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _newVault) returns()
func (_Vault *VaultTransactor) Migrate(opts *bind.TransactOpts, _newVault common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "migrate", _newVault)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _newVault) returns()
func (_Vault *VaultSession) Migrate(_newVault common.Address) (*types.Transaction, error) {
	return _Vault.Contract.Migrate(&_Vault.TransactOpts, _newVault)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _newVault) returns()
func (_Vault *VaultTransactorSession) Migrate(_newVault common.Address) (*types.Transaction, error) {
	return _Vault.Contract.Migrate(&_Vault.TransactOpts, _newVault)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Vault *VaultTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Vault *VaultSession) Pause() (*types.Transaction, error) {
	return _Vault.Contract.Pause(&_Vault.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Vault *VaultTransactorSession) Pause() (*types.Transaction, error) {
	return _Vault.Contract.Pause(&_Vault.TransactOpts)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_Vault *VaultTransactor) Retire(opts *bind.TransactOpts, _successor common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "retire", _successor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_Vault *VaultSession) Retire(_successor common.Address) (*types.Transaction, error) {
	return _Vault.Contract.Retire(&_Vault.TransactOpts, _successor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_Vault *VaultTransactorSession) Retire(_successor common.Address) (*types.Transaction, error) {
	return _Vault.Contract.Retire(&_Vault.TransactOpts, _successor)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Vault *VaultTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Vault *VaultSession) Unpause() (*types.Transaction, error) {
	return _Vault.Contract.Unpause(&_Vault.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Vault *VaultTransactorSession) Unpause() (*types.Transaction, error) {
	return _Vault.Contract.Unpause(&_Vault.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf8b39745.
//
// Solidity: function withdraw(bytes inst, uint256[2] heights, bytes32[][2] instPaths, bool[][2] instPathIsLefts, bytes32[2] instRoots, bytes32[2] blkData, uint256[][2] sigIdxs, uint8[][2] sigVs, bytes32[][2] sigRs, bytes32[][2] sigSs) returns()
func (_Vault *VaultTransactor) Withdraw(opts *bind.TransactOpts, inst []byte, heights [2]*big.Int, instPaths [2][][32]byte, instPathIsLefts [2][]bool, instRoots [2][32]byte, blkData [2][32]byte, sigIdxs [2][]*big.Int, sigVs [2][]uint8, sigRs [2][][32]byte, sigSs [2][][32]byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "withdraw", inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf8b39745.
//
// Solidity: function withdraw(bytes inst, uint256[2] heights, bytes32[][2] instPaths, bool[][2] instPathIsLefts, bytes32[2] instRoots, bytes32[2] blkData, uint256[][2] sigIdxs, uint8[][2] sigVs, bytes32[][2] sigRs, bytes32[][2] sigSs) returns()
func (_Vault *VaultSession) Withdraw(inst []byte, heights [2]*big.Int, instPaths [2][][32]byte, instPathIsLefts [2][]bool, instRoots [2][32]byte, blkData [2][32]byte, sigIdxs [2][]*big.Int, sigVs [2][]uint8, sigRs [2][][32]byte, sigSs [2][][32]byte) (*types.Transaction, error) {
	return _Vault.Contract.Withdraw(&_Vault.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf8b39745.
//
// Solidity: function withdraw(bytes inst, uint256[2] heights, bytes32[][2] instPaths, bool[][2] instPathIsLefts, bytes32[2] instRoots, bytes32[2] blkData, uint256[][2] sigIdxs, uint8[][2] sigVs, bytes32[][2] sigRs, bytes32[][2] sigSs) returns()
func (_Vault *VaultTransactorSession) Withdraw(inst []byte, heights [2]*big.Int, instPaths [2][][32]byte, instPathIsLefts [2][]bool, instRoots [2][32]byte, blkData [2][32]byte, sigIdxs [2][]*big.Int, sigVs [2][]uint8, sigRs [2][][32]byte, sigSs [2][][32]byte) (*types.Transaction, error) {
	return _Vault.Contract.Withdraw(&_Vault.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// VaultClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the Vault contract.
type VaultClaimIterator struct {
	Event *VaultClaim // Event containing the contract specifics and raw log

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
func (it *VaultClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultClaim)
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
		it.Event = new(VaultClaim)
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
func (it *VaultClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultClaim represents a Claim event raised by the Vault contract.
type VaultClaim struct {
	Claimer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_Vault *VaultFilterer) FilterClaim(opts *bind.FilterOpts) (*VaultClaimIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &VaultClaimIterator{contract: _Vault.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_Vault *VaultFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *VaultClaim) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultClaim)
				if err := _Vault.contract.UnpackLog(event, "Claim", log); err != nil {
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
// Solidity: event Deposit(address token, string incognitoAddress, uint256 amount)
func (_Vault *VaultFilterer) FilterDeposit(opts *bind.FilterOpts) (*VaultDepositIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &VaultDepositIterator{contract: _Vault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e.
//
// Solidity: event Deposit(address token, string incognitoAddress, uint256 amount)
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

// VaultExtendIterator is returned from FilterExtend and is used to iterate over the raw logs and unpacked data for Extend events raised by the Vault contract.
type VaultExtendIterator struct {
	Event *VaultExtend // Event containing the contract specifics and raw log

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
func (it *VaultExtendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultExtend)
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
		it.Event = new(VaultExtend)
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
func (it *VaultExtendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultExtendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultExtend represents a Extend event raised by the Vault contract.
type VaultExtend struct {
	Ndays *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterExtend is a free log retrieval operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_Vault *VaultFilterer) FilterExtend(opts *bind.FilterOpts) (*VaultExtendIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return &VaultExtendIterator{contract: _Vault.contract, event: "Extend", logs: logs, sub: sub}, nil
}

// WatchExtend is a free log subscription operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_Vault *VaultFilterer) WatchExtend(opts *bind.WatchOpts, sink chan<- *VaultExtend) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultExtend)
				if err := _Vault.contract.UnpackLog(event, "Extend", log); err != nil {
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

// VaultMigrateIterator is returned from FilterMigrate and is used to iterate over the raw logs and unpacked data for Migrate events raised by the Vault contract.
type VaultMigrateIterator struct {
	Event *VaultMigrate // Event containing the contract specifics and raw log

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
func (it *VaultMigrateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultMigrate)
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
		it.Event = new(VaultMigrate)
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
func (it *VaultMigrateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultMigrateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultMigrate represents a Migrate event raised by the Vault contract.
type VaultMigrate struct {
	NewVault common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMigrate is a free log retrieval operation binding the contract event 0xd58a618a39de682696ea37dd9a6bf9c793afa426fa1438e75c3966e3b541e45a.
//
// Solidity: event Migrate(address newVault)
func (_Vault *VaultFilterer) FilterMigrate(opts *bind.FilterOpts) (*VaultMigrateIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Migrate")
	if err != nil {
		return nil, err
	}
	return &VaultMigrateIterator{contract: _Vault.contract, event: "Migrate", logs: logs, sub: sub}, nil
}

// WatchMigrate is a free log subscription operation binding the contract event 0xd58a618a39de682696ea37dd9a6bf9c793afa426fa1438e75c3966e3b541e45a.
//
// Solidity: event Migrate(address newVault)
func (_Vault *VaultFilterer) WatchMigrate(opts *bind.WatchOpts, sink chan<- *VaultMigrate) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Migrate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultMigrate)
				if err := _Vault.contract.UnpackLog(event, "Migrate", log); err != nil {
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

// VaultPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Vault contract.
type VaultPausedIterator struct {
	Event *VaultPaused // Event containing the contract specifics and raw log

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
func (it *VaultPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultPaused)
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
		it.Event = new(VaultPaused)
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
func (it *VaultPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultPaused represents a Paused event raised by the Vault contract.
type VaultPaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_Vault *VaultFilterer) FilterPaused(opts *bind.FilterOpts) (*VaultPausedIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &VaultPausedIterator{contract: _Vault.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_Vault *VaultFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *VaultPaused) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultPaused)
				if err := _Vault.contract.UnpackLog(event, "Paused", log); err != nil {
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

// VaultUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Vault contract.
type VaultUnpausedIterator struct {
	Event *VaultUnpaused // Event containing the contract specifics and raw log

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
func (it *VaultUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultUnpaused)
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
		it.Event = new(VaultUnpaused)
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
func (it *VaultUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultUnpaused represents a Unpaused event raised by the Vault contract.
type VaultUnpaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_Vault *VaultFilterer) FilterUnpaused(opts *bind.FilterOpts) (*VaultUnpausedIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &VaultUnpausedIterator{contract: _Vault.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_Vault *VaultFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *VaultUnpaused) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultUnpaused)
				if err := _Vault.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
// Solidity: event Withdraw(address token, address to, uint256 amount)
func (_Vault *VaultFilterer) FilterWithdraw(opts *bind.FilterOpts) (*VaultWithdrawIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &VaultWithdrawIterator{contract: _Vault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address token, address to, uint256 amount)
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
