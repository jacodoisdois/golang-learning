package main

import "fmt"

func main() {

	numbers := [...]int{1, 2, 3, 4, 5, 6}

	for i, number := range numbers {
		fmt.Printf("%v) %v\n", i, number)
	}

	for _, number := range numbers {
		fmt.Printf("%v\n", number)
	}

}
