package main

import "fmt"

func main(){
	
	x := 10
	fmt.Printf("Value of x: %d\n", x)
	fmt.Printf("Address of x: %d\n", &x)

	// pointer variable
	var ptr *int = &x
	fmt.Printf("Pointer value: %d\n", ptr)
	fmt.Printf("Value at pointer: %d\n", *ptr)

	// change value through pointer
	*ptr = 20
	fmt.Printf("New x value: %d\n", x)

	// practical example
	name := "Jeff"
	fmt.Printf("Name: %s", name)
	updateName(&name, "Tom")
	fmt.Printf("Name: %s", name)
}

func updateName(n *string, newName string) {
	*n = newName
}