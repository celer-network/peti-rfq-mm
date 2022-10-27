package rfqmm

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/celer-network/endpoint-proxy/endpointproxy"
	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/rfq-mm/sdk/bindings/interfaces/ierc20"
	"github.com/celer-network/rfq-mm/sdk/bindings/interfaces/iweth"
	"github.com/celer-network/rfq-mm/sdk/bindings/rfq"
	"github.com/celer-network/rfq-mm/sdk/common"
	"github.com/celer-network/rfq-mm/sdk/eth"
	"github.com/celer-network/rfq-mm/sdk/service/rfqmm/proto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	maxPendingTxNum     = 20 // max number of tx in pending status (already in txpool)
	maxSubmittingTxNum  = 10 // max number of tx being submitted (not in txpool yet)
	updateMsgPeriod     = 24 * time.Hour
	gasPriceValidPeriod = 5 * time.Minute
)

type RfqMmChainConfig struct {
	ChainId                                             uint64
	Name, Gateway                                       string
	BlkInterval, BlkDelay, MaxBlkDelta, ForwardBlkDelay uint64
	GasLimit                                            uint64
	AddGasEstimateRatio                                 float64
	// Legacy gas price flags
	AddGasGwei   uint64
	MinGasGwei   uint64
	MaxGasGwei   uint64
	ForceGasGwei string
	// EIP-1559 gas price flags
	MaxFeePerGasGwei         uint64
	MaxPriorityFeePerGasGwei uint64
	// if ProxyPort > 0, a proxy with this port will be created to support some special chain such as harmony, celo.
	// chainID will be used to determined which type proxy to create, so make sure the chainID is supported in the "endpoint-proxy"
	// create a proxy to the Gateway, and eth-client will be created to "127.0.0.1:ProxyPort"
	// more detail, https://github.com/celer-network/endpoint-proxy
	ProxyPort int

	Rfq    string
	Native *common.Token
}

type ChainManager struct {
	chains   map[uint64]*Chain
	eventIDs map[string]eth.Hash
}

func NewChainManager(configs []*RfqMmChainConfig) *ChainManager {
	chains := make(map[uint64]*Chain)
	for _, config := range configs {
		chain := NewChain(config)
		chains[chain.ChainId] = chain
	}
	cm := &ChainManager{chains: chains, eventIDs: map[string]eth.Hash{}}
	cm.startUpdateMsgFees(updateMsgPeriod)
	return cm
}

func (cm *ChainManager) startUpdateMsgFees(duration time.Duration) {
	// run once
	cm.updateMsgFees()
	go func() {
		ticker := time.NewTicker(duration)
		for range ticker.C {
			cm.updateMsgFees()
		}
	}()
}

func (cm *ChainManager) updateMsgFees() {
	for _, chain := range cm.chains {
		fee, err := chain.GetMsgFee(nil, eth.ZeroHash.Bytes())
		if err != nil {
			log.Errorf("failed to get msg fee on %d", chain.ChainId)
			continue
		}
		chain.MsgFee = fee
	}
}

func (cm *ChainManager) GetChain(chainId uint64) (*Chain, error) {
	chain, ok := cm.chains[chainId]
	if !ok {
		return nil, proto.NewErr(proto.ErrCode_ERROR_CHAIN_MANAGER, fmt.Sprintf("no chain %d", chainId))
	}
	return chain, nil
}

var _ ChainQuerier = &ChainManager{}

func (cm *ChainManager) GetRfqFee(srcChainId, dstChainId uint64, amount *big.Int) (*big.Int, error) {
	chain, err := cm.GetChain(srcChainId)
	if err != nil {
		return nil, err
	}
	return chain.GetRfqFee(nil, dstChainId, amount)
}

func (cm *ChainManager) GetMsgFee(chainId uint64) (*big.Int, error) {
	chain, err := cm.GetChain(chainId)
	if err != nil {
		return nil, err
	}
	return chain.MsgFee, nil
}

func (cm *ChainManager) GetNativeToken(chainId uint64) (*common.Token, error) {
	chain, err := cm.GetChain(chainId)
	if err != nil {
		return nil, err
	}
	return chain.NativeToken, nil
}

func (cm *ChainManager) GetGasPrice(chainId uint64) (*big.Int, error) {
	chain, err := cm.GetChain(chainId)
	if err != nil {
		return nil, err
	}
	return chain.GetGasPrice(), nil
}

