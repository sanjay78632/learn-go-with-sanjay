package main

import "fmt"

func main() {
    //  Array (fixed size)
    var marks [3]int = [3]int{85, 90, 75}
    fmt.Println("Marks Array:", marks)

    //  Slice (dynamic array)
    students := []string{"Sanjay", "Anjali", "Rahul"}
    students = append(students, "Priya")
    fmt.Println("Student Slice:", students)

    //  Looping through slice
    for i, name := range students {
        fmt.Printf("%d: %s\n", i+1, name)
    }

    //  Map (key-value pair)
    grades := map[string]string{
        "Sanjay": "A",
        "Anjali": "B",
        "Rahul":  "C",
    }
    grades["Priya"] = "A+"

    //  Accessing and looping maps
    fmt.Println("\nStudent Grades:")
    for name, grade := range grades {
        fmt.Printf("%s got grade %s\n", name, grade)
    }
}
