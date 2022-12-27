package rfqmm

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"
	"sort"
	"strings"
	"sync"
	"time"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/peti-rfq-mm/sdk/common"
	"github.com/celer-network/peti-rfq-mm/sdk/eth"
	"github.com/celer-network/peti-rfq-mm/sdk/service/rfqmm/proto"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	LiqOpAny = iota
	LiqOpReserve
	LiqOpConfirm
)

type LiqManager struct {
	LPs map[uint64]*LiqProvider
}

func NewLiqManager(configs []*LPConfig) *LiqManager {
	lps := make(map[uint64]*LiqProvider)
	for _, config := range configs {
		lp := NewLiqProvider(config)
		lp.log()
		lps[config.ChainId] = lp
	}
	return &LiqManager{LPs: lps}
}

func (d *LiqManager) GetLiqNeedApprove(chainId uint64) ([]*common.Token, []*big.Int, error) {
	lp, err := d.GetLP(chainId)
	if err != nil {
		return nil, nil, err
	}
	tokens, amounts := lp.getLiqNeedApprove()
	return tokens, amounts, nil
}

func (d *LiqManager) GetChains() []uint64 {
	res := make([]uint64, 0)
	for k := range d.LPs {
		res = append(res, k)
	}
	return res
}

func (d *LiqManager) GetTokens() map[uint64][]*common.Token {
	res := make(map[uint64][]*common.Token, 0)
	for _, lp := range d.LPs {
		res[lp.chainId] = lp.getTokens()
	}
	return res
}

func (d *LiqManager) GetLiquidityProvider(chainId uint64) (eth.Addr, error) {
	lp, err := d.GetLP(chainId)
	if err != nil {
		return eth.ZeroAddr, err
	}
	return lp.address, nil
}

func (d *LiqManager) AskForFreezing(chainId uint64, token eth.Addr, amount *big.Int) (int64, error) {
	lp, err := d.GetLP(chainId)
	if err != nil {
		return 0, err
	}
	// clear liquidity
	lp.clear()
	available, err := lp.getAvailableLiquidity(eth.Addr2Hex(token))
	if err != nil {
		return 0, err
	}
	if amount.Cmp(available) == 1 {
		return 0, proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_MANAGER, fmt.Sprintf("no sufficient liquidity to freeze, token %s, chain %d", token, chainId))
	}
	return lp.getFreezeTime(eth.Addr2Hex(token))
}

func (d *LiqManager) ReserveLiquidity(chainId uint64, token eth.Addr, amount *big.Int, until int64, hash eth.Hash) error {
	lp, err := d.GetLP(chainId)
	if err != nil {
		return err
	}
	return lp.reserveLiquidity(eth.Addr2Hex(token), amount, until, hash)
}

func (d *LiqManager) ConfirmLiquidity(chainId uint64, token eth.Addr, amount *big.Int, until int64, hash eth.Hash) error {
	lp, err := d.GetLP(chainId)
	if err != nil {
		return err
	}
	return lp.confirmLiquidity(eth.Addr2Hex(token), amount, until, hash)
}

func (d *LiqManager) UnfreezeLiquidity(chainId uint64, hash eth.Hash) error {
	lp, err := d.GetLP(chainId)
	if err != nil {
		return err
	}
	lp.unfreezeLiquidity(hash)
	return nil
}

func (d *LiqManager) TransferOutLiquidity(chainId uint64, token eth.Addr, amount *big.Int, hash eth.Hash) error {
	lp, err := d.GetLP(chainId)
	if err != nil {
		return err
	}
	return lp.transferOutLiquidity(eth.Addr2Hex(token), amount, hash)
}

func (d *LiqManager) ReleaseInLiquidity(chainId uint64, token eth.Addr, amount *big.Int) error {
	lp, err := d.GetLP(chainId)
	if err != nil {
		return err
	}
	return lp.releaseInLiquidity(eth.Addr2Hex(token), amount)
}

func (d *LiqManager) ReleaseNative(chainId uint64) (bool, error) {
	lp, err := d.GetLP(chainId)
	if err != nil {
		return false, err
	}
	return lp.releaseNative, nil
}

