// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iweth

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

// IwethMetaData contains all meta data concerning the Iweth contract.
var IwethMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IwethABI is the input ABI used to generate the binding from.
// Deprecated: Use IwethMetaData.ABI instead.
var IwethABI = IwethMetaData.ABI

// Iweth is an auto generated Go binding around an Ethereum contract.
type Iweth struct {
	IwethCaller     // Read-only binding to the contract
	IwethTransactor // Write-only binding to the contract
	IwethFilterer   // Log filterer for contract events
}

// IwethCaller is an auto generated read-only Go binding around an Ethereum contract.
type IwethCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IwethTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IwethTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IwethFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IwethFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IwethSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IwethSession struct {
	Contract     *Iweth            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IwethCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IwethCallerSession struct {
	Contract *IwethCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IwethTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IwethTransactorSession struct {
	Contract     *IwethTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IwethRaw is an auto generated low-level Go binding around an Ethereum contract.
type IwethRaw struct {
	Contract *Iweth // Generic contract binding to access the raw methods on
}

// IwethCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IwethCallerRaw struct {
	Contract *IwethCaller // Generic read-only contract binding to access the raw methods on
}

// IwethTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IwethTransactorRaw struct {
	Contract *IwethTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIweth creates a new instance of Iweth, bound to a specific deployed contract.
func NewIweth(address common.Address, backend bind.ContractBackend) (*Iweth, error) {
	contract, err := bindIweth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iweth{IwethCaller: IwethCaller{contract: contract}, IwethTransactor: IwethTransactor{contract: contract}, IwethFilterer: IwethFilterer{contract: contract}}, nil
}

// NewIwethCaller creates a new read-only instance of Iweth, bound to a specific deployed contract.
func NewIwethCaller(address common.Address, caller bind.ContractCaller) (*IwethCaller, error) {
	contract, err := bindIweth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IwethCaller{contract: contract}, nil
}

// NewIwethTransactor creates a new write-only instance of Iweth, bound to a specific deployed contract.
func NewIwethTransactor(address common.Address, transactor bind.ContractTransactor) (*IwethTransactor, error) {
	contract, err := bindIweth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IwethTransactor{contract: contract}, nil
}

// NewIwethFilterer creates a new log filterer instance of Iweth, bound to a specific deployed contract.
func NewIwethFilterer(address common.Address, filterer bind.ContractFilterer) (*IwethFilterer, error) {
	contract, err := bindIweth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IwethFilterer{contract: contract}, nil
}

// bindIweth binds a generic wrapper to an already deployed contract.
func bindIweth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IwethABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iweth *IwethRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iweth.Contract.IwethCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iweth *IwethRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iweth.Contract.IwethTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iweth *IwethRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iweth.Contract.IwethTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iweth *IwethCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iweth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iweth *IwethTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iweth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iweth *IwethTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iweth.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Iweth *IwethTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iweth.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Iweth *IwethSession) Deposit() (*types.Transaction, error) {
	return _Iweth.Contract.Deposit(&_Iweth.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Iweth *IwethTransactorSession) Deposit() (*types.Transaction, error) {
	return _Iweth.Contract.Deposit(&_Iweth.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_Iweth *IwethTransactor) Withdraw(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _Iweth.contract.Transact(opts, "withdraw", arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_Iweth *IwethSession) Withdraw(arg0 *big.Int) (*types.Transaction, error) {
	return _Iweth.Contract.Withdraw(&_Iweth.TransactOpts, arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_Iweth *IwethTransactorSession) Withdraw(arg0 *big.Int) (*types.Transaction, error) {
	return _Iweth.Contract.Withdraw(&_Iweth.TransactOpts, arg0)
}
