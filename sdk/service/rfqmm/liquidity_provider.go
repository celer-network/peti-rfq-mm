package rfqmm

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/peti-rfq-mm/sdk/bindings/interfaces/ierc20"
	"github.com/celer-network/peti-rfq-mm/sdk/bindings/rfq"
	"github.com/celer-network/peti-rfq-mm/sdk/common"
	"github.com/celer-network/peti-rfq-mm/sdk/eth"
	"github.com/celer-network/peti-rfq-mm/sdk/service/rfqmm/proto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	dstTransfer          = "DstTransfer"
	srcRelease           = "SrcRelease"
	sameChainTransfer    = "SameChainTransfer"
	NativeTokenReference = "ffffffffffffffffffffffffffffffffffffffff"
	TPPolicyAll          = "All"
	TPPolicyPrefixAny2Of = "Any2Of="
	TPPolicyPrefixOneOf  = "OneOf="
)

var _ LiquidityProvider = &DefaultLiquidityProvider{}

type DefaultLiquidityProvider struct {
	paused       bool
	txrs         map[uint64]*ethutils.Transactor
	chainManager *ChainManager
	liqManager   *LiqManager
	tokenPair    map[string]bool
}

func NewDefaultLiquidityProvider(cm *ChainManager, lm *LiqManager) *DefaultLiquidityProvider {
	lp := &DefaultLiquidityProvider{
		paused:       false,
		txrs:         make(map[uint64]*ethutils.Transactor),
		chainManager: cm,
		liqManager:   lm,
		tokenPair:    make(map[string]bool),
	}
	// construct transactor for each chain
	for _, chainId := range lm.GetChains() {
		addr, signer, _ := lm.GetSigner(chainId)
		chain, err := cm.GetChain(chainId)
		if err != nil {
			log.Errorf("GetChain err:%s", err)
			continue
		}
		lp.txrs[chainId] = ethutils.NewTransactorByExternalSigner(addr, signer, chain.Client, big.NewInt(int64(chain.ChainId)), chain.TxOptions...)
	}
	// update liquidity amount
	lm.UpdateLiqAmt(cm)

	// approve to rfq contract
	lp.approveERC20ToRfq()
	return lp
}

func (d DefaultLiquidityProvider) IsPaused() bool {
	return d.paused
}

func (d DefaultLiquidityProvider) GetTokens() []*common.Token {
	tokensMap := d.liqManager.GetTokens()
	res := make([]*common.Token, 0)
	native := eth.Hex2Addr(NativeTokenReference)
	for chainId, tokens := range tokensMap {
		wNative, err := d.chainManager.GetNativeWrap(chainId)
		if err != nil {
			continue
		}
		addWrappedNative := false
		for _, token := range tokens {
			if token.GetAddr() == native || token.GetAddr() == wNative.GetAddr() {
				addWrappedNative = true
			} else {
				res = append(res, token)
			}
		}
		if addWrappedNative && wNative.GetAddr() != eth.ZeroAddr {
			res = append(res, wNative)
		}
	}
	return res
}

// policy str is one of {"All", "Any2Of=<chainId>-<symbol>,<chainId>-<symbol>...", "OneOf=<chainId>-<symbol>,<chainId>-<symbol>"}
func (d *DefaultLiquidityProvider) SetupTokenPairs(policies []string) {
	for _, policy := range policies {
		if policy == TPPolicyAll {
			d.setupTokenPairsAll()
			return
		} else if strings.HasPrefix(policy, TPPolicyPrefixAny2Of) {
			arg := strings.TrimPrefix(policy, TPPolicyPrefixAny2Of)
			d.setupTokenPairsAny2Of(strings.Split(arg, ","))
		} else if strings.HasPrefix(policy, TPPolicyPrefixOneOf) {
			arg := strings.TrimPrefix(policy, TPPolicyPrefixOneOf)
			d.setupTokenPairsOneOf(strings.Split(arg, ","))
		}
		continue
	}
}

func (d DefaultLiquidityProvider) HasTokenPair(srcToken, dstToken *common.Token) bool {
	key := genTokenPairKey(srcToken, dstToken)
	return d.tokenPair[key]
}

