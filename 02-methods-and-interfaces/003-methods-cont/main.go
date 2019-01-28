package main

import (
	"fmt"
	"math"
)

// MyFloat is a float64 ext
type MyFloat float64

// Abs computes absolute value
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
