package main

import "fmt"

func main() {
	var a int = 42 // Declare an integer variable 'a'
	var p *int     // Declare a pointer to an integer

	p = &a // Assign the address of 'a' to the pointer 'p'

	fmt.Println("Value of a:", a)             // Print the value of 'a'
	fmt.Println("Address of a:", &a)          // Print the address of 'a'
	fmt.Println("Value of p:", p)             // Print the value of 'p' (address of 'a')
	fmt.Println("Value pointed to by p:", *p) // Print the value of 'p

	*p = 21 // Change the value of 'a' through the pointer 'p'

	fmt.Println("Value of a:", a) // Print the value of 'a'

	var x int = 5
	var y int = 10

	fmt.Printf("Before swap, value of a : %d\n", x)
	fmt.Printf("Before swap, value of b : %d\n", y)

	swap(&x, &y)

	fmt.Printf("After swap, value of a : %d\n", x)
	fmt.Printf("After swap, value of b : %d\n", y)
}

func swap(x, y *int) {
	var temp = *x
	*x = *y
	*y = temp

	fmt.Println("After swap", *x, *y)
}
