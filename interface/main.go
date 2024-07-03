package main

import "fmt"

type salaryCalculator interface {
	calculateSalary() float64
	report()
}

type PermanentEmployee struct {
	id          int
	basicSalary float64
	commission  float64
}

type ContractEmployee struct {
	id          int
	basicSalary float64
}

func (p PermanentEmployee) calculateSalary() float64 {
	return p.basicSalary + (p.commission/100)*p.basicSalary
}

func (c ContractEmployee) calculateSalary() float64 {
	return c.basicSalary
}

func (p PermanentEmployee) report() {
	fmt.Printf("Employee ID %d earns USD %f per month \n", p.id, p.calculateSalary())
}

func (c ContractEmployee) report() {
	fmt.Printf("Employee ID %d earns USD %f per month \n", c.id, c.calculateSalary())
}

func displayReport(calcuate salaryCalculator) {
	calcuate.report()
}

func printAny(value interface{}) {
	fmt.Println(value)
}

func printAnyValued(value interface{}) {
	switch value.(type) {
	case bool:
		fmt.Printf("Type: %T, Value: %v\n", value, value.(bool))
	case int:
		fmt.Printf("Type: %T, Value: %d\n", value, value.(int))
	case string:
		fmt.Printf("Type: %T, Value: %s\n", value, value)
	default:
		fmt.Println("Unsupported type")
	}
	fmt.Println("----------------------------------------")
}

func printAnyValued1(value interface{}) {
	v1, ok := value.(int)
	if ok {
		fmt.Printf("Type: %T, Value: %d\n", value, v1)
	}
	v2, ok := value.(string)
	if ok {
		fmt.Printf("Type: %T, Value: %s\n", value, v2)
	} else {
		fmt.Println("Unsupported type")
	}
	fmt.Println("----------------------------------------")
}

func main() {
	printAnyValued(true)
	printAnyValued(12)
	printAnyValued("hello")
	fmt.Println("----------------------------------------")
	printAny(true)
	printAny(12)
	printAny("hello")
	p1 := PermanentEmployee{1, 5000, 10}
	c1 := ContractEmployee{2, 3000}

	displayReport(p1)
	displayReport(c1)

	var calcutator salaryCalculator
	calcutator = PermanentEmployee{1, 5000, 10}
	calcutator.report()

	calcutator = ContractEmployee{1, 3000}
	calcutator.report()
}
