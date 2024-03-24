package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	// First instance using struct literal style without names
	emp1 := Employee{"John", "Doe", 1001}

	// Second instance using struct literal style with names
	emp2 := Employee{firstName: "Alice", lastName: "Smith", id: 1002}

	// Third instance using var declaration and dot notation to populate fields
	var emp3 Employee
	emp3.firstName = "Bob"
	emp3.lastName = "Jones"
	emp3.id = 1003

	fmt.Println("Employee 1:", emp1)
	fmt.Println("Employee 2:", emp2)
	fmt.Println("Employee 3:", emp3)
}
