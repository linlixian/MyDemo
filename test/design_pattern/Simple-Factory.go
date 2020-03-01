// 简单工厂
package design_pattern

import "fmt"

// AB interface
type AB interface {
	Say(name string) string
}

// A .
type A struct{}

// Say .
func (*A) Say(name string) string {
	return fmt.Sprintf("我是A实例, %s", name)
}

// B .
type B struct{}

// Say .
func (*B) Say(name string) string {
	return fmt.Sprintf("我是B实例, %s", name)
}

// NewAB 根据参数不同返回不同实例
func NewAB(t int) AB {
	if t == 1 {
		return &A{}
	}
	return &B{}
}
