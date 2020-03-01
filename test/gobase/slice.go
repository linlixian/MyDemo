// slice的扩容和共享
package main

import (
	"fmt"
)

type s struct {
	a int
}

// // 使用指针进行传递
// func main() {
// 	var arr = [5]int{1, 2, 3, 4, 5}
// 	fmt.Println(sum(&arr))
// }
// func sum(arr *[5]int) int {
// 	s := 0
// 	for i := 0; i < len(arr); i++ {
// 		s += arr[i]
// 	}
// 	return s
// }

// // 使用切片进行传递
// func main() {
// 	var arr = [5]int{1, 2, 3, 4, 5}
// 	fmt.Println(sum(arr[:]))
// }
// func sum(arr []int) int {
// 	s := 0
// 	for i := 0; i < len(arr); i++ {
// 		s += arr[i]
// 	}
// 	return s
// }

func main() {
	a1 := s{1}
	a2 := s{2}
	var list []s
	list = append(list, a1, a2)
	fmt.Println(list)

	for _, v := range list { // range是拷贝
		v.a += 1
	}
	fmt.Println(list)

	for k, _ := range list {
		list[k].a += 1
	}
	fmt.Println(list)
	// s1 := []int{0, 1, 2, 3, 8: 123} // 通过初始化表达式构造，可使用索引号。即下标最大是8，存的是123
	// fmt.Println(s1, len(s1), cap(s1))

	// data := [...]int{0, 1, 2, 3, 4, 10: 0}
	// s := data[:2:3]         // 3代表cap
	// s = append(s, 100, 200) // 一次 append 两个值，超出 s.cap 限制，如果cap改成4，不会扩容，data也会改变
	// fmt.Println(s, len(s), cap(s), data)
	// fmt.Println(&s[0], &data[0]) // 比对底层数组起始指针

	s := make([]int, 0, 1)
	c := cap(s)

	for i := 0; i < 50; i++ {
		s = append(s, i)
		if n := cap(s); n > c {
			fmt.Printf("cap: %d -> %d\n", c, n)
			c = n
		}
	}
	/*通常以 2 倍容量重新分配底层数组
	  cap: 1 -> 2
	  cap: 2 -> 4
	  cap: 4 -> 8
	  cap: 8 -> 16
	  cap: 16 -> 32
	  cap: 32 -> 64
	*/

	// slice的共享和扩容
	/*
			我们知道数据结构中切片是非常高效的，可以直接寻址，但是有个缺陷，难以扩容。
			所以slice被设计为指向切片的指针，在需要扩容时，会将底层切片上的值复制到一个更大的切片上然后指向这个新切片。
		slice有个特性是允许多个slice指向同一个底层切片，这是一个有用的特性，在很多场景下都能通过这个特性实现 no copy 而提高效率。
		但共享同时意味着不安全。b在追加3时实际上覆盖了a[1]，导致c变成了[3 4]。
	*/
	// a := make([]int, 2, 2)
	// a[0], a[1] = 1, 2
	// fmt.Println(a)

	// b := append(a[0:1], 3) // a[0:1]即{1,待定}，append 3 后把原来的2覆盖
	// fmt.Println("a=", a)   // [1 3]
	// fmt.Println("b=", b)   // [1 3]

	// fmt.Println(a[1:2]) // [3]

	// c := append(a[1:2], 4) // a[1:2]即{待定,3}，因为3后面没空间，所以append 4 后新建了一个切片即为c={3，4}
	// fmt.Println("a=", a)   // [1 3]

	// fmt.Println(b, c)
	// 解决方法,使用copy
	// a := make([]int, 2, 2)
	// a[0], a[1] = 1, 2

	// b := make([]int, 1)
	// copy(b, a[0:1])
	// b = append(b, 3)
	// fmt.Println("a=", a) // [1 2]
	// fmt.Println("b=", b) // [1 3]

	// c := make([]int, 1)
	// copy(c, a[1:2])
	// c = append(c, 4)

	// fmt.Println(c) // [2,4]

	// copy注意点
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice s1 : %v\n", s1)
	s2 := make([]int, 10)
	fmt.Printf("slice s2 : %v\n", s2)
	copy(s2, s1)
	fmt.Printf("copied slice s1 : %v\n", s1)
	fmt.Printf("copied slice s2 : %v\n", s2)
	s3 := []int{1, 2, 3}
	fmt.Printf("slice s3 : %v\n", s3)
	s3 = append(s3, s2...)
	fmt.Printf("appended slice s3 : %v\n", s3)
	s3 = append(s3, 4, 5, 6)
	fmt.Printf("last slice s3 : %v\n", s3)
	/*
			slice s1 : [1 2 3 4 5]
		slice s2 : [0 0 0 0 0 0 0 0 0 0]
		copied slice s1 : [1 2 3 4 5]
		copied slice s2 : [1 2 3 4 5 0 0 0 0 0]
		slice s3 : [1 2 3]
		appended slice s3 : [1 2 3 1 2 3 4 5 0 0 0 0 0]
		last slice s3 : [1 2 3 1 2 3 4 5 0 0 0 0 0 4 5 6] // 注意0也是拷贝的数据，不是cap
	*/

	var a [10]int
	fmt.Println(a) // [0 0 0 0 0 0 0 0 0 0]  数组不用通过make初始化内存

	str := "hello"
	// str[0] = "x" // 编译错误
	fmt.Println(str)

	// s := make([]int) // 编译错误

	saa := new([]int)                 // new可以初始化任何类型,初始化大小为0
	fmt.Println(len(*saa), cap(*saa)) // 0,0
	fmt.Println(saa)                  // &[]

}
