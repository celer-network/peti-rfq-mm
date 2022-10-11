package proto

import (
	"math/big"

	"github.com/celer-network/rfq-mm/sdk/bindings/rfq"
	"github.com/celer-network/rfq-mm/sdk/eth"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func (p *Price) GetSrcAmt() *big.Int {
	srcAmt, _ := big.NewInt(0).SetString(p.GetSrcAmount(), 10)
	return srcAmt
}

func (p *Price) GetSrcReleaseAmt() *big.Int {
	releaseAmt, _ := big.NewInt(0).SetString(p.GetSrcReleaseAmount(), 10)
	return releaseAmt
}

func (p *Price) GetDstAmt() *big.Int {
	dstAmt, _ := big.NewInt(0).SetString(p.GetDstAmount(), 10)
	return dstAmt
}

func (p *Price) GetSigBytes() []byte {
	return eth.Hex2Bytes(p.GetSig())
}

func (q *Price) GetMMAddr() eth.Addr {
	return eth.Hex2Addr(q.GetMmAddr())
}

func (p *Price) GetSrcChainId() uint64 {
	return p.GetSrcToken().GetChainId()
}

func (p *Price) GetDstChainId() uint64 {
	return p.GetDstToken().GetChainId()
}

func (p *Price) EncodeSignData() []byte {
	mmAddr := eth.Hex2Addr(p.GetMmAddr())
	srcToken := eth.Hex2Addr(p.GetSrcToken().GetAddress())
	dstToken := eth.Hex2Addr(p.GetDstToken().GetAddress())
	srcAmt := p.GetSrcAmt()
	recvAmt := p.GetDstAmt()
	releaseAmt := p.GetSrcReleaseAmt()
	return solsha3.SoliditySHA3(
		[]string{"string", "address", "uint64", "uint64", "address", "uint256", "uint64", "address", "uint256", "uint256"},
		[]interface{}{"rfq price", mmAddr, p.GetValidThru(), p.GetSrcToken().GetChainId(), srcToken, srcAmt, p.GetDstToken().GetChainId(), dstToken, recvAmt, releaseAmt},
	)
}

func (q *Quote) GetQuoteHash() eth.Hash {
	return eth.Hex2Hash(q.GetHash())
}

func (q *Quote) GetSrcAmt() *big.Int {
	srcAmt, _ := big.NewInt(0).SetString(q.GetSrcAmount(), 10)
	return srcAmt
}

func (q *Quote) GetDstAmt() *big.Int {
	dstAmt, _ := big.NewInt(0).SetString(q.GetDstAmount(), 10)
	return dstAmt
}

func (q *Quote) GetSrcReleaseAmt() *big.Int {
	srcReleaseAmt, _ := big.NewInt(0).SetString(q.GetSrcReleaseAmount(), 10)
	return srcReleaseAmt
}

func (q *Quote) GetSenderAddr() eth.Addr {
	return eth.Hex2Addr(q.GetSender())
}

func (q *Quote) GetReceiverAddr() eth.Addr {
	return eth.Hex2Addr(q.GetReceiver())
}

func (q *Quote) GetRefundToAddr() eth.Addr {
	return eth.Hex2Addr(q.GetRefundTo())
}

func (q *Quote) GetMMAddr() eth.Addr {
	return eth.Hex2Addr(q.GetMmAddr())
}

func (q *Quote) GetSrcChainId() uint64 {
	return q.GetSrcToken().GetChainId()
}

func (q *Quote) GetDstChainId() uint64 {
	return q.GetDstToken().GetChainId()
}

func (q *Quote) EncodeQuoteHash() eth.Hash {
	st := q.GetSrcToken()
	rt := q.GetDstToken()
	var quoteHash = solsha3.SoliditySHA3(
		[]string{"uint64", "address", "uint256", "uint256", "uint64", "address", "uint256", "uint64", "uint64", "address", "address", "address", "address"},
		st.ChainId, st.Address, q.SrcAmount, q.SrcReleaseAmount, rt.ChainId, rt.Address, q.DstAmount, q.DstDeadline, q.Nonce, q.Sender, q.Receiver, q.RefundTo, q.MmAddr)
	return eth.Bytes2Hash(quoteHash)
}

func (q *Quote) ValidateQuoteHash() bool {
	return q.GetQuoteHash() == q.EncodeQuoteHash()
}

func (q *Quote) ToQuoteOnChain() rfq.RFQQuote {
	return rfq.RFQQuote{
		SrcChainId:        q.GetSrcChainId(),
		SrcToken:          q.SrcToken.GetAddr(),
		SrcAmount:         q.GetSrcAmt(),
		SrcReleaseAmount:  q.GetSrcReleaseAmt(),
		DstChainId:        q.GetDstChainId(),
		DstToken:          q.DstToken.GetAddr(),
		DstAmount:         q.GetDstAmt(),
		Deadline:          uint64(q.DstDeadline),
		Nonce:             q.Nonce,
		Sender:            q.GetSenderAddr(),
		Receiver:          q.GetReceiverAddr(),
		RefundTo:          q.GetRefundToAddr(),
		LiquidityProvider: q.GetMMAddr(),
	}
}
