package main

import (
	"fmt"
)

func main() {
	// Original map
	original := map[string]int{
		"apple":  5,
		"banana": 8,
		"orange": 3,
	}

	// 1. Shallow copy using assignment
	shallowCopy := original
	fmt.Println("Original:", original)
	fmt.Println("Shallow copy:", shallowCopy)

	// Modifying shallow copy
	shallowCopy["apple"] = 10
	fmt.Println("\nAfter modifying shallow copy:")
	fmt.Println("Original:", original)
	fmt.Println("Shallow copy:", shallowCopy)

	// 2. Deep copy using a loop
	deepCopy := make(map[string]int)
	for k, v := range original {
		deepCopy[k] = v
	}
	// Modifying deep copy
	deepCopy["banana"] = 12
	fmt.Println("\nAfter modifying deep copy:")
	fmt.Println("Original:", original)
	fmt.Println("Deep copy:", deepCopy)
}
