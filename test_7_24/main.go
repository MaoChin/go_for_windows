package main

import (
	"fmt"
)

func main() {
	// // 数组
	// a := [...]int{1, 2, 3, 4, 5}
	// // 由数组得到切片
	// s := a[:]
	// // 利用append 删除s切片下标为2的值
	// s = append(s[:2], s[3:]...)
	// fmt.Println(s, len(s), cap(s))
	// // a的值也会改！！！s是a的引用，对s的改变会改变a
	// fmt.Println(a, len(a), cap(a))
	// s := make([]int, 5, 10)
	// fmt.Println(s)

	// m1 := make(map[string]int, 10)
	// m1["sdfsa"] = 12
	// m1["kkk"] = 1
	// fmt.Println(m1)
	// fmt.Println(m1["kkk"])
	// fmt.Println(m1["as"])

	// m := make(map[int]string, 10)
	// m[34] = "safsa"
	// m[1] = "safdsfasdfsa"
	// m[222] = "kkk"
	// // 遍历
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }
	// // 以key为顺序遍历
	// keys := make([]int, 0, 10)
	// fmt.Println(keys)
	// for key := range m {
	// 	keys = append(keys, key)
	// }
	// fmt.Println(keys)
	// sort.Ints(keys)
	// fmt.Println(keys)
	// for _, k := range keys {
	// 	fmt.Printf("%d:%s\n", k, m[k])
	// }

	// 值是map类型的切片
	// var s = make([]map[int]string, 9, 10)
	// s[0] = make(map[int]string, 2)
	// s[0][10] = "sadfsad"
	// s[0][3434] = "nk"
	// fmt.Println(s)

	// // value 是切片类型的map
	// var mp = make(map[string][]int, 10)
	// mp["asfsadf"] = []int{1, 2, 3, 4}
	// fmt.Println(mp)

	// 统计一个句子中字符出现的次数
	// s := "sdahhfsadfasdf jsdklf,sadfsa.,.asdf.as /sdafsad'sadfsadf'"

	// mp := make(map[rune]int, 256)
	// for _, v := range s {
	// 	mp[v]++
	// }
	// for k, v := range mp {
	// 	fmt.Printf("%c:%d\n", k, v)
	// }

	// 函数
	// s := sum(1, 3)
	// a := 1
	// a++

	// fmt.Println(s)

	// s := "how do you do"
	// s1 := strings.Split(s, " ")
	// mp := make(map[string]int, 100)
	// for _, v := range s1 {
	// 	mp[v]++
	// }
	// fmt.Println(mp)

	// 判断回文
	s := "asdsidsa"
	l := len(s)
	flag := true
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			flag = false
			break
		}
	}
	fmt.Println(flag)
}
func sum(x int, y int) (ret int) {
	return x + y
}
