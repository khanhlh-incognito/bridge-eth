// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zrxtrade

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

// ZrxtradeABI is the input ABI used to generate the binding from.
const ZrxtradeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wETH\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zeroProxy\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_incognitoSmartContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"ETH_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"incognitoSmartContract\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callDataHex\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_forwarder\",\"type\":\"address\"}],\"name\":\"trade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawWrapETH\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ZrxtradeBin is the compiled bytecode used for deploying new contracts.
const ZrxtradeBin = `0x608060405234801561001057600080fd5b50604051610a78380380610a7883398101604081905261002f91610085565b60018054600160a060020a03938416600160a060020a0319918216179091556000805492841692821692909217909155600280549390921692169190911790556100fa565b805161007f816100e3565b92915050565b60008060006060848603121561009a57600080fd5b60006100a68686610074565b93505060206100b786828701610074565b92505060406100c886828701610074565b9150509250925092565b6000600160a060020a03821661007f565b6100ec816100d2565b81146100f757600080fd5b50565b61096f806101096000396000f3fe60806040526004361061005b577c010000000000000000000000000000000000000000000000000000000060003504630e0569a6811461005d57806372e94bf61461007d57806398c36152146100a8578063b42a644b146100c9575b005b34801561006957600080fd5b5061005b6100783660046106fe565b6100eb565b34801561008957600080fd5b50610092610169565b60405161009f919061080f565b60405180910390f35b6100bb6100b6366004610671565b61016e565b60405161009f929190610801565b3480156100d557600080fd5b506100de61026a565b60405161009f91906107bd565b6002546040517f2e1a7d4d000000000000000000000000000000000000000000000000000000008152600160a060020a0390911690632e1a7d4d9061013490849060040161081d565b600060405180830381600087803b15801561014e57600080fd5b505af1158015610162573d6000803e3d6000fd5b5050505050565b600081565b600080548190600160a060020a0316331461018857600080fd5b6001546101a0908890600160a060020a031688610279565b6000839050600081600160a060020a031634876040516101c0919061079c565b60006040518083038185875af1925050503d80600081146101fd576040519150601f19603f3d011682016040523d82523d6000602084013e610202565b606091505b505090508061021057600080fd5b6000600160a060020a0388166102455761022a60006103d2565b9050610235816100eb565b61024060008261040f565b61025a565b61024e886103d2565b905061025a888261040f565b9699969850959650505050505050565b600054600160a060020a031681565b600160a060020a038316156103cd576040517f095ea7b3000000000000000000000000000000000000000000000000000000008152600160a060020a0384169063095ea7b3906102d09085906000906004016107e6565b602060405180830381600087803b1580156102ea57600080fd5b505af11580156102fe573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610322919081019061064b565b61032b57600080fd5b6040517f095ea7b3000000000000000000000000000000000000000000000000000000008152600160a060020a0384169063095ea7b3906103729085908590600401610801565b602060405180830381600087803b15801561038c57600080fd5b505af11580156103a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506103c4919081019061064b565b6103cd57600080fd5b505050565b6000600160a060020a0382166103fe576002546103f790600160a060020a031661050d565b905061040a565b6104078261050d565b90505b919050565b600160a060020a03821661046a57303181111561042b57600080fd5b60008054604051600160a060020a039091169183156108fc02918491818181858888f19350505050158015610464573d6000803e3d6000fd5b50610509565b6000546040517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a038481169263a9059cbb926104b7929091169085906004016107cb565b602060405180830381600087803b1580156104d157600080fd5b505af11580156104e5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506103cd919081019061064b565b5050565b6000600160a060020a0382166105255750303161040a565b6040517f70a08231000000000000000000000000000000000000000000000000000000008152600160a060020a038316906370a082319061056a9030906004016107af565b60206040518083038186803b15801561058257600080fd5b505afa158015610596573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610407919081019061071c565b80356105c5816108fa565b92915050565b80516105c581610911565b600082601f8301126105e757600080fd5b81356105fa6105f582610852565b61082b565b9150808252602083016020830185838301111561061657600080fd5b6106218382846108be565b50505092915050565b80356105c58161091a565b80356105c581610923565b80516105c581610923565b60006020828403121561065d57600080fd5b600061066984846105cb565b949350505050565b600080600080600060a0868803121561068957600080fd5b6000610695888861062a565b95505060206106a688828901610635565b94505060406106b78882890161062a565b935050606086013567ffffffffffffffff8111156106d457600080fd5b6106e0888289016105d6565b92505060806106f1888289016105ba565b9150509295509295909350565b60006020828403121561071057600080fd5b60006106698484610635565b60006020828403121561072e57600080fd5b60006106698484610640565b610743816108a8565b82525050565b6107438161087e565b600061075d8261087a565b610767818561040a565b93506107778185602086016108ca565b9290920192915050565b6107438161088e565b610743816108b3565b610743816108a5565b60006107a88284610752565b9392505050565b602081016105c5828461073a565b602081016105c58284610749565b604081016107d9828561073a565b6107a86020830184610793565b604081016107f48285610749565b6107a8602083018461078a565b604081016107d98285610749565b602081016105c58284610781565b602081016105c58284610793565b60405181810167ffffffffffffffff8111828210171561084a57600080fd5b604052919050565b600067ffffffffffffffff82111561086957600080fd5b506020601f91909101601f19160190565b5190565b600061040782610899565b151590565b60006104078261087e565b600160a060020a031690565b90565b60006104078261088e565b6000610407826108a5565b82818337506000910152565b60005b838110156108e55781810151838201526020016108cd565b838111156108f4576000848401525b50505050565b6109038161087e565b811461090e57600080fd5b50565b61090381610889565b6109038161088e565b610903816108a556fea365627a7a723158204ed2ec7169bf9c66727abef9a628e54d94b5d1aad0bda19f3e5d684ac199bf406c6578706572696d656e74616cf564736f6c634300050c0040`

