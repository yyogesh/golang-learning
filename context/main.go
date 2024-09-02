package main

import "context"

func main1() {
	// Create a new context
	ctx := context.Background()

	// Add a value to the context
	ctx = context.WithValue(ctx, "userId", 123)

	printValue(ctx)
}

func printValue(ctx context.Context) {
	userId := ctx.Value("userId")
	if userId != nil {
		println("User ID:", userId.(int))
	} else {
		println("No user ID found")
	}
	// Output: User ID: 123
}
