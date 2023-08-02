package mylog

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 开多个 goroutine 并发的向文件中写！！否则串行化读写磁盘很慢！！

// 向文件中打印
type LogFile struct {
	level       LogLevel
	filePath    string
	fileName    string
	fileDes     *os.File
	errFileDes  *os.File // error以上级别错误再单独存一分
	maxFileSize int64

	// 加入chan
	logCh chan *LogMsg
}

// 输出内容
type LogMsg struct {
	level     LogLevel
	msg       string
	funcName  string
	fileName  string
	timestame string
	line      int
}

// 构造函数
func NewLogFile(levelStr, filePath, fileName string, maxFileSize int64) *LogFile {
	level, err := parseLogLevel(&levelStr)
	if err != nil {
		panic(err)
	}
	fullFile := path.Join(filePath, fileName)
	fileDes, err := os.OpenFile(fullFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("OpenFile error:%#v\n", err)
	}

	errFileDes, err := os.OpenFile(fullFile+".err", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("OpenFile error:%#v\n", err)
	}

	log := &LogFile{
		level:       level,
		filePath:    filePath,
		fileName:    fileName,
		fileDes:     fileDes,
		errFileDes:  errFileDes,
		maxFileSize: maxFileSize,
		logCh:       make(chan *LogMsg, 10000),
	}
	go log.logPrintBackup()
	return log
}

func (l *LogFile) init() {
}

func (l *LogFile) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("file.Stat() error:%#v\n", err)
		return false
	}
	return fileInfo.Size() >= l.maxFileSize
}

func (l *LogFile) splitFile(file *os.File) *os.File {
	// 需要切割文件，就是把日志打印到新的文件里
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("file.Stat() error:%#v", err)
		return nil
	}
	fileName := fileInfo.Name()

	file.Close()

	// 备份一下这个已经满了的文件
	now := time.Now().Format("20060102150405")
	logName := path.Join(l.filePath, fileName)
	// 拼接一个新的名字
	newLogName := fmt.Sprintf("%s.bak%s", logName, now)
	os.Rename(logName, newLogName)
	newFile, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("os.OpenFile error:%#v\n", err)
		return nil
	}
	return newFile
}

func (l *LogFile) logPrintBackup() {
	for {
		if l.checkSize(l.fileDes) {
			newFile := l.splitFile(l.fileDes)
			if newFile != nil {
				l.fileDes = newFile
			}
		}

		select {
		// 再从管道里取值
		case logCon := <-l.logCh:
			logInfo := fmt.Sprintf("[%s] [%s] [%s:%s:%d] %s\n", logCon.timestame,
				getLogString(logCon.level), logCon.funcName, logCon.fileName,
				logCon.line, logCon.msg)

			fmt.Fprint(l.fileDes, logInfo)

			if logCon.level >= ERROR {
				if l.checkSize(l.errFileDes) {
					newFile := l.splitFile(l.errFileDes)
					if newFile != nil {
						l.errFileDes = newFile
					}
				}
				fmt.Fprint(l.errFileDes, logInfo)
			}
		default:
			time.Sleep(time.Second)
		}
	}
}

func (l *LogFile) logPrint(logLevel LogLevel, msg *string, args ...interface{}) {
	now := time.Now()
	*msg = fmt.Sprintf(*msg, args...)
	// 这个3表示主函数那里要有三层函数调用才能调到 runtime.Caller
	funcName, fileName, lineNum := getInfo(3)

	// 只需要把要写的内容放到管道里交给后台另一个goroutine
	logCon := &LogMsg{
		level:     logLevel,
		msg:       *msg,
		funcName:  funcName,
		fileName:  fileName,
		timestame: now.Format("2006-01-02 15:04:05"),
		line:      lineNum,
	}
	// 为防止阻塞，这里用一下select
	select {
	// 防止这里阻塞住了
	case l.logCh <- logCon:

	// 直接把这条日志扔了
	default:
	}
}

// 成员函数,判断该条日志是否符合打印要求
func (l *LogFile) enable(logLevel LogLevel) bool {
	return logLevel >= l.level
}

func (l *LogFile) Debug(msg string, args ...interface{}) {
	if l.enable(DEBUG) {
		l.logPrint(DEBUG, &msg, args...)
	}
}

func (l *LogFile) Info(msg string, args ...interface{}) {
	if l.enable(INFO) {
		l.logPrint(INFO, &msg, args...)
	}
}

func (l *LogFile) Warning(msg string, args ...interface{}) {
	if l.enable(WARNING) {
		l.logPrint(WARNING, &msg, args...)
	}
}

func (l *LogFile) Error(msg string, args ...interface{}) {
	if l.enable(ERROR) {
		l.logPrint(ERROR, &msg, args...)
	}
}

func (l *LogFile) Fatal(msg string, args ...interface{}) {
	if l.enable(FATAL) {
		l.logPrint(FATAL, &msg, args...)
	}
}
