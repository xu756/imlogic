package xerr

var message map[uint32]string

func init() {
	// 全局码 错误消息
	message = make(map[uint32]string)
	message[OK] = "ok"
	message[SystemErrCode] = "系统错误，请联系管理员"
	message[DbErrCode] = "数据库错误"
	message[Param] = "参数错误"
	message[DbConnectErr] = "数据库连接失败"
	message[RedisConnectErr] = "redis连接失败"
	message[RoleErr] = "角色错误"
	message[ExportExcelErr] = "导出excel错误"

	// 用户模块码 错误消息
	message[UserExist] = "账号已存在"
	message[UserPasswordErr] = "账号密码错误"
	message[UserLoginOut] = "账号未登录"
	message[UserExpired] = "账号登录过期"
	message[UserDeleteErr] = "删除用户失败,请联系管理员"
	message[UserNotExist] = "账号不存在"
	message[UserRegisterErr] = "注册失败"
	message[UserUpdateErr] = "更新用户信息失败"
	message[UserEmailExist] = "邮箱已存在"
	message[UserEmailNotExist] = "邮箱不存在"

	// jwt
	message[JwtCreateErr] = "jwt生成失败"
	message[JwtParseErr] = "jwt解析失败"
	message[JwtAuthEmpty] = "请求头中auth为空"
	message[JwtAuthErr] = "请求头中auth格式错误"
	message[CaptchaErr] = "验证码错误"
	message[CaptchaExpired] = "验证码过期"
	message[NotSameUser] = "不是同一个用户"

	// Subject
	message[SubjectExist] = "课程已存在"
	message[SubjectFindErr] = " 课程查找错误"
	message[SubjectUpdateErr] = "课程更新失败"

	// Mark
	message[MarkExist] = "成绩已存在"
	message[MarkCreateErr] = "成绩创建失败"
	message[MarkFindErr] = "成绩查找失败"

	// Emails
	message[SendEmailErr] = "邮件发送失败"
	message[SendSmsErr] = "短信发送失败"
	message[CaptchaNotExist] = "验证码已过期"

	// File
	message[FileSaveErr] = "文件保存失败"
	message[FileNotExist] = "文件不存在"
	message[UploadFileErr] = "文件上传失败，请重试"
	message[CreatDirErr] = "创建文件夹失败,请联系管理员"
	message[TeacherConNotDown] = "不支持老师下载学生成绩"
	message[FileTypeErr] = "文件类型错误，只支持xlsx"
	message[ExcelReadErr] = "Excel文件读取失败"
	message[ExcelFileErr] = "Excel文件错误，请重新下载模板文件再重新上传"
}
