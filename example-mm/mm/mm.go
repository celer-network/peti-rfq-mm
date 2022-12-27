package mm

import (
	"crypto/tls"

	"github.com/celer-network/goutils/log"
	rfqserver "github.com/celer-network/peti-rfq-mm/sdk/service/rfq"
	rfqproto "github.com/celer-network/peti-rfq-mm/sdk/service/rfq/proto"
	"github.com/celer-network/peti-rfq-mm/sdk/service/rfqmm"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	LPConfig         = "lp"
	ChainConfig      = "multichain"
	FeeConfig        = "fee"
	PriceProviderUrl = "priceprovider.url"
	RfqServerUrl     = "rfqserver.url"
	RfqServerApikey  = "rfqserver.apikey"
	RequestSigner    = "requestsigner"
	MMConfig         = "mm"
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
	var signerConfig *rfqmm.SignerConfig
	err = viper.UnmarshalKey(RequestSigner, &signerConfig)
	if err != nil {
		log.Fatalf("failed to load signer configs:%v", err)
		return nil
	}
	rsChainId := signerConfig.ChainId
	requestSigner, _ := lm.GetLP(rsChainId)
	if signerConfig.Keystore != "" {
		requestSigner.SetSigner(rfqmm.NewSigner(signerConfig))
	}

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
	serverConfig := new(rfqmm.ServerConfig)
	err = viper.UnmarshalKey(MMConfig, serverConfig)
	if err != nil {
		log.Fatalf("failed to load mm server configs:%v", err)
		return nil
	}

	server := rfqmm.NewServer(serverConfig, client, cm, lp, ac, requestSigner)

	return &ExampleMM{
		Server:             server,
		DefaultLiqProvider: lp,
		Client:             client,
	}
}
