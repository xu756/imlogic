package xerr

// 全局错误码
const (
	SILENT       uint32 = 0 // 不提示
	WarnMessage  uint32 = 1 // 警告提示
	ERRORMESSAGE uint32 = 2 // 报错提示
	NOTIFICATION uint32 = 3 // 通知
	REDIRECT     uint32 = 9 // 页面跳转

)
