package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
loop:
	for {
		data, _, _ := reader.ReadLine()
		if len(data) == 0 {
			break
		}
		count, _ := strconv.Atoi(string(data))
		m := make(map[int]int)
		var arr []int
		for i := 0; i < count; i++ {
			data1, _, _ := reader.ReadLine()
			if len(data) == 0 {
				break loop
			}
			strs := strings.Split(string(data1), " ")
			k, _ := strconv.Atoi(strs[0])
			v, _ := strconv.Atoi(strs[1])
			if _, ok := m[k]; ok {
				m[k] += v
			} else {
				m[k] = v
				arr = append(arr, k)

			}
		}
		sort(arr)
		for _, v := range arr {
			fmt.Println(v, m[v])
		}
	}

}

func sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
}
