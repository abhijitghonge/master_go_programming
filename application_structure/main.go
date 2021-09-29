package main

import "fmt"

var secondInHour = 3600

func main() {
	fmt.Println("Hello Go World!!")
	distance := 60.8
	fmt.Printf("The distance in miles is %f\n", distance*0.62137)

	var a int = 7
	var b float64 = 3.5
	const c int = 10

	fmt.Println(a, b, c)

	var (
		p int
		q float64
		r rune
	)

	_, _, _ = p, q, r
}
