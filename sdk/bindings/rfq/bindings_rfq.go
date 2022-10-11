// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rfq

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

// RFQQuote is an auto generated low-level Go binding around an user-defined struct.
type RFQQuote struct {
	SrcChainId        uint64
	SrcToken          common.Address
	SrcAmount         *big.Int
	SrcReleaseAmount  *big.Int
	DstChainId        uint64
	DstToken          common.Address
	DstAmount         *big.Int
	Deadline          uint64
	Nonce             uint64
	Sender            common.Address
	Receiver          common.Address
	RefundTo          common.Address
	LiquidityProvider common.Address
}

// RfqMetaData contains all meta data concerning the Rfq contract.
var RfqMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"quoteHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DstTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"treasuryAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"chainIds\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"uint32[]\",\"name\":\"feePercs\",\"type\":\"uint32[]\"}],\"name\":\"FeePercUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"messageBus\",\"type\":\"address\"}],\"name\":\"MessageBusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"quoteHash\",\"type\":\"bytes32\"}],\"name\":\"RefundInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"quoteHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Refunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"chainIds\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"remoteRfqContracts\",\"type\":\"address[]\"}],\"name\":\"RfqContractsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"quoteHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcReleaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structRFQ.Quote\",\"name\":\"quote\",\"type\":\"tuple\"}],\"name\":\"SrcDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"quoteHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SrcReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"treasuryAddr\",\"type\":\"address\"}],\"name\":\"TreasuryAddrUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"collectFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcReleaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"}],\"internalType\":\"structRFQ.Quote\",\"name\":\"_quote\",\"type\":\"tuple\"}],\"name\":\"dstTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcReleaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"}],\"internalType\":\"structRFQ.Quote\",\"name\":\"_quote\",\"type\":\"tuple\"}],\"name\":\"dstTransferNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_sender\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessage\",\"outputs\":[{\"internalType\":\"enumIMessageReceiverApp.ExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"executeMessage\",\"outputs\":[{\"internalType\":\"enumIMessageReceiverApp.ExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessageWithTransfer\",\"outputs\":[{\"internalType\":\"enumIMessageReceiverApp.ExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessageWithTransferFallback\",\"outputs\":[{\"internalType\":\"enumIMessageReceiverApp.ExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessageWithTransferRefund\",\"outputs\":[{\"internalType\":\"enumIMessageReceiverApp.ExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcReleaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"}],\"internalType\":\"structRFQ.Quote\",\"name\":\"_quote\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_execMsgCallData\",\"type\":\"bytes\"}],\"name\":\"executeRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcReleaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"}],\"internalType\":\"structRFQ.Quote\",\"name\":\"_quote\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_execMsgCallData\",\"type\":\"bytes\"}],\"name\":\"executeRefundNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePercGlobal\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"feePercOverride\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"getMsgFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcReleaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"}],\"internalType\":\"structRFQ.Quote\",\"name\":\"_quote\",\"type\":\"tuple\"}],\"name\":\"getQuoteHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"getRfqFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"governors\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageBus\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeWrap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pausers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"protocolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"quotes\",\"outputs\":[{\"internalType\":\"enumRFQ.QuoteStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"remoteRfqContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcReleaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"}],\"internalType\":\"structRFQ.Quote\",\"name\":\"_quote\",\"type\":\"tuple\"}],\"name\":\"requestRefund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcReleaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"}],\"internalType\":\"structRFQ.Quote\",\"name\":\"_quote\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"_releaseNative\",\"type\":\"bool\"}],\"name\":\"sameChainTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcReleaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"}],\"internalType\":\"structRFQ.Quote\",\"name\":\"_quote\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"_releaseNative\",\"type\":\"bool\"}],\"name\":\"sameChainTransferNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"_chainIds\",\"type\":\"uint64[]\"},{\"internalType\":\"uint32[]\",\"name\":\"_feePercs\",\"type\":\"uint32[]\"}],\"name\":\"setFeePerc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"}],\"name\":\"setMessageBus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"_chainIds\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"_remoteRfqContracts\",\"type\":\"address[]\"}],\"name\":\"setRemoteRfqContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasuryAddr\",\"type\":\"address\"}],\"name\":\"setTreasuryAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcReleaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"}],\"internalType\":\"structRFQ.Quote\",\"name\":\"_quote\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_submissionDeadline\",\"type\":\"uint64\"}],\"name\":\"srcDeposit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcReleaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"}],\"internalType\":\"structRFQ.Quote\",\"name\":\"_quote\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_submissionDeadline\",\"type\":\"uint64\"}],\"name\":\"srcDepositNative\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcReleaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"}],\"internalType\":\"structRFQ.Quote\",\"name\":\"_quote\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_execMsgCallData\",\"type\":\"bytes\"}],\"name\":\"srcRelease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcReleaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"}],\"internalType\":\"structRFQ.Quote\",\"name\":\"_quote\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_execMsgCallData\",\"type\":\"bytes\"}],\"name\":\"srcReleaseNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"unconsumedMsg\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162004c7038038062004c70833981016040819052620000349162000267565b6200003f336200008d565b6001805460ff60a01b191690556200005733620000dd565b6200006233620001a7565b6001600481905580546001600160a01b0319166001600160a01b039290921691909117905562000299565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b03811660009081526002602052604090205460ff16156200014c5760405162461bcd60e51b815260206004820152601960248201527f4163636f756e7420697320616c7265616479207061757365720000000000000060448201526064015b60405180910390fd5b6001600160a01b038116600081815260026020908152604091829020805460ff1916600117905590519182527f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f891015b60405180910390a150565b6001600160a01b03811660009081526003602052604090205460ff1615620002125760405162461bcd60e51b815260206004820152601b60248201527f4163636f756e7420697320616c726561647920676f7665726e6f720000000000604482015260640162000143565b6001600160a01b038116600081815260036020908152604091829020805460ff1916600117905590519182527fdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b591016200019c565b6000602082840312156200027a57600080fd5b81516001600160a01b03811681146200029257600080fd5b9392505050565b6149c780620002a96000396000f3fe6080604052600436106103135760003560e01c806382dc1ec41161019a578063cbac44df116100e1578063e43581b81161008a578063f2fde38b11610064578063f2fde38b146108f6578063f8b839e514610916578063fbe42fea1461093657600080fd5b8063e43581b81461088a578063ed9830bb146108c3578063eecdac88146108d657600080fd5b8063df1f64ef116100bb578063df1f64ef14610815578063e026049c14610845578063e3eece261461085a57600080fd5b8063cbac44df1461079f578063cc47e400146107bf578063d1ce8b45146107f557600080fd5b8063a5ea10cb11610143578063af4ab1381161011d578063af4ab1381461073f578063b62b31e41461075f578063c78e33a21461078c57600080fd5b8063a5ea10cb146106ec578063a7e05b9c146106ff578063ab9341fd1461071f57600080fd5b80639c649fdf116101745780639c649fdf14610699578063a1a227fa146106ac578063a2bdb899146106cc57600080fd5b806382dc1ec4146106465780638456cb59146106665780638da5cb5b1461067b57600080fd5b806346fbf68e1161025e57806369b59e751161020757806379c7efd3116101e157806379c7efd3146106035780637cd2bffc1461055c57806380f51c121461061657600080fd5b806369b59e75146105ae5780636b2c0f55146105ce5780636ef8d66d146105ee57600080fd5b80635ab7afc6116102385780635ab7afc61461055c5780635c975abb1461056f5780636609870d1461058e57600080fd5b806346fbf68e146104d3578063547cad121461051c57806355e9e3d21461053c57600080fd5b806325329eaf116102c05780633e07d1721161029a5780633e07d1721461046b5780633f4ba83a1461049e578063457bfa2f146104b357600080fd5b806325329eaf146103d657806330d9a62a146104135780633c4a25d01461044b57600080fd5b80630bd930b4116102f15780630bd930b4146103755780631000cd9e146103ae5780631e9c5748146103c357600080fd5b8063063ce4e5146103185780630a54aacd146103415780630bcb498214610362575b600080fd5b61032b610326366004613ecb565b610956565b6040516103389190613f84565b60405180910390f35b61035461034f366004613fb0565b6109c2565b604051908152602001610338565b61032b610370366004613fe6565b610bc3565b34801561038157600080fd5b5060095461039990600160a01b900463ffffffff1681565b60405163ffffffff9091168152602001610338565b6103c16103bc366004614055565b610c29565b005b6103c16103d1366004614080565b610e6d565b3480156103e257600080fd5b506104066103f13660046140ba565b60086020526000908152604090205460ff1681565b60405161033891906140d3565b34801561041f57600080fd5b50600954610433906001600160a01b031681565b6040516001600160a01b039091168152602001610338565b34801561045757600080fd5b506103c16104663660046140e7565b6110b2565b34801561047757600080fd5b50610399610486366004614102565b600a6020526000908152604090205463ffffffff1681565b3480156104aa57600080fd5b506103c1611127565b3480156104bf57600080fd5b50600554610433906001600160a01b031681565b3480156104df57600080fd5b5061050c6104ee3660046140e7565b6001600160a01b031660009081526002602052604090205460ff1690565b6040519015158152602001610338565b34801561052857600080fd5b506103c16105373660046140e7565b611190565b34801561054857600080fd5b5061035461055736600461411d565b61124e565b61032b61056a36600461415f565b6112d8565b34801561057b57600080fd5b50600154600160a01b900460ff1661050c565b34801561059a57600080fd5b506103c16105a93660046141ef565b611340565b3480156105ba57600080fd5b506103c16105c93660046140e7565b611511565b3480156105da57600080fd5b506103c16105e93660046140e7565b6115ea565b3480156105fa57600080fd5b506103c161165c565b6103c1610611366004614055565b611665565b34801561062257600080fd5b5061050c6106313660046140e7565b60026020526000908152604090205460ff1681565b34801561065257600080fd5b506103c16106613660046140e7565b61193f565b34801561067257600080fd5b506103c16119b1565b34801561068757600080fd5b506000546001600160a01b0316610433565b61032b6106a7366004614245565b611a18565b3480156106b857600080fd5b50600154610433906001600160a01b031681565b3480156106d857600080fd5b506103c16106e73660046141ef565b611b45565b6103c16106fa366004614055565b611c10565b34801561070b57600080fd5b506103c161071a3660046140e7565b611d17565b34801561072b57600080fd5b506103c161073a3660046142d5565b611dce565b34801561074b57600080fd5b506103c161075a3660046141ef565b612053565b34801561076b57600080fd5b5061035461077a3660046140e7565b600b6020526000908152604090205481565b6103c161079a366004614080565b612185565b3480156107ab57600080fd5b506103c16107ba3660046142d5565b612288565b3480156107cb57600080fd5b506104336107da366004614102565b6006602052600090815260409020546001600160a01b031681565b34801561080157600080fd5b50610354610810366004614055565b61241d565b34801561082157600080fd5b5061050c6108303660046140ba565b60076020526000908152604090205460ff1681565b34801561085157600080fd5b506103c1612511565b34801561086657600080fd5b5061050c6108753660046140e7565b60036020526000908152604090205460ff1681565b34801561089657600080fd5b5061050c6108a53660046140e7565b6001600160a01b031660009081526003602052604090205460ff1690565b6103546108d1366004613fb0565b61251a565b3480156108e257600080fd5b506103c16108f13660046140e7565b612592565b34801561090257600080fd5b506103c16109113660046140e7565b612604565b34801561092257600080fd5b506103c16109313660046141ef565b6126f2565b34801561094257600080fd5b50610354610951366004614341565b6127f2565b6001546000906001600160a01b031633146109b85760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d657373616765206275730000000000000060448201526064015b60405180910390fd5b9695505050505050565b600154600090600160a01b900460ff1615610a125760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016109af565b6005546001600160a01b0316610a6a5760405162461bcd60e51b815260206004820152601860248201527f5266713a206e61746976652077726170206e6f7420736574000000000000000060448201526064016109af565b6005546001600160a01b0316610a8660408501602086016140e7565b6001600160a01b031614610adc5760405162461bcd60e51b815260206004820152601760248201527f5266713a2073726320746f6b656e206d69736d6174636800000000000000000060448201526064016109af565b8260400135341015610b305760405162461bcd60e51b815260206004820152601860248201527f5266713a20696e73756666696369656e7420616d6f756e74000000000000000060448201526064016109af565b6000610b4a8484610b45604083013534614381565b61284f565b9050600560009054906101000a90046001600160a01b03166001600160a01b031663d0e30db085604001356040518263ffffffff1660e01b81526004016000604051808303818588803b158015610ba057600080fd5b505af1158015610bb4573d6000803e3d6000fd5b50939450505050505b92915050565b6001546000906001600160a01b03163314610c205760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d657373616765206275730000000000000060448201526064016109af565b95945050505050565b600154600160a01b900460ff1615610c765760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016109af565b6005546001600160a01b0316610c9260c0830160a084016140e7565b6001600160a01b031614610ce85760405162461bcd60e51b815260206004820152601760248201527f5266713a2064737420746f6b656e206d69736d6174636800000000000000000060448201526064016109af565b8060c00135341015610d3c5760405162461bcd60e51b815260206004820152601860248201527f5266713a20696e73756666696369656e7420616d6f756e74000000000000000060448201526064016109af565b600080610d4883612c9a565b6000828152600860208181526040808420805460ff1916909317909255905193955091935091610d7d91859160019101614398565b60408051601f198184030181528282528051602091820120908301520160408051601f198184030181529190529050610dd182610dbd6020870187614102565b83610dcc60c089013534614381565b612f3d565b610df0610de6610160860161014087016140e7565b8560c00135612f5f565b7fb97bb040c4582b3252c1079bcea2a781f656ef09ceb53be48b2d615c61198bc583610e24610160870161014088016140e7565b610e3460c0880160a089016140e7565b604080519384526001600160a01b03928316602085015291169082015260c086013560608201526080015b60405180910390a150505050565b600154600160a01b900460ff1615610eba5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016109af565b610eca60a0830160808401614102565b67ffffffffffffffff16610ee16020840184614102565b67ffffffffffffffff1614610f385760405162461bcd60e51b815260206004820152601860248201527f5266713a206e6f742073616d6520636861696e2073776170000000000000000060448201526064016109af565b6005546001600160a01b0316610f5460c0840160a085016140e7565b6001600160a01b031614610faa5760405162461bcd60e51b815260206004820152601760248201527f5266713a2064737420746f6b656e206d69736d6174636800000000000000000060448201526064016109af565b8160c00135341015610ffe5760405162461bcd60e51b815260206004820152601860248201527f5266713a20696e73756666696369656e7420616d6f756e74000000000000000060448201526064016109af565b600061100983612c9a565b50905061102b611021610160850161014086016140e7565b8460c00135612f5f565b611036838284613070565b7fb97bb040c4582b3252c1079bcea2a781f656ef09ceb53be48b2d615c61198bc58161106a610160860161014087016140e7565b61107a60c0870160a088016140e7565b604080519384526001600160a01b03928316602085015291169082015260c085013560608201526080015b60405180910390a1505050565b336110c56000546001600160a01b031690565b6001600160a01b03161461111b5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016109af565b611124816131c3565b50565b3360009081526002602052604090205460ff166111865760405162461bcd60e51b815260206004820152601460248201527f43616c6c6572206973206e6f742070617573657200000000000000000000000060448201526064016109af565b61118e613280565b565b336111a36000546001600160a01b031690565b6001600160a01b0316146111f95760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016109af565b600180546001600160a01b0319166001600160a01b0383169081179091556040519081527f3f8223bcd8b3b875473e9f9e14e1ad075451a2b5ffd31591655da9a01516bf5e906020015b60405180910390a150565b60015460405163299aee5160e11b81526000916001600160a01b031690635335dca29061128190869086906004016143b7565b60206040518083038186803b15801561129957600080fd5b505afa1580156112ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112d191906143e6565b9392505050565b6001546000906001600160a01b031633146113355760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d657373616765206275730000000000000060448201526064016109af565b979650505050505050565b600260045414156113935760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016109af565b6002600455600154600160a01b900460ff16156113e55760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016109af565b6005546001600160a01b031661140160408501602086016140e7565b6001600160a01b0316146114575760405162461bcd60e51b815260206004820152601760248201527f5266713a2073726320746f6b656e206d69736d6174636800000000000000000060448201526064016109af565b600080611465858585613326565b6000828152600860205260409020805460ff1916600517905590925090506114a2611498610160870161014088016140e7565b86604001356134bf565b7f2e0668a62a5f556368dca9c7113e20f2852c05155548243804bf714ce72b25a682826114d56040890160208a016140e7565b604080519384526001600160a01b039283166020850152911682820152870135606082015260800160405180910390a150506001600455505050565b6009546001600160a01b03166115695760405162461bcd60e51b815260206004820152601d60248201527f5266713a2074726561737572792061646472657373206e6f742073657400000060448201526064016109af565b6001600160a01b038082166000818152600b602052604081208054919055600954909261159892911683613629565b600954604080516001600160a01b039283168152918416602083015281018290527ff228de527fc1b9843baac03b9a04565473a263375950e63435d4138464386f469060600160405180910390a15050565b336115fd6000546001600160a01b031690565b6001600160a01b0316146116535760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016109af565b611124816136b9565b61118e336136b9565b600154600160a01b900460ff16156116b25760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016109af565b426116c4610100830160e08401614102565b67ffffffffffffffff16106117255760405162461bcd60e51b815260206004820152602160248201527f5266713a207472616e7366657220646561646c696e65206e6f742070617373656044820152601960fa1b60648201526084016109af565b67ffffffffffffffff461661174060a0830160808401614102565b67ffffffffffffffff16146117975760405162461bcd60e51b815260206004820152601960248201527f5266713a2064737420636861696e4964206d69736d617463680000000000000060448201526064016109af565b60006006816117a96020850185614102565b67ffffffffffffffff1681526020810191909152604001600020546001600160a01b031690508061181c5760405162461bcd60e51b815260206004820152601d60248201527f5266713a207372632072667120636f6e7472616374206e6f742073657400000060448201526064016109af565b60006118278361241d565b90506000808281526008602081905260409091205460ff169081111561184f5761184f613f5e565b1461189c5760405162461bcd60e51b815260206004820152601b60248201527f5266713a2071756f746520616c7265616479206578656375746564000000000060448201526064016109af565b6000818152600860209081526040808320805460ff19166006179055516118c891849160029101614398565b60408051601f198184030181528282528051602091820120908301520160408051601f19818403018152919052905061190f836119086020870187614102565b8334612f3d565b6040518281527f7cdd4403cff3a09d96c1ffe4ad1cc5c195e9053463a55edfc2944644ec02211890602001610e5f565b336119526000546001600160a01b031690565b6001600160a01b0316146119a85760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016109af565b61112481613772565b3360009081526002602052604090205460ff16611a105760405162461bcd60e51b815260206004820152601460248201527f43616c6c6572206973206e6f742070617573657200000000000000000000000060448201526064016109af565b61118e61382f565b6001546000906001600160a01b03163314611a755760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d657373616765206275730000000000000060448201526064016109af565b60208314611ac55760405162461bcd60e51b815260206004820152601d60248201527f5266713a20696e636f7272656374206d657373616765206c656e67746800000060448201526064016109af565b67ffffffffffffffff85166000908152600660205260409020546001600160a01b039081169087168114611afd576002915050610c20565b600160076000611b10602082898b6143ff565b611b1991614429565b81526020810191909152604001600020805460ff19169115159190911790555060019695505050505050565b60026004541415611b985760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016109af565b6002600455600154600160a01b900460ff1615611bea5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016109af565b6000611bf78484846138b7565b9050611c0584826000613070565b505060016004555050565b600154600160a01b900460ff1615611c5d5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016109af565b600080611c6983612c9a565b6000828152600860209081526040808320805460ff191660071790555193955091935091611c9c91859160019101614398565b60408051601f198184030181528282528051602091820120908301520160408051601f198184030181529190529050611cdc826119086020870187614102565b610df033611cf2610160870161014088016140e7565b60c08701803590611d069060a08a016140e7565b6001600160a01b0316929190613946565b33611d2a6000546001600160a01b031690565b6001600160a01b031614611d805760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016109af565b600980546001600160a01b0319166001600160a01b0383169081179091556040519081527fb17659014001857e7557191ad74dc9e967b181eaed0895975325e3848debbc4290602001611243565b3360009081526003602052604090205460ff16611e2d5760405162461bcd60e51b815260206004820152601660248201527f43616c6c6572206973206e6f7420676f7665726e6f720000000000000000000060448201526064016109af565b828114611e7c5760405162461bcd60e51b815260206004820152601460248201527f5266713a206c656e677468206d69736d6174636800000000000000000000000060448201526064016109af565b60005b8381101561201d57620f4240838383818110611e9d57611e9d614447565b9050602002016020810190611eb29190614471565b63ffffffff1610611f055760405162461bcd60e51b815260206004820152601d60248201527f5266713a206665652070657263656e7461676520746f6f206c6172676500000060448201526064016109af565b848482818110611f1757611f17614447565b9050602002016020810190611f2c9190614102565b67ffffffffffffffff16611f8657828282818110611f4c57611f4c614447565b9050602002016020810190611f619190614471565b600960146101000a81548163ffffffff021916908363ffffffff16021790555061200b565b828282818110611f9857611f98614447565b9050602002016020810190611fad9190614471565b600a6000878785818110611fc357611fc3614447565b9050602002016020810190611fd89190614102565b67ffffffffffffffff1681526020810191909152604001600020805463ffffffff191663ffffffff929092169190911790555b806120158161448c565b915050611e7f565b507f541df5e570cf10ffe04899eebd9eebebd1c54e2bd4af9f24b23fb4a40c6ea00b84848484604051610e5f94939291906144ef565b600260045414156120a65760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016109af565b6002600455600154600160a01b900460ff16156120f85760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016109af565b6005546001600160a01b031661211460408501602086016140e7565b6001600160a01b03161461216a5760405162461bcd60e51b815260206004820152601760248201527f5266713a2073726320746f6b656e206d69736d6174636800000000000000000060448201526064016109af565b60006121778484846138b7565b9050611c0584826001613070565b600154600160a01b900460ff16156121d25760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016109af565b6121e260a0830160808401614102565b67ffffffffffffffff166121f96020840184614102565b67ffffffffffffffff16146122505760405162461bcd60e51b815260206004820152601860248201527f5266713a206e6f742073616d6520636861696e2073776170000000000000000060448201526064016109af565b600061225b83612c9a565b50905061102b33612274610160860161014087016140e7565b60c08601803590611d069060a089016140e7565b3361229b6000546001600160a01b031690565b6001600160a01b0316146122f15760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016109af565b8281146123405760405162461bcd60e51b815260206004820152601460248201527f5266713a206c656e677468206d69736d6174636800000000000000000000000060448201526064016109af565b60005b838110156123e75782828281811061235d5761235d614447565b905060200201602081019061237291906140e7565b6006600087878581811061238857612388614447565b905060200201602081019061239d9190614102565b67ffffffffffffffff168152602081019190915260400160002080546001600160a01b0319166001600160a01b0392909216919091179055806123df8161448c565b915050612343565b507fb4739c640c5970d8ce88b6c31f3706099efca660e282d47b0a267a0bb572d8b784848484604051610e5f949392919061454f565b600061242c6020830183614102565b61243c60408401602085016140e7565b6040840135606085013561245660a0870160808801614102565b61246660c0880160a089016140e7565b60c088013561247c6101008a0160e08b01614102565b61248e6101208b016101008c01614102565b6124a06101408c016101208d016140e7565b6124b26101608d016101408e016140e7565b6124c46101808e016101608f016140e7565b8d6101800160208101906124d891906140e7565b6040516020016124f49d9c9b9a999897969594939291906145a5565b604051602081830303815290604052805190602001209050919050565b61118e3361397e565b600154600090600160a01b900460ff161561256a5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016109af565b600061257784843461284f565b90506112d1333060408701803590611d069060208a016140e7565b336125a56000546001600160a01b031690565b6001600160a01b0316146125fb5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016109af565b6111248161397e565b336126176000546001600160a01b031690565b6001600160a01b03161461266d5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016109af565b6001600160a01b0381166126e95760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016109af565b61112481613a37565b600260045414156127455760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016109af565b6002600455600154600160a01b900460ff16156127975760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016109af565b6000806127a5858585613326565b6000828152600860205260409020805492945090925060049160ff191660018302179055506114a281604087018035906127e29060208a016140e7565b6001600160a01b03169190613629565b67ffffffffffffffff82166000908152600a602052604081205463ffffffff16806128295750600954600160a01b900463ffffffff165b620f424061283d63ffffffff8316856146d4565b61284791906146f3565b949350505050565b6000428367ffffffffffffffff16116128aa5760405162461bcd60e51b815260206004820152601f60248201527f5266713a207375626d697373696f6e20646561646c696e65207061737365640060448201526064016109af565b60006128be610160860161014087016140e7565b6001600160a01b0316141580156128f0575060006128e46101a0860161018087016140e7565b6001600160a01b031614155b6129625760405162461bcd60e51b815260206004820152602a60248201527f5266713a20696e76616c6964207265636569766572206f72206c69717569646960448201527f747950726f76696465720000000000000000000000000000000000000000000060648201526084016109af565b67ffffffffffffffff461661297a6020860186614102565b67ffffffffffffffff16146129d15760405162461bcd60e51b815260206004820152601960248201527f5266713a2073726320636861696e4964206d69736d617463680000000000000060448201526064016109af565b336129e4610140860161012087016140e7565b6001600160a01b031614612a3a5760405162461bcd60e51b815260206004820152601460248201527f5266713a2073656e646572206d69736d6174636800000000000000000000000060448201526064016109af565b6000612a458561241d565b90506000808281526008602081905260409091205460ff1690811115612a6d57612a6d613f5e565b14612aba5760405162461bcd60e51b815260206004820152601660248201527f5266713a2071756f74652068617368206578697374730000000000000000000060448201526064016109af565b6000612ad9612acf60a0880160808901614102565b87604001356127f2565b9050612aed60608701356040880135614381565b811115612b3c5760405162461bcd60e51b815260206004820152601e60248201527f5266713a20696e73756666696369656e742070726f746f636f6c20666565000060448201526064016109af565b6000828152600860205260409020805460ff19166001179055612b6560a0870160808801614102565b67ffffffffffffffff16612b7c6020880188614102565b67ffffffffffffffff1614612c58576000600681612ba060a08a0160808b01614102565b67ffffffffffffffff1681526020810191909152604001600020546001600160a01b0316905080612c135760405162461bcd60e51b815260206004820152601960248201527f5266713a2064737420636f6e7472616374206e6f74207365740000000000000060448201526064016109af565b600083604051602001612c2891815260200190565b60408051601f198184030181529190529050612c5582612c4e60a08b0160808c01614102565b8389612f3d565b50505b7f3e4de2d1674631d426ae2a89635b421e6d40a31d27681afdf0eed67e81d07bcb8287604051612c89929190614715565b60405180910390a150949350505050565b60008042612caf610100850160e08601614102565b67ffffffffffffffff1611612d065760405162461bcd60e51b815260206004820152601d60248201527f5266713a207472616e7366657220646561646c696e652070617373656400000060448201526064016109af565b67ffffffffffffffff4616612d2160a0850160808601614102565b67ffffffffffffffff1614612d785760405162461bcd60e51b815260206004820152601960248201527f5266713a2064737420636861696e4964206d69736d617463680000000000000060448201526064016109af565b6000612d838461241d565b90506000600681612d976020880188614102565b67ffffffffffffffff1681526020810191909152604001600020546001600160a01b03169050612dcd60a0860160808701614102565b67ffffffffffffffff16612de46020870187614102565b67ffffffffffffffff1614612ebf57600082815260086020819052604082205460ff1690811115612e1757612e17613f5e565b14612e645760405162461bcd60e51b815260206004820152601b60248201527f5266713a2071756f746520616c7265616479206578656375746564000000000060448201526064016109af565b6001600160a01b038116612eba5760405162461bcd60e51b815260206004820152601d60248201527f5266713a206473742072667120636f6e7472616374206e6f742073657400000060448201526064016109af565b612f33565b600160008381526008602081905260409091205460ff1690811115612ee657612ee6613f5e565b14612f335760405162461bcd60e51b815260206004820152601d60248201527f5266713a206e6f206465706f736974206f6e2073616d6520636861696e00000060448201526064016109af565b9094909350915050565b600154612f59908590859085906001600160a01b031685613a87565b50505050565b6005546001600160a01b0316612fb75760405162461bcd60e51b815260206004820152601860248201527f5266713a206e61746976652077726170206e6f7420736574000000000000000060448201526064016109af565b6000826001600160a01b03168261c35090604051600060405180830381858888f193505050503d8060008114613009576040519150601f19603f3d011682016040523d82523d6000602084013e61300e565b606091505b505090508061306b5760405162461bcd60e51b8152602060048201526024808201527f5266713a206661696c656420746f207472616e73666572206e6174697665207460448201526337b5b2b760e11b60648201526084016109af565b505050565b61308260608401356040850135614381565b600b600061309660408701602088016140e7565b6001600160a01b03166001600160a01b0316815260200190815260200160002060008282546130c59190614893565b9091555050801561310d576000828152600860205260409020805460ff191660031790556131086130fe6101a0850161018086016140e7565b84606001356134bf565b613150565b6000828152600860205260409020805460ff1916600217905561315061313b6101a0850161018086016140e7565b60608501356127e260408701602088016140e7565b7ff29b32a17c591b8b3b1216ce0ffb67c07f3478e99b50c5ccf8602878b1ee6376826131846101a0860161018087016140e7565b61319460408701602088016140e7565b604080519384526001600160a01b039283166020850152911690820152606080860135908201526080016110a5565b6001600160a01b03811660009081526003602052604090205460ff161561322c5760405162461bcd60e51b815260206004820152601b60248201527f4163636f756e7420697320616c726561647920676f7665726e6f72000000000060448201526064016109af565b6001600160a01b038116600081815260036020908152604091829020805460ff1916600117905590519182527fdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b59101611243565b600154600160a01b900460ff166132d95760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f742070617573656400000000000000000000000060448201526064016109af565b6001805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b60008060006133348661241d565b9050600160008281526008602081905260409091205460ff169081111561335d5761335d613f5e565b146133aa5760405162461bcd60e51b815260206004820152601960248201527f5266713a20696e636f72726563742071756f746520686173680000000000000060448201526064016109af565b6133ba60a0870160808801614102565b67ffffffffffffffff166133d16020880188614102565b67ffffffffffffffff16146133f2576133ed8585836002613af2565b613465565b42613404610100880160e08901614102565b67ffffffffffffffff16106134655760405162461bcd60e51b815260206004820152602160248201527f5266713a207472616e7366657220646561646c696e65206e6f742070617373656044820152601960fa1b60648201526084016109af565b60008061347a61018089016101608a016140e7565b6001600160a01b03161461349f5761349a610180880161016089016140e7565b6134b1565b6134b1610140880161012089016140e7565b919791965090945050505050565b6005546001600160a01b03166135175760405162461bcd60e51b815260206004820152601860248201527f5266713a206e61746976652077726170206e6f7420736574000000000000000060448201526064016109af565b600554604051632e1a7d4d60e01b8152600481018390526001600160a01b0390911690632e1a7d4d90602401600060405180830381600087803b15801561355d57600080fd5b505af1158015613571573d6000803e3d6000fd5b505050506000826001600160a01b03168261c35090604051600060405180830381858888f193505050503d80600081146135c7576040519150601f19603f3d011682016040523d82523d6000602084013e6135cc565b606091505b505090508061306b5760405162461bcd60e51b8152602060048201526024808201527f5266713a206661696c656420746f207769746864726177206e6174697665207460448201526337b5b2b760e11b60648201526084016109af565b6040516001600160a01b03831660248201526044810182905261306b90849063a9059cbb60e01b906064015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152613c65565b6001600160a01b03811660009081526002602052604090205460ff166137215760405162461bcd60e51b815260206004820152601560248201527f4163636f756e74206973206e6f7420706175736572000000000000000000000060448201526064016109af565b6001600160a01b038116600081815260026020908152604091829020805460ff1916905590519182527fcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e9101611243565b6001600160a01b03811660009081526002602052604090205460ff16156137db5760405162461bcd60e51b815260206004820152601960248201527f4163636f756e7420697320616c7265616479207061757365720000000000000060448201526064016109af565b6001600160a01b038116600081815260026020908152604091829020805460ff1916600117905590519182527f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f89101611243565b600154600160a01b900460ff161561387c5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016109af565b6001805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586133093390565b6000806138c38561241d565b9050600160008281526008602081905260409091205460ff16908111156138ec576138ec613f5e565b146139395760405162461bcd60e51b815260206004820152601960248201527f5266713a20696e636f72726563742071756f746520686173680000000000000060448201526064016109af565b6128478484836001613af2565b6040516001600160a01b0380851660248301528316604482015260648101829052612f599085906323b872dd60e01b90608401613655565b6001600160a01b03811660009081526003602052604090205460ff166139e65760405162461bcd60e51b815260206004820152601760248201527f4163636f756e74206973206e6f7420676f7665726e6f7200000000000000000060448201526064016109af565b6001600160a01b038116600081815260036020908152604091829020805460ff1916905590519182527f1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b9101611243565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b604051634f9e72ad60e11b81526001600160a01b03831690639f3ce55a908390613ab990899089908990600401614903565b6000604051808303818588803b158015613ad257600080fd5b505af1158015613ae6573d6000803e3d6000fd5b50505050505050505050565b60008282604051602001613b07929190614398565b60408051601f1981840301815291815281516020928301206000818152600790935291205490915060ff16613beb576001546040516000916001600160a01b031690613b569088908890614935565b6000604051808303816000865af19150503d8060008114613b93576040519150601f19603f3d011682016040523d82523d6000602084013e613b98565b606091505b5050905080613be95760405162461bcd60e51b815260206004820152601260248201527f65786563757465206d7367206661696c6564000000000000000000000000000060448201526064016109af565b505b60008181526007602052604090205460ff16613c495760405162461bcd60e51b815260206004820152601060248201527f5266713a20696e76616c6964206d73670000000000000000000000000000000060448201526064016109af565b6000908152600760205260409020805460ff1916905550505050565b6000613cba826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316613d4a9092919063ffffffff16565b80519091501561306b5780806020019051810190613cd89190614945565b61306b5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016109af565b60606128478484600085856001600160a01b0385163b613dac5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016109af565b600080866001600160a01b03168587604051613dc89190614962565b60006040518083038185875af1925050503d8060008114613e05576040519150601f19603f3d011682016040523d82523d6000602084013e613e0a565b606091505b509150915061133582828660608315613e245750816112d1565b825115613e345782518084602001fd5b8160405162461bcd60e51b81526004016109af919061497e565b60008083601f840112613e6057600080fd5b50813567ffffffffffffffff811115613e7857600080fd5b602083019150836020828501011115613e9057600080fd5b9250929050565b803567ffffffffffffffff81168114613eaf57600080fd5b919050565b80356001600160a01b0381168114613eaf57600080fd5b60008060008060008060808789031215613ee457600080fd5b863567ffffffffffffffff80821115613efc57600080fd5b613f088a838b01613e4e565b9098509650869150613f1c60208a01613e97565b95506040890135915080821115613f3257600080fd5b50613f3f89828a01613e4e565b9094509250613f52905060608801613eb4565b90509295509295509295565b634e487b7160e01b600052602160045260246000fd5b6003811061112457611124613f5e565b60208101613f9183613f74565b91905290565b60006101a08284031215613faa57600080fd5b50919050565b6000806101c08385031215613fc457600080fd5b613fce8484613f97565b9150613fdd6101a08401613e97565b90509250929050565b600080600080600060808688031215613ffe57600080fd5b61400786613eb4565b945060208601359350604086013567ffffffffffffffff81111561402a57600080fd5b61403688828901613e4e565b9094509250614049905060608701613eb4565b90509295509295909350565b60006101a0828403121561406857600080fd5b6112d18383613f97565b801515811461112457600080fd5b6000806101c0838503121561409457600080fd5b61409e8484613f97565b91506101a08301356140af81614072565b809150509250929050565b6000602082840312156140cc57600080fd5b5035919050565b6020810160098310613f9157613f91613f5e565b6000602082840312156140f957600080fd5b6112d182613eb4565b60006020828403121561411457600080fd5b6112d182613e97565b6000806020838503121561413057600080fd5b823567ffffffffffffffff81111561414757600080fd5b61415385828601613e4e565b90969095509350505050565b600080600080600080600060c0888a03121561417a57600080fd5b61418388613eb4565b965061419160208901613eb4565b9550604088013594506141a660608901613e97565b9350608088013567ffffffffffffffff8111156141c257600080fd5b6141ce8a828b01613e4e565b90945092506141e1905060a08901613eb4565b905092959891949750929550565b60008060006101c0848603121561420557600080fd5b61420f8585613f97565b92506101a084013567ffffffffffffffff81111561422c57600080fd5b61423886828701613e4e565b9497909650939450505050565b60008060008060006080868803121561425d57600080fd5b61426686613eb4565b945061427460208701613e97565b9350604086013567ffffffffffffffff81111561402a57600080fd5b60008083601f8401126142a257600080fd5b50813567ffffffffffffffff8111156142ba57600080fd5b6020830191508360208260051b8501011115613e9057600080fd5b600080600080604085870312156142eb57600080fd5b843567ffffffffffffffff8082111561430357600080fd5b61430f88838901614290565b9096509450602087013591508082111561432857600080fd5b5061433587828801614290565b95989497509550505050565b6000806040838503121561435457600080fd5b61435d83613e97565b946020939093013593505050565b634e487b7160e01b600052601160045260246000fd5b6000828210156143935761439361436b565b500390565b8281526143a482613f74565b60f89190911b6020820152602101919050565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b6000602082840312156143f857600080fd5b5051919050565b6000808585111561440f57600080fd5b8386111561441c57600080fd5b5050820193919092039150565b80356020831015610bbd57600019602084900360031b1b1692915050565b634e487b7160e01b600052603260045260246000fd5b803563ffffffff81168114613eaf57600080fd5b60006020828403121561448357600080fd5b6112d18261445d565b60006000198214156144a0576144a061436b565b5060010190565b8183526000602080850194508260005b858110156144e45767ffffffffffffffff6144d183613e97565b16875295820195908201906001016144b7565b509495945050505050565b6040815260006145036040830186886144a7565b8281036020848101919091528482528591810160005b868110156145425763ffffffff61452f8561445d565b1682529282019290820190600101614519565b5098975050505050505050565b6040815260006145636040830186886144a7565b8281036020848101919091528482528591810160005b86811015614542576001600160a01b0361459285613eb4565b1682529282019290820190600101614579565b6001600160c01b03198e60c01b1681526bffffffffffffffffffffffff198d60601b1660088201528b601c8201528a603c8201526145f2605c82018b60c01b6001600160c01b0319169052565b614610606482018a60601b6bffffffffffffffffffffffff19169052565b87607882015261462f609882018860c01b6001600160c01b0319169052565b61464860a082018760c01b6001600160c01b0319169052565b61466660a882018660601b6bffffffffffffffffffffffff19169052565b61468460bc82018560601b6bffffffffffffffffffffffff19169052565b6146a260d082018460601b6bffffffffffffffffffffffff19169052565b6146c060e482018360601b6bffffffffffffffffffffffff19169052565b60f8019d9c50505050505050505050505050565b60008160001904831182151516156146ee576146ee61436b565b500290565b60008261471057634e487b7160e01b600052601260045260246000fd5b500490565b8281526101c0810161473b6020830161472d85613e97565b67ffffffffffffffff169052565b61474760208401613eb4565b6001600160a01b038116604084015250604083013560608301526060830135608083015261477760808401613e97565b67ffffffffffffffff811660a08401525061479460a08401613eb4565b6001600160a01b03811660c08401525060c083013560e08301526147ba60e08401613e97565b6101006147d28185018367ffffffffffffffff169052565b6147dd818601613e97565b9150506101206147f88185018367ffffffffffffffff169052565b614803818601613eb4565b91505061014061481d818501836001600160a01b03169052565b614828818601613eb4565b915050610160614842818501836001600160a01b03169052565b61484d818601613eb4565b915050610180614867818501836001600160a01b03169052565b614872818601613eb4565b91505061488b6101a08401826001600160a01b03169052565b509392505050565b600082198211156148a6576148a661436b565b500190565b60005b838110156148c65781810151838201526020016148ae565b83811115612f595750506000910152565b600081518084526148ef8160208601602086016148ab565b601f01601f19169290920160200192915050565b6001600160a01b038416815267ffffffffffffffff83166020820152606060408201526000610c2060608301846148d7565b8183823760009101908152919050565b60006020828403121561495757600080fd5b81516112d181614072565b600082516149748184602087016148ab565b9190910192915050565b6020815260006112d160208301846148d756fea2646970667358221220bf993f90d1d6810d2bd770d3417d07f4e9eb75fbaa85b21dafe1bb157df0954764736f6c63430008090033",
}

