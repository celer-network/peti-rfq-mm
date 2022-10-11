// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cbridge

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

// MsgDataTypesRouteInfo is an auto generated low-level Go binding around an user-defined struct.
type MsgDataTypesRouteInfo struct {
	Sender     common.Address
	Receiver   common.Address
	SrcChainId uint64
	SrcTxHash  [32]byte
}

// MsgDataTypesTransferInfo is an auto generated low-level Go binding around an user-defined struct.
type MsgDataTypesTransferInfo struct {
	T          uint8
	Sender     common.Address
	Receiver   common.Address
	Token      common.Address
	Amount     *big.Int
	Wdseq      uint64
	SrcChainId uint64
	RefId      [32]byte
	SrcTxHash  [32]byte
}

// IMessageBusMetaData contains all meta data concerning the IMessageBus contract.
var IMessageBusMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"calcFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.RouteInfo\",\"name\":\"_route\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMsgDataTypes.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"wdseq\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.TransferInfo\",\"name\":\"_transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessageWithTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMsgDataTypes.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"wdseq\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.TransferInfo\",\"name\":\"_transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessageWithTransferRefund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePerByte\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegBridgeV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegVaultV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_srcBridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_srcTransferId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendMessageWithTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_cumulativeFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"withdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IMessageBusABI is the input ABI used to generate the binding from.
// Deprecated: Use IMessageBusMetaData.ABI instead.
var IMessageBusABI = IMessageBusMetaData.ABI

// IMessageBus is an auto generated Go binding around an Ethereum contract.
type IMessageBus struct {
	IMessageBusCaller     // Read-only binding to the contract
	IMessageBusTransactor // Write-only binding to the contract
	IMessageBusFilterer   // Log filterer for contract events
}

// IMessageBusCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMessageBusCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageBusTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMessageBusTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageBusFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMessageBusFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageBusSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMessageBusSession struct {
	Contract     *IMessageBus      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMessageBusCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMessageBusCallerSession struct {
	Contract *IMessageBusCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IMessageBusTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMessageBusTransactorSession struct {
	Contract     *IMessageBusTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IMessageBusRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMessageBusRaw struct {
	Contract *IMessageBus // Generic contract binding to access the raw methods on
}

// IMessageBusCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMessageBusCallerRaw struct {
	Contract *IMessageBusCaller // Generic read-only contract binding to access the raw methods on
}

// IMessageBusTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMessageBusTransactorRaw struct {
	Contract *IMessageBusTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMessageBus creates a new instance of IMessageBus, bound to a specific deployed contract.
func NewIMessageBus(address common.Address, backend bind.ContractBackend) (*IMessageBus, error) {
	contract, err := bindIMessageBus(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMessageBus{IMessageBusCaller: IMessageBusCaller{contract: contract}, IMessageBusTransactor: IMessageBusTransactor{contract: contract}, IMessageBusFilterer: IMessageBusFilterer{contract: contract}}, nil
}

// NewIMessageBusCaller creates a new read-only instance of IMessageBus, bound to a specific deployed contract.
func NewIMessageBusCaller(address common.Address, caller bind.ContractCaller) (*IMessageBusCaller, error) {
	contract, err := bindIMessageBus(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageBusCaller{contract: contract}, nil
}

// NewIMessageBusTransactor creates a new write-only instance of IMessageBus, bound to a specific deployed contract.
func NewIMessageBusTransactor(address common.Address, transactor bind.ContractTransactor) (*IMessageBusTransactor, error) {
	contract, err := bindIMessageBus(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageBusTransactor{contract: contract}, nil
}

// NewIMessageBusFilterer creates a new log filterer instance of IMessageBus, bound to a specific deployed contract.
func NewIMessageBusFilterer(address common.Address, filterer bind.ContractFilterer) (*IMessageBusFilterer, error) {
	contract, err := bindIMessageBus(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMessageBusFilterer{contract: contract}, nil
}

// bindIMessageBus binds a generic wrapper to an already deployed contract.
func bindIMessageBus(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMessageBusABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageBus *IMessageBusRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageBus.Contract.IMessageBusCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageBus *IMessageBusRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageBus.Contract.IMessageBusTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageBus *IMessageBusRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageBus.Contract.IMessageBusTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageBus *IMessageBusCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageBus.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageBus *IMessageBusTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageBus.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageBus *IMessageBusTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageBus.Contract.contract.Transact(opts, method, params...)
}

// CalcFee is a free data retrieval call binding the contract method 0x5335dca2.
//
// Solidity: function calcFee(bytes _message) view returns(uint256)
func (_IMessageBus *IMessageBusCaller) CalcFee(opts *bind.CallOpts, _message []byte) (*big.Int, error) {
	var out []interface{}
	err := _IMessageBus.contract.Call(opts, &out, "calcFee", _message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcFee is a free data retrieval call binding the contract method 0x5335dca2.
//
// Solidity: function calcFee(bytes _message) view returns(uint256)
func (_IMessageBus *IMessageBusSession) CalcFee(_message []byte) (*big.Int, error) {
	return _IMessageBus.Contract.CalcFee(&_IMessageBus.CallOpts, _message)
}

// CalcFee is a free data retrieval call binding the contract method 0x5335dca2.
//
// Solidity: function calcFee(bytes _message) view returns(uint256)
func (_IMessageBus *IMessageBusCallerSession) CalcFee(_message []byte) (*big.Int, error) {
	return _IMessageBus.Contract.CalcFee(&_IMessageBus.CallOpts, _message)
}

// FeeBase is a free data retrieval call binding the contract method 0x95e911a8.
//
// Solidity: function feeBase() view returns(uint256)
func (_IMessageBus *IMessageBusCaller) FeeBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IMessageBus.contract.Call(opts, &out, "feeBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeBase is a free data retrieval call binding the contract method 0x95e911a8.
//
// Solidity: function feeBase() view returns(uint256)
func (_IMessageBus *IMessageBusSession) FeeBase() (*big.Int, error) {
	return _IMessageBus.Contract.FeeBase(&_IMessageBus.CallOpts)
}

// FeeBase is a free data retrieval call binding the contract method 0x95e911a8.
//
// Solidity: function feeBase() view returns(uint256)
func (_IMessageBus *IMessageBusCallerSession) FeeBase() (*big.Int, error) {
	return _IMessageBus.Contract.FeeBase(&_IMessageBus.CallOpts)
}

// FeePerByte is a free data retrieval call binding the contract method 0xf60bbe2a.
//
// Solidity: function feePerByte() view returns(uint256)
func (_IMessageBus *IMessageBusCaller) FeePerByte(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IMessageBus.contract.Call(opts, &out, "feePerByte")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeePerByte is a free data retrieval call binding the contract method 0xf60bbe2a.
//
// Solidity: function feePerByte() view returns(uint256)
func (_IMessageBus *IMessageBusSession) FeePerByte() (*big.Int, error) {
	return _IMessageBus.Contract.FeePerByte(&_IMessageBus.CallOpts)
}

// FeePerByte is a free data retrieval call binding the contract method 0xf60bbe2a.
//
// Solidity: function feePerByte() view returns(uint256)
func (_IMessageBus *IMessageBusCallerSession) FeePerByte() (*big.Int, error) {
	return _IMessageBus.Contract.FeePerByte(&_IMessageBus.CallOpts)
}

// LiquidityBridge is a free data retrieval call binding the contract method 0x82980dc4.
//
// Solidity: function liquidityBridge() view returns(address)
func (_IMessageBus *IMessageBusCaller) LiquidityBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IMessageBus.contract.Call(opts, &out, "liquidityBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidityBridge is a free data retrieval call binding the contract method 0x82980dc4.
//
// Solidity: function liquidityBridge() view returns(address)
func (_IMessageBus *IMessageBusSession) LiquidityBridge() (common.Address, error) {
	return _IMessageBus.Contract.LiquidityBridge(&_IMessageBus.CallOpts)
}

// LiquidityBridge is a free data retrieval call binding the contract method 0x82980dc4.
//
// Solidity: function liquidityBridge() view returns(address)
func (_IMessageBus *IMessageBusCallerSession) LiquidityBridge() (common.Address, error) {
	return _IMessageBus.Contract.LiquidityBridge(&_IMessageBus.CallOpts)
}

// PegBridge is a free data retrieval call binding the contract method 0xdfa2dbaf.
//
// Solidity: function pegBridge() view returns(address)
func (_IMessageBus *IMessageBusCaller) PegBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IMessageBus.contract.Call(opts, &out, "pegBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegBridge is a free data retrieval call binding the contract method 0xdfa2dbaf.
//
// Solidity: function pegBridge() view returns(address)
func (_IMessageBus *IMessageBusSession) PegBridge() (common.Address, error) {
	return _IMessageBus.Contract.PegBridge(&_IMessageBus.CallOpts)
}

// PegBridge is a free data retrieval call binding the contract method 0xdfa2dbaf.
//
// Solidity: function pegBridge() view returns(address)
func (_IMessageBus *IMessageBusCallerSession) PegBridge() (common.Address, error) {
	return _IMessageBus.Contract.PegBridge(&_IMessageBus.CallOpts)
}

// PegBridgeV2 is a free data retrieval call binding the contract method 0x95b12c27.
//
// Solidity: function pegBridgeV2() view returns(address)
func (_IMessageBus *IMessageBusCaller) PegBridgeV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IMessageBus.contract.Call(opts, &out, "pegBridgeV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegBridgeV2 is a free data retrieval call binding the contract method 0x95b12c27.
//
// Solidity: function pegBridgeV2() view returns(address)
func (_IMessageBus *IMessageBusSession) PegBridgeV2() (common.Address, error) {
	return _IMessageBus.Contract.PegBridgeV2(&_IMessageBus.CallOpts)
}

// PegBridgeV2 is a free data retrieval call binding the contract method 0x95b12c27.
//
// Solidity: function pegBridgeV2() view returns(address)
func (_IMessageBus *IMessageBusCallerSession) PegBridgeV2() (common.Address, error) {
	return _IMessageBus.Contract.PegBridgeV2(&_IMessageBus.CallOpts)
}

// PegVault is a free data retrieval call binding the contract method 0xd8257d17.
//
// Solidity: function pegVault() view returns(address)
func (_IMessageBus *IMessageBusCaller) PegVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IMessageBus.contract.Call(opts, &out, "pegVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegVault is a free data retrieval call binding the contract method 0xd8257d17.
//
// Solidity: function pegVault() view returns(address)
func (_IMessageBus *IMessageBusSession) PegVault() (common.Address, error) {
	return _IMessageBus.Contract.PegVault(&_IMessageBus.CallOpts)
}

// PegVault is a free data retrieval call binding the contract method 0xd8257d17.
//
// Solidity: function pegVault() view returns(address)
func (_IMessageBus *IMessageBusCallerSession) PegVault() (common.Address, error) {
	return _IMessageBus.Contract.PegVault(&_IMessageBus.CallOpts)
}

// PegVaultV2 is a free data retrieval call binding the contract method 0xc66a9c5a.
//
// Solidity: function pegVaultV2() view returns(address)
func (_IMessageBus *IMessageBusCaller) PegVaultV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IMessageBus.contract.Call(opts, &out, "pegVaultV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegVaultV2 is a free data retrieval call binding the contract method 0xc66a9c5a.
//
// Solidity: function pegVaultV2() view returns(address)
func (_IMessageBus *IMessageBusSession) PegVaultV2() (common.Address, error) {
	return _IMessageBus.Contract.PegVaultV2(&_IMessageBus.CallOpts)
}

// PegVaultV2 is a free data retrieval call binding the contract method 0xc66a9c5a.
//
// Solidity: function pegVaultV2() view returns(address)
func (_IMessageBus *IMessageBusCallerSession) PegVaultV2() (common.Address, error) {
	return _IMessageBus.Contract.PegVaultV2(&_IMessageBus.CallOpts)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x468a2d04.
//
// Solidity: function executeMessage(bytes _message, (address,address,uint64,bytes32) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_IMessageBus *IMessageBusTransactor) ExecuteMessage(opts *bind.TransactOpts, _message []byte, _route MsgDataTypesRouteInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.contract.Transact(opts, "executeMessage", _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x468a2d04.
//
// Solidity: function executeMessage(bytes _message, (address,address,uint64,bytes32) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_IMessageBus *IMessageBusSession) ExecuteMessage(_message []byte, _route MsgDataTypesRouteInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.Contract.ExecuteMessage(&_IMessageBus.TransactOpts, _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x468a2d04.
//
// Solidity: function executeMessage(bytes _message, (address,address,uint64,bytes32) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_IMessageBus *IMessageBusTransactorSession) ExecuteMessage(_message []byte, _route MsgDataTypesRouteInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.Contract.ExecuteMessage(&_IMessageBus.TransactOpts, _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x3f395aff.
//
// Solidity: function executeMessageWithTransfer(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_IMessageBus *IMessageBusTransactor) ExecuteMessageWithTransfer(opts *bind.TransactOpts, _message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.contract.Transact(opts, "executeMessageWithTransfer", _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x3f395aff.
//
// Solidity: function executeMessageWithTransfer(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_IMessageBus *IMessageBusSession) ExecuteMessageWithTransfer(_message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.Contract.ExecuteMessageWithTransfer(&_IMessageBus.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x3f395aff.
//
// Solidity: function executeMessageWithTransfer(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_IMessageBus *IMessageBusTransactorSession) ExecuteMessageWithTransfer(_message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.Contract.ExecuteMessageWithTransfer(&_IMessageBus.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x7b80ab20.
//
// Solidity: function executeMessageWithTransferRefund(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_IMessageBus *IMessageBusTransactor) ExecuteMessageWithTransferRefund(opts *bind.TransactOpts, _message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.contract.Transact(opts, "executeMessageWithTransferRefund", _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x7b80ab20.
//
// Solidity: function executeMessageWithTransferRefund(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_IMessageBus *IMessageBusSession) ExecuteMessageWithTransferRefund(_message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.Contract.ExecuteMessageWithTransferRefund(&_IMessageBus.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x7b80ab20.
//
// Solidity: function executeMessageWithTransferRefund(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_IMessageBus *IMessageBusTransactorSession) ExecuteMessageWithTransferRefund(_message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.Contract.ExecuteMessageWithTransferRefund(&_IMessageBus.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_IMessageBus *IMessageBusTransactor) SendMessage(opts *bind.TransactOpts, _receiver common.Address, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _IMessageBus.contract.Transact(opts, "sendMessage", _receiver, _dstChainId, _message)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_IMessageBus *IMessageBusSession) SendMessage(_receiver common.Address, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _IMessageBus.Contract.SendMessage(&_IMessageBus.TransactOpts, _receiver, _dstChainId, _message)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_IMessageBus *IMessageBusTransactorSession) SendMessage(_receiver common.Address, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _IMessageBus.Contract.SendMessage(&_IMessageBus.TransactOpts, _receiver, _dstChainId, _message)
}

// SendMessageWithTransfer is a paid mutator transaction binding the contract method 0x4289fbb3.
//
// Solidity: function sendMessageWithTransfer(address _receiver, uint256 _dstChainId, address _srcBridge, bytes32 _srcTransferId, bytes _message) payable returns()
func (_IMessageBus *IMessageBusTransactor) SendMessageWithTransfer(opts *bind.TransactOpts, _receiver common.Address, _dstChainId *big.Int, _srcBridge common.Address, _srcTransferId [32]byte, _message []byte) (*types.Transaction, error) {
	return _IMessageBus.contract.Transact(opts, "sendMessageWithTransfer", _receiver, _dstChainId, _srcBridge, _srcTransferId, _message)
}

// SendMessageWithTransfer is a paid mutator transaction binding the contract method 0x4289fbb3.
//
// Solidity: function sendMessageWithTransfer(address _receiver, uint256 _dstChainId, address _srcBridge, bytes32 _srcTransferId, bytes _message) payable returns()
func (_IMessageBus *IMessageBusSession) SendMessageWithTransfer(_receiver common.Address, _dstChainId *big.Int, _srcBridge common.Address, _srcTransferId [32]byte, _message []byte) (*types.Transaction, error) {
	return _IMessageBus.Contract.SendMessageWithTransfer(&_IMessageBus.TransactOpts, _receiver, _dstChainId, _srcBridge, _srcTransferId, _message)
}

// SendMessageWithTransfer is a paid mutator transaction binding the contract method 0x4289fbb3.
//
// Solidity: function sendMessageWithTransfer(address _receiver, uint256 _dstChainId, address _srcBridge, bytes32 _srcTransferId, bytes _message) payable returns()
func (_IMessageBus *IMessageBusTransactorSession) SendMessageWithTransfer(_receiver common.Address, _dstChainId *big.Int, _srcBridge common.Address, _srcTransferId [32]byte, _message []byte) (*types.Transaction, error) {
	return _IMessageBus.Contract.SendMessageWithTransfer(&_IMessageBus.TransactOpts, _receiver, _dstChainId, _srcBridge, _srcTransferId, _message)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x2ff4c411.
//
// Solidity: function withdrawFee(address _account, uint256 _cumulativeFee, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_IMessageBus *IMessageBusTransactor) WithdrawFee(opts *bind.TransactOpts, _account common.Address, _cumulativeFee *big.Int, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.contract.Transact(opts, "withdrawFee", _account, _cumulativeFee, _sigs, _signers, _powers)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x2ff4c411.
//
// Solidity: function withdrawFee(address _account, uint256 _cumulativeFee, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_IMessageBus *IMessageBusSession) WithdrawFee(_account common.Address, _cumulativeFee *big.Int, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.Contract.WithdrawFee(&_IMessageBus.TransactOpts, _account, _cumulativeFee, _sigs, _signers, _powers)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x2ff4c411.
//
// Solidity: function withdrawFee(address _account, uint256 _cumulativeFee, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_IMessageBus *IMessageBusTransactorSession) WithdrawFee(_account common.Address, _cumulativeFee *big.Int, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.Contract.WithdrawFee(&_IMessageBus.TransactOpts, _account, _cumulativeFee, _sigs, _signers, _powers)
}

// MsgDataTypesMetaData contains all meta data concerning the MsgDataTypes contract.
var MsgDataTypesMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220623342e0418be6ac3b4de01a1763c1d8311a3483683f2a366814784af245669d64736f6c634300080f0033",
}

// MsgDataTypesABI is the input ABI used to generate the binding from.
// Deprecated: Use MsgDataTypesMetaData.ABI instead.
var MsgDataTypesABI = MsgDataTypesMetaData.ABI

// MsgDataTypesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MsgDataTypesMetaData.Bin instead.
var MsgDataTypesBin = MsgDataTypesMetaData.Bin

// DeployMsgDataTypes deploys a new Ethereum contract, binding an instance of MsgDataTypes to it.
func DeployMsgDataTypes(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MsgDataTypes, error) {
	parsed, err := MsgDataTypesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MsgDataTypesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MsgDataTypes{MsgDataTypesCaller: MsgDataTypesCaller{contract: contract}, MsgDataTypesTransactor: MsgDataTypesTransactor{contract: contract}, MsgDataTypesFilterer: MsgDataTypesFilterer{contract: contract}}, nil
}

// MsgDataTypes is an auto generated Go binding around an Ethereum contract.
type MsgDataTypes struct {
	MsgDataTypesCaller     // Read-only binding to the contract
	MsgDataTypesTransactor // Write-only binding to the contract
	MsgDataTypesFilterer   // Log filterer for contract events
}

// MsgDataTypesCaller is an auto generated read-only Go binding around an Ethereum contract.
type MsgDataTypesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MsgDataTypesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MsgDataTypesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MsgDataTypesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MsgDataTypesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MsgDataTypesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MsgDataTypesSession struct {
	Contract     *MsgDataTypes     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MsgDataTypesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MsgDataTypesCallerSession struct {
	Contract *MsgDataTypesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MsgDataTypesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MsgDataTypesTransactorSession struct {
	Contract     *MsgDataTypesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MsgDataTypesRaw is an auto generated low-level Go binding around an Ethereum contract.
type MsgDataTypesRaw struct {
	Contract *MsgDataTypes // Generic contract binding to access the raw methods on
}

// MsgDataTypesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MsgDataTypesCallerRaw struct {
	Contract *MsgDataTypesCaller // Generic read-only contract binding to access the raw methods on
}

// MsgDataTypesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MsgDataTypesTransactorRaw struct {
	Contract *MsgDataTypesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMsgDataTypes creates a new instance of MsgDataTypes, bound to a specific deployed contract.
func NewMsgDataTypes(address common.Address, backend bind.ContractBackend) (*MsgDataTypes, error) {
	contract, err := bindMsgDataTypes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MsgDataTypes{MsgDataTypesCaller: MsgDataTypesCaller{contract: contract}, MsgDataTypesTransactor: MsgDataTypesTransactor{contract: contract}, MsgDataTypesFilterer: MsgDataTypesFilterer{contract: contract}}, nil
}

// NewMsgDataTypesCaller creates a new read-only instance of MsgDataTypes, bound to a specific deployed contract.
func NewMsgDataTypesCaller(address common.Address, caller bind.ContractCaller) (*MsgDataTypesCaller, error) {
	contract, err := bindMsgDataTypes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MsgDataTypesCaller{contract: contract}, nil
}

// NewMsgDataTypesTransactor creates a new write-only instance of MsgDataTypes, bound to a specific deployed contract.
func NewMsgDataTypesTransactor(address common.Address, transactor bind.ContractTransactor) (*MsgDataTypesTransactor, error) {
	contract, err := bindMsgDataTypes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MsgDataTypesTransactor{contract: contract}, nil
}

// NewMsgDataTypesFilterer creates a new log filterer instance of MsgDataTypes, bound to a specific deployed contract.
func NewMsgDataTypesFilterer(address common.Address, filterer bind.ContractFilterer) (*MsgDataTypesFilterer, error) {
	contract, err := bindMsgDataTypes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MsgDataTypesFilterer{contract: contract}, nil
}

// bindMsgDataTypes binds a generic wrapper to an already deployed contract.
func bindMsgDataTypes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MsgDataTypesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MsgDataTypes *MsgDataTypesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MsgDataTypes.Contract.MsgDataTypesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MsgDataTypes *MsgDataTypesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MsgDataTypes.Contract.MsgDataTypesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MsgDataTypes *MsgDataTypesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MsgDataTypes.Contract.MsgDataTypesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MsgDataTypes *MsgDataTypesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MsgDataTypes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MsgDataTypes *MsgDataTypesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MsgDataTypes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MsgDataTypes *MsgDataTypesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MsgDataTypes.Contract.contract.Transact(opts, method, params...)
}
