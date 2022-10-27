package proto

import (
	"fmt"

	"github.com/celer-network/peti-rfq-mm/sdk/common"
)

type Err struct {
	Code ErrCode
	Msg  string
}

func NewErr(code ErrCode, msg string) *Err {
	return &Err{
		Code: code,
		Msg:  msg,
	}
}

func (e Err) Error() string {
	return fmt.Sprintf("%s: %s", e.Code.String(), e.Msg)
}

func (e *Err) ToCommonErr() *common.Err {
	return &common.Err{Code: uint32(e.Code), Msg: e.Msg}
}
