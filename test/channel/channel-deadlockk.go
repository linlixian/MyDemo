// 如果协程运行完都没有cs <- cakeName，却有<-cs，则会死锁
// 如果协程有cs <- cakeName，却没有检测到<-cs，则会死锁
// chan for-range 必须有close才能退出协程，否则会死锁(for 同理)
// k,v:=range arr  每次v都是同一个地址，指针相等
// v:=range arr  只接受一个返回值，则这个返回值是key
package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func makeCakeAndSend(cs chan string, count int) {
	for i := 1; i <= count; i++ {
		cakeName := "Strawberry Cake " + strconv.Itoa(i)
		cs <- cakeName //send a strawberry cake
		fmt.Println("Strawberry Cake ", strconv.Itoa(i))
	}
	time.Sleep(1 * 1e9)

	// cakeName := "Strawberry Cake 1"
	// cs <- cakeName //send a strawberry cake
	// time.Sleep(5 * 1e9)
	// cs <- cakeName //send a strawberry cake

	close(cs)
}

func receiveCakeAndPack(cs chan string, w *sync.WaitGroup) {
	for s := range cs {
		fmt.Println("Packing received cake: ", s)
	}

	// fmt.Println("Packing received cake: ", <-cs)
	// fmt.Println("wait...")
	// fmt.Println("Packing received cake: ", <-cs)

	w.Done()
}

func main() {
	w := &sync.WaitGroup{}

	cs := make(chan string)
	go makeCakeAndSend(cs, 5)
	// time.Sleep(2 * 1e9)
	w.Add(1)
	go receiveCakeAndPack(cs, w)
	w.Wait()

	strs := []string{"one", "two", "three"}

	for _, s := range strs {
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Printf("%s ", s) // three three three
		}()
	}
	time.Sleep(3 * time.Second)

	x := []string{"a", "b", "c"}
	for v := range x {
		fmt.Print(v)
	}

}
