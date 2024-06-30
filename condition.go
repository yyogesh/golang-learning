package main

import "fmt"

func main1() {
	num := 6

	if num%2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}

	if num%2 == 0 {
		fmt.Println("The number is even")
		return
	}

	fmt.Println("The number is odd")

	age := 10
	ticketPrice := 0
	if age < 5 {
		ticketPrice = 0
	} else if age >= 5 && age <= 22 {
		ticketPrice = 10
	} else {
		ticketPrice = 15
	}
	fmt.Printf("Ticket price is $%d", ticketPrice)
}
