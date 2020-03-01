// 装饰者模式 (新建装饰类，类属性中有被装饰的类（组合），然后实现被装饰类一样的方法)
// 拉面加蛋加肠 SetNoddles 输出新价格和名称
package main

import "fmt"

// 面
type Noddles interface {
	Description() string
	Price() float32
}

// 拉面
type Ramen struct {
	name  string
	price float32
}

func (p Ramen) Description() string {
	return p.name
}

func (p Ramen) Price() float32 {
	return p.price
}

// 加鸡蛋

type Egg struct {
	noddles Noddles
	name    string
	price   float32
}

func (p Egg) SetNoddles(noddles Noddles) {
	p.noddles = noddles
}

func (p Egg) Description() string {
	return p.noddles.Description() + "+" + p.name
}

func (p Egg) Price() float32 {
	return p.noddles.Price() + p.price
}

// 加香肠
type Sausage struct {
	noddles Noddles
	name    string
	price   float32
}

func (p Sausage) SetNoddles(noddles Noddles) {
	p.noddles = noddles
}

func (p Sausage) Description() string {
	return p.noddles.Description() + "+" + p.name
}

func (p Sausage) Price() float32 {
	return p.noddles.Price() + p.price
}

func main() {
	ramen := Ramen{name: "ramen", price: 10}
	egg := Egg{noddles: ramen, name: "egg", price: 2}
	sausage := Sausage{noddles: egg, name: "Sausage", price: 2}

	fmt.Println(sausage.Description())
	fmt.Println(sausage.Price())
}
