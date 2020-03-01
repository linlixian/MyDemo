package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	sd := time.Now().Format("2006-01-02")
	s := sd + " " + strconv.Itoa(10) + ":00:00"
	st, _ := time.ParseInLocation("2006-01-02 15:04:05", s, time.Local) // 系统的时区cst 领先UTC八个小时
	stt, _ := time.Parse("2006-01-02 15:04:05", s)                      // 世界时区 utc
	fmt.Println(st)
	fmt.Println(stt)

}
