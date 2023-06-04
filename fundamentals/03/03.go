package main

import "fmt"

func main() {
	fmt.Print("Same ")
	fmt.Println("line.")

	fmt.Println("New line")

	x := 3.141516
	xs := fmt.Sprint(x)

	fmt.Println("The x value is: ", xs)
	fmt.Println("The x value is: ", x)

	fmt.Printf("The x value is: %f", x)
}
