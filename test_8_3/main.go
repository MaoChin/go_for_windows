package main

import (
	"flag"
	"fmt"
)

func main() {
	// 创建标记位参数,而后在命令行中指定name的具体值，就可以解析到name里！
	// 若不指定，name就是aaa的地址！name是个指针！！
	name := flag.String("name", "aaa", "姓名") // 最后一个参数是提示信息
	age := flag.Int("age", 10, "年龄")

	var num int
	flag.IntVar(&num, "num", 111, "数字")

	flag.Parse()
	fmt.Println(*name)
	fmt.Println(*age)

	fmt.Println(num)
}
