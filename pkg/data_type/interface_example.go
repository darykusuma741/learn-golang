package data_type

import "fmt"

// Interface
type ExampleInterface interface {
	SayHello() string
	GetAge() int
}

// Struct Person
type Person struct {
	Name string
	Age  int
}

func (p Person) SayHello() string {
	return "Hello, my name is " + p.Name
}

func (p Person) GetAge() int {
	return p.Age
}

// Struct Student
type Student struct {
	Name  string
	Grade int
}

func (s Student) SayHello() string {
	return "Hi, I am student " + s.Name
}

func (s Student) GetAge() int {
	return s.Grade // just for demo
}

// Function using interface
func InterfaceExample() {
	example := Person{Name: "Andi", Age: 25}
	var exampleInterface ExampleInterface = example

	fmt.Println(exampleInterface.SayHello())
	fmt.Println("Age via interface:", exampleInterface.GetAge())

	// Using function with interface
	printInfo := func(e ExampleInterface) {
		fmt.Println(e.SayHello())
		fmt.Println("Age:", e.GetAge())
	}

	studentExample := Student{Name: "Budi", Grade: 10}
	fmt.Println("\nStudent implementing same interface:")
	printInfo(studentExample)
}
