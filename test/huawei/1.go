/*
输入描述:
输入多行，先输入随机整数的个数，再输入相应个数的整数

输出描述:
返回多行，处理后的结果
*/
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
		tmp, _, _ := reader.ReadLine()
		if len(tmp) == 0 {
			break
		}
		count, _ := strconv.Atoi(string(tmp))
		var arry [1001]bool
		for i := 0; i < count; i++ {
			tmp, _, _ := reader.ReadLine()
			num, _ := strconv.Atoi(string(tmp))
			arry[num] = true
		}
		for k, v := range arry {
			if v == true {
				fmt.Println(k)
			}
		}
	}
}
