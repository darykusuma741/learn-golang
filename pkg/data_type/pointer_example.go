package data_type

import (
	"fmt"
)

func PointerExample() {
	// Normal variable
	example := 10
	fmt.Println("Example value:", example)
	fmt.Println("Example address:", &example)

	// Pointer variable
	var pointerExample *int = &example
	fmt.Println("\nPointerExample points to address:", pointerExample)
	fmt.Println("Value via pointer:", *pointerExample)

	// Changing value via pointer
	*pointerExample = 20
	fmt.Println("\nAfter changing via pointer:")
	fmt.Println("Example value:", example)
	fmt.Println("Value via pointer:", *pointerExample)

	// Pointer to struct
	type ExampleStruct struct {
		Name string
		Age  int
	}

	structExample := ExampleStruct{Name: "Andi", Age: 25}
	pointerStruct := &structExample
	fmt.Println("\nStruct Example:", structExample)

	// Updating struct via pointer
	pointerStruct.Age = 26
	fmt.Println("Updated Struct Example via pointer:", structExample)

	// Function using pointer
	increment := func(val *int) {
		*val += 1
	}
	increment(&example)
	fmt.Println("\nExample after increment function:", example)
}
