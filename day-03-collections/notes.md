Day 3 – Arrays, Slices, and Maps in Go
 Arrays
Arrays are fixed-size collections of the same data type.

Once declared, size cannot be changed.

Syntax: var arr [size]Type

Example:
var marks [3]int = [3]int{85, 90, 75}
fmt.Println(marks[1]) // Output: 90

Use case: When the number of elements is known and fixed.

 Slices
Slices are dynamic arrays.

You can grow them using the append() function.

Syntax: []Type

Example:
students := []string{"Sanjay", "Anjali"}
students = append(students, "Rahul")
fmt.Println(students) // ["Sanjay", "Anjali", "Rahul"]

Looping through slice:
for i, name := range students {
 fmt.Printf("%d: %s\n", i+1, name)
}

Maps
Maps are key-value pairs.

Syntax: map[KeyType]ValueType

Example:
grades := map[string]string{
 "Sanjay": "A",
 "Anjali": "B",
}
grades["Rahul"] = "C"

Looping through map:
for name, grade := range grades {
 fmt.Printf("%s got grade %s\n", name, grade)
}

Access: fmt.Println(grades["Sanjay"])

Delete: delete(grades, "Anjali")

 Range (Looping)
Used with arrays, slices, and maps.

For slices/arrays: gives index, value

For maps: gives key, value

Examples:
for index, value := range someSlice {
 fmt.Println(index, value)
}

for key, value := range someMap {
 fmt.Println(key, value)
}

 Summary Table
Array – Fixed-size list – Use when data is constant
Slice – Dynamic list – Most commonly used
Map – Key-value pair – For lookups and relational data
Range – Loop helper – Used for iteration