package result

type ResponseSuccessBean struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}
type NullJson struct{}

func Success(data interface{}) *ResponseSuccessBean {
	return &ResponseSuccessBean{true, data}
}

type ResponseErrorBean struct {
	ErrorCode    uint32 `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

func Error(errCode uint32, errMsg string) *ResponseErrorBean {
	return &ResponseErrorBean{errCode, errMsg}
}
