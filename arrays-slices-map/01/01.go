package main

import "fmt"

func main() {
	// array fixed spaces and with the same type
	var scores [3]float64

	scores[0], scores[1], scores[2] = 2.2, 3.93, 4.3

	fmt.Println(scores)

	total := 0.0

	for i := 0; i < len(scores); i++ {
		total += scores[i]
	}

	gradeAverage := total / float64(len(scores))

	fmt.Printf("Your grade average is: %.2f", gradeAverage)

}
