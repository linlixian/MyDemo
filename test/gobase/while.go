//golang 里面没有while关键字，可以用for
package main

import (
	"fmt"
)

func main() {
	i := 0
	// while
	for i < 10 {
		i++
		fmt.Println(i)
	}

	//do  while
	for {
		i++
		if i >= 10 {
			break
		}
	}
	fmt.Println(i)

}
