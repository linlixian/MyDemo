package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	s := MD5("123456")
	fmt.Print(s)
}

func MD5(data string) string {
	m := md5.Sum([]byte(data))
	return hex.EncodeToString(m[:])
}
