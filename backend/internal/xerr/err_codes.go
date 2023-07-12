package xerr

// OK 成功返回
const OK uint32 = 200

// 全局错误码
const (
	SystemError uint32 = 207 // 系统错误
	MsgError    uint32 = 112 // 需要提醒用户的错误
	LogOutError uint32 = 113 // 重新登录
)
