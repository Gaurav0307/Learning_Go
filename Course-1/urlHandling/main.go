package main

import (
	"fmt"
	"net/url"
)

func main() {
	// Create a URL object from a string
	myUrl := "https://example.com/path/to/page?param1=value1&param2=value2#fragment"
	parsedUrl, err := url.Parse(myUrl)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Type of url: %T\n", parsedUrl)

	// Accessing URL components
	fmt.Println("Scheme:", parsedUrl.Scheme)
	fmt.Println("Host:", parsedUrl.Host)
	fmt.Println("Path:", parsedUrl.Path)
	fmt.Println("RawQuery:", parsedUrl.RawQuery)
	fmt.Println("Query:", parsedUrl.Query())
	fmt.Println("Fragment:", parsedUrl.Fragment)

	// Modifying URL components
	parsedUrl.Scheme = "https"
	parsedUrl.Host = "www.example.com"
	parsedUrl.Path = "/new/path"
	parsedUrl.RawQuery = "param3=value3&param4=value4"
	parsedUrl.Fragment = "new-fragment"

	fmt.Println("\nModified URL:")
	fmt.Println("Scheme:", parsedUrl.Scheme)
	fmt.Println("Host:", parsedUrl.Host)
	fmt.Println("Path:", parsedUrl.Path)
	fmt.Println("RawQuery:", parsedUrl.RawQuery)
	fmt.Println("Query:", parsedUrl.Query())
	fmt.Println("Fragment:", parsedUrl.Fragment)
}
