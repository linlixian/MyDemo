// 给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。

// 说明:

// 初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
// 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
// 示例:

// 输入:
// nums1 = [1,2,3,0,0,0], m = 3
// nums2 = [2,5,6],       n = 3

// 输出: [1,2,2,3,5,6]
package main

import (
	"fmt"
)

func main() {
	nums1 := []int{0}
	nums2 := []int{1}
	merge(nums1, 0, nums2, 1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	idx1 := m - 1
	idx2 := n - 1
	i := m + n - 1
	fmt.Println("idx1=", idx1)
	fmt.Println("idx2=", idx2)
	for idx1 >= 0 && idx2 >= 0 {
		fmt.Println("idx1=", idx1)
		fmt.Println("idx2=", idx2)

		if nums1[idx1] < nums2[idx2] {
			nums1[i] = nums2[idx2]
			idx2--
			i--
		} else if nums1[idx1] >= nums2[idx2] {
			nums1[i] = nums1[idx1]
			idx1--
			i--
		}
	}
	if idx1 < 0 {
		for idx2 >= 0 {
			nums1[idx2] = nums2[idx2]
			idx2--
		}
	}
	fmt.Println("nums5=", nums1)
}
