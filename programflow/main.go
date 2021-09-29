package main

import "fmt"

func main() {
	var count int
	numbers := 50
	for i := 1; i <= numbers; i++ {
		if i%7 == 0 {
			fmt.Printf("%d, ", i)
			count++
		}
		if count == 3 {
			break
		}
	}

	fmt.Printf("\n")

	count = 0
	numbers = 500

	for i := 1; i <= numbers; i++ {
		if i%5 == 0 && i%7 == 0 {
			fmt.Printf("%v,", i)
		}
	}

	fmt.Println()

	age := 18

	switch {
	case age < 0 || age > 100:
		fmt.Println("Invalid Age")
	case age < 18:
		fmt.Println("You are minor!")
	case age == 18:
		fmt.Println("Congratulations! You've just become major!")
	default:
		fmt.Println("You are major!")
	}

}
