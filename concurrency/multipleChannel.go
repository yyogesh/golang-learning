package main

import "fmt"

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func main2() {
	number := 123 // 1, 4 , 9 ==> 14
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)

	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares, cubes, squares+cubes)
	fmt.Println("Ending main function")
}
