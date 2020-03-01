// 求众数
// 给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

// 你可以假设数组是非空的，并且给定的数组总是存在众数。

// 示例 1:

// 输入: [3,2,3]
// 输出: 3
// 示例 2:

// 输入: [2,2,1,1,1,2,2]
// 输出: 2
package main

import (
	"fmt"
)

// 大于一半则加减后一定大于1
func main() {
	nums1 := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums1))
}

func majorityElement(nums []int) int {
	idx := 0
	num := nums[0]
	for i := 0; i < len(nums); i++ {
		if num == nums[i] {
			idx++
		} else {
			idx--
			if idx == 0 { // 关键
				num = nums[i]
				idx++
			}
		}
	}
	return num
}
