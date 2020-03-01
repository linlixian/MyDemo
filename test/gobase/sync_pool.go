// 对象池(减少gc，提高性能)
// goroutine之间通信：channel，context，全局变量，对象池，存redis
package main

import (
	"fmt"
	"sync"
	"time"
)

// 一个[]byte的对象池，每个对象为一个[]byte
var bytePool = sync.Pool{
	New: func() interface{} {
		b := make([]byte, 1024)
		return &b
	},
}

func main() {
	a := time.Now().Unix()
	// 不使用对象池
	for i := 0; i < 1000; i++ {
		obj := make([]byte, 1024)
		_ = obj
	}
	b := time.Now().Unix()
	// 使用对象池
	for i := 0; i < 1000; i++ {
		obj := bytePool.Get().(*[]byte)
		_ = obj
		bytePool.Put(obj)
	}
	c := time.Now().Unix()
	fmt.Println("without pool ", b-a, "s")
	fmt.Println("with    pool ", c-b, "s")

	obj1 := bytePool.Get().(*[]byte)
	*obj1 = []byte("asdasd")
	bytePool.Put(obj1)

	obj11 := bytePool.Get().(*[]byte)
	fmt.Println("111111111", string(*obj11)) //  asdasd

}

// without pool  34 s
// with    pool  24 s

//上面代码的运行结果显示使用对象池很明显提升了性能
/*
Get方法并不会对获取到的对象值做任何的保证，因为放入本地池中的值有可能会在任何时候被删除，但是不通知调用者。
放入共享池中的值有可能被其他的goroutine偷走。 所以对象池比较适合用来存储一些临时切状态无关的数据，
但是不适合用来存储数据库连接的实例，因为存入对象池重的值有可能会在垃圾回收时被删除掉，这违反了数据库连接池建立的初衷。
根据上面的说法，Golang的对象池严格意义上来说是一个临时的对象池，适用于储存一些会在goroutine间分享的临时对象。
主要作用是减少GC，提高性能。在Golang中最常见的使用场景是fmt包中的输出缓冲区。
*/
