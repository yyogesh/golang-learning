package main

import (
	"encoding/json"
	"fmt"
)

// Base struct
type Animal struct {
	Name string `json:"dog_name"`
	Age  int    `json:"age"`
}

// Derived struct (inheritance through embedding)
type Dog struct {
	Animal
	Breed string
}

func main() {
	dog1 := Dog{
		Animal: Animal{Name: "Buddy", Age: 5},
		Breed:  "Labrador",
	}

	dog2 := Dog{
		Animal: Animal{Name: "Buddy", Age: 5},
		Breed:  "Labrador",
	}

	dog3 := Dog{
		Animal: Animal{Name: "Buddy", Age: 5},
		Breed:  "Golden Retriever",
	}

	if dog1 == dog2 {
		fmt.Println("Dog1 and Dog2 are same.")
	} else {
		fmt.Println("Dog1 and Dog2 are different.")
	}

	if dog2 == dog3 {
		fmt.Println("Dog2 and Dog3 are same.")
	} else {
		fmt.Println("Dog2 and Dog3 are different.")
	}

	dogJson, err := json.Marshal(dog1)
	if err != nil {
		fmt.Println("Error marshaling Dog1:", err)
	} else {
		fmt.Println("Dog1 as JSON:", string(dogJson))
	}

	var jsonString = `{"Name":"Buddy","Age":5,"Breed":"Labrador"}`
	var dog4 Dog
	err1 := json.Unmarshal([]byte(jsonString), &dog4)

	if err1 != nil {
		fmt.Println("Error unmarshaling JSON:", err1)
	} else {
		fmt.Println("Dog4:", dog4)
	}

	fmt.Println("Unmarshaled dog4:", dog4)

	dog6 := new(Dog)
	fmt.Println("dog 6 :", dog6)

	dog4.SetDogName("Buddy").SetDogAge(6)
}

func (d *Dog) SetDogName(name string) *Dog {
	d.Name = name
	return d
}

func (d *Dog) SetDogAge(age int) *Dog {
	d.Age = age
	return d
}
