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
var ZrxtradeBin = "0x608060405234801561001057600080fd5b50604051610a36380380610a3683398101604081905261002f91610085565b60018054600160a060020a03938416600160a060020a0319918216179091556000805492841692821692909217909155600280549390921692169190911790556100fa565b805161007f816100e3565b92915050565b60008060006060848603121561009a57600080fd5b60006100a68686610074565b93505060206100b786828701610074565b92505060406100c886828701610074565b9150509250925092565b6000600160a060020a03821661007f565b6100ec816100d2565b81146100f757600080fd5b50565b61092d806101096000396000f3fe60806040526004361061005b577c010000000000000000000000000000000000000000000000000000000060003504630e0569a6811461005d57806372e94bf61461007d57806398c36152146100a8578063b42a644b146100c9575b005b34801561006957600080fd5b5061005b6100783660046106c2565b6100eb565b34801561008957600080fd5b50610092610169565b60405161009f91906107db565b60405180910390f35b6100bb6100b6366004610635565b61016e565b60405161009f9291906107cd565b3480156100d557600080fd5b506100de61026a565b60405161009f9190610789565b6002546040517f2e1a7d4d000000000000000000000000000000000000000000000000000000008152600160a060020a0390911690632e1a7d4d906101349084906004016107e9565b600060405180830381600087803b15801561014e57600080fd5b505af1158015610162573d6000803e3d6000fd5b5050505050565b600081565b600080548190600160a060020a0316331461018857600080fd5b6001546101a0908890600160a060020a031688610279565b6000839050600081600160a060020a031634876040516101c09190610768565b60006040518083038185875af1925050503d80600081146101fd576040519150601f19603f3d011682016040523d82523d6000602084013e610202565b606091505b505090508061021057600080fd5b6000600160a060020a0388166102455761022a60006103a2565b9050610235816100eb565b6102406000826103df565b61025a565b61024e886103a2565b905061025a88826103df565b9699969850959650505050505050565b600054600160a060020a031681565b600160a060020a0383161561039d576040517f095ea7b3000000000000000000000000000000000000000000000000000000008152600160a060020a0384169063095ea7b3906102d09085906000906004016107b2565b600060405180830381600087803b1580156102ea57600080fd5b505af11580156102fe573d6000803e3d6000fd5b5050505061030a6104ce565b61031357600080fd5b6040517f095ea7b3000000000000000000000000000000000000000000000000000000008152600160a060020a0384169063095ea7b39061035a90859085906004016107cd565b600060405180830381600087803b15801561037457600080fd5b505af1158015610388573d6000803e3d6000fd5b505050506103946104ce565b61039d57600080fd5b505050565b6000600160a060020a0382166103ce576002546103c790600160a060020a0316610502565b90506103da565b6103d782610502565b90505b919050565b600160a060020a03821661043a5730318111156103fb57600080fd5b60008054604051600160a060020a039091169183156108fc02918491818181858888f19350505050158015610434573d6000803e3d6000fd5b506104ca565b6000546040517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a038481169263a9059cbb9261048792909116908590600401610797565b600060405180830381600087803b1580156104a157600080fd5b505af11580156104b5573d6000803e3d6000fd5b505050506104c16104ce565b6104ca57600080fd5b5050565b6000803d80156104e557602081146104ee576104fa565b600191506104fa565b60206000803e60005191505b501515905090565b6000600160a060020a03821661051a575030316103da565b6040517f70a08231000000000000000000000000000000000000000000000000000000008152600160a060020a038316906370a082319061055f90309060040161077b565b60206040518083038186803b15801561057757600080fd5b505afa15801561058b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506103d791908101906106e8565b80356105ba816108c1565b92915050565b600082601f8301126105d157600080fd5b81356105e46105df8261081e565b6107f7565b9150808252602083016020830185838301111561060057600080fd5b61060b838284610885565b50505092915050565b80356105ba816108d8565b80356105ba816108e1565b80516105ba816108e1565b600080600080600060a0868803121561064d57600080fd5b60006106598888610614565b955050602061066a8882890161061f565b945050604061067b88828901610614565b935050606086013567ffffffffffffffff81111561069857600080fd5b6106a4888289016105c0565b92505060806106b5888289016105af565b9150509295509295909350565b6000602082840312156106d457600080fd5b60006106e0848461061f565b949350505050565b6000602082840312156106fa57600080fd5b60006106e0848461062a565b61070f8161086f565b82525050565b61070f8161084a565b600061072982610846565b61073381856103da565b9350610743818560208601610891565b9290920192915050565b61070f81610855565b61070f8161087a565b61070f8161086c565b6000610774828461071e565b9392505050565b602081016105ba8284610706565b602081016105ba8284610715565b604081016107a58285610706565b610774602083018461075f565b604081016107c08285610715565b6107746020830184610756565b604081016107a58285610715565b602081016105ba828461074d565b602081016105ba828461075f565b60405181810167ffffffffffffffff8111828210171561081657600080fd5b604052919050565b600067ffffffffffffffff82111561083557600080fd5b506020601f91909101601f19160190565b5190565b60006103d782610860565b60006103d78261084a565b600160a060020a031690565b90565b60006103d782610855565b60006103d78261086c565b82818337506000910152565b60005b838110156108ac578181015183820152602001610894565b838111156108bb576000848401525b50505050565b6108ca8161084a565b81146108d557600080fd5b50565b6108ca81610855565b6108ca8161086c56fea365627a7a723158207560a9b9bfbc13e8436531b5f74602c4a246178403df75e7df06faa84e4720b06c6578706572696d656e74616cf564736f6c634300050c0040"

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
