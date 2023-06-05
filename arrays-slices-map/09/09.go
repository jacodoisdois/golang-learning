package main

import "fmt"

func main() {
	workersByLetter := map[string]map[string]float64{
		"G": {
			"Gabriel": 1234.23,
		},
		"P": {
			"Peter": 123123.23,
		},
	}

	for letter, workers := range workersByLetter {
		fmt.Println(letter, workers)
	}

	delete(workersByLetter, "G")

	fmt.Println(workersByLetter)
}
