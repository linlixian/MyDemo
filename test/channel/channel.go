package main

import (
	"fmt"
)

type Stu struct {
	name string
}

func main() {
	// 接口存取
	var StuInterChain chan interface{}
	StuInterChain = make(chan interface{}, 10)
	stuInit := Stu{
		name: "Murphy",
	}
	//存
	StuInterChain <- &stuInit
	//取
	mFetchStu := <-StuInterChain
	fmt.Printf("interface 类型 chan : %v\n", mFetchStu)

	//转
	var mStuConvert *Stu
	mStuConvert, ok := mFetchStu.(*Stu)
	if !ok {
		fmt.Println("cannot convert")
		return
	}
	fmt.Printf("interface chan转 *struct chan : %v\n", mStuConvert)
	fmt.Printf("interface chan转 *struct chan 取值 : %v\n", *(mStuConvert))

	// 内置函数 len 返回未被读取的缓冲元素数量，cap 返回缓冲区大小。
	ch1 := make(chan int)
	ch2 := make(chan int, 3)

	// ch1 <- 1 // 如果协程有cs <- cakeName，却没有检测到<-cs，则会死锁

	fmt.Println(len(ch1), cap(ch1)) // 0 0
	fmt.Println(len(ch2), cap(ch2)) // 0 3

	// channel对nil的chan操作都会造成死锁
	// select case 会跳过阻塞或者报错的通道case
	ch2 = nil
	ch1 = nil
	// for {
	select {
	case k1 := <-ch2:
		fmt.Println(k1)
	case k2 := <-ch1:
		fmt.Println(k2)
	default:
		fmt.Println("chan")
	}
	// }

	// channel单向 ：可以将 channel 隐式转换为单向队列，只收或只发。
	// 不能将单向 channel 转换为普通 channel。
	c := make(chan interface{}, 3)

	var send chan<- interface{} = c // send-only
	var recv <-chan interface{} = c // receive-only

	send <- stuInit
	// <-send               // Error: receive from send-only type chan<- int
	// recv <- 2           // Error: send to receive-only type <-chan int
	send <- stuInit
	send <- stuInit
	close(send) // close不能关闭只收通道

	for {
		val, ok := <-recv
		fmt.Println(ok)
		fmt.Println(val)
		if ok { // 如果通道关闭了，只有当chan中数据都取完后，再取时ok才为false
			fmt.Println(val)
		} else {
			break
		}
	}

}
