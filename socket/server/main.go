package main

import (
	"fmt"
	"net"
)

func handlerTask(conn net.Conn) {
	b := make([]byte, 128)
	// 实际读到n个字节
	for {
		n, err := conn.Read(b)
		if err != nil {
			fmt.Printf("Read error:%#v\n", err)
			return
		}
		fmt.Print(string(b[:n]))
	}
}

// tcp server
// func main() {
// 	// 底层应该就是 linux 的网络接口：socket->bind->listen->accept
// 	listener, err := net.Listen("tcp", "127.0.0.1:8080")
// 	if err != nil {
// 		fmt.Printf("Listen error:%#v\n", err)
// 		return
// 	}
// 	defer listener.Close()
//
// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			fmt.Printf("Accept error:%#v\n", err)
// 			return
// 		}
// 		go handlerTask(conn)
// 	}
// }

// UDP server
func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8080,
	})
	if err != nil {
		fmt.Printf("ListenUdp error:%v\n", err)
		return
	}
	defer conn.Close()

	// 读数据
	msg := make([]byte, 1024)
	for {
		// 对端addr
		n, clientAddr, err := conn.ReadFromUDP(msg)
		if err != nil {
			fmt.Printf("Read error: %v\n", err)
			return
		}
		fmt.Print(string(msg[:n]))

		// 向对端发送数据
		_, err = conn.WriteToUDP((msg[:n]), clientAddr)
		if err != nil {
			fmt.Printf("Write error: %v\n", err)
			return
		}
	}
}
