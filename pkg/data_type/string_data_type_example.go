package data_type

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func StringDataTypeExample() {
	// Normal string
	example := "Andi"
	fmt.Println("Example:", example)

	// Raw string literal
	raw := `This is a raw string.
	It can span multiple lines without using \n.
	Special characters like \t or \" are not processed.`
	fmt.Println("\nRaw String:\n", raw)

	// Multi-line string with escaped newline
	multiLine := "Line 1\nLine 2\nLine 3"
	fmt.Println("\nMulti-line String:\n", multiLine)

	// String concatenation
	greeting := "Hello, " + example + "!"
	fmt.Println("\nGreeting Example:", greeting)

	// Length of string in bytes
	fmt.Println("Length in bytes:", len(example))

	// Length in characters (runes, Unicode safe)
	fmt.Println("Length in characters:", utf8.RuneCountInString(example))

	// Accessing characters (runes)
	for i, r := range example {
		fmt.Printf("Character %d: %c\n", i, r)
	}

	// Strings package functions
	fmt.Println("\nUppercase Example:", strings.ToUpper(example))
	fmt.Println("Lowercase Example:", strings.ToLower(example))
	fmt.Println("Contains 'nd'? :", strings.Contains(example, "nd"))
	fmt.Println("Replace 'Andi' with 'Dary':", strings.Replace(example, "Andi", "Dary", 1))

	// Formatting string
	formatted := fmt.Sprintf("Welcome %s to the Go world!", example)
	fmt.Println(formatted)

	// Joining and splitting
	words := []string{"Go", "is", "fun"}
	joined := strings.Join(words, " ")
	fmt.Println("Joined words Example:", joined)

	split := strings.Split(joined, " ")
	fmt.Println("Split words Example:", split)
}
