package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func md5Hash(s string) string {
	h := md5.New()
	io.WriteString(h, "hello md5 world")

	// base-16, lowercase
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {
	fmt.Println(md5Hash("Hello world"))
}
