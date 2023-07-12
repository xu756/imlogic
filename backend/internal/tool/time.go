package tool

import (
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

var cstSh *time.Location

func init() {
	var err error
	cstSh, err = time.LoadLocation("Asia/Shanghai") //上海
	if err != nil {
		logx.Errorf("时区设置失败，错误信息: %v", err)
	}
}

// TimeNowInTimeZone
//
//	获得当前时间，默认上海时间
func TimeNowInTimeZone() time.Time {
	return time.Now().In(cstSh)
}
