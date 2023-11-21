package tool

import (
	"time"
)

//func init() {
//	var err error
//	cstSh, err = time.LoadLocation("Asia/Shanghai") //上海
//	if err != nil {
//		log.Printf("时区设置失败，错误信息: %v\n", err)
//	}
//}

// TimeNow 获取当前时间
func TimeNow() time.Time {
	return time.Now()
}

// TimeNowUnixMilli 获取当前时间毫秒
func TimeNowUnixMilli() int64 {
	return TimeNow().UnixMilli()
}

// TimeNowString 返回当前时间的字符串
func TimeNowString() string {
	return TimeNow().Format("2006-01-02 15:04:05")
}
