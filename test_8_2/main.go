package main

import (
	"fmt"
	"net"
)

// var lock sync.Mutex
// var rwLock sync.RWMutex
// var mp sync.Map
// func main() {
// 	// lock.Lock()
// 	// rwLock.
// 	mp.Store(1, 1)
// 	atomic.
// }

func main() {
	// 底层应该就是 linux 的网络接口：socket->bind->listen->accept
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Printf("Listen error:%#v\n", err)
		return
	}
	conn, err := listener.Accept()
	if err != nil {
		fmt.Printf("Accept error:%#v\n", err)
		return
	}
	b := make([]byte, 128)
	// 实际读到n个字节
	n, err := conn.Read(b)
	if err != nil {
		fmt.Printf("Read error:%#v\n", err)
		return
	}
	fmt.Println(string(b[:n]))
}
