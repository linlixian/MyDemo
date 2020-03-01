// 3、下面的代码会输出什么，并说明原因

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// 运行结果：

// i:  9
// i:  10
// i:  10
// i:  10
// i:  10
// i:  10
// i:  10
// i:  10
// i:  10
// i:  10
// i:  10
// i:  0
// i:  1
// i:  2
// i:  3
// i:  4
// i:  5
// i:  6
// i:  7
// i:  8
// 答：
// 将随机输出数字，但前面一个循环中并不会输出所有值。

// 解析：

// 实际上第一行是否设置CPU为1都不会影响后续代码。

// 2017年7月25日：将GOMAXPROCS设置为1，将影响goroutine的并发，后续代码中的go func()相当于串行执行。

// 两个for循环内部go func 调用参数i的方式是不同的，导致结果完全不同。这也是新手容易遇到的坑。

// 第一个go func中i是外部for的一个变量，地址不变化。遍历完成后，最终i=10。故go func执行时，i的值始终是10（10次遍历很快完成）。

// 第二个go func中i是函数参数，与外部for中的i完全是两个变量。尾部(i)将发生值拷贝，go func内部指向值拷贝地址。
