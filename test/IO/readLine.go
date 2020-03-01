package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine() // 没有读取到则阻塞
	tmp, _, _ := reader.ReadLine()  // 没有读取到则阻塞

	count := 0
	for _, v := range data {
		if v == tmp[0] || v == tmp[0]-32 || v-32 == tmp[0] { // 32 是ascll码大小写相差数，这里不区分大小写
			count++
		}
	}
	fmt.Println(count)
}
