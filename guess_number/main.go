package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 生成[0,100)的随机数
	// go1.20已经设置了
	rand.Seed(time.Now().UnixNano()) // 设置随机数种子
	reallyNum := rand.Intn(101)
	fmt.Println(reallyNum)
	var input int
	fmt.Print("Please input a number between 0 and 100:")
	fmt.Scan(&input)
	for input != reallyNum {
		if input > reallyNum {
			fmt.Println("Your input is bigger than answer!Please input again!")
			fmt.Scan(&input)
		} else if input < reallyNum {
			fmt.Println("Your input is smaller than answer!Please input again!")
			fmt.Scan(&input)
		}
	}
	fmt.Println("You are right!")
}
