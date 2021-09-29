package main

import (
	"fmt"
)

func power(i int, ch chan int) {
	ch <- i * i
}

func main() {
	c := make(chan int, 10)

	defer close(c)

	for i := 1; i <= 50; i++ {
		fmt.Printf("SEnding for %d\n", i)
		go power(i, c)
		fmt.Printf("after sending for %d\n", i)

	}

	for i := 1; i < 50; i++ {
		fmt.Printf("The power of %d is %d\n", i, <-c)
	}
}
