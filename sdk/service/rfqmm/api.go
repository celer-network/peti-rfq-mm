package rfqmm

import (
	"context"
	"math/big"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/rfq-mm/sdk/eth"
	"github.com/celer-network/rfq-mm/sdk/service/rfqmm/proto"
	"google.golang.org/grpc"
)

const BestPeriodMultiplier = 1.2

func (c *Client) Price(ctx context.Context, in *proto.PriceRequest, opts ...grpc.CallOption) (*proto.PriceResponse, error) {
	if ok, reason := validatePriceRequest(in); !ok {
		return &proto.PriceResponse{Err: proto.NewErr(proto.ErrCode_ERROR_INVALID_ARGUMENTS, reason).ToCommonErr()}, nil
	}
	return c.ApiClient.Price(ctx, in, opts...)
}

func (c *Client) Quote(ctx context.Context, in *proto.QuoteRequest, opts ...grpc.CallOption) (*proto.QuoteResponse, error) {
	if ok, reason := validateQuoteRequest(in); !ok {
		return &proto.QuoteResponse{Err: proto.NewErr(proto.ErrCode_ERROR_INVALID_ARGUMENTS, reason).ToCommonErr()}, nil
	}
	return c.ApiClient.Quote(ctx, in, opts...)
}

func (s *Server) Price(ctx context.Context, request *proto.PriceRequest) (response *proto.PriceResponse, err error) {
	if ok, reason := validatePriceRequest(request); !ok {
		return &proto.PriceResponse{Err: proto.NewErr(proto.ErrCode_ERROR_INVALID_ARGUMENTS, reason).ToCommonErr()}, nil
	}
	sendAmount := new(big.Int)
	releaseAmount := new(big.Int)
	receiveAmount := new(big.Int)
	fee := new(big.Int)
	// switch mod, one is sendAmt => receiveAmt, the other one is receiveAmt => sendAmt
	if request.SrcAmount == "" {
		// todo, not supported now
		receiveAmount.SetString(request.DstAmount, 10)
		sendAmount, releaseAmount, fee, err = s.AmountCalculator.CalSendAmt(request.SrcToken, request.DstToken, receiveAmount)
		if err != nil {
			return &proto.PriceResponse{Err: err.(*proto.Err).ToCommonErr()}, nil
		}
	} else {
		sendAmount.SetString(request.SrcAmount, 10)
		receiveAmount, releaseAmount, fee, err = s.AmountCalculator.CalRecvAmt(request.SrcToken, request.DstToken, sendAmount)
		if err != nil {
			return &proto.PriceResponse{Err: err.(*proto.Err).ToCommonErr()}, nil
		}
	}
	mmAddr, err := s.LiquidityProvider.GetLiquidityProviderAddr(request.SrcToken.ChainId)
	if err != nil {
		return &proto.PriceResponse{Err: err.(*proto.Err).ToCommonErr()}, nil
	}
	dstTokenAddr := request.DstToken.GetAddr()
	freezeTime, err := s.LiquidityProvider.AskForFreezing(request.DstToken.ChainId, dstTokenAddr, receiveAmount, request.DstNative)
	if err != nil {
		return &proto.PriceResponse{Err: err.(*proto.Err).ToCommonErr()}, nil
	}

	price := &proto.Price{
		SrcToken:          request.SrcToken,
		SrcAmount:         sendAmount.String(),
		SrcReleaseAmount:  releaseAmount.String(),
		DstToken:          request.DstToken,
		DstAmount:         receiveAmount.String(),
		FeeAmount:         fee.String(),
		ValidThru:         time.Now().Unix() + s.Config.PriceValidPeriod,
		MmAddr:            mmAddr.String(),
		Sig:               "",
		SrcDepositPeriod:  int64(float64(freezeTime) / BestPeriodMultiplier),
		DstTransferPeriod: int64(BestPeriodMultiplier * float64(s.Config.DstTransferPeriod)),
	}
	sigBytes, err := s.RequestSigner.Sign(price.EncodeSignData())
	if err != nil {
		return &proto.PriceResponse{Err: err.(*proto.Err).ToCommonErr()}, nil
	}
	price.Sig = eth.Bytes2Hex(sigBytes)
	return &proto.PriceResponse{Price: price}, nil
}

