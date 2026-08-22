package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if false {
		var firstName string
		fmt.Println("Please enter your first name:")
		// Scan function reads input from the standard input (keyboard) and stores it in the provided variable.
		// The & operator is used to pass the address of the variable to the Scan function.
		// This allows the Scan function to modify the value of the variable directly.
		// The Scan function can read only one value at a time, so if the user enters multiple words, only the first word will be stored in the variable.
		// The Scan function can read multiple values separated by spaces, but in this case, we are only reading one value (the first name).
		// Note: The Scan function will wait for the user to input a value and press Enter before proceeding.
		fmt.Scan(&firstName)
		fmt.Println("Hello, " + firstName + "!")
	} else {
		var fullName string
		fmt.Println("Enter your full name:")
		// Create a new reader that reads from the standard input (keyboard).
		reader := bufio.NewReader(os.Stdin)
		// The ReadString function reads input from the standard input (keyboard) until it encounters a newline character ('\n').
		// It returns the input as a string, including the newline character.
		// The second return value is an error, which we are ignoring in this case by using the blank identifier (_).
		// Note: The ReadString function will wait for the user to input a value and press Enter before proceeding.
		fullName, _ = reader.ReadString('\n')
		// The TrimSpace function removes any leading and trailing whitespace characters (including the newline character) from the input string.
		fullName = strings.TrimSpace(fullName)
		fmt.Println("Nice to meet you, " + fullName + "!")

		// Reset the reader to read from the standard input (keyboard) again.
		reader.Reset(os.Stdin)

		var age int
		fmt.Print("Enter your age: ")
		ageInput, _ := reader.ReadString('\n')
		// The Atoi function converts a string to an integer. It returns the integer value and an error if the conversion fails.
		age, _ = strconv.Atoi(strings.TrimSpace(ageInput))
		fmt.Println("You are " + strconv.Itoa(age) + " years old.")	
	}
}