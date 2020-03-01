package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	port := "9090"
	Start(port)
}

// 启动服务器
func Start(port string) {
	host := ":" + port

	// 获取tcp地址
	tcpAddr, err := net.ResolveTCPAddr("tcp4", host)
	if err != nil {
		log.Printf("resolve tcp addr failed: %v\n", err)
		return
	}

	// 监听
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Printf("listen tcp port failed: %v\n", err)
		return
	}

	// 建立连接池，用于广播消息
	conns := make(map[string]net.Conn)

	// 消息通道
	messageChan := make(chan string, 10)

	// 广播消息
	go BroadMessages(&conns, messageChan)

	// 启动
	for {
		fmt.Printf("listening port %s ...\n", port) // 监听阻塞，当收到一个连接请求后再向下执行
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Printf("Accept failed:%v\n", err)
			continue
		}

		// 把每个客户端连接扔进连接池
		conns[conn.RemoteAddr().String()] = conn
		fmt.Println(conns)

		// 处理消息
		go Handler(conn, &conns, messageChan) //每次新增一个客户端连接，这里执行一个协程，messageChan是所有客户端共用的
	}
}

// 向所有连接上的乡亲们发广播
func BroadMessages(conns *map[string]net.Conn, messages chan string) {
	for {

		// 不断从通道里读取消息
		msg := <-messages
		fmt.Println(msg)

		// 向所有的乡亲们发消息
		for key, conn := range *conns {
			fmt.Println("connection is connected from ", key)
			_, err := conn.Write([]byte(msg))
			if err != nil {
				log.Printf("broad message to %s failed: %v\n", key, err)
				delete(*conns, key)
			}
		}
	}
}

// 处理客户端发到服务端的消息，将其扔到通道中
func Handler(conn net.Conn, conns *map[string]net.Conn, messages chan string) {
	fmt.Println("connect from client ", conn.RemoteAddr().String())

	buf := make([]byte, 1024)
	for {
		length, err := conn.Read(buf)
		if err != nil {
			log.Printf("read client message failed:%v\n", err)
			delete(*conns, conn.RemoteAddr().String()) // 客户端断开后(四次挥手)，删除连接池的连接（严格来说还需要心跳机制来确定连接是否异常断开）
			conn.Close()                               // 如果不关闭连接会导致close_wait过多
			break
		}

		// 1.把收到的消息写到通道中，这里使用通道可以提高效率，因为这样可以使用goroutine，实现同时读取数据并广播数据
		//   如果不用通道，则只能接受一个数据并广播，然后才能再接收下一个数据并广播
		// 2.为了同步处理从各个客户端传来的数据，防止并发错误（所有客户端发送的信息都会到这个通道中，然后可以一一广播）
		recvStr := string(buf[0:length])
		messages <- recvStr
	}
}