// DeployZrxtrade deploys a new Ethereum contract, binding an instance of Zrxtrade to it.
func DeployZrxtrade(auth *bind.TransactOpts, backend bind.ContractBackend, _wETH common.Address, _zeroProxy common.Address, _incognitoSmartContract common.Address) (common.Address, *types.Transaction, *Zrxtrade, error) {
	parsed, err := abi.JSON(strings.NewReader(ZrxtradeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZrxtradeBin), backend, _wETH, _zeroProxy, _incognitoSmartContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Zrxtrade{ZrxtradeCaller: ZrxtradeCaller{contract: contract}, ZrxtradeTransactor: ZrxtradeTransactor{contract: contract}, ZrxtradeFilterer: ZrxtradeFilterer{contract: contract}}, nil
}

// Zrxtrade is an auto generated Go binding around an Ethereum contract.
type Zrxtrade struct {
	ZrxtradeCaller     // Read-only binding to the contract
	ZrxtradeTransactor // Write-only binding to the contract
	ZrxtradeFilterer   // Log filterer for contract events
}

// ZrxtradeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZrxtradeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZrxtradeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZrxtradeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZrxtradeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZrxtradeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZrxtradeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZrxtradeSession struct {
	Contract     *Zrxtrade         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZrxtradeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZrxtradeCallerSession struct {
	Contract *ZrxtradeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ZrxtradeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZrxtradeTransactorSession struct {
	Contract     *ZrxtradeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ZrxtradeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZrxtradeRaw struct {
	Contract *Zrxtrade // Generic contract binding to access the raw methods on
}

// ZrxtradeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZrxtradeCallerRaw struct {
	Contract *ZrxtradeCaller // Generic read-only contract binding to access the raw methods on
}

// ZrxtradeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZrxtradeTransactorRaw struct {
	Contract *ZrxtradeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZrxtrade creates a new instance of Zrxtrade, bound to a specific deployed contract.
func NewZrxtrade(address common.Address, backend bind.ContractBackend) (*Zrxtrade, error) {
	contract, err := bindZrxtrade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Zrxtrade{ZrxtradeCaller: ZrxtradeCaller{contract: contract}, ZrxtradeTransactor: ZrxtradeTransactor{contract: contract}, ZrxtradeFilterer: ZrxtradeFilterer{contract: contract}}, nil
}

// NewZrxtradeCaller creates a new read-only instance of Zrxtrade, bound to a specific deployed contract.
func NewZrxtradeCaller(address common.Address, caller bind.ContractCaller) (*ZrxtradeCaller, error) {
	contract, err := bindZrxtrade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZrxtradeCaller{contract: contract}, nil
}

// NewZrxtradeTransactor creates a new write-only instance of Zrxtrade, bound to a specific deployed contract.
func NewZrxtradeTransactor(address common.Address, transactor bind.ContractTransactor) (*ZrxtradeTransactor, error) {
	contract, err := bindZrxtrade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZrxtradeTransactor{contract: contract}, nil
}

// NewZrxtradeFilterer creates a new log filterer instance of Zrxtrade, bound to a specific deployed contract.
func NewZrxtradeFilterer(address common.Address, filterer bind.ContractFilterer) (*ZrxtradeFilterer, error) {
	contract, err := bindZrxtrade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZrxtradeFilterer{contract: contract}, nil
}

// bindZrxtrade binds a generic wrapper to an already deployed contract.
func bindZrxtrade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZrxtradeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zrxtrade *ZrxtradeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Zrxtrade.Contract.ZrxtradeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zrxtrade *ZrxtradeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zrxtrade.Contract.ZrxtradeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zrxtrade *ZrxtradeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zrxtrade.Contract.ZrxtradeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zrxtrade *ZrxtradeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Zrxtrade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zrxtrade *ZrxtradeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zrxtrade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zrxtrade *ZrxtradeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zrxtrade.Contract.contract.Transact(opts, method, params...)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_Zrxtrade *ZrxtradeCaller) ETHCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Zrxtrade.contract.Call(opts, out, "ETH_CONTRACT_ADDRESS")
	return *ret0, err
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_Zrxtrade *ZrxtradeSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _Zrxtrade.Contract.ETHCONTRACTADDRESS(&_Zrxtrade.CallOpts)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_Zrxtrade *ZrxtradeCallerSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _Zrxtrade.Contract.ETHCONTRACTADDRESS(&_Zrxtrade.CallOpts)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_Zrxtrade *ZrxtradeCaller) IncognitoSmartContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Zrxtrade.contract.Call(opts, out, "incognitoSmartContract")
	return *ret0, err
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_Zrxtrade *ZrxtradeSession) IncognitoSmartContract() (common.Address, error) {
	return _Zrxtrade.Contract.IncognitoSmartContract(&_Zrxtrade.CallOpts)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_Zrxtrade *ZrxtradeCallerSession) IncognitoSmartContract() (common.Address, error) {
	return _Zrxtrade.Contract.IncognitoSmartContract(&_Zrxtrade.CallOpts)
}

