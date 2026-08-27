package myUtils

import "fmt"

func PrintMessage(message string) {
	fmt.Println(message)
}

func AddNumbers(a int, b int) int {
	return a + b
}

// privateFunction is a private function that cannot be accessed outside of this package.
// It is only used within this package.
// Note: In Go, functions that start with a lowercase letter are considered private to the package.
func privateFunction() {
	fmt.Println("This is a private function.")
}