func (d *LiqManager) UpdateLiqAmt(querier ChainQuerier) {
	for chainId, lp := range d.LPs {
		for _, liq := range lp.liqs {
			if liq.amount != nil {
				continue
			}
			var balance *big.Int
			var err error
			if liq.token.GetAddr() == eth.Hex2Addr(NativeTokenReference) {
				balance, err = querier.GetNativeBalance(chainId, lp.address)
			} else {
				balance, err = querier.GetERC20Balance(chainId, liq.token.GetAddr(), lp.address)
			}
			if err != nil {
				log.Errorf("GetBalance err:%s", err)
				continue
			}
			liq.amount = balance
			log.Infof("Liquidity amount of %s(%s) on %d is updated to %s", liq.token.Symbol, liq.token.Address, chainId, balance.String())
		}
	}
}

func (d *LiqManager) GetLP(chainId uint64) (*LiqProvider, error) {
	if lp, found := d.LPs[chainId]; !found {
		return nil, proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_MANAGER, fmt.Sprintf("no liquidity provider on %d", chainId))
	} else {
		return lp, nil
	}
}

func (d *LiqManager) GetSigner(chainId uint64) (eth.Addr, ethutils.Signer, error) {
	lp, err := d.GetLP(chainId)
	if err != nil {
		return eth.ZeroAddr, nil, err
	}
	return lp.address, lp.signer, nil
}

type LiqProvider struct {
	mux           sync.RWMutex
	signer        ethutils.Signer
	signerAddress eth.Addr
	chainId       uint64
	address       eth.Addr
	liqs          map[string]*Liquidity
	// sorted slice by LiqOpDetail.Until in ascending order
	liqOps []*LiqOpDetail
	// to minimize searching cost when doing unfreeze
	hashToUntil   map[eth.Hash]int64
	releaseNative bool
}

type LiqOpDetail struct {
	Type   int
	Until  int64
	Token  string
	Amount *big.Int
	Hash   eth.Hash
}

type LPConfig struct {
	ChainId       uint64
	Keystore      string
	Passphrase    string
	Liqs          []*LiquidityConfig
	ReleaseNative bool
}

type LiquidityConfig struct {
	Address    string
	Symbol     string
	Amount     string
	Approve    string
	Decimals   int32
	FreezeTime int64
}

type SignerConfig struct {
	ChainId    uint64
	Keystore   string
	Passphrase string
}

func NewSigner(config *SignerConfig) (ethutils.Signer, eth.Addr) {
	signer, addr, err := createSigner(config.Keystore, config.Passphrase, big.NewInt(int64(config.ChainId)))
	if err != nil {
		panic(err)
	}
	return signer, addr
}

func NewLiqProvider(config *LPConfig) *LiqProvider {
	signer, addr, err := createSigner(config.Keystore, config.Passphrase, big.NewInt(int64(config.ChainId)))
	if err != nil {
		panic(err)
	}
	liqs := make(map[string]*Liquidity)
	for _, liq := range config.Liqs {
		amount, _ := new(big.Int).SetString(liq.Amount, 10)
		approved, _ := new(big.Int).SetString(liq.Approve, 10)
		token := &common.Token{ChainId: config.ChainId, Address: liq.Address, Symbol: liq.Symbol, Decimals: liq.Decimals}
		liquidity := &Liquidity{
			amount:     amount,
			reserved:   big.NewInt(0),
			confirmed:  big.NewInt(0),
			approved:   approved,
			token:      token,
			freezeTime: liq.FreezeTime,
		}
		liqs[eth.FormatAddrHex(liq.Address)] = liquidity
	}
	return &LiqProvider{
		signer:        signer,
		chainId:       config.ChainId,
		address:       addr,
		signerAddress: addr,
		liqs:          liqs,
		liqOps:        make([]*LiqOpDetail, 0),
		hashToUntil:   make(map[eth.Hash]int64),
		releaseNative: config.ReleaseNative,
	}
}

func (lp *LiqProvider) SetSigner(signer ethutils.Signer, addr eth.Addr) {
	lp.signer = signer
	lp.signerAddress = addr
}