// Trade is a paid mutator transaction binding the contract method 0x98c36152.
//
// Solidity: function trade(address srcToken, uint256 amount, address destToken, bytes callDataHex, address _forwarder) returns(address, uint256)
func (_Zrxtrade *ZrxtradeTransactor) Trade(opts *bind.TransactOpts, srcToken common.Address, amount *big.Int, destToken common.Address, callDataHex []byte, _forwarder common.Address) (*types.Transaction, error) {
	return _Zrxtrade.contract.Transact(opts, "trade", srcToken, amount, destToken, callDataHex, _forwarder)
}

// Trade is a paid mutator transaction binding the contract method 0x98c36152.
//
// Solidity: function trade(address srcToken, uint256 amount, address destToken, bytes callDataHex, address _forwarder) returns(address, uint256)
func (_Zrxtrade *ZrxtradeSession) Trade(srcToken common.Address, amount *big.Int, destToken common.Address, callDataHex []byte, _forwarder common.Address) (*types.Transaction, error) {
	return _Zrxtrade.Contract.Trade(&_Zrxtrade.TransactOpts, srcToken, amount, destToken, callDataHex, _forwarder)
}

// Trade is a paid mutator transaction binding the contract method 0x98c36152.
//
// Solidity: function trade(address srcToken, uint256 amount, address destToken, bytes callDataHex, address _forwarder) returns(address, uint256)
func (_Zrxtrade *ZrxtradeTransactorSession) Trade(srcToken common.Address, amount *big.Int, destToken common.Address, callDataHex []byte, _forwarder common.Address) (*types.Transaction, error) {
	return _Zrxtrade.Contract.Trade(&_Zrxtrade.TransactOpts, srcToken, amount, destToken, callDataHex, _forwarder)
}

// WithdrawWrapETH is a paid mutator transaction binding the contract method 0x0e0569a6.
//
// Solidity: function withdrawWrapETH(uint256 amount) returns()
func (_Zrxtrade *ZrxtradeTransactor) WithdrawWrapETH(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Zrxtrade.contract.Transact(opts, "withdrawWrapETH", amount)
}

// WithdrawWrapETH is a paid mutator transaction binding the contract method 0x0e0569a6.
//
// Solidity: function withdrawWrapETH(uint256 amount) returns()
func (_Zrxtrade *ZrxtradeSession) WithdrawWrapETH(amount *big.Int) (*types.Transaction, error) {
	return _Zrxtrade.Contract.WithdrawWrapETH(&_Zrxtrade.TransactOpts, amount)
}

// WithdrawWrapETH is a paid mutator transaction binding the contract method 0x0e0569a6.
//
// Solidity: function withdrawWrapETH(uint256 amount) returns()
func (_Zrxtrade *ZrxtradeTransactorSession) WithdrawWrapETH(amount *big.Int) (*types.Transaction, error) {
	return _Zrxtrade.Contract.WithdrawWrapETH(&_Zrxtrade.TransactOpts, amount)
}
