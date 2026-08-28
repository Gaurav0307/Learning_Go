package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Make a GET request to a website
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	fmt.Printf("Type of response : %T\n", res)
	// fmt.Printf("Response : %v\n", res)

	// Read the response body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Type of data : %T\n", data)
	// fmt.Printf("Data : %v\n", data)
	fmt.Printf("Data : %s\n", string(data))
}
