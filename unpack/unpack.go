package main

import (
	"fmt"
)

func arbFunction(a ...string) {
	for _, v := range a {
		fmt.Println(v)
	}
}

func main() {
	var stringslice = []string{"hello", "dear", "world"}
	arbFunction(stringslice...)
}
