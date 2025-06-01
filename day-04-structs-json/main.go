package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Student struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Grade    string  `json:"grade"`
	Enrolled bool    `json:"enrolled"`
}

func main() {
	var students []Student

	for i := 1; i <= 3; i++ {
		var s Student
		var enrolledInput string

		fmt.Printf("\nEnter details for student %d\n", i)
		fmt.Print("Name: ")
		fmt.Scanln(&s.Name)
		fmt.Print("Age: ")
		fmt.Scanln(&s.Age)
		fmt.Print("Grade: ")
		fmt.Scanln(&s.Grade)
		fmt.Print("Is Enrolled? (true/false): ")
		fmt.Scanln(&enrolledInput)

		s.Enrolled = (enrolledInput == "true")
		students = append(students, s)
	}

	// Encode students slice to JSON and save to file
	file, err := os.Create("students.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty JSON
	err = encoder.Encode(students)
	if err != nil {
		fmt.Println("Error writing JSON:", err)
		return
	}

	fmt.Println("\n✅ Student data saved to students.json")
}