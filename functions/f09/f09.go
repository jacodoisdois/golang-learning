package main

import "fmt"

func getApprovedValue(number int) int {
	defer fmt.Println("End!")
	if number > 5000 {
		fmt.Println("High value")
	} else {
		fmt.Println("Low value")
	}
	return number
}

func main() {
	fmt.Println(getApprovedValue(5000))
}
