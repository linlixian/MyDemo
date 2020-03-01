package main

import (
	"fmt"
)

func main() {
	// 也可以选择是否在创建时指定该map的初始存储能力，如创建了一个初始存储能力为5的map
	// m := make(map[string]string, 5)
	// // 长度：获取键值对数量。
	// len := len(m)
	// fmt.Printf("map's len is %v\n", len)
	// val := m["key"]
	// fmt.Printf(val)
	// // 删除，如果 key 不存在，不会出错。
	// delete(m, "key1")
	// if _, ok := m["key1"]; ok {
	// 	delete(m, "key1")
	// 	fmt.Printf("deleted key1 map m : %v\n", m)
	// }

	// map/slice/chan 都必须初始化分配空间，否则会报错
	items := make([][]int, 5)
	item := make([]int, 5)
	itemsss := make([]chan int, 5)
	itemss := make([]map[int]int, 5)
	// map不初始会报错
	for i := 0; i < 5; i++ {
		itemss[i] = make(map[int]int)
	}
	// 第二维数组不初始化，直接使用itemss[0][0]会数组越界，但append不会报错
	for i := 0; i < 5; i++ {
		items[i] = make([]int, 5)
	}
	for i := 0; i < 5; i++ {
		itemsss[i] = make(chan int, 5)
	}
	items[0][0] = 0
	itemss[0][0] = 0
	// item[1] = 3
	itemsss[0] <- 1
	a := <-itemsss[0]
	fmt.Println(a)
	fmt.Println(items)
	fmt.Println(item)
	fmt.Println(itemss)

	// map排序 先获取所有key，把key进行排序，再按照排序好的key，进行遍历
	// m := map[string]string{"q": "q", "w": "w", "e": "e", "r": "r", "t": "t", "y": "y"}
	// var slice []string
	// for k, _ := range m {
	// 	slice = append(slice, k)
	// }
	// fmt.Printf("clise string is : %v\n", slice)
	// sort.Strings(slice[:])
	// fmt.Printf("sorted slice string is : %v\n", slice)
	// for _, v := range slice {
	// 	fmt.Println(m[v])
	// }

	type user struct{ name string }
	/*
	   当 map 因扩张而重新哈希时，各键值项存储位置都会发生改变。
	   因此，map 被设计成 not addressable。
	   类似 m[1].name 这种期望透过原 value 指针修改成员的行为自然会被禁 。
	*/
	m := map[int]user{ //n

		1: {"user1"},
	}
	// m[1].name = "Tom" // ./main.go:16:12: cannot assign to struct field m[1].name in map
	fmt.Println(m)
	fmt.Println(m[1].name)

	// 正确做法是完整替换 value 或使用指针。
	u := m[1]
	u.name = "Tom"
	m[1] = u // 替换 value。

	m2 := map[int]*user{
		1: &user{"user1"},
	}

	m2[1].name = "Jack" // 返回的是指针复制品。透过指针修改原对象是允许的。
	fmt.Println(m2)
	fmt.Println(m2[1].name)
}
