package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	score := 6.9
	finalNote := int(score)
	fmt.Println(finalNote)

	// int to string
	fmt.Println("Something", strconv.Itoa(40))

	// string to int

	num, er := strconv.Atoi("123")
	fmt.Println(num)
	fmt.Println(er)
}
