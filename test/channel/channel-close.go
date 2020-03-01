// chan设nil可以执行select的其他case，因为nil会堵塞
// 引用类型传到channel和传到方法一样，也是值拷贝。只不过该值是地址，所以看起来好像是引用传递(range也是同理)
// 不管是指针、引用类型还是其他类型参数，都是值拷贝传递，区别在于拷贝的目标对象
package main

import "fmt"

func main() {
	c := make(chan []int, 10)
	// c := make(chan int, 10)
	// c = nil
	var a = []int{1, 2, 3}
	// a := 1
	c <- a
	// close(c)

loop:
	for {
		select {
		case i, isClose := <-c:
			if !isClose {
				fmt.Println("channel closed!")
				c = nil
				continue
			}
			fmt.Println("1111111", i)
			i[0] = 5
			// i = 2
			fmt.Println("1111111", i)

		default:
			fmt.Println("22222222")
			break loop // 如果不加loop，则只是跳出select
			// or goto loop1

		}

		// i, isClose := <-c
		// if !isClose {
		// 	fmt.Println("channel closed!")
		// 	break
		// }
		// fmt.Println(i)
	}
	fmt.Println("3333333", a)

}
