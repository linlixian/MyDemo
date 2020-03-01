/*
输入描述:
输入一个十六进制的数值字符串。

输出描述:
输出该数值的十进制字符串。
*/
// str[2:] 因为输入的0xA的0x前缀代表16进制，所以要去掉
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		data, _, _ := reader.ReadLine()
		if len(data) == 0 {
			break
		}
		str := string(data)
		/*
			参数1 数字的字符串形式

			参数2 数字字符串的进制 比如二进制 八进制 十进制 十六进制

			参数3 返回结果的bit大小 也就是int8 int16 int32 int64
		*/
		a, _ := strconv.ParseInt(str[2:], 16, 0)
		fmt.Println(a)
	}

}
