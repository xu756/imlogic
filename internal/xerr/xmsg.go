package xerr

var message map[uint32]string

func init() {
	// 全局码 错误消息
	message = make(map[uint32]string)
	message[OK] = "ok"
	message[SystemErrCode] = "系统错误"
	message[DbErrCode] = "数据库错误"
	message[Param] = "参数错误"
	message[DbConnectErr] = "数据库连接失败"
	message[RedisConnectErr] = "redis连接失败"

	// 用户模块码 错误消息
	message[UserExist] = "用户已存在"
	message[UserLoginOut] = "用户未登录"
	message[UserExpired] = "用户登录过期"
	message[UserMobileExist] = "手机号已存在"
	message[UserNotExist] = "用户不存在"
	message[UserRegisterErr] = "注册失败"
	message[UserUpdateErr] = "更新用户信息失败"
	// jwt
	message[JwtCreateErr] = "jwt生成失败"
	message[JwtParseErr] = "jwt解析失败"
	message[JwtAuthEmpty] = "请求头中auth为空"
	message[JwtAuthErr] = "请求头中auth格式错误"

}
