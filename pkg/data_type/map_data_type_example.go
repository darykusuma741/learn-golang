package data_type

import (
	"fmt"
)

func MapDataTypeExample() {
	// Creating a map
	example := map[string]int{
		"apple":  5,
		"banana": 10,
		"orange": 7,
	}
	fmt.Println("Example Map:", example)

	// Accessing values by key
	fmt.Println("Number of apples:", example["apple"])

	// Updating values
	example["banana"] = 15
	fmt.Println("Updated Map:", example)

	// Adding a new key-value pair
	example["grape"] = 12
	fmt.Println("After adding grape:", example)

	// Deleting a key-value pair
	delete(example, "orange")
	fmt.Println("After deleting orange:", example)

	// Checking if a key exists
	value, exists := example["orange"]
	fmt.Println("Does orange exist?", exists, "Value:", value)

	// Looping through a map
	fmt.Println("\nLoop through map:")
	for key, val := range example {
		fmt.Printf("Key: %s, Value: %d\n", key, val)
	}

	// Map with different value types
	mixedExample := map[string]interface{}{
		"name":   "Andi",
		"age":    25,
		"isMale": true,
	}
	fmt.Println("\nMixed Map Example:", mixedExample)
}
