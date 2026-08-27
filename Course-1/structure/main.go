package main

import "fmt"

// Struct declaration
type Person struct {
	Name string
	Gender string
	Age  int
	Height float64
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	Street string
	City string
	State string
	ZipCode string
}

type Employee struct {
	Employee_Details Person
	Employee_Contact Contact
	Employee_Address Address
}

func main() {
	// Struct type variable
	var person1 Person

	fmt.Println("Person 1:",person1)

	// Struct initialization
	person1.Name = "John"
	person1.Gender = "Male"
	person1.Age = 30
	person1.Height = 1.78

	fmt.Println("Person 1:",person1)

	// Struct initialization using struct literal
	person2 := Person {"Jane", "Female", 25, 1.65}

	fmt.Println("Person 2:", person2)

	// Struct initialization using struct literal
	person3 := Person {
		Name: "Bob", 
		Gender: "Male", 
		Age: 40, 
		Height: 1.80,
	}

	fmt.Println("Person 3:", person3)

	// Using the new function to create a struct variable
	var person4 = new(Person)

	person4.Name = "Alice"
	person4.Gender = "Female"
	person4.Age = 35
	person4.Height = 1.70

	fmt.Println("Person 4:", person4)

	// Struct type variable
	var employee1 Employee;

	// Struct initialization
	employee1.Employee_Details = Person {
		Name: "John",
		Gender: "Male",
		Age: 30,
		Height: 1.78,
	}

	employee1.Employee_Contact.Email = "john@emailcom"
	employee1.Employee_Contact.Phone = "123-456-7890"

	employee1.Employee_Address = Address {
		Street: "123 Main St",
		City: "New York",
		State: "NY",
		ZipCode: "10001",
	}

	fmt.Println("Employee 1 Details:", employee1.Employee_Details)
	fmt.Println("Employee 1 Contact:", employee1.Employee_Contact)
	fmt.Println("Employee 1 Address:", employee1.Employee_Address)

	fmt.Println("Employee 1 City:", employee1.Employee_Address.City)

	fmt.Println("Employee 1:", employee1)

	// Struct initialization using struct literal
	var employee2 = Employee {	
		Employee_Details: Person {	
			Name: "Bob",
			Gender: "Male",
			Age: 40,
			Height: 1.80,
		},
		Employee_Contact: Contact {
			Email: "bob@email.com",
			Phone: "123-456-7890",
		},
		Employee_Address: Address {
			Street: "01 Park St",
			City: "New York",
			State: "NY",
			ZipCode: "10001",
		},
	}

	fmt.Println("Employee 2:", employee2)
}
