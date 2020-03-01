// 主程序结束goroutine也结束
package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			fmt.Println("waork..")
		}
	}()

	time.Sleep(3 * time.Second)

}
