package main

import "fmt"

func main() {
	approveds := make(map[int]string)

	approveds[123123123] = "Maria"
	approveds[3213] = "John"
	approveds[3211] = "Ana"

	fmt.Println(approveds)

	for cpf, name := range approveds {
		fmt.Println(cpf, name)
	}

	delete(approveds, 3213)
	fmt.Println(approveds)
}
