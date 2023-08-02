package mylog

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

type LogLevel uint8

const (
	UNKNOWN LogLevel = iota
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

// string -> LogLevel
func parseLogLevel(str *string) (LogLevel, error) {
	s := strings.ToUpper(*str)
	switch s {
	case "DEBUG":
		return DEBUG, nil
	case "INFO":
		return INFO, nil
	case "WARNING":
		return WARNING, nil
	case "ERROR":
		return ERROR, nil
	case "FATAL":
		return FATAL, nil
	default:
		return UNKNOWN, nil
	}
}

// LogLevel -> string
func getLogString(logLevel LogLevel) string {
	switch logLevel {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "DEBUG"
	}
}

func getInfo(skip int) (funcName string, fileName string, lineNum int) {
	// 获取调用者的函数名、文件名、行号
	pc, fileName, lineNum, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime.Caller error!")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	funcName = strings.Split(funcName, ".")[1]
	fileName = path.Base(fileName)
	return
}
