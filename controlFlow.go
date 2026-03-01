package main

import "fmt"

func main() {
	
	age := 23

	// If-else
	if age >= 18 {
		fmt.Printf("Adult \n")
	} else {
		fmt.Printf("THIS IS A CHILD \n")
	}

	// If with initialization
	if score := 85; score >= 80{
		fmt.Printf("Grade: A\n")
	}

	// For loop (Go only has 'for', no 'while')
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// While-style loop
	count := 0
	for count < 3 {
		fmt.Printf("Count: %d\n", count)
		count++
	}

	// Switch Statement
	day := "Sunday"
	switch day {
	case "Monday":
		fmt.Printf("Monday Blues\n")
	case "Tuesday":
		fmt.Printf("Today sucks\n")
	case "Wednesday":
		fmt.Printf("Humpday Funday\n")
	case "Thursday":
		fmt.Printf("Almost Friday\n")
	case "Friday":
		fmt.Printf("TGIF\n")
	case "Saturday":
		fmt.Printf("Yay Saturday\n")
	case "Sunday":
		fmt.Printf("Boooo! Its almost Monday\n")
	default:
		fmt.Printf("IDK WHAT DAY IT IS MAN! LEAVE ME ALONE!\n")
	}
	
	// Switch without expression (like if-else chain)
    temperature := 25
    switch {
    case temperature < 0:
        fmt.Println("Freezing\n")
    case temperature < 20:
        fmt.Println("Cold\n")
    case temperature < 30:
        fmt.Println("Comfortable\n")
    default:
        fmt.Println("Hot\n")
    }
}