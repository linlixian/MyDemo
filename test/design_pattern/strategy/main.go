//工厂模式中只管生产实例，具体怎么使用工厂实例由调用方决定，
//策略模式是将生成实例的使用策略放在策略类中配置后才提供调用方使用。
//工厂模式调用方可以直接调用工厂实例的方法属性等，策略模式不能直接调用实例的方法属性，需要在策略类中封装策略后调用。
package main

import (
	// "MyProject/models/design_pattern/strategy/computer/computer"
	// "MyProject/test/design_pattern/strategy/strategy/strategy"
	"MyProject/models/computer"
	"MyProject/models/strategy"
	"flag"
	"fmt"
)

var stra *string = flag.String("type", "a", "input the strategy")
var num1 *int = flag.Int("num1", 1, "input num1")
var num2 *int = flag.Int("num2", 1, "input num2")

func init() {
	flag.Parse()
}

func main() {
	com := computer.Computer{Num1: *num1, Num2: *num2}
	fmt.Println(*stra)

	strate := strategy.NewStrategy(*stra)

	com.SetStrategy(strate)
	fmt.Println(com.Do())
}
