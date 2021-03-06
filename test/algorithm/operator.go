// 运算符

// 假定 A = 60; B = 13; 其二进制数转换为：
// A = 0011 1100

// B = 0000 1101

// -----------------

// 10进制转为2进制方法：
// 10进制数不断除以2。genhao把每一个新的商数除以二，然后把余数写在被除数的右边。直到商数为0为止写出新的二进制数字。
// 从最下面的余数开始，按顺序读到最上面。
/*
   2√8￣     0
   2√4￣     0
   2√2￣     0
   2√1￣     1    所以8的二进制数为1000
*/
// 2进制转为10进制方法：
/*
   1*2^3+0*2^2+0*2^1++0*2^0=8 的二进制数为1000
*/
// A&B = 0000 1100  // 与操作

// A|B = 0011 1101  // 或操作

// A^B = 0011 0001  // 异或操作

package main

import "fmt"

func main() {

	var a uint = 60 /* 60 = 0011 1100 */
	var b uint = 13 /* 13 = 0000 1101 */
	var c uint = 0

	c = a & b /* 12 = 0000 1100 */
	fmt.Printf("第一行 - c 的值为 %d\n", c)

	c = a | b /* 61 = 0011 1101 */
	fmt.Printf("第二行 - c 的值为 %d\n", c)

	c = a ^ b /* 49 = 0011 0001 */
	fmt.Printf("第三行 - c 的值为 %d\n", c)

	c = a << 2 /* 240 = 1111 0000 */
	fmt.Printf("第四行 - c 的值为 %d\n", c)

	c = a >> 2 /* 15 = 0000 1111 */
	fmt.Printf("第五行 - c 的值为 %d\n", c)

	var asd int = 13
	zxc := asd >> 1
	fmt.Println(zxc)
}

/*
第一行 - c 的值为 12
第二行 - c 的值为 61
第三行 - c 的值为 49
第四行 - c 的值为 240
第五行 - c 的值为 15
*/
