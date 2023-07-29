package main

import (
	"fmt"
	"reflect"
	"sync"
)

func f(x interface{}) {
	// 获取空接口的具体值
	v := reflect.ValueOf(x)
	// 取出v的值并修改
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(100)
	}
}

// func main() {
// 	var a int64
// 	a = 10
// 	// 要传地址
// 	f(&a)
// }

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

// func main() {
// 	stu1 := student{
// 		Name:  "小王子",
// 		Score: 90,
// 	}
//
// 	t := reflect.TypeOf(stu1)
// 	fmt.Println(t.Name(), t.Kind()) // student struct
// 	// 通过for循环遍历结构体的所有字段信息
// 	for i := 0; i < t.NumField(); i++ {
// 		field := t.Field(i)
// 		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
// 	}
//
// 	// 通过字段名获取指定结构体字段信息
// 	if scoreField, ok := t.FieldByName("Score"); ok {
// 		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
// 	}
// }

func f2() {
	fmt.Println("aaa")
}

var wg sync.WaitGroup

var a chan int

func main() {
	// str := "111"
	// a, _ := strconv.ParseInt(str, 10, 64)
	// fmt.Printf("%v, %T\n", a, a)

	// go f2()

	// time.Sleep(time.Second)
	// fmt.Println("bbb")
	// fmt.Println(runtime.NumCPU())
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		fmt.Println(i)
	// 	}()
	// }
	// fmt.Println("aaaaaaaaaaa")
	// wg.Wait()

	a = make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-a
		fmt.Println("接收", x)
	}()
	a <- 10
	fmt.Println("发送10")
	wg.Wait()
}