func (d DefaultLiquidityProvider) GetLiquidityProviderAddr(chainId uint64) (eth.Addr, error) {
	return d.liqManager.GetLiquidityProvider(chainId)
}

func (d DefaultLiquidityProvider) AskForFreezing(chainId uint64, token eth.Addr, amount *big.Int, isNative bool) (int64, error) {
	if d.paused {
		return 0, proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_PROVIDER, "liquidity provider is paused due to some serious error")
	}
	if isNative {
		native, err := d.substituteNativeToken(chainId, token)
		if err != nil {
			return 0, err
		}
		return d.liqManager.AskForFreezing(chainId, native, amount)
	}
	return d.liqManager.AskForFreezing(chainId, token, amount)
}

func (d DefaultLiquidityProvider) FreezeLiquidity(chainId uint64, token eth.Addr, amount *big.Int, until int64, hash eth.Hash, isNative bool) error {
	if d.paused {
		return proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_PROVIDER, "liquidity provider is paused due to some serious error")
	}
	if isNative {
		native, err := d.substituteNativeToken(chainId, token)
		if err != nil {
			return err
		}
		return d.liqManager.ReserveLiquidity(chainId, native, amount, until, hash)
	}
	return d.liqManager.ReserveLiquidity(chainId, token, amount, until, hash)
}

func (d DefaultLiquidityProvider) UnfreezeLiquidity(chainId uint64, hash eth.Hash) error {
	return d.liqManager.UnfreezeLiquidity(chainId, hash)
}

func (d *DefaultLiquidityProvider) DstTransfer(transferNative bool, _quote rfq.RFQQuote, opts ...ethutils.TxOption) (eth.Hash, error) {
	if d.paused {
		return eth.ZeroHash, proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_PROVIDER, "liquidity provider is paused due to some serious error")
	}
	quoteHash := _quote.Hash()
	// check if it's a same chain swap
	if _quote.DstChainId == _quote.SrcChainId {
		return d.sameChainTransfer(transferNative, _quote, opts...)
	}
	chain, err := d.chainManager.GetChain(_quote.DstChainId)
	if err != nil {
		return eth.ZeroHash, err
	}
	// confirm liquidity before dst transfer
	err = d.confirmLiquidity(_quote.DstChainId, _quote.DstToken, _quote.DstAmount, int64(_quote.Deadline), quoteHash, transferNative)
	if err != nil {
		return eth.ZeroHash, err
	}
	txr, ok := d.txrs[_quote.DstChainId]
	if !ok {
		return eth.ZeroHash, proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_PROVIDER, fmt.Sprintf("no transactor for chain %d", _quote.DstChainId))
	}
	var method ethutils.TxMethod
	if transferNative {
		method = func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return chain.RfqContract.DstTransferNative(opts, _quote)
		}
		opts = append(opts, ethutils.WithEthValue(new(big.Int).Add(_quote.DstAmount, chain.MsgFee)))
	} else {
		method = func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return chain.RfqContract.DstTransfer(opts, _quote)
		}
		opts = append(opts, ethutils.WithEthValue(chain.MsgFee))
	}
	tx, err := txr.Transact(d.genTxHandler(dstTransfer, _quote, transferNative, false), method, opts...)
	if err != nil {
		return eth.ZeroHash, proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_PROVIDER, fmt.Sprintf("Transact err:%s", err))
	}
	return tx.Hash(), nil
}

