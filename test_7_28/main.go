package main

import (
	"fmt"
	"time"
)

func main() {
	// file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND,
	// 	0666)
	// if err != nil {
	// 	fmt.Printf("OpenFile error:%#v\n", err)
	// 	return
	// }
	// defer file.Close()
	// file.Write([]byte("ada\n"))
	// file.WriteString("asdad")

	// writer := bufio.NewWriter(file)
	// writer.WriteString("fasfas\n")
	// writer.Flush()

	// var a interface{}
	// var b int8
	// b = 10

	// a = b
	// c, ok := a.(int16)
	// if !ok {
	// 	fmt.Println("a is not a int16")
	// } else {
	// 	fmt.Println("a is a int8")
	// 	fmt.Println(c)
	// }

	now := time.Now()
	// time := time.Unix()
	fmt.Printf("%T\n", now) // time.Time
	fmt.Println(now)
}
