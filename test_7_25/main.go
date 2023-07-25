package main

import "fmt"

// func f1() int {
// 	x := 5
// 	defer func() {
// 		x++
// 	}() // 直接调用该函数
// 	return x
// }
// func f2() (ret int) {
// 	defer func() {
// 		ret++
// 	}()
// 	return 2
// }
// func f3() (ret int) {
// 	x := 5
// 	defer func() {
// 		x++ // 修改的是x
// 	}()
// 	return x
// }
// func f4() (ret int) {
// 	// 传参是值拷贝，对函数里变量的改变不改变外面
// 	defer func(ret int) {
// 		ret++
// 	}(ret) // 把ret当参数,里面改的是一份拷贝
// 	return 5
// }
// func f5(x int) func(int) int{
// 	return func(y int) int{
// 		return x + y
// 	}
// }
//
//
// func main() {
// 	// fmt.Println("sajfksld")
// 	// defer fmt.Println("最后执行")
// 	// fmt.Println("kjfhjsakdhf")
//
// 	// fmt.Println(f1())
// 	// fmt.Println(f2())
// 	// fmt.Println(f3())
// 	// fmt.Println(f4())
//
// 	ret := f5(100)
// 	fmt.Printf("%T\n", ret)
// 	ret2 := ret(200)
// 	fmt.Printf("%T\n", ret2)
// 	fmt.Println(ret2)
//
// }

// func f1(f func()) {
// 	// ...
// 	fmt.Println("I am the f1")
// 	f() // 闭包后f1里调用的这个f就是f2
// 	// ...
// }
// func f2(x, y int) {
// 	// ...
// 	fmt.Println("I am the f2")
// 	fmt.Println(x + y)
// }
//
// // f3参数是f2的那一套，返回值是f1的参数
// func f3(f func(int, int), x, y int) func() {
// 	return func() {
// 		// 这里就相当于调用了传入的函数
// 		f(x, y)
// 	}
// }
//
// func main() {
// 	// 想要 f1(f2), 借助f3  闭包
// 	// f3的返回值(匿名函数)刚好是f1的参数，这样把原来需要传递两个int类型的参数包装成了一个匿名函数传给f1
// 	f1(f3(f2, 100, 200))
// 	// f3(f2, 100, 200)() // 这样就相当于调用f2
// }

// func main() {
// panic && recover
// if(...){
// 	// 出现错误
// 	panic("错误信息")  // 直接报错退出程序
// }
// fmt.Printf("%b\n", 10)
// fmt.Printf("%q\n", 65)
// var s string
// var a int
// fmt.Scan(&s, &a)
// // fmt.Scanln(&s, &a)

// fmt.Println(s, a)

// s2 := fmt.Sprintf("name:%s, age:%d", "asdfas", 12)
// fmt.Println(s2)
// }

var (
	coins = 50
	users = []string{"Mattew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter",
		"Giana", "Adriano", "Aaron", "Elizabeth"}
	distribution = make(map[string]int, len(users))
)

func destributionCoin() int {
	for _, str := range users {
		for _, s := range str {
			switch {
			case s == 'e' || s == 'E':
				if coins > 0 {
					distribution[str] += 1
					coins -= 1
				}
			case s == 'i' || s == 'I':
				if coins > 0 {
					distribution[str] += 2
					coins -= 2
				}
			case s == 'o' || s == 'O':
				if coins > 0 {
					distribution[str] += 3
					coins -= 3
				}
			case s == 'u' || s == 'U':
				if coins > 0 {
					distribution[str] += 4
					coins -= 4
				}
			}
		}
	}
	return coins
}
func main() {
	left := destributionCoin()
	fmt.Printf("剩余：%d\n", left)
	fmt.Println(distribution)
}
