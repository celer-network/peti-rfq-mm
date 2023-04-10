package rfqmm

import (
	"fmt"
	"math"
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/peti-rfq-mm/sdk/common"
	"github.com/celer-network/peti-rfq-mm/sdk/eth"
	"github.com/celer-network/peti-rfq-mm/sdk/service/rfqmm/proto"
)

var _ AmountCalculator = &DefaultAmtCalculator{}

// DefaultAmtCalculator is a default implementation of interface AmountCalculator
type DefaultAmtCalculator struct {
	// fixed cost related fields
	// how much gas charged on dst chain, gas used by DstTransfer
	DstGasCost uint64
	// how much gas charged on src chain, gas used by SrcRelease
	SrcGasCost uint64
	// map from chain id to charging gas price, only used when Querier(ChainQuerier.GetGasPrice) returns error or 0
	GasPrice map[uint64]uint64

	// personalized percentage fee related fieds
	// 100% = 1000000
	// global fee percentage configuration, will be used when no override matches
	FeePercGlobal uint32
	// override per pairs of chain id
	PerChainPairOverride map[uint64]map[uint64]uint32
	// override per pairs of <chainId-tokenAddr>
	PerTokenPairOverride map[string]map[string]uint32

	// helper
	Querier       ChainQuerier
	PriceProvider PriceProvider
}

type FeeConfig struct {
	DstGasCost     uint64
	SrcGasCost     uint64
	PercGlobal     uint32
	ChainOverrides []*ChainOverride
	TokenOverrides []*TokenOverride
	GasPrices      []*GasPrice
}

type ChainOverride struct {
	SrcChainId, DstChainId uint64
	Perc                   uint32
}

type TokenOverride struct {
	SrcChainId, DstChainId uint64
	SrcToken, DstToken     string
	Perc                   uint32
}

type GasPrice struct {
	ChainId uint64
	Price   uint64
}

type PriceProvider interface {
	GetPrice(token *common.Token) (float64, error)
}

func NewDefaultAmtCalculator(feeConfig *FeeConfig, querier ChainQuerier, priceProvider PriceProvider) *DefaultAmtCalculator {
	ac := &DefaultAmtCalculator{
		DstGasCost:           feeConfig.DstGasCost,
		SrcGasCost:           feeConfig.SrcGasCost,
		FeePercGlobal:        feeConfig.PercGlobal,
		PerChainPairOverride: make(map[uint64]map[uint64]uint32),
		PerTokenPairOverride: make(map[string]map[string]uint32),
		GasPrice:             make(map[uint64]uint64),
		Querier:              querier,
		PriceProvider:        priceProvider,
	}
	err := ac.SetPerChainPairFeePercOverride(feeConfig.ChainOverrides)
	if err != nil {
		panic(err)
	}
	err = ac.SetPerTokenPairFeePercOverride(feeConfig.TokenOverrides)
	if err != nil {
		panic(err)
	}
	ac.SetGasPrice(feeConfig.GasPrices)
	return ac
}

// SetDstGasCost Method sets gas cost charged by MM on dst chain.
func (ac *DefaultAmtCalculator) SetDstGasCost(gasCost uint64) {
	ac.DstGasCost = gasCost
}

// SetSrcGasCost Method sets gas cost charged by MM on src chain.
func (ac *DefaultAmtCalculator) SetSrcGasCost(gasCost uint64) {
	ac.SrcGasCost = gasCost
}

// SetGlobalFeePerc Method sets global fee percentage, of which maximum is 1000000(=100%).
func (ac *DefaultAmtCalculator) SetGlobalFeePerc(feePerc uint32) error {
	if feePerc > 1e6 {
		return proto.NewErr(proto.ErrCode_ERROR_INVALID_ARGUMENTS, "fee percentage too large")
	}
	ac.FeePercGlobal = feePerc
	return nil
}

// SetPerChainPairFeePercOverride Method override fee percentage per chain pair.
func (ac *DefaultAmtCalculator) SetPerChainPairFeePercOverride(overrides []*ChainOverride) error {
	for _, override := range overrides {
		if override.Perc > 1e6 {
			return proto.NewErr(proto.ErrCode_ERROR_INVALID_ARGUMENTS, "fee percentage too large")
		}
		if ac.PerChainPairOverride[override.SrcChainId] == nil {
			ac.PerChainPairOverride[override.SrcChainId] = map[uint64]uint32{override.DstChainId: override.Perc}
		} else {
			ac.PerChainPairOverride[override.SrcChainId][override.DstChainId] = override.Perc
		}
	}
	return nil
}

