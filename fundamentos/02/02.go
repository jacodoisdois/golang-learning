package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2
	area := PI * math.Pow(raio, 2)

	fmt.Println(PI)
	fmt.Println("The circunference area is: ", area)

	const (
		a = 2
		b = 1
	)

	var (
		c = 3
		d = 4
	)

	e := c + d

	fmt.Println("The sum of C and D is", e)

	g, h := 2, false

	fmt.Println(g, h)
	fmt.Println(a, b)
}
