//思路：从顶点（0,0）开始，分别向左右上下遍历，遇见1设置为0，遇见0停止，遍历到4个方向都为1则加1
package main

import "fmt"

func main() {
	var graph1 = [5][5]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}

	// graph1 = [5][5]int{
	//     {1, 1, 1, 1, 0},
	//     {1, 1, 0, 1, 0},
	//     {1, 1, 0, 0, 0},
	//     {0, 0, 0, 0, 0},
	// }

	var count int

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if graph1[i][j] == 1 {
				search(&graph1, i, j)
				count++
			}
		}
	}
	fmt.Println(count)
}

func search(graph1 *[5][5]int, i, j int) {
	if i < 0 || i >= 5 || j >= 5 || j < 0 || graph1[i][j] == 0 {
		return

	}
	graph1[i][j] = 0
	search(graph1, i-1, j)
	search(graph1, i+1, j)
	search(graph1, i, j+1)
	search(graph1, i, j-1)
}