func (d *DefaultLiquidityProvider) SrcRelease(_quote rfq.RFQQuote, _execMsgCallData []byte, opts ...ethutils.TxOption) (eth.Hash, error) {
	if d.paused {
		return eth.ZeroHash, proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_PROVIDER, "liquidity provider is paused due to some serious error")
	}
	chain, err := d.chainManager.GetChain(_quote.SrcChainId)
	if err != nil {
		return eth.ZeroHash, err
	}
	txr, ok := d.txrs[_quote.SrcChainId]
	if !ok {
		return eth.ZeroHash, proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_PROVIDER, fmt.Sprintf("no transactor for chain %d", _quote.SrcChainId))
	}
	// determine release native or not
	releaseNative := false
	if chain.NativeWrap.GetAddr() == _quote.SrcToken {
		releaseNative, err = d.liqManager.ReleaseNative(_quote.SrcChainId)
		if err != nil {
			return eth.ZeroHash, err
		}
	}
	var method ethutils.TxMethod
	if releaseNative {
		method = func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return chain.RfqContract.SrcReleaseNative(opts, _quote, _execMsgCallData)
		}
	} else {
		method = func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return chain.RfqContract.SrcRelease(opts, _quote, _execMsgCallData)
		}
	}
	tx, err := txr.Transact(d.genTxHandler(srcRelease, _quote, false, releaseNative), method, opts...)
	if err != nil {
		return eth.ZeroHash, proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_PROVIDER, fmt.Sprintf("Transact err:%s", err))
	}
	return tx.Hash(), nil
}

func (d *DefaultLiquidityProvider) sameChainTransfer(transferNative bool, _quote rfq.RFQQuote, opts ...ethutils.TxOption) (eth.Hash, error) {
	chain, err := d.chainManager.GetChain(_quote.DstChainId)
	if err != nil {
		return eth.ZeroHash, err
	}
	quoteHash := _quote.Hash()
	// confirm liquidity before same chain transfer
	err = d.confirmLiquidity(_quote.DstChainId, _quote.DstToken, _quote.DstAmount, int64(_quote.Deadline), quoteHash, transferNative)
	if err != nil {
		return eth.ZeroHash, err
	}
	txr, ok := d.txrs[_quote.DstChainId]
	if !ok {
		return eth.ZeroHash, proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_PROVIDER, fmt.Sprintf("no transactor for chain %d", _quote.DstChainId))
	}
	// determine release native or not
	releaseNative := false
	if chain.NativeWrap.GetAddr() == _quote.SrcToken {
		releaseNative, err = d.liqManager.ReleaseNative(_quote.SrcChainId)
		if err != nil {
			return eth.ZeroHash, err
		}
	}
	var method ethutils.TxMethod
	if transferNative {
		method = func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return chain.RfqContract.SameChainTransferNative(opts, _quote, releaseNative)
		}
		opts = append(opts, ethutils.WithEthValue(_quote.DstAmount))
	} else {
		method = func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return chain.RfqContract.SameChainTransfer(opts, _quote, releaseNative)
		}
	}
	tx, err := txr.Transact(
		d.genTxHandler(sameChainTransfer, _quote, transferNative, releaseNative), method, opts...)
	if err != nil {
		return eth.ZeroHash, proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_PROVIDER, fmt.Sprintf("Transact err:%s", err))
	}
	return tx.Hash(), nil
}

