/*
造车
（建造者接口包含多个传入属性返回建造者接口的方法，还有一个返回产品的builder方法）
建造者模式将一个复杂对象的构建与它的表示分离，使得同样的构建过程可以创建不同的表示。
建造者模式是一步一步创建一个复杂的对象，它允许用户只通过指定复杂对象的类型和内容就可以构建它们，用户不需要知道内部的具体构建细节。
应用：excel对象添加表头
*/
package main

import "fmt"

// Car 产品角色(汽车)
type Car struct {
	Brand string
	Type  string
	Color string
}

// Drive 骑乘显示
func (this *Car) Drive() error {
	fmt.Printf("A %s %s %s car is running on the road!\n", this.Color, this.Type, this.Brand)
	return nil
}

// Builder 建造者角色
type Builder interface {
	PaintColor(color string) Builder
	AddBrand(brand string) Builder
	SetType(t string) Builder
	Build() Car
}

// ConcreteBuilder 具体的建造者
type ConcreteBuilder struct {
	ACar Car
}

func (this *ConcreteBuilder) PaintColor(cor string) Builder {
	this.ACar.Color = cor
	return this
}

func (this *ConcreteBuilder) AddBrand(bnd string) Builder {
	this.ACar.Brand = bnd
	return this
}

func (this *ConcreteBuilder) SetType(t string) Builder {
	this.ACar.Type = t
	return this
}

func (this *ConcreteBuilder) Build() Car {
	return this.ACar
}

// Director 角色
type Director struct {
	Builder Builder
}

func main() {
	dr := Director{&ConcreteBuilder{}}
	adCar := dr.Builder.SetType("SUV").AddBrand("奥迪").PaintColor("white").Build()
	adCar.Drive()

	bwCar := dr.Builder.SetType("sporting").AddBrand("宝马").PaintColor("red").Build()
	bwCar.Drive()
}
