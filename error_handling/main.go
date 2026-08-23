package main

import "fmt"

// divide is a function that takes two float64 parameters and returns their division result as a float64 and 
// an error if the second parameter is zero.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	var a float64 = 10
	var b float64 = 0
	var result float64

	// Calling the divide function and handle the error if it occurs.
	result, err := divide(a, b)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("The result of dividing %.1f by %.1f is %.2f.\n", a, b, result)
	}

	// Calling the divide function but not handling the error. The underscore (_) is used to ignore the error value.
	result, _ = divide(a, b)
	fmt.Printf("The result of dividing %.1f by %.1f is %.2f.\n", a, b, result)
}
