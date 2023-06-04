package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// Integers
	fmt.Println(1, 2, 3000)
	fmt.Println("Literal integer is:", reflect.TypeOf(32000))

	// without sign (just positives): uint8 uint16 uint32 uint64
	var nB byte = 255
	fmt.Println("Byte is:", reflect.TypeOf(nB))

	// with sign int8, int16, int32, int64
	nC := math.MaxInt64
	fmt.Println("The max value of int64 is:", nC)
	fmt.Println("The typeof nC is:", reflect.TypeOf(nC))

	var nD rune = 'a'
	fmt.Println("The typeof nD(rune) is:", reflect.TypeOf(nD))
	fmt.Println(nD)

	// Real numbers
	var nX float32 = 49.99
	fmt.Println("Type of x is:", reflect.TypeOf(nX))
	fmt.Println("Type of 49.99(literal) is:", reflect.TypeOf(49.99))

	// boolean
	bo := true
	fmt.Println("Type of x is:", reflect.TypeOf(bo))
	fmt.Println(!bo)

	s1 := "Test"
	fmt.Println(s1 + "!")
	fmt.Println("The length of s1 is:", len(s1))

	s2 := `T
	e
	s
	t
	e`

	fmt.Println(s2)
}
