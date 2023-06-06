package main

import "fmt"

type product struct {
	name     string
	price    float64
	discount float64
}

func (p product) priceWithDiscount() float64 {
	return p.price * (1 - p.discount)
}

func main() {

	p1 := product{
		name:     "Pencil",
		price:    1,
		discount: 0.25,
	}

	fmt.Println(p1.priceWithDiscount())
}