// SetPerTokenPairFeePercOverride Method override fee percentage per token pair.
func (ac *DefaultAmtCalculator) SetPerTokenPairFeePercOverride(overrides []*TokenOverride) error {
	for _, override := range overrides {
		if override.Perc > 1e6 {
			return proto.NewErr(proto.ErrCode_ERROR_INVALID_ARGUMENTS, "fee percentage too large")
		}
		inStr := fmt.Sprintf("%d-%s", override.SrcChainId, eth.FormatAddrHex(override.SrcToken))
		outStr := fmt.Sprintf("%d-%s", override.DstChainId, eth.FormatAddrHex(override.DstToken))
		if ac.PerTokenPairOverride[inStr] == nil {
			ac.PerTokenPairOverride[inStr] = map[string]uint32{outStr: override.Perc}
		} else {
			ac.PerTokenPairOverride[inStr][outStr] = override.Perc
		}
	}
	return nil
}

// SetGasPrice Method sets gas price charged for each gas in wei by MM.
func (ac *DefaultAmtCalculator) SetGasPrice(prices []*GasPrice) {
	for _, gasPrice := range prices {
		ac.GasPrice[gasPrice.ChainId] = gasPrice.Price
	}
}

func (ac *DefaultAmtCalculator) SetMinFeeUsdValue(minFeeUsdValue uint64) {
	//todo
}

func (ac *DefaultAmtCalculator) SetMaxFeeUsdValue(maxFeeUsdValue uint64) {
	//todo
}

// base fee is composed of 1. src gas cost(only cross chain swap) 2. dst gas cost 3. dst msg fee(only cross chain swap)
func (ac *DefaultAmtCalculator) calBaseFee(tokenIn, tokenOut *common.Token) (baseFee *big.Int, err error) {
	chainIn := tokenIn.ChainId
	chainOut := tokenOut.ChainId
	tokenInPrice, err := ac.PriceProvider.GetPrice(tokenIn)
	if err != nil {
		return
	}

	// 1 src gas cost
	nativeIn, err := ac.Querier.GetNativeWrap(chainIn)
	if err != nil {
		return
	}
	nativeInPrice, err := ac.PriceProvider.GetPrice(nativeIn)
	if err != nil {
		return
	}
	srcGasPrice, err := ac.Querier.GetGasPrice(chainIn)
	if err != nil || srcGasPrice.Sign() == 0 {
		log.Warnf("Fail to get gas price on chain %d, err:%v", chainIn, err)
		srcGasPrice = big.NewInt(int64(ac.GasPrice[chainIn]))
	}
	srcGasCost := big.NewInt(int64(ac.SrcGasCost))
	// represent src gas cost by src token amount
	srcGasCostInIn := convertAmount(new(big.Int).Mul(srcGasCost, srcGasPrice), nativeInPrice, tokenInPrice, tokenIn.Decimals-nativeIn.Decimals)

	// 2 dst gas cost
	nativeOut, err := ac.Querier.GetNativeWrap(chainOut)
	if err != nil {
		return
	}
	nativeOutPrice, err := ac.PriceProvider.GetPrice(nativeOut)
	if err != nil {
		return
	}
	dstGasPrice, err := ac.Querier.GetGasPrice(chainOut)
	if err != nil || dstGasPrice.Sign() == 0 {
		log.Warnf("Fail to get gas price on chain %d, err:%v", chainOut, err)
		dstGasPrice = big.NewInt(int64(ac.GasPrice[chainOut]))
	}
	dstGasCost := big.NewInt(int64(ac.DstGasCost))
	// represent dst gas cost by src token amount
	dstGasCostInIn := convertAmount(new(big.Int).Mul(dstGasCost, dstGasPrice), nativeOutPrice, tokenInPrice, tokenIn.Decimals-nativeOut.Decimals)

	// 3 dst msg fee
	msgFeeAmt, _ := ac.Querier.GetMsgFee(chainOut)
	// represent dst msg fee by src token amount
	msgFeeInIn := convertAmount(msgFeeAmt, nativeOutPrice, tokenInPrice, tokenIn.Decimals-nativeOut.Decimals)

	// sum all cost
	if chainIn != chainOut {
		// all of 3
		baseFee = new(big.Int).Add(srcGasCostInIn, dstGasCostInIn)
		baseFee.Add(baseFee, msgFeeInIn)
	} else {
		baseFee = new(big.Int).Set(dstGasCostInIn)
	}
	return
}

