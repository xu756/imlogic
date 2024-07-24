package logic

type LoginReq struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Code      string `json:"code"`
	Mobile    string `json:"mobile"`
	Captcha   string `json:"captcha"`
	SessionId string `json:"session_id,required"`
}

type SendCaptchaReq struct {
	Mobile    string `json:"mobile,required"`
	SessionId string `json:"session_id,required"`
}

type SendCaptchaRes struct {
	Expire int64 `json:"expire"`
}
