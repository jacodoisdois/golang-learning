package main

import (
	"fmt"
	"strings"
)

type person struct {
	name    string
	surname string
}

func (p person) getFullName() string {
	return p.name + " " + p.surname
}

func (p *person) setFullName(fullName string) {
	pieces := strings.Split(fullName, " ")

	p.name = pieces[0]
	p.surname = pieces[1]
}

func main() {

	p1 := person{"John", "Silva"}

	fmt.Println(p1.getFullName())

	p1.setFullName("Mario Dias")

	fmt.Println(p1.getFullName())

}
