package main

import "fmt"

func main() {
	fmt.Println("sadfasdfhasl")
	// && || !
	// & | ^
	// >> <<
	// 运算符和C一模一样

	// arr
	// var arr [10]int32
	// var a1 [5]bool
	// for i,v := range arr{
	// 	fmt.Printf("%d:%d\n", i, v)
	// }
	// fmt.Printf("%T\n", arr)
	// fmt.Printf("%T\n", a1)
	// fmt.Println(arr, a1)

	// a := [...]int32{1,2,3,4,5}
	// for i := 0; i < len(a); i++{
	// 	fmt.Println(a[i])
	// }
	// fmt.Println(a)
	// citys := [...]string{"safsad", "上海", "NewYork"}
	// for i, v := range citys{
	// 	fmt.Println(i, v)
	// }

	// arr := [...]int{3245,546,4575,567}
	// sum := 0
	// fmt.Printf("%T\n", sum)
	// for i := 0; i < len(arr); i++{
	// 	sum += arr[i]
	// }
	// fmt.Println(sum)

	// arr := [...]int{1,3,5,7,8}
	// for i := 0; i < len(arr); i++{
	// 	for j := i + 1; j < len(arr); j++{
	// 		if arr[i] + arr[j] == 8{
	// 			fmt.Println(i,j)
	//     }
	// 	}
	// }

	// 由数组得到切片
	// a := [...]int{1,3,45,56}
	// // 用a的[0, 3)下标的值初始化a1
	// a1 := a[0:3]
	// // fmt.Println(a)
	// // fmt.Println(a1)
	// a2 := a[1:2]
	// fmt.Println(len(a), cap(a))   // 4 4
	// fmt.Println(len(a1), cap(a1)) // 3 4
	// fmt.Println(len(a2), cap(a2)) // 1 3

	// a := make([]int, 5, 10)
	// fmt.Printf("%T, %v, %d, %d\n", a, a, len(a), cap(a))

	// 切片的赋值
	// a := []int{1, 3, 4, 5, 65}
	// b := a
	// b[1] = 3000
	// // fmt.Println(a)
	// // fmt.Println(b)
	// fmt.Printf("%T, %v, %d, %d", a, a, len(a), cap(a))
	// fmt.Println()
	// // 2倍扩容
	// a = append(a, 1111)
	// fmt.Printf("%T, %v, %d, %d", a, a, len(a), cap(a))
	// c := []int{5, 6, 7}
	// // ... 表示把 c 拆开
	// a = append(a, c...)
	// fmt.Println()
	// fmt.Printf("%T, %v, %d, %d", a, a, len(a), cap(a))
	// fmt.Println()

	// // copy
	// a1 := make([]int, 10)
	// copy(a1, a)
	// a2 := a
	// fmt.Println(a)
	// fmt.Println(a1)
	// fmt.Println(a2)
	// a[0] = 10000000
	// fmt.Println()
	// fmt.Println(a)
	// fmt.Println(a1)
	// fmt.Println(a2)
	a := []int{1,2,3,4,5}
	fmt.Println(a)
	a = append(a[0:2], a[3:]...)  // 删除下标为2的元素
	fmt.Println(a)


}
