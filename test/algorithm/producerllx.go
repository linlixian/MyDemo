package main

import (
	"fmt"
	"sync"
)

func productor(chanData chan int, i int, w *sync.WaitGroup) {
	chanData <- i
	fmt.Println("生产：", i)
	w.Done()
}

func consumer(chanData chan int, i int, w *sync.WaitGroup) {
	fmt.Println("消费：", <-chanData)
	w.Done()
}

func main() {
	chanData := make(chan int, 10)
	w := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		w.Add(1)
		go productor(chanData, i, w)
	}
	for i := 0; i < 10; i++ {
		w.Add(1)
		go consumer(chanData, i, w)
	}
	w.Wait()
}
