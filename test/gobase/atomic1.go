package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {

	var t1 uint64 = 0
	var t2 uint64 = 0

	endChan := make(chan int)
	for i := 0; i < 1000; i++ {
		go func() {
			for i := 0; i < 10000; i++ {
				atomic.AddUint64(&t1, 1)
				t2 += 1
			}
			endChan <- 1
		}()
	}

	for i := 0; i < 1000; i++ {
		<-endChan
	}

	// 测试非原子操作造成的值不正确
	// t1= 10000000
	// t2= 8513393
	fmt.Println("t1=", t1)
	fmt.Println("t2=", t2)

	// 性能测试
	func() {
		var t1 uint64 = 0

		startTime := time.Now()
		for i := 0; i < 1000000000; i++ {
			t1 += 1
		}
		endTime := time.Now()
		fmt.Println("非原子操作耗时：", endTime.Sub(startTime))
		// 非原子操作耗时： 535.0303ms

	}()

	func() {
		var t1 uint64 = 0

		startTime := time.Now()
		for i := 0; i < 1000000000; i++ {
			atomic.AddUint64(&t1, 1)
		}
		endTime := time.Now()
		fmt.Println("原子操作耗时：", endTime.Sub(startTime))
		//原子操作耗时： 14.7758413s
	}()
}
