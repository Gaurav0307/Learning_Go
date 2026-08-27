package main

import "fmt"

func main() {
	fmt.Println("Statement-1")
	// defer defers the execution of a function until the surrounding function returns
	// defer statements are executed in LIFO order
	defer fmt.Println("Statement-2")
	defer fmt.Println("statement-3")
	defer fmt.Println("Statement-4")
	fmt.Println("Statement-5")
}