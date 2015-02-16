package main

import (
	"circle"
	"fmt"
)

func main() {
	var rad float64
	rad = 20.343434
	fmt.Println("Radius, Area", rad, circle.Area(rad))
}
