package proto

import (
	"math/big"

	"github.com/celer-network/rfq-mm/sdk/eth"
)

func (r *PriceRequest) GetSrcAmt() *big.Int {
	sendAmt, _ := big.NewInt(0).SetString(r.GetSrcAmount(), 10)
	return sendAmt
}

func (r *PriceRequest) GetDstAmt() *big.Int {
	receiveAmt, _ := big.NewInt(0).SetString(r.GetDstAmount(), 10)
	return receiveAmt
}

func (r *QuoteRequest) GetSenderAddr() eth.Addr {
	return eth.Hex2Addr(r.GetSender())
}

func (r *QuoteRequest) GetReceiverAddr() eth.Addr {
	return eth.Hex2Addr(r.GetReceiver())
}

func (r *QuoteRequest) GetRefundToAddr() eth.Addr {
	return eth.Hex2Addr(r.GetRefundTo())
}
