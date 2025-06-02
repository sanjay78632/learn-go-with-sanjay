## Day 5 – Interfaces, Packages, and Modules in Go

---

### ✅ Interfaces
- Interfaces define behavior expected from types.
- They don’t implement methods — types that satisfy them do.
```go
type Shape interface {
	Area() float64
}
