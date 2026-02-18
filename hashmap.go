package main

import "fmt"

func main() {

	// create hash-map
	student := make(map[string]string)
	student["name"] = "Jeff"
	student["university"] = "Bruh Uni"
	student["major"] = "fuckall"

	fmt.Println(student)

	// map literals
	scores := map[string]int{
		"Anatomy": 88,
		"Physio":  99,
		"Biochem": 100,
	}

	// Access values
	fmt.Println("Anatomy score:", scores["Anatomy"])

	// Check if key exist
	score, exists := scores["Cardio"]
	if exists {
		fmt.Printf("Cardio: %d\n", score)
	} else {
		fmt.Printf("Cardio score not found\n")
	}

	// Delete key
	delete(scores, "Biochem")

	// Iterate over map
	for subject, score := range scores {
		fmt.Printf("%s: %d\n", subject, score)
	}
}
