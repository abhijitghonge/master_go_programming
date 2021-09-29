package main

import (
	"fmt"
	"math"
)

func main() {
	a, b := 4, 2
	r := (a + b) / (a - b) * 2

	fmt.Println(r)
	f := float32(math.MaxFloat32)
	fmt.Println(f)
	f = f * 1.2
	fmt.Println(f)
}
