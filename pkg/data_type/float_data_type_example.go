package data_type

import (
	"fmt"
	"math"
)

func FloatDataTypeExample() {
	// Normal float declaration
	example := 3.1415
	fmt.Println("Example:", example)

	// Float32 and Float64
	var example32 float32 = 3.14
	var example64 float64 = 3.14159265359
	fmt.Println("Example Float32:", example32)
	fmt.Println("Example Float64:", example64)

	// Arithmetic operations
	sum := example + 2.5
	fmt.Println("Sum (example + 2.5):", sum)

	diff := example - 1.0
	fmt.Println("Difference (example - 1.0):", diff)

	product := example * 2
	fmt.Println("Product (example * 2):", product)

	quotient := example / 2
	fmt.Println("Quotient (example / 2):", quotient)

	// Using math package
	square := math.Pow(example, 2)
	fmt.Println("Square (example^2):", square)

	sqrt := math.Sqrt(example)
	fmt.Println("Square root of example:", sqrt)

	absVal := math.Abs(-example)
	fmt.Println("Absolute value:", absVal)

	ceil := math.Ceil(example)
	fmt.Println("Ceiling value:", ceil)

	floor := math.Floor(example)
	fmt.Println("Floor value:", floor)

	round := math.Round(example)
	fmt.Println("Rounded value:", round)

	// Comparisons
	isGreater := example > 2.0
	fmt.Println("Is example > 2.0?", isGreater)

	isEqual := example == 3.1415
	fmt.Println("Is example equal to 3.1415?", isEqual)
}
