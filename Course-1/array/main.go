package main

import "fmt"

func main() {
	// Declare an array of strings with a length of 5
	var names [5]string
	// Assign values to the array elements
	names[0] = "Alice"
	names[1] = "Bob"
	names[2] = "Charlie"
	// Print the array using the default format
	fmt.Println("Names :", names)
	// Print the array using the %v verb to display the values
	fmt.Printf("Names : %v\n", names)
	// Print the array using the %q verb to quote the strings
	fmt.Printf("Names : %q\n", names)

	// Declare and initialize an array of strings with a length of 5
	var persons = [5]string{"David", "Eve", "Frank"}
	// Print the array using the default format
	fmt.Println("Persons :", persons)
	// Print the array using the %v verb to display the values
	fmt.Printf("Persons : %v\n", persons)
	// Print the array using the %q verb to quote the strings
	fmt.Printf("Persons : %q\n", persons)

	// Declare and initialize an array of integers with a length of 8
	var numbers = [8]int{1, 2, 3, 4, 5}
	// Print the array using the default format
	fmt.Println("Numbers :", numbers)
	// Print the array using the %v verb to display the values
	fmt.Printf("Numbers : %v\n", numbers)

	// Get the length of the array
	fmt.Println("Length of numbers array:", len(numbers))

	// Accessing an element of the array
	fmt.Println("Element at 2nd index of numbers array is", numbers[2])
}
