package xerr

// 返回给用户的
const (
	OK = 200 + iota
	ParamErrCode
	SystemErrCode
	RoleErrCode
	WarnCode
)

// 系统错误 弹窗提示
const (
	DbErrCode = 4001 + iota
	RedisErrCode
	SmsErrCode
	EmailErrCode
	UnmarshalErrCode
)
