package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Sum =>", a+b)
	fmt.Println("Subtraction =>", a-b)
	fmt.Println("Division =>", a/b)
	fmt.Println("Multiplication =>", a*b)
	fmt.Println("Module =>", a%b)

	// bitwise
	fmt.Println("E => ", a&b)
	fmt.Println("Or => ", a|b)
	fmt.Println("Xor => ", a^b)

	c := 3.0
	d := 2.0

	fmt.Println("Higher => ", math.Max(float64(c), float64(d)))
	fmt.Println("Minor => ", math.Min(c, d))
	fmt.Println("Exponenciation => ", math.Pow(c, d))
}
