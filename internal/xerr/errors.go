package xerr

import (
	"errors"
	"fmt"
	types "imlogic/common/types"
	"imlogic/internal/xlog"

	"google.golang.org/grpc/codes"
)

/**
常用通用固定错误
*/

type CodeError struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func GetCodeError(code int32) CodeError {
	return CodeError{
		Code: code,
		Msg:  message[code],
	}
}

func (e CodeError) Error() string {
	return fmt.Sprintf("code:%d,msg:%s", e.Code, e.Msg)
}
func (e CodeError) RpcErr() error {
	return errors.New(e.Msg)
}

func (e CodeError) GetCode() int32 {
	return e.Code
}

func GetMsg(code int32) string {
	return message[code]
}

func WarnMsg(format string, v ...interface{}) error {
	msg := fmt.Sprintf(format, v...)
	return CodeError{
		Code: WarnCode,
		Msg:  msg,
	}
}

func RoleErr(format string, v ...interface{}) error {
	msg := fmt.Sprintf(format, v...)
	xlog.ErrLog(types.RedisErrCode, msg, nil)
	return CodeError{
		Code: RoleErrCode,
		Msg:  message[RoleErrCode],
	}
}

func DbErr(err error, format string, v ...interface{}) error {
	msg := fmt.Sprintf(format, v...)
	xlog.ErrLog(types.DbErrCode, msg, err)
	return CodeError{
		Code: SystemErrCode,
		Msg:  message[SystemErrCode],
	}
}

func RedisErr(err error, format string, v ...interface{}) error {
	msg := fmt.Sprintf(format, v...)
	xlog.ErrLog(types.RedisErrCode, msg, err)
	return CodeError{
		Code: SystemErrCode,
		Msg:  message[SystemErrCode],
	}
}

func UnmarshalErr(msg string) error {
	xlog.ErrLog(types.UnmarshalErrCode, msg, nil)
	return CodeError{
		Code: SystemErrCode,
		Msg:  message[SystemErrCode],
	}

}

func UploadImageErr(err error, format string, v ...interface{}) error {
	msg := fmt.Sprintf(format, v...)
	xlog.ErrLog(types.UploadImageFail, msg, err)
	return CodeError{
		Code: SystemErrCode,
		Msg:  message[SystemErrCode],
	}
}

func SmsErr(err error, format string, v ...interface{}) error {
	msg := fmt.Sprintf(format, v...)
	xlog.ErrLog(types.SmsErrCode, msg, err)
	return CodeError{
		Code: SystemErrCode,
		Msg:  message[SystemErrCode],
	}
}

func GetErrorCode(err error) (codes.Code, bool) {
	var xErr CodeError
	if !errors.As(err, &xErr) {
		return 0, false
	}
	return codes.Code(xErr.Code), true
}
