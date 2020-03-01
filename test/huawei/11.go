package main

import "fmt"

func main() {
	resi := [3]int{1, 2, 3}
	resi2 := [3]int{1, 3, 2}
	m1 := make(map[[3]int]bool)
	m1[resi] = true
	m1[resi2] = true
	fmt.Println(m1)
}
