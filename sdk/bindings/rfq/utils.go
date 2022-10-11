package rfq

import (
	"fmt"
	"strings"

	"github.com/celer-network/rfq-mm/sdk/eth"
	"github.com/ethereum/go-ethereum/accounts/abi"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

const (
	EventNameSrcDeposited        = "SrcDeposited"
	EventNameDstTransferred      = "DstTransferred"
	EventNameRefundInitiated     = "RefundInitiated"
	EventNameSrcReleased         = "SrcReleased"
	EventNameRefunded            = "Refunded"
	EventNameRfqContractsUpdated = "RfqContractsUpdated"
	EventNameFeePercUpdated      = "FeePercUpdated"
	EventNameTreasuryAddrUpdated = "TreasuryAddrUpdated"
	EventNameFeeCollected        = "FeeCollected"
)

const (
	QuoteStatusNull uint8 = iota
	QuoteStatusSrcDeposited
	QuoteStatusSrcReleased
	QuoteStatusSrcReleasedNative
	QuoteStatusSrcRefunded
	QuoteStatusSrcRefundedNative
	QuoteStatusDstRefundInitiated
	QuoteStatusDstTransferred
	QuoteStatusDstTransferredNative
)

var quoteStatusName = map[uint8]string{
	0: "QuoteStatusNull",
	1: "QuoteStatusSrcDeposited",
	2: "QuoteStatusSrcReleased",
	3: "QuoteStatusSrcReleasedNative",
	4: "QuoteStatusSrcRefunded",
	5: "QuoteStatusSrcRefundedNative",
	6: "QuoteStatusDstRefundInitiated",
	7: "QuoteStatusDstTransferred",
	8: "QuoteStatusDstTransferredNative",
}

func GetQuoteStatusName(status uint8) string {
	return quoteStatusName[status]
}

func (r *RfqSrcDeposited) PrettyLog(chainId uint64) string {
	return fmt.Sprintf("RfqSrcDeposited, QuoteHash:%s, LiquidityProvider:%s, chainId:%d", eth.Hash(r.QuoteHash).String(), r.Quote.LiquidityProvider.String(), chainId)
}

func (r *RfqDstTransferred) PrettyLog(chainId uint64) string {
	return fmt.Sprintf("RfqDstTransferred, QuoteHash:%s, chainId:%d", eth.Hash(r.QuoteHash).String(), chainId)
}

func (r *RfqRefundInitiated) PrettyLog(chainId uint64) string {
	return fmt.Sprintf("RfqRefundInitiated, QuoteHash:%s, chainId:%d", eth.Hash(r.QuoteHash).String(), chainId)
}

func (r *RfqRefunded) PrettyLog(chainId uint64) string {
	return fmt.Sprintf("RfqRefunded, QuoteHash:%s, chainId:%d", eth.Hash(r.QuoteHash).String(), chainId)
}

func (r *RfqSrcReleased) PrettyLog(chainId uint64) string {
	return fmt.Sprintf("RfqSrcReleased, QuoteHash:%s, amt:%s, LiquidityProvider:%s, srcToken:%s, chainId:%d", eth.Hash(r.QuoteHash).String(), r.Amount.String(), r.LiquidityProvider.String(), r.SrcToken.String(), chainId)
}

func (q *RFQQuote) String() string {
	return fmt.Sprintf("RfqQuote, SrcChainId:%d, SrcToken:%s, SrcAmount:%s, SrcReleaseAmount:%s, DstChainId:%d, DstToken:%s, DstAmount:%s, Deadline:%d, Nonce:%d, Sender:%s, Receiver:%s, RefundTo:%s, LiquidityProvider:%s",
		q.SrcChainId, q.SrcToken.String(), q.SrcAmount.String(), q.SrcReleaseAmount.String(), q.DstChainId, q.DstToken, q.DstAmount.String(), q.Deadline, q.Nonce, q.Sender.String(), q.Receiver.String(), q.RefundTo.String(), q.LiquidityProvider.String())
}

func (q *RFQQuote) Hash() eth.Hash {
	var quoteHash = solsha3.SoliditySHA3(
		[]string{"uint64", "address", "uint256", "uint256", "uint64", "address", "uint256", "uint64", "uint64", "address", "address", "address", "address"},
		q.SrcChainId, q.SrcToken, q.SrcAmount, q.SrcReleaseAmount, q.DstChainId, q.DstToken, q.DstAmount, q.Deadline, q.Nonce, q.Sender, q.Receiver, q.RefundTo, q.LiquidityProvider)
	return eth.Bytes2Hash(quoteHash)
}

func GetEventId(name string) eth.Hash {
	contractAbi, _ := abi.JSON(strings.NewReader(RfqABI))
	return contractAbi.Events[name].ID
}
