package data_type

import (
	"fmt"
)

func SliceDataTypeExample() {
	// Creating a slice
	example := []int{1, 2, 3}
	fmt.Println("Example Slice:", example)

	// Appending elements
	example = append(example, 4)
	example = append(example, 5, 6)
	fmt.Println("After append:", example)

	// Accessing elements
	fmt.Println("First element:", example[0])
	fmt.Println("Last element:", example[len(example)-1])

	// Updating elements
	example[1] = 20
	fmt.Println("After update:", example)

	// Slicing a slice
	subExample := example[1:4] // elements index 1 to 3
	fmt.Println("Sub-slice (index 1 to 3):", subExample)

	// Looping through slice
	fmt.Println("\nLoop through slice:")
	for i, val := range example {
		fmt.Printf("Index %d: %d\n", i, val)
	}

	// Slice of strings
	stringExample := []string{"Go", "is", "fun"}
	stringExample = append(stringExample, "and", "powerful")
	fmt.Println("\nString Slice Example:", stringExample)

	// Loop through string slice
	fmt.Println("\nLoop through string slice:")
	for i, val := range stringExample {
		fmt.Printf("Index %d: %s\n", i, val)
	}

	// Multi-dimensional slice
	matrixExample := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("\n2D Slice Example:", matrixExample)

	// Access 2D slice elements
	fmt.Println("Element at row 1, col 2:", matrixExample[1][2])
}
