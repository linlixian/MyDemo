package main

import (
	"fmt"
	"strconv"
)

// "strings"

//"zcm_crm_zj/models/crm"

// -------------------
//      Interface
// -------------------

func main() {
	s := "0xa"
	// b := []byte(s)
	fmt.Println([]byte(s))
	/*
		参数1 数字的字符串形式

		参数2 数字字符串的进制 比如二进制 八进制 十进制 十六进制

		参数3 返回结果的bit大小 也就是int8 int16 int32 int64
	*/
	base, _ := strconv.ParseInt(s[2:], 16, 0) // str[2:] 因为输入的0xA的0x前缀代表16进制，所以要去掉。a,A都表示16进制的10
	fmt.Println(base)

	xvv := strconv.FormatInt(base, 16) // 参数16 表示16进制
	fmt.Println(xvv)

	ss := "13.1"
	x, _ := strconv.ParseFloat(ss, 64) // 参数64 表示返回结果的bit大小
	fmt.Println(x)

	adsd := strconv.FormatFloat(x, 'f', 2, 64) //参数2表示保留2位小数 参数64 表示返回结果的bit大小
	fmt.Println(adsd)

}
