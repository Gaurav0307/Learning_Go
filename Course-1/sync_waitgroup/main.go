package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Main function started")

	// WaitGroup is used to wait for a group of goroutines to finish
	wg := sync.WaitGroup{}

	// Add 3 to the WaitGroup which represents 3 goroutines
	wg.Add(3)

	// Start 3 goroutines
	go printNumber(1, &wg)
	go printNumber(2, &wg)
	go printNumber(3, &wg)

	fmt.Println("In between")

	// Wait for the WaitGroup to finish
	wg.Wait()

	fmt.Println("Main function ended")
}

func printNumber(number int, wg *sync.WaitGroup) {
	// Decrement the WaitGroup when the goroutine finishes
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		fmt.Println("Number", number, ":", i)
		// sleep for 1 second
		time.Sleep(time.Second)
	}
}
