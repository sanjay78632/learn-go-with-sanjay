## Day 4 – Structs, Pointers, and JSON in Go

---

### ✅ Structs
- Used to define custom data types.
- Syntax: type Name struct { fields... }
```go
type Student struct {
    Name string
    Age  int
    Grade string
}

also a way to create a Variable in go 
s := Student{Name: "Sanjay", Age: 24, Grade: "A"}

Pointers
A pointer stores the memory address of a variable.
x := 5
ptr := &x  // pointer to x
fmt.Println(*ptr) // dereferencing = 5
