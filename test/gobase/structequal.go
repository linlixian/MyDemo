// 结构体比较
// 两个结构体可以使用 == 或 != 运算符进行比较，但不支持 > 或 <。
// 如果结构体的所有成员变量都是可比较的，那么结构体就可比较
package main

import (
	"fmt"
)

type Student struct {
	id   int
	name string
}

type Student1 struct {
	id   int
	name string
}

type Student2 struct {
	id   int
	name string
	arr  []int
}

func main() {
	// arr := []int{1, 2}
	//比较
	s1 := Student{1, "yy"}
	s2 := Student{2, "yang"}
	s3 := Student{1, "yy"}
	s5 := Student1{1, "yy"}
	var s6 interface{}
	s6 = s1
	s7 := interface{}(s5)
	// s8 := Student2{1, "yy", arr}
	// s9 := Student2{1, "yy", arr}
	fmt.Println(s1 == s2) //false
	fmt.Println(s1 == s3) //true
	// fmt.Println(s1 == s5) //编译不过
	fmt.Println(s1 == Student(s5))  // true 强转后可以比较
	fmt.Println(s1 == s6.(Student)) // true interface类型只能通过.(a)形式强转
	fmt.Println(s7 == s6)           // true interface类型也能比较，比较的是具体实例
	// fmt.Println(s8 == s9)           //  因为有不可比较的元素 slice ，所以此结构体不能比较

	//赋值
	var s4 Student
	s4 = s1
	fmt.Println(s4) //{1 yy}

	// var a int
	// b := interface{}(a) // interface{}是可以指向任意对象的Any类型.包括基本类型
}
