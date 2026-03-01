package main

import "fmt"

func main() {

	// Fixed Sized Arrays
	var courses [4]string
	courses[0] = "Anatomy"
	courses[1] = "Physiology"
	courses[2] = "Biochemistry"
	courses[3] = "We Don't talk about this one"

	grades := [4]int{85, 90, 88, 92}

	for i := range courses {
		fmt.Printf("Course: %s\n", courses[i])
	}

	for j := range grades {
		fmt.Printf("Grade: %d\n", grades[j])
	}

	// Slices - Dynamically Sized --> Most Commonly Used
	subjects := []string{"Cardiology", "Respiratory", "Renal"}
	fmt.Println(subjects)

	// Slice Operations
	// Append to slice
	subjects = append(subjects, "Neurology")
	fmt.Println(subjects)

	fmt.Println(subjects[1:3])
	fmt.Println(subjects[:2])
	fmt.Println(subjects[2:])

	// Make slices with capacity
	numbers := make([]int, 0, 5) // length 0, capacity 5
	numbers = append(numbers, 1, 2, 3, 4)
	fmt.Println(numbers, "Length:", len(numbers), "Capacity:", cap(numbers))

	// iterate over slice
	for index, value := range numbers {
		fmt.Printf("Index %d: %d \n", index, value)
	}

}
