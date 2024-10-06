package xlog

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"imlogic/common/types"

	"imlogic/internal/enum"
	"io"
	"log"
	"os"
	"path"
	"time"
)

//func InitHertzLog(serviceName string) *hertzzap.Logger {
//	logger := hertzzap.NewLogger(hertzzap.WithZapOptions(zap.WithFatalHook(zapcore.WriteThenPanic)))
//	logger.SetOutput(LogOutFile(serviceName))
//	return logger
//}

func LogOutFile(serviceName string) io.Writer {
	// 可定制的输出目录。
	var logFilePath string
	logFilePath = "./log/" + serviceName + "/"
	if err := os.MkdirAll(logFilePath, 0o777); err != nil {
		log.Println(err.Error())
		return nil
	}
	// 将文件名设置为日期
	logFileName := time.Now().Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
			return nil
		}
	}
	// 提供压缩和删除
	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    20,    // 一个文件最大可达 20M。
		MaxBackups: 5,     // 最多同时保存 5 个文件。
		MaxAge:     10,    // 一个文件最多可以保存 10 天。
		Compress:   false, // 用 gzip 压缩。
	}
	//iw := io.MultiWriter(lumberjackLogger, os.Stdout)
	iw := io.MultiWriter(lumberjackLogger)

	return iw
}

func SystemErrLog(logType enum.SystemLogType, content string, err error) {
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}
	log.Print(content, errMsg)
	go Send(map[string]interface{}{
		"logType": logType,
		"content": content,
		"errMsg":  errMsg,
	})
}

func UserLog(logType enum.SystemLogType, userId int64, format string, v ...interface{}) {
	content := fmt.Sprintf(format, v...)
	go Send(map[string]interface{}{
		"logType": logType,
		"content": content,
		"userId":  userId,
	})
}

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
