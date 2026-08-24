package main

import "fmt"

func main() {
	// Declare a slice of integers
	var slice []int
	// Alternatively, you can initialize the slice using a composite literal
	// slice := []int{}

	// Print the initial slice, its data type, length, and capacity
	fmt.Println("Initial slice:", slice)
	fmt.Printf("Initial slice (using %%v): %v\n", slice)
	fmt.Printf("Data type of the slice: %T\n", slice)
	fmt.Printf("Length of the slice: %d\n", len(slice))
	fmt.Printf("Capacity of the slice: %d\n", cap(slice))

	// Append elements to the slice
	slice = append(slice, 1, 2, 3, 4, 5)

	// Print the updated slice, its length, and capacity
	fmt.Println("Slice after appending elements:", slice)
	fmt.Printf("Length of the slice: %d\n", len(slice))
	fmt.Printf("Capacity of the slice: %d\n", cap(slice))

	fmt.Printf("-------------------------------------------------------------------------\n")

	// Create a slice of integers with initial values
	numbers := []int{1, 2, 3, 4, 5}

	// Print the original slice, its data type, length, and capacity
	fmt.Println("Original slice:", numbers)
	fmt.Printf("Original slice (using %%v): %v\n", numbers)
	fmt.Printf("Data type of the slice: %T\n", numbers)
	fmt.Printf("Length of the slice: %d\n", len(numbers))
	fmt.Printf("Capacity of the slice: %d\n", cap(numbers))

	// Append a new element to the slice
	numbers = append(numbers, 6, 7, 8)

	// Print the updated slice, its length, and capacity
	fmt.Println("Updated slice after appending elements:", numbers)
	fmt.Printf("Length of the updated slice: %d\n", len(numbers))
	fmt.Printf("Capacity of the updated slice: %d\n", cap(numbers))

	fmt.Printf("-------------------------------------------------------------------------\n")

	// Create a slice using the make function
	var values = make([]int, 5, 10) // Create a slice of integers with length 5 and capacity 10
	
	// Print the slice, its length, and capacity
	fmt.Println("Slice created using make:", values)
	fmt.Printf("Length of the slice: %d\n", len(values))
	fmt.Printf("Capacity of the slice: %d\n", cap(values))

	// Append elements to the slice
	values = append(values, 1, 2, 3, 4, 5, 6)

	// Print the updated slice, its length, and capacity
	fmt.Println("Slice after appending elements:", values)
	fmt.Printf("Length of the slice: %d\n", len(values))
	fmt.Printf("Capacity of the slice: %d\n", cap(values))

	// Modify elements of the slice
	values[0] = 1
	values[1] = 2
	values[2] = 3

	// Print the updated slice after modifying elements
	fmt.Println("Slice after modifying elements:", values)
	fmt.Printf("Length of the slice: %d\n", len(values))
	fmt.Printf("Capacity of the slice: %d\n", cap(values))
}
