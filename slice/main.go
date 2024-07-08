package main

import (
	"fmt"
	"reflect"
)

func main() {
	number := [5]int{1, 2, 3, 4, 5}
	fmt.Println(reflect.ValueOf(number).Kind())

	// Create a slice using a slice literal
	number1 := []int{1, 2, 3, 4, 5}
	fmt.Println(reflect.ValueOf(number1).Kind())

	arr := [5]int{1, 2, 3, 4, 5}
	// Create a slice from the array
	slice := arr[1:4]
	fmt.Println(reflect.ValueOf(slice).Kind())
	fmt.Println("Slice from array:", slice)
}
