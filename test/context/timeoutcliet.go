// 三种方式实现请求超时放弃，time.timer,time.after,context.withtimeout
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

type ResPack struct {
	r   *http.Response
	err error
}

func work(ctx context.Context) {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	defer wg.Done()
	c := make(chan ResPack, 1)

	req, _ := http.NewRequest("GET", "http://localhost:9200", nil)
	req.Close = false
	go func() {
		resp, err := client.Do(req)
		pack := ResPack{r: resp, err: err}
		c <- pack
	}()
	// t := time.NewTimer(3 * time.Second)
	s := time.Now()
	select {
	// case <-ctx.Done():
	// 	tr.CancelRequest(req)
	// 	fmt.Println("Timeout! ctx.Done()")
	// case <-time.After(time.Second * 3):
	// 	fmt.Println("Timeout!  time.After")
	// case <-t.C:
	// 	fmt.Println("Timeout!  time.NewTimer")
	case res := <-c:
		if res.err != nil {
			fmt.Println(res.err)
			return
		}
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Printf("Server Response: %s", out)
	}
	e := time.Now()
	fmt.Println("taked time is ", e.Sub(s).Seconds())
	return
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	wg.Add(1)
	go work(ctx)
	wg.Wait()
	fmt.Println("Finished")
}
