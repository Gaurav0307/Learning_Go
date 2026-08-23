package main

import "fmt"

// simpleFunction is a simple function that prints a message to the console and 
// does not take any parameters or return any values.
func simpleFunction() {
	fmt.Println("This is a simple function.")
}

// addNumbers is a function that takes two integer parameters and returns their sum as an integer.
func addNumbers(a int, b int) int {
	return a + b
}

// addNumbersWithNamedReturn is a function that takes two integer parameters and 
// returns their sum as an integer using a named return value.
func addNumbersWithNamedReturn(a int, b int) (sum int) {
	sum = a + b
	return // Named return value is used here, so we can just use 'return' without specifying the variable.
}

// addNumbersWithNamedReturnAndShortHand is a function that takes two integer parameters and 
// returns their sum as an integer using a named return value.
func addNumbersWithNamedReturnAndShortHand(a , b int) (sum int) {
	return a + b
}

func main() {
	simpleFunction()

	result := addNumbers(5, 10)
	fmt.Printf("The sum of 5 and 10 is %d.\n", result)

	result2 := addNumbersWithNamedReturn(5, 10)
	fmt.Printf("The sum of 5 and 10 (using named return) is %d.\n", result2)

	result3 := addNumbersWithNamedReturnAndShortHand(5, 10)
	fmt.Printf("The sum of 5 and 10 (using named return with shorthand) is %d.\n", result3)
}
