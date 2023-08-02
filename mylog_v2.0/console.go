package mylog

import (
	"fmt"
	"time"
)

// 向终端中打印
type LogConsole struct {
	level LogLevel
}

// 构造函数
func NewLogConsole(str string) *LogConsole {
	level, err := parseLogLevel(&str)
	if err != nil {
		panic(err)
	}
	return &LogConsole{
		level: level,
	}
}

func (l *LogConsole) logPrint(logLevel LogLevel, msg *string, args ...interface{}) {
	now := time.Now()
	*msg = fmt.Sprintf(*msg, args...)
	// 这个3表示主函数那里要有三层函数调用才能调到 runtime.Caller
	funcName, fileName, lineNum := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05.000"),
		getLogString(logLevel), funcName, fileName, lineNum, *msg)
}

// 成员函数,判断该条日志是否符合打印要求
func (l *LogConsole) enable(logLevel LogLevel) bool {
	return logLevel >= l.level
}

func (l *LogConsole) Debug(msg string, args ...interface{}) {
	if l.enable(DEBUG) {
		l.logPrint(DEBUG, &msg, args...)
	}
}

func (l *LogConsole) Info(msg string, args ...interface{}) {
	if l.enable(INFO) {
		l.logPrint(INFO, &msg, args...)
	}
}

func (l *LogConsole) Warning(msg string, args ...interface{}) {
	if l.enable(WARNING) {
		l.logPrint(WARNING, &msg, args...)
	}
}

func (l *LogConsole) Error(msg string, args ...interface{}) {
	if l.enable(ERROR) {
		l.logPrint(ERROR, &msg, args...)
	}
}

func (l *LogConsole) Fatal(msg string, args ...interface{}) {
	if l.enable(FATAL) {
		l.logPrint(FATAL, &msg, args...)
	}
}
