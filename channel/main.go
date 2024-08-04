package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

// Example 1
// func printStatments(ch1 chan int) {
// 	fmt.Println("statements")
// 	time.Sleep(4 * time.Second)
// 	ch1 <- 1 // writing on channel again
// }

// func main() {
// 	ch := make(chan int)

// 	// ch <- 10 // writing on channel

// 	// value := <-ch // reading from channel

// 	// channel communication between goroutines
// 	go printStatments(ch)

// 	value := <-ch // reading from channel

// 	fmt.Println(value)
// }

// Example 2
// func main() {
// 	x := 20
// 	y := 10

// 	sumCh := make(chan int)
// 	subCh := make(chan int)
// 	go calcuateSum(x, y, sumCh)
// 	go calcuateSub(x, y, subCh)
// 	sumResult, subResult := <-sumCh, <-subCh
// 	fmt.Printf("Total is: %d\n", sumResult+subResult)
// }

// func calcuateSum(x, y int, sumCh chan int) {
// 	sum := x + y
// 	sumCh <- sum
// }

// func calcuateSub(x, y int, subCh chan int) {
// 	sub := x - y
// 	subCh <- sub
// }

// Example 3 // close channel

// func main() {
// 	ch := make(chan int)
// 	go printFunc(ch)
// 	for v := range ch {
// 		fmt.Println("Received ", v)
// 	}

// 	for {
// 		v, ok := <-ch
// 		if ok == false {
// 			break
// 		}
// 		fmt.Println("Received ", v, ok)
// 	}

// }

// func printFunc(chnl chan int) {
// 	for i := 0; i < 10; i++ {
// 		chnl <- i
// 	}
// 	close(chnl)
// }

// Example 4 // uni directional

// Channel Direction Types
// Bidirectional Channel: Can send and receive values.
// Send-Only Channel: Can only send values.
// Receive-Only Channel: Can only receive values.
// func main() {
// 	// ch := make(chan int) Bidirectional Channel
// 	// ch <- 10 // Send value to channel
// 	// value := <-ch // Receive value from channel

// 	// chSend := make(chan<- int) // Send-Only Channel

// 	// chReceive := make(<-chan int) // Receive-Only Channel
// 	ch := make(chan int)
// 	go sendData(ch)
// 	receiveData(ch)
// }

// // A function that sends numbers to the channel
// func sendData(sendCh chan<- int) {
// 	for i := 0; i < 5; i++ {
// 		sendCh <- i
// 	}
// 	close(sendCh)
// }

// // A function that receives numbers from the channel
// func receiveData(recvCh <-chan int) {
// 	for value := range recvCh {
// 		fmt.Println(value)
// 	}
// }

// Example 5 Buffered Channel

// func main() {
// 	// ch := make(chan int, 2) // Create a buffered channel with a capacity of 2
// 	// ch <- 1                 // Send value 1
// 	// ch <- 2                 // Send value 2

// 	// fmt.Println(<-ch) // Receive and print value 1
// 	// fmt.Println(<-ch) // Receive and print value 2

// 	ch1 := make(chan int, 3) // Buffered channel with capacity 3
// 	go producer(ch1)
// 	consumer(ch1)

// }

// func producer(ch chan int) {
// 	for i := 0; i < 5; i++ {
// 		ch <- i
// 		fmt.Println("Sent:", i)
// 	}
// 	close(ch)
// }

// func consumer(ch chan int) {
// 	for val := range ch {
// 		fmt.Println("Received:", val)
// 		time.Sleep(time.Second) // Simulate processing time
// 	}
// }

// Example 6

func main() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)

}

func producer(ch chan int) {
	for i := 0; i < 100; i++ {
		result := work()
		ch <- result
	}
	close(ch)
}

func work() int {
	time.Sleep(time.Second)
	return rand.IntN(100)
}

func consumer(ch chan int) {
	for val := range ch {
		fmt.Println("Received:", val)
	}
}
