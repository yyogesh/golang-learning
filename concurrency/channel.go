package main

import (
	"fmt"
	"time"
)

func hello1(done chan bool) {
	fmt.Println("Hello, world!")
	time.Sleep(4 * time.Second)
	done <- true // write to channel
}

func main3() {
	done := make(chan bool)
	// done <- true // write to channel
	// isDone := <- done // reading from channel
	fmt.Println("Main function")
	go hello1(done)
	<-done
	// time.Sleep(1 * time.Second)
	fmt.Println("Ending main function")
}
