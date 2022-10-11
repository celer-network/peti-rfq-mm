package mm

import (
	"encoding/json"
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/rfq-mm/sdk/common"
	"github.com/celer-network/rfq-mm/sdk/service/rfqmm"
	mmproto "github.com/celer-network/rfq-mm/sdk/service/rfqmm/proto"
	"gopkg.in/resty.v1"
)

type Price struct {
	//UpdateEpoch uint64        `json:"updateEpoch,omitempty"`
	AssetPrice []*AssetPrice `json:"assetPrice,omitempty"`
}

type AssetPrice struct {
	Symbol string `json:"symbol,omitempty"`
	//ChainIds     []uint64 `json:"chainIds,omitempty"`
	Price        uint32 `json:"price,omitempty"`
	ExtraPower10 uint32 `json:"extraPower10,omitempty"`
}

type ExamplePriceProvider struct {
	priceUrl       string
	updateCtl      chan bool
	updateDuration time.Duration
	priceCache     map[string]float64
	mux            sync.RWMutex
}

var _ rfqmm.PriceProvider = &ExamplePriceProvider{}

func NewExamplePriceProvider(url string) *ExamplePriceProvider {
	return &ExamplePriceProvider{
		priceUrl:       url,
		updateCtl:      make(chan bool),
		updateDuration: 30 * time.Minute,
		priceCache:     make(map[string]float64),
	}
}

func (pp *ExamplePriceProvider) GetPrice(token *common.Token) (float64, error) {
	pp.mux.RLock()
	defer pp.mux.RUnlock()
	if price, ok := pp.priceCache[token.Symbol]; !ok {
		return 0, mmproto.NewErr(mmproto.ErrCode_ERROR_PRICE_PROVIDER, fmt.Sprintf("no price for token %s", token.Symbol))
	} else {
		return price, nil
	}
}

func (pp *ExamplePriceProvider) StopUpdatePrice() {
	pp.updateCtl <- true
}

func (pp *ExamplePriceProvider) UpdatePrice() {
	go func() {
		if pp.updateCtl == nil {
			log.Panicln("nil update control channel")
		}
		// do once
		err := pp.updatePrice()
		if err != nil {
			log.Errorf("updatePrice err:%s", err)
		}
		ticker := time.NewTicker(pp.updateDuration)
		for {
			select {
			case <-ticker.C:
				err := pp.updatePrice()
				if err != nil {
					log.Errorf("updatePrice err:%s", err)
				}
			case <-pp.updateCtl:
				return
			}
		}
	}()
}

func (pp *ExamplePriceProvider) updatePrice() error {
	cp := &Price{}
	client := resty.New()
	response, err := client.R().Get(pp.priceUrl)
	if err != nil {
		return fmt.Errorf("fail to get price change json from s3: %v", err)
	}
	if response.StatusCode() != 200 {
		return fmt.Errorf("fail to get price change err status:%d", response.StatusCode())
	}
	err = json.Unmarshal(response.Body(), cp)
	//unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	//err = unmarshaler.Unmarshal(strings.NewReader(response.String()), cp)
	if err != nil {
		return fmt.Errorf("fail to get price change json from s3: %v", err)
	}
	pp.mux.Lock()
	defer pp.mux.Unlock()
	for _, asset := range cp.AssetPrice {
		price := float64(asset.Price) / math.Pow(10, float64(asset.ExtraPower10+4))
		pp.priceCache[asset.Symbol] = price
	}
	return nil
}
