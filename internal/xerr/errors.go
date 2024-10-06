package xerr

import (
	"errors"
	"fmt"
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
	//msg := fmt.Sprintf(format, v...)
	//xlog.SystemErrLog(enum.RedisErrCode, msg, nil)
	return CodeError{
		Code: RoleErrCode,
		Msg:  message[RoleErrCode],
	}
}

func DbErr(err error, format string, v ...interface{}) error {
	//msg := fmt.Sprintf(format, v...)
	//xlog.SystemErrLog(enum.DbErrCode, msg, err)
	return CodeError{
		Code: SystemErrCode,
		Msg:  message[SystemErrCode],
	}
}

func RedisErr(err error, format string, v ...interface{}) error {
	//msg := fmt.Sprintf(format, v...)
	//xlog.SystemErrLog(enum.RedisErrCode, msg, err)
	return CodeError{
		Code: SystemErrCode,
		Msg:  message[SystemErrCode],
	}
}

func WxApiErr(msg string, err error) error {
	//xlog.SystemErrLog(enum.WxApiErrCode, msg, err)
	return CodeError{
		Code: SystemErrCode,
		Msg:  message[SystemErrCode],
	}
}

func TtApiErr(msg string, err error) error {
	//xlog.SystemErrLog(enum.TtApiErrCode, msg, err)
	return CodeError{
		Code: SystemErrCode,
		Msg:  message[SystemErrCode],
	}
}

func EmailErr(msg string, err error) error {
	//xlog.SystemErrLog(enum.EmailErrCode, msg, err)
	return CodeError{
		Code: SystemErrCode,
		Msg:  message[SystemErrCode],
	}

}

func SmsErr(msg string, err error) error {
	//xlog.SystemErrLog(enum.SmsErrCode, msg, err)
	return CodeError{
		Code: SystemErrCode,
		Msg:  message[SystemErrCode],
	}
}

func UnmarshalErr(msg string) error {
	//xlog.SystemErrLog(enum.UnmarshalErrCode, msg, nil)
	return CodeError{
		Code: SystemErrCode,
		Msg:  message[SystemErrCode],
	}

}

func UploadImageErr(err error, format string, v ...interface{}) error {
	//msg := fmt.Sprintf(format, v...)
	//xlog.SystemErrLog(enum.UploadImageFail, msg, err)
	return CodeError{
		Code: SystemErrCode,
		Msg:  message[SystemErrCode],
	}
}
