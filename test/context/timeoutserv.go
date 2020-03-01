package main

import (
	"fmt"
	"net/http"
	"time"
)

func lazyHandler(w http.ResponseWriter, req *http.Request) {
	// ranNum := rand.Intn(2)
	ranNum := 0
	if ranNum == 0 {
		time.Sleep(6 * time.Second)
		fmt.Fprintf(w, "slow response, %d\n", ranNum)
		fmt.Printf("slow response, %d\n", ranNum)
		return
	}
	fmt.Fprintf(w, "quick response12, %d\n", ranNum)
	fmt.Printf("quick response11, %d\n", ranNum)
	return
}

func main() {
	http.HandleFunc("/", lazyHandler)
	http.ListenAndServe(":9200", nil)
}
