package main

import "fmt"

type sportive interface {
	setTurboOn()
}

type ferrari struct {
	model        string
	turbo        bool
	currentSpeed int
}

func (f *ferrari) setTurboOn() {
	f.turbo = true
}
func main() {
	car := ferrari{"S30", false, 0}

	car.setTurboOn()
	var car2 sportive = &ferrari{"S40", false, 0}
	car2.setTurboOn()

	fmt.Println(car, car2)
}
