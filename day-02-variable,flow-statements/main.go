package main

import (
	"fmt"
	"strconv"
)

func main() {
	//  Constant Variable
	const university = "SKIPS University"

	//  Using a for loop to collect info for 3 students
	for i := 1; i <= 3; i++ {
		fmt.Printf("\nStudent %d\n", i)

		// Variable declaration using 'var'
		var name, scoreInput string

		// Getting user input
		fmt.Print("Enter name: ")
		fmt.Scanln(&name)
		//purposely taking the score in string to see the type conversion
		fmt.Print("Enter score (0-100): ")
		fmt.Scanln(&scoreInput)

		// 
		score, _ := strconv.Atoi(scoreInput)

		
		var grade string
		if score >= 90 {
			grade = "A"
		} else if score >= 75 {
			grade = "B"
		} else if score >= 60 {
			grade = "C"
		} else {
			grade = "D"
		}


		fmt.Println("🎓 Summary:")
		fmt.Println("Name:", name)
		fmt.Println("Score:", score)
		fmt.Println("Grade:", grade)
		fmt.Println("University:", university)
	}
}
