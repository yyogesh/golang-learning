package main

import (
	"context"
	"fmt"
	"time"
)

func main1() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // 5 seconds timeout

	defer cancel()

	go cookRice(ctx)
	go fryVegetables(ctx)

	<-ctx.Done()
	fmt.Println("Time's up! Stop cooking.")
}

func cookRice(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopping rice cooking")
			return
		default:
			// Continue cooking rice
		}
	}
}

func fryVegetables(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopping vegetable frying")
			return
		default:
			// Continue frying vegetables
		}
	}
}
