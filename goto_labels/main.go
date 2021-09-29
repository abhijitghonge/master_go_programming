package main

import "fmt"

func main() {
	//** LABEL STATEMENT **//

	// declaring a variable
	// there is no conflict name between variable and label
	outer := 19
	_ = outer

	// declaring two arrays
	people := [5]string{"Helen", "Mark", "Brenda", "Antonio", "Michael"}
	friends := [2]string{"Mark", "Marry"}

	// searching for a single friend in a list of people.
outer:
	for index, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("Found a friend: %v at index %v", friend, index)
				break outer
			}
		}
	}

}
