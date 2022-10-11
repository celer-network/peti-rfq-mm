package mm

import (
	"crypto/tls"

	"github.com/celer-network/goutils/log"
	rfqserver "github.com/celer-network/rfq-mm/sdk/service/rfq"
	rfqproto "github.com/celer-network/rfq-mm/sdk/service/rfq/proto"
	"github.com/celer-network/rfq-mm/sdk/service/rfqmm"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	LPConfig             = "lp"
	ChainConfig          = "multichain"
	FeeConfig            = "fee"
	PriceProviderUrl     = "priceprovider.url"
	RfqServerUrl         = "rfqserver.url"
	RfqServerApikey      = "rfqserver.apikey"
	RequestSigner        = "mm.requestsigner"
	PriceValidPeriod     = "mm.pricevalidperiod"
	SecureTransferPeriod = "mm.securetransferperiod"
	ReportPeriod         = "mm.reportperiod"
	ProcessPeriod        = "mm.processperiod"
)

type ExampleMM struct {
	*rfqmm.Server
	DefaultLiqProvider *rfqmm.DefaultLiquidityProvider
	Client             rfqproto.MMApiClient
}

func NewExampleMM() *ExampleMM {
	// new Liquidity Manager
	var lpConfig []*rfqmm.LPConfig
	err := viper.UnmarshalKey(LPConfig, &lpConfig)
	if err != nil {
		log.Fatalf("failed to load liquidity-provider configs:%v", err)
		return nil
	}
	lm := rfqmm.NewLiqManager(lpConfig)

	// get a signer from Liquidity Manager by chain id
	rsChainId := viper.GetUint64(RequestSigner)
	requestSigner, _ := lm.GetLP(rsChainId)

	// new Chain Manager
	var chainConfig []*rfqmm.RfqMmChainConfig
	err = viper.UnmarshalKey(ChainConfig, &chainConfig)
	if err != nil {
		log.Fatalf("failed to load multichain configs:%v", err)
		return nil
	}
	cm := rfqmm.NewChainManager(chainConfig)

	// new default liquidity provider
	lp := rfqmm.NewDefaultLiquidityProvider(cm, lm)

	// new Amount Calculator
	feeConfig := new(rfqmm.FeeConfig)
	err = viper.UnmarshalKey(FeeConfig, feeConfig)
	if err != nil {
		log.Fatalf("failed to load mmfee configs:%v", err)
		return nil
	}
	// prepare a Price Provider
	priceUrl := viper.GetString(PriceProviderUrl)
	priceProvider := NewExamplePriceProvider(priceUrl)
	priceProvider.UpdatePrice()
	ac := rfqmm.NewDefaultAmtCalculator(feeConfig, cm, priceProvider)

	// new Client of Rfq server
	rfqServerUrl := viper.GetString(RfqServerUrl)
	rfqServerApiKey := viper.GetString(RfqServerApikey)
	client := rfqserver.NewClient(rfqServerUrl, rfqServerApiKey, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))

	// new Server
	processPeriod := viper.GetInt64(ProcessPeriod)
	reportPeriod := viper.GetInt64(ReportPeriod)
	priceValidPeriod := viper.GetInt64(PriceValidPeriod)
	secureTransferPeriod := viper.GetInt64(SecureTransferPeriod)
	serverConfig := &rfqmm.ServerConfig{
		ReportRetryPeriod: reportPeriod,
		ProcessPeriod:     processPeriod,
		PriceValidPeriod:  priceValidPeriod,
		DstTransferPeriod: secureTransferPeriod,
	}
	server := rfqmm.NewServer(serverConfig, client, cm, lp, ac, requestSigner)

	return &ExampleMM{
		Server:             server,
		DefaultLiqProvider: lp,
		Client:             client,
	}
}