// RfqABI is the input ABI used to generate the binding from.
// Deprecated: Use RfqMetaData.ABI instead.
var RfqABI = RfqMetaData.ABI

// RfqBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RfqMetaData.Bin instead.
var RfqBin = RfqMetaData.Bin

// DeployRfq deploys a new Ethereum contract, binding an instance of Rfq to it.
func DeployRfq(auth *bind.TransactOpts, backend bind.ContractBackend, _messageBus common.Address) (common.Address, *types.Transaction, *Rfq, error) {
	parsed, err := RfqMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RfqBin), backend, _messageBus)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Rfq{RfqCaller: RfqCaller{contract: contract}, RfqTransactor: RfqTransactor{contract: contract}, RfqFilterer: RfqFilterer{contract: contract}}, nil
}

// Rfq is an auto generated Go binding around an Ethereum contract.
type Rfq struct {
	RfqCaller     // Read-only binding to the contract
	RfqTransactor // Write-only binding to the contract
	RfqFilterer   // Log filterer for contract events
}

// RfqCaller is an auto generated read-only Go binding around an Ethereum contract.
type RfqCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RfqTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RfqTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RfqFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RfqFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RfqSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RfqSession struct {
	Contract     *Rfq              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RfqCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RfqCallerSession struct {
	Contract *RfqCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RfqTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RfqTransactorSession struct {
	Contract     *RfqTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RfqRaw is an auto generated low-level Go binding around an Ethereum contract.
type RfqRaw struct {
	Contract *Rfq // Generic contract binding to access the raw methods on
}

// RfqCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RfqCallerRaw struct {
	Contract *RfqCaller // Generic read-only contract binding to access the raw methods on
}

// RfqTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RfqTransactorRaw struct {
	Contract *RfqTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRfq creates a new instance of Rfq, bound to a specific deployed contract.
func NewRfq(address common.Address, backend bind.ContractBackend) (*Rfq, error) {
	contract, err := bindRfq(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rfq{RfqCaller: RfqCaller{contract: contract}, RfqTransactor: RfqTransactor{contract: contract}, RfqFilterer: RfqFilterer{contract: contract}}, nil
}

// NewRfqCaller creates a new read-only instance of Rfq, bound to a specific deployed contract.
func NewRfqCaller(address common.Address, caller bind.ContractCaller) (*RfqCaller, error) {
	contract, err := bindRfq(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RfqCaller{contract: contract}, nil
}

// NewRfqTransactor creates a new write-only instance of Rfq, bound to a specific deployed contract.
func NewRfqTransactor(address common.Address, transactor bind.ContractTransactor) (*RfqTransactor, error) {
	contract, err := bindRfq(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RfqTransactor{contract: contract}, nil
}

// NewRfqFilterer creates a new log filterer instance of Rfq, bound to a specific deployed contract.
func NewRfqFilterer(address common.Address, filterer bind.ContractFilterer) (*RfqFilterer, error) {
	contract, err := bindRfq(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RfqFilterer{contract: contract}, nil
}

// bindRfq binds a generic wrapper to an already deployed contract.
func bindRfq(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RfqABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rfq *RfqRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rfq.Contract.RfqCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rfq *RfqRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rfq.Contract.RfqTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rfq *RfqRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rfq.Contract.RfqTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rfq *RfqCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rfq.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rfq *RfqTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rfq.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rfq *RfqTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rfq.Contract.contract.Transact(opts, method, params...)
}

// FeePercGlobal is a free data retrieval call binding the contract method 0x0bd930b4.
//
// Solidity: function feePercGlobal() view returns(uint32)
func (_Rfq *RfqCaller) FeePercGlobal(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "feePercGlobal")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// FeePercGlobal is a free data retrieval call binding the contract method 0x0bd930b4.
//
// Solidity: function feePercGlobal() view returns(uint32)
func (_Rfq *RfqSession) FeePercGlobal() (uint32, error) {
	return _Rfq.Contract.FeePercGlobal(&_Rfq.CallOpts)
}

// FeePercGlobal is a free data retrieval call binding the contract method 0x0bd930b4.
//
// Solidity: function feePercGlobal() view returns(uint32)
func (_Rfq *RfqCallerSession) FeePercGlobal() (uint32, error) {
	return _Rfq.Contract.FeePercGlobal(&_Rfq.CallOpts)
}

// FeePercOverride is a free data retrieval call binding the contract method 0x3e07d172.
//
// Solidity: function feePercOverride(uint64 ) view returns(uint32)
func (_Rfq *RfqCaller) FeePercOverride(opts *bind.CallOpts, arg0 uint64) (uint32, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "feePercOverride", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// FeePercOverride is a free data retrieval call binding the contract method 0x3e07d172.
//
// Solidity: function feePercOverride(uint64 ) view returns(uint32)
func (_Rfq *RfqSession) FeePercOverride(arg0 uint64) (uint32, error) {
	return _Rfq.Contract.FeePercOverride(&_Rfq.CallOpts, arg0)
}

// FeePercOverride is a free data retrieval call binding the contract method 0x3e07d172.
//
// Solidity: function feePercOverride(uint64 ) view returns(uint32)
func (_Rfq *RfqCallerSession) FeePercOverride(arg0 uint64) (uint32, error) {
	return _Rfq.Contract.FeePercOverride(&_Rfq.CallOpts, arg0)
}

// GetMsgFee is a free data retrieval call binding the contract method 0x55e9e3d2.
//
// Solidity: function getMsgFee(bytes _message) view returns(uint256)
func (_Rfq *RfqCaller) GetMsgFee(opts *bind.CallOpts, _message []byte) (*big.Int, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "getMsgFee", _message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMsgFee is a free data retrieval call binding the contract method 0x55e9e3d2.
//
// Solidity: function getMsgFee(bytes _message) view returns(uint256)
func (_Rfq *RfqSession) GetMsgFee(_message []byte) (*big.Int, error) {
	return _Rfq.Contract.GetMsgFee(&_Rfq.CallOpts, _message)
}

// GetMsgFee is a free data retrieval call binding the contract method 0x55e9e3d2.
//
// Solidity: function getMsgFee(bytes _message) view returns(uint256)
func (_Rfq *RfqCallerSession) GetMsgFee(_message []byte) (*big.Int, error) {
	return _Rfq.Contract.GetMsgFee(&_Rfq.CallOpts, _message)
}

// GetQuoteHash is a free data retrieval call binding the contract method 0xd1ce8b45.
//
// Solidity: function getQuoteHash((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote) pure returns(bytes32)
func (_Rfq *RfqCaller) GetQuoteHash(opts *bind.CallOpts, _quote RFQQuote) ([32]byte, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "getQuoteHash", _quote)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetQuoteHash is a free data retrieval call binding the contract method 0xd1ce8b45.
//
// Solidity: function getQuoteHash((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote) pure returns(bytes32)
func (_Rfq *RfqSession) GetQuoteHash(_quote RFQQuote) ([32]byte, error) {
	return _Rfq.Contract.GetQuoteHash(&_Rfq.CallOpts, _quote)
}

// GetQuoteHash is a free data retrieval call binding the contract method 0xd1ce8b45.
//
// Solidity: function getQuoteHash((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote) pure returns(bytes32)
func (_Rfq *RfqCallerSession) GetQuoteHash(_quote RFQQuote) ([32]byte, error) {
	return _Rfq.Contract.GetQuoteHash(&_Rfq.CallOpts, _quote)
}

// GetRfqFee is a free data retrieval call binding the contract method 0xfbe42fea.
//
// Solidity: function getRfqFee(uint64 _chainId, uint256 _amount) view returns(uint256)
func (_Rfq *RfqCaller) GetRfqFee(opts *bind.CallOpts, _chainId uint64, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "getRfqFee", _chainId, _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRfqFee is a free data retrieval call binding the contract method 0xfbe42fea.
//
// Solidity: function getRfqFee(uint64 _chainId, uint256 _amount) view returns(uint256)
func (_Rfq *RfqSession) GetRfqFee(_chainId uint64, _amount *big.Int) (*big.Int, error) {
	return _Rfq.Contract.GetRfqFee(&_Rfq.CallOpts, _chainId, _amount)
}

// GetRfqFee is a free data retrieval call binding the contract method 0xfbe42fea.
//
// Solidity: function getRfqFee(uint64 _chainId, uint256 _amount) view returns(uint256)
func (_Rfq *RfqCallerSession) GetRfqFee(_chainId uint64, _amount *big.Int) (*big.Int, error) {
	return _Rfq.Contract.GetRfqFee(&_Rfq.CallOpts, _chainId, _amount)
}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_Rfq *RfqCaller) Governors(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "governors", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_Rfq *RfqSession) Governors(arg0 common.Address) (bool, error) {
	return _Rfq.Contract.Governors(&_Rfq.CallOpts, arg0)
}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_Rfq *RfqCallerSession) Governors(arg0 common.Address) (bool, error) {
	return _Rfq.Contract.Governors(&_Rfq.CallOpts, arg0)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_Rfq *RfqCaller) IsGovernor(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "isGovernor", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_Rfq *RfqSession) IsGovernor(_account common.Address) (bool, error) {
	return _Rfq.Contract.IsGovernor(&_Rfq.CallOpts, _account)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_Rfq *RfqCallerSession) IsGovernor(_account common.Address) (bool, error) {
	return _Rfq.Contract.IsGovernor(&_Rfq.CallOpts, _account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_Rfq *RfqCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "isPauser", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_Rfq *RfqSession) IsPauser(account common.Address) (bool, error) {
	return _Rfq.Contract.IsPauser(&_Rfq.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_Rfq *RfqCallerSession) IsPauser(account common.Address) (bool, error) {
	return _Rfq.Contract.IsPauser(&_Rfq.CallOpts, account)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_Rfq *RfqCaller) MessageBus(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "messageBus")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_Rfq *RfqSession) MessageBus() (common.Address, error) {
	return _Rfq.Contract.MessageBus(&_Rfq.CallOpts)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_Rfq *RfqCallerSession) MessageBus() (common.Address, error) {
	return _Rfq.Contract.MessageBus(&_Rfq.CallOpts)
}

// NativeWrap is a free data retrieval call binding the contract method 0x457bfa2f.
//
// Solidity: function nativeWrap() view returns(address)
func (_Rfq *RfqCaller) NativeWrap(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "nativeWrap")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeWrap is a free data retrieval call binding the contract method 0x457bfa2f.
//
// Solidity: function nativeWrap() view returns(address)
func (_Rfq *RfqSession) NativeWrap() (common.Address, error) {
	return _Rfq.Contract.NativeWrap(&_Rfq.CallOpts)
}

// NativeWrap is a free data retrieval call binding the contract method 0x457bfa2f.
//
// Solidity: function nativeWrap() view returns(address)
func (_Rfq *RfqCallerSession) NativeWrap() (common.Address, error) {
	return _Rfq.Contract.NativeWrap(&_Rfq.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rfq *RfqCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rfq *RfqSession) Owner() (common.Address, error) {
	return _Rfq.Contract.Owner(&_Rfq.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rfq *RfqCallerSession) Owner() (common.Address, error) {
	return _Rfq.Contract.Owner(&_Rfq.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Rfq *RfqCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Rfq *RfqSession) Paused() (bool, error) {
	return _Rfq.Contract.Paused(&_Rfq.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Rfq *RfqCallerSession) Paused() (bool, error) {
	return _Rfq.Contract.Paused(&_Rfq.CallOpts)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_Rfq *RfqCaller) Pausers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "pausers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_Rfq *RfqSession) Pausers(arg0 common.Address) (bool, error) {
	return _Rfq.Contract.Pausers(&_Rfq.CallOpts, arg0)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_Rfq *RfqCallerSession) Pausers(arg0 common.Address) (bool, error) {
	return _Rfq.Contract.Pausers(&_Rfq.CallOpts, arg0)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb62b31e4.
//
// Solidity: function protocolFee(address ) view returns(uint256)
func (_Rfq *RfqCaller) ProtocolFee(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "protocolFee", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFee is a free data retrieval call binding the contract method 0xb62b31e4.
//
// Solidity: function protocolFee(address ) view returns(uint256)
func (_Rfq *RfqSession) ProtocolFee(arg0 common.Address) (*big.Int, error) {
	return _Rfq.Contract.ProtocolFee(&_Rfq.CallOpts, arg0)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb62b31e4.
//
// Solidity: function protocolFee(address ) view returns(uint256)
func (_Rfq *RfqCallerSession) ProtocolFee(arg0 common.Address) (*big.Int, error) {
	return _Rfq.Contract.ProtocolFee(&_Rfq.CallOpts, arg0)
}

// Quotes is a free data retrieval call binding the contract method 0x25329eaf.
//
// Solidity: function quotes(bytes32 ) view returns(uint8)
func (_Rfq *RfqCaller) Quotes(opts *bind.CallOpts, arg0 [32]byte) (uint8, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "quotes", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Quotes is a free data retrieval call binding the contract method 0x25329eaf.
//
// Solidity: function quotes(bytes32 ) view returns(uint8)
func (_Rfq *RfqSession) Quotes(arg0 [32]byte) (uint8, error) {
	return _Rfq.Contract.Quotes(&_Rfq.CallOpts, arg0)
}

// Quotes is a free data retrieval call binding the contract method 0x25329eaf.
//
// Solidity: function quotes(bytes32 ) view returns(uint8)
func (_Rfq *RfqCallerSession) Quotes(arg0 [32]byte) (uint8, error) {
	return _Rfq.Contract.Quotes(&_Rfq.CallOpts, arg0)
}

// RemoteRfqContracts is a free data retrieval call binding the contract method 0xcc47e400.
//
// Solidity: function remoteRfqContracts(uint64 ) view returns(address)
func (_Rfq *RfqCaller) RemoteRfqContracts(opts *bind.CallOpts, arg0 uint64) (common.Address, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "remoteRfqContracts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RemoteRfqContracts is a free data retrieval call binding the contract method 0xcc47e400.
//
// Solidity: function remoteRfqContracts(uint64 ) view returns(address)
func (_Rfq *RfqSession) RemoteRfqContracts(arg0 uint64) (common.Address, error) {
	return _Rfq.Contract.RemoteRfqContracts(&_Rfq.CallOpts, arg0)
}

// RemoteRfqContracts is a free data retrieval call binding the contract method 0xcc47e400.
//
// Solidity: function remoteRfqContracts(uint64 ) view returns(address)
func (_Rfq *RfqCallerSession) RemoteRfqContracts(arg0 uint64) (common.Address, error) {
	return _Rfq.Contract.RemoteRfqContracts(&_Rfq.CallOpts, arg0)
}

// TreasuryAddr is a free data retrieval call binding the contract method 0x30d9a62a.
//
// Solidity: function treasuryAddr() view returns(address)
func (_Rfq *RfqCaller) TreasuryAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "treasuryAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryAddr is a free data retrieval call binding the contract method 0x30d9a62a.
//
// Solidity: function treasuryAddr() view returns(address)
func (_Rfq *RfqSession) TreasuryAddr() (common.Address, error) {
	return _Rfq.Contract.TreasuryAddr(&_Rfq.CallOpts)
}

// TreasuryAddr is a free data retrieval call binding the contract method 0x30d9a62a.
//
// Solidity: function treasuryAddr() view returns(address)
func (_Rfq *RfqCallerSession) TreasuryAddr() (common.Address, error) {
	return _Rfq.Contract.TreasuryAddr(&_Rfq.CallOpts)
}

// UnconsumedMsg is a free data retrieval call binding the contract method 0xdf1f64ef.
//
// Solidity: function unconsumedMsg(bytes32 ) view returns(bool)
func (_Rfq *RfqCaller) UnconsumedMsg(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Rfq.contract.Call(opts, &out, "unconsumedMsg", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UnconsumedMsg is a free data retrieval call binding the contract method 0xdf1f64ef.
//
// Solidity: function unconsumedMsg(bytes32 ) view returns(bool)
func (_Rfq *RfqSession) UnconsumedMsg(arg0 [32]byte) (bool, error) {
	return _Rfq.Contract.UnconsumedMsg(&_Rfq.CallOpts, arg0)
}

// UnconsumedMsg is a free data retrieval call binding the contract method 0xdf1f64ef.
//
// Solidity: function unconsumedMsg(bytes32 ) view returns(bool)
func (_Rfq *RfqCallerSession) UnconsumedMsg(arg0 [32]byte) (bool, error) {
	return _Rfq.Contract.UnconsumedMsg(&_Rfq.CallOpts, arg0)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_Rfq *RfqTransactor) AddGovernor(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "addGovernor", _account)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_Rfq *RfqSession) AddGovernor(_account common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.AddGovernor(&_Rfq.TransactOpts, _account)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_Rfq *RfqTransactorSession) AddGovernor(_account common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.AddGovernor(&_Rfq.TransactOpts, _account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Rfq *RfqTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Rfq *RfqSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.AddPauser(&_Rfq.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Rfq *RfqTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.AddPauser(&_Rfq.TransactOpts, account)
}

// CollectFee is a paid mutator transaction binding the contract method 0x69b59e75.
//
// Solidity: function collectFee(address _token) returns()
func (_Rfq *RfqTransactor) CollectFee(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "collectFee", _token)
}

// CollectFee is a paid mutator transaction binding the contract method 0x69b59e75.
//
// Solidity: function collectFee(address _token) returns()
func (_Rfq *RfqSession) CollectFee(_token common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.CollectFee(&_Rfq.TransactOpts, _token)
}

// CollectFee is a paid mutator transaction binding the contract method 0x69b59e75.
//
// Solidity: function collectFee(address _token) returns()
func (_Rfq *RfqTransactorSession) CollectFee(_token common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.CollectFee(&_Rfq.TransactOpts, _token)
}

// DstTransfer is a paid mutator transaction binding the contract method 0xa5ea10cb.
//
// Solidity: function dstTransfer((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote) payable returns()
func (_Rfq *RfqTransactor) DstTransfer(opts *bind.TransactOpts, _quote RFQQuote) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "dstTransfer", _quote)
}

// DstTransfer is a paid mutator transaction binding the contract method 0xa5ea10cb.
//
// Solidity: function dstTransfer((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote) payable returns()
func (_Rfq *RfqSession) DstTransfer(_quote RFQQuote) (*types.Transaction, error) {
	return _Rfq.Contract.DstTransfer(&_Rfq.TransactOpts, _quote)
}

// DstTransfer is a paid mutator transaction binding the contract method 0xa5ea10cb.
//
// Solidity: function dstTransfer((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote) payable returns()
func (_Rfq *RfqTransactorSession) DstTransfer(_quote RFQQuote) (*types.Transaction, error) {
	return _Rfq.Contract.DstTransfer(&_Rfq.TransactOpts, _quote)
}

// DstTransferNative is a paid mutator transaction binding the contract method 0x1000cd9e.
//
// Solidity: function dstTransferNative((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote) payable returns()
func (_Rfq *RfqTransactor) DstTransferNative(opts *bind.TransactOpts, _quote RFQQuote) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "dstTransferNative", _quote)
}

// DstTransferNative is a paid mutator transaction binding the contract method 0x1000cd9e.
//
// Solidity: function dstTransferNative((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote) payable returns()
func (_Rfq *RfqSession) DstTransferNative(_quote RFQQuote) (*types.Transaction, error) {
	return _Rfq.Contract.DstTransferNative(&_Rfq.TransactOpts, _quote)
}

// DstTransferNative is a paid mutator transaction binding the contract method 0x1000cd9e.
//
// Solidity: function dstTransferNative((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote) payable returns()
func (_Rfq *RfqTransactorSession) DstTransferNative(_quote RFQQuote) (*types.Transaction, error) {
	return _Rfq.Contract.DstTransferNative(&_Rfq.TransactOpts, _quote)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x063ce4e5.
//
// Solidity: function executeMessage(bytes _sender, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_Rfq *RfqTransactor) ExecuteMessage(opts *bind.TransactOpts, _sender []byte, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "executeMessage", _sender, _srcChainId, _message, _executor)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x063ce4e5.
//
// Solidity: function executeMessage(bytes _sender, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_Rfq *RfqSession) ExecuteMessage(_sender []byte, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.ExecuteMessage(&_Rfq.TransactOpts, _sender, _srcChainId, _message, _executor)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x063ce4e5.
//
// Solidity: function executeMessage(bytes _sender, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_Rfq *RfqTransactorSession) ExecuteMessage(_sender []byte, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.ExecuteMessage(&_Rfq.TransactOpts, _sender, _srcChainId, _message, _executor)
}

// ExecuteMessage0 is a paid mutator transaction binding the contract method 0x9c649fdf.
//
// Solidity: function executeMessage(address _sender, uint64 _srcChainId, bytes _message, address ) payable returns(uint8)
func (_Rfq *RfqTransactor) ExecuteMessage0(opts *bind.TransactOpts, _sender common.Address, _srcChainId uint64, _message []byte, arg3 common.Address) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "executeMessage0", _sender, _srcChainId, _message, arg3)
}

// ExecuteMessage0 is a paid mutator transaction binding the contract method 0x9c649fdf.
//
// Solidity: function executeMessage(address _sender, uint64 _srcChainId, bytes _message, address ) payable returns(uint8)
func (_Rfq *RfqSession) ExecuteMessage0(_sender common.Address, _srcChainId uint64, _message []byte, arg3 common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.ExecuteMessage0(&_Rfq.TransactOpts, _sender, _srcChainId, _message, arg3)
}

// ExecuteMessage0 is a paid mutator transaction binding the contract method 0x9c649fdf.
//
// Solidity: function executeMessage(address _sender, uint64 _srcChainId, bytes _message, address ) payable returns(uint8)
func (_Rfq *RfqTransactorSession) ExecuteMessage0(_sender common.Address, _srcChainId uint64, _message []byte, arg3 common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.ExecuteMessage0(&_Rfq.TransactOpts, _sender, _srcChainId, _message, arg3)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x7cd2bffc.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_Rfq *RfqTransactor) ExecuteMessageWithTransfer(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "executeMessageWithTransfer", _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x7cd2bffc.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_Rfq *RfqSession) ExecuteMessageWithTransfer(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.ExecuteMessageWithTransfer(&_Rfq.TransactOpts, _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x7cd2bffc.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_Rfq *RfqTransactorSession) ExecuteMessageWithTransfer(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.ExecuteMessageWithTransfer(&_Rfq.TransactOpts, _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0x5ab7afc6.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_Rfq *RfqTransactor) ExecuteMessageWithTransferFallback(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "executeMessageWithTransferFallback", _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0x5ab7afc6.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_Rfq *RfqSession) ExecuteMessageWithTransferFallback(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.ExecuteMessageWithTransferFallback(&_Rfq.TransactOpts, _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0x5ab7afc6.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_Rfq *RfqTransactorSession) ExecuteMessageWithTransferFallback(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.ExecuteMessageWithTransferFallback(&_Rfq.TransactOpts, _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x0bcb4982.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message, address _executor) payable returns(uint8)
func (_Rfq *RfqTransactor) ExecuteMessageWithTransferRefund(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "executeMessageWithTransferRefund", _token, _amount, _message, _executor)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x0bcb4982.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message, address _executor) payable returns(uint8)
func (_Rfq *RfqSession) ExecuteMessageWithTransferRefund(_token common.Address, _amount *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.ExecuteMessageWithTransferRefund(&_Rfq.TransactOpts, _token, _amount, _message, _executor)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x0bcb4982.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message, address _executor) payable returns(uint8)
func (_Rfq *RfqTransactorSession) ExecuteMessageWithTransferRefund(_token common.Address, _amount *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.ExecuteMessageWithTransferRefund(&_Rfq.TransactOpts, _token, _amount, _message, _executor)
}

// ExecuteRefund is a paid mutator transaction binding the contract method 0xf8b839e5.
//
// Solidity: function executeRefund((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bytes _execMsgCallData) returns()
func (_Rfq *RfqTransactor) ExecuteRefund(opts *bind.TransactOpts, _quote RFQQuote, _execMsgCallData []byte) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "executeRefund", _quote, _execMsgCallData)
}

// ExecuteRefund is a paid mutator transaction binding the contract method 0xf8b839e5.
//
// Solidity: function executeRefund((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bytes _execMsgCallData) returns()
func (_Rfq *RfqSession) ExecuteRefund(_quote RFQQuote, _execMsgCallData []byte) (*types.Transaction, error) {
	return _Rfq.Contract.ExecuteRefund(&_Rfq.TransactOpts, _quote, _execMsgCallData)
}

// ExecuteRefund is a paid mutator transaction binding the contract method 0xf8b839e5.
//
// Solidity: function executeRefund((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bytes _execMsgCallData) returns()
func (_Rfq *RfqTransactorSession) ExecuteRefund(_quote RFQQuote, _execMsgCallData []byte) (*types.Transaction, error) {
	return _Rfq.Contract.ExecuteRefund(&_Rfq.TransactOpts, _quote, _execMsgCallData)
}

// ExecuteRefundNative is a paid mutator transaction binding the contract method 0x6609870d.
//
// Solidity: function executeRefundNative((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bytes _execMsgCallData) returns()
func (_Rfq *RfqTransactor) ExecuteRefundNative(opts *bind.TransactOpts, _quote RFQQuote, _execMsgCallData []byte) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "executeRefundNative", _quote, _execMsgCallData)
}

// ExecuteRefundNative is a paid mutator transaction binding the contract method 0x6609870d.
//
// Solidity: function executeRefundNative((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bytes _execMsgCallData) returns()
func (_Rfq *RfqSession) ExecuteRefundNative(_quote RFQQuote, _execMsgCallData []byte) (*types.Transaction, error) {
	return _Rfq.Contract.ExecuteRefundNative(&_Rfq.TransactOpts, _quote, _execMsgCallData)
}

// ExecuteRefundNative is a paid mutator transaction binding the contract method 0x6609870d.
//
// Solidity: function executeRefundNative((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bytes _execMsgCallData) returns()
func (_Rfq *RfqTransactorSession) ExecuteRefundNative(_quote RFQQuote, _execMsgCallData []byte) (*types.Transaction, error) {
	return _Rfq.Contract.ExecuteRefundNative(&_Rfq.TransactOpts, _quote, _execMsgCallData)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Rfq *RfqTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Rfq *RfqSession) Pause() (*types.Transaction, error) {
	return _Rfq.Contract.Pause(&_Rfq.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Rfq *RfqTransactorSession) Pause() (*types.Transaction, error) {
	return _Rfq.Contract.Pause(&_Rfq.TransactOpts)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_Rfq *RfqTransactor) RemoveGovernor(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "removeGovernor", _account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_Rfq *RfqSession) RemoveGovernor(_account common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.RemoveGovernor(&_Rfq.TransactOpts, _account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_Rfq *RfqTransactorSession) RemoveGovernor(_account common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.RemoveGovernor(&_Rfq.TransactOpts, _account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_Rfq *RfqTransactor) RemovePauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "removePauser", account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_Rfq *RfqSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.RemovePauser(&_Rfq.TransactOpts, account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_Rfq *RfqTransactorSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.RemovePauser(&_Rfq.TransactOpts, account)
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_Rfq *RfqTransactor) RenounceGovernor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "renounceGovernor")
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_Rfq *RfqSession) RenounceGovernor() (*types.Transaction, error) {
	return _Rfq.Contract.RenounceGovernor(&_Rfq.TransactOpts)
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_Rfq *RfqTransactorSession) RenounceGovernor() (*types.Transaction, error) {
	return _Rfq.Contract.RenounceGovernor(&_Rfq.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Rfq *RfqTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Rfq *RfqSession) RenouncePauser() (*types.Transaction, error) {
	return _Rfq.Contract.RenouncePauser(&_Rfq.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Rfq *RfqTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _Rfq.Contract.RenouncePauser(&_Rfq.TransactOpts)
}

// RequestRefund is a paid mutator transaction binding the contract method 0x79c7efd3.
//
// Solidity: function requestRefund((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote) payable returns()
func (_Rfq *RfqTransactor) RequestRefund(opts *bind.TransactOpts, _quote RFQQuote) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "requestRefund", _quote)
}

// RequestRefund is a paid mutator transaction binding the contract method 0x79c7efd3.
//
// Solidity: function requestRefund((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote) payable returns()
func (_Rfq *RfqSession) RequestRefund(_quote RFQQuote) (*types.Transaction, error) {
	return _Rfq.Contract.RequestRefund(&_Rfq.TransactOpts, _quote)
}

// RequestRefund is a paid mutator transaction binding the contract method 0x79c7efd3.
//
// Solidity: function requestRefund((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote) payable returns()
func (_Rfq *RfqTransactorSession) RequestRefund(_quote RFQQuote) (*types.Transaction, error) {
	return _Rfq.Contract.RequestRefund(&_Rfq.TransactOpts, _quote)
}

// SameChainTransfer is a paid mutator transaction binding the contract method 0xc78e33a2.
//
// Solidity: function sameChainTransfer((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bool _releaseNative) payable returns()
func (_Rfq *RfqTransactor) SameChainTransfer(opts *bind.TransactOpts, _quote RFQQuote, _releaseNative bool) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "sameChainTransfer", _quote, _releaseNative)
}

// SameChainTransfer is a paid mutator transaction binding the contract method 0xc78e33a2.
//
// Solidity: function sameChainTransfer((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bool _releaseNative) payable returns()
func (_Rfq *RfqSession) SameChainTransfer(_quote RFQQuote, _releaseNative bool) (*types.Transaction, error) {
	return _Rfq.Contract.SameChainTransfer(&_Rfq.TransactOpts, _quote, _releaseNative)
}

// SameChainTransfer is a paid mutator transaction binding the contract method 0xc78e33a2.
//
// Solidity: function sameChainTransfer((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bool _releaseNative) payable returns()
func (_Rfq *RfqTransactorSession) SameChainTransfer(_quote RFQQuote, _releaseNative bool) (*types.Transaction, error) {
	return _Rfq.Contract.SameChainTransfer(&_Rfq.TransactOpts, _quote, _releaseNative)
}

// SameChainTransferNative is a paid mutator transaction binding the contract method 0x1e9c5748.
//
// Solidity: function sameChainTransferNative((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bool _releaseNative) payable returns()
func (_Rfq *RfqTransactor) SameChainTransferNative(opts *bind.TransactOpts, _quote RFQQuote, _releaseNative bool) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "sameChainTransferNative", _quote, _releaseNative)
}

// SameChainTransferNative is a paid mutator transaction binding the contract method 0x1e9c5748.
//
// Solidity: function sameChainTransferNative((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bool _releaseNative) payable returns()
func (_Rfq *RfqSession) SameChainTransferNative(_quote RFQQuote, _releaseNative bool) (*types.Transaction, error) {
	return _Rfq.Contract.SameChainTransferNative(&_Rfq.TransactOpts, _quote, _releaseNative)
}

// SameChainTransferNative is a paid mutator transaction binding the contract method 0x1e9c5748.
//
// Solidity: function sameChainTransferNative((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bool _releaseNative) payable returns()
func (_Rfq *RfqTransactorSession) SameChainTransferNative(_quote RFQQuote, _releaseNative bool) (*types.Transaction, error) {
	return _Rfq.Contract.SameChainTransferNative(&_Rfq.TransactOpts, _quote, _releaseNative)
}

// SetFeePerc is a paid mutator transaction binding the contract method 0xab9341fd.
//
// Solidity: function setFeePerc(uint64[] _chainIds, uint32[] _feePercs) returns()
func (_Rfq *RfqTransactor) SetFeePerc(opts *bind.TransactOpts, _chainIds []uint64, _feePercs []uint32) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "setFeePerc", _chainIds, _feePercs)
}

// SetFeePerc is a paid mutator transaction binding the contract method 0xab9341fd.
//
// Solidity: function setFeePerc(uint64[] _chainIds, uint32[] _feePercs) returns()
func (_Rfq *RfqSession) SetFeePerc(_chainIds []uint64, _feePercs []uint32) (*types.Transaction, error) {
	return _Rfq.Contract.SetFeePerc(&_Rfq.TransactOpts, _chainIds, _feePercs)
}

// SetFeePerc is a paid mutator transaction binding the contract method 0xab9341fd.
//
// Solidity: function setFeePerc(uint64[] _chainIds, uint32[] _feePercs) returns()
func (_Rfq *RfqTransactorSession) SetFeePerc(_chainIds []uint64, _feePercs []uint32) (*types.Transaction, error) {
	return _Rfq.Contract.SetFeePerc(&_Rfq.TransactOpts, _chainIds, _feePercs)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_Rfq *RfqTransactor) SetMessageBus(opts *bind.TransactOpts, _messageBus common.Address) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "setMessageBus", _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_Rfq *RfqSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.SetMessageBus(&_Rfq.TransactOpts, _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_Rfq *RfqTransactorSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.SetMessageBus(&_Rfq.TransactOpts, _messageBus)
}

// SetRemoteRfqContracts is a paid mutator transaction binding the contract method 0xcbac44df.
//
// Solidity: function setRemoteRfqContracts(uint64[] _chainIds, address[] _remoteRfqContracts) returns()
func (_Rfq *RfqTransactor) SetRemoteRfqContracts(opts *bind.TransactOpts, _chainIds []uint64, _remoteRfqContracts []common.Address) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "setRemoteRfqContracts", _chainIds, _remoteRfqContracts)
}

// SetRemoteRfqContracts is a paid mutator transaction binding the contract method 0xcbac44df.
//
// Solidity: function setRemoteRfqContracts(uint64[] _chainIds, address[] _remoteRfqContracts) returns()
func (_Rfq *RfqSession) SetRemoteRfqContracts(_chainIds []uint64, _remoteRfqContracts []common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.SetRemoteRfqContracts(&_Rfq.TransactOpts, _chainIds, _remoteRfqContracts)
}

// SetRemoteRfqContracts is a paid mutator transaction binding the contract method 0xcbac44df.
//
// Solidity: function setRemoteRfqContracts(uint64[] _chainIds, address[] _remoteRfqContracts) returns()
func (_Rfq *RfqTransactorSession) SetRemoteRfqContracts(_chainIds []uint64, _remoteRfqContracts []common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.SetRemoteRfqContracts(&_Rfq.TransactOpts, _chainIds, _remoteRfqContracts)
}

// SetTreasuryAddr is a paid mutator transaction binding the contract method 0xa7e05b9c.
//
// Solidity: function setTreasuryAddr(address _treasuryAddr) returns()
func (_Rfq *RfqTransactor) SetTreasuryAddr(opts *bind.TransactOpts, _treasuryAddr common.Address) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "setTreasuryAddr", _treasuryAddr)
}

// SetTreasuryAddr is a paid mutator transaction binding the contract method 0xa7e05b9c.
//
// Solidity: function setTreasuryAddr(address _treasuryAddr) returns()
func (_Rfq *RfqSession) SetTreasuryAddr(_treasuryAddr common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.SetTreasuryAddr(&_Rfq.TransactOpts, _treasuryAddr)
}

// SetTreasuryAddr is a paid mutator transaction binding the contract method 0xa7e05b9c.
//
// Solidity: function setTreasuryAddr(address _treasuryAddr) returns()
func (_Rfq *RfqTransactorSession) SetTreasuryAddr(_treasuryAddr common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.SetTreasuryAddr(&_Rfq.TransactOpts, _treasuryAddr)
}

// SrcDeposit is a paid mutator transaction binding the contract method 0xed9830bb.
//
// Solidity: function srcDeposit((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, uint64 _submissionDeadline) payable returns(bytes32)
func (_Rfq *RfqTransactor) SrcDeposit(opts *bind.TransactOpts, _quote RFQQuote, _submissionDeadline uint64) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "srcDeposit", _quote, _submissionDeadline)
}

