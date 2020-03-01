// 使用sort包排序，需要实现 Swap  Len  Less 方法
// sort的实现是快排
package main

import (
	"fmt"
	"sort"
)

type SubList []int

func (p SubList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p SubList) Len() int {
	return len(p)
}

// index i should sort before the element with index j.
func (p SubList) Less(i, j int) bool {
	return p[i] < p[j] // 当p[i] < p[j],Less为true,p[i]排在p[j]之前。从小到大
	// return p[i] > p[j] // 当p[i] > p[j],Less为true,p[i]排在p[j]之前。从大到小
}
func main() {
	st_list := SubList{}
	st_list = []int{2, 3, 1, 4}
	sort.Sort(st_list) // 排序
	fmt.Println(st_list)
	sort.Sort(sort.Reverse(st_list)) // 倒排 reverse
	fmt.Println(st_list)

	st_list1 := []int{2, 3, 1, 4}
	sort.Ints(st_list1) // sort 自带int排序
	fmt.Println(st_list1)

	st_list2 := []string{"a", "c", "b", "a"}
	sort.Strings(st_list2) // sort 自带string排序
	fmt.Println(st_list2)
	sort.Sort(sort.Reverse(sort.StringSlice(st_list2))) // string反转排序
	fmt.Println(st_list2)

	ints := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0} // 基本类型反转排序
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println(ints) // [9 8 7 6 5 4 3 2 1 0]

}
