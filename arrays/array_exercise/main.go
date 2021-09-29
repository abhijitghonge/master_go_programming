package main

import "fmt"

func main() {
	var cities = [2]string{"Mumbai", "Pune"}

	fmt.Printf("the cities:%#v\n", cities)

	grades := [3]float64{2: 23.4, 1: 45.6}

	fmt.Printf("grades: %v\n", grades)

	taskDone := [...]bool{}

	fmt.Println("tasks done: ", taskDone)

	for i := 0; i < len(cities); i++ {
		fmt.Printf("%v,", cities[i])
	}

	fmt.Println()

	for index, city := range cities {
		fmt.Printf("index:%v, city:%v\n", index, city)
	}

	nums := [...]int{30, -1, -6, 90, -6}

	fmt.Printf("Even numbers:")
	for _, num := range nums {
		if num > 0 && num%2 == 0 {
			fmt.Printf("%v,", num)
		}
	}
}
