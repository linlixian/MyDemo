// break 是跳出最里面的一层for循环和select 这种可以循环的语句，
package main

import "fmt"

func main() {
	values := []int{2, 2, 1}
	singleNumber(values)
	// fmt.Println(singleNumber(values))
}

func singleNumber(nums []int) int {
	var r int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			break
		}
		fmt.Println(nums[i])
	}
	return r
}
