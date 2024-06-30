package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Seed the random number generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Generate a random number between 1 and 100
	target := r.Intn(100) + 1
	// fmt.Println(target)

	describe()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter your guess: ")
		// Read user input

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("An error occurred while reading input. Please try again.", err)
			continue
		}

		input = strings.TrimSpace(input)

		// Convert the input to an integer
		// Itoa integer to string
		guess, err := strconv.Atoi(input)

		if err != nil || guess < 0 || guess > 100 {
			fmt.Println("Invalid input. Please enter a number between 1 and 100.")
			continue
		}

		if guess < target {
			fmt.Println("Too low!")
		} else if guess > target {
			fmt.Println("Too high!")
		} else {
			fmt.Println("Congratulations! You guessed the correct number.")
			break
		}
	}
	fmt.Println("Thanks for playing the Number Guessing Game!")
}

func describe() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I have selected a random number between 1 and 100.")
	fmt.Println("Can you guess what it is?")
}
