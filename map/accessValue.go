package main

import "fmt"

func main() {
	// Initialize a map
	fruits := map[string]int{
		"apple":  5,
		"banana": 8,
		"orange": 3,
	}

	// 1. Basic value access
	appleCount := fruits["apple"]
	fmt.Println("Number of apples:", appleCount)

	// 2. Accessing a non-existent key
	grapeCount := fruits["grape"]
	fmt.Println("Number of grapes:", grapeCount) // Prints 0 (zero value for int)

	// 3. Checking for key existence using comma ok idiom
	orangeCount, exists := fruits["orange"]
	if exists {
		fmt.Println("Number of oranges:", orangeCount)
	} else {
		fmt.Println("No oranges in the map")
	}

	// 4. Looping through the map
	fmt.Println("All fruits:")
	for fruit, count := range fruits {
		fmt.Printf("%s: %d\n", fruit, count)
	}

	// 5. Checking length of the map
	fmt.Println("Number of fruit types:", len(fruits))

	// 6. Accessing nested map values
	people := map[string]map[string]string{
		"Alice": {"phone": "123-456-7890", "email": "alice@example.com"},
		"Bob":   {"phone": "987-654-3210", "email": "bob@example.com"},
	}

	if personInfo, ok := people["Alice"]; ok {
		if phone, ok := personInfo["phone"]; ok {
			fmt.Println("Alice's phone number:", phone)
		}
	}

}
