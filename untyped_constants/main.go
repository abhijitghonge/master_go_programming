package main

import (
	"fmt"
)

func main() {
	const a float64 = 5.1

	const b = 6.7

	const d = 5 > 10

	fmt.Println(d)

	const (
		x = -(iota + 2) //-2
		_               //skip -3
		y               //-4
		z               //-5
	)

	fmt.Println(x, y, z)

	//Rune Type
	var r rune = 'f'
	fmt.Printf("%T\n", r)
}