func (lp *LiqProvider) log() {
	log.Infof("Configuration of liquidity provider on chain %d", lp.chainId)
	log.Infof("\taddr:%s", lp.address)
	log.Infof("\tliqs:")
	for _, liq := range lp.liqs {
		log.Infof("\t\t%s, %s, %s, freeze time %d seconds", liq.token.Symbol, liq.token.Address, liq.amount, liq.freezeTime)
	}
}

func (lp *LiqProvider) getAvailableLiquidity(token string) (*big.Int, error) {
	lp.mux.RLock()
	defer lp.mux.RUnlock()
	if liq, found := lp.liqs[token]; found {
		return liq.available(), nil
	} else {
		return nil, proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_MANAGER, fmt.Sprintf("not supported token %s", token))
	}
}

// clear all liq ops that are time out
func (lp *LiqProvider) clear() {
	lp.mux.Lock()
	defer lp.mux.Unlock()
	dropTo := 0
	now := time.Now().Unix()
	for i, detail := range lp.liqOps {
		if detail.Until < now {
			dropTo = i + 1
			switch detail.Type {
			case LiqOpConfirm:
				lp.liqs[detail.Token].cancelConfirmation(detail.Amount)
			case LiqOpReserve:
				lp.liqs[detail.Token].cancelReservation(detail.Amount)
			}
			delete(lp.hashToUntil, detail.Hash)
			continue
		}
		break
	}
	if dropTo == len(lp.liqOps) {
		lp.liqOps = make([]*LiqOpDetail, 0)
	} else {
		lp.liqOps = lp.liqOps[dropTo:]
	}
}

func (lp *LiqProvider) reserveLiquidity(token string, amount *big.Int, until int64, hash eth.Hash) error {
	lp.mux.Lock()
	defer lp.mux.Unlock()
	// check hash map at first
	if _, ok := lp.hashToUntil[hash]; ok {
		return nil
	}
	if liq, found := lp.liqs[token]; found {
		err := liq.reserve(amount)
		if err != nil {
			return err
		}
		lp.hashToUntil[hash] = until
		lp.liqOps = append(lp.liqOps, &LiqOpDetail{
			Type:   LiqOpReserve,
			Until:  until,
			Token:  token,
			Amount: amount,
			Hash:   hash,
		})
		// as liqOps is a sorted array, we should sort the slice after appending
		sort.Slice(lp.liqOps, func(i, j int) bool {
			return lp.liqOps[i].Until < lp.liqOps[j].Until
		})
		return nil
	} else {
		return proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_MANAGER, fmt.Sprintf("not supported token %s", token))
	}
}

func (lp *LiqProvider) confirmLiquidity(token string, amount *big.Int, until int64, hash eth.Hash) error {
	lp.mux.Lock()
	defer lp.mux.Unlock()
	// try to release reserved liquidity before confirm it
	lp.privateUnfreezeLiquidity(hash, LiqOpReserve)
	// check hash map at first
	if _, ok := lp.hashToUntil[hash]; ok {
		return nil
	}
	if liq, found := lp.liqs[token]; found {
		err := liq.confirm(amount)
		if err != nil {
			return err
		}
		lp.hashToUntil[hash] = until
		lp.liqOps = append(lp.liqOps, &LiqOpDetail{
			Type:   LiqOpConfirm,
			Until:  until,
			Token:  token,
			Amount: amount,
			Hash:   hash,
		})
		// as liqOps is a sorted array, we should sort the slice after appending
		sort.Slice(lp.liqOps, func(i, j int) bool {
			return lp.liqOps[i].Until < lp.liqOps[j].Until
		})
		return nil
	} else {
		return proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_MANAGER, fmt.Sprintf("not supported token %s", token))
	}
}

// unfreeze a frozen liquidity with specified hash, can be used for releasing reserved/confirmed liquidity
func (lp *LiqProvider) unfreezeLiquidity(hash eth.Hash) {
	lp.mux.Lock()
	defer lp.mux.Unlock()
	lp.privateUnfreezeLiquidity(hash, LiqOpAny)
}

