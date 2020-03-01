package main

import "fmt"

type Slice []int

func NewSlice() Slice {
	return make(Slice, 0)
}
func (s *Slice) Add(elem int) *Slice {
	*s = append(*s, elem)
	fmt.Print(elem)
	return s
}
func main() {
	s := NewSlice()

	defer s.Add(1).Add(4).Add(2) // s.add(1)，s.add(4)相当于表达式的一部分先运算
	s.Add(3)

}
