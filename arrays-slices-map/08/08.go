package main

import "fmt"

func main() {
	workersAndSalaries := map[string]float64{
		"John":  1323.23,
		"Maria": 99999.23,
		"Pedro": 12332.32,
	}

	workersAndSalaries["Rafael"] = 1350
	delete(workersAndSalaries, "Doesn't exist")

	for worker, salary := range workersAndSalaries {
		fmt.Println(worker, salary)
	}
}
