package main

import "fmt"

func f4(p1, p2 string) string {
	return "F4"
}

func f5() (string, string) {
	return "V1", "V2"
}

func main() {
	fmt.Println(f4("a", "2"))
	r1, _ := f5()

	fmt.Println(r1)
}
