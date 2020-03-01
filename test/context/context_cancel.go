/*
通过context，我们可以方便地对同一个请求所产生地goroutine进行约束管理，可以设定超时、deadline，甚至是取消这个请求相关的所有goroutine。
形象地说，假如一个请求过来，需要A去做事情，而A让B去做一些事情，B让C去做一些事情，A、B、C是三个有关联的goroutine，
那么问题来了：假如在A、B、C还在处理事情的时候请求被取消了，那么该如何优雅地同时关闭goroutine A、B、C呢？这个时候就轮到context包上场了
(即可以处理goroutine的（内存）泄露
还可以传一个stopchannel,close通知其他goroutine return
)
*/
package main

import (
	"context"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go http.ListenAndServe(":8989", nil)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()
	log.Println(A(ctx))
}

func C(ctx context.Context) string {
	select {
	case <-ctx.Done():
		return "C Done"
	}
	return ""
}

func B(ctx context.Context) string {
	ctx, _ = context.WithCancel(ctx)
	go log.Println(C(ctx))
	select {
	case <-ctx.Done():
		return "B Done"
	}
	return ""
}

func A(ctx context.Context) string {
	go log.Println(B(ctx))
	select {
	case <-ctx.Done():
		return "A Done"
	}
	return ""
}
