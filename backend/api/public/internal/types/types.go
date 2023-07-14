// Code generated by goctl. DO NOT EDIT.
package types

type GetCodeReq struct {
	SessionId string `json:"session_id"` //会话id
	Sign      string `json:"sign"`       //签名
	Timestamp int64  `json:"timestamp"`  //时间戳
}

type GetCodeRes struct {
	Expire int64  `json:"expire"` //过期时间
	Img    string `json:"img"`    //图片base64
}

type LoginReq struct {
	Username  string `json:"username"`   //用户名
	Password  string `json:"password"`   //密码
	SessionId string `json:"session_id"` //会话id
	Code      string `json:"code"`       //验证码
}

type LoginRes struct {
	Token  string `json:"token"`  //token
	Expire int64  `json:"expire"` //过期时间
}
