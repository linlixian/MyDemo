// 观察者模式 （目标接口有注册方法和发送方法，观察者接口有接受方法，发送方法中调用接收方法）
// （新闻订阅）
package main

import (
	"container/list"
	"fmt"
)

//抽象主题
type Subject interface {
	Add(o Observer)
	Send(str string)
}

//抽象观察者
type Observer interface {
	Receive(str string)
}

//-------------天气主题------------------
type NewsSubject struct {
	title string
	l     *list.List
}

//天气主题非侵入式实现抽象主题
func (sub *NewsSubject) Add(o Observer) {
	sub.l.PushBack(o)
}

//天气主题非侵入式实现抽象主题
func (sub *NewsSubject) Send(str string) {
	for i := sub.l.Front(); i != nil; i = i.Next() {
		(i.Value).(Observer).Receive(sub.title + "发送的：" + str)
	}
}

//-------------热点主题------------------
type HotSubject struct {
	title string
	l     *list.List
}

//热点主题非侵入式实现抽象主题
func (sub *HotSubject) Add(o Observer) {
	sub.l.PushBack(o)
}

//热点主题非侵入式实现抽象主题
func (sub *HotSubject) Send(str string) {
	for i := sub.l.Front(); i != nil; i = i.Next() {
		(i.Value).(Observer).Receive(sub.title + "发送的：" + str)
	}
}

//-------------a观察者------------------
type aObserver struct {
	name string
}

//a观察者非侵入式实现抽象观察者
func (o *aObserver) Receive(str string) {
	fmt.Println("A类观察者【" + o.name + "】接收" + str)
}

//-------------b观察者------------------
type bObserver struct {
	name string
}

//b观察者非侵入式实现抽象观察者
func (o *bObserver) Receive(str string) {
	fmt.Println("B类观察者【" + o.name + "】接收" + str)
}

func main() {
	a := &aObserver{
		name: "张三",
	}

	b := &bObserver{
		name: "李四",
	}

	//新闻允许a和b类型观察者订阅
	news := NewsSubject{
		l:     list.New(),
		title: "武汉新闻",
	}
	news.Add(a)
	news.Add(b)

	//热点只允许b类型观察者订阅
	hot := HotSubject{
		l:     list.New(),
		title: "孝感热点",
	}
	hot.Add(b)

	news.Send("全省暴雨红色警报")
	hot.Send("全市停工停课")
}
