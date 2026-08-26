package main

import "fmt"

// modifyValueByReference modifies the value of the variable pointed to by the pointer ptr
func modifyValueByReference(ptr *int) {
	*ptr = *ptr + 10
}

func main() {
	// a is a variable of type int
	var a int = 10
	// p1 is a pointer to an int. It points to the memory address of a
	var p1 *int = &a
	fmt.Println("a =", a, ", p1 =", p1, ", &a =", &a, ", *p1 =", *p1)

	// b is a variable of type int
	b := 20
	// p2 is a pointer to an int. It points to the memory address of b
	p2 := &b
	fmt.Println("b =", b, ", p2 =", p2, ", &b =", &b, ", *p2 =", *p2)

	// p3 is a pointer to an int. It does not point to any memory address. It is nil by default.
	var p3 *int
	if p3 == nil {
		fmt.Println("p3 is nil")
	}

	// value is a variable of type int
	var value = 10;

	fmt.Println("value =", value)

	// modifyValueByReference modifies the value of the variable pointed to by the pointer
	modifyValueByReference(&value)

	fmt.Println("After modifying value =", value)
}
