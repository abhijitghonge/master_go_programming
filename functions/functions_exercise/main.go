package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func cube(x float64) float64 {
	return math.Pow(x, 3)
}

func factorial(x int) int {
	if x == 0 {
		return 0
	}

	if x == 1 {
		return 1
	}
	return x * factorial(x-1)
}

func sum(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	return x + sum(x-1)
}

func f(x rune) int {
	first, err := strconv.Atoi(fmt.Sprintf("%c", x))
	if err != nil {
		log.Fatal(err)
	}

	second, err := strconv.Atoi(fmt.Sprintf("%c%c", x, x))
	if err != nil {
		log.Fatal(err)
	}

	third, err := strconv.Atoi(fmt.Sprintf("%c%c%c", x, x, x))
	if err != nil {
		log.Fatal(err)
	}

	return first + second + third
}

func searchItem(items []string, toSearch string) bool {
	for _, item := range items {
		if toSearch == item {
			return true
		}
	}

	return false
}

func searchCaseIns(items []string, toSearch string) bool {
	for _, item := range items {
		if strings.EqualFold(item, toSearch) {
			return true
		}
	}

	return false
}

func main() {
	fmt.Printf("The cube: %v\n", cube(2))
	fmt.Printf("The factorial: %v\n", factorial(4))
	fmt.Printf("The sum: %d\n", sum(10))
	fmt.Printf("The myFunc: %d\n", f('5'))
	fmt.Printf("The myFunc: %d\n", f('9'))

	animals := []string{"lion", "tiger", "bear"}
	result := searchItem(animals, "bear")
	fmt.Printf("Found %v in %v\n", result, animals)

	animals = []string{"lion", "tiger", "bear"}
	result = searchCaseIns(animals, "beAR")
	fmt.Printf("Found %v in %v\n", result, animals)

}
