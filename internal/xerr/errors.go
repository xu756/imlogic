package xerr

import "fmt"

/**
常用通用固定错误
*/

type CodeError struct {
	Code uint32
	Msg  string
}

func (c CodeError) Error() string {
	return fmt.Sprintf("code:%d msg:%s", c.Code, c.Msg)
}

func (c CodeError) GetCode() uint32 {
	return c.Code
}
func (c CodeError) GetMsg() string {
	return c.Msg
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
func DbErr() error {
	return CodeError{
		Code: DbErrCode,
		Msg:  message[DbErrCode],
	}

}

func SystemErr() error {
	return CodeError{
		Code: SystemErrCode,
		Msg:  message[SystemErrCode],
	}
}

func ParamErr() error {
	return CodeError{
		Code: Param,
		Msg:  message[Param],
	}
}
func UserExpire() error {
	return CodeError{
		Code: UserExpired,
		Msg:  message[UserExpired],
	}
}
