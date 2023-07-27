package main

import (
	"fmt"
	"os"
)

type person struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

// func main() {
// json
// p1 := person{
// 	Name: "safsad",
// 	Id:   1,
// }
// j, err := json.Marshal(p1)
// if err != nil {
// 	fmt.Printf("Marshal err:%s\n", err)
// 	return
// }
// fmt.Println(string(j))

// str := `{"name":"safs","id":11}`
// var p2 person
// // 反序列化
// err1 := json.Unmarshal([]byte(str), &p2)
// if err1 != nil {
// 	fmt.Printf("Marshal err:%s\n", err1)
// 	return
// }
// fmt.Println(p2)

// var s1 string
// var s2 string
// fmt.Scan(&s1)
// fmt.Scan(&s2)
// fmt.Print(s1, s2)
// }

// 在函数传参时就可以使用speaker这个类型来当形参
// 实参可以是实现了speak方法的任意类
type speaker interface {
	// 具体的方法声明,所有有这些方法的类都可以通过speaker这个interface来统一
	speak()
	// ...
}

// cat类有speak这个方法
type cat struct {
	name string
}

func (c cat) speak() {
	// ...
}

// dog类有speak这个方法
type dog struct {
	name string
}

func (d dog) speak() {
	// ...
}

// pig类有speak这个方法
type pig struct {
	name string
}

func (p pig) speak() {
	// ...
}

func amimalSpeak(x speaker) {
	x.speak()
}

// func main() {
// 	c := cat{name: "aaa"}
// 	d := dog{name: "sss"}
// 	p := pig{name: "ddd"}
// 	// 实参可以是实现了speak方法的任何类
// 	amimalSpeak(c)
// 	amimalSpeak(p)
// 	amimalSpeak(d)
//
// 	// 可以直接赋值！speaker就是抽象类型
// 	var am speaker
// 	fmt.Printf("%T\n", am) // main.speaker
// 	am = c
// 	fmt.Printf("%T\n", am) // main.cat
// 	fmt.Println(am)
// 	am = d
// 	fmt.Printf("%T\n", am) // main.dog
// 	fmt.Println(am)
// 	am = p
// 	fmt.Printf("%T\n", am) // main.pig
// 	fmt.Println(am)
// }

func main() {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Printf("open file error:%v\n", err)
		return
	}
	// 关闭文件描述符，并设置成defer,这也是defer的应用场景
	defer file.Close()
	tmp := make([]byte, 128)
	// 读到tmp里，返回值n表示实际读到的字节数
	n, err := file.Read(tmp)
	if err != nil {
		// ...
	}
	fmt.Println(n)

}
