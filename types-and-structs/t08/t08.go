package main

import "fmt"

type Sporty interface {
	TurboOn()
}

type Luxurious interface {
	ParallelPark()
}

type LuxurySport interface {
	Sporty
	Luxurious
}

type Bmw7 struct{}

func (b Bmw7) TurboOn() {
	fmt.Println("Engaging turbo... Fasten your seatbelts!")
}

func (b Bmw7) ParallelPark() {
	fmt.Println("Initiating parallel parking mode... Please stand clear!")
}

func main() {
	var luxuryVehicle LuxurySport = Bmw7{}
	luxuryVehicle.TurboOn()
	luxuryVehicle.ParallelPark()
}
