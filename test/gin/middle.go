package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.Use(PrintRequest)
	// r.Use(PrintResponse)
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.String(200, "pong")
	// 	fmt.Println("this is request1111111")

	// })
	r.GET("/", func(ctx *gin.Context) {

		fmt.Println(1)
		// 这个是最后一个，因此就不必要调用 ctx.Next() 了，这里是我们常用的控制器方法
	}, func(ctx *gin.Context) {
		fmt.Println(2)

		ctx.Next()

		fmt.Println(5)
		// 这里是我们的中间件
	}, func(ctx *gin.Context) {
		fmt.Println(3)
		ctx.Next()
		fmt.Println(4)

	})
	// add your handlers here
	r.Run()
}

func PrintRequest(c *gin.Context) {
	fmt.Println("this is request222222222222")
}

func PrintResponse(c *gin.Context) {
	now := time.Now()
	// 先执行handler下面中间件和handler
	c.Next()
	fmt.Printf("this is response3333333333333, cost: %d",
		time.Since(now).Nanoseconds())
}
