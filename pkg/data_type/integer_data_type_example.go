package data_type

import (
	"fmt"
)

func IntegerDataTypeExample() {
	// Normal integer declaration
	example := 42
	fmt.Println("Example:", example)

	// Different integer types
	var example8 int8 = 127
	var example16 int16 = 32000
	var example32 int32 = 2000000000
	var example64 int64 = 9000000000000000000
	fmt.Println("Example int8:", example8)
	fmt.Println("Example int16:", example16)
	fmt.Println("Example int32:", example32)
	fmt.Println("Example int64:", example64)

	// Unsigned integers
	var uexample uint = 100
	var uexample8 uint8 = 255
	var uexample16 uint16 = 65000
	fmt.Println("Example uint:", uexample)
	fmt.Println("Example uint8:", uexample8)
	fmt.Println("Example uint16:", uexample16)

	// Arithmetic operations
	sum := example + 10
	diff := example - 5
	product := example * 2
	quotient := example / 3
	modulus := example % 7
	fmt.Println("\nSum (example + 10):", sum)
	fmt.Println("Difference (example - 5):", diff)
	fmt.Println("Product (example * 2):", product)
	fmt.Println("Quotient (example / 3):", quotient)
	fmt.Println("Modulus (example % 7):", modulus)

	// Increment & Decrement
	example++
	fmt.Println("\nAfter Increment:", example)
	example--
	fmt.Println("After Decrement:", example)

	// Comparison
	fmt.Println("\nIs example > 50?", example > 50)
	fmt.Println("Is example == 42?", example == 42)
	fmt.Println("Is example != 10?", example != 10)
}
