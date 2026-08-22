package main

import "fmt"

func main() {
	name := "John"
	age := 30
	height := 5.6789

	// PrintLn function prints the provided arguments followed by a newline character.
	// It automatically adds spaces between the arguments.
	fmt.Println("Name:", name, ", Age:", age, ", Height:", height)

	// Print function prints the provided arguments without adding a newline character.
	// It does not automatically add spaces between the arguments.
	fmt.Print("Name: ", name, ", Age: ", age, ", Height: ", height)

	// Printf function allows formatted output using format specifiers.
	// %s is used for strings, %d for integers, and %f for floating-point numbers.
	fmt.Printf("Name: %s, Age: %d, Height: %.1f\n", name, age, height)

	// You can also use Printf to format the output with different formatting options.
	// For example, you can specify the width and precision for floating-point numbers.
	fmt.Printf("Height with 2 decimal places: %.2f\n", height)

	// the %T format specifier can be used to print the type of a variable.
	fmt.Printf("The Data-Type of variable 'name' is %T\n", name);
}