/*
输入描述:
输入一个正浮点数值

输出描述:
输出该数值的近似整数值
*/
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		data, _, _ := reader.ReadLine()
		if len(data) == 0 {
			break
		}
		f, _ := strconv.ParseFloat(string(data), 64)
		f = math.Floor(f + 0.5)
		fmt.Print(f)
	}

}
