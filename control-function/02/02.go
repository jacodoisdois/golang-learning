package main

import "fmt"

func scoreConcept(number float64) string {
	if number >= 9 && number <= 10 {
		return "A"
	} else if number >= 8 && number < 9 {
		return "B"
	} else if number >= 5 && number < 8 {
		return "C"
	} else if number >= 3 && number < 5 {
		return "D"
	}

	return "F"
}

func main() {
	fmt.Println(scoreConcept(2.4))
	fmt.Println(scoreConcept(6))
	fmt.Println(scoreConcept(9))
	fmt.Println(scoreConcept(7))
}
