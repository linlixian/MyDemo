package main

import (
	"fmt"
)

// 归并排序（适合大文件排序）
// 把数组对半拆分，递归拆分成一个个只含有两个元素的数组，然后分别排序，最后合并，合并算法是把两个小块的最小值对比，更小的放到一个新数组，然后被取的数组指针后移，再和另一个小块的最小值对比。。
// 最后新数组存放了这两个小块的正确排序
func hebing(arr []int, left int, mid int, right int) {
	var arr2 [10]int
	i, j := left, mid+1
	cnt := 0
	for i <= mid && j <= right {
		if arr[i] < arr[j] {
			arr2[cnt] = arr[i]
			cnt += 1
			i += 1
		} else {
			arr2[cnt] = arr[j]
			cnt += 1
			j += 1
		}
	}
	for i <= mid {
		arr2[cnt] = arr[i]
		cnt += 1
		i += 1
	}
	for j <= right {
		arr2[cnt] = arr[j]
		cnt += 1
		j += 1
	}
	for tm1, tm2 := 0, left; tm2 <= right; tm1, tm2 = tm1+1, tm2+1 {
		arr[tm2] = arr2[tm1]
	}
}

func mergesort(arr []int, left int, right int) {
	if left >= right {
		return
	}
	mid := (left + right) / 2
	fmt.Println(left, mid, right)
	// 左半边递归到底，返回两个数，然后合并排序
	mergesort(arr, left, mid)
	mergesort(arr, mid+1, right)
	hebing(arr, left, mid, right)
}

func main() {
	arr := []int{1, 10, 6, 3, 7, 9, 5, 4, 2, 11}
	mergesort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
