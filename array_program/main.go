package main

import (
	findminandmax "array_program/findMinAndMax"
	reversearray "array_program/reverseArray"
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	max, min := findminandmax.FindMinAndMax(arr)
	fmt.Printf("Max: %d, Min: %d\n", max, min)

	arr1 := []int{1, 2, 3, 4, 5}
	fmt.Println("Original Array:", arr1)
	fmt.Println("Reversed Array:", reversearray.ReverseArray(arr1))

}
