// 10、以下代码能编译过去吗？为什么？

package main

import (
	"fmt"
)

type People interface {
	Speak(string) string
}
type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}
func main() {
	var peo People = Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}

// 答： 编译失败，值类型 Student{} 未实现接口People的方法，不能定义为 People类型。

// 解析：

// 考题中的 func (stu Stduent) Speak(think string) (talk string) 是表示结构类型Student的指针有提供该方法，但该方法并不属于结构类型Student的方法。因为struct是值类型。

// 修改方法：

// 定义为指针：

// var peo People = &Stduent{}
// 方法定义在值类型上，指针类型本身是包含值类型的方法。

// func (stu Stduent) Speak(think string) (talk string) {
// 	//...
// }
