package xerr

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/remote"
)

/**
常用通用固定错误
*/

type CodeError struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func (e CodeError) Error() string {
	return fmt.Sprintf("code:%d,msg:%s", e.Code, e.Msg)
}

func (e CodeError) GetCode() int32 {
	return e.Code
}

func GetMsg(code int32) string {
	return message[code]

}

func NewErr(code int32, msg string) error {
	return CodeError{
		Code: code,
		Msg:  msg,
	}
}

func NewSprintfErr(code int32, format string, a ...interface{}) error {
	err := CodeError{
		Code: code,
		Msg:  fmt.Sprintf(format, a...),
	}
	return remote.NewTransError(code, err)
}

func ErrMsg(code int32) error {

	err := CodeError{
		Code: code,
		Msg:  message[code],
	}
	return remote.NewTransError(code, err)

}
func ParamErr() error {
	err := CodeError{
		Code: Param,
		Msg:  message[Param],
	}
	return remote.NewTransError(Param, err)

}

func SystemErr() error {
	err := CodeError{
		Code: SystemErrCode,
		Msg:  message[SystemErrCode],
	}
	return remote.NewTransError(SystemErrCode, err)
}
