package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	f1 := 5.2
	f1 = math.Ceil(f1)
	fmt.Println(f1)

	dddd := "asdasd"
	fmt.Println(dddd[1:3]) // 字符串截取(不算右边的index)
	fmt.Println(len(dddd)) // 字符串截取
	bq := []string{"a", "s", "d", "a", "s", "d"}
	fmt.Println(bq[0:1]) // 字符串截取(不算右边的index)

	// var f float64
	// fmt.Println(float64(6) / 60 * 1.14)
	f := float64(6) / 60 * 1.14
	fmt.Println(f)

	fmt.Println(SubFloatToFloat(float64(6)/60*1.14, 4))
	fmt.Println(decimal(float64(6)/60*1.14, 4))

	// 1. 两位小数乘100强转int, 比期望值少了1
	var d float64 = 1129.6
	fmt.Println(int64(d * 100))         // 112959 错误
	fmt.Println(int64(Round(d*100, 2))) // 112960 正确

	// 2. 两个浮点数相加减，可能不准确
	x := 74.96
	y := 20.48
	b := x - y
	fmt.Println(b)                     //  错误
	fmt.Println(Round2(b, 2))          //  正确
	fmt.Println(Round(b, 2))           //  正确
	fmt.Println(SubFloatToFloat(b, 2)) //  正确
	fmt.Println(decimal(b, 2))         //  正确

	// 3. float32和float64直接互转会精度丢失, 四舍五入后错误
	var a1 float32 = 80.45
	var b1 float64
	b1 = float64(a1)
	// 打印出值, 强转后出现偏差
	fmt.Println(b1)                     //output:80.44999694824219
	fmt.Println(Round2(b1, 2))          // 正确 80.45
	fmt.Println(Round(b1, 2))           //  正确
	fmt.Println(SubFloatToFloat(b1, 2)) //  正确
	fmt.Println(decimal(b1, 2))         //  正确

	// 4.int64转float64在数值很大的时候出现偏差
	var c int64 = 987654321098765432
	fmt.Printf("%.f\n", float64(c)) //output:987654321098765440
}

func Round2(f float64, n int) float64 {
	floatStr := fmt.Sprintf("%."+strconv.Itoa(n)+"f", f)
	inst, _ := strconv.ParseFloat(floatStr, 64)
	return inst
}

func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

//截取小数点后几位
func SubFloatToFloat(f float64, m int) float64 {
	newn := SubFloatToString(f, m)
	newf, _ := strconv.ParseFloat(newn, 64)
	return newf
}

//截取小数点后几位
func SubFloatToString(f float64, m int) string {
	n := strconv.FormatFloat(f, 'f', -1, 64)
	if n == "" {
		return ""
	}
	if m >= len(n) {
		return n
	}
	newn := strings.Split(n, ".")
	if m == 0 {
		return newn[0]
	}
	if len(newn) < 2 || m >= len(newn[1]) {
		return n
	}
	return newn[0] + "." + newn[1][:m]
}

func decimal(f float64, len int) float64 {
	p := math.Pow(10, float64(len+1))
	p2 := math.Pow(10, float64(len))
	f2 := f * p
	return float64(int(f2/10)) / p2
}
