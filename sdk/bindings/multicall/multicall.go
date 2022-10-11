// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package multicall

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IMulticall2Call is an auto generated low-level Go binding around an user-defined struct.
type IMulticall2Call struct {
	Target   common.Address
	CallData []byte
}

// IMulticall2Result is an auto generated low-level Go binding around an user-defined struct.
type IMulticall2Result struct {
	Success    bool
	ReturnData []byte
}

// IMulticall2MetaData contains all meta data concerning the IMulticall2 contract.
var IMulticall2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"requireSuccess\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structIMulticall2.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"tryBlockAndAggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structIMulticall2.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IMulticall2ABI is the input ABI used to generate the binding from.
// Deprecated: Use IMulticall2MetaData.ABI instead.
var IMulticall2ABI = IMulticall2MetaData.ABI

// IMulticall2 is an auto generated Go binding around an Ethereum contract.
type IMulticall2 struct {
	IMulticall2Caller     // Read-only binding to the contract
	IMulticall2Transactor // Write-only binding to the contract
	IMulticall2Filterer   // Log filterer for contract events
}

// IMulticall2Caller is an auto generated read-only Go binding around an Ethereum contract.
type IMulticall2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMulticall2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IMulticall2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMulticall2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMulticall2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMulticall2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMulticall2Session struct {
	Contract     *IMulticall2      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMulticall2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMulticall2CallerSession struct {
	Contract *IMulticall2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IMulticall2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMulticall2TransactorSession struct {
	Contract     *IMulticall2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IMulticall2Raw is an auto generated low-level Go binding around an Ethereum contract.
type IMulticall2Raw struct {
	Contract *IMulticall2 // Generic contract binding to access the raw methods on
}

// IMulticall2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMulticall2CallerRaw struct {
	Contract *IMulticall2Caller // Generic read-only contract binding to access the raw methods on
}

// IMulticall2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMulticall2TransactorRaw struct {
	Contract *IMulticall2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIMulticall2 creates a new instance of IMulticall2, bound to a specific deployed contract.
func NewIMulticall2(address common.Address, backend bind.ContractBackend) (*IMulticall2, error) {
	contract, err := bindIMulticall2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMulticall2{IMulticall2Caller: IMulticall2Caller{contract: contract}, IMulticall2Transactor: IMulticall2Transactor{contract: contract}, IMulticall2Filterer: IMulticall2Filterer{contract: contract}}, nil
}

// NewIMulticall2Caller creates a new read-only instance of IMulticall2, bound to a specific deployed contract.
func NewIMulticall2Caller(address common.Address, caller bind.ContractCaller) (*IMulticall2Caller, error) {
	contract, err := bindIMulticall2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMulticall2Caller{contract: contract}, nil
}

// NewIMulticall2Transactor creates a new write-only instance of IMulticall2, bound to a specific deployed contract.
func NewIMulticall2Transactor(address common.Address, transactor bind.ContractTransactor) (*IMulticall2Transactor, error) {
	contract, err := bindIMulticall2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMulticall2Transactor{contract: contract}, nil
}

// NewIMulticall2Filterer creates a new log filterer instance of IMulticall2, bound to a specific deployed contract.
func NewIMulticall2Filterer(address common.Address, filterer bind.ContractFilterer) (*IMulticall2Filterer, error) {
	contract, err := bindIMulticall2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMulticall2Filterer{contract: contract}, nil
}

// bindIMulticall2 binds a generic wrapper to an already deployed contract.
func bindIMulticall2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMulticall2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMulticall2 *IMulticall2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMulticall2.Contract.IMulticall2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMulticall2 *IMulticall2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMulticall2.Contract.IMulticall2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMulticall2 *IMulticall2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMulticall2.Contract.IMulticall2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMulticall2 *IMulticall2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMulticall2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMulticall2 *IMulticall2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMulticall2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMulticall2 *IMulticall2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMulticall2.Contract.contract.Transact(opts, method, params...)
}

// TryBlockAndAggregate is a paid mutator transaction binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_IMulticall2 *IMulticall2Transactor) TryBlockAndAggregate(opts *bind.TransactOpts, requireSuccess bool, calls []IMulticall2Call) (*types.Transaction, error) {
	return _IMulticall2.contract.Transact(opts, "tryBlockAndAggregate", requireSuccess, calls)
}

// TryBlockAndAggregate is a paid mutator transaction binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_IMulticall2 *IMulticall2Session) TryBlockAndAggregate(requireSuccess bool, calls []IMulticall2Call) (*types.Transaction, error) {
	return _IMulticall2.Contract.TryBlockAndAggregate(&_IMulticall2.TransactOpts, requireSuccess, calls)
}

// TryBlockAndAggregate is a paid mutator transaction binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_IMulticall2 *IMulticall2TransactorSession) TryBlockAndAggregate(requireSuccess bool, calls []IMulticall2Call) (*types.Transaction, error) {
	return _IMulticall2.Contract.TryBlockAndAggregate(&_IMulticall2.TransactOpts, requireSuccess, calls)
}
