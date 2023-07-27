package main

import (
	"fmt"
	"os"
)

type student struct {
	name string
	id   int
}

var allstudent map[int]*student

// 构造
func newStudent(name string, id int) *student {
	return &student{
		name: name,
		id:   id,
	}
}
func showStudent() {
	for k, v := range allstudent {
		fmt.Printf("学号：%d, 姓名：%s\n", k, v.name)
	}
	fmt.Println()
}
func addStudent() {
	var (
		id   int
		name string
	)
	fmt.Printf("请输入要添加的学生姓名：")
	fmt.Scanln(&name)
	fmt.Printf("请输入要添加的学生学号：")
	fmt.Scanln(&id)

	_, ok := allstudent[id]
	if !ok {
		newS := newStudent(name, id)
		allstudent[id] = newS
		fmt.Println("添加成功！")
		showStudent()
	} else {
		fmt.Println("该学生已经存在！")
		showStudent()
	}
}
func deleteStudent() {
	var id int
	fmt.Printf("请输入要添加的学生学号：")
	fmt.Scanln(&id)
	_, ok := allstudent[id]
	if !ok {
		fmt.Println("要删除的学生不存在！")
		showStudent()
	} else {
		delete(allstudent, id)
		fmt.Println("删除成功！")
		showStudent()
	}
}

func main() {
	allstudent = make(map[int]*student, 100)
	for {
		fmt.Println("学生管理系统")
		fmt.Print(` 
    1. 查看所有学生
		2. 添加学生
    3. 删除学生
    4. 退出
    `)
		var input int
		fmt.Print("请输入你的操作：")
		fmt.Scanln(&input)
		switch input {
		case 1:
			showStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			fmt.Println("退出！")
			os.Exit(1)
		default:
			fmt.Println("你的输入错误，请重新输入！")
		}
	}
}
