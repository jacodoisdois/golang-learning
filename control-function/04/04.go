package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
	fmt.Println("----------")

	for i := 0; i <= 20; i += 2 {
		fmt.Println(i)
	}
	fmt.Println("----------")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}

	for {
		fmt.Println("...")
		time.Sleep(time.Second)
	}
}
