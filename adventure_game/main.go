package main

import (
	"adventure_game/game"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	currentRoom := "entrance"
	gameRunning := true
	inventory := []string{}

	fmt.Println("Welcome to the Adventure Game!")
	fmt.Println("Type 'help' to see the list of commands.")

	for gameRunning {
		// https://github.com/yyogesh/golang-learning
		// Print current room description
		currentRoomDescription(currentRoom)
		fmt.Print("What do you want to do? ")

		// Read user input
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		parts := strings.SplitN(input, " ", 2)
		// go north

		switch parts[0] {
		case "help":
			fmt.Println("Available commands: help, go <direction>, look, take <item>, inventory, exit")
		case "go":
			if len(parts) < 2 {
				fmt.Println("Go where?")
			} else {
				newRoom := game.Move(currentRoom, parts[1])
				if newRoom == "" {
					fmt.Println("You can't go that way.")
				} else {
					currentRoom = newRoom
					if currentRoom == "treasure room" {
						fmt.Println("Congratulations! You've found the treasure and won the game!")
						gameRunning = false
					}
				}
			}
		case "look":
			currentRoomDescription(currentRoom)
		case "take":
			if len(parts) > 1 {
				item := parts[1]
				if game.TakeItem(currentRoom, &inventory, item) {
					fmt.Printf("You took the %s.\n", item)
				} else {
					fmt.Println("There is no such item here.")
				}
			} else {
				fmt.Println("Take what?")
			}
		case "inventory":
			if len(inventory) > 0 {
				fmt.Printf("You have: %s\n", strings.Join(inventory, ", "))
			} else {
				fmt.Println("Your inventory is empty.")
			}
		case "exit":
			fmt.Println("Thanks for playing the Adventure Game!")
			gameRunning = false
		default:
			fmt.Println("Type 'help' for a list of commands.")
		}
	}
}

func currentRoomDescription(currentRoom string) {
	game.DescribeRoom(currentRoom)
	roomItems := game.GetRoomItems(currentRoom)
	if len(roomItems) > 0 {
		fmt.Printf("You see: %s\n", strings.Join(roomItems, ", "))
	}
}
