package main

import "fmt"

func printMessage(message string) {
	fmt.Println(message)
}

func sum(a, b int) int {
	return a + b
}

func rectProps(l, b float64) (float64, float64) {
	var area = l * b
	var perimeter = 2 * (l + b)
	return area, perimeter
}

// Named return values
func rectNameProps(l, b float64) (area, perimeter float64) {
	area = l * b
	perimeter = 2 * (l + b)
	return
}

const name = "demo"

func init() {
	fmt.Println("Initializing")
}

func init() {
	fmt.Println("Initializing 1")
}

var (
	sum1 = func(a, b int) int {
		return a + b
	}
)

func displayMessage1(message string) func(string) string {
	return func(msg string) string {
		return "Hello " + message + " " + msg
	}
}

func calcuateArea(l float64) func(b float64) float64 {
	return func(b float64) float64 {
		return l * b
	}
}

func main2() {
	fmt.Println("Starting")
	printMessage("This is printing functionality")
	fmt.Println(sum(1, 2))
	area, perimeter := rectProps(10.8, 5.6)
	fmt.Println(area, perimeter)

	area1, _ := rectProps(4, 5)
	fmt.Println(area1)

	area2, perimeter1 := rectNameProps(10.8, 5.6)
	fmt.Println(area2, perimeter1)

	// IIFE
	func() {
		fmt.Println("Anonymous function")
	}()

	var displayMessage = func() {
		fmt.Println("Anonymous function")
	}

	displayMessage()

	fmt.Println(sum1(2, 4))

	func(x int) {
		fmt.Println("Anonymous function")
		fmt.Println(x * x)
	}(2)

	fmt.Println("asdf", func(x int) int {
		fmt.Println("Anonymous function")
		return x * x
	}(2))

	// resultMessage := displayMessage1("How are you?")
	// fmt.Println(resultMessage())

	fmt.Println(displayMessage1("How are you?")("what is up?"))

	fmt.Println(calcuateArea(5)(2))
}
