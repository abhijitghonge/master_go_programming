// run this program in a terminal with arguments
// ex: go run main.go 5 7.1 9.9 10
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 || len(os.Args) > 11 {
		log.Fatal("Please enter 2 to 10 numbers as input")
	}

	args := os.Args[1:]
	sum, product := 0, 1
	for _, input := range args {
		if value, err := strconv.Atoi(input); err != nil {
			fmt.Println(err)
			continue //skip parsing
		} else {
			sum = sum + value
			product = product * value
		}

	}

	fmt.Printf("Sum:%d Product:%d\n", sum, product)

	friends := []string{"Marry", "John", "Paul", "Diana"}
	// f := make([]string, len(friends))
	f := []string{}

	f = append(f, friends...)

	friends[0] = "Abhijit"

	fmt.Printf("The Friends: %#v\n", friends)
	fmt.Printf("Copy of friends: %#v\n", f)

	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}

	newYears := append(years[:3], years[len(years)-3:]...)
	fmt.Printf("New years: %#v\n", newYears)

}
