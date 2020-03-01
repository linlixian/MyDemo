package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	t := time.Now()
	h, _ := time.ParseDuration("-1h")
	h1 := t.Add(h)
	d := t.Sub(h1)
	fmt.Println(int(d.Seconds()))

	fmt.Println(SecToTime(3600))
	tt := time.Now()
	fmt.Println(tt)
	// fmt.Println(time.Now().Add(1 * time.Minute))
	asdasd := time.Duration(tt.Second()) * time.Second
	ss := time.Now().Add(0 * time.Minute).Add(-asdasd)
	fmt.Println(ss)
	fmt.Println(tt.Format("2006-01-02 15:04:01"))
	fmt.Println(time.Now().Add(1*time.Minute).Unix() - time.Now().Unix())
	fmt.Println(time.Now().Minute())
	fmt.Println(time.Now().Add(-5 * time.Minute))

	tempV, _ := time.Parse("15:04:05", "18:40:00")
	fmt.Println(tempV)

}

// 秒数转换为时分秒格式
func SecToTime(time int) string {
	timeStr := ""
	hour := 0
	minute := 0
	second := 0
	if time <= 0 {
		return "00:00"
	} else {
		minute = time / 60
		if minute < 60 {
			second = time % 60
			timeStr = unitFormat(minute) + ":" + unitFormat(second)
		} else {
			hour = minute / 60
			if hour > 99 {
				return "99:59:59"
			}
			minute = minute % 60
			second = time - hour*3600 - minute*60
			timeStr = unitFormat(hour) + ":" + unitFormat(minute) + ":" + unitFormat(second)
		}
	}
	return timeStr
}

func unitFormat(i int) string {
	retStr := ""
	if i >= 0 && i < 10 {
		retStr = "0" + strconv.Itoa(i)
	} else {
		retStr = "" + strconv.Itoa(i)
	}
	return retStr
}
