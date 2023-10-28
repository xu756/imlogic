package xerr

// 全局码
const (
	OK = 200 + iota
	SystemErrCode
	DbErrCode
	Param
	DbConnectErr
	RedisConnectErr
)

// 用户模块码
const (
	UserExist = 210 + iota
	UserLoginOut
	UserExpired
	UserMobileExist
	UserNotExist
	UserRegisterErr
	UserUpdateErr
)

// JWT 模块码
const (
	JwtCreateErr = 220 + iota
	JwtParseErr
	JwtAuthEmpty
	JwtAuthErr
)
