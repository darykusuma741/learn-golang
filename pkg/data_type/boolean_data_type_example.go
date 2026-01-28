package data_type

import (
	"fmt"
)

func BooleanDataTypeExample() {
	// Normal boolean declaration
	example := true
	fmt.Println("Example:", example)

	// Boolean false
	var exampleFalse bool = false
	fmt.Println("Example False:", exampleFalse)

	// Boolean expressions
	a := 10
	b := 20

	greater := a > b
	less := a < b
	equal := a == b
	notEqual := a != b
	andExample := (a < b) && (b > 15)
	orExample := (a > b) || (b > 15)
	notExample := !(a < b)

	fmt.Println("\nGreater (a > b):", greater)
	fmt.Println("Less (a < b):", less)
	fmt.Println("Equal (a == b):", equal)
	fmt.Println("Not Equal (a != b):", notEqual)
	fmt.Println("AND Example ((a < b) && (b > 15)):", andExample)
	fmt.Println("OR Example ((a > b) || (b > 15)):", orExample)
	fmt.Println("NOT Example (!(a < b)):", notExample)

	// Using boolean in if statement
	if example {
		fmt.Println("\nExample is true!")
	} else {
		fmt.Println("\nExample is false!")
	}

	// Boolean in loops
	count := 0
	for example && count < 3 {
		fmt.Println("Loop iteration:", count)
		count++
		if count == 2 {
			example = false // break loop using boolean
		}
	}
	fmt.Println("Loop ended, example:", example)
}