func (cm *ChainManager) GetERC20Balance(chainId uint64, token, account eth.Addr) (*big.Int, error) {
	chain, err := cm.GetChain(chainId)
	if err != nil {
		return nil, err
	}
	return chain.GetERC20Balance(nil, token, account)
}

func (cm *ChainManager) GetNativeBalance(chainId uint64, account eth.Addr) (*big.Int, error) {
	chain, err := cm.GetChain(chainId)
	if err != nil {
		return nil, err
	}
	return chain.GetNativeBalance(account)
}

func (cm *ChainManager) GetQuoteStatus(chainId uint64, quoteHash eth.Hash) (uint8, error) {
	chain, err := cm.GetChain(chainId)
	if err != nil {
		return 0, err
	}
	return chain.GetQuoteStatus(nil, quoteHash)
}

func (cm *ChainManager) VerifyRfqEvent(chainId uint64, tx eth.Hash, evName string) (bool, error) {
	chain, err := cm.GetChain(chainId)
	if err != nil {
		return false, err
	}
	id, ok := cm.eventIDs[evName]
	if !ok {
		id = rfq.GetEventId(evName)
		if id == eth.ZeroHash {
			return false, proto.NewErr(proto.ErrCode_ERROR_CHAIN_MANAGER, fmt.Sprintf("no ID found for event %s", evName))
		}
		cm.eventIDs[evName] = id
	}
	return chain.VerifyRfqEvent(tx, id, evName)
}

type Chain struct {
	*ethclient.Client
	ChainId       uint64
	BlockDelay    uint64
	NativeToken   *common.Token
	RfqContract   *rfq.Rfq
	RfqAddress    eth.Addr
	IWETH         *iweth.Iweth
	MsgFee        *big.Int
	TxOptions     []ethutils.TxOption
	GasPriceCache *dataCache
}

func NewChain(config *RfqMmChainConfig) *Chain {
	var ec *ethclient.Client
	var err error
	if config.ProxyPort > 0 {
		if err = endpointproxy.StartProxy(config.Gateway, config.ChainId, int(config.ProxyPort)); err != nil {
			log.Fatalln("can not start proxy for chain:", config.ChainId, "gateway:", config.Gateway, "port:", config.ProxyPort, "err:", err)
		}
		ec, err = ethclient.Dial(fmt.Sprintf("http://127.0.0.1:%d", config.ProxyPort))
		if err != nil {
			log.Fatalln("chainId", config.ChainId, "dial", config.Gateway, "err:", err)
		}
	} else {
		ec, err = ethclient.Dial(config.Gateway)
		if err != nil {
			log.Fatalln("chainId", config.ChainId, "dial", config.Gateway, "err:", err)
		}
	}
	chid, err := ec.ChainID(context.Background())
	if err != nil {
		log.Fatalf("get chainid %d err: %s", config.ChainId, err)
	}
	if chid.Uint64() != config.ChainId {
		log.Fatalf("chainid mismatch! cfg has %d but onchain has %d", config.ChainId, chid.Uint64())
	}
	if eth.Hex2Addr(config.Rfq) == eth.ZeroAddr {
		log.Fatalf("no rfq contract on %d", chid)
	}
	rfqContract, err := rfq.NewRfq(eth.Hex2Addr(config.Rfq), ec)
	if err != nil {
		log.Fatalln("rfq contract at", config.Rfq, "err:", err)
	}
	iwethContract, err := iweth.NewIweth(eth.Hex2Addr(config.Native.Address), ec)
	if err != nil {
		log.Fatalln("rfq contract at", config.Rfq, "err:", err)
	}
	txOptions := []ethutils.TxOption{
		ethutils.WithBlockDelay(config.BlkDelay),
		ethutils.WithPollingInterval(time.Duration(config.BlkInterval) * time.Second * 4),
		ethutils.WithAddGasEstimateRatio(float64(config.AddGasEstimateRatio)),
		ethutils.WithGasLimit(config.GasLimit),
		ethutils.WithAddGasGwei(config.AddGasGwei),
		ethutils.WithMaxGasGwei(config.MaxGasGwei),
		ethutils.WithMinGasGwei(config.MinGasGwei),
		ethutils.WithForceGasGwei(config.ForceGasGwei),
		ethutils.WithMaxFeePerGasGwei(config.MaxFeePerGasGwei),
		ethutils.WithMaxPriorityFeePerGasGwei(config.MaxPriorityFeePerGasGwei),
		ethutils.WithMaxPendingTxNum(maxPendingTxNum),
		ethutils.WithMaxSubmittingTxNum(maxSubmittingTxNum),
	}
	gasPriceCache := newDataCache(big.NewInt(0).Bytes(), gasPriceValidPeriod, func() ([]byte, error) {
		price, err := ec.SuggestGasPrice(context.Background())
		if err != nil {
			return big.NewInt(0).Bytes(), err
		} else {
			return price.Bytes(), nil
		}
	})
	chain := &Chain{
		Client:        ec,
		ChainId:       config.ChainId,
		BlockDelay:    config.BlkDelay,
		NativeToken:   config.Native,
		RfqContract:   rfqContract,
		RfqAddress:    eth.Hex2Addr(config.Rfq),
		IWETH:         iwethContract,
		MsgFee:        big.NewInt(0),
		TxOptions:     txOptions,
		GasPriceCache: gasPriceCache,
	}
	return chain
}

