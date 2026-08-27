package main

import (
	"fmt"
	"strconv"
)

func main() {
	// float to int conversion
	fmt.Printf("Value = %d, Type = %T\n", int(10.0), int(10.0))

	// int to float conversion
	fmt.Printf("Value = %f, Type = %T\n", float64(10), float64(10))

	// int to string conversion (to ASCII characters)
	fmt.Printf("Value = %s, Type = %T\n", string(65), string(65))

	// int to string conversion
	str := strconv.Itoa(65)
	fmt.Printf("Value = %s, Type = %T\n", str, str)

	// string to int conversion
	i, _ := strconv.Atoi("65")
	fmt.Printf("Value = %d, Type = %T\n", i, i)

	// float to string conversion
	s := strconv.FormatFloat(3.14, 'f', 2, 64)
	fmt.Printf("Value = %s, Type = %T\n", s, s)

	// string to float conversion
	f, _ := strconv.ParseFloat("3.14", 64)
	fmt.Printf("Value = %f, Type = %T\n", f, f)

	// string to bool conversion
	b, _ := strconv.ParseBool("true")
	fmt.Printf("Value = %t, Type = %T\n", b, b)

	// bool to string conversion
	s1 := strconv.FormatBool(true)
	fmt.Printf("Value = %s, Type = %T\n", s1, s1)
}