package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello, world!")
}

func main1() {
	fmt.Println("Main function")
	go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("Ending main function")
}
