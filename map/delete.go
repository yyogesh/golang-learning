package main

import "fmt"

func main() {
	// Initialize a map
	fruits := map[string]int{
		"apple":  5,
		"banana": 8,
		"orange": 3,
		"grape":  4,
		"mango":  2,
	}

	fmt.Println("Original map:", fruits)

	// 1. Basic deletion
	delete(fruits, "orange")
	fmt.Println("After deleting orange:", fruits)

	// 2. Deleting a non-existent key (safe operation)
	delete(fruits, "pear")
	fmt.Println("After trying to delete non-existent pear:", fruits)

	// 3. Deleting all elements (clearing the map)
	for key := range fruits {
		delete(fruits, key)
	}
	fmt.Println("After deleting all elements:", fruits)

	fruits = make(map[string]int)
	fmt.Println("After recreating the map:", fruits)

	// 4. Deleting from a nil map (safe operation)
	var nilMap map[string]int
	delete(nilMap, "something")
	fmt.Println("After deleting from nil map:", nilMap)

}