func (d *DefaultLiquidityProvider) genTxHandler(methodName string, _quote rfq.RFQQuote, transferNative, releaseNative bool) *ethutils.TransactionStateHandler {
	quoteHash := _quote.Hash()
	switch methodName {
	case dstTransfer:
		return &ethutils.TransactionStateHandler{
			OnMined: func(receipt *ethtypes.Receipt) {
				if receipt.Status == ethtypes.ReceiptStatusSuccessful {
					log.Infof("%s succeeded, tx %x. quote hash %s", methodName, receipt.TxHash, quoteHash)
					err := d.transferOutLiquidity(_quote.DstChainId, _quote.DstToken, _quote.DstAmount, quoteHash, transferNative)
					if err != nil {
						log.Errorf("TransferOutLiquidity err:%s", err)
					}
				} else {
					log.Errorf("%s failed, tx %x. quote hash %s", methodName, receipt.TxHash, quoteHash)
					err := d.UnfreezeLiquidity(_quote.DstChainId, quoteHash)
					if err != nil {
						log.Errorf("UnfreezeLiquidity err:%s", err)
					}
					d.pause()
				}
			},
			OnError: func(tx *ethtypes.Transaction, err error) {
				log.Warnf("%s err: %s. quote hash %s", methodName, err, quoteHash)
			},
		}
	case srcRelease:
		return &ethutils.TransactionStateHandler{
			OnMined: func(receipt *ethtypes.Receipt) {
				if receipt.Status == ethtypes.ReceiptStatusSuccessful {
					log.Infof("%s succeeded, tx %x. quote hash %s", methodName, receipt.TxHash, quoteHash)
					err := d.releaseInLiquidity(_quote.SrcChainId, _quote.SrcToken, _quote.SrcReleaseAmount, releaseNative)
					if err != nil {
						log.Errorf("ReleaseInLiquidity err:%s", err)
					}
				} else {
					log.Errorf("%s failed, tx %x. quote hash %s", methodName, receipt.TxHash, quoteHash)
					d.pause()
				}
			},
			OnError: func(tx *ethtypes.Transaction, err error) {
				log.Warnf("%s err: %s. quote hash %s", methodName, err, quoteHash)
			},
		}
	case sameChainTransfer:
		return &ethutils.TransactionStateHandler{
			OnMined: func(receipt *ethtypes.Receipt) {
				if receipt.Status == ethtypes.ReceiptStatusSuccessful {
					log.Infof("%s succeeded, tx %x. quote hash %s", methodName, receipt.TxHash, quoteHash)
					err := d.transferOutLiquidity(_quote.DstChainId, _quote.DstToken, _quote.DstAmount, quoteHash, transferNative)
					if err != nil {
						log.Errorf("TransferOutLiquidity err:%s", err)
					}
					err = d.releaseInLiquidity(_quote.SrcChainId, _quote.SrcToken, _quote.SrcReleaseAmount, releaseNative)
					if err != nil {
						log.Errorf("ReleaseInLiquidity err:%s", err)
					}
				} else {
					log.Errorf("%s failed, tx %x. quote hash %s", methodName, receipt.TxHash, quoteHash)
					err := d.UnfreezeLiquidity(_quote.DstChainId, quoteHash)
					if err != nil {
						log.Errorf("UnfreezeLiquidity err:%s", err)
					}
					d.pause()
				}
			},
			OnError: func(tx *ethtypes.Transaction, err error) {
				log.Warnf("%s err: %s. quote hash %s", methodName, err, quoteHash)
			},
		}
	default:
		return nil
	}
}

func (d DefaultLiquidityProvider) approveERC20ToRfq() {
	for chainId, txr := range d.txrs {
		chain, _ := d.chainManager.GetChain(chainId)
		tokens, amounts, _ := d.liqManager.GetLiqNeedApprove(chainId)
		for i, token := range tokens {
			amount := amounts[i]
			tx, err := txr.Transact(
				&ethutils.TransactionStateHandler{
					OnMined: func(receipt *ethtypes.Receipt) {
						if receipt.Status == ethtypes.ReceiptStatusSuccessful {
							log.Infof("ApproveERC20ToRfq succeeded, tx %x, amount %s", receipt.TxHash, amount)
						} else {
							log.Warnf("ApproveERC20ToRfq failed, tx %x, amount %s", receipt.TxHash, amount)
						}
					},
					OnError: func(tx *ethtypes.Transaction, err error) {
						log.Warnf("ApproveERC20ToRfq err: %s. amount %s", err, amount.String())
					},
				},
				func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
					erc20, err := ierc20.NewIerc20Transactor(token.GetAddr(), transactor)
					if err != nil {
						return nil, fmt.Errorf("erc20 contract at %s err %s", token, err)
					}
					return erc20.Approve(opts, chain.RfqAddress, amount)
				},
			)
			if err != nil {
				log.Errorf("ApproveERC20ToRfq err:%s", err)
			} else {
				log.Infof("Approve %s of %s on %d to RFQ contract %x, txHash %x", amount, token, chainId, chain.RfqAddress, tx.Hash())
			}
		}
	}
}

