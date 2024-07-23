package main

import "fmt"

func main() {
	// Assigning a function to a variable
	greet := func(name string) {
		fmt.Printf("Hello, %s!\n", name)
	}

	// Calling the function through the variable
	greet("Alice")

	// Functions as parameters:

	// Define some operations
	add := func(a, b int) int { return a + b }
	multiply := func(a, b int) int { return a * b }

	// Use the functions as arguments
	result1 := applyOperation(5, 3, add)
	result2 := applyOperation(5, 3, multiply)

	fmt.Printf("5 + 3 = %d\n", result1)
	fmt.Printf("5 * 3 = %d\n", result2)

	// Functions as return types:
	englishGreeter := getGreeter("Hello")
	spanishGreeter := getGreeter("Hola")

	englishGreeter("Bob")
	spanishGreeter("Carlos")

	// Custom function types:

	// Create functions that match the MathFunc type
	add2 := MathFunc(func(a, b int) int { return a + b })
	subtract2 := MathFunc(func(a, b int) int { return a - b })

	fmt.Printf("10 + 5 = %d\n", operate(10, 5, add2))
	fmt.Printf("10 - 5 = %d\n", operate(10, 5, subtract2))

	fmt.Printf("10 + 5 = %d\n", operate(10, 5, func(a, b int) int { return a + b }))

	// Variadic Function:
	fmt.Println(sum(1, 2))          // 3
	fmt.Println(sum(1, 2, 3, 4, 5)) // 15
	fmt.Println(sum())              // 0
}

func sum(num ...int) int {
	total := 0

	for _, n := range num {
		total += n
	}

	return total
}

// Define a custom function type
type MathFunc func(int, int) int

// Function that uses the custom function type
func operate(x, y int, operation MathFunc) int {
	return operation(x, y)
}

func applyOperation(a, b int, fn func(int, int) int) int {
	return fn(a, b)
}

// Function that returns another function
func getGreeter(greeting string) func(string) {
	return func(name string) {
		fmt.Printf("%s, %s!\n", greeting, name)
	}
}
