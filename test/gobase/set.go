//  【golang】Go实现set类型

package main

import (
	"fmt"
	"sort"
)

type Empty struct{}

var empty Empty

type Set struct {
	m map[int]Empty
}

//返回一个set
func SetFactory() *Set {
	return &Set{
		m: map[int]Empty{},
	}
}

//添加元素
func (s *Set) Add(val int) {
	s.m[val] = empty
}

//删除元素
func (s *Set) Remove(val int) {
	delete(s.m, val)
}

//获取长度
func (s *Set) Len() int {
	return len(s.m)
}

//清空set
func (s *Set) Clear() {
	s.m = make(map[int]Empty)
}

//遍历set
func (s *Set) Traverse() {
	for v := range s.m {
		fmt.Println(v)
	}
}

//排序输出
func (s *Set) SortTraverse() {
	vals := make([]int, 0, s.Len())

	for v := range s.m {
		vals = append(vals, v)
	}

	//排序
	sort.Ints(vals)

	for _, v := range vals {
		fmt.Println(v)
	}
}
