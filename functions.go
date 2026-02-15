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

func multiply(a int, b int) {
	fmt.Printf("%d multiplied by %d is %d \n", a, b, a*b)
}

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}

// Named return values
func calculate(a, b int) (sum int, product int) {
    sum = a + b
    product = a * b
    return  // returns sum and product automatically
}

func main() {
	greet("Alice")
	add(10, 5)
	subtract(10, 5)
	divide(10, 5)
	multiply(10, 5)

	sum, product := calculate(10, 5)
	fmt.Printf("Sum: %d, Product: %d\n", sum, product)
}
