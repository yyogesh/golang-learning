package main

import (
	"fmt"
	"time"
)

// select Statment channel
func main() {
	// select {
	// case <-ch1:
	// 	// Code to execute when ch1 is ready for receiving

	// case ch2 <- value:
	// 	// Code to execute when ch2 is ready for sending

	// default:
	// 	// Code to execute when none of the above cases are ready
	// }

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Message from ch1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Message from ch2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout")
	default:
		fmt.Println("No channels ready, doing other work")
		time.Sleep(500 * time.Millisecond) // Simulate doing other work
	}
}
