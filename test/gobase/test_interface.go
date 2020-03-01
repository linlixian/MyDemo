package main

import (
	"fmt"
	"reflect"
)

type persion interface {
	sayHello(hello_word string) string
	sayFuck(hello_word string) string
}

type sb interface {
	sayHello111(hello_word string) string
}

type stu struct {
	name        string
	school_name string
}

type teacher struct {
	teacher_id string
	years      int
}

func (t teacher) sayHello(hello_word string) string {
	return "my tech id is " + t.teacher_id + " and hello --> " + hello_word
}

func (s stu) sayHello(hello_word string) string {
	return "my stu name is " + s.name + " and hello --> " + hello_word
}

func (t teacher) sayFuck(hello_word string) string {
	return "my tech id is " + t.teacher_id + " and fuck --> " + hello_word
}

func (s stu) sayFuck(hello_word string) string {
	return "my stu name is " + s.name + " and fuck --> " + hello_word
}

func main() {
	stu_test := stu{}
	stu_test.name = "yuyong"
	stu_test.school_name = "a school"
	teh_test := new(teacher)
	teh_test.teacher_id = "1001"
	teh_test.years = 4

	var people [2]persion
	people[0] = stu_test // 向上转型
	people[1] = teh_test // 向上转型

	fmt.Println(reflect.TypeOf(people[0]))  // main.stu
	fmt.Println(reflect.ValueOf(people[0])) // {yuyong a school}
	check := people[0].(stu)                // 向下转型
	check1 := people[1].(*teacher)          // 向下转型
	fmt.Println(check.sayHello("1111"))     // (s stu) sayHello
	fmt.Println(check1.sayHello("1111"))    // (s teacher) sayHello

	var people1 persion
	people1 = stu_test                              // 向上转型
	fmt.Println(people1.sayFuck("1111"))            // 1111
	fmt.Println(persion(&stu_test).sayFuck("1111")) // 1111
	fmt.Println(persion(stu_test).sayFuck("1111"))  // 1111

	// new(stu)等同于&stu{}
	stu1 := new(stu)
	stu2 := &stu{}
	fmt.Println(stu1)
	fmt.Println(stu2)

	channels := make([]*teacher, 10)
	t1 := &teacher{}
	channels[0] = t1

	xcas1 := make(map[string]([]*stu))
	xcas2 := make(map[string][]*interface{})
	xcas3 := map[string]stu{}
	xcas4 := map[string]interface{}{}
	fmt.Println(xcas1)
	fmt.Println(xcas2)
	fmt.Println(xcas3)
	fmt.Println(xcas4)

	// range迭代是复制数据,如果是指针类型，复制地址，操作地址改变数据，原数据也会改变

	channels1 := make([]*teacher, 2)
	t11 := &teacher{
		teacher_id: "a",
		years:      1,
	}
	channels1[0] = t11
	channels1[1] = t11
	fmt.Println(channels1[0].years)

	for _, v := range channels1 {
		v.years = 2
	}
	fmt.Println(channels1[0].years)

}
