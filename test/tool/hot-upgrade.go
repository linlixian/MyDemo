/*
2、打开两个终端
先启动第一个终端
一个终端输入请求： "http://127.0.0.1:8080/hello?sleep=20"

3、热升级
kill -HUP 进程ID
启动第二个终端

4、打开一个终端输入请求： "http://127.0.0.1:8080/hello?sleep=0"

我们可以看到这样的结果，第一个请求等待20s，但是处理他的是老的进程，热升级之后，第一个请求还在执行，最后会输出老的进程ID，而第二次请求，输出的是新的进程ID
*/
package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/astaxie/beego/grace"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("WORLD!"))
	// w.Write([]byte("ospid:" + strconv.Itoa(os.Getpid())))
	r.ParseForm()
	a := r.Form["sleep"][0]
	i, _ := strconv.Atoi(a)
	time.Sleep(time.Duration(i) * time.Second)
	w.Write([]byte("ospid:" + strconv.Itoa(os.Getpid())))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handler)

	err := grace.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Println(err)
	}
	log.Println("Server on 8080 stopped")
	os.Exit(0)
}
