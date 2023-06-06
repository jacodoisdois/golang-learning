package main

import "fmt"

func printApproveds(approveds ...string) {
	fmt.Println("List of approveds")
	for i, approved := range approveds {
		fmt.Printf("%v# place: %v\n", i+1, approved)
	}
}

func main() {
	printApproveds("João", "Maria")
}
