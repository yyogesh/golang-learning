package main

import "fmt"

func main() {
	// Initialize an empty map
	fruits := make(map[string]int)

	// 1. Adding new key-value pairs
	fruits["apple"] = 5
	fruits["banana"] = 8
	fruits["orange"] = 3
	fmt.Println("After adding:", fruits)

	// 2. Updating existing key-value pairs
	fruits["apple"] = 7
	fmt.Println("After updating apple:", fruits)

	// 3. Adding a new pair and updating in a single line
	fruits["grape"] = 4
	fruits["grape"]++
	fmt.Println("After adding and incrementing grape:", fruits)

	// 4. Using shorthand declaration to add/update
	watermelon := 2
	fruits["watermelon"] = watermelon
	fmt.Println("After adding watermelon:", fruits)

	// 5. Conditional update (increment if exists, otherwise set to 1)
	if value, exists := fruits["kiwi"]; exists {
		fruits["kiwi"] = value + 1
	} else {
		fruits["kiwi"] = 1
	}

	fmt.Println("After conditional update for kiwi:", fruits)

	// 6 Using map index expression in larger expressions
	fruits["total"] = fruits["apple"] + fruits["banana"] + fruits["orange"]
	fmt.Println("After calculating total:", fruits)

	// 7. Updating multiple values at once
	for _, fruit := range []string{"apple", "banana", "orange"} {
		fruits[fruit] *= 2
	}
	fmt.Println("After doubling some values:", fruits)

}
