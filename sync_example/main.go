package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var once sync.Once

func f1(job chan<- int64) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		x := rand.Int63n(time.Now().UnixNano())
		fmt.Println("随机数：", x)
		job <- x
	}
	close(job)
}
func f2(job <-chan int64, result chan<- int64) {
	defer wg.Done()
	for i := range job {
		ret := int64(0)
		for i != 0 {
			ret += i % 10
			i /= 10
		}
		result <- ret
	}
}
func main() {
	job := make(chan int64, 100)
	result := make(chan int64, 100)
	wg.Add(1)
	go f1(job)

	// 开启多个 goroutine 执行任务
	for i := 0; i < 24; i++ {
		wg.Add(1)
		go f2(job, result)
	}
	wg.Wait()
	close(result)

	for i := range result {
		fmt.Println(i)
	}
}
