package main

import (
	"fmt"
)

func main() {
	for {
		var count int
		n, _ := fmt.Scan(&count)
		if n == 0 {
			return
		}
		var str string
		var strs [200]string
		for i := 0; i < count; i++ {
			n, _ := fmt.Scan(&str)
			if n == 0 {
				return
			}
			rs := []rune(str)
			if rs[0] < 97 {
				rs[0] -= 32
			}
			if strs[rs[0]] == "" {
				strs[rs[0]] = str
			} else {
				strs[rs[0]] = strs[rs[0]] + "\n" + str
			}
		}
		for _, v := range strs {
			if v != "" {
				fmt.Println(v)
			}

		}

	}
}