// private method without lock&unlock
func (lp *LiqProvider) privateUnfreezeLiquidity(hash eth.Hash, opType int) {
	if hash == eth.ZeroHash {
		return
	}
	// check hash map at first
	until, ok := lp.hashToUntil[hash]
	if !ok {
		return
	}
	for i, detail := range lp.liqOps {
		if detail.Until < until {
			continue
		} else if detail.Until > until {
			return
		} else if detail.Hash == hash {
			if opType == LiqOpAny || opType == detail.Type {
				if i+1 == len(lp.liqOps) {
					lp.liqOps = lp.liqOps[:i]
				} else {
					lp.liqOps = append(lp.liqOps[:i], lp.liqOps[i+1:]...)
				}
				switch detail.Type {
				case LiqOpConfirm:
					lp.liqs[detail.Token].cancelConfirmation(detail.Amount)
				case LiqOpReserve:
					lp.liqs[detail.Token].cancelReservation(detail.Amount)
				}
				delete(lp.hashToUntil, hash)
			}
			return
		}
	}
}

func (lp *LiqProvider) transferOutLiquidity(token string, amount *big.Int, hash eth.Hash) error {
	lp.mux.Lock()
	defer lp.mux.Unlock()
	// try to release confirmed liquidity before transfer out
	lp.privateUnfreezeLiquidity(hash, LiqOpConfirm)
	if liq, found := lp.liqs[token]; found {
		liq.transferOut(amount)
		return nil
	} else {
		return proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_MANAGER, fmt.Sprintf("not supported token %s", token))
	}
}

func (lp *LiqProvider) releaseInLiquidity(token string, amount *big.Int) error {
	lp.mux.Lock()
	defer lp.mux.Unlock()
	if liq, found := lp.liqs[token]; found {
		liq.releaseIn(amount)
		return nil
	} else {
		return proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_MANAGER, fmt.Sprintf("not supported token %s", token))
	}
}

func (lp *LiqProvider) getTokens() []*common.Token {
	tokens := make([]*common.Token, 0)
	for _, liq := range lp.liqs {
		tokens = append(tokens, liq.token)
	}
	return tokens
}

func (lp *LiqProvider) getFreezeTime(token string) (int64, error) {
	if liq, found := lp.liqs[token]; found {
		return liq.freezeTime, nil
	} else {
		return 0, proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_MANAGER, fmt.Sprintf("not supported token %s", token))
	}
}

func (lp *LiqProvider) getLiqNeedApprove() ([]*common.Token, []*big.Int) {
	tokens := make([]*common.Token, 0)
	amounts := make([]*big.Int, 0)
	for _, liq := range lp.liqs {
		if liq.approved != nil {
			tokens = append(tokens, liq.token)
			amounts = append(amounts, liq.approved)
		}
	}
	return tokens, amounts
}

var _ RequestSigner = &LiqProvider{}

func (lp *LiqProvider) Sign(data []byte) ([]byte, error) {
	sig, err := lp.signer.SignEthMessage(data)
	if err != nil {
		return nil, proto.NewErr(proto.ErrCode_ERROR_REQUEST_SIGNER, err.Error())
	}
	return sig, nil
}

func (lp *LiqProvider) Verify(data, sig []byte) bool {
	addr, err := ethutils.RecoverSigner(data, sig)
	if err != nil {
		return false
	}
	if lp.address != addr {
		return false
	}
	return true
}

type Liquidity struct {
	amount     *big.Int
	reserved   *big.Int
	confirmed  *big.Int
	approved   *big.Int
	token      *common.Token
	freezeTime int64
}

func (liq *Liquidity) Status() string {
	return fmt.Sprintf("[%s/%s/%s]", liq.confirmed, new(big.Int).Add(liq.confirmed, liq.reserved), liq.amount)
}

func (liq *Liquidity) available() *big.Int {
	return new(big.Int).Sub(liq.amount, new(big.Int).Add(liq.reserved, liq.confirmed))
}

func (liq *Liquidity) reserve(amount *big.Int) error {
	after := new(big.Int).Add(liq.reserved, amount)
	if new(big.Int).Add(after, liq.confirmed).Cmp(liq.amount) == 1 {
		return proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_MANAGER, fmt.Sprintf("no sufficient %s on %d to reserve", liq.token.Symbol, liq.token.ChainId))
	}
	liq.reserved = after
	log.Debugf("Liquidity of %s on %d reserved, amount %s. %s", liq.token.Symbol, liq.token.ChainId, amount, liq.Status())
	return nil
}