// some notes about amount:
// fee = baseFee + mmFee (fee is used for statistical analysis)
// srcAmount - rfqFee - fee = dstAmount
// for light mm:
//   srcAmount - rfqFee - baseFee = srcReleaseAmount
// for default mm:
//   srcAmount - rfqFee = srcReleaseAmount

// CalRecvAmt Method returns
//   - amountOut: how much `tokenOut` will be received by User
//   - releaseAmt: how much `tokenIn` will be received by MM
//   - fee: how much `tokenIn` is charged as fee in total.
func (ac *DefaultAmtCalculator) CalRecvAmt(tokenIn, tokenOut *common.Token, amountIn, baseFeeForLMM *big.Int, isLightMM bool) (amountOut, releaseAmt, fee *big.Int, err error) {
	tokenInPrice, err := ac.PriceProvider.GetPrice(tokenIn)
	if err != nil {
		return
	}
	tokenOutPrice, err := ac.PriceProvider.GetPrice(tokenOut)
	if err != nil {
		return
	}
	// calculate rfq protocol fee which is paid in src token
	rfqFeeAmt, err := ac.Querier.GetRfqFee(tokenIn.ChainId, tokenOut.ChainId, amountIn)
	if err != nil {
		return
	}

	// calculate base fee, of which unit is src token
	baseFeeAmt := new(big.Int).Set(baseFeeForLMM)
	if !isLightMM {
		baseFeeAmt, err = ac.calBaseFee(tokenIn, tokenOut)
		if err != nil {
			return
		}
	}

	// calculate fee required by mm, of which unit is src token
	mmFeeAmt := ac.calMmFee(tokenIn, tokenOut, amountIn)

	// calculate total fee, described as a sum of base fee and mm fee
	fee = new(big.Int).Add(mmFeeAmt, baseFeeAmt)

	// calculate release amount
	releaseAmt = new(big.Int).Sub(amountIn, rfqFeeAmt)
	if isLightMM {
		releaseAmt = new(big.Int).Sub(releaseAmt, baseFeeAmt)
	}

	// calculate amount out
	dstTransferAmt := new(big.Int).Sub(amountIn, rfqFeeAmt)
	dstTransferAmt.Sub(dstTransferAmt, fee)
	amountOut = convertAmount(dstTransferAmt, tokenInPrice, tokenOutPrice, tokenOut.Decimals-tokenIn.Decimals)
	return
}

// CalSendAmt Method returns
//   - amountIn: how much `tokenIn` should be sent by User
//   - releaseAmt: how much `tokenIn` will be received by MM
//   - fee: how much `tokenIn` is charged as fee in total.
func (ac *DefaultAmtCalculator) CalSendAmt(tokenIn, tokenOut *common.Token, amountOut, baseFeeForLMM *big.Int, isLightMM bool) (amountIn *big.Int, releaseAmt *big.Int, fee *big.Int, err error) {
	// TODO
	return nil, nil, nil, fmt.Errorf("not supported now")
}

// calMmFee returns fee required by mm, of which unit is tokenIn
func (ac *DefaultAmtCalculator) calMmFee(tokenIn, tokenOut *common.Token, amountIn *big.Int) *big.Int {
	perc := ac.FeePercGlobal
	if p, found := ac.PerChainPairOverride[tokenIn.ChainId][tokenOut.ChainId]; found {
		perc = p
	}
	inStr := fmt.Sprintf("%d-%s", tokenIn.ChainId, eth.FormatAddrHex(tokenIn.Address))
	outStr := fmt.Sprintf("%d-%s", tokenOut.ChainId, eth.FormatAddrHex(tokenOut.Address))
	if p, found := ac.PerTokenPairOverride[inStr][outStr]; found {
		perc = p
	}
	mmFee := new(big.Int).Mul(amountIn, big.NewInt(int64(perc)))
	mmFee.Div(mmFee, big.NewInt(1e6))
	return mmFee
}

// decimalDiff = tokenOut.decimal - tokenIn.decimal
func convertAmount(amountIn *big.Int, priceIn, priceOut float64, decimalDiff int32) *big.Int {
	value := new(big.Float).Mul(big.NewFloat(priceIn), new(big.Float).SetInt(amountIn))
	amountOut := new(big.Float).Quo(value, big.NewFloat(priceOut))
	amountOut.Mul(amountOut, big.NewFloat(math.Pow(10, float64(decimalDiff))))
	res, _ := amountOut.Int(nil)
	return res
}
