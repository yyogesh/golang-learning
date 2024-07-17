package game

import "fmt"

func DescribeRoom(room string) {
	switch room {
	case "entrance":
		fmt.Println("You are at the entrance of a dark cave.")
	case "hallway":
		fmt.Println("You are in a narrow hallway. It continues to the north.")
	case "library":
		fmt.Println("You are in a dusty library filled with old books.")
	case "chamber":
		fmt.Println("You are in a large chamber with high ceilings.")
	case "treasure room":
		fmt.Println("You are in the treasure room. You've found the treasure!")
	default:
		fmt.Println("You are in an unknown area.")
	}
}
