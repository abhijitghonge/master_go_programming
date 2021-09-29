package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	go func() {
		c <- "ONE"
	}()

	go func() {
		c <- "TWO"
	}()

	time.Sleep(time.Second)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c:
			fmt.Println("Received", msg1)
		case msg2 := <-c:
			fmt.Println("Received", msg2)

		}
	}
}
