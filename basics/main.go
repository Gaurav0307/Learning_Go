package main

import (
	"basics/myUtils"
	"fmt"
)

func main() {
	// Print a message to the console
	fmt.Println("Hello, World!")

	// Use the PrintMessage function from the myUtils package
	myUtils.PrintMessage("Learning Go is fun!")

	// Use the AddNumbers function from the myUtils package
	result := myUtils.AddNumbers(3, 5)
	fmt.Println("3 + 5 = ", result)

	// Variable declaration and initialization
	var a int = 10 // Declare a variable 'a' of type int and initialize it with the value 10
	var b = 20     // Declare a variable 'b' of type int and initialize it with the value 20
	fmt.Println("a =", a, ", b =", b)

	var c float32 = 3.14 // Declare a variable 'c' of type float32 and initialize it with the value 3.14
	var d = 2.71          // Declare a variable 'd' of type float64 and initialize it with the value 2.71
	fmt.Println("c =", c, ", d =", d)

	var e string = "Hello" // Declare a variable 'e' of type string and initialize it with the value "Hello"
	var f = "World" // Declare a variable 'f' of type string and initialize it with the value "World"
	fmt.Println("e =", e, ", f =", f)

	var g bool = true // Declare a variable 'g' of type bool and initialize it with the value true
	var h = false // Declare a variable 'h' of type bool and initialize it with the value false
	fmt.Println("g =", g, ", h =", h)

	const pi = 3.14159 // Declare a constant 'pi' of type float64 and initialize it with the value 3.14159
	fmt.Println("pi =", pi)

	// Short variable declaration and initialization
	name := "Alice" // Declare a variable 'name' of type string and initialize it with the value "Alice".
	age := 30 // Declare a variable 'age' of type int and initialize it with the value 30
	fmt.Println("Name:", name, ", Age:", age)

	// Private and public variables
	privateVar := "This is a private variable" // This variable is private to the main package and cannot be accessed outside of this package. All private variables in Go start with a lowercase letter.
	fmt.Println("Private Variable:", privateVar)

	PublicVar := "This is a public variable" // This variable is public and can be accessed from other packages. All public variables in Go start with an uppercase letter.
	fmt.Println("Public Variable:", PublicVar)

}