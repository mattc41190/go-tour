package main

import "fmt"

// Vertex represents two points on a coordinate plane
type Vertex struct {
	X, Y float64
}

// Scale -- Method on type version -- (Scales a Vertex)
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// ScaleFunc -- Func which accepts type pf pointer version -- (Scales a Vertex)
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)        // v == Vertex{6,8} //  <-- Pointer is optional on methods
	ScaleFunc(&v, 10) //v == Vertex{60,80} // <-- Pointer is required on functions

	p := &Vertex{4, 3}
	p.Scale(3)      // p == &Vertex{12, 9}
	ScaleFunc(p, 8) // p== &Vertex{96, 72}

	fmt.Println(v, p)

}
