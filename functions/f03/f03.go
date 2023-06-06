package main

import "fmt"

func handle(p1, p2 int) (second int, first int) {
	first = p1
	second = p2
	return
}

func main() {
	a, b := handle(2, 3)

	fmt.Println(a, b)
}
