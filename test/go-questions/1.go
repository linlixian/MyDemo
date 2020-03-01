// 2、以下代码有什么问题，说明原因
package main

import (
	"fmt"
)

type student struct {
	Name string
	Age  int
}

func pase_student() map[string]*student {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	return m
}
func main() {
	students := pase_student()
	for k, v := range students {
		fmt.Printf("key=%s,value=%v \n", k, v)
	}
}

// 运行结果：
// key=zhou,value=&{wang 22}
// key=li,value=&{wang 22}
// key=wang,value=&{wang 22}
// 答：
// 输出的均是相同的值：&{wang 22}
// 解析：
// 因为for range遍历时，变量stu指针不变，每次遍历仅进行struct值拷贝，故m[stu.Name]=&stu实际上一致指向同一个指针，最终该指针的值为遍历的最后一个struct的值拷贝。形同如下代码：
// var stu student
// for _, stu = range stus {
// 	m[stu.Name] = &stu
// }
// 修正方案，取数组中原始值的指针：
// for i, _ := range stus {
// 	stu := stus[i]
// 	m[stu.Name] = &stu
// }
