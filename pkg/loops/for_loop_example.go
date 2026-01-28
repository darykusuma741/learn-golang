package loops

import "fmt"

func ForLoopAdvancedExample() {
	// ==============================
	// 1. Basic for loop
	// ==============================
	fmt.Println("1. Basic for loop:")
	n := 5
	for i := range make([]struct{}, n) {
		fmt.Println("Iteration", i)
	}

	// ==============================
	// 2. For as while loop
	// ==============================
	fmt.Println("\n2. For as while loop:")
	count := 0
	for count < 5 {
		fmt.Println("Count:", count)
		count++
	}

	// ==============================
	// 3. Infinite loop with break
	// ==============================
	fmt.Println("\n3. Infinite loop with break:")
	sum := 0
	for {
		sum += 2
		fmt.Println("Sum:", sum)
		if sum >= 10 {
			break
		}
	}

	// ==============================
	// 4. For-range loop with slice/array
	// ==============================
	fmt.Println("\n4. For-range loop with slice:")
	numbers := []int{10, 20, 30, 40}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	fmt.Println("\nFor-range loop with array:")
	arr := [3]string{"Go", "Python", "JavaScript"}
	for index, val := range arr {
		fmt.Printf("Index: %d, Value: %s\n", index, val)
	}

	// ==============================
	// 5. For-range loop with map
	// ==============================
	fmt.Println("\n5. For-range loop with map:")
	countryCapital := map[string]string{
		"Indonesia": "Jakarta",
		"USA":       "Washington D.C.",
		"Japan":     "Tokyo",
	}
	for country, capital := range countryCapital {
		fmt.Printf("Country: %s, Capital: %s\n", country, capital)
	}

	// ==============================
	// 6. Nested for loops
	// ==============================
	fmt.Println("\n6. Nested for loops (multiplication table):")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d x %d = %d\n", i, j, i*j)
		}
		fmt.Println("---")
	}

	// ==============================
	// 7. Using continue
	// ==============================
	fmt.Println("\n7. Using continue (skip odd numbers):")
	for i := range make([]struct{}, n) {
		if i%2 != 0 {
			continue // skip odd numbers
		}
		fmt.Println("Even number:", i)
	}

	// ==============================
	// 8. Iterating string safely (runes)
	// ==============================
	fmt.Println("\n8. Iterating string safely (unicode runes):")
	text := "Go✨Lang"
	for index, char := range text {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}

	// ==============================
	// 9. For loop with short statement
	// ==============================
	fmt.Println("\n9. For loop with short statement:")
	for i := range make([]struct{}, n) {
		square := i * i
		fmt.Printf("Number: %d, Square: %d\n", i, square)
	}

	// ==============================
	// 10. Looping n times without creating a slice
	// ==============================
	fmt.Println("\n10. Looping n times without creating a slice:")
	for i := range make([]struct{}, n) { // modern idiomatic loop
		fmt.Println("Iteration:", i)
	}

	// ==============================
	// 11. Looping backward (decrement)
	// ==============================
	fmt.Println("\n11. Looping backward (decrement):")
	for i := 5; i > 0; i-- {
		fmt.Println("i:", i)
	}

	// ==============================
	// 12. Break outer loop with label (nested loops)
	// ==============================
	fmt.Println("\n12. Break outer loop with label:")
OuterLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i*j > 4 {
				break OuterLoop
			}
			fmt.Println("i:", i, "j:", j)
		}
	}

	// ==============================
	// 13. Continue outer loop with label (nested loops)
	// ==============================
	fmt.Println("\n13. Continue outer loop with label:")
OuterLoop2:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j == 2 {
				continue OuterLoop2
			}
			fmt.Println("i:", i, "j:", j)
		}
	}

	// ==============================
	// 14. Iterating string safely with rune (unicode safe)
	// ==============================
	fmt.Println("\n14. Iterating string safely with rune:")
	for i, r := range text {
		fmt.Printf("Byte index: %d, Rune: %c\n", i, r)
	}

	// ==============================
	// 15. Looping over slice with index only
	// ==============================
	fmt.Println("\n15. Looping over slice with index only:")
	for i := range numbers {
		fmt.Println("Index only:", i)
	}

}
