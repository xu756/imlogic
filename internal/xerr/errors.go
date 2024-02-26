package xerr

import (
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

func (e CodeError) GetCode() int32 {
	return e.Code
}

func GetMsg(code int32) string {
	return message[code]
}

func ErrMsg(code int32) error {
	return CodeError{
		Code: code,
		Msg:  message[code],
	}

}
func ParamErr() error {
	return CodeError{
		Code: Param,
		Msg:  message[Param],
	}

}

func SystemErr() error {
	return CodeError{
		Code: SystemErrCode,
		Msg:  message[SystemErrCode],
	}
}

func DbErr(code int32, msg string) error {
	return CodeError{
		Code: code,
		Msg:  msg,
	}
}

func DbConnectErr() error {
	return DbErr(DbConnect, message[DbConnect])
}

func DbFindErr() error {
	//return CodeError{
	//	Code: DbFind,
	//	Msg:  message[DbFind],
	//}
	return DbErr(DbFind, message[DbFind])
}

func DbCreateErr() error {
	//return CodeError{
	//	Code: DbCreate,
	//	Msg:  message[DbCreate],
	//}
	return DbErr(DbCreate, message[DbCreate])
}

func DbUpdateErr() error {
	//return CodeError{
	//	Code: DbUpdate,
	//	Msg:  message[DbUpdate],
	//}
	return DbErr(DbUpdate, message[DbUpdate])
}

func DbDeleteErr() error {
	//return CodeError{
	//	Code: DbDelete,
	//	Msg:  message[DbDelete],
	//}
	return DbErr(DbDelete, message[DbDelete])

}
