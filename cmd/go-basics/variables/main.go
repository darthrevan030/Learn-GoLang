package main

import "fmt"

func main() {

	// explicit type declaration
	var name string = "Alice"
	var age int = 30
	var isStudent bool = true

	// short variable declaration (type inference)
	city := "San Francisco"
	salary := 75000.50

	// type inference with explicit type declaration
	var country = "USA"
	var university = "University of Go"

	// print the variables
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Student:", isStudent)
	fmt.Println("City:", city)
	fmt.Println("Salary:", salary)
	fmt.Println("Country:", country)
	fmt.Println("University:", university)

	var x, y, z int = 1, 2, 3
	fmt.Println("x:", x, "y:", y, "z:", z)

}