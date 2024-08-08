package main

import (
	"fmt"
	"sync"
	"time"
)

func main1() {
	var schoolTrip sync.WaitGroup
	fmt.Println("Planning the school trip...")

	// Add 3 tasks to the WaitGroup
	schoolTrip.Add(3)

	// Arjun's task
	go func() {
		time.Sleep(2 * time.Second) // Simulate time taken to get snacks
		fmt.Println("Arjun: Snacks are ready!")
		schoolTrip.Done()
	}()

	// Priya's task
	go func() {
		time.Sleep(3 * time.Second) // Simulate time taken to book transportation
		fmt.Println("Priya: Bus is booked!")
		schoolTrip.Done()
	}()

	// Rahul's task
	go func() {
		time.Sleep(1 * time.Second) // Simulate time taken to plan activities
		fmt.Println("Rahul: Activities are planned!")
		schoolTrip.Done()
	}()

	// Wait for all tasks to complete
	schoolTrip.Wait()
	fmt.Println("Great! We're ready for the school trip!")
}
