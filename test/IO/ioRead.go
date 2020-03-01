// bufio.NewReader是固定内存空间读取,类似 net.conn的Read()
//  ioutil.ReadAll是直接读取所有
package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
	file, _, err := r.FormFile("file") // multipart.file 类型

	// if err == nil {
	// 	data, err := ioutil.ReadAll(file) // 与bufio.NewReader一样可以读取文件内容
	// 	if err == nil {
	// 		fmt.Println(len(data))
	// 		fmt.Fprintln(w, string(data))
	// 	}
	// }
	// fmt.Println(err)

	// file, err := os.Open("./file.txt")  // multipart.file 类型
	// if err != nil {
	// 	fmt.Printf("os open ./file.txt err : %v\n", err)
	// }
	// if file != nil {
	// 	defer func(file *os.File) { file.Close() }(file)
	// }

	read1 := bufio.NewReader(file)

	var b1 = make([]byte, 102)
	readByte1, err := read1.Read(b1)
	if err != nil {
		fmt.Printf("read err : %v\n", err)
	}
	fmt.Printf("read success , 读取 %d 字节\n读取的内容：\n%s\n", readByte1, string(b1))

	var line []byte
	for {
		data, prefix, err := read1.ReadLine()
		if err == io.EOF {
			// fmt.Println(err)
			break
		}

		line = append(line, data...)
		if !prefix {
			// fmt.Printf("data:%s\n", string(line))
		}
	}
	fmt.Println(string(line))
}

// func main() {
// 	http.HandleFunc("/", sayHello) //注册URI路径与相应的处理函数
// 	log.Println("【默认项目】服务启动成功 监听端口 80")
// 	er := http.ListenAndServe("0.0.0.0:80", nil)
// 	if er != nil {
// 		log.Fatal("ListenAndServe: ", er)
// 	}

// }

func main() {
	file, err := os.Open("./test/IO/scan.go") // 等同于 test/IO/scan.go
	defer file.Close()
	if err != nil {
		fmt.Println("open file failed:", err)
		return
	}
	read1 := bufio.NewReader(file)

	var b1 = make([]byte, 102)
	for {
		readByte1, err := read1.Read(b1)
		if err != nil {
			fmt.Printf("read err : %v\n", err.Error())
			break
		}
		fmt.Printf("read success , 读取 %d 字节\n读取的内容：\n%s\n", readByte1, string(b1))
	}
}
