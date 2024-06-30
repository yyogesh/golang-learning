package user

import "fmt"

type User struct {
	FirstName string
	LastName  string
	Age       int
}

func (u User) GetUserDetail() {
	fmt.Printf("Hello, my name is %s %s. I am %d years old.\n", u.FirstName, u.LastName, u.Age)
}

func (u *User) UpdateAge(newAge int) {
	u.Age = newAge
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

func New(firstName, lastName string, age int) User {
	return User{FirstName: firstName, LastName: lastName, Age: age}
}
