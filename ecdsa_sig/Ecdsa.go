// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ecdsa_sig

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

// ECDSAABI is the input ABI used to generate the binding from.
const ECDSAABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"int128\",\"name\":\"numSig\",\"type\":\"int128\"},{\"internalType\":\"address[10]\",\"name\":\"committee\",\"type\":\"address[10]\"},{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[10]\",\"name\":\"v\",\"type\":\"uint256[10]\"},{\"internalType\":\"bytes32[10]\",\"name\":\"r\",\"type\":\"bytes32[10]\"},{\"internalType\":\"bytes32[10]\",\"name\":\"s\",\"type\":\"bytes32[10]\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ECDSAFuncSigs maps the 4-byte function signature to its string representation.
var ECDSAFuncSigs = map[string]string{
	"b7ae9d2e": "verify(int128,address[10],bytes32,uint256[10],bytes32[10],bytes32[10])",
}

// ECDSABin is the compiled bytecode used for deploying new contracts.
var ECDSABin = "0x608060405234801561001057600080fd5b50610258806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063b7ae9d2e14610030575b600080fd5b61011f600480360361054081101561004757600080fd5b604080516101408181019092528335600f0b939283019291610160830191906020840190600a90839083908082843760009201919091525050604080516101408181019092529295843595909490936101608201935091602090910190600a90839083908082843760009201919091525050604080516101408181019092529295949381810193925090600a90839083908082843760009201919091525050604080516101408181019092529295949381810193925090600a9083908390808284376000920191909152509194506101339350505050565b604080519115158252519081900360200190f35b6000805b87600f0b811015610213578681600a811061014e57fe5b60200201516001600160a01b03166001878784600a811061016b57fe5b60200201518785600a811061017c57fe5b60200201518786600a811061018d57fe5b602002015160405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa1580156101e9573d6000803e3d6000fd5b505050602060405103516001600160a01b03161461020b576000915050610219565b600101610137565b50600190505b969550505050505056fea265627a7a7231582007e3a58ada9cd05d246faae492b16a63303a0047bdfffcd22306e4857527d45a64736f6c634300050b0032"

// DeployECDSA deploys a new Ethereum contract, binding an instance of ECDSA to it.
func DeployECDSA(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECDSA, error) {
	parsed, err := abi.JSON(strings.NewReader(ECDSAABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ECDSABin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// ECDSA is an auto generated Go binding around an Ethereum contract.
type ECDSA struct {
	ECDSACaller     // Read-only binding to the contract
	ECDSATransactor // Write-only binding to the contract
	ECDSAFilterer   // Log filterer for contract events
}

// ECDSACaller is an auto generated read-only Go binding around an Ethereum contract.
type ECDSACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSATransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECDSATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECDSAFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECDSASession struct {
	Contract     *ECDSA            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECDSACallerSession struct {
	Contract *ECDSACaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ECDSATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECDSATransactorSession struct {
	Contract     *ECDSATransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSARaw is an auto generated low-level Go binding around an Ethereum contract.
type ECDSARaw struct {
	Contract *ECDSA // Generic contract binding to access the raw methods on
}

// ECDSACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECDSACallerRaw struct {
	Contract *ECDSACaller // Generic read-only contract binding to access the raw methods on
}

// ECDSATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECDSATransactorRaw struct {
	Contract *ECDSATransactor // Generic write-only contract binding to access the raw methods on
}

// NewECDSA creates a new instance of ECDSA, bound to a specific deployed contract.
func NewECDSA(address common.Address, backend bind.ContractBackend) (*ECDSA, error) {
	contract, err := bindECDSA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// NewECDSACaller creates a new read-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSACaller(address common.Address, caller bind.ContractCaller) (*ECDSACaller, error) {
	contract, err := bindECDSA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSACaller{contract: contract}, nil
}

// NewECDSATransactor creates a new write-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSATransactor(address common.Address, transactor bind.ContractTransactor) (*ECDSATransactor, error) {
	contract, err := bindECDSA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSATransactor{contract: contract}, nil
}

// NewECDSAFilterer creates a new log filterer instance of ECDSA, bound to a specific deployed contract.
func NewECDSAFilterer(address common.Address, filterer bind.ContractFilterer) (*ECDSAFilterer, error) {
	contract, err := bindECDSA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECDSAFilterer{contract: contract}, nil
}

// bindECDSA binds a generic wrapper to an already deployed contract.
func bindECDSA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECDSAABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSARaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.ECDSACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSACallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transact(opts, method, params...)
}

// Verify is a free data retrieval call binding the contract method 0xb7ae9d2e.
//
// Solidity: function verify(int128 numSig, address[10] committee, bytes32 msgHash, uint256[10] v, bytes32[10] r, bytes32[10] s) constant returns(bool)
func (_ECDSA *ECDSACaller) Verify(opts *bind.CallOpts, numSig *big.Int, committee [10]common.Address, msgHash [32]byte, v [10]*big.Int, r [10][32]byte, s [10][32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ECDSA.contract.Call(opts, out, "verify", numSig, committee, msgHash, v, r, s)
	return *ret0, err
}

// Verify is a free data retrieval call binding the contract method 0xb7ae9d2e.
//
// Solidity: function verify(int128 numSig, address[10] committee, bytes32 msgHash, uint256[10] v, bytes32[10] r, bytes32[10] s) constant returns(bool)
func (_ECDSA *ECDSASession) Verify(numSig *big.Int, committee [10]common.Address, msgHash [32]byte, v [10]*big.Int, r [10][32]byte, s [10][32]byte) (bool, error) {
	return _ECDSA.Contract.Verify(&_ECDSA.CallOpts, numSig, committee, msgHash, v, r, s)
}

// Verify is a free data retrieval call binding the contract method 0xb7ae9d2e.
//
// Solidity: function verify(int128 numSig, address[10] committee, bytes32 msgHash, uint256[10] v, bytes32[10] r, bytes32[10] s) constant returns(bool)
func (_ECDSA *ECDSACallerSession) Verify(numSig *big.Int, committee [10]common.Address, msgHash [32]byte, v [10]*big.Int, r [10][32]byte, s [10][32]byte) (bool, error) {
	return _ECDSA.Contract.Verify(&_ECDSA.CallOpts, numSig, committee, msgHash, v, r, s)
}
