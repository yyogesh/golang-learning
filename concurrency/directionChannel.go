package main

import "fmt"

// Channel Direction Types
// Bidirectional Channel: Can send and receive values.
// Send-Only Channel: Can only send values.
// Receive-Only Channel: Can only receive values.
func main() {
	// ch := make(chan int) Bidirectional Channel
	// ch <- 10 // Send value to channel
	// value := <-ch // Receive value from channel

	// chSend := make(chan<- int) // Send-Only Channel
	// chSend <- 10

	// chReceive := make(<-chan int) // Receive-Only Channel
	// value := <-chReceive
	ch := make(chan int)
	go sendData(ch)
	receiveData(ch)
}

// A function that sends numbers to the channel
func sendData(sendCh chan<- int) {
	for i := 0; i < 5; i++ {
		sendCh <- i
	}
	close(sendCh)
}

// A function that receives numbers from the channel
func receiveData(recvCh <-chan int) {
	for value := range recvCh {
		fmt.Println(value)
	}
}
