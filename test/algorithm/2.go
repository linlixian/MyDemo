package main

import "fmt"

// 搜索二维矩阵 II
// 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：

// 每行的元素从左到右升序排列。
// 每列的元素从上到下升序排列。
// 示例:

// 现有矩阵 matrix 如下：

// [
//   [1,   4,  7, 11, 15],
//   [2,   5,  8, 12, 19],
//   [3,   6,  9, 16, 22],
//   [10, 13, 14, 17, 24],
//   [18, 21, 23, 26, 30]
// ]
// 给定 target = 5，返回 true。

// 给定 target = 20，返回 false。

func main() {
	// nums1 := [][]int{
	// 	{1, 4, 7, 11, 15},
	// 	{2, 5, 8, 12, 19},
	// 	{3, 6, 9, 16, 22},
	// 	{10, 13, 14, 17, 24},
	// 	{18, 21, 23, 26, 30}}
	nums1 := [][]int{}
	fmt.Println(searchMatrix(nums1, 31))
}

func searchMatrix(matrix [][]int, target int) bool {
	r := 0
	c := 0
	if len(matrix) > 0 {
		c = len(matrix[0]) - 1
	}
	for c >= 0 && r <= len(matrix)-1 {
		if matrix[r][c] == target {
			return true
		} else if matrix[r][c] > target {
			c--
		} else if matrix[r][c] < target {
			r++
		}
	}
	return false
}
