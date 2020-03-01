// channel感知到已关闭然后安全退出协程
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 1)

	go func() {

	loop:
		for {
			select {
			case n, ok := <-c:
				fmt.Println("读取到数据 ", n)
				if !ok {
					fmt.Println("channel关闭了")
					break loop
				}
			case <-time.After(time.Second * 2):
				fmt.Println("超时")
				break loop
			}
		}
		fmt.Println("安全退出goroutine")
	}()
	c <- 1

	// close(c)
	time.Sleep(time.Second * 4)

}