// SrcDeposit is a paid mutator transaction binding the contract method 0xed9830bb.
//
// Solidity: function srcDeposit((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, uint64 _submissionDeadline) payable returns(bytes32)
func (_Rfq *RfqSession) SrcDeposit(_quote RFQQuote, _submissionDeadline uint64) (*types.Transaction, error) {
	return _Rfq.Contract.SrcDeposit(&_Rfq.TransactOpts, _quote, _submissionDeadline)
}

// SrcDeposit is a paid mutator transaction binding the contract method 0xed9830bb.
//
// Solidity: function srcDeposit((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, uint64 _submissionDeadline) payable returns(bytes32)
func (_Rfq *RfqTransactorSession) SrcDeposit(_quote RFQQuote, _submissionDeadline uint64) (*types.Transaction, error) {
	return _Rfq.Contract.SrcDeposit(&_Rfq.TransactOpts, _quote, _submissionDeadline)
}

// SrcDepositNative is a paid mutator transaction binding the contract method 0x0a54aacd.
//
// Solidity: function srcDepositNative((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, uint64 _submissionDeadline) payable returns(bytes32)
func (_Rfq *RfqTransactor) SrcDepositNative(opts *bind.TransactOpts, _quote RFQQuote, _submissionDeadline uint64) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "srcDepositNative", _quote, _submissionDeadline)
}

