package main

import (
	"MyProject/test/socket/tcp-pack/proto"
	"bufio"
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn) // 流水一样不断流入数据,所以可以按顺序获取包,此方法适用于粘包处理
	// buf := make([]byte, 1024)

	for {

		//用一个固定大小的缓存循环接收会出错，因为接受一次数据后，无法填满buf
		// _, err := conn.Read(buf)
		// log.Println("11111", string(buf))
		// if err != nil {
		// 	log.Printf("recv server msg failed: %v\n", err)
		// 	break
		// }
		// reader := bufio.NewReader(bytes.NewReader(buf))

		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		fmt.Println("收到client发来的数据：", msg)
	}
}

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		fmt.Printf("listening port %s ...\n", 30000)
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}
