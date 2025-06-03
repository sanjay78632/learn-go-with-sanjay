package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// Flags
	name := flag.String("name", "", "Student's name")
	age := flag.Int("age", 0, "Student's age")
	mode := flag.String("mode", "append", "Mode: overwrite/append")

	flag.Parse()

	// Basic validation
	if *name == "" || *age == 0 {
		log.Fatal("Both --name and --age are required.")
	}

	// Log setup
	log.SetPrefix("STUDENT-LOG: ")
	log.SetFlags(log.Ldate | log.Ltime)

	// File write mode
	var file *os.File
	var err error
	if *mode == "overwrite" {
		file, err = os.Create("students.txt") // overwrites
	} else {
		file, err = os.OpenFile("students.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	}

	if err != nil {
		log.Fatalf("Could not open file: %v", err)
	}
	defer file.Close()

	entry := fmt.Sprintf("Name: %s, Age: %d\n", *name, *age)

	if _, err := file.WriteString(entry); err != nil {
		log.Fatalf("Failed to write: %v", err)
	}

	log.Println("Student data saved successfully.")
}