// SrcDepositNative is a paid mutator transaction binding the contract method 0x0a54aacd.
//
// Solidity: function srcDepositNative((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, uint64 _submissionDeadline) payable returns(bytes32)
func (_Rfq *RfqSession) SrcDepositNative(_quote RFQQuote, _submissionDeadline uint64) (*types.Transaction, error) {
	return _Rfq.Contract.SrcDepositNative(&_Rfq.TransactOpts, _quote, _submissionDeadline)
}

// SrcDepositNative is a paid mutator transaction binding the contract method 0x0a54aacd.
//
// Solidity: function srcDepositNative((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, uint64 _submissionDeadline) payable returns(bytes32)
func (_Rfq *RfqTransactorSession) SrcDepositNative(_quote RFQQuote, _submissionDeadline uint64) (*types.Transaction, error) {
	return _Rfq.Contract.SrcDepositNative(&_Rfq.TransactOpts, _quote, _submissionDeadline)
}

// SrcRelease is a paid mutator transaction binding the contract method 0xa2bdb899.
//
// Solidity: function srcRelease((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bytes _execMsgCallData) returns()
func (_Rfq *RfqTransactor) SrcRelease(opts *bind.TransactOpts, _quote RFQQuote, _execMsgCallData []byte) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "srcRelease", _quote, _execMsgCallData)
}

