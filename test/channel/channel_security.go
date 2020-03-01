package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var count0 int32 = 0
var count = 0
var count1 = 0
var count2 = 0
var count3 = 0
var ss sync.Mutex

func main() {
	w := &sync.WaitGroup{}
	startTime := time.Now()

	//一百个协程去累加
	for i := 0; i < 2000000; i++ {
		w.Add(1)
		go add1(w)
	}
	w.Wait()
	endTime := time.Now()
	fmt.Println("原子操作耗时：", endTime.Sub(startTime))
	fmt.Println("count0====>", count0)

	startTime = time.Now()
	chs := make(chan int, 2000000)

	//也是用一百个协程累加，并且每个协程都使用channel来通信
	for i := 0; i < 2000000; i++ {
		w.Add(1)
		go add2(chs, w)
	}
	w.Wait()
	close(chs)
	for ch := range chs {
		count += ch
		//fmt.Println(i,"====>",count)
	}
	endTime = time.Now()
	fmt.Println("channel通信操作耗时：", endTime.Sub(startTime))

	fmt.Println("count====>", count)

	startTime = time.Now()
	chsa := make(chan int, 1)
	chsa <- 1
	for i := 0; i < 2000000; i++ {
		w.Add(1)
		go add3(chsa, w)
	}
	w.Wait()
	close(chsa)
	endTime = time.Now()
	fmt.Println("channel阻塞操作（相当于锁但比锁慢很多）耗时：", endTime.Sub(startTime))
	fmt.Println("count====>", count1)

	startTime = time.Now()
	for i := 0; i < 2000000; i++ {
		w.Add(1)
		go add4(w)
	}
	w.Wait()
	endTime = time.Now()
	fmt.Println("互斥锁操作耗时：", endTime.Sub(startTime))

	fmt.Println("count====>", count2)

	startTime = time.Now()
	chsaa := make([]chan int, 2000000)
	for i := 0; i < 2000000; i++ {
		chsaa[i] = make(chan int, 1)
		w.Add(1)

		go add5(w, chsaa[i])
	}
	w.Wait()
	for _, ch := range chsaa {
		close(ch)
		a := <-ch
		count3 += a
		//fmt.Println(i,"====>",count)
	}
	endTime = time.Now()
	fmt.Println("channel数组操作耗时：", endTime.Sub(startTime))

	fmt.Println("count====>", count3)
}

//线程不安全的方式
func add1(w *sync.WaitGroup) {
	// count0++
	atomic.AddInt32(&count0, 1)
	w.Done()
}

//线程安全的方式
func add2(ch chan int, w *sync.WaitGroup) {
	ch <- 1
	w.Done()
}

//线程安全的方式
func add3(ch chan int, w *sync.WaitGroup) {
	<-ch
	count1++
	ch <- 1
	w.Done()
}

//线程安全的方式
func add4(w *sync.WaitGroup) {
	ss.Lock()
	defer ss.Unlock()
	count2++
	w.Done()
}

//线程安全的方式
func add5(w *sync.WaitGroup, ch chan int) {
	ch <- 1
	w.Done()
}
