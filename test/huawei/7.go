package main

import (
	"fmt"
)

func main() {
	for {
		var n string
		x, _ := fmt.Scan(&n)
		if x == 0 {
			return
		}
		arr := []rune(n)
		var a []rune
		m := make(map[rune]int)
		for i := len(arr) - 1; i >= 0; i-- {
			if _, ok := m[arr[i]]; !ok {
				m[arr[i]] = 1
				a = append(a, arr[i])
			}
		}
		fmt.Print(string(a))
	}

}
