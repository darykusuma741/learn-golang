package data_type

import (
	"fmt"
)

func ArrayDataTypeExample() {
	// Array of integers
	var example [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Example Array:", example)

	// Accessing array elements
	fmt.Println("First element:", example[0])
	fmt.Println("Last element:", example[len(example)-1])

	// Updating array elements
	example[2] = 10
	fmt.Println("Updated Array:", example)

	// Looping through array
	fmt.Println("\nLoop through array:")
	for i, val := range example {
		fmt.Printf("Index %d: %d\n", i, val)
	}

	// Array of strings
	stringExample := [3]string{"Go", "is", "fun"}
	fmt.Println("\nString Array Example:", stringExample)

	// Loop through string array
	fmt.Println("\nLoop through string array:")
	for i, val := range stringExample {
		fmt.Printf("Index %d: %s\n", i, val)
	}

	// Multi-dimensional array
	matrixExample := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("\n2D Array Example:", matrixExample)

	// Access 2D array elements
	fmt.Println("Element at row 1, col 2:", matrixExample[1][2])
}
