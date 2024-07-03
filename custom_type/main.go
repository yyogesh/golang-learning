package main

import "fmt"

type MyInt int

func (a MyInt) isEven() bool {
	return a%2 == 0 // Check if the number is even
}

func main() {
	var a MyInt = 5
	fmt.Println(a) // Output: 5

	fmt.Println(a.isEven()) // Output: true

	// var b = 10;

	var b MyInt = MyInt(5) //
	fmt.Println(b)         // Output: 5

	var c int = int(a)
	fmt.Println(c) // Output: 5
}
