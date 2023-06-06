package main

import "fmt"

type score float64

func (s score) between(start, end float64) bool {
	return float64(s) >= start && float64(s) <= end
}

func scoreForConcept(s score) string {
	if s.between(9, 10) {
		return "A"
	} else if s.between(7, 8) {
		return "B"
	}

	return "F"
}
func main() {
	fmt.Println(scoreForConcept(2))
}
