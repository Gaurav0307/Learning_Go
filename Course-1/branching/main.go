package main

import "fmt"

func main() {
	x := 10
	// Using if-else statement
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is less than or equal to 5")
	}

	y := 3
	// Using if-else-if statement
	if y > 5 {
		fmt.Println("y is greater than 5")
	} else if y == 5 {
		fmt.Println("y is equal to 5")
	} else {
		fmt.Println("y is less than 5")
	}

	day := 5
	// Using switch-case statement
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	month := "March"
	// Using switch-case statement with multiple cases and multiple values
	switch month {
	case "January", "February", "March":
		fmt.Println("Winter Season")
	case "April", "May", "June":
		fmt.Println("Spring Season")
	case "July", "August", "September":
		fmt.Println("Summer Season")
	case "October", "November", "December":
		fmt.Println("Autumn Season")
	default:
		fmt.Println("It's another season")
	}

	temperature := 25
	// Using switch-case statement with conditions
	switch {
	case temperature < 0:
		fmt.Println("Freezing weather")
	case temperature >= 0 && temperature < 10:
		fmt.Println("Very Cold weather")
	case temperature >= 10 && temperature < 20:
		fmt.Println("Cold weather")
	case temperature >= 20 && temperature < 30:
		fmt.Println("Warm weather")
	case temperature >= 30 && temperature < 40:
		fmt.Println("Its Hot")
	default:
		fmt.Println("Its Very Hot")
	}
}
