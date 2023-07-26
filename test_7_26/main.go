package main

import (
	"fmt"
)

// type person struct {
// 	name  string
// 	age   int
// 	hobby []string
// }
//
// func main() {
// 	// fmt.Println("jsakdl")
// 	var p person
// 	p.age = 10
// 	p.hobby = []string{"asdfsa", "sfsad"}
// 	p.name = "jj"
// 	fmt.Printf("%T\n", p) // main.person
// 	fmt.Println(p)
//
// 	var p1 = new(person)
// 	p1.age = 111
//
// 	var p2 = person{
// 		name: "safsda",
// 	}
// 	fmt.Println(p2)
//
// 	var p3 = person{
// 		"asjkfsad",
// 		12,
// 		[]string{"sdafsadk"},
// 	}
// 	fmt.Println(p3)
//
// 	p4 := person{
// 		name: "asfsa",
// 	}
// 	fmt.Println(p4)
//
// }

type dog struct {
	name string
	age  int
}

// 构造函数
func newDog(name string, age int) *dog {
	return &dog{
		name: name,
		age:  age,
	}
}

// 成员函数
func (d dog) bark() {
	fmt.Printf("%s:汪汪汪~~~", d.name)
}
func main() {
	// d1 := dog{
	// 	name: "aaa",
	// 	age:  1,
	// }
	d1 := newDog("aaa", 1)
	fmt.Printf("%T\n", d1)
	d1.bark()

}