func (c Chain) GetChainId() uint64 {
	return c.ChainId
}

func (c Chain) GetNativeToken() *common.Token {
	return c.NativeToken
}

func (c Chain) GetRfqFee(opts *bind.CallOpts, _chainId uint64, _amount *big.Int) (*big.Int, error) {
	return c.RfqContract.GetRfqFee(opts, _chainId, _amount)
}

func (c Chain) GetMsgFee(opts *bind.CallOpts, _message []byte) (*big.Int, error) {
	return c.RfqContract.GetMsgFee(opts, _message)
}

func (c Chain) GetQuoteStatus(opts *bind.CallOpts, _quoteHash eth.Hash) (uint8, error) {
	return c.RfqContract.Quotes(opts, _quoteHash)
}

func (c Chain) GetERC20Balance(opts *bind.CallOpts, token, account eth.Addr) (*big.Int, error) {
	erc20, err := ierc20.NewIerc20(token, c.Client)
	if err != nil {
		return nil, fmt.Errorf("erc20 contract at %s err %s", token, err)
	}
	return erc20.BalanceOf(opts, account)
}

func (c Chain) GetNativeBalance(account eth.Addr) (*big.Int, error) {
	return c.Client.BalanceAt(context.Background(), account, nil)
}

func (c Chain) GetGasPrice() *big.Int {
	return new(big.Int).SetBytes(c.GasPriceCache.get())
}

func (c Chain) VerifyRfqEvent(tx, evID eth.Hash, evName string) (bool, error) {
	receipt, err := c.Client.TransactionReceipt(context.Background(), tx)
	if err != nil {
		return false, proto.NewErr(proto.ErrCode_ERROR_CHAIN_MANAGER, fmt.Sprintf("get TransactionReceipt err:%s", err))
	}
	var expectedLog *ethtypes.Log
	for _, log := range receipt.Logs {
		if len(log.Topics) > 0 && log.Topics[0] == evID && log.Address == c.RfqAddress {
			expectedLog = log
		}
	}
	if expectedLog == nil {
		log.Errorf("No event %s found from tx %x on chain %d", evName, tx, c.ChainId)
		return false, nil
	}
	if expectedLog.Removed {
		log.Errorf("Event %s from tx %x on chain %d is removed", evName, tx, c.ChainId)
		return false, nil
	}
	// make sure event isn't too recent
	blk, err := c.Client.BlockNumber(context.Background())
	if expectedLog.BlockNumber > blk-c.BlockDelay {
		return false, proto.NewErr(proto.ErrCode_ERROR_CHAIN_MANAGER, fmt.Sprintf("Event block %d too soon, should only up to block %d", expectedLog.BlockNumber, blk-c.BlockDelay))
	}
	return true, nil
}

// a very simple cache, has only one method, which is "fresh data if necessary and return it".
type dataCache struct {
	data        []byte
	updateTime  time.Time
	validPeriod time.Duration
	updateFunc  func() ([]byte, error)
}

func newDataCache(initData []byte, validPeriod time.Duration, updateFunc func() ([]byte, error)) *dataCache {
	return &dataCache{
		data:        initData,
		updateTime:  time.Unix(0, 0),
		validPeriod: validPeriod,
		updateFunc:  updateFunc,
	}
}

func (c *dataCache) get() []byte {
	if time.Now().Sub(c.updateTime) > c.validPeriod {
		data, err := c.updateFunc()
		if err != nil {
			log.Warnf("failed to update cache, err:%v", err)
		} else {
			c.data = data
			c.updateTime = time.Now()
		}
	}
	return c.data
}
