package main

import (
	"fmt"
	"time"
)

func main9() {
	// ch := make(chan int, 2) // Create a buffered channel with a capacity of 2

	// ch <- 1 // Send value 1
	// ch <- 2 // Send value 2

	// fmt.Println(<-ch) // Receive and print value 1
	// fmt.Println(<-ch) // Receive and print value 2

	ch1 := make(chan int, 3) // Buffered channel with capacity 3
	go producer(ch1)
	consumer(ch1)
}

func producer(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("Sent:", i)
	}
	close(ch)
}

func consumer(ch chan int) {
	for val := range ch {
		fmt.Println("Received:", val)
		time.Sleep(time.Second) // Simulate processing time
	}
}
