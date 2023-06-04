package main

import "fmt"

func scoreConcept(n float64) string {
	var score = int(n)

	switch score {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "F"
	default:
		return "Invalid score provided"
	}

}

func main() {
	fmt.Println(scoreConcept(6.9))
	fmt.Println(scoreConcept(7))
	fmt.Println(scoreConcept(2))
	fmt.Println(scoreConcept(10))
	fmt.Println(scoreConcept(4))
	fmt.Println(scoreConcept(-1))
}
