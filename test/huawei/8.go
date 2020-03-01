package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	for {
		reader := bufio.NewReader(os.Stdin)
		data, _, _ := reader.ReadLine()
		if len(data) == 0 {
			break
		}
		var num int
		var arr [128]bool
		for _, v := range data {
			if v >= 0 && v <= 127 {
				if !arr[v] {
					arr[v] = true
					num++
				}

			}
		}
		fmt.Println(num)
	}

}
