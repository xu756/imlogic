package result

import (
	"github.com/xu756/imlogic/internal/tool"
	"time"
)

type ResponseSuccessBean struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}
type NullJson struct{}

func successRes(data interface{}) *ResponseSuccessBean {
	return &ResponseSuccessBean{true, data}
}

type ResponseErrorBean struct {
	Success      bool      `json:"success"`
	ErrorCode    uint32    `json:"errorCode"`
	ErrorMessage string    `json:"errorMessage"`
	Timestamp    time.Time `json:"timestamp"`
}

func errorRes(errCode uint32, errMsg string) *ResponseErrorBean {
	return &ResponseErrorBean{false, errCode, errMsg, tool.TimeNow()}
}
