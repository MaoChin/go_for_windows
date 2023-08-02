package main

import (
	"time"

	mylog "github.com/go_study/mylog_v2.0"
)

func main() {
	// 参数表示只打印比该级别高的信息
	log := mylog.NewLogFile("DEBUG", "./", "test.log", 2*1024)
	for {
		id := 10
		log.Debug("debug, id:%d", id)
		log.Info("info")
		log.Warning("warning")
		log.Error("error")
		log.Fatal("fatal")
		time.Sleep(time.Second)
	}
}
