package xerr

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/remote"
)

/**
常用通用固定错误
*/

type CodeError struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

func (e CodeError) Error() string {
	return fmt.Sprintf("code:%d,msg:%s", e.Code, e.Msg)
}

func (e CodeError) GetCode() uint32 {
	return e.Code
}

func NewErr(code uint32, msg string) error {
	return CodeError{
		Code: code,
		Msg:  msg,
	}
}

func NewSprintfErr(code uint32, format string, a ...interface{}) error {
	return CodeError{
		Code: code,
		Msg:  fmt.Sprintf(format, a...),
	}
}

func ErrMsg(code uint32) error {

	return CodeError{
		Code: code,
		Msg:  message[code],
	}
}
func ParamErr() error {
	//return CodeError{
	//	Code: Param,
	//	Msg:  message[Param],
	//}

}

func SystemErr() error {
	//return CodeError{
	//	Code: SystemErrCode,
	//	Msg:  message[SystemErrCode],
	//}
	return remote.NewTransErrorWithMsg(SystemErrCode, message[SystemErrCode])
}
