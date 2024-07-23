package types

type SystemLogType string

const (
	// SuccessCode 成功
	SuccessCode SystemLogType = "Success"
	// DbErrCode 数据库错误
	DbErrCode SystemLogType = "DbErrCode"
	// RedisErrCode Redis错误
	RedisErrCode SystemLogType = "RedisErrCode"
	// SmsErrCode 短信错误
	SmsErrCode SystemLogType = "SmsErrCode"
	// EmailErrCode 邮件错误
	EmailErrCode SystemLogType = "EmailErrCode"
	// WxApiErrCode 微信API错误
	WxApiErrCode SystemLogType = "WxApiErrCode"
	// TtApiErrCode 头条API错误
	TtApiErrCode SystemLogType = "TtApiErrCode"
	// WxPayErrCode 微信支付错误
	WxPayErrCode SystemLogType = "WxPayErrCode"
	// UnmarshalErrCode 反序列化错误
	UnmarshalErrCode SystemLogType = "UnmarshalErrCode"
	// WechatLoginSuccess 微信登录成功
	WechatLoginSuccess SystemLogType = "WechatLogin"
	// ToutiaoLoginSuccess 头条登录成功
	ToutiaoLoginSuccess SystemLogType = "ToutiaoLoginSuccess"
	// WechatLoginFail 微信登录失败
	WechatLoginFail SystemLogType = "WechatLoginFail"
	// ToutiaoLoginFail 头条登录失败
	ToutiaoLoginFail SystemLogType = "ToutiaoLoginFail"
	// AccountLoginSuccess 账号登录成功
	AccountLoginSuccess SystemLogType = "AccountLoginSuccess"
	// AccountLoginFail 账号登录失败
	AccountLoginFail SystemLogType = "AccountLoginFail"
	// UserChangeAvatar 用户更改头像
	UserChangeAvatar SystemLogType = "UserChangeAvatar"
	// UserChangeNickname 用户更改昵称
	UserChangeNickname SystemLogType = "UserChangeNickname"
	// UserChangePassword 用户更改密码
	UserChangePassword SystemLogType = "UserChangePassword"
	// UserChangeMobile 用户更改手机号
	UserChangeMobile SystemLogType = "UserChangeMobile"
	// UploadImageFail 上传图片错误
	UploadImageFail SystemLogType = "UploadImageFail"
)

func (SystemLogType) Values() (kinds []string) {
	for _, s := range []SystemLogType{
		SuccessCode,
		DbErrCode,
		RedisErrCode,
		SmsErrCode,
		EmailErrCode,
		WxPayErrCode,
		UnmarshalErrCode,
		WxApiErrCode,
		TtApiErrCode,
		WechatLoginSuccess,
		ToutiaoLoginSuccess,
		WechatLoginFail,
		ToutiaoLoginFail,
		AccountLoginSuccess,
		AccountLoginFail,
		UserChangeAvatar,
		UserChangeNickname,
		UserChangePassword,
		UserChangeMobile,
		UploadImageFail,
	} {
		kinds = append(kinds, string(s))
	}
	return
}
