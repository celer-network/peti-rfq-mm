package rfqmm

import (
	"fmt"
	"math"
	"math/big"

	"github.com/celer-network/rfq-mm/sdk/common"
	"github.com/celer-network/rfq-mm/sdk/eth"
	"github.com/celer-network/rfq-mm/sdk/service/rfqmm/proto"
)

var _ AmountCalculator = &DefaultAmtCalculator{}

type DefaultAmtCalculator struct {
	DstGasCost uint64
	SrcGasCost uint64
	// 100% = 1000000
	FeePercGlobal        uint32
	PerChainPairOverride map[uint64]map[uint64]uint32
	PerTokenPairOverride map[string]map[string]uint32
	GasPrice             map[uint64]uint64

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

func (ac *DefaultAmtCalculator) SetDstGasCost(gasCost uint64) {
	ac.DstGasCost = gasCost
}

func (ac *DefaultAmtCalculator) SetSrcGasCost(gasCost uint64) {
	ac.SrcGasCost = gasCost
}

func (ac *DefaultAmtCalculator) SetGlobalFeePerc(feePerc uint32) error {
	if feePerc > 1e6 {
		return proto.NewErr(proto.ErrCode_ERROR_INVALID_ARGUMENTS, "feePercGlobal too large")
	}
	ac.FeePercGlobal = feePerc
	return nil
}

func (ac *DefaultAmtCalculator) SetPerChainPairFeePercOverride(overrides []*ChainOverride) error {
	for _, override := range overrides {
		if override.Perc > 1e6 {
			return proto.NewErr(proto.ErrCode_ERROR_INVALID_ARGUMENTS, "feePercGlobal too large")
		}
		if ac.PerChainPairOverride[override.SrcChainId] == nil {
			ac.PerChainPairOverride[override.SrcChainId] = map[uint64]uint32{override.DstChainId: override.Perc}
		} else {
			ac.PerChainPairOverride[override.SrcChainId][override.DstChainId] = override.Perc
		}
	}
	return nil
}

func (ac *DefaultAmtCalculator) SetPerTokenPairFeePercOverride(overrides []*TokenOverride) error {
	for _, override := range overrides {
		if override.Perc > 1e6 {
			return proto.NewErr(proto.ErrCode_ERROR_INVALID_ARGUMENTS, "feePercGlobal too large")
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

func (ac *DefaultAmtCalculator) CalRecvAmt(tokenIn, tokenOut *common.Token, amountIn *big.Int) (amountOut, releaseAmt, fee *big.Int, err error) {
	tokenInPrice, err := ac.PriceProvider.GetPrice(tokenIn)
	if err != nil {
		return
	}
	tokenOutPrice, err := ac.PriceProvider.GetPrice(tokenOut)
	if err != nil {
		return
	}
	nativeIn, err := ac.Querier.GetNativeToken(tokenIn.ChainId)
	if err != nil {
		return
	}
	nativeInPrice, err := ac.PriceProvider.GetPrice(nativeIn)
	if err != nil {
		return
	}
	nativeOut, err := ac.Querier.GetNativeToken(tokenOut.ChainId)
	if err != nil {
		return
	}
	nativeOutPrice, err := ac.PriceProvider.GetPrice(nativeOut)
	if err != nil {
		return
	}
	rfqFeeAmt, err := ac.Querier.GetRfqFee(tokenIn.ChainId, tokenOut.ChainId, amountIn)
	if err != nil {
		return
	}
	releaseAmt = new(big.Int).Sub(amountIn, rfqFeeAmt)
	chainIn := tokenIn.ChainId
	chainOut := tokenOut.ChainId

	msgFeeAmt, _ := ac.Querier.GetMsgFee(chainOut)
	dstGasCost := big.NewInt(int64(ac.DstGasCost))
	srcGasCost := big.NewInt(int64(ac.SrcGasCost))
	mmFeeAmt := ac.calMmFee(tokenIn, tokenOut, amountIn)
	msgFeeInIn := convertAmount(msgFeeAmt, nativeOutPrice, tokenInPrice, tokenIn.Decimals-nativeOut.Decimals)
	dstGasCostInIn := convertAmount(new(big.Int).Mul(dstGasCost, big.NewInt(int64(ac.GasPrice[chainOut]))), nativeOutPrice, tokenInPrice, tokenIn.Decimals-nativeOut.Decimals)
	srcGasCostInIn := convertAmount(new(big.Int).Mul(srcGasCost, big.NewInt(int64(ac.GasPrice[chainIn]))), nativeInPrice, tokenInPrice, tokenIn.Decimals-nativeIn.Decimals)
	if tokenIn.ChainId != tokenOut.ChainId {
		fee = new(big.Int).Add(mmFeeAmt, msgFeeInIn)
		fee.Add(fee, dstGasCostInIn)
		fee.Add(fee, srcGasCostInIn)
	} else {
		// in same chain swap scenario, no msgFee and no srcGasCost
		fee = new(big.Int).Add(mmFeeAmt, dstGasCostInIn)
	}
	amountOut = convertAmount(new(big.Int).Sub(releaseAmt, fee), tokenInPrice, tokenOutPrice, tokenOut.Decimals-tokenIn.Decimals)
	return
}

func (ac *DefaultAmtCalculator) CalSendAmt(tokenIn, tokenOut *common.Token, amountOut *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return nil, nil, nil, nil
}

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
