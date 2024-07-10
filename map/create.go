package main

import "fmt"

func main() {
	// 1. Declaring and initializing an empty map
	var emptyMap map[string]int
	fmt.Println("Initializing empty map", emptyMap)

	// 2. Creating a map using make()
	ages := make(map[string]int)
	fmt.Println("Map created with make():", ages)

	// 3. Creating and initializing a map using map literal
	population := map[string]int{
		"New York":    8419000,
		"Los Angeles": 3898000,
		"Chicago":     2746000,
	}
	fmt.Println("Map created with literal:", population)

	// 4. Creating a map with initial capacity
	scores := make(map[string]int, 5)
	fmt.Println("Map with initial capacity:", scores)

	// 5. Creating a map with non-string keys
	coordinates := map[int][]float64{
		1: {40.7128, -74.0060},  // New York
		2: {34.0522, -118.2437}, // Los Angeles
	}
	fmt.Println("Map with non-string keys:", coordinates)

	// 6. Creating a map of maps
	nested := map[string]map[string]int{
		"person1": {"age": 30, "salary": 50000},
		"person2": {"age": 25, "salary": 45000},
	}
	fmt.Println("Nested map:", nested)
}
