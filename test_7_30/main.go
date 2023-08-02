package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var once sync.Once

func f1(a chan int) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		a <- i
	}
	// 别忘了关闭通道
	close(a)
}
func f2(a, b chan int) {
	defer wg.Done()
	// for {
	// 	x, ok := <-a
	// 	if !ok {
	// 		break
	// 	}
	// 	b <- (x * x)
	// }
	for i := range a {
		b <- (i * i)
	}
	// 别忘了关闭通道, 这样确保该操作只执行一次！
	once.Do(func() { close(b) })
}
func main() {
	var a, b chan int
	a = make(chan int, 100)
	b = make(chan int, 100)
	wg.Add(3)

	// 本身就是引用类型，不用传地址
	// go f1(&a)
	// go f2(&a, &b)
	go f1(a)
	go f2(a, b)
	go f2(a, b)

	wg.Wait()
	for i := range b {
		fmt.Println(i)
	}
}
