package main

import (
	"sync"
	"time"
)

var m *sync.RWMutex

// 读的时候别的线程可以读但不能写，写的时候别的线程什么都不能干
// 读写锁主要作用是在写的时候不让别人读，读的时候让别人读 RLock 搭配 Lock 使用
// 如果用普通锁则读的时候别的线程什么都不能干，写的时候别的线程什么都不能干
func main() {
	m = new(sync.RWMutex)

	// 多个同时读
	go read(1)
	go read(2)

	// // 写的时候啥也不能干
	go write(1)
	go read(3)
	go write(3)

	time.Sleep(7 * time.Second)
}

func read(i int) {
	m.RLock()
	println(i, "read start")
	println(i, "reading")
	time.Sleep(1 * time.Second)
	println(i, "read over")
	m.RUnlock()
}

func write(i int) {
	m.Lock()
	println(i, "write start")
	println(i, "writing")
	time.Sleep(1 * time.Second)
	println(i, "write over")
	m.Unlock()
}