// SrcRelease is a paid mutator transaction binding the contract method 0xa2bdb899.
//
// Solidity: function srcRelease((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bytes _execMsgCallData) returns()
func (_Rfq *RfqSession) SrcRelease(_quote RFQQuote, _execMsgCallData []byte) (*types.Transaction, error) {
	return _Rfq.Contract.SrcRelease(&_Rfq.TransactOpts, _quote, _execMsgCallData)
}

// SrcRelease is a paid mutator transaction binding the contract method 0xa2bdb899.
//
// Solidity: function srcRelease((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bytes _execMsgCallData) returns()
func (_Rfq *RfqTransactorSession) SrcRelease(_quote RFQQuote, _execMsgCallData []byte) (*types.Transaction, error) {
	return _Rfq.Contract.SrcRelease(&_Rfq.TransactOpts, _quote, _execMsgCallData)
}

// SrcReleaseNative is a paid mutator transaction binding the contract method 0xaf4ab138.
//
// Solidity: function srcReleaseNative((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bytes _execMsgCallData) returns()
func (_Rfq *RfqTransactor) SrcReleaseNative(opts *bind.TransactOpts, _quote RFQQuote, _execMsgCallData []byte) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "srcReleaseNative", _quote, _execMsgCallData)
}

// SrcReleaseNative is a paid mutator transaction binding the contract method 0xaf4ab138.
//
// Solidity: function srcReleaseNative((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bytes _execMsgCallData) returns()
func (_Rfq *RfqSession) SrcReleaseNative(_quote RFQQuote, _execMsgCallData []byte) (*types.Transaction, error) {
	return _Rfq.Contract.SrcReleaseNative(&_Rfq.TransactOpts, _quote, _execMsgCallData)
}

