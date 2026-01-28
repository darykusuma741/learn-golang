package data_type

import "fmt"

func DataType() {
	// ==============================
	// Multiple variable declaration
	// ==============================
	var a, b, c int = 1, 2, 3
	fmt.Println("a:", a, "b:", b, "c:", c) // Print multiple variables

	// ==============================
	// Short variable declaration (type inferred)
	// ==============================
	country := "Indonesia"
	fmt.Println("Country:", country)

	// ==============================
	// Demonstrating String Data Type
	// ==============================
	fmt.Println("\n=============================")
	fmt.Println("STRING DATA TYPE EXAMPLES")
	fmt.Println("=============================")
	StringDataTypeExample()

	// ==============================
	// Demonstrating Float Data Type
	// ==============================
	fmt.Println("\n=============================")
	fmt.Println("FLOAT DATA TYPE EXAMPLES")
	fmt.Println("=============================")
	FloatDataTypeExample()

	// ==============================
	// Demonstrating Integer Data Type
	// ==============================
	fmt.Println("\n=============================")
	fmt.Println("INTEGER DATA TYPE EXAMPLES")
	fmt.Println("=============================")
	IntegerDataTypeExample()

	// ==============================
	// Demonstrating Boolean Data Type
	// ==============================
	fmt.Println("\n=============================")
	fmt.Println("BOOLEAN DATA TYPE EXAMPLES")
	fmt.Println("=============================")
	BooleanDataTypeExample()

	// ==============================
	// Demonstrating Array Data Type
	// ==============================
	fmt.Println("\n=============================")
	fmt.Println("ARRAY DATA TYPE EXAMPLES")
	fmt.Println("=============================")
	ArrayDataTypeExample()

	// ==============================
	// Demonstrating Slice Data Type
	// ==============================
	fmt.Println("\n=============================")
	fmt.Println("SLICE DATA TYPE EXAMPLES")
	fmt.Println("=============================")
	SliceDataTypeExample()

	// ==============================
	// Demonstrating Map Data Type
	// ==============================
	fmt.Println("\n=============================")
	fmt.Println("MAP DATA TYPE EXAMPLES")
	fmt.Println("=============================")
	MapDataTypeExample()

	// ==============================
	// Demonstrating Struct Data Type
	// ==============================
	fmt.Println("\n=============================")
	fmt.Println("STRUCT DATA TYPE EXAMPLES")
	fmt.Println("=============================")
	StructDataTypeExample()

	// ==============================
	// Demonstrating Pointer
	// ==============================
	fmt.Println("\n=============================")
	fmt.Println("POINTER EXAMPLES")
	fmt.Println("=============================")
	PointerExample()

	// ==============================
	// Demonstrating Interface
	// ==============================
	fmt.Println("\n=============================")
	fmt.Println("INTERFACE EXAMPLES")
	fmt.Println("=============================")
	InterfaceExample()
}
