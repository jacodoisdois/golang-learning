package main

import "fmt"

func multiply(a, b int) int {
	return a * b
}

func exec(function func(int, int) int, p1, p2 int) int {
	return function(p1, p2)
}

func main() {
	fmt.Println(exec(multiply, 3, 5))
}
