package data_type

import (
	"fmt"
)

// Defining a struct
type ExampleStruct struct {
	Name   string
	Age    int
	IsMale bool
	Skills []string
	Scores map[string]int
}

func StructDataTypeExample() {
	// Creating an instance of the struct
	example := ExampleStruct{
		Name:   "Andi",
		Age:    25,
		IsMale: true,
		Skills: []string{"Go", "JavaScript", "Python"},
		Scores: map[string]int{
			"math":    90,
			"english": 85,
		},
	}

	// Accessing struct fields
	fmt.Println("Name:", example.Name)
	fmt.Println("Age:", example.Age)
	fmt.Println("IsMale:", example.IsMale)
	fmt.Println("Skills:", example.Skills)
	fmt.Println("Scores:", example.Scores)

	// Updating struct fields
	example.Age = 26
	example.Skills = append(example.Skills, "Java")
	example.Scores["science"] = 95

	fmt.Println("\nUpdated Struct Example:", example)

	// Loop through slice inside struct
	fmt.Println("\nSkills List:")
	for i, skill := range example.Skills {
		fmt.Printf("%d: %s\n", i, skill)
	}

	// Loop through map inside struct
	fmt.Println("\nScores List:")
	for subject, score := range example.Scores {
		fmt.Printf("%s: %d\n", subject, score)
	}
}
