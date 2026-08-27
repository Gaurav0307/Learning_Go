package main

import (
	"fmt"
	"strings"
)

func main() {
	// Split function splits a string into a slice of substrings based on a separator character.
	s := "Apple,Banana,Coconut,Date,Orange,Strawberry"
	data := strings.Split(s, ",")
	fmt.Println("Data:", data)

	// Join function joins a slice of strings into a single string using a separator character.
	joined := strings.Join(data, ", ")
	fmt.Println("Joined:", joined)

	// Replace function replaces all occurrences of a substring in a string with another substring.
	replaced := strings.Replace(joined, "Date", "Durian", -1)
	fmt.Println("Replaced:", replaced)

	// Count function counts the number of occurrences of a substring in a string.
	s = "One two two three three three four four four four five five five five five"
	count := strings.Count(s, "four")
	fmt.Println("Count (four):", count)

	// Index function returns the index of the first occurrence of a substring in a string.
	index := strings.Index(s, "four")
	fmt.Println("Index (four):", index)

	// TrimSpace function removes leading and trailing whitespace characters from a string.
	str := "  Hello,  World!  "
	trimmed := strings.TrimSpace(str)
	fmt.Println("Trimmed:", trimmed)

	// Join function joins two strings with a separator character.
	str1 := "Hello"
	str2 := "World"
	concatenated := strings.Join([]string{str1, str2}, " ")
	fmt.Println("Concatenated:", concatenated)

	// Repeat function repeats a string a specified number of times.
	repeated := strings.Repeat("Hello", 3)
	fmt.Println("Repeated:", repeated)

	// ToLower function converts a string to lowercase.
	lowercase := strings.ToLower(str)
	fmt.Println("Lowercase:", lowercase)

	// ToUpper function converts a string to uppercase.
	uppercase := strings.ToUpper(str)
	fmt.Println("Uppercase:", uppercase)

	// ToTitle function converts a string to title case.
	contains := strings.Contains(s, "four")
	fmt.Println("Contains:", contains)
}