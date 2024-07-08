package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 10
	fmt.Println("Value of a ", a)
	var b [3]int
	fmt.Println("Value of b ", b)

	var c [3]int
	c[0] = 10
	c[1] = 20
	c[2] = 30
	fmt.Println("Value of c is: ", c)

	d := [3]int{10, 20}
	fmt.Println("Value of d is: ", d)

	e := [...]int{10, 20}
	fmt.Println("Value of d is: ", e)

	fmt.Println(reflect.ValueOf(e).Kind())

	fmt.Println("--------------------------------")

	f := [...]float64{67.7, 89.9, 21, 78}
	for i := 0; i < len(f); i++ {
		fmt.Printf("Value of h at index %d is: %.2f\n", i, f[i])
	}

	fmt.Println("--------------------------------")
	for index, item := range f {
		fmt.Printf("Value of f at index %d is: %.2f\n", index, item)
	}

	fmt.Println("--------------------------------")
	g := [...]string{"red", "green", "blue"}
	h := g
	h[0] = "white"
	fmt.Println("--------------------------------")
	fmt.Println("Value of g before calling function is: ", g, h)

	fmt.Println("--------------------------------")

	fmt.Println("Value of g before calling function is: ", g)
	changeColor(g)
	fmt.Println("Value of g after calling function is: ", g)

	i := [...]string{"red", "green", "blue"}
	if itemExists(i, "yellow") {
		fmt.Println("Item exists in array")
	} else {
		fmt.Println("Item does not exist in array")
	}

	fmt.Println("--------------------------------")

	j := [...]string{"red", "green", "blue", "yellow", "black", "orange", "cyan", "white"}
	fmt.Println(j)
	fmt.Println(j[:2])
	fmt.Println(j[2:5])
	fmt.Println(j[2:])
	fmt.Println(j[0:2])
	fmt.Println(j[len(j)-1])
	fmt.Println(j[0 : len(j)-1])
	fmt.Println(j[:len(j)/2])
	fmt.Println(j[:])
	fmt.Println(j[0:])
	fmt.Println(j[0:len(j)])

	fmt.Println("--------------------------------")
	var matrix = [2][2]int{{1, 2}, {3, 4}}
	fmt.Println("Matrix:", matrix)
	for _, v1 := range matrix {
		for _, v2 := range v1 {
			fmt.Print(v2, " ")
		}
		fmt.Println()
	}

	fmt.Println("Sum of the matrix ", sumOfArray(matrix))

}

func sumOfArray(arr [2][2]int) int {
	sum := 0
	for _, v1 := range arr {
		for _, v2 := range v1 {
			sum += v2
		}
	}
	return sum
}

func itemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array {
		fmt.Println("Input is not an array")
		panic("Invalid data type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}
	return false
}

func changeColor(color [3]string) {
	color[0] = "yellow"
	fmt.Println("Value of color inside function is: ", color)
}
