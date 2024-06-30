package main

import (
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		// return 0, errors.New("can't divide by zero")
		return 0, fmt.Errorf("can't divide by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(4, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result1, err := divide(4, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result1:", result1)
	}
}
