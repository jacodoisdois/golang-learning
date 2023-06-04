package main

import "fmt"

func printResult(score float64) {
	if score >= 7 {
		fmt.Println("Approved! Score:", score)
	} else {
		fmt.Println("Disapproved! Score:", score)
	}
}

func main() {
	printResult(2)
	printResult(8)
}
