package main

import (
	"fmt"
	"math"
)

// Vertex is a point on a coordinate plane
type Vertex struct {
	X, Y float64
}

// Abs computes math.Sqrt(v.X*v.X + v.Y*v.Y)
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale scales both points of a Vertex
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