// SrcReleaseNative is a paid mutator transaction binding the contract method 0xaf4ab138.
//
// Solidity: function srcReleaseNative((uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) _quote, bytes _execMsgCallData) returns()
func (_Rfq *RfqTransactorSession) SrcReleaseNative(_quote RFQQuote, _execMsgCallData []byte) (*types.Transaction, error) {
	return _Rfq.Contract.SrcReleaseNative(&_Rfq.TransactOpts, _quote, _execMsgCallData)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rfq *RfqTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rfq *RfqSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.TransferOwnership(&_Rfq.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rfq *RfqTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Rfq.Contract.TransferOwnership(&_Rfq.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Rfq *RfqTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rfq.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Rfq *RfqSession) Unpause() (*types.Transaction, error) {
	return _Rfq.Contract.Unpause(&_Rfq.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Rfq *RfqTransactorSession) Unpause() (*types.Transaction, error) {
	return _Rfq.Contract.Unpause(&_Rfq.TransactOpts)
}

// RfqDstTransferredIterator is returned from FilterDstTransferred and is used to iterate over the raw logs and unpacked data for DstTransferred events raised by the Rfq contract.
type RfqDstTransferredIterator struct {
	Event *RfqDstTransferred // Event containing the contract specifics and raw log

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
func (it *RfqDstTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqDstTransferred)
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
		it.Event = new(RfqDstTransferred)
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
func (it *RfqDstTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqDstTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqDstTransferred represents a DstTransferred event raised by the Rfq contract.
type RfqDstTransferred struct {
	QuoteHash [32]byte
	Receiver  common.Address
	DstToken  common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDstTransferred is a free log retrieval operation binding the contract event 0xb97bb040c4582b3252c1079bcea2a781f656ef09ceb53be48b2d615c61198bc5.
//
// Solidity: event DstTransferred(bytes32 quoteHash, address receiver, address dstToken, uint256 amount)
func (_Rfq *RfqFilterer) FilterDstTransferred(opts *bind.FilterOpts) (*RfqDstTransferredIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "DstTransferred")
	if err != nil {
		return nil, err
	}
	return &RfqDstTransferredIterator{contract: _Rfq.contract, event: "DstTransferred", logs: logs, sub: sub}, nil
}

// WatchDstTransferred is a free log subscription operation binding the contract event 0xb97bb040c4582b3252c1079bcea2a781f656ef09ceb53be48b2d615c61198bc5.
//
// Solidity: event DstTransferred(bytes32 quoteHash, address receiver, address dstToken, uint256 amount)
func (_Rfq *RfqFilterer) WatchDstTransferred(opts *bind.WatchOpts, sink chan<- *RfqDstTransferred) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "DstTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqDstTransferred)
				if err := _Rfq.contract.UnpackLog(event, "DstTransferred", log); err != nil {
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

// ParseDstTransferred is a log parse operation binding the contract event 0xb97bb040c4582b3252c1079bcea2a781f656ef09ceb53be48b2d615c61198bc5.
//
// Solidity: event DstTransferred(bytes32 quoteHash, address receiver, address dstToken, uint256 amount)
func (_Rfq *RfqFilterer) ParseDstTransferred(log types.Log) (*RfqDstTransferred, error) {
	event := new(RfqDstTransferred)
	if err := _Rfq.contract.UnpackLog(event, "DstTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqFeeCollectedIterator is returned from FilterFeeCollected and is used to iterate over the raw logs and unpacked data for FeeCollected events raised by the Rfq contract.
type RfqFeeCollectedIterator struct {
	Event *RfqFeeCollected // Event containing the contract specifics and raw log

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
func (it *RfqFeeCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqFeeCollected)
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
		it.Event = new(RfqFeeCollected)
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
func (it *RfqFeeCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqFeeCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqFeeCollected represents a FeeCollected event raised by the Rfq contract.
type RfqFeeCollected struct {
	TreasuryAddr common.Address
	Token        common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFeeCollected is a free log retrieval operation binding the contract event 0xf228de527fc1b9843baac03b9a04565473a263375950e63435d4138464386f46.
//
// Solidity: event FeeCollected(address treasuryAddr, address token, uint256 amount)
func (_Rfq *RfqFilterer) FilterFeeCollected(opts *bind.FilterOpts) (*RfqFeeCollectedIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "FeeCollected")
	if err != nil {
		return nil, err
	}
	return &RfqFeeCollectedIterator{contract: _Rfq.contract, event: "FeeCollected", logs: logs, sub: sub}, nil
}

// WatchFeeCollected is a free log subscription operation binding the contract event 0xf228de527fc1b9843baac03b9a04565473a263375950e63435d4138464386f46.
//
// Solidity: event FeeCollected(address treasuryAddr, address token, uint256 amount)
func (_Rfq *RfqFilterer) WatchFeeCollected(opts *bind.WatchOpts, sink chan<- *RfqFeeCollected) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "FeeCollected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqFeeCollected)
				if err := _Rfq.contract.UnpackLog(event, "FeeCollected", log); err != nil {
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

// ParseFeeCollected is a log parse operation binding the contract event 0xf228de527fc1b9843baac03b9a04565473a263375950e63435d4138464386f46.
//
// Solidity: event FeeCollected(address treasuryAddr, address token, uint256 amount)
func (_Rfq *RfqFilterer) ParseFeeCollected(log types.Log) (*RfqFeeCollected, error) {
	event := new(RfqFeeCollected)
	if err := _Rfq.contract.UnpackLog(event, "FeeCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqFeePercUpdatedIterator is returned from FilterFeePercUpdated and is used to iterate over the raw logs and unpacked data for FeePercUpdated events raised by the Rfq contract.
type RfqFeePercUpdatedIterator struct {
	Event *RfqFeePercUpdated // Event containing the contract specifics and raw log

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
func (it *RfqFeePercUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqFeePercUpdated)
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
		it.Event = new(RfqFeePercUpdated)
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
func (it *RfqFeePercUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqFeePercUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqFeePercUpdated represents a FeePercUpdated event raised by the Rfq contract.
type RfqFeePercUpdated struct {
	ChainIds []uint64
	FeePercs []uint32
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFeePercUpdated is a free log retrieval operation binding the contract event 0x541df5e570cf10ffe04899eebd9eebebd1c54e2bd4af9f24b23fb4a40c6ea00b.
//
// Solidity: event FeePercUpdated(uint64[] chainIds, uint32[] feePercs)
func (_Rfq *RfqFilterer) FilterFeePercUpdated(opts *bind.FilterOpts) (*RfqFeePercUpdatedIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "FeePercUpdated")
	if err != nil {
		return nil, err
	}
	return &RfqFeePercUpdatedIterator{contract: _Rfq.contract, event: "FeePercUpdated", logs: logs, sub: sub}, nil
}

// WatchFeePercUpdated is a free log subscription operation binding the contract event 0x541df5e570cf10ffe04899eebd9eebebd1c54e2bd4af9f24b23fb4a40c6ea00b.
//
// Solidity: event FeePercUpdated(uint64[] chainIds, uint32[] feePercs)
func (_Rfq *RfqFilterer) WatchFeePercUpdated(opts *bind.WatchOpts, sink chan<- *RfqFeePercUpdated) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "FeePercUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqFeePercUpdated)
				if err := _Rfq.contract.UnpackLog(event, "FeePercUpdated", log); err != nil {
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

// ParseFeePercUpdated is a log parse operation binding the contract event 0x541df5e570cf10ffe04899eebd9eebebd1c54e2bd4af9f24b23fb4a40c6ea00b.
//
// Solidity: event FeePercUpdated(uint64[] chainIds, uint32[] feePercs)
func (_Rfq *RfqFilterer) ParseFeePercUpdated(log types.Log) (*RfqFeePercUpdated, error) {
	event := new(RfqFeePercUpdated)
	if err := _Rfq.contract.UnpackLog(event, "FeePercUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqGovernorAddedIterator is returned from FilterGovernorAdded and is used to iterate over the raw logs and unpacked data for GovernorAdded events raised by the Rfq contract.
type RfqGovernorAddedIterator struct {
	Event *RfqGovernorAdded // Event containing the contract specifics and raw log

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
func (it *RfqGovernorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqGovernorAdded)
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
		it.Event = new(RfqGovernorAdded)
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
func (it *RfqGovernorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqGovernorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqGovernorAdded represents a GovernorAdded event raised by the Rfq contract.
type RfqGovernorAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGovernorAdded is a free log retrieval operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_Rfq *RfqFilterer) FilterGovernorAdded(opts *bind.FilterOpts) (*RfqGovernorAddedIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "GovernorAdded")
	if err != nil {
		return nil, err
	}
	return &RfqGovernorAddedIterator{contract: _Rfq.contract, event: "GovernorAdded", logs: logs, sub: sub}, nil
}

// WatchGovernorAdded is a free log subscription operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_Rfq *RfqFilterer) WatchGovernorAdded(opts *bind.WatchOpts, sink chan<- *RfqGovernorAdded) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "GovernorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqGovernorAdded)
				if err := _Rfq.contract.UnpackLog(event, "GovernorAdded", log); err != nil {
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

// ParseGovernorAdded is a log parse operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_Rfq *RfqFilterer) ParseGovernorAdded(log types.Log) (*RfqGovernorAdded, error) {
	event := new(RfqGovernorAdded)
	if err := _Rfq.contract.UnpackLog(event, "GovernorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqGovernorRemovedIterator is returned from FilterGovernorRemoved and is used to iterate over the raw logs and unpacked data for GovernorRemoved events raised by the Rfq contract.
type RfqGovernorRemovedIterator struct {
	Event *RfqGovernorRemoved // Event containing the contract specifics and raw log

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
func (it *RfqGovernorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqGovernorRemoved)
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
		it.Event = new(RfqGovernorRemoved)
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
func (it *RfqGovernorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqGovernorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqGovernorRemoved represents a GovernorRemoved event raised by the Rfq contract.
type RfqGovernorRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGovernorRemoved is a free log retrieval operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_Rfq *RfqFilterer) FilterGovernorRemoved(opts *bind.FilterOpts) (*RfqGovernorRemovedIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "GovernorRemoved")
	if err != nil {
		return nil, err
	}
	return &RfqGovernorRemovedIterator{contract: _Rfq.contract, event: "GovernorRemoved", logs: logs, sub: sub}, nil
}

// WatchGovernorRemoved is a free log subscription operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_Rfq *RfqFilterer) WatchGovernorRemoved(opts *bind.WatchOpts, sink chan<- *RfqGovernorRemoved) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "GovernorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqGovernorRemoved)
				if err := _Rfq.contract.UnpackLog(event, "GovernorRemoved", log); err != nil {
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

// ParseGovernorRemoved is a log parse operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_Rfq *RfqFilterer) ParseGovernorRemoved(log types.Log) (*RfqGovernorRemoved, error) {
	event := new(RfqGovernorRemoved)
	if err := _Rfq.contract.UnpackLog(event, "GovernorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqMessageBusUpdatedIterator is returned from FilterMessageBusUpdated and is used to iterate over the raw logs and unpacked data for MessageBusUpdated events raised by the Rfq contract.
type RfqMessageBusUpdatedIterator struct {
	Event *RfqMessageBusUpdated // Event containing the contract specifics and raw log

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
func (it *RfqMessageBusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqMessageBusUpdated)
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
		it.Event = new(RfqMessageBusUpdated)
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
func (it *RfqMessageBusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqMessageBusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqMessageBusUpdated represents a MessageBusUpdated event raised by the Rfq contract.
type RfqMessageBusUpdated struct {
	MessageBus common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMessageBusUpdated is a free log retrieval operation binding the contract event 0x3f8223bcd8b3b875473e9f9e14e1ad075451a2b5ffd31591655da9a01516bf5e.
//
// Solidity: event MessageBusUpdated(address messageBus)
func (_Rfq *RfqFilterer) FilterMessageBusUpdated(opts *bind.FilterOpts) (*RfqMessageBusUpdatedIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "MessageBusUpdated")
	if err != nil {
		return nil, err
	}
	return &RfqMessageBusUpdatedIterator{contract: _Rfq.contract, event: "MessageBusUpdated", logs: logs, sub: sub}, nil
}

// WatchMessageBusUpdated is a free log subscription operation binding the contract event 0x3f8223bcd8b3b875473e9f9e14e1ad075451a2b5ffd31591655da9a01516bf5e.
//
// Solidity: event MessageBusUpdated(address messageBus)
func (_Rfq *RfqFilterer) WatchMessageBusUpdated(opts *bind.WatchOpts, sink chan<- *RfqMessageBusUpdated) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "MessageBusUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqMessageBusUpdated)
				if err := _Rfq.contract.UnpackLog(event, "MessageBusUpdated", log); err != nil {
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

// ParseMessageBusUpdated is a log parse operation binding the contract event 0x3f8223bcd8b3b875473e9f9e14e1ad075451a2b5ffd31591655da9a01516bf5e.
//
// Solidity: event MessageBusUpdated(address messageBus)
func (_Rfq *RfqFilterer) ParseMessageBusUpdated(log types.Log) (*RfqMessageBusUpdated, error) {
	event := new(RfqMessageBusUpdated)
	if err := _Rfq.contract.UnpackLog(event, "MessageBusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Rfq contract.
type RfqOwnershipTransferredIterator struct {
	Event *RfqOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RfqOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqOwnershipTransferred)
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
		it.Event = new(RfqOwnershipTransferred)
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
func (it *RfqOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqOwnershipTransferred represents a OwnershipTransferred event raised by the Rfq contract.
type RfqOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rfq *RfqFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RfqOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RfqOwnershipTransferredIterator{contract: _Rfq.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rfq *RfqFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RfqOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqOwnershipTransferred)
				if err := _Rfq.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rfq *RfqFilterer) ParseOwnershipTransferred(log types.Log) (*RfqOwnershipTransferred, error) {
	event := new(RfqOwnershipTransferred)
	if err := _Rfq.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Rfq contract.
type RfqPausedIterator struct {
	Event *RfqPaused // Event containing the contract specifics and raw log

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
func (it *RfqPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqPaused)
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
		it.Event = new(RfqPaused)
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
func (it *RfqPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqPaused represents a Paused event raised by the Rfq contract.
type RfqPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Rfq *RfqFilterer) FilterPaused(opts *bind.FilterOpts) (*RfqPausedIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &RfqPausedIterator{contract: _Rfq.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Rfq *RfqFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *RfqPaused) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqPaused)
				if err := _Rfq.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Rfq *RfqFilterer) ParsePaused(log types.Log) (*RfqPaused, error) {
	event := new(RfqPaused)
	if err := _Rfq.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqPauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the Rfq contract.
type RfqPauserAddedIterator struct {
	Event *RfqPauserAdded // Event containing the contract specifics and raw log

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
func (it *RfqPauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqPauserAdded)
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
		it.Event = new(RfqPauserAdded)
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
func (it *RfqPauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqPauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqPauserAdded represents a PauserAdded event raised by the Rfq contract.
type RfqPauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_Rfq *RfqFilterer) FilterPauserAdded(opts *bind.FilterOpts) (*RfqPauserAddedIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return &RfqPauserAddedIterator{contract: _Rfq.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_Rfq *RfqFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *RfqPauserAdded) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqPauserAdded)
				if err := _Rfq.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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

// ParsePauserAdded is a log parse operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_Rfq *RfqFilterer) ParsePauserAdded(log types.Log) (*RfqPauserAdded, error) {
	event := new(RfqPauserAdded)
	if err := _Rfq.contract.UnpackLog(event, "PauserAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqPauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the Rfq contract.
type RfqPauserRemovedIterator struct {
	Event *RfqPauserRemoved // Event containing the contract specifics and raw log

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
func (it *RfqPauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqPauserRemoved)
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
		it.Event = new(RfqPauserRemoved)
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
func (it *RfqPauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqPauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqPauserRemoved represents a PauserRemoved event raised by the Rfq contract.
type RfqPauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_Rfq *RfqFilterer) FilterPauserRemoved(opts *bind.FilterOpts) (*RfqPauserRemovedIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return &RfqPauserRemovedIterator{contract: _Rfq.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_Rfq *RfqFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *RfqPauserRemoved) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqPauserRemoved)
				if err := _Rfq.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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

// ParsePauserRemoved is a log parse operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_Rfq *RfqFilterer) ParsePauserRemoved(log types.Log) (*RfqPauserRemoved, error) {
	event := new(RfqPauserRemoved)
	if err := _Rfq.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqRefundInitiatedIterator is returned from FilterRefundInitiated and is used to iterate over the raw logs and unpacked data for RefundInitiated events raised by the Rfq contract.
type RfqRefundInitiatedIterator struct {
	Event *RfqRefundInitiated // Event containing the contract specifics and raw log

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
func (it *RfqRefundInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqRefundInitiated)
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
		it.Event = new(RfqRefundInitiated)
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
func (it *RfqRefundInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqRefundInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqRefundInitiated represents a RefundInitiated event raised by the Rfq contract.
type RfqRefundInitiated struct {
	QuoteHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRefundInitiated is a free log retrieval operation binding the contract event 0x7cdd4403cff3a09d96c1ffe4ad1cc5c195e9053463a55edfc2944644ec022118.
//
// Solidity: event RefundInitiated(bytes32 quoteHash)
func (_Rfq *RfqFilterer) FilterRefundInitiated(opts *bind.FilterOpts) (*RfqRefundInitiatedIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "RefundInitiated")
	if err != nil {
		return nil, err
	}
	return &RfqRefundInitiatedIterator{contract: _Rfq.contract, event: "RefundInitiated", logs: logs, sub: sub}, nil
}

// WatchRefundInitiated is a free log subscription operation binding the contract event 0x7cdd4403cff3a09d96c1ffe4ad1cc5c195e9053463a55edfc2944644ec022118.
//
// Solidity: event RefundInitiated(bytes32 quoteHash)
func (_Rfq *RfqFilterer) WatchRefundInitiated(opts *bind.WatchOpts, sink chan<- *RfqRefundInitiated) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "RefundInitiated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqRefundInitiated)
				if err := _Rfq.contract.UnpackLog(event, "RefundInitiated", log); err != nil {
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

// ParseRefundInitiated is a log parse operation binding the contract event 0x7cdd4403cff3a09d96c1ffe4ad1cc5c195e9053463a55edfc2944644ec022118.
//
// Solidity: event RefundInitiated(bytes32 quoteHash)
func (_Rfq *RfqFilterer) ParseRefundInitiated(log types.Log) (*RfqRefundInitiated, error) {
	event := new(RfqRefundInitiated)
	if err := _Rfq.contract.UnpackLog(event, "RefundInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqRefundedIterator is returned from FilterRefunded and is used to iterate over the raw logs and unpacked data for Refunded events raised by the Rfq contract.
type RfqRefundedIterator struct {
	Event *RfqRefunded // Event containing the contract specifics and raw log

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
func (it *RfqRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqRefunded)
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
		it.Event = new(RfqRefunded)
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
func (it *RfqRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqRefunded represents a Refunded event raised by the Rfq contract.
type RfqRefunded struct {
	QuoteHash [32]byte
	RefundTo  common.Address
	SrcToken  common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRefunded is a free log retrieval operation binding the contract event 0x2e0668a62a5f556368dca9c7113e20f2852c05155548243804bf714ce72b25a6.
//
// Solidity: event Refunded(bytes32 quoteHash, address refundTo, address srcToken, uint256 amount)
func (_Rfq *RfqFilterer) FilterRefunded(opts *bind.FilterOpts) (*RfqRefundedIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "Refunded")
	if err != nil {
		return nil, err
	}
	return &RfqRefundedIterator{contract: _Rfq.contract, event: "Refunded", logs: logs, sub: sub}, nil
}

// WatchRefunded is a free log subscription operation binding the contract event 0x2e0668a62a5f556368dca9c7113e20f2852c05155548243804bf714ce72b25a6.
//
// Solidity: event Refunded(bytes32 quoteHash, address refundTo, address srcToken, uint256 amount)
func (_Rfq *RfqFilterer) WatchRefunded(opts *bind.WatchOpts, sink chan<- *RfqRefunded) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "Refunded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqRefunded)
				if err := _Rfq.contract.UnpackLog(event, "Refunded", log); err != nil {
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

// ParseRefunded is a log parse operation binding the contract event 0x2e0668a62a5f556368dca9c7113e20f2852c05155548243804bf714ce72b25a6.
//
// Solidity: event Refunded(bytes32 quoteHash, address refundTo, address srcToken, uint256 amount)
func (_Rfq *RfqFilterer) ParseRefunded(log types.Log) (*RfqRefunded, error) {
	event := new(RfqRefunded)
	if err := _Rfq.contract.UnpackLog(event, "Refunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqRfqContractsUpdatedIterator is returned from FilterRfqContractsUpdated and is used to iterate over the raw logs and unpacked data for RfqContractsUpdated events raised by the Rfq contract.
type RfqRfqContractsUpdatedIterator struct {
	Event *RfqRfqContractsUpdated // Event containing the contract specifics and raw log

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
func (it *RfqRfqContractsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqRfqContractsUpdated)
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
		it.Event = new(RfqRfqContractsUpdated)
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
func (it *RfqRfqContractsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqRfqContractsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqRfqContractsUpdated represents a RfqContractsUpdated event raised by the Rfq contract.
type RfqRfqContractsUpdated struct {
	ChainIds           []uint64
	RemoteRfqContracts []common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRfqContractsUpdated is a free log retrieval operation binding the contract event 0xb4739c640c5970d8ce88b6c31f3706099efca660e282d47b0a267a0bb572d8b7.
//
// Solidity: event RfqContractsUpdated(uint64[] chainIds, address[] remoteRfqContracts)
func (_Rfq *RfqFilterer) FilterRfqContractsUpdated(opts *bind.FilterOpts) (*RfqRfqContractsUpdatedIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "RfqContractsUpdated")
	if err != nil {
		return nil, err
	}
	return &RfqRfqContractsUpdatedIterator{contract: _Rfq.contract, event: "RfqContractsUpdated", logs: logs, sub: sub}, nil
}

// WatchRfqContractsUpdated is a free log subscription operation binding the contract event 0xb4739c640c5970d8ce88b6c31f3706099efca660e282d47b0a267a0bb572d8b7.
//
// Solidity: event RfqContractsUpdated(uint64[] chainIds, address[] remoteRfqContracts)
func (_Rfq *RfqFilterer) WatchRfqContractsUpdated(opts *bind.WatchOpts, sink chan<- *RfqRfqContractsUpdated) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "RfqContractsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqRfqContractsUpdated)
				if err := _Rfq.contract.UnpackLog(event, "RfqContractsUpdated", log); err != nil {
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

// ParseRfqContractsUpdated is a log parse operation binding the contract event 0xb4739c640c5970d8ce88b6c31f3706099efca660e282d47b0a267a0bb572d8b7.
//
// Solidity: event RfqContractsUpdated(uint64[] chainIds, address[] remoteRfqContracts)
func (_Rfq *RfqFilterer) ParseRfqContractsUpdated(log types.Log) (*RfqRfqContractsUpdated, error) {
	event := new(RfqRfqContractsUpdated)
	if err := _Rfq.contract.UnpackLog(event, "RfqContractsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqSrcDepositedIterator is returned from FilterSrcDeposited and is used to iterate over the raw logs and unpacked data for SrcDeposited events raised by the Rfq contract.
type RfqSrcDepositedIterator struct {
	Event *RfqSrcDeposited // Event containing the contract specifics and raw log

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
func (it *RfqSrcDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqSrcDeposited)
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
		it.Event = new(RfqSrcDeposited)
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
func (it *RfqSrcDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqSrcDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqSrcDeposited represents a SrcDeposited event raised by the Rfq contract.
type RfqSrcDeposited struct {
	QuoteHash [32]byte
	Quote     RFQQuote
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSrcDeposited is a free log retrieval operation binding the contract event 0x3e4de2d1674631d426ae2a89635b421e6d40a31d27681afdf0eed67e81d07bcb.
//
// Solidity: event SrcDeposited(bytes32 quoteHash, (uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) quote)
func (_Rfq *RfqFilterer) FilterSrcDeposited(opts *bind.FilterOpts) (*RfqSrcDepositedIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "SrcDeposited")
	if err != nil {
		return nil, err
	}
	return &RfqSrcDepositedIterator{contract: _Rfq.contract, event: "SrcDeposited", logs: logs, sub: sub}, nil
}

// WatchSrcDeposited is a free log subscription operation binding the contract event 0x3e4de2d1674631d426ae2a89635b421e6d40a31d27681afdf0eed67e81d07bcb.
//
// Solidity: event SrcDeposited(bytes32 quoteHash, (uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) quote)
func (_Rfq *RfqFilterer) WatchSrcDeposited(opts *bind.WatchOpts, sink chan<- *RfqSrcDeposited) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "SrcDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqSrcDeposited)
				if err := _Rfq.contract.UnpackLog(event, "SrcDeposited", log); err != nil {
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

// ParseSrcDeposited is a log parse operation binding the contract event 0x3e4de2d1674631d426ae2a89635b421e6d40a31d27681afdf0eed67e81d07bcb.
//
// Solidity: event SrcDeposited(bytes32 quoteHash, (uint64,address,uint256,uint256,uint64,address,uint256,uint64,uint64,address,address,address,address) quote)
func (_Rfq *RfqFilterer) ParseSrcDeposited(log types.Log) (*RfqSrcDeposited, error) {
	event := new(RfqSrcDeposited)
	if err := _Rfq.contract.UnpackLog(event, "SrcDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqSrcReleasedIterator is returned from FilterSrcReleased and is used to iterate over the raw logs and unpacked data for SrcReleased events raised by the Rfq contract.
type RfqSrcReleasedIterator struct {
	Event *RfqSrcReleased // Event containing the contract specifics and raw log

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
func (it *RfqSrcReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqSrcReleased)
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
		it.Event = new(RfqSrcReleased)
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
func (it *RfqSrcReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqSrcReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqSrcReleased represents a SrcReleased event raised by the Rfq contract.
type RfqSrcReleased struct {
	QuoteHash         [32]byte
	LiquidityProvider common.Address
	SrcToken          common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSrcReleased is a free log retrieval operation binding the contract event 0xf29b32a17c591b8b3b1216ce0ffb67c07f3478e99b50c5ccf8602878b1ee6376.
//
// Solidity: event SrcReleased(bytes32 quoteHash, address liquidityProvider, address srcToken, uint256 amount)
func (_Rfq *RfqFilterer) FilterSrcReleased(opts *bind.FilterOpts) (*RfqSrcReleasedIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "SrcReleased")
	if err != nil {
		return nil, err
	}
	return &RfqSrcReleasedIterator{contract: _Rfq.contract, event: "SrcReleased", logs: logs, sub: sub}, nil
}

// WatchSrcReleased is a free log subscription operation binding the contract event 0xf29b32a17c591b8b3b1216ce0ffb67c07f3478e99b50c5ccf8602878b1ee6376.
//
// Solidity: event SrcReleased(bytes32 quoteHash, address liquidityProvider, address srcToken, uint256 amount)
func (_Rfq *RfqFilterer) WatchSrcReleased(opts *bind.WatchOpts, sink chan<- *RfqSrcReleased) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "SrcReleased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqSrcReleased)
				if err := _Rfq.contract.UnpackLog(event, "SrcReleased", log); err != nil {
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

// ParseSrcReleased is a log parse operation binding the contract event 0xf29b32a17c591b8b3b1216ce0ffb67c07f3478e99b50c5ccf8602878b1ee6376.
//
// Solidity: event SrcReleased(bytes32 quoteHash, address liquidityProvider, address srcToken, uint256 amount)
func (_Rfq *RfqFilterer) ParseSrcReleased(log types.Log) (*RfqSrcReleased, error) {
	event := new(RfqSrcReleased)
	if err := _Rfq.contract.UnpackLog(event, "SrcReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqTreasuryAddrUpdatedIterator is returned from FilterTreasuryAddrUpdated and is used to iterate over the raw logs and unpacked data for TreasuryAddrUpdated events raised by the Rfq contract.
type RfqTreasuryAddrUpdatedIterator struct {
	Event *RfqTreasuryAddrUpdated // Event containing the contract specifics and raw log

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
func (it *RfqTreasuryAddrUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqTreasuryAddrUpdated)
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
		it.Event = new(RfqTreasuryAddrUpdated)
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
func (it *RfqTreasuryAddrUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqTreasuryAddrUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqTreasuryAddrUpdated represents a TreasuryAddrUpdated event raised by the Rfq contract.
type RfqTreasuryAddrUpdated struct {
	TreasuryAddr common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTreasuryAddrUpdated is a free log retrieval operation binding the contract event 0xb17659014001857e7557191ad74dc9e967b181eaed0895975325e3848debbc42.
//
// Solidity: event TreasuryAddrUpdated(address treasuryAddr)
func (_Rfq *RfqFilterer) FilterTreasuryAddrUpdated(opts *bind.FilterOpts) (*RfqTreasuryAddrUpdatedIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "TreasuryAddrUpdated")
	if err != nil {
		return nil, err
	}
	return &RfqTreasuryAddrUpdatedIterator{contract: _Rfq.contract, event: "TreasuryAddrUpdated", logs: logs, sub: sub}, nil
}

// WatchTreasuryAddrUpdated is a free log subscription operation binding the contract event 0xb17659014001857e7557191ad74dc9e967b181eaed0895975325e3848debbc42.
//
// Solidity: event TreasuryAddrUpdated(address treasuryAddr)
func (_Rfq *RfqFilterer) WatchTreasuryAddrUpdated(opts *bind.WatchOpts, sink chan<- *RfqTreasuryAddrUpdated) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "TreasuryAddrUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqTreasuryAddrUpdated)
				if err := _Rfq.contract.UnpackLog(event, "TreasuryAddrUpdated", log); err != nil {
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

// ParseTreasuryAddrUpdated is a log parse operation binding the contract event 0xb17659014001857e7557191ad74dc9e967b181eaed0895975325e3848debbc42.
//
// Solidity: event TreasuryAddrUpdated(address treasuryAddr)
func (_Rfq *RfqFilterer) ParseTreasuryAddrUpdated(log types.Log) (*RfqTreasuryAddrUpdated, error) {
	event := new(RfqTreasuryAddrUpdated)
	if err := _Rfq.contract.UnpackLog(event, "TreasuryAddrUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RfqUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Rfq contract.
type RfqUnpausedIterator struct {
	Event *RfqUnpaused // Event containing the contract specifics and raw log

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
func (it *RfqUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RfqUnpaused)
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
		it.Event = new(RfqUnpaused)
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
func (it *RfqUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RfqUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RfqUnpaused represents a Unpaused event raised by the Rfq contract.
type RfqUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Rfq *RfqFilterer) FilterUnpaused(opts *bind.FilterOpts) (*RfqUnpausedIterator, error) {

	logs, sub, err := _Rfq.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &RfqUnpausedIterator{contract: _Rfq.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Rfq *RfqFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *RfqUnpaused) (event.Subscription, error) {

	logs, sub, err := _Rfq.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RfqUnpaused)
				if err := _Rfq.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Rfq *RfqFilterer) ParseUnpaused(log types.Log) (*RfqUnpaused, error) {
	event := new(RfqUnpaused)
	if err := _Rfq.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
