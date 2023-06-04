package main

import (
	"fmt"
	"time"
)

func byType(i interface{}) string {
	switch i.(type) {
	case int:
		return "Integer"
	case float32, float64:
		return "Real"
	case string:
		return "String"
	case func():
		return "function"
	default:
		return "Invalid data type"
	}
}
func main() {
	fmt.Println(byType("Teste"))
	fmt.Println(byType(3.2))
	fmt.Println(byType(3))
	fmt.Println(byType(func() {}))
	fmt.Println(byType(time.Now()))
}
