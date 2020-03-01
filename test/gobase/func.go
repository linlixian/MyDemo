// 只有 func (u *user) changeEmail0  才能改变原本的属性值 func (u user)只是操作的副本
package main

import (
	"fmt"
)

type user struct {
	name  string
	email string
}

func (u user) changeEmail0(email string) {
	u.email = email
	fmt.Println("in func", u.email)
}
func (u user) notify() { fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email) }

func main() {

	bill := user{"Bill", "bill@email.com"}
	bill.notify()
	bill.changeEmail0("bill@newdomain0.com") //只对bill的一个副本进行修改，不会改的bill本身
	bill.notify()
	//运行结果
	// Sending User Email To Bill<bill@email.com>
	// in func bill@newdomain0.com
	// Sending User Email To Bill<bill@email.com>

	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()
	//lisa为类型值的指针，而notify方法的接受者为类型值,编译器会做如下操作：(*lisa).notify()
	lisa.changeEmail0("lisa@comcast0.com")
	//lisa为类型值的指针，而changeEmail0方法的接受者为类型值，编译器实际会做如下操作：(*lisa).changeEmail("lisa@comcast0.com"), 这次操作的是lisa指向的值的“副本”，不会改变lisa本身
	lisa.notify()
	//运行结果：
	// Sending User Email To Lisa<lisa@email.com>
	// in func lisa@comcast0.com
	// Sending User Email To Lisa<lisa@email.com>

	add(1, 2)
	add(1, 3, 7)
	// add([]int{1, 2}) 写法错误
	add([]int{1, 3, 7}...) // 和append一样需要加...

	//
	var fragment Fragment = new(GetPodAction)
	// var fragmen2t Fragment = &GetPodAction{} //  指针对象既算实现了（结构体）.方法,也算实现了（*结构体）.方法
	// var fragment3 Fragment = GetPodAction{}  // 非指针对象只算实现了（结构体）.方法,不算实现了（*结构体）.方法
	var fragment1 Fragment1 = fragment // 只要两个接口拥有相同的方法列表（次序不同不要紧），那么它们就是等价的，可以相互赋值
	fmt.Println(fragment1)
}

// type user struct {
// 	name  string
// 	email string
// }
// func (u *user) changeEmail(email string){ u.email = email }
// func (u user) notify(){	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email) }
// bill := user{"Bill", "bill@email.com"}
// bill.notify()
// bill.changeEmail("bill@newdomain.com")
// //bill为类型值，而changeEmail方法的接收者为指针，编译器实际会做如下操作：
// // (&bill).changeEmail("bill@newdomain.com")，对指针指向的值进行修改，会改变bill本身的字段（属性）值
// bill.notify()
// // 运行结果：
// // Sending User Email To Bill<bill@email.com>
// // Sending User Email To Bill<bill@newdomain.com>
// bill := &user{"Bill", "bill@email.com"}
// bill.notify()
// bill.changeEmail("bill@newdomain.com")
// // 运行结果：
// // Sending User Email To Bill<bill@email.com>
// // Sending User Email To Bill<bill@newdomain.com>

func add(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

type Fragment interface {
	Exec(transInfo *TransInfo) error
}

type Fragment1 interface {
	Exec(transInfo *TransInfo) error
}
type GetPodAction struct {
}
type TransInfo struct {
}

func (g GetPodAction) Exec(transInfo *TransInfo) error {
	return nil
}

type g int

func (c g) Exec(transInfo *TransInfo) error { // 编译通过
	return nil
}

// func (g int) Exec(transInfo *TransInfo) error {   // 编译不通过，int不是自定义的类型
// 	return nil
// }
