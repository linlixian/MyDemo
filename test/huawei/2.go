/*
输入描述:
连续输入字符串(输入2次,每个字符串长度小于100)

输出描述:
输出到长度为8的新字符串数组
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := 0
	for n < 2 {
		tmp, _, _ := reader.ReadLine()
		flag := 0
		length := len(tmp)
		for length > 8 {
			fmt.Println(string(tmp[flag : flag+8]))
			flag += 8
			length -= 8
		}
		tmp1 := string(tmp[flag : flag+length])
		for i := 0; i < 8-length; i++ {
			tmp1 += "0"
		}
		fmt.Println(string(tmp1))
		n++
	}
}