// wrap or unwrap before dst transfer is not supported for now.
func (d DefaultLiquidityProvider) wrapNative(chainId uint64, amount *big.Int, opts ...ethutils.TxOption) (eth.Hash, error) {
	chain, err := d.chainManager.GetChain(chainId)
	if err != nil {
		return eth.ZeroHash, err
	}
	txr, ok := d.txrs[chainId]
	if !ok {
		return eth.ZeroHash, proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_PROVIDER, fmt.Sprintf("no transactor for chain %d", chainId))
	}
	tx, err := txr.Transact(
		&ethutils.TransactionStateHandler{
			OnMined: func(receipt *ethtypes.Receipt) {
				if receipt.Status == ethtypes.ReceiptStatusSuccessful {
					log.Infof("WrapNative succeeded, tx %x, amount %s", receipt.TxHash, amount.String())
				} else {
					log.Warnf("WrapNative failed, tx %x, amount %s", receipt.TxHash, amount.String())
				}
			},
			OnError: func(tx *ethtypes.Transaction, err error) {
				log.Warnf("WrapNative err: %s. amount %s", err, amount.String())
			},
		},
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return chain.IWETH.Deposit(opts)
		},
		append(opts, ethutils.WithEthValue(amount))...,
	)
	if err != nil {
		return eth.ZeroHash, err
	}
	return tx.Hash(), nil
}

func (d DefaultLiquidityProvider) unwrapNative(chainId uint64, amount *big.Int, opts ...ethutils.TxOption) (eth.Hash, error) {
	chain, err := d.chainManager.GetChain(chainId)
	if err != nil {
		return eth.ZeroHash, err
	}
	txr, ok := d.txrs[chainId]
	if !ok {
		return eth.ZeroHash, proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_PROVIDER, fmt.Sprintf("no transactor for chain %d", chainId))
	}
	tx, err := txr.Transact(
		&ethutils.TransactionStateHandler{
			OnMined: func(receipt *ethtypes.Receipt) {
				if receipt.Status == ethtypes.ReceiptStatusSuccessful {
					log.Infof("UnwrapNative succeeded, tx %x, amount %s", receipt.TxHash, amount.String())
				} else {
					log.Warnf("UnwrapNative failed, tx %x, amount %s", receipt.TxHash, amount.String())
				}
			},
			OnError: func(tx *ethtypes.Transaction, err error) {
				log.Warnf("UnwrapNative err: %s. amount %s", err, amount.String())
			},
		},
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return chain.IWETH.Withdraw(opts, amount)
		},
		opts...,
	)
	if err != nil {
		return eth.ZeroHash, err
	}
	return tx.Hash(), nil
}

func (d DefaultLiquidityProvider) substituteNativeToken(chainId uint64, wrap eth.Addr) (eth.Addr, error) {
	expectedWrap, err := d.chainManager.GetNativeWrap(chainId)
	if err != nil {
		return eth.ZeroAddr, err
	}
	if expectedWrap == nil || expectedWrap.GetAddr() == eth.ZeroAddr || expectedWrap.GetAddr() == wrap {
		return eth.Hex2Addr(NativeTokenReference), nil
	} else {
		return eth.ZeroAddr, proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_PROVIDER, fmt.Sprintf("invalid wrap %x, should be %x", wrap, expectedWrap.GetAddr()))
	}
}

func (d DefaultLiquidityProvider) confirmLiquidity(chainId uint64, token eth.Addr, amount *big.Int, until int64, hash eth.Hash, isNative bool) error {
	if isNative {
		native, err := d.substituteNativeToken(chainId, token)
		if err != nil {
			return err
		}
		return d.liqManager.ConfirmLiquidity(chainId, native, amount, until, hash)
	}
	return d.liqManager.ConfirmLiquidity(chainId, token, amount, until, hash)
}

func (d DefaultLiquidityProvider) transferOutLiquidity(chainId uint64, token eth.Addr, amount *big.Int, hash eth.Hash, isNative bool) error {
	if isNative {
		native, err := d.substituteNativeToken(chainId, token)
		if err != nil {
			return err
		}
		return d.liqManager.TransferOutLiquidity(chainId, native, amount, hash)
	}
	return d.liqManager.TransferOutLiquidity(chainId, token, amount, hash)
}

