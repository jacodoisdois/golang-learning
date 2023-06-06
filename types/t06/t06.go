package main

import "fmt"

type printable interface {
	toString() string
}

type person struct {
	name    string
	surname string
}

type product struct {
	name  string
	price float64
}

func (p person) toString() string {
	return p.name + " " + p.surname
}

func (p product) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.name, p.price)
}

func print(x printable) {
	fmt.Println(x.toString())
}

func main() {

	anything := product{"Pasta", 99}
	person := person{"John", "Mario"}
	print(person)
	print(anything)
}
