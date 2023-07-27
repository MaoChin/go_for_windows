package main

import "fmt"

type student struct {
	name string
	id   int
}

// 使用一个学生管理的结构体，包含所有学生和方法
type studentManager struct {
	allStudent map[int]*student
}

// 构造
func newstudentManager() *studentManager {
	tmp := make(map[int]*student, 100)
	return &studentManager{
		allStudent: tmp,
	}
}

// 构造
func newStudent(name string, id int) *student {
	return &student{
		name: name,
		id:   id,
	}
}
func (s *studentManager) showStudent() {
	for k, v := range s.allStudent {
		fmt.Printf("学号：%d, 姓名：%s\n", k, v.name)
	}
	fmt.Println()
}
func (s *studentManager) addStudent() {
	var (
		id   int
		name string
	)
	fmt.Printf("请输入要添加的学生姓名：")
	fmt.Scan(&name)
	fmt.Printf("请输入要添加的学生学号：")
	fmt.Scan(&id)

	_, ok := s.allStudent[id]
	if !ok {
		newS := newStudent(name, id)
		s.allStudent[id] = newS
		fmt.Println("添加成功！")
		s.showStudent()
	} else {
		fmt.Println("该学生已经存在！")
		s.showStudent()
	}
}
func (s *studentManager) changeStudent() {
	s.showStudent()
	fmt.Print("请输入要修改的学生ID:")
	var Id int
	fmt.Scan(&Id)
	_, ok := s.allStudent[Id]
	if !ok {
		fmt.Println("要修改的学生Id不存在!")
		return
	}
	fmt.Print("请输入修改后的名字：")
	var newName string
	fmt.Scan(&newName)
	s.allStudent[Id].name = newName
	fmt.Println("修改成功")
	s.showStudent()
}

func (s *studentManager) deleteStudent() {
	var id int
	fmt.Printf("请输入要删除的学生学号：")
	fmt.Scan(&id)
	_, ok := s.allStudent[id]
	if !ok {
		fmt.Println("要删除的学生不存在！")
		s.showStudent()
	} else {
		delete(s.allStudent, id)
		fmt.Println("删除成功！")
		s.showStudent()
	}
}
