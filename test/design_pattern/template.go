// 模板模式  （定义模板person结构体，实现模版公共方法）
// 男人和女人穿衣模式不一样，person结构体的Specific IPerson
package main

import (
	"fmt"
)

// template method design pattern
func main() {
	var p *Person = &Person{}

	p.Specific = &Boy{}
	p.SetName("qibin")
	p.Out()

	p.Specific = &Girl{}
	p.SetName("loader")
	p.Out()
}

type IPerson interface {
	SetName(name string)
	BeforeOut()
	Out()
}

type Person struct {
	Specific IPerson
	name     string
}

func (this *Person) SetName(name string) {
	this.name = name
}

func (this *Person) Out() {
	this.BeforeOut()
	fmt.Println(this.name + " go out...")
}

func (this *Person) BeforeOut() {
	if this.Specific == nil {
		return
	}

	this.Specific.BeforeOut()
}

type Boy struct {
	Person
}

func (_ *Boy) BeforeOut() {
	fmt.Println("get up..")
}

type Girl struct {
	Person
}

func (_ *Girl) BeforeOut() {
	fmt.Println("dress up..")
}
