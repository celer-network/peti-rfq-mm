package common

import (
	"github.com/ethereum/go-ethereum/common"
)

func (t *Token) GetAddr() common.Address {
	return common.HexToAddress(t.Address)
}

func (t *Token) EqualBasically(tt *Token) bool {
	return (t.ChainId == tt.ChainId) && (t.GetAddr() == tt.GetAddr())
}