func (s *Server) Quote(ctx context.Context, request *proto.QuoteRequest) (response *proto.QuoteResponse, err error) {
	defer func() {
		if response.Err == nil {
			log.Infof("Quote with success, quote %s", request.Quote.String())
		} else {
			log.Errorf("Quote with failure, err:%s, quote %s", response.Err.String(), request.Quote.String())
		}
	}()
	if ok, reason := validateQuoteRequest(request); !ok {
		return &proto.QuoteResponse{Err: proto.NewErr(proto.ErrCode_ERROR_INVALID_ARGUMENTS, reason).ToCommonErr()}, nil
	}
	price := request.Price
	quote := request.Quote
	srcAmt := price.GetSrcAmt()
	if !s.RequestSigner.Verify(price.EncodeSignData(), eth.Hex2Bytes(price.Sig)) {
		return &proto.QuoteResponse{Err: proto.NewErr(proto.ErrCode_ERROR_INVALID_ARGUMENTS, "invalid sig").ToCommonErr()}, nil
	}
	if !quote.ValidateQuoteHash() {
		return &proto.QuoteResponse{Err: proto.NewErr(proto.ErrCode_ERROR_INVALID_ARGUMENTS, "invalid quote hash").ToCommonErr()}, nil
	}
	rfqFee, err := s.ChainCaller.GetRfqFee(price.GetSrcChainId(), price.GetDstChainId(), srcAmt)
	if err != nil {
		return &proto.QuoteResponse{Err: err.(*proto.Err).ToCommonErr()}, nil
	}
	if new(big.Int).Sub(srcAmt, rfqFee).Cmp(price.GetSrcReleaseAmt()) == -1 {
		return &proto.QuoteResponse{Err: proto.NewErr(proto.ErrCode_ERROR_INVALID_ARGUMENTS, "incorrect src release amount").ToCommonErr()}, nil
	}
	dstAmt := price.GetDstAmt()
	dstTokenAddr := request.Price.DstToken.GetAddr()
	freezeTime, err := s.LiquidityProvider.AskForFreezing(price.GetDstChainId(), dstTokenAddr, dstAmt, request.DstNative)
	if err != nil {
		return &proto.QuoteResponse{Err: err.(*proto.Err).ToCommonErr()}, nil
	}
	if time.Now().Unix()+freezeTime < quote.SrcDeadline {
		return &proto.QuoteResponse{Err: proto.NewErr(proto.ErrCode_ERROR_INVALID_ARGUMENTS, "srcDeadline too large").ToCommonErr()}, nil
	}
	if time.Now().Unix()+s.Config.DstTransferPeriod > quote.DstDeadline {
		return &proto.QuoteResponse{Err: proto.NewErr(proto.ErrCode_ERROR_INVALID_ARGUMENTS, "dstDeadline too small").ToCommonErr()}, nil
	}
	sigBytes, err := s.RequestSigner.Sign(quote.GetQuoteHash().Bytes())
	if err != nil {
		return &proto.QuoteResponse{Err: err.(*proto.Err).ToCommonErr()}, nil
	}
	// no freeze before user deposit token
	//err = s.LiquidityProvider.FreezeLiquidity(price.GetDstChainId(), dstTokenAddr, dstAmt, quote.SrcDeadline, quote.GetQuoteHash(), request.DstNative)
	//if err != nil {
	//	return &proto.QuoteResponse{Err: err.(*proto.Err).ToCommonErr()}, nil
	//}
	return &proto.QuoteResponse{QuoteSig: eth.Bytes2Hex(sigBytes)}, nil
}

func validatePriceRequest(request *proto.PriceRequest) (bool, string) {
	if request.SrcToken == nil || request.DstToken == nil {
		return false, "SrcToken or DstToken is nil"
	}
	if request.SrcAmount == "" && request.DstAmount == "" {
		return false, "SrcAmount and DstAmount are both empty"
	}
	if request.SrcAmount == "" {
		if _, ok := new(big.Int).SetString(request.DstAmount, 10); !ok {
			return false, "invalid SrcAmount"
		}
	} else {
		if _, ok := new(big.Int).SetString(request.SrcAmount, 10); !ok {
			return false, "invalid DstAmount"
		}
	}
	return true, ""
}

func validateQuoteRequest(request *proto.QuoteRequest) (bool, string) {
	price := request.Price
	quote := request.Quote
	if request.Price == nil || request.Quote == nil {
		return false, "price or quote is nil"
	}
	if price.SrcToken == nil || price.DstToken == nil {
		return false, "price.SrcToken or price.DstToken is nil"
	}
	if price.SrcAmount == "" || price.DstAmount == "" || price.SrcReleaseAmount == "" {
		return false, "price.SrcAmount, price.DstAmount or price.SrcReleaseAmount is empty"
	}
	if time.Now().Unix() > price.ValidThru {
		return false, "past price valid time"
	}
	if price.GetMMAddr() != quote.GetMMAddr() {
		return false, "mm addr mismatch"
	}
	if !quote.SrcToken.EqualBasically(price.SrcToken) || !quote.DstToken.EqualBasically(price.DstToken) {
		return false, "token in price and quote mismatch"
	}
	if quote.SrcAmount != price.SrcAmount || quote.DstAmount != price.DstAmount || quote.SrcReleaseAmount != price.SrcReleaseAmount {
		return false, "amount in price and quote mismatch"
	}
	if quote.Sender == "" || quote.Receiver == "" || quote.MmAddr == "" {
		return false, "quote.Sender, quote.Receiver or quote.MmAddr is empty"
	}
	if time.Now().Unix() > quote.SrcDeadline {
		return false, "past src deadline"
	}
	if quote.DstDeadline < quote.SrcDeadline {
		return false, "dst deadline is earlier than src deadline"
	}
	if !quote.ValidateQuoteHash() {
		return false, "quote hahs mismatch"
	}
	return true, ""
}
