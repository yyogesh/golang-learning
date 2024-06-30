package main

import (
	"errors"
	"fmt"
	"struct_example/user"
)

func main() {
	println("Struct")

	var FirstName = "first"
	var LastName = "last"
	var Age = 20

	fmt.Println(errors.New("asdfasf"))

	fmt.Println(user.New("Ravi", "kumar", 10))

	var ajay = user.User{}
	user.DisplayStructsProperties(ajay)

	var abc = user.User{
		FirstName: "abc",
		LastName:  "xyz",
		Age:       20,
	}

	user.DisplayStructsProperties(abc)

	var abc1 = user.User{
		FirstName: "abc1",
		LastName:  "xyz",
	}

	user.DisplayStructsProperties(abc1)

	// var abc2 = user.User{
	// 	"xyz",
	// 	"abc2",
	// 	20,
	// }

	// user.DisplayStructsProperties(abc2)

	var abc3 = user.User{
		FirstName: FirstName,
		LastName:  LastName,
		Age:       Age,
	}

	user.DisplayStructsPointerProperties(&abc3)

	abc3.GetUserDetail()
	abc3.UpdateAge(21)
	abc3.GetUserDetail()

	fmt.Println("************************************")

	user.DisplayProperties(FirstName, LastName, Age)
}
