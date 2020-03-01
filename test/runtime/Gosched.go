package main

import (
	"fmt"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(13) //使用多核
}

func main() {
	exit := make(chan int)
	go func() {
		defer close(exit)
		go func() {
			fmt.Println("b")
		}()
	}()

	for i := 0; i < 100; i++ {
		fmt.Println("a:", i)

		if i == 50 {
			runtime.Gosched() //切换任务
		}
	}
	<-exit
	num := runtime.NumCPU()
	fmt.Println(num)
}
