package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int32  `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"passward"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int32  `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}

var fileData []string

func readFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Open file error: %#v\n", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		// 去掉每一行首尾的空格
		line = strings.TrimSpace(line)
		if err == io.EOF {
			if len(line) > 0 && line[0] != '[' && line[0] != '#' &&
				line[0] != ';' && line[0] != '\n' {
				// 有效数据
				fileData = append(fileData, line)
			}
			break
		}
		if err != nil {
			fmt.Printf("ReadString error:%#v\n", err)
			return
		}
		if len(line) > 0 && line[0] != '[' && line[0] != '#' &&
			line[0] != ';' && line[0] != '\r' {
			// 有效数据
			fileData = append(fileData, line)
		}
	}
	fmt.Println(fileData)
}

func loadIni(data interface{}) (err error) {
	// 1. 参数校验，v必须是结构体指针
	v := reflect.TypeOf(data)
	if v.Kind() != reflect.Ptr {
		// 格式化输出到err
		// err = fmt.Errorf("参数错误：不是指针类型")
		// 使用这个也可以
		err = errors.New("参数错误：不是指针类型")
		return
	}
	// 判断值的类型
	if v.Elem().Kind() != reflect.Struct {
		err = errors.New("参数错误：不是结构体指针")
		return
	}
	// 2. 遍历得到的有效数据fileData
}
func main() {
	var mc MysqlConfig

	// 1. 读 .ini 文件并去掉注释和[]这些不需要的数据,并对有效数据按 '=' 分割
	readFile("./conf.ini")
	// 3. 进行赋值，要传地址！！
	err := loadIni(&mc)
	fmt.Println(mc)
}
