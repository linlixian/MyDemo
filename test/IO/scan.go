// scan  以空格为分割，不断循环读取字符串
package main

import (
	"fmt"
)

func main() {
	var str, tmp string
	for {
		n, _ := fmt.Scan(&str)
		fmt.Println(str)
		fmt.Println(n)

		if n != 0 {
			tmp = str
		}
		if n == 0 {
			fmt.Println(len(tmp))
			break
		}
	}
}
