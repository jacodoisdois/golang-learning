package main

import "fmt"

type car struct {
	name            string
	currentVelocity int
}

type ferrari struct {
	car
	turboOn bool
}

func main() {
	f := ferrari{}
	f.name = "F30"
	f.currentVelocity = 0
	f.turboOn = true

	fmt.Println(f)
}
