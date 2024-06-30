package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	println("Struct")

	var FirstName = "first"
	var LastName = "last"
	var Age = 20

	var ajay = User{}
	DisplayStructsProperties(ajay)

	var abc = User{
		FirstName: "abc",
		LastName:  "xyz",
		Age:       20,
	}

	DisplayStructsProperties(abc)

	var abc1 = User{
		FirstName: "abc1",
		LastName:  "xyz",
	}

	DisplayStructsProperties(abc1)

	var abc2 = User{
		"xyz",
		"abc2",
		20,
	}

	DisplayStructsProperties(abc2)

	var abc3 = User{
		FirstName,
		LastName,
		Age,
	}

	DisplayStructsPointerProperties(&abc3)

	DisplayProperties(FirstName, LastName, Age)
}

func DisplayProperties(FirstName string, LastName string, Age int) {
	fmt.Printf("Hello, my name is %s %s. I am %d years old.\n", FirstName, LastName, Age)
	fmt.Println("************************************")
}

func DisplayStructsProperties(u User) {
	fmt.Printf("Hello, my name is %s %s. I am %d years old.\n", u.FirstName, u.LastName, u.Age)
	fmt.Println("************************************")
}

func DisplayStructsPointerProperties(u *User) {
	fmt.Printf("Hello, my name is %s %s. I am %d years old.\n", (*u).FirstName, u.LastName, u.Age)
	fmt.Printf("Hello, my name is %s %s. I am %d years old.\n", u.FirstName, u.LastName, u.Age)
	fmt.Println("************************************")
}
