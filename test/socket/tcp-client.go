package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	Start(":9090") // 客户端随机生成一个端口去和9090连接
}

func Start(tcpAddrStr string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", tcpAddrStr)
	if err != nil {
		log.Printf("Resolve tcp addr failed: %v\n", err)
		return
	}

	// 向服务器拨号
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Printf("Dial to server failed: %v\n", err)
		return
	}

	// 使用goroutine收发数据可以提高效率，即边发边收
	// 向服务器发消息
	go SendMsg(conn)

	// 接收来自服务器端的广播消息
	buf := make([]byte, 1024)
	for {
		//阻塞等待用户发送的数据,因为不知道传过来的数据多大，所以用一个固定大小的缓存循环接收
		length, err := conn.Read(buf)
		if err != nil {
			log.Printf("recv server msg failed: %v\n", err)
			conn.Close()
			os.Exit(0)
			break
		}

		fmt.Println(string(buf[0:length]))
	}
}

// 向服务器端发消息
func SendMsg(conn net.Conn) {
	username := conn.LocalAddr().String()
	for {
		var input string

		// 接收输入消息，放到input变量中
		fmt.Scanln(&input)

		if input == "/q" || input == "/quit" {
			fmt.Println("Byebye ...")
			conn.Close()
			os.Exit(0)
		}

		// 只处理有内容的消息
		if len(input) > 0 {
			msg := username + " say:" + input
			_, err := conn.Write([]byte(msg)) //tcp传输字节流
			if err != nil {
				conn.Close()
				break
			}
		}
	}
}
