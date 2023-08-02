package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// tcp client
// func main() {
// 	// 一样的，底层应该也就是linux的接口
// 	// socket->connect
// 	conn, err := net.Dial("tcp", "127.0.0.1:8080")
// 	if err != nil {
// 		fmt.Printf("Dial error:%#v\n", err)
// 		return
// 	}
// 	// 别忘了关闭连接
// 	defer conn.Close()
//
// 	reader := bufio.NewReader(os.Stdin)
// 	for {
// 		fmt.Print("please input# ")
// 		msg, err := reader.ReadString('\n')
// 		if err != nil {
// 			fmt.Printf("Read error:%#v\n", err)
// 			return
// 		}
// 		_, err = conn.Write([]byte(msg))
// 		if err != nil {
// 			fmt.Printf("Write error:%#v\n", err)
// 			return
// 		}
// 	}
// }

// UDP client
func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8080,
	})
	if err != nil {
		fmt.Printf("Dial error:%v\n", err)
		return
	}

	defer conn.Close()

	// 发送数据
	reader := bufio.NewReader(os.Stdin)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Read error:%v\n", err)
			return
		}
		_, err = conn.Write([]byte(msg))
		if err != nil {
			fmt.Printf("Write error:%v\n", err)
			return
		}

		// 接受回复的数据
		recv := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(recv)
		if err != nil {
			fmt.Printf("ReadFormUDP error:%v\n", err)
			return
		}
		fmt.Printf("收到回复：%v", string(recv[:n]))
	}
}
