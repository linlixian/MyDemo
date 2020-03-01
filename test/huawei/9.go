package main

import (
	"fmt"
)

func main() {
	for {
		var num int
		var count int
		var out []int
		for count < 10 {
			n, _ := fmt.Scan(&num)
			if n == 0 || num == 0 {
				break
			}
			var x int
			for num > 1 {
				x = x + num/3
				num = num/3 + num%3
				if num == 2 {
					x += 1
					num = 0
				}
			}
			out = append(out, x)

			count++
		}
		if len(out) > 0 {
			for _, v := range out {
				fmt.Println(v)
			}
		}

	}
}
