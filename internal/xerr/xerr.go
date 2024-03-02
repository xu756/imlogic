package xerr

// 全局码
const (
	OK = 200 + iota
	SystemErrCode
	DbFind
	DbCreate
	DbUpdate
	DbDelete
	Param
	DbConnect
	RedisConnectErr
	RoleErr
	RedisErr
)

// 用户模块码
const (
	UserExist = 220 + iota
	UserPasswordErr
	UserLoginOut
	UserExpired
	UserDeleteErr
	UserNotExist
	UserRegisterErr
	UserUpdateErr
	UserMobileNotExist
)

// JWT 模块码
const (
	JwtCreateErr = 230 + iota
	JwtParseErr
	JwtAuthEmpty
	JwtAuthErr
	CaptchaErr
	CaptchaExpired
	NotSameUser
)

// Send 发送信息
const (
	SendEmailErr = 250 + iota
	SendSmsErr
	CaptchaNotExist
)
