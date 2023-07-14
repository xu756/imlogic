package xerr

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

/**
常用通用固定错误
*/

type CodeError struct {
	errCode uint32
	errMsg  string
}

// GetErrCode 返回给前端的错误码
func (e *CodeError) GetErrCode() uint32 {
	return e.errCode
}

// GetErrMsg 返回给前端显示端错误信息
func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMsg)
}

// NewSystemError 系统错误
func NewSystemError(err interface{}) *CodeError {
	logx.Error("【系统错误】%v", err)
	return &CodeError{errCode: NOTIFICATION, errMsg: "系统错误,请联系管理员"}
}

// NewDbErr 数据库错误
func NewDbErr(msg string, err interface{}) *CodeError {
	logx.Error("【数据库错误  %s】---%v", msg, err)
	return &CodeError{errCode: ERRORMESSAGE, errMsg: "系统错误,请联系管理员"}
}

// NewMsgError 通用错误（返回给用户的）
func NewMsgError(msg string) *CodeError {
	return &CodeError{errCode: WarnMessage, errMsg: msg}
}

// LogOut  重新登录
func LogOut() *CodeError {
	return &CodeError{errCode: REDIRECT, errMsg: "身份验证失败，请重新登录"}

}
