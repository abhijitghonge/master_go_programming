/////////////////////////////////
// Intro to Channels
// Go Playground: https://play.golang.org/p/Uc7iiqVeZLL
/////////////////////////////////

package main

import "fmt"

func readString(ch chan string, sentence string) {
	ch <- sentence
}

func readInto(ch chan int, num int) {
	ch <- num
}

func main() {
	// Declaring a channel of type `chan int`
	var c1 chan int
	fmt.Println(c1) // => nil (its zero value is nil)

	// Initializing the channel
	c1 = make(chan int)
	fmt.Println(c1) // => 0xc000078060 (the channel stores an address)

	// Declaring and initializing a channel at the same time
	c2 := make(chan int)
	_ = c2

	// Declaring and initilizing a RECEIVE-ONLY channel
	c3 := make(<-chan string)

	// Declaring and initilizing a SEND-ONLY channel
	c4 := make(chan string)

	fmt.Printf("%T, %T, %T\n", c1, c3, c4) // => chan int, <-chan string, chan<- string

	//** The "arrow" indicates the direction of data flow!! **//

	go readInto(c1, 10)

	// Receiving a value from the channel
	num := <-c1

	fmt.Println("int received:", num)

	go readString(c4, " I love golang")
	// Waiting for a value to be sent into the channel and print out that value
	fmt.Println("Value received: ", <-c4)

	// Closing a channel
	close(c1)
	close(c4)

}