func (d DefaultLiquidityProvider) releaseInLiquidity(chainId uint64, token eth.Addr, amount *big.Int, isNative bool) error {
	if isNative {
		native, err := d.substituteNativeToken(chainId, token)
		if err != nil {
			return err
		}
		return d.liqManager.ReleaseInLiquidity(chainId, native, amount)
	}
	return d.liqManager.ReleaseInLiquidity(chainId, token, amount)
}

func (d *DefaultLiquidityProvider) pause() {
	d.paused = true
}

func (d *DefaultLiquidityProvider) setupTokenPairsAll() {
	tokens := d.GetTokens()
	if len(tokens) <= 1 {
		return
	}
	log.Debugf("setup token pairs with policy All")
	logStr := "Token pairs:"
	for i := 0; i < len(tokens)-1; i++ {
		for j := i + 1; j < len(tokens); j++ {
			logStr += fmt.Sprintf(" %d-%s>>%d-%s |", tokens[i].ChainId, tokens[i].Symbol, tokens[j].ChainId, tokens[j].Symbol)
			logStr += fmt.Sprintf(" %d-%s>>%d-%s |", tokens[j].ChainId, tokens[j].Symbol, tokens[i].ChainId, tokens[i].Symbol)
			d.tokenPair[genTokenPairKey(tokens[i], tokens[j])] = true
			d.tokenPair[genTokenPairKey(tokens[j], tokens[i])] = true
		}
	}
	log.Debugf(logStr)
}

func (d *DefaultLiquidityProvider) setupTokenPairsAny2Of(list []string) {
	if len(list) <= 1 {
		return
	}
	tokens := d.getTokensByStrList(list)
	if len(tokens) <= 1 {
		return
	}
	log.Debugf("setup token pairs with policy Any2Of")
	logStr := "Token pairs:"
	for i := 0; i < len(tokens)-1; i++ {
		for j := i + 1; j < len(tokens); j++ {
			logStr += fmt.Sprintf(" %d-%s>>%d-%s |", tokens[i].ChainId, tokens[i].Symbol, tokens[j].ChainId, tokens[j].Symbol)
			logStr += fmt.Sprintf(" %d-%s>>%d-%s |", tokens[j].ChainId, tokens[j].Symbol, tokens[i].ChainId, tokens[i].Symbol)
			d.tokenPair[genTokenPairKey(tokens[i], tokens[j])] = true
			d.tokenPair[genTokenPairKey(tokens[j], tokens[i])] = true
		}
	}
	log.Debugf(logStr)
}

func (d *DefaultLiquidityProvider) setupTokenPairsOneOf(list []string) {
	if len(list) != 2 {
		return
	}
	tokens := d.getTokensByStrList(list)
	if len(tokens) != 2 {
		return
	}
	log.Debugf("setup token pairs with policy OneOf")
	log.Debugf("Token pairs: %d-%s>>%d-%s", tokens[0].ChainId, tokens[0].Symbol, tokens[1].ChainId, tokens[1].Symbol)
	d.tokenPair[genTokenPairKey(tokens[0], tokens[1])] = true
}

// string within list should be in format of [chainId]-[symbol]
func (d DefaultLiquidityProvider) getTokensByStrList(list []string) []*common.Token {
	tokens := make([]*common.Token, 0)
	supportedTokens := d.GetTokens()
	for _, str := range list {
		splitRes := strings.Split(str, "-")
		if len(splitRes) != 2 {
			continue
		}
		chainId, err := strconv.Atoi(splitRes[0])
		if err != nil {
			continue
		}
		symbol := splitRes[1]
		for _, token := range supportedTokens {
			if token.ChainId == uint64(chainId) && token.Symbol == symbol {
				tokens = append(tokens, token)
			}
		}
	}
	return tokens
}

func genTokenPairKey(srcToken, dstToken *common.Token) string {
	return fmt.Sprintf("%d-%s-%d-%d-%s-%d", srcToken.ChainId, eth.FormatAddrHex(srcToken.Address), srcToken.Decimals,
		dstToken.ChainId, eth.FormatAddrHex(dstToken.Address), dstToken.Decimals)
}
