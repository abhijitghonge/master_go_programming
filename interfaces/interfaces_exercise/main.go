package main

import "fmt"

type vehicle interface {
	License() string
	Name() string
}

type car struct {
	licenceNo string
	brand     string
}

func (c car) License() string {
	return c.licenceNo
}

func (c car) Name() string {
	return c.brand
}

func main() {

	var audi vehicle = car{licenceNo: "MHOG345R", brand: "Audi"}

	fmt.Println("Car license:", audi.License())
	fmt.Println("Car brand:", audi.Name())

	var empty interface{}

	empty = 64

	fmt.Printf("%T\n", empty)

	empty = 65.78
	fmt.Printf("%T\n", empty)

	empty = []int{1, 2, 3, 4}
	fmt.Printf("%T\n", empty)

	empty = append(empty.([]int), 4, 5)
	fmt.Printf("%v\n", empty)

}
