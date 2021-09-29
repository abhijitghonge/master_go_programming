package main

import "fmt"

func main() {
	type gradeInfo struct {
		grade  int
		course string
	}
	type person struct {
		name   string
		age    int
		colors []string
		grades gradeInfo
	}

	me := person{name: "Abhijit", age: 40, colors: []string{
		"Red", "Blue", "Black"},
		grades: gradeInfo{
			grade:  98,
			course: "Golang",
		},
	}
	you := person{name: "Marius", age: 50, colors: []string{
		"black",
		"pink",
	},
		grades: gradeInfo{
			grade:  99,
			course: "Golang",
		},
	}

	fmt.Printf("%v\n", me)
	fmt.Printf("%+v\n", you)

	you.name = "Andrei"

	me.colors = append(me.colors, "Olive")
	fmt.Printf("%v\n", me.colors)

	for k, v := range me.colors {
		fmt.Printf("index:%v, color:%v\n", k, v)
	}

}
