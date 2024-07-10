package main

import (
	"fmt"
	"sort"
)

func main() {
	// Initialize a map
	population := map[string]int{
		"New York":    8419000,
		"Los Angeles": 3898000,
		"Chicago":     2716000,
		"Houston":     2313000,
		"Phoenix":     1680000,
	}

	// 1. Basic iteration using for range
	fmt.Println("1. Basic iteration:")
	for city, pop := range population {
		fmt.Printf("%s: %d\n", city, pop)
	}

	// 2.  Iterating in a specific order (sorted keys)
	fmt.Println("\n4. Iterating in sorted order:")

	var cities []string
	for city := range population {
		cities = append(cities, city)
	}

	sort.Strings(cities)

	for _, city := range cities {
		fmt.Printf("%s: %d\n", city, population[city])
	}

	// 3. Using a slice to maintain insertion order
	fmt.Println("\n5. Maintaining insertion order:")
	orderedCities := []string{"New York", "Los Angeles", "Chicago", "Houston", "Phoenix"}

	for _, city := range orderedCities {
		fmt.Printf("%s: %d\n", city, population[city])
	}

}
