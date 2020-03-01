// 快排， 设头尾变量了l和r，以a[l]为基准数据t，从a[r]开始和t比较，a[r]>=t则r--,否则a[r]和a[l]数据互换，然后从a[l]开始和t比较，做相同操作l++，最终l=r时，
// 基准数据就换到了最终正确的位置。然后使用递归把基准数据两边的数组再进行递归操作
package main

import "fmt"

func main() {
	arry := []int{1, 10, 6, 3, 7, 9, 5, 4, 2, 11}
	quickSort1(arry, 0, len(arry)-1)
	fmt.Println(arry)

}

func quickSort1(arr []int, l, r int) {
	if l >= r {
		return
	}
	left := l
	right := r
	t := arr[l]
	for l < r {
		for arr[r] >= t && r > l {
			r--
		}
		arr[l], arr[r] = arr[r], arr[l]
		for arr[l] <= t && r > l {
			l++
		}
		arr[l], arr[r] = arr[r], arr[l]
	}
	quickSort1(arr, left, l-1)
	quickSort1(arr, l+1, right)
}
