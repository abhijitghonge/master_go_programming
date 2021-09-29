package main

import "fmt"

func main() {
	var m1 map[int]int

	fmt.Println("Map:", m1)

	m2 := map[int]string{
		1: "Anna",
		2: "Appa",
	}

	m2[10] = "Abba"
	fmt.Printf("Existing:%v\n", m2[10])
	fmt.Printf("Non Existing:%v\n", m2[3])

	for k, v := range m2 {
		fmt.Printf("Key:%v, Value:%v\n", k, v)
	}

}
