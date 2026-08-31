package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Main function started")

	// Start 3 goroutines
	go printNumber(1)
	go printNumber(2)
	go printNumber(3)

	fmt.Println("In between")

	seconds := 3
	fmt.Printf("Waiting for %d seconds...\n", seconds)
	time.Sleep(time.Second * time.Duration(seconds))

	fmt.Println("Main function ended")
}

func printNumber(number int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Number", number, ":", i)
		// sleep for 1 second
		time.Sleep(time.Second)
	}
}
