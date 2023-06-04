package main

import "fmt"

func shopping(job1, job2 bool) (bool, bool, bool) {
	buyTv50 := job1 && job2
	buyTv32 := job1 != job2
	buyIcecream := job1 || job2

	return buyTv50, buyTv32, buyIcecream
}

func main() {
	tv50, tv32, icecream := shopping(true, true)

	fmt.Printf("Tv50: %v\nTv32: %v\nIcecream: %v", tv50, tv32, icecream)
}
