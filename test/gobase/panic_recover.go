/*
当启动多个goroutine时，如果其中一个goroutine异常了，并且我们并没有对进行异常处理，那么整个程序都会终止，
所以我们在编写程序时候最好每个goroutine所运行的函数都做异常处理，异常处理采用recover
先处理defer，再处理panic
*/
package main

import (
	"fmt"
	"time"
)

func A() {
	defer func() {
		if r := recover(); r != nil {

			fmt.Println("Recovered in A", r)
		}
	}()

	fmt.Println("Calling A.")
	B(-1)
	fmt.Println("Returned normally from g.")
}

func B(i int) {
	if i > -1 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
		panic(fmt.Sprintf("%v", 123))
	}
	defer fmt.Println("Defer in B", i)
	fmt.Println("Printing in B", i)
	B(i + 1) //这里是一个递归函数函数。
}

func C() {
	for {
		time.Sleep(1 * time.Second)

		fmt.Println("c is working")
	}
}

func main() {
	go A()
	go C()
	time.Sleep(5 * time.Second)
	fmt.Println("程序结束！")
}
