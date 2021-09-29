package main

import "fmt"

type money float64

type book struct {
	price float64
	title string
}

func vat(value float64) float64 {
	return value * 0.09
}

func (b *book) discount() {
	(*b).price = (*b).price * 0.9
}

func (m money) print() {
	fmt.Printf("%.2f\n", m)
}

func (m money) sprintf() string {
	return fmt.Sprintf("%.2f", m)
}

func main() {
	var m money = 64.567

	m.print()

	v := m.sprintf()

	fmt.Println(v)

	fmt.Println("The vat:", vat(float64(m)))

	b := book{price: 34.567, title: "Pscycology of money"}
	fmt.Println(b)
	b.discount()
	fmt.Printf(" Discount available: %v\n", b.price)

}
