package main

import (
	"fmt"
	"os"
)

func main() {
	// 先初始化一下
	studentManager := newstudentManager()
	for {
		fmt.Println("学生管理系统")
		fmt.Print(` 
			1. 查看所有学生
			2. 添加学生
			3. 删除学生
			4. 修改学生信息
			5. 退出
		`)
		var input int
		fmt.Print("请输入你的操作：")
		fmt.Scan(&input)
		switch input {
		case 1:
			studentManager.showStudent()
		case 2:
			studentManager.addStudent()
		case 3:
			studentManager.deleteStudent()
		case 4:
			studentManager.changeStudent()
		case 5:
			fmt.Println("退出！")
			os.Exit(1)
		default:
			fmt.Println("你的输入错误，请重新输入！")
		}
	}
}
