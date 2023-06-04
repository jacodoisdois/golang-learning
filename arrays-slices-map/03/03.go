package main

import (
	"fmt"
	"reflect"
)

func main() {
	array := [3]int{1, 2, 3}
	slice := []int{1, 2, 3}

	fmt.Println(array, slice)
	fmt.Println(reflect.TypeOf(array), reflect.TypeOf(slice))

	array2 := [5]int{1, 3, 4, 5, 6}
	slice2 := array[1:3]

	fmt.Println(array2, slice2)
}
