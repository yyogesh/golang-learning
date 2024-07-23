package main

import (
	"fmt"
	"oop_example/circle"
	"oop_example/rectangle"
	"oop_example/shape"
)

func PrintShapeInfo(s shape.Shape) {
	fmt.Printf("Area: %0.2f\n", s.Area())
	fmt.Printf("Perimeter: %0.2f\n", s.Perimeter())
}

func main() {
	rect := rectangle.New(100, 50)

	fmt.Println("Area:", rect.Area())

	bigRect := rectangle.New(100, 50)

	fmt.Println("Area of bigRect:", bigRect.Area())

	bigColorRect := rectangle.NewColoredRectangle(100, 50, "red")

	fmt.Println("Area of bigColorRect:", bigColorRect.Area())

	c := circle.Circle{Radius: 7}

	// Polymorphism in action
	PrintShapeInfo(c)
	PrintShapeInfo(bigRect)

}
