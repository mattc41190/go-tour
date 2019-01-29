package main

import (
	"fmt"
	"math"
)

// Vertex is two points on a coordinate plane
type Vertex struct {
	X, Y float64
}

// Abs calculates abs of a vertex
func (v Vertex) Abs() float64 { // <-- Takes a value receiver not a pointer receiver
	return math.Sqrt(v.X*v.X + v.Y*v.Y)

}

// AbsFunc calculates abs of a vertex
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4} // <-- Value
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}       // Pointer
	fmt.Println(p.Abs())     // <-- But, wait I thought Abs had a value type receiver -- WHHHHAAAAAATT ;)
	fmt.Println(AbsFunc(*p)) // <-- Still need to read through to the actual value when passing it as a normal argument
}
