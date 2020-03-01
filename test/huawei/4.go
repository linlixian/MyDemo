/*
输入描述:
输入一个long型整数

输出描述:
按照从小到大的顺序输出它的所有质数的因子，以空格隔开。最后一个数后面也要有空格。
*/
// uint 没有符号
// 质数：只能被1和它本身整除，1不是质数
package main

import (
	"fmt"
)

func main() {
	for {
		var num, tmp, i uint
		_, err := fmt.Scan(&num)
		if err != nil {
			break
		}
		tmp = num
		for i = 2; i <= tmp; i++ {
			if num == 1 {
				break
			}
			for num%i == 0 {
				num /= i
				fmt.Print(i, " ")
			}
		}
		fmt.Println()
	}
}
