// 一种是 uint8 类型，或者叫 byte 型，代表了 ASCII 码的一个字符。
// 另一种是 rune 类型，代表一个 UTF-8 字符，当需要处理中文、日文或者其他复合字符时，则需要用到 rune 类型。rune 类型等价于 int32 类型。
// 当字符串只有数字 和字母时，rune 和 byte的大小相同（UTF-8包含ASCII 码）
// Go 语言的字符串是不可变的。修改字符串时，可以将字符串转换为 []byte 进行修改。[]byte 和 string 可以通过强制类型转换互转。
// []byte和[]rune可以和string互相转换，byte和rune可以直接赋值字符，如'x'
package main

import (
	"fmt"
	"strings"
)

func main() {
	var apps []string
	apps = []string{"我", "2", "abcd"}
	fmt.Println(apps[2][0:1]) // a
	// apps[2][0:1] = "x"        // 编译不过
	// apps[2][0] = rune('x')  // 编译不过
	// apps[2][0] = byte('x')  // 编译不过
	// apps[2][0] = 'x'  // 编译不过
	fmt.Println(apps[2][0]) // 97
	angleBytes := []byte(apps[2])
	for i := 0; i <= 3; i++ {
		if angleBytes[i] == 'c' {
			angleBytes[i] = 'x' // byte和rune可以直接赋值字符，如'x'
		}
	}
	fmt.Println(string(angleBytes))

	fmt.Println(fmt.Sprintf("%c", apps[2][0])) // a  %c表示字符型格式符，%s表示字符串
	asd := []rune(apps[2])
	fmt.Println(asd[0])         // 97
	fmt.Println(string(asd[0])) // a

	appname := strings.Join(apps, "")
	//     b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(appname)
	fmt.Println(len(appname))
	apps = strings.Split(appname, ",")
	fmt.Println(apps)

	xcv := "3"
	fmt.Println(len(xcv))
	asdads := []rune(xcv)
	fmt.Println(asdads)
	asdads[0] = rune('b')
	fmt.Println(string(asdads))

	// xcv[0] = "asd" // 编译不通过

	var s string = "hello go"
	//返回count个s串联的字符串
	s3 := strings.Repeat(s, 2)
	fmt.Println(s3) //hello gohello go

	//返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串
	s4 := strings.Replace(s, "o", "e", -1)
	fmt.Println(s4) //helle ge

	//去除字符串的空格符，并且按照空格分割返回slice
	s7 := " hello go "
	s8 := strings.Fields(s7)
	fmt.Println(s8)      //[hello go]
	fmt.Println(len(s8)) //

	//字符串切割，如果后一个参数为空，则按字符切割，返回字符串切片
	s9 := strings.Split(s7, " ")
	fmt.Println(s9)      //[hello go]
	fmt.Println(len(s9)) //
	s10 := strings.Split(s7, "")
	fmt.Println(s10)      //[hello go]
	fmt.Println(len(s10)) //

	// 输出格式化的字符串。`Sprintf` 则格式化并返回一个字 符串而不带任何输出。Printf是打印
	// sad := fmt.Sprintf("是字符串 %s ", "string")

	// var buffer bytes.Buffer

	// s := time.Now()
	// for i := 0; i < 100000; i++ {
	// 	buffer.WriteString("test is here\n")
	// }
	// buffer.String() // 拼接结果
	// e := time.Now()
	// fmt.Println("taked time is ", e.Sub(s).Seconds())

	// s = time.Now()
	// str := ""
	// for i := 0; i < 100000; i++ {
	// 	str += "test is here\n"
	// }
	// e = time.Now()
	// fmt.Println("taked time is ", e.Sub(s).Seconds())

	// s = time.Now()
	// var sl []string
	// for i := 0; i < 100000; i++ {
	// 	sl = append(sl, "test is here\n")
	// }
	// strings.Join(sl, "")
	// e = time.Now()
	// fmt.Println("taked time is", e.Sub(s).Seconds())

	// []byte和[]rune可以和string互相转换
	var data [10]rune
	data[0] = 'T'
	data[1] = 'E'
	var strs string = string(data[:]) //string(data) 编译不过string 不能直接和byte数组转换，string可以和byte的切片转换
	fmt.Println(strs)
	fmt.Println(data)
}
