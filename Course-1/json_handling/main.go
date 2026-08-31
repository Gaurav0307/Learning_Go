package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"isAdult"`
}

func main() {
	person1 := Person{Name: "John", Age: 30, IsAdult: true}
	fmt.Println("Person 1:", person1)

	// Convert struct to JSON
	jsonData, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("JSON data:", string(jsonData))

	// Convert JSON to struct
	var person2 Person
	err = json.Unmarshal(jsonData, &person2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Person 2:", person2)
}
