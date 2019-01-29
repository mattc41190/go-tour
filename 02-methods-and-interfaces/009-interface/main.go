package main

import (
	"fmt"
	"math"
)

// Abser is for types that implement Abs
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser

	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // builds because type value MyFloat implements Abs
	a = &v // builds because type pointer to Vertex implements Abs

	// Derivation from Go Tour. If you want to see the original implementation uncomment the line below.
	// a = v // does not build because the type value of type Vertex does not implement Abs, only its pointer does.

	fmt.Println(a.Abs())
}

// MyFloat is a custom `float64`
type MyFloat float64

// Abs returns the absolute value of a MyFloat
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return (float64(-f))
	}
	return float64(f)
}

// Vertex represent two points on a coordinate plane
type Vertex struct {
	X, Y float64
}

// Abs returns math.Sqrt(v.X*v.X + v.Y*v.Y)
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
