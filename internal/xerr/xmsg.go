package xerr

var message map[int32]string

func init() {
	// 全局码 错误消息
	message = make(map[int32]string)
	message[OK] = "success"
	message[ParamErrCode] = "param error"
	message[SystemErrCode] = "system error"
	message[RoleErrCode] = "role error"
	message[WarnCode] = "warn"

}
