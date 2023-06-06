package main

import "fmt"

type product struct {
	productID int
	quantity  int
	price     float64
}
type order struct {
	customerID int
	products   []product
}

func (o order) totalValue() float64 {
	total := 0.0
	for _, product := range o.products {
		total += product.price * float64(product.quantity)
	}
	return total
}
func main() {
	order := order{
		customerID: 1,
		products: []product{
			{productID: 1, quantity: 2, price: 12.10},
			{2, 1, 23.49},
			{11, 100, 3.13},
		},
	}
	fmt.Printf("The total value of the order is $ %.2f", order.totalValue())
}
