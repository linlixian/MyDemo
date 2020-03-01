// 继承：如果只有子结构体有唯一此方法，则可以直接 父.方法()调用，如果此方法在子结构体和父结构体都存在，则只能 父.子.方法()调用
package main

import (
	"fmt"
)

type Car struct {
	Name string
	Age  int
}

// func (c *Car) Set(name string, age int) {
// 	c.Name = name
// 	c.Age = age
// }

func (c *Car2) Set(name string) {
	c.Name = name
}

type Car2 struct {
	Name string
}

//Go有匿名字段特性
type Train struct {
	Car
	Car2
	// createTime time.Time
	//count int   正常写法，Go的特性可以写成
	int
}

// //给Train加方法，t指定接受变量的名字，变量可以叫this，t，p
// func (t *Train) Set(name string) {
// 	t.int = 1111
// }

func main() {
	var train Train
	// train.int = 300 //这里用的匿名字段写法，给Age赋值
	train.Set("1010")
	// train.Car.Set("huas123", 1010)
	// train.Car2.Set("huas1")
	fmt.Println(train)

}
