package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ToDo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func get() {
	// Make a GET request to an API
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("API call failed:", err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("API call failed with status code:", res.StatusCode)
		return
	}

	fmt.Println("Response Status:", res.Status)

	var todo ToDo

	if false {
		// Method 1

		// Read the response body
		data, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Failed to read response body:", err)
			return
		}

		fmt.Printf("Type of data: %T\n", data)
		fmt.Printf("JSON Data: %s\n", string(data))

		// Unmarshal the JSON data into the ToDo struct
		err = json.Unmarshal(data, &todo)
		if err != nil {
			fmt.Println("Failed to unmarshal JSON:", err)
			return
		}
	} else {
		// Method 2

		// Decode the JSON response into the ToDo struct
		err = json.NewDecoder(res.Body).Decode(&todo)
		if err != nil {
			fmt.Println("Failed to decode JSON:", err)
			return
		}
	}

	fmt.Println("ToDo:", todo)
	fmt.Println("User ID:", todo.UserID)
	fmt.Println("ID:", todo.ID)
	fmt.Println("Title:", todo.Title)
	fmt.Println("Completed:", todo.Completed)
}

func post() {
	// Create a ToDo struct
	todo := ToDo{
		UserID:    1,
		ID:        1,
		Title:     "Buy groceries",
		Completed: false,
	}

	// Marshal the ToDo struct into JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Failed to marshal JSON:", err)
		return
	}

	// Convert the JSON data to a string
	stringData := string(jsonData)

	// Create a new reader that reads from the string
	reader := strings.NewReader(stringData)

	// Make a POST request to an API
	res, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json", reader)
	if err != nil {
		fmt.Println("API call failed:", err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		fmt.Println("API call failed with status code:", res.StatusCode)
		return
	}

	fmt.Println("Response Status:", res.Status)

	fmt.Println("ToDo created successfully")

	// Read the response body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return
	}

	fmt.Printf("Type of data: %T\n", data)
	fmt.Printf("JSON Data: %s\n", string(data))
}

func put() {
	todo := ToDo{
		UserID:    1,
		ID:        1,
		Title:     "Buy groceries",
		Completed: true,
	}

	// Marshal the ToDo struct into JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Failed to marshal JSON:", err)
		return
	}

	// Convert the JSON data to a string
	stringData := string(jsonData)

	// Create a new reader that reads from the string
	reader := strings.NewReader(stringData)

	// Make a PUT request to an API
	req, err := http.NewRequest(http.MethodPut, "https://jsonplaceholder.typicode.com/todos/1", reader)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the request headers
	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP client
	client := &http.Client{}

	// Make the API call
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("API call failed:", err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("API call failed with status code:", res.StatusCode)
		return
	}

	fmt.Println("Response Status:", res.Status)

	fmt.Println("ToDo updated successfully")

	// Read the response body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return
	}

	fmt.Printf("Type of data: %T\n", data)
	fmt.Printf("JSON Data: %s\n", string(data))
}

func delete() {
	req, err := http.NewRequest(http.MethodDelete, "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Create a new HTTP client
	client := &http.Client{}

	// Make the API call
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("API call failed:", err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("API call failed with status code:", res.StatusCode)
		return
	}

	fmt.Println("Response Status:", res.Status)

	fmt.Println("ToDo deleted successfully")

	// Read the response body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return
	}

	fmt.Printf("Type of data: %T\n", data)
	fmt.Printf("JSON Data: %s\n", string(data))
}

func main() {
	// get()
	// post()
	// put()
	delete()
}
