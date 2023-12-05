package xerr

var message map[int32]string

func init() {
	// 全局码 错误消息
	message = make(map[int32]string)
	message[OK] = "ok"
	message[SystemErrCode] = "系统错误，请联系管理员"
	message[DbErrCode] = "数据库错误"
	message[DbFindErr] = "数据库查询错误"
	message[DbCreateErr] = "数据库创建错误"
	message[DbUpdateErr] = "数据库更新错误"
	message[DbDeleteErr] = "数据库删除错误"
	message[Param] = "参数错误"
	message[DbConnectErr] = "数据库连接失败"
	message[RedisConnectErr] = "redis连接失败"
	message[RoleErr] = "角色错误"

	// 用户模块码 错误消息
	message[UserExist] = "账号已存在"
	message[UserPasswordErr] = "密码错误"
	message[UserLoginOut] = "账号未登录"
	message[UserExpired] = "账号登录过期"
	message[UserDeleteErr] = "删除用户失败,请联系管理员"
	message[UserNotExist] = "账号不存在"
	message[UserRegisterErr] = "注册失败"
	message[UserUpdateErr] = "更新用户信息失败"
	message[UserMobileNotExist] = "手机号不存在"
	// jwt模块码 错误消息
	message[JwtCreateErr] = "jwt生成失败"
	message[JwtParseErr] = "jwt解析失败"
	message[JwtAuthEmpty] = "请求头中auth为空"
	message[JwtAuthErr] = "请求头中auth格式错误"
	message[CaptchaErr] = "验证码错误"
	message[CaptchaExpired] = "验证码过期"
	message[NotSameUser] = "不是同一个用户"
	// Emails
	message[SendEmailErr] = "邮件发送失败"
	message[SendSmsErr] = "短信发送失败"
	message[CaptchaNotExist] = "验证码已过期"
	//Upgrade
	message[UpgradeErr] = "websocket升级失败"
}
