package rfqmm

import (
	"math/big"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/peti-rfq-mm/sdk/eth"
	"github.com/celer-network/peti-rfq-mm/sdk/service/rfqmm/proto"
)

type RequestSignerConfig struct {
	ChainId    uint64
	Keystore   string
	Passphrase string
}

type DefaultRequestSigner struct {
	Signer  ethutils.Signer
	Address eth.Addr
}

func NewRequestSigner(config *RequestSignerConfig) RequestSigner {
	signer, addr, err := createSigner(config.Keystore, config.Passphrase, big.NewInt(int64(config.ChainId)))
	if err != nil {
		panic(err)
	}
	return &DefaultRequestSigner{
		Signer:  signer,
		Address: addr,
	}
}

var _ RequestSigner = &DefaultRequestSigner{}

func (rs *DefaultRequestSigner) Sign(data []byte) ([]byte, error) {
	sig, err := rs.Signer.SignEthMessage(data)
	if err != nil {
		return nil, proto.NewErr(proto.ErrCode_ERROR_REQUEST_SIGNER, err.Error())
	}
	return sig, nil
}

func (rs *DefaultRequestSigner) Verify(data, sig []byte) bool {
	addr, err := ethutils.RecoverSigner(data, sig)
	if err != nil {
		return false
	}
	if rs.Address != addr {
		return false
	}
	return true
}
