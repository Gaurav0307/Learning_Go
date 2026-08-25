package main

import "fmt"

func main() {
	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println("i:", i)
	}

	j := 0
	// while loop
	for j < 5 {
		fmt.Println("j:", j)
		j++
	}

	numbers := []int{1, 2, 3, 4, 5}
	// for-each loop
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	txt := "Hello, World!"
	// for loop with string
	for index, char := range txt {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}
}
