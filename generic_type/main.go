package main

import "fmt"

func sum[T int | float64](a, b T) T {
	return a + b
}

func dispaly[T any](a, b T) T {
	fmt.Println("value of ", a, b)
	return a
}

// Generic function to swap the values of two variables
func Swap[T any](a, b T) (T, T) {
	return b, a
}

func main() {
	// <T>

	fmt.Println(sum(5, 5), sum(5.5, 6.2))

	dispaly(5, 5)
	dispaly("123", "asdf")
	dispaly(true, true)

	x, y := 1, 2
	fmt.Printf("Before Swap: x = %d, y = %d\n", x, y)
	x, y = Swap(x, y)
	fmt.Printf("After Swap: x = %d, y = %d\n", x, y)

	str1, str2 := "hello", "world"
	fmt.Printf("Before Swap: str1 = %s, str2 = %s\n", str1, str2)
	str1, str2 = Swap(str1, str2)
	fmt.Printf("After Swap: str1 = %s, str2 = %s\n", str1, str2)
}
