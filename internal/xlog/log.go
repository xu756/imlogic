package xlog

import (
	types "imlogic/common/types"
	"log"
)

func ErrLog(logType types.SystemLogType, content string, err error) {
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}
	log.Print(content, errMsg)
}

func Printf(logType types.SystemLogType, format string, v ...any) {
	log.Printf(format, v...)
}

func Fatal(logType types.SystemLogType, v ...any) {
	log.Fatal(v...)
}
