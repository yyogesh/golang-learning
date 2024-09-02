package main

import (
	"context"
	"fmt"
	"time"
)

func main12() {
	// Create a base context with some order information
	orderCtx := context.WithValue(context.Background(), "orderId", "12345")

	// Create a cancellable context for the main chef
	chefCtx, chefCancel := context.WithCancel(orderCtx)
	defer chefCancel()

	// Start the salad and grill stations with the order context
	go saladStation(orderCtx)
	go grillStation(orderCtx)

	// Main chef oversees for 5 seconds, then leaves
	go mainChef(chefCtx)

	// Wait for 10 seconds total
	time.Sleep(2 * time.Second)
	fmt.Println("Restaurant closing!")
}

func mainChef(ctx context.Context) {
	fmt.Println("Main chef started overseeing")
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Main chef is leaving")
		return
	case <-ctx.Done():
		return
	}
}

func saladStation(ctx context.Context) {
	orderId := ctx.Value("orderId")
	fmt.Printf("Salad station started for order %v\n", orderId)
	for {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("Salad prepared!")
		case <-ctx.Done():
			fmt.Println("Salad station closing")
			return
		}
	}
}

func grillStation(ctx context.Context) {
	orderId := ctx.Value("orderId")
	fmt.Printf("Grill station started for order %v\n", orderId)
	for {
		select {
		case <-time.After(3 * time.Second):
			fmt.Println("Dish grilled!")
		case <-ctx.Done():
			fmt.Println("Grill station closing")
			return
		}
	}
}
