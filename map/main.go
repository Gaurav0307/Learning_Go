package main

import "fmt"

func main() {
	// Create a map to store student marks
	studentMarks := make(map[string]int)

	// Add student marks
	studentMarks["Alice"] = 85
	studentMarks["Bob"] = 92
	studentMarks["Charlie"] = 78
	// studentMarks["David"] = 88


	// Access and print student marks
	fmt.Println("Bob's marks:", studentMarks["Bob"])

	// Update Bob's marks
	studentMarks["Bob"] = 95

	fmt.Println("Updated Bob's marks:", studentMarks["Bob"])

	// Iterate over the map and print all student marks
	fmt.Println("Student Marks:")
	for student, marks := range studentMarks {
		fmt.Printf("%s: %d\n", student, marks)
	}

	// Check and delete student marks if exists
	if marks, exists := studentMarks["David"]; exists {
		delete(studentMarks, "David")
		fmt.Println("Deleted David's marks:", marks)
	}

	// Create a map to store person heights
	personHeight := map[string]float64{"Alice": 1.5, "Bob": 1.8, "Charlie": 1.7}
	
	// Iterate over the map and print all person heights
	fmt.Println("Person Heights:")
	for person, height := range personHeight {
		fmt.Printf("%s: %.2f\n", person, height)
	}
}