package main

import (
	"fmt"
	"myapp/shapes"
)


func printArea(s shapes.Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {
	r := shapes.Rectangle{Width: 10, Height: 5}
	c := shapes.Circle{Radius: 7}

	fmt.Println("🔷 Rectangle:")
	printArea(r)

	fmt.Println("\n🔵 Circle:")
	printArea(c)
}
