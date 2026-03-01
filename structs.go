package main

import "fmt"

// define a struct
type Patient struct {
	Name string
	Age int
	BloodType string
	HeartRate int
}

// Method on struct
func (p Patient) IsAdult() bool {
	return p.Age >=21
}

// method with pointer receiver (can modify the struct)
func (p *Patient) UpdateHeartRate(newRate int) {
	p.HeartRate = newRate
}

func main()  {
	// create struct
	patient1 := Patient{
		Name: "John",
		Age: 22,
		BloodType: "O-",
		HeartRate: 75,

	}

	fmt.Println(patient1)
	fmt.Printf("Is Adult: %t\n", patient1.IsAdult())

	// update heart rate
	patient1.UpdateHeartRate(99)
	fmt.Printf("New Heart Rate: %d\n", patient1.HeartRate)

	//short declaration
	patient2 := Patient{"Jane", 90, "B-", 120}
	fmt.Println(patient2.Name)

}