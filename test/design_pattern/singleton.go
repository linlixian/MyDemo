package main

import (
	"fmt"
	"sync"
)

var o sync.Once // 只执行一次
var m *manager

func GetInstance() *manager {
	o.Do(func() {
		m = &manager{}
	})

	// if m == nil {
	// 	m = &manager{}
	// }
	return m
}

type manager struct {
}

func main() {

	fmt.Println("111111111", GetInstance())

}