func (liq *Liquidity) cancelReservation(amount *big.Int) {
	after := new(big.Int).Sub(liq.reserved, amount)
	if after.Sign() == -1 {
		after.SetInt64(0)
		log.Warnf("Reserved liquidity of %s on %d is negative after releasing %s. %s", liq.token.Symbol, liq.token.ChainId, amount, liq.Status())
	}
	liq.reserved = after
	log.Debugf("Liquidity of %s on %d released, amount %s. %s", liq.token.Symbol, liq.token.ChainId, amount, liq.Status())
}

func (liq *Liquidity) confirm(amount *big.Int) error {
	after := new(big.Int).Add(liq.confirmed, amount)
	if new(big.Int).Add(after, liq.reserved).Cmp(liq.amount) == 1 {
		return proto.NewErr(proto.ErrCode_ERROR_LIQUIDITY_MANAGER, fmt.Sprintf("no sufficient %s on %d to confirm", liq.token.Symbol, liq.token.ChainId))
	}
	liq.confirmed = after
	log.Debugf("Liquidity of %s on %d confirmed, amount %s. %s", liq.token.Symbol, liq.token.ChainId, amount, liq.Status())
	return nil
}

func (liq *Liquidity) cancelConfirmation(amount *big.Int) {
	after := new(big.Int).Sub(liq.confirmed, amount)
	if after.Sign() == -1 {
		after.SetInt64(0)
		log.Warnf("Confirmed liquidity of %s on %d is negative after releasing %s. %s", liq.token.Symbol, liq.token.ChainId, amount, liq.Status())
	}
	liq.confirmed = after
	log.Debugf("Liquidity of %s on %d released, amount %s. %s", liq.token.Symbol, liq.token.ChainId, amount, liq.Status())
}

func (liq *Liquidity) transferOut(amount *big.Int) {
	after := new(big.Int).Sub(liq.amount, amount)
	if after.Sign() == -1 {
		log.Panicf("Liquidity of %s on %d is negative after transfer out %s. %s", liq.token.Symbol, liq.token.ChainId, amount, liq.Status())
	}
	liq.amount = after
	log.Debugf("Liquidity of %s on %d deducted by %s. %s", liq.token.Symbol, liq.token.ChainId, amount, liq.Status())
}

func (liq *Liquidity) releaseIn(amount *big.Int) {
	liq.amount = new(big.Int).Add(liq.amount, amount)
	log.Debugf("Liquidity of %s on %d added by %s. %s", liq.token.Symbol, liq.token.ChainId, amount, liq.Status())
}

func createSigner(ksfile, passphrase string, chainid *big.Int) (ethutils.Signer, eth.Addr, error) {
	if strings.HasPrefix(ksfile, "awskms") {
		kmskeyinfo := strings.SplitN(ksfile, ":", 3)
		if len(kmskeyinfo) != 3 {
			return nil, eth.ZeroAddr, fmt.Errorf("%s has wrong format", ksfile)
		}
		awskeysec := []string{"", ""}
		if passphrase != "" {
			awskeysec = strings.SplitN(passphrase, ":", 2)
			if len(awskeysec) != 2 {
				return nil, eth.ZeroAddr, fmt.Errorf("%s has wrong format", passphrase)
			}
		}
		kmsSigner, err := ethutils.NewKmsSigner(kmskeyinfo[1], kmskeyinfo[2], awskeysec[0], awskeysec[1], chainid)
		if err != nil {
			return nil, eth.ZeroAddr, err
		}
		return kmsSigner, kmsSigner.Addr, nil
	}
	ksBytes, err := ioutil.ReadFile(ksfile)
	if err != nil {
		return nil, eth.ZeroAddr, err
	}
	key, err := keystore.DecryptKey(ksBytes, passphrase)
	if err != nil {
		return nil, eth.ZeroAddr, err
	}
	signer, err := ethutils.NewSigner(hex.EncodeToString(crypto.FromECDSA(key.PrivateKey)), chainid)
	return signer, key.Address, err
}
