package main

import (
	"errors"
	"unsafe"
)

type MyInt int

var i int = 1

// var j MyInt = i // 类型不一致，不能自动转型

func asdasd() {
	// var a int = 10
	// var a = 10 // 基本类型不写声明也可以初始化
	// a := 10

	// 以下切片声明方式都正确
	// x := []int{1, 2, 3, 4, 5, 6}
	// x := []int{
	// 	1, 2, 3,
	// 	4, 5, 6}
	// x := []int{
	// 	1, 2, 3,
	// 	4, 5, 6,
	// }

	// b := 1
	// ++b  go没有++b这种，和java不同
	// c := b++  这种写法也是错的

	const zero = 0.0

}

type stu struct {
}

// 常量值必须是编译期可确定的字符、字符串、布尔或数字类型的值。
const (
	ERR_ELEM_EXIST    stu   = stu{}                            // 编译错误
	ERR_ELEM_NT_EXIST error = errors.New("element not exists") // 编译错误
	a                       = "hello world"
	b                       = len(a)
	c                       = unsafe.Sizeof(b)
)
