package main

import "fmt"

func main() {
	var a = 4
	var b = 5.4
	a = int(b)

	fmt.Println(a,b)

	var value int
	var price float64
	var name string
	var done bool

	fmt.Println(value, price, name, done)
}