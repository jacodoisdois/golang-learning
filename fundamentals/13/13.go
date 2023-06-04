package main

import "fmt"

func main() {
	i := 1

	var pointer *int = nil

	pointer = &i

	fmt.Println(pointer)

	*pointer++
	fmt.Println(*pointer)

	pointer = &i
	fmt.Println(pointer, *pointer, i, &i)

	i = 3
	fmt.Println(pointer, *pointer, i, &i)
}
