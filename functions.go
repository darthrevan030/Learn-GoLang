package main

import "fmt"

func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func add(a int, b int) {
	fmt.Printf("The sum of %d and %d is %d\n", a, b, a+b)
}

func subtract(a int, b int) {
	fmt.Printf("%d minus %d is %d \n", a, b, a-b)
}

func divide(a int, b int) {
	fmt.Printf("%d divided by %d is %d \n", a, b, a/b)
}

func multiply(a int, b int) {
	fmt.Printf("%d multiplied by %d is %d \n", a, b, a*b)
}

func 

func main() {
	greet("Alice")
	add(10, 5)
	subtract(10, 5)
	divide(10, 5)
	multiply(10, 5)
